// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunVideoScriptGenerateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RunVideoScriptGenerateResponseBody
	GetCode() *string
	SetHeader(v *RunVideoScriptGenerateResponseBodyHeader) *RunVideoScriptGenerateResponseBody
	GetHeader() *RunVideoScriptGenerateResponseBodyHeader
	SetHttpStatusCode(v string) *RunVideoScriptGenerateResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *RunVideoScriptGenerateResponseBody
	GetMessage() *string
	SetPayload(v *RunVideoScriptGenerateResponseBodyPayload) *RunVideoScriptGenerateResponseBody
	GetPayload() *RunVideoScriptGenerateResponseBodyPayload
	SetRequestId(v string) *RunVideoScriptGenerateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RunVideoScriptGenerateResponseBody
	GetSuccess() *bool
}

type RunVideoScriptGenerateResponseBody struct {
	// example:
	//
	// NoPermission
	Code   *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Header *RunVideoScriptGenerateResponseBodyHeader `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	// example:
	//
	// 403
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// You are not authorized to perform this action.
	Message *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	Payload *RunVideoScriptGenerateResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// F2F366D6-E9FE-1006-BB70-2C650896AAB5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RunVideoScriptGenerateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunVideoScriptGenerateResponseBody) GoString() string {
	return s.String()
}

func (s *RunVideoScriptGenerateResponseBody) GetCode() *string {
	return s.Code
}

func (s *RunVideoScriptGenerateResponseBody) GetHeader() *RunVideoScriptGenerateResponseBodyHeader {
	return s.Header
}

func (s *RunVideoScriptGenerateResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *RunVideoScriptGenerateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RunVideoScriptGenerateResponseBody) GetPayload() *RunVideoScriptGenerateResponseBodyPayload {
	return s.Payload
}

func (s *RunVideoScriptGenerateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunVideoScriptGenerateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RunVideoScriptGenerateResponseBody) SetCode(v string) *RunVideoScriptGenerateResponseBody {
	s.Code = &v
	return s
}

func (s *RunVideoScriptGenerateResponseBody) SetHeader(v *RunVideoScriptGenerateResponseBodyHeader) *RunVideoScriptGenerateResponseBody {
	s.Header = v
	return s
}

func (s *RunVideoScriptGenerateResponseBody) SetHttpStatusCode(v string) *RunVideoScriptGenerateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RunVideoScriptGenerateResponseBody) SetMessage(v string) *RunVideoScriptGenerateResponseBody {
	s.Message = &v
	return s
}

func (s *RunVideoScriptGenerateResponseBody) SetPayload(v *RunVideoScriptGenerateResponseBodyPayload) *RunVideoScriptGenerateResponseBody {
	s.Payload = v
	return s
}

func (s *RunVideoScriptGenerateResponseBody) SetRequestId(v string) *RunVideoScriptGenerateResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunVideoScriptGenerateResponseBody) SetSuccess(v bool) *RunVideoScriptGenerateResponseBody {
	s.Success = &v
	return s
}

func (s *RunVideoScriptGenerateResponseBody) Validate() error {
	return dara.Validate(s)
}

type RunVideoScriptGenerateResponseBodyHeader struct {
	// example:
	//
	// ScriptNumberExceed
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// result-generated
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 400
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// example:
	//
	// b057f2fa-2277-477b-babf-cbc062307828
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 2150451a17191950923411783e2927
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunVideoScriptGenerateResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunVideoScriptGenerateResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunVideoScriptGenerateResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunVideoScriptGenerateResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunVideoScriptGenerateResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunVideoScriptGenerateResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunVideoScriptGenerateResponseBodyHeader) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunVideoScriptGenerateResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunVideoScriptGenerateResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunVideoScriptGenerateResponseBodyHeader) SetErrorCode(v string) *RunVideoScriptGenerateResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunVideoScriptGenerateResponseBodyHeader) SetErrorMessage(v string) *RunVideoScriptGenerateResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunVideoScriptGenerateResponseBodyHeader) SetEvent(v string) *RunVideoScriptGenerateResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunVideoScriptGenerateResponseBodyHeader) SetSessionId(v string) *RunVideoScriptGenerateResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunVideoScriptGenerateResponseBodyHeader) SetStatusCode(v int32) *RunVideoScriptGenerateResponseBodyHeader {
	s.StatusCode = &v
	return s
}

func (s *RunVideoScriptGenerateResponseBodyHeader) SetTaskId(v string) *RunVideoScriptGenerateResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunVideoScriptGenerateResponseBodyHeader) SetTraceId(v string) *RunVideoScriptGenerateResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunVideoScriptGenerateResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunVideoScriptGenerateResponseBodyPayload struct {
	Output *RunVideoScriptGenerateResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunVideoScriptGenerateResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunVideoScriptGenerateResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunVideoScriptGenerateResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunVideoScriptGenerateResponseBodyPayload) GetOutput() *RunVideoScriptGenerateResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunVideoScriptGenerateResponseBodyPayload) GetUsage() *RunVideoScriptGenerateResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunVideoScriptGenerateResponseBodyPayload) SetOutput(v *RunVideoScriptGenerateResponseBodyPayloadOutput) *RunVideoScriptGenerateResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunVideoScriptGenerateResponseBodyPayload) SetUsage(v *RunVideoScriptGenerateResponseBodyPayloadUsage) *RunVideoScriptGenerateResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunVideoScriptGenerateResponseBodyPayload) Validate() error {
	return dara.Validate(s)
}

type RunVideoScriptGenerateResponseBodyPayloadOutput struct {
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RunVideoScriptGenerateResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunVideoScriptGenerateResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunVideoScriptGenerateResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunVideoScriptGenerateResponseBodyPayloadOutput) SetText(v string) *RunVideoScriptGenerateResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunVideoScriptGenerateResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunVideoScriptGenerateResponseBodyPayloadUsage struct {
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

func (s RunVideoScriptGenerateResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunVideoScriptGenerateResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunVideoScriptGenerateResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunVideoScriptGenerateResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunVideoScriptGenerateResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunVideoScriptGenerateResponseBodyPayloadUsage) SetInputTokens(v int64) *RunVideoScriptGenerateResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunVideoScriptGenerateResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunVideoScriptGenerateResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunVideoScriptGenerateResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunVideoScriptGenerateResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunVideoScriptGenerateResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
