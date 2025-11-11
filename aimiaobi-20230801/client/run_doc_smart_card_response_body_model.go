// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunDocSmartCardResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeader(v *RunDocSmartCardResponseBodyHeader) *RunDocSmartCardResponseBody
	GetHeader() *RunDocSmartCardResponseBodyHeader
	SetPayload(v *RunDocSmartCardResponseBodyPayload) *RunDocSmartCardResponseBody
	GetPayload() *RunDocSmartCardResponseBodyPayload
	SetRequestId(v string) *RunDocSmartCardResponseBody
	GetRequestId() *string
}

type RunDocSmartCardResponseBody struct {
	Header  *RunDocSmartCardResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunDocSmartCardResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunDocSmartCardResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunDocSmartCardResponseBody) GoString() string {
	return s.String()
}

func (s *RunDocSmartCardResponseBody) GetHeader() *RunDocSmartCardResponseBodyHeader {
	return s.Header
}

func (s *RunDocSmartCardResponseBody) GetPayload() *RunDocSmartCardResponseBodyPayload {
	return s.Payload
}

func (s *RunDocSmartCardResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunDocSmartCardResponseBody) SetHeader(v *RunDocSmartCardResponseBodyHeader) *RunDocSmartCardResponseBody {
	s.Header = v
	return s
}

func (s *RunDocSmartCardResponseBody) SetPayload(v *RunDocSmartCardResponseBodyPayload) *RunDocSmartCardResponseBody {
	s.Payload = v
	return s
}

func (s *RunDocSmartCardResponseBody) SetRequestId(v string) *RunDocSmartCardResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunDocSmartCardResponseBody) Validate() error {
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

type RunDocSmartCardResponseBodyHeader struct {
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
	// 07181f55-2311-48af-8048-132a77dee020
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 8d55b429d7c6d321fcff54823e8d317b
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 0abb781c17337107444473701ed7c3
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunDocSmartCardResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunDocSmartCardResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunDocSmartCardResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunDocSmartCardResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunDocSmartCardResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunDocSmartCardResponseBodyHeader) GetEventInfo() *string {
	return s.EventInfo
}

func (s *RunDocSmartCardResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunDocSmartCardResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunDocSmartCardResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunDocSmartCardResponseBodyHeader) SetErrorCode(v string) *RunDocSmartCardResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunDocSmartCardResponseBodyHeader) SetErrorMessage(v string) *RunDocSmartCardResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunDocSmartCardResponseBodyHeader) SetEvent(v string) *RunDocSmartCardResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunDocSmartCardResponseBodyHeader) SetEventInfo(v string) *RunDocSmartCardResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunDocSmartCardResponseBodyHeader) SetSessionId(v string) *RunDocSmartCardResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunDocSmartCardResponseBodyHeader) SetTaskId(v string) *RunDocSmartCardResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunDocSmartCardResponseBodyHeader) SetTraceId(v string) *RunDocSmartCardResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunDocSmartCardResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunDocSmartCardResponseBodyPayload struct {
	Output *RunDocSmartCardResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunDocSmartCardResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunDocSmartCardResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunDocSmartCardResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunDocSmartCardResponseBodyPayload) GetOutput() *RunDocSmartCardResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunDocSmartCardResponseBodyPayload) GetUsage() *RunDocSmartCardResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunDocSmartCardResponseBodyPayload) SetOutput(v *RunDocSmartCardResponseBodyPayloadOutput) *RunDocSmartCardResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunDocSmartCardResponseBodyPayload) SetUsage(v *RunDocSmartCardResponseBodyPayloadUsage) *RunDocSmartCardResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunDocSmartCardResponseBodyPayload) Validate() error {
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

type RunDocSmartCardResponseBodyPayloadOutput struct {
	Content *string   `json:"Content,omitempty" xml:"Content,omitempty"`
	Tags    []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s RunDocSmartCardResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunDocSmartCardResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunDocSmartCardResponseBodyPayloadOutput) GetContent() *string {
	return s.Content
}

func (s *RunDocSmartCardResponseBodyPayloadOutput) GetTags() []*string {
	return s.Tags
}

func (s *RunDocSmartCardResponseBodyPayloadOutput) SetContent(v string) *RunDocSmartCardResponseBodyPayloadOutput {
	s.Content = &v
	return s
}

func (s *RunDocSmartCardResponseBodyPayloadOutput) SetTags(v []*string) *RunDocSmartCardResponseBodyPayloadOutput {
	s.Tags = v
	return s
}

func (s *RunDocSmartCardResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunDocSmartCardResponseBodyPayloadUsage struct {
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

func (s RunDocSmartCardResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunDocSmartCardResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunDocSmartCardResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunDocSmartCardResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunDocSmartCardResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunDocSmartCardResponseBodyPayloadUsage) SetInputTokens(v int64) *RunDocSmartCardResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunDocSmartCardResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunDocSmartCardResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunDocSmartCardResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunDocSmartCardResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunDocSmartCardResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
