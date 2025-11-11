// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSummaryGenerateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnd(v bool) *RunSummaryGenerateResponseBody
	GetEnd() *bool
	SetHeader(v *RunSummaryGenerateResponseBodyHeader) *RunSummaryGenerateResponseBody
	GetHeader() *RunSummaryGenerateResponseBodyHeader
	SetPayload(v *RunSummaryGenerateResponseBodyPayload) *RunSummaryGenerateResponseBody
	GetPayload() *RunSummaryGenerateResponseBodyPayload
	SetRequestId(v string) *RunSummaryGenerateResponseBody
	GetRequestId() *string
}

type RunSummaryGenerateResponseBody struct {
	End     *bool                                  `json:"End,omitempty" xml:"End,omitempty"`
	Header  *RunSummaryGenerateResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunSummaryGenerateResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// d3be9981-ca2d-4e17-bf31-1c0a628e9f99
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunSummaryGenerateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunSummaryGenerateResponseBody) GoString() string {
	return s.String()
}

func (s *RunSummaryGenerateResponseBody) GetEnd() *bool {
	return s.End
}

func (s *RunSummaryGenerateResponseBody) GetHeader() *RunSummaryGenerateResponseBodyHeader {
	return s.Header
}

func (s *RunSummaryGenerateResponseBody) GetPayload() *RunSummaryGenerateResponseBodyPayload {
	return s.Payload
}

func (s *RunSummaryGenerateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunSummaryGenerateResponseBody) SetEnd(v bool) *RunSummaryGenerateResponseBody {
	s.End = &v
	return s
}

func (s *RunSummaryGenerateResponseBody) SetHeader(v *RunSummaryGenerateResponseBodyHeader) *RunSummaryGenerateResponseBody {
	s.Header = v
	return s
}

func (s *RunSummaryGenerateResponseBody) SetPayload(v *RunSummaryGenerateResponseBodyPayload) *RunSummaryGenerateResponseBody {
	s.Payload = v
	return s
}

func (s *RunSummaryGenerateResponseBody) SetRequestId(v string) *RunSummaryGenerateResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunSummaryGenerateResponseBody) Validate() error {
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

type RunSummaryGenerateResponseBodyHeader struct {
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

func (s RunSummaryGenerateResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunSummaryGenerateResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunSummaryGenerateResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunSummaryGenerateResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunSummaryGenerateResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunSummaryGenerateResponseBodyHeader) GetEventInfo() *string {
	return s.EventInfo
}

func (s *RunSummaryGenerateResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunSummaryGenerateResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunSummaryGenerateResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunSummaryGenerateResponseBodyHeader) SetErrorCode(v string) *RunSummaryGenerateResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunSummaryGenerateResponseBodyHeader) SetErrorMessage(v string) *RunSummaryGenerateResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunSummaryGenerateResponseBodyHeader) SetEvent(v string) *RunSummaryGenerateResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunSummaryGenerateResponseBodyHeader) SetEventInfo(v string) *RunSummaryGenerateResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunSummaryGenerateResponseBodyHeader) SetSessionId(v string) *RunSummaryGenerateResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunSummaryGenerateResponseBodyHeader) SetTaskId(v string) *RunSummaryGenerateResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunSummaryGenerateResponseBodyHeader) SetTraceId(v string) *RunSummaryGenerateResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunSummaryGenerateResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunSummaryGenerateResponseBodyPayload struct {
	Output *RunSummaryGenerateResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunSummaryGenerateResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunSummaryGenerateResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunSummaryGenerateResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunSummaryGenerateResponseBodyPayload) GetOutput() *RunSummaryGenerateResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunSummaryGenerateResponseBodyPayload) GetUsage() *RunSummaryGenerateResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunSummaryGenerateResponseBodyPayload) SetOutput(v *RunSummaryGenerateResponseBodyPayloadOutput) *RunSummaryGenerateResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunSummaryGenerateResponseBodyPayload) SetUsage(v *RunSummaryGenerateResponseBodyPayloadUsage) *RunSummaryGenerateResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunSummaryGenerateResponseBodyPayload) Validate() error {
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

type RunSummaryGenerateResponseBodyPayloadOutput struct {
	// example:
	//
	// 这是测试输出
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RunSummaryGenerateResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunSummaryGenerateResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunSummaryGenerateResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunSummaryGenerateResponseBodyPayloadOutput) SetText(v string) *RunSummaryGenerateResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunSummaryGenerateResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunSummaryGenerateResponseBodyPayloadUsage struct {
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

func (s RunSummaryGenerateResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunSummaryGenerateResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunSummaryGenerateResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunSummaryGenerateResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunSummaryGenerateResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunSummaryGenerateResponseBodyPayloadUsage) SetInputTokens(v int64) *RunSummaryGenerateResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunSummaryGenerateResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunSummaryGenerateResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunSummaryGenerateResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunSummaryGenerateResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunSummaryGenerateResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
