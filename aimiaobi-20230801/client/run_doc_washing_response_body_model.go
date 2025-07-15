// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunDocWashingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnd(v bool) *RunDocWashingResponseBody
	GetEnd() *bool
	SetHeader(v *RunDocWashingResponseBodyHeader) *RunDocWashingResponseBody
	GetHeader() *RunDocWashingResponseBodyHeader
	SetPayload(v *RunDocWashingResponseBodyPayload) *RunDocWashingResponseBody
	GetPayload() *RunDocWashingResponseBodyPayload
	SetRequestId(v string) *RunDocWashingResponseBody
	GetRequestId() *string
}

type RunDocWashingResponseBody struct {
	// example:
	//
	// false
	End     *bool                             `json:"End,omitempty" xml:"End,omitempty"`
	Header  *RunDocWashingResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunDocWashingResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunDocWashingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunDocWashingResponseBody) GoString() string {
	return s.String()
}

func (s *RunDocWashingResponseBody) GetEnd() *bool {
	return s.End
}

func (s *RunDocWashingResponseBody) GetHeader() *RunDocWashingResponseBodyHeader {
	return s.Header
}

func (s *RunDocWashingResponseBody) GetPayload() *RunDocWashingResponseBodyPayload {
	return s.Payload
}

func (s *RunDocWashingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunDocWashingResponseBody) SetEnd(v bool) *RunDocWashingResponseBody {
	s.End = &v
	return s
}

func (s *RunDocWashingResponseBody) SetHeader(v *RunDocWashingResponseBodyHeader) *RunDocWashingResponseBody {
	s.Header = v
	return s
}

func (s *RunDocWashingResponseBody) SetPayload(v *RunDocWashingResponseBodyPayload) *RunDocWashingResponseBody {
	s.Payload = v
	return s
}

func (s *RunDocWashingResponseBody) SetRequestId(v string) *RunDocWashingResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunDocWashingResponseBody) Validate() error {
	return dara.Validate(s)
}

type RunDocWashingResponseBodyHeader struct {
	// example:
	//
	// task-finished
	Event     *string `json:"Event,omitempty" xml:"Event,omitempty"`
	EventInfo *string `json:"EventInfo,omitempty" xml:"EventInfo,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20247a52-23e2-46fb-943d-309cdee2bc6d
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 2150451a17191950923411783e2927
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunDocWashingResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunDocWashingResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunDocWashingResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunDocWashingResponseBodyHeader) GetEventInfo() *string {
	return s.EventInfo
}

func (s *RunDocWashingResponseBodyHeader) GetRequestId() *string {
	return s.RequestId
}

func (s *RunDocWashingResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunDocWashingResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunDocWashingResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunDocWashingResponseBodyHeader) SetEvent(v string) *RunDocWashingResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunDocWashingResponseBodyHeader) SetEventInfo(v string) *RunDocWashingResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunDocWashingResponseBodyHeader) SetRequestId(v string) *RunDocWashingResponseBodyHeader {
	s.RequestId = &v
	return s
}

func (s *RunDocWashingResponseBodyHeader) SetSessionId(v string) *RunDocWashingResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunDocWashingResponseBodyHeader) SetTaskId(v string) *RunDocWashingResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunDocWashingResponseBodyHeader) SetTraceId(v string) *RunDocWashingResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunDocWashingResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunDocWashingResponseBodyPayload struct {
	Output *RunDocWashingResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunDocWashingResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunDocWashingResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunDocWashingResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunDocWashingResponseBodyPayload) GetOutput() *RunDocWashingResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunDocWashingResponseBodyPayload) GetUsage() *RunDocWashingResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunDocWashingResponseBodyPayload) SetOutput(v *RunDocWashingResponseBodyPayloadOutput) *RunDocWashingResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunDocWashingResponseBodyPayload) SetUsage(v *RunDocWashingResponseBodyPayloadUsage) *RunDocWashingResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunDocWashingResponseBodyPayload) Validate() error {
	return dara.Validate(s)
}

type RunDocWashingResponseBodyPayloadOutput struct {
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RunDocWashingResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunDocWashingResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunDocWashingResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunDocWashingResponseBodyPayloadOutput) SetText(v string) *RunDocWashingResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunDocWashingResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunDocWashingResponseBodyPayloadUsage struct {
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

func (s RunDocWashingResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunDocWashingResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunDocWashingResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunDocWashingResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunDocWashingResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunDocWashingResponseBodyPayloadUsage) SetInputTokens(v int64) *RunDocWashingResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunDocWashingResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunDocWashingResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunDocWashingResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunDocWashingResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunDocWashingResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
