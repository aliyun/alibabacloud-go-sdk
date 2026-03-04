// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMediaConvertTransConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAdjDarMethod(v string) *MediaConvertTransConfig
	GetAdjDarMethod() *string
	SetIsCheckAudioBitrate(v bool) *MediaConvertTransConfig
	GetIsCheckAudioBitrate() *bool
	SetIsCheckAudioBitrateFail(v bool) *MediaConvertTransConfig
	GetIsCheckAudioBitrateFail() *bool
	SetIsCheckReso(v bool) *MediaConvertTransConfig
	GetIsCheckReso() *bool
	SetIsCheckResoFail(v bool) *MediaConvertTransConfig
	GetIsCheckResoFail() *bool
	SetIsCheckVideoBitrate(v bool) *MediaConvertTransConfig
	GetIsCheckVideoBitrate() *bool
	SetIsCheckVideoBitrateFail(v bool) *MediaConvertTransConfig
	GetIsCheckVideoBitrateFail() *bool
	SetTransMode(v string) *MediaConvertTransConfig
	GetTransMode() *string
}

type MediaConvertTransConfig struct {
	// The method for adjusting the resolution. This parameter takes effect only if both the Width and Height parameters are specified. You can use this parameter together with the LongShortMode parameter.
	//
	// 	- Valid values: rescale, crop, pad, and none.
	//
	// 	- Default value: none.
	//
	// *
	//
	// example:
	//
	// rescale
	AdjDarMethod *string `json:"AdjDarMethod,omitempty" xml:"AdjDarMethod,omitempty"`
	// Specifies whether to check the audio bitrate. Mutually exclusive with IsCheckVideoBitrateFail. IsCheckVideoBitrateFail has a higher priority.
	//
	// 	- true: checks the audio bitrate. If the source bitrate is lower than the target, the task will transcode using the source bitrate.
	//
	// 	- false: does not check the audio bitrate.
	//
	// 	- Default value:
	//
	//     	- If this parameter is not specified and the codec of the output audio is different from that of the input audio, the default value is false.
	//
	//     	- If this parameter is not specified and the codec of the output audio is the same as that of the input audio, the default value is true.
	IsCheckAudioBitrate *bool `json:"IsCheckAudioBitrate,omitempty" xml:"IsCheckAudioBitrate,omitempty"`
	// Specifies whether to check the audio bitrate. Mutually exclusive with IsCheckAudioBitrate. This parameter has a higher priority.
	//
	// 	- true: checks the audio bitrate. If the source bitrate is lower than the target, the task will fail.
	//
	// 	- false: does not check the audio bitrate.
	//
	// 	- Default value: false.
	IsCheckAudioBitrateFail *bool `json:"IsCheckAudioBitrateFail,omitempty" xml:"IsCheckAudioBitrateFail,omitempty"`
	// Specifies whether to check the video resolution. Mutually exclusive with IsCheckResoFail. IsCheckResoFail has a higher priority.
	//
	// 	- true: checks the video resolution. If the source resolution (width or height) is smaller than the target, the task will transcode using the source resolution.
	//
	// 	- false: does not check the video resolution.
	//
	// 	- Default value: false.
	IsCheckReso *bool `json:"IsCheckReso,omitempty" xml:"IsCheckReso,omitempty"`
	// Specifies whether to check the video resolution. Mutually exclusive with IsCheckReso. This parameter has a higher priority.
	//
	// 	- true: checks the video resolution. If the source resolution (width or height) is smaller than the target, the task will fail.
	//
	// 	- false: does not check the video resolution.
	//
	// 	- Default value: false.
	IsCheckResoFail *bool `json:"IsCheckResoFail,omitempty" xml:"IsCheckResoFail,omitempty"`
	// Specifies whether to check the video bitrate. Mutually exclusive with IsCheckVideoBitrateFail. IsCheckVideoBitrateFail has a higher priority.
	//
	// 	- true: checks the video bitrate. If the source bitrate is lower than the target, the task will transcode using the source bitrate.
	//
	// 	- false: does not check the video bitrate.
	//
	// 	- Default value: false.
	IsCheckVideoBitrate *bool `json:"IsCheckVideoBitrate,omitempty" xml:"IsCheckVideoBitrate,omitempty"`
	// Specifies whether to check the video bitrate. Mutually exclusive with IsCheckVideoBitrate. This parameter has a higher priority.
	//
	// 	- true: checks the video bitrate. If the source bitrate is lower than the target, the task will fail.
	//
	// 	- false: does not check the video bitrate.
	//
	// 	- Default value: false.
	IsCheckVideoBitrateFail *bool `json:"IsCheckVideoBitrateFail,omitempty" xml:"IsCheckVideoBitrateFail,omitempty"`
	// The video bitrate control mode. This is only valid for H.264, H.265, and AV1 codecs and must be used with the correct Bitrate or Crf settings. Valid values:
	//
	// 	- CBR: Constant bitrate mode.
	//
	// 	- onepass: Typically used for Average bitrate mode (ABR). Faster than twopass.
	//
	// 	- twopass: Typically used for variable bitrate mode (VBR). Slower than onepass.
	//
	// 	- fixCRF: Constant Rate Factor (quality-based) mode.
	//
	// 	- If you specify the Bitrate parameter, the default value of this parameter is onepass. If you do not specify the Bitrate parameter, the default value of this parameter is fixCRF, and the default value of the Crf parameter is used.
	//
	// example:
	//
	// onepass
	TransMode *string `json:"TransMode,omitempty" xml:"TransMode,omitempty"`
}

func (s MediaConvertTransConfig) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertTransConfig) GoString() string {
	return s.String()
}

func (s *MediaConvertTransConfig) GetAdjDarMethod() *string {
	return s.AdjDarMethod
}

func (s *MediaConvertTransConfig) GetIsCheckAudioBitrate() *bool {
	return s.IsCheckAudioBitrate
}

func (s *MediaConvertTransConfig) GetIsCheckAudioBitrateFail() *bool {
	return s.IsCheckAudioBitrateFail
}

func (s *MediaConvertTransConfig) GetIsCheckReso() *bool {
	return s.IsCheckReso
}

func (s *MediaConvertTransConfig) GetIsCheckResoFail() *bool {
	return s.IsCheckResoFail
}

func (s *MediaConvertTransConfig) GetIsCheckVideoBitrate() *bool {
	return s.IsCheckVideoBitrate
}

func (s *MediaConvertTransConfig) GetIsCheckVideoBitrateFail() *bool {
	return s.IsCheckVideoBitrateFail
}

func (s *MediaConvertTransConfig) GetTransMode() *string {
	return s.TransMode
}

func (s *MediaConvertTransConfig) SetAdjDarMethod(v string) *MediaConvertTransConfig {
	s.AdjDarMethod = &v
	return s
}

func (s *MediaConvertTransConfig) SetIsCheckAudioBitrate(v bool) *MediaConvertTransConfig {
	s.IsCheckAudioBitrate = &v
	return s
}

func (s *MediaConvertTransConfig) SetIsCheckAudioBitrateFail(v bool) *MediaConvertTransConfig {
	s.IsCheckAudioBitrateFail = &v
	return s
}

func (s *MediaConvertTransConfig) SetIsCheckReso(v bool) *MediaConvertTransConfig {
	s.IsCheckReso = &v
	return s
}

func (s *MediaConvertTransConfig) SetIsCheckResoFail(v bool) *MediaConvertTransConfig {
	s.IsCheckResoFail = &v
	return s
}

func (s *MediaConvertTransConfig) SetIsCheckVideoBitrate(v bool) *MediaConvertTransConfig {
	s.IsCheckVideoBitrate = &v
	return s
}

func (s *MediaConvertTransConfig) SetIsCheckVideoBitrateFail(v bool) *MediaConvertTransConfig {
	s.IsCheckVideoBitrateFail = &v
	return s
}

func (s *MediaConvertTransConfig) SetTransMode(v string) *MediaConvertTransConfig {
	s.TransMode = &v
	return s
}

func (s *MediaConvertTransConfig) Validate() error {
	return dara.Validate(s)
}
