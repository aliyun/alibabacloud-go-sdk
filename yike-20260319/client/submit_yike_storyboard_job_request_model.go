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
	SetKeepOriginDialogue(v bool) *SubmitYikeStoryboardJobRequest
	GetKeepOriginDialogue() *bool
	SetModelParams(v string) *SubmitYikeStoryboardJobRequest
	GetModelParams() *string
	SetNarrationVoiceId(v string) *SubmitYikeStoryboardJobRequest
	GetNarrationVoiceId() *string
	SetNeedCaption(v bool) *SubmitYikeStoryboardJobRequest
	GetNeedCaption() *bool
	SetResolution(v string) *SubmitYikeStoryboardJobRequest
	GetResolution() *string
	SetShotPromptLang(v string) *SubmitYikeStoryboardJobRequest
	GetShotPromptLang() *string
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
	// The aspect ratio of the output video. Valid values:
	//
	// - 16:9
	//
	// - 9:16
	//
	// - 4:3
	//
	// - 3:4
	//
	// example:
	//
	// 4:3
	AspectRatio *string `json:"AspectRatio,omitempty" xml:"AspectRatio,omitempty"`
	// The execution mode for storyboard generation. Valid values:
	//
	// - FullPipeline: full pipeline generation, which includes storyboard generation and shot video generation.
	//
	// - StoryboardOnly: generates only the storyboard.
	//
	// example:
	//
	// FullPipeline
	ExecMode *string `json:"ExecMode,omitempty" xml:"ExecMode,omitempty"`
	// The OSS URL of the file. The URL must point to a file with a .txt or .doc extension.
	//
	// example:
	//
	// http://test.oss-cn-shanghai.aliyuncs.com/test.txt
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	// Specifies whether to retain the original dialogue during final video composition. Default value: True.
	//
	// example:
	//
	// True
	KeepOriginDialogue *bool `json:"KeepOriginDialogue,omitempty" xml:"KeepOriginDialogue,omitempty"`
	// The model parameters in JSON format.
	//
	//  "AudioEnable": false disables audio.
	//
	// example:
	//
	// {
	//
	//   "AudioEnable": false
	//
	// }
	ModelParams *string `json:"ModelParams,omitempty" xml:"ModelParams,omitempty"`
	// The narration voice ID. Valid values:
	//
	// - sys_GracefulPoisedWoman: mature graceful female
	//
	// - sys_ElderlyWistfulWoman: wistful elderly female
	//
	// - sys_SweetBrightGirl: sweet bright girl
	//
	// - sys_YoungGracefulWoman: gentle graceful female
	//
	// - sys_MaturePoisedWoman: poised mature female
	//
	// - sys_MatureWiseWoman: elegant wise female
	//
	// - sys_CalmDeepMale: calm deep male
	//
	// - sys_SereneIntellect: serene intellectual male
	//
	// - sys_MajesticBaritone: majestic baritone male
	//
	// - sys_GravellySoulful: gravelly soulful male
	//
	// - sys_ClassicYoungMan: classic narration male
	//
	// - sys_WiseYoungMan: wise narration male
	//
	// - sys_ClassicYoungWoman: classic narration female
	//
	// - sys_IntellectualYoungWoman: intellectual narration female
	//
	// - sys_GentleYoungMan: gentle narration male
	//
	// - sys_thoughtfulBoy: thoughtful boy
	//
	// - sys_RichBassMale: rich bass male
	//
	// - sys_ClassicMiddleAgedWoman: classic middle-aged narration female
	//
	// example:
	//
	// sys_YoungGracefulWoman
	NarrationVoiceId *string `json:"NarrationVoiceId,omitempty" xml:"NarrationVoiceId,omitempty"`
	NeedCaption      *bool   `json:"NeedCaption,omitempty" xml:"NeedCaption,omitempty"`
	// The resolution of the output video. Valid values:
	//
	// - 720P
	//
	// - 1080P
	//
	// - 2K
	//
	// - 4K
	//
	// example:
	//
	// 720P
	Resolution     *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	ShotPromptLang *string `json:"ShotPromptLang,omitempty" xml:"ShotPromptLang,omitempty"`
	// The storyboard shot generation mode. Valid values:
	//
	// - multi: multi-reference video generation
	//
	// - default: image-to-video generation
	//
	// example:
	//
	// multi
	ShotPromptMode *string `json:"ShotPromptMode,omitempty" xml:"ShotPromptMode,omitempty"`
	// The shot split mode. Valid values:
	//
	// - firstPersonNarration: narration commentary mode
	//
	// example:
	//
	// firstPersonNarration
	ShotSplitMode *string `json:"ShotSplitMode,omitempty" xml:"ShotSplitMode,omitempty"`
	// Specifies whether to skip failed shots. Default value: True.
	//
	// example:
	//
	// True
	SkipFailureShot *bool `json:"SkipFailureShot,omitempty" xml:"SkipFailureShot,omitempty"`
	// The type of the material source. Valid values:
	//
	// - Novel: novel
	//
	// example:
	//
	// Novel
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The storyboard style ID. Valid values:
	//
	// - RealisticPhotographyPro: realistic photography Pro
	//
	// - RealisticGuzhuangPro: realistic ancient costume Pro
	//
	// - RealisticXianxiaPro: realistic Xianxia Pro
	//
	// - RealisticWesternPro: Western realistic Pro
	//
	// - RealisticPhotography: realistic photography
	//
	// - RealisticGuzhuang: realistic ancient costume
	//
	// - RealisticXianxia: realistic Xianxia
	//
	// - RealisticWasteland: realistic wasteland
	//
	// - RealisticEra: realistic vintage
	//
	// - GuofengAnime: 2D Chinese-style anime
	//
	// - GuofengAnime3D: 3D Chinese-style anime
	//
	// - AncientRomanceAnime: anime ancient romance
	//
	// - PostApocalypticAnime: anime post-apocalyptic
	//
	// - Cartoon3D: 3D cartoon
	//
	// - Photorealistic3D: photorealistic 3D rendering
	//
	// - SciFiRealism: sci-fi realism
	//
	// - Chibi3D: 3D chibi
	//
	// - ShojoManga: Japanese manga
	//
	// - NewPeriodAnime: new era Japanese anime
	//
	// - FairyTale2D: 2D fairy tale
	//
	// - Wasteland2D: 2D wasteland
	//
	// - InkWuxia: ink wash Wuxia
	//
	// - ShadiaoMeme: panda head meme style
	//
	// - Chibi2D: 2D chibi
	//
	// - Ghibli: Ghibli
	//
	// - SciFiComic: cyberpunk
	//
	// - AmericanSuperhero: American superhero
	//
	// example:
	//
	// RealisticPhotography
	StyleId *string `json:"StyleId,omitempty" xml:"StyleId,omitempty"`
	// The task title. If not specified, a default title is automatically generated based on the date. The title cannot exceed 128 bytes in length and must be UTF-8 encoded.
	//
	// example:
	//
	// test-title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The custom settings in JSON format. Fields include:
	//
	// - NotifyAddress: the callback URL for task completion. MNS callbacks and HTTP callbacks are supported.
	//
	// example:
	//
	// {
	//
	//   "NotifyAddress": "https://www.callback.com"
	//
	// }
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The video model. Valid values:
	//
	// - wan2.6-r2v-flash
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

func (s *SubmitYikeStoryboardJobRequest) GetKeepOriginDialogue() *bool {
	return s.KeepOriginDialogue
}

func (s *SubmitYikeStoryboardJobRequest) GetModelParams() *string {
	return s.ModelParams
}

func (s *SubmitYikeStoryboardJobRequest) GetNarrationVoiceId() *string {
	return s.NarrationVoiceId
}

func (s *SubmitYikeStoryboardJobRequest) GetNeedCaption() *bool {
	return s.NeedCaption
}

func (s *SubmitYikeStoryboardJobRequest) GetResolution() *string {
	return s.Resolution
}

func (s *SubmitYikeStoryboardJobRequest) GetShotPromptLang() *string {
	return s.ShotPromptLang
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

func (s *SubmitYikeStoryboardJobRequest) SetKeepOriginDialogue(v bool) *SubmitYikeStoryboardJobRequest {
	s.KeepOriginDialogue = &v
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

func (s *SubmitYikeStoryboardJobRequest) SetNeedCaption(v bool) *SubmitYikeStoryboardJobRequest {
	s.NeedCaption = &v
	return s
}

func (s *SubmitYikeStoryboardJobRequest) SetResolution(v string) *SubmitYikeStoryboardJobRequest {
	s.Resolution = &v
	return s
}

func (s *SubmitYikeStoryboardJobRequest) SetShotPromptLang(v string) *SubmitYikeStoryboardJobRequest {
	s.ShotPromptLang = &v
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
