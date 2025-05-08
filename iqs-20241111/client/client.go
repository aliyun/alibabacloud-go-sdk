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
	PageItems    []*ScorePageItem `json:"pageItems,omitempty" xml:"pageItems,omitempty" type:"Repeated"`
	QueryContext *QueryContext    `json:"queryContext,omitempty" xml:"queryContext,omitempty"`
	// example:
	//
	// 123456
	RequestId         *string            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SceneItems        []*SceneItem       `json:"sceneItems,omitempty" xml:"sceneItems,omitempty" type:"Repeated"`
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

func (s *GenericSearchResult) SetQueryContext(v *QueryContext) *GenericSearchResult {
	s.QueryContext = v
	return s
}

func (s *GenericSearchResult) SetRequestId(v string) *GenericSearchResult {
	s.RequestId = &v
	return s
}

func (s *GenericSearchResult) SetSceneItems(v []*SceneItem) *GenericSearchResult {
	s.SceneItems = v
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

type GlobalPageItem struct {
	// This parameter is required.
	//
	// example:
	//
	// https://baijiahao.baidu.com/s?id=1787881554557805096
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	// example:
	//
	// 100km/h-0制动能力上，仅有33.3m，不黑不吹，单看这个，小米SU7确实表现不错。而续航方面，101kWh电池容量，实现CLTC续航800km，还有现5分钟补能220km，15分钟补能510km的800V高压平台。而在...
	Snippet *string `json:"snippet,omitempty" xml:"snippet,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 小米SU7售价22.99万元起 高管亲自辟谣：发布前不会有价格
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GlobalPageItem) String() string {
	return tea.Prettify(s)
}

func (s GlobalPageItem) GoString() string {
	return s.String()
}

func (s *GlobalPageItem) SetLink(v string) *GlobalPageItem {
	s.Link = &v
	return s
}

func (s *GlobalPageItem) SetSnippet(v string) *GlobalPageItem {
	s.Snippet = &v
	return s
}

func (s *GlobalPageItem) SetTitle(v string) *GlobalPageItem {
	s.Title = &v
	return s
}

type GlobalQueryContext struct {
	OriginalQuery *GlobalQueryContextOriginalQuery `json:"originalQuery,omitempty" xml:"originalQuery,omitempty" type:"Struct"`
}

func (s GlobalQueryContext) String() string {
	return tea.Prettify(s)
}

func (s GlobalQueryContext) GoString() string {
	return s.String()
}

func (s *GlobalQueryContext) SetOriginalQuery(v *GlobalQueryContextOriginalQuery) *GlobalQueryContext {
	s.OriginalQuery = v
	return s
}

type GlobalQueryContextOriginalQuery struct {
	Page      *string `json:"page,omitempty" xml:"page,omitempty"`
	Query     *string `json:"query,omitempty" xml:"query,omitempty"`
	TimeRange *string `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
}

func (s GlobalQueryContextOriginalQuery) String() string {
	return tea.Prettify(s)
}

func (s GlobalQueryContextOriginalQuery) GoString() string {
	return s.String()
}

func (s *GlobalQueryContextOriginalQuery) SetPage(v string) *GlobalQueryContextOriginalQuery {
	s.Page = &v
	return s
}

func (s *GlobalQueryContextOriginalQuery) SetQuery(v string) *GlobalQueryContextOriginalQuery {
	s.Query = &v
	return s
}

func (s *GlobalQueryContextOriginalQuery) SetTimeRange(v string) *GlobalQueryContextOriginalQuery {
	s.TimeRange = &v
	return s
}

type GlobalSceneItem struct {
	Detail *string `json:"detail,omitempty" xml:"detail,omitempty"`
	Type   *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GlobalSceneItem) String() string {
	return tea.Prettify(s)
}

func (s GlobalSceneItem) GoString() string {
	return s.String()
}

func (s *GlobalSceneItem) SetDetail(v string) *GlobalSceneItem {
	s.Detail = &v
	return s
}

func (s *GlobalSceneItem) SetType(v string) *GlobalSceneItem {
	s.Type = &v
	return s
}

type GlobalSearchInformation struct {
	SearchTime *int64 `json:"searchTime,omitempty" xml:"searchTime,omitempty"`
	Total      *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GlobalSearchInformation) String() string {
	return tea.Prettify(s)
}

func (s GlobalSearchInformation) GoString() string {
	return s.String()
}

func (s *GlobalSearchInformation) SetSearchTime(v int64) *GlobalSearchInformation {
	s.SearchTime = &v
	return s
}

func (s *GlobalSearchInformation) SetTotal(v int64) *GlobalSearchInformation {
	s.Total = &v
	return s
}

type GlobalSearchResult struct {
	PageItems    []*GlobalPageItem   `json:"pageItems,omitempty" xml:"pageItems,omitempty" type:"Repeated"`
	QueryContext *GlobalQueryContext `json:"queryContext,omitempty" xml:"queryContext,omitempty"`
	// example:
	//
	// 123456
	RequestId         *string                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SceneItems        []*GlobalSceneItem       `json:"sceneItems,omitempty" xml:"sceneItems,omitempty" type:"Repeated"`
	SearchInformation *GlobalSearchInformation `json:"searchInformation,omitempty" xml:"searchInformation,omitempty"`
}

func (s GlobalSearchResult) String() string {
	return tea.Prettify(s)
}

func (s GlobalSearchResult) GoString() string {
	return s.String()
}

func (s *GlobalSearchResult) SetPageItems(v []*GlobalPageItem) *GlobalSearchResult {
	s.PageItems = v
	return s
}

func (s *GlobalSearchResult) SetQueryContext(v *GlobalQueryContext) *GlobalSearchResult {
	s.QueryContext = v
	return s
}

func (s *GlobalSearchResult) SetRequestId(v string) *GlobalSearchResult {
	s.RequestId = &v
	return s
}

func (s *GlobalSearchResult) SetSceneItems(v []*GlobalSceneItem) *GlobalSearchResult {
	s.SceneItems = v
	return s
}

func (s *GlobalSearchResult) SetSearchInformation(v *GlobalSearchInformation) *GlobalSearchResult {
	s.SearchInformation = v
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

type QueryContext struct {
	OriginalQuery *QueryContextOriginalQuery `json:"originalQuery,omitempty" xml:"originalQuery,omitempty" type:"Struct"`
	Rewrite       *QueryContextRewrite       `json:"rewrite,omitempty" xml:"rewrite,omitempty" type:"Struct"`
}

func (s QueryContext) String() string {
	return tea.Prettify(s)
}

func (s QueryContext) GoString() string {
	return s.String()
}

func (s *QueryContext) SetOriginalQuery(v *QueryContextOriginalQuery) *QueryContext {
	s.OriginalQuery = v
	return s
}

func (s *QueryContext) SetRewrite(v *QueryContextRewrite) *QueryContext {
	s.Rewrite = v
	return s
}

type QueryContextOriginalQuery struct {
	Industry  *string `json:"industry,omitempty" xml:"industry,omitempty"`
	Page      *string `json:"page,omitempty" xml:"page,omitempty"`
	Query     *string `json:"query,omitempty" xml:"query,omitempty"`
	TimeRange *string `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
}

func (s QueryContextOriginalQuery) String() string {
	return tea.Prettify(s)
}

func (s QueryContextOriginalQuery) GoString() string {
	return s.String()
}

func (s *QueryContextOriginalQuery) SetIndustry(v string) *QueryContextOriginalQuery {
	s.Industry = &v
	return s
}

func (s *QueryContextOriginalQuery) SetPage(v string) *QueryContextOriginalQuery {
	s.Page = &v
	return s
}

func (s *QueryContextOriginalQuery) SetQuery(v string) *QueryContextOriginalQuery {
	s.Query = &v
	return s
}

func (s *QueryContextOriginalQuery) SetTimeRange(v string) *QueryContextOriginalQuery {
	s.TimeRange = &v
	return s
}

type QueryContextRewrite struct {
	Enabled   *bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
	TimeRange *string `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
}

func (s QueryContextRewrite) String() string {
	return tea.Prettify(s)
}

func (s QueryContextRewrite) GoString() string {
	return s.String()
}

func (s *QueryContextRewrite) SetEnabled(v bool) *QueryContextRewrite {
	s.Enabled = &v
	return s
}

func (s *QueryContextRewrite) SetTimeRange(v string) *QueryContextRewrite {
	s.TimeRange = &v
	return s
}

type RequestContents struct {
	MainText     *bool `json:"mainText,omitempty" xml:"mainText,omitempty"`
	MarkdownText *bool `json:"markdownText,omitempty" xml:"markdownText,omitempty"`
	RerankScore  *bool `json:"rerankScore,omitempty" xml:"rerankScore,omitempty"`
	Summary      *bool `json:"summary,omitempty" xml:"summary,omitempty"`
}

func (s RequestContents) String() string {
	return tea.Prettify(s)
}

func (s RequestContents) GoString() string {
	return s.String()
}

func (s *RequestContents) SetMainText(v bool) *RequestContents {
	s.MainText = &v
	return s
}

func (s *RequestContents) SetMarkdownText(v bool) *RequestContents {
	s.MarkdownText = &v
	return s
}

func (s *RequestContents) SetRerankScore(v bool) *RequestContents {
	s.RerankScore = &v
	return s
}

func (s *RequestContents) SetSummary(v bool) *RequestContents {
	s.Summary = &v
	return s
}

type SceneItem struct {
	Detail *string `json:"detail,omitempty" xml:"detail,omitempty"`
	Type   *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SceneItem) String() string {
	return tea.Prettify(s)
}

func (s SceneItem) GoString() string {
	return s.String()
}

func (s *SceneItem) SetDetail(v string) *SceneItem {
	s.Detail = &v
	return s
}

func (s *SceneItem) SetType(v string) *SceneItem {
	s.Type = &v
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
	// example:
	//
	// https://s2.zimgs.cn/ims?kt=url&at=smstruct&key=aHR0cHM6Ly9ndy5hbGljZG4uY29tL0wxLzcyMy8xNTY1MjU2NjAwLzJhL2YwL2I0LzJhZjBiNDQxMGI5YmVlMDVjOGVlNGJmODk3MTNkNTFjLnBuZw==&sign=yx:CUlNNQVJQjFrk3Kxt2F3KWhTOFU=&tv=400_400
	HostLogo *string `json:"hostLogo,omitempty" xml:"hostLogo,omitempty"`
	// example:
	//
	// 新华网
	Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty"`
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
	MainText     *string `json:"mainText,omitempty" xml:"mainText,omitempty"`
	MarkdownText *string `json:"markdownText,omitempty" xml:"markdownText,omitempty"`
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
	// example:
	//
	// 权威媒体
	SiteLabel *string `json:"siteLabel,omitempty" xml:"siteLabel,omitempty"`
	// example:
	//
	// 100km/h-0制动能力上，仅有33.3m，不黑不吹，单看这个，小米SU7确实表现不错。而续航方面，101kWh电池容量，实现CLTC续航800km，还有现5分钟补能220km，15分钟补能510km的800V高压平台。而在...
	Snippet *string `json:"snippet,omitempty" xml:"snippet,omitempty"`
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
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

func (s *ScorePageItem) SetHostLogo(v string) *ScorePageItem {
	s.HostLogo = &v
	return s
}

func (s *ScorePageItem) SetHostname(v string) *ScorePageItem {
	s.Hostname = &v
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

func (s *ScorePageItem) SetMarkdownText(v string) *ScorePageItem {
	s.MarkdownText = &v
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

func (s *ScorePageItem) SetSiteLabel(v string) *ScorePageItem {
	s.SiteLabel = &v
	return s
}

func (s *ScorePageItem) SetSnippet(v string) *ScorePageItem {
	s.Snippet = &v
	return s
}

func (s *ScorePageItem) SetSummary(v string) *ScorePageItem {
	s.Summary = &v
	return s
}

func (s *ScorePageItem) SetTitle(v string) *ScorePageItem {
	s.Title = &v
	return s
}

type SearchCredits struct {
	GenericTextSearch *int32 `json:"genericTextSearch,omitempty" xml:"genericTextSearch,omitempty"`
}

func (s SearchCredits) String() string {
	return tea.Prettify(s)
}

func (s SearchCredits) GoString() string {
	return s.String()
}

func (s *SearchCredits) SetGenericTextSearch(v int32) *SearchCredits {
	s.GenericTextSearch = &v
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

type UnifiedCostCredits struct {
	Search     *SearchCredits     `json:"search,omitempty" xml:"search,omitempty"`
	ValueAdded *ValueAddedCredits `json:"valueAdded,omitempty" xml:"valueAdded,omitempty"`
}

func (s UnifiedCostCredits) String() string {
	return tea.Prettify(s)
}

func (s UnifiedCostCredits) GoString() string {
	return s.String()
}

func (s *UnifiedCostCredits) SetSearch(v *SearchCredits) *UnifiedCostCredits {
	s.Search = v
	return s
}

func (s *UnifiedCostCredits) SetValueAdded(v *ValueAddedCredits) *UnifiedCostCredits {
	s.ValueAdded = v
	return s
}

type UnifiedOriginalQuery struct {
	Query     *string `json:"query,omitempty" xml:"query,omitempty"`
	TimeRange *string `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
}

func (s UnifiedOriginalQuery) String() string {
	return tea.Prettify(s)
}

func (s UnifiedOriginalQuery) GoString() string {
	return s.String()
}

func (s *UnifiedOriginalQuery) SetQuery(v string) *UnifiedOriginalQuery {
	s.Query = &v
	return s
}

func (s *UnifiedOriginalQuery) SetTimeRange(v string) *UnifiedOriginalQuery {
	s.TimeRange = &v
	return s
}

type UnifiedPageItem struct {
	HostLogo      *string   `json:"hostLogo,omitempty" xml:"hostLogo,omitempty"`
	Hostname      *string   `json:"hostname,omitempty" xml:"hostname,omitempty"`
	Images        []*string `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	Link          *string   `json:"link,omitempty" xml:"link,omitempty"`
	MainText      *string   `json:"mainText,omitempty" xml:"mainText,omitempty"`
	MarkdownText  *string   `json:"markdownText,omitempty" xml:"markdownText,omitempty"`
	PublishedTime *string   `json:"publishedTime,omitempty" xml:"publishedTime,omitempty"`
	RerankScore   *float64  `json:"rerankScore,omitempty" xml:"rerankScore,omitempty"`
	// example:
	//
	// 2025-04-07T10:15:30.123+08:00
	Snippet *string `json:"snippet,omitempty" xml:"snippet,omitempty"`
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	Title   *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s UnifiedPageItem) String() string {
	return tea.Prettify(s)
}

func (s UnifiedPageItem) GoString() string {
	return s.String()
}

func (s *UnifiedPageItem) SetHostLogo(v string) *UnifiedPageItem {
	s.HostLogo = &v
	return s
}

func (s *UnifiedPageItem) SetHostname(v string) *UnifiedPageItem {
	s.Hostname = &v
	return s
}

func (s *UnifiedPageItem) SetImages(v []*string) *UnifiedPageItem {
	s.Images = v
	return s
}

func (s *UnifiedPageItem) SetLink(v string) *UnifiedPageItem {
	s.Link = &v
	return s
}

func (s *UnifiedPageItem) SetMainText(v string) *UnifiedPageItem {
	s.MainText = &v
	return s
}

func (s *UnifiedPageItem) SetMarkdownText(v string) *UnifiedPageItem {
	s.MarkdownText = &v
	return s
}

func (s *UnifiedPageItem) SetPublishedTime(v string) *UnifiedPageItem {
	s.PublishedTime = &v
	return s
}

func (s *UnifiedPageItem) SetRerankScore(v float64) *UnifiedPageItem {
	s.RerankScore = &v
	return s
}

func (s *UnifiedPageItem) SetSnippet(v string) *UnifiedPageItem {
	s.Snippet = &v
	return s
}

func (s *UnifiedPageItem) SetSummary(v string) *UnifiedPageItem {
	s.Summary = &v
	return s
}

func (s *UnifiedPageItem) SetTitle(v string) *UnifiedPageItem {
	s.Title = &v
	return s
}

type UnifiedQueryContext struct {
	EngineType    *string               `json:"engineType,omitempty" xml:"engineType,omitempty"`
	OriginalQuery *UnifiedOriginalQuery `json:"originalQuery,omitempty" xml:"originalQuery,omitempty"`
	Rewrite       *UnifiedRewrite       `json:"rewrite,omitempty" xml:"rewrite,omitempty"`
}

func (s UnifiedQueryContext) String() string {
	return tea.Prettify(s)
}

func (s UnifiedQueryContext) GoString() string {
	return s.String()
}

func (s *UnifiedQueryContext) SetEngineType(v string) *UnifiedQueryContext {
	s.EngineType = &v
	return s
}

func (s *UnifiedQueryContext) SetOriginalQuery(v *UnifiedOriginalQuery) *UnifiedQueryContext {
	s.OriginalQuery = v
	return s
}

func (s *UnifiedQueryContext) SetRewrite(v *UnifiedRewrite) *UnifiedQueryContext {
	s.Rewrite = v
	return s
}

type UnifiedRewrite struct {
	Enabled   *bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
	TimeRange *string `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
}

func (s UnifiedRewrite) String() string {
	return tea.Prettify(s)
}

func (s UnifiedRewrite) GoString() string {
	return s.String()
}

func (s *UnifiedRewrite) SetEnabled(v bool) *UnifiedRewrite {
	s.Enabled = &v
	return s
}

func (s *UnifiedRewrite) SetTimeRange(v string) *UnifiedRewrite {
	s.TimeRange = &v
	return s
}

type UnifiedSceneItem struct {
	Detail *string `json:"detail,omitempty" xml:"detail,omitempty"`
	Type   *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UnifiedSceneItem) String() string {
	return tea.Prettify(s)
}

func (s UnifiedSceneItem) GoString() string {
	return s.String()
}

func (s *UnifiedSceneItem) SetDetail(v string) *UnifiedSceneItem {
	s.Detail = &v
	return s
}

func (s *UnifiedSceneItem) SetType(v string) *UnifiedSceneItem {
	s.Type = &v
	return s
}

type UnifiedSearchInformation struct {
	SearchTime *int64 `json:"searchTime,omitempty" xml:"searchTime,omitempty"`
}

func (s UnifiedSearchInformation) String() string {
	return tea.Prettify(s)
}

func (s UnifiedSearchInformation) GoString() string {
	return s.String()
}

func (s *UnifiedSearchInformation) SetSearchTime(v int64) *UnifiedSearchInformation {
	s.SearchTime = &v
	return s
}

type UnifiedSearchRequest struct {
	Category   *string          `json:"category,omitempty" xml:"category,omitempty"`
	Contents   *RequestContents `json:"contents,omitempty" xml:"contents,omitempty"`
	EngineType *string          `json:"engineType,omitempty" xml:"engineType,omitempty"`
	Query      *string          `json:"query,omitempty" xml:"query,omitempty"`
	TimeRange  *string          `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
}

func (s UnifiedSearchRequest) String() string {
	return tea.Prettify(s)
}

func (s UnifiedSearchRequest) GoString() string {
	return s.String()
}

func (s *UnifiedSearchRequest) SetCategory(v string) *UnifiedSearchRequest {
	s.Category = &v
	return s
}

func (s *UnifiedSearchRequest) SetContents(v *RequestContents) *UnifiedSearchRequest {
	s.Contents = v
	return s
}

func (s *UnifiedSearchRequest) SetEngineType(v string) *UnifiedSearchRequest {
	s.EngineType = &v
	return s
}

func (s *UnifiedSearchRequest) SetQuery(v string) *UnifiedSearchRequest {
	s.Query = &v
	return s
}

func (s *UnifiedSearchRequest) SetTimeRange(v string) *UnifiedSearchRequest {
	s.TimeRange = &v
	return s
}

type UnifiedSearchResponse struct {
	CostCredits       *UnifiedCostCredits       `json:"costCredits,omitempty" xml:"costCredits,omitempty"`
	PageItems         []*UnifiedPageItem        `json:"pageItems,omitempty" xml:"pageItems,omitempty" type:"Repeated"`
	QueryContext      *UnifiedQueryContext      `json:"queryContext,omitempty" xml:"queryContext,omitempty"`
	RequestId         *string                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SceneItems        []*UnifiedSceneItem       `json:"sceneItems,omitempty" xml:"sceneItems,omitempty" type:"Repeated"`
	SearchInformation *UnifiedSearchInformation `json:"searchInformation,omitempty" xml:"searchInformation,omitempty"`
}

func (s UnifiedSearchResponse) String() string {
	return tea.Prettify(s)
}

func (s UnifiedSearchResponse) GoString() string {
	return s.String()
}

func (s *UnifiedSearchResponse) SetCostCredits(v *UnifiedCostCredits) *UnifiedSearchResponse {
	s.CostCredits = v
	return s
}

func (s *UnifiedSearchResponse) SetPageItems(v []*UnifiedPageItem) *UnifiedSearchResponse {
	s.PageItems = v
	return s
}

func (s *UnifiedSearchResponse) SetQueryContext(v *UnifiedQueryContext) *UnifiedSearchResponse {
	s.QueryContext = v
	return s
}

func (s *UnifiedSearchResponse) SetRequestId(v string) *UnifiedSearchResponse {
	s.RequestId = &v
	return s
}

func (s *UnifiedSearchResponse) SetSceneItems(v []*UnifiedSceneItem) *UnifiedSearchResponse {
	s.SceneItems = v
	return s
}

func (s *UnifiedSearchResponse) SetSearchInformation(v *UnifiedSearchInformation) *UnifiedSearchResponse {
	s.SearchInformation = v
	return s
}

type ValueAddedCredits struct {
	Advanced *int32 `json:"advanced,omitempty" xml:"advanced,omitempty"`
	Summary  *int32 `json:"summary,omitempty" xml:"summary,omitempty"`
}

func (s ValueAddedCredits) String() string {
	return tea.Prettify(s)
}

func (s ValueAddedCredits) GoString() string {
	return s.String()
}

func (s *ValueAddedCredits) SetAdvanced(v int32) *ValueAddedCredits {
	s.Advanced = &v
	return s
}

func (s *ValueAddedCredits) SetSummary(v int32) *ValueAddedCredits {
	s.Summary = &v
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

type AiSearchRequest struct {
	// example:
	//
	// finance
	Industry *string `json:"industry,omitempty" xml:"industry,omitempty"`
	// example:
	//
	// {\\"total_count\\": 6851, \\"page_number\\": 54, \\"page_size\\": 100}
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// This parameter is required.
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// example:
	//
	// 17dc8bcd-f34a-46d1-a7a3-0fa3d1ce3824
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// OneWeek
	TimeRange *string `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
}

func (s AiSearchRequest) String() string {
	return tea.Prettify(s)
}

func (s AiSearchRequest) GoString() string {
	return s.String()
}

func (s *AiSearchRequest) SetIndustry(v string) *AiSearchRequest {
	s.Industry = &v
	return s
}

func (s *AiSearchRequest) SetPage(v int32) *AiSearchRequest {
	s.Page = &v
	return s
}

func (s *AiSearchRequest) SetQuery(v string) *AiSearchRequest {
	s.Query = &v
	return s
}

func (s *AiSearchRequest) SetSessionId(v string) *AiSearchRequest {
	s.SessionId = &v
	return s
}

func (s *AiSearchRequest) SetTimeRange(v string) *AiSearchRequest {
	s.TimeRange = &v
	return s
}

type AiSearchResponseBody struct {
	Header *AiSearchResponseBodyHeader `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	// example:
	//
	// {"header":{"eventId":"6f617de0-204f-406f-a9be-34779c06d498","event":"on_common_search_start","responseTime":120},"payload":"","requestId":"715d01a0-de7e-42c3-abca-b901fcd79b39"}
	Payload *string `json:"payload,omitempty" xml:"payload,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ECB2144C-E277-5434-80E6-12D26678D364
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AiSearchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AiSearchResponseBody) GoString() string {
	return s.String()
}

func (s *AiSearchResponseBody) SetHeader(v *AiSearchResponseBodyHeader) *AiSearchResponseBody {
	s.Header = v
	return s
}

func (s *AiSearchResponseBody) SetPayload(v string) *AiSearchResponseBody {
	s.Payload = &v
	return s
}

func (s *AiSearchResponseBody) SetRequestId(v string) *AiSearchResponseBody {
	s.RequestId = &v
	return s
}

type AiSearchResponseBodyHeader struct {
	// example:
	//
	// on_common_search_end
	Event *string `json:"event,omitempty" xml:"event,omitempty"`
	// example:
	//
	// 988021f0-951a-43d0-ba4d-785359e7e7be
	EventId      *string                                 `json:"eventId,omitempty" xml:"eventId,omitempty"`
	QueryContext *AiSearchResponseBodyHeaderQueryContext `json:"queryContext,omitempty" xml:"queryContext,omitempty" type:"Struct"`
	// example:
	//
	// 1293
	ResponseTime *int64 `json:"responseTime,omitempty" xml:"responseTime,omitempty"`
}

func (s AiSearchResponseBodyHeader) String() string {
	return tea.Prettify(s)
}

func (s AiSearchResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *AiSearchResponseBodyHeader) SetEvent(v string) *AiSearchResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *AiSearchResponseBodyHeader) SetEventId(v string) *AiSearchResponseBodyHeader {
	s.EventId = &v
	return s
}

func (s *AiSearchResponseBodyHeader) SetQueryContext(v *AiSearchResponseBodyHeaderQueryContext) *AiSearchResponseBodyHeader {
	s.QueryContext = v
	return s
}

func (s *AiSearchResponseBodyHeader) SetResponseTime(v int64) *AiSearchResponseBodyHeader {
	s.ResponseTime = &v
	return s
}

type AiSearchResponseBodyHeaderQueryContext struct {
	OriginalQuery *AiSearchResponseBodyHeaderQueryContextOriginalQuery `json:"originalQuery,omitempty" xml:"originalQuery,omitempty" type:"Struct"`
	Rewrite       *AiSearchResponseBodyHeaderQueryContextRewrite       `json:"rewrite,omitempty" xml:"rewrite,omitempty" type:"Struct"`
}

func (s AiSearchResponseBodyHeaderQueryContext) String() string {
	return tea.Prettify(s)
}

func (s AiSearchResponseBodyHeaderQueryContext) GoString() string {
	return s.String()
}

func (s *AiSearchResponseBodyHeaderQueryContext) SetOriginalQuery(v *AiSearchResponseBodyHeaderQueryContextOriginalQuery) *AiSearchResponseBodyHeaderQueryContext {
	s.OriginalQuery = v
	return s
}

func (s *AiSearchResponseBodyHeaderQueryContext) SetRewrite(v *AiSearchResponseBodyHeaderQueryContextRewrite) *AiSearchResponseBodyHeaderQueryContext {
	s.Rewrite = v
	return s
}

type AiSearchResponseBodyHeaderQueryContextOriginalQuery struct {
	Industry  *string `json:"industry,omitempty" xml:"industry,omitempty"`
	Page      *int32  `json:"page,omitempty" xml:"page,omitempty"`
	Query     *string `json:"query,omitempty" xml:"query,omitempty"`
	TimeRange *string `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
}

func (s AiSearchResponseBodyHeaderQueryContextOriginalQuery) String() string {
	return tea.Prettify(s)
}

func (s AiSearchResponseBodyHeaderQueryContextOriginalQuery) GoString() string {
	return s.String()
}

func (s *AiSearchResponseBodyHeaderQueryContextOriginalQuery) SetIndustry(v string) *AiSearchResponseBodyHeaderQueryContextOriginalQuery {
	s.Industry = &v
	return s
}

func (s *AiSearchResponseBodyHeaderQueryContextOriginalQuery) SetPage(v int32) *AiSearchResponseBodyHeaderQueryContextOriginalQuery {
	s.Page = &v
	return s
}

func (s *AiSearchResponseBodyHeaderQueryContextOriginalQuery) SetQuery(v string) *AiSearchResponseBodyHeaderQueryContextOriginalQuery {
	s.Query = &v
	return s
}

func (s *AiSearchResponseBodyHeaderQueryContextOriginalQuery) SetTimeRange(v string) *AiSearchResponseBodyHeaderQueryContextOriginalQuery {
	s.TimeRange = &v
	return s
}

type AiSearchResponseBodyHeaderQueryContextRewrite struct {
	Enabled   *bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
	TimeRange *string `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
}

func (s AiSearchResponseBodyHeaderQueryContextRewrite) String() string {
	return tea.Prettify(s)
}

func (s AiSearchResponseBodyHeaderQueryContextRewrite) GoString() string {
	return s.String()
}

func (s *AiSearchResponseBodyHeaderQueryContextRewrite) SetEnabled(v bool) *AiSearchResponseBodyHeaderQueryContextRewrite {
	s.Enabled = &v
	return s
}

func (s *AiSearchResponseBodyHeaderQueryContextRewrite) SetTimeRange(v string) *AiSearchResponseBodyHeaderQueryContextRewrite {
	s.TimeRange = &v
	return s
}

type AiSearchResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AiSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AiSearchResponse) String() string {
	return tea.Prettify(s)
}

func (s AiSearchResponse) GoString() string {
	return s.String()
}

func (s *AiSearchResponse) SetHeaders(v map[string]*string) *AiSearchResponse {
	s.Headers = v
	return s
}

func (s *AiSearchResponse) SetStatusCode(v int32) *AiSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *AiSearchResponse) SetBody(v *AiSearchResponseBody) *AiSearchResponse {
	s.Body = v
	return s
}

type GenericAdvancedSearchRequest struct {
	Industry *string `json:"industry,omitempty" xml:"industry,omitempty"`
	// This parameter is required.
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// example:
	//
	// job-4065bee3-e7aa-49fc-aad2-a8e3a7fd6acd
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// OneWeek
	TimeRange *string `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
}

func (s GenericAdvancedSearchRequest) String() string {
	return tea.Prettify(s)
}

func (s GenericAdvancedSearchRequest) GoString() string {
	return s.String()
}

func (s *GenericAdvancedSearchRequest) SetIndustry(v string) *GenericAdvancedSearchRequest {
	s.Industry = &v
	return s
}

func (s *GenericAdvancedSearchRequest) SetQuery(v string) *GenericAdvancedSearchRequest {
	s.Query = &v
	return s
}

func (s *GenericAdvancedSearchRequest) SetSessionId(v string) *GenericAdvancedSearchRequest {
	s.SessionId = &v
	return s
}

func (s *GenericAdvancedSearchRequest) SetTimeRange(v string) *GenericAdvancedSearchRequest {
	s.TimeRange = &v
	return s
}

type GenericAdvancedSearchResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenericSearchResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenericAdvancedSearchResponse) String() string {
	return tea.Prettify(s)
}

func (s GenericAdvancedSearchResponse) GoString() string {
	return s.String()
}

func (s *GenericAdvancedSearchResponse) SetHeaders(v map[string]*string) *GenericAdvancedSearchResponse {
	s.Headers = v
	return s
}

func (s *GenericAdvancedSearchResponse) SetStatusCode(v int32) *GenericAdvancedSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *GenericAdvancedSearchResponse) SetBody(v *GenericSearchResult) *GenericAdvancedSearchResponse {
	s.Body = v
	return s
}

type GenericSearchRequest struct {
	EnableRerank *bool   `json:"enableRerank,omitempty" xml:"enableRerank,omitempty"`
	Industry     *string `json:"industry,omitempty" xml:"industry,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// This parameter is required.
	Query              *string `json:"query,omitempty" xml:"query,omitempty"`
	ReturnMainText     *bool   `json:"returnMainText,omitempty" xml:"returnMainText,omitempty"`
	ReturnMarkdownText *bool   `json:"returnMarkdownText,omitempty" xml:"returnMarkdownText,omitempty"`
	ReturnSummary      *bool   `json:"returnSummary,omitempty" xml:"returnSummary,omitempty"`
	SessionId          *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// OneWeek
	TimeRange *string `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
}

func (s GenericSearchRequest) String() string {
	return tea.Prettify(s)
}

func (s GenericSearchRequest) GoString() string {
	return s.String()
}

func (s *GenericSearchRequest) SetEnableRerank(v bool) *GenericSearchRequest {
	s.EnableRerank = &v
	return s
}

func (s *GenericSearchRequest) SetIndustry(v string) *GenericSearchRequest {
	s.Industry = &v
	return s
}

func (s *GenericSearchRequest) SetPage(v int32) *GenericSearchRequest {
	s.Page = &v
	return s
}

func (s *GenericSearchRequest) SetQuery(v string) *GenericSearchRequest {
	s.Query = &v
	return s
}

func (s *GenericSearchRequest) SetReturnMainText(v bool) *GenericSearchRequest {
	s.ReturnMainText = &v
	return s
}

func (s *GenericSearchRequest) SetReturnMarkdownText(v bool) *GenericSearchRequest {
	s.ReturnMarkdownText = &v
	return s
}

func (s *GenericSearchRequest) SetReturnSummary(v bool) *GenericSearchRequest {
	s.ReturnSummary = &v
	return s
}

func (s *GenericSearchRequest) SetSessionId(v string) *GenericSearchRequest {
	s.SessionId = &v
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

type GlobalSearchRequest struct {
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// This parameter is required.
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// example:
	//
	// OneWeek
	TimeRange *string `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
}

func (s GlobalSearchRequest) String() string {
	return tea.Prettify(s)
}

func (s GlobalSearchRequest) GoString() string {
	return s.String()
}

func (s *GlobalSearchRequest) SetPage(v int32) *GlobalSearchRequest {
	s.Page = &v
	return s
}

func (s *GlobalSearchRequest) SetPageSize(v int32) *GlobalSearchRequest {
	s.PageSize = &v
	return s
}

func (s *GlobalSearchRequest) SetQuery(v string) *GlobalSearchRequest {
	s.Query = &v
	return s
}

func (s *GlobalSearchRequest) SetTimeRange(v string) *GlobalSearchRequest {
	s.TimeRange = &v
	return s
}

type GlobalSearchResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GlobalSearchResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GlobalSearchResponse) String() string {
	return tea.Prettify(s)
}

func (s GlobalSearchResponse) GoString() string {
	return s.String()
}

func (s *GlobalSearchResponse) SetHeaders(v map[string]*string) *GlobalSearchResponse {
	s.Headers = v
	return s
}

func (s *GlobalSearchResponse) SetStatusCode(v int32) *GlobalSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *GlobalSearchResponse) SetBody(v *GlobalSearchResult) *GlobalSearchResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("iqs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// # AI搜索流式接口
//
// @param request - AiSearchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AiSearchResponse
func (client *Client) AiSearchWithOptions(request *AiSearchRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AiSearchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Industry)) {
		query["industry"] = request.Industry
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["sessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.TimeRange)) {
		query["timeRange"] = request.TimeRange
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AiSearch"),
		Version:     tea.String("2024-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/linked-retrieval/linked-retrieval-entry/v3/linkedRetrieval/commands/aiSearch"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AiSearchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # AI搜索流式接口
//
// @param request - AiSearchRequest
//
// @return AiSearchResponse
func (client *Client) AiSearch(request *AiSearchRequest) (_result *AiSearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AiSearchResponse{}
	_body, _err := client.AiSearchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 增强版通用搜索
//
// @param request - GenericAdvancedSearchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenericAdvancedSearchResponse
func (client *Client) GenericAdvancedSearchWithOptions(request *GenericAdvancedSearchRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GenericAdvancedSearchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Industry)) {
		query["industry"] = request.Industry
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["sessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.TimeRange)) {
		query["timeRange"] = request.TimeRange
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenericAdvancedSearch"),
		Version:     tea.String("2024-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/linked-retrieval/linked-retrieval-entry/v2/linkedRetrieval/commands/genericAdvancedSearch"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GenericAdvancedSearchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 增强版通用搜索
//
// @param request - GenericAdvancedSearchRequest
//
// @return GenericAdvancedSearchResponse
func (client *Client) GenericAdvancedSearch(request *GenericAdvancedSearchRequest) (_result *GenericAdvancedSearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenericAdvancedSearchResponse{}
	_body, _err := client.GenericAdvancedSearchWithOptions(request, headers, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.EnableRerank)) {
		query["enableRerank"] = request.EnableRerank
	}

	if !tea.BoolValue(util.IsUnset(request.Industry)) {
		query["industry"] = request.Industry
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnMainText)) {
		query["returnMainText"] = request.ReturnMainText
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnMarkdownText)) {
		query["returnMarkdownText"] = request.ReturnMarkdownText
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnSummary)) {
		query["returnSummary"] = request.ReturnSummary
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["sessionId"] = request.SessionId
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
		Version:     tea.String("2024-11-11"),
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

// Summary:
//
// 通晓搜索-出海版(全球信息搜索)
//
// @param request - GlobalSearchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GlobalSearchResponse
func (client *Client) GlobalSearchWithOptions(request *GlobalSearchRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GlobalSearchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

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
		Action:      tea.String("GlobalSearch"),
		Version:     tea.String("2024-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/linked-retrieval/linked-retrieval-entry/v1/iqs/search/global"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GlobalSearchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通晓搜索-出海版(全球信息搜索)
//
// @param request - GlobalSearchRequest
//
// @return GlobalSearchResponse
func (client *Client) GlobalSearch(request *GlobalSearchRequest) (_result *GlobalSearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GlobalSearchResponse{}
	_body, _err := client.GlobalSearchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
