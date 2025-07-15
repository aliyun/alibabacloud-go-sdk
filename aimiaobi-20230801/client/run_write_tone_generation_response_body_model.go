// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunWriteToneGenerationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeader(v *RunWriteToneGenerationResponseBodyHeader) *RunWriteToneGenerationResponseBody
	GetHeader() *RunWriteToneGenerationResponseBodyHeader
	SetPayload(v *RunWriteToneGenerationResponseBodyPayload) *RunWriteToneGenerationResponseBody
	GetPayload() *RunWriteToneGenerationResponseBodyPayload
	SetRequestId(v string) *RunWriteToneGenerationResponseBody
	GetRequestId() *string
}

type RunWriteToneGenerationResponseBody struct {
	Header  *RunWriteToneGenerationResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunWriteToneGenerationResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// FB698445-61DA-5361-BF73-1C5F1157E888
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunWriteToneGenerationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunWriteToneGenerationResponseBody) GoString() string {
	return s.String()
}

func (s *RunWriteToneGenerationResponseBody) GetHeader() *RunWriteToneGenerationResponseBodyHeader {
	return s.Header
}

func (s *RunWriteToneGenerationResponseBody) GetPayload() *RunWriteToneGenerationResponseBodyPayload {
	return s.Payload
}

func (s *RunWriteToneGenerationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunWriteToneGenerationResponseBody) SetHeader(v *RunWriteToneGenerationResponseBodyHeader) *RunWriteToneGenerationResponseBody {
	s.Header = v
	return s
}

func (s *RunWriteToneGenerationResponseBody) SetPayload(v *RunWriteToneGenerationResponseBodyPayload) *RunWriteToneGenerationResponseBody {
	s.Payload = v
	return s
}

func (s *RunWriteToneGenerationResponseBody) SetRequestId(v string) *RunWriteToneGenerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunWriteToneGenerationResponseBody) Validate() error {
	return dara.Validate(s)
}

type RunWriteToneGenerationResponseBodyHeader struct {
	// example:
	//
	// AccessForbid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// xxx
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// task-failed
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// F1953EE6-157C-40DC-BBF1-87C98AC27C51
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// F1953EE6-157C-40DC-BBF1-87C98AC27C51
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// F1953EE6-157C-40DC-BBF1-87C98AC27C51
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunWriteToneGenerationResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunWriteToneGenerationResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunWriteToneGenerationResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunWriteToneGenerationResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunWriteToneGenerationResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunWriteToneGenerationResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunWriteToneGenerationResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunWriteToneGenerationResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunWriteToneGenerationResponseBodyHeader) SetErrorCode(v string) *RunWriteToneGenerationResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunWriteToneGenerationResponseBodyHeader) SetErrorMessage(v string) *RunWriteToneGenerationResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunWriteToneGenerationResponseBodyHeader) SetEvent(v string) *RunWriteToneGenerationResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunWriteToneGenerationResponseBodyHeader) SetSessionId(v string) *RunWriteToneGenerationResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunWriteToneGenerationResponseBodyHeader) SetTaskId(v string) *RunWriteToneGenerationResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunWriteToneGenerationResponseBodyHeader) SetTraceId(v string) *RunWriteToneGenerationResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunWriteToneGenerationResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunWriteToneGenerationResponseBodyPayload struct {
	Output *RunWriteToneGenerationResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunWriteToneGenerationResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunWriteToneGenerationResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunWriteToneGenerationResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunWriteToneGenerationResponseBodyPayload) GetOutput() *RunWriteToneGenerationResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunWriteToneGenerationResponseBodyPayload) GetUsage() *RunWriteToneGenerationResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunWriteToneGenerationResponseBodyPayload) SetOutput(v *RunWriteToneGenerationResponseBodyPayloadOutput) *RunWriteToneGenerationResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunWriteToneGenerationResponseBodyPayload) SetUsage(v *RunWriteToneGenerationResponseBodyPayloadUsage) *RunWriteToneGenerationResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunWriteToneGenerationResponseBodyPayload) Validate() error {
	return dara.Validate(s)
}

type RunWriteToneGenerationResponseBodyPayloadOutput struct {
	// example:
	//
	// xxx
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RunWriteToneGenerationResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunWriteToneGenerationResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunWriteToneGenerationResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunWriteToneGenerationResponseBodyPayloadOutput) SetText(v string) *RunWriteToneGenerationResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunWriteToneGenerationResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunWriteToneGenerationResponseBodyPayloadUsage struct {
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

func (s RunWriteToneGenerationResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunWriteToneGenerationResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunWriteToneGenerationResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunWriteToneGenerationResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunWriteToneGenerationResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunWriteToneGenerationResponseBodyPayloadUsage) SetInputTokens(v int64) *RunWriteToneGenerationResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunWriteToneGenerationResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunWriteToneGenerationResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunWriteToneGenerationResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunWriteToneGenerationResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunWriteToneGenerationResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
