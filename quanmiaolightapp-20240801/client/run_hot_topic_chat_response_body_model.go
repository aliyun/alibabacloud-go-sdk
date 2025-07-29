// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunHotTopicChatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeader(v *RunHotTopicChatResponseBodyHeader) *RunHotTopicChatResponseBody
	GetHeader() *RunHotTopicChatResponseBodyHeader
	SetPayload(v *RunHotTopicChatResponseBodyPayload) *RunHotTopicChatResponseBody
	GetPayload() *RunHotTopicChatResponseBodyPayload
	SetRequestId(v string) *RunHotTopicChatResponseBody
	GetRequestId() *string
}

type RunHotTopicChatResponseBody struct {
	Header  *RunHotTopicChatResponseBodyHeader  `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload *RunHotTopicChatResponseBodyPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
	// example:
	//
	// 04DA1A52-4E51-56CB-BA64-FDDA0B53BAE8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RunHotTopicChatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunHotTopicChatResponseBody) GoString() string {
	return s.String()
}

func (s *RunHotTopicChatResponseBody) GetHeader() *RunHotTopicChatResponseBodyHeader {
	return s.Header
}

func (s *RunHotTopicChatResponseBody) GetPayload() *RunHotTopicChatResponseBodyPayload {
	return s.Payload
}

func (s *RunHotTopicChatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunHotTopicChatResponseBody) SetHeader(v *RunHotTopicChatResponseBodyHeader) *RunHotTopicChatResponseBody {
	s.Header = v
	return s
}

func (s *RunHotTopicChatResponseBody) SetPayload(v *RunHotTopicChatResponseBodyPayload) *RunHotTopicChatResponseBody {
	s.Payload = v
	return s
}

func (s *RunHotTopicChatResponseBody) SetRequestId(v string) *RunHotTopicChatResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunHotTopicChatResponseBody) Validate() error {
	return dara.Validate(s)
}

type RunHotTopicChatResponseBodyHeader struct {
	// example:
	//
	// InvalidParam
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// xx
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// task-finished
	Event *string `json:"event,omitempty" xml:"event,omitempty"`
	// example:
	//
	// xx
	EventInfo *string `json:"eventInfo,omitempty" xml:"eventInfo,omitempty"`
	// example:
	//
	// xxx
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// d3be9981-ca2d-4e17-bf31-1c0a628e9f99
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 2150451a17191950923411783e2927
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s RunHotTopicChatResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunHotTopicChatResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunHotTopicChatResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunHotTopicChatResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunHotTopicChatResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunHotTopicChatResponseBodyHeader) GetEventInfo() *string {
	return s.EventInfo
}

func (s *RunHotTopicChatResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunHotTopicChatResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunHotTopicChatResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunHotTopicChatResponseBodyHeader) SetErrorCode(v string) *RunHotTopicChatResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunHotTopicChatResponseBodyHeader) SetErrorMessage(v string) *RunHotTopicChatResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunHotTopicChatResponseBodyHeader) SetEvent(v string) *RunHotTopicChatResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunHotTopicChatResponseBodyHeader) SetEventInfo(v string) *RunHotTopicChatResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunHotTopicChatResponseBodyHeader) SetSessionId(v string) *RunHotTopicChatResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunHotTopicChatResponseBodyHeader) SetTaskId(v string) *RunHotTopicChatResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunHotTopicChatResponseBodyHeader) SetTraceId(v string) *RunHotTopicChatResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunHotTopicChatResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunHotTopicChatResponseBodyPayload struct {
	Output *RunHotTopicChatResponseBodyPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Usage  *RunHotTopicChatResponseBodyPayloadUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunHotTopicChatResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunHotTopicChatResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunHotTopicChatResponseBodyPayload) GetOutput() *RunHotTopicChatResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunHotTopicChatResponseBodyPayload) GetUsage() *RunHotTopicChatResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunHotTopicChatResponseBodyPayload) SetOutput(v *RunHotTopicChatResponseBodyPayloadOutput) *RunHotTopicChatResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunHotTopicChatResponseBodyPayload) SetUsage(v *RunHotTopicChatResponseBodyPayloadUsage) *RunHotTopicChatResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunHotTopicChatResponseBodyPayload) Validate() error {
	return dara.Validate(s)
}

type RunHotTopicChatResponseBodyPayloadOutput struct {
	Articles          []*RunHotTopicChatResponseBodyPayloadOutputArticles          `json:"articles,omitempty" xml:"articles,omitempty" type:"Repeated"`
	Category          *string                                                      `json:"category,omitempty" xml:"category,omitempty"`
	HotTopicSummaries []*RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries `json:"hotTopicSummaries,omitempty" xml:"hotTopicSummaries,omitempty" type:"Repeated"`
	Keyword           *string                                                      `json:"keyword,omitempty" xml:"keyword,omitempty"`
	Location          *string                                                      `json:"location,omitempty" xml:"location,omitempty"`
	MultimodalMedias  []*RunHotTopicChatResponseBodyPayloadOutputMultimodalMedias  `json:"multimodalMedias,omitempty" xml:"multimodalMedias,omitempty" type:"Repeated"`
	RecommendQueries  []*string                                                    `json:"recommendQueries,omitempty" xml:"recommendQueries,omitempty" type:"Repeated"`
	SearchQuery       *string                                                      `json:"searchQuery,omitempty" xml:"searchQuery,omitempty"`
	// example:
	//
	// xx
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s RunHotTopicChatResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunHotTopicChatResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunHotTopicChatResponseBodyPayloadOutput) GetArticles() []*RunHotTopicChatResponseBodyPayloadOutputArticles {
	return s.Articles
}

func (s *RunHotTopicChatResponseBodyPayloadOutput) GetCategory() *string {
	return s.Category
}

func (s *RunHotTopicChatResponseBodyPayloadOutput) GetHotTopicSummaries() []*RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries {
	return s.HotTopicSummaries
}

func (s *RunHotTopicChatResponseBodyPayloadOutput) GetKeyword() *string {
	return s.Keyword
}

func (s *RunHotTopicChatResponseBodyPayloadOutput) GetLocation() *string {
	return s.Location
}

func (s *RunHotTopicChatResponseBodyPayloadOutput) GetMultimodalMedias() []*RunHotTopicChatResponseBodyPayloadOutputMultimodalMedias {
	return s.MultimodalMedias
}

func (s *RunHotTopicChatResponseBodyPayloadOutput) GetRecommendQueries() []*string {
	return s.RecommendQueries
}

func (s *RunHotTopicChatResponseBodyPayloadOutput) GetSearchQuery() *string {
	return s.SearchQuery
}

func (s *RunHotTopicChatResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunHotTopicChatResponseBodyPayloadOutput) SetArticles(v []*RunHotTopicChatResponseBodyPayloadOutputArticles) *RunHotTopicChatResponseBodyPayloadOutput {
	s.Articles = v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutput) SetCategory(v string) *RunHotTopicChatResponseBodyPayloadOutput {
	s.Category = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutput) SetHotTopicSummaries(v []*RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) *RunHotTopicChatResponseBodyPayloadOutput {
	s.HotTopicSummaries = v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutput) SetKeyword(v string) *RunHotTopicChatResponseBodyPayloadOutput {
	s.Keyword = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutput) SetLocation(v string) *RunHotTopicChatResponseBodyPayloadOutput {
	s.Location = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutput) SetMultimodalMedias(v []*RunHotTopicChatResponseBodyPayloadOutputMultimodalMedias) *RunHotTopicChatResponseBodyPayloadOutput {
	s.MultimodalMedias = v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutput) SetRecommendQueries(v []*string) *RunHotTopicChatResponseBodyPayloadOutput {
	s.RecommendQueries = v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutput) SetSearchQuery(v string) *RunHotTopicChatResponseBodyPayloadOutput {
	s.SearchQuery = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutput) SetText(v string) *RunHotTopicChatResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunHotTopicChatResponseBodyPayloadOutputArticles struct {
	// example:
	//
	// xxx
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 2024-09-22 16:45:06
	PubTime          *string  `json:"pubTime,omitempty" xml:"pubTime,omitempty"`
	Score            *float64 `json:"score,omitempty" xml:"score,omitempty"`
	SearchSourceName *string  `json:"searchSourceName,omitempty" xml:"searchSourceName,omitempty"`
	Select           *bool    `json:"select,omitempty" xml:"select,omitempty"`
	// example:
	//
	// xx
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// example:
	//
	// test
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// http://xxx
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s RunHotTopicChatResponseBodyPayloadOutputArticles) String() string {
	return dara.Prettify(s)
}

func (s RunHotTopicChatResponseBodyPayloadOutputArticles) GoString() string {
	return s.String()
}

func (s *RunHotTopicChatResponseBodyPayloadOutputArticles) GetContent() *string {
	return s.Content
}

func (s *RunHotTopicChatResponseBodyPayloadOutputArticles) GetPubTime() *string {
	return s.PubTime
}

func (s *RunHotTopicChatResponseBodyPayloadOutputArticles) GetScore() *float64 {
	return s.Score
}

func (s *RunHotTopicChatResponseBodyPayloadOutputArticles) GetSearchSourceName() *string {
	return s.SearchSourceName
}

func (s *RunHotTopicChatResponseBodyPayloadOutputArticles) GetSelect() *bool {
	return s.Select
}

func (s *RunHotTopicChatResponseBodyPayloadOutputArticles) GetSummary() *string {
	return s.Summary
}

func (s *RunHotTopicChatResponseBodyPayloadOutputArticles) GetTitle() *string {
	return s.Title
}

func (s *RunHotTopicChatResponseBodyPayloadOutputArticles) GetUrl() *string {
	return s.Url
}

func (s *RunHotTopicChatResponseBodyPayloadOutputArticles) SetContent(v string) *RunHotTopicChatResponseBodyPayloadOutputArticles {
	s.Content = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputArticles) SetPubTime(v string) *RunHotTopicChatResponseBodyPayloadOutputArticles {
	s.PubTime = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputArticles) SetScore(v float64) *RunHotTopicChatResponseBodyPayloadOutputArticles {
	s.Score = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputArticles) SetSearchSourceName(v string) *RunHotTopicChatResponseBodyPayloadOutputArticles {
	s.SearchSourceName = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputArticles) SetSelect(v bool) *RunHotTopicChatResponseBodyPayloadOutputArticles {
	s.Select = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputArticles) SetSummary(v string) *RunHotTopicChatResponseBodyPayloadOutputArticles {
	s.Summary = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputArticles) SetTitle(v string) *RunHotTopicChatResponseBodyPayloadOutputArticles {
	s.Title = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputArticles) SetUrl(v string) *RunHotTopicChatResponseBodyPayloadOutputArticles {
	s.Url = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputArticles) Validate() error {
	return dara.Validate(s)
}

type RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries struct {
	// example:
	//
	// 100000
	CustomHotValue *float64 `json:"customHotValue,omitempty" xml:"customHotValue,omitempty"`
	// example:
	//
	// xxx
	CustomTextSummary *string `json:"customTextSummary,omitempty" xml:"customTextSummary,omitempty"`
	// example:
	//
	// xx
	HotTopic *string `json:"hotTopic,omitempty" xml:"hotTopic,omitempty"`
	// example:
	//
	// 2024-09-13_08
	HotTopicVersion *string `json:"hotTopicVersion,omitempty" xml:"hotTopicVersion,omitempty"`
	// example:
	//
	// 100000
	HotValue *float64                                                           `json:"hotValue,omitempty" xml:"hotValue,omitempty"`
	Images   []*RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesImages `json:"images,omitempty" xml:"images,omitempty" type:"Repeated"`
	News     []*RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesNews   `json:"news,omitempty" xml:"news,omitempty" type:"Repeated"`
	PubTime  *string                                                            `json:"pubTime,omitempty" xml:"pubTime,omitempty"`
	// example:
	//
	// xxx
	TextSummary *string `json:"textSummary,omitempty" xml:"textSummary,omitempty"`
	Url         *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) String() string {
	return dara.Prettify(s)
}

func (s RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) GoString() string {
	return s.String()
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) GetCustomHotValue() *float64 {
	return s.CustomHotValue
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) GetCustomTextSummary() *string {
	return s.CustomTextSummary
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) GetHotTopic() *string {
	return s.HotTopic
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) GetHotTopicVersion() *string {
	return s.HotTopicVersion
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) GetHotValue() *float64 {
	return s.HotValue
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) GetImages() []*RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesImages {
	return s.Images
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) GetNews() []*RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesNews {
	return s.News
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) GetPubTime() *string {
	return s.PubTime
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) GetTextSummary() *string {
	return s.TextSummary
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) GetUrl() *string {
	return s.Url
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) SetCustomHotValue(v float64) *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries {
	s.CustomHotValue = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) SetCustomTextSummary(v string) *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries {
	s.CustomTextSummary = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) SetHotTopic(v string) *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries {
	s.HotTopic = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) SetHotTopicVersion(v string) *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries {
	s.HotTopicVersion = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) SetHotValue(v float64) *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries {
	s.HotValue = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) SetImages(v []*RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesImages) *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries {
	s.Images = v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) SetNews(v []*RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesNews) *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries {
	s.News = v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) SetPubTime(v string) *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries {
	s.PubTime = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) SetTextSummary(v string) *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries {
	s.TextSummary = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) SetUrl(v string) *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries {
	s.Url = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummaries) Validate() error {
	return dara.Validate(s)
}

type RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesImages struct {
	// example:
	//
	// http://xxx
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesImages) String() string {
	return dara.Prettify(s)
}

func (s RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesImages) GoString() string {
	return s.String()
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesImages) GetUrl() *string {
	return s.Url
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesImages) SetUrl(v string) *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesImages {
	s.Url = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesImages) Validate() error {
	return dara.Validate(s)
}

type RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesNews struct {
	// example:
	//
	// xxx
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// http://xxx
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesNews) String() string {
	return dara.Prettify(s)
}

func (s RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesNews) GoString() string {
	return s.String()
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesNews) GetTitle() *string {
	return s.Title
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesNews) GetUrl() *string {
	return s.Url
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesNews) SetTitle(v string) *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesNews {
	s.Title = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesNews) SetUrl(v string) *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesNews {
	s.Url = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputHotTopicSummariesNews) Validate() error {
	return dara.Validate(s)
}

type RunHotTopicChatResponseBodyPayloadOutputMultimodalMedias struct {
	// example:
	//
	// http://xxxx
	FileUrl *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// example:
	//
	// image
	MediaType *string  `json:"mediaType,omitempty" xml:"mediaType,omitempty"`
	SortScore *float64 `json:"sortScore,omitempty" xml:"sortScore,omitempty"`
}

func (s RunHotTopicChatResponseBodyPayloadOutputMultimodalMedias) String() string {
	return dara.Prettify(s)
}

func (s RunHotTopicChatResponseBodyPayloadOutputMultimodalMedias) GoString() string {
	return s.String()
}

func (s *RunHotTopicChatResponseBodyPayloadOutputMultimodalMedias) GetFileUrl() *string {
	return s.FileUrl
}

func (s *RunHotTopicChatResponseBodyPayloadOutputMultimodalMedias) GetMediaType() *string {
	return s.MediaType
}

func (s *RunHotTopicChatResponseBodyPayloadOutputMultimodalMedias) GetSortScore() *float64 {
	return s.SortScore
}

func (s *RunHotTopicChatResponseBodyPayloadOutputMultimodalMedias) SetFileUrl(v string) *RunHotTopicChatResponseBodyPayloadOutputMultimodalMedias {
	s.FileUrl = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputMultimodalMedias) SetMediaType(v string) *RunHotTopicChatResponseBodyPayloadOutputMultimodalMedias {
	s.MediaType = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputMultimodalMedias) SetSortScore(v float64) *RunHotTopicChatResponseBodyPayloadOutputMultimodalMedias {
	s.SortScore = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadOutputMultimodalMedias) Validate() error {
	return dara.Validate(s)
}

type RunHotTopicChatResponseBodyPayloadUsage struct {
	// example:
	//
	// 100
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 100
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 200
	TotalTokens *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s RunHotTopicChatResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunHotTopicChatResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunHotTopicChatResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunHotTopicChatResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunHotTopicChatResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunHotTopicChatResponseBodyPayloadUsage) SetInputTokens(v int64) *RunHotTopicChatResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunHotTopicChatResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunHotTopicChatResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunHotTopicChatResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
