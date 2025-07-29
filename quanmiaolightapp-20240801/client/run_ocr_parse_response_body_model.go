// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunOcrParseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeader(v *RunOcrParseResponseBodyHeader) *RunOcrParseResponseBody
	GetHeader() *RunOcrParseResponseBodyHeader
	SetPayload(v *RunOcrParseResponseBodyPayload) *RunOcrParseResponseBody
	GetPayload() *RunOcrParseResponseBodyPayload
	SetRequestId(v string) *RunOcrParseResponseBody
	GetRequestId() *string
}

type RunOcrParseResponseBody struct {
	Header  *RunOcrParseResponseBodyHeader  `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload *RunOcrParseResponseBodyPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 117F5ABE-CF02-5502-9A3F-E56BC9081A64
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RunOcrParseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunOcrParseResponseBody) GoString() string {
	return s.String()
}

func (s *RunOcrParseResponseBody) GetHeader() *RunOcrParseResponseBodyHeader {
	return s.Header
}

func (s *RunOcrParseResponseBody) GetPayload() *RunOcrParseResponseBodyPayload {
	return s.Payload
}

func (s *RunOcrParseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunOcrParseResponseBody) SetHeader(v *RunOcrParseResponseBodyHeader) *RunOcrParseResponseBody {
	s.Header = v
	return s
}

func (s *RunOcrParseResponseBody) SetPayload(v *RunOcrParseResponseBodyPayload) *RunOcrParseResponseBody {
	s.Payload = v
	return s
}

func (s *RunOcrParseResponseBody) SetRequestId(v string) *RunOcrParseResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunOcrParseResponseBody) Validate() error {
	return dara.Validate(s)
}

type RunOcrParseResponseBodyHeader struct {
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

func (s RunOcrParseResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunOcrParseResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunOcrParseResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunOcrParseResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunOcrParseResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunOcrParseResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunOcrParseResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunOcrParseResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunOcrParseResponseBodyHeader) SetErrorCode(v string) *RunOcrParseResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunOcrParseResponseBodyHeader) SetErrorMessage(v string) *RunOcrParseResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunOcrParseResponseBodyHeader) SetEvent(v string) *RunOcrParseResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunOcrParseResponseBodyHeader) SetSessionId(v string) *RunOcrParseResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunOcrParseResponseBodyHeader) SetTaskId(v string) *RunOcrParseResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunOcrParseResponseBodyHeader) SetTraceId(v string) *RunOcrParseResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunOcrParseResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunOcrParseResponseBodyPayload struct {
	Output *RunOcrParseResponseBodyPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Usage  *RunOcrParseResponseBodyPayloadUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunOcrParseResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunOcrParseResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunOcrParseResponseBodyPayload) GetOutput() *RunOcrParseResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunOcrParseResponseBodyPayload) GetUsage() *RunOcrParseResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunOcrParseResponseBodyPayload) SetOutput(v *RunOcrParseResponseBodyPayloadOutput) *RunOcrParseResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunOcrParseResponseBodyPayload) SetUsage(v *RunOcrParseResponseBodyPayloadUsage) *RunOcrParseResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunOcrParseResponseBodyPayload) Validate() error {
	return dara.Validate(s)
}

type RunOcrParseResponseBodyPayloadOutput struct {
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s RunOcrParseResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunOcrParseResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunOcrParseResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunOcrParseResponseBodyPayloadOutput) SetText(v string) *RunOcrParseResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunOcrParseResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunOcrParseResponseBodyPayloadUsage struct {
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

func (s RunOcrParseResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunOcrParseResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunOcrParseResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunOcrParseResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunOcrParseResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunOcrParseResponseBodyPayloadUsage) SetInputTokens(v int64) *RunOcrParseResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunOcrParseResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunOcrParseResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunOcrParseResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunOcrParseResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunOcrParseResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
