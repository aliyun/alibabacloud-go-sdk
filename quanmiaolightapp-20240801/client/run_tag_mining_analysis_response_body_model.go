// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunTagMiningAnalysisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeader(v *RunTagMiningAnalysisResponseBodyHeader) *RunTagMiningAnalysisResponseBody
	GetHeader() *RunTagMiningAnalysisResponseBodyHeader
	SetPayload(v *RunTagMiningAnalysisResponseBodyPayload) *RunTagMiningAnalysisResponseBody
	GetPayload() *RunTagMiningAnalysisResponseBodyPayload
	SetRequestId(v string) *RunTagMiningAnalysisResponseBody
	GetRequestId() *string
}

type RunTagMiningAnalysisResponseBody struct {
	Header  *RunTagMiningAnalysisResponseBodyHeader  `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload *RunTagMiningAnalysisResponseBodyPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 085BE2D2-BB7E-59A6-B688-F2CB32124E7F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RunTagMiningAnalysisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunTagMiningAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *RunTagMiningAnalysisResponseBody) GetHeader() *RunTagMiningAnalysisResponseBodyHeader {
	return s.Header
}

func (s *RunTagMiningAnalysisResponseBody) GetPayload() *RunTagMiningAnalysisResponseBodyPayload {
	return s.Payload
}

func (s *RunTagMiningAnalysisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunTagMiningAnalysisResponseBody) SetHeader(v *RunTagMiningAnalysisResponseBodyHeader) *RunTagMiningAnalysisResponseBody {
	s.Header = v
	return s
}

func (s *RunTagMiningAnalysisResponseBody) SetPayload(v *RunTagMiningAnalysisResponseBodyPayload) *RunTagMiningAnalysisResponseBody {
	s.Payload = v
	return s
}

func (s *RunTagMiningAnalysisResponseBody) SetRequestId(v string) *RunTagMiningAnalysisResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunTagMiningAnalysisResponseBody) Validate() error {
	return dara.Validate(s)
}

type RunTagMiningAnalysisResponseBodyHeader struct {
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

func (s RunTagMiningAnalysisResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunTagMiningAnalysisResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunTagMiningAnalysisResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunTagMiningAnalysisResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunTagMiningAnalysisResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunTagMiningAnalysisResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunTagMiningAnalysisResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunTagMiningAnalysisResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunTagMiningAnalysisResponseBodyHeader) SetErrorCode(v string) *RunTagMiningAnalysisResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunTagMiningAnalysisResponseBodyHeader) SetErrorMessage(v string) *RunTagMiningAnalysisResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunTagMiningAnalysisResponseBodyHeader) SetEvent(v string) *RunTagMiningAnalysisResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunTagMiningAnalysisResponseBodyHeader) SetSessionId(v string) *RunTagMiningAnalysisResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunTagMiningAnalysisResponseBodyHeader) SetTaskId(v string) *RunTagMiningAnalysisResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunTagMiningAnalysisResponseBodyHeader) SetTraceId(v string) *RunTagMiningAnalysisResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunTagMiningAnalysisResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunTagMiningAnalysisResponseBodyPayload struct {
	Output *RunTagMiningAnalysisResponseBodyPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Usage  *RunTagMiningAnalysisResponseBodyPayloadUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunTagMiningAnalysisResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunTagMiningAnalysisResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunTagMiningAnalysisResponseBodyPayload) GetOutput() *RunTagMiningAnalysisResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunTagMiningAnalysisResponseBodyPayload) GetUsage() *RunTagMiningAnalysisResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunTagMiningAnalysisResponseBodyPayload) SetOutput(v *RunTagMiningAnalysisResponseBodyPayloadOutput) *RunTagMiningAnalysisResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunTagMiningAnalysisResponseBodyPayload) SetUsage(v *RunTagMiningAnalysisResponseBodyPayloadUsage) *RunTagMiningAnalysisResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunTagMiningAnalysisResponseBodyPayload) Validate() error {
	return dara.Validate(s)
}

type RunTagMiningAnalysisResponseBodyPayloadOutput struct {
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s RunTagMiningAnalysisResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunTagMiningAnalysisResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunTagMiningAnalysisResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunTagMiningAnalysisResponseBodyPayloadOutput) SetText(v string) *RunTagMiningAnalysisResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunTagMiningAnalysisResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunTagMiningAnalysisResponseBodyPayloadUsage struct {
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

func (s RunTagMiningAnalysisResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunTagMiningAnalysisResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunTagMiningAnalysisResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunTagMiningAnalysisResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunTagMiningAnalysisResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunTagMiningAnalysisResponseBodyPayloadUsage) SetInputTokens(v int64) *RunTagMiningAnalysisResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunTagMiningAnalysisResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunTagMiningAnalysisResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunTagMiningAnalysisResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunTagMiningAnalysisResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunTagMiningAnalysisResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
