// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunHotwordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeader(v *RunHotwordResponseBodyHeader) *RunHotwordResponseBody
	GetHeader() *RunHotwordResponseBodyHeader
	SetPayload(v *RunHotwordResponseBodyPayload) *RunHotwordResponseBody
	GetPayload() *RunHotwordResponseBodyPayload
	SetRequestId(v string) *RunHotwordResponseBody
	GetRequestId() *string
}

type RunHotwordResponseBody struct {
	Header  *RunHotwordResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunHotwordResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunHotwordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunHotwordResponseBody) GoString() string {
	return s.String()
}

func (s *RunHotwordResponseBody) GetHeader() *RunHotwordResponseBodyHeader {
	return s.Header
}

func (s *RunHotwordResponseBody) GetPayload() *RunHotwordResponseBodyPayload {
	return s.Payload
}

func (s *RunHotwordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunHotwordResponseBody) SetHeader(v *RunHotwordResponseBodyHeader) *RunHotwordResponseBody {
	s.Header = v
	return s
}

func (s *RunHotwordResponseBody) SetPayload(v *RunHotwordResponseBodyPayload) *RunHotwordResponseBody {
	s.Payload = v
	return s
}

func (s *RunHotwordResponseBody) SetRequestId(v string) *RunHotwordResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunHotwordResponseBody) Validate() error {
	return dara.Validate(s)
}

type RunHotwordResponseBodyHeader struct {
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// success
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// finished
	Event     *string `json:"Event,omitempty" xml:"Event,omitempty"`
	EventInfo *string `json:"EventInfo,omitempty" xml:"EventInfo,omitempty"`
	// example:
	//
	// 92e16ccb-92b6-4894-abbf-fc6e2929a0df
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 0abb7e3217356108993888059ee72b
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunHotwordResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunHotwordResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunHotwordResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunHotwordResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunHotwordResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunHotwordResponseBodyHeader) GetEventInfo() *string {
	return s.EventInfo
}

func (s *RunHotwordResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunHotwordResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunHotwordResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunHotwordResponseBodyHeader) SetErrorCode(v string) *RunHotwordResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunHotwordResponseBodyHeader) SetErrorMessage(v string) *RunHotwordResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunHotwordResponseBodyHeader) SetEvent(v string) *RunHotwordResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunHotwordResponseBodyHeader) SetEventInfo(v string) *RunHotwordResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunHotwordResponseBodyHeader) SetSessionId(v string) *RunHotwordResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunHotwordResponseBodyHeader) SetTaskId(v string) *RunHotwordResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunHotwordResponseBodyHeader) SetTraceId(v string) *RunHotwordResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunHotwordResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunHotwordResponseBodyPayload struct {
	Output *RunHotwordResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunHotwordResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunHotwordResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunHotwordResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunHotwordResponseBodyPayload) GetOutput() *RunHotwordResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunHotwordResponseBodyPayload) GetUsage() *RunHotwordResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunHotwordResponseBodyPayload) SetOutput(v *RunHotwordResponseBodyPayloadOutput) *RunHotwordResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunHotwordResponseBodyPayload) SetUsage(v *RunHotwordResponseBodyPayloadUsage) *RunHotwordResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunHotwordResponseBodyPayload) Validate() error {
	return dara.Validate(s)
}

type RunHotwordResponseBodyPayloadOutput struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s RunHotwordResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunHotwordResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunHotwordResponseBodyPayloadOutput) GetContent() *string {
	return s.Content
}

func (s *RunHotwordResponseBodyPayloadOutput) SetContent(v string) *RunHotwordResponseBodyPayloadOutput {
	s.Content = &v
	return s
}

func (s *RunHotwordResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunHotwordResponseBodyPayloadUsage struct {
	// example:
	//
	// 100
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 1
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// example:
	//
	// 101
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunHotwordResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunHotwordResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunHotwordResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunHotwordResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunHotwordResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunHotwordResponseBodyPayloadUsage) SetInputTokens(v int64) *RunHotwordResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunHotwordResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunHotwordResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunHotwordResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunHotwordResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunHotwordResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
