// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunTranslateGenerationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeader(v *RunTranslateGenerationResponseBodyHeader) *RunTranslateGenerationResponseBody
	GetHeader() *RunTranslateGenerationResponseBodyHeader
	SetPayload(v *RunTranslateGenerationResponseBodyPayload) *RunTranslateGenerationResponseBody
	GetPayload() *RunTranslateGenerationResponseBodyPayload
	SetRequestId(v string) *RunTranslateGenerationResponseBody
	GetRequestId() *string
}

type RunTranslateGenerationResponseBody struct {
	Header  *RunTranslateGenerationResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunTranslateGenerationResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// DA021073-17CE-5CCF-9FEB-93226C766887
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunTranslateGenerationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunTranslateGenerationResponseBody) GoString() string {
	return s.String()
}

func (s *RunTranslateGenerationResponseBody) GetHeader() *RunTranslateGenerationResponseBodyHeader {
	return s.Header
}

func (s *RunTranslateGenerationResponseBody) GetPayload() *RunTranslateGenerationResponseBodyPayload {
	return s.Payload
}

func (s *RunTranslateGenerationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunTranslateGenerationResponseBody) SetHeader(v *RunTranslateGenerationResponseBodyHeader) *RunTranslateGenerationResponseBody {
	s.Header = v
	return s
}

func (s *RunTranslateGenerationResponseBody) SetPayload(v *RunTranslateGenerationResponseBodyPayload) *RunTranslateGenerationResponseBody {
	s.Payload = v
	return s
}

func (s *RunTranslateGenerationResponseBody) SetRequestId(v string) *RunTranslateGenerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunTranslateGenerationResponseBody) Validate() error {
	return dara.Validate(s)
}

type RunTranslateGenerationResponseBodyHeader struct {
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
	// 91C2B2B8-7D12-4A8D-A724-1E576D30C096
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 0abb781d17146157564845243e20b5
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunTranslateGenerationResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunTranslateGenerationResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunTranslateGenerationResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunTranslateGenerationResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunTranslateGenerationResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunTranslateGenerationResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunTranslateGenerationResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunTranslateGenerationResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunTranslateGenerationResponseBodyHeader) SetErrorCode(v string) *RunTranslateGenerationResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunTranslateGenerationResponseBodyHeader) SetErrorMessage(v string) *RunTranslateGenerationResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunTranslateGenerationResponseBodyHeader) SetEvent(v string) *RunTranslateGenerationResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunTranslateGenerationResponseBodyHeader) SetSessionId(v string) *RunTranslateGenerationResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunTranslateGenerationResponseBodyHeader) SetTaskId(v string) *RunTranslateGenerationResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunTranslateGenerationResponseBodyHeader) SetTraceId(v string) *RunTranslateGenerationResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunTranslateGenerationResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunTranslateGenerationResponseBodyPayload struct {
	Output *RunTranslateGenerationResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunTranslateGenerationResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunTranslateGenerationResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunTranslateGenerationResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunTranslateGenerationResponseBodyPayload) GetOutput() *RunTranslateGenerationResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunTranslateGenerationResponseBodyPayload) GetUsage() *RunTranslateGenerationResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunTranslateGenerationResponseBodyPayload) SetOutput(v *RunTranslateGenerationResponseBodyPayloadOutput) *RunTranslateGenerationResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunTranslateGenerationResponseBodyPayload) SetUsage(v *RunTranslateGenerationResponseBodyPayloadUsage) *RunTranslateGenerationResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunTranslateGenerationResponseBodyPayload) Validate() error {
	return dara.Validate(s)
}

type RunTranslateGenerationResponseBodyPayloadOutput struct {
	// example:
	//
	// xx
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RunTranslateGenerationResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunTranslateGenerationResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunTranslateGenerationResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunTranslateGenerationResponseBodyPayloadOutput) SetText(v string) *RunTranslateGenerationResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunTranslateGenerationResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunTranslateGenerationResponseBodyPayloadUsage struct {
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

func (s RunTranslateGenerationResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunTranslateGenerationResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunTranslateGenerationResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunTranslateGenerationResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunTranslateGenerationResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunTranslateGenerationResponseBodyPayloadUsage) SetInputTokens(v int64) *RunTranslateGenerationResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunTranslateGenerationResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunTranslateGenerationResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunTranslateGenerationResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunTranslateGenerationResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunTranslateGenerationResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
