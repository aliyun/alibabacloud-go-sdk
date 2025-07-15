// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunWritingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnd(v bool) *RunWritingResponseBody
	GetEnd() *bool
	SetHeader(v *RunWritingResponseBodyHeader) *RunWritingResponseBody
	GetHeader() *RunWritingResponseBodyHeader
	SetPayload(v *RunWritingResponseBodyPayload) *RunWritingResponseBody
	GetPayload() *RunWritingResponseBodyPayload
	SetRequestId(v string) *RunWritingResponseBody
	GetRequestId() *string
}

type RunWritingResponseBody struct {
	End     *bool                          `json:"End,omitempty" xml:"End,omitempty"`
	Header  *RunWritingResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunWritingResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunWritingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunWritingResponseBody) GoString() string {
	return s.String()
}

func (s *RunWritingResponseBody) GetEnd() *bool {
	return s.End
}

func (s *RunWritingResponseBody) GetHeader() *RunWritingResponseBodyHeader {
	return s.Header
}

func (s *RunWritingResponseBody) GetPayload() *RunWritingResponseBodyPayload {
	return s.Payload
}

func (s *RunWritingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunWritingResponseBody) SetEnd(v bool) *RunWritingResponseBody {
	s.End = &v
	return s
}

func (s *RunWritingResponseBody) SetHeader(v *RunWritingResponseBodyHeader) *RunWritingResponseBody {
	s.Header = v
	return s
}

func (s *RunWritingResponseBody) SetPayload(v *RunWritingResponseBodyPayload) *RunWritingResponseBody {
	s.Payload = v
	return s
}

func (s *RunWritingResponseBody) SetRequestId(v string) *RunWritingResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunWritingResponseBody) Validate() error {
	return dara.Validate(s)
}

type RunWritingResponseBodyHeader struct {
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

func (s RunWritingResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunWritingResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunWritingResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunWritingResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunWritingResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunWritingResponseBodyHeader) GetOriginSessionId() *string {
	return s.OriginSessionId
}

func (s *RunWritingResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunWritingResponseBodyHeader) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunWritingResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunWritingResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunWritingResponseBodyHeader) SetErrorCode(v string) *RunWritingResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunWritingResponseBodyHeader) SetErrorMessage(v string) *RunWritingResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunWritingResponseBodyHeader) SetEvent(v string) *RunWritingResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunWritingResponseBodyHeader) SetOriginSessionId(v string) *RunWritingResponseBodyHeader {
	s.OriginSessionId = &v
	return s
}

func (s *RunWritingResponseBodyHeader) SetSessionId(v string) *RunWritingResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunWritingResponseBodyHeader) SetStatusCode(v int32) *RunWritingResponseBodyHeader {
	s.StatusCode = &v
	return s
}

func (s *RunWritingResponseBodyHeader) SetTaskId(v string) *RunWritingResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunWritingResponseBodyHeader) SetTraceId(v string) *RunWritingResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunWritingResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunWritingResponseBodyPayload struct {
	Output *RunWritingResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunWritingResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunWritingResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunWritingResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunWritingResponseBodyPayload) GetOutput() *RunWritingResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunWritingResponseBodyPayload) GetUsage() *RunWritingResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunWritingResponseBodyPayload) SetOutput(v *RunWritingResponseBodyPayloadOutput) *RunWritingResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunWritingResponseBodyPayload) SetUsage(v *RunWritingResponseBodyPayloadUsage) *RunWritingResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunWritingResponseBodyPayload) Validate() error {
	return dara.Validate(s)
}

type RunWritingResponseBodyPayloadOutput struct {
	Articles []*RunWritingResponseBodyPayloadOutputArticles `json:"Articles,omitempty" xml:"Articles,omitempty" type:"Repeated"`
	// example:
	//
	// 文章精排之后的片段
	MiniDoc []*string `json:"MiniDoc,omitempty" xml:"MiniDoc,omitempty" type:"Repeated"`
	// example:
	//
	// 大模型改变世界
	SearchQuery *string `json:"SearchQuery,omitempty" xml:"SearchQuery,omitempty"`
	// example:
	//
	// 文本生成结果
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RunWritingResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunWritingResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunWritingResponseBodyPayloadOutput) GetArticles() []*RunWritingResponseBodyPayloadOutputArticles {
	return s.Articles
}

func (s *RunWritingResponseBodyPayloadOutput) GetMiniDoc() []*string {
	return s.MiniDoc
}

func (s *RunWritingResponseBodyPayloadOutput) GetSearchQuery() *string {
	return s.SearchQuery
}

func (s *RunWritingResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunWritingResponseBodyPayloadOutput) SetArticles(v []*RunWritingResponseBodyPayloadOutputArticles) *RunWritingResponseBodyPayloadOutput {
	s.Articles = v
	return s
}

func (s *RunWritingResponseBodyPayloadOutput) SetMiniDoc(v []*string) *RunWritingResponseBodyPayloadOutput {
	s.MiniDoc = v
	return s
}

func (s *RunWritingResponseBodyPayloadOutput) SetSearchQuery(v string) *RunWritingResponseBodyPayloadOutput {
	s.SearchQuery = &v
	return s
}

func (s *RunWritingResponseBodyPayloadOutput) SetText(v string) *RunWritingResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunWritingResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunWritingResponseBodyPayloadOutputArticles struct {
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

func (s RunWritingResponseBodyPayloadOutputArticles) String() string {
	return dara.Prettify(s)
}

func (s RunWritingResponseBodyPayloadOutputArticles) GoString() string {
	return s.String()
}

func (s *RunWritingResponseBodyPayloadOutputArticles) GetAuthor() *string {
	return s.Author
}

func (s *RunWritingResponseBodyPayloadOutputArticles) GetContent() *string {
	return s.Content
}

func (s *RunWritingResponseBodyPayloadOutputArticles) GetDocId() *string {
	return s.DocId
}

func (s *RunWritingResponseBodyPayloadOutputArticles) GetDocUuid() *string {
	return s.DocUuid
}

func (s *RunWritingResponseBodyPayloadOutputArticles) GetPubTime() *string {
	return s.PubTime
}

func (s *RunWritingResponseBodyPayloadOutputArticles) GetSource() *string {
	return s.Source
}

func (s *RunWritingResponseBodyPayloadOutputArticles) GetSummary() *string {
	return s.Summary
}

func (s *RunWritingResponseBodyPayloadOutputArticles) GetTag() *string {
	return s.Tag
}

func (s *RunWritingResponseBodyPayloadOutputArticles) GetTitle() *string {
	return s.Title
}

func (s *RunWritingResponseBodyPayloadOutputArticles) GetUrl() *string {
	return s.Url
}

func (s *RunWritingResponseBodyPayloadOutputArticles) SetAuthor(v string) *RunWritingResponseBodyPayloadOutputArticles {
	s.Author = &v
	return s
}

func (s *RunWritingResponseBodyPayloadOutputArticles) SetContent(v string) *RunWritingResponseBodyPayloadOutputArticles {
	s.Content = &v
	return s
}

func (s *RunWritingResponseBodyPayloadOutputArticles) SetDocId(v string) *RunWritingResponseBodyPayloadOutputArticles {
	s.DocId = &v
	return s
}

func (s *RunWritingResponseBodyPayloadOutputArticles) SetDocUuid(v string) *RunWritingResponseBodyPayloadOutputArticles {
	s.DocUuid = &v
	return s
}

func (s *RunWritingResponseBodyPayloadOutputArticles) SetPubTime(v string) *RunWritingResponseBodyPayloadOutputArticles {
	s.PubTime = &v
	return s
}

func (s *RunWritingResponseBodyPayloadOutputArticles) SetSource(v string) *RunWritingResponseBodyPayloadOutputArticles {
	s.Source = &v
	return s
}

func (s *RunWritingResponseBodyPayloadOutputArticles) SetSummary(v string) *RunWritingResponseBodyPayloadOutputArticles {
	s.Summary = &v
	return s
}

func (s *RunWritingResponseBodyPayloadOutputArticles) SetTag(v string) *RunWritingResponseBodyPayloadOutputArticles {
	s.Tag = &v
	return s
}

func (s *RunWritingResponseBodyPayloadOutputArticles) SetTitle(v string) *RunWritingResponseBodyPayloadOutputArticles {
	s.Title = &v
	return s
}

func (s *RunWritingResponseBodyPayloadOutputArticles) SetUrl(v string) *RunWritingResponseBodyPayloadOutputArticles {
	s.Url = &v
	return s
}

func (s *RunWritingResponseBodyPayloadOutputArticles) Validate() error {
	return dara.Validate(s)
}

type RunWritingResponseBodyPayloadUsage struct {
	// example:
	//
	// 1
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 1
	OutputTokens *int64            `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	TokenMap     map[string]*int64 `json:"TokenMap,omitempty" xml:"TokenMap,omitempty"`
	// example:
	//
	// 2
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunWritingResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunWritingResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunWritingResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunWritingResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunWritingResponseBodyPayloadUsage) GetTokenMap() map[string]*int64 {
	return s.TokenMap
}

func (s *RunWritingResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunWritingResponseBodyPayloadUsage) SetInputTokens(v int64) *RunWritingResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunWritingResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunWritingResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunWritingResponseBodyPayloadUsage) SetTokenMap(v map[string]*int64) *RunWritingResponseBodyPayloadUsage {
	s.TokenMap = v
	return s
}

func (s *RunWritingResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunWritingResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunWritingResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
