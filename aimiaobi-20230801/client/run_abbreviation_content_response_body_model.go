// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunAbbreviationContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnd(v bool) *RunAbbreviationContentResponseBody
	GetEnd() *bool
	SetHeader(v *RunAbbreviationContentResponseBodyHeader) *RunAbbreviationContentResponseBody
	GetHeader() *RunAbbreviationContentResponseBodyHeader
	SetPayload(v *RunAbbreviationContentResponseBodyPayload) *RunAbbreviationContentResponseBody
	GetPayload() *RunAbbreviationContentResponseBodyPayload
	SetRequestId(v string) *RunAbbreviationContentResponseBody
	GetRequestId() *string
}

type RunAbbreviationContentResponseBody struct {
	End     *bool                                      `json:"End,omitempty" xml:"End,omitempty"`
	Header  *RunAbbreviationContentResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunAbbreviationContentResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// d3be9981-ca2d-4e17-bf31-1c0a628e9f99
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunAbbreviationContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunAbbreviationContentResponseBody) GoString() string {
	return s.String()
}

func (s *RunAbbreviationContentResponseBody) GetEnd() *bool {
	return s.End
}

func (s *RunAbbreviationContentResponseBody) GetHeader() *RunAbbreviationContentResponseBodyHeader {
	return s.Header
}

func (s *RunAbbreviationContentResponseBody) GetPayload() *RunAbbreviationContentResponseBodyPayload {
	return s.Payload
}

func (s *RunAbbreviationContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunAbbreviationContentResponseBody) SetEnd(v bool) *RunAbbreviationContentResponseBody {
	s.End = &v
	return s
}

func (s *RunAbbreviationContentResponseBody) SetHeader(v *RunAbbreviationContentResponseBodyHeader) *RunAbbreviationContentResponseBody {
	s.Header = v
	return s
}

func (s *RunAbbreviationContentResponseBody) SetPayload(v *RunAbbreviationContentResponseBodyPayload) *RunAbbreviationContentResponseBody {
	s.Payload = v
	return s
}

func (s *RunAbbreviationContentResponseBody) SetRequestId(v string) *RunAbbreviationContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunAbbreviationContentResponseBody) Validate() error {
	return dara.Validate(s)
}

type RunAbbreviationContentResponseBodyHeader struct {
	// example:
	//
	// 403
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Pop sign mismatch, please check.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// result-generated
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// 模型生成事件
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
	// 2150451a17191950923411783e2927
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunAbbreviationContentResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunAbbreviationContentResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunAbbreviationContentResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunAbbreviationContentResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunAbbreviationContentResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunAbbreviationContentResponseBodyHeader) GetEventInfo() *string {
	return s.EventInfo
}

func (s *RunAbbreviationContentResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunAbbreviationContentResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunAbbreviationContentResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunAbbreviationContentResponseBodyHeader) SetErrorCode(v string) *RunAbbreviationContentResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunAbbreviationContentResponseBodyHeader) SetErrorMessage(v string) *RunAbbreviationContentResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunAbbreviationContentResponseBodyHeader) SetEvent(v string) *RunAbbreviationContentResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunAbbreviationContentResponseBodyHeader) SetEventInfo(v string) *RunAbbreviationContentResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunAbbreviationContentResponseBodyHeader) SetSessionId(v string) *RunAbbreviationContentResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunAbbreviationContentResponseBodyHeader) SetTaskId(v string) *RunAbbreviationContentResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunAbbreviationContentResponseBodyHeader) SetTraceId(v string) *RunAbbreviationContentResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunAbbreviationContentResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunAbbreviationContentResponseBodyPayload struct {
	Output *RunAbbreviationContentResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunAbbreviationContentResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunAbbreviationContentResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunAbbreviationContentResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunAbbreviationContentResponseBodyPayload) GetOutput() *RunAbbreviationContentResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunAbbreviationContentResponseBodyPayload) GetUsage() *RunAbbreviationContentResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunAbbreviationContentResponseBodyPayload) SetOutput(v *RunAbbreviationContentResponseBodyPayloadOutput) *RunAbbreviationContentResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunAbbreviationContentResponseBodyPayload) SetUsage(v *RunAbbreviationContentResponseBodyPayloadUsage) *RunAbbreviationContentResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunAbbreviationContentResponseBodyPayload) Validate() error {
	return dara.Validate(s)
}

type RunAbbreviationContentResponseBodyPayloadOutput struct {
	// example:
	//
	// 这是测试输出
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RunAbbreviationContentResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunAbbreviationContentResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunAbbreviationContentResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunAbbreviationContentResponseBodyPayloadOutput) SetText(v string) *RunAbbreviationContentResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunAbbreviationContentResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunAbbreviationContentResponseBodyPayloadUsage struct {
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

func (s RunAbbreviationContentResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunAbbreviationContentResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunAbbreviationContentResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunAbbreviationContentResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunAbbreviationContentResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunAbbreviationContentResponseBodyPayloadUsage) SetInputTokens(v int64) *RunAbbreviationContentResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunAbbreviationContentResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunAbbreviationContentResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunAbbreviationContentResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunAbbreviationContentResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunAbbreviationContentResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
