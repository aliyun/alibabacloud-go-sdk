// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMediaConvertVideo interface {
	dara.Model
	String() string
	GoString() string
	SetBitrate(v int32) *MediaConvertVideo
	GetBitrate() *int32
	SetBufsize(v int32) *MediaConvertVideo
	GetBufsize() *int32
	SetCodec(v string) *MediaConvertVideo
	GetCodec() *string
	SetCrf(v interface{}) *MediaConvertVideo
	GetCrf() interface{}
	SetCrop(v string) *MediaConvertVideo
	GetCrop() *string
	SetFps(v interface{}) *MediaConvertVideo
	GetFps() interface{}
	SetGop(v interface{}) *MediaConvertVideo
	GetGop() interface{}
	SetHeight(v int32) *MediaConvertVideo
	GetHeight() *int32
	SetLongShortMode(v bool) *MediaConvertVideo
	GetLongShortMode() *bool
	SetMaxFps(v interface{}) *MediaConvertVideo
	GetMaxFps() interface{}
	SetMaxrate(v int32) *MediaConvertVideo
	GetMaxrate() *int32
	SetPad(v string) *MediaConvertVideo
	GetPad() *string
	SetProfile(v string) *MediaConvertVideo
	GetProfile() *string
	SetQscale(v int32) *MediaConvertVideo
	GetQscale() *int32
	SetRemove(v bool) *MediaConvertVideo
	GetRemove() *bool
	SetScanMode(v string) *MediaConvertVideo
	GetScanMode() *string
	SetWidth(v int32) *MediaConvertVideo
	GetWidth() *int32
}

type MediaConvertVideo struct {
	// The average bitrate of the output. If you use the CRB, ABR, or VBR bitrate control mode, you must specify Bitrate, and you must set TransMode to a valid value.
	//
	// 	- Unit: Kbps.
	//
	// 	- Valid values: -1 and [10,50000]. A value of -1 indicates that the original bitrate of the input is used.
	//
	// Best practices:
	//
	// 	- CBR: Set TransMode to CBR and Bitrate, Maxrate, and Bufsize to the same value.
	//
	// 	- ABR: Set TransMode to onepass and specify Bitrate. You can also specify Maxrate and Bufsize to control the bitrate range.
	//
	// 	- VBR: Set TransMode to twopass and specify Maxrate/BitrateBnd and Bufsize.
	//
	// example:
	//
	// 6000
	Bitrate *int32 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The buffer size. It controls bitrate fluctuations. A larger value allows for more bitrate variation and potentially higher video quality.
	//
	// 	- Unit: Kbps.
	//
	// 	- Valid values: [1000,128000].
	//
	// 	- Default value: 6000.
	//
	// example:
	//
	// 20000
	Bufsize *int32 `json:"Bufsize,omitempty" xml:"Bufsize,omitempty"`
	// The video codec.
	//
	// 	- Valid values: H.264, H.265, AV1, GIF, and WEBP.
	//
	// 	- Default value: H.264.
	//
	// example:
	//
	// H.264
	Codec *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	// The quality control factor. To use the CRF mode, you must specify Crf and set TransMode to fixCRF. A larger value means lower quality and a higher compression ratio.
	//
	// 	- Valid values: [20,51].
	//
	// 	- If Codec is set to H.264, the default value is 23. If Codec is set to H.265, the default value is 26. If Codec is set to AV1, the default value is 32.
	//
	// Best Practice:
	//
	// 	- A value of 0 specifies lossless quality. A value of 51 specifies the worst quality. A recommended range is [23, 29]. You can adjust the value based on the complexity of the image. If you increase the value by six, the bitrate is reduced by half. Under the same definition, you can set the value for an animated cartoon higher than that for a shot video.
	//
	// 	- CRF targets perceptual quality, not a fixed bitrate. Use it with Maxrate and Bufsize to control bitrate fluctuations.
	//
	// example:
	//
	// 23
	Crf interface{} `json:"Crf,omitempty" xml:"Crf,omitempty"`
	// Crops the video frame. You can set automatic black border removal or custom cropping.
	//
	// 	- Specify this parameter if the input resolution is greater than the output resolution. Do not specify AdjDarMethod if this parameter is specified.
	//
	// 	- To automatically remove black borders, set this parameter to border.
	//
	// 	- To use custom cropping, set the parameter in the format of {width}:{height}:{left}:{top}.
	//
	//     	- width: the width of the output video.
	//
	//     	- height: the height of the output video.
	//
	//     	- left: the distance between the left border of the output and that of the original frame.
	//
	//     	- top: the distance between the top border of the output and that of the original frame.
	//
	// Example: 1920:800:0:140.
	//
	// example:
	//
	// 1920:800:0:140
	Crop *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	// The frame rate of the video stream.
	//
	// 	- Unit: frames per second (FPS).
	//
	// 	- Valid values: (0,60].
	//
	// 	- Default value: the frame rate of the input file. If the original frame rate exceeds 60 FPS, 60 is used.
	//
	// 	- Recommended values: 24, 25, and 30.
	//
	// example:
	//
	// 25
	Fps interface{} `json:"Fps,omitempty" xml:"Fps,omitempty"`
	// The keyframe interval.
	//
	// 	- Set by time: The time interval, in seconds. Valid values: [1,100000].
	//
	// 	- Set by frame count: The number of frames. Valid values: [1,100000].
	//
	// 	- Default value: 10s.
	//
	// Best practice: Set this parameter to 2-7s to improve the Time-to-First-Frame and seeking performance.
	//
	// example:
	//
	// 10s
	Gop interface{} `json:"Gop,omitempty" xml:"Gop,omitempty"`
	// The height or short edge of the output. If LongShortMode is set to false or left empty, this parameter specifies the height of the output. If LongShortMode is set to true, this parameter specifies the short edge of the output.
	//
	// 	- Unit: pixel.
	//
	// 	- Valid values: [128,4096]. The value must be an even number.
	//
	// 	- Default value:
	//
	//     	- If neither Width nor Height is specified, the dimension of the input is used.
	//
	//     	- If only Width is specified, Height is auto-scaled.
	//
	// example:
	//
	// 1080
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// Specifies whether to enable orientation-adaptive scaling. This parameter takes effect if at least one of the Width and Height parameters is specified. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// 	- Default value: false.
	//
	// Best practice: Enable this feature when your inputs include both horizontal and vertical videos. This prevents videos from stretching.
	//
	// example:
	//
	// true
	LongShortMode *bool `json:"LongShortMode,omitempty" xml:"LongShortMode,omitempty"`
	// The maximum frame rate.
	//
	// example:
	//
	// 25
	MaxFps interface{} `json:"MaxFps,omitempty" xml:"MaxFps,omitempty"`
	// The maximum bitrate of the output.
	//
	// 	- Unit: Kbps.
	//
	// 	- Valid values: [10,50000].
	//
	// example:
	//
	// 10000
	Maxrate *int32 `json:"Maxrate,omitempty" xml:"Maxrate,omitempty"`
	// The black borders added to the video.
	//
	// 	- Specify this parameter if the input resolution is smaller than the output resolution. If you specify this parameter, do not specify IsCheckReso, IsCheckResoFail, or AdjDarMethod.
	//
	// 	- Format: {width}:{height}:{left}:{top}.
	//
	//     	- width: the width of the output video.
	//
	//     	- height: the height of the output video.
	//
	//     	- left: the distance between the left border of the output and that of the original frame.
	//
	//     	- top: the distance between the top border of the output and that of the original frame.
	//
	// Example: 1920:1080:0:140.
	//
	// example:
	//
	// 1920:1080:0:140
	Pad *string `json:"Pad,omitempty" xml:"Pad,omitempty"`
	// The codec profile.
	//
	// 	- This parameter takes effect only if Codec is set to H.264.
	//
	// 	- Valid values: baseline, main, and high.
	//
	// 	- Default value: high.
	//
	// Best practice: For multi-bitrate streaming, use baseline for the lowest quality rendition to ensure maximum compatibility with older devices. Use main or high for other renditions.
	//
	// example:
	//
	// high
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The video quality scale. This parameter applies to VBR mode. A higher value means lower video quality and a higher compression ratio.
	//
	// 	- This parameter takes effect only if Codec is set to H.264.
	//
	// 	- Valid values: [0,51].
	//
	// example:
	//
	// 13
	Qscale *int32 `json:"Qscale,omitempty" xml:"Qscale,omitempty"`
	// Specifies whether to delete the video stream. Valid values:
	//
	// 	- true: deletes the video stream. All video-related parameters are invalid.
	//
	// 	- false: retains the video stream.
	//
	// 	- Default value: false.
	//
	// example:
	//
	// false
	Remove *bool `json:"Remove,omitempty" xml:"Remove,omitempty"`
	// The scan mode. Valid values:
	//
	// 	- If you leave this parameter empty, the output follows the source\\"s original scan mode.
	//
	// 	- auto: automatic deinterlacing
	//
	// 	- progressive
	//
	// 	- interlaced
	//
	// 	- By default, this parameter is left empty.
	//
	// Best practice: The interlaced scan mode saves data traffic than the progressive scan mode but provides poor image quality. Therefore, the latter is commonly used in mainstream video production.
	//
	// 	- If you set ScanMode to progressive or interlaced, but this scan mode does not match that of the input, the video fails to be transcoded.
	//
	// 	- To improve compatibility, leave this parameter empty or set it to auto.
	//
	// example:
	//
	// auto
	ScanMode *string `json:"ScanMode,omitempty" xml:"ScanMode,omitempty"`
	// The width or long edge of the output. If LongShortMode is set to false or left empty, this parameter specifies the width of the output. If LongShortMode is set to true, this parameter specifies the long edge of the output.
	//
	// 	- Unit: pixel.
	//
	// 	- Valid values: [128,4096]. The value must be an even number.
	//
	// 	- Default value:
	//
	//     	- If neither Width nor Height is specified, the dimension of the input is used.
	//
	//     	- If only Height is specified, Width is auto-scaled.
	//
	// example:
	//
	// 1920
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s MediaConvertVideo) String() string {
	return dara.Prettify(s)
}

func (s MediaConvertVideo) GoString() string {
	return s.String()
}

func (s *MediaConvertVideo) GetBitrate() *int32 {
	return s.Bitrate
}

func (s *MediaConvertVideo) GetBufsize() *int32 {
	return s.Bufsize
}

func (s *MediaConvertVideo) GetCodec() *string {
	return s.Codec
}

func (s *MediaConvertVideo) GetCrf() interface{} {
	return s.Crf
}

func (s *MediaConvertVideo) GetCrop() *string {
	return s.Crop
}

func (s *MediaConvertVideo) GetFps() interface{} {
	return s.Fps
}

func (s *MediaConvertVideo) GetGop() interface{} {
	return s.Gop
}

func (s *MediaConvertVideo) GetHeight() *int32 {
	return s.Height
}

func (s *MediaConvertVideo) GetLongShortMode() *bool {
	return s.LongShortMode
}

func (s *MediaConvertVideo) GetMaxFps() interface{} {
	return s.MaxFps
}

func (s *MediaConvertVideo) GetMaxrate() *int32 {
	return s.Maxrate
}

func (s *MediaConvertVideo) GetPad() *string {
	return s.Pad
}

func (s *MediaConvertVideo) GetProfile() *string {
	return s.Profile
}

func (s *MediaConvertVideo) GetQscale() *int32 {
	return s.Qscale
}

func (s *MediaConvertVideo) GetRemove() *bool {
	return s.Remove
}

func (s *MediaConvertVideo) GetScanMode() *string {
	return s.ScanMode
}

func (s *MediaConvertVideo) GetWidth() *int32 {
	return s.Width
}

func (s *MediaConvertVideo) SetBitrate(v int32) *MediaConvertVideo {
	s.Bitrate = &v
	return s
}

func (s *MediaConvertVideo) SetBufsize(v int32) *MediaConvertVideo {
	s.Bufsize = &v
	return s
}

func (s *MediaConvertVideo) SetCodec(v string) *MediaConvertVideo {
	s.Codec = &v
	return s
}

func (s *MediaConvertVideo) SetCrf(v interface{}) *MediaConvertVideo {
	s.Crf = v
	return s
}

func (s *MediaConvertVideo) SetCrop(v string) *MediaConvertVideo {
	s.Crop = &v
	return s
}

func (s *MediaConvertVideo) SetFps(v interface{}) *MediaConvertVideo {
	s.Fps = v
	return s
}

func (s *MediaConvertVideo) SetGop(v interface{}) *MediaConvertVideo {
	s.Gop = v
	return s
}

func (s *MediaConvertVideo) SetHeight(v int32) *MediaConvertVideo {
	s.Height = &v
	return s
}

func (s *MediaConvertVideo) SetLongShortMode(v bool) *MediaConvertVideo {
	s.LongShortMode = &v
	return s
}

func (s *MediaConvertVideo) SetMaxFps(v interface{}) *MediaConvertVideo {
	s.MaxFps = v
	return s
}

func (s *MediaConvertVideo) SetMaxrate(v int32) *MediaConvertVideo {
	s.Maxrate = &v
	return s
}

func (s *MediaConvertVideo) SetPad(v string) *MediaConvertVideo {
	s.Pad = &v
	return s
}

func (s *MediaConvertVideo) SetProfile(v string) *MediaConvertVideo {
	s.Profile = &v
	return s
}

func (s *MediaConvertVideo) SetQscale(v int32) *MediaConvertVideo {
	s.Qscale = &v
	return s
}

func (s *MediaConvertVideo) SetRemove(v bool) *MediaConvertVideo {
	s.Remove = &v
	return s
}

func (s *MediaConvertVideo) SetScanMode(v string) *MediaConvertVideo {
	s.ScanMode = &v
	return s
}

func (s *MediaConvertVideo) SetWidth(v int32) *MediaConvertVideo {
	s.Width = &v
	return s
}

func (s *MediaConvertVideo) Validate() error {
	return dara.Validate(s)
}
