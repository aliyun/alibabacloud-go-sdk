// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunKeywordsExtractionGenerationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeader(v *RunKeywordsExtractionGenerationResponseBodyHeader) *RunKeywordsExtractionGenerationResponseBody
	GetHeader() *RunKeywordsExtractionGenerationResponseBodyHeader
	SetPayload(v *RunKeywordsExtractionGenerationResponseBodyPayload) *RunKeywordsExtractionGenerationResponseBody
	GetPayload() *RunKeywordsExtractionGenerationResponseBodyPayload
	SetRequestId(v string) *RunKeywordsExtractionGenerationResponseBody
	GetRequestId() *string
}

type RunKeywordsExtractionGenerationResponseBody struct {
	Header  *RunKeywordsExtractionGenerationResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunKeywordsExtractionGenerationResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// 419F3FBE-5C8D-5949-AC29-E9615235D15A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunKeywordsExtractionGenerationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunKeywordsExtractionGenerationResponseBody) GoString() string {
	return s.String()
}

func (s *RunKeywordsExtractionGenerationResponseBody) GetHeader() *RunKeywordsExtractionGenerationResponseBodyHeader {
	return s.Header
}

func (s *RunKeywordsExtractionGenerationResponseBody) GetPayload() *RunKeywordsExtractionGenerationResponseBodyPayload {
	return s.Payload
}

func (s *RunKeywordsExtractionGenerationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunKeywordsExtractionGenerationResponseBody) SetHeader(v *RunKeywordsExtractionGenerationResponseBodyHeader) *RunKeywordsExtractionGenerationResponseBody {
	s.Header = v
	return s
}

func (s *RunKeywordsExtractionGenerationResponseBody) SetPayload(v *RunKeywordsExtractionGenerationResponseBodyPayload) *RunKeywordsExtractionGenerationResponseBody {
	s.Payload = v
	return s
}

func (s *RunKeywordsExtractionGenerationResponseBody) SetRequestId(v string) *RunKeywordsExtractionGenerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunKeywordsExtractionGenerationResponseBody) Validate() error {
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

type RunKeywordsExtractionGenerationResponseBodyHeader struct {
	// example:
	//
	// AccessForbid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// xx
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// task-failed
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// 1a3d7c9f-3a6d-4e49-b176-2d8721a27397
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 8d55b429d7c6d321fcff54823e8d317b
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 210bc4e817219607963985396de8bd
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunKeywordsExtractionGenerationResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunKeywordsExtractionGenerationResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunKeywordsExtractionGenerationResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunKeywordsExtractionGenerationResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunKeywordsExtractionGenerationResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunKeywordsExtractionGenerationResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunKeywordsExtractionGenerationResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunKeywordsExtractionGenerationResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunKeywordsExtractionGenerationResponseBodyHeader) SetErrorCode(v string) *RunKeywordsExtractionGenerationResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunKeywordsExtractionGenerationResponseBodyHeader) SetErrorMessage(v string) *RunKeywordsExtractionGenerationResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunKeywordsExtractionGenerationResponseBodyHeader) SetEvent(v string) *RunKeywordsExtractionGenerationResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunKeywordsExtractionGenerationResponseBodyHeader) SetSessionId(v string) *RunKeywordsExtractionGenerationResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunKeywordsExtractionGenerationResponseBodyHeader) SetTaskId(v string) *RunKeywordsExtractionGenerationResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunKeywordsExtractionGenerationResponseBodyHeader) SetTraceId(v string) *RunKeywordsExtractionGenerationResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunKeywordsExtractionGenerationResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunKeywordsExtractionGenerationResponseBodyPayload struct {
	Output *RunKeywordsExtractionGenerationResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunKeywordsExtractionGenerationResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunKeywordsExtractionGenerationResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunKeywordsExtractionGenerationResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunKeywordsExtractionGenerationResponseBodyPayload) GetOutput() *RunKeywordsExtractionGenerationResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunKeywordsExtractionGenerationResponseBodyPayload) GetUsage() *RunKeywordsExtractionGenerationResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunKeywordsExtractionGenerationResponseBodyPayload) SetOutput(v *RunKeywordsExtractionGenerationResponseBodyPayloadOutput) *RunKeywordsExtractionGenerationResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunKeywordsExtractionGenerationResponseBodyPayload) SetUsage(v *RunKeywordsExtractionGenerationResponseBodyPayloadUsage) *RunKeywordsExtractionGenerationResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunKeywordsExtractionGenerationResponseBodyPayload) Validate() error {
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

type RunKeywordsExtractionGenerationResponseBodyPayloadOutput struct {
	// example:
	//
	// xxx
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RunKeywordsExtractionGenerationResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunKeywordsExtractionGenerationResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunKeywordsExtractionGenerationResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunKeywordsExtractionGenerationResponseBodyPayloadOutput) SetText(v string) *RunKeywordsExtractionGenerationResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunKeywordsExtractionGenerationResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunKeywordsExtractionGenerationResponseBodyPayloadUsage struct {
	// example:
	//
	// 1
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 1
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// example:
	//
	// 2
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunKeywordsExtractionGenerationResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunKeywordsExtractionGenerationResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunKeywordsExtractionGenerationResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunKeywordsExtractionGenerationResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunKeywordsExtractionGenerationResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunKeywordsExtractionGenerationResponseBodyPayloadUsage) SetInputTokens(v int64) *RunKeywordsExtractionGenerationResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunKeywordsExtractionGenerationResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunKeywordsExtractionGenerationResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunKeywordsExtractionGenerationResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunKeywordsExtractionGenerationResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunKeywordsExtractionGenerationResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
