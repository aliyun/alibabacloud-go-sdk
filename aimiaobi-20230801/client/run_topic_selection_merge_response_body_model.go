// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunTopicSelectionMergeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnd(v bool) *RunTopicSelectionMergeResponseBody
	GetEnd() *bool
	SetHeader(v *RunTopicSelectionMergeResponseBodyHeader) *RunTopicSelectionMergeResponseBody
	GetHeader() *RunTopicSelectionMergeResponseBodyHeader
	SetPayload(v *RunTopicSelectionMergeResponseBodyPayload) *RunTopicSelectionMergeResponseBody
	GetPayload() *RunTopicSelectionMergeResponseBodyPayload
	SetRequestId(v string) *RunTopicSelectionMergeResponseBody
	GetRequestId() *string
}

type RunTopicSelectionMergeResponseBody struct {
	// example:
	//
	// true
	End     *bool                                      `json:"End,omitempty" xml:"End,omitempty"`
	Header  *RunTopicSelectionMergeResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunTopicSelectionMergeResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunTopicSelectionMergeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunTopicSelectionMergeResponseBody) GoString() string {
	return s.String()
}

func (s *RunTopicSelectionMergeResponseBody) GetEnd() *bool {
	return s.End
}

func (s *RunTopicSelectionMergeResponseBody) GetHeader() *RunTopicSelectionMergeResponseBodyHeader {
	return s.Header
}

func (s *RunTopicSelectionMergeResponseBody) GetPayload() *RunTopicSelectionMergeResponseBodyPayload {
	return s.Payload
}

func (s *RunTopicSelectionMergeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunTopicSelectionMergeResponseBody) SetEnd(v bool) *RunTopicSelectionMergeResponseBody {
	s.End = &v
	return s
}

func (s *RunTopicSelectionMergeResponseBody) SetHeader(v *RunTopicSelectionMergeResponseBodyHeader) *RunTopicSelectionMergeResponseBody {
	s.Header = v
	return s
}

func (s *RunTopicSelectionMergeResponseBody) SetPayload(v *RunTopicSelectionMergeResponseBodyPayload) *RunTopicSelectionMergeResponseBody {
	s.Payload = v
	return s
}

func (s *RunTopicSelectionMergeResponseBody) SetRequestId(v string) *RunTopicSelectionMergeResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunTopicSelectionMergeResponseBody) Validate() error {
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

type RunTopicSelectionMergeResponseBodyHeader struct {
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

func (s RunTopicSelectionMergeResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunTopicSelectionMergeResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunTopicSelectionMergeResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunTopicSelectionMergeResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunTopicSelectionMergeResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunTopicSelectionMergeResponseBodyHeader) GetOriginSessionId() *string {
	return s.OriginSessionId
}

func (s *RunTopicSelectionMergeResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunTopicSelectionMergeResponseBodyHeader) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunTopicSelectionMergeResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunTopicSelectionMergeResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunTopicSelectionMergeResponseBodyHeader) SetErrorCode(v string) *RunTopicSelectionMergeResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunTopicSelectionMergeResponseBodyHeader) SetErrorMessage(v string) *RunTopicSelectionMergeResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunTopicSelectionMergeResponseBodyHeader) SetEvent(v string) *RunTopicSelectionMergeResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunTopicSelectionMergeResponseBodyHeader) SetOriginSessionId(v string) *RunTopicSelectionMergeResponseBodyHeader {
	s.OriginSessionId = &v
	return s
}

func (s *RunTopicSelectionMergeResponseBodyHeader) SetSessionId(v string) *RunTopicSelectionMergeResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunTopicSelectionMergeResponseBodyHeader) SetStatusCode(v int32) *RunTopicSelectionMergeResponseBodyHeader {
	s.StatusCode = &v
	return s
}

func (s *RunTopicSelectionMergeResponseBodyHeader) SetTaskId(v string) *RunTopicSelectionMergeResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunTopicSelectionMergeResponseBodyHeader) SetTraceId(v string) *RunTopicSelectionMergeResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunTopicSelectionMergeResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunTopicSelectionMergeResponseBodyPayload struct {
	Output *RunTopicSelectionMergeResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunTopicSelectionMergeResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunTopicSelectionMergeResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunTopicSelectionMergeResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunTopicSelectionMergeResponseBodyPayload) GetOutput() *RunTopicSelectionMergeResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunTopicSelectionMergeResponseBodyPayload) GetUsage() *RunTopicSelectionMergeResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunTopicSelectionMergeResponseBodyPayload) SetOutput(v *RunTopicSelectionMergeResponseBodyPayloadOutput) *RunTopicSelectionMergeResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunTopicSelectionMergeResponseBodyPayload) SetUsage(v *RunTopicSelectionMergeResponseBodyPayloadUsage) *RunTopicSelectionMergeResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunTopicSelectionMergeResponseBodyPayload) Validate() error {
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

type RunTopicSelectionMergeResponseBodyPayloadOutput struct {
	// example:
	//
	// 文本生成结果
	Text  *string         `json:"Text,omitempty" xml:"Text,omitempty"`
	Topic *TopicSelection `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s RunTopicSelectionMergeResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunTopicSelectionMergeResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunTopicSelectionMergeResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunTopicSelectionMergeResponseBodyPayloadOutput) GetTopic() *TopicSelection {
	return s.Topic
}

func (s *RunTopicSelectionMergeResponseBodyPayloadOutput) SetText(v string) *RunTopicSelectionMergeResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunTopicSelectionMergeResponseBodyPayloadOutput) SetTopic(v *TopicSelection) *RunTopicSelectionMergeResponseBodyPayloadOutput {
	s.Topic = v
	return s
}

func (s *RunTopicSelectionMergeResponseBodyPayloadOutput) Validate() error {
	if s.Topic != nil {
		if err := s.Topic.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunTopicSelectionMergeResponseBodyPayloadUsage struct {
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

func (s RunTopicSelectionMergeResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunTopicSelectionMergeResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunTopicSelectionMergeResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunTopicSelectionMergeResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunTopicSelectionMergeResponseBodyPayloadUsage) GetTokenMap() map[string]*int64 {
	return s.TokenMap
}

func (s *RunTopicSelectionMergeResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunTopicSelectionMergeResponseBodyPayloadUsage) SetInputTokens(v int64) *RunTopicSelectionMergeResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunTopicSelectionMergeResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunTopicSelectionMergeResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunTopicSelectionMergeResponseBodyPayloadUsage) SetTokenMap(v map[string]*int64) *RunTopicSelectionMergeResponseBodyPayloadUsage {
	s.TokenMap = v
	return s
}

func (s *RunTopicSelectionMergeResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunTopicSelectionMergeResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunTopicSelectionMergeResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
