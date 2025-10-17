// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunMarketingInformationWritingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnd(v bool) *RunMarketingInformationWritingResponseBody
	GetEnd() *bool
	SetHeader(v *RunMarketingInformationWritingResponseBodyHeader) *RunMarketingInformationWritingResponseBody
	GetHeader() *RunMarketingInformationWritingResponseBodyHeader
	SetPayload(v *RunMarketingInformationWritingResponseBodyPayload) *RunMarketingInformationWritingResponseBody
	GetPayload() *RunMarketingInformationWritingResponseBodyPayload
}

type RunMarketingInformationWritingResponseBody struct {
	// example:
	//
	// 2024-06-21T10:29:52+08:00
	End     *bool                                              `json:"end,omitempty" xml:"end,omitempty"`
	Header  *RunMarketingInformationWritingResponseBodyHeader  `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload *RunMarketingInformationWritingResponseBodyPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
}

func (s RunMarketingInformationWritingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunMarketingInformationWritingResponseBody) GoString() string {
	return s.String()
}

func (s *RunMarketingInformationWritingResponseBody) GetEnd() *bool {
	return s.End
}

func (s *RunMarketingInformationWritingResponseBody) GetHeader() *RunMarketingInformationWritingResponseBodyHeader {
	return s.Header
}

func (s *RunMarketingInformationWritingResponseBody) GetPayload() *RunMarketingInformationWritingResponseBodyPayload {
	return s.Payload
}

func (s *RunMarketingInformationWritingResponseBody) SetEnd(v bool) *RunMarketingInformationWritingResponseBody {
	s.End = &v
	return s
}

func (s *RunMarketingInformationWritingResponseBody) SetHeader(v *RunMarketingInformationWritingResponseBodyHeader) *RunMarketingInformationWritingResponseBody {
	s.Header = v
	return s
}

func (s *RunMarketingInformationWritingResponseBody) SetPayload(v *RunMarketingInformationWritingResponseBodyPayload) *RunMarketingInformationWritingResponseBody {
	s.Payload = v
	return s
}

func (s *RunMarketingInformationWritingResponseBody) Validate() error {
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

type RunMarketingInformationWritingResponseBodyHeader struct {
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// result-generated
	Event     *string `json:"event,omitempty" xml:"event,omitempty"`
	EventInfo *string `json:"eventInfo,omitempty" xml:"eventInfo,omitempty"`
	// example:
	//
	// 436BC5AE-0573-59D8-9803-6B5FDCD3BBA1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// uqubxgqzlnf4exfektij032lgb3yvix678p232n56387aqxdo4uutdt9wstqzovvz6j3ho7wnbgye785u79yn5q3euqmsvzmqdn3nmfq2826oscjvsi43kof8b8uxufpp1x97jjukk6jd3183hy8ni6hqpskuhuascpd
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// 13312125943232
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 213e20e517049392478441155e8b2a
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s RunMarketingInformationWritingResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunMarketingInformationWritingResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunMarketingInformationWritingResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunMarketingInformationWritingResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunMarketingInformationWritingResponseBodyHeader) GetEventInfo() *string {
	return s.EventInfo
}

func (s *RunMarketingInformationWritingResponseBodyHeader) GetRequestId() *string {
	return s.RequestId
}

func (s *RunMarketingInformationWritingResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunMarketingInformationWritingResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunMarketingInformationWritingResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunMarketingInformationWritingResponseBodyHeader) SetErrorMessage(v string) *RunMarketingInformationWritingResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunMarketingInformationWritingResponseBodyHeader) SetEvent(v string) *RunMarketingInformationWritingResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunMarketingInformationWritingResponseBodyHeader) SetEventInfo(v string) *RunMarketingInformationWritingResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunMarketingInformationWritingResponseBodyHeader) SetRequestId(v string) *RunMarketingInformationWritingResponseBodyHeader {
	s.RequestId = &v
	return s
}

func (s *RunMarketingInformationWritingResponseBodyHeader) SetSessionId(v string) *RunMarketingInformationWritingResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunMarketingInformationWritingResponseBodyHeader) SetTaskId(v string) *RunMarketingInformationWritingResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunMarketingInformationWritingResponseBodyHeader) SetTraceId(v string) *RunMarketingInformationWritingResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunMarketingInformationWritingResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunMarketingInformationWritingResponseBodyPayload struct {
	Output *RunMarketingInformationWritingResponseBodyPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Usage  *RunMarketingInformationWritingResponseBodyPayloadUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunMarketingInformationWritingResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunMarketingInformationWritingResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunMarketingInformationWritingResponseBodyPayload) GetOutput() *RunMarketingInformationWritingResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunMarketingInformationWritingResponseBodyPayload) GetUsage() *RunMarketingInformationWritingResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunMarketingInformationWritingResponseBodyPayload) SetOutput(v *RunMarketingInformationWritingResponseBodyPayloadOutput) *RunMarketingInformationWritingResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunMarketingInformationWritingResponseBodyPayload) SetUsage(v *RunMarketingInformationWritingResponseBodyPayloadUsage) *RunMarketingInformationWritingResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunMarketingInformationWritingResponseBodyPayload) Validate() error {
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

type RunMarketingInformationWritingResponseBodyPayloadOutput struct {
	ReasonContent *string `json:"reasonContent,omitempty" xml:"reasonContent,omitempty"`
	Text          *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s RunMarketingInformationWritingResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunMarketingInformationWritingResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunMarketingInformationWritingResponseBodyPayloadOutput) GetReasonContent() *string {
	return s.ReasonContent
}

func (s *RunMarketingInformationWritingResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunMarketingInformationWritingResponseBodyPayloadOutput) SetReasonContent(v string) *RunMarketingInformationWritingResponseBodyPayloadOutput {
	s.ReasonContent = &v
	return s
}

func (s *RunMarketingInformationWritingResponseBodyPayloadOutput) SetText(v string) *RunMarketingInformationWritingResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunMarketingInformationWritingResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunMarketingInformationWritingResponseBodyPayloadUsage struct {
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

func (s RunMarketingInformationWritingResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunMarketingInformationWritingResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunMarketingInformationWritingResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunMarketingInformationWritingResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunMarketingInformationWritingResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunMarketingInformationWritingResponseBodyPayloadUsage) SetInputTokens(v int64) *RunMarketingInformationWritingResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunMarketingInformationWritingResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunMarketingInformationWritingResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunMarketingInformationWritingResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunMarketingInformationWritingResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunMarketingInformationWritingResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
