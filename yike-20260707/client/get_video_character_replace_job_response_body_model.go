// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoCharacterReplaceJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetVideoCharacterReplaceJobResponseBody
	GetAppId() *string
	SetAppParams(v string) *GetVideoCharacterReplaceJobResponseBody
	GetAppParams() *string
	SetExecutionFinishTime(v string) *GetVideoCharacterReplaceJobResponseBody
	GetExecutionFinishTime() *string
	SetExecutionStartTime(v string) *GetVideoCharacterReplaceJobResponseBody
	GetExecutionStartTime() *string
	SetFolderId(v string) *GetVideoCharacterReplaceJobResponseBody
	GetFolderId() *string
	SetJobId(v string) *GetVideoCharacterReplaceJobResponseBody
	GetJobId() *string
	SetProjectId(v string) *GetVideoCharacterReplaceJobResponseBody
	GetProjectId() *string
	SetRequestId(v string) *GetVideoCharacterReplaceJobResponseBody
	GetRequestId() *string
	SetResult(v *GetVideoCharacterReplaceJobResponseBodyResult) *GetVideoCharacterReplaceJobResponseBody
	GetResult() *GetVideoCharacterReplaceJobResponseBodyResult
	SetStatus(v string) *GetVideoCharacterReplaceJobResponseBody
	GetStatus() *string
}

type GetVideoCharacterReplaceJobResponseBody struct {
	// The ID of the video character replacement application. Use a configured application.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application input parameters, which are serialized into a JSON string. The field names and material requirements are subject to the input definition of the corresponding AppId.
	//
	// example:
	//
	// {"LoadImage.1.TargetImage":"1a7852b032c371f0b64fe6f6c74b****"}
	AppParams *string `json:"AppParams,omitempty" xml:"AppParams,omitempty"`
	// The time when the task finished execution, in ISO 8601 format with time zone.
	//
	// example:
	//
	// 2026-02-06T18:54:18.809+08:00
	ExecutionFinishTime *string `json:"ExecutionFinishTime,omitempty" xml:"ExecutionFinishTime,omitempty"`
	// The time when the task started execution, in ISO 8601 format with time zone.
	//
	// example:
	//
	// 2026-02-06T18:53:34.001+08:00
	ExecutionStartTime *string `json:"ExecutionStartTime,omitempty" xml:"ExecutionStartTime,omitempty"`
	// The folder ID.
	//
	// example:
	//
	// folder-u3ilwhoc36ux9a****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The task ID, which is returned by SubmitVideoCharacterReplaceJob.
	//
	// example:
	//
	// ag_83baeea38b534ee1937b8a4acd08****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// pd_061716****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The request ID, which is used to locate and troubleshoot the request.
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task output results, which are categorized by image, video, and audio.
	Result *GetVideoCharacterReplaceJobResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The task status. Valid values:
	//
	// - Created: The task is created.
	//
	// - Queuing: The task is queuing.
	//
	// - Executing: The task is executing.
	//
	// - Finished: The task is finished.
	//
	// - Failed: The task failed.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetVideoCharacterReplaceJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVideoCharacterReplaceJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoCharacterReplaceJobResponseBody) GetAppId() *string {
	return s.AppId
}

func (s *GetVideoCharacterReplaceJobResponseBody) GetAppParams() *string {
	return s.AppParams
}

func (s *GetVideoCharacterReplaceJobResponseBody) GetExecutionFinishTime() *string {
	return s.ExecutionFinishTime
}

func (s *GetVideoCharacterReplaceJobResponseBody) GetExecutionStartTime() *string {
	return s.ExecutionStartTime
}

func (s *GetVideoCharacterReplaceJobResponseBody) GetFolderId() *string {
	return s.FolderId
}

func (s *GetVideoCharacterReplaceJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *GetVideoCharacterReplaceJobResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetVideoCharacterReplaceJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVideoCharacterReplaceJobResponseBody) GetResult() *GetVideoCharacterReplaceJobResponseBodyResult {
	return s.Result
}

func (s *GetVideoCharacterReplaceJobResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetVideoCharacterReplaceJobResponseBody) SetAppId(v string) *GetVideoCharacterReplaceJobResponseBody {
	s.AppId = &v
	return s
}

func (s *GetVideoCharacterReplaceJobResponseBody) SetAppParams(v string) *GetVideoCharacterReplaceJobResponseBody {
	s.AppParams = &v
	return s
}

func (s *GetVideoCharacterReplaceJobResponseBody) SetExecutionFinishTime(v string) *GetVideoCharacterReplaceJobResponseBody {
	s.ExecutionFinishTime = &v
	return s
}

func (s *GetVideoCharacterReplaceJobResponseBody) SetExecutionStartTime(v string) *GetVideoCharacterReplaceJobResponseBody {
	s.ExecutionStartTime = &v
	return s
}

func (s *GetVideoCharacterReplaceJobResponseBody) SetFolderId(v string) *GetVideoCharacterReplaceJobResponseBody {
	s.FolderId = &v
	return s
}

func (s *GetVideoCharacterReplaceJobResponseBody) SetJobId(v string) *GetVideoCharacterReplaceJobResponseBody {
	s.JobId = &v
	return s
}

func (s *GetVideoCharacterReplaceJobResponseBody) SetProjectId(v string) *GetVideoCharacterReplaceJobResponseBody {
	s.ProjectId = &v
	return s
}

func (s *GetVideoCharacterReplaceJobResponseBody) SetRequestId(v string) *GetVideoCharacterReplaceJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoCharacterReplaceJobResponseBody) SetResult(v *GetVideoCharacterReplaceJobResponseBodyResult) *GetVideoCharacterReplaceJobResponseBody {
	s.Result = v
	return s
}

func (s *GetVideoCharacterReplaceJobResponseBody) SetStatus(v string) *GetVideoCharacterReplaceJobResponseBody {
	s.Status = &v
	return s
}

func (s *GetVideoCharacterReplaceJobResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVideoCharacterReplaceJobResponseBodyResult struct {
	// The list of output audio.
	AudioResult []*GetVideoCharacterReplaceJobResponseBodyResultAudioResult `json:"AudioResult,omitempty" xml:"AudioResult,omitempty" type:"Repeated"`
	// The list of output images.
	ImageResult []*GetVideoCharacterReplaceJobResponseBodyResultImageResult `json:"ImageResult,omitempty" xml:"ImageResult,omitempty" type:"Repeated"`
	// The list of output videos.
	VideoResult []*GetVideoCharacterReplaceJobResponseBodyResultVideoResult `json:"VideoResult,omitempty" xml:"VideoResult,omitempty" type:"Repeated"`
}

func (s GetVideoCharacterReplaceJobResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetVideoCharacterReplaceJobResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetVideoCharacterReplaceJobResponseBodyResult) GetAudioResult() []*GetVideoCharacterReplaceJobResponseBodyResultAudioResult {
	return s.AudioResult
}

func (s *GetVideoCharacterReplaceJobResponseBodyResult) GetImageResult() []*GetVideoCharacterReplaceJobResponseBodyResultImageResult {
	return s.ImageResult
}

func (s *GetVideoCharacterReplaceJobResponseBodyResult) GetVideoResult() []*GetVideoCharacterReplaceJobResponseBodyResultVideoResult {
	return s.VideoResult
}

func (s *GetVideoCharacterReplaceJobResponseBodyResult) SetAudioResult(v []*GetVideoCharacterReplaceJobResponseBodyResultAudioResult) *GetVideoCharacterReplaceJobResponseBodyResult {
	s.AudioResult = v
	return s
}

func (s *GetVideoCharacterReplaceJobResponseBodyResult) SetImageResult(v []*GetVideoCharacterReplaceJobResponseBodyResultImageResult) *GetVideoCharacterReplaceJobResponseBodyResult {
	s.ImageResult = v
	return s
}

func (s *GetVideoCharacterReplaceJobResponseBodyResult) SetVideoResult(v []*GetVideoCharacterReplaceJobResponseBodyResultVideoResult) *GetVideoCharacterReplaceJobResponseBodyResult {
	s.VideoResult = v
	return s
}

func (s *GetVideoCharacterReplaceJobResponseBodyResult) Validate() error {
	if s.AudioResult != nil {
		for _, item := range s.AudioResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ImageResult != nil {
		for _, item := range s.ImageResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VideoResult != nil {
		for _, item := range s.VideoResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetVideoCharacterReplaceJobResponseBodyResultAudioResult struct {
	// The output media asset ID.
	//
	// example:
	//
	// 1a7852b032c371f0b64fe6f6c74b****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The download URL of the output media asset.
	//
	// example:
	//
	// https://outin-***.oss-cn-shanghai.aliyuncs.com/stream/48555e8b-181dd5a8c07/48555e8b-181dd5a8c07.mp3
	OutputUrl *string `json:"OutputUrl,omitempty" xml:"OutputUrl,omitempty"`
}

func (s GetVideoCharacterReplaceJobResponseBodyResultAudioResult) String() string {
	return dara.Prettify(s)
}

func (s GetVideoCharacterReplaceJobResponseBodyResultAudioResult) GoString() string {
	return s.String()
}

func (s *GetVideoCharacterReplaceJobResponseBodyResultAudioResult) GetMediaId() *string {
	return s.MediaId
}

func (s *GetVideoCharacterReplaceJobResponseBodyResultAudioResult) GetOutputUrl() *string {
	return s.OutputUrl
}

func (s *GetVideoCharacterReplaceJobResponseBodyResultAudioResult) SetMediaId(v string) *GetVideoCharacterReplaceJobResponseBodyResultAudioResult {
	s.MediaId = &v
	return s
}

func (s *GetVideoCharacterReplaceJobResponseBodyResultAudioResult) SetOutputUrl(v string) *GetVideoCharacterReplaceJobResponseBodyResultAudioResult {
	s.OutputUrl = &v
	return s
}

func (s *GetVideoCharacterReplaceJobResponseBodyResultAudioResult) Validate() error {
	return dara.Validate(s)
}

type GetVideoCharacterReplaceJobResponseBodyResultImageResult struct {
	// The output media asset ID.
	//
	// example:
	//
	// 1a7852b032c371f0b64fe6f6c74b****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The download URL of the output media asset.
	//
	// example:
	//
	// https://outin-***.oss-cn-shanghai.aliyuncs.com/stream/48555e8b-181dd5a8c07/48555e8b-181dd5a8c07.png
	OutputUrl *string `json:"OutputUrl,omitempty" xml:"OutputUrl,omitempty"`
}

func (s GetVideoCharacterReplaceJobResponseBodyResultImageResult) String() string {
	return dara.Prettify(s)
}

func (s GetVideoCharacterReplaceJobResponseBodyResultImageResult) GoString() string {
	return s.String()
}

func (s *GetVideoCharacterReplaceJobResponseBodyResultImageResult) GetMediaId() *string {
	return s.MediaId
}

func (s *GetVideoCharacterReplaceJobResponseBodyResultImageResult) GetOutputUrl() *string {
	return s.OutputUrl
}

func (s *GetVideoCharacterReplaceJobResponseBodyResultImageResult) SetMediaId(v string) *GetVideoCharacterReplaceJobResponseBodyResultImageResult {
	s.MediaId = &v
	return s
}

func (s *GetVideoCharacterReplaceJobResponseBodyResultImageResult) SetOutputUrl(v string) *GetVideoCharacterReplaceJobResponseBodyResultImageResult {
	s.OutputUrl = &v
	return s
}

func (s *GetVideoCharacterReplaceJobResponseBodyResultImageResult) Validate() error {
	return dara.Validate(s)
}

type GetVideoCharacterReplaceJobResponseBodyResultVideoResult struct {
	// The output media asset ID.
	//
	// example:
	//
	// 1a7852b032c371f0b64fe6f6c74b****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The download URL of the output media asset.
	//
	// example:
	//
	// https://outin-***.oss-cn-shanghai.aliyuncs.com/stream/48555e8b-181dd5a8c07/48555e8b-181dd5a8c07.mp4
	OutputUrl *string `json:"OutputUrl,omitempty" xml:"OutputUrl,omitempty"`
}

func (s GetVideoCharacterReplaceJobResponseBodyResultVideoResult) String() string {
	return dara.Prettify(s)
}

func (s GetVideoCharacterReplaceJobResponseBodyResultVideoResult) GoString() string {
	return s.String()
}

func (s *GetVideoCharacterReplaceJobResponseBodyResultVideoResult) GetMediaId() *string {
	return s.MediaId
}

func (s *GetVideoCharacterReplaceJobResponseBodyResultVideoResult) GetOutputUrl() *string {
	return s.OutputUrl
}

func (s *GetVideoCharacterReplaceJobResponseBodyResultVideoResult) SetMediaId(v string) *GetVideoCharacterReplaceJobResponseBodyResultVideoResult {
	s.MediaId = &v
	return s
}

func (s *GetVideoCharacterReplaceJobResponseBodyResultVideoResult) SetOutputUrl(v string) *GetVideoCharacterReplaceJobResponseBodyResultVideoResult {
	s.OutputUrl = &v
	return s
}

func (s *GetVideoCharacterReplaceJobResponseBodyResultVideoResult) Validate() error {
	return dara.Validate(s)
}
