// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetTaskResultResponseBodyData) *GetTaskResultResponseBody
	GetData() *GetTaskResultResponseBodyData
	SetRequestId(v string) *GetTaskResultResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetTaskResultResponseBody
	GetSuccess() *string
}

type GetTaskResultResponseBody struct {
	Data *GetTaskResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 968A8634-FA2C-5381-9B3E-C552DED7E8BF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetTaskResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskResultResponseBody) GetData() *GetTaskResultResponseBodyData {
	return s.Data
}

func (s *GetTaskResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskResultResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetTaskResultResponseBody) SetData(v *GetTaskResultResponseBodyData) *GetTaskResultResponseBody {
	s.Data = v
	return s
}

func (s *GetTaskResultResponseBody) SetRequestId(v string) *GetTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskResultResponseBody) SetSuccess(v string) *GetTaskResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetTaskResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTaskResultResponseBodyData struct {
	AsrResult        []*GetTaskResultResponseBodyDataAsrResult `json:"asrResult,omitempty" xml:"asrResult,omitempty" type:"Repeated"`
	Extra            *string                                   `json:"extra,omitempty" xml:"extra,omitempty"`
	RagErrorMessage  *string                                   `json:"ragErrorMessage,omitempty" xml:"ragErrorMessage,omitempty"`
	RagResult        *string                                   `json:"ragResult,omitempty" xml:"ragResult,omitempty"`
	RagStatus        *string                                   `json:"ragStatus,omitempty" xml:"ragStatus,omitempty"`
	TaskErrorMessage *string                                   `json:"taskErrorMessage,omitempty" xml:"taskErrorMessage,omitempty"`
	// example:
	//
	// 20240905-********-93E9-5D45-B4EF-045743A34071
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// FINISH
	TaskStatus *string                             `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
	Text       *string                             `json:"text,omitempty" xml:"text,omitempty"`
	Usage      *GetTaskResultResponseBodyDataUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetTaskResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTaskResultResponseBodyData) GetAsrResult() []*GetTaskResultResponseBodyDataAsrResult {
	return s.AsrResult
}

func (s *GetTaskResultResponseBodyData) GetExtra() *string {
	return s.Extra
}

func (s *GetTaskResultResponseBodyData) GetRagErrorMessage() *string {
	return s.RagErrorMessage
}

func (s *GetTaskResultResponseBodyData) GetRagResult() *string {
	return s.RagResult
}

func (s *GetTaskResultResponseBodyData) GetRagStatus() *string {
	return s.RagStatus
}

func (s *GetTaskResultResponseBodyData) GetTaskErrorMessage() *string {
	return s.TaskErrorMessage
}

func (s *GetTaskResultResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTaskResultResponseBodyData) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *GetTaskResultResponseBodyData) GetText() *string {
	return s.Text
}

func (s *GetTaskResultResponseBodyData) GetUsage() *GetTaskResultResponseBodyDataUsage {
	return s.Usage
}

func (s *GetTaskResultResponseBodyData) SetAsrResult(v []*GetTaskResultResponseBodyDataAsrResult) *GetTaskResultResponseBodyData {
	s.AsrResult = v
	return s
}

func (s *GetTaskResultResponseBodyData) SetExtra(v string) *GetTaskResultResponseBodyData {
	s.Extra = &v
	return s
}

func (s *GetTaskResultResponseBodyData) SetRagErrorMessage(v string) *GetTaskResultResponseBodyData {
	s.RagErrorMessage = &v
	return s
}

func (s *GetTaskResultResponseBodyData) SetRagResult(v string) *GetTaskResultResponseBodyData {
	s.RagResult = &v
	return s
}

func (s *GetTaskResultResponseBodyData) SetRagStatus(v string) *GetTaskResultResponseBodyData {
	s.RagStatus = &v
	return s
}

func (s *GetTaskResultResponseBodyData) SetTaskErrorMessage(v string) *GetTaskResultResponseBodyData {
	s.TaskErrorMessage = &v
	return s
}

func (s *GetTaskResultResponseBodyData) SetTaskId(v string) *GetTaskResultResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetTaskResultResponseBodyData) SetTaskStatus(v string) *GetTaskResultResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *GetTaskResultResponseBodyData) SetText(v string) *GetTaskResultResponseBodyData {
	s.Text = &v
	return s
}

func (s *GetTaskResultResponseBodyData) SetUsage(v *GetTaskResultResponseBodyDataUsage) *GetTaskResultResponseBodyData {
	s.Usage = v
	return s
}

func (s *GetTaskResultResponseBodyData) Validate() error {
	if s.AsrResult != nil {
		for _, item := range s.AsrResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTaskResultResponseBodyDataAsrResult struct {
	Begin        *int64  `json:"begin,omitempty" xml:"begin,omitempty"`
	EmotionValue *int32  `json:"emotionValue,omitempty" xml:"emotionValue,omitempty"`
	End          *int64  `json:"end,omitempty" xml:"end,omitempty"`
	Role         *string `json:"role,omitempty" xml:"role,omitempty"`
	RoleName     *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
	SpeechRate   *int32  `json:"speechRate,omitempty" xml:"speechRate,omitempty"`
	Words        *string `json:"words,omitempty" xml:"words,omitempty"`
}

func (s GetTaskResultResponseBodyDataAsrResult) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResultResponseBodyDataAsrResult) GoString() string {
	return s.String()
}

func (s *GetTaskResultResponseBodyDataAsrResult) GetBegin() *int64 {
	return s.Begin
}

func (s *GetTaskResultResponseBodyDataAsrResult) GetEmotionValue() *int32 {
	return s.EmotionValue
}

func (s *GetTaskResultResponseBodyDataAsrResult) GetEnd() *int64 {
	return s.End
}

func (s *GetTaskResultResponseBodyDataAsrResult) GetRole() *string {
	return s.Role
}

func (s *GetTaskResultResponseBodyDataAsrResult) GetRoleName() *string {
	return s.RoleName
}

func (s *GetTaskResultResponseBodyDataAsrResult) GetSpeechRate() *int32 {
	return s.SpeechRate
}

func (s *GetTaskResultResponseBodyDataAsrResult) GetWords() *string {
	return s.Words
}

func (s *GetTaskResultResponseBodyDataAsrResult) SetBegin(v int64) *GetTaskResultResponseBodyDataAsrResult {
	s.Begin = &v
	return s
}

func (s *GetTaskResultResponseBodyDataAsrResult) SetEmotionValue(v int32) *GetTaskResultResponseBodyDataAsrResult {
	s.EmotionValue = &v
	return s
}

func (s *GetTaskResultResponseBodyDataAsrResult) SetEnd(v int64) *GetTaskResultResponseBodyDataAsrResult {
	s.End = &v
	return s
}

func (s *GetTaskResultResponseBodyDataAsrResult) SetRole(v string) *GetTaskResultResponseBodyDataAsrResult {
	s.Role = &v
	return s
}

func (s *GetTaskResultResponseBodyDataAsrResult) SetRoleName(v string) *GetTaskResultResponseBodyDataAsrResult {
	s.RoleName = &v
	return s
}

func (s *GetTaskResultResponseBodyDataAsrResult) SetSpeechRate(v int32) *GetTaskResultResponseBodyDataAsrResult {
	s.SpeechRate = &v
	return s
}

func (s *GetTaskResultResponseBodyDataAsrResult) SetWords(v string) *GetTaskResultResponseBodyDataAsrResult {
	s.Words = &v
	return s
}

func (s *GetTaskResultResponseBodyDataAsrResult) Validate() error {
	return dara.Validate(s)
}

type GetTaskResultResponseBodyDataUsage struct {
	Rag *GetTaskResultResponseBodyDataUsageRag `json:"rag,omitempty" xml:"rag,omitempty" type:"Struct"`
}

func (s GetTaskResultResponseBodyDataUsage) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResultResponseBodyDataUsage) GoString() string {
	return s.String()
}

func (s *GetTaskResultResponseBodyDataUsage) GetRag() *GetTaskResultResponseBodyDataUsageRag {
	return s.Rag
}

func (s *GetTaskResultResponseBodyDataUsage) SetRag(v *GetTaskResultResponseBodyDataUsageRag) *GetTaskResultResponseBodyDataUsage {
	s.Rag = v
	return s
}

func (s *GetTaskResultResponseBodyDataUsage) Validate() error {
	if s.Rag != nil {
		if err := s.Rag.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTaskResultResponseBodyDataUsageRag struct {
	Adaptive      *GetTaskResultResponseBodyDataUsageRagAdaptive      `json:"adaptive,omitempty" xml:"adaptive,omitempty" type:"Struct"`
	DialogSummary *GetTaskResultResponseBodyDataUsageRagDialogSummary `json:"dialogSummary,omitempty" xml:"dialogSummary,omitempty" type:"Struct"`
}

func (s GetTaskResultResponseBodyDataUsageRag) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResultResponseBodyDataUsageRag) GoString() string {
	return s.String()
}

func (s *GetTaskResultResponseBodyDataUsageRag) GetAdaptive() *GetTaskResultResponseBodyDataUsageRagAdaptive {
	return s.Adaptive
}

func (s *GetTaskResultResponseBodyDataUsageRag) GetDialogSummary() *GetTaskResultResponseBodyDataUsageRagDialogSummary {
	return s.DialogSummary
}

func (s *GetTaskResultResponseBodyDataUsageRag) SetAdaptive(v *GetTaskResultResponseBodyDataUsageRagAdaptive) *GetTaskResultResponseBodyDataUsageRag {
	s.Adaptive = v
	return s
}

func (s *GetTaskResultResponseBodyDataUsageRag) SetDialogSummary(v *GetTaskResultResponseBodyDataUsageRagDialogSummary) *GetTaskResultResponseBodyDataUsageRag {
	s.DialogSummary = v
	return s
}

func (s *GetTaskResultResponseBodyDataUsageRag) Validate() error {
	if s.Adaptive != nil {
		if err := s.Adaptive.Validate(); err != nil {
			return err
		}
	}
	if s.DialogSummary != nil {
		if err := s.DialogSummary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTaskResultResponseBodyDataUsageRagAdaptive struct {
	InputTokens  *int32 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	InvokeCount  *int32 `json:"invokeCount,omitempty" xml:"invokeCount,omitempty"`
	OutputTokens *int32 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
}

func (s GetTaskResultResponseBodyDataUsageRagAdaptive) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResultResponseBodyDataUsageRagAdaptive) GoString() string {
	return s.String()
}

func (s *GetTaskResultResponseBodyDataUsageRagAdaptive) GetInputTokens() *int32 {
	return s.InputTokens
}

func (s *GetTaskResultResponseBodyDataUsageRagAdaptive) GetInvokeCount() *int32 {
	return s.InvokeCount
}

func (s *GetTaskResultResponseBodyDataUsageRagAdaptive) GetOutputTokens() *int32 {
	return s.OutputTokens
}

func (s *GetTaskResultResponseBodyDataUsageRagAdaptive) SetInputTokens(v int32) *GetTaskResultResponseBodyDataUsageRagAdaptive {
	s.InputTokens = &v
	return s
}

func (s *GetTaskResultResponseBodyDataUsageRagAdaptive) SetInvokeCount(v int32) *GetTaskResultResponseBodyDataUsageRagAdaptive {
	s.InvokeCount = &v
	return s
}

func (s *GetTaskResultResponseBodyDataUsageRagAdaptive) SetOutputTokens(v int32) *GetTaskResultResponseBodyDataUsageRagAdaptive {
	s.OutputTokens = &v
	return s
}

func (s *GetTaskResultResponseBodyDataUsageRagAdaptive) Validate() error {
	return dara.Validate(s)
}

type GetTaskResultResponseBodyDataUsageRagDialogSummary struct {
	InputTokens  *int32 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	InvokeCount  *int32 `json:"invokeCount,omitempty" xml:"invokeCount,omitempty"`
	OutputTokens *int32 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
}

func (s GetTaskResultResponseBodyDataUsageRagDialogSummary) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResultResponseBodyDataUsageRagDialogSummary) GoString() string {
	return s.String()
}

func (s *GetTaskResultResponseBodyDataUsageRagDialogSummary) GetInputTokens() *int32 {
	return s.InputTokens
}

func (s *GetTaskResultResponseBodyDataUsageRagDialogSummary) GetInvokeCount() *int32 {
	return s.InvokeCount
}

func (s *GetTaskResultResponseBodyDataUsageRagDialogSummary) GetOutputTokens() *int32 {
	return s.OutputTokens
}

func (s *GetTaskResultResponseBodyDataUsageRagDialogSummary) SetInputTokens(v int32) *GetTaskResultResponseBodyDataUsageRagDialogSummary {
	s.InputTokens = &v
	return s
}

func (s *GetTaskResultResponseBodyDataUsageRagDialogSummary) SetInvokeCount(v int32) *GetTaskResultResponseBodyDataUsageRagDialogSummary {
	s.InvokeCount = &v
	return s
}

func (s *GetTaskResultResponseBodyDataUsageRagDialogSummary) SetOutputTokens(v int32) *GetTaskResultResponseBodyDataUsageRagDialogSummary {
	s.OutputTokens = &v
	return s
}

func (s *GetTaskResultResponseBodyDataUsageRagDialogSummary) Validate() error {
	return dara.Validate(s)
}
