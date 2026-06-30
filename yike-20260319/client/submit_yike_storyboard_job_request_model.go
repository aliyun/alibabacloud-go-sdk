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
	// The aspect ratio of the output video. Valid values:
	//
	// - 16:9
	//
	// - 9:16
	//
	// - 4:3
	//
	// - 3:4.
	//
	// example:
	//
	// 4:3
	AspectRatio *string `json:"AspectRatio,omitempty" xml:"AspectRatio,omitempty"`
	// The execution mode for storyboard generation. Valid values:
	//
	// - FullPipeline: Full pipeline generation, which includes both storyboard generation and shot video generation.
	//
	// - StoryboardOnly: Generates only the storyboard.
	//
	// example:
	//
	// FullPipeline
	ExecMode *string `json:"ExecMode,omitempty" xml:"ExecMode,omitempty"`
	// The OSS URL of the file. Only URLs with the .txt or .doc file name extension are supported.
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
	// "AudioEnable": false disables audio.
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
	// - sys_GracefulPoisedWoman: Graceful Poised Woman
	//
	// - sys_ElderlyWistfulWoman: Elderly Wistful Woman
	//
	// - sys_SweetBrightGirl: Sweet Bright Girl
	//
	// - sys_YoungGracefulWoman: Young Graceful Woman
	//
	// - sys_MaturePoisedWoman: Mature Poised Woman
	//
	// - sys_MatureWiseWoman: Mature Wise Woman
	//
	// - sys_CalmDeepMale: Calm Deep Male
	//
	// - sys_SereneIntellect: Serene Intellect
	//
	// - sys_MajesticBaritone: Majestic Baritone
	//
	// - sys_GravellySoulful: Gravelly Soulful
	//
	// - sys_ClassicYoungMan: Classic Young Man Narrator
	//
	// - sys_WiseYoungMan: Wise Young Man Narrator
	//
	// - sys_ClassicYoungWoman: Classic Young Woman Narrator
	//
	// - sys_IntellectualYoungWoman: Intellectual Young Woman Narrator
	//
	// - sys_GentleYoungMan: Gentle Young Man Narrator
	//
	// - sys_thoughtfulBoy: Thoughtful Boy
	//
	// - sys_RichBassMale: Rich Bass Male
	//
	// - sys_ClassicMiddleAgedWoman: Classic Middle-Aged Woman Narrator.
	//
	// example:
	//
	// sys_YoungGracefulWoman
	NarrationVoiceId *string `json:"NarrationVoiceId,omitempty" xml:"NarrationVoiceId,omitempty"`
	// The resolution of the output video. Valid values:
	//
	// - 720P
	//
	// - 1080P
	//
	// - 2K
	//
	// - 4K.
	//
	// example:
	//
	// 720P
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	// The storyboard shot generation mode. Valid values:
	//
	// - multi: multi-reference video generation
	//
	// - default: image-to-video generation.
	//
	// example:
	//
	// multi
	ShotPromptMode *string `json:"ShotPromptMode,omitempty" xml:"ShotPromptMode,omitempty"`
	// The shot split mode. Valid values:
	//
	// - firstPersonNarration: narration mode.
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
	// - Novel: novel.
	//
	// example:
	//
	// Novel
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The storyboard style ID. Valid values:
	//
	// - RealisticPhotographyPro: Realistic Photography Pro
	//
	// - RealisticGuzhuangPro: Realistic Ancient Costume Pro
	//
	// - RealisticXianxiaPro: Realistic Xianxia Pro
	//
	// - RealisticWesternPro: Western Realism Pro
	//
	// - RealisticPhotography: Realistic Photography
	//
	// - RealisticGuzhuang: Realistic Ancient Costume
	//
	// - RealisticXianxia: Realistic Xianxia
	//
	// - RealisticWasteland: Realistic Wasteland
	//
	// - RealisticEra: Realistic Vintage
	//
	// - GuofengAnime: 2D Chinese-style Anime
	//
	// - GuofengAnime3D: 3D Chinese-style Anime
	//
	// - AncientRomanceAnime: Anime Ancient Romance
	//
	// - PostApocalypticAnime: Anime Post-Apocalyptic
	//
	// - Cartoon3D: 3D Cartoon
	//
	// - Photorealistic3D: Photorealistic 3D Rendering
	//
	// - SciFiRealism: Sci-Fi Realism
	//
	// - Chibi3D: 3D Chibi
	//
	// - ShojoManga: Japanese Manga
	//
	// - NewPeriodAnime: New Era Japanese Anime
	//
	// - FairyTale2D: 2D Fairy Tale
	//
	// - Wasteland2D: 2D Wasteland
	//
	// - InkWuxia: Ink Wash Wuxia
	//
	// - ShadiaoMeme: Panda Head Meme
	//
	// - Chibi2D: 2D Chibi
	//
	// - Ghibli: Ghibli
	//
	// - SciFiComic: Cyberpunk
	//
	// - AmericanSuperhero: American Superhero.
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
	// The custom settings in JSON format.
	//
	// - NotifyAddress specifies the callback for task completion. Both MNS callbacks and HTTP callbacks are supported.
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
	// - wan2.6-r2v-flash.
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
