// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunNetworkContentAuditResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeader(v *RunNetworkContentAuditResponseBodyHeader) *RunNetworkContentAuditResponseBody
	GetHeader() *RunNetworkContentAuditResponseBodyHeader
	SetPayload(v *RunNetworkContentAuditResponseBodyPayload) *RunNetworkContentAuditResponseBody
	GetPayload() *RunNetworkContentAuditResponseBodyPayload
	SetRequestId(v string) *RunNetworkContentAuditResponseBody
	GetRequestId() *string
}

type RunNetworkContentAuditResponseBody struct {
	Header  *RunNetworkContentAuditResponseBodyHeader  `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload *RunNetworkContentAuditResponseBodyPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 5D0E915E-655D-59A8-894F-93873F73AAE5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RunNetworkContentAuditResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunNetworkContentAuditResponseBody) GoString() string {
	return s.String()
}

func (s *RunNetworkContentAuditResponseBody) GetHeader() *RunNetworkContentAuditResponseBodyHeader {
	return s.Header
}

func (s *RunNetworkContentAuditResponseBody) GetPayload() *RunNetworkContentAuditResponseBodyPayload {
	return s.Payload
}

func (s *RunNetworkContentAuditResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunNetworkContentAuditResponseBody) SetHeader(v *RunNetworkContentAuditResponseBodyHeader) *RunNetworkContentAuditResponseBody {
	s.Header = v
	return s
}

func (s *RunNetworkContentAuditResponseBody) SetPayload(v *RunNetworkContentAuditResponseBodyPayload) *RunNetworkContentAuditResponseBody {
	s.Payload = v
	return s
}

func (s *RunNetworkContentAuditResponseBody) SetRequestId(v string) *RunNetworkContentAuditResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunNetworkContentAuditResponseBody) Validate() error {
	return dara.Validate(s)
}

type RunNetworkContentAuditResponseBodyHeader struct {
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

func (s RunNetworkContentAuditResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunNetworkContentAuditResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunNetworkContentAuditResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunNetworkContentAuditResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunNetworkContentAuditResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunNetworkContentAuditResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunNetworkContentAuditResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunNetworkContentAuditResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunNetworkContentAuditResponseBodyHeader) SetErrorCode(v string) *RunNetworkContentAuditResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunNetworkContentAuditResponseBodyHeader) SetErrorMessage(v string) *RunNetworkContentAuditResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunNetworkContentAuditResponseBodyHeader) SetEvent(v string) *RunNetworkContentAuditResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunNetworkContentAuditResponseBodyHeader) SetSessionId(v string) *RunNetworkContentAuditResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunNetworkContentAuditResponseBodyHeader) SetTaskId(v string) *RunNetworkContentAuditResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunNetworkContentAuditResponseBodyHeader) SetTraceId(v string) *RunNetworkContentAuditResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunNetworkContentAuditResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunNetworkContentAuditResponseBodyPayload struct {
	Output *RunNetworkContentAuditResponseBodyPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Usage  *RunNetworkContentAuditResponseBodyPayloadUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunNetworkContentAuditResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunNetworkContentAuditResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunNetworkContentAuditResponseBodyPayload) GetOutput() *RunNetworkContentAuditResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunNetworkContentAuditResponseBodyPayload) GetUsage() *RunNetworkContentAuditResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunNetworkContentAuditResponseBodyPayload) SetOutput(v *RunNetworkContentAuditResponseBodyPayloadOutput) *RunNetworkContentAuditResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunNetworkContentAuditResponseBodyPayload) SetUsage(v *RunNetworkContentAuditResponseBodyPayloadUsage) *RunNetworkContentAuditResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunNetworkContentAuditResponseBodyPayload) Validate() error {
	return dara.Validate(s)
}

type RunNetworkContentAuditResponseBodyPayloadOutput struct {
	// example:
	//
	// xxxx
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s RunNetworkContentAuditResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunNetworkContentAuditResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunNetworkContentAuditResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunNetworkContentAuditResponseBodyPayloadOutput) SetText(v string) *RunNetworkContentAuditResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunNetworkContentAuditResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunNetworkContentAuditResponseBodyPayloadUsage struct {
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

func (s RunNetworkContentAuditResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunNetworkContentAuditResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunNetworkContentAuditResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunNetworkContentAuditResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunNetworkContentAuditResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunNetworkContentAuditResponseBodyPayloadUsage) SetInputTokens(v int64) *RunNetworkContentAuditResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunNetworkContentAuditResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunNetworkContentAuditResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunNetworkContentAuditResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunNetworkContentAuditResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunNetworkContentAuditResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
