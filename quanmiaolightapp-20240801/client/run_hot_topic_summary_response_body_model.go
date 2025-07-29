// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunHotTopicSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeader(v *RunHotTopicSummaryResponseBodyHeader) *RunHotTopicSummaryResponseBody
	GetHeader() *RunHotTopicSummaryResponseBodyHeader
	SetPayload(v *RunHotTopicSummaryResponseBodyPayload) *RunHotTopicSummaryResponseBody
	GetPayload() *RunHotTopicSummaryResponseBodyPayload
	SetRequestId(v string) *RunHotTopicSummaryResponseBody
	GetRequestId() *string
}

type RunHotTopicSummaryResponseBody struct {
	Header  *RunHotTopicSummaryResponseBodyHeader  `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload *RunHotTopicSummaryResponseBodyPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 5D0E915E-655D-59A8-894F-93873F73AAE5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RunHotTopicSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunHotTopicSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *RunHotTopicSummaryResponseBody) GetHeader() *RunHotTopicSummaryResponseBodyHeader {
	return s.Header
}

func (s *RunHotTopicSummaryResponseBody) GetPayload() *RunHotTopicSummaryResponseBodyPayload {
	return s.Payload
}

func (s *RunHotTopicSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunHotTopicSummaryResponseBody) SetHeader(v *RunHotTopicSummaryResponseBodyHeader) *RunHotTopicSummaryResponseBody {
	s.Header = v
	return s
}

func (s *RunHotTopicSummaryResponseBody) SetPayload(v *RunHotTopicSummaryResponseBodyPayload) *RunHotTopicSummaryResponseBody {
	s.Payload = v
	return s
}

func (s *RunHotTopicSummaryResponseBody) SetRequestId(v string) *RunHotTopicSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunHotTopicSummaryResponseBody) Validate() error {
	return dara.Validate(s)
}

type RunHotTopicSummaryResponseBodyHeader struct {
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

func (s RunHotTopicSummaryResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunHotTopicSummaryResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunHotTopicSummaryResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunHotTopicSummaryResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunHotTopicSummaryResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunHotTopicSummaryResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunHotTopicSummaryResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunHotTopicSummaryResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunHotTopicSummaryResponseBodyHeader) SetErrorCode(v string) *RunHotTopicSummaryResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunHotTopicSummaryResponseBodyHeader) SetErrorMessage(v string) *RunHotTopicSummaryResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunHotTopicSummaryResponseBodyHeader) SetEvent(v string) *RunHotTopicSummaryResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunHotTopicSummaryResponseBodyHeader) SetSessionId(v string) *RunHotTopicSummaryResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunHotTopicSummaryResponseBodyHeader) SetTaskId(v string) *RunHotTopicSummaryResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunHotTopicSummaryResponseBodyHeader) SetTraceId(v string) *RunHotTopicSummaryResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunHotTopicSummaryResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunHotTopicSummaryResponseBodyPayload struct {
	Output *RunHotTopicSummaryResponseBodyPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Usage  *RunHotTopicSummaryResponseBodyPayloadUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunHotTopicSummaryResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunHotTopicSummaryResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunHotTopicSummaryResponseBodyPayload) GetOutput() *RunHotTopicSummaryResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunHotTopicSummaryResponseBodyPayload) GetUsage() *RunHotTopicSummaryResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunHotTopicSummaryResponseBodyPayload) SetOutput(v *RunHotTopicSummaryResponseBodyPayloadOutput) *RunHotTopicSummaryResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunHotTopicSummaryResponseBodyPayload) SetUsage(v *RunHotTopicSummaryResponseBodyPayloadUsage) *RunHotTopicSummaryResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunHotTopicSummaryResponseBodyPayload) Validate() error {
	return dara.Validate(s)
}

type RunHotTopicSummaryResponseBodyPayloadOutput struct {
	Text    *string `json:"text,omitempty" xml:"text,omitempty"`
	TopicId *string `json:"topicId,omitempty" xml:"topicId,omitempty"`
}

func (s RunHotTopicSummaryResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunHotTopicSummaryResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunHotTopicSummaryResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunHotTopicSummaryResponseBodyPayloadOutput) GetTopicId() *string {
	return s.TopicId
}

func (s *RunHotTopicSummaryResponseBodyPayloadOutput) SetText(v string) *RunHotTopicSummaryResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunHotTopicSummaryResponseBodyPayloadOutput) SetTopicId(v string) *RunHotTopicSummaryResponseBodyPayloadOutput {
	s.TopicId = &v
	return s
}

func (s *RunHotTopicSummaryResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunHotTopicSummaryResponseBodyPayloadUsage struct {
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

func (s RunHotTopicSummaryResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunHotTopicSummaryResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunHotTopicSummaryResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunHotTopicSummaryResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunHotTopicSummaryResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunHotTopicSummaryResponseBodyPayloadUsage) SetInputTokens(v int64) *RunHotTopicSummaryResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunHotTopicSummaryResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunHotTopicSummaryResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunHotTopicSummaryResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunHotTopicSummaryResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunHotTopicSummaryResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
