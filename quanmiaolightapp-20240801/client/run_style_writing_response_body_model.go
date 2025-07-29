// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunStyleWritingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnd(v bool) *RunStyleWritingResponseBody
	GetEnd() *bool
	SetHeader(v *RunStyleWritingResponseBodyHeader) *RunStyleWritingResponseBody
	GetHeader() *RunStyleWritingResponseBodyHeader
	SetPayload(v *RunStyleWritingResponseBodyPayload) *RunStyleWritingResponseBody
	GetPayload() *RunStyleWritingResponseBodyPayload
}

type RunStyleWritingResponseBody struct {
	// example:
	//
	// true
	End *bool `json:"end,omitempty" xml:"end,omitempty"`
	// example:
	//
	// {"event":"task-progress-start-generating","sessionId":"3cd10828-0e42-471c-8f1a-931cde20b035","taskId":"d3be9981-ca2d-4e17-bf31-1c0a628e9f99","traceId":"66bef4a7f5d61ff3c43f3b710574e175"}
	Header  *RunStyleWritingResponseBodyHeader  `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload *RunStyleWritingResponseBodyPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
}

func (s RunStyleWritingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunStyleWritingResponseBody) GoString() string {
	return s.String()
}

func (s *RunStyleWritingResponseBody) GetEnd() *bool {
	return s.End
}

func (s *RunStyleWritingResponseBody) GetHeader() *RunStyleWritingResponseBodyHeader {
	return s.Header
}

func (s *RunStyleWritingResponseBody) GetPayload() *RunStyleWritingResponseBodyPayload {
	return s.Payload
}

func (s *RunStyleWritingResponseBody) SetEnd(v bool) *RunStyleWritingResponseBody {
	s.End = &v
	return s
}

func (s *RunStyleWritingResponseBody) SetHeader(v *RunStyleWritingResponseBodyHeader) *RunStyleWritingResponseBody {
	s.Header = v
	return s
}

func (s *RunStyleWritingResponseBody) SetPayload(v *RunStyleWritingResponseBodyPayload) *RunStyleWritingResponseBody {
	s.Payload = v
	return s
}

func (s *RunStyleWritingResponseBody) Validate() error {
	return dara.Validate(s)
}

type RunStyleWritingResponseBodyHeader struct {
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
	// task-progress-start-generating
	Event     *string `json:"event,omitempty" xml:"event,omitempty"`
	EventInfo *string `json:"eventInfo,omitempty" xml:"eventInfo,omitempty"`
	// example:
	//
	// 0EB27AE3-CA53-5FAE-83C6-EE66CA4DF5DF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 3cd10828-0e42-471c-8f1a-931cde20b035
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// d3be9981-ca2d-4e17-bf31-1c0a628e9f99
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 2150451a17191950923411783e2927
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s RunStyleWritingResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunStyleWritingResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunStyleWritingResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunStyleWritingResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunStyleWritingResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunStyleWritingResponseBodyHeader) GetEventInfo() *string {
	return s.EventInfo
}

func (s *RunStyleWritingResponseBodyHeader) GetRequestId() *string {
	return s.RequestId
}

func (s *RunStyleWritingResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunStyleWritingResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunStyleWritingResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunStyleWritingResponseBodyHeader) SetErrorCode(v string) *RunStyleWritingResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunStyleWritingResponseBodyHeader) SetErrorMessage(v string) *RunStyleWritingResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunStyleWritingResponseBodyHeader) SetEvent(v string) *RunStyleWritingResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunStyleWritingResponseBodyHeader) SetEventInfo(v string) *RunStyleWritingResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunStyleWritingResponseBodyHeader) SetRequestId(v string) *RunStyleWritingResponseBodyHeader {
	s.RequestId = &v
	return s
}

func (s *RunStyleWritingResponseBodyHeader) SetSessionId(v string) *RunStyleWritingResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunStyleWritingResponseBodyHeader) SetTaskId(v string) *RunStyleWritingResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunStyleWritingResponseBodyHeader) SetTraceId(v string) *RunStyleWritingResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunStyleWritingResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunStyleWritingResponseBodyPayload struct {
	Output *RunStyleWritingResponseBodyPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	// example:
	//
	// {
	//
	//         "inputTokens": 1816,
	//
	//         "outputTokens": 96,
	//
	//         "totalTokens": 1912
	//
	//     }
	Usage *RunStyleWritingResponseBodyPayloadUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunStyleWritingResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunStyleWritingResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunStyleWritingResponseBodyPayload) GetOutput() *RunStyleWritingResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunStyleWritingResponseBodyPayload) GetUsage() *RunStyleWritingResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunStyleWritingResponseBodyPayload) SetOutput(v *RunStyleWritingResponseBodyPayloadOutput) *RunStyleWritingResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunStyleWritingResponseBodyPayload) SetUsage(v *RunStyleWritingResponseBodyPayloadUsage) *RunStyleWritingResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunStyleWritingResponseBodyPayload) Validate() error {
	return dara.Validate(s)
}

type RunStyleWritingResponseBodyPayloadOutput struct {
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s RunStyleWritingResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunStyleWritingResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunStyleWritingResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunStyleWritingResponseBodyPayloadOutput) SetText(v string) *RunStyleWritingResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunStyleWritingResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunStyleWritingResponseBodyPayloadUsage struct {
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

func (s RunStyleWritingResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunStyleWritingResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunStyleWritingResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunStyleWritingResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunStyleWritingResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunStyleWritingResponseBodyPayloadUsage) SetInputTokens(v int64) *RunStyleWritingResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunStyleWritingResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunStyleWritingResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunStyleWritingResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunStyleWritingResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunStyleWritingResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
