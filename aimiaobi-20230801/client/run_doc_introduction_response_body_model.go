// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunDocIntroductionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeader(v *RunDocIntroductionResponseBodyHeader) *RunDocIntroductionResponseBody
	GetHeader() *RunDocIntroductionResponseBodyHeader
	SetPayload(v *RunDocIntroductionResponseBodyPayload) *RunDocIntroductionResponseBody
	GetPayload() *RunDocIntroductionResponseBodyPayload
	SetRequestId(v string) *RunDocIntroductionResponseBody
	GetRequestId() *string
}

type RunDocIntroductionResponseBody struct {
	Header  *RunDocIntroductionResponseBodyHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *RunDocIntroductionResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// example:
	//
	// C9B5BEA6-E8C4-5861-BE37-D906D516510E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunDocIntroductionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunDocIntroductionResponseBody) GoString() string {
	return s.String()
}

func (s *RunDocIntroductionResponseBody) GetHeader() *RunDocIntroductionResponseBodyHeader {
	return s.Header
}

func (s *RunDocIntroductionResponseBody) GetPayload() *RunDocIntroductionResponseBodyPayload {
	return s.Payload
}

func (s *RunDocIntroductionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunDocIntroductionResponseBody) SetHeader(v *RunDocIntroductionResponseBodyHeader) *RunDocIntroductionResponseBody {
	s.Header = v
	return s
}

func (s *RunDocIntroductionResponseBody) SetPayload(v *RunDocIntroductionResponseBodyPayload) *RunDocIntroductionResponseBody {
	s.Payload = v
	return s
}

func (s *RunDocIntroductionResponseBody) SetRequestId(v string) *RunDocIntroductionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunDocIntroductionResponseBody) Validate() error {
	return dara.Validate(s)
}

type RunDocIntroductionResponseBodyHeader struct {
	// example:
	//
	// success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// success
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// task-started
	Event     *string `json:"Event,omitempty" xml:"Event,omitempty"`
	EventInfo *string `json:"EventInfo,omitempty" xml:"EventInfo,omitempty"`
	// example:
	//
	// 20247a52-23e2-46fb-943d-309cdee2bc6d
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 8a9cecb7-6d20-32db-8823-5882c217b647
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 0bd58ea2-dc38-45da-ac02-17f05cb9040b
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunDocIntroductionResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunDocIntroductionResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunDocIntroductionResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunDocIntroductionResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunDocIntroductionResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunDocIntroductionResponseBodyHeader) GetEventInfo() *string {
	return s.EventInfo
}

func (s *RunDocIntroductionResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunDocIntroductionResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunDocIntroductionResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunDocIntroductionResponseBodyHeader) SetErrorCode(v string) *RunDocIntroductionResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunDocIntroductionResponseBodyHeader) SetErrorMessage(v string) *RunDocIntroductionResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunDocIntroductionResponseBodyHeader) SetEvent(v string) *RunDocIntroductionResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunDocIntroductionResponseBodyHeader) SetEventInfo(v string) *RunDocIntroductionResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunDocIntroductionResponseBodyHeader) SetSessionId(v string) *RunDocIntroductionResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunDocIntroductionResponseBodyHeader) SetTaskId(v string) *RunDocIntroductionResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunDocIntroductionResponseBodyHeader) SetTraceId(v string) *RunDocIntroductionResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunDocIntroductionResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunDocIntroductionResponseBodyPayload struct {
	Output *RunDocIntroductionResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *RunDocIntroductionResponseBodyPayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunDocIntroductionResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunDocIntroductionResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunDocIntroductionResponseBodyPayload) GetOutput() *RunDocIntroductionResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunDocIntroductionResponseBodyPayload) GetUsage() *RunDocIntroductionResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunDocIntroductionResponseBodyPayload) SetOutput(v *RunDocIntroductionResponseBodyPayloadOutput) *RunDocIntroductionResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunDocIntroductionResponseBodyPayload) SetUsage(v *RunDocIntroductionResponseBodyPayloadUsage) *RunDocIntroductionResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunDocIntroductionResponseBodyPayload) Validate() error {
	return dara.Validate(s)
}

type RunDocIntroductionResponseBodyPayloadOutput struct {
	Introductions []*RunDocIntroductionResponseBodyPayloadOutputIntroductions `json:"Introductions,omitempty" xml:"Introductions,omitempty" type:"Repeated"`
	KeyPoint      *string                                                     `json:"KeyPoint,omitempty" xml:"KeyPoint,omitempty"`
	Summary       *string                                                     `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s RunDocIntroductionResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunDocIntroductionResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunDocIntroductionResponseBodyPayloadOutput) GetIntroductions() []*RunDocIntroductionResponseBodyPayloadOutputIntroductions {
	return s.Introductions
}

func (s *RunDocIntroductionResponseBodyPayloadOutput) GetKeyPoint() *string {
	return s.KeyPoint
}

func (s *RunDocIntroductionResponseBodyPayloadOutput) GetSummary() *string {
	return s.Summary
}

func (s *RunDocIntroductionResponseBodyPayloadOutput) SetIntroductions(v []*RunDocIntroductionResponseBodyPayloadOutputIntroductions) *RunDocIntroductionResponseBodyPayloadOutput {
	s.Introductions = v
	return s
}

func (s *RunDocIntroductionResponseBodyPayloadOutput) SetKeyPoint(v string) *RunDocIntroductionResponseBodyPayloadOutput {
	s.KeyPoint = &v
	return s
}

func (s *RunDocIntroductionResponseBodyPayloadOutput) SetSummary(v string) *RunDocIntroductionResponseBodyPayloadOutput {
	s.Summary = &v
	return s
}

func (s *RunDocIntroductionResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunDocIntroductionResponseBodyPayloadOutputIntroductions struct {
	Blocks []*RunDocIntroductionResponseBodyPayloadOutputIntroductionsBlocks `json:"Blocks,omitempty" xml:"Blocks,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	StartPageId *int32  `json:"StartPageId,omitempty" xml:"StartPageId,omitempty"`
	Summary     *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s RunDocIntroductionResponseBodyPayloadOutputIntroductions) String() string {
	return dara.Prettify(s)
}

func (s RunDocIntroductionResponseBodyPayloadOutputIntroductions) GoString() string {
	return s.String()
}

func (s *RunDocIntroductionResponseBodyPayloadOutputIntroductions) GetBlocks() []*RunDocIntroductionResponseBodyPayloadOutputIntroductionsBlocks {
	return s.Blocks
}

func (s *RunDocIntroductionResponseBodyPayloadOutputIntroductions) GetStartPageId() *int32 {
	return s.StartPageId
}

func (s *RunDocIntroductionResponseBodyPayloadOutputIntroductions) GetSummary() *string {
	return s.Summary
}

func (s *RunDocIntroductionResponseBodyPayloadOutputIntroductions) GetTitle() *string {
	return s.Title
}

func (s *RunDocIntroductionResponseBodyPayloadOutputIntroductions) SetBlocks(v []*RunDocIntroductionResponseBodyPayloadOutputIntroductionsBlocks) *RunDocIntroductionResponseBodyPayloadOutputIntroductions {
	s.Blocks = v
	return s
}

func (s *RunDocIntroductionResponseBodyPayloadOutputIntroductions) SetStartPageId(v int32) *RunDocIntroductionResponseBodyPayloadOutputIntroductions {
	s.StartPageId = &v
	return s
}

func (s *RunDocIntroductionResponseBodyPayloadOutputIntroductions) SetSummary(v string) *RunDocIntroductionResponseBodyPayloadOutputIntroductions {
	s.Summary = &v
	return s
}

func (s *RunDocIntroductionResponseBodyPayloadOutputIntroductions) SetTitle(v string) *RunDocIntroductionResponseBodyPayloadOutputIntroductions {
	s.Title = &v
	return s
}

func (s *RunDocIntroductionResponseBodyPayloadOutputIntroductions) Validate() error {
	return dara.Validate(s)
}

type RunDocIntroductionResponseBodyPayloadOutputIntroductionsBlocks struct {
	// example:
	//
	// 0
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// 1200
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 600
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 10
	PageId *int32 `json:"PageId,omitempty" xml:"PageId,omitempty"`
	// example:
	//
	// 600
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// example:
	//
	// 10
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 10
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RunDocIntroductionResponseBodyPayloadOutputIntroductionsBlocks) String() string {
	return dara.Prettify(s)
}

func (s RunDocIntroductionResponseBodyPayloadOutputIntroductionsBlocks) GoString() string {
	return s.String()
}

func (s *RunDocIntroductionResponseBodyPayloadOutputIntroductionsBlocks) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *RunDocIntroductionResponseBodyPayloadOutputIntroductionsBlocks) GetEndTime() *int64 {
	return s.EndTime
}

func (s *RunDocIntroductionResponseBodyPayloadOutputIntroductionsBlocks) GetHeight() *int32 {
	return s.Height
}

func (s *RunDocIntroductionResponseBodyPayloadOutputIntroductionsBlocks) GetPageId() *int32 {
	return s.PageId
}

func (s *RunDocIntroductionResponseBodyPayloadOutputIntroductionsBlocks) GetWidth() *int32 {
	return s.Width
}

func (s *RunDocIntroductionResponseBodyPayloadOutputIntroductionsBlocks) GetX() *int32 {
	return s.X
}

func (s *RunDocIntroductionResponseBodyPayloadOutputIntroductionsBlocks) GetY() *int32 {
	return s.Y
}

func (s *RunDocIntroductionResponseBodyPayloadOutputIntroductionsBlocks) SetBeginTime(v int64) *RunDocIntroductionResponseBodyPayloadOutputIntroductionsBlocks {
	s.BeginTime = &v
	return s
}

func (s *RunDocIntroductionResponseBodyPayloadOutputIntroductionsBlocks) SetEndTime(v int64) *RunDocIntroductionResponseBodyPayloadOutputIntroductionsBlocks {
	s.EndTime = &v
	return s
}

func (s *RunDocIntroductionResponseBodyPayloadOutputIntroductionsBlocks) SetHeight(v int32) *RunDocIntroductionResponseBodyPayloadOutputIntroductionsBlocks {
	s.Height = &v
	return s
}

func (s *RunDocIntroductionResponseBodyPayloadOutputIntroductionsBlocks) SetPageId(v int32) *RunDocIntroductionResponseBodyPayloadOutputIntroductionsBlocks {
	s.PageId = &v
	return s
}

func (s *RunDocIntroductionResponseBodyPayloadOutputIntroductionsBlocks) SetWidth(v int32) *RunDocIntroductionResponseBodyPayloadOutputIntroductionsBlocks {
	s.Width = &v
	return s
}

func (s *RunDocIntroductionResponseBodyPayloadOutputIntroductionsBlocks) SetX(v int32) *RunDocIntroductionResponseBodyPayloadOutputIntroductionsBlocks {
	s.X = &v
	return s
}

func (s *RunDocIntroductionResponseBodyPayloadOutputIntroductionsBlocks) SetY(v int32) *RunDocIntroductionResponseBodyPayloadOutputIntroductionsBlocks {
	s.Y = &v
	return s
}

func (s *RunDocIntroductionResponseBodyPayloadOutputIntroductionsBlocks) Validate() error {
	return dara.Validate(s)
}

type RunDocIntroductionResponseBodyPayloadUsage struct {
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

func (s RunDocIntroductionResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunDocIntroductionResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunDocIntroductionResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunDocIntroductionResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunDocIntroductionResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunDocIntroductionResponseBodyPayloadUsage) SetInputTokens(v int64) *RunDocIntroductionResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunDocIntroductionResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunDocIntroductionResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunDocIntroductionResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunDocIntroductionResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunDocIntroductionResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
