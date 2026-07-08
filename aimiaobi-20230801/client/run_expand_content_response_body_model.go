// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunExpandContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnd(v bool) *RunExpandContentResponseBody
	GetEnd() *bool
	SetHeader(v *RunExpandContentResponseBodyHeader) *RunExpandContentResponseBody
	GetHeader() *RunExpandContentResponseBodyHeader
	SetPayload(v *RunExpandContentResponseBodyPayload) *RunExpandContentResponseBody
	GetPayload() *RunExpandContentResponseBodyPayload
	SetRequestId(v string) *RunExpandContentResponseBody
	GetRequestId() *string
}

type RunExpandContentResponseBody struct {
	// Whether the output is complete. True indicates completion.
	End *bool `json:"End,omitempty" xml:"End,omitempty"`
	// Streaming output header, containing general return information.
	Header *RunExpandContentResponseBodyHeader `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	// Payload of the returned result, JSON structure.
	Payload *RunExpandContentResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// Unique request ID.
	//
	// example:
	//
	// d3be9981-ca2d-4e17-bf31-1c0a628e9f99
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunExpandContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunExpandContentResponseBody) GoString() string {
	return s.String()
}

func (s *RunExpandContentResponseBody) GetEnd() *bool {
	return s.End
}

func (s *RunExpandContentResponseBody) GetHeader() *RunExpandContentResponseBodyHeader {
	return s.Header
}

func (s *RunExpandContentResponseBody) GetPayload() *RunExpandContentResponseBodyPayload {
	return s.Payload
}

func (s *RunExpandContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunExpandContentResponseBody) SetEnd(v bool) *RunExpandContentResponseBody {
	s.End = &v
	return s
}

func (s *RunExpandContentResponseBody) SetHeader(v *RunExpandContentResponseBodyHeader) *RunExpandContentResponseBody {
	s.Header = v
	return s
}

func (s *RunExpandContentResponseBody) SetPayload(v *RunExpandContentResponseBodyPayload) *RunExpandContentResponseBody {
	s.Payload = v
	return s
}

func (s *RunExpandContentResponseBody) SetRequestId(v string) *RunExpandContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunExpandContentResponseBody) Validate() error {
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

type RunExpandContentResponseBodyHeader struct {
	// Error code for exceptions.
	//
	// example:
	//
	// 403
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Error message for exceptions.
	//
	// example:
	//
	// Pop sign mismatch, please check.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Event type.
	//
	// example:
	//
	// result-generated
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// Event description.
	//
	// example:
	//
	// 模型生成事件
	EventInfo *string `json:"EventInfo,omitempty" xml:"EventInfo,omitempty"`
	// Session ID for a single session.
	//
	// example:
	//
	// 3cd10828-0e42-471c-8f1a-931cde20b035
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// Task ID for a single generation task.
	//
	// example:
	//
	// d3be9981-ca2d-4e17-bf31-1c0a628e9f99
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Trace ID for the link.
	//
	// example:
	//
	// 2150451a17191950923411783e2927
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunExpandContentResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunExpandContentResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunExpandContentResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunExpandContentResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunExpandContentResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunExpandContentResponseBodyHeader) GetEventInfo() *string {
	return s.EventInfo
}

func (s *RunExpandContentResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunExpandContentResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunExpandContentResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunExpandContentResponseBodyHeader) SetErrorCode(v string) *RunExpandContentResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunExpandContentResponseBodyHeader) SetErrorMessage(v string) *RunExpandContentResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunExpandContentResponseBodyHeader) SetEvent(v string) *RunExpandContentResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunExpandContentResponseBodyHeader) SetEventInfo(v string) *RunExpandContentResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunExpandContentResponseBodyHeader) SetSessionId(v string) *RunExpandContentResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunExpandContentResponseBodyHeader) SetTaskId(v string) *RunExpandContentResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunExpandContentResponseBodyHeader) SetTraceId(v string) *RunExpandContentResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunExpandContentResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunExpandContentResponseBodyPayload struct {
	// Output content object.
	Output *RunExpandContentResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// Large Language Model (LLM) token usage information.
	Usage *RunExpandContentResponseBodyPayloadUsage `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunExpandContentResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunExpandContentResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunExpandContentResponseBodyPayload) GetOutput() *RunExpandContentResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunExpandContentResponseBodyPayload) GetUsage() *RunExpandContentResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunExpandContentResponseBodyPayload) SetOutput(v *RunExpandContentResponseBodyPayloadOutput) *RunExpandContentResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunExpandContentResponseBodyPayload) SetUsage(v *RunExpandContentResponseBodyPayloadUsage) *RunExpandContentResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunExpandContentResponseBodyPayload) Validate() error {
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

type RunExpandContentResponseBodyPayloadOutput struct {
	// Output content.
	//
	// example:
	//
	// 这是测试输出
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RunExpandContentResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunExpandContentResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunExpandContentResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunExpandContentResponseBodyPayloadOutput) SetText(v string) *RunExpandContentResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunExpandContentResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunExpandContentResponseBodyPayloadUsage struct {
	// Number of input tokens.
	//
	// example:
	//
	// 100
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// Number of output tokens.
	//
	// example:
	//
	// 100
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// Total number of tokens.
	//
	// example:
	//
	// 200
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunExpandContentResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunExpandContentResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunExpandContentResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunExpandContentResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunExpandContentResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunExpandContentResponseBodyPayloadUsage) SetInputTokens(v int64) *RunExpandContentResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunExpandContentResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunExpandContentResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunExpandContentResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunExpandContentResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunExpandContentResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
