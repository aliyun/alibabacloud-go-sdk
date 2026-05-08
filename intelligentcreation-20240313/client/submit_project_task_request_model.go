// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitProjectTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrames(v []*SubmitProjectTaskRequestFrames) *SubmitProjectTaskRequest
	GetFrames() []*SubmitProjectTaskRequestFrames
	SetScaleType(v string) *SubmitProjectTaskRequest
	GetScaleType() *string
	SetSubtitleTag(v int32) *SubmitProjectTaskRequest
	GetSubtitleTag() *int32
	SetTransparentBackground(v int32) *SubmitProjectTaskRequest
	GetTransparentBackground() *int32
}

type SubmitProjectTaskRequest struct {
	// frame
	Frames []*SubmitProjectTaskRequestFrames `json:"frames,omitempty" xml:"frames,omitempty" type:"Repeated"`
	// example:
	//
	// 9:16
	ScaleType *string `json:"scaleType,omitempty" xml:"scaleType,omitempty"`
	// example:
	//
	// 1
	SubtitleTag           *int32 `json:"subtitleTag,omitempty" xml:"subtitleTag,omitempty"`
	TransparentBackground *int32 `json:"transparentBackground,omitempty" xml:"transparentBackground,omitempty"`
}

func (s SubmitProjectTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitProjectTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitProjectTaskRequest) GetFrames() []*SubmitProjectTaskRequestFrames {
	return s.Frames
}

func (s *SubmitProjectTaskRequest) GetScaleType() *string {
	return s.ScaleType
}

func (s *SubmitProjectTaskRequest) GetSubtitleTag() *int32 {
	return s.SubtitleTag
}

func (s *SubmitProjectTaskRequest) GetTransparentBackground() *int32 {
	return s.TransparentBackground
}

func (s *SubmitProjectTaskRequest) SetFrames(v []*SubmitProjectTaskRequestFrames) *SubmitProjectTaskRequest {
	s.Frames = v
	return s
}

func (s *SubmitProjectTaskRequest) SetScaleType(v string) *SubmitProjectTaskRequest {
	s.ScaleType = &v
	return s
}

func (s *SubmitProjectTaskRequest) SetSubtitleTag(v int32) *SubmitProjectTaskRequest {
	s.SubtitleTag = &v
	return s
}

func (s *SubmitProjectTaskRequest) SetTransparentBackground(v int32) *SubmitProjectTaskRequest {
	s.TransparentBackground = &v
	return s
}

func (s *SubmitProjectTaskRequest) Validate() error {
	if s.Frames != nil {
		for _, item := range s.Frames {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubmitProjectTaskRequestFrames struct {
	// example:
	//
	// 1
	Index       *int32                                     `json:"index,omitempty" xml:"index,omitempty"`
	Layers      []*SubmitProjectTaskRequestFramesLayers    `json:"layers,omitempty" xml:"layers,omitempty" type:"Repeated"`
	Subtitle    *SubmitProjectTaskRequestFramesSubtitle    `json:"subtitle,omitempty" xml:"subtitle,omitempty" type:"Struct"`
	VideoScript *SubmitProjectTaskRequestFramesVideoScript `json:"videoScript,omitempty" xml:"videoScript,omitempty" type:"Struct"`
}

func (s SubmitProjectTaskRequestFrames) String() string {
	return dara.Prettify(s)
}

func (s SubmitProjectTaskRequestFrames) GoString() string {
	return s.String()
}

func (s *SubmitProjectTaskRequestFrames) GetIndex() *int32 {
	return s.Index
}

func (s *SubmitProjectTaskRequestFrames) GetLayers() []*SubmitProjectTaskRequestFramesLayers {
	return s.Layers
}

func (s *SubmitProjectTaskRequestFrames) GetSubtitle() *SubmitProjectTaskRequestFramesSubtitle {
	return s.Subtitle
}

func (s *SubmitProjectTaskRequestFrames) GetVideoScript() *SubmitProjectTaskRequestFramesVideoScript {
	return s.VideoScript
}

func (s *SubmitProjectTaskRequestFrames) SetIndex(v int32) *SubmitProjectTaskRequestFrames {
	s.Index = &v
	return s
}

func (s *SubmitProjectTaskRequestFrames) SetLayers(v []*SubmitProjectTaskRequestFramesLayers) *SubmitProjectTaskRequestFrames {
	s.Layers = v
	return s
}

func (s *SubmitProjectTaskRequestFrames) SetSubtitle(v *SubmitProjectTaskRequestFramesSubtitle) *SubmitProjectTaskRequestFrames {
	s.Subtitle = v
	return s
}

func (s *SubmitProjectTaskRequestFrames) SetVideoScript(v *SubmitProjectTaskRequestFramesVideoScript) *SubmitProjectTaskRequestFrames {
	s.VideoScript = v
	return s
}

func (s *SubmitProjectTaskRequestFrames) Validate() error {
	if s.Layers != nil {
		for _, item := range s.Layers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Subtitle != nil {
		if err := s.Subtitle.Validate(); err != nil {
			return err
		}
	}
	if s.VideoScript != nil {
		if err := s.VideoScript.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitProjectTaskRequestFramesLayers struct {
	// example:
	//
	// 222
	Height *int32 `json:"height,omitempty" xml:"height,omitempty"`
	// example:
	//
	// 1
	Index    *int32                                        `json:"index,omitempty" xml:"index,omitempty"`
	Material *SubmitProjectTaskRequestFramesLayersMaterial `json:"material,omitempty" xml:"material,omitempty" type:"Struct"`
	// example:
	//
	// 11
	PositionX *int32 `json:"positionX,omitempty" xml:"positionX,omitempty"`
	// example:
	//
	// 22
	PositionY *int32 `json:"positionY,omitempty" xml:"positionY,omitempty"`
	// example:
	//
	// ANCHOR
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 111
	Width *int32 `json:"width,omitempty" xml:"width,omitempty"`
}

func (s SubmitProjectTaskRequestFramesLayers) String() string {
	return dara.Prettify(s)
}

func (s SubmitProjectTaskRequestFramesLayers) GoString() string {
	return s.String()
}

func (s *SubmitProjectTaskRequestFramesLayers) GetHeight() *int32 {
	return s.Height
}

func (s *SubmitProjectTaskRequestFramesLayers) GetIndex() *int32 {
	return s.Index
}

func (s *SubmitProjectTaskRequestFramesLayers) GetMaterial() *SubmitProjectTaskRequestFramesLayersMaterial {
	return s.Material
}

func (s *SubmitProjectTaskRequestFramesLayers) GetPositionX() *int32 {
	return s.PositionX
}

func (s *SubmitProjectTaskRequestFramesLayers) GetPositionY() *int32 {
	return s.PositionY
}

func (s *SubmitProjectTaskRequestFramesLayers) GetType() *string {
	return s.Type
}

func (s *SubmitProjectTaskRequestFramesLayers) GetWidth() *int32 {
	return s.Width
}

func (s *SubmitProjectTaskRequestFramesLayers) SetHeight(v int32) *SubmitProjectTaskRequestFramesLayers {
	s.Height = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesLayers) SetIndex(v int32) *SubmitProjectTaskRequestFramesLayers {
	s.Index = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesLayers) SetMaterial(v *SubmitProjectTaskRequestFramesLayersMaterial) *SubmitProjectTaskRequestFramesLayers {
	s.Material = v
	return s
}

func (s *SubmitProjectTaskRequestFramesLayers) SetPositionX(v int32) *SubmitProjectTaskRequestFramesLayers {
	s.PositionX = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesLayers) SetPositionY(v int32) *SubmitProjectTaskRequestFramesLayers {
	s.PositionY = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesLayers) SetType(v string) *SubmitProjectTaskRequestFramesLayers {
	s.Type = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesLayers) SetWidth(v int32) *SubmitProjectTaskRequestFramesLayers {
	s.Width = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesLayers) Validate() error {
	if s.Material != nil {
		if err := s.Material.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitProjectTaskRequestFramesLayersMaterial struct {
	AnchorStyleLevel *string `json:"anchorStyleLevel,omitempty" xml:"anchorStyleLevel,omitempty"`
	// example:
	//
	// video/mp4
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// example:
	//
	// 38863
	Id    *string                                           `json:"id,omitempty" xml:"id,omitempty"`
	Mask  *SubmitProjectTaskRequestFramesLayersMaterialMask `json:"mask,omitempty" xml:"mask,omitempty" type:"Struct"`
	Speed *string                                           `json:"speed,omitempty" xml:"speed,omitempty"`
	// example:
	//
	// https://xxx
	Url    *string `json:"url,omitempty" xml:"url,omitempty"`
	Volume *int32  `json:"volume,omitempty" xml:"volume,omitempty"`
}

func (s SubmitProjectTaskRequestFramesLayersMaterial) String() string {
	return dara.Prettify(s)
}

func (s SubmitProjectTaskRequestFramesLayersMaterial) GoString() string {
	return s.String()
}

func (s *SubmitProjectTaskRequestFramesLayersMaterial) GetAnchorStyleLevel() *string {
	return s.AnchorStyleLevel
}

func (s *SubmitProjectTaskRequestFramesLayersMaterial) GetFormat() *string {
	return s.Format
}

func (s *SubmitProjectTaskRequestFramesLayersMaterial) GetId() *string {
	return s.Id
}

func (s *SubmitProjectTaskRequestFramesLayersMaterial) GetMask() *SubmitProjectTaskRequestFramesLayersMaterialMask {
	return s.Mask
}

func (s *SubmitProjectTaskRequestFramesLayersMaterial) GetSpeed() *string {
	return s.Speed
}

func (s *SubmitProjectTaskRequestFramesLayersMaterial) GetUrl() *string {
	return s.Url
}

func (s *SubmitProjectTaskRequestFramesLayersMaterial) GetVolume() *int32 {
	return s.Volume
}

func (s *SubmitProjectTaskRequestFramesLayersMaterial) SetAnchorStyleLevel(v string) *SubmitProjectTaskRequestFramesLayersMaterial {
	s.AnchorStyleLevel = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesLayersMaterial) SetFormat(v string) *SubmitProjectTaskRequestFramesLayersMaterial {
	s.Format = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesLayersMaterial) SetId(v string) *SubmitProjectTaskRequestFramesLayersMaterial {
	s.Id = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesLayersMaterial) SetMask(v *SubmitProjectTaskRequestFramesLayersMaterialMask) *SubmitProjectTaskRequestFramesLayersMaterial {
	s.Mask = v
	return s
}

func (s *SubmitProjectTaskRequestFramesLayersMaterial) SetSpeed(v string) *SubmitProjectTaskRequestFramesLayersMaterial {
	s.Speed = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesLayersMaterial) SetUrl(v string) *SubmitProjectTaskRequestFramesLayersMaterial {
	s.Url = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesLayersMaterial) SetVolume(v int32) *SubmitProjectTaskRequestFramesLayersMaterial {
	s.Volume = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesLayersMaterial) Validate() error {
	if s.Mask != nil {
		if err := s.Mask.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitProjectTaskRequestFramesLayersMaterialMask struct {
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s SubmitProjectTaskRequestFramesLayersMaterialMask) String() string {
	return dara.Prettify(s)
}

func (s SubmitProjectTaskRequestFramesLayersMaterialMask) GoString() string {
	return s.String()
}

func (s *SubmitProjectTaskRequestFramesLayersMaterialMask) GetUrl() *string {
	return s.Url
}

func (s *SubmitProjectTaskRequestFramesLayersMaterialMask) SetUrl(v string) *SubmitProjectTaskRequestFramesLayersMaterialMask {
	s.Url = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesLayersMaterialMask) Validate() error {
	return dara.Validate(s)
}

type SubmitProjectTaskRequestFramesSubtitle struct {
	// example:
	//
	// BottomLeft
	Alignment *string `json:"alignment,omitempty" xml:"alignment,omitempty"`
	// example:
	//
	// #ffffff
	BackgroundColor *string `json:"backgroundColor,omitempty" xml:"backgroundColor,omitempty"`
	// example:
	//
	// SimSun
	Font *string `json:"font,omitempty" xml:"font,omitempty"`
	// example:
	//
	// #ffffff
	FontColor *string `json:"fontColor,omitempty" xml:"fontColor,omitempty"`
	// example:
	//
	// 32
	FontSize *int32 `json:"fontSize,omitempty" xml:"fontSize,omitempty"`
	// example:
	//
	// 11
	MaxCharLength *int32 `json:"maxCharLength,omitempty" xml:"maxCharLength,omitempty"`
	// example:
	//
	// 2
	PositionX *int32 `json:"positionX,omitempty" xml:"positionX,omitempty"`
	// example:
	//
	// 1
	PositionY *int32 `json:"positionY,omitempty" xml:"positionY,omitempty"`
	// example:
	//
	// 22
	TextHeight *int32 `json:"textHeight,omitempty" xml:"textHeight,omitempty"`
	// example:
	//
	// 11
	TextWidth *int32 `json:"textWidth,omitempty" xml:"textWidth,omitempty"`
}

func (s SubmitProjectTaskRequestFramesSubtitle) String() string {
	return dara.Prettify(s)
}

func (s SubmitProjectTaskRequestFramesSubtitle) GoString() string {
	return s.String()
}

func (s *SubmitProjectTaskRequestFramesSubtitle) GetAlignment() *string {
	return s.Alignment
}

func (s *SubmitProjectTaskRequestFramesSubtitle) GetBackgroundColor() *string {
	return s.BackgroundColor
}

func (s *SubmitProjectTaskRequestFramesSubtitle) GetFont() *string {
	return s.Font
}

func (s *SubmitProjectTaskRequestFramesSubtitle) GetFontColor() *string {
	return s.FontColor
}

func (s *SubmitProjectTaskRequestFramesSubtitle) GetFontSize() *int32 {
	return s.FontSize
}

func (s *SubmitProjectTaskRequestFramesSubtitle) GetMaxCharLength() *int32 {
	return s.MaxCharLength
}

func (s *SubmitProjectTaskRequestFramesSubtitle) GetPositionX() *int32 {
	return s.PositionX
}

func (s *SubmitProjectTaskRequestFramesSubtitle) GetPositionY() *int32 {
	return s.PositionY
}

func (s *SubmitProjectTaskRequestFramesSubtitle) GetTextHeight() *int32 {
	return s.TextHeight
}

func (s *SubmitProjectTaskRequestFramesSubtitle) GetTextWidth() *int32 {
	return s.TextWidth
}

func (s *SubmitProjectTaskRequestFramesSubtitle) SetAlignment(v string) *SubmitProjectTaskRequestFramesSubtitle {
	s.Alignment = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesSubtitle) SetBackgroundColor(v string) *SubmitProjectTaskRequestFramesSubtitle {
	s.BackgroundColor = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesSubtitle) SetFont(v string) *SubmitProjectTaskRequestFramesSubtitle {
	s.Font = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesSubtitle) SetFontColor(v string) *SubmitProjectTaskRequestFramesSubtitle {
	s.FontColor = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesSubtitle) SetFontSize(v int32) *SubmitProjectTaskRequestFramesSubtitle {
	s.FontSize = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesSubtitle) SetMaxCharLength(v int32) *SubmitProjectTaskRequestFramesSubtitle {
	s.MaxCharLength = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesSubtitle) SetPositionX(v int32) *SubmitProjectTaskRequestFramesSubtitle {
	s.PositionX = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesSubtitle) SetPositionY(v int32) *SubmitProjectTaskRequestFramesSubtitle {
	s.PositionY = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesSubtitle) SetTextHeight(v int32) *SubmitProjectTaskRequestFramesSubtitle {
	s.TextHeight = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesSubtitle) SetTextWidth(v int32) *SubmitProjectTaskRequestFramesSubtitle {
	s.TextWidth = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesSubtitle) Validate() error {
	return dara.Validate(s)
}

type SubmitProjectTaskRequestFramesVideoScript struct {
	// example:
	//
	// https://xxx
	AudioUrl   *string `json:"audioUrl,omitempty" xml:"audioUrl,omitempty"`
	Emotion    *string `json:"emotion,omitempty" xml:"emotion,omitempty"`
	PitchRate  *string `json:"pitchRate,omitempty" xml:"pitchRate,omitempty"`
	SpeechOpen *bool   `json:"speechOpen,omitempty" xml:"speechOpen,omitempty"`
	// example:
	//
	// 2.0
	SpeedRate   *string `json:"speedRate,omitempty" xml:"speedRate,omitempty"`
	TextContent *string `json:"textContent,omitempty" xml:"textContent,omitempty"`
	// example:
	//
	// TEXT
	Type          *string `json:"type,omitempty" xml:"type,omitempty"`
	VoiceLanguage *string `json:"voiceLanguage,omitempty" xml:"voiceLanguage,omitempty"`
	// example:
	//
	// 11
	VoiceTemplateId *int64 `json:"voiceTemplateId,omitempty" xml:"voiceTemplateId,omitempty"`
	// example:
	//
	// 20
	Volume *int32 `json:"volume,omitempty" xml:"volume,omitempty"`
}

func (s SubmitProjectTaskRequestFramesVideoScript) String() string {
	return dara.Prettify(s)
}

func (s SubmitProjectTaskRequestFramesVideoScript) GoString() string {
	return s.String()
}

func (s *SubmitProjectTaskRequestFramesVideoScript) GetAudioUrl() *string {
	return s.AudioUrl
}

func (s *SubmitProjectTaskRequestFramesVideoScript) GetEmotion() *string {
	return s.Emotion
}

func (s *SubmitProjectTaskRequestFramesVideoScript) GetPitchRate() *string {
	return s.PitchRate
}

func (s *SubmitProjectTaskRequestFramesVideoScript) GetSpeechOpen() *bool {
	return s.SpeechOpen
}

func (s *SubmitProjectTaskRequestFramesVideoScript) GetSpeedRate() *string {
	return s.SpeedRate
}

func (s *SubmitProjectTaskRequestFramesVideoScript) GetTextContent() *string {
	return s.TextContent
}

func (s *SubmitProjectTaskRequestFramesVideoScript) GetType() *string {
	return s.Type
}

func (s *SubmitProjectTaskRequestFramesVideoScript) GetVoiceLanguage() *string {
	return s.VoiceLanguage
}

func (s *SubmitProjectTaskRequestFramesVideoScript) GetVoiceTemplateId() *int64 {
	return s.VoiceTemplateId
}

func (s *SubmitProjectTaskRequestFramesVideoScript) GetVolume() *int32 {
	return s.Volume
}

func (s *SubmitProjectTaskRequestFramesVideoScript) SetAudioUrl(v string) *SubmitProjectTaskRequestFramesVideoScript {
	s.AudioUrl = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesVideoScript) SetEmotion(v string) *SubmitProjectTaskRequestFramesVideoScript {
	s.Emotion = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesVideoScript) SetPitchRate(v string) *SubmitProjectTaskRequestFramesVideoScript {
	s.PitchRate = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesVideoScript) SetSpeechOpen(v bool) *SubmitProjectTaskRequestFramesVideoScript {
	s.SpeechOpen = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesVideoScript) SetSpeedRate(v string) *SubmitProjectTaskRequestFramesVideoScript {
	s.SpeedRate = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesVideoScript) SetTextContent(v string) *SubmitProjectTaskRequestFramesVideoScript {
	s.TextContent = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesVideoScript) SetType(v string) *SubmitProjectTaskRequestFramesVideoScript {
	s.Type = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesVideoScript) SetVoiceLanguage(v string) *SubmitProjectTaskRequestFramesVideoScript {
	s.VoiceLanguage = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesVideoScript) SetVoiceTemplateId(v int64) *SubmitProjectTaskRequestFramesVideoScript {
	s.VoiceTemplateId = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesVideoScript) SetVolume(v int32) *SubmitProjectTaskRequestFramesVideoScript {
	s.Volume = &v
	return s
}

func (s *SubmitProjectTaskRequestFramesVideoScript) Validate() error {
	return dara.Validate(s)
}
