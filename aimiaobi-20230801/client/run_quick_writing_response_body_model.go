// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunQuickWritingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnd(v bool) *RunQuickWritingResponseBody
	GetEnd() *bool
	SetHeader(v *RunQuickWritingResponseBodyHeader) *RunQuickWritingResponseBody
	GetHeader() *RunQuickWritingResponseBodyHeader
	SetPayload(v *RunQuickWritingResponseBodyPayload) *RunQuickWritingResponseBody
	GetPayload() *RunQuickWritingResponseBodyPayload
	SetRequestId(v string) *RunQuickWritingResponseBody
	GetRequestId() *string
}

type RunQuickWritingResponseBody struct {
	// example:
	//
	// false
	End     *bool                               `json:"End,omitempty" xml:"End,omitempty"`
	Header  *RunQuickWritingResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunQuickWritingResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunQuickWritingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunQuickWritingResponseBody) GoString() string {
	return s.String()
}

func (s *RunQuickWritingResponseBody) GetEnd() *bool {
	return s.End
}

func (s *RunQuickWritingResponseBody) GetHeader() *RunQuickWritingResponseBodyHeader {
	return s.Header
}

func (s *RunQuickWritingResponseBody) GetPayload() *RunQuickWritingResponseBodyPayload {
	return s.Payload
}

func (s *RunQuickWritingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunQuickWritingResponseBody) SetEnd(v bool) *RunQuickWritingResponseBody {
	s.End = &v
	return s
}

func (s *RunQuickWritingResponseBody) SetHeader(v *RunQuickWritingResponseBodyHeader) *RunQuickWritingResponseBody {
	s.Header = v
	return s
}

func (s *RunQuickWritingResponseBody) SetPayload(v *RunQuickWritingResponseBodyPayload) *RunQuickWritingResponseBody {
	s.Payload = v
	return s
}

func (s *RunQuickWritingResponseBody) SetRequestId(v string) *RunQuickWritingResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunQuickWritingResponseBody) Validate() error {
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

type RunQuickWritingResponseBodyHeader struct {
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

func (s RunQuickWritingResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunQuickWritingResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunQuickWritingResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunQuickWritingResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunQuickWritingResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunQuickWritingResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunQuickWritingResponseBodyHeader) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunQuickWritingResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunQuickWritingResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunQuickWritingResponseBodyHeader) SetErrorCode(v string) *RunQuickWritingResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunQuickWritingResponseBodyHeader) SetErrorMessage(v string) *RunQuickWritingResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunQuickWritingResponseBodyHeader) SetEvent(v string) *RunQuickWritingResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunQuickWritingResponseBodyHeader) SetSessionId(v string) *RunQuickWritingResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunQuickWritingResponseBodyHeader) SetStatusCode(v int32) *RunQuickWritingResponseBodyHeader {
	s.StatusCode = &v
	return s
}

func (s *RunQuickWritingResponseBodyHeader) SetTaskId(v string) *RunQuickWritingResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunQuickWritingResponseBodyHeader) SetTraceId(v string) *RunQuickWritingResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunQuickWritingResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunQuickWritingResponseBodyPayload struct {
	Output *RunQuickWritingResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunQuickWritingResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunQuickWritingResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunQuickWritingResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunQuickWritingResponseBodyPayload) GetOutput() *RunQuickWritingResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunQuickWritingResponseBodyPayload) GetUsage() *RunQuickWritingResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunQuickWritingResponseBodyPayload) SetOutput(v *RunQuickWritingResponseBodyPayloadOutput) *RunQuickWritingResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunQuickWritingResponseBodyPayload) SetUsage(v *RunQuickWritingResponseBodyPayloadUsage) *RunQuickWritingResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunQuickWritingResponseBodyPayload) Validate() error {
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

type RunQuickWritingResponseBodyPayloadOutput struct {
	// example:
	//
	// 文本生成结果
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RunQuickWritingResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunQuickWritingResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunQuickWritingResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunQuickWritingResponseBodyPayloadOutput) SetText(v string) *RunQuickWritingResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunQuickWritingResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunQuickWritingResponseBodyPayloadUsage struct {
	// example:
	//
	// 78
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 34
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// example:
	//
	// 38
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunQuickWritingResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunQuickWritingResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunQuickWritingResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunQuickWritingResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunQuickWritingResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunQuickWritingResponseBodyPayloadUsage) SetInputTokens(v int64) *RunQuickWritingResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunQuickWritingResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunQuickWritingResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunQuickWritingResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunQuickWritingResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunQuickWritingResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
