// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunMarketingInformationExtractResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnd(v bool) *RunMarketingInformationExtractResponseBody
	GetEnd() *bool
	SetHeader(v *RunMarketingInformationExtractResponseBodyHeader) *RunMarketingInformationExtractResponseBody
	GetHeader() *RunMarketingInformationExtractResponseBodyHeader
	SetPayload(v *RunMarketingInformationExtractResponseBodyPayload) *RunMarketingInformationExtractResponseBody
	GetPayload() *RunMarketingInformationExtractResponseBodyPayload
}

type RunMarketingInformationExtractResponseBody struct {
	// example:
	//
	// {\\"TimeZone\\": \\"Asia/Shanghai\\", \\"DateTime\\": \\"2024-03-07T17:00:09+08:00\\"}
	End     *bool                                              `json:"end,omitempty" xml:"end,omitempty"`
	Header  *RunMarketingInformationExtractResponseBodyHeader  `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload *RunMarketingInformationExtractResponseBodyPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
}

func (s RunMarketingInformationExtractResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunMarketingInformationExtractResponseBody) GoString() string {
	return s.String()
}

func (s *RunMarketingInformationExtractResponseBody) GetEnd() *bool {
	return s.End
}

func (s *RunMarketingInformationExtractResponseBody) GetHeader() *RunMarketingInformationExtractResponseBodyHeader {
	return s.Header
}

func (s *RunMarketingInformationExtractResponseBody) GetPayload() *RunMarketingInformationExtractResponseBodyPayload {
	return s.Payload
}

func (s *RunMarketingInformationExtractResponseBody) SetEnd(v bool) *RunMarketingInformationExtractResponseBody {
	s.End = &v
	return s
}

func (s *RunMarketingInformationExtractResponseBody) SetHeader(v *RunMarketingInformationExtractResponseBodyHeader) *RunMarketingInformationExtractResponseBody {
	s.Header = v
	return s
}

func (s *RunMarketingInformationExtractResponseBody) SetPayload(v *RunMarketingInformationExtractResponseBodyPayload) *RunMarketingInformationExtractResponseBody {
	s.Payload = v
	return s
}

func (s *RunMarketingInformationExtractResponseBody) Validate() error {
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

type RunMarketingInformationExtractResponseBodyHeader struct {
	// example:
	//
	// result-generated
	Event     *string `json:"event,omitempty" xml:"event,omitempty"`
	EventInfo *string `json:"eventInfo,omitempty" xml:"eventInfo,omitempty"`
	// example:
	//
	// F08C71C0-9399-548C-838B-1DA01DE211B0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 121dlsga4o7golrl1hojazg0u9lvytjc17ebgzzj2u4zukgh122tfg7wj1e6a1vcowy1ewzinauxriai9atcr6r323mm9ddbr0bg5m61ij8hxnf8664tstlfkfol6m8luc4shs3gums7l46uauyy0xndqmhdjtdon6coyhb4x17bo762bg9e3tb2geufg2
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// 12826092918145
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 2150432017236011824686132ecdbc
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s RunMarketingInformationExtractResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunMarketingInformationExtractResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunMarketingInformationExtractResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunMarketingInformationExtractResponseBodyHeader) GetEventInfo() *string {
	return s.EventInfo
}

func (s *RunMarketingInformationExtractResponseBodyHeader) GetRequestId() *string {
	return s.RequestId
}

func (s *RunMarketingInformationExtractResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunMarketingInformationExtractResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunMarketingInformationExtractResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunMarketingInformationExtractResponseBodyHeader) SetEvent(v string) *RunMarketingInformationExtractResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunMarketingInformationExtractResponseBodyHeader) SetEventInfo(v string) *RunMarketingInformationExtractResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunMarketingInformationExtractResponseBodyHeader) SetRequestId(v string) *RunMarketingInformationExtractResponseBodyHeader {
	s.RequestId = &v
	return s
}

func (s *RunMarketingInformationExtractResponseBodyHeader) SetSessionId(v string) *RunMarketingInformationExtractResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunMarketingInformationExtractResponseBodyHeader) SetTaskId(v string) *RunMarketingInformationExtractResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunMarketingInformationExtractResponseBodyHeader) SetTraceId(v string) *RunMarketingInformationExtractResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunMarketingInformationExtractResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunMarketingInformationExtractResponseBodyPayload struct {
	Output *RunMarketingInformationExtractResponseBodyPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Usage  *RunMarketingInformationExtractResponseBodyPayloadUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunMarketingInformationExtractResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunMarketingInformationExtractResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunMarketingInformationExtractResponseBodyPayload) GetOutput() *RunMarketingInformationExtractResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunMarketingInformationExtractResponseBodyPayload) GetUsage() *RunMarketingInformationExtractResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunMarketingInformationExtractResponseBodyPayload) SetOutput(v *RunMarketingInformationExtractResponseBodyPayloadOutput) *RunMarketingInformationExtractResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunMarketingInformationExtractResponseBodyPayload) SetUsage(v *RunMarketingInformationExtractResponseBodyPayloadUsage) *RunMarketingInformationExtractResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunMarketingInformationExtractResponseBodyPayload) Validate() error {
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

type RunMarketingInformationExtractResponseBodyPayloadOutput struct {
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s RunMarketingInformationExtractResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunMarketingInformationExtractResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunMarketingInformationExtractResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunMarketingInformationExtractResponseBodyPayloadOutput) SetText(v string) *RunMarketingInformationExtractResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunMarketingInformationExtractResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunMarketingInformationExtractResponseBodyPayloadUsage struct {
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

func (s RunMarketingInformationExtractResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunMarketingInformationExtractResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunMarketingInformationExtractResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunMarketingInformationExtractResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunMarketingInformationExtractResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunMarketingInformationExtractResponseBodyPayloadUsage) SetInputTokens(v int64) *RunMarketingInformationExtractResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunMarketingInformationExtractResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunMarketingInformationExtractResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunMarketingInformationExtractResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunMarketingInformationExtractResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunMarketingInformationExtractResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
