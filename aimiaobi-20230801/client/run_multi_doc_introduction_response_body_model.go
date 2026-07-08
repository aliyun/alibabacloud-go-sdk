// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunMultiDocIntroductionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeader(v *RunMultiDocIntroductionResponseBodyHeader) *RunMultiDocIntroductionResponseBody
	GetHeader() *RunMultiDocIntroductionResponseBodyHeader
	SetPayload(v *RunMultiDocIntroductionResponseBodyPayload) *RunMultiDocIntroductionResponseBody
	GetPayload() *RunMultiDocIntroductionResponseBodyPayload
	SetRequestId(v string) *RunMultiDocIntroductionResponseBody
	GetRequestId() *string
}

type RunMultiDocIntroductionResponseBody struct {
	// Response header.
	Header *RunMultiDocIntroductionResponseBodyHeader `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	// Response body.
	Payload *RunMultiDocIntroductionResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// ID of the request.
	//
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunMultiDocIntroductionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunMultiDocIntroductionResponseBody) GoString() string {
	return s.String()
}

func (s *RunMultiDocIntroductionResponseBody) GetHeader() *RunMultiDocIntroductionResponseBodyHeader {
	return s.Header
}

func (s *RunMultiDocIntroductionResponseBody) GetPayload() *RunMultiDocIntroductionResponseBodyPayload {
	return s.Payload
}

func (s *RunMultiDocIntroductionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunMultiDocIntroductionResponseBody) SetHeader(v *RunMultiDocIntroductionResponseBodyHeader) *RunMultiDocIntroductionResponseBody {
	s.Header = v
	return s
}

func (s *RunMultiDocIntroductionResponseBody) SetPayload(v *RunMultiDocIntroductionResponseBodyPayload) *RunMultiDocIntroductionResponseBody {
	s.Payload = v
	return s
}

func (s *RunMultiDocIntroductionResponseBody) SetRequestId(v string) *RunMultiDocIntroductionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunMultiDocIntroductionResponseBody) Validate() error {
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

type RunMultiDocIntroductionResponseBodyHeader struct {
	// Error code.
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Error message.
	//
	// example:
	//
	// Message does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Server-sent event (SSE) type.
	//
	// example:
	//
	// finished
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// Event description.
	//
	// example:
	//
	// 模型生成事件
	EventInfo *string `json:"EventInfo,omitempty" xml:"EventInfo,omitempty"`
	// Session ID.
	//
	// example:
	//
	// 92e16ccb-92b6-4894-abbf-fc6e2929a0df
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// Task ID.
	//
	// example:
	//
	// b057f2fa-2277-477b-babf-cbc062307828
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// End-to-end trace ID.
	//
	// example:
	//
	// 46e5c2b5-0877-4f09-bd91-ab0cf314e48b
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunMultiDocIntroductionResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunMultiDocIntroductionResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunMultiDocIntroductionResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunMultiDocIntroductionResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunMultiDocIntroductionResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunMultiDocIntroductionResponseBodyHeader) GetEventInfo() *string {
	return s.EventInfo
}

func (s *RunMultiDocIntroductionResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunMultiDocIntroductionResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunMultiDocIntroductionResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunMultiDocIntroductionResponseBodyHeader) SetErrorCode(v string) *RunMultiDocIntroductionResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunMultiDocIntroductionResponseBodyHeader) SetErrorMessage(v string) *RunMultiDocIntroductionResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunMultiDocIntroductionResponseBodyHeader) SetEvent(v string) *RunMultiDocIntroductionResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunMultiDocIntroductionResponseBodyHeader) SetEventInfo(v string) *RunMultiDocIntroductionResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunMultiDocIntroductionResponseBodyHeader) SetSessionId(v string) *RunMultiDocIntroductionResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunMultiDocIntroductionResponseBodyHeader) SetTaskId(v string) *RunMultiDocIntroductionResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunMultiDocIntroductionResponseBodyHeader) SetTraceId(v string) *RunMultiDocIntroductionResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunMultiDocIntroductionResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunMultiDocIntroductionResponseBodyPayload struct {
	// Output data.
	Output *RunMultiDocIntroductionResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// Token usage.
	Usage *RunMultiDocIntroductionResponseBodyPayloadUsage `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunMultiDocIntroductionResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunMultiDocIntroductionResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunMultiDocIntroductionResponseBodyPayload) GetOutput() *RunMultiDocIntroductionResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunMultiDocIntroductionResponseBodyPayload) GetUsage() *RunMultiDocIntroductionResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunMultiDocIntroductionResponseBodyPayload) SetOutput(v *RunMultiDocIntroductionResponseBodyPayloadOutput) *RunMultiDocIntroductionResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunMultiDocIntroductionResponseBodyPayload) SetUsage(v *RunMultiDocIntroductionResponseBodyPayloadUsage) *RunMultiDocIntroductionResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunMultiDocIntroductionResponseBodyPayload) Validate() error {
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

type RunMultiDocIntroductionResponseBodyPayloadOutput struct {
	// Key point information.
	KeyPoints []*RunMultiDocIntroductionResponseBodyPayloadOutputKeyPoints `json:"KeyPoints,omitempty" xml:"KeyPoints,omitempty" type:"Repeated"`
	// Outline-style summary.
	//
	// example:
	//
	// 大纲摘要内容
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s RunMultiDocIntroductionResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunMultiDocIntroductionResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunMultiDocIntroductionResponseBodyPayloadOutput) GetKeyPoints() []*RunMultiDocIntroductionResponseBodyPayloadOutputKeyPoints {
	return s.KeyPoints
}

func (s *RunMultiDocIntroductionResponseBodyPayloadOutput) GetSummary() *string {
	return s.Summary
}

func (s *RunMultiDocIntroductionResponseBodyPayloadOutput) SetKeyPoints(v []*RunMultiDocIntroductionResponseBodyPayloadOutputKeyPoints) *RunMultiDocIntroductionResponseBodyPayloadOutput {
	s.KeyPoints = v
	return s
}

func (s *RunMultiDocIntroductionResponseBodyPayloadOutput) SetSummary(v string) *RunMultiDocIntroductionResponseBodyPayloadOutput {
	s.Summary = &v
	return s
}

func (s *RunMultiDocIntroductionResponseBodyPayloadOutput) Validate() error {
	if s.KeyPoints != nil {
		for _, item := range s.KeyPoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunMultiDocIntroductionResponseBodyPayloadOutputKeyPoints struct {
	// Key point.
	//
	// example:
	//
	// 关键点信息
	KeyPoint *string `json:"KeyPoint,omitempty" xml:"KeyPoint,omitempty"`
	// Source of the information.
	//
	// example:
	//
	// 信息来源
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s RunMultiDocIntroductionResponseBodyPayloadOutputKeyPoints) String() string {
	return dara.Prettify(s)
}

func (s RunMultiDocIntroductionResponseBodyPayloadOutputKeyPoints) GoString() string {
	return s.String()
}

func (s *RunMultiDocIntroductionResponseBodyPayloadOutputKeyPoints) GetKeyPoint() *string {
	return s.KeyPoint
}

func (s *RunMultiDocIntroductionResponseBodyPayloadOutputKeyPoints) GetSource() *string {
	return s.Source
}

func (s *RunMultiDocIntroductionResponseBodyPayloadOutputKeyPoints) SetKeyPoint(v string) *RunMultiDocIntroductionResponseBodyPayloadOutputKeyPoints {
	s.KeyPoint = &v
	return s
}

func (s *RunMultiDocIntroductionResponseBodyPayloadOutputKeyPoints) SetSource(v string) *RunMultiDocIntroductionResponseBodyPayloadOutputKeyPoints {
	s.Source = &v
	return s
}

func (s *RunMultiDocIntroductionResponseBodyPayloadOutputKeyPoints) Validate() error {
	return dara.Validate(s)
}

type RunMultiDocIntroductionResponseBodyPayloadUsage struct {
	// Number of input tokens.
	//
	// example:
	//
	// 65
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// Number of output tokens.
	//
	// example:
	//
	// 100
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// Total number of tokens.
	//
	// example:
	//
	// 165
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunMultiDocIntroductionResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunMultiDocIntroductionResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunMultiDocIntroductionResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunMultiDocIntroductionResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunMultiDocIntroductionResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunMultiDocIntroductionResponseBodyPayloadUsage) SetInputTokens(v int64) *RunMultiDocIntroductionResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunMultiDocIntroductionResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunMultiDocIntroductionResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunMultiDocIntroductionResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunMultiDocIntroductionResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunMultiDocIntroductionResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
