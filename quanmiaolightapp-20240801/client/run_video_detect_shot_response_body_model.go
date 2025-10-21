// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunVideoDetectShotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeader(v *RunVideoDetectShotResponseBodyHeader) *RunVideoDetectShotResponseBody
	GetHeader() *RunVideoDetectShotResponseBodyHeader
	SetPayload(v *RunVideoDetectShotResponseBodyPayload) *RunVideoDetectShotResponseBody
	GetPayload() *RunVideoDetectShotResponseBodyPayload
	SetRequestId(v string) *RunVideoDetectShotResponseBody
	GetRequestId() *string
}

type RunVideoDetectShotResponseBody struct {
	Header  *RunVideoDetectShotResponseBodyHeader  `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload *RunVideoDetectShotResponseBodyPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 58868FD6-53D7-5ACD-80F7-854C8EA256EF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RunVideoDetectShotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunVideoDetectShotResponseBody) GoString() string {
	return s.String()
}

func (s *RunVideoDetectShotResponseBody) GetHeader() *RunVideoDetectShotResponseBodyHeader {
	return s.Header
}

func (s *RunVideoDetectShotResponseBody) GetPayload() *RunVideoDetectShotResponseBodyPayload {
	return s.Payload
}

func (s *RunVideoDetectShotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunVideoDetectShotResponseBody) SetHeader(v *RunVideoDetectShotResponseBodyHeader) *RunVideoDetectShotResponseBody {
	s.Header = v
	return s
}

func (s *RunVideoDetectShotResponseBody) SetPayload(v *RunVideoDetectShotResponseBodyPayload) *RunVideoDetectShotResponseBody {
	s.Payload = v
	return s
}

func (s *RunVideoDetectShotResponseBody) SetRequestId(v string) *RunVideoDetectShotResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunVideoDetectShotResponseBody) Validate() error {
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

type RunVideoDetectShotResponseBodyHeader struct {
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
	// task-finished
	Event     *string `json:"event,omitempty" xml:"event,omitempty"`
	EventInfo *string `json:"eventInfo,omitempty" xml:"eventInfo,omitempty"`
	// example:
	//
	// 14d15c78c4c34d428212f4d923d4ede1
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// xxxx
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 3b5287b317477940746851672dca0c
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s RunVideoDetectShotResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunVideoDetectShotResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunVideoDetectShotResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunVideoDetectShotResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunVideoDetectShotResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunVideoDetectShotResponseBodyHeader) GetEventInfo() *string {
	return s.EventInfo
}

func (s *RunVideoDetectShotResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunVideoDetectShotResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunVideoDetectShotResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunVideoDetectShotResponseBodyHeader) SetErrorCode(v string) *RunVideoDetectShotResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunVideoDetectShotResponseBodyHeader) SetErrorMessage(v string) *RunVideoDetectShotResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunVideoDetectShotResponseBodyHeader) SetEvent(v string) *RunVideoDetectShotResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunVideoDetectShotResponseBodyHeader) SetEventInfo(v string) *RunVideoDetectShotResponseBodyHeader {
	s.EventInfo = &v
	return s
}

func (s *RunVideoDetectShotResponseBodyHeader) SetSessionId(v string) *RunVideoDetectShotResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunVideoDetectShotResponseBodyHeader) SetTaskId(v string) *RunVideoDetectShotResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunVideoDetectShotResponseBodyHeader) SetTraceId(v string) *RunVideoDetectShotResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunVideoDetectShotResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunVideoDetectShotResponseBodyPayload struct {
	Output *RunVideoDetectShotResponseBodyPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Usage  *RunVideoDetectShotResponseBodyPayloadUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s RunVideoDetectShotResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunVideoDetectShotResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunVideoDetectShotResponseBodyPayload) GetOutput() *RunVideoDetectShotResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunVideoDetectShotResponseBodyPayload) GetUsage() *RunVideoDetectShotResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunVideoDetectShotResponseBodyPayload) SetOutput(v *RunVideoDetectShotResponseBodyPayloadOutput) *RunVideoDetectShotResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunVideoDetectShotResponseBodyPayload) SetUsage(v *RunVideoDetectShotResponseBodyPayloadUsage) *RunVideoDetectShotResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunVideoDetectShotResponseBodyPayload) Validate() error {
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

type RunVideoDetectShotResponseBodyPayloadOutput struct {
	VideoSplitResult *RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResult `json:"videoSplitResult,omitempty" xml:"videoSplitResult,omitempty" type:"Struct"`
}

func (s RunVideoDetectShotResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunVideoDetectShotResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunVideoDetectShotResponseBodyPayloadOutput) GetVideoSplitResult() *RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResult {
	return s.VideoSplitResult
}

func (s *RunVideoDetectShotResponseBodyPayloadOutput) SetVideoSplitResult(v *RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResult) *RunVideoDetectShotResponseBodyPayloadOutput {
	s.VideoSplitResult = v
	return s
}

func (s *RunVideoDetectShotResponseBodyPayloadOutput) Validate() error {
	if s.VideoSplitResult != nil {
		if err := s.VideoSplitResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResult struct {
	// example:
	//
	// xxx
	ReasonText *string `json:"reasonText,omitempty" xml:"reasonText,omitempty"`
	// example:
	//
	// xxx
	Text                   *string                                                                              `json:"text,omitempty" xml:"text,omitempty"`
	VideoParts             []map[string]*string                                                                 `json:"videoParts,omitempty" xml:"videoParts,omitempty" type:"Repeated"`
	VideoRecognitionResult []*RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResultVideoRecognitionResult `json:"videoRecognitionResult,omitempty" xml:"videoRecognitionResult,omitempty" type:"Repeated"`
}

func (s RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResult) String() string {
	return dara.Prettify(s)
}

func (s RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResult) GoString() string {
	return s.String()
}

func (s *RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResult) GetReasonText() *string {
	return s.ReasonText
}

func (s *RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResult) GetText() *string {
	return s.Text
}

func (s *RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResult) GetVideoParts() []map[string]*string {
	return s.VideoParts
}

func (s *RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResult) GetVideoRecognitionResult() []*RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResultVideoRecognitionResult {
	return s.VideoRecognitionResult
}

func (s *RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResult) SetReasonText(v string) *RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResult {
	s.ReasonText = &v
	return s
}

func (s *RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResult) SetText(v string) *RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResult {
	s.Text = &v
	return s
}

func (s *RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResult) SetVideoParts(v []map[string]*string) *RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResult {
	s.VideoParts = v
	return s
}

func (s *RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResult) SetVideoRecognitionResult(v []*RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResultVideoRecognitionResult) *RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResult {
	s.VideoRecognitionResult = v
	return s
}

func (s *RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResult) Validate() error {
	if s.VideoRecognitionResult != nil {
		for _, item := range s.VideoRecognitionResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResultVideoRecognitionResult struct {
	// example:
	//
	// xxx
	Asr *string `json:"asr,omitempty" xml:"asr,omitempty"`
	// example:
	//
	// 1755742611000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// xxx
	Ocr *string `json:"ocr,omitempty" xml:"ocr,omitempty"`
	// example:
	//
	// 1756433675000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// xxx
	Vl *string `json:"vl,omitempty" xml:"vl,omitempty"`
}

func (s RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResultVideoRecognitionResult) String() string {
	return dara.Prettify(s)
}

func (s RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResultVideoRecognitionResult) GoString() string {
	return s.String()
}

func (s *RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResultVideoRecognitionResult) GetAsr() *string {
	return s.Asr
}

func (s *RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResultVideoRecognitionResult) GetEndTime() *int64 {
	return s.EndTime
}

func (s *RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResultVideoRecognitionResult) GetOcr() *string {
	return s.Ocr
}

func (s *RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResultVideoRecognitionResult) GetStartTime() *int64 {
	return s.StartTime
}

func (s *RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResultVideoRecognitionResult) GetVl() *string {
	return s.Vl
}

func (s *RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResultVideoRecognitionResult) SetAsr(v string) *RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResultVideoRecognitionResult {
	s.Asr = &v
	return s
}

func (s *RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResultVideoRecognitionResult) SetEndTime(v int64) *RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResultVideoRecognitionResult {
	s.EndTime = &v
	return s
}

func (s *RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResultVideoRecognitionResult) SetOcr(v string) *RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResultVideoRecognitionResult {
	s.Ocr = &v
	return s
}

func (s *RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResultVideoRecognitionResult) SetStartTime(v int64) *RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResultVideoRecognitionResult {
	s.StartTime = &v
	return s
}

func (s *RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResultVideoRecognitionResult) SetVl(v string) *RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResultVideoRecognitionResult {
	s.Vl = &v
	return s
}

func (s *RunVideoDetectShotResponseBodyPayloadOutputVideoSplitResultVideoRecognitionResult) Validate() error {
	return dara.Validate(s)
}

type RunVideoDetectShotResponseBodyPayloadUsage struct {
	// example:
	//
	// 4546
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 820
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 5366
	TotalTokens *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s RunVideoDetectShotResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunVideoDetectShotResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunVideoDetectShotResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunVideoDetectShotResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunVideoDetectShotResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunVideoDetectShotResponseBodyPayloadUsage) SetInputTokens(v int64) *RunVideoDetectShotResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunVideoDetectShotResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunVideoDetectShotResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunVideoDetectShotResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunVideoDetectShotResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunVideoDetectShotResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
