// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitYikeStoryboardJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAspectRatio(v string) *SubmitYikeStoryboardJobRequest
	GetAspectRatio() *string
	SetExecMode(v string) *SubmitYikeStoryboardJobRequest
	GetExecMode() *string
	SetFileURL(v string) *SubmitYikeStoryboardJobRequest
	GetFileURL() *string
	SetModelParams(v string) *SubmitYikeStoryboardJobRequest
	GetModelParams() *string
	SetNarrationVoiceId(v string) *SubmitYikeStoryboardJobRequest
	GetNarrationVoiceId() *string
	SetResolution(v string) *SubmitYikeStoryboardJobRequest
	GetResolution() *string
	SetShotPromptMode(v string) *SubmitYikeStoryboardJobRequest
	GetShotPromptMode() *string
	SetShotSplitMode(v string) *SubmitYikeStoryboardJobRequest
	GetShotSplitMode() *string
	SetSkipFailureShot(v bool) *SubmitYikeStoryboardJobRequest
	GetSkipFailureShot() *bool
	SetSourceType(v string) *SubmitYikeStoryboardJobRequest
	GetSourceType() *string
	SetStyleId(v string) *SubmitYikeStoryboardJobRequest
	GetStyleId() *string
	SetTitle(v string) *SubmitYikeStoryboardJobRequest
	GetTitle() *string
	SetUserData(v string) *SubmitYikeStoryboardJobRequest
	GetUserData() *string
	SetVideoModel(v string) *SubmitYikeStoryboardJobRequest
	GetVideoModel() *string
}

type SubmitYikeStoryboardJobRequest struct {
	// The aspect ratio of the output video.
	//
	// example:
	//
	// 4:3
	AspectRatio *string `json:"AspectRatio,omitempty" xml:"AspectRatio,omitempty"`
	// The storyboard generation execution mode.
	//
	// - `FullPipeline`: Executes the full generation pipeline, including both storyboard creation and shot video generation.
	//
	// - `StoryboardOnly`: Generates only the storyboard.
	//
	// example:
	//
	// FullPipeline
	ExecMode *string `json:"ExecMode,omitempty" xml:"ExecMode,omitempty"`
	// The OSS address of the file.
	//
	// example:
	//
	// http://test.oss-cn-shanghai.aliyuncs.com/test.mp4
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	// Parameters for the model, in JSON format.
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
	// The resolution of the output video.
	//
	// example:
	//
	// 720P
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	// The storyboard shot generation mode.
	//
	// example:
	//
	// multi
	ShotPromptMode *string `json:"ShotPromptMode,omitempty" xml:"ShotPromptMode,omitempty"`
	// The shot split mode.
	//
	// example:
	//
	// firstPersonNarration
	ShotSplitMode *string `json:"ShotSplitMode,omitempty" xml:"ShotSplitMode,omitempty"`
	// Specifies whether to skip a failed shot. The default value is `true`.
	//
	// example:
	//
	// True
	SkipFailureShot *bool `json:"SkipFailureShot,omitempty" xml:"SkipFailureShot,omitempty"`
	// The source type.
	//
	// example:
	//
	// Novel
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The storyboard style ID.
	//
	// example:
	//
	// RealisticPhotography
	StyleId *string `json:"StyleId,omitempty" xml:"StyleId,omitempty"`
	// The job title. It must be a UTF-8 encoded string of up to 128 bytes. If you do not specify a title, the system generates a default one based on the date.
	//
	// example:
	//
	// test-title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// Custom settings in JSON format. This parameter can contain the following field:
	//
	// - The `NotifyAddress` field specifies the callback URL that is invoked when the job is complete. Both MNS and HTTP callbacks are supported.
	//
	// example:
	//
	// {
	//
	//   "NotifyAddress": "https://www.callback.com"
	//
	// }
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The video model.
	//
	// example:
	//
	// wan2.6-r2v-flash
	VideoModel *string `json:"VideoModel,omitempty" xml:"VideoModel,omitempty"`
}

func (s SubmitYikeStoryboardJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitYikeStoryboardJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitYikeStoryboardJobRequest) GetAspectRatio() *string {
	return s.AspectRatio
}

func (s *SubmitYikeStoryboardJobRequest) GetExecMode() *string {
	return s.ExecMode
}

func (s *SubmitYikeStoryboardJobRequest) GetFileURL() *string {
	return s.FileURL
}

func (s *SubmitYikeStoryboardJobRequest) GetModelParams() *string {
	return s.ModelParams
}

func (s *SubmitYikeStoryboardJobRequest) GetNarrationVoiceId() *string {
	return s.NarrationVoiceId
}

func (s *SubmitYikeStoryboardJobRequest) GetResolution() *string {
	return s.Resolution
}

func (s *SubmitYikeStoryboardJobRequest) GetShotPromptMode() *string {
	return s.ShotPromptMode
}

func (s *SubmitYikeStoryboardJobRequest) GetShotSplitMode() *string {
	return s.ShotSplitMode
}

func (s *SubmitYikeStoryboardJobRequest) GetSkipFailureShot() *bool {
	return s.SkipFailureShot
}

func (s *SubmitYikeStoryboardJobRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *SubmitYikeStoryboardJobRequest) GetStyleId() *string {
	return s.StyleId
}

func (s *SubmitYikeStoryboardJobRequest) GetTitle() *string {
	return s.Title
}

func (s *SubmitYikeStoryboardJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitYikeStoryboardJobRequest) GetVideoModel() *string {
	return s.VideoModel
}

func (s *SubmitYikeStoryboardJobRequest) SetAspectRatio(v string) *SubmitYikeStoryboardJobRequest {
	s.AspectRatio = &v
	return s
}

func (s *SubmitYikeStoryboardJobRequest) SetExecMode(v string) *SubmitYikeStoryboardJobRequest {
	s.ExecMode = &v
	return s
}

func (s *SubmitYikeStoryboardJobRequest) SetFileURL(v string) *SubmitYikeStoryboardJobRequest {
	s.FileURL = &v
	return s
}

func (s *SubmitYikeStoryboardJobRequest) SetModelParams(v string) *SubmitYikeStoryboardJobRequest {
	s.ModelParams = &v
	return s
}

func (s *SubmitYikeStoryboardJobRequest) SetNarrationVoiceId(v string) *SubmitYikeStoryboardJobRequest {
	s.NarrationVoiceId = &v
	return s
}

func (s *SubmitYikeStoryboardJobRequest) SetResolution(v string) *SubmitYikeStoryboardJobRequest {
	s.Resolution = &v
	return s
}

func (s *SubmitYikeStoryboardJobRequest) SetShotPromptMode(v string) *SubmitYikeStoryboardJobRequest {
	s.ShotPromptMode = &v
	return s
}

func (s *SubmitYikeStoryboardJobRequest) SetShotSplitMode(v string) *SubmitYikeStoryboardJobRequest {
	s.ShotSplitMode = &v
	return s
}

func (s *SubmitYikeStoryboardJobRequest) SetSkipFailureShot(v bool) *SubmitYikeStoryboardJobRequest {
	s.SkipFailureShot = &v
	return s
}

func (s *SubmitYikeStoryboardJobRequest) SetSourceType(v string) *SubmitYikeStoryboardJobRequest {
	s.SourceType = &v
	return s
}

func (s *SubmitYikeStoryboardJobRequest) SetStyleId(v string) *SubmitYikeStoryboardJobRequest {
	s.StyleId = &v
	return s
}

func (s *SubmitYikeStoryboardJobRequest) SetTitle(v string) *SubmitYikeStoryboardJobRequest {
	s.Title = &v
	return s
}

func (s *SubmitYikeStoryboardJobRequest) SetUserData(v string) *SubmitYikeStoryboardJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitYikeStoryboardJobRequest) SetVideoModel(v string) *SubmitYikeStoryboardJobRequest {
	s.VideoModel = &v
	return s
}

func (s *SubmitYikeStoryboardJobRequest) Validate() error {
	return dara.Validate(s)
}
