# brain-decoder
一个基于 GPT 的文本序列化工具，可以将输入的文本按照定义的结构体序列化

# Install
```shell
go get github.com/burybell/brain_decoder
```
## example
```go
type Article struct {
	Title                      string   `json:"title"`
	Subtitle                   string   `json:"subtitle"`
	Tags                       []string `json:"tags"`
	Summary                    string   `json:"summary"`
	Topics                     []string `json:"topics"`
	Kind                       string   `json:"kind" jsonschema:"enum=新闻报道,enum=科学论文,enum=文学作品,enum=学术论文,enum=技术指南,enum=历史研究"`
	WordCount                  int      `json:"wordCount" jsonschema:"description=文章的总汉字字数"`
	Highlight                  []string `json:"highlight" jsonschema:"description=文章的中可能会吸引人眼球的词语"`
	Rank                       int      `json:"rank" jsonschema:"enum=1,enum=2,enum=3,enum=4,enum=5"`
	RankReason                 string   `json:"rankReason" jsonschema:"description=评为当前分数的理由"`
	PredictingLikesReceived    int      `json:"predictingLikesReceived" jsonschema:"description=预测当前文章获取多少点赞"`
	LikeAudience               []string `json:"likeAudience"  jsonschema:"description=点赞的人群"`
	PredictingCommentsReceived int      `json:"PredictingCommentsReceived" jsonschema:"description=预测当前文章会有多少人评论"`
	CommentAudience            []string `json:"CommentAudience"  jsonschema:"description=评论的人群"`
	PossibleComments           []string `json:"possibleComments" jsonschema:"description=可能出现的评论，可以带emoji"`
	Source                     string   `json:"source" jsonschema:"description=猜测文章来源"`
}

func init() {
    brain_decoder.DefaultPrompt = brain_decoder.ChinesePrompt
    brain_decoder.OpenAIClient = openai.NewClient(os.Getenv("OPENAI_KEY"))
}

func main() {
    var source = `
全国政协委员、对外经济贸易大学教授孙洁持续关注积极应对人口老龄化相关政策的不断完善，在今年的全国两会上，她将就完善老年助餐体系、建设老年友好型社会等方面提出建议。
春节期间，全国政协委员孙洁来到北京的一些街道、社区养老机构开展实地调研，进一步完善今年两会的提案内容。去年一年，她通过发放2000多份调查问卷、召开多场座谈会围绕完善居家养老服务体系展开调研。
去年，国家关于基本养老服务体系建设、老年助餐服务、养老人才发展等一系列政策密集出台。孙洁在调研中发现，老年助餐服务因价格优惠、品种丰富等特点受到老年人的欢迎，但是还面临物业费用高、上门送餐人员短缺等方面的压力。
北京康养老年福养老服务中心党支部书记 张文峰：社区餐厅从价格上要亲民，同时现在针对社区的空巢老人、高龄独居老人我们要开展送餐服务，送餐成本还很高，人力还很不足。
今年两会，孙洁将就如何让老年助餐服务具有造血功能、实现可持续运营等提出具体建议。
全国政协委员 孙洁：老年助餐服务面临的最大的问题应该是它的可持续经营，所以建议政府相关部门能够出台政策鼓励老年餐厅多元化经营，同时吸引或者引导社会资本来支持老年就餐服务。
近年来，孙洁从积极应对人口老龄化的经济保障、服务保障、制度建设等多个角度持续向全国政协提交提案，其中多项建议被有关部门采纳。过去一年，孙洁聚焦银发经济，从老年人刚需的适老化产品着手，先后到多家养老机构、康复辅具企业展开调研。今年两会期间，她还将提交一项提案，建议从适老化改造的供给端和需求端同时发力，推动公共场所、社区和家庭适老化改造。
全国政协委员 孙洁：一方面是有居家的适老化改造，还有社区的适老化改造，交通出行的数字鸿沟这一块的适老化改造，要实现一个老年友好型的社会的建设。
`
    var article Article
    err := brain_decoder.Unmarshal([]byte(source), &article)
    if err != nil {
        panic(err)
    }
    bs, err := json.MarshalIndent(article, "", "\t")
    if err != nil {
        panic(err)
    }
    fmt.Println(string(bs))
    // output
    //{
    //	"title": "全国政协委员孙洁关注完善老年助餐体系 提出建议",
    //	"subtitle": "今年两会将提交关于老年助餐服务的具体建议",
    //	"tags": [
    //		"全国政协委员",
    //		"人口老龄化",
    //		"老年助餐服务",
    //		"建议"
    //	],
    //	"summary": "全国政协委员、对外经济贸易大学教授孙洁持续关注和积极应对人口老龄化相关政策的不断完善，今年将就完善老年助餐体系、建设老年友好型社会等方面提出建议。",
    //	"topics": [
    //		"人口老龄化",
    //		"老年助餐服务",
    //		"社会养老保障"
    //	],
    //	"kind": "新闻报道",
    //	"wordCount": 340,
    //	"highlight": [
    //		"老年助餐服务",
    //		"积极应对人口老龄化",
    //		"完善政策",
    //		"建设老年友好型社会"
    //	],
    //	"rank": 4,
    //	"rankReason": "全国政协委员孙洁关注老年助餐服务，建议具体、有益社会",
    //	"predictingLikesReceived": 120,
    //	"likeAudience": [
    //		"政府部门",
    //		"社会资本",
    //		"老年人"
    //	],
    //	"PredictingCommentsReceived": 25,
    //	"CommentAudience": [
    //		"政策制定者",
    //		"社会公众",
    //		"养老服务机构"
    //	],
    //	"possibleComments": [
    //		"很有建设性的提案",
    //		"老年助餐服务很重要",
    //		"需要加强实施"
    //	],
    //	"source": "新华社"
    //}
}
```