// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunGenerateQuestionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeader(v *RunGenerateQuestionsResponseBodyHeader) *RunGenerateQuestionsResponseBody
	GetHeader() *RunGenerateQuestionsResponseBodyHeader
	SetPayload(v *RunGenerateQuestionsResponseBodyPayload) *RunGenerateQuestionsResponseBody
	GetPayload() *RunGenerateQuestionsResponseBodyPayload
	SetRequestId(v string) *RunGenerateQuestionsResponseBody
	GetRequestId() *string
}

type RunGenerateQuestionsResponseBody struct {
	Header  *RunGenerateQuestionsResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunGenerateQuestionsResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunGenerateQuestionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunGenerateQuestionsResponseBody) GoString() string {
	return s.String()
}

func (s *RunGenerateQuestionsResponseBody) GetHeader() *RunGenerateQuestionsResponseBodyHeader {
	return s.Header
}

func (s *RunGenerateQuestionsResponseBody) GetPayload() *RunGenerateQuestionsResponseBodyPayload {
	return s.Payload
}

func (s *RunGenerateQuestionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunGenerateQuestionsResponseBody) SetHeader(v *RunGenerateQuestionsResponseBodyHeader) *RunGenerateQuestionsResponseBody {
	s.Header = v
	return s
}

func (s *RunGenerateQuestionsResponseBody) SetPayload(v *RunGenerateQuestionsResponseBodyPayload) *RunGenerateQuestionsResponseBody {
	s.Payload = v
	return s
}

func (s *RunGenerateQuestionsResponseBody) SetRequestId(v string) *RunGenerateQuestionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunGenerateQuestionsResponseBody) Validate() error {
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

type RunGenerateQuestionsResponseBodyHeader struct {
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
	// 3cd10828-0e42-471c-8f1a-931cde20b035
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// d3be9981-ca2d-4e17-bf31-1c0a628e9f99
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 0bc3b4b417362160345997589e5f6e
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunGenerateQuestionsResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunGenerateQuestionsResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunGenerateQuestionsResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunGenerateQuestionsResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunGenerateQuestionsResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunGenerateQuestionsResponseBodyHeader) GetEventInfo() *string {
	return s.EventInfo
}

func (s *RunGenerateQuestionsResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunGenerateQuestionsResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunGenerateQuestionsResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunGenerateQuestionsResponseBodyHeader) SetErrorCode(v string) *RunGenerateQuestionsResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunGenerateQuestionsResponseBodyHeader) SetErrorMessage(v string) *RunGenerateQuestionsResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunGenerateQuestionsResponseBodyHeader) SetEvent(v string) *RunGenerateQuestionsResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunGenerateQuestionsResponseBodyHeader) SetEventInfo(v string) *RunGenerateQuestionsResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunGenerateQuestionsResponseBodyHeader) SetSessionId(v string) *RunGenerateQuestionsResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunGenerateQuestionsResponseBodyHeader) SetTaskId(v string) *RunGenerateQuestionsResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunGenerateQuestionsResponseBodyHeader) SetTraceId(v string) *RunGenerateQuestionsResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunGenerateQuestionsResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunGenerateQuestionsResponseBodyPayload struct {
	Output *RunGenerateQuestionsResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunGenerateQuestionsResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunGenerateQuestionsResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunGenerateQuestionsResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunGenerateQuestionsResponseBodyPayload) GetOutput() *RunGenerateQuestionsResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunGenerateQuestionsResponseBodyPayload) GetUsage() *RunGenerateQuestionsResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunGenerateQuestionsResponseBodyPayload) SetOutput(v *RunGenerateQuestionsResponseBodyPayloadOutput) *RunGenerateQuestionsResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunGenerateQuestionsResponseBodyPayload) SetUsage(v *RunGenerateQuestionsResponseBodyPayloadUsage) *RunGenerateQuestionsResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunGenerateQuestionsResponseBodyPayload) Validate() error {
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

type RunGenerateQuestionsResponseBodyPayloadOutput struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s RunGenerateQuestionsResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunGenerateQuestionsResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunGenerateQuestionsResponseBodyPayloadOutput) GetContent() *string {
	return s.Content
}

func (s *RunGenerateQuestionsResponseBodyPayloadOutput) SetContent(v string) *RunGenerateQuestionsResponseBodyPayloadOutput {
	s.Content = &v
	return s
}

func (s *RunGenerateQuestionsResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunGenerateQuestionsResponseBodyPayloadUsage struct {
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

func (s RunGenerateQuestionsResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunGenerateQuestionsResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunGenerateQuestionsResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunGenerateQuestionsResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunGenerateQuestionsResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunGenerateQuestionsResponseBodyPayloadUsage) SetInputTokens(v int64) *RunGenerateQuestionsResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunGenerateQuestionsResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunGenerateQuestionsResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunGenerateQuestionsResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunGenerateQuestionsResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunGenerateQuestionsResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
