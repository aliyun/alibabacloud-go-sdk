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
	// - true: Disabled. The output file does not contain a video stream.
	//
	// - false (default): Not disabled.
	//
	// example:
	//
	// false
	DisableVideo *bool `json:"DisableVideo,omitempty" xml:"DisableVideo,omitempty"`
	// The video filter parameters. This parameter does not take effect when **TranscodeVideo*	- is empty or **TranscodeVideo.Codec*	- is set to copy.
	//
	// > This parameter is not supported for the GenerateVideoPlaylist API.
	FilterVideo *TargetVideoFilterVideo `json:"FilterVideo,omitempty" xml:"FilterVideo,omitempty" type:"Struct"`
	// The list of video stream index numbers to process from the source file. An empty value (default) indicates that the video stream with the smallest index number (the first video stream) is processed. An index number greater than 100 indicates that all video streams are processed.
	//
	// - Example: `[0,1]` processes video streams with index numbers 0 and 1. `[1]` processes the video stream with index number 1. `[101]` processes all video streams.
	//
	// > Only video streams with existing index numbers are processed. If a video stream corresponding to an index number does not exist, that index number is ignored.
	Stream []*int32 `json:"Stream,omitempty" xml:"Stream,omitempty" type:"Repeated"`
	// The video transcoding parameters. An empty value indicates that video processing is disabled and the output file does not contain a video stream.
	//
	// > Setting this parameter to an empty value to disable video processing is not recommended.
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
	// Applies mosaic processing to a rectangular area of the video to remove logos or station watermarks.
	Delogos []*TargetVideoFilterVideoDelogos `json:"Delogos,omitempty" xml:"Delogos,omitempty" type:"Repeated"`
	// The video desensitization configuration.
	//
	// 	Notice:
	//
	// - This parameter is applicable only to the CreateMediaConvertTask API.
	Desensitization *TargetVideoFilterVideoDesensitization `json:"Desensitization,omitempty" xml:"Desensitization,omitempty" type:"Struct"`
	// The video playback speed setting. Valid values: [0.5,1.0]. Default value: 1.0.
	//
	// > - This is the ratio of the transcoded media file playback speed to the source media file default playback speed, not speed-up transcoding.
	//
	// 	Notice:
	//
	// - This parameter is applicable only to the CreateMediaConvertTask API.
	//
	// example:
	//
	// 1.0
	Speed *float32 `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// The list of video watermarks.
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
	// The duration for which the mosaic is applied, in seconds (s). The default value is until the end of the video.
	//
	// example:
	//
	// 15
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The meanings differ depending on whether the value is an integer or a decimal:
	//
	// - 0 (default): Both the offset in pixels and the ratio of horizontal offset to the output resolution height are 0.
	//
	// - Integer: The offset in pixels (px). Valid values: [1,4096].
	//
	// - Decimal: The ratio of horizontal offset to the output resolution height. Valid values: (0,1).
	//
	// example:
	//
	// 0
	Dx *float32 `json:"Dx,omitempty" xml:"Dx,omitempty"`
	// Default value: 0. The meanings differ depending on whether the value is an integer or a decimal:
	//
	// - 0 (default): Both the offset in pixels and the ratio of vertical offset to the output resolution height are 0.
	//
	// - Integer: The offset in pixels (px). Valid values: [1,4096].
	//
	// - Decimal: The ratio of vertical offset to the output resolution height. Valid values: (0,1).
	//
	// example:
	//
	// 0
	Dy *float32 `json:"Dy,omitempty" xml:"Dy,omitempty"`
	// The height of the mosaic. The default value is the decimal 1.0, which fills the entire output video height. The meanings differ depending on whether the value is an integer or a decimal:
	//
	// - Integer: The height in pixels (px). Valid values: [1,4096].
	//
	// - Decimal: The ratio relative to the output video resolution height. Valid values: (0,1).
	//
	// example:
	//
	// 40
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The reference position for adding the mosaic. Valid values:
	//
	// - topleft (default): top-left corner
	//
	// - topright: top-right corner
	//
	// - bottomright: bottom-right corner
	//
	// - bottomleft: bottom-left corner
	//
	// example:
	//
	// topleft
	ReferPos *string `json:"ReferPos,omitempty" xml:"ReferPos,omitempty"`
	// The start time for adding the mosaic, in seconds (s). The default value is the start time of the video.
	//
	// example:
	//
	// 0
	StartTime *float64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The width of the mosaic. The default value is the decimal 1.0, which fills the entire output video width. The meanings differ depending on whether the value is an integer or a decimal:
	//
	// - Integer: The width in pixels (px). Valid values: [1,4096].
	//
	// - Decimal: The ratio relative to the output video resolution width. Valid values: (0,1).
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
	// The face desensitization configuration.
	//
	// > This feature is in public preview. If you have any questions, join the DingTalk group for feedback. For the DingTalk group number, see [Contact us](https://help.aliyun.com/document_detail/84454.html).
	Face *TargetVideoFilterVideoDesensitizationFace `json:"Face,omitempty" xml:"Face,omitempty" type:"Struct"`
	// The license plate desensitization configuration.
	//
	// > This feature is in public preview. If you have any questions, join the DingTalk group for feedback. For the DingTalk group number, see [Contact us](https://help.aliyun.com/document_detail/84454.html).
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
	BlurRadius *int32 `json:"BlurRadius,omitempty" xml:"BlurRadius,omitempty"`
	// The face confidence threshold, which sets the lower limit of confidence for face recognition. If the confidence value of a detected face is lower than this threshold, the face is not desensitized.
	//
	// - Valid values: 0.0 to 1.0.
	//
	// - Default value: 0.0 (no confidence filtering is performed).
	//
	// example:
	//
	// 0.4
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The minimum face size threshold, which sets the minimum size of faces to be desensitized. If the width or height of a detected face is smaller than this threshold, the face is not desensitized. Unit: pixels. Default value: 0, which indicates no size restriction on faces.
	//
	// example:
	//
	// 0.4
	MinSize      *int32   `json:"MinSize,omitempty" xml:"MinSize,omitempty"`
	ScaleRatio   *float32 `json:"ScaleRatio,omitempty" xml:"ScaleRatio,omitempty"`
	Transparency *float32 `json:"Transparency,omitempty" xml:"Transparency,omitempty"`
}

func (s TargetVideoFilterVideoDesensitizationFace) String() string {
	return dara.Prettify(s)
}

func (s TargetVideoFilterVideoDesensitizationFace) GoString() string {
	return s.String()
}

func (s *TargetVideoFilterVideoDesensitizationFace) GetBlurRadius() *int32 {
	return s.BlurRadius
}

func (s *TargetVideoFilterVideoDesensitizationFace) GetConfidence() *float32 {
	return s.Confidence
}

func (s *TargetVideoFilterVideoDesensitizationFace) GetMinSize() *int32 {
	return s.MinSize
}

func (s *TargetVideoFilterVideoDesensitizationFace) GetScaleRatio() *float32 {
	return s.ScaleRatio
}

func (s *TargetVideoFilterVideoDesensitizationFace) GetTransparency() *float32 {
	return s.Transparency
}

func (s *TargetVideoFilterVideoDesensitizationFace) SetBlurRadius(v int32) *TargetVideoFilterVideoDesensitizationFace {
	s.BlurRadius = &v
	return s
}

func (s *TargetVideoFilterVideoDesensitizationFace) SetConfidence(v float32) *TargetVideoFilterVideoDesensitizationFace {
	s.Confidence = &v
	return s
}

func (s *TargetVideoFilterVideoDesensitizationFace) SetMinSize(v int32) *TargetVideoFilterVideoDesensitizationFace {
	s.MinSize = &v
	return s
}

func (s *TargetVideoFilterVideoDesensitizationFace) SetScaleRatio(v float32) *TargetVideoFilterVideoDesensitizationFace {
	s.ScaleRatio = &v
	return s
}

func (s *TargetVideoFilterVideoDesensitizationFace) SetTransparency(v float32) *TargetVideoFilterVideoDesensitizationFace {
	s.Transparency = &v
	return s
}

func (s *TargetVideoFilterVideoDesensitizationFace) Validate() error {
	return dara.Validate(s)
}

type TargetVideoFilterVideoDesensitizationLicensePlate struct {
	BlurRadius *int32 `json:"BlurRadius,omitempty" xml:"BlurRadius,omitempty"`
	// The license plate confidence threshold, which sets the lower limit of confidence for license plate recognition. If the confidence value of a detected license plate is lower than this threshold, the license plate is not desensitized.
	//
	// - Valid values: 0.0 to 1.0.
	//
	// - Default value: 0.0 (no confidence filtering is performed).
	//
	// example:
	//
	// 0.4
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The minimum license plate size threshold, which sets the minimum size of license plates to be desensitized. If the width or height of a detected license plate is smaller than this threshold, the license plate is not desensitized. Unit: pixels. Default value: 0, which indicates no size restriction on license plates.
	//
	// example:
	//
	// 0.4
	MinSize      *int32   `json:"MinSize,omitempty" xml:"MinSize,omitempty"`
	ScaleRatio   *float32 `json:"ScaleRatio,omitempty" xml:"ScaleRatio,omitempty"`
	Transparency *float32 `json:"Transparency,omitempty" xml:"Transparency,omitempty"`
}

func (s TargetVideoFilterVideoDesensitizationLicensePlate) String() string {
	return dara.Prettify(s)
}

func (s TargetVideoFilterVideoDesensitizationLicensePlate) GoString() string {
	return s.String()
}

func (s *TargetVideoFilterVideoDesensitizationLicensePlate) GetBlurRadius() *int32 {
	return s.BlurRadius
}

func (s *TargetVideoFilterVideoDesensitizationLicensePlate) GetConfidence() *float32 {
	return s.Confidence
}

func (s *TargetVideoFilterVideoDesensitizationLicensePlate) GetMinSize() *int32 {
	return s.MinSize
}

func (s *TargetVideoFilterVideoDesensitizationLicensePlate) GetScaleRatio() *float32 {
	return s.ScaleRatio
}

func (s *TargetVideoFilterVideoDesensitizationLicensePlate) GetTransparency() *float32 {
	return s.Transparency
}

func (s *TargetVideoFilterVideoDesensitizationLicensePlate) SetBlurRadius(v int32) *TargetVideoFilterVideoDesensitizationLicensePlate {
	s.BlurRadius = &v
	return s
}

func (s *TargetVideoFilterVideoDesensitizationLicensePlate) SetConfidence(v float32) *TargetVideoFilterVideoDesensitizationLicensePlate {
	s.Confidence = &v
	return s
}

func (s *TargetVideoFilterVideoDesensitizationLicensePlate) SetMinSize(v int32) *TargetVideoFilterVideoDesensitizationLicensePlate {
	s.MinSize = &v
	return s
}

func (s *TargetVideoFilterVideoDesensitizationLicensePlate) SetScaleRatio(v float32) *TargetVideoFilterVideoDesensitizationLicensePlate {
	s.ScaleRatio = &v
	return s
}

func (s *TargetVideoFilterVideoDesensitizationLicensePlate) SetTransparency(v float32) *TargetVideoFilterVideoDesensitizationLicensePlate {
	s.Transparency = &v
	return s
}

func (s *TargetVideoFilterVideoDesensitizationLicensePlate) Validate() error {
	return dara.Validate(s)
}

type TargetVideoFilterVideoWatermarks struct {
	// The border color of the watermark text. The format is #RRGGBB. Default value: #000000. Values such as "red" and "green" are also supported.
	//
	// 	Notice:  This parameter takes effect only when the `Type` parameter is set to `text`.</notice>
	//
	// example:
	//
	// red
	BorderColor *string `json:"BorderColor,omitempty" xml:"BorderColor,omitempty"`
	// The border width of the text watermark, in pixels (px). The value must be an integer. Valid values: [0,4096]. Default value: 0.
	//
	// 	Notice:  This parameter takes effect only when the `Type` parameter is set to `text`.</notice>
	//
	// example:
	//
	// 2
	BorderWidth *int32 `json:"BorderWidth,omitempty" xml:"BorderWidth,omitempty"`
	// The content of the text watermark. Default value: empty.
	//
	// 	Notice:  This parameter takes effect only when the `Type` parameter is set to `text`.</notice>
	//
	// example:
	//
	// example
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The duration for which the watermark is displayed, in seconds (s). The default value is until the end of the video.
	//
	// example:
	//
	// 0
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The meanings differ depending on whether the value is an integer or a decimal:
	//
	// - 0 (default): Both the offset in pixels and the ratio of horizontal offset to the output resolution height are 0.
	//
	// - Integer: The offset in pixels (px). Valid values: [1,4096].
	//
	// - Decimal: The ratio of horizontal offset to the output resolution height. Valid values: (0,1).
	//
	// example:
	//
	// 0
	Dx *float32 `json:"Dx,omitempty" xml:"Dx,omitempty"`
	// The meanings differ depending on whether the value is an integer or a decimal:
	//
	// - 0 (default): Both the offset in pixels and the ratio of vertical offset to the output resolution height are 0.
	//
	// - Integer: The offset in pixels (px). Valid values: [1,4096].
	//
	// - Decimal: The ratio of vertical offset to the output resolution height. Valid values: (0,1).
	//
	// example:
	//
	// 0
	Dy *float32 `json:"Dy,omitempty" xml:"Dy,omitempty"`
	// The font opacity of the text watermark. Valid values: (0,1]. Default value: 1, which indicates fully opaque.
	//
	// 	Notice:  This parameter takes effect only when the `Type` parameter is set to `text`.</notice>
	//
	// example:
	//
	// 0.8
	FontApha *float32 `json:"FontApha,omitempty" xml:"FontApha,omitempty"`
	// The font color of the watermark text. The format is #RRGGBB. Default value: #000000. Values such as "red" and "green" are also supported.
	//
	// 	Notice:  This parameter takes effect only when the `Type` parameter is set to `text`.</notice>
	//
	// example:
	//
	// red
	FontColor *string `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	// The font name of the text watermark. Valid values:
	//
	// - SourceHanSans-Regular (default)
	//
	// - SourceHanSans-Bold
	//
	// - SourceHanSerif-Regular
	//
	// - SourceHanSerif-Bold
	//
	// 	Notice:  This parameter takes effect only when the `Type` parameter is set to `text`.</notice>
	//
	// example:
	//
	// SourceHanSans-Bold
	FontName *string `json:"FontName,omitempty" xml:"FontName,omitempty"`
	// The font size of the text watermark. Default value: 16. The value must be an integer. Valid values: (4,120).
	//
	// 	Notice:  This parameter takes effect only when the `Type` parameter is set to `text`.</notice>
	//
	// example:
	//
	// 18
	FontSize *int32 `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	// The height of the watermark image. The default value is the original height of the watermark image. The meanings differ depending on whether the value is an integer or a decimal:
	//
	// - Integer: The height in pixels (px). Valid values: [1,4096].
	//
	// - Decimal: The ratio relative to the output video resolution height. Valid values: (0,1).
	//
	// example:
	//
	// 40
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The reference position for adding the watermark. Valid values:
	//
	// - topleft (default): top-left corner
	//
	// - topright: top-right corner
	//
	// - bottomright: bottom-right corner
	//
	// - bottomleft: bottom-left corner
	//
	// example:
	//
	// topleft
	ReferPos *string `json:"ReferPos,omitempty" xml:"ReferPos,omitempty"`
	// The start time for adding the watermark, in seconds (s). The default value is the start time of the video.
	//
	// example:
	//
	// 0
	StartTime *float64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The watermark type. Valid values:
	//
	// - text (default): text watermark.
	//
	// - file: image or animated image watermark.
	//
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The OSS URI of the watermark file. Supported formats are PNG and MOV.
	//
	// The OSS URI format is `oss://<bucket>/<object>`, where `<bucket>` is the name of an OSS bucket in the same region as the current project, and `<object>` is the full path of the file including the file name extension.
	//
	// 	Notice:  This parameter takes effect only when the `Type` parameter is set to `file`.</notice>
	//
	// example:
	//
	// oss://test-bucket/watermark.jpg
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
	// The width of the watermark image. The default value is the original width of the watermark image. The meanings differ depending on whether the value is an integer or a decimal:
	//
	// - Integer: The width in pixels (px). Valid values: [1,4096].
	//
	// - Decimal: The ratio relative to the output video resolution width. Valid values: (0,1).
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
	// Specifies whether to enable adaptive long/short side mode. Valid values:
	//
	// - true: Enabled. The format of the **Resolution*	- parameter is `LongSide×ShortSide`.
	//
	// - false (default): Disabled. The format of the **Resolution*	- parameter is `Width×Height`.
	//
	// example:
	//
	// true
	AdaptiveResolutionDirection *bool `json:"AdaptiveResolutionDirection,omitempty" xml:"AdaptiveResolutionDirection,omitempty"`
	// The number of consecutive B-frames. Default value: 3.
	//
	// example:
	//
	// 3
	BFrames *int32 `json:"BFrames,omitempty" xml:"BFrames,omitempty"`
	// The video stream bitrate, in bits per second (bit/s).
	//
	// > This parameter is mutually exclusive with **CRF**. If both this parameter and **CRF*	- are empty, encoding is performed with a **CRF*	- value of 23.
	//
	// example:
	//
	// 128000
	Bitrate *int32 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The video bitrate option. Valid values:
	//
	// - fixed: Always uses the specified target video bitrate.
	//
	// - adaptive: Uses the source video bitrate when it is lower than the specified target video bitrate.
	//
	// - fall: Returns a failure when the source video bitrate is lower than the specified target video bitrate.
	//
	// Default value:
	//
	// - For the CreateMediaConvert API, the default value is fixed.
	//
	// - For the GenerateVideoPlaylist API, the default value is adaptive.
	//
	// > This parameter must be set together with the **Bitrate*	- parameter.
	//
	// example:
	//
	// fixed
	BitrateOption *string `json:"BitrateOption,omitempty" xml:"BitrateOption,omitempty"`
	// The decoding buffer size for variable bitrate, in bits per second (bps).
	//
	// > This parameter takes effect only when used together with the **CRF*	- parameter.
	//
	// example:
	//
	// 4000000
	BufferSize *int32 `json:"BufferSize,omitempty" xml:"BufferSize,omitempty"`
	// Specifies the constant quality mode. This parameter is mutually exclusive with **Bitrate**. Valid values: [0,51]. A higher value results in lower quality. Recommended values: [18,38].
	//
	// example:
	//
	// 18
	CRF *float32 `json:"CRF,omitempty" xml:"CRF,omitempty"`
	// The video encoding format. Valid values:
	//
	// - For the CreateMediaConvert API: copy (default), h264, h265, vp9.
	//
	// <warning>When this parameter is set to copy, the video streams to be processed are directly copied to the output file, and other parameters under **TranscodeVideo*	- do not take effect. copy cannot be used for video concatenation and is typically used for container format conversion scenarios.</warning>
	//
	// - For the GenerateVideoPlaylist API: h264 (default), h265.
	//
	// example:
	//
	// h264
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	// The video frame rate. The default value is the same as the source video.
	//
	// example:
	//
	// 25
	FrameRate *float32 `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	// The frame rate option. Valid values:
	//
	// - fixed: Always uses the specified target video frame rate.
	//
	// - adaptive: Uses the source video frame rate when it is lower than the specified target video frame rate.
	//
	// - fall: Returns a failure when the source video frame rate is lower than the specified target video frame rate.
	//
	// Default value:
	//
	// - For the CreateMediaConvert API, the default value is fixed.
	//
	// - For the GenerateVideoPlaylist API, the default value is adaptive.
	//
	// > This parameter must be set together with the **FrameRate*	- parameter.
	//
	// example:
	//
	// fixed
	FrameRateOption *string `json:"FrameRateOption,omitempty" xml:"FrameRateOption,omitempty"`
	// The number of frames between keyframes. Default value: 150.
	//
	// > This parameter is not supported for the GenerateVideoPlaylist API.
	//
	// example:
	//
	// 60
	GOPSize *int32 `json:"GOPSize,omitempty" xml:"GOPSize,omitempty"`
	// The maximum bitrate limit for variable bitrate. When using this parameter, you must specify the BufferSize parameter.
	//
	// > This parameter takes effect only when used together with the **CRF*	- parameter.
	//
	// example:
	//
	// 128000
	MaxBitrate *int32 `json:"MaxBitrate,omitempty" xml:"MaxBitrate,omitempty"`
	// The pixel format. The default value is the same as the source video. Valid values:
	//
	// - yuv420p
	//
	// - yuv422p
	//
	// - yuv444p
	//
	// - yuv420p10le
	//
	// - yuv422p10le
	//
	// - yuv444p10le
	//
	// - yuva420p
	//
	// > yuva420p is available only for the CreateMediaConvert API, and the **Codec*	- parameter must be set to vp9.
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
	// The resolution of the output video in the format of `WidthxHeight`. The default value is the same as the playback resolution of the source video. You can configure both width and height, or configure only width or height. You can also use the **AdaptiveResolutionDirection*	- parameter to configure both long and short sides, or configure only the long side or short side. The value range for a single side is (0,4096].
	//
	// - Example 1: If **AdaptiveResolutionDirection*	- is false, `1280x720` sets the width to 1280 and the height to 720. `1280x` sets the width to 1280 and keeps the height the same as the source video. `x720` sets the height to 720 and keeps the width the same as the source video.
	//
	// - Example 2: If **AdaptiveResolutionDirection*	- is true, `1280x720` sets the long side to 1280 and the short side to 720. `1280x` sets the long side to 1280 and keeps the short side the same as the source video. `x720` sets the short side to 720 and keeps the long side the same as the source video.
	//
	// > If the source video contains rotation information, the width/height and long/short side determination is based on the post-rotation state, which is the playback resolution.
	//
	// example:
	//
	// 640x480
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	// The resolution option. Valid values:
	//
	// - fixed: Always uses the specified target video resolution.
	//
	// - adaptive: Uses the source video resolution when the source video resolution area is smaller than the specified target video resolution area.
	//
	// - fall: Returns a failure when the source video resolution area is smaller than the specified target video resolution area.
	//
	// Default value:
	//
	// - For the CreateMediaConvert API, the default value is fixed.
	//
	// - For the GenerateVideoPlaylist API, the default value is adaptive.
	//
	// > This parameter must be set together with the **Resolution*	- parameter.
	//
	// example:
	//
	// fixed
	ResolutionOption *string `json:"ResolutionOption,omitempty" xml:"ResolutionOption,omitempty"`
	// The clockwise rotation angle of the video in degrees. Valid values:
	//
	// - 0 (default)
	//
	// - 90
	//
	// - 180
	//
	// - 270
	//
	// example:
	//
	// 90
	Rotation *int32 `json:"Rotation,omitempty" xml:"Rotation,omitempty"`
	// The scaling mode. Valid values:
	//
	// - stretch (default): Fixes the width/height or long/short sides and forcibly scales the video to fill the blank area by stretching.
	//
	// - crop: Scales proportionally to the minimum resolution that extends beyond the specified width/height or long/short side rectangle, and then center-crops the excess area.
	//
	// - fill: Scales proportionally to the maximum resolution within the specified width/height or long/short side rectangle, and then center-fills the blank area with black.
	//
	// - fit: Scales proportionally to the maximum resolution within the specified width/height or long/short side rectangle.
	//
	// > This parameter must be set together with the **Resolution*	- parameter.
	//
	// example:
	//
	// crop
	ScaleType *string `json:"ScaleType,omitempty" xml:"ScaleType,omitempty"`
	// Enables the Narrowband HD mode. Valid values:
	//
	// 0: Default value. Disabled.
	//
	// 1: Uses the Narrowband HD mode for transcoding.
	//
	// > For optimal results, use the officially recommended Bitrate or CRF parameters for video transcoding with Narrowband HD.
	//
	// >
	//
	// 	Notice: Narrowband HD supports only H.264/H.265 formats, only yuv420p, 8-bit depth, and does not support multi-target video transcoding output or video concatenation. For more information, see [Narrowband HD overview](https://help.aliyun.com/document_detail/2984556.html).
	//
	// example:
	//
	// 0
	VideoSlim *int32 `json:"VideoSlim,omitempty" xml:"VideoSlim,omitempty"`
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
