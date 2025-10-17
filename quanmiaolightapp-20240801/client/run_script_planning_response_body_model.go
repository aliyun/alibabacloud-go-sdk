// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunScriptPlanningResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnd(v bool) *RunScriptPlanningResponseBody
	GetEnd() *bool
	SetHeader(v *RunScriptPlanningResponseBodyHeader) *RunScriptPlanningResponseBody
	GetHeader() *RunScriptPlanningResponseBodyHeader
	SetPayload(v *RunScriptPlanningResponseBodyPayload) *RunScriptPlanningResponseBody
	GetPayload() *RunScriptPlanningResponseBodyPayload
}

type RunScriptPlanningResponseBody struct {
	End     *bool                                 `json:"end,omitempty" xml:"end,omitempty"`
	Header  *RunScriptPlanningResponseBodyHeader  `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload *RunScriptPlanningResponseBodyPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
}

func (s RunScriptPlanningResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunScriptPlanningResponseBody) GoString() string {
	return s.String()
}

func (s *RunScriptPlanningResponseBody) GetEnd() *bool {
	return s.End
}

func (s *RunScriptPlanningResponseBody) GetHeader() *RunScriptPlanningResponseBodyHeader {
	return s.Header
}

func (s *RunScriptPlanningResponseBody) GetPayload() *RunScriptPlanningResponseBodyPayload {
	return s.Payload
}

func (s *RunScriptPlanningResponseBody) SetEnd(v bool) *RunScriptPlanningResponseBody {
	s.End = &v
	return s
}

func (s *RunScriptPlanningResponseBody) SetHeader(v *RunScriptPlanningResponseBodyHeader) *RunScriptPlanningResponseBody {
	s.Header = v
	return s
}

func (s *RunScriptPlanningResponseBody) SetPayload(v *RunScriptPlanningResponseBodyPayload) *RunScriptPlanningResponseBody {
	s.Payload = v
	return s
}

func (s *RunScriptPlanningResponseBody) Validate() error {
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

type RunScriptPlanningResponseBodyHeader struct {
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

func (s RunScriptPlanningResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunScriptPlanningResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunScriptPlanningResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunScriptPlanningResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunScriptPlanningResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunScriptPlanningResponseBodyHeader) GetEventInfo() *string {
	return s.EventInfo
}

func (s *RunScriptPlanningResponseBodyHeader) GetRequestId() *string {
	return s.RequestId
}

func (s *RunScriptPlanningResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunScriptPlanningResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunScriptPlanningResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunScriptPlanningResponseBodyHeader) SetErrorCode(v string) *RunScriptPlanningResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunScriptPlanningResponseBodyHeader) SetErrorMessage(v string) *RunScriptPlanningResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunScriptPlanningResponseBodyHeader) SetEvent(v string) *RunScriptPlanningResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunScriptPlanningResponseBodyHeader) SetEventInfo(v string) *RunScriptPlanningResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunScriptPlanningResponseBodyHeader) SetRequestId(v string) *RunScriptPlanningResponseBodyHeader {
	s.RequestId = &v
	return s
}

func (s *RunScriptPlanningResponseBodyHeader) SetSessionId(v string) *RunScriptPlanningResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunScriptPlanningResponseBodyHeader) SetTaskId(v string) *RunScriptPlanningResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunScriptPlanningResponseBodyHeader) SetTraceId(v string) *RunScriptPlanningResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunScriptPlanningResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunScriptPlanningResponseBodyPayload struct {
	Output *RunScriptPlanningResponseBodyPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Usage  *RunScriptPlanningResponseBodyPayloadUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunScriptPlanningResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunScriptPlanningResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunScriptPlanningResponseBodyPayload) GetOutput() *RunScriptPlanningResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunScriptPlanningResponseBodyPayload) GetUsage() *RunScriptPlanningResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunScriptPlanningResponseBodyPayload) SetOutput(v *RunScriptPlanningResponseBodyPayloadOutput) *RunScriptPlanningResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunScriptPlanningResponseBodyPayload) SetUsage(v *RunScriptPlanningResponseBodyPayloadUsage) *RunScriptPlanningResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunScriptPlanningResponseBodyPayload) Validate() error {
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

type RunScriptPlanningResponseBodyPayloadOutput struct {
	// example:
	//
	// 这是测试输出
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s RunScriptPlanningResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunScriptPlanningResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunScriptPlanningResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunScriptPlanningResponseBodyPayloadOutput) SetText(v string) *RunScriptPlanningResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunScriptPlanningResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunScriptPlanningResponseBodyPayloadUsage struct {
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

func (s RunScriptPlanningResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunScriptPlanningResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunScriptPlanningResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunScriptPlanningResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunScriptPlanningResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunScriptPlanningResponseBodyPayloadUsage) SetInputTokens(v int64) *RunScriptPlanningResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunScriptPlanningResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunScriptPlanningResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunScriptPlanningResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunScriptPlanningResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunScriptPlanningResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
