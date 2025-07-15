// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCustomHotTopicAnalysisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeader(v *RunCustomHotTopicAnalysisResponseBodyHeader) *RunCustomHotTopicAnalysisResponseBody
	GetHeader() *RunCustomHotTopicAnalysisResponseBodyHeader
	SetPayload(v *RunCustomHotTopicAnalysisResponseBodyPayload) *RunCustomHotTopicAnalysisResponseBody
	GetPayload() *RunCustomHotTopicAnalysisResponseBodyPayload
	SetRequestId(v string) *RunCustomHotTopicAnalysisResponseBody
	GetRequestId() *string
}

type RunCustomHotTopicAnalysisResponseBody struct {
	Header  *RunCustomHotTopicAnalysisResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunCustomHotTopicAnalysisResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunCustomHotTopicAnalysisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunCustomHotTopicAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *RunCustomHotTopicAnalysisResponseBody) GetHeader() *RunCustomHotTopicAnalysisResponseBodyHeader {
	return s.Header
}

func (s *RunCustomHotTopicAnalysisResponseBody) GetPayload() *RunCustomHotTopicAnalysisResponseBodyPayload {
	return s.Payload
}

func (s *RunCustomHotTopicAnalysisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunCustomHotTopicAnalysisResponseBody) SetHeader(v *RunCustomHotTopicAnalysisResponseBodyHeader) *RunCustomHotTopicAnalysisResponseBody {
	s.Header = v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBody) SetPayload(v *RunCustomHotTopicAnalysisResponseBodyPayload) *RunCustomHotTopicAnalysisResponseBody {
	s.Payload = v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBody) SetRequestId(v string) *RunCustomHotTopicAnalysisResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBody) Validate() error {
	return dara.Validate(s)
}

type RunCustomHotTopicAnalysisResponseBodyHeader struct {
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
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 全链路ID
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunCustomHotTopicAnalysisResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunCustomHotTopicAnalysisResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunCustomHotTopicAnalysisResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunCustomHotTopicAnalysisResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunCustomHotTopicAnalysisResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunCustomHotTopicAnalysisResponseBodyHeader) GetOriginSessionId() *string {
	return s.OriginSessionId
}

func (s *RunCustomHotTopicAnalysisResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunCustomHotTopicAnalysisResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunCustomHotTopicAnalysisResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunCustomHotTopicAnalysisResponseBodyHeader) SetErrorCode(v string) *RunCustomHotTopicAnalysisResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyHeader) SetErrorMessage(v string) *RunCustomHotTopicAnalysisResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyHeader) SetEvent(v string) *RunCustomHotTopicAnalysisResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyHeader) SetOriginSessionId(v string) *RunCustomHotTopicAnalysisResponseBodyHeader {
	s.OriginSessionId = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyHeader) SetSessionId(v string) *RunCustomHotTopicAnalysisResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyHeader) SetTaskId(v string) *RunCustomHotTopicAnalysisResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyHeader) SetTraceId(v string) *RunCustomHotTopicAnalysisResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunCustomHotTopicAnalysisResponseBodyPayload struct {
	Output *RunCustomHotTopicAnalysisResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunCustomHotTopicAnalysisResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunCustomHotTopicAnalysisResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunCustomHotTopicAnalysisResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayload) GetOutput() *RunCustomHotTopicAnalysisResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayload) GetUsage() *RunCustomHotTopicAnalysisResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayload) SetOutput(v *RunCustomHotTopicAnalysisResponseBodyPayloadOutput) *RunCustomHotTopicAnalysisResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayload) SetUsage(v *RunCustomHotTopicAnalysisResponseBodyPayloadUsage) *RunCustomHotTopicAnalysisResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayload) Validate() error {
	return dara.Validate(s)
}

type RunCustomHotTopicAnalysisResponseBodyPayloadOutput struct {
	Articles []*RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles `json:"Articles,omitempty" xml:"Articles,omitempty" type:"Repeated"`
	AskUser  []*string                                                     `json:"AskUser,omitempty" xml:"AskUser,omitempty" type:"Repeated"`
	// example:
	//
	// 异步任务ID
	AsyncTaskId *string `json:"AsyncTaskId,omitempty" xml:"AsyncTaskId,omitempty"`
	// example:
	//
	// 自定义选题视角
	Attitude *string `json:"Attitude,omitempty" xml:"Attitude,omitempty"`
	// example:
	//
	// 大模型改变世界
	SearchQuery *string `json:"SearchQuery,omitempty" xml:"SearchQuery,omitempty"`
	// example:
	//
	// 文本生成结果
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// 话题ID
	TopicId *string `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
}

func (s RunCustomHotTopicAnalysisResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunCustomHotTopicAnalysisResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutput) GetArticles() []*RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles {
	return s.Articles
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutput) GetAskUser() []*string {
	return s.AskUser
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutput) GetAsyncTaskId() *string {
	return s.AsyncTaskId
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutput) GetAttitude() *string {
	return s.Attitude
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutput) GetSearchQuery() *string {
	return s.SearchQuery
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutput) GetTopicId() *string {
	return s.TopicId
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutput) SetArticles(v []*RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) *RunCustomHotTopicAnalysisResponseBodyPayloadOutput {
	s.Articles = v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutput) SetAskUser(v []*string) *RunCustomHotTopicAnalysisResponseBodyPayloadOutput {
	s.AskUser = v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutput) SetAsyncTaskId(v string) *RunCustomHotTopicAnalysisResponseBodyPayloadOutput {
	s.AsyncTaskId = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutput) SetAttitude(v string) *RunCustomHotTopicAnalysisResponseBodyPayloadOutput {
	s.Attitude = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutput) SetSearchQuery(v string) *RunCustomHotTopicAnalysisResponseBodyPayloadOutput {
	s.SearchQuery = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutput) SetText(v string) *RunCustomHotTopicAnalysisResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutput) SetTopicId(v string) *RunCustomHotTopicAnalysisResponseBodyPayloadOutput {
	s.TopicId = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles struct {
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
	// a2103fcfbd5441f1991c72f8834833e3
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// 2024-08-27 14:50:47
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

func (s RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) String() string {
	return dara.Prettify(s)
}

func (s RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) GoString() string {
	return s.String()
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) GetAuthor() *string {
	return s.Author
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) GetContent() *string {
	return s.Content
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) GetDocId() *string {
	return s.DocId
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) GetDocUuid() *string {
	return s.DocUuid
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) GetPubTime() *string {
	return s.PubTime
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) GetSource() *string {
	return s.Source
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) GetSummary() *string {
	return s.Summary
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) GetTag() *string {
	return s.Tag
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) GetTitle() *string {
	return s.Title
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) GetUrl() *string {
	return s.Url
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) SetAuthor(v string) *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles {
	s.Author = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) SetContent(v string) *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles {
	s.Content = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) SetDocId(v string) *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles {
	s.DocId = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) SetDocUuid(v string) *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles {
	s.DocUuid = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) SetPubTime(v string) *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles {
	s.PubTime = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) SetSource(v string) *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles {
	s.Source = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) SetSummary(v string) *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles {
	s.Summary = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) SetTag(v string) *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles {
	s.Tag = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) SetTitle(v string) *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles {
	s.Title = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) SetUrl(v string) *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles {
	s.Url = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadOutputArticles) Validate() error {
	return dara.Validate(s)
}

type RunCustomHotTopicAnalysisResponseBodyPayloadUsage struct {
	// example:
	//
	// 60
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 13
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// example:
	//
	// 73
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunCustomHotTopicAnalysisResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunCustomHotTopicAnalysisResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadUsage) SetInputTokens(v int64) *RunCustomHotTopicAnalysisResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunCustomHotTopicAnalysisResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunCustomHotTopicAnalysisResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunCustomHotTopicAnalysisResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
