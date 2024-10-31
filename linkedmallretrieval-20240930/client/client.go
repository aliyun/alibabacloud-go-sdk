// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AISearchQuery struct {
	Card         *string `json:"card,omitempty" xml:"card,omitempty"`
	CardQuery    *string `json:"cardQuery,omitempty" xml:"cardQuery,omitempty"`
	Page         *int32  `json:"page,omitempty" xml:"page,omitempty"`
	Query        *string `json:"query,omitempty" xml:"query,omitempty"`
	SearchEngine *string `json:"searchEngine,omitempty" xml:"searchEngine,omitempty"`
	SearchPlan   *string `json:"searchPlan,omitempty" xml:"searchPlan,omitempty"`
	SessionId    *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	TimeRange    *string `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
}

func (s AISearchQuery) String() string {
	return tea.Prettify(s)
}

func (s AISearchQuery) GoString() string {
	return s.String()
}

func (s *AISearchQuery) SetCard(v string) *AISearchQuery {
	s.Card = &v
	return s
}

func (s *AISearchQuery) SetCardQuery(v string) *AISearchQuery {
	s.CardQuery = &v
	return s
}

func (s *AISearchQuery) SetPage(v int32) *AISearchQuery {
	s.Page = &v
	return s
}

func (s *AISearchQuery) SetQuery(v string) *AISearchQuery {
	s.Query = &v
	return s
}

func (s *AISearchQuery) SetSearchEngine(v string) *AISearchQuery {
	s.SearchEngine = &v
	return s
}

func (s *AISearchQuery) SetSearchPlan(v string) *AISearchQuery {
	s.SearchPlan = &v
	return s
}

func (s *AISearchQuery) SetSessionId(v string) *AISearchQuery {
	s.SessionId = &v
	return s
}

func (s *AISearchQuery) SetTimeRange(v string) *AISearchQuery {
	s.TimeRange = &v
	return s
}

type GenericSearchResult struct {
	PageItems []*ScorePageItem `json:"pageItems,omitempty" xml:"pageItems,omitempty" type:"Repeated"`
	// example:
	//
	// 123456
	RequestId         *string            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SearchInformation *SearchInformation `json:"searchInformation,omitempty" xml:"searchInformation,omitempty"`
	WeiboItems        []*WeiboItem       `json:"weiboItems,omitempty" xml:"weiboItems,omitempty" type:"Repeated"`
}

func (s GenericSearchResult) String() string {
	return tea.Prettify(s)
}

func (s GenericSearchResult) GoString() string {
	return s.String()
}

func (s *GenericSearchResult) SetPageItems(v []*ScorePageItem) *GenericSearchResult {
	s.PageItems = v
	return s
}

func (s *GenericSearchResult) SetRequestId(v string) *GenericSearchResult {
	s.RequestId = &v
	return s
}

func (s *GenericSearchResult) SetSearchInformation(v *SearchInformation) *GenericSearchResult {
	s.SearchInformation = v
	return s
}

func (s *GenericSearchResult) SetWeiboItems(v []*WeiboItem) *GenericSearchResult {
	s.WeiboItems = v
	return s
}

type IncludeImage struct {
	// example:
	//
	// 324
	Height *int32 `json:"height,omitempty" xml:"height,omitempty"`
	// example:
	//
	// http://k.sinaimg.cn/n/sinakd20121/594/w2048h946/20240328/74cf-32c0d62e843df76567d760b4459d57c1.jpg/w700d1q75cms.jpg
	ImageLink *string `json:"imageLink,omitempty" xml:"imageLink,omitempty"`
	// example:
	//
	// 700
	Width *int32 `json:"width,omitempty" xml:"width,omitempty"`
}

func (s IncludeImage) String() string {
	return tea.Prettify(s)
}

func (s IncludeImage) GoString() string {
	return s.String()
}

func (s *IncludeImage) SetHeight(v int32) *IncludeImage {
	s.Height = &v
	return s
}

func (s *IncludeImage) SetImageLink(v string) *IncludeImage {
	s.ImageLink = &v
	return s
}

func (s *IncludeImage) SetWidth(v int32) *IncludeImage {
	s.Width = &v
	return s
}

type ScorePageItem struct {
	// This parameter is required.
	//
	// example:
	//
	// structure_web_info
	CardType *string `json:"cardType,omitempty" xml:"cardType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// baijiahao.baidu.com
	DisplayLink *string `json:"displayLink,omitempty" xml:"displayLink,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100km/h-0制动能力上，仅有33.3m，不黑不吹，单看这个，<em>小米SU7</em>确实表现不错。而续航方面，101kWh电池容量，实现CLTC续航800km，还有现5分钟补能220km，15分钟补能510km的800V高压平台。而在...
	HtmlSnippet *string `json:"htmlSnippet,omitempty" xml:"htmlSnippet,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// <em>小米</em>SU7售价22.99万元起 高管亲自辟谣：发布前不会有<em>价格</em>-百家号
	HtmlTitle *string         `json:"htmlTitle,omitempty" xml:"htmlTitle,omitempty"`
	Images    []*IncludeImage `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// https://baijiahao.baidu.com/s?id=1787881554557805096
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// example:
	//
	// 昨天	，	小米	汽车	没有	发布	，	但	相关	的	信息	透露	的	差	不	多	了	。
	//
	// 		在	发布	会	现场	，	雷军	直接	称	小米	S	U	7	对	标	特斯拉	保时捷	，	有	100	项	行业	领先	，	可见	“	遥遥	领先	”	已经	不再	是	华为	专利	了	？
	//
	//
	//
	// 		而	在	介绍	技术	时	，	雷军	也	从	电机	、	电池	、	大	压铸	、	自动	驾驶	、	智能	座舱	等	五	大	方面	展开	，	充分	展示	了	小米	汽车	的	技术	以及	技术	储备	，		 		能	堆	的	料	，	全都	堆	上去	了	…	…
	//
	// 		大家	比较	感	兴趣	的	性能	方面	，	2	.	78	s	的	0	-	100	km	/	h	加速	，	265	km	/	h	的	最高	时速
	MainText *string `json:"mainText,omitempty" xml:"mainText,omitempty"`
	// example:
	//
	// text/html
	Mime    *string            `json:"mime,omitempty" xml:"mime,omitempty"`
	PageMap map[string]*string `json:"pageMap,omitempty" xml:"pageMap,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1704426524000
	PublishTime *int64 `json:"publishTime,omitempty" xml:"publishTime,omitempty"`
	// example:
	//
	// 0.234325235
	Score *float64 `json:"score,omitempty" xml:"score,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 小米SU7售价22.99万元起 高管亲自辟谣：发布前不会有价格
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ScorePageItem) String() string {
	return tea.Prettify(s)
}

func (s ScorePageItem) GoString() string {
	return s.String()
}

func (s *ScorePageItem) SetCardType(v string) *ScorePageItem {
	s.CardType = &v
	return s
}

func (s *ScorePageItem) SetDisplayLink(v string) *ScorePageItem {
	s.DisplayLink = &v
	return s
}

func (s *ScorePageItem) SetHtmlSnippet(v string) *ScorePageItem {
	s.HtmlSnippet = &v
	return s
}

func (s *ScorePageItem) SetHtmlTitle(v string) *ScorePageItem {
	s.HtmlTitle = &v
	return s
}

func (s *ScorePageItem) SetImages(v []*IncludeImage) *ScorePageItem {
	s.Images = v
	return s
}

func (s *ScorePageItem) SetLink(v string) *ScorePageItem {
	s.Link = &v
	return s
}

func (s *ScorePageItem) SetMainText(v string) *ScorePageItem {
	s.MainText = &v
	return s
}

func (s *ScorePageItem) SetMime(v string) *ScorePageItem {
	s.Mime = &v
	return s
}

func (s *ScorePageItem) SetPageMap(v map[string]*string) *ScorePageItem {
	s.PageMap = v
	return s
}

func (s *ScorePageItem) SetPublishTime(v int64) *ScorePageItem {
	s.PublishTime = &v
	return s
}

func (s *ScorePageItem) SetScore(v float64) *ScorePageItem {
	s.Score = &v
	return s
}

func (s *ScorePageItem) SetTitle(v string) *ScorePageItem {
	s.Title = &v
	return s
}

type SearchInformation struct {
	SearchTime *int64 `json:"searchTime,omitempty" xml:"searchTime,omitempty"`
	Total      *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s SearchInformation) String() string {
	return tea.Prettify(s)
}

func (s SearchInformation) GoString() string {
	return s.String()
}

func (s *SearchInformation) SetSearchTime(v int64) *SearchInformation {
	s.SearchTime = &v
	return s
}

func (s *SearchInformation) SetTotal(v int64) *SearchInformation {
	s.Total = &v
	return s
}

type WeiboItem struct {
	// This parameter is required.
	//
	// example:
	//
	// weibo_strong
	CardType *string `json:"cardType,omitempty" xml:"cardType,omitempty"`
	// example:
	//
	// https://m.weibo.cn/u/7761793874?luicode=20000061&lfid=5024099350350075
	HomepageLink *string `json:"homepageLink,omitempty" xml:"homepageLink,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 【小调查：你会买<em>小米SU7</em>吗？】#小米SU7路测覆盖300多城市#4月17日，@小米汽车 发文称SU7道路测试覆盖全国300多个城市，涵盖极寒，极热天气，总里程数高达540万公里，目前仍在进行中。  网页链接
	HtmlSnippet *string   `json:"htmlSnippet,omitempty" xml:"htmlSnippet,omitempty"`
	Images      []*string `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// https://m.weibo.cn/detail/5024099350350075?wm=90194_90009
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1小时前
	PublishDisplayTime *string `json:"publishDisplayTime,omitempty" xml:"publishDisplayTime,omitempty"`
	// example:
	//
	// 白鹿科技
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s WeiboItem) String() string {
	return tea.Prettify(s)
}

func (s WeiboItem) GoString() string {
	return s.String()
}

func (s *WeiboItem) SetCardType(v string) *WeiboItem {
	s.CardType = &v
	return s
}

func (s *WeiboItem) SetHomepageLink(v string) *WeiboItem {
	s.HomepageLink = &v
	return s
}

func (s *WeiboItem) SetHtmlSnippet(v string) *WeiboItem {
	s.HtmlSnippet = &v
	return s
}

func (s *WeiboItem) SetImages(v []*string) *WeiboItem {
	s.Images = v
	return s
}

func (s *WeiboItem) SetLink(v string) *WeiboItem {
	s.Link = &v
	return s
}

func (s *WeiboItem) SetPublishDisplayTime(v string) *WeiboItem {
	s.PublishDisplayTime = &v
	return s
}

func (s *WeiboItem) SetUsername(v string) *WeiboItem {
	s.Username = &v
	return s
}

type GenericSearchRequest struct {
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// example:
	//
	// OneMonth
	TimeRange *string `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
}

func (s GenericSearchRequest) String() string {
	return tea.Prettify(s)
}

func (s GenericSearchRequest) GoString() string {
	return s.String()
}

func (s *GenericSearchRequest) SetQuery(v string) *GenericSearchRequest {
	s.Query = &v
	return s
}

func (s *GenericSearchRequest) SetTimeRange(v string) *GenericSearchRequest {
	s.TimeRange = &v
	return s
}

type GenericSearchResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenericSearchResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenericSearchResponse) String() string {
	return tea.Prettify(s)
}

func (s GenericSearchResponse) GoString() string {
	return s.String()
}

func (s *GenericSearchResponse) SetHeaders(v map[string]*string) *GenericSearchResponse {
	s.Headers = v
	return s
}

func (s *GenericSearchResponse) SetStatusCode(v int32) *GenericSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *GenericSearchResponse) SetBody(v *GenericSearchResult) *GenericSearchResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("linkedmallretrieval"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通用搜索
//
// @param request - GenericSearchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenericSearchResponse
func (client *Client) GenericSearchWithOptions(request *GenericSearchRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GenericSearchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.TimeRange)) {
		query["timeRange"] = request.TimeRange
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenericSearch"),
		Version:     tea.String("2024-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/linked-retrieval/linked-retrieval-entry/v2/linkedRetrieval/commands/genericSearch"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GenericSearchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通用搜索
//
// @param request - GenericSearchRequest
//
// @return GenericSearchResponse
func (client *Client) GenericSearch(request *GenericSearchRequest) (_result *GenericSearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenericSearchResponse{}
	_body, _err := client.GenericSearchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
