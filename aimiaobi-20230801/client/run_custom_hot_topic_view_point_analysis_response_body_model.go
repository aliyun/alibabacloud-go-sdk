// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCustomHotTopicViewPointAnalysisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeader(v *RunCustomHotTopicViewPointAnalysisResponseBodyHeader) *RunCustomHotTopicViewPointAnalysisResponseBody
	GetHeader() *RunCustomHotTopicViewPointAnalysisResponseBodyHeader
	SetPayload(v *RunCustomHotTopicViewPointAnalysisResponseBodyPayload) *RunCustomHotTopicViewPointAnalysisResponseBody
	GetPayload() *RunCustomHotTopicViewPointAnalysisResponseBodyPayload
	SetRequestId(v string) *RunCustomHotTopicViewPointAnalysisResponseBody
	GetRequestId() *string
}

type RunCustomHotTopicViewPointAnalysisResponseBody struct {
	Header  *RunCustomHotTopicViewPointAnalysisResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunCustomHotTopicViewPointAnalysisResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunCustomHotTopicViewPointAnalysisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunCustomHotTopicViewPointAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBody) GetHeader() *RunCustomHotTopicViewPointAnalysisResponseBodyHeader {
	return s.Header
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBody) GetPayload() *RunCustomHotTopicViewPointAnalysisResponseBodyPayload {
	return s.Payload
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBody) SetHeader(v *RunCustomHotTopicViewPointAnalysisResponseBodyHeader) *RunCustomHotTopicViewPointAnalysisResponseBody {
	s.Header = v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBody) SetPayload(v *RunCustomHotTopicViewPointAnalysisResponseBodyPayload) *RunCustomHotTopicViewPointAnalysisResponseBody {
	s.Payload = v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBody) SetRequestId(v string) *RunCustomHotTopicViewPointAnalysisResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBody) Validate() error {
	if s.Header != nil {
		if err := s.Header.Validate(); err != nil {
			return err
		}
	}
	if s.Payload != nil {
		if err := s.Payload.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunCustomHotTopicViewPointAnalysisResponseBodyHeader struct {
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

func (s RunCustomHotTopicViewPointAnalysisResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunCustomHotTopicViewPointAnalysisResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyHeader) GetOriginSessionId() *string {
	return s.OriginSessionId
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyHeader) SetErrorCode(v string) *RunCustomHotTopicViewPointAnalysisResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyHeader) SetErrorMessage(v string) *RunCustomHotTopicViewPointAnalysisResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyHeader) SetEvent(v string) *RunCustomHotTopicViewPointAnalysisResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyHeader) SetOriginSessionId(v string) *RunCustomHotTopicViewPointAnalysisResponseBodyHeader {
	s.OriginSessionId = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyHeader) SetSessionId(v string) *RunCustomHotTopicViewPointAnalysisResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyHeader) SetTaskId(v string) *RunCustomHotTopicViewPointAnalysisResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyHeader) SetTraceId(v string) *RunCustomHotTopicViewPointAnalysisResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunCustomHotTopicViewPointAnalysisResponseBodyPayload struct {
	Output *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunCustomHotTopicViewPointAnalysisResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunCustomHotTopicViewPointAnalysisResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayload) GetOutput() *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayload) GetUsage() *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayload) SetOutput(v *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput) *RunCustomHotTopicViewPointAnalysisResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayload) SetUsage(v *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadUsage) *RunCustomHotTopicViewPointAnalysisResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayload) Validate() error {
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
			return err
		}
	}
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput struct {
	Articles []*RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutputArticles `json:"Articles,omitempty" xml:"Articles,omitempty" type:"Repeated"`
	AskUser  []*string                                                              `json:"AskUser,omitempty" xml:"AskUser,omitempty" type:"Repeated"`
	// example:
	//
	// 异步任务ID
	AsyncTaskId *string `json:"AsyncTaskId,omitempty" xml:"AsyncTaskId,omitempty"`
	// example:
	//
	// 模型生成的自定义选题视角的观点
	Attitude *string `json:"Attitude,omitempty" xml:"Attitude,omitempty"`
	// example:
	//
	// xxxxxx
	CustomViewPointId *string `json:"CustomViewPointId,omitempty" xml:"CustomViewPointId,omitempty"`
	// example:
	//
	// 文本生成结果
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// 话题ID
	TopicId *string `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
}

func (s RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput) GetArticles() []*RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutputArticles {
	return s.Articles
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput) GetAskUser() []*string {
	return s.AskUser
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput) GetAsyncTaskId() *string {
	return s.AsyncTaskId
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput) GetAttitude() *string {
	return s.Attitude
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput) GetCustomViewPointId() *string {
	return s.CustomViewPointId
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput) GetTopicId() *string {
	return s.TopicId
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput) SetArticles(v []*RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutputArticles) *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput {
	s.Articles = v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput) SetAskUser(v []*string) *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput {
	s.AskUser = v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput) SetAsyncTaskId(v string) *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput {
	s.AsyncTaskId = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput) SetAttitude(v string) *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput {
	s.Attitude = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput) SetCustomViewPointId(v string) *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput {
	s.CustomViewPointId = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput) SetText(v string) *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput) SetTopicId(v string) *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput {
	s.TopicId = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutput) Validate() error {
	if s.Articles != nil {
		for _, item := range s.Articles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutputArticles struct {
	Author  *string `json:"Author,omitempty" xml:"Author,omitempty"`
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	PubTime *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	Source  *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	Title   *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Url     *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutputArticles) String() string {
	return dara.Prettify(s)
}

func (s RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutputArticles) GoString() string {
	return s.String()
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutputArticles) GetAuthor() *string {
	return s.Author
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutputArticles) GetContent() *string {
	return s.Content
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutputArticles) GetPubTime() *string {
	return s.PubTime
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutputArticles) GetSource() *string {
	return s.Source
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutputArticles) GetSummary() *string {
	return s.Summary
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutputArticles) GetTitle() *string {
	return s.Title
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutputArticles) GetUrl() *string {
	return s.Url
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutputArticles) SetAuthor(v string) *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutputArticles {
	s.Author = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutputArticles) SetContent(v string) *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutputArticles {
	s.Content = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutputArticles) SetPubTime(v string) *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutputArticles {
	s.PubTime = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutputArticles) SetSource(v string) *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutputArticles {
	s.Source = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutputArticles) SetSummary(v string) *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutputArticles {
	s.Summary = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutputArticles) SetTitle(v string) *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutputArticles {
	s.Title = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutputArticles) SetUrl(v string) *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutputArticles {
	s.Url = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadOutputArticles) Validate() error {
	return dara.Validate(s)
}

type RunCustomHotTopicViewPointAnalysisResponseBodyPayloadUsage struct {
	// example:
	//
	// 51
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 79
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// example:
	//
	// 130
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunCustomHotTopicViewPointAnalysisResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunCustomHotTopicViewPointAnalysisResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadUsage) SetInputTokens(v int64) *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunCustomHotTopicViewPointAnalysisResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
