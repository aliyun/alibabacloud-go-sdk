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
	// Specifies whether to disable video stream generation. Valid values:
	//
	// 	- true: disables video stream generation. No video stream is included in the output file.
	//
	// 	- false: does not disable video stream generation. This is the default value.
	//
	// example:
	//
	// false
	DisableVideo *bool `json:"DisableVideo,omitempty" xml:"DisableVideo,omitempty"`
	// The video processing parameters. This parameter is invalid if **TranscodeVideo*	- is left empty or **TranscodeVideo.Codec*	- is set to copy.
	//
	// > This parameter is not available to the GenerateVideoPlaylist operation.
	FilterVideo *TargetVideoFilterVideo `json:"FilterVideo,omitempty" xml:"FilterVideo,omitempty" type:"Struct"`
	// The index numbers of video streams. If you do not specify this parameter, the first video stream (the one with the smallest index number) is processed. If the array contains an element greater than 100, all video streams are processed.
	//
	// 	- For example, you can set the parameter to `[0,1]` to process video streams with index numbers 0 and 1, `[1]` to process only the video stream with the index number 1, and `[101]` to process all video streams.
	//
	// > If you specify an index number but no video stream with the index number is found, the index number is ignored.
	Stream []*int32 `json:"Stream,omitempty" xml:"Stream,omitempty" type:"Repeated"`
	// The video transcoding parameters. If you do not specify this parameter, no video streams are included in the output file.
	//
	// > We recommend that you do not use this parameter to disable video stream generation.
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
	// The configurations for blurring a rectangular area of the video. This parameter is used to remove logos from videos, such as TV channel logos.
	Delogos []*TargetVideoFilterVideoDelogos `json:"Delogos,omitempty" xml:"Delogos,omitempty" type:"Repeated"`
	// The video anonymization settings.
	//
	// >
	//
	// 	- This parameter only applies to the CreateMediaConvertTask operation.
	Desensitization *TargetVideoFilterVideoDesensitization `json:"Desensitization,omitempty" xml:"Desensitization,omitempty" type:"Struct"`
	// The video playback speed. Valid values: [0.5,1.0]. Default value: 1.0.
	//
	// >
	//
	// 	- This parameter specifies the ratio of the playback speed of the output media file to the default playback speed of the source media file. It does not indicate transcoding acceleration.
	//
	// >
	//
	// 	- This parameter only applies to the CreateMediaConvertTask operation.
	//
	// example:
	//
	// 1.0
	Speed *float32 `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// The video watermarks.
	Watermarks []*TargetVideoFilterVideoWatermarks `json:"Watermarks,omitempty" xml:"Watermarks,omitempty" type:"Repeated"`
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
	// The duration of the blur in seconds. By default, the blur lasts until the end of the video.
	//
	// example:
	//
	// 15
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The value of the parameter can be an integer or a decimal.
	//
	// 	- 0: indicates that both the offset in pixels and the ratio of the horizontal offset relative to the height of the target resolution are 0. This is the default value.
	//
	// 	- An integer: the offset in pixels. Valid values: [1,4096].
	//
	// 	- A decimal: the ratio of the horizontal offset relative to the height of the target resolution. Valid values: (0,1).
	//
	// example:
	//
	// 0
	Dx *float32 `json:"Dx,omitempty" xml:"Dx,omitempty"`
	// Default value: 0. The value of the parameter can be an integer or a decimal.
	//
	// 	- 0: indicates that both the offset in pixels and the ratio of the vertical offset relative to the height of the target resolution are 0. This is the default value.
	//
	// 	- An integer: the offset in pixels. Valid values: [1,4096].
	//
	// 	- A decimal: the ratio of the vertical offset relative to the height of the target resolution. Valid values: (0,1).
	//
	// example:
	//
	// 0
	Dy *float32 `json:"Dy,omitempty" xml:"Dy,omitempty"`
	// The height of the blur. The default value is 1.0, which specifies that the blur is as high as the video. The value can be a decimal or an integer.
	//
	// 	- An integer: the number of pixels. Valid values: [1,4096].
	//
	// 	- A decimal: the ratio relative to the height of the target resolution. Valid values: (0,1).
	//
	// example:
	//
	// 40
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The reference position of the blur. Valid values:
	//
	// 	- topleft: the upper-left corner. This is the default value.
	//
	// 	- topright: the upper-right corner.
	//
	// 	- bottomright: the lower-right corner.
	//
	// 	- bottomleft: the lower-left corner.
	//
	// example:
	//
	// topleft
	ReferPos *string `json:"ReferPos,omitempty" xml:"ReferPos,omitempty"`
	// The start time of blurring (in seconds). By default, the blur begins at the start time of the video.
	//
	// example:
	//
	// 0
	StartTime *float64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The width of the blur. The default value is 1.0, which specifies that the blur is as wide as the video. The value can be a decimal or an integer.
	//
	// 	- An integer: the number of pixels. Valid values: [1,4096].
	//
	// 	- A decimal: the ratio relative to the width of the target resolution. Valid values: (0,1).
	//
	// example:
	//
	// 100
	Width *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
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
	// The settings for face anonymization.
	Face *TargetVideoFilterVideoDesensitizationFace `json:"Face,omitempty" xml:"Face,omitempty" type:"Struct"`
	// The settings for license plate anonymization.
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
	// The minimum confidence threshold. Valid values: 0 to 1. Default value: 0.
	//
	// example:
	//
	// 0
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// This parameter does not take effect if the width or height of the bounding box of a face falls below the specified minimum threshold. Default value: 0.
	//
	// example:
	//
	// 0
	MinSize *int32 `json:"MinSize,omitempty" xml:"MinSize,omitempty"`
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
	// The minimum confidence threshold. Valid values: 0 to 1. Default value: 0.
	//
	// example:
	//
	// 0
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The minimum threshold for license plate size. This parameter does not take effect if the width or height of the bounding box of a license plate falls below the specified minimum threshold. Default value: 0.
	//
	// example:
	//
	// 0
	MinSize *int32 `json:"MinSize,omitempty" xml:"MinSize,omitempty"`
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
	// The color of the text watermark border. You can specify a color name, such as "red" or "green", or an RGB color code. The default color is #000000.
	//
	// > This parameter takes effect only when the Type parameter is set to text.
	//
	// example:
	//
	// red
	BorderColor *string `json:"BorderColor,omitempty" xml:"BorderColor,omitempty"`
	// The width of the text watermark border. Unit: pixels. The value must be an integer within the [0,4096] range. Default value: 0.
	//
	// > This parameter takes effect only when the Type parameter is set to text.
	//
	// example:
	//
	// 2
	BorderWidth *int32 `json:"BorderWidth,omitempty" xml:"BorderWidth,omitempty"`
	// The content of the text watermark. By default, this parameter is left empty.
	//
	// > This parameter takes effect only when the Type parameter is set to text.
	//
	// example:
	//
	// example
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The duration of watermarking (in seconds). By default, the watermark lasts until the video ends.
	//
	// example:
	//
	// 0
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The value of the parameter can be an integer or a decimal.
	//
	// 	- 0: indicates that both the offset in pixels and the ratio of the horizontal offset relative to the height of the target resolution are 0. This is the default value.
	//
	// 	- An integer: the offset in pixels. Valid values: [1,4096].
	//
	// 	- A decimal: the ratio of the horizontal offset relative to the height of the target resolution. Valid values: (0,1).
	//
	// example:
	//
	// 0
	Dx *float32 `json:"Dx,omitempty" xml:"Dx,omitempty"`
	// The value of the parameter can be an integer or a decimal.
	//
	// 	- 0: indicates that both the offset in pixels and the ratio of the vertical offset relative to the height of the target resolution are 0. This is the default value.
	//
	// 	- An integer: the offset in pixels. Valid values: [1,4096].
	//
	// 	- A decimal: the ratio of the vertical offset relative to the height of the target resolution. Valid values: (0,1).
	//
	// example:
	//
	// 0
	Dy *float32 `json:"Dy,omitempty" xml:"Dy,omitempty"`
	// The font transparency of the text watermark. Valid values: (0,1]. Default value: 1, which specifies that the text watermark is fully opaque.
	//
	// > This parameter takes effect only when the Type parameter is set to text.
	//
	// example:
	//
	// 0.8
	FontApha *float32 `json:"FontApha,omitempty" xml:"FontApha,omitempty"`
	// The color of the text watermark. You can specify a color name, such as "red" or "green", or an RGB color code. The default color is #000000.
	//
	// > This parameter takes effect only when the Type parameter is set to text.
	//
	// example:
	//
	// red
	FontColor *string `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	// The font of the text watermark. Valid values:
	//
	// 	- SourceHanSans-Regular (default)
	//
	// 	- SourceHanSans-Bold
	//
	// 	- SourceHanSerif-Regular
	//
	// 	- SourceHanSerif-Bold
	//
	// > This parameter takes effect only when the Type parameter is set to text.
	//
	// example:
	//
	// SourceHanSans-Bold
	FontName *string `json:"FontName,omitempty" xml:"FontName,omitempty"`
	// The size of the text watermark. Default value: 16. The value must be an integer within the (4,120) range.
	//
	// > This parameter takes effect only when the Type parameter is set to text.
	//
	// example:
	//
	// 18
	FontSize *int32 `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	// The height of the image watermark. By default, the height is the same as the height of the original watermark file. The value of the parameter can be an integer or a decimal.
	//
	// 	- An integer: the number of pixels excluding the height of the logo. Valid values: [1,4096].
	//
	// 	- A decimal: the ratio relative to the height of the target resolution. Valid values: (0,1).
	//
	// example:
	//
	// 40
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The reference position for adding the watermark. Valid values:
	//
	// 	- topleft: the upper-left corner. This is the default value.
	//
	// 	- topright: the upper-right corner.
	//
	// 	- bottomright: the lower-right corner.
	//
	// 	- bottomleft: the lower-left corner.
	//
	// example:
	//
	// topleft
	ReferPos *string `json:"ReferPos,omitempty" xml:"ReferPos,omitempty"`
	// The start time of watermarking (in seconds). By default, the watermark begins at the start time of the video.
	//
	// example:
	//
	// 0
	StartTime *float64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The watermark type. Valid values:
	//
	// 	- text: a text watermark. This is the default value.
	//
	// 	- file: a still or animated image watermark.
	//
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The Object Storage Service (OSS) URI of the watermark file. The watermark file can be a PNG or MOV file.
	//
	// The URI is in the `oss://<bucket>/<object>` format, where `<bucket>` is the name of the bucket in the same region as the current project and `<object>` is the path of the object with the extension included.
	//
	// > This parameter takes effect only when the Type parameter is set to file.
	//
	// example:
	//
	// oss://test-bucket/watermark
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
	// The width of the image watermark. By default, the width is the same as the width of the original watermark file. The value of the parameter can be an integer or a decimal.
	//
	// 	- An integer: the number of pixels excluding the width of the logo. Valid values: [1,4096].
	//
	// 	- A decimal: the ratio relative to the width of the target resolution. Valid values: (0,1).
	//
	// example:
	//
	// 80
	Width *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
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
	// Specifies whether to enable adaptation to resolution based on long and short sides. Valid values:
	//
	// 	- true: The format of the **Resolution*	- parameter is `LongSide×ShortSide`. This is the default value.
	//
	// 	- false: The format of the **Resolution*	- parameter is `Width×Height`.
	//
	// example:
	//
	// true
	AdaptiveResolutionDirection *bool `json:"AdaptiveResolutionDirection,omitempty" xml:"AdaptiveResolutionDirection,omitempty"`
	// The number of consecutive B frames. Default value: 3.
	//
	// example:
	//
	// 3
	BFrames *int32 `json:"BFrames,omitempty" xml:"BFrames,omitempty"`
	// The bitrate of the video stream. Unit: bit/s.
	//
	// > This parameter and the **CRF*	- parameter are mutually exclusive. If you leave both the parameters empty, the **CRF*	- parameter with a value of 23 applies.
	//
	// example:
	//
	// 128000
	Bitrate *int32 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The video bitrate option. Valid values:
	//
	// 	- fixed: always uses the target bitrate.
	//
	// 	- adaptive: uses the source bitrate when the source bitrate is less than the target bitrate.
	//
	// 	- fall: returns a failure when the source bitrate is less than the target bitrate.
	//
	// Default value:
	//
	// 	- fixed for the CreateMediaConvert operation.
	//
	// 	- adaptive for the GenerateVideoPlaylist operation.
	//
	// > This parameter must be used together with the **Bitrate*	- parameter.
	//
	// example:
	//
	// fixed
	BitrateOption *string `json:"BitrateOption,omitempty" xml:"BitrateOption,omitempty"`
	// The size of the buffer for decoding when the dynamic bitrate is used. Unit: bit/s.
	//
	// > This parameter must be used together with the **CRF*	- parameter.
	//
	// example:
	//
	// 4000000
	BufferSize *int32 `json:"BufferSize,omitempty" xml:"BufferSize,omitempty"`
	// The constant rate factor (CRF) of the video. The value of this parameter falls within the [0,51] range. A greater indicates lower quality. We recommend that you specify a value within the [18,38] range. This parameter and the **Bitrate*	- parameter are mutually exclusive.
	//
	// example:
	//
	// 18
	CRF *float32 `json:"CRF,omitempty" xml:"CRF,omitempty"`
	// The video coding format. Valid values:
	//
	// 	- copy, h264, h265, and vp9 for the CreateMediaConvert operation. The default value is copy.
	//
	//     **
	//
	//     **Warning **When you set the parameter to copy, the video stream is directly copied into the output file and all other parameters in TranscodeVideo do not take effect. The copy value is commonly used in container format conversion scenarios. You cannot use this value in video merging scenarios.
	//
	// 	- h264 and h265 for the GenerateVideoPlaylist operation. The default value is h264.
	//
	// example:
	//
	// h264
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	// The target frame rate. By default, the target frame rate is the same as the source frame rate.
	//
	// example:
	//
	// 25
	FrameRate *float32 `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	// The frame rate option. Valid values:
	//
	// 	- fixed: always uses the target frame rate.
	//
	// 	- adaptive: uses the source frame rate when the source frame rate is less than the target frame rate.
	//
	// 	- fall: returns a failure if the source frame rate is less than the target frame rate.
	//
	// Default value:
	//
	// 	- fixed for the CreateMediaConvert operation.
	//
	// 	- adaptive for the GenerateVideoPlaylist operation.
	//
	// > This parameter must be used together with the **FrameRate*	- parameter.
	//
	// example:
	//
	// fixed
	FrameRateOption *string `json:"FrameRateOption,omitempty" xml:"FrameRateOption,omitempty"`
	// The keyframe interval. Default value: 150.
	//
	// > This parameter is not available to the GenerateVideoPlaylist operation.
	//
	// example:
	//
	// 60
	GOPSize *int32 `json:"GOPSize,omitempty" xml:"GOPSize,omitempty"`
	// The maximum bitrate when the dynamic bitrate is used. If you specify this parameter, you must also specify the BufferSize parameter.
	//
	// > This parameter must be used together with the **CRF*	- parameter.
	//
	// example:
	//
	// 128000
	MaxBitrate *int32 `json:"MaxBitrate,omitempty" xml:"MaxBitrate,omitempty"`
	// The pixel format. By default, the pixel format matches that of the source video. Valid values:
	//
	// 	- yuv420p
	//
	// 	- yuv422p
	//
	// 	- yuv444p
	//
	// 	- yuv420p10le
	//
	// 	- yuv422p10le
	//
	// 	- yuv444p10le
	//
	// 	- yuva420p
	//
	// > You can set the parameter to yuva420p only when you call the CreateMediaConvert operation and set the **Codec*	- parameter to vp9.
	//
	// example:
	//
	// yuv420p
	PixelFormat *string `json:"PixelFormat,omitempty" xml:"PixelFormat,omitempty"`
	// The number of reference frames. Default value: 2.
	//
	// example:
	//
	// 2
	Refs *int32 `json:"Refs,omitempty" xml:"Refs,omitempty"`
	// The target resolution in the `WidthxHeight` format. By default, the target resolution is the same as the source resolution. You can specify both width and height, or only one of them. You can use this parameter together with the **AdaptiveResolutionDirection*	- parameter to set both the long side and short side or one of them. The value of each side falls within the range of (0,4096].
	//
	// 	- Example 1: If **AdaptiveResolutionDirection*	- is set to false, `1280x720` specifies a width of 1280 pixels and a height of 720 pixels, `1280x` specifies a width of 1280 pixels and the same height as the source video, and `x720` specifies a height of 720 pixels and the same width as the source video.
	//
	// 	- Example 2: If **AdaptiveResolutionDirection*	- is set to true, `1280x720` specifies a long side of 1280 pixels and a short side of 720 pixels, `1280x` specifies a long side of 1280 pixels and the same short side as the source video, and `x720` specifies a short side of 720 pixels and the same long side as the source video.
	//
	// > If the source video contains rotation information, the width, height, long side, and short side depend on the frame after rotation (the playback resolution).
	//
	// example:
	//
	// 640x480
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	// The resolution option. Valid values:
	//
	// 	- fixed: always uses the target video resolution.
	//
	// 	- adaptive: uses the source video resolution when the source video resolution is less than the target video resolution.
	//
	// 	- fall: returns a failure when the source video resolution is less than the target video resolution.
	//
	// Default value:
	//
	// 	- fixed for the CreateMediaConvert operation.
	//
	// 	- adaptive for the GenerateVideoPlaylist operation.
	//
	// > This parameter must be used together with the **Resolution*	- parameter.
	//
	// example:
	//
	// fixed
	ResolutionOption *string `json:"ResolutionOption,omitempty" xml:"ResolutionOption,omitempty"`
	// The degrees to rotate the video clockwise. Valid values:
	//
	// 	- 0 (default)
	//
	// 	- 90
	//
	// 	- 180
	//
	// 	- 270
	//
	// example:
	//
	// 90
	Rotation *int32 `json:"Rotation,omitempty" xml:"Rotation,omitempty"`
	// The resizing mode. Valid values:
	//
	// 	- stretch: forcibly stretches the video based on the specified width and height or long side and short side to fill any remaining space. This is the default value.
	//
	// 	- crop: proportionally resizes the video to the minimum resolution outside the rectangular shape based on the specified width and height or long side and short side, and crops the parts beyond the rectangular shape from the center.
	//
	// 	- fill: proportionally resizes the video to the maximum resolution within the rectangular shape based on the specified width and height or long side and short side, and fills empty space with black from the center.
	//
	// 	- fit: proportionally resizes the video to the maximum resolution that fits within the specified width and height or long side and short side.
	//
	// > This parameter must be used together with the **Resolution*	- parameter.
	//
	// example:
	//
	// crop
	ScaleType *string `json:"ScaleType,omitempty" xml:"ScaleType,omitempty"`
	VideoSlim *int32  `json:"VideoSlim,omitempty" xml:"VideoSlim,omitempty"`
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
