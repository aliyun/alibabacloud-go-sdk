// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunVideoAnalysisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeader(v *RunVideoAnalysisResponseBodyHeader) *RunVideoAnalysisResponseBody
	GetHeader() *RunVideoAnalysisResponseBodyHeader
	SetPayload(v *RunVideoAnalysisResponseBodyPayload) *RunVideoAnalysisResponseBody
	GetPayload() *RunVideoAnalysisResponseBodyPayload
	SetRequestId(v string) *RunVideoAnalysisResponseBody
	GetRequestId() *string
}

type RunVideoAnalysisResponseBody struct {
	Header  *RunVideoAnalysisResponseBodyHeader  `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload *RunVideoAnalysisResponseBodyPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
	// example:
	//
	// 117F5ABE-CF02-5502-9A3F-E56BC9081A64
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RunVideoAnalysisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunVideoAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBody) GetHeader() *RunVideoAnalysisResponseBodyHeader {
	return s.Header
}

func (s *RunVideoAnalysisResponseBody) GetPayload() *RunVideoAnalysisResponseBodyPayload {
	return s.Payload
}

func (s *RunVideoAnalysisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunVideoAnalysisResponseBody) SetHeader(v *RunVideoAnalysisResponseBodyHeader) *RunVideoAnalysisResponseBody {
	s.Header = v
	return s
}

func (s *RunVideoAnalysisResponseBody) SetPayload(v *RunVideoAnalysisResponseBodyPayload) *RunVideoAnalysisResponseBody {
	s.Payload = v
	return s
}

func (s *RunVideoAnalysisResponseBody) SetRequestId(v string) *RunVideoAnalysisResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunVideoAnalysisResponseBody) Validate() error {
	return dara.Validate(s)
}

type RunVideoAnalysisResponseBodyHeader struct {
	// example:
	//
	// InvalidParam
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// Pop sign mismatch, please check log.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// task-progress-start-generating
	Event     *string `json:"event,omitempty" xml:"event,omitempty"`
	EventInfo *string `json:"eventInfo,omitempty" xml:"eventInfo,omitempty"`
	// example:
	//
	// xxx
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// xxx
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 2150432017236011824686132ecdbc
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s RunVideoAnalysisResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunVideoAnalysisResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunVideoAnalysisResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunVideoAnalysisResponseBodyHeader) GetEventInfo() *string {
	return s.EventInfo
}

func (s *RunVideoAnalysisResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunVideoAnalysisResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunVideoAnalysisResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunVideoAnalysisResponseBodyHeader) SetErrorCode(v string) *RunVideoAnalysisResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyHeader) SetErrorMessage(v string) *RunVideoAnalysisResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyHeader) SetEvent(v string) *RunVideoAnalysisResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyHeader) SetEventInfo(v string) *RunVideoAnalysisResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyHeader) SetSessionId(v string) *RunVideoAnalysisResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyHeader) SetTaskId(v string) *RunVideoAnalysisResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyHeader) SetTraceId(v string) *RunVideoAnalysisResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunVideoAnalysisResponseBodyPayload struct {
	Output *RunVideoAnalysisResponseBodyPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Usage  *RunVideoAnalysisResponseBodyPayloadUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunVideoAnalysisResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayload) GetOutput() *RunVideoAnalysisResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunVideoAnalysisResponseBodyPayload) GetUsage() *RunVideoAnalysisResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunVideoAnalysisResponseBodyPayload) SetOutput(v *RunVideoAnalysisResponseBodyPayloadOutput) *RunVideoAnalysisResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayload) SetUsage(v *RunVideoAnalysisResponseBodyPayloadUsage) *RunVideoAnalysisResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayload) Validate() error {
	return dara.Validate(s)
}

type RunVideoAnalysisResponseBodyPayloadOutput struct {
	// example:
	//
	// http://
	ResultJsonFileUrl              *string                                                                  `json:"resultJsonFileUrl,omitempty" xml:"resultJsonFileUrl,omitempty"`
	VideoAnalysisResult            *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult            `json:"videoAnalysisResult,omitempty" xml:"videoAnalysisResult,omitempty" type:"Struct"`
	VideoCaptionResult             *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResult             `json:"videoCaptionResult,omitempty" xml:"videoCaptionResult,omitempty" type:"Struct"`
	VideoGenerateResult            *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult            `json:"videoGenerateResult,omitempty" xml:"videoGenerateResult,omitempty" type:"Struct"`
	VideoGenerateResults           []*RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults         `json:"videoGenerateResults,omitempty" xml:"videoGenerateResults,omitempty" type:"Repeated"`
	VideoMindMappingGenerateResult *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult `json:"videoMindMappingGenerateResult,omitempty" xml:"videoMindMappingGenerateResult,omitempty" type:"Struct"`
	VideoShotSnapshotResult        *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResult        `json:"videoShotSnapshotResult,omitempty" xml:"videoShotSnapshotResult,omitempty" type:"Struct"`
	VideoTitleGenerateResult       *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult       `json:"videoTitleGenerateResult,omitempty" xml:"videoTitleGenerateResult,omitempty" type:"Struct"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutput) GetResultJsonFileUrl() *string {
	return s.ResultJsonFileUrl
}

func (s *RunVideoAnalysisResponseBodyPayloadOutput) GetVideoAnalysisResult() *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult {
	return s.VideoAnalysisResult
}

func (s *RunVideoAnalysisResponseBodyPayloadOutput) GetVideoCaptionResult() *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResult {
	return s.VideoCaptionResult
}

func (s *RunVideoAnalysisResponseBodyPayloadOutput) GetVideoGenerateResult() *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult {
	return s.VideoGenerateResult
}

func (s *RunVideoAnalysisResponseBodyPayloadOutput) GetVideoGenerateResults() []*RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults {
	return s.VideoGenerateResults
}

func (s *RunVideoAnalysisResponseBodyPayloadOutput) GetVideoMindMappingGenerateResult() *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult {
	return s.VideoMindMappingGenerateResult
}

func (s *RunVideoAnalysisResponseBodyPayloadOutput) GetVideoShotSnapshotResult() *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResult {
	return s.VideoShotSnapshotResult
}

func (s *RunVideoAnalysisResponseBodyPayloadOutput) GetVideoTitleGenerateResult() *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult {
	return s.VideoTitleGenerateResult
}

func (s *RunVideoAnalysisResponseBodyPayloadOutput) SetResultJsonFileUrl(v string) *RunVideoAnalysisResponseBodyPayloadOutput {
	s.ResultJsonFileUrl = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutput) SetVideoAnalysisResult(v *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult) *RunVideoAnalysisResponseBodyPayloadOutput {
	s.VideoAnalysisResult = v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutput) SetVideoCaptionResult(v *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResult) *RunVideoAnalysisResponseBodyPayloadOutput {
	s.VideoCaptionResult = v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutput) SetVideoGenerateResult(v *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult) *RunVideoAnalysisResponseBodyPayloadOutput {
	s.VideoGenerateResult = v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutput) SetVideoGenerateResults(v []*RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults) *RunVideoAnalysisResponseBodyPayloadOutput {
	s.VideoGenerateResults = v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutput) SetVideoMindMappingGenerateResult(v *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult) *RunVideoAnalysisResponseBodyPayloadOutput {
	s.VideoMindMappingGenerateResult = v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutput) SetVideoShotSnapshotResult(v *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResult) *RunVideoAnalysisResponseBodyPayloadOutput {
	s.VideoShotSnapshotResult = v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutput) SetVideoTitleGenerateResult(v *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult) *RunVideoAnalysisResponseBodyPayloadOutput {
	s.VideoTitleGenerateResult = v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult struct {
	// example:
	//
	// true
	GenerateFinished *bool `json:"generateFinished,omitempty" xml:"generateFinished,omitempty"`
	// example:
	//
	// qwen-vl-max
	ModelId                  *string                                                                                 `json:"modelId,omitempty" xml:"modelId,omitempty"`
	Text                     *string                                                                                 `json:"text,omitempty" xml:"text,omitempty"`
	Usage                    *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultUsage                      `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
	VideoShotAnalysisResults []*RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultVideoShotAnalysisResults `json:"videoShotAnalysisResults,omitempty" xml:"videoShotAnalysisResults,omitempty" type:"Repeated"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult) String() string {
	return dara.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult) GetGenerateFinished() *bool {
	return s.GenerateFinished
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult) GetModelId() *string {
	return s.ModelId
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult) GetText() *string {
	return s.Text
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult) GetUsage() *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultUsage {
	return s.Usage
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult) GetVideoShotAnalysisResults() []*RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultVideoShotAnalysisResults {
	return s.VideoShotAnalysisResults
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult) SetGenerateFinished(v bool) *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult {
	s.GenerateFinished = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult) SetModelId(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult {
	s.ModelId = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult) SetText(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult {
	s.Text = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult) SetUsage(v *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultUsage) *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult {
	s.Usage = v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult) SetVideoShotAnalysisResults(v []*RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultVideoShotAnalysisResults) *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult {
	s.VideoShotAnalysisResults = v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResult) Validate() error {
	return dara.Validate(s)
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultUsage struct {
	// example:
	//
	// 1
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 1
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 1
	TotalTokens *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultUsage) String() string {
	return dara.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultUsage) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultUsage) SetInputTokens(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultUsage {
	s.InputTokens = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultUsage) SetOutputTokens(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultUsage) SetTotalTokens(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultUsage) Validate() error {
	return dara.Validate(s)
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultVideoShotAnalysisResults struct {
	// example:
	//
	// 10000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 1000
	StartTime *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Text      *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultVideoShotAnalysisResults) String() string {
	return dara.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultVideoShotAnalysisResults) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultVideoShotAnalysisResults) GetEndTime() *int64 {
	return s.EndTime
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultVideoShotAnalysisResults) GetStartTime() *int64 {
	return s.StartTime
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultVideoShotAnalysisResults) GetText() *string {
	return s.Text
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultVideoShotAnalysisResults) SetEndTime(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultVideoShotAnalysisResults {
	s.EndTime = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultVideoShotAnalysisResults) SetStartTime(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultVideoShotAnalysisResults {
	s.StartTime = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultVideoShotAnalysisResults) SetText(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultVideoShotAnalysisResults {
	s.Text = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoAnalysisResultVideoShotAnalysisResults) Validate() error {
	return dara.Validate(s)
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResult struct {
	// example:
	//
	// true
	GenerateFinished *bool                                                                       `json:"generateFinished,omitempty" xml:"generateFinished,omitempty"`
	VideoCaptions    []*RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions `json:"videoCaptions,omitempty" xml:"videoCaptions,omitempty" type:"Repeated"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResult) String() string {
	return dara.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResult) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResult) GetGenerateFinished() *bool {
	return s.GenerateFinished
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResult) GetVideoCaptions() []*RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions {
	return s.VideoCaptions
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResult) SetGenerateFinished(v bool) *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResult {
	s.GenerateFinished = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResult) SetVideoCaptions(v []*RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions) *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResult {
	s.VideoCaptions = v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResult) Validate() error {
	return dara.Validate(s)
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions struct {
	// example:
	//
	// 1710432000000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 00:01
	EndTimeFormat *string `json:"endTimeFormat,omitempty" xml:"endTimeFormat,omitempty"`
	// example:
	//
	// 张三
	Speaker *string `json:"speaker,omitempty" xml:"speaker,omitempty"`
	// example:
	//
	// 0
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 00:01
	StartTimeFormat *string `json:"startTimeFormat,omitempty" xml:"startTimeFormat,omitempty"`
	// example:
	//
	// xxx
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions) String() string {
	return dara.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions) GetEndTime() *int64 {
	return s.EndTime
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions) GetEndTimeFormat() *string {
	return s.EndTimeFormat
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions) GetSpeaker() *string {
	return s.Speaker
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions) GetStartTime() *int64 {
	return s.StartTime
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions) GetStartTimeFormat() *string {
	return s.StartTimeFormat
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions) GetText() *string {
	return s.Text
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions) SetEndTime(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions {
	s.EndTime = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions) SetEndTimeFormat(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions {
	s.EndTimeFormat = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions) SetSpeaker(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions {
	s.Speaker = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions) SetStartTime(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions {
	s.StartTime = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions) SetStartTimeFormat(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions {
	s.StartTimeFormat = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions) SetText(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions {
	s.Text = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoCaptionResultVideoCaptions) Validate() error {
	return dara.Validate(s)
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult struct {
	// example:
	//
	// true
	GenerateFinished *bool  `json:"generateFinished,omitempty" xml:"generateFinished,omitempty"`
	Index            *int32 `json:"index,omitempty" xml:"index,omitempty"`
	// example:
	//
	// qwen-max
	ModelId     *string                                                            `json:"modelId,omitempty" xml:"modelId,omitempty"`
	ModelReduce *bool                                                              `json:"modelReduce,omitempty" xml:"modelReduce,omitempty"`
	ReasonText  *string                                                            `json:"reasonText,omitempty" xml:"reasonText,omitempty"`
	Text        *string                                                            `json:"text,omitempty" xml:"text,omitempty"`
	Usage       *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult) String() string {
	return dara.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult) GetGenerateFinished() *bool {
	return s.GenerateFinished
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult) GetIndex() *int32 {
	return s.Index
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult) GetModelId() *string {
	return s.ModelId
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult) GetModelReduce() *bool {
	return s.ModelReduce
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult) GetReasonText() *string {
	return s.ReasonText
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult) GetText() *string {
	return s.Text
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult) GetUsage() *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultUsage {
	return s.Usage
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult) SetGenerateFinished(v bool) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult {
	s.GenerateFinished = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult) SetIndex(v int32) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult {
	s.Index = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult) SetModelId(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult {
	s.ModelId = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult) SetModelReduce(v bool) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult {
	s.ModelReduce = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult) SetReasonText(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult {
	s.ReasonText = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult) SetText(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult {
	s.Text = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult) SetUsage(v *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultUsage) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult {
	s.Usage = v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResult) Validate() error {
	return dara.Validate(s)
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultUsage struct {
	// example:
	//
	// 1
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 1
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 1
	TotalTokens *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultUsage) String() string {
	return dara.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultUsage) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultUsage) SetInputTokens(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultUsage {
	s.InputTokens = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultUsage) SetOutputTokens(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultUsage) SetTotalTokens(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultUsage) Validate() error {
	return dara.Validate(s)
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults struct {
	GenerateFinished *bool                                                               `json:"generateFinished,omitempty" xml:"generateFinished,omitempty"`
	Index            *int32                                                              `json:"index,omitempty" xml:"index,omitempty"`
	ModelId          *string                                                             `json:"modelId,omitempty" xml:"modelId,omitempty"`
	ReasonText       *string                                                             `json:"reasonText,omitempty" xml:"reasonText,omitempty"`
	Text             *string                                                             `json:"text,omitempty" xml:"text,omitempty"`
	Usage            *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultsUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults) String() string {
	return dara.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults) GetGenerateFinished() *bool {
	return s.GenerateFinished
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults) GetIndex() *int32 {
	return s.Index
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults) GetModelId() *string {
	return s.ModelId
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults) GetReasonText() *string {
	return s.ReasonText
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults) GetText() *string {
	return s.Text
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults) GetUsage() *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultsUsage {
	return s.Usage
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults) SetGenerateFinished(v bool) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults {
	s.GenerateFinished = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults) SetIndex(v int32) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults {
	s.Index = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults) SetModelId(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults {
	s.ModelId = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults) SetReasonText(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults {
	s.ReasonText = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults) SetText(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults {
	s.Text = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults) SetUsage(v *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultsUsage) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults {
	s.Usage = v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResults) Validate() error {
	return dara.Validate(s)
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultsUsage struct {
	InputTokens  *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	TotalTokens  *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultsUsage) String() string {
	return dara.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultsUsage) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultsUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultsUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultsUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultsUsage) SetInputTokens(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultsUsage {
	s.InputTokens = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultsUsage) SetOutputTokens(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultsUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultsUsage) SetTotalTokens(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultsUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoGenerateResultsUsage) Validate() error {
	return dara.Validate(s)
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult struct {
	// example:
	//
	// true
	GenerateFinished *bool `json:"generateFinished,omitempty" xml:"generateFinished,omitempty"`
	// example:
	//
	// true
	ModelId           *string                                                                                     `json:"modelId,omitempty" xml:"modelId,omitempty"`
	ModelReduce       *bool                                                                                       `json:"modelReduce,omitempty" xml:"modelReduce,omitempty"`
	Text              *string                                                                                     `json:"text,omitempty" xml:"text,omitempty"`
	Usage             *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultUsage               `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
	VideoMindMappings []*RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappings `json:"videoMindMappings,omitempty" xml:"videoMindMappings,omitempty" type:"Repeated"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult) String() string {
	return dara.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult) GetGenerateFinished() *bool {
	return s.GenerateFinished
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult) GetModelId() *string {
	return s.ModelId
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult) GetModelReduce() *bool {
	return s.ModelReduce
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult) GetText() *string {
	return s.Text
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult) GetUsage() *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultUsage {
	return s.Usage
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult) GetVideoMindMappings() []*RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappings {
	return s.VideoMindMappings
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult) SetGenerateFinished(v bool) *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult {
	s.GenerateFinished = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult) SetModelId(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult {
	s.ModelId = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult) SetModelReduce(v bool) *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult {
	s.ModelReduce = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult) SetText(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult {
	s.Text = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult) SetUsage(v *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultUsage) *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult {
	s.Usage = v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult) SetVideoMindMappings(v []*RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappings) *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult {
	s.VideoMindMappings = v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResult) Validate() error {
	return dara.Validate(s)
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultUsage struct {
	// example:
	//
	// 1
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 1
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 1
	TotalTokens *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultUsage) String() string {
	return dara.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultUsage) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultUsage) SetInputTokens(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultUsage {
	s.InputTokens = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultUsage) SetOutputTokens(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultUsage) SetTotalTokens(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultUsage) Validate() error {
	return dara.Validate(s)
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappings struct {
	ChildNodes []*RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes `json:"childNodes,omitempty" xml:"childNodes,omitempty" type:"Repeated"`
	Name       *string                                                                                               `json:"name,omitempty" xml:"name,omitempty"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappings) String() string {
	return dara.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappings) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappings) GetChildNodes() []*RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes {
	return s.ChildNodes
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappings) GetName() *string {
	return s.Name
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappings) SetChildNodes(v []*RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes) *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappings {
	s.ChildNodes = v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappings) SetName(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappings {
	s.Name = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappings) Validate() error {
	return dara.Validate(s)
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes struct {
	ChildNodes []*RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodesChildNodes `json:"childNodes,omitempty" xml:"childNodes,omitempty" type:"Repeated"`
	Name       *string                                                                                                         `json:"name,omitempty" xml:"name,omitempty"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes) String() string {
	return dara.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes) GetChildNodes() []*RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodesChildNodes {
	return s.ChildNodes
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes) GetName() *string {
	return s.Name
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes) SetChildNodes(v []*RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodesChildNodes) *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes {
	s.ChildNodes = v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes) SetName(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes {
	s.Name = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes) Validate() error {
	return dara.Validate(s)
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodesChildNodes struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodesChildNodes) String() string {
	return dara.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodesChildNodes) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodesChildNodes) GetName() *string {
	return s.Name
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodesChildNodes) SetName(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodesChildNodes {
	s.Name = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodesChildNodes) Validate() error {
	return dara.Validate(s)
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResult struct {
	VideoShots []*RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots `json:"videoShots,omitempty" xml:"videoShots,omitempty" type:"Repeated"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResult) String() string {
	return dara.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResult) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResult) GetVideoShots() []*RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots {
	return s.VideoShots
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResult) SetVideoShots(v []*RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots) *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResult {
	s.VideoShots = v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResult) Validate() error {
	return dara.Validate(s)
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots struct {
	EndTime         *int64                                                                                      `json:"endTime,omitempty" xml:"endTime,omitempty"`
	EndTimeFormat   *string                                                                                     `json:"endTimeFormat,omitempty" xml:"endTimeFormat,omitempty"`
	StartTime       *int64                                                                                      `json:"startTime,omitempty" xml:"startTime,omitempty"`
	StartTimeFormat *string                                                                                     `json:"startTimeFormat,omitempty" xml:"startTimeFormat,omitempty"`
	VideoSnapshots  []*RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShotsVideoSnapshots `json:"videoSnapshots,omitempty" xml:"videoSnapshots,omitempty" type:"Repeated"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots) String() string {
	return dara.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots) GetEndTime() *int64 {
	return s.EndTime
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots) GetEndTimeFormat() *string {
	return s.EndTimeFormat
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots) GetStartTime() *int64 {
	return s.StartTime
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots) GetStartTimeFormat() *string {
	return s.StartTimeFormat
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots) GetVideoSnapshots() []*RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShotsVideoSnapshots {
	return s.VideoSnapshots
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots) SetEndTime(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots {
	s.EndTime = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots) SetEndTimeFormat(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots {
	s.EndTimeFormat = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots) SetStartTime(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots {
	s.StartTime = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots) SetStartTimeFormat(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots {
	s.StartTimeFormat = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots) SetVideoSnapshots(v []*RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShotsVideoSnapshots) *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots {
	s.VideoSnapshots = v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShots) Validate() error {
	return dara.Validate(s)
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShotsVideoSnapshots struct {
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShotsVideoSnapshots) String() string {
	return dara.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShotsVideoSnapshots) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShotsVideoSnapshots) GetUrl() *string {
	return s.Url
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShotsVideoSnapshots) SetUrl(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShotsVideoSnapshots {
	s.Url = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoShotSnapshotResultVideoShotsVideoSnapshots) Validate() error {
	return dara.Validate(s)
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult struct {
	// example:
	//
	// true
	GenerateFinished *bool `json:"generateFinished,omitempty" xml:"generateFinished,omitempty"`
	// example:
	//
	// qwen-max
	ModelId     *string                                                                 `json:"modelId,omitempty" xml:"modelId,omitempty"`
	ModelReduce *bool                                                                   `json:"modelReduce,omitempty" xml:"modelReduce,omitempty"`
	Text        *string                                                                 `json:"text,omitempty" xml:"text,omitempty"`
	Usage       *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResultUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult) String() string {
	return dara.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult) GetGenerateFinished() *bool {
	return s.GenerateFinished
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult) GetModelId() *string {
	return s.ModelId
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult) GetModelReduce() *bool {
	return s.ModelReduce
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult) GetText() *string {
	return s.Text
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult) GetUsage() *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResultUsage {
	return s.Usage
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult) SetGenerateFinished(v bool) *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult {
	s.GenerateFinished = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult) SetModelId(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult {
	s.ModelId = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult) SetModelReduce(v bool) *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult {
	s.ModelReduce = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult) SetText(v string) *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult {
	s.Text = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult) SetUsage(v *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResultUsage) *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult {
	s.Usage = v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResult) Validate() error {
	return dara.Validate(s)
}

type RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResultUsage struct {
	// example:
	//
	// 1
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 1
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 2
	TotalTokens *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResultUsage) String() string {
	return dara.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResultUsage) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResultUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResultUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResultUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResultUsage) SetInputTokens(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResultUsage {
	s.InputTokens = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResultUsage) SetOutputTokens(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResultUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResultUsage) SetTotalTokens(v int64) *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResultUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadOutputVideoTitleGenerateResultUsage) Validate() error {
	return dara.Validate(s)
}

type RunVideoAnalysisResponseBodyPayloadUsage struct {
	InputTokens  *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	TotalTokens  *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s RunVideoAnalysisResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunVideoAnalysisResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunVideoAnalysisResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunVideoAnalysisResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunVideoAnalysisResponseBodyPayloadUsage) SetInputTokens(v int64) *RunVideoAnalysisResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunVideoAnalysisResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunVideoAnalysisResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunVideoAnalysisResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
