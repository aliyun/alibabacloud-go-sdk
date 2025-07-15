// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSearchSimilarArticlesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeader(v *RunSearchSimilarArticlesResponseBodyHeader) *RunSearchSimilarArticlesResponseBody
	GetHeader() *RunSearchSimilarArticlesResponseBodyHeader
	SetPayload(v *RunSearchSimilarArticlesResponseBodyPayload) *RunSearchSimilarArticlesResponseBody
	GetPayload() *RunSearchSimilarArticlesResponseBodyPayload
	SetRequestId(v string) *RunSearchSimilarArticlesResponseBody
	GetRequestId() *string
}

type RunSearchSimilarArticlesResponseBody struct {
	Header  *RunSearchSimilarArticlesResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunSearchSimilarArticlesResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunSearchSimilarArticlesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunSearchSimilarArticlesResponseBody) GoString() string {
	return s.String()
}

func (s *RunSearchSimilarArticlesResponseBody) GetHeader() *RunSearchSimilarArticlesResponseBodyHeader {
	return s.Header
}

func (s *RunSearchSimilarArticlesResponseBody) GetPayload() *RunSearchSimilarArticlesResponseBodyPayload {
	return s.Payload
}

func (s *RunSearchSimilarArticlesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunSearchSimilarArticlesResponseBody) SetHeader(v *RunSearchSimilarArticlesResponseBodyHeader) *RunSearchSimilarArticlesResponseBody {
	s.Header = v
	return s
}

func (s *RunSearchSimilarArticlesResponseBody) SetPayload(v *RunSearchSimilarArticlesResponseBodyPayload) *RunSearchSimilarArticlesResponseBody {
	s.Payload = v
	return s
}

func (s *RunSearchSimilarArticlesResponseBody) SetRequestId(v string) *RunSearchSimilarArticlesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunSearchSimilarArticlesResponseBody) Validate() error {
	return dara.Validate(s)
}

type RunSearchSimilarArticlesResponseBodyHeader struct {
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
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RunSearchSimilarArticlesResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunSearchSimilarArticlesResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunSearchSimilarArticlesResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunSearchSimilarArticlesResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunSearchSimilarArticlesResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunSearchSimilarArticlesResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunSearchSimilarArticlesResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunSearchSimilarArticlesResponseBodyHeader) SetErrorCode(v string) *RunSearchSimilarArticlesResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunSearchSimilarArticlesResponseBodyHeader) SetErrorMessage(v string) *RunSearchSimilarArticlesResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunSearchSimilarArticlesResponseBodyHeader) SetEvent(v string) *RunSearchSimilarArticlesResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunSearchSimilarArticlesResponseBodyHeader) SetSessionId(v string) *RunSearchSimilarArticlesResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunSearchSimilarArticlesResponseBodyHeader) SetTaskId(v string) *RunSearchSimilarArticlesResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunSearchSimilarArticlesResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunSearchSimilarArticlesResponseBodyPayload struct {
	Output *RunSearchSimilarArticlesResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunSearchSimilarArticlesResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunSearchSimilarArticlesResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunSearchSimilarArticlesResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunSearchSimilarArticlesResponseBodyPayload) GetOutput() *RunSearchSimilarArticlesResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunSearchSimilarArticlesResponseBodyPayload) GetUsage() *RunSearchSimilarArticlesResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunSearchSimilarArticlesResponseBodyPayload) SetOutput(v *RunSearchSimilarArticlesResponseBodyPayloadOutput) *RunSearchSimilarArticlesResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunSearchSimilarArticlesResponseBodyPayload) SetUsage(v *RunSearchSimilarArticlesResponseBodyPayloadUsage) *RunSearchSimilarArticlesResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunSearchSimilarArticlesResponseBodyPayload) Validate() error {
	return dara.Validate(s)
}

type RunSearchSimilarArticlesResponseBodyPayloadOutput struct {
	Articles []*RunSearchSimilarArticlesResponseBodyPayloadOutputArticles `json:"Articles,omitempty" xml:"Articles,omitempty" type:"Repeated"`
	// example:
	//
	// 文本生成结果
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RunSearchSimilarArticlesResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunSearchSimilarArticlesResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunSearchSimilarArticlesResponseBodyPayloadOutput) GetArticles() []*RunSearchSimilarArticlesResponseBodyPayloadOutputArticles {
	return s.Articles
}

func (s *RunSearchSimilarArticlesResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunSearchSimilarArticlesResponseBodyPayloadOutput) SetArticles(v []*RunSearchSimilarArticlesResponseBodyPayloadOutputArticles) *RunSearchSimilarArticlesResponseBodyPayloadOutput {
	s.Articles = v
	return s
}

func (s *RunSearchSimilarArticlesResponseBodyPayloadOutput) SetText(v string) *RunSearchSimilarArticlesResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunSearchSimilarArticlesResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunSearchSimilarArticlesResponseBodyPayloadOutputArticles struct {
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// a26c2c1
	DocUuid *string `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	// example:
	//
	// 2025-01-16 18:07:22
	PubTime          *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	SearchSourceName *string `json:"SearchSourceName,omitempty" xml:"SearchSourceName,omitempty"`
	// example:
	//
	// xxx.com
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// xxx
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	Title   *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// https://xxx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RunSearchSimilarArticlesResponseBodyPayloadOutputArticles) String() string {
	return dara.Prettify(s)
}

func (s RunSearchSimilarArticlesResponseBodyPayloadOutputArticles) GoString() string {
	return s.String()
}

func (s *RunSearchSimilarArticlesResponseBodyPayloadOutputArticles) GetDocId() *string {
	return s.DocId
}

func (s *RunSearchSimilarArticlesResponseBodyPayloadOutputArticles) GetDocUuid() *string {
	return s.DocUuid
}

func (s *RunSearchSimilarArticlesResponseBodyPayloadOutputArticles) GetPubTime() *string {
	return s.PubTime
}

func (s *RunSearchSimilarArticlesResponseBodyPayloadOutputArticles) GetSearchSourceName() *string {
	return s.SearchSourceName
}

func (s *RunSearchSimilarArticlesResponseBodyPayloadOutputArticles) GetSource() *string {
	return s.Source
}

func (s *RunSearchSimilarArticlesResponseBodyPayloadOutputArticles) GetSummary() *string {
	return s.Summary
}

func (s *RunSearchSimilarArticlesResponseBodyPayloadOutputArticles) GetTitle() *string {
	return s.Title
}

func (s *RunSearchSimilarArticlesResponseBodyPayloadOutputArticles) GetUrl() *string {
	return s.Url
}

func (s *RunSearchSimilarArticlesResponseBodyPayloadOutputArticles) SetDocId(v string) *RunSearchSimilarArticlesResponseBodyPayloadOutputArticles {
	s.DocId = &v
	return s
}

func (s *RunSearchSimilarArticlesResponseBodyPayloadOutputArticles) SetDocUuid(v string) *RunSearchSimilarArticlesResponseBodyPayloadOutputArticles {
	s.DocUuid = &v
	return s
}

func (s *RunSearchSimilarArticlesResponseBodyPayloadOutputArticles) SetPubTime(v string) *RunSearchSimilarArticlesResponseBodyPayloadOutputArticles {
	s.PubTime = &v
	return s
}

func (s *RunSearchSimilarArticlesResponseBodyPayloadOutputArticles) SetSearchSourceName(v string) *RunSearchSimilarArticlesResponseBodyPayloadOutputArticles {
	s.SearchSourceName = &v
	return s
}

func (s *RunSearchSimilarArticlesResponseBodyPayloadOutputArticles) SetSource(v string) *RunSearchSimilarArticlesResponseBodyPayloadOutputArticles {
	s.Source = &v
	return s
}

func (s *RunSearchSimilarArticlesResponseBodyPayloadOutputArticles) SetSummary(v string) *RunSearchSimilarArticlesResponseBodyPayloadOutputArticles {
	s.Summary = &v
	return s
}

func (s *RunSearchSimilarArticlesResponseBodyPayloadOutputArticles) SetTitle(v string) *RunSearchSimilarArticlesResponseBodyPayloadOutputArticles {
	s.Title = &v
	return s
}

func (s *RunSearchSimilarArticlesResponseBodyPayloadOutputArticles) SetUrl(v string) *RunSearchSimilarArticlesResponseBodyPayloadOutputArticles {
	s.Url = &v
	return s
}

func (s *RunSearchSimilarArticlesResponseBodyPayloadOutputArticles) Validate() error {
	return dara.Validate(s)
}

type RunSearchSimilarArticlesResponseBodyPayloadUsage struct {
	// example:
	//
	// 81
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 9
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// example:
	//
	// 50
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunSearchSimilarArticlesResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunSearchSimilarArticlesResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunSearchSimilarArticlesResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunSearchSimilarArticlesResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunSearchSimilarArticlesResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunSearchSimilarArticlesResponseBodyPayloadUsage) SetInputTokens(v int64) *RunSearchSimilarArticlesResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunSearchSimilarArticlesResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunSearchSimilarArticlesResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunSearchSimilarArticlesResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunSearchSimilarArticlesResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunSearchSimilarArticlesResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
