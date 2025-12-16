// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoAnalysisTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetVideoAnalysisTaskResponseBody
	GetCode() *string
	SetData(v *GetVideoAnalysisTaskResponseBodyData) *GetVideoAnalysisTaskResponseBody
	GetData() *GetVideoAnalysisTaskResponseBodyData
	SetHttpStatusCode(v int32) *GetVideoAnalysisTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetVideoAnalysisTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetVideoAnalysisTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetVideoAnalysisTaskResponseBody
	GetSuccess() *bool
}

type GetVideoAnalysisTaskResponseBody struct {
	// example:
	//
	// successful
	Code *string                               `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetVideoAnalysisTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5D0E915E-655D-59A8-894F-93873F73AAE5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetVideoAnalysisTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetVideoAnalysisTaskResponseBody) GetData() *GetVideoAnalysisTaskResponseBodyData {
	return s.Data
}

func (s *GetVideoAnalysisTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetVideoAnalysisTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetVideoAnalysisTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVideoAnalysisTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetVideoAnalysisTaskResponseBody) SetCode(v string) *GetVideoAnalysisTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBody) SetData(v *GetVideoAnalysisTaskResponseBodyData) *GetVideoAnalysisTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBody) SetHttpStatusCode(v int32) *GetVideoAnalysisTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBody) SetMessage(v string) *GetVideoAnalysisTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBody) SetRequestId(v string) *GetVideoAnalysisTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBody) SetSuccess(v bool) *GetVideoAnalysisTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVideoAnalysisTaskResponseBodyData struct {
	// example:
	//
	// Access was denied, message: No such namespace namespaces/mjp-test-default.
	ErrorMessage *string                                      `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Header       *GetVideoAnalysisTaskResponseBodyDataHeader  `json:"header,omitempty" xml:"header,omitempty" type:"Struct"`
	Payload      *GetVideoAnalysisTaskResponseBodyDataPayload `json:"payload,omitempty" xml:"payload,omitempty" type:"Struct"`
	// example:
	//
	// 3feb69ed02d9b1a17d0f1a942675d300
	TaskId      *string                                          `json:"taskId,omitempty" xml:"taskId,omitempty"`
	TaskRunInfo *GetVideoAnalysisTaskResponseBodyDataTaskRunInfo `json:"taskRunInfo,omitempty" xml:"taskRunInfo,omitempty" type:"Struct"`
	// example:
	//
	// SUCCESSED
	TaskStatus *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
}

func (s GetVideoAnalysisTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetVideoAnalysisTaskResponseBodyData) GetHeader() *GetVideoAnalysisTaskResponseBodyDataHeader {
	return s.Header
}

func (s *GetVideoAnalysisTaskResponseBodyData) GetPayload() *GetVideoAnalysisTaskResponseBodyDataPayload {
	return s.Payload
}

func (s *GetVideoAnalysisTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetVideoAnalysisTaskResponseBodyData) GetTaskRunInfo() *GetVideoAnalysisTaskResponseBodyDataTaskRunInfo {
	return s.TaskRunInfo
}

func (s *GetVideoAnalysisTaskResponseBodyData) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *GetVideoAnalysisTaskResponseBodyData) SetErrorMessage(v string) *GetVideoAnalysisTaskResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyData) SetHeader(v *GetVideoAnalysisTaskResponseBodyDataHeader) *GetVideoAnalysisTaskResponseBodyData {
	s.Header = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyData) SetPayload(v *GetVideoAnalysisTaskResponseBodyDataPayload) *GetVideoAnalysisTaskResponseBodyData {
	s.Payload = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyData) SetTaskId(v string) *GetVideoAnalysisTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyData) SetTaskRunInfo(v *GetVideoAnalysisTaskResponseBodyDataTaskRunInfo) *GetVideoAnalysisTaskResponseBodyData {
	s.TaskRunInfo = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyData) SetTaskStatus(v string) *GetVideoAnalysisTaskResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyData) Validate() error {
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

type GetVideoAnalysisTaskResponseBodyDataHeader struct {
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

func (s GetVideoAnalysisTaskResponseBodyDataHeader) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataHeader) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetVideoAnalysisTaskResponseBodyDataHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetVideoAnalysisTaskResponseBodyDataHeader) GetEvent() *string {
	return s.Event
}

func (s *GetVideoAnalysisTaskResponseBodyDataHeader) GetEventInfo() *string {
	return s.EventInfo
}

func (s *GetVideoAnalysisTaskResponseBodyDataHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *GetVideoAnalysisTaskResponseBodyDataHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *GetVideoAnalysisTaskResponseBodyDataHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *GetVideoAnalysisTaskResponseBodyDataHeader) SetErrorCode(v string) *GetVideoAnalysisTaskResponseBodyDataHeader {
	s.ErrorCode = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataHeader) SetErrorMessage(v string) *GetVideoAnalysisTaskResponseBodyDataHeader {
	s.ErrorMessage = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataHeader) SetEvent(v string) *GetVideoAnalysisTaskResponseBodyDataHeader {
	s.Event = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataHeader) SetEventInfo(v string) *GetVideoAnalysisTaskResponseBodyDataHeader {
	s.EventInfo = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataHeader) SetSessionId(v string) *GetVideoAnalysisTaskResponseBodyDataHeader {
	s.SessionId = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataHeader) SetTaskId(v string) *GetVideoAnalysisTaskResponseBodyDataHeader {
	s.TaskId = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataHeader) SetTraceId(v string) *GetVideoAnalysisTaskResponseBodyDataHeader {
	s.TraceId = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataHeader) Validate() error {
	return dara.Validate(s)
}

type GetVideoAnalysisTaskResponseBodyDataPayload struct {
	Output *GetVideoAnalysisTaskResponseBodyDataPayloadOutput `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Usage  *GetVideoAnalysisTaskResponseBodyDataPayloadUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayload) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayload) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayload) GetOutput() *GetVideoAnalysisTaskResponseBodyDataPayloadOutput {
	return s.Output
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayload) GetUsage() *GetVideoAnalysisTaskResponseBodyDataPayloadUsage {
	return s.Usage
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayload) SetOutput(v *GetVideoAnalysisTaskResponseBodyDataPayloadOutput) *GetVideoAnalysisTaskResponseBodyDataPayload {
	s.Output = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayload) SetUsage(v *GetVideoAnalysisTaskResponseBodyDataPayloadUsage) *GetVideoAnalysisTaskResponseBodyDataPayload {
	s.Usage = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayload) Validate() error {
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

type GetVideoAnalysisTaskResponseBodyDataPayloadOutput struct {
	AddDatasetDocumentsResult      *GetVideoAnalysisTaskResponseBodyDataPayloadOutputAddDatasetDocumentsResult      `json:"addDatasetDocumentsResult,omitempty" xml:"addDatasetDocumentsResult,omitempty" type:"Struct"`
	ResultJsonFileUrl              *string                                                                          `json:"resultJsonFileUrl,omitempty" xml:"resultJsonFileUrl,omitempty"`
	VideoAnalysisResult            *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResult            `json:"videoAnalysisResult,omitempty" xml:"videoAnalysisResult,omitempty" type:"Struct"`
	VideoCalculatorResult          *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResult          `json:"videoCalculatorResult,omitempty" xml:"videoCalculatorResult,omitempty" type:"Struct"`
	VideoCaptionResult             *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResult             `json:"videoCaptionResult,omitempty" xml:"videoCaptionResult,omitempty" type:"Struct"`
	VideoGenerateResult            *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult            `json:"videoGenerateResult,omitempty" xml:"videoGenerateResult,omitempty" type:"Struct"`
	VideoGenerateResults           []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults         `json:"videoGenerateResults,omitempty" xml:"videoGenerateResults,omitempty" type:"Repeated"`
	VideoMindMappingGenerateResult *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResult `json:"videoMindMappingGenerateResult,omitempty" xml:"videoMindMappingGenerateResult,omitempty" type:"Struct"`
	VideoRoleRecognitionResult     *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResult     `json:"videoRoleRecognitionResult,omitempty" xml:"videoRoleRecognitionResult,omitempty" type:"Struct"`
	VideoTitleGenerateResult       *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResult       `json:"videoTitleGenerateResult,omitempty" xml:"videoTitleGenerateResult,omitempty" type:"Struct"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutput) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutput) GetAddDatasetDocumentsResult() *GetVideoAnalysisTaskResponseBodyDataPayloadOutputAddDatasetDocumentsResult {
	return s.AddDatasetDocumentsResult
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutput) GetResultJsonFileUrl() *string {
	return s.ResultJsonFileUrl
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutput) GetVideoAnalysisResult() *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResult {
	return s.VideoAnalysisResult
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutput) GetVideoCalculatorResult() *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResult {
	return s.VideoCalculatorResult
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutput) GetVideoCaptionResult() *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResult {
	return s.VideoCaptionResult
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutput) GetVideoGenerateResult() *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult {
	return s.VideoGenerateResult
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutput) GetVideoGenerateResults() []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults {
	return s.VideoGenerateResults
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutput) GetVideoMindMappingGenerateResult() *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResult {
	return s.VideoMindMappingGenerateResult
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutput) GetVideoRoleRecognitionResult() *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResult {
	return s.VideoRoleRecognitionResult
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutput) GetVideoTitleGenerateResult() *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResult {
	return s.VideoTitleGenerateResult
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutput) SetAddDatasetDocumentsResult(v *GetVideoAnalysisTaskResponseBodyDataPayloadOutputAddDatasetDocumentsResult) *GetVideoAnalysisTaskResponseBodyDataPayloadOutput {
	s.AddDatasetDocumentsResult = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutput) SetResultJsonFileUrl(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutput {
	s.ResultJsonFileUrl = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutput) SetVideoAnalysisResult(v *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResult) *GetVideoAnalysisTaskResponseBodyDataPayloadOutput {
	s.VideoAnalysisResult = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutput) SetVideoCalculatorResult(v *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResult) *GetVideoAnalysisTaskResponseBodyDataPayloadOutput {
	s.VideoCalculatorResult = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutput) SetVideoCaptionResult(v *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResult) *GetVideoAnalysisTaskResponseBodyDataPayloadOutput {
	s.VideoCaptionResult = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutput) SetVideoGenerateResult(v *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult) *GetVideoAnalysisTaskResponseBodyDataPayloadOutput {
	s.VideoGenerateResult = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutput) SetVideoGenerateResults(v []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults) *GetVideoAnalysisTaskResponseBodyDataPayloadOutput {
	s.VideoGenerateResults = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutput) SetVideoMindMappingGenerateResult(v *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResult) *GetVideoAnalysisTaskResponseBodyDataPayloadOutput {
	s.VideoMindMappingGenerateResult = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutput) SetVideoRoleRecognitionResult(v *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResult) *GetVideoAnalysisTaskResponseBodyDataPayloadOutput {
	s.VideoRoleRecognitionResult = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutput) SetVideoTitleGenerateResult(v *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResult) *GetVideoAnalysisTaskResponseBodyDataPayloadOutput {
	s.VideoTitleGenerateResult = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutput) Validate() error {
	if s.AddDatasetDocumentsResult != nil {
		if err := s.AddDatasetDocumentsResult.Validate(); err != nil {
			return err
		}
	}
	if s.VideoAnalysisResult != nil {
		if err := s.VideoAnalysisResult.Validate(); err != nil {
			return err
		}
	}
	if s.VideoCalculatorResult != nil {
		if err := s.VideoCalculatorResult.Validate(); err != nil {
			return err
		}
	}
	if s.VideoCaptionResult != nil {
		if err := s.VideoCaptionResult.Validate(); err != nil {
			return err
		}
	}
	if s.VideoGenerateResult != nil {
		if err := s.VideoGenerateResult.Validate(); err != nil {
			return err
		}
	}
	if s.VideoGenerateResults != nil {
		for _, item := range s.VideoGenerateResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VideoMindMappingGenerateResult != nil {
		if err := s.VideoMindMappingGenerateResult.Validate(); err != nil {
			return err
		}
	}
	if s.VideoRoleRecognitionResult != nil {
		if err := s.VideoRoleRecognitionResult.Validate(); err != nil {
			return err
		}
	}
	if s.VideoTitleGenerateResult != nil {
		if err := s.VideoTitleGenerateResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputAddDatasetDocumentsResult struct {
	DocId        *string `json:"docId,omitempty" xml:"docId,omitempty"`
	DocUuid      *string `json:"docUuid,omitempty" xml:"docUuid,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Status       *int64  `json:"status,omitempty" xml:"status,omitempty"`
	Title        *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputAddDatasetDocumentsResult) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputAddDatasetDocumentsResult) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputAddDatasetDocumentsResult) GetDocId() *string {
	return s.DocId
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputAddDatasetDocumentsResult) GetDocUuid() *string {
	return s.DocUuid
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputAddDatasetDocumentsResult) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputAddDatasetDocumentsResult) GetStatus() *int64 {
	return s.Status
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputAddDatasetDocumentsResult) GetTitle() *string {
	return s.Title
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputAddDatasetDocumentsResult) SetDocId(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputAddDatasetDocumentsResult {
	s.DocId = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputAddDatasetDocumentsResult) SetDocUuid(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputAddDatasetDocumentsResult {
	s.DocUuid = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputAddDatasetDocumentsResult) SetErrorMessage(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputAddDatasetDocumentsResult {
	s.ErrorMessage = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputAddDatasetDocumentsResult) SetStatus(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputAddDatasetDocumentsResult {
	s.Status = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputAddDatasetDocumentsResult) SetTitle(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputAddDatasetDocumentsResult {
	s.Title = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputAddDatasetDocumentsResult) Validate() error {
	return dara.Validate(s)
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResult struct {
	// example:
	//
	// true
	GenerateFinished *bool `json:"generateFinished,omitempty" xml:"generateFinished,omitempty"`
	// example:
	//
	// xxx
	Text                     *string                                                                                         `json:"text,omitempty" xml:"text,omitempty"`
	Usage                    *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultUsage                      `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
	VideoShotAnalysisResults []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultVideoShotAnalysisResults `json:"videoShotAnalysisResults,omitempty" xml:"videoShotAnalysisResults,omitempty" type:"Repeated"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResult) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResult) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResult) GetGenerateFinished() *bool {
	return s.GenerateFinished
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResult) GetText() *string {
	return s.Text
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResult) GetUsage() *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultUsage {
	return s.Usage
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResult) GetVideoShotAnalysisResults() []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultVideoShotAnalysisResults {
	return s.VideoShotAnalysisResults
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResult) SetGenerateFinished(v bool) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResult {
	s.GenerateFinished = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResult) SetText(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResult {
	s.Text = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResult) SetUsage(v *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultUsage) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResult {
	s.Usage = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResult) SetVideoShotAnalysisResults(v []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultVideoShotAnalysisResults) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResult {
	s.VideoShotAnalysisResults = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResult) Validate() error {
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	if s.VideoShotAnalysisResults != nil {
		for _, item := range s.VideoShotAnalysisResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultUsage struct {
	ImageTokens *int64 `json:"imageTokens,omitempty" xml:"imageTokens,omitempty"`
	// example:
	//
	// 0
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 0
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 0
	TotalTokens *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultUsage) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultUsage) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultUsage) GetImageTokens() *int64 {
	return s.ImageTokens
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultUsage) SetImageTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultUsage {
	s.ImageTokens = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultUsage) SetInputTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultUsage {
	s.InputTokens = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultUsage) SetOutputTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultUsage {
	s.OutputTokens = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultUsage) SetTotalTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultUsage {
	s.TotalTokens = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultUsage) Validate() error {
	return dara.Validate(s)
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultVideoShotAnalysisResults struct {
	// example:
	//
	// 1710432000000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 2024-10-05 06:22:00
	StartTime *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Text      *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultVideoShotAnalysisResults) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultVideoShotAnalysisResults) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultVideoShotAnalysisResults) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultVideoShotAnalysisResults) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultVideoShotAnalysisResults) GetText() *string {
	return s.Text
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultVideoShotAnalysisResults) SetEndTime(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultVideoShotAnalysisResults {
	s.EndTime = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultVideoShotAnalysisResults) SetStartTime(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultVideoShotAnalysisResults {
	s.StartTime = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultVideoShotAnalysisResults) SetText(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultVideoShotAnalysisResults {
	s.Text = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoAnalysisResultVideoShotAnalysisResults) Validate() error {
	return dara.Validate(s)
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResult struct {
	Items []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResultItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResult) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResult) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResult) GetItems() []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResultItems {
	return s.Items
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResult) SetItems(v []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResultItems) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResult {
	s.Items = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResult) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResultItems struct {
	InputExpense *float64 `json:"inputExpense,omitempty" xml:"inputExpense,omitempty"`
	InputToken   *int64   `json:"inputToken,omitempty" xml:"inputToken,omitempty"`
	// example:
	//
	// xxx
	Name          *string  `json:"name,omitempty" xml:"name,omitempty"`
	OutputExpense *float64 `json:"outputExpense,omitempty" xml:"outputExpense,omitempty"`
	OutputToken   *int64   `json:"outputToken,omitempty" xml:"outputToken,omitempty"`
	Time          *int64   `json:"time,omitempty" xml:"time,omitempty"`
	TimeExpense   *float64 `json:"timeExpense,omitempty" xml:"timeExpense,omitempty"`
	// example:
	//
	// 0.098
	TotalExpense *float64 `json:"totalExpense,omitempty" xml:"totalExpense,omitempty"`
	Type         *string  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResultItems) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResultItems) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResultItems) GetInputExpense() *float64 {
	return s.InputExpense
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResultItems) GetInputToken() *int64 {
	return s.InputToken
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResultItems) GetName() *string {
	return s.Name
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResultItems) GetOutputExpense() *float64 {
	return s.OutputExpense
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResultItems) GetOutputToken() *int64 {
	return s.OutputToken
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResultItems) GetTime() *int64 {
	return s.Time
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResultItems) GetTimeExpense() *float64 {
	return s.TimeExpense
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResultItems) GetTotalExpense() *float64 {
	return s.TotalExpense
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResultItems) GetType() *string {
	return s.Type
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResultItems) SetInputExpense(v float64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResultItems {
	s.InputExpense = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResultItems) SetInputToken(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResultItems {
	s.InputToken = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResultItems) SetName(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResultItems {
	s.Name = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResultItems) SetOutputExpense(v float64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResultItems {
	s.OutputExpense = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResultItems) SetOutputToken(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResultItems {
	s.OutputToken = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResultItems) SetTime(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResultItems {
	s.Time = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResultItems) SetTimeExpense(v float64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResultItems {
	s.TimeExpense = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResultItems) SetTotalExpense(v float64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResultItems {
	s.TotalExpense = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResultItems) SetType(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResultItems {
	s.Type = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCalculatorResultItems) Validate() error {
	return dara.Validate(s)
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResult struct {
	// example:
	//
	// true
	GenerateFinished *bool                                                                               `json:"generateFinished,omitempty" xml:"generateFinished,omitempty"`
	VideoCaptions    []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions `json:"videoCaptions,omitempty" xml:"videoCaptions,omitempty" type:"Repeated"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResult) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResult) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResult) GetGenerateFinished() *bool {
	return s.GenerateFinished
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResult) GetVideoCaptions() []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions {
	return s.VideoCaptions
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResult) SetGenerateFinished(v bool) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResult {
	s.GenerateFinished = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResult) SetVideoCaptions(v []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResult {
	s.VideoCaptions = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResult) Validate() error {
	if s.VideoCaptions != nil {
		for _, item := range s.VideoCaptions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions struct {
	// example:
	//
	// 1736129678000
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
	// 00:01
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 2025-01-07 11:52:06
	StartTimeFormat *string `json:"startTimeFormat,omitempty" xml:"startTimeFormat,omitempty"`
	// example:
	//
	// xxxx
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions) GetEndTimeFormat() *string {
	return s.EndTimeFormat
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions) GetSpeaker() *string {
	return s.Speaker
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions) GetStartTimeFormat() *string {
	return s.StartTimeFormat
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions) GetText() *string {
	return s.Text
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions) SetEndTime(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions {
	s.EndTime = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions) SetEndTimeFormat(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions {
	s.EndTimeFormat = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions) SetSpeaker(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions {
	s.Speaker = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions) SetStartTime(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions {
	s.StartTime = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions) SetStartTimeFormat(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions {
	s.StartTimeFormat = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions) SetText(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions {
	s.Text = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoCaptionResultVideoCaptions) Validate() error {
	return dara.Validate(s)
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult struct {
	// example:
	//
	// true
	GenerateFinished *bool   `json:"generateFinished,omitempty" xml:"generateFinished,omitempty"`
	Index            *int32  `json:"index,omitempty" xml:"index,omitempty"`
	ModelId          *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	ModelReduce      *bool   `json:"modelReduce,omitempty" xml:"modelReduce,omitempty"`
	ReasonText       *string `json:"reasonText,omitempty" xml:"reasonText,omitempty"`
	// example:
	//
	// xxx
	Text  *string                                                                    `json:"text,omitempty" xml:"text,omitempty"`
	Usage *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult) GetGenerateFinished() *bool {
	return s.GenerateFinished
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult) GetIndex() *int32 {
	return s.Index
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult) GetModelId() *string {
	return s.ModelId
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult) GetModelReduce() *bool {
	return s.ModelReduce
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult) GetReasonText() *string {
	return s.ReasonText
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult) GetText() *string {
	return s.Text
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult) GetUsage() *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultUsage {
	return s.Usage
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult) SetGenerateFinished(v bool) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult {
	s.GenerateFinished = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult) SetIndex(v int32) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult {
	s.Index = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult) SetModelId(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult {
	s.ModelId = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult) SetModelReduce(v bool) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult {
	s.ModelReduce = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult) SetReasonText(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult {
	s.ReasonText = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult) SetText(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult {
	s.Text = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult) SetUsage(v *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultUsage) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult {
	s.Usage = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResult) Validate() error {
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultUsage struct {
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

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultUsage) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultUsage) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultUsage) SetInputTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultUsage {
	s.InputTokens = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultUsage) SetOutputTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultUsage {
	s.OutputTokens = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultUsage) SetTotalTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultUsage {
	s.TotalTokens = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultUsage) Validate() error {
	return dara.Validate(s)
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults struct {
	GenerateFinished *bool                                                                       `json:"generateFinished,omitempty" xml:"generateFinished,omitempty"`
	Index            *int32                                                                      `json:"index,omitempty" xml:"index,omitempty"`
	ModelId          *string                                                                     `json:"modelId,omitempty" xml:"modelId,omitempty"`
	ReasonText       *string                                                                     `json:"reasonText,omitempty" xml:"reasonText,omitempty"`
	Text             *string                                                                     `json:"text,omitempty" xml:"text,omitempty"`
	Usage            *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultsUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults) GetGenerateFinished() *bool {
	return s.GenerateFinished
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults) GetIndex() *int32 {
	return s.Index
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults) GetModelId() *string {
	return s.ModelId
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults) GetReasonText() *string {
	return s.ReasonText
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults) GetText() *string {
	return s.Text
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults) GetUsage() *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultsUsage {
	return s.Usage
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults) SetGenerateFinished(v bool) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults {
	s.GenerateFinished = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults) SetIndex(v int32) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults {
	s.Index = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults) SetModelId(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults {
	s.ModelId = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults) SetReasonText(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults {
	s.ReasonText = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults) SetText(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults {
	s.Text = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults) SetUsage(v *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultsUsage) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults {
	s.Usage = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResults) Validate() error {
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultsUsage struct {
	InputTokens  *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	TotalTokens  *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultsUsage) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultsUsage) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultsUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultsUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultsUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultsUsage) SetInputTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultsUsage {
	s.InputTokens = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultsUsage) SetOutputTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultsUsage {
	s.OutputTokens = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultsUsage) SetTotalTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultsUsage {
	s.TotalTokens = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoGenerateResultsUsage) Validate() error {
	return dara.Validate(s)
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResult struct {
	// example:
	//
	// true
	GenerateFinished  *bool                                                                                               `json:"generateFinished,omitempty" xml:"generateFinished,omitempty"`
	Text              *string                                                                                             `json:"text,omitempty" xml:"text,omitempty"`
	Usage             *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultUsage               `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
	VideoMindMappings []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappings `json:"videoMindMappings,omitempty" xml:"videoMindMappings,omitempty" type:"Repeated"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResult) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResult) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResult) GetGenerateFinished() *bool {
	return s.GenerateFinished
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResult) GetText() *string {
	return s.Text
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResult) GetUsage() *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultUsage {
	return s.Usage
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResult) GetVideoMindMappings() []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappings {
	return s.VideoMindMappings
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResult) SetGenerateFinished(v bool) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResult {
	s.GenerateFinished = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResult) SetText(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResult {
	s.Text = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResult) SetUsage(v *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultUsage) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResult {
	s.Usage = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResult) SetVideoMindMappings(v []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappings) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResult {
	s.VideoMindMappings = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResult) Validate() error {
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	if s.VideoMindMappings != nil {
		for _, item := range s.VideoMindMappings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultUsage struct {
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

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultUsage) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultUsage) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultUsage) SetInputTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultUsage {
	s.InputTokens = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultUsage) SetOutputTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultUsage {
	s.OutputTokens = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultUsage) SetTotalTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultUsage {
	s.TotalTokens = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultUsage) Validate() error {
	return dara.Validate(s)
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappings struct {
	ChildNodes []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes `json:"childNodes,omitempty" xml:"childNodes,omitempty" type:"Repeated"`
	Name       *string                                                                                                       `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappings) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappings) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappings) GetChildNodes() []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes {
	return s.ChildNodes
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappings) GetName() *string {
	return s.Name
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappings) SetChildNodes(v []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappings {
	s.ChildNodes = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappings) SetName(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappings {
	s.Name = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappings) Validate() error {
	if s.ChildNodes != nil {
		for _, item := range s.ChildNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes struct {
	ChildNodes []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodesChildNodes `json:"childNodes,omitempty" xml:"childNodes,omitempty" type:"Repeated"`
	Name       *string                                                                                                                 `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes) GetChildNodes() []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodesChildNodes {
	return s.ChildNodes
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes) GetName() *string {
	return s.Name
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes) SetChildNodes(v []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodesChildNodes) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes {
	s.ChildNodes = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes) SetName(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes {
	s.Name = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodes) Validate() error {
	if s.ChildNodes != nil {
		for _, item := range s.ChildNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodesChildNodes struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodesChildNodes) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodesChildNodes) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodesChildNodes) GetName() *string {
	return s.Name
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodesChildNodes) SetName(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodesChildNodes {
	s.Name = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoMindMappingGenerateResultVideoMindMappingsChildNodesChildNodes) Validate() error {
	return dara.Validate(s)
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResult struct {
	VideoRoles []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRoles `json:"videoRoles,omitempty" xml:"videoRoles,omitempty" type:"Repeated"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResult) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResult) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResult) GetVideoRoles() []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRoles {
	return s.VideoRoles
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResult) SetVideoRoles(v []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRoles) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResult {
	s.VideoRoles = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResult) Validate() error {
	if s.VideoRoles != nil {
		for _, item := range s.VideoRoles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRoles struct {
	IsAutoRecognition *bool                                                                                                 `json:"isAutoRecognition,omitempty" xml:"isAutoRecognition,omitempty"`
	Ratio             *float32                                                                                              `json:"ratio,omitempty" xml:"ratio,omitempty"`
	RoleInfo          *string                                                                                               `json:"roleInfo,omitempty" xml:"roleInfo,omitempty"`
	RoleName          *string                                                                                               `json:"roleName,omitempty" xml:"roleName,omitempty"`
	TimeIntervals     []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRolesTimeIntervals `json:"timeIntervals,omitempty" xml:"timeIntervals,omitempty" type:"Repeated"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRoles) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRoles) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRoles) GetIsAutoRecognition() *bool {
	return s.IsAutoRecognition
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRoles) GetRatio() *float32 {
	return s.Ratio
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRoles) GetRoleInfo() *string {
	return s.RoleInfo
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRoles) GetRoleName() *string {
	return s.RoleName
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRoles) GetTimeIntervals() []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRolesTimeIntervals {
	return s.TimeIntervals
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRoles) SetIsAutoRecognition(v bool) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRoles {
	s.IsAutoRecognition = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRoles) SetRatio(v float32) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRoles {
	s.Ratio = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRoles) SetRoleInfo(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRoles {
	s.RoleInfo = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRoles) SetRoleName(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRoles {
	s.RoleName = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRoles) SetTimeIntervals(v []*GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRolesTimeIntervals) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRoles {
	s.TimeIntervals = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRoles) Validate() error {
	if s.TimeIntervals != nil {
		for _, item := range s.TimeIntervals {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRolesTimeIntervals struct {
	EndTime   *int64  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	StartTime *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Timestamp *int64  `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	Url       *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRolesTimeIntervals) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRolesTimeIntervals) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRolesTimeIntervals) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRolesTimeIntervals) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRolesTimeIntervals) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRolesTimeIntervals) GetUrl() *string {
	return s.Url
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRolesTimeIntervals) SetEndTime(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRolesTimeIntervals {
	s.EndTime = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRolesTimeIntervals) SetStartTime(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRolesTimeIntervals {
	s.StartTime = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRolesTimeIntervals) SetTimestamp(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRolesTimeIntervals {
	s.Timestamp = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRolesTimeIntervals) SetUrl(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRolesTimeIntervals {
	s.Url = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoRoleRecognitionResultVideoRolesTimeIntervals) Validate() error {
	return dara.Validate(s)
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResult struct {
	// example:
	//
	// true
	GenerateFinished *bool `json:"generateFinished,omitempty" xml:"generateFinished,omitempty"`
	// example:
	//
	// xxxx
	Text  *string                                                                         `json:"text,omitempty" xml:"text,omitempty"`
	Usage *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResultUsage `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResult) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResult) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResult) GetGenerateFinished() *bool {
	return s.GenerateFinished
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResult) GetText() *string {
	return s.Text
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResult) GetUsage() *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResultUsage {
	return s.Usage
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResult) SetGenerateFinished(v bool) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResult {
	s.GenerateFinished = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResult) SetText(v string) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResult {
	s.Text = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResult) SetUsage(v *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResultUsage) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResult {
	s.Usage = v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResult) Validate() error {
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResultUsage struct {
	// example:
	//
	// 0
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 0
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 0
	TotalTokens *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResultUsage) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResultUsage) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResultUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResultUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResultUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResultUsage) SetInputTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResultUsage {
	s.InputTokens = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResultUsage) SetOutputTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResultUsage {
	s.OutputTokens = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResultUsage) SetTotalTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResultUsage {
	s.TotalTokens = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadOutputVideoTitleGenerateResultUsage) Validate() error {
	return dara.Validate(s)
}

type GetVideoAnalysisTaskResponseBodyDataPayloadUsage struct {
	// example:
	//
	// 0
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 0
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// 0
	TotalTokens *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataPayloadUsage) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadUsage) SetInputTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadUsage) SetOutputTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadUsage) SetTotalTokens(v int64) *GetVideoAnalysisTaskResponseBodyDataPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataPayloadUsage) Validate() error {
	return dara.Validate(s)
}

type GetVideoAnalysisTaskResponseBodyDataTaskRunInfo struct {
	// example:
	//
	// true
	ConcurrentChargeEnable *bool `json:"concurrentChargeEnable,omitempty" xml:"concurrentChargeEnable,omitempty"`
	// example:
	//
	// 1
	ResponseTime *int64 `json:"responseTime,omitempty" xml:"responseTime,omitempty"`
}

func (s GetVideoAnalysisTaskResponseBodyDataTaskRunInfo) String() string {
	return dara.Prettify(s)
}

func (s GetVideoAnalysisTaskResponseBodyDataTaskRunInfo) GoString() string {
	return s.String()
}

func (s *GetVideoAnalysisTaskResponseBodyDataTaskRunInfo) GetConcurrentChargeEnable() *bool {
	return s.ConcurrentChargeEnable
}

func (s *GetVideoAnalysisTaskResponseBodyDataTaskRunInfo) GetResponseTime() *int64 {
	return s.ResponseTime
}

func (s *GetVideoAnalysisTaskResponseBodyDataTaskRunInfo) SetConcurrentChargeEnable(v bool) *GetVideoAnalysisTaskResponseBodyDataTaskRunInfo {
	s.ConcurrentChargeEnable = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataTaskRunInfo) SetResponseTime(v int64) *GetVideoAnalysisTaskResponseBodyDataTaskRunInfo {
	s.ResponseTime = &v
	return s
}

func (s *GetVideoAnalysisTaskResponseBodyDataTaskRunInfo) Validate() error {
	return dara.Validate(s)
}
