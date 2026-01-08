// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoCreationTaskResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetVideoCreationTaskResultResponseBody
	GetCode() *string
	SetData(v *GetVideoCreationTaskResultResponseBodyData) *GetVideoCreationTaskResultResponseBody
	GetData() *GetVideoCreationTaskResultResponseBodyData
	SetMessage(v string) *GetVideoCreationTaskResultResponseBody
	GetMessage() *string
	SetRetryAble(v bool) *GetVideoCreationTaskResultResponseBody
	GetRetryAble() *bool
	SetSuccess(v bool) *GetVideoCreationTaskResultResponseBody
	GetSuccess() *bool
}

type GetVideoCreationTaskResultResponseBody struct {
	// example:
	//
	// 0
	Code *string                                     `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetVideoCreationTaskResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 成功
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RetryAble *bool   `json:"retryAble,omitempty" xml:"retryAble,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetVideoCreationTaskResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVideoCreationTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoCreationTaskResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetVideoCreationTaskResultResponseBody) GetData() *GetVideoCreationTaskResultResponseBodyData {
	return s.Data
}

func (s *GetVideoCreationTaskResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetVideoCreationTaskResultResponseBody) GetRetryAble() *bool {
	return s.RetryAble
}

func (s *GetVideoCreationTaskResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetVideoCreationTaskResultResponseBody) SetCode(v string) *GetVideoCreationTaskResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetVideoCreationTaskResultResponseBody) SetData(v *GetVideoCreationTaskResultResponseBodyData) *GetVideoCreationTaskResultResponseBody {
	s.Data = v
	return s
}

func (s *GetVideoCreationTaskResultResponseBody) SetMessage(v string) *GetVideoCreationTaskResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetVideoCreationTaskResultResponseBody) SetRetryAble(v bool) *GetVideoCreationTaskResultResponseBody {
	s.RetryAble = &v
	return s
}

func (s *GetVideoCreationTaskResultResponseBody) SetSuccess(v bool) *GetVideoCreationTaskResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetVideoCreationTaskResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVideoCreationTaskResultResponseBodyData struct {
	// example:
	//
	// 1
	EstimatedWaitTimeInSeconds *int64                                              `json:"estimatedWaitTimeInSeconds,omitempty" xml:"estimatedWaitTimeInSeconds,omitempty"`
	FileInfo                   *GetVideoCreationTaskResultResponseBodyDataFileInfo `json:"fileInfo,omitempty" xml:"fileInfo,omitempty" type:"Struct"`
	// example:
	//
	// xxx
	FinishTime               *string                                                             `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	MediaDetectionTaskResult *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResult `json:"mediaDetectionTaskResult,omitempty" xml:"mediaDetectionTaskResult,omitempty" type:"Struct"`
	// example:
	//
	// xxx
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 1
	StatusReason *string `json:"statusReason,omitempty" xml:"statusReason,omitempty"`
	// example:
	//
	// xxx
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// xxx
	TaskStatus *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
	// example:
	//
	// xxx
	VideoUrl *string `json:"videoUrl,omitempty" xml:"videoUrl,omitempty"`
}

func (s GetVideoCreationTaskResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetVideoCreationTaskResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVideoCreationTaskResultResponseBodyData) GetEstimatedWaitTimeInSeconds() *int64 {
	return s.EstimatedWaitTimeInSeconds
}

func (s *GetVideoCreationTaskResultResponseBodyData) GetFileInfo() *GetVideoCreationTaskResultResponseBodyDataFileInfo {
	return s.FileInfo
}

func (s *GetVideoCreationTaskResultResponseBodyData) GetFinishTime() *string {
	return s.FinishTime
}

func (s *GetVideoCreationTaskResultResponseBodyData) GetMediaDetectionTaskResult() *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResult {
	return s.MediaDetectionTaskResult
}

func (s *GetVideoCreationTaskResultResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *GetVideoCreationTaskResultResponseBodyData) GetStatusReason() *string {
	return s.StatusReason
}

func (s *GetVideoCreationTaskResultResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetVideoCreationTaskResultResponseBodyData) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *GetVideoCreationTaskResultResponseBodyData) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *GetVideoCreationTaskResultResponseBodyData) SetEstimatedWaitTimeInSeconds(v int64) *GetVideoCreationTaskResultResponseBodyData {
	s.EstimatedWaitTimeInSeconds = &v
	return s
}

func (s *GetVideoCreationTaskResultResponseBodyData) SetFileInfo(v *GetVideoCreationTaskResultResponseBodyDataFileInfo) *GetVideoCreationTaskResultResponseBodyData {
	s.FileInfo = v
	return s
}

func (s *GetVideoCreationTaskResultResponseBodyData) SetFinishTime(v string) *GetVideoCreationTaskResultResponseBodyData {
	s.FinishTime = &v
	return s
}

func (s *GetVideoCreationTaskResultResponseBodyData) SetMediaDetectionTaskResult(v *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResult) *GetVideoCreationTaskResultResponseBodyData {
	s.MediaDetectionTaskResult = v
	return s
}

func (s *GetVideoCreationTaskResultResponseBodyData) SetStartTime(v string) *GetVideoCreationTaskResultResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetVideoCreationTaskResultResponseBodyData) SetStatusReason(v string) *GetVideoCreationTaskResultResponseBodyData {
	s.StatusReason = &v
	return s
}

func (s *GetVideoCreationTaskResultResponseBodyData) SetTaskId(v string) *GetVideoCreationTaskResultResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetVideoCreationTaskResultResponseBodyData) SetTaskStatus(v string) *GetVideoCreationTaskResultResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *GetVideoCreationTaskResultResponseBodyData) SetVideoUrl(v string) *GetVideoCreationTaskResultResponseBodyData {
	s.VideoUrl = &v
	return s
}

func (s *GetVideoCreationTaskResultResponseBodyData) Validate() error {
	if s.FileInfo != nil {
		if err := s.FileInfo.Validate(); err != nil {
			return err
		}
	}
	if s.MediaDetectionTaskResult != nil {
		if err := s.MediaDetectionTaskResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVideoCreationTaskResultResponseBodyDataFileInfo struct {
	// example:
	//
	// xxx
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// example:
	//
	// xxx
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// xxx
	FileTraceId *string `json:"fileTraceId,omitempty" xml:"fileTraceId,omitempty"`
	// example:
	//
	// xxx
	OssKey *string `json:"ossKey,omitempty" xml:"ossKey,omitempty"`
}

func (s GetVideoCreationTaskResultResponseBodyDataFileInfo) String() string {
	return dara.Prettify(s)
}

func (s GetVideoCreationTaskResultResponseBodyDataFileInfo) GoString() string {
	return s.String()
}

func (s *GetVideoCreationTaskResultResponseBodyDataFileInfo) GetFileId() *string {
	return s.FileId
}

func (s *GetVideoCreationTaskResultResponseBodyDataFileInfo) GetFileName() *string {
	return s.FileName
}

func (s *GetVideoCreationTaskResultResponseBodyDataFileInfo) GetFileTraceId() *string {
	return s.FileTraceId
}

func (s *GetVideoCreationTaskResultResponseBodyDataFileInfo) GetOssKey() *string {
	return s.OssKey
}

func (s *GetVideoCreationTaskResultResponseBodyDataFileInfo) SetFileId(v string) *GetVideoCreationTaskResultResponseBodyDataFileInfo {
	s.FileId = &v
	return s
}

func (s *GetVideoCreationTaskResultResponseBodyDataFileInfo) SetFileName(v string) *GetVideoCreationTaskResultResponseBodyDataFileInfo {
	s.FileName = &v
	return s
}

func (s *GetVideoCreationTaskResultResponseBodyDataFileInfo) SetFileTraceId(v string) *GetVideoCreationTaskResultResponseBodyDataFileInfo {
	s.FileTraceId = &v
	return s
}

func (s *GetVideoCreationTaskResultResponseBodyDataFileInfo) SetOssKey(v string) *GetVideoCreationTaskResultResponseBodyDataFileInfo {
	s.OssKey = &v
	return s
}

func (s *GetVideoCreationTaskResultResponseBodyDataFileInfo) Validate() error {
	return dara.Validate(s)
}

type GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResult struct {
	// example:
	//
	// xxx
	DetectionConclusion *string                                                                            `json:"detectionConclusion,omitempty" xml:"detectionConclusion,omitempty"`
	DetectionResult     *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultDetectionResult `json:"detectionResult,omitempty" xml:"detectionResult,omitempty" type:"Struct"`
	FileInfo            *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultFileInfo        `json:"fileInfo,omitempty" xml:"fileInfo,omitempty" type:"Struct"`
	// example:
	//
	// xxx
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// xxx
	TaskStatus *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
}

func (s GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResult) String() string {
	return dara.Prettify(s)
}

func (s GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResult) GoString() string {
	return s.String()
}

func (s *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResult) GetDetectionConclusion() *string {
	return s.DetectionConclusion
}

func (s *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResult) GetDetectionResult() *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultDetectionResult {
	return s.DetectionResult
}

func (s *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResult) GetFileInfo() *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultFileInfo {
	return s.FileInfo
}

func (s *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResult) GetTaskId() *string {
	return s.TaskId
}

func (s *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResult) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResult) SetDetectionConclusion(v string) *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResult {
	s.DetectionConclusion = &v
	return s
}

func (s *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResult) SetDetectionResult(v *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultDetectionResult) *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResult {
	s.DetectionResult = v
	return s
}

func (s *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResult) SetFileInfo(v *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultFileInfo) *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResult {
	s.FileInfo = v
	return s
}

func (s *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResult) SetTaskId(v string) *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResult {
	s.TaskId = &v
	return s
}

func (s *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResult) SetTaskStatus(v string) *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResult {
	s.TaskStatus = &v
	return s
}

func (s *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResult) Validate() error {
	if s.DetectionResult != nil {
		if err := s.DetectionResult.Validate(); err != nil {
			return err
		}
	}
	if s.FileInfo != nil {
		if err := s.FileInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultDetectionResult struct {
	DetectionDetails []*GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultDetectionResultDetectionDetails `json:"detectionDetails,omitempty" xml:"detectionDetails,omitempty" type:"Repeated"`
	Suggestions      []*string                                                                                            `json:"suggestions,omitempty" xml:"suggestions,omitempty" type:"Repeated"`
}

func (s GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultDetectionResult) String() string {
	return dara.Prettify(s)
}

func (s GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultDetectionResult) GoString() string {
	return s.String()
}

func (s *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultDetectionResult) GetDetectionDetails() []*GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultDetectionResultDetectionDetails {
	return s.DetectionDetails
}

func (s *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultDetectionResult) GetSuggestions() []*string {
	return s.Suggestions
}

func (s *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultDetectionResult) SetDetectionDetails(v []*GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultDetectionResultDetectionDetails) *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultDetectionResult {
	s.DetectionDetails = v
	return s
}

func (s *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultDetectionResult) SetSuggestions(v []*string) *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultDetectionResult {
	s.Suggestions = v
	return s
}

func (s *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultDetectionResult) Validate() error {
	if s.DetectionDetails != nil {
		for _, item := range s.DetectionDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultDetectionResultDetectionDetails struct {
	// example:
	//
	// xxx
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 0.8
	Confidence *float64 `json:"confidence,omitempty" xml:"confidence,omitempty"`
	Pass       *bool    `json:"pass,omitempty" xml:"pass,omitempty"`
	// example:
	//
	// xxx
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultDetectionResultDetectionDetails) String() string {
	return dara.Prettify(s)
}

func (s GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultDetectionResultDetectionDetails) GoString() string {
	return s.String()
}

func (s *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultDetectionResultDetectionDetails) GetCode() *string {
	return s.Code
}

func (s *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultDetectionResultDetectionDetails) GetConfidence() *float64 {
	return s.Confidence
}

func (s *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultDetectionResultDetectionDetails) GetPass() *bool {
	return s.Pass
}

func (s *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultDetectionResultDetectionDetails) GetReason() *string {
	return s.Reason
}

func (s *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultDetectionResultDetectionDetails) SetCode(v string) *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultDetectionResultDetectionDetails {
	s.Code = &v
	return s
}

func (s *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultDetectionResultDetectionDetails) SetConfidence(v float64) *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultDetectionResultDetectionDetails {
	s.Confidence = &v
	return s
}

func (s *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultDetectionResultDetectionDetails) SetPass(v bool) *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultDetectionResultDetectionDetails {
	s.Pass = &v
	return s
}

func (s *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultDetectionResultDetectionDetails) SetReason(v string) *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultDetectionResultDetectionDetails {
	s.Reason = &v
	return s
}

func (s *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultDetectionResultDetectionDetails) Validate() error {
	return dara.Validate(s)
}

type GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultFileInfo struct {
	// example:
	//
	// xxx
	FileId *string `json:"fileId,omitempty" xml:"fileId,omitempty"`
	// example:
	//
	// xxx
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// xxx
	FileTraceId *string `json:"fileTraceId,omitempty" xml:"fileTraceId,omitempty"`
	// example:
	//
	// xxx
	OssKey *string `json:"ossKey,omitempty" xml:"ossKey,omitempty"`
}

func (s GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultFileInfo) String() string {
	return dara.Prettify(s)
}

func (s GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultFileInfo) GoString() string {
	return s.String()
}

func (s *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultFileInfo) GetFileId() *string {
	return s.FileId
}

func (s *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultFileInfo) GetFileName() *string {
	return s.FileName
}

func (s *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultFileInfo) GetFileTraceId() *string {
	return s.FileTraceId
}

func (s *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultFileInfo) GetOssKey() *string {
	return s.OssKey
}

func (s *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultFileInfo) SetFileId(v string) *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultFileInfo {
	s.FileId = &v
	return s
}

func (s *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultFileInfo) SetFileName(v string) *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultFileInfo {
	s.FileName = &v
	return s
}

func (s *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultFileInfo) SetFileTraceId(v string) *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultFileInfo {
	s.FileTraceId = &v
	return s
}

func (s *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultFileInfo) SetOssKey(v string) *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultFileInfo {
	s.OssKey = &v
	return s
}

func (s *GetVideoCreationTaskResultResponseBodyDataMediaDetectionTaskResultFileInfo) Validate() error {
	return dara.Validate(s)
}
