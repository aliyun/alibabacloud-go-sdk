// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncCreateClipsTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdaptMode(v string) *AsyncCreateClipsTaskShrinkRequest
	GetAdaptMode() *string
	SetAlignment(v string) *AsyncCreateClipsTaskShrinkRequest
	GetAlignment() *string
	SetCloseMusic(v bool) *AsyncCreateClipsTaskShrinkRequest
	GetCloseMusic() *bool
	SetCloseSubtitle(v bool) *AsyncCreateClipsTaskShrinkRequest
	GetCloseSubtitle() *bool
	SetCloseVoice(v bool) *AsyncCreateClipsTaskShrinkRequest
	GetCloseVoice() *bool
	SetClosingCreditsUrl(v string) *AsyncCreateClipsTaskShrinkRequest
	GetClosingCreditsUrl() *string
	SetColorWordsShrink(v string) *AsyncCreateClipsTaskShrinkRequest
	GetColorWordsShrink() *string
	SetCosyVoiceAppKey(v string) *AsyncCreateClipsTaskShrinkRequest
	GetCosyVoiceAppKey() *string
	SetCosyVoiceToken(v string) *AsyncCreateClipsTaskShrinkRequest
	GetCosyVoiceToken() *string
	SetCustomVoiceStyle(v string) *AsyncCreateClipsTaskShrinkRequest
	GetCustomVoiceStyle() *string
	SetCustomVoiceUrl(v string) *AsyncCreateClipsTaskShrinkRequest
	GetCustomVoiceUrl() *string
	SetCustomVoiceVolume(v int32) *AsyncCreateClipsTaskShrinkRequest
	GetCustomVoiceVolume() *int32
	SetHeight(v int32) *AsyncCreateClipsTaskShrinkRequest
	GetHeight() *int32
	SetHighDefSourceVideosShrink(v string) *AsyncCreateClipsTaskShrinkRequest
	GetHighDefSourceVideosShrink() *string
	SetMusicStyle(v string) *AsyncCreateClipsTaskShrinkRequest
	GetMusicStyle() *string
	SetMusicUrl(v string) *AsyncCreateClipsTaskShrinkRequest
	GetMusicUrl() *string
	SetMusicVolume(v int32) *AsyncCreateClipsTaskShrinkRequest
	GetMusicVolume() *int32
	SetOpeningCreditsUrl(v string) *AsyncCreateClipsTaskShrinkRequest
	GetOpeningCreditsUrl() *string
	SetStickersShrink(v string) *AsyncCreateClipsTaskShrinkRequest
	GetStickersShrink() *string
	SetSubtitleFontSize(v int32) *AsyncCreateClipsTaskShrinkRequest
	GetSubtitleFontSize() *int32
	SetTaskId(v string) *AsyncCreateClipsTaskShrinkRequest
	GetTaskId() *string
	SetTextWidth(v string) *AsyncCreateClipsTaskShrinkRequest
	GetTextWidth() *string
	SetVoiceStyle(v string) *AsyncCreateClipsTaskShrinkRequest
	GetVoiceStyle() *string
	SetVoiceVolume(v int32) *AsyncCreateClipsTaskShrinkRequest
	GetVoiceVolume() *int32
	SetWidth(v int32) *AsyncCreateClipsTaskShrinkRequest
	GetWidth() *int32
	SetWorkspaceId(v string) *AsyncCreateClipsTaskShrinkRequest
	GetWorkspaceId() *string
}

type AsyncCreateClipsTaskShrinkRequest struct {
	// example:
	//
	// AutoWrap：自动换行
	//
	// AutoScale：自动缩放
	//
	// AutoWrapAtSpaces：只在空格位置自动换行（适用于纯英文字幕自动换行场景）
	AdaptMode *string `json:"AdaptMode,omitempty" xml:"AdaptMode,omitempty"`
	// example:
	//
	// 支持设置：
	//
	// TopLeft：视频左上角
	//
	// TopCenter：视频竖直中轴线上侧
	//
	// TopRight：视频右上角
	//
	// CenterLeft：视频水平中轴线左侧
	//
	// CenterCenter：视频中心位置
	//
	// CenterRight：视频水平中轴线右侧
	//
	// BottomLeft：视频左下角
	//
	// BottomCenter：视频竖直中轴线下侧
	//
	// BottomRight：视频右下角
	//
	// 若需要在不同对齐方式下准确定位字幕位置，建议设置以下对齐方式：
	//
	// Left，左对齐，X、Y传入字幕左上角顶点相对于视频左上角的坐标
	//
	// Center，居中对齐，X、Y传入字幕中轴线上边界交点相对于视频左上角的坐标
	//
	// Right，右对齐，X、Y传入字幕右上角顶点相对于视频左上角的坐标
	Alignment *string `json:"Alignment,omitempty" xml:"Alignment,omitempty"`
	// Specifies whether to disable the background music.
	//
	// example:
	//
	// true
	CloseMusic *bool `json:"CloseMusic,omitempty" xml:"CloseMusic,omitempty"`
	// Specifies whether to disable the subtitles.
	CloseSubtitle *bool `json:"CloseSubtitle,omitempty" xml:"CloseSubtitle,omitempty"`
	// Specifies whether to disable the narration voice.
	//
	// example:
	//
	// false
	CloseVoice *bool `json:"CloseVoice,omitempty" xml:"CloseVoice,omitempty"`
	// The URL of the closing credits video.
	//
	// example:
	//
	// http://xxx/xxx.mp4
	ClosingCreditsUrl *string `json:"ClosingCreditsUrl,omitempty" xml:"ClosingCreditsUrl,omitempty"`
	// The array of animated text elements.
	ColorWordsShrink *string `json:"ColorWords,omitempty" xml:"ColorWords,omitempty"`
	// The AppKey of CosyVoice.
	//
	// example:
	//
	// ddgsase
	CosyVoiceAppKey *string `json:"CosyVoiceAppKey,omitempty" xml:"CosyVoiceAppKey,omitempty"`
	// The token of CosyVoice.
	//
	// example:
	//
	// xxsfazs
	CosyVoiceToken *string `json:"CosyVoiceToken,omitempty" xml:"CosyVoiceToken,omitempty"`
	// The voice tone of CosyVoice.
	//
	// example:
	//
	// longxian_normal
	CustomVoiceStyle *string `json:"CustomVoiceStyle,omitempty" xml:"CustomVoiceStyle,omitempty"`
	// The URL of the custom audio track.
	//
	// example:
	//
	// http://xxx/xxx.mp4
	CustomVoiceUrl *string `json:"CustomVoiceUrl,omitempty" xml:"CustomVoiceUrl,omitempty"`
	// The volume of the custom audio track.
	//
	// example:
	//
	// 0
	CustomVoiceVolume *int32 `json:"CustomVoiceVolume,omitempty" xml:"CustomVoiceVolume,omitempty"`
	// The height of the video.
	//
	// example:
	//
	// 1920
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The list of high-definition video structures.
	HighDefSourceVideosShrink *string `json:"HighDefSourceVideos,omitempty" xml:"HighDefSourceVideos,omitempty"`
	// The type of recommended music.
	//
	// example:
	//
	// 浪漫, 美食,国风,轻快,动感,舒缓,搞怪,时尚
	MusicStyle *string `json:"MusicStyle,omitempty" xml:"MusicStyle,omitempty"`
	// The URL of the background music.
	//
	// example:
	//
	// http://music.mp4
	MusicUrl *string `json:"MusicUrl,omitempty" xml:"MusicUrl,omitempty"`
	// The volume of the background music.
	//
	// example:
	//
	// 0-10，默认5
	MusicVolume *int32 `json:"MusicVolume,omitempty" xml:"MusicVolume,omitempty"`
	// The URL of the opening credits video.
	//
	// example:
	//
	// http://xxx/xxx.mp4
	OpeningCreditsUrl *string `json:"OpeningCreditsUrl,omitempty" xml:"OpeningCreditsUrl,omitempty"`
	// The array of sticker structures.
	StickersShrink *string `json:"Stickers,omitempty" xml:"Stickers,omitempty"`
	// The font size of the subtitles.
	//
	// example:
	//
	// 默认120
	SubtitleFontSize *int32 `json:"SubtitleFontSize,omitempty" xml:"SubtitleFontSize,omitempty"`
	// The unique ID of the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 17a299a9-f223-4707-b0dd-4c22519bddf5
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 将按照该值设置文本框宽度进行自动换行或缩放。不填写时，会按照视频宽度进行自动换行或缩放。当值大于0小于等于1时，表示相对输出视频的宽度，当值大于1时，表示绝对像素值。
	TextWidth *string `json:"TextWidth,omitempty" xml:"TextWidth,omitempty"`
	// The type of narration voice.
	//
	// example:
	//
	// 甜美女声
	//
	// 中国台湾话女声
	//
	// 舌尖男声
	//
	// 新闻男声
	//
	// 激昂解说
	//
	// 标准女声
	//
	// 悬疑解说
	//
	// 广告男声
	//
	// 温柔女声
	//
	// 资讯女声
	//
	// 新闻女声
	//
	// 萝莉女声
	//
	// 磁性男声
	VoiceStyle *string `json:"VoiceStyle,omitempty" xml:"VoiceStyle,omitempty"`
	// The volume of the narration voice.
	//
	// example:
	//
	// 0-10，默认5
	VoiceVolume *int32 `json:"VoiceVolume,omitempty" xml:"VoiceVolume,omitempty"`
	// The width of the video.
	//
	// example:
	//
	// 1080
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// The [Bailian workspace ID](https://help.aliyun.com/document_detail/2782167.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// llm-ipe7d81yq4sl5jmk
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s AsyncCreateClipsTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AsyncCreateClipsTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetAdaptMode() *string {
	return s.AdaptMode
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetAlignment() *string {
	return s.Alignment
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetCloseMusic() *bool {
	return s.CloseMusic
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetCloseSubtitle() *bool {
	return s.CloseSubtitle
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetCloseVoice() *bool {
	return s.CloseVoice
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetClosingCreditsUrl() *string {
	return s.ClosingCreditsUrl
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetColorWordsShrink() *string {
	return s.ColorWordsShrink
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetCosyVoiceAppKey() *string {
	return s.CosyVoiceAppKey
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetCosyVoiceToken() *string {
	return s.CosyVoiceToken
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetCustomVoiceStyle() *string {
	return s.CustomVoiceStyle
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetCustomVoiceUrl() *string {
	return s.CustomVoiceUrl
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetCustomVoiceVolume() *int32 {
	return s.CustomVoiceVolume
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetHeight() *int32 {
	return s.Height
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetHighDefSourceVideosShrink() *string {
	return s.HighDefSourceVideosShrink
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetMusicStyle() *string {
	return s.MusicStyle
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetMusicUrl() *string {
	return s.MusicUrl
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetMusicVolume() *int32 {
	return s.MusicVolume
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetOpeningCreditsUrl() *string {
	return s.OpeningCreditsUrl
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetStickersShrink() *string {
	return s.StickersShrink
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetSubtitleFontSize() *int32 {
	return s.SubtitleFontSize
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetTextWidth() *string {
	return s.TextWidth
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetVoiceStyle() *string {
	return s.VoiceStyle
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetVoiceVolume() *int32 {
	return s.VoiceVolume
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetWidth() *int32 {
	return s.Width
}

func (s *AsyncCreateClipsTaskShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetAdaptMode(v string) *AsyncCreateClipsTaskShrinkRequest {
	s.AdaptMode = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetAlignment(v string) *AsyncCreateClipsTaskShrinkRequest {
	s.Alignment = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetCloseMusic(v bool) *AsyncCreateClipsTaskShrinkRequest {
	s.CloseMusic = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetCloseSubtitle(v bool) *AsyncCreateClipsTaskShrinkRequest {
	s.CloseSubtitle = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetCloseVoice(v bool) *AsyncCreateClipsTaskShrinkRequest {
	s.CloseVoice = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetClosingCreditsUrl(v string) *AsyncCreateClipsTaskShrinkRequest {
	s.ClosingCreditsUrl = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetColorWordsShrink(v string) *AsyncCreateClipsTaskShrinkRequest {
	s.ColorWordsShrink = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetCosyVoiceAppKey(v string) *AsyncCreateClipsTaskShrinkRequest {
	s.CosyVoiceAppKey = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetCosyVoiceToken(v string) *AsyncCreateClipsTaskShrinkRequest {
	s.CosyVoiceToken = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetCustomVoiceStyle(v string) *AsyncCreateClipsTaskShrinkRequest {
	s.CustomVoiceStyle = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetCustomVoiceUrl(v string) *AsyncCreateClipsTaskShrinkRequest {
	s.CustomVoiceUrl = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetCustomVoiceVolume(v int32) *AsyncCreateClipsTaskShrinkRequest {
	s.CustomVoiceVolume = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetHeight(v int32) *AsyncCreateClipsTaskShrinkRequest {
	s.Height = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetHighDefSourceVideosShrink(v string) *AsyncCreateClipsTaskShrinkRequest {
	s.HighDefSourceVideosShrink = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetMusicStyle(v string) *AsyncCreateClipsTaskShrinkRequest {
	s.MusicStyle = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetMusicUrl(v string) *AsyncCreateClipsTaskShrinkRequest {
	s.MusicUrl = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetMusicVolume(v int32) *AsyncCreateClipsTaskShrinkRequest {
	s.MusicVolume = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetOpeningCreditsUrl(v string) *AsyncCreateClipsTaskShrinkRequest {
	s.OpeningCreditsUrl = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetStickersShrink(v string) *AsyncCreateClipsTaskShrinkRequest {
	s.StickersShrink = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetSubtitleFontSize(v int32) *AsyncCreateClipsTaskShrinkRequest {
	s.SubtitleFontSize = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetTaskId(v string) *AsyncCreateClipsTaskShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetTextWidth(v string) *AsyncCreateClipsTaskShrinkRequest {
	s.TextWidth = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetVoiceStyle(v string) *AsyncCreateClipsTaskShrinkRequest {
	s.VoiceStyle = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetVoiceVolume(v int32) *AsyncCreateClipsTaskShrinkRequest {
	s.VoiceVolume = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetWidth(v int32) *AsyncCreateClipsTaskShrinkRequest {
	s.Width = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) SetWorkspaceId(v string) *AsyncCreateClipsTaskShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *AsyncCreateClipsTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
