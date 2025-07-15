// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunTextPolishingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeader(v *RunTextPolishingResponseBodyHeader) *RunTextPolishingResponseBody
	GetHeader() *RunTextPolishingResponseBodyHeader
	SetPayload(v *RunTextPolishingResponseBodyPayload) *RunTextPolishingResponseBody
	GetPayload() *RunTextPolishingResponseBodyPayload
	SetRequestId(v string) *RunTextPolishingResponseBody
	GetRequestId() *string
}

type RunTextPolishingResponseBody struct {
	Header  *RunTextPolishingResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunTextPolishingResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunTextPolishingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunTextPolishingResponseBody) GoString() string {
	return s.String()
}

func (s *RunTextPolishingResponseBody) GetHeader() *RunTextPolishingResponseBodyHeader {
	return s.Header
}

func (s *RunTextPolishingResponseBody) GetPayload() *RunTextPolishingResponseBodyPayload {
	return s.Payload
}

func (s *RunTextPolishingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunTextPolishingResponseBody) SetHeader(v *RunTextPolishingResponseBodyHeader) *RunTextPolishingResponseBody {
	s.Header = v
	return s
}

func (s *RunTextPolishingResponseBody) SetPayload(v *RunTextPolishingResponseBodyPayload) *RunTextPolishingResponseBody {
	s.Payload = v
	return s
}

func (s *RunTextPolishingResponseBody) SetRequestId(v string) *RunTextPolishingResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunTextPolishingResponseBody) Validate() error {
	return dara.Validate(s)
}

type RunTextPolishingResponseBodyHeader struct {
	// example:
	//
	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 错误信息
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// task-started
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	OriginSessionId *string `json:"OriginSessionId,omitempty" xml:"OriginSessionId,omitempty"`
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
	// 全链路ID
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunTextPolishingResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunTextPolishingResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunTextPolishingResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunTextPolishingResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunTextPolishingResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunTextPolishingResponseBodyHeader) GetOriginSessionId() *string {
	return s.OriginSessionId
}

func (s *RunTextPolishingResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunTextPolishingResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunTextPolishingResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunTextPolishingResponseBodyHeader) SetErrorCode(v string) *RunTextPolishingResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunTextPolishingResponseBodyHeader) SetErrorMessage(v string) *RunTextPolishingResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunTextPolishingResponseBodyHeader) SetEvent(v string) *RunTextPolishingResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunTextPolishingResponseBodyHeader) SetOriginSessionId(v string) *RunTextPolishingResponseBodyHeader {
	s.OriginSessionId = &v
	return s
}

func (s *RunTextPolishingResponseBodyHeader) SetSessionId(v string) *RunTextPolishingResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunTextPolishingResponseBodyHeader) SetTaskId(v string) *RunTextPolishingResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunTextPolishingResponseBodyHeader) SetTraceId(v string) *RunTextPolishingResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunTextPolishingResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunTextPolishingResponseBodyPayload struct {
	Output *RunTextPolishingResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunTextPolishingResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunTextPolishingResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunTextPolishingResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunTextPolishingResponseBodyPayload) GetOutput() *RunTextPolishingResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunTextPolishingResponseBodyPayload) GetUsage() *RunTextPolishingResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunTextPolishingResponseBodyPayload) SetOutput(v *RunTextPolishingResponseBodyPayloadOutput) *RunTextPolishingResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunTextPolishingResponseBodyPayload) SetUsage(v *RunTextPolishingResponseBodyPayloadUsage) *RunTextPolishingResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunTextPolishingResponseBodyPayload) Validate() error {
	return dara.Validate(s)
}

type RunTextPolishingResponseBodyPayloadOutput struct {
	// example:
	//
	// 文本生成结果
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RunTextPolishingResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunTextPolishingResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunTextPolishingResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunTextPolishingResponseBodyPayloadOutput) SetText(v string) *RunTextPolishingResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunTextPolishingResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunTextPolishingResponseBodyPayloadUsage struct {
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

func (s RunTextPolishingResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunTextPolishingResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunTextPolishingResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunTextPolishingResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunTextPolishingResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunTextPolishingResponseBodyPayloadUsage) SetInputTokens(v int64) *RunTextPolishingResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunTextPolishingResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunTextPolishingResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunTextPolishingResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunTextPolishingResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunTextPolishingResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
