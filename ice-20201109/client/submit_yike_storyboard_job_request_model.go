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
	// example:
	//
	// 4:3
	AspectRatio *string `json:"AspectRatio,omitempty" xml:"AspectRatio,omitempty"`
	// example:
	//
	// http://test.oss-cn-shanghai.aliyuncs.com/test.mp4
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
	// 720P
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
	// {
	//
	//   "NotifyAddress": "https://www.callback.com"
	//
	// }
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
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
