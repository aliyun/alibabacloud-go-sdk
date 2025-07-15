// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunDocBrainmapResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeader(v *RunDocBrainmapResponseBodyHeader) *RunDocBrainmapResponseBody
	GetHeader() *RunDocBrainmapResponseBodyHeader
	SetPayload(v *RunDocBrainmapResponseBodyPayload) *RunDocBrainmapResponseBody
	GetPayload() *RunDocBrainmapResponseBodyPayload
	SetRequestId(v string) *RunDocBrainmapResponseBody
	GetRequestId() *string
}

type RunDocBrainmapResponseBody struct {
	Header  *RunDocBrainmapResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunDocBrainmapResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunDocBrainmapResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunDocBrainmapResponseBody) GoString() string {
	return s.String()
}

func (s *RunDocBrainmapResponseBody) GetHeader() *RunDocBrainmapResponseBodyHeader {
	return s.Header
}

func (s *RunDocBrainmapResponseBody) GetPayload() *RunDocBrainmapResponseBodyPayload {
	return s.Payload
}

func (s *RunDocBrainmapResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunDocBrainmapResponseBody) SetHeader(v *RunDocBrainmapResponseBodyHeader) *RunDocBrainmapResponseBody {
	s.Header = v
	return s
}

func (s *RunDocBrainmapResponseBody) SetPayload(v *RunDocBrainmapResponseBodyPayload) *RunDocBrainmapResponseBody {
	s.Payload = v
	return s
}

func (s *RunDocBrainmapResponseBody) SetRequestId(v string) *RunDocBrainmapResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunDocBrainmapResponseBody) Validate() error {
	return dara.Validate(s)
}

type RunDocBrainmapResponseBodyHeader struct {
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
	// task-failed
	Event     *string `json:"Event,omitempty" xml:"Event,omitempty"`
	EventInfo *string `json:"EventInfo,omitempty" xml:"EventInfo,omitempty"`
	// example:
	//
	// 07181f55-2311-48af-8048-132a77dee020
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 161816
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 6427cdf4-2ffe-4d05-b0ef-c4adceea90f4
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunDocBrainmapResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunDocBrainmapResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunDocBrainmapResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunDocBrainmapResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunDocBrainmapResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunDocBrainmapResponseBodyHeader) GetEventInfo() *string {
	return s.EventInfo
}

func (s *RunDocBrainmapResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunDocBrainmapResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunDocBrainmapResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunDocBrainmapResponseBodyHeader) SetErrorCode(v string) *RunDocBrainmapResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunDocBrainmapResponseBodyHeader) SetErrorMessage(v string) *RunDocBrainmapResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunDocBrainmapResponseBodyHeader) SetEvent(v string) *RunDocBrainmapResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunDocBrainmapResponseBodyHeader) SetEventInfo(v string) *RunDocBrainmapResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunDocBrainmapResponseBodyHeader) SetSessionId(v string) *RunDocBrainmapResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunDocBrainmapResponseBodyHeader) SetTaskId(v string) *RunDocBrainmapResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunDocBrainmapResponseBodyHeader) SetTraceId(v string) *RunDocBrainmapResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunDocBrainmapResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunDocBrainmapResponseBodyPayload struct {
	Output *RunDocBrainmapResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunDocBrainmapResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunDocBrainmapResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunDocBrainmapResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunDocBrainmapResponseBodyPayload) GetOutput() *RunDocBrainmapResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunDocBrainmapResponseBodyPayload) GetUsage() *RunDocBrainmapResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunDocBrainmapResponseBodyPayload) SetOutput(v *RunDocBrainmapResponseBodyPayloadOutput) *RunDocBrainmapResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunDocBrainmapResponseBodyPayload) SetUsage(v *RunDocBrainmapResponseBodyPayloadUsage) *RunDocBrainmapResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunDocBrainmapResponseBodyPayload) Validate() error {
	return dara.Validate(s)
}

type RunDocBrainmapResponseBodyPayloadOutput struct {
	// example:
	//
	// {"xxxx":"xxx"}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s RunDocBrainmapResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunDocBrainmapResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunDocBrainmapResponseBodyPayloadOutput) GetContent() *string {
	return s.Content
}

func (s *RunDocBrainmapResponseBodyPayloadOutput) SetContent(v string) *RunDocBrainmapResponseBodyPayloadOutput {
	s.Content = &v
	return s
}

func (s *RunDocBrainmapResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunDocBrainmapResponseBodyPayloadUsage struct {
	// example:
	//
	// 65
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 100
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// example:
	//
	// 165
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunDocBrainmapResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunDocBrainmapResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunDocBrainmapResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunDocBrainmapResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunDocBrainmapResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunDocBrainmapResponseBodyPayloadUsage) SetInputTokens(v int64) *RunDocBrainmapResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunDocBrainmapResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunDocBrainmapResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunDocBrainmapResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunDocBrainmapResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunDocBrainmapResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
