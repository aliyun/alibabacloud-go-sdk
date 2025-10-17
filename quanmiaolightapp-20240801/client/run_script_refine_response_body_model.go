// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunScriptRefineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnd(v bool) *RunScriptRefineResponseBody
	GetEnd() *bool
	SetHeader(v *RunScriptRefineResponseBodyHeader) *RunScriptRefineResponseBody
	GetHeader() *RunScriptRefineResponseBodyHeader
	SetPayload(v *RunScriptRefineResponseBodyPayload) *RunScriptRefineResponseBody
	GetPayload() *RunScriptRefineResponseBodyPayload
}

type RunScriptRefineResponseBody struct {
	End     *bool                               `json:"end,omitempty" xml:"end,omitempty"`
	Header  *RunScriptRefineResponseBodyHeader  `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload *RunScriptRefineResponseBodyPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
}

func (s RunScriptRefineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunScriptRefineResponseBody) GoString() string {
	return s.String()
}

func (s *RunScriptRefineResponseBody) GetEnd() *bool {
	return s.End
}

func (s *RunScriptRefineResponseBody) GetHeader() *RunScriptRefineResponseBodyHeader {
	return s.Header
}

func (s *RunScriptRefineResponseBody) GetPayload() *RunScriptRefineResponseBodyPayload {
	return s.Payload
}

func (s *RunScriptRefineResponseBody) SetEnd(v bool) *RunScriptRefineResponseBody {
	s.End = &v
	return s
}

func (s *RunScriptRefineResponseBody) SetHeader(v *RunScriptRefineResponseBodyHeader) *RunScriptRefineResponseBody {
	s.Header = v
	return s
}

func (s *RunScriptRefineResponseBody) SetPayload(v *RunScriptRefineResponseBodyPayload) *RunScriptRefineResponseBody {
	s.Payload = v
	return s
}

func (s *RunScriptRefineResponseBody) Validate() error {
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

type RunScriptRefineResponseBodyHeader struct {
	// example:
	//
	// 403
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// Pop sign mismatch, please check log.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// result-generated
	Event     *string `json:"event,omitempty" xml:"event,omitempty"`
	EventInfo *string `json:"eventInfo,omitempty" xml:"eventInfo,omitempty"`
	// example:
	//
	// F8A35034-EDCF-5C50-95A5-1044316F36E3
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 17dc8bcd-f34a-46d1-a7a3-0fa3d1ce3824
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// 14356391-6c6c-40d5-b80a-8ecd03b69d72
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 2150432017236011824686132ecdbc
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s RunScriptRefineResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunScriptRefineResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunScriptRefineResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunScriptRefineResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunScriptRefineResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunScriptRefineResponseBodyHeader) GetEventInfo() *string {
	return s.EventInfo
}

func (s *RunScriptRefineResponseBodyHeader) GetRequestId() *string {
	return s.RequestId
}

func (s *RunScriptRefineResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunScriptRefineResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunScriptRefineResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunScriptRefineResponseBodyHeader) SetErrorCode(v string) *RunScriptRefineResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunScriptRefineResponseBodyHeader) SetErrorMessage(v string) *RunScriptRefineResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunScriptRefineResponseBodyHeader) SetEvent(v string) *RunScriptRefineResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunScriptRefineResponseBodyHeader) SetEventInfo(v string) *RunScriptRefineResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunScriptRefineResponseBodyHeader) SetRequestId(v string) *RunScriptRefineResponseBodyHeader {
	s.RequestId = &v
	return s
}

func (s *RunScriptRefineResponseBodyHeader) SetSessionId(v string) *RunScriptRefineResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunScriptRefineResponseBodyHeader) SetTaskId(v string) *RunScriptRefineResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunScriptRefineResponseBodyHeader) SetTraceId(v string) *RunScriptRefineResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunScriptRefineResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunScriptRefineResponseBodyPayload struct {
	Output *RunScriptRefineResponseBodyPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Usage  *RunScriptRefineResponseBodyPayloadUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunScriptRefineResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunScriptRefineResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunScriptRefineResponseBodyPayload) GetOutput() *RunScriptRefineResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunScriptRefineResponseBodyPayload) GetUsage() *RunScriptRefineResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunScriptRefineResponseBodyPayload) SetOutput(v *RunScriptRefineResponseBodyPayloadOutput) *RunScriptRefineResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunScriptRefineResponseBodyPayload) SetUsage(v *RunScriptRefineResponseBodyPayloadUsage) *RunScriptRefineResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunScriptRefineResponseBodyPayload) Validate() error {
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

type RunScriptRefineResponseBodyPayloadOutput struct {
	Content []map[string]*string `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	Outline *string              `json:"outline,omitempty" xml:"outline,omitempty"`
	Role    *string              `json:"role,omitempty" xml:"role,omitempty"`
	Scene   *string              `json:"scene,omitempty" xml:"scene,omitempty"`
	Summary *string              `json:"summary,omitempty" xml:"summary,omitempty"`
	// example:
	//
	// xx
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s RunScriptRefineResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunScriptRefineResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunScriptRefineResponseBodyPayloadOutput) GetContent() []map[string]*string {
	return s.Content
}

func (s *RunScriptRefineResponseBodyPayloadOutput) GetOutline() *string {
	return s.Outline
}

func (s *RunScriptRefineResponseBodyPayloadOutput) GetRole() *string {
	return s.Role
}

func (s *RunScriptRefineResponseBodyPayloadOutput) GetScene() *string {
	return s.Scene
}

func (s *RunScriptRefineResponseBodyPayloadOutput) GetSummary() *string {
	return s.Summary
}

func (s *RunScriptRefineResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunScriptRefineResponseBodyPayloadOutput) SetContent(v []map[string]*string) *RunScriptRefineResponseBodyPayloadOutput {
	s.Content = v
	return s
}

func (s *RunScriptRefineResponseBodyPayloadOutput) SetOutline(v string) *RunScriptRefineResponseBodyPayloadOutput {
	s.Outline = &v
	return s
}

func (s *RunScriptRefineResponseBodyPayloadOutput) SetRole(v string) *RunScriptRefineResponseBodyPayloadOutput {
	s.Role = &v
	return s
}

func (s *RunScriptRefineResponseBodyPayloadOutput) SetScene(v string) *RunScriptRefineResponseBodyPayloadOutput {
	s.Scene = &v
	return s
}

func (s *RunScriptRefineResponseBodyPayloadOutput) SetSummary(v string) *RunScriptRefineResponseBodyPayloadOutput {
	s.Summary = &v
	return s
}

func (s *RunScriptRefineResponseBodyPayloadOutput) SetText(v string) *RunScriptRefineResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunScriptRefineResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunScriptRefineResponseBodyPayloadUsage struct {
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

func (s RunScriptRefineResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunScriptRefineResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunScriptRefineResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunScriptRefineResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunScriptRefineResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunScriptRefineResponseBodyPayloadUsage) SetInputTokens(v int64) *RunScriptRefineResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunScriptRefineResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunScriptRefineResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunScriptRefineResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunScriptRefineResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunScriptRefineResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
