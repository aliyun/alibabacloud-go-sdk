// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunBookIntroductionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeader(v *RunBookIntroductionResponseBodyHeader) *RunBookIntroductionResponseBody
	GetHeader() *RunBookIntroductionResponseBodyHeader
	SetPayload(v *RunBookIntroductionResponseBodyPayload) *RunBookIntroductionResponseBody
	GetPayload() *RunBookIntroductionResponseBodyPayload
	SetRequestId(v string) *RunBookIntroductionResponseBody
	GetRequestId() *string
}

type RunBookIntroductionResponseBody struct {
	Header  *RunBookIntroductionResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunBookIntroductionResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunBookIntroductionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunBookIntroductionResponseBody) GoString() string {
	return s.String()
}

func (s *RunBookIntroductionResponseBody) GetHeader() *RunBookIntroductionResponseBodyHeader {
	return s.Header
}

func (s *RunBookIntroductionResponseBody) GetPayload() *RunBookIntroductionResponseBodyPayload {
	return s.Payload
}

func (s *RunBookIntroductionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunBookIntroductionResponseBody) SetHeader(v *RunBookIntroductionResponseBodyHeader) *RunBookIntroductionResponseBody {
	s.Header = v
	return s
}

func (s *RunBookIntroductionResponseBody) SetPayload(v *RunBookIntroductionResponseBodyPayload) *RunBookIntroductionResponseBody {
	s.Payload = v
	return s
}

func (s *RunBookIntroductionResponseBody) SetRequestId(v string) *RunBookIntroductionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunBookIntroductionResponseBody) Validate() error {
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

type RunBookIntroductionResponseBodyHeader struct {
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
	// 411c4dfa-2168-4379-a902-675d67f453f8
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// d3be9981-ca2d-4e17-bf31-1c0a628e9f99
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 46e5c2b5-0877-4f09-bd91-ab0cf314e48b
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunBookIntroductionResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunBookIntroductionResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunBookIntroductionResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunBookIntroductionResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunBookIntroductionResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunBookIntroductionResponseBodyHeader) GetEventInfo() *string {
	return s.EventInfo
}

func (s *RunBookIntroductionResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunBookIntroductionResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunBookIntroductionResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunBookIntroductionResponseBodyHeader) SetErrorCode(v string) *RunBookIntroductionResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunBookIntroductionResponseBodyHeader) SetErrorMessage(v string) *RunBookIntroductionResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunBookIntroductionResponseBodyHeader) SetEvent(v string) *RunBookIntroductionResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunBookIntroductionResponseBodyHeader) SetEventInfo(v string) *RunBookIntroductionResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunBookIntroductionResponseBodyHeader) SetSessionId(v string) *RunBookIntroductionResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunBookIntroductionResponseBodyHeader) SetTaskId(v string) *RunBookIntroductionResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunBookIntroductionResponseBodyHeader) SetTraceId(v string) *RunBookIntroductionResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunBookIntroductionResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunBookIntroductionResponseBodyPayload struct {
	Output *RunBookIntroductionResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunBookIntroductionResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunBookIntroductionResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunBookIntroductionResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunBookIntroductionResponseBodyPayload) GetOutput() *RunBookIntroductionResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunBookIntroductionResponseBodyPayload) GetUsage() *RunBookIntroductionResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunBookIntroductionResponseBodyPayload) SetOutput(v *RunBookIntroductionResponseBodyPayloadOutput) *RunBookIntroductionResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunBookIntroductionResponseBodyPayload) SetUsage(v *RunBookIntroductionResponseBodyPayloadUsage) *RunBookIntroductionResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunBookIntroductionResponseBodyPayload) Validate() error {
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

type RunBookIntroductionResponseBodyPayloadOutput struct {
	KeyPoint *string `json:"KeyPoint,omitempty" xml:"KeyPoint,omitempty"`
	Summary  *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s RunBookIntroductionResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunBookIntroductionResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunBookIntroductionResponseBodyPayloadOutput) GetKeyPoint() *string {
	return s.KeyPoint
}

func (s *RunBookIntroductionResponseBodyPayloadOutput) GetSummary() *string {
	return s.Summary
}

func (s *RunBookIntroductionResponseBodyPayloadOutput) SetKeyPoint(v string) *RunBookIntroductionResponseBodyPayloadOutput {
	s.KeyPoint = &v
	return s
}

func (s *RunBookIntroductionResponseBodyPayloadOutput) SetSummary(v string) *RunBookIntroductionResponseBodyPayloadOutput {
	s.Summary = &v
	return s
}

func (s *RunBookIntroductionResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunBookIntroductionResponseBodyPayloadUsage struct {
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

func (s RunBookIntroductionResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunBookIntroductionResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunBookIntroductionResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunBookIntroductionResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunBookIntroductionResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunBookIntroductionResponseBodyPayloadUsage) SetInputTokens(v int64) *RunBookIntroductionResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunBookIntroductionResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunBookIntroductionResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunBookIntroductionResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunBookIntroductionResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunBookIntroductionResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
