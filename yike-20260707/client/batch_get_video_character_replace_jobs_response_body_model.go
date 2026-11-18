// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetVideoCharacterReplaceJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobList(v []*BatchGetVideoCharacterReplaceJobsResponseBodyJobList) *BatchGetVideoCharacterReplaceJobsResponseBody
	GetJobList() []*BatchGetVideoCharacterReplaceJobsResponseBodyJobList
	SetRequestId(v string) *BatchGetVideoCharacterReplaceJobsResponseBody
	GetRequestId() *string
}

type BatchGetVideoCharacterReplaceJobsResponseBody struct {
	// The list of queried jobs.
	JobList []*BatchGetVideoCharacterReplaceJobsResponseBodyJobList `json:"JobList,omitempty" xml:"JobList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchGetVideoCharacterReplaceJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchGetVideoCharacterReplaceJobsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBody) GetJobList() []*BatchGetVideoCharacterReplaceJobsResponseBodyJobList {
	return s.JobList
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBody) SetJobList(v []*BatchGetVideoCharacterReplaceJobsResponseBodyJobList) *BatchGetVideoCharacterReplaceJobsResponseBody {
	s.JobList = v
	return s
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBody) SetRequestId(v string) *BatchGetVideoCharacterReplaceJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBody) Validate() error {
	if s.JobList != nil {
		for _, item := range s.JobList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchGetVideoCharacterReplaceJobsResponseBodyJobList struct {
	// The ID of the video character replacement application. Make sure to use a configured application.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application input parameters specified when the job was submitted. The value is a JSON-serialized string. This corresponds to the AppParams parameter in the submit operation.
	//
	// example:
	//
	// {"LoadImage.1.TargetImage":"1a7852b032c371f0b64fe6f6c74b****"}
	AppInputConfig *string `json:"AppInputConfig,omitempty" xml:"AppInputConfig,omitempty"`
	// The time when the job finished execution. The time is in the ISO 8601 format with a time zone.
	//
	// example:
	//
	// 2026-02-06T18:54:18.809+08:00
	ExecutionFinishTime *string `json:"ExecutionFinishTime,omitempty" xml:"ExecutionFinishTime,omitempty"`
	// The time when the job started execution. The time is in the ISO 8601 format with a time zone.
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
	// The job ID, which is returned by the SubmitVideoCharacterReplaceJob operation.
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
	// The output results of the job, categorized by image, video, and audio.
	Result *BatchGetVideoCharacterReplaceJobsResponseBodyJobListResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The job status. Valid values:
	//
	// - Created: The job is created.
	//
	// - Queuing: The job is in the queue.
	//
	// - Executing: The job is being executed.
	//
	// - Finished: The job is finished.
	//
	// - Failed: The job failed.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s BatchGetVideoCharacterReplaceJobsResponseBodyJobList) String() string {
	return dara.Prettify(s)
}

func (s BatchGetVideoCharacterReplaceJobsResponseBodyJobList) GoString() string {
	return s.String()
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobList) GetAppId() *string {
	return s.AppId
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobList) GetAppInputConfig() *string {
	return s.AppInputConfig
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobList) GetExecutionFinishTime() *string {
	return s.ExecutionFinishTime
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobList) GetExecutionStartTime() *string {
	return s.ExecutionStartTime
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobList) GetFolderId() *string {
	return s.FolderId
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobList) GetJobId() *string {
	return s.JobId
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobList) GetProjectId() *string {
	return s.ProjectId
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobList) GetResult() *BatchGetVideoCharacterReplaceJobsResponseBodyJobListResult {
	return s.Result
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobList) GetStatus() *string {
	return s.Status
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobList) SetAppId(v string) *BatchGetVideoCharacterReplaceJobsResponseBodyJobList {
	s.AppId = &v
	return s
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobList) SetAppInputConfig(v string) *BatchGetVideoCharacterReplaceJobsResponseBodyJobList {
	s.AppInputConfig = &v
	return s
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobList) SetExecutionFinishTime(v string) *BatchGetVideoCharacterReplaceJobsResponseBodyJobList {
	s.ExecutionFinishTime = &v
	return s
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobList) SetExecutionStartTime(v string) *BatchGetVideoCharacterReplaceJobsResponseBodyJobList {
	s.ExecutionStartTime = &v
	return s
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobList) SetFolderId(v string) *BatchGetVideoCharacterReplaceJobsResponseBodyJobList {
	s.FolderId = &v
	return s
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobList) SetJobId(v string) *BatchGetVideoCharacterReplaceJobsResponseBodyJobList {
	s.JobId = &v
	return s
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobList) SetProjectId(v string) *BatchGetVideoCharacterReplaceJobsResponseBodyJobList {
	s.ProjectId = &v
	return s
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobList) SetResult(v *BatchGetVideoCharacterReplaceJobsResponseBodyJobListResult) *BatchGetVideoCharacterReplaceJobsResponseBodyJobList {
	s.Result = v
	return s
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobList) SetStatus(v string) *BatchGetVideoCharacterReplaceJobsResponseBodyJobList {
	s.Status = &v
	return s
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobList) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchGetVideoCharacterReplaceJobsResponseBodyJobListResult struct {
	// The list of output audio files.
	AudioResult []*BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultAudioResult `json:"AudioResult,omitempty" xml:"AudioResult,omitempty" type:"Repeated"`
	// The list of output images.
	ImageResult []*BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultImageResult `json:"ImageResult,omitempty" xml:"ImageResult,omitempty" type:"Repeated"`
	// The list of output videos.
	VideoResult []*BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultVideoResult `json:"VideoResult,omitempty" xml:"VideoResult,omitempty" type:"Repeated"`
}

func (s BatchGetVideoCharacterReplaceJobsResponseBodyJobListResult) String() string {
	return dara.Prettify(s)
}

func (s BatchGetVideoCharacterReplaceJobsResponseBodyJobListResult) GoString() string {
	return s.String()
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobListResult) GetAudioResult() []*BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultAudioResult {
	return s.AudioResult
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobListResult) GetImageResult() []*BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultImageResult {
	return s.ImageResult
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobListResult) GetVideoResult() []*BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultVideoResult {
	return s.VideoResult
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobListResult) SetAudioResult(v []*BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultAudioResult) *BatchGetVideoCharacterReplaceJobsResponseBodyJobListResult {
	s.AudioResult = v
	return s
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobListResult) SetImageResult(v []*BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultImageResult) *BatchGetVideoCharacterReplaceJobsResponseBodyJobListResult {
	s.ImageResult = v
	return s
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobListResult) SetVideoResult(v []*BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultVideoResult) *BatchGetVideoCharacterReplaceJobsResponseBodyJobListResult {
	s.VideoResult = v
	return s
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobListResult) Validate() error {
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

type BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultAudioResult struct {
	// The ID of the output media asset.
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

func (s BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultAudioResult) String() string {
	return dara.Prettify(s)
}

func (s BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultAudioResult) GoString() string {
	return s.String()
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultAudioResult) GetMediaId() *string {
	return s.MediaId
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultAudioResult) GetOutputUrl() *string {
	return s.OutputUrl
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultAudioResult) SetMediaId(v string) *BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultAudioResult {
	s.MediaId = &v
	return s
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultAudioResult) SetOutputUrl(v string) *BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultAudioResult {
	s.OutputUrl = &v
	return s
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultAudioResult) Validate() error {
	return dara.Validate(s)
}

type BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultImageResult struct {
	// The ID of the output media asset.
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

func (s BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultImageResult) String() string {
	return dara.Prettify(s)
}

func (s BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultImageResult) GoString() string {
	return s.String()
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultImageResult) GetMediaId() *string {
	return s.MediaId
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultImageResult) GetOutputUrl() *string {
	return s.OutputUrl
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultImageResult) SetMediaId(v string) *BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultImageResult {
	s.MediaId = &v
	return s
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultImageResult) SetOutputUrl(v string) *BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultImageResult {
	s.OutputUrl = &v
	return s
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultImageResult) Validate() error {
	return dara.Validate(s)
}

type BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultVideoResult struct {
	// The ID of the output media asset.
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

func (s BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultVideoResult) String() string {
	return dara.Prettify(s)
}

func (s BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultVideoResult) GoString() string {
	return s.String()
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultVideoResult) GetMediaId() *string {
	return s.MediaId
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultVideoResult) GetOutputUrl() *string {
	return s.OutputUrl
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultVideoResult) SetMediaId(v string) *BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultVideoResult {
	s.MediaId = &v
	return s
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultVideoResult) SetOutputUrl(v string) *BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultVideoResult {
	s.OutputUrl = &v
	return s
}

func (s *BatchGetVideoCharacterReplaceJobsResponseBodyJobListResultVideoResult) Validate() error {
	return dara.Validate(s)
}
