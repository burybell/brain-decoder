package brain_decoder

import (
	"encoding/json"
	"fmt"
	"github.com/invopop/jsonschema"
	"github.com/xeipuuv/gojsonschema"
	"io"
)

type Decoder struct {
	r     io.Reader
	g     func(prompt string) ([]byte, error)
	p     func(source string, schema string) string
	tries int
}

func NewDecoder(
	r io.Reader,
	g func(prompt string) ([]byte, error),
	p func(source string, schema string) string,
	tries int,
) *Decoder {
	return &Decoder{r: r, g: g, p: p, tries: tries}
}

func (enc *Decoder) Encode(v any) error {

	// 生成schema
	reflect := jsonschema.Reflect(v)
	schema, err := json.Marshal(reflect)
	if err != nil {
		return err
	}

	// 读取待处理字符串
	source, err := io.ReadAll(enc.r)
	if err != nil {
		return err
	}

	var errs = make([]error, 0)
	for i := 0; i < enc.tries; i++ {
		// 处理
		prompt := enc.p(string(source), string(schema))
		output, err := enc.g(prompt)
		if err != nil {
			return err
		}

		schemaLoader := gojsonschema.NewBytesLoader(schema)
		documentLoader := gojsonschema.NewBytesLoader(output)
		validate, err := gojsonschema.Validate(schemaLoader, documentLoader)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		if validate.Valid() {
			return json.Unmarshal(output, &v)
		} else {
			errs = append(errs, fmt.Errorf("%+v", validate.Errors()))
		}
	}
	return fmt.Errorf("%+v", errs)
}
