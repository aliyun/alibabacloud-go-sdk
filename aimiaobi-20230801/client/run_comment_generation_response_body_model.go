// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCommentGenerationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnd(v bool) *RunCommentGenerationResponseBody
	GetEnd() *bool
	SetHeader(v *RunCommentGenerationResponseBodyHeader) *RunCommentGenerationResponseBody
	GetHeader() *RunCommentGenerationResponseBodyHeader
	SetPayload(v *RunCommentGenerationResponseBodyPayload) *RunCommentGenerationResponseBody
	GetPayload() *RunCommentGenerationResponseBodyPayload
	SetRequestId(v string) *RunCommentGenerationResponseBody
	GetRequestId() *string
}

type RunCommentGenerationResponseBody struct {
	End     *bool                                    `json:"End,omitempty" xml:"End,omitempty"`
	Header  *RunCommentGenerationResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunCommentGenerationResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunCommentGenerationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunCommentGenerationResponseBody) GoString() string {
	return s.String()
}

func (s *RunCommentGenerationResponseBody) GetEnd() *bool {
	return s.End
}

func (s *RunCommentGenerationResponseBody) GetHeader() *RunCommentGenerationResponseBodyHeader {
	return s.Header
}

func (s *RunCommentGenerationResponseBody) GetPayload() *RunCommentGenerationResponseBodyPayload {
	return s.Payload
}

func (s *RunCommentGenerationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunCommentGenerationResponseBody) SetEnd(v bool) *RunCommentGenerationResponseBody {
	s.End = &v
	return s
}

func (s *RunCommentGenerationResponseBody) SetHeader(v *RunCommentGenerationResponseBodyHeader) *RunCommentGenerationResponseBody {
	s.Header = v
	return s
}

func (s *RunCommentGenerationResponseBody) SetPayload(v *RunCommentGenerationResponseBodyPayload) *RunCommentGenerationResponseBody {
	s.Payload = v
	return s
}

func (s *RunCommentGenerationResponseBody) SetRequestId(v string) *RunCommentGenerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunCommentGenerationResponseBody) Validate() error {
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

type RunCommentGenerationResponseBodyHeader struct {
	// example:
	//
	// result-generated
	Event     *string `json:"Event,omitempty" xml:"Event,omitempty"`
	EventInfo *string `json:"EventInfo,omitempty" xml:"EventInfo,omitempty"`
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// 0bd58ea2-dc38-45da-ac02-17f05cb9040b
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunCommentGenerationResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunCommentGenerationResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunCommentGenerationResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunCommentGenerationResponseBodyHeader) GetEventInfo() *string {
	return s.EventInfo
}

func (s *RunCommentGenerationResponseBodyHeader) GetRequestId() *string {
	return s.RequestId
}

func (s *RunCommentGenerationResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunCommentGenerationResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunCommentGenerationResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunCommentGenerationResponseBodyHeader) SetEvent(v string) *RunCommentGenerationResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunCommentGenerationResponseBodyHeader) SetEventInfo(v string) *RunCommentGenerationResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunCommentGenerationResponseBodyHeader) SetRequestId(v string) *RunCommentGenerationResponseBodyHeader {
	s.RequestId = &v
	return s
}

func (s *RunCommentGenerationResponseBodyHeader) SetSessionId(v string) *RunCommentGenerationResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunCommentGenerationResponseBodyHeader) SetTaskId(v string) *RunCommentGenerationResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunCommentGenerationResponseBodyHeader) SetTraceId(v string) *RunCommentGenerationResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunCommentGenerationResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunCommentGenerationResponseBodyPayload struct {
	Output *RunCommentGenerationResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunCommentGenerationResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunCommentGenerationResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunCommentGenerationResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunCommentGenerationResponseBodyPayload) GetOutput() *RunCommentGenerationResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunCommentGenerationResponseBodyPayload) GetUsage() *RunCommentGenerationResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunCommentGenerationResponseBodyPayload) SetOutput(v *RunCommentGenerationResponseBodyPayloadOutput) *RunCommentGenerationResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunCommentGenerationResponseBodyPayload) SetUsage(v *RunCommentGenerationResponseBodyPayloadUsage) *RunCommentGenerationResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunCommentGenerationResponseBodyPayload) Validate() error {
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

type RunCommentGenerationResponseBodyPayloadOutput struct {
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RunCommentGenerationResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunCommentGenerationResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunCommentGenerationResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunCommentGenerationResponseBodyPayloadOutput) SetText(v string) *RunCommentGenerationResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunCommentGenerationResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunCommentGenerationResponseBodyPayloadUsage struct {
	// example:
	//
	// 1
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 2
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// example:
	//
	// 3
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunCommentGenerationResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunCommentGenerationResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunCommentGenerationResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunCommentGenerationResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunCommentGenerationResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunCommentGenerationResponseBodyPayloadUsage) SetInputTokens(v int64) *RunCommentGenerationResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunCommentGenerationResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunCommentGenerationResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunCommentGenerationResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunCommentGenerationResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunCommentGenerationResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
