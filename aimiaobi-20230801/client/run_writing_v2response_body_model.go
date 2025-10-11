// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunWritingV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnd(v bool) *RunWritingV2ResponseBody
	GetEnd() *bool
	SetHeader(v *RunWritingV2ResponseBodyHeader) *RunWritingV2ResponseBody
	GetHeader() *RunWritingV2ResponseBodyHeader
	SetPayload(v *RunWritingV2ResponseBodyPayload) *RunWritingV2ResponseBody
	GetPayload() *RunWritingV2ResponseBodyPayload
	SetRequestId(v string) *RunWritingV2ResponseBody
	GetRequestId() *string
}

type RunWritingV2ResponseBody struct {
	// example:
	//
	// true
	End     *bool                            `json:"End,omitempty" xml:"End,omitempty"`
	Header  *RunWritingV2ResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunWritingV2ResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunWritingV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunWritingV2ResponseBody) GoString() string {
	return s.String()
}

func (s *RunWritingV2ResponseBody) GetEnd() *bool {
	return s.End
}

func (s *RunWritingV2ResponseBody) GetHeader() *RunWritingV2ResponseBodyHeader {
	return s.Header
}

func (s *RunWritingV2ResponseBody) GetPayload() *RunWritingV2ResponseBodyPayload {
	return s.Payload
}

func (s *RunWritingV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunWritingV2ResponseBody) SetEnd(v bool) *RunWritingV2ResponseBody {
	s.End = &v
	return s
}

func (s *RunWritingV2ResponseBody) SetHeader(v *RunWritingV2ResponseBodyHeader) *RunWritingV2ResponseBody {
	s.Header = v
	return s
}

func (s *RunWritingV2ResponseBody) SetPayload(v *RunWritingV2ResponseBodyPayload) *RunWritingV2ResponseBody {
	s.Payload = v
	return s
}

func (s *RunWritingV2ResponseBody) SetRequestId(v string) *RunWritingV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunWritingV2ResponseBody) Validate() error {
	return dara.Validate(s)
}

type RunWritingV2ResponseBodyHeader struct {
	// example:
	//
	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 错误信息
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// task-started
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	OriginSessionId *string `json:"OriginSessionId,omitempty" xml:"OriginSessionId,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 400
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 全链路ID
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunWritingV2ResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunWritingV2ResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunWritingV2ResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunWritingV2ResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunWritingV2ResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunWritingV2ResponseBodyHeader) GetOriginSessionId() *string {
	return s.OriginSessionId
}

func (s *RunWritingV2ResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunWritingV2ResponseBodyHeader) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunWritingV2ResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunWritingV2ResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunWritingV2ResponseBodyHeader) SetErrorCode(v string) *RunWritingV2ResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunWritingV2ResponseBodyHeader) SetErrorMessage(v string) *RunWritingV2ResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunWritingV2ResponseBodyHeader) SetEvent(v string) *RunWritingV2ResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunWritingV2ResponseBodyHeader) SetOriginSessionId(v string) *RunWritingV2ResponseBodyHeader {
	s.OriginSessionId = &v
	return s
}

func (s *RunWritingV2ResponseBodyHeader) SetSessionId(v string) *RunWritingV2ResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunWritingV2ResponseBodyHeader) SetStatusCode(v int32) *RunWritingV2ResponseBodyHeader {
	s.StatusCode = &v
	return s
}

func (s *RunWritingV2ResponseBodyHeader) SetTaskId(v string) *RunWritingV2ResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunWritingV2ResponseBodyHeader) SetTraceId(v string) *RunWritingV2ResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunWritingV2ResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunWritingV2ResponseBodyPayload struct {
	Output *RunWritingV2ResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunWritingV2ResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunWritingV2ResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunWritingV2ResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunWritingV2ResponseBodyPayload) GetOutput() *RunWritingV2ResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunWritingV2ResponseBodyPayload) GetUsage() *RunWritingV2ResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunWritingV2ResponseBodyPayload) SetOutput(v *RunWritingV2ResponseBodyPayloadOutput) *RunWritingV2ResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunWritingV2ResponseBodyPayload) SetUsage(v *RunWritingV2ResponseBodyPayloadUsage) *RunWritingV2ResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunWritingV2ResponseBodyPayload) Validate() error {
	return dara.Validate(s)
}

type RunWritingV2ResponseBodyPayloadOutput struct {
	Articles             []*RunWritingV2ResponseBodyPayloadOutputArticles `json:"Articles,omitempty" xml:"Articles,omitempty" type:"Repeated"`
	GenerateTraceability *GenerateTraceability                            `json:"GenerateTraceability,omitempty" xml:"GenerateTraceability,omitempty"`
	// example:
	//
	// 文章精排之后的片段
	MiniDoc  []*string         `json:"MiniDoc,omitempty" xml:"MiniDoc,omitempty" type:"Repeated"`
	Outlines []*WritingOutline `json:"Outlines,omitempty" xml:"Outlines,omitempty" type:"Repeated"`
	// example:
	//
	// 大模型改变世界
	SearchQuery  *string              `json:"SearchQuery,omitempty" xml:"SearchQuery,omitempty"`
	SearchResult *OutlineSearchResult `json:"SearchResult,omitempty" xml:"SearchResult,omitempty"`
	// example:
	//
	// 文本生成结果
	Text  *string `json:"Text,omitempty" xml:"Text,omitempty"`
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s RunWritingV2ResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunWritingV2ResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunWritingV2ResponseBodyPayloadOutput) GetArticles() []*RunWritingV2ResponseBodyPayloadOutputArticles {
	return s.Articles
}

func (s *RunWritingV2ResponseBodyPayloadOutput) GetGenerateTraceability() *GenerateTraceability {
	return s.GenerateTraceability
}

func (s *RunWritingV2ResponseBodyPayloadOutput) GetMiniDoc() []*string {
	return s.MiniDoc
}

func (s *RunWritingV2ResponseBodyPayloadOutput) GetOutlines() []*WritingOutline {
	return s.Outlines
}

func (s *RunWritingV2ResponseBodyPayloadOutput) GetSearchQuery() *string {
	return s.SearchQuery
}

func (s *RunWritingV2ResponseBodyPayloadOutput) GetSearchResult() *OutlineSearchResult {
	return s.SearchResult
}

func (s *RunWritingV2ResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunWritingV2ResponseBodyPayloadOutput) GetTitle() *string {
	return s.Title
}

func (s *RunWritingV2ResponseBodyPayloadOutput) SetArticles(v []*RunWritingV2ResponseBodyPayloadOutputArticles) *RunWritingV2ResponseBodyPayloadOutput {
	s.Articles = v
	return s
}

func (s *RunWritingV2ResponseBodyPayloadOutput) SetGenerateTraceability(v *GenerateTraceability) *RunWritingV2ResponseBodyPayloadOutput {
	s.GenerateTraceability = v
	return s
}

func (s *RunWritingV2ResponseBodyPayloadOutput) SetMiniDoc(v []*string) *RunWritingV2ResponseBodyPayloadOutput {
	s.MiniDoc = v
	return s
}

func (s *RunWritingV2ResponseBodyPayloadOutput) SetOutlines(v []*WritingOutline) *RunWritingV2ResponseBodyPayloadOutput {
	s.Outlines = v
	return s
}

func (s *RunWritingV2ResponseBodyPayloadOutput) SetSearchQuery(v string) *RunWritingV2ResponseBodyPayloadOutput {
	s.SearchQuery = &v
	return s
}

func (s *RunWritingV2ResponseBodyPayloadOutput) SetSearchResult(v *OutlineSearchResult) *RunWritingV2ResponseBodyPayloadOutput {
	s.SearchResult = v
	return s
}

func (s *RunWritingV2ResponseBodyPayloadOutput) SetText(v string) *RunWritingV2ResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunWritingV2ResponseBodyPayloadOutput) SetTitle(v string) *RunWritingV2ResponseBodyPayloadOutput {
	s.Title = &v
	return s
}

func (s *RunWritingV2ResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunWritingV2ResponseBodyPayloadOutputArticles struct {
	// example:
	//
	// 作者
	Author *string `json:"Author,omitempty" xml:"Author,omitempty"`
	// example:
	//
	// 文章内容
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 文档-自定义的唯一ID
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// 98229f6001cf4deeb1668191d4eccc75
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// 2024-08-28 11:38:28
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	// example:
	//
	// 央视网
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 文章摘要
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// 文章标签
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// example:
	//
	// 文章标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// https://www.example.com/aaa.docx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RunWritingV2ResponseBodyPayloadOutputArticles) String() string {
	return dara.Prettify(s)
}

func (s RunWritingV2ResponseBodyPayloadOutputArticles) GoString() string {
	return s.String()
}

func (s *RunWritingV2ResponseBodyPayloadOutputArticles) GetAuthor() *string {
	return s.Author
}

func (s *RunWritingV2ResponseBodyPayloadOutputArticles) GetContent() *string {
	return s.Content
}

func (s *RunWritingV2ResponseBodyPayloadOutputArticles) GetDocId() *string {
	return s.DocId
}

func (s *RunWritingV2ResponseBodyPayloadOutputArticles) GetDocUuid() *string {
	return s.DocUuid
}

func (s *RunWritingV2ResponseBodyPayloadOutputArticles) GetPubTime() *string {
	return s.PubTime
}

func (s *RunWritingV2ResponseBodyPayloadOutputArticles) GetSource() *string {
	return s.Source
}

func (s *RunWritingV2ResponseBodyPayloadOutputArticles) GetSummary() *string {
	return s.Summary
}

func (s *RunWritingV2ResponseBodyPayloadOutputArticles) GetTag() *string {
	return s.Tag
}

func (s *RunWritingV2ResponseBodyPayloadOutputArticles) GetTitle() *string {
	return s.Title
}

func (s *RunWritingV2ResponseBodyPayloadOutputArticles) GetUrl() *string {
	return s.Url
}

func (s *RunWritingV2ResponseBodyPayloadOutputArticles) SetAuthor(v string) *RunWritingV2ResponseBodyPayloadOutputArticles {
	s.Author = &v
	return s
}

func (s *RunWritingV2ResponseBodyPayloadOutputArticles) SetContent(v string) *RunWritingV2ResponseBodyPayloadOutputArticles {
	s.Content = &v
	return s
}

func (s *RunWritingV2ResponseBodyPayloadOutputArticles) SetDocId(v string) *RunWritingV2ResponseBodyPayloadOutputArticles {
	s.DocId = &v
	return s
}

func (s *RunWritingV2ResponseBodyPayloadOutputArticles) SetDocUuid(v string) *RunWritingV2ResponseBodyPayloadOutputArticles {
	s.DocUuid = &v
	return s
}

func (s *RunWritingV2ResponseBodyPayloadOutputArticles) SetPubTime(v string) *RunWritingV2ResponseBodyPayloadOutputArticles {
	s.PubTime = &v
	return s
}

func (s *RunWritingV2ResponseBodyPayloadOutputArticles) SetSource(v string) *RunWritingV2ResponseBodyPayloadOutputArticles {
	s.Source = &v
	return s
}

func (s *RunWritingV2ResponseBodyPayloadOutputArticles) SetSummary(v string) *RunWritingV2ResponseBodyPayloadOutputArticles {
	s.Summary = &v
	return s
}

func (s *RunWritingV2ResponseBodyPayloadOutputArticles) SetTag(v string) *RunWritingV2ResponseBodyPayloadOutputArticles {
	s.Tag = &v
	return s
}

func (s *RunWritingV2ResponseBodyPayloadOutputArticles) SetTitle(v string) *RunWritingV2ResponseBodyPayloadOutputArticles {
	s.Title = &v
	return s
}

func (s *RunWritingV2ResponseBodyPayloadOutputArticles) SetUrl(v string) *RunWritingV2ResponseBodyPayloadOutputArticles {
	s.Url = &v
	return s
}

func (s *RunWritingV2ResponseBodyPayloadOutputArticles) Validate() error {
	return dara.Validate(s)
}

type RunWritingV2ResponseBodyPayloadUsage struct {
	// example:
	//
	// 78
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 34
	OutputTokens *int64            `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	TokenMap     map[string]*int64 `json:"TokenMap,omitempty" xml:"TokenMap,omitempty"`
	// example:
	//
	// 38
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunWritingV2ResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunWritingV2ResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunWritingV2ResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunWritingV2ResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunWritingV2ResponseBodyPayloadUsage) GetTokenMap() map[string]*int64 {
	return s.TokenMap
}

func (s *RunWritingV2ResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunWritingV2ResponseBodyPayloadUsage) SetInputTokens(v int64) *RunWritingV2ResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunWritingV2ResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunWritingV2ResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunWritingV2ResponseBodyPayloadUsage) SetTokenMap(v map[string]*int64) *RunWritingV2ResponseBodyPayloadUsage {
	s.TokenMap = v
	return s
}

func (s *RunWritingV2ResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunWritingV2ResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunWritingV2ResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
