// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunStyleFeatureAnalysisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnd(v bool) *RunStyleFeatureAnalysisResponseBody
	GetEnd() *bool
	SetHeader(v *RunStyleFeatureAnalysisResponseBodyHeader) *RunStyleFeatureAnalysisResponseBody
	GetHeader() *RunStyleFeatureAnalysisResponseBodyHeader
	SetPayload(v *RunStyleFeatureAnalysisResponseBodyPayload) *RunStyleFeatureAnalysisResponseBody
	GetPayload() *RunStyleFeatureAnalysisResponseBodyPayload
	SetRequestId(v string) *RunStyleFeatureAnalysisResponseBody
	GetRequestId() *string
}

type RunStyleFeatureAnalysisResponseBody struct {
	End     *bool                                       `json:"End,omitempty" xml:"End,omitempty"`
	Header  *RunStyleFeatureAnalysisResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunStyleFeatureAnalysisResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// d3be9981-ca2d-4e17-bf31-1c0a628e9f99
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunStyleFeatureAnalysisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunStyleFeatureAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *RunStyleFeatureAnalysisResponseBody) GetEnd() *bool {
	return s.End
}

func (s *RunStyleFeatureAnalysisResponseBody) GetHeader() *RunStyleFeatureAnalysisResponseBodyHeader {
	return s.Header
}

func (s *RunStyleFeatureAnalysisResponseBody) GetPayload() *RunStyleFeatureAnalysisResponseBodyPayload {
	return s.Payload
}

func (s *RunStyleFeatureAnalysisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunStyleFeatureAnalysisResponseBody) SetEnd(v bool) *RunStyleFeatureAnalysisResponseBody {
	s.End = &v
	return s
}

func (s *RunStyleFeatureAnalysisResponseBody) SetHeader(v *RunStyleFeatureAnalysisResponseBodyHeader) *RunStyleFeatureAnalysisResponseBody {
	s.Header = v
	return s
}

func (s *RunStyleFeatureAnalysisResponseBody) SetPayload(v *RunStyleFeatureAnalysisResponseBodyPayload) *RunStyleFeatureAnalysisResponseBody {
	s.Payload = v
	return s
}

func (s *RunStyleFeatureAnalysisResponseBody) SetRequestId(v string) *RunStyleFeatureAnalysisResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunStyleFeatureAnalysisResponseBody) Validate() error {
	return dara.Validate(s)
}

type RunStyleFeatureAnalysisResponseBodyHeader struct {
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

func (s RunStyleFeatureAnalysisResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunStyleFeatureAnalysisResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunStyleFeatureAnalysisResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunStyleFeatureAnalysisResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunStyleFeatureAnalysisResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunStyleFeatureAnalysisResponseBodyHeader) GetEventInfo() *string {
	return s.EventInfo
}

func (s *RunStyleFeatureAnalysisResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunStyleFeatureAnalysisResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunStyleFeatureAnalysisResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunStyleFeatureAnalysisResponseBodyHeader) SetErrorCode(v string) *RunStyleFeatureAnalysisResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunStyleFeatureAnalysisResponseBodyHeader) SetErrorMessage(v string) *RunStyleFeatureAnalysisResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunStyleFeatureAnalysisResponseBodyHeader) SetEvent(v string) *RunStyleFeatureAnalysisResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunStyleFeatureAnalysisResponseBodyHeader) SetEventInfo(v string) *RunStyleFeatureAnalysisResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunStyleFeatureAnalysisResponseBodyHeader) SetSessionId(v string) *RunStyleFeatureAnalysisResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunStyleFeatureAnalysisResponseBodyHeader) SetTaskId(v string) *RunStyleFeatureAnalysisResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunStyleFeatureAnalysisResponseBodyHeader) SetTraceId(v string) *RunStyleFeatureAnalysisResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunStyleFeatureAnalysisResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunStyleFeatureAnalysisResponseBodyPayload struct {
	Output *RunStyleFeatureAnalysisResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunStyleFeatureAnalysisResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunStyleFeatureAnalysisResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunStyleFeatureAnalysisResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunStyleFeatureAnalysisResponseBodyPayload) GetOutput() *RunStyleFeatureAnalysisResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunStyleFeatureAnalysisResponseBodyPayload) GetUsage() *RunStyleFeatureAnalysisResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunStyleFeatureAnalysisResponseBodyPayload) SetOutput(v *RunStyleFeatureAnalysisResponseBodyPayloadOutput) *RunStyleFeatureAnalysisResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunStyleFeatureAnalysisResponseBodyPayload) SetUsage(v *RunStyleFeatureAnalysisResponseBodyPayloadUsage) *RunStyleFeatureAnalysisResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunStyleFeatureAnalysisResponseBodyPayload) Validate() error {
	return dara.Validate(s)
}

type RunStyleFeatureAnalysisResponseBodyPayloadOutput struct {
	// example:
	//
	// 这是测试输出
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RunStyleFeatureAnalysisResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunStyleFeatureAnalysisResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunStyleFeatureAnalysisResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunStyleFeatureAnalysisResponseBodyPayloadOutput) SetText(v string) *RunStyleFeatureAnalysisResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunStyleFeatureAnalysisResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunStyleFeatureAnalysisResponseBodyPayloadUsage struct {
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

func (s RunStyleFeatureAnalysisResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunStyleFeatureAnalysisResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunStyleFeatureAnalysisResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunStyleFeatureAnalysisResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunStyleFeatureAnalysisResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunStyleFeatureAnalysisResponseBodyPayloadUsage) SetInputTokens(v int64) *RunStyleFeatureAnalysisResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunStyleFeatureAnalysisResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunStyleFeatureAnalysisResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunStyleFeatureAnalysisResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunStyleFeatureAnalysisResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunStyleFeatureAnalysisResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
