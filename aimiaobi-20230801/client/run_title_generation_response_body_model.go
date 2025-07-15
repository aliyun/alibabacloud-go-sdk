// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunTitleGenerationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RunTitleGenerationResponseBody
	GetCode() *string
	SetHeader(v *RunTitleGenerationResponseBodyHeader) *RunTitleGenerationResponseBody
	GetHeader() *RunTitleGenerationResponseBodyHeader
	SetHttpStatusCode(v string) *RunTitleGenerationResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *RunTitleGenerationResponseBody
	GetMessage() *string
	SetPayload(v *RunTitleGenerationResponseBodyPayload) *RunTitleGenerationResponseBody
	GetPayload() *RunTitleGenerationResponseBodyPayload
	SetRequestId(v string) *RunTitleGenerationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RunTitleGenerationResponseBody
	GetSuccess() *bool
}

type RunTitleGenerationResponseBody struct {
	Code           *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Header         *RunTitleGenerationResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	HttpStatusCode *string                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	Payload        *RunTitleGenerationResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// 94512A33-8EC1-5452-A793-5C91F18ED2F0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RunTitleGenerationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunTitleGenerationResponseBody) GoString() string {
	return s.String()
}

func (s *RunTitleGenerationResponseBody) GetCode() *string {
	return s.Code
}

func (s *RunTitleGenerationResponseBody) GetHeader() *RunTitleGenerationResponseBodyHeader {
	return s.Header
}

func (s *RunTitleGenerationResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *RunTitleGenerationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RunTitleGenerationResponseBody) GetPayload() *RunTitleGenerationResponseBodyPayload {
	return s.Payload
}

func (s *RunTitleGenerationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunTitleGenerationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RunTitleGenerationResponseBody) SetCode(v string) *RunTitleGenerationResponseBody {
	s.Code = &v
	return s
}

func (s *RunTitleGenerationResponseBody) SetHeader(v *RunTitleGenerationResponseBodyHeader) *RunTitleGenerationResponseBody {
	s.Header = v
	return s
}

func (s *RunTitleGenerationResponseBody) SetHttpStatusCode(v string) *RunTitleGenerationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RunTitleGenerationResponseBody) SetMessage(v string) *RunTitleGenerationResponseBody {
	s.Message = &v
	return s
}

func (s *RunTitleGenerationResponseBody) SetPayload(v *RunTitleGenerationResponseBodyPayload) *RunTitleGenerationResponseBody {
	s.Payload = v
	return s
}

func (s *RunTitleGenerationResponseBody) SetRequestId(v string) *RunTitleGenerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunTitleGenerationResponseBody) SetSuccess(v bool) *RunTitleGenerationResponseBody {
	s.Success = &v
	return s
}

func (s *RunTitleGenerationResponseBody) Validate() error {
	return dara.Validate(s)
}

type RunTitleGenerationResponseBodyHeader struct {
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
	// 3f7045e099474ba28ceca1b4eb6d6e21
	SessionId  *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	StatusCode *int32  `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// example:
	//
	// 50a1cc8e-717e-4a2b-a76b-dc9734a8564b
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 0a3d448f17000139741898287e0eb3
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunTitleGenerationResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunTitleGenerationResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunTitleGenerationResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunTitleGenerationResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunTitleGenerationResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunTitleGenerationResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunTitleGenerationResponseBodyHeader) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunTitleGenerationResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunTitleGenerationResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunTitleGenerationResponseBodyHeader) SetErrorCode(v string) *RunTitleGenerationResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunTitleGenerationResponseBodyHeader) SetErrorMessage(v string) *RunTitleGenerationResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunTitleGenerationResponseBodyHeader) SetEvent(v string) *RunTitleGenerationResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunTitleGenerationResponseBodyHeader) SetSessionId(v string) *RunTitleGenerationResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunTitleGenerationResponseBodyHeader) SetStatusCode(v int32) *RunTitleGenerationResponseBodyHeader {
	s.StatusCode = &v
	return s
}

func (s *RunTitleGenerationResponseBodyHeader) SetTaskId(v string) *RunTitleGenerationResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunTitleGenerationResponseBodyHeader) SetTraceId(v string) *RunTitleGenerationResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunTitleGenerationResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunTitleGenerationResponseBodyPayload struct {
	Output *RunTitleGenerationResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunTitleGenerationResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunTitleGenerationResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunTitleGenerationResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunTitleGenerationResponseBodyPayload) GetOutput() *RunTitleGenerationResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunTitleGenerationResponseBodyPayload) GetUsage() *RunTitleGenerationResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunTitleGenerationResponseBodyPayload) SetOutput(v *RunTitleGenerationResponseBodyPayloadOutput) *RunTitleGenerationResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunTitleGenerationResponseBodyPayload) SetUsage(v *RunTitleGenerationResponseBodyPayloadUsage) *RunTitleGenerationResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunTitleGenerationResponseBodyPayload) Validate() error {
	return dara.Validate(s)
}

type RunTitleGenerationResponseBodyPayloadOutput struct {
	// example:
	//
	// xxx
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RunTitleGenerationResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunTitleGenerationResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunTitleGenerationResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunTitleGenerationResponseBodyPayloadOutput) SetText(v string) *RunTitleGenerationResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunTitleGenerationResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunTitleGenerationResponseBodyPayloadUsage struct {
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

func (s RunTitleGenerationResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunTitleGenerationResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunTitleGenerationResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunTitleGenerationResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunTitleGenerationResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunTitleGenerationResponseBodyPayloadUsage) SetInputTokens(v int64) *RunTitleGenerationResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunTitleGenerationResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunTitleGenerationResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunTitleGenerationResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunTitleGenerationResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunTitleGenerationResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
