// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunEssayCorrectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeader(v *RunEssayCorrectionResponseBodyHeader) *RunEssayCorrectionResponseBody
	GetHeader() *RunEssayCorrectionResponseBodyHeader
	SetPayload(v *RunEssayCorrectionResponseBodyPayload) *RunEssayCorrectionResponseBody
	GetPayload() *RunEssayCorrectionResponseBodyPayload
	SetRequestId(v string) *RunEssayCorrectionResponseBody
	GetRequestId() *string
}

type RunEssayCorrectionResponseBody struct {
	Header  *RunEssayCorrectionResponseBodyHeader  `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload *RunEssayCorrectionResponseBodyPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 117F5ABE-CF02-5502-9A3F-E56BC9081A64
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RunEssayCorrectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunEssayCorrectionResponseBody) GoString() string {
	return s.String()
}

func (s *RunEssayCorrectionResponseBody) GetHeader() *RunEssayCorrectionResponseBodyHeader {
	return s.Header
}

func (s *RunEssayCorrectionResponseBody) GetPayload() *RunEssayCorrectionResponseBodyPayload {
	return s.Payload
}

func (s *RunEssayCorrectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunEssayCorrectionResponseBody) SetHeader(v *RunEssayCorrectionResponseBodyHeader) *RunEssayCorrectionResponseBody {
	s.Header = v
	return s
}

func (s *RunEssayCorrectionResponseBody) SetPayload(v *RunEssayCorrectionResponseBodyPayload) *RunEssayCorrectionResponseBody {
	s.Payload = v
	return s
}

func (s *RunEssayCorrectionResponseBody) SetRequestId(v string) *RunEssayCorrectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunEssayCorrectionResponseBody) Validate() error {
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

type RunEssayCorrectionResponseBodyHeader struct {
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

func (s RunEssayCorrectionResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunEssayCorrectionResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunEssayCorrectionResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunEssayCorrectionResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunEssayCorrectionResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunEssayCorrectionResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunEssayCorrectionResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunEssayCorrectionResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunEssayCorrectionResponseBodyHeader) SetErrorCode(v string) *RunEssayCorrectionResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunEssayCorrectionResponseBodyHeader) SetErrorMessage(v string) *RunEssayCorrectionResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunEssayCorrectionResponseBodyHeader) SetEvent(v string) *RunEssayCorrectionResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunEssayCorrectionResponseBodyHeader) SetSessionId(v string) *RunEssayCorrectionResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunEssayCorrectionResponseBodyHeader) SetTaskId(v string) *RunEssayCorrectionResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunEssayCorrectionResponseBodyHeader) SetTraceId(v string) *RunEssayCorrectionResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunEssayCorrectionResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunEssayCorrectionResponseBodyPayload struct {
	Output *RunEssayCorrectionResponseBodyPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Usage  *RunEssayCorrectionResponseBodyPayloadUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunEssayCorrectionResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunEssayCorrectionResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunEssayCorrectionResponseBodyPayload) GetOutput() *RunEssayCorrectionResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunEssayCorrectionResponseBodyPayload) GetUsage() *RunEssayCorrectionResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunEssayCorrectionResponseBodyPayload) SetOutput(v *RunEssayCorrectionResponseBodyPayloadOutput) *RunEssayCorrectionResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunEssayCorrectionResponseBodyPayload) SetUsage(v *RunEssayCorrectionResponseBodyPayloadUsage) *RunEssayCorrectionResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunEssayCorrectionResponseBodyPayload) Validate() error {
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

type RunEssayCorrectionResponseBodyPayloadOutput struct {
	DimensionResults []*RunEssayCorrectionResponseBodyPayloadOutputDimensionResults `json:"dimensionResults,omitempty" xml:"dimensionResults,omitempty" type:"Repeated"`
	// example:
	//
	// 整体表现良好，建议在论述深度上进一步加强。
	OverallComment *string `json:"overallComment,omitempty" xml:"overallComment,omitempty"`
	// example:
	//
	// 首先分析文章结构，发现开头、正文、结尾完整...
	ReasoningContent *string `json:"reasoningContent,omitempty" xml:"reasoningContent,omitempty"`
	// example:
	//
	// 50
	Score *int32 `json:"score,omitempty" xml:"score,omitempty"`
	// example:
	//
	// 本文整体结构清晰，语言流畅...
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s RunEssayCorrectionResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunEssayCorrectionResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunEssayCorrectionResponseBodyPayloadOutput) GetDimensionResults() []*RunEssayCorrectionResponseBodyPayloadOutputDimensionResults {
	return s.DimensionResults
}

func (s *RunEssayCorrectionResponseBodyPayloadOutput) GetOverallComment() *string {
	return s.OverallComment
}

func (s *RunEssayCorrectionResponseBodyPayloadOutput) GetReasoningContent() *string {
	return s.ReasoningContent
}

func (s *RunEssayCorrectionResponseBodyPayloadOutput) GetScore() *int32 {
	return s.Score
}

func (s *RunEssayCorrectionResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunEssayCorrectionResponseBodyPayloadOutput) SetDimensionResults(v []*RunEssayCorrectionResponseBodyPayloadOutputDimensionResults) *RunEssayCorrectionResponseBodyPayloadOutput {
	s.DimensionResults = v
	return s
}

func (s *RunEssayCorrectionResponseBodyPayloadOutput) SetOverallComment(v string) *RunEssayCorrectionResponseBodyPayloadOutput {
	s.OverallComment = &v
	return s
}

func (s *RunEssayCorrectionResponseBodyPayloadOutput) SetReasoningContent(v string) *RunEssayCorrectionResponseBodyPayloadOutput {
	s.ReasoningContent = &v
	return s
}

func (s *RunEssayCorrectionResponseBodyPayloadOutput) SetScore(v int32) *RunEssayCorrectionResponseBodyPayloadOutput {
	s.Score = &v
	return s
}

func (s *RunEssayCorrectionResponseBodyPayloadOutput) SetText(v string) *RunEssayCorrectionResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunEssayCorrectionResponseBodyPayloadOutput) Validate() error {
	if s.DimensionResults != nil {
		for _, item := range s.DimensionResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunEssayCorrectionResponseBodyPayloadOutputDimensionResults struct {
	// example:
	//
	// 文章内容较为完整，涵盖了题目的核心要求，但部分论述略显简略。
	Analysis *string `json:"analysis,omitempty" xml:"analysis,omitempty"`
	// example:
	//
	// 30
	MaxScore *float64 `json:"maxScore,omitempty" xml:"maxScore,omitempty"`
	// example:
	//
	// 内容完整度
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 25.5
	Score *float64 `json:"score,omitempty" xml:"score,omitempty"`
}

func (s RunEssayCorrectionResponseBodyPayloadOutputDimensionResults) String() string {
	return dara.Prettify(s)
}

func (s RunEssayCorrectionResponseBodyPayloadOutputDimensionResults) GoString() string {
	return s.String()
}

func (s *RunEssayCorrectionResponseBodyPayloadOutputDimensionResults) GetAnalysis() *string {
	return s.Analysis
}

func (s *RunEssayCorrectionResponseBodyPayloadOutputDimensionResults) GetMaxScore() *float64 {
	return s.MaxScore
}

func (s *RunEssayCorrectionResponseBodyPayloadOutputDimensionResults) GetName() *string {
	return s.Name
}

func (s *RunEssayCorrectionResponseBodyPayloadOutputDimensionResults) GetScore() *float64 {
	return s.Score
}

func (s *RunEssayCorrectionResponseBodyPayloadOutputDimensionResults) SetAnalysis(v string) *RunEssayCorrectionResponseBodyPayloadOutputDimensionResults {
	s.Analysis = &v
	return s
}

func (s *RunEssayCorrectionResponseBodyPayloadOutputDimensionResults) SetMaxScore(v float64) *RunEssayCorrectionResponseBodyPayloadOutputDimensionResults {
	s.MaxScore = &v
	return s
}

func (s *RunEssayCorrectionResponseBodyPayloadOutputDimensionResults) SetName(v string) *RunEssayCorrectionResponseBodyPayloadOutputDimensionResults {
	s.Name = &v
	return s
}

func (s *RunEssayCorrectionResponseBodyPayloadOutputDimensionResults) SetScore(v float64) *RunEssayCorrectionResponseBodyPayloadOutputDimensionResults {
	s.Score = &v
	return s
}

func (s *RunEssayCorrectionResponseBodyPayloadOutputDimensionResults) Validate() error {
	return dara.Validate(s)
}

type RunEssayCorrectionResponseBodyPayloadUsage struct {
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

func (s RunEssayCorrectionResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunEssayCorrectionResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunEssayCorrectionResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunEssayCorrectionResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunEssayCorrectionResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunEssayCorrectionResponseBodyPayloadUsage) SetInputTokens(v int64) *RunEssayCorrectionResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunEssayCorrectionResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunEssayCorrectionResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunEssayCorrectionResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunEssayCorrectionResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunEssayCorrectionResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
