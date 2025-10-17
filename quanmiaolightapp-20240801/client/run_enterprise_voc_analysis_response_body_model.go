// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunEnterpriseVocAnalysisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeader(v *RunEnterpriseVocAnalysisResponseBodyHeader) *RunEnterpriseVocAnalysisResponseBody
	GetHeader() *RunEnterpriseVocAnalysisResponseBodyHeader
	SetPayload(v *RunEnterpriseVocAnalysisResponseBodyPayload) *RunEnterpriseVocAnalysisResponseBody
	GetPayload() *RunEnterpriseVocAnalysisResponseBodyPayload
	SetRequestId(v string) *RunEnterpriseVocAnalysisResponseBody
	GetRequestId() *string
}

type RunEnterpriseVocAnalysisResponseBody struct {
	Header  *RunEnterpriseVocAnalysisResponseBodyHeader  `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload *RunEnterpriseVocAnalysisResponseBodyPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 49483FFC-0CB9-5163-8D3E-234E276E6DA8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RunEnterpriseVocAnalysisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunEnterpriseVocAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *RunEnterpriseVocAnalysisResponseBody) GetHeader() *RunEnterpriseVocAnalysisResponseBodyHeader {
	return s.Header
}

func (s *RunEnterpriseVocAnalysisResponseBody) GetPayload() *RunEnterpriseVocAnalysisResponseBodyPayload {
	return s.Payload
}

func (s *RunEnterpriseVocAnalysisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunEnterpriseVocAnalysisResponseBody) SetHeader(v *RunEnterpriseVocAnalysisResponseBodyHeader) *RunEnterpriseVocAnalysisResponseBody {
	s.Header = v
	return s
}

func (s *RunEnterpriseVocAnalysisResponseBody) SetPayload(v *RunEnterpriseVocAnalysisResponseBodyPayload) *RunEnterpriseVocAnalysisResponseBody {
	s.Payload = v
	return s
}

func (s *RunEnterpriseVocAnalysisResponseBody) SetRequestId(v string) *RunEnterpriseVocAnalysisResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunEnterpriseVocAnalysisResponseBody) Validate() error {
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

type RunEnterpriseVocAnalysisResponseBodyHeader struct {
	// example:
	//
	// AccessForbidden
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// task-finished
	Event *string `json:"event,omitempty" xml:"event,omitempty"`
	// example:
	//
	// xxxx
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// xxxx
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// xxxxx
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s RunEnterpriseVocAnalysisResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunEnterpriseVocAnalysisResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunEnterpriseVocAnalysisResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunEnterpriseVocAnalysisResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunEnterpriseVocAnalysisResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunEnterpriseVocAnalysisResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunEnterpriseVocAnalysisResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunEnterpriseVocAnalysisResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunEnterpriseVocAnalysisResponseBodyHeader) SetErrorCode(v string) *RunEnterpriseVocAnalysisResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunEnterpriseVocAnalysisResponseBodyHeader) SetErrorMessage(v string) *RunEnterpriseVocAnalysisResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunEnterpriseVocAnalysisResponseBodyHeader) SetEvent(v string) *RunEnterpriseVocAnalysisResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunEnterpriseVocAnalysisResponseBodyHeader) SetSessionId(v string) *RunEnterpriseVocAnalysisResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunEnterpriseVocAnalysisResponseBodyHeader) SetTaskId(v string) *RunEnterpriseVocAnalysisResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunEnterpriseVocAnalysisResponseBodyHeader) SetTraceId(v string) *RunEnterpriseVocAnalysisResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunEnterpriseVocAnalysisResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunEnterpriseVocAnalysisResponseBodyPayload struct {
	Output *RunEnterpriseVocAnalysisResponseBodyPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Usage  *RunEnterpriseVocAnalysisResponseBodyPayloadUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunEnterpriseVocAnalysisResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunEnterpriseVocAnalysisResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunEnterpriseVocAnalysisResponseBodyPayload) GetOutput() *RunEnterpriseVocAnalysisResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunEnterpriseVocAnalysisResponseBodyPayload) GetUsage() *RunEnterpriseVocAnalysisResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunEnterpriseVocAnalysisResponseBodyPayload) SetOutput(v *RunEnterpriseVocAnalysisResponseBodyPayloadOutput) *RunEnterpriseVocAnalysisResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunEnterpriseVocAnalysisResponseBodyPayload) SetUsage(v *RunEnterpriseVocAnalysisResponseBodyPayloadUsage) *RunEnterpriseVocAnalysisResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunEnterpriseVocAnalysisResponseBodyPayload) Validate() error {
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

type RunEnterpriseVocAnalysisResponseBodyPayloadOutput struct {
	FilterResult  *RunEnterpriseVocAnalysisResponseBodyPayloadOutputFilterResult `json:"filterResult,omitempty" xml:"filterResult,omitempty" type:"Struct"`
	ReasonContent *string                                                        `json:"reasonContent,omitempty" xml:"reasonContent,omitempty"`
	Text          *string                                                        `json:"text,omitempty" xml:"text,omitempty"`
}

func (s RunEnterpriseVocAnalysisResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunEnterpriseVocAnalysisResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunEnterpriseVocAnalysisResponseBodyPayloadOutput) GetFilterResult() *RunEnterpriseVocAnalysisResponseBodyPayloadOutputFilterResult {
	return s.FilterResult
}

func (s *RunEnterpriseVocAnalysisResponseBodyPayloadOutput) GetReasonContent() *string {
	return s.ReasonContent
}

func (s *RunEnterpriseVocAnalysisResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunEnterpriseVocAnalysisResponseBodyPayloadOutput) SetFilterResult(v *RunEnterpriseVocAnalysisResponseBodyPayloadOutputFilterResult) *RunEnterpriseVocAnalysisResponseBodyPayloadOutput {
	s.FilterResult = v
	return s
}

func (s *RunEnterpriseVocAnalysisResponseBodyPayloadOutput) SetReasonContent(v string) *RunEnterpriseVocAnalysisResponseBodyPayloadOutput {
	s.ReasonContent = &v
	return s
}

func (s *RunEnterpriseVocAnalysisResponseBodyPayloadOutput) SetText(v string) *RunEnterpriseVocAnalysisResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunEnterpriseVocAnalysisResponseBodyPayloadOutput) Validate() error {
	if s.FilterResult != nil {
		if err := s.FilterResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunEnterpriseVocAnalysisResponseBodyPayloadOutputFilterResult struct {
	FilterResults []*RunEnterpriseVocAnalysisResponseBodyPayloadOutputFilterResultFilterResults `json:"filterResults,omitempty" xml:"filterResults,omitempty" type:"Repeated"`
}

func (s RunEnterpriseVocAnalysisResponseBodyPayloadOutputFilterResult) String() string {
	return dara.Prettify(s)
}

func (s RunEnterpriseVocAnalysisResponseBodyPayloadOutputFilterResult) GoString() string {
	return s.String()
}

func (s *RunEnterpriseVocAnalysisResponseBodyPayloadOutputFilterResult) GetFilterResults() []*RunEnterpriseVocAnalysisResponseBodyPayloadOutputFilterResultFilterResults {
	return s.FilterResults
}

func (s *RunEnterpriseVocAnalysisResponseBodyPayloadOutputFilterResult) SetFilterResults(v []*RunEnterpriseVocAnalysisResponseBodyPayloadOutputFilterResultFilterResults) *RunEnterpriseVocAnalysisResponseBodyPayloadOutputFilterResult {
	s.FilterResults = v
	return s
}

func (s *RunEnterpriseVocAnalysisResponseBodyPayloadOutputFilterResult) Validate() error {
	if s.FilterResults != nil {
		for _, item := range s.FilterResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunEnterpriseVocAnalysisResponseBodyPayloadOutputFilterResultFilterResults struct {
	// example:
	//
	// true
	Hit      *bool   `json:"hit,omitempty" xml:"hit,omitempty"`
	TagName  *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s RunEnterpriseVocAnalysisResponseBodyPayloadOutputFilterResultFilterResults) String() string {
	return dara.Prettify(s)
}

func (s RunEnterpriseVocAnalysisResponseBodyPayloadOutputFilterResultFilterResults) GoString() string {
	return s.String()
}

func (s *RunEnterpriseVocAnalysisResponseBodyPayloadOutputFilterResultFilterResults) GetHit() *bool {
	return s.Hit
}

func (s *RunEnterpriseVocAnalysisResponseBodyPayloadOutputFilterResultFilterResults) GetTagName() *string {
	return s.TagName
}

func (s *RunEnterpriseVocAnalysisResponseBodyPayloadOutputFilterResultFilterResults) GetTagValue() *string {
	return s.TagValue
}

func (s *RunEnterpriseVocAnalysisResponseBodyPayloadOutputFilterResultFilterResults) SetHit(v bool) *RunEnterpriseVocAnalysisResponseBodyPayloadOutputFilterResultFilterResults {
	s.Hit = &v
	return s
}

func (s *RunEnterpriseVocAnalysisResponseBodyPayloadOutputFilterResultFilterResults) SetTagName(v string) *RunEnterpriseVocAnalysisResponseBodyPayloadOutputFilterResultFilterResults {
	s.TagName = &v
	return s
}

func (s *RunEnterpriseVocAnalysisResponseBodyPayloadOutputFilterResultFilterResults) SetTagValue(v string) *RunEnterpriseVocAnalysisResponseBodyPayloadOutputFilterResultFilterResults {
	s.TagValue = &v
	return s
}

func (s *RunEnterpriseVocAnalysisResponseBodyPayloadOutputFilterResultFilterResults) Validate() error {
	return dara.Validate(s)
}

type RunEnterpriseVocAnalysisResponseBodyPayloadUsage struct {
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

func (s RunEnterpriseVocAnalysisResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunEnterpriseVocAnalysisResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunEnterpriseVocAnalysisResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunEnterpriseVocAnalysisResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunEnterpriseVocAnalysisResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunEnterpriseVocAnalysisResponseBodyPayloadUsage) SetInputTokens(v int64) *RunEnterpriseVocAnalysisResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunEnterpriseVocAnalysisResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunEnterpriseVocAnalysisResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunEnterpriseVocAnalysisResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunEnterpriseVocAnalysisResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunEnterpriseVocAnalysisResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
