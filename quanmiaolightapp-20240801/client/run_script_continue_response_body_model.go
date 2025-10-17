// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunScriptContinueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnd(v bool) *RunScriptContinueResponseBody
	GetEnd() *bool
	SetHeader(v *RunScriptContinueResponseBodyHeader) *RunScriptContinueResponseBody
	GetHeader() *RunScriptContinueResponseBodyHeader
	SetPayload(v *RunScriptContinueResponseBodyPayload) *RunScriptContinueResponseBody
	GetPayload() *RunScriptContinueResponseBodyPayload
}

type RunScriptContinueResponseBody struct {
	End     *bool                                 `json:"end,omitempty" xml:"end,omitempty"`
	Header  *RunScriptContinueResponseBodyHeader  `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload *RunScriptContinueResponseBodyPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
}

func (s RunScriptContinueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunScriptContinueResponseBody) GoString() string {
	return s.String()
}

func (s *RunScriptContinueResponseBody) GetEnd() *bool {
	return s.End
}

func (s *RunScriptContinueResponseBody) GetHeader() *RunScriptContinueResponseBodyHeader {
	return s.Header
}

func (s *RunScriptContinueResponseBody) GetPayload() *RunScriptContinueResponseBodyPayload {
	return s.Payload
}

func (s *RunScriptContinueResponseBody) SetEnd(v bool) *RunScriptContinueResponseBody {
	s.End = &v
	return s
}

func (s *RunScriptContinueResponseBody) SetHeader(v *RunScriptContinueResponseBodyHeader) *RunScriptContinueResponseBody {
	s.Header = v
	return s
}

func (s *RunScriptContinueResponseBody) SetPayload(v *RunScriptContinueResponseBodyPayload) *RunScriptContinueResponseBody {
	s.Payload = v
	return s
}

func (s *RunScriptContinueResponseBody) Validate() error {
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

type RunScriptContinueResponseBodyHeader struct {
	// example:
	//
	// 403
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// Pop sign mismatch, please check.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// result-generated
	Event *string `json:"event,omitempty" xml:"event,omitempty"`
	// example:
	//
	// 模型生成事件
	EventInfo *string `json:"eventInfo,omitempty" xml:"eventInfo,omitempty"`
	// example:
	//
	// 0EB27AE3-CA53-5FAE-83C6-EE66CA4DF5DF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 3cd10828-0e42-471c-8f1a-931cde20b035
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// d3be9981-ca2d-4e17-bf31-1c0a628e9f99
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 2150451a17191950923411783e2927
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s RunScriptContinueResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunScriptContinueResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunScriptContinueResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunScriptContinueResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunScriptContinueResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunScriptContinueResponseBodyHeader) GetEventInfo() *string {
	return s.EventInfo
}

func (s *RunScriptContinueResponseBodyHeader) GetRequestId() *string {
	return s.RequestId
}

func (s *RunScriptContinueResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunScriptContinueResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunScriptContinueResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunScriptContinueResponseBodyHeader) SetErrorCode(v string) *RunScriptContinueResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunScriptContinueResponseBodyHeader) SetErrorMessage(v string) *RunScriptContinueResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunScriptContinueResponseBodyHeader) SetEvent(v string) *RunScriptContinueResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunScriptContinueResponseBodyHeader) SetEventInfo(v string) *RunScriptContinueResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunScriptContinueResponseBodyHeader) SetRequestId(v string) *RunScriptContinueResponseBodyHeader {
	s.RequestId = &v
	return s
}

func (s *RunScriptContinueResponseBodyHeader) SetSessionId(v string) *RunScriptContinueResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunScriptContinueResponseBodyHeader) SetTaskId(v string) *RunScriptContinueResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunScriptContinueResponseBodyHeader) SetTraceId(v string) *RunScriptContinueResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunScriptContinueResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunScriptContinueResponseBodyPayload struct {
	Output *RunScriptContinueResponseBodyPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Usage  *RunScriptContinueResponseBodyPayloadUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunScriptContinueResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunScriptContinueResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunScriptContinueResponseBodyPayload) GetOutput() *RunScriptContinueResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunScriptContinueResponseBodyPayload) GetUsage() *RunScriptContinueResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunScriptContinueResponseBodyPayload) SetOutput(v *RunScriptContinueResponseBodyPayloadOutput) *RunScriptContinueResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunScriptContinueResponseBodyPayload) SetUsage(v *RunScriptContinueResponseBodyPayloadUsage) *RunScriptContinueResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunScriptContinueResponseBodyPayload) Validate() error {
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

type RunScriptContinueResponseBodyPayloadOutput struct {
	// example:
	//
	// 这是测试输出
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s RunScriptContinueResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunScriptContinueResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunScriptContinueResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunScriptContinueResponseBodyPayloadOutput) SetText(v string) *RunScriptContinueResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunScriptContinueResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunScriptContinueResponseBodyPayloadUsage struct {
	// example:
	//
	// 100
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 100
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 200
	TotalTokens *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s RunScriptContinueResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunScriptContinueResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunScriptContinueResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunScriptContinueResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunScriptContinueResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunScriptContinueResponseBodyPayloadUsage) SetInputTokens(v int64) *RunScriptContinueResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunScriptContinueResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunScriptContinueResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunScriptContinueResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunScriptContinueResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunScriptContinueResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
