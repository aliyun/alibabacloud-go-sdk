// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunDocSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeader(v *RunDocSummaryResponseBodyHeader) *RunDocSummaryResponseBody
	GetHeader() *RunDocSummaryResponseBodyHeader
	SetPayload(v *RunDocSummaryResponseBodyPayload) *RunDocSummaryResponseBody
	GetPayload() *RunDocSummaryResponseBodyPayload
	SetRequestId(v string) *RunDocSummaryResponseBody
	GetRequestId() *string
}

type RunDocSummaryResponseBody struct {
	Header  *RunDocSummaryResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunDocSummaryResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// 3259D344-E871-5DE0-8FFE-CDA21F8D4382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunDocSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunDocSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *RunDocSummaryResponseBody) GetHeader() *RunDocSummaryResponseBodyHeader {
	return s.Header
}

func (s *RunDocSummaryResponseBody) GetPayload() *RunDocSummaryResponseBodyPayload {
	return s.Payload
}

func (s *RunDocSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunDocSummaryResponseBody) SetHeader(v *RunDocSummaryResponseBodyHeader) *RunDocSummaryResponseBody {
	s.Header = v
	return s
}

func (s *RunDocSummaryResponseBody) SetPayload(v *RunDocSummaryResponseBodyPayload) *RunDocSummaryResponseBody {
	s.Payload = v
	return s
}

func (s *RunDocSummaryResponseBody) SetRequestId(v string) *RunDocSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunDocSummaryResponseBody) Validate() error {
	return dara.Validate(s)
}

type RunDocSummaryResponseBodyHeader struct {
	// example:
	//
	// success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// success
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// task-started
	Event     *string `json:"Event,omitempty" xml:"Event,omitempty"`
	EventInfo *string `json:"EventInfo,omitempty" xml:"EventInfo,omitempty"`
	// example:
	//
	// 92e16ccb-92b6-4894-abbf-fc6e2929a0df
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// b057f2fa-2277-477b-babf-cbc062307828
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 2150451a17191950923411783e2927
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunDocSummaryResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunDocSummaryResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunDocSummaryResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunDocSummaryResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunDocSummaryResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunDocSummaryResponseBodyHeader) GetEventInfo() *string {
	return s.EventInfo
}

func (s *RunDocSummaryResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunDocSummaryResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunDocSummaryResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunDocSummaryResponseBodyHeader) SetErrorCode(v string) *RunDocSummaryResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunDocSummaryResponseBodyHeader) SetErrorMessage(v string) *RunDocSummaryResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunDocSummaryResponseBodyHeader) SetEvent(v string) *RunDocSummaryResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunDocSummaryResponseBodyHeader) SetEventInfo(v string) *RunDocSummaryResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunDocSummaryResponseBodyHeader) SetSessionId(v string) *RunDocSummaryResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunDocSummaryResponseBodyHeader) SetTaskId(v string) *RunDocSummaryResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunDocSummaryResponseBodyHeader) SetTraceId(v string) *RunDocSummaryResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunDocSummaryResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunDocSummaryResponseBodyPayload struct {
	Output *RunDocSummaryResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunDocSummaryResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunDocSummaryResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunDocSummaryResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunDocSummaryResponseBodyPayload) GetOutput() *RunDocSummaryResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunDocSummaryResponseBodyPayload) GetUsage() *RunDocSummaryResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunDocSummaryResponseBodyPayload) SetOutput(v *RunDocSummaryResponseBodyPayloadOutput) *RunDocSummaryResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunDocSummaryResponseBodyPayload) SetUsage(v *RunDocSummaryResponseBodyPayloadUsage) *RunDocSummaryResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunDocSummaryResponseBodyPayload) Validate() error {
	return dara.Validate(s)
}

type RunDocSummaryResponseBodyPayloadOutput struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s RunDocSummaryResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunDocSummaryResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunDocSummaryResponseBodyPayloadOutput) GetContent() *string {
	return s.Content
}

func (s *RunDocSummaryResponseBodyPayloadOutput) SetContent(v string) *RunDocSummaryResponseBodyPayloadOutput {
	s.Content = &v
	return s
}

func (s *RunDocSummaryResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunDocSummaryResponseBodyPayloadUsage struct {
	// example:
	//
	// 100
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 100
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// example:
	//
	// 200
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunDocSummaryResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunDocSummaryResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunDocSummaryResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunDocSummaryResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunDocSummaryResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunDocSummaryResponseBodyPayloadUsage) SetInputTokens(v int64) *RunDocSummaryResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunDocSummaryResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunDocSummaryResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunDocSummaryResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunDocSummaryResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunDocSummaryResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
