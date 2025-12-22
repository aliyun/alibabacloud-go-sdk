// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageDetectionTaskResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetImageDetectionTaskResultResponseBody
	GetCode() *string
	SetData(v *GetImageDetectionTaskResultResponseBodyData) *GetImageDetectionTaskResultResponseBody
	GetData() *GetImageDetectionTaskResultResponseBodyData
	SetMessage(v string) *GetImageDetectionTaskResultResponseBody
	GetMessage() *string
	SetRetryAble(v bool) *GetImageDetectionTaskResultResponseBody
	GetRetryAble() *bool
	SetSuccess(v bool) *GetImageDetectionTaskResultResponseBody
	GetSuccess() *bool
}

type GetImageDetectionTaskResultResponseBody struct {
	// example:
	//
	// 0
	Code *string                                      `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetImageDetectionTaskResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 成功
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RetryAble *bool   `json:"retryAble,omitempty" xml:"retryAble,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetImageDetectionTaskResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetImageDetectionTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageDetectionTaskResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetImageDetectionTaskResultResponseBody) GetData() *GetImageDetectionTaskResultResponseBodyData {
	return s.Data
}

func (s *GetImageDetectionTaskResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetImageDetectionTaskResultResponseBody) GetRetryAble() *bool {
	return s.RetryAble
}

func (s *GetImageDetectionTaskResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetImageDetectionTaskResultResponseBody) SetCode(v string) *GetImageDetectionTaskResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBody) SetData(v *GetImageDetectionTaskResultResponseBodyData) *GetImageDetectionTaskResultResponseBody {
	s.Data = v
	return s
}

func (s *GetImageDetectionTaskResultResponseBody) SetMessage(v string) *GetImageDetectionTaskResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBody) SetRetryAble(v bool) *GetImageDetectionTaskResultResponseBody {
	s.RetryAble = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBody) SetSuccess(v bool) *GetImageDetectionTaskResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetImageDetectionTaskResultResponseBodyData struct {
	// example:
	//
	// xxx
	DetectionConclusion *string                                                     `json:"detectionConclusion,omitempty" xml:"detectionConclusion,omitempty"`
	DetectionResult     *GetImageDetectionTaskResultResponseBodyDataDetectionResult `json:"detectionResult,omitempty" xml:"detectionResult,omitempty" type:"Struct"`
	FileInfo            *GetImageDetectionTaskResultResponseBodyDataFileInfo        `json:"fileInfo,omitempty" xml:"fileInfo,omitempty" type:"Struct"`
	// example:
	//
	// xxx
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// xxx
	TaskStatus *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
}

func (s GetImageDetectionTaskResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetImageDetectionTaskResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetImageDetectionTaskResultResponseBodyData) GetDetectionConclusion() *string {
	return s.DetectionConclusion
}

func (s *GetImageDetectionTaskResultResponseBodyData) GetDetectionResult() *GetImageDetectionTaskResultResponseBodyDataDetectionResult {
	return s.DetectionResult
}

func (s *GetImageDetectionTaskResultResponseBodyData) GetFileInfo() *GetImageDetectionTaskResultResponseBodyDataFileInfo {
	return s.FileInfo
}

func (s *GetImageDetectionTaskResultResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetImageDetectionTaskResultResponseBodyData) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *GetImageDetectionTaskResultResponseBodyData) SetDetectionConclusion(v string) *GetImageDetectionTaskResultResponseBodyData {
	s.DetectionConclusion = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyData) SetDetectionResult(v *GetImageDetectionTaskResultResponseBodyDataDetectionResult) *GetImageDetectionTaskResultResponseBodyData {
	s.DetectionResult = v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyData) SetFileInfo(v *GetImageDetectionTaskResultResponseBodyDataFileInfo) *GetImageDetectionTaskResultResponseBodyData {
	s.FileInfo = v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyData) SetTaskId(v string) *GetImageDetectionTaskResultResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyData) SetTaskStatus(v string) *GetImageDetectionTaskResultResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyData) Validate() error {
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

type GetImageDetectionTaskResultResponseBodyDataDetectionResult struct {
	DetectionDetails []*GetImageDetectionTaskResultResponseBodyDataDetectionResultDetectionDetails `json:"detectionDetails,omitempty" xml:"detectionDetails,omitempty" type:"Repeated"`
	Suggestions      []*string                                                                     `json:"suggestions,omitempty" xml:"suggestions,omitempty" type:"Repeated"`
}

func (s GetImageDetectionTaskResultResponseBodyDataDetectionResult) String() string {
	return dara.Prettify(s)
}

func (s GetImageDetectionTaskResultResponseBodyDataDetectionResult) GoString() string {
	return s.String()
}

func (s *GetImageDetectionTaskResultResponseBodyDataDetectionResult) GetDetectionDetails() []*GetImageDetectionTaskResultResponseBodyDataDetectionResultDetectionDetails {
	return s.DetectionDetails
}

func (s *GetImageDetectionTaskResultResponseBodyDataDetectionResult) GetSuggestions() []*string {
	return s.Suggestions
}

func (s *GetImageDetectionTaskResultResponseBodyDataDetectionResult) SetDetectionDetails(v []*GetImageDetectionTaskResultResponseBodyDataDetectionResultDetectionDetails) *GetImageDetectionTaskResultResponseBodyDataDetectionResult {
	s.DetectionDetails = v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyDataDetectionResult) SetSuggestions(v []*string) *GetImageDetectionTaskResultResponseBodyDataDetectionResult {
	s.Suggestions = v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyDataDetectionResult) Validate() error {
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

type GetImageDetectionTaskResultResponseBodyDataDetectionResultDetectionDetails struct {
	// example:
	//
	// x
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// xxx
	Confidence *float64 `json:"confidence,omitempty" xml:"confidence,omitempty"`
	Pass       *bool    `json:"pass,omitempty" xml:"pass,omitempty"`
	// example:
	//
	// xxx
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
}

func (s GetImageDetectionTaskResultResponseBodyDataDetectionResultDetectionDetails) String() string {
	return dara.Prettify(s)
}

func (s GetImageDetectionTaskResultResponseBodyDataDetectionResultDetectionDetails) GoString() string {
	return s.String()
}

func (s *GetImageDetectionTaskResultResponseBodyDataDetectionResultDetectionDetails) GetCode() *string {
	return s.Code
}

func (s *GetImageDetectionTaskResultResponseBodyDataDetectionResultDetectionDetails) GetConfidence() *float64 {
	return s.Confidence
}

func (s *GetImageDetectionTaskResultResponseBodyDataDetectionResultDetectionDetails) GetPass() *bool {
	return s.Pass
}

func (s *GetImageDetectionTaskResultResponseBodyDataDetectionResultDetectionDetails) GetReason() *string {
	return s.Reason
}

func (s *GetImageDetectionTaskResultResponseBodyDataDetectionResultDetectionDetails) SetCode(v string) *GetImageDetectionTaskResultResponseBodyDataDetectionResultDetectionDetails {
	s.Code = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyDataDetectionResultDetectionDetails) SetConfidence(v float64) *GetImageDetectionTaskResultResponseBodyDataDetectionResultDetectionDetails {
	s.Confidence = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyDataDetectionResultDetectionDetails) SetPass(v bool) *GetImageDetectionTaskResultResponseBodyDataDetectionResultDetectionDetails {
	s.Pass = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyDataDetectionResultDetectionDetails) SetReason(v string) *GetImageDetectionTaskResultResponseBodyDataDetectionResultDetectionDetails {
	s.Reason = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyDataDetectionResultDetectionDetails) Validate() error {
	return dara.Validate(s)
}

type GetImageDetectionTaskResultResponseBodyDataFileInfo struct {
	// example:
	//
	// xx
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

func (s GetImageDetectionTaskResultResponseBodyDataFileInfo) String() string {
	return dara.Prettify(s)
}

func (s GetImageDetectionTaskResultResponseBodyDataFileInfo) GoString() string {
	return s.String()
}

func (s *GetImageDetectionTaskResultResponseBodyDataFileInfo) GetFileId() *string {
	return s.FileId
}

func (s *GetImageDetectionTaskResultResponseBodyDataFileInfo) GetFileName() *string {
	return s.FileName
}

func (s *GetImageDetectionTaskResultResponseBodyDataFileInfo) GetFileTraceId() *string {
	return s.FileTraceId
}

func (s *GetImageDetectionTaskResultResponseBodyDataFileInfo) GetOssKey() *string {
	return s.OssKey
}

func (s *GetImageDetectionTaskResultResponseBodyDataFileInfo) SetFileId(v string) *GetImageDetectionTaskResultResponseBodyDataFileInfo {
	s.FileId = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyDataFileInfo) SetFileName(v string) *GetImageDetectionTaskResultResponseBodyDataFileInfo {
	s.FileName = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyDataFileInfo) SetFileTraceId(v string) *GetImageDetectionTaskResultResponseBodyDataFileInfo {
	s.FileTraceId = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyDataFileInfo) SetOssKey(v string) *GetImageDetectionTaskResultResponseBodyDataFileInfo {
	s.OssKey = &v
	return s
}

func (s *GetImageDetectionTaskResultResponseBodyDataFileInfo) Validate() error {
	return dara.Validate(s)
}
