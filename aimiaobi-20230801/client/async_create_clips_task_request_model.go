// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncCreateClipsTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloseMusic(v bool) *AsyncCreateClipsTaskRequest
	GetCloseMusic() *bool
	SetCloseSubtitle(v bool) *AsyncCreateClipsTaskRequest
	GetCloseSubtitle() *bool
	SetCloseVoice(v bool) *AsyncCreateClipsTaskRequest
	GetCloseVoice() *bool
	SetClosingCreditsUrl(v string) *AsyncCreateClipsTaskRequest
	GetClosingCreditsUrl() *string
	SetColorWords(v []*AsyncCreateClipsTaskRequestColorWords) *AsyncCreateClipsTaskRequest
	GetColorWords() []*AsyncCreateClipsTaskRequestColorWords
	SetCosyVoiceAppKey(v string) *AsyncCreateClipsTaskRequest
	GetCosyVoiceAppKey() *string
	SetCosyVoiceToken(v string) *AsyncCreateClipsTaskRequest
	GetCosyVoiceToken() *string
	SetCustomVoiceStyle(v string) *AsyncCreateClipsTaskRequest
	GetCustomVoiceStyle() *string
	SetCustomVoiceUrl(v string) *AsyncCreateClipsTaskRequest
	GetCustomVoiceUrl() *string
	SetCustomVoiceVolume(v int32) *AsyncCreateClipsTaskRequest
	GetCustomVoiceVolume() *int32
	SetHeight(v int32) *AsyncCreateClipsTaskRequest
	GetHeight() *int32
	SetHighDefSourceVideos(v []*AsyncCreateClipsTaskRequestHighDefSourceVideos) *AsyncCreateClipsTaskRequest
	GetHighDefSourceVideos() []*AsyncCreateClipsTaskRequestHighDefSourceVideos
	SetMusicStyle(v string) *AsyncCreateClipsTaskRequest
	GetMusicStyle() *string
	SetMusicUrl(v string) *AsyncCreateClipsTaskRequest
	GetMusicUrl() *string
	SetMusicVolume(v int32) *AsyncCreateClipsTaskRequest
	GetMusicVolume() *int32
	SetOpeningCreditsUrl(v string) *AsyncCreateClipsTaskRequest
	GetOpeningCreditsUrl() *string
	SetStickers(v []*AsyncCreateClipsTaskRequestStickers) *AsyncCreateClipsTaskRequest
	GetStickers() []*AsyncCreateClipsTaskRequestStickers
	SetSubtitleFontSize(v int32) *AsyncCreateClipsTaskRequest
	GetSubtitleFontSize() *int32
	SetTaskId(v string) *AsyncCreateClipsTaskRequest
	GetTaskId() *string
	SetVoiceStyle(v string) *AsyncCreateClipsTaskRequest
	GetVoiceStyle() *string
	SetVoiceVolume(v int32) *AsyncCreateClipsTaskRequest
	GetVoiceVolume() *int32
	SetWidth(v int32) *AsyncCreateClipsTaskRequest
	GetWidth() *int32
	SetWorkspaceId(v string) *AsyncCreateClipsTaskRequest
	GetWorkspaceId() *string
}

type AsyncCreateClipsTaskRequest struct {
	CloseMusic        *bool                                    `json:"CloseMusic,omitempty" xml:"CloseMusic,omitempty"`
	CloseSubtitle     *bool                                    `json:"CloseSubtitle,omitempty" xml:"CloseSubtitle,omitempty"`
	CloseVoice        *bool                                    `json:"CloseVoice,omitempty" xml:"CloseVoice,omitempty"`
	ClosingCreditsUrl *string                                  `json:"ClosingCreditsUrl,omitempty" xml:"ClosingCreditsUrl,omitempty"`
	ColorWords        []*AsyncCreateClipsTaskRequestColorWords `json:"ColorWords,omitempty" xml:"ColorWords,omitempty" type:"Repeated"`
	CosyVoiceAppKey   *string                                  `json:"CosyVoiceAppKey,omitempty" xml:"CosyVoiceAppKey,omitempty"`
	CosyVoiceToken    *string                                  `json:"CosyVoiceToken,omitempty" xml:"CosyVoiceToken,omitempty"`
	CustomVoiceStyle  *string                                  `json:"CustomVoiceStyle,omitempty" xml:"CustomVoiceStyle,omitempty"`
	// example:
	//
	// http://xxx/xxx.mp4
	CustomVoiceUrl *string `json:"CustomVoiceUrl,omitempty" xml:"CustomVoiceUrl,omitempty"`
	// example:
	//
	// 0
	CustomVoiceVolume *int32 `json:"CustomVoiceVolume,omitempty" xml:"CustomVoiceVolume,omitempty"`
	// example:
	//
	// 1920
	Height              *int32                                            `json:"Height,omitempty" xml:"Height,omitempty"`
	HighDefSourceVideos []*AsyncCreateClipsTaskRequestHighDefSourceVideos `json:"HighDefSourceVideos,omitempty" xml:"HighDefSourceVideos,omitempty" type:"Repeated"`
	MusicStyle          *string                                           `json:"MusicStyle,omitempty" xml:"MusicStyle,omitempty"`
	// example:
	//
	// http://music.mp4
	MusicUrl          *string                                `json:"MusicUrl,omitempty" xml:"MusicUrl,omitempty"`
	MusicVolume       *int32                                 `json:"MusicVolume,omitempty" xml:"MusicVolume,omitempty"`
	OpeningCreditsUrl *string                                `json:"OpeningCreditsUrl,omitempty" xml:"OpeningCreditsUrl,omitempty"`
	Stickers          []*AsyncCreateClipsTaskRequestStickers `json:"Stickers,omitempty" xml:"Stickers,omitempty" type:"Repeated"`
	SubtitleFontSize  *int32                                 `json:"SubtitleFontSize,omitempty" xml:"SubtitleFontSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 17a299a9-f223-4707-b0dd-4c22519bddf5
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	VoiceStyle  *string `json:"VoiceStyle,omitempty" xml:"VoiceStyle,omitempty"`
	VoiceVolume *int32  `json:"VoiceVolume,omitempty" xml:"VoiceVolume,omitempty"`
	// example:
	//
	// 1080
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-ipe7d81yq4sl5jmk
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s AsyncCreateClipsTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s AsyncCreateClipsTaskRequest) GoString() string {
	return s.String()
}

func (s *AsyncCreateClipsTaskRequest) GetCloseMusic() *bool {
	return s.CloseMusic
}

func (s *AsyncCreateClipsTaskRequest) GetCloseSubtitle() *bool {
	return s.CloseSubtitle
}

func (s *AsyncCreateClipsTaskRequest) GetCloseVoice() *bool {
	return s.CloseVoice
}

func (s *AsyncCreateClipsTaskRequest) GetClosingCreditsUrl() *string {
	return s.ClosingCreditsUrl
}

func (s *AsyncCreateClipsTaskRequest) GetColorWords() []*AsyncCreateClipsTaskRequestColorWords {
	return s.ColorWords
}

func (s *AsyncCreateClipsTaskRequest) GetCosyVoiceAppKey() *string {
	return s.CosyVoiceAppKey
}

func (s *AsyncCreateClipsTaskRequest) GetCosyVoiceToken() *string {
	return s.CosyVoiceToken
}

func (s *AsyncCreateClipsTaskRequest) GetCustomVoiceStyle() *string {
	return s.CustomVoiceStyle
}

func (s *AsyncCreateClipsTaskRequest) GetCustomVoiceUrl() *string {
	return s.CustomVoiceUrl
}

func (s *AsyncCreateClipsTaskRequest) GetCustomVoiceVolume() *int32 {
	return s.CustomVoiceVolume
}

func (s *AsyncCreateClipsTaskRequest) GetHeight() *int32 {
	return s.Height
}

func (s *AsyncCreateClipsTaskRequest) GetHighDefSourceVideos() []*AsyncCreateClipsTaskRequestHighDefSourceVideos {
	return s.HighDefSourceVideos
}

func (s *AsyncCreateClipsTaskRequest) GetMusicStyle() *string {
	return s.MusicStyle
}

func (s *AsyncCreateClipsTaskRequest) GetMusicUrl() *string {
	return s.MusicUrl
}

func (s *AsyncCreateClipsTaskRequest) GetMusicVolume() *int32 {
	return s.MusicVolume
}

func (s *AsyncCreateClipsTaskRequest) GetOpeningCreditsUrl() *string {
	return s.OpeningCreditsUrl
}

func (s *AsyncCreateClipsTaskRequest) GetStickers() []*AsyncCreateClipsTaskRequestStickers {
	return s.Stickers
}

func (s *AsyncCreateClipsTaskRequest) GetSubtitleFontSize() *int32 {
	return s.SubtitleFontSize
}

func (s *AsyncCreateClipsTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *AsyncCreateClipsTaskRequest) GetVoiceStyle() *string {
	return s.VoiceStyle
}

func (s *AsyncCreateClipsTaskRequest) GetVoiceVolume() *int32 {
	return s.VoiceVolume
}

func (s *AsyncCreateClipsTaskRequest) GetWidth() *int32 {
	return s.Width
}

func (s *AsyncCreateClipsTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AsyncCreateClipsTaskRequest) SetCloseMusic(v bool) *AsyncCreateClipsTaskRequest {
	s.CloseMusic = &v
	return s
}

func (s *AsyncCreateClipsTaskRequest) SetCloseSubtitle(v bool) *AsyncCreateClipsTaskRequest {
	s.CloseSubtitle = &v
	return s
}

func (s *AsyncCreateClipsTaskRequest) SetCloseVoice(v bool) *AsyncCreateClipsTaskRequest {
	s.CloseVoice = &v
	return s
}

func (s *AsyncCreateClipsTaskRequest) SetClosingCreditsUrl(v string) *AsyncCreateClipsTaskRequest {
	s.ClosingCreditsUrl = &v
	return s
}

func (s *AsyncCreateClipsTaskRequest) SetColorWords(v []*AsyncCreateClipsTaskRequestColorWords) *AsyncCreateClipsTaskRequest {
	s.ColorWords = v
	return s
}

func (s *AsyncCreateClipsTaskRequest) SetCosyVoiceAppKey(v string) *AsyncCreateClipsTaskRequest {
	s.CosyVoiceAppKey = &v
	return s
}

func (s *AsyncCreateClipsTaskRequest) SetCosyVoiceToken(v string) *AsyncCreateClipsTaskRequest {
	s.CosyVoiceToken = &v
	return s
}

func (s *AsyncCreateClipsTaskRequest) SetCustomVoiceStyle(v string) *AsyncCreateClipsTaskRequest {
	s.CustomVoiceStyle = &v
	return s
}

func (s *AsyncCreateClipsTaskRequest) SetCustomVoiceUrl(v string) *AsyncCreateClipsTaskRequest {
	s.CustomVoiceUrl = &v
	return s
}

func (s *AsyncCreateClipsTaskRequest) SetCustomVoiceVolume(v int32) *AsyncCreateClipsTaskRequest {
	s.CustomVoiceVolume = &v
	return s
}

func (s *AsyncCreateClipsTaskRequest) SetHeight(v int32) *AsyncCreateClipsTaskRequest {
	s.Height = &v
	return s
}

func (s *AsyncCreateClipsTaskRequest) SetHighDefSourceVideos(v []*AsyncCreateClipsTaskRequestHighDefSourceVideos) *AsyncCreateClipsTaskRequest {
	s.HighDefSourceVideos = v
	return s
}

func (s *AsyncCreateClipsTaskRequest) SetMusicStyle(v string) *AsyncCreateClipsTaskRequest {
	s.MusicStyle = &v
	return s
}

func (s *AsyncCreateClipsTaskRequest) SetMusicUrl(v string) *AsyncCreateClipsTaskRequest {
	s.MusicUrl = &v
	return s
}

func (s *AsyncCreateClipsTaskRequest) SetMusicVolume(v int32) *AsyncCreateClipsTaskRequest {
	s.MusicVolume = &v
	return s
}

func (s *AsyncCreateClipsTaskRequest) SetOpeningCreditsUrl(v string) *AsyncCreateClipsTaskRequest {
	s.OpeningCreditsUrl = &v
	return s
}

func (s *AsyncCreateClipsTaskRequest) SetStickers(v []*AsyncCreateClipsTaskRequestStickers) *AsyncCreateClipsTaskRequest {
	s.Stickers = v
	return s
}

func (s *AsyncCreateClipsTaskRequest) SetSubtitleFontSize(v int32) *AsyncCreateClipsTaskRequest {
	s.SubtitleFontSize = &v
	return s
}

func (s *AsyncCreateClipsTaskRequest) SetTaskId(v string) *AsyncCreateClipsTaskRequest {
	s.TaskId = &v
	return s
}

func (s *AsyncCreateClipsTaskRequest) SetVoiceStyle(v string) *AsyncCreateClipsTaskRequest {
	s.VoiceStyle = &v
	return s
}

func (s *AsyncCreateClipsTaskRequest) SetVoiceVolume(v int32) *AsyncCreateClipsTaskRequest {
	s.VoiceVolume = &v
	return s
}

func (s *AsyncCreateClipsTaskRequest) SetWidth(v int32) *AsyncCreateClipsTaskRequest {
	s.Width = &v
	return s
}

func (s *AsyncCreateClipsTaskRequest) SetWorkspaceId(v string) *AsyncCreateClipsTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *AsyncCreateClipsTaskRequest) Validate() error {
	if s.ColorWords != nil {
		for _, item := range s.ColorWords {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.HighDefSourceVideos != nil {
		for _, item := range s.HighDefSourceVideos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Stickers != nil {
		for _, item := range s.Stickers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AsyncCreateClipsTaskRequestColorWords struct {
	Content          *string `json:"Content,omitempty" xml:"Content,omitempty"`
	EffectColorStyle *string `json:"EffectColorStyle,omitempty" xml:"EffectColorStyle,omitempty"`
	FontSize         *int32  `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	// example:
	//
	// 0
	TimelineIn *int32 `json:"TimelineIn,omitempty" xml:"TimelineIn,omitempty"`
	// example:
	//
	// 5
	TimelineOut *int32 `json:"TimelineOut,omitempty" xml:"TimelineOut,omitempty"`
	// example:
	//
	// 0.2
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 0.5
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s AsyncCreateClipsTaskRequestColorWords) String() string {
	return dara.Prettify(s)
}

func (s AsyncCreateClipsTaskRequestColorWords) GoString() string {
	return s.String()
}

func (s *AsyncCreateClipsTaskRequestColorWords) GetContent() *string {
	return s.Content
}

func (s *AsyncCreateClipsTaskRequestColorWords) GetEffectColorStyle() *string {
	return s.EffectColorStyle
}

func (s *AsyncCreateClipsTaskRequestColorWords) GetFontSize() *int32 {
	return s.FontSize
}

func (s *AsyncCreateClipsTaskRequestColorWords) GetTimelineIn() *int32 {
	return s.TimelineIn
}

func (s *AsyncCreateClipsTaskRequestColorWords) GetTimelineOut() *int32 {
	return s.TimelineOut
}

func (s *AsyncCreateClipsTaskRequestColorWords) GetX() *float32 {
	return s.X
}

func (s *AsyncCreateClipsTaskRequestColorWords) GetY() *float32 {
	return s.Y
}

func (s *AsyncCreateClipsTaskRequestColorWords) SetContent(v string) *AsyncCreateClipsTaskRequestColorWords {
	s.Content = &v
	return s
}

func (s *AsyncCreateClipsTaskRequestColorWords) SetEffectColorStyle(v string) *AsyncCreateClipsTaskRequestColorWords {
	s.EffectColorStyle = &v
	return s
}

func (s *AsyncCreateClipsTaskRequestColorWords) SetFontSize(v int32) *AsyncCreateClipsTaskRequestColorWords {
	s.FontSize = &v
	return s
}

func (s *AsyncCreateClipsTaskRequestColorWords) SetTimelineIn(v int32) *AsyncCreateClipsTaskRequestColorWords {
	s.TimelineIn = &v
	return s
}

func (s *AsyncCreateClipsTaskRequestColorWords) SetTimelineOut(v int32) *AsyncCreateClipsTaskRequestColorWords {
	s.TimelineOut = &v
	return s
}

func (s *AsyncCreateClipsTaskRequestColorWords) SetX(v float32) *AsyncCreateClipsTaskRequestColorWords {
	s.X = &v
	return s
}

func (s *AsyncCreateClipsTaskRequestColorWords) SetY(v float32) *AsyncCreateClipsTaskRequestColorWords {
	s.Y = &v
	return s
}

func (s *AsyncCreateClipsTaskRequestColorWords) Validate() error {
	return dara.Validate(s)
}

type AsyncCreateClipsTaskRequestHighDefSourceVideos struct {
	VideoId   *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
	VideoName *string `json:"VideoName,omitempty" xml:"VideoName,omitempty"`
	VideoUrl  *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s AsyncCreateClipsTaskRequestHighDefSourceVideos) String() string {
	return dara.Prettify(s)
}

func (s AsyncCreateClipsTaskRequestHighDefSourceVideos) GoString() string {
	return s.String()
}

func (s *AsyncCreateClipsTaskRequestHighDefSourceVideos) GetVideoId() *string {
	return s.VideoId
}

func (s *AsyncCreateClipsTaskRequestHighDefSourceVideos) GetVideoName() *string {
	return s.VideoName
}

func (s *AsyncCreateClipsTaskRequestHighDefSourceVideos) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *AsyncCreateClipsTaskRequestHighDefSourceVideos) SetVideoId(v string) *AsyncCreateClipsTaskRequestHighDefSourceVideos {
	s.VideoId = &v
	return s
}

func (s *AsyncCreateClipsTaskRequestHighDefSourceVideos) SetVideoName(v string) *AsyncCreateClipsTaskRequestHighDefSourceVideos {
	s.VideoName = &v
	return s
}

func (s *AsyncCreateClipsTaskRequestHighDefSourceVideos) SetVideoUrl(v string) *AsyncCreateClipsTaskRequestHighDefSourceVideos {
	s.VideoUrl = &v
	return s
}

func (s *AsyncCreateClipsTaskRequestHighDefSourceVideos) Validate() error {
	return dara.Validate(s)
}

type AsyncCreateClipsTaskRequestStickers struct {
	// example:
	//
	// 10
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 8
	DyncFrames *int32 `json:"DyncFrames,omitempty" xml:"DyncFrames,omitempty"`
	// example:
	//
	// 100
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 5
	TimelineIn *int32 `json:"TimelineIn,omitempty" xml:"TimelineIn,omitempty"`
	// example:
	//
	// http://xxx/xxx.gif
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// example:
	//
	// 100
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// example:
	//
	// 200
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 200
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s AsyncCreateClipsTaskRequestStickers) String() string {
	return dara.Prettify(s)
}

func (s AsyncCreateClipsTaskRequestStickers) GoString() string {
	return s.String()
}

func (s *AsyncCreateClipsTaskRequestStickers) GetDuration() *int32 {
	return s.Duration
}

func (s *AsyncCreateClipsTaskRequestStickers) GetDyncFrames() *int32 {
	return s.DyncFrames
}

func (s *AsyncCreateClipsTaskRequestStickers) GetHeight() *int32 {
	return s.Height
}

func (s *AsyncCreateClipsTaskRequestStickers) GetTimelineIn() *int32 {
	return s.TimelineIn
}

func (s *AsyncCreateClipsTaskRequestStickers) GetUrl() *string {
	return s.Url
}

func (s *AsyncCreateClipsTaskRequestStickers) GetWidth() *int32 {
	return s.Width
}

func (s *AsyncCreateClipsTaskRequestStickers) GetX() *float32 {
	return s.X
}

func (s *AsyncCreateClipsTaskRequestStickers) GetY() *float32 {
	return s.Y
}

func (s *AsyncCreateClipsTaskRequestStickers) SetDuration(v int32) *AsyncCreateClipsTaskRequestStickers {
	s.Duration = &v
	return s
}

func (s *AsyncCreateClipsTaskRequestStickers) SetDyncFrames(v int32) *AsyncCreateClipsTaskRequestStickers {
	s.DyncFrames = &v
	return s
}

func (s *AsyncCreateClipsTaskRequestStickers) SetHeight(v int32) *AsyncCreateClipsTaskRequestStickers {
	s.Height = &v
	return s
}

func (s *AsyncCreateClipsTaskRequestStickers) SetTimelineIn(v int32) *AsyncCreateClipsTaskRequestStickers {
	s.TimelineIn = &v
	return s
}

func (s *AsyncCreateClipsTaskRequestStickers) SetUrl(v string) *AsyncCreateClipsTaskRequestStickers {
	s.Url = &v
	return s
}

func (s *AsyncCreateClipsTaskRequestStickers) SetWidth(v int32) *AsyncCreateClipsTaskRequestStickers {
	s.Width = &v
	return s
}

func (s *AsyncCreateClipsTaskRequestStickers) SetX(v float32) *AsyncCreateClipsTaskRequestStickers {
	s.X = &v
	return s
}

func (s *AsyncCreateClipsTaskRequestStickers) SetY(v float32) *AsyncCreateClipsTaskRequestStickers {
	s.Y = &v
	return s
}

func (s *AsyncCreateClipsTaskRequestStickers) Validate() error {
	return dara.Validate(s)
}
