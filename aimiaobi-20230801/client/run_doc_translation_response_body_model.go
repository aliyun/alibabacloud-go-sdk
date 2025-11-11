// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunDocTranslationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeader(v *RunDocTranslationResponseBodyHeader) *RunDocTranslationResponseBody
	GetHeader() *RunDocTranslationResponseBodyHeader
	SetPayload(v *RunDocTranslationResponseBodyPayload) *RunDocTranslationResponseBody
	GetPayload() *RunDocTranslationResponseBodyPayload
	SetRequestId(v string) *RunDocTranslationResponseBody
	GetRequestId() *string
}

type RunDocTranslationResponseBody struct {
	Header  *RunDocTranslationResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunDocTranslationResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunDocTranslationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunDocTranslationResponseBody) GoString() string {
	return s.String()
}

func (s *RunDocTranslationResponseBody) GetHeader() *RunDocTranslationResponseBodyHeader {
	return s.Header
}

func (s *RunDocTranslationResponseBody) GetPayload() *RunDocTranslationResponseBodyPayload {
	return s.Payload
}

func (s *RunDocTranslationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunDocTranslationResponseBody) SetHeader(v *RunDocTranslationResponseBodyHeader) *RunDocTranslationResponseBody {
	s.Header = v
	return s
}

func (s *RunDocTranslationResponseBody) SetPayload(v *RunDocTranslationResponseBodyPayload) *RunDocTranslationResponseBody {
	s.Payload = v
	return s
}

func (s *RunDocTranslationResponseBody) SetRequestId(v string) *RunDocTranslationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunDocTranslationResponseBody) Validate() error {
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

type RunDocTranslationResponseBodyHeader struct {
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
	// 411c4dfa-2168-4379-a902-675d67f453f8
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 50a1cc8e-717e-4a2b-a76b-dc9734a8564b
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// ebd19b12-0cae-488f-9e41-5a1c825f545b
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunDocTranslationResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunDocTranslationResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunDocTranslationResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunDocTranslationResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunDocTranslationResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunDocTranslationResponseBodyHeader) GetEventInfo() *string {
	return s.EventInfo
}

func (s *RunDocTranslationResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunDocTranslationResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunDocTranslationResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunDocTranslationResponseBodyHeader) SetErrorCode(v string) *RunDocTranslationResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunDocTranslationResponseBodyHeader) SetErrorMessage(v string) *RunDocTranslationResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunDocTranslationResponseBodyHeader) SetEvent(v string) *RunDocTranslationResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunDocTranslationResponseBodyHeader) SetEventInfo(v string) *RunDocTranslationResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunDocTranslationResponseBodyHeader) SetSessionId(v string) *RunDocTranslationResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunDocTranslationResponseBodyHeader) SetTaskId(v string) *RunDocTranslationResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunDocTranslationResponseBodyHeader) SetTraceId(v string) *RunDocTranslationResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunDocTranslationResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunDocTranslationResponseBodyPayload struct {
	Output *RunDocTranslationResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunDocTranslationResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunDocTranslationResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunDocTranslationResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunDocTranslationResponseBodyPayload) GetOutput() *RunDocTranslationResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunDocTranslationResponseBodyPayload) GetUsage() *RunDocTranslationResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunDocTranslationResponseBodyPayload) SetOutput(v *RunDocTranslationResponseBodyPayloadOutput) *RunDocTranslationResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunDocTranslationResponseBodyPayload) SetUsage(v *RunDocTranslationResponseBodyPayloadUsage) *RunDocTranslationResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunDocTranslationResponseBodyPayload) Validate() error {
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

type RunDocTranslationResponseBodyPayloadOutput struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s RunDocTranslationResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunDocTranslationResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunDocTranslationResponseBodyPayloadOutput) GetContent() *string {
	return s.Content
}

func (s *RunDocTranslationResponseBodyPayloadOutput) SetContent(v string) *RunDocTranslationResponseBodyPayloadOutput {
	s.Content = &v
	return s
}

func (s *RunDocTranslationResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunDocTranslationResponseBodyPayloadUsage struct {
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

func (s RunDocTranslationResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunDocTranslationResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunDocTranslationResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunDocTranslationResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunDocTranslationResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunDocTranslationResponseBodyPayloadUsage) SetInputTokens(v int64) *RunDocTranslationResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunDocTranslationResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunDocTranslationResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunDocTranslationResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunDocTranslationResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunDocTranslationResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
