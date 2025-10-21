// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoDetectShotTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetVideoDetectShotTaskResponseBody
	GetCode() *string
	SetData(v *GetVideoDetectShotTaskResponseBodyData) *GetVideoDetectShotTaskResponseBody
	GetData() *GetVideoDetectShotTaskResponseBodyData
	SetHttpStatusCode(v int32) *GetVideoDetectShotTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetVideoDetectShotTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetVideoDetectShotTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetVideoDetectShotTaskResponseBody
	GetSuccess() *bool
}

type GetVideoDetectShotTaskResponseBody struct {
	// example:
	//
	// successful
	Code *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetVideoDetectShotTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 117F5ABE-CF02-5502-9A3F-E56BC9081A64
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetVideoDetectShotTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVideoDetectShotTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoDetectShotTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetVideoDetectShotTaskResponseBody) GetData() *GetVideoDetectShotTaskResponseBodyData {
	return s.Data
}

func (s *GetVideoDetectShotTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetVideoDetectShotTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetVideoDetectShotTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVideoDetectShotTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetVideoDetectShotTaskResponseBody) SetCode(v string) *GetVideoDetectShotTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetVideoDetectShotTaskResponseBody) SetData(v *GetVideoDetectShotTaskResponseBodyData) *GetVideoDetectShotTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetVideoDetectShotTaskResponseBody) SetHttpStatusCode(v int32) *GetVideoDetectShotTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetVideoDetectShotTaskResponseBody) SetMessage(v string) *GetVideoDetectShotTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetVideoDetectShotTaskResponseBody) SetRequestId(v string) *GetVideoDetectShotTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoDetectShotTaskResponseBody) SetSuccess(v bool) *GetVideoDetectShotTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetVideoDetectShotTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVideoDetectShotTaskResponseBodyData struct {
	// example:
	//
	// Failed to proxy flink ui request, message: An error occurred: Invalid UUID string: jobsn.
	ErrorMessage *string                                        `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Header       *GetVideoDetectShotTaskResponseBodyDataHeader  `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload      *GetVideoDetectShotTaskResponseBodyDataPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
	// example:
	//
	// 3feb69ed02d9b1a17d0f1a942675d300
	TaskId      *string                                            `json:"taskId,omitempty" xml:"taskId,omitempty"`
	TaskRunInfo *GetVideoDetectShotTaskResponseBodyDataTaskRunInfo `json:"taskRunInfo,omitempty" xml:"taskRunInfo,omitempty" type:"Struct"`
	// example:
	//
	// SUCCESSED
	TaskStatus *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
}

func (s GetVideoDetectShotTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetVideoDetectShotTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVideoDetectShotTaskResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetVideoDetectShotTaskResponseBodyData) GetHeader() *GetVideoDetectShotTaskResponseBodyDataHeader {
	return s.Header
}

func (s *GetVideoDetectShotTaskResponseBodyData) GetPayload() *GetVideoDetectShotTaskResponseBodyDataPayload {
	return s.Payload
}

func (s *GetVideoDetectShotTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetVideoDetectShotTaskResponseBodyData) GetTaskRunInfo() *GetVideoDetectShotTaskResponseBodyDataTaskRunInfo {
	return s.TaskRunInfo
}

func (s *GetVideoDetectShotTaskResponseBodyData) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *GetVideoDetectShotTaskResponseBodyData) SetErrorMessage(v string) *GetVideoDetectShotTaskResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetVideoDetectShotTaskResponseBodyData) SetHeader(v *GetVideoDetectShotTaskResponseBodyDataHeader) *GetVideoDetectShotTaskResponseBodyData {
	s.Header = v
	return s
}

func (s *GetVideoDetectShotTaskResponseBodyData) SetPayload(v *GetVideoDetectShotTaskResponseBodyDataPayload) *GetVideoDetectShotTaskResponseBodyData {
	s.Payload = v
	return s
}

func (s *GetVideoDetectShotTaskResponseBodyData) SetTaskId(v string) *GetVideoDetectShotTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetVideoDetectShotTaskResponseBodyData) SetTaskRunInfo(v *GetVideoDetectShotTaskResponseBodyDataTaskRunInfo) *GetVideoDetectShotTaskResponseBodyData {
	s.TaskRunInfo = v
	return s
}

func (s *GetVideoDetectShotTaskResponseBodyData) SetTaskStatus(v string) *GetVideoDetectShotTaskResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *GetVideoDetectShotTaskResponseBodyData) Validate() error {
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
	if s.TaskRunInfo != nil {
		if err := s.TaskRunInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVideoDetectShotTaskResponseBodyDataHeader struct {
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// Deduct task already success,Please do not resubmit.token \\"369e8f2c-d283-424a-96c4-c83efe08c89e\\"
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// TIMEOUT_CLOSE_ORDER
	Event *string `json:"event,omitempty" xml:"event,omitempty"`
	// example:
	//
	// xxx
	EventInfo *string `json:"eventInfo,omitempty" xml:"eventInfo,omitempty"`
	// example:
	//
	// d5c38cf6-a4bf-4a57-a697-9f449926f0c9
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// 6e223291-729b-4e84-9271-c13ada1a776b
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// 215045f817272303448235204efdef
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s GetVideoDetectShotTaskResponseBodyDataHeader) String() string {
	return dara.Prettify(s)
}

func (s GetVideoDetectShotTaskResponseBodyDataHeader) GoString() string {
	return s.String()
}

func (s *GetVideoDetectShotTaskResponseBodyDataHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetVideoDetectShotTaskResponseBodyDataHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetVideoDetectShotTaskResponseBodyDataHeader) GetEvent() *string {
	return s.Event
}

func (s *GetVideoDetectShotTaskResponseBodyDataHeader) GetEventInfo() *string {
	return s.EventInfo
}

func (s *GetVideoDetectShotTaskResponseBodyDataHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *GetVideoDetectShotTaskResponseBodyDataHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *GetVideoDetectShotTaskResponseBodyDataHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *GetVideoDetectShotTaskResponseBodyDataHeader) SetErrorCode(v string) *GetVideoDetectShotTaskResponseBodyDataHeader {
	s.ErrorCode = &v
	return s
}

func (s *GetVideoDetectShotTaskResponseBodyDataHeader) SetErrorMessage(v string) *GetVideoDetectShotTaskResponseBodyDataHeader {
	s.ErrorMessage = &v
	return s
}

func (s *GetVideoDetectShotTaskResponseBodyDataHeader) SetEvent(v string) *GetVideoDetectShotTaskResponseBodyDataHeader {
	s.Event = &v
	return s
}

func (s *GetVideoDetectShotTaskResponseBodyDataHeader) SetEventInfo(v string) *GetVideoDetectShotTaskResponseBodyDataHeader {
	s.EventInfo = &v
	return s
}

func (s *GetVideoDetectShotTaskResponseBodyDataHeader) SetSessionId(v string) *GetVideoDetectShotTaskResponseBodyDataHeader {
	s.SessionId = &v
	return s
}

func (s *GetVideoDetectShotTaskResponseBodyDataHeader) SetTaskId(v string) *GetVideoDetectShotTaskResponseBodyDataHeader {
	s.TaskId = &v
	return s
}

func (s *GetVideoDetectShotTaskResponseBodyDataHeader) SetTraceId(v string) *GetVideoDetectShotTaskResponseBodyDataHeader {
	s.TraceId = &v
	return s
}

func (s *GetVideoDetectShotTaskResponseBodyDataHeader) Validate() error {
	return dara.Validate(s)
}

type GetVideoDetectShotTaskResponseBodyDataPayload struct {
	Output *GetVideoDetectShotTaskResponseBodyDataPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Usage  *GetVideoDetectShotTaskResponseBodyDataPayloadUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetVideoDetectShotTaskResponseBodyDataPayload) String() string {
	return dara.Prettify(s)
}

func (s GetVideoDetectShotTaskResponseBodyDataPayload) GoString() string {
	return s.String()
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayload) GetOutput() *GetVideoDetectShotTaskResponseBodyDataPayloadOutput {
	return s.Output
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayload) GetUsage() *GetVideoDetectShotTaskResponseBodyDataPayloadUsage {
	return s.Usage
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayload) SetOutput(v *GetVideoDetectShotTaskResponseBodyDataPayloadOutput) *GetVideoDetectShotTaskResponseBodyDataPayload {
	s.Output = v
	return s
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayload) SetUsage(v *GetVideoDetectShotTaskResponseBodyDataPayloadUsage) *GetVideoDetectShotTaskResponseBodyDataPayload {
	s.Usage = v
	return s
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayload) Validate() error {
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

type GetVideoDetectShotTaskResponseBodyDataPayloadOutput struct {
	VideoSplitResult *GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResult `json:"videoSplitResult,omitempty" xml:"videoSplitResult,omitempty" type:"Struct"`
}

func (s GetVideoDetectShotTaskResponseBodyDataPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s GetVideoDetectShotTaskResponseBodyDataPayloadOutput) GoString() string {
	return s.String()
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayloadOutput) GetVideoSplitResult() *GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResult {
	return s.VideoSplitResult
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayloadOutput) SetVideoSplitResult(v *GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResult) *GetVideoDetectShotTaskResponseBodyDataPayloadOutput {
	s.VideoSplitResult = v
	return s
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayloadOutput) Validate() error {
	if s.VideoSplitResult != nil {
		if err := s.VideoSplitResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResult struct {
	// example:
	//
	// xxx
	ReasonText *string `json:"reasonText,omitempty" xml:"reasonText,omitempty"`
	// example:
	//
	// xxx
	Text                   *string                                                                                      `json:"text,omitempty" xml:"text,omitempty"`
	VideoParts             []map[string]*string                                                                         `json:"videoParts,omitempty" xml:"videoParts,omitempty" type:"Repeated"`
	VideoRecognitionResult []*GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResultVideoRecognitionResult `json:"videoRecognitionResult,omitempty" xml:"videoRecognitionResult,omitempty" type:"Repeated"`
}

func (s GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResult) String() string {
	return dara.Prettify(s)
}

func (s GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResult) GoString() string {
	return s.String()
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResult) GetReasonText() *string {
	return s.ReasonText
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResult) GetText() *string {
	return s.Text
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResult) GetVideoParts() []map[string]*string {
	return s.VideoParts
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResult) GetVideoRecognitionResult() []*GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResultVideoRecognitionResult {
	return s.VideoRecognitionResult
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResult) SetReasonText(v string) *GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResult {
	s.ReasonText = &v
	return s
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResult) SetText(v string) *GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResult {
	s.Text = &v
	return s
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResult) SetVideoParts(v []map[string]*string) *GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResult {
	s.VideoParts = v
	return s
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResult) SetVideoRecognitionResult(v []*GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResultVideoRecognitionResult) *GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResult {
	s.VideoRecognitionResult = v
	return s
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResult) Validate() error {
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

type GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResultVideoRecognitionResult struct {
	// example:
	//
	// xxx
	Asr *string `json:"asr,omitempty" xml:"asr,omitempty"`
	// example:
	//
	// 1748483740000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// xxx
	Ocr *string `json:"ocr,omitempty" xml:"ocr,omitempty"`
	// example:
	//
	// 1758108425000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// xxx
	Vl *string `json:"vl,omitempty" xml:"vl,omitempty"`
}

func (s GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResultVideoRecognitionResult) String() string {
	return dara.Prettify(s)
}

func (s GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResultVideoRecognitionResult) GoString() string {
	return s.String()
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResultVideoRecognitionResult) GetAsr() *string {
	return s.Asr
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResultVideoRecognitionResult) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResultVideoRecognitionResult) GetOcr() *string {
	return s.Ocr
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResultVideoRecognitionResult) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResultVideoRecognitionResult) GetVl() *string {
	return s.Vl
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResultVideoRecognitionResult) SetAsr(v string) *GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResultVideoRecognitionResult {
	s.Asr = &v
	return s
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResultVideoRecognitionResult) SetEndTime(v int64) *GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResultVideoRecognitionResult {
	s.EndTime = &v
	return s
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResultVideoRecognitionResult) SetOcr(v string) *GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResultVideoRecognitionResult {
	s.Ocr = &v
	return s
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResultVideoRecognitionResult) SetStartTime(v int64) *GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResultVideoRecognitionResult {
	s.StartTime = &v
	return s
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResultVideoRecognitionResult) SetVl(v string) *GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResultVideoRecognitionResult {
	s.Vl = &v
	return s
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayloadOutputVideoSplitResultVideoRecognitionResult) Validate() error {
	return dara.Validate(s)
}

type GetVideoDetectShotTaskResponseBodyDataPayloadUsage struct {
	// example:
	//
	// 36
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 13
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 49
	TotalTokens *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s GetVideoDetectShotTaskResponseBodyDataPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s GetVideoDetectShotTaskResponseBodyDataPayloadUsage) GoString() string {
	return s.String()
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayloadUsage) SetInputTokens(v int64) *GetVideoDetectShotTaskResponseBodyDataPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayloadUsage) SetOutputTokens(v int64) *GetVideoDetectShotTaskResponseBodyDataPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayloadUsage) SetTotalTokens(v int64) *GetVideoDetectShotTaskResponseBodyDataPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *GetVideoDetectShotTaskResponseBodyDataPayloadUsage) Validate() error {
	return dara.Validate(s)
}

type GetVideoDetectShotTaskResponseBodyDataTaskRunInfo struct {
	// example:
	//
	// true
	ConcurrentChargeEnable *bool `json:"concurrentChargeEnable,omitempty" xml:"concurrentChargeEnable,omitempty"`
	// example:
	//
	// 1000
	ResponseTime *int64 `json:"responseTime,omitempty" xml:"responseTime,omitempty"`
}

func (s GetVideoDetectShotTaskResponseBodyDataTaskRunInfo) String() string {
	return dara.Prettify(s)
}

func (s GetVideoDetectShotTaskResponseBodyDataTaskRunInfo) GoString() string {
	return s.String()
}

func (s *GetVideoDetectShotTaskResponseBodyDataTaskRunInfo) GetConcurrentChargeEnable() *bool {
	return s.ConcurrentChargeEnable
}

func (s *GetVideoDetectShotTaskResponseBodyDataTaskRunInfo) GetResponseTime() *int64 {
	return s.ResponseTime
}

func (s *GetVideoDetectShotTaskResponseBodyDataTaskRunInfo) SetConcurrentChargeEnable(v bool) *GetVideoDetectShotTaskResponseBodyDataTaskRunInfo {
	s.ConcurrentChargeEnable = &v
	return s
}

func (s *GetVideoDetectShotTaskResponseBodyDataTaskRunInfo) SetResponseTime(v int64) *GetVideoDetectShotTaskResponseBodyDataTaskRunInfo {
	s.ResponseTime = &v
	return s
}

func (s *GetVideoDetectShotTaskResponseBodyDataTaskRunInfo) Validate() error {
	return dara.Validate(s)
}
