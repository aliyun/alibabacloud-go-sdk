// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTargetVideo interface {
	dara.Model
	String() string
	GoString() string
	SetDisableVideo(v bool) *TargetVideo
	GetDisableVideo() *bool
	SetFilterVideo(v *TargetVideoFilterVideo) *TargetVideo
	GetFilterVideo() *TargetVideoFilterVideo
	SetStream(v []*int32) *TargetVideo
	GetStream() []*int32
	SetTranscodeVideo(v *TargetVideoTranscodeVideo) *TargetVideo
	GetTranscodeVideo() *TargetVideoTranscodeVideo
}

type TargetVideo struct {
	DisableVideo   *bool                      `json:"DisableVideo,omitempty" xml:"DisableVideo,omitempty"`
	FilterVideo    *TargetVideoFilterVideo    `json:"FilterVideo,omitempty" xml:"FilterVideo,omitempty" type:"Struct"`
	Stream         []*int32                   `json:"Stream,omitempty" xml:"Stream,omitempty" type:"Repeated"`
	TranscodeVideo *TargetVideoTranscodeVideo `json:"TranscodeVideo,omitempty" xml:"TranscodeVideo,omitempty" type:"Struct"`
}

func (s TargetVideo) String() string {
	return dara.Prettify(s)
}

func (s TargetVideo) GoString() string {
	return s.String()
}

func (s *TargetVideo) GetDisableVideo() *bool {
	return s.DisableVideo
}

func (s *TargetVideo) GetFilterVideo() *TargetVideoFilterVideo {
	return s.FilterVideo
}

func (s *TargetVideo) GetStream() []*int32 {
	return s.Stream
}

func (s *TargetVideo) GetTranscodeVideo() *TargetVideoTranscodeVideo {
	return s.TranscodeVideo
}

func (s *TargetVideo) SetDisableVideo(v bool) *TargetVideo {
	s.DisableVideo = &v
	return s
}

func (s *TargetVideo) SetFilterVideo(v *TargetVideoFilterVideo) *TargetVideo {
	s.FilterVideo = v
	return s
}

func (s *TargetVideo) SetStream(v []*int32) *TargetVideo {
	s.Stream = v
	return s
}

func (s *TargetVideo) SetTranscodeVideo(v *TargetVideoTranscodeVideo) *TargetVideo {
	s.TranscodeVideo = v
	return s
}

func (s *TargetVideo) Validate() error {
	if s.FilterVideo != nil {
		if err := s.FilterVideo.Validate(); err != nil {
			return err
		}
	}
	if s.TranscodeVideo != nil {
		if err := s.TranscodeVideo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TargetVideoFilterVideo struct {
	Delogos         []*TargetVideoFilterVideoDelogos       `json:"Delogos,omitempty" xml:"Delogos,omitempty" type:"Repeated"`
	Desensitization *TargetVideoFilterVideoDesensitization `json:"Desensitization,omitempty" xml:"Desensitization,omitempty" type:"Struct"`
	Speed           *float32                               `json:"Speed,omitempty" xml:"Speed,omitempty"`
	Watermarks      []*TargetVideoFilterVideoWatermarks    `json:"Watermarks,omitempty" xml:"Watermarks,omitempty" type:"Repeated"`
}

func (s TargetVideoFilterVideo) String() string {
	return dara.Prettify(s)
}

func (s TargetVideoFilterVideo) GoString() string {
	return s.String()
}

func (s *TargetVideoFilterVideo) GetDelogos() []*TargetVideoFilterVideoDelogos {
	return s.Delogos
}

func (s *TargetVideoFilterVideo) GetDesensitization() *TargetVideoFilterVideoDesensitization {
	return s.Desensitization
}

func (s *TargetVideoFilterVideo) GetSpeed() *float32 {
	return s.Speed
}

func (s *TargetVideoFilterVideo) GetWatermarks() []*TargetVideoFilterVideoWatermarks {
	return s.Watermarks
}

func (s *TargetVideoFilterVideo) SetDelogos(v []*TargetVideoFilterVideoDelogos) *TargetVideoFilterVideo {
	s.Delogos = v
	return s
}

func (s *TargetVideoFilterVideo) SetDesensitization(v *TargetVideoFilterVideoDesensitization) *TargetVideoFilterVideo {
	s.Desensitization = v
	return s
}

func (s *TargetVideoFilterVideo) SetSpeed(v float32) *TargetVideoFilterVideo {
	s.Speed = &v
	return s
}

func (s *TargetVideoFilterVideo) SetWatermarks(v []*TargetVideoFilterVideoWatermarks) *TargetVideoFilterVideo {
	s.Watermarks = v
	return s
}

func (s *TargetVideoFilterVideo) Validate() error {
	if s.Delogos != nil {
		for _, item := range s.Delogos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Desensitization != nil {
		if err := s.Desensitization.Validate(); err != nil {
			return err
		}
	}
	if s.Watermarks != nil {
		for _, item := range s.Watermarks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TargetVideoFilterVideoDelogos struct {
	Duration  *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Dx        *float32 `json:"Dx,omitempty" xml:"Dx,omitempty"`
	Dy        *float32 `json:"Dy,omitempty" xml:"Dy,omitempty"`
	Height    *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	ReferPos  *string  `json:"ReferPos,omitempty" xml:"ReferPos,omitempty"`
	StartTime *float64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Width     *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s TargetVideoFilterVideoDelogos) String() string {
	return dara.Prettify(s)
}

func (s TargetVideoFilterVideoDelogos) GoString() string {
	return s.String()
}

func (s *TargetVideoFilterVideoDelogos) GetDuration() *float64 {
	return s.Duration
}

func (s *TargetVideoFilterVideoDelogos) GetDx() *float32 {
	return s.Dx
}

func (s *TargetVideoFilterVideoDelogos) GetDy() *float32 {
	return s.Dy
}

func (s *TargetVideoFilterVideoDelogos) GetHeight() *float32 {
	return s.Height
}

func (s *TargetVideoFilterVideoDelogos) GetReferPos() *string {
	return s.ReferPos
}

func (s *TargetVideoFilterVideoDelogos) GetStartTime() *float64 {
	return s.StartTime
}

func (s *TargetVideoFilterVideoDelogos) GetWidth() *float32 {
	return s.Width
}

func (s *TargetVideoFilterVideoDelogos) SetDuration(v float64) *TargetVideoFilterVideoDelogos {
	s.Duration = &v
	return s
}

func (s *TargetVideoFilterVideoDelogos) SetDx(v float32) *TargetVideoFilterVideoDelogos {
	s.Dx = &v
	return s
}

func (s *TargetVideoFilterVideoDelogos) SetDy(v float32) *TargetVideoFilterVideoDelogos {
	s.Dy = &v
	return s
}

func (s *TargetVideoFilterVideoDelogos) SetHeight(v float32) *TargetVideoFilterVideoDelogos {
	s.Height = &v
	return s
}

func (s *TargetVideoFilterVideoDelogos) SetReferPos(v string) *TargetVideoFilterVideoDelogos {
	s.ReferPos = &v
	return s
}

func (s *TargetVideoFilterVideoDelogos) SetStartTime(v float64) *TargetVideoFilterVideoDelogos {
	s.StartTime = &v
	return s
}

func (s *TargetVideoFilterVideoDelogos) SetWidth(v float32) *TargetVideoFilterVideoDelogos {
	s.Width = &v
	return s
}

func (s *TargetVideoFilterVideoDelogos) Validate() error {
	return dara.Validate(s)
}

type TargetVideoFilterVideoDesensitization struct {
	Face         *TargetVideoFilterVideoDesensitizationFace         `json:"Face,omitempty" xml:"Face,omitempty" type:"Struct"`
	LicensePlate *TargetVideoFilterVideoDesensitizationLicensePlate `json:"LicensePlate,omitempty" xml:"LicensePlate,omitempty" type:"Struct"`
}

func (s TargetVideoFilterVideoDesensitization) String() string {
	return dara.Prettify(s)
}

func (s TargetVideoFilterVideoDesensitization) GoString() string {
	return s.String()
}

func (s *TargetVideoFilterVideoDesensitization) GetFace() *TargetVideoFilterVideoDesensitizationFace {
	return s.Face
}

func (s *TargetVideoFilterVideoDesensitization) GetLicensePlate() *TargetVideoFilterVideoDesensitizationLicensePlate {
	return s.LicensePlate
}

func (s *TargetVideoFilterVideoDesensitization) SetFace(v *TargetVideoFilterVideoDesensitizationFace) *TargetVideoFilterVideoDesensitization {
	s.Face = v
	return s
}

func (s *TargetVideoFilterVideoDesensitization) SetLicensePlate(v *TargetVideoFilterVideoDesensitizationLicensePlate) *TargetVideoFilterVideoDesensitization {
	s.LicensePlate = v
	return s
}

func (s *TargetVideoFilterVideoDesensitization) Validate() error {
	if s.Face != nil {
		if err := s.Face.Validate(); err != nil {
			return err
		}
	}
	if s.LicensePlate != nil {
		if err := s.LicensePlate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TargetVideoFilterVideoDesensitizationFace struct {
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	MinSize    *int32   `json:"MinSize,omitempty" xml:"MinSize,omitempty"`
}

func (s TargetVideoFilterVideoDesensitizationFace) String() string {
	return dara.Prettify(s)
}

func (s TargetVideoFilterVideoDesensitizationFace) GoString() string {
	return s.String()
}

func (s *TargetVideoFilterVideoDesensitizationFace) GetConfidence() *float32 {
	return s.Confidence
}

func (s *TargetVideoFilterVideoDesensitizationFace) GetMinSize() *int32 {
	return s.MinSize
}

func (s *TargetVideoFilterVideoDesensitizationFace) SetConfidence(v float32) *TargetVideoFilterVideoDesensitizationFace {
	s.Confidence = &v
	return s
}

func (s *TargetVideoFilterVideoDesensitizationFace) SetMinSize(v int32) *TargetVideoFilterVideoDesensitizationFace {
	s.MinSize = &v
	return s
}

func (s *TargetVideoFilterVideoDesensitizationFace) Validate() error {
	return dara.Validate(s)
}

type TargetVideoFilterVideoDesensitizationLicensePlate struct {
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	MinSize    *int32   `json:"MinSize,omitempty" xml:"MinSize,omitempty"`
}

func (s TargetVideoFilterVideoDesensitizationLicensePlate) String() string {
	return dara.Prettify(s)
}

func (s TargetVideoFilterVideoDesensitizationLicensePlate) GoString() string {
	return s.String()
}

func (s *TargetVideoFilterVideoDesensitizationLicensePlate) GetConfidence() *float32 {
	return s.Confidence
}

func (s *TargetVideoFilterVideoDesensitizationLicensePlate) GetMinSize() *int32 {
	return s.MinSize
}

func (s *TargetVideoFilterVideoDesensitizationLicensePlate) SetConfidence(v float32) *TargetVideoFilterVideoDesensitizationLicensePlate {
	s.Confidence = &v
	return s
}

func (s *TargetVideoFilterVideoDesensitizationLicensePlate) SetMinSize(v int32) *TargetVideoFilterVideoDesensitizationLicensePlate {
	s.MinSize = &v
	return s
}

func (s *TargetVideoFilterVideoDesensitizationLicensePlate) Validate() error {
	return dara.Validate(s)
}

type TargetVideoFilterVideoWatermarks struct {
	BorderColor *string  `json:"BorderColor,omitempty" xml:"BorderColor,omitempty"`
	BorderWidth *int32   `json:"BorderWidth,omitempty" xml:"BorderWidth,omitempty"`
	Content     *string  `json:"Content,omitempty" xml:"Content,omitempty"`
	Duration    *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Dx          *float32 `json:"Dx,omitempty" xml:"Dx,omitempty"`
	Dy          *float32 `json:"Dy,omitempty" xml:"Dy,omitempty"`
	FontApha    *float32 `json:"FontApha,omitempty" xml:"FontApha,omitempty"`
	FontColor   *string  `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	FontName    *string  `json:"FontName,omitempty" xml:"FontName,omitempty"`
	FontSize    *int32   `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	Height      *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	ReferPos    *string  `json:"ReferPos,omitempty" xml:"ReferPos,omitempty"`
	StartTime   *float64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Type        *string  `json:"Type,omitempty" xml:"Type,omitempty"`
	URI         *string  `json:"URI,omitempty" xml:"URI,omitempty"`
	Width       *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s TargetVideoFilterVideoWatermarks) String() string {
	return dara.Prettify(s)
}

func (s TargetVideoFilterVideoWatermarks) GoString() string {
	return s.String()
}

func (s *TargetVideoFilterVideoWatermarks) GetBorderColor() *string {
	return s.BorderColor
}

func (s *TargetVideoFilterVideoWatermarks) GetBorderWidth() *int32 {
	return s.BorderWidth
}

func (s *TargetVideoFilterVideoWatermarks) GetContent() *string {
	return s.Content
}

func (s *TargetVideoFilterVideoWatermarks) GetDuration() *float64 {
	return s.Duration
}

func (s *TargetVideoFilterVideoWatermarks) GetDx() *float32 {
	return s.Dx
}

func (s *TargetVideoFilterVideoWatermarks) GetDy() *float32 {
	return s.Dy
}

func (s *TargetVideoFilterVideoWatermarks) GetFontApha() *float32 {
	return s.FontApha
}

func (s *TargetVideoFilterVideoWatermarks) GetFontColor() *string {
	return s.FontColor
}

func (s *TargetVideoFilterVideoWatermarks) GetFontName() *string {
	return s.FontName
}

func (s *TargetVideoFilterVideoWatermarks) GetFontSize() *int32 {
	return s.FontSize
}

func (s *TargetVideoFilterVideoWatermarks) GetHeight() *float32 {
	return s.Height
}

func (s *TargetVideoFilterVideoWatermarks) GetReferPos() *string {
	return s.ReferPos
}

func (s *TargetVideoFilterVideoWatermarks) GetStartTime() *float64 {
	return s.StartTime
}

func (s *TargetVideoFilterVideoWatermarks) GetType() *string {
	return s.Type
}

func (s *TargetVideoFilterVideoWatermarks) GetURI() *string {
	return s.URI
}

func (s *TargetVideoFilterVideoWatermarks) GetWidth() *float32 {
	return s.Width
}

func (s *TargetVideoFilterVideoWatermarks) SetBorderColor(v string) *TargetVideoFilterVideoWatermarks {
	s.BorderColor = &v
	return s
}

func (s *TargetVideoFilterVideoWatermarks) SetBorderWidth(v int32) *TargetVideoFilterVideoWatermarks {
	s.BorderWidth = &v
	return s
}

func (s *TargetVideoFilterVideoWatermarks) SetContent(v string) *TargetVideoFilterVideoWatermarks {
	s.Content = &v
	return s
}

func (s *TargetVideoFilterVideoWatermarks) SetDuration(v float64) *TargetVideoFilterVideoWatermarks {
	s.Duration = &v
	return s
}

func (s *TargetVideoFilterVideoWatermarks) SetDx(v float32) *TargetVideoFilterVideoWatermarks {
	s.Dx = &v
	return s
}

func (s *TargetVideoFilterVideoWatermarks) SetDy(v float32) *TargetVideoFilterVideoWatermarks {
	s.Dy = &v
	return s
}

func (s *TargetVideoFilterVideoWatermarks) SetFontApha(v float32) *TargetVideoFilterVideoWatermarks {
	s.FontApha = &v
	return s
}

func (s *TargetVideoFilterVideoWatermarks) SetFontColor(v string) *TargetVideoFilterVideoWatermarks {
	s.FontColor = &v
	return s
}

func (s *TargetVideoFilterVideoWatermarks) SetFontName(v string) *TargetVideoFilterVideoWatermarks {
	s.FontName = &v
	return s
}

func (s *TargetVideoFilterVideoWatermarks) SetFontSize(v int32) *TargetVideoFilterVideoWatermarks {
	s.FontSize = &v
	return s
}

func (s *TargetVideoFilterVideoWatermarks) SetHeight(v float32) *TargetVideoFilterVideoWatermarks {
	s.Height = &v
	return s
}

func (s *TargetVideoFilterVideoWatermarks) SetReferPos(v string) *TargetVideoFilterVideoWatermarks {
	s.ReferPos = &v
	return s
}

func (s *TargetVideoFilterVideoWatermarks) SetStartTime(v float64) *TargetVideoFilterVideoWatermarks {
	s.StartTime = &v
	return s
}

func (s *TargetVideoFilterVideoWatermarks) SetType(v string) *TargetVideoFilterVideoWatermarks {
	s.Type = &v
	return s
}

func (s *TargetVideoFilterVideoWatermarks) SetURI(v string) *TargetVideoFilterVideoWatermarks {
	s.URI = &v
	return s
}

func (s *TargetVideoFilterVideoWatermarks) SetWidth(v float32) *TargetVideoFilterVideoWatermarks {
	s.Width = &v
	return s
}

func (s *TargetVideoFilterVideoWatermarks) Validate() error {
	return dara.Validate(s)
}

type TargetVideoTranscodeVideo struct {
	AdaptiveResolutionDirection *bool    `json:"AdaptiveResolutionDirection,omitempty" xml:"AdaptiveResolutionDirection,omitempty"`
	BFrames                     *int32   `json:"BFrames,omitempty" xml:"BFrames,omitempty"`
	Bitrate                     *int32   `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	BitrateOption               *string  `json:"BitrateOption,omitempty" xml:"BitrateOption,omitempty"`
	BufferSize                  *int32   `json:"BufferSize,omitempty" xml:"BufferSize,omitempty"`
	CRF                         *float32 `json:"CRF,omitempty" xml:"CRF,omitempty"`
	Codec                       *string  `json:"Codec,omitempty" xml:"Codec,omitempty"`
	FrameRate                   *float32 `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	FrameRateOption             *string  `json:"FrameRateOption,omitempty" xml:"FrameRateOption,omitempty"`
	GOPSize                     *int32   `json:"GOPSize,omitempty" xml:"GOPSize,omitempty"`
	MaxBitrate                  *int32   `json:"MaxBitrate,omitempty" xml:"MaxBitrate,omitempty"`
	PixelFormat                 *string  `json:"PixelFormat,omitempty" xml:"PixelFormat,omitempty"`
	Refs                        *int32   `json:"Refs,omitempty" xml:"Refs,omitempty"`
	Resolution                  *string  `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	ResolutionOption            *string  `json:"ResolutionOption,omitempty" xml:"ResolutionOption,omitempty"`
	Rotation                    *int32   `json:"Rotation,omitempty" xml:"Rotation,omitempty"`
	ScaleType                   *string  `json:"ScaleType,omitempty" xml:"ScaleType,omitempty"`
	VideoSlim                   *int32   `json:"VideoSlim,omitempty" xml:"VideoSlim,omitempty"`
}

func (s TargetVideoTranscodeVideo) String() string {
	return dara.Prettify(s)
}

func (s TargetVideoTranscodeVideo) GoString() string {
	return s.String()
}

func (s *TargetVideoTranscodeVideo) GetAdaptiveResolutionDirection() *bool {
	return s.AdaptiveResolutionDirection
}

func (s *TargetVideoTranscodeVideo) GetBFrames() *int32 {
	return s.BFrames
}

func (s *TargetVideoTranscodeVideo) GetBitrate() *int32 {
	return s.Bitrate
}

func (s *TargetVideoTranscodeVideo) GetBitrateOption() *string {
	return s.BitrateOption
}

func (s *TargetVideoTranscodeVideo) GetBufferSize() *int32 {
	return s.BufferSize
}

func (s *TargetVideoTranscodeVideo) GetCRF() *float32 {
	return s.CRF
}

func (s *TargetVideoTranscodeVideo) GetCodec() *string {
	return s.Codec
}

func (s *TargetVideoTranscodeVideo) GetFrameRate() *float32 {
	return s.FrameRate
}

func (s *TargetVideoTranscodeVideo) GetFrameRateOption() *string {
	return s.FrameRateOption
}

func (s *TargetVideoTranscodeVideo) GetGOPSize() *int32 {
	return s.GOPSize
}

func (s *TargetVideoTranscodeVideo) GetMaxBitrate() *int32 {
	return s.MaxBitrate
}

func (s *TargetVideoTranscodeVideo) GetPixelFormat() *string {
	return s.PixelFormat
}

func (s *TargetVideoTranscodeVideo) GetRefs() *int32 {
	return s.Refs
}

func (s *TargetVideoTranscodeVideo) GetResolution() *string {
	return s.Resolution
}

func (s *TargetVideoTranscodeVideo) GetResolutionOption() *string {
	return s.ResolutionOption
}

func (s *TargetVideoTranscodeVideo) GetRotation() *int32 {
	return s.Rotation
}

func (s *TargetVideoTranscodeVideo) GetScaleType() *string {
	return s.ScaleType
}

func (s *TargetVideoTranscodeVideo) GetVideoSlim() *int32 {
	return s.VideoSlim
}

func (s *TargetVideoTranscodeVideo) SetAdaptiveResolutionDirection(v bool) *TargetVideoTranscodeVideo {
	s.AdaptiveResolutionDirection = &v
	return s
}

func (s *TargetVideoTranscodeVideo) SetBFrames(v int32) *TargetVideoTranscodeVideo {
	s.BFrames = &v
	return s
}

func (s *TargetVideoTranscodeVideo) SetBitrate(v int32) *TargetVideoTranscodeVideo {
	s.Bitrate = &v
	return s
}

func (s *TargetVideoTranscodeVideo) SetBitrateOption(v string) *TargetVideoTranscodeVideo {
	s.BitrateOption = &v
	return s
}

func (s *TargetVideoTranscodeVideo) SetBufferSize(v int32) *TargetVideoTranscodeVideo {
	s.BufferSize = &v
	return s
}

func (s *TargetVideoTranscodeVideo) SetCRF(v float32) *TargetVideoTranscodeVideo {
	s.CRF = &v
	return s
}

func (s *TargetVideoTranscodeVideo) SetCodec(v string) *TargetVideoTranscodeVideo {
	s.Codec = &v
	return s
}

func (s *TargetVideoTranscodeVideo) SetFrameRate(v float32) *TargetVideoTranscodeVideo {
	s.FrameRate = &v
	return s
}

func (s *TargetVideoTranscodeVideo) SetFrameRateOption(v string) *TargetVideoTranscodeVideo {
	s.FrameRateOption = &v
	return s
}

func (s *TargetVideoTranscodeVideo) SetGOPSize(v int32) *TargetVideoTranscodeVideo {
	s.GOPSize = &v
	return s
}

func (s *TargetVideoTranscodeVideo) SetMaxBitrate(v int32) *TargetVideoTranscodeVideo {
	s.MaxBitrate = &v
	return s
}

func (s *TargetVideoTranscodeVideo) SetPixelFormat(v string) *TargetVideoTranscodeVideo {
	s.PixelFormat = &v
	return s
}

func (s *TargetVideoTranscodeVideo) SetRefs(v int32) *TargetVideoTranscodeVideo {
	s.Refs = &v
	return s
}

func (s *TargetVideoTranscodeVideo) SetResolution(v string) *TargetVideoTranscodeVideo {
	s.Resolution = &v
	return s
}

func (s *TargetVideoTranscodeVideo) SetResolutionOption(v string) *TargetVideoTranscodeVideo {
	s.ResolutionOption = &v
	return s
}

func (s *TargetVideoTranscodeVideo) SetRotation(v int32) *TargetVideoTranscodeVideo {
	s.Rotation = &v
	return s
}

func (s *TargetVideoTranscodeVideo) SetScaleType(v string) *TargetVideoTranscodeVideo {
	s.ScaleType = &v
	return s
}

func (s *TargetVideoTranscodeVideo) SetVideoSlim(v int32) *TargetVideoTranscodeVideo {
	s.VideoSlim = &v
	return s
}

func (s *TargetVideoTranscodeVideo) Validate() error {
	return dara.Validate(s)
}
