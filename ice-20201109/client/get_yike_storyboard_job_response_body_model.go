// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikeStoryboardJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetYikeStoryboardJobResponseBody
	GetJobId() *string
	SetJobParams(v *GetYikeStoryboardJobResponseBodyJobParams) *GetYikeStoryboardJobResponseBody
	GetJobParams() *GetYikeStoryboardJobResponseBodyJobParams
	SetJobResult(v *GetYikeStoryboardJobResponseBodyJobResult) *GetYikeStoryboardJobResponseBody
	GetJobResult() *GetYikeStoryboardJobResponseBodyJobResult
	SetJobStatus(v string) *GetYikeStoryboardJobResponseBody
	GetJobStatus() *string
	SetRequestId(v string) *GetYikeStoryboardJobResponseBody
	GetRequestId() *string
}

type GetYikeStoryboardJobResponseBody struct {
	// example:
	//
	// ****a046-263c-3560-978a-fb287782****
	JobId     *string                                    `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobParams *GetYikeStoryboardJobResponseBodyJobParams `json:"JobParams,omitempty" xml:"JobParams,omitempty" type:"Struct"`
	JobResult *GetYikeStoryboardJobResponseBodyJobResult `json:"JobResult,omitempty" xml:"JobResult,omitempty" type:"Struct"`
	// example:
	//
	// Succeeded
	JobStatus *string `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetYikeStoryboardJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetYikeStoryboardJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetYikeStoryboardJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *GetYikeStoryboardJobResponseBody) GetJobParams() *GetYikeStoryboardJobResponseBodyJobParams {
	return s.JobParams
}

func (s *GetYikeStoryboardJobResponseBody) GetJobResult() *GetYikeStoryboardJobResponseBodyJobResult {
	return s.JobResult
}

func (s *GetYikeStoryboardJobResponseBody) GetJobStatus() *string {
	return s.JobStatus
}

func (s *GetYikeStoryboardJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetYikeStoryboardJobResponseBody) SetJobId(v string) *GetYikeStoryboardJobResponseBody {
	s.JobId = &v
	return s
}

func (s *GetYikeStoryboardJobResponseBody) SetJobParams(v *GetYikeStoryboardJobResponseBodyJobParams) *GetYikeStoryboardJobResponseBody {
	s.JobParams = v
	return s
}

func (s *GetYikeStoryboardJobResponseBody) SetJobResult(v *GetYikeStoryboardJobResponseBodyJobResult) *GetYikeStoryboardJobResponseBody {
	s.JobResult = v
	return s
}

func (s *GetYikeStoryboardJobResponseBody) SetJobStatus(v string) *GetYikeStoryboardJobResponseBody {
	s.JobStatus = &v
	return s
}

func (s *GetYikeStoryboardJobResponseBody) SetRequestId(v string) *GetYikeStoryboardJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetYikeStoryboardJobResponseBody) Validate() error {
	if s.JobParams != nil {
		if err := s.JobParams.Validate(); err != nil {
			return err
		}
	}
	if s.JobResult != nil {
		if err := s.JobResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetYikeStoryboardJobResponseBodyJobParams struct {
	// example:
	//
	// 16:9
	AspectRatio *string `json:"AspectRatio,omitempty" xml:"AspectRatio,omitempty"`
	// example:
	//
	// https://test.oss-cn-shanghai.aliyuncs.com/test.mp4
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	// example:
	//
	// {
	//
	//   "AudioEnable": false
	//
	// }
	ModelParams *string `json:"ModelParams,omitempty" xml:"ModelParams,omitempty"`
	// example:
	//
	// sys_YoungGracefulWoman
	NarrationVoiceId *string `json:"NarrationVoiceId,omitempty" xml:"NarrationVoiceId,omitempty"`
	// example:
	//
	// 1K
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	// example:
	//
	// multi
	ShotPromptMode *string `json:"ShotPromptMode,omitempty" xml:"ShotPromptMode,omitempty"`
	// example:
	//
	// firstPersonNarration
	ShotSplitMode *string `json:"ShotSplitMode,omitempty" xml:"ShotSplitMode,omitempty"`
	// example:
	//
	// Novel
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// RealisticPhotography
	StyleId *string `json:"StyleId,omitempty" xml:"StyleId,omitempty"`
	// example:
	//
	// test-title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// wan2.6-r2v-flash
	VideoModel *string `json:"VideoModel,omitempty" xml:"VideoModel,omitempty"`
}

func (s GetYikeStoryboardJobResponseBodyJobParams) String() string {
	return dara.Prettify(s)
}

func (s GetYikeStoryboardJobResponseBodyJobParams) GoString() string {
	return s.String()
}

func (s *GetYikeStoryboardJobResponseBodyJobParams) GetAspectRatio() *string {
	return s.AspectRatio
}

func (s *GetYikeStoryboardJobResponseBodyJobParams) GetFileURL() *string {
	return s.FileURL
}

func (s *GetYikeStoryboardJobResponseBodyJobParams) GetModelParams() *string {
	return s.ModelParams
}

func (s *GetYikeStoryboardJobResponseBodyJobParams) GetNarrationVoiceId() *string {
	return s.NarrationVoiceId
}

func (s *GetYikeStoryboardJobResponseBodyJobParams) GetResolution() *string {
	return s.Resolution
}

func (s *GetYikeStoryboardJobResponseBodyJobParams) GetShotPromptMode() *string {
	return s.ShotPromptMode
}

func (s *GetYikeStoryboardJobResponseBodyJobParams) GetShotSplitMode() *string {
	return s.ShotSplitMode
}

func (s *GetYikeStoryboardJobResponseBodyJobParams) GetSourceType() *string {
	return s.SourceType
}

func (s *GetYikeStoryboardJobResponseBodyJobParams) GetStyleId() *string {
	return s.StyleId
}

func (s *GetYikeStoryboardJobResponseBodyJobParams) GetTitle() *string {
	return s.Title
}

func (s *GetYikeStoryboardJobResponseBodyJobParams) GetVideoModel() *string {
	return s.VideoModel
}

func (s *GetYikeStoryboardJobResponseBodyJobParams) SetAspectRatio(v string) *GetYikeStoryboardJobResponseBodyJobParams {
	s.AspectRatio = &v
	return s
}

func (s *GetYikeStoryboardJobResponseBodyJobParams) SetFileURL(v string) *GetYikeStoryboardJobResponseBodyJobParams {
	s.FileURL = &v
	return s
}

func (s *GetYikeStoryboardJobResponseBodyJobParams) SetModelParams(v string) *GetYikeStoryboardJobResponseBodyJobParams {
	s.ModelParams = &v
	return s
}

func (s *GetYikeStoryboardJobResponseBodyJobParams) SetNarrationVoiceId(v string) *GetYikeStoryboardJobResponseBodyJobParams {
	s.NarrationVoiceId = &v
	return s
}

func (s *GetYikeStoryboardJobResponseBodyJobParams) SetResolution(v string) *GetYikeStoryboardJobResponseBodyJobParams {
	s.Resolution = &v
	return s
}

func (s *GetYikeStoryboardJobResponseBodyJobParams) SetShotPromptMode(v string) *GetYikeStoryboardJobResponseBodyJobParams {
	s.ShotPromptMode = &v
	return s
}

func (s *GetYikeStoryboardJobResponseBodyJobParams) SetShotSplitMode(v string) *GetYikeStoryboardJobResponseBodyJobParams {
	s.ShotSplitMode = &v
	return s
}

func (s *GetYikeStoryboardJobResponseBodyJobParams) SetSourceType(v string) *GetYikeStoryboardJobResponseBodyJobParams {
	s.SourceType = &v
	return s
}

func (s *GetYikeStoryboardJobResponseBodyJobParams) SetStyleId(v string) *GetYikeStoryboardJobResponseBodyJobParams {
	s.StyleId = &v
	return s
}

func (s *GetYikeStoryboardJobResponseBodyJobParams) SetTitle(v string) *GetYikeStoryboardJobResponseBodyJobParams {
	s.Title = &v
	return s
}

func (s *GetYikeStoryboardJobResponseBodyJobParams) SetVideoModel(v string) *GetYikeStoryboardJobResponseBodyJobParams {
	s.VideoModel = &v
	return s
}

func (s *GetYikeStoryboardJobResponseBodyJobParams) Validate() error {
	return dara.Validate(s)
}

type GetYikeStoryboardJobResponseBodyJobResult struct {
	// example:
	//
	// [\\"st_2053348871\\"]
	ExceptionStoryboardIds *string `json:"ExceptionStoryboardIds,omitempty" xml:"ExceptionStoryboardIds,omitempty"`
	// example:
	//
	// [{\\"errorCode\\":\\"NoMediaData\\",\\"storyboardId\\":\\"st_2118280473\\",\\"shotId\\":\\"54\\"}]
	FailureShotList *string `json:"FailureShotList,omitempty" xml:"FailureShotList,omitempty"`
	// example:
	//
	// https://test.oss-cn-shanghai.aliyuncs.com/test.mp4
	OutputUrl *string `json:"OutputUrl,omitempty" xml:"OutputUrl,omitempty"`
	// example:
	//
	// [\\"st_2118280473\\"]
	SuccessStoryboardList *string `json:"SuccessStoryboardList,omitempty" xml:"SuccessStoryboardList,omitempty"`
}

func (s GetYikeStoryboardJobResponseBodyJobResult) String() string {
	return dara.Prettify(s)
}

func (s GetYikeStoryboardJobResponseBodyJobResult) GoString() string {
	return s.String()
}

func (s *GetYikeStoryboardJobResponseBodyJobResult) GetExceptionStoryboardIds() *string {
	return s.ExceptionStoryboardIds
}

func (s *GetYikeStoryboardJobResponseBodyJobResult) GetFailureShotList() *string {
	return s.FailureShotList
}

func (s *GetYikeStoryboardJobResponseBodyJobResult) GetOutputUrl() *string {
	return s.OutputUrl
}

func (s *GetYikeStoryboardJobResponseBodyJobResult) GetSuccessStoryboardList() *string {
	return s.SuccessStoryboardList
}

func (s *GetYikeStoryboardJobResponseBodyJobResult) SetExceptionStoryboardIds(v string) *GetYikeStoryboardJobResponseBodyJobResult {
	s.ExceptionStoryboardIds = &v
	return s
}

func (s *GetYikeStoryboardJobResponseBodyJobResult) SetFailureShotList(v string) *GetYikeStoryboardJobResponseBodyJobResult {
	s.FailureShotList = &v
	return s
}

func (s *GetYikeStoryboardJobResponseBodyJobResult) SetOutputUrl(v string) *GetYikeStoryboardJobResponseBodyJobResult {
	s.OutputUrl = &v
	return s
}

func (s *GetYikeStoryboardJobResponseBodyJobResult) SetSuccessStoryboardList(v string) *GetYikeStoryboardJobResponseBodyJobResult {
	s.SuccessStoryboardList = &v
	return s
}

func (s *GetYikeStoryboardJobResponseBodyJobResult) Validate() error {
	return dara.Validate(s)
}
