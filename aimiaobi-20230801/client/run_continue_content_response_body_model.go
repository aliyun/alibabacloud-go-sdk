// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunContinueContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnd(v bool) *RunContinueContentResponseBody
	GetEnd() *bool
	SetHeader(v *RunContinueContentResponseBodyHeader) *RunContinueContentResponseBody
	GetHeader() *RunContinueContentResponseBodyHeader
	SetPayload(v *RunContinueContentResponseBodyPayload) *RunContinueContentResponseBody
	GetPayload() *RunContinueContentResponseBodyPayload
	SetRequestId(v string) *RunContinueContentResponseBody
	GetRequestId() *string
}

type RunContinueContentResponseBody struct {
	End     *bool                                  `json:"End,omitempty" xml:"End,omitempty"`
	Header  *RunContinueContentResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunContinueContentResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// d3be9981-ca2d-4e17-bf31-1c0a628e9f99
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunContinueContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunContinueContentResponseBody) GoString() string {
	return s.String()
}

func (s *RunContinueContentResponseBody) GetEnd() *bool {
	return s.End
}

func (s *RunContinueContentResponseBody) GetHeader() *RunContinueContentResponseBodyHeader {
	return s.Header
}

func (s *RunContinueContentResponseBody) GetPayload() *RunContinueContentResponseBodyPayload {
	return s.Payload
}

func (s *RunContinueContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunContinueContentResponseBody) SetEnd(v bool) *RunContinueContentResponseBody {
	s.End = &v
	return s
}

func (s *RunContinueContentResponseBody) SetHeader(v *RunContinueContentResponseBodyHeader) *RunContinueContentResponseBody {
	s.Header = v
	return s
}

func (s *RunContinueContentResponseBody) SetPayload(v *RunContinueContentResponseBodyPayload) *RunContinueContentResponseBody {
	s.Payload = v
	return s
}

func (s *RunContinueContentResponseBody) SetRequestId(v string) *RunContinueContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunContinueContentResponseBody) Validate() error {
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

type RunContinueContentResponseBodyHeader struct {
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

func (s RunContinueContentResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunContinueContentResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunContinueContentResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunContinueContentResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunContinueContentResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunContinueContentResponseBodyHeader) GetEventInfo() *string {
	return s.EventInfo
}

func (s *RunContinueContentResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunContinueContentResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunContinueContentResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunContinueContentResponseBodyHeader) SetErrorCode(v string) *RunContinueContentResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunContinueContentResponseBodyHeader) SetErrorMessage(v string) *RunContinueContentResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunContinueContentResponseBodyHeader) SetEvent(v string) *RunContinueContentResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunContinueContentResponseBodyHeader) SetEventInfo(v string) *RunContinueContentResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunContinueContentResponseBodyHeader) SetSessionId(v string) *RunContinueContentResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunContinueContentResponseBodyHeader) SetTaskId(v string) *RunContinueContentResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunContinueContentResponseBodyHeader) SetTraceId(v string) *RunContinueContentResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunContinueContentResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunContinueContentResponseBodyPayload struct {
	Output *RunContinueContentResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunContinueContentResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunContinueContentResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunContinueContentResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunContinueContentResponseBodyPayload) GetOutput() *RunContinueContentResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunContinueContentResponseBodyPayload) GetUsage() *RunContinueContentResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunContinueContentResponseBodyPayload) SetOutput(v *RunContinueContentResponseBodyPayloadOutput) *RunContinueContentResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunContinueContentResponseBodyPayload) SetUsage(v *RunContinueContentResponseBodyPayloadUsage) *RunContinueContentResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunContinueContentResponseBodyPayload) Validate() error {
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

type RunContinueContentResponseBodyPayloadOutput struct {
	// example:
	//
	// 这是测试输出
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RunContinueContentResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunContinueContentResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunContinueContentResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunContinueContentResponseBodyPayloadOutput) SetText(v string) *RunContinueContentResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunContinueContentResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunContinueContentResponseBodyPayloadUsage struct {
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

func (s RunContinueContentResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunContinueContentResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunContinueContentResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunContinueContentResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunContinueContentResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunContinueContentResponseBodyPayloadUsage) SetInputTokens(v int64) *RunContinueContentResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunContinueContentResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunContinueContentResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunContinueContentResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunContinueContentResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunContinueContentResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
