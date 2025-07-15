// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunBookBrainmapResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeader(v *RunBookBrainmapResponseBodyHeader) *RunBookBrainmapResponseBody
	GetHeader() *RunBookBrainmapResponseBodyHeader
	SetPayload(v *RunBookBrainmapResponseBodyPayload) *RunBookBrainmapResponseBody
	GetPayload() *RunBookBrainmapResponseBodyPayload
	SetRequestId(v string) *RunBookBrainmapResponseBody
	GetRequestId() *string
}

type RunBookBrainmapResponseBody struct {
	Header  *RunBookBrainmapResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunBookBrainmapResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunBookBrainmapResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunBookBrainmapResponseBody) GoString() string {
	return s.String()
}

func (s *RunBookBrainmapResponseBody) GetHeader() *RunBookBrainmapResponseBodyHeader {
	return s.Header
}

func (s *RunBookBrainmapResponseBody) GetPayload() *RunBookBrainmapResponseBodyPayload {
	return s.Payload
}

func (s *RunBookBrainmapResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunBookBrainmapResponseBody) SetHeader(v *RunBookBrainmapResponseBodyHeader) *RunBookBrainmapResponseBody {
	s.Header = v
	return s
}

func (s *RunBookBrainmapResponseBody) SetPayload(v *RunBookBrainmapResponseBodyPayload) *RunBookBrainmapResponseBody {
	s.Payload = v
	return s
}

func (s *RunBookBrainmapResponseBody) SetRequestId(v string) *RunBookBrainmapResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunBookBrainmapResponseBody) Validate() error {
	return dara.Validate(s)
}

type RunBookBrainmapResponseBodyHeader struct {
	// example:
	//
	// Success
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
	// 3cd10828-0e42-471c-8f1a-931cde20b035
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 0bc1409b17210096103458421ec62e
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunBookBrainmapResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunBookBrainmapResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunBookBrainmapResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunBookBrainmapResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunBookBrainmapResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunBookBrainmapResponseBodyHeader) GetEventInfo() *string {
	return s.EventInfo
}

func (s *RunBookBrainmapResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunBookBrainmapResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunBookBrainmapResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunBookBrainmapResponseBodyHeader) SetErrorCode(v string) *RunBookBrainmapResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunBookBrainmapResponseBodyHeader) SetErrorMessage(v string) *RunBookBrainmapResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunBookBrainmapResponseBodyHeader) SetEvent(v string) *RunBookBrainmapResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunBookBrainmapResponseBodyHeader) SetEventInfo(v string) *RunBookBrainmapResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunBookBrainmapResponseBodyHeader) SetSessionId(v string) *RunBookBrainmapResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunBookBrainmapResponseBodyHeader) SetTaskId(v string) *RunBookBrainmapResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunBookBrainmapResponseBodyHeader) SetTraceId(v string) *RunBookBrainmapResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunBookBrainmapResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunBookBrainmapResponseBodyPayload struct {
	Output *RunBookBrainmapResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunBookBrainmapResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunBookBrainmapResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunBookBrainmapResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunBookBrainmapResponseBodyPayload) GetOutput() *RunBookBrainmapResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunBookBrainmapResponseBodyPayload) GetUsage() *RunBookBrainmapResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunBookBrainmapResponseBodyPayload) SetOutput(v *RunBookBrainmapResponseBodyPayloadOutput) *RunBookBrainmapResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunBookBrainmapResponseBodyPayload) SetUsage(v *RunBookBrainmapResponseBodyPayloadUsage) *RunBookBrainmapResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunBookBrainmapResponseBodyPayload) Validate() error {
	return dara.Validate(s)
}

type RunBookBrainmapResponseBodyPayloadOutput struct {
	// example:
	//
	// {"xxxx":"xxx"}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s RunBookBrainmapResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunBookBrainmapResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunBookBrainmapResponseBodyPayloadOutput) GetContent() *string {
	return s.Content
}

func (s *RunBookBrainmapResponseBodyPayloadOutput) SetContent(v string) *RunBookBrainmapResponseBodyPayloadOutput {
	s.Content = &v
	return s
}

func (s *RunBookBrainmapResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunBookBrainmapResponseBodyPayloadUsage struct {
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

func (s RunBookBrainmapResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunBookBrainmapResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunBookBrainmapResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunBookBrainmapResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunBookBrainmapResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunBookBrainmapResponseBodyPayloadUsage) SetInputTokens(v int64) *RunBookBrainmapResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunBookBrainmapResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunBookBrainmapResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunBookBrainmapResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunBookBrainmapResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunBookBrainmapResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
