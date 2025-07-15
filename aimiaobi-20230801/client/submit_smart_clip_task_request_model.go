// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSmartClipTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEditingConfig(v *SubmitSmartClipTaskRequestEditingConfig) *SubmitSmartClipTaskRequest
	GetEditingConfig() *SubmitSmartClipTaskRequestEditingConfig
	SetExtendParam(v string) *SubmitSmartClipTaskRequest
	GetExtendParam() *string
	SetInputConfig(v *SubmitSmartClipTaskRequestInputConfig) *SubmitSmartClipTaskRequest
	GetInputConfig() *SubmitSmartClipTaskRequestInputConfig
	SetOutputConfig(v *SubmitSmartClipTaskRequestOutputConfig) *SubmitSmartClipTaskRequest
	GetOutputConfig() *SubmitSmartClipTaskRequestOutputConfig
	SetWorkspaceId(v string) *SubmitSmartClipTaskRequest
	GetWorkspaceId() *string
}

type SubmitSmartClipTaskRequest struct {
	EditingConfig *SubmitSmartClipTaskRequestEditingConfig `json:"EditingConfig,omitempty" xml:"EditingConfig,omitempty" type:"Struct"`
	ExtendParam   *string                                  `json:"ExtendParam,omitempty" xml:"ExtendParam,omitempty"`
	// This parameter is required.
	InputConfig  *SubmitSmartClipTaskRequestInputConfig  `json:"InputConfig,omitempty" xml:"InputConfig,omitempty" type:"Struct"`
	OutputConfig *SubmitSmartClipTaskRequestOutputConfig `json:"OutputConfig,omitempty" xml:"OutputConfig,omitempty" type:"Struct"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SubmitSmartClipTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmartClipTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitSmartClipTaskRequest) GetEditingConfig() *SubmitSmartClipTaskRequestEditingConfig {
	return s.EditingConfig
}

func (s *SubmitSmartClipTaskRequest) GetExtendParam() *string {
	return s.ExtendParam
}

func (s *SubmitSmartClipTaskRequest) GetInputConfig() *SubmitSmartClipTaskRequestInputConfig {
	return s.InputConfig
}

func (s *SubmitSmartClipTaskRequest) GetOutputConfig() *SubmitSmartClipTaskRequestOutputConfig {
	return s.OutputConfig
}

func (s *SubmitSmartClipTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SubmitSmartClipTaskRequest) SetEditingConfig(v *SubmitSmartClipTaskRequestEditingConfig) *SubmitSmartClipTaskRequest {
	s.EditingConfig = v
	return s
}

func (s *SubmitSmartClipTaskRequest) SetExtendParam(v string) *SubmitSmartClipTaskRequest {
	s.ExtendParam = &v
	return s
}

func (s *SubmitSmartClipTaskRequest) SetInputConfig(v *SubmitSmartClipTaskRequestInputConfig) *SubmitSmartClipTaskRequest {
	s.InputConfig = v
	return s
}

func (s *SubmitSmartClipTaskRequest) SetOutputConfig(v *SubmitSmartClipTaskRequestOutputConfig) *SubmitSmartClipTaskRequest {
	s.OutputConfig = v
	return s
}

func (s *SubmitSmartClipTaskRequest) SetWorkspaceId(v string) *SubmitSmartClipTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SubmitSmartClipTaskRequest) Validate() error {
	return dara.Validate(s)
}

type SubmitSmartClipTaskRequestEditingConfig struct {
	BackgroundMusicConfig *SubmitSmartClipTaskRequestEditingConfigBackgroundMusicConfig `json:"BackgroundMusicConfig,omitempty" xml:"BackgroundMusicConfig,omitempty" type:"Struct"`
	MediaConfig           *SubmitSmartClipTaskRequestEditingConfigMediaConfig           `json:"MediaConfig,omitempty" xml:"MediaConfig,omitempty" type:"Struct"`
	SpeechConfig          *SubmitSmartClipTaskRequestEditingConfigSpeechConfig          `json:"SpeechConfig,omitempty" xml:"SpeechConfig,omitempty" type:"Struct"`
	TitleConfig           *SubmitSmartClipTaskRequestEditingConfigTitleConfig           `json:"TitleConfig,omitempty" xml:"TitleConfig,omitempty" type:"Struct"`
}

func (s SubmitSmartClipTaskRequestEditingConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmartClipTaskRequestEditingConfig) GoString() string {
	return s.String()
}

func (s *SubmitSmartClipTaskRequestEditingConfig) GetBackgroundMusicConfig() *SubmitSmartClipTaskRequestEditingConfigBackgroundMusicConfig {
	return s.BackgroundMusicConfig
}

func (s *SubmitSmartClipTaskRequestEditingConfig) GetMediaConfig() *SubmitSmartClipTaskRequestEditingConfigMediaConfig {
	return s.MediaConfig
}

func (s *SubmitSmartClipTaskRequestEditingConfig) GetSpeechConfig() *SubmitSmartClipTaskRequestEditingConfigSpeechConfig {
	return s.SpeechConfig
}

func (s *SubmitSmartClipTaskRequestEditingConfig) GetTitleConfig() *SubmitSmartClipTaskRequestEditingConfigTitleConfig {
	return s.TitleConfig
}

func (s *SubmitSmartClipTaskRequestEditingConfig) SetBackgroundMusicConfig(v *SubmitSmartClipTaskRequestEditingConfigBackgroundMusicConfig) *SubmitSmartClipTaskRequestEditingConfig {
	s.BackgroundMusicConfig = v
	return s
}

func (s *SubmitSmartClipTaskRequestEditingConfig) SetMediaConfig(v *SubmitSmartClipTaskRequestEditingConfigMediaConfig) *SubmitSmartClipTaskRequestEditingConfig {
	s.MediaConfig = v
	return s
}

func (s *SubmitSmartClipTaskRequestEditingConfig) SetSpeechConfig(v *SubmitSmartClipTaskRequestEditingConfigSpeechConfig) *SubmitSmartClipTaskRequestEditingConfig {
	s.SpeechConfig = v
	return s
}

func (s *SubmitSmartClipTaskRequestEditingConfig) SetTitleConfig(v *SubmitSmartClipTaskRequestEditingConfigTitleConfig) *SubmitSmartClipTaskRequestEditingConfig {
	s.TitleConfig = v
	return s
}

func (s *SubmitSmartClipTaskRequestEditingConfig) Validate() error {
	return dara.Validate(s)
}

type SubmitSmartClipTaskRequestEditingConfigBackgroundMusicConfig struct {
	Style *string `json:"Style,omitempty" xml:"Style,omitempty"`
	// example:
	//
	// 0.2
	Volume *float64 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s SubmitSmartClipTaskRequestEditingConfigBackgroundMusicConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmartClipTaskRequestEditingConfigBackgroundMusicConfig) GoString() string {
	return s.String()
}

func (s *SubmitSmartClipTaskRequestEditingConfigBackgroundMusicConfig) GetStyle() *string {
	return s.Style
}

func (s *SubmitSmartClipTaskRequestEditingConfigBackgroundMusicConfig) GetVolume() *float64 {
	return s.Volume
}

func (s *SubmitSmartClipTaskRequestEditingConfigBackgroundMusicConfig) SetStyle(v string) *SubmitSmartClipTaskRequestEditingConfigBackgroundMusicConfig {
	s.Style = &v
	return s
}

func (s *SubmitSmartClipTaskRequestEditingConfigBackgroundMusicConfig) SetVolume(v float64) *SubmitSmartClipTaskRequestEditingConfigBackgroundMusicConfig {
	s.Volume = &v
	return s
}

func (s *SubmitSmartClipTaskRequestEditingConfigBackgroundMusicConfig) Validate() error {
	return dara.Validate(s)
}

type SubmitSmartClipTaskRequestEditingConfigMediaConfig struct {
	Volume *float64 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s SubmitSmartClipTaskRequestEditingConfigMediaConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmartClipTaskRequestEditingConfigMediaConfig) GoString() string {
	return s.String()
}

func (s *SubmitSmartClipTaskRequestEditingConfigMediaConfig) GetVolume() *float64 {
	return s.Volume
}

func (s *SubmitSmartClipTaskRequestEditingConfigMediaConfig) SetVolume(v float64) *SubmitSmartClipTaskRequestEditingConfigMediaConfig {
	s.Volume = &v
	return s
}

func (s *SubmitSmartClipTaskRequestEditingConfigMediaConfig) Validate() error {
	return dara.Validate(s)
}

type SubmitSmartClipTaskRequestEditingConfigSpeechConfig struct {
	AsrConfig *SubmitSmartClipTaskRequestEditingConfigSpeechConfigAsrConfig `json:"AsrConfig,omitempty" xml:"AsrConfig,omitempty" type:"Struct"`
	// example:
	//
	// 0
	SpeechRate *float64 `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	Style      *string  `json:"Style,omitempty" xml:"Style,omitempty"`
	Voice      *string  `json:"Voice,omitempty" xml:"Voice,omitempty"`
	// example:
	//
	// 0.5
	Volume *float64 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s SubmitSmartClipTaskRequestEditingConfigSpeechConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmartClipTaskRequestEditingConfigSpeechConfig) GoString() string {
	return s.String()
}

func (s *SubmitSmartClipTaskRequestEditingConfigSpeechConfig) GetAsrConfig() *SubmitSmartClipTaskRequestEditingConfigSpeechConfigAsrConfig {
	return s.AsrConfig
}

func (s *SubmitSmartClipTaskRequestEditingConfigSpeechConfig) GetSpeechRate() *float64 {
	return s.SpeechRate
}

func (s *SubmitSmartClipTaskRequestEditingConfigSpeechConfig) GetStyle() *string {
	return s.Style
}

func (s *SubmitSmartClipTaskRequestEditingConfigSpeechConfig) GetVoice() *string {
	return s.Voice
}

func (s *SubmitSmartClipTaskRequestEditingConfigSpeechConfig) GetVolume() *float64 {
	return s.Volume
}

func (s *SubmitSmartClipTaskRequestEditingConfigSpeechConfig) SetAsrConfig(v *SubmitSmartClipTaskRequestEditingConfigSpeechConfigAsrConfig) *SubmitSmartClipTaskRequestEditingConfigSpeechConfig {
	s.AsrConfig = v
	return s
}

func (s *SubmitSmartClipTaskRequestEditingConfigSpeechConfig) SetSpeechRate(v float64) *SubmitSmartClipTaskRequestEditingConfigSpeechConfig {
	s.SpeechRate = &v
	return s
}

func (s *SubmitSmartClipTaskRequestEditingConfigSpeechConfig) SetStyle(v string) *SubmitSmartClipTaskRequestEditingConfigSpeechConfig {
	s.Style = &v
	return s
}

func (s *SubmitSmartClipTaskRequestEditingConfigSpeechConfig) SetVoice(v string) *SubmitSmartClipTaskRequestEditingConfigSpeechConfig {
	s.Voice = &v
	return s
}

func (s *SubmitSmartClipTaskRequestEditingConfigSpeechConfig) SetVolume(v float64) *SubmitSmartClipTaskRequestEditingConfigSpeechConfig {
	s.Volume = &v
	return s
}

func (s *SubmitSmartClipTaskRequestEditingConfigSpeechConfig) Validate() error {
	return dara.Validate(s)
}

type SubmitSmartClipTaskRequestEditingConfigSpeechConfigAsrConfig struct {
	Alignment *string `json:"Alignment,omitempty" xml:"Alignment,omitempty"`
	// example:
	//
	// SimSun
	Font *string `json:"Font,omitempty" xml:"Font,omitempty"`
	// example:
	//
	// #ffffff
	FontColor *string `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	// example:
	//
	// 0
	FontSize *string `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	// example:
	//
	// 0
	Spacing *string  `json:"Spacing,omitempty" xml:"Spacing,omitempty"`
	X       *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y       *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s SubmitSmartClipTaskRequestEditingConfigSpeechConfigAsrConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmartClipTaskRequestEditingConfigSpeechConfigAsrConfig) GoString() string {
	return s.String()
}

func (s *SubmitSmartClipTaskRequestEditingConfigSpeechConfigAsrConfig) GetAlignment() *string {
	return s.Alignment
}

func (s *SubmitSmartClipTaskRequestEditingConfigSpeechConfigAsrConfig) GetFont() *string {
	return s.Font
}

func (s *SubmitSmartClipTaskRequestEditingConfigSpeechConfigAsrConfig) GetFontColor() *string {
	return s.FontColor
}

func (s *SubmitSmartClipTaskRequestEditingConfigSpeechConfigAsrConfig) GetFontSize() *string {
	return s.FontSize
}

func (s *SubmitSmartClipTaskRequestEditingConfigSpeechConfigAsrConfig) GetSpacing() *string {
	return s.Spacing
}

func (s *SubmitSmartClipTaskRequestEditingConfigSpeechConfigAsrConfig) GetX() *float32 {
	return s.X
}

func (s *SubmitSmartClipTaskRequestEditingConfigSpeechConfigAsrConfig) GetY() *float32 {
	return s.Y
}

func (s *SubmitSmartClipTaskRequestEditingConfigSpeechConfigAsrConfig) SetAlignment(v string) *SubmitSmartClipTaskRequestEditingConfigSpeechConfigAsrConfig {
	s.Alignment = &v
	return s
}

func (s *SubmitSmartClipTaskRequestEditingConfigSpeechConfigAsrConfig) SetFont(v string) *SubmitSmartClipTaskRequestEditingConfigSpeechConfigAsrConfig {
	s.Font = &v
	return s
}

func (s *SubmitSmartClipTaskRequestEditingConfigSpeechConfigAsrConfig) SetFontColor(v string) *SubmitSmartClipTaskRequestEditingConfigSpeechConfigAsrConfig {
	s.FontColor = &v
	return s
}

func (s *SubmitSmartClipTaskRequestEditingConfigSpeechConfigAsrConfig) SetFontSize(v string) *SubmitSmartClipTaskRequestEditingConfigSpeechConfigAsrConfig {
	s.FontSize = &v
	return s
}

func (s *SubmitSmartClipTaskRequestEditingConfigSpeechConfigAsrConfig) SetSpacing(v string) *SubmitSmartClipTaskRequestEditingConfigSpeechConfigAsrConfig {
	s.Spacing = &v
	return s
}

func (s *SubmitSmartClipTaskRequestEditingConfigSpeechConfigAsrConfig) SetX(v float32) *SubmitSmartClipTaskRequestEditingConfigSpeechConfigAsrConfig {
	s.X = &v
	return s
}

func (s *SubmitSmartClipTaskRequestEditingConfigSpeechConfigAsrConfig) SetY(v float32) *SubmitSmartClipTaskRequestEditingConfigSpeechConfigAsrConfig {
	s.Y = &v
	return s
}

func (s *SubmitSmartClipTaskRequestEditingConfigSpeechConfigAsrConfig) Validate() error {
	return dara.Validate(s)
}

type SubmitSmartClipTaskRequestEditingConfigTitleConfig struct {
	// example:
	//
	// TopLeft
	Alignment *string `json:"Alignment,omitempty" xml:"Alignment,omitempty"`
	// example:
	//
	// 2
	TimelineIn *float32 `json:"TimelineIn,omitempty" xml:"TimelineIn,omitempty"`
	// example:
	//
	// 3
	TimelineOut *float32 `json:"TimelineOut,omitempty" xml:"TimelineOut,omitempty"`
	// example:
	//
	// 100
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 100
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s SubmitSmartClipTaskRequestEditingConfigTitleConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmartClipTaskRequestEditingConfigTitleConfig) GoString() string {
	return s.String()
}

func (s *SubmitSmartClipTaskRequestEditingConfigTitleConfig) GetAlignment() *string {
	return s.Alignment
}

func (s *SubmitSmartClipTaskRequestEditingConfigTitleConfig) GetTimelineIn() *float32 {
	return s.TimelineIn
}

func (s *SubmitSmartClipTaskRequestEditingConfigTitleConfig) GetTimelineOut() *float32 {
	return s.TimelineOut
}

func (s *SubmitSmartClipTaskRequestEditingConfigTitleConfig) GetX() *float32 {
	return s.X
}

func (s *SubmitSmartClipTaskRequestEditingConfigTitleConfig) GetY() *float32 {
	return s.Y
}

func (s *SubmitSmartClipTaskRequestEditingConfigTitleConfig) SetAlignment(v string) *SubmitSmartClipTaskRequestEditingConfigTitleConfig {
	s.Alignment = &v
	return s
}

func (s *SubmitSmartClipTaskRequestEditingConfigTitleConfig) SetTimelineIn(v float32) *SubmitSmartClipTaskRequestEditingConfigTitleConfig {
	s.TimelineIn = &v
	return s
}

func (s *SubmitSmartClipTaskRequestEditingConfigTitleConfig) SetTimelineOut(v float32) *SubmitSmartClipTaskRequestEditingConfigTitleConfig {
	s.TimelineOut = &v
	return s
}

func (s *SubmitSmartClipTaskRequestEditingConfigTitleConfig) SetX(v float32) *SubmitSmartClipTaskRequestEditingConfigTitleConfig {
	s.X = &v
	return s
}

func (s *SubmitSmartClipTaskRequestEditingConfigTitleConfig) SetY(v float32) *SubmitSmartClipTaskRequestEditingConfigTitleConfig {
	s.Y = &v
	return s
}

func (s *SubmitSmartClipTaskRequestEditingConfigTitleConfig) Validate() error {
	return dara.Validate(s)
}

type SubmitSmartClipTaskRequestInputConfig struct {
	BackgroundMusics []*SubmitSmartClipTaskRequestInputConfigBackgroundMusics `json:"BackgroundMusics,omitempty" xml:"BackgroundMusics,omitempty" type:"Repeated"`
	SpeechTexts      []*string                                                `json:"SpeechTexts,omitempty" xml:"SpeechTexts,omitempty" type:"Repeated"`
	Stickers         []*SubmitSmartClipTaskRequestInputConfigStickers         `json:"Stickers,omitempty" xml:"Stickers,omitempty" type:"Repeated"`
	Titles           []*string                                                `json:"Titles,omitempty" xml:"Titles,omitempty" type:"Repeated"`
	// This parameter is required.
	VideoIds []*SubmitSmartClipTaskRequestInputConfigVideoIds `json:"VideoIds,omitempty" xml:"VideoIds,omitempty" type:"Repeated"`
}

func (s SubmitSmartClipTaskRequestInputConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmartClipTaskRequestInputConfig) GoString() string {
	return s.String()
}

func (s *SubmitSmartClipTaskRequestInputConfig) GetBackgroundMusics() []*SubmitSmartClipTaskRequestInputConfigBackgroundMusics {
	return s.BackgroundMusics
}

func (s *SubmitSmartClipTaskRequestInputConfig) GetSpeechTexts() []*string {
	return s.SpeechTexts
}

func (s *SubmitSmartClipTaskRequestInputConfig) GetStickers() []*SubmitSmartClipTaskRequestInputConfigStickers {
	return s.Stickers
}

func (s *SubmitSmartClipTaskRequestInputConfig) GetTitles() []*string {
	return s.Titles
}

func (s *SubmitSmartClipTaskRequestInputConfig) GetVideoIds() []*SubmitSmartClipTaskRequestInputConfigVideoIds {
	return s.VideoIds
}

func (s *SubmitSmartClipTaskRequestInputConfig) SetBackgroundMusics(v []*SubmitSmartClipTaskRequestInputConfigBackgroundMusics) *SubmitSmartClipTaskRequestInputConfig {
	s.BackgroundMusics = v
	return s
}

func (s *SubmitSmartClipTaskRequestInputConfig) SetSpeechTexts(v []*string) *SubmitSmartClipTaskRequestInputConfig {
	s.SpeechTexts = v
	return s
}

func (s *SubmitSmartClipTaskRequestInputConfig) SetStickers(v []*SubmitSmartClipTaskRequestInputConfigStickers) *SubmitSmartClipTaskRequestInputConfig {
	s.Stickers = v
	return s
}

func (s *SubmitSmartClipTaskRequestInputConfig) SetTitles(v []*string) *SubmitSmartClipTaskRequestInputConfig {
	s.Titles = v
	return s
}

func (s *SubmitSmartClipTaskRequestInputConfig) SetVideoIds(v []*SubmitSmartClipTaskRequestInputConfigVideoIds) *SubmitSmartClipTaskRequestInputConfig {
	s.VideoIds = v
	return s
}

func (s *SubmitSmartClipTaskRequestInputConfig) Validate() error {
	return dara.Validate(s)
}

type SubmitSmartClipTaskRequestInputConfigBackgroundMusics struct {
	// This parameter is required.
	//
	// example:
	//
	// oss://default/bucket-name/filepath/video.mp3
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// fileKey
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitSmartClipTaskRequestInputConfigBackgroundMusics) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmartClipTaskRequestInputConfigBackgroundMusics) GoString() string {
	return s.String()
}

func (s *SubmitSmartClipTaskRequestInputConfigBackgroundMusics) GetId() *string {
	return s.Id
}

func (s *SubmitSmartClipTaskRequestInputConfigBackgroundMusics) GetType() *string {
	return s.Type
}

func (s *SubmitSmartClipTaskRequestInputConfigBackgroundMusics) SetId(v string) *SubmitSmartClipTaskRequestInputConfigBackgroundMusics {
	s.Id = &v
	return s
}

func (s *SubmitSmartClipTaskRequestInputConfigBackgroundMusics) SetType(v string) *SubmitSmartClipTaskRequestInputConfigBackgroundMusics {
	s.Type = &v
	return s
}

func (s *SubmitSmartClipTaskRequestInputConfigBackgroundMusics) Validate() error {
	return dara.Validate(s)
}

type SubmitSmartClipTaskRequestInputConfigStickers struct {
	// This parameter is required.
	//
	// example:
	//
	// 0.5
	Height *float64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// This parameter is required.
	StickerId *SubmitSmartClipTaskRequestInputConfigStickersStickerId `json:"StickerId,omitempty" xml:"StickerId,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 0.5
	Width *float64 `json:"Width,omitempty" xml:"Width,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.5
	X *float64 `json:"X,omitempty" xml:"X,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.5
	Y *float64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s SubmitSmartClipTaskRequestInputConfigStickers) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmartClipTaskRequestInputConfigStickers) GoString() string {
	return s.String()
}

func (s *SubmitSmartClipTaskRequestInputConfigStickers) GetHeight() *float64 {
	return s.Height
}

func (s *SubmitSmartClipTaskRequestInputConfigStickers) GetStickerId() *SubmitSmartClipTaskRequestInputConfigStickersStickerId {
	return s.StickerId
}

func (s *SubmitSmartClipTaskRequestInputConfigStickers) GetWidth() *float64 {
	return s.Width
}

func (s *SubmitSmartClipTaskRequestInputConfigStickers) GetX() *float64 {
	return s.X
}

func (s *SubmitSmartClipTaskRequestInputConfigStickers) GetY() *float64 {
	return s.Y
}

func (s *SubmitSmartClipTaskRequestInputConfigStickers) SetHeight(v float64) *SubmitSmartClipTaskRequestInputConfigStickers {
	s.Height = &v
	return s
}

func (s *SubmitSmartClipTaskRequestInputConfigStickers) SetStickerId(v *SubmitSmartClipTaskRequestInputConfigStickersStickerId) *SubmitSmartClipTaskRequestInputConfigStickers {
	s.StickerId = v
	return s
}

func (s *SubmitSmartClipTaskRequestInputConfigStickers) SetWidth(v float64) *SubmitSmartClipTaskRequestInputConfigStickers {
	s.Width = &v
	return s
}

func (s *SubmitSmartClipTaskRequestInputConfigStickers) SetX(v float64) *SubmitSmartClipTaskRequestInputConfigStickers {
	s.X = &v
	return s
}

func (s *SubmitSmartClipTaskRequestInputConfigStickers) SetY(v float64) *SubmitSmartClipTaskRequestInputConfigStickers {
	s.Y = &v
	return s
}

func (s *SubmitSmartClipTaskRequestInputConfigStickers) Validate() error {
	return dara.Validate(s)
}

type SubmitSmartClipTaskRequestInputConfigStickersStickerId struct {
	// This parameter is required.
	//
	// example:
	//
	// oss://default/bucket-name/filepath/sticker.png
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// fileKey
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitSmartClipTaskRequestInputConfigStickersStickerId) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmartClipTaskRequestInputConfigStickersStickerId) GoString() string {
	return s.String()
}

func (s *SubmitSmartClipTaskRequestInputConfigStickersStickerId) GetId() *string {
	return s.Id
}

func (s *SubmitSmartClipTaskRequestInputConfigStickersStickerId) GetType() *string {
	return s.Type
}

func (s *SubmitSmartClipTaskRequestInputConfigStickersStickerId) SetId(v string) *SubmitSmartClipTaskRequestInputConfigStickersStickerId {
	s.Id = &v
	return s
}

func (s *SubmitSmartClipTaskRequestInputConfigStickersStickerId) SetType(v string) *SubmitSmartClipTaskRequestInputConfigStickersStickerId {
	s.Type = &v
	return s
}

func (s *SubmitSmartClipTaskRequestInputConfigStickersStickerId) Validate() error {
	return dara.Validate(s)
}

type SubmitSmartClipTaskRequestInputConfigVideoIds struct {
	// This parameter is required.
	//
	// example:
	//
	// oss://default/bucket-name/filepath/video.mp4
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// fileKey
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitSmartClipTaskRequestInputConfigVideoIds) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmartClipTaskRequestInputConfigVideoIds) GoString() string {
	return s.String()
}

func (s *SubmitSmartClipTaskRequestInputConfigVideoIds) GetId() *string {
	return s.Id
}

func (s *SubmitSmartClipTaskRequestInputConfigVideoIds) GetType() *string {
	return s.Type
}

func (s *SubmitSmartClipTaskRequestInputConfigVideoIds) SetId(v string) *SubmitSmartClipTaskRequestInputConfigVideoIds {
	s.Id = &v
	return s
}

func (s *SubmitSmartClipTaskRequestInputConfigVideoIds) SetType(v string) *SubmitSmartClipTaskRequestInputConfigVideoIds {
	s.Type = &v
	return s
}

func (s *SubmitSmartClipTaskRequestInputConfigVideoIds) Validate() error {
	return dara.Validate(s)
}

type SubmitSmartClipTaskRequestOutputConfig struct {
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// test_{index}.mp4
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 1080
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 120
	MaxDuration *int32 `json:"MaxDuration,omitempty" xml:"MaxDuration,omitempty"`
	// example:
	//
	// true
	SaveToGeneratedContent *bool `json:"SaveToGeneratedContent,omitempty" xml:"SaveToGeneratedContent,omitempty"`
	// example:
	//
	// 1920
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s SubmitSmartClipTaskRequestOutputConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmartClipTaskRequestOutputConfig) GoString() string {
	return s.String()
}

func (s *SubmitSmartClipTaskRequestOutputConfig) GetCount() *int32 {
	return s.Count
}

func (s *SubmitSmartClipTaskRequestOutputConfig) GetFileName() *string {
	return s.FileName
}

func (s *SubmitSmartClipTaskRequestOutputConfig) GetHeight() *int32 {
	return s.Height
}

func (s *SubmitSmartClipTaskRequestOutputConfig) GetMaxDuration() *int32 {
	return s.MaxDuration
}

func (s *SubmitSmartClipTaskRequestOutputConfig) GetSaveToGeneratedContent() *bool {
	return s.SaveToGeneratedContent
}

func (s *SubmitSmartClipTaskRequestOutputConfig) GetWidth() *int32 {
	return s.Width
}

func (s *SubmitSmartClipTaskRequestOutputConfig) SetCount(v int32) *SubmitSmartClipTaskRequestOutputConfig {
	s.Count = &v
	return s
}

func (s *SubmitSmartClipTaskRequestOutputConfig) SetFileName(v string) *SubmitSmartClipTaskRequestOutputConfig {
	s.FileName = &v
	return s
}

func (s *SubmitSmartClipTaskRequestOutputConfig) SetHeight(v int32) *SubmitSmartClipTaskRequestOutputConfig {
	s.Height = &v
	return s
}

func (s *SubmitSmartClipTaskRequestOutputConfig) SetMaxDuration(v int32) *SubmitSmartClipTaskRequestOutputConfig {
	s.MaxDuration = &v
	return s
}

func (s *SubmitSmartClipTaskRequestOutputConfig) SetSaveToGeneratedContent(v bool) *SubmitSmartClipTaskRequestOutputConfig {
	s.SaveToGeneratedContent = &v
	return s
}

func (s *SubmitSmartClipTaskRequestOutputConfig) SetWidth(v int32) *SubmitSmartClipTaskRequestOutputConfig {
	s.Width = &v
	return s
}

func (s *SubmitSmartClipTaskRequestOutputConfig) Validate() error {
	return dara.Validate(s)
}
