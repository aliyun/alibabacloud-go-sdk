// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunBookSmartCardResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeader(v *RunBookSmartCardResponseBodyHeader) *RunBookSmartCardResponseBody
	GetHeader() *RunBookSmartCardResponseBodyHeader
	SetPayload(v *RunBookSmartCardResponseBodyPayload) *RunBookSmartCardResponseBody
	GetPayload() *RunBookSmartCardResponseBodyPayload
	SetRequestId(v string) *RunBookSmartCardResponseBody
	GetRequestId() *string
}

type RunBookSmartCardResponseBody struct {
	Header  *RunBookSmartCardResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunBookSmartCardResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunBookSmartCardResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunBookSmartCardResponseBody) GoString() string {
	return s.String()
}

func (s *RunBookSmartCardResponseBody) GetHeader() *RunBookSmartCardResponseBodyHeader {
	return s.Header
}

func (s *RunBookSmartCardResponseBody) GetPayload() *RunBookSmartCardResponseBodyPayload {
	return s.Payload
}

func (s *RunBookSmartCardResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunBookSmartCardResponseBody) SetHeader(v *RunBookSmartCardResponseBodyHeader) *RunBookSmartCardResponseBody {
	s.Header = v
	return s
}

func (s *RunBookSmartCardResponseBody) SetPayload(v *RunBookSmartCardResponseBodyPayload) *RunBookSmartCardResponseBody {
	s.Payload = v
	return s
}

func (s *RunBookSmartCardResponseBody) SetRequestId(v string) *RunBookSmartCardResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunBookSmartCardResponseBody) Validate() error {
	if s.Header != nil {
		if err := s.Header.Validate(); err != nil {
			return err
		}
	}
	if s.Payload != nil {
		if err := s.Payload.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunBookSmartCardResponseBodyHeader struct {
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
	// finished
	Event     *string `json:"Event,omitempty" xml:"Event,omitempty"`
	EventInfo *string `json:"EventInfo,omitempty" xml:"EventInfo,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 1a0e898717105546647125853d4f54
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunBookSmartCardResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunBookSmartCardResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunBookSmartCardResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunBookSmartCardResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunBookSmartCardResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunBookSmartCardResponseBodyHeader) GetEventInfo() *string {
	return s.EventInfo
}

func (s *RunBookSmartCardResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunBookSmartCardResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunBookSmartCardResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunBookSmartCardResponseBodyHeader) SetErrorCode(v string) *RunBookSmartCardResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunBookSmartCardResponseBodyHeader) SetErrorMessage(v string) *RunBookSmartCardResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunBookSmartCardResponseBodyHeader) SetEvent(v string) *RunBookSmartCardResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunBookSmartCardResponseBodyHeader) SetEventInfo(v string) *RunBookSmartCardResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunBookSmartCardResponseBodyHeader) SetSessionId(v string) *RunBookSmartCardResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunBookSmartCardResponseBodyHeader) SetTaskId(v string) *RunBookSmartCardResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunBookSmartCardResponseBodyHeader) SetTraceId(v string) *RunBookSmartCardResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunBookSmartCardResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunBookSmartCardResponseBodyPayload struct {
	Output *RunBookSmartCardResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunBookSmartCardResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunBookSmartCardResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunBookSmartCardResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunBookSmartCardResponseBodyPayload) GetOutput() *RunBookSmartCardResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunBookSmartCardResponseBodyPayload) GetUsage() *RunBookSmartCardResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunBookSmartCardResponseBodyPayload) SetOutput(v *RunBookSmartCardResponseBodyPayloadOutput) *RunBookSmartCardResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunBookSmartCardResponseBodyPayload) SetUsage(v *RunBookSmartCardResponseBodyPayloadUsage) *RunBookSmartCardResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunBookSmartCardResponseBodyPayload) Validate() error {
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
			return err
		}
	}
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunBookSmartCardResponseBodyPayloadOutput struct {
	Content *string   `json:"Content,omitempty" xml:"Content,omitempty"`
	Tags    []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s RunBookSmartCardResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunBookSmartCardResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunBookSmartCardResponseBodyPayloadOutput) GetContent() *string {
	return s.Content
}

func (s *RunBookSmartCardResponseBodyPayloadOutput) GetTags() []*string {
	return s.Tags
}

func (s *RunBookSmartCardResponseBodyPayloadOutput) SetContent(v string) *RunBookSmartCardResponseBodyPayloadOutput {
	s.Content = &v
	return s
}

func (s *RunBookSmartCardResponseBodyPayloadOutput) SetTags(v []*string) *RunBookSmartCardResponseBodyPayloadOutput {
	s.Tags = v
	return s
}

func (s *RunBookSmartCardResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunBookSmartCardResponseBodyPayloadUsage struct {
	// example:
	//
	// 1
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 100
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// example:
	//
	// 101
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunBookSmartCardResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunBookSmartCardResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunBookSmartCardResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunBookSmartCardResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunBookSmartCardResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunBookSmartCardResponseBodyPayloadUsage) SetInputTokens(v int64) *RunBookSmartCardResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunBookSmartCardResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunBookSmartCardResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunBookSmartCardResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunBookSmartCardResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunBookSmartCardResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
