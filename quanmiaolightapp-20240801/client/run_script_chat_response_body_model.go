// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunScriptChatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnd(v bool) *RunScriptChatResponseBody
	GetEnd() *bool
	SetHeader(v *RunScriptChatResponseBodyHeader) *RunScriptChatResponseBody
	GetHeader() *RunScriptChatResponseBodyHeader
	SetPayload(v *RunScriptChatResponseBodyPayload) *RunScriptChatResponseBody
	GetPayload() *RunScriptChatResponseBodyPayload
}

type RunScriptChatResponseBody struct {
	// example:
	//
	// true
	End     *bool                             `json:"end,omitempty" xml:"end,omitempty"`
	Header  *RunScriptChatResponseBodyHeader  `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload *RunScriptChatResponseBodyPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
}

func (s RunScriptChatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunScriptChatResponseBody) GoString() string {
	return s.String()
}

func (s *RunScriptChatResponseBody) GetEnd() *bool {
	return s.End
}

func (s *RunScriptChatResponseBody) GetHeader() *RunScriptChatResponseBodyHeader {
	return s.Header
}

func (s *RunScriptChatResponseBody) GetPayload() *RunScriptChatResponseBodyPayload {
	return s.Payload
}

func (s *RunScriptChatResponseBody) SetEnd(v bool) *RunScriptChatResponseBody {
	s.End = &v
	return s
}

func (s *RunScriptChatResponseBody) SetHeader(v *RunScriptChatResponseBodyHeader) *RunScriptChatResponseBody {
	s.Header = v
	return s
}

func (s *RunScriptChatResponseBody) SetPayload(v *RunScriptChatResponseBodyPayload) *RunScriptChatResponseBody {
	s.Payload = v
	return s
}

func (s *RunScriptChatResponseBody) Validate() error {
	return dara.Validate(s)
}

type RunScriptChatResponseBodyHeader struct {
	// example:
	//
	// 403
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// Pop sign mismatch, please check log.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// result-generated
	Event     *string `json:"event,omitempty" xml:"event,omitempty"`
	EventInfo *string `json:"eventInfo,omitempty" xml:"eventInfo,omitempty"`
	// example:
	//
	// F8A35034-EDCF-5C50-95A5-1044316F36E3
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 147648697127_914847410985_1730600302167
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// d3be9981-ca2d-4e17-bf31-1c0a628e9f99
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 2150432017236011824686132ecdbc
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s RunScriptChatResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunScriptChatResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunScriptChatResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunScriptChatResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunScriptChatResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunScriptChatResponseBodyHeader) GetEventInfo() *string {
	return s.EventInfo
}

func (s *RunScriptChatResponseBodyHeader) GetRequestId() *string {
	return s.RequestId
}

func (s *RunScriptChatResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunScriptChatResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunScriptChatResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunScriptChatResponseBodyHeader) SetErrorCode(v string) *RunScriptChatResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunScriptChatResponseBodyHeader) SetErrorMessage(v string) *RunScriptChatResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunScriptChatResponseBodyHeader) SetEvent(v string) *RunScriptChatResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunScriptChatResponseBodyHeader) SetEventInfo(v string) *RunScriptChatResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunScriptChatResponseBodyHeader) SetRequestId(v string) *RunScriptChatResponseBodyHeader {
	s.RequestId = &v
	return s
}

func (s *RunScriptChatResponseBodyHeader) SetSessionId(v string) *RunScriptChatResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunScriptChatResponseBodyHeader) SetTaskId(v string) *RunScriptChatResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunScriptChatResponseBodyHeader) SetTraceId(v string) *RunScriptChatResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunScriptChatResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunScriptChatResponseBodyPayload struct {
	Output *RunScriptChatResponseBodyPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Usage  *RunScriptChatResponseBodyPayloadUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunScriptChatResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunScriptChatResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunScriptChatResponseBodyPayload) GetOutput() *RunScriptChatResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunScriptChatResponseBodyPayload) GetUsage() *RunScriptChatResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunScriptChatResponseBodyPayload) SetOutput(v *RunScriptChatResponseBodyPayloadOutput) *RunScriptChatResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunScriptChatResponseBodyPayload) SetUsage(v *RunScriptChatResponseBodyPayloadUsage) *RunScriptChatResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunScriptChatResponseBodyPayload) Validate() error {
	return dara.Validate(s)
}

type RunScriptChatResponseBodyPayloadOutput struct {
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s RunScriptChatResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunScriptChatResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunScriptChatResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunScriptChatResponseBodyPayloadOutput) SetText(v string) *RunScriptChatResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunScriptChatResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunScriptChatResponseBodyPayloadUsage struct {
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

func (s RunScriptChatResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunScriptChatResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunScriptChatResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunScriptChatResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunScriptChatResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunScriptChatResponseBodyPayloadUsage) SetInputTokens(v int64) *RunScriptChatResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunScriptChatResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunScriptChatResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunScriptChatResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunScriptChatResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunScriptChatResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
