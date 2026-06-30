// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikeStoryboardJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobCredit(v *GetYikeStoryboardJobResponseBodyJobCredit) *GetYikeStoryboardJobResponseBody
	GetJobCredit() *GetYikeStoryboardJobResponseBodyJobCredit
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
	// The credit consumption.
	JobCredit *GetYikeStoryboardJobResponseBodyJobCredit `json:"JobCredit,omitempty" xml:"JobCredit,omitempty" type:"Struct"`
	// The task ID. You can obtain this value from the response of [SubmitPackageJob](https://help.aliyun.com/document_detail/461964.html).
	//
	// example:
	//
	// ****a046-263c-3560-978a-fb287782****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The algorithm job parameters. This is a JSON object whose content varies depending on the algorithm.
	JobParams *GetYikeStoryboardJobResponseBodyJobParams `json:"JobParams,omitempty" xml:"JobParams,omitempty" type:"Struct"`
	// The task result.
	JobResult *GetYikeStoryboardJobResponseBodyJobResult `json:"JobResult,omitempty" xml:"JobResult,omitempty" type:"Struct"`
	// The task status. Valid values:
	//
	// - **Succeeded**: The task is successful.
	//
	// - **Failed**: The task failed.
	//
	// - **Running**: The task is running.
	//
	// example:
	//
	// Succeeded
	JobStatus *string `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// The request ID.
	//
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

func (s *GetYikeStoryboardJobResponseBody) GetJobCredit() *GetYikeStoryboardJobResponseBodyJobCredit {
	return s.JobCredit
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

func (s *GetYikeStoryboardJobResponseBody) SetJobCredit(v *GetYikeStoryboardJobResponseBodyJobCredit) *GetYikeStoryboardJobResponseBody {
	s.JobCredit = v
	return s
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
	if s.JobCredit != nil {
		if err := s.JobCredit.Validate(); err != nil {
			return err
		}
	}
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

type GetYikeStoryboardJobResponseBodyJobCredit struct {
	// The credit consumption for element image generation.
	//
	// example:
	//
	// 10.0
	ElementImageGeneration *string `json:"ElementImageGeneration,omitempty" xml:"ElementImageGeneration,omitempty"`
	// The total credit consumption.
	//
	// example:
	//
	// 200.2
	TotalCreditCost *string `json:"TotalCreditCost,omitempty" xml:"TotalCreditCost,omitempty"`
	// The credit consumption for video composition.
	//
	// example:
	//
	// 10.2
	VideoComposition *string `json:"VideoComposition,omitempty" xml:"VideoComposition,omitempty"`
	// The credit consumption for video generation.
	//
	// example:
	//
	// 180.0
	VideoGeneration *string `json:"VideoGeneration,omitempty" xml:"VideoGeneration,omitempty"`
}

func (s GetYikeStoryboardJobResponseBodyJobCredit) String() string {
	return dara.Prettify(s)
}

func (s GetYikeStoryboardJobResponseBodyJobCredit) GoString() string {
	return s.String()
}

func (s *GetYikeStoryboardJobResponseBodyJobCredit) GetElementImageGeneration() *string {
	return s.ElementImageGeneration
}

func (s *GetYikeStoryboardJobResponseBodyJobCredit) GetTotalCreditCost() *string {
	return s.TotalCreditCost
}

func (s *GetYikeStoryboardJobResponseBodyJobCredit) GetVideoComposition() *string {
	return s.VideoComposition
}

func (s *GetYikeStoryboardJobResponseBodyJobCredit) GetVideoGeneration() *string {
	return s.VideoGeneration
}

func (s *GetYikeStoryboardJobResponseBodyJobCredit) SetElementImageGeneration(v string) *GetYikeStoryboardJobResponseBodyJobCredit {
	s.ElementImageGeneration = &v
	return s
}

func (s *GetYikeStoryboardJobResponseBodyJobCredit) SetTotalCreditCost(v string) *GetYikeStoryboardJobResponseBodyJobCredit {
	s.TotalCreditCost = &v
	return s
}

func (s *GetYikeStoryboardJobResponseBodyJobCredit) SetVideoComposition(v string) *GetYikeStoryboardJobResponseBodyJobCredit {
	s.VideoComposition = &v
	return s
}

func (s *GetYikeStoryboardJobResponseBodyJobCredit) SetVideoGeneration(v string) *GetYikeStoryboardJobResponseBodyJobCredit {
	s.VideoGeneration = &v
	return s
}

func (s *GetYikeStoryboardJobResponseBodyJobCredit) Validate() error {
	return dara.Validate(s)
}

type GetYikeStoryboardJobResponseBodyJobParams struct {
	// The aspect ratio of the video.
	//
	// example:
	//
	// 16:9
	AspectRatio *string `json:"AspectRatio,omitempty" xml:"AspectRatio,omitempty"`
	// The OSS URL of the file.
	//
	// example:
	//
	// https://test.oss-cn-shanghai.aliyuncs.com/test.mp4
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	// The random seed.
	//
	// example:
	//
	// {
	//
	//   "AudioEnable": false
	//
	// }
	ModelParams *string `json:"ModelParams,omitempty" xml:"ModelParams,omitempty"`
	// The narration voice ID.
	//
	// example:
	//
	// sys_YoungGracefulWoman
	NarrationVoiceId *string `json:"NarrationVoiceId,omitempty" xml:"NarrationVoiceId,omitempty"`
	// The resolution of the generated video.
	//
	// example:
	//
	// 1K
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	// The shot prompt generation mode.
	//
	// example:
	//
	// multi
	ShotPromptMode *string `json:"ShotPromptMode,omitempty" xml:"ShotPromptMode,omitempty"`
	// The shot splitting mode.
	//
	// example:
	//
	// firstPersonNarration
	ShotSplitMode *string `json:"ShotSplitMode,omitempty" xml:"ShotSplitMode,omitempty"`
	// The source type.
	//
	// example:
	//
	// Novel
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The style ID.
	//
	// example:
	//
	// RealisticPhotography
	StyleId *string `json:"StyleId,omitempty" xml:"StyleId,omitempty"`
	// The task title. Requirements:
	//
	// - The title cannot exceed 128 bytes in length.
	//
	// - The title must be UTF-8 encoded.
	//
	// example:
	//
	// test-title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The video model.
	//
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
	// The list of abnormal storyboard IDs.
	//
	// example:
	//
	// [\\"st_2053348871\\"]
	ExceptionStoryboardIds *string `json:"ExceptionStoryboardIds,omitempty" xml:"ExceptionStoryboardIds,omitempty"`
	// The list of failed shots.
	//
	// example:
	//
	// [{\\"errorCode\\":\\"NoMediaData\\",\\"storyboardId\\":\\"st_2118280473\\",\\"shotId\\":\\"54\\"}]
	FailureShotList *string `json:"FailureShotList,omitempty" xml:"FailureShotList,omitempty"`
	// The downloadable OSS URL.
	//
	// example:
	//
	// https://test.oss-cn-shanghai.aliyuncs.com/test.mp4
	OutputUrl *string `json:"OutputUrl,omitempty" xml:"OutputUrl,omitempty"`
	// The URL of the SRT file.
	//
	// example:
	//
	// https://bucket.oss-cn-shanghai.aliyuncs.com/test/110412818/6bf24c75285142f395464d4b9c2bcf07.srt?Expires=1778220836&OSSAccessKeyId=*******&Signature=*******
	SrtFileUrl *string `json:"SrtFileUrl,omitempty" xml:"SrtFileUrl,omitempty"`
	// The detailed storyboard information for the storyboard generation task.
	//
	// example:
	//
	// [{\\"storyboardId\\":\\"st_1541525214\\",\\"title\\":\\"test_1\\",\\"status\\":\\"Produced\\",\\"subStatus\\":\\"ProduceSucc\\"},{\\"storyboardId\\":\\"st_1633435355\\",\\"title\\":\\"test_2\\",\\"status\\":\\"Produced\\",\\"subStatus\\":\\"ProduceSucc\\"}]
	StoryboardInfoList *string `json:"StoryboardInfoList,omitempty" xml:"StoryboardInfoList,omitempty"`
	// The list of successful storyboard IDs, separated by commas.
	//
	// example:
	//
	// st_2118280473, st_2118280471
	SuccessStoryboardIds *string `json:"SuccessStoryboardIds,omitempty" xml:"SuccessStoryboardIds,omitempty"`
	// The list of successful storyboards.
	//
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

func (s *GetYikeStoryboardJobResponseBodyJobResult) GetSrtFileUrl() *string {
	return s.SrtFileUrl
}

func (s *GetYikeStoryboardJobResponseBodyJobResult) GetStoryboardInfoList() *string {
	return s.StoryboardInfoList
}

func (s *GetYikeStoryboardJobResponseBodyJobResult) GetSuccessStoryboardIds() *string {
	return s.SuccessStoryboardIds
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

func (s *GetYikeStoryboardJobResponseBodyJobResult) SetSrtFileUrl(v string) *GetYikeStoryboardJobResponseBodyJobResult {
	s.SrtFileUrl = &v
	return s
}

func (s *GetYikeStoryboardJobResponseBodyJobResult) SetStoryboardInfoList(v string) *GetYikeStoryboardJobResponseBodyJobResult {
	s.StoryboardInfoList = &v
	return s
}

func (s *GetYikeStoryboardJobResponseBodyJobResult) SetSuccessStoryboardIds(v string) *GetYikeStoryboardJobResponseBodyJobResult {
	s.SuccessStoryboardIds = &v
	return s
}

func (s *GetYikeStoryboardJobResponseBodyJobResult) SetSuccessStoryboardList(v string) *GetYikeStoryboardJobResponseBodyJobResult {
	s.SuccessStoryboardList = &v
	return s
}

func (s *GetYikeStoryboardJobResponseBodyJobResult) Validate() error {
	return dara.Validate(s)
}
