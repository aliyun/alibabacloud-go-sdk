// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunEssayCorrectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeader(v *RunEssayCorrectionResponseBodyHeader) *RunEssayCorrectionResponseBody
	GetHeader() *RunEssayCorrectionResponseBodyHeader
	SetPayload(v *RunEssayCorrectionResponseBodyPayload) *RunEssayCorrectionResponseBody
	GetPayload() *RunEssayCorrectionResponseBodyPayload
	SetRequestId(v string) *RunEssayCorrectionResponseBody
	GetRequestId() *string
}

type RunEssayCorrectionResponseBody struct {
	Header  *RunEssayCorrectionResponseBodyHeader  `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload *RunEssayCorrectionResponseBodyPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 117F5ABE-CF02-5502-9A3F-E56BC9081A64
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RunEssayCorrectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunEssayCorrectionResponseBody) GoString() string {
	return s.String()
}

func (s *RunEssayCorrectionResponseBody) GetHeader() *RunEssayCorrectionResponseBodyHeader {
	return s.Header
}

func (s *RunEssayCorrectionResponseBody) GetPayload() *RunEssayCorrectionResponseBodyPayload {
	return s.Payload
}

func (s *RunEssayCorrectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunEssayCorrectionResponseBody) SetHeader(v *RunEssayCorrectionResponseBodyHeader) *RunEssayCorrectionResponseBody {
	s.Header = v
	return s
}

func (s *RunEssayCorrectionResponseBody) SetPayload(v *RunEssayCorrectionResponseBodyPayload) *RunEssayCorrectionResponseBody {
	s.Payload = v
	return s
}

func (s *RunEssayCorrectionResponseBody) SetRequestId(v string) *RunEssayCorrectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunEssayCorrectionResponseBody) Validate() error {
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

type RunEssayCorrectionResponseBodyHeader struct {
	// example:
	//
	// AccessForbidden
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// task-finished
	Event *string `json:"event,omitempty" xml:"event,omitempty"`
	// example:
	//
	// xxxx
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// xxxx
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// xxxxx
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s RunEssayCorrectionResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunEssayCorrectionResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunEssayCorrectionResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunEssayCorrectionResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunEssayCorrectionResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunEssayCorrectionResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunEssayCorrectionResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunEssayCorrectionResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunEssayCorrectionResponseBodyHeader) SetErrorCode(v string) *RunEssayCorrectionResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunEssayCorrectionResponseBodyHeader) SetErrorMessage(v string) *RunEssayCorrectionResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunEssayCorrectionResponseBodyHeader) SetEvent(v string) *RunEssayCorrectionResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunEssayCorrectionResponseBodyHeader) SetSessionId(v string) *RunEssayCorrectionResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunEssayCorrectionResponseBodyHeader) SetTaskId(v string) *RunEssayCorrectionResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunEssayCorrectionResponseBodyHeader) SetTraceId(v string) *RunEssayCorrectionResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunEssayCorrectionResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunEssayCorrectionResponseBodyPayload struct {
	Output *RunEssayCorrectionResponseBodyPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Usage  *RunEssayCorrectionResponseBodyPayloadUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunEssayCorrectionResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunEssayCorrectionResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunEssayCorrectionResponseBodyPayload) GetOutput() *RunEssayCorrectionResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunEssayCorrectionResponseBodyPayload) GetUsage() *RunEssayCorrectionResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunEssayCorrectionResponseBodyPayload) SetOutput(v *RunEssayCorrectionResponseBodyPayloadOutput) *RunEssayCorrectionResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunEssayCorrectionResponseBodyPayload) SetUsage(v *RunEssayCorrectionResponseBodyPayloadUsage) *RunEssayCorrectionResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunEssayCorrectionResponseBodyPayload) Validate() error {
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

type RunEssayCorrectionResponseBodyPayloadOutput struct {
	// example:
	//
	// 50
	Score *int32  `json:"score,omitempty" xml:"score,omitempty"`
	Text  *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s RunEssayCorrectionResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunEssayCorrectionResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunEssayCorrectionResponseBodyPayloadOutput) GetScore() *int32 {
	return s.Score
}

func (s *RunEssayCorrectionResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunEssayCorrectionResponseBodyPayloadOutput) SetScore(v int32) *RunEssayCorrectionResponseBodyPayloadOutput {
	s.Score = &v
	return s
}

func (s *RunEssayCorrectionResponseBodyPayloadOutput) SetText(v string) *RunEssayCorrectionResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunEssayCorrectionResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunEssayCorrectionResponseBodyPayloadUsage struct {
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

func (s RunEssayCorrectionResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunEssayCorrectionResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunEssayCorrectionResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunEssayCorrectionResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunEssayCorrectionResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunEssayCorrectionResponseBodyPayloadUsage) SetInputTokens(v int64) *RunEssayCorrectionResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunEssayCorrectionResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunEssayCorrectionResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunEssayCorrectionResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunEssayCorrectionResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunEssayCorrectionResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
