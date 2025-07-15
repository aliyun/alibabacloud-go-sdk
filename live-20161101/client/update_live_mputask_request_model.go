// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveMPUTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateLiveMPUTaskRequest
	GetAppId() *string
	SetChannelId(v string) *UpdateLiveMPUTaskRequest
	GetChannelId() *string
	SetMixMode(v string) *UpdateLiveMPUTaskRequest
	GetMixMode() *string
	SetMultiStreamURL(v []*UpdateLiveMPUTaskRequestMultiStreamURL) *UpdateLiveMPUTaskRequest
	GetMultiStreamURL() []*UpdateLiveMPUTaskRequestMultiStreamURL
	SetSeiParams(v *UpdateLiveMPUTaskRequestSeiParams) *UpdateLiveMPUTaskRequest
	GetSeiParams() *UpdateLiveMPUTaskRequestSeiParams
	SetSingleSubParams(v *UpdateLiveMPUTaskRequestSingleSubParams) *UpdateLiveMPUTaskRequest
	GetSingleSubParams() *UpdateLiveMPUTaskRequestSingleSubParams
	SetStreamURL(v string) *UpdateLiveMPUTaskRequest
	GetStreamURL() *string
	SetTaskId(v string) *UpdateLiveMPUTaskRequest
	GetTaskId() *string
	SetTranscodeParams(v *UpdateLiveMPUTaskRequestTranscodeParams) *UpdateLiveMPUTaskRequest
	GetTranscodeParams() *UpdateLiveMPUTaskRequestTranscodeParams
}

type UpdateLiveMPUTaskRequest struct {
	// The application ID. You can specify only one application ID. The ID can be up to 64 characters in length and can contain letters, digits, underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The channel ID. You can specify only one channel ID. The ID can be up to 64 characters in length and can contain letters, digits, underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// yourChannelId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The stream mixing mode. Valid values:
	//
	// 	- **0**: the single-stream relay mode. In this mode, the service only relays the original single stream, but does not transcode mixed streams. You do not need to set parameters for mixed-stream transcoding.
	//
	// 	- **1*	- (default): the mixed-stream relay mode.
	//
	// example:
	//
	// 0
	MixMode *string `json:"MixMode,omitempty" xml:"MixMode,omitempty"`
	// The multiple ingest URLs to relay. This parameter allows you to specify multiple ingest URLs.
	MultiStreamURL []*UpdateLiveMPUTaskRequestMultiStreamURL `json:"MultiStreamURL,omitempty" xml:"MultiStreamURL,omitempty" type:"Repeated"`
	// The supplemental enhancement information (SEI) parameters.
	SeiParams *UpdateLiveMPUTaskRequestSeiParams `json:"SeiParams,omitempty" xml:"SeiParams,omitempty" type:"Struct"`
	// The single-stream relay parameters. These parameters are required if you set MixMode to 0.
	SingleSubParams *UpdateLiveMPUTaskRequestSingleSubParams `json:"SingleSubParams,omitempty" xml:"SingleSubParams,omitempty" type:"Struct"`
	// The ingest URL. You can specify only one ingest URL in the Real-Time Messaging Protocol (RTMP) format. The URL can be up to 2,048 characters in length. For information about the generation rules of ingest URLs, see [Ingest and streaming URLs](https://help.aliyun.com/document_detail/199339.html).
	//
	// >
	//
	// 	- If the ingest URL is under a domain name for which hotlink protection is enabled, you must include an access token in the URL.
	//
	// 	- You cannot use the same ingest URL in different tasks.
	//
	// 	- You cannot use the same ingest URL within 10 seconds after a task is stopped.
	//
	// example:
	//
	// rtmp://example.com/live/stream
	StreamURL *string `json:"StreamURL,omitempty" xml:"StreamURL,omitempty"`
	// The task ID. You can specify only one task ID. The ID can be up to 55 characters in length and can contain letters, digits, underscores (_), and hyphens (-). The ID must be unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourTaskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The mixed-stream relay parameters. These parameters are required if you set MixMode to 1.
	TranscodeParams *UpdateLiveMPUTaskRequestTranscodeParams `json:"TranscodeParams,omitempty" xml:"TranscodeParams,omitempty" type:"Struct"`
}

func (s UpdateLiveMPUTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveMPUTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateLiveMPUTaskRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateLiveMPUTaskRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *UpdateLiveMPUTaskRequest) GetMixMode() *string {
	return s.MixMode
}

func (s *UpdateLiveMPUTaskRequest) GetMultiStreamURL() []*UpdateLiveMPUTaskRequestMultiStreamURL {
	return s.MultiStreamURL
}

func (s *UpdateLiveMPUTaskRequest) GetSeiParams() *UpdateLiveMPUTaskRequestSeiParams {
	return s.SeiParams
}

func (s *UpdateLiveMPUTaskRequest) GetSingleSubParams() *UpdateLiveMPUTaskRequestSingleSubParams {
	return s.SingleSubParams
}

func (s *UpdateLiveMPUTaskRequest) GetStreamURL() *string {
	return s.StreamURL
}

func (s *UpdateLiveMPUTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateLiveMPUTaskRequest) GetTranscodeParams() *UpdateLiveMPUTaskRequestTranscodeParams {
	return s.TranscodeParams
}

func (s *UpdateLiveMPUTaskRequest) SetAppId(v string) *UpdateLiveMPUTaskRequest {
	s.AppId = &v
	return s
}

func (s *UpdateLiveMPUTaskRequest) SetChannelId(v string) *UpdateLiveMPUTaskRequest {
	s.ChannelId = &v
	return s
}

func (s *UpdateLiveMPUTaskRequest) SetMixMode(v string) *UpdateLiveMPUTaskRequest {
	s.MixMode = &v
	return s
}

func (s *UpdateLiveMPUTaskRequest) SetMultiStreamURL(v []*UpdateLiveMPUTaskRequestMultiStreamURL) *UpdateLiveMPUTaskRequest {
	s.MultiStreamURL = v
	return s
}

func (s *UpdateLiveMPUTaskRequest) SetSeiParams(v *UpdateLiveMPUTaskRequestSeiParams) *UpdateLiveMPUTaskRequest {
	s.SeiParams = v
	return s
}

func (s *UpdateLiveMPUTaskRequest) SetSingleSubParams(v *UpdateLiveMPUTaskRequestSingleSubParams) *UpdateLiveMPUTaskRequest {
	s.SingleSubParams = v
	return s
}

func (s *UpdateLiveMPUTaskRequest) SetStreamURL(v string) *UpdateLiveMPUTaskRequest {
	s.StreamURL = &v
	return s
}

func (s *UpdateLiveMPUTaskRequest) SetTaskId(v string) *UpdateLiveMPUTaskRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateLiveMPUTaskRequest) SetTranscodeParams(v *UpdateLiveMPUTaskRequestTranscodeParams) *UpdateLiveMPUTaskRequest {
	s.TranscodeParams = v
	return s
}

func (s *UpdateLiveMPUTaskRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateLiveMPUTaskRequestMultiStreamURL struct {
	// Specifies whether to perform stream relay by using Alibaba Cloud CDN. Valid values:
	//
	// 	- false: performs stream relay by using a CDN service that is not Alibaba Cloud CDN.
	//
	// 	- true: performs stream relay by using Alibaba Cloud CDN.
	//
	// >  The default value of this parameter is false.
	//
	// example:
	//
	// false
	IsAliCdn *bool `json:"IsAliCdn,omitempty" xml:"IsAliCdn,omitempty"`
	// The ingest URL. Only the RTMP format is supported. The URL can be up to 2,048 characters in length. For information about the generation rules of ingest URLs, see [Ingest and streaming URLs](https://help.aliyun.com/document_detail/199339.html).
	//
	// example:
	//
	// rtmp://example.com/live/stream****
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s UpdateLiveMPUTaskRequestMultiStreamURL) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveMPUTaskRequestMultiStreamURL) GoString() string {
	return s.String()
}

func (s *UpdateLiveMPUTaskRequestMultiStreamURL) GetIsAliCdn() *bool {
	return s.IsAliCdn
}

func (s *UpdateLiveMPUTaskRequestMultiStreamURL) GetURL() *string {
	return s.URL
}

func (s *UpdateLiveMPUTaskRequestMultiStreamURL) SetIsAliCdn(v bool) *UpdateLiveMPUTaskRequestMultiStreamURL {
	s.IsAliCdn = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestMultiStreamURL) SetURL(v string) *UpdateLiveMPUTaskRequestMultiStreamURL {
	s.URL = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestMultiStreamURL) Validate() error {
	return dara.Validate(s)
}

type UpdateLiveMPUTaskRequestSeiParams struct {
	// The layout and volume SEI. If you leave this parameter empty, the default layout and volume SEI is used.
	LayoutVolume *UpdateLiveMPUTaskRequestSeiParamsLayoutVolume `json:"LayoutVolume,omitempty" xml:"LayoutVolume,omitempty" type:"Struct"`
	// Specifies whether to pass through the SEI.
	PassThrough *UpdateLiveMPUTaskRequestSeiParamsPassThrough `json:"PassThrough,omitempty" xml:"PassThrough,omitempty" type:"Struct"`
}

func (s UpdateLiveMPUTaskRequestSeiParams) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveMPUTaskRequestSeiParams) GoString() string {
	return s.String()
}

func (s *UpdateLiveMPUTaskRequestSeiParams) GetLayoutVolume() *UpdateLiveMPUTaskRequestSeiParamsLayoutVolume {
	return s.LayoutVolume
}

func (s *UpdateLiveMPUTaskRequestSeiParams) GetPassThrough() *UpdateLiveMPUTaskRequestSeiParamsPassThrough {
	return s.PassThrough
}

func (s *UpdateLiveMPUTaskRequestSeiParams) SetLayoutVolume(v *UpdateLiveMPUTaskRequestSeiParamsLayoutVolume) *UpdateLiveMPUTaskRequestSeiParams {
	s.LayoutVolume = v
	return s
}

func (s *UpdateLiveMPUTaskRequestSeiParams) SetPassThrough(v *UpdateLiveMPUTaskRequestSeiParamsPassThrough) *UpdateLiveMPUTaskRequestSeiParams {
	s.PassThrough = v
	return s
}

func (s *UpdateLiveMPUTaskRequestSeiParams) Validate() error {
	return dara.Validate(s)
}

type UpdateLiveMPUTaskRequestSeiParamsLayoutVolume struct {
	// Specifies whether to include the SEI in an Instantaneous Decoder Refresh (IDR) frame. Valid values:
	//
	// 	- **0**: does not include the SEI.
	//
	// 	- **1**: includes the SEI.
	//
	// example:
	//
	// 0
	FollowIdr *string `json:"FollowIdr,omitempty" xml:"FollowIdr,omitempty"`
	// The interval at which the SEI is sent. Valid values: [1000,5000]. Unit: milliseconds.
	//
	// example:
	//
	// 1000
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s UpdateLiveMPUTaskRequestSeiParamsLayoutVolume) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveMPUTaskRequestSeiParamsLayoutVolume) GoString() string {
	return s.String()
}

func (s *UpdateLiveMPUTaskRequestSeiParamsLayoutVolume) GetFollowIdr() *string {
	return s.FollowIdr
}

func (s *UpdateLiveMPUTaskRequestSeiParamsLayoutVolume) GetInterval() *string {
	return s.Interval
}

func (s *UpdateLiveMPUTaskRequestSeiParamsLayoutVolume) SetFollowIdr(v string) *UpdateLiveMPUTaskRequestSeiParamsLayoutVolume {
	s.FollowIdr = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestSeiParamsLayoutVolume) SetInterval(v string) *UpdateLiveMPUTaskRequestSeiParamsLayoutVolume {
	s.Interval = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestSeiParamsLayoutVolume) Validate() error {
	return dara.Validate(s)
}

type UpdateLiveMPUTaskRequestSeiParamsPassThrough struct {
	// Specifies whether to include the SEI in an IDR frame. Valid values:
	//
	// 	- **0**: does not include the SEI.
	//
	// 	- **1**: includes the SEI.
	//
	// example:
	//
	// 0
	FollowIdr *string `json:"FollowIdr,omitempty" xml:"FollowIdr,omitempty"`
	// The interval at which the SEI is sent. Valid values: [1000,5000]. Unit: milliseconds.
	//
	// example:
	//
	// 1000
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The payload content of the SEI.
	//
	// example:
	//
	// yourPayloadContent
	PayloadContent *string `json:"PayloadContent,omitempty" xml:"PayloadContent,omitempty"`
	// The key of the payload content of the SEI. If you do not specify this parameter, the default value udd is used.
	//
	// example:
	//
	// yourPayloadContentKey
	PayloadContentKey *string `json:"PayloadContentKey,omitempty" xml:"PayloadContentKey,omitempty"`
}

func (s UpdateLiveMPUTaskRequestSeiParamsPassThrough) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveMPUTaskRequestSeiParamsPassThrough) GoString() string {
	return s.String()
}

func (s *UpdateLiveMPUTaskRequestSeiParamsPassThrough) GetFollowIdr() *string {
	return s.FollowIdr
}

func (s *UpdateLiveMPUTaskRequestSeiParamsPassThrough) GetInterval() *string {
	return s.Interval
}

func (s *UpdateLiveMPUTaskRequestSeiParamsPassThrough) GetPayloadContent() *string {
	return s.PayloadContent
}

func (s *UpdateLiveMPUTaskRequestSeiParamsPassThrough) GetPayloadContentKey() *string {
	return s.PayloadContentKey
}

func (s *UpdateLiveMPUTaskRequestSeiParamsPassThrough) SetFollowIdr(v string) *UpdateLiveMPUTaskRequestSeiParamsPassThrough {
	s.FollowIdr = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestSeiParamsPassThrough) SetInterval(v string) *UpdateLiveMPUTaskRequestSeiParamsPassThrough {
	s.Interval = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestSeiParamsPassThrough) SetPayloadContent(v string) *UpdateLiveMPUTaskRequestSeiParamsPassThrough {
	s.PayloadContent = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestSeiParamsPassThrough) SetPayloadContentKey(v string) *UpdateLiveMPUTaskRequestSeiParamsPassThrough {
	s.PayloadContentKey = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestSeiParamsPassThrough) Validate() error {
	return dara.Validate(s)
}

type UpdateLiveMPUTaskRequestSingleSubParams struct {
	// The type of the video source. This parameter is valid only when you set StreamType to 2. Valid values:
	//
	// 	- **camera*	- (default)
	//
	// 	- **shareScreen**
	//
	// example:
	//
	// camera
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The type of the stream that you want to relay. Valid values:
	//
	// 	- **0*	- (default): original stream
	//
	// 	- **1**: only the audio track
	//
	// 	- **2**: only the video track
	//
	// example:
	//
	// 0
	StreamType *string `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	// The user ID. In the single-stream relay mode, you can relay only one stream in a request.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourSubUserId
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateLiveMPUTaskRequestSingleSubParams) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveMPUTaskRequestSingleSubParams) GoString() string {
	return s.String()
}

func (s *UpdateLiveMPUTaskRequestSingleSubParams) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateLiveMPUTaskRequestSingleSubParams) GetStreamType() *string {
	return s.StreamType
}

func (s *UpdateLiveMPUTaskRequestSingleSubParams) GetUserId() *string {
	return s.UserId
}

func (s *UpdateLiveMPUTaskRequestSingleSubParams) SetSourceType(v string) *UpdateLiveMPUTaskRequestSingleSubParams {
	s.SourceType = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestSingleSubParams) SetStreamType(v string) *UpdateLiveMPUTaskRequestSingleSubParams {
	s.StreamType = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestSingleSubParams) SetUserId(v string) *UpdateLiveMPUTaskRequestSingleSubParams {
	s.UserId = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestSingleSubParams) Validate() error {
	return dara.Validate(s)
}

type UpdateLiveMPUTaskRequestTranscodeParams struct {
	// The global background image.
	Background *UpdateLiveMPUTaskRequestTranscodeParamsBackground `json:"Background,omitempty" xml:"Background,omitempty" type:"Struct"`
	// The encoding parameters for the output stream.
	EncodeParams *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams `json:"EncodeParams,omitempty" xml:"EncodeParams,omitempty" type:"Struct"`
	// The video layout information.
	//
	// >  If video transcoding is required, you must specify the video layout information, including the x-coordinate and y-coordinate, the width and height, and the layer. For audio-only transcoding, leave the video layout information empty.
	Layout *UpdateLiveMPUTaskRequestTranscodeParamsLayout `json:"Layout,omitempty" xml:"Layout,omitempty" type:"Struct"`
	// The information about the users whose streams are subscribed to. If you leave this parameter empty, streams from all users are mixed.
	UserInfos []*UpdateLiveMPUTaskRequestTranscodeParamsUserInfos `json:"UserInfos,omitempty" xml:"UserInfos,omitempty" type:"Repeated"`
}

func (s UpdateLiveMPUTaskRequestTranscodeParams) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveMPUTaskRequestTranscodeParams) GoString() string {
	return s.String()
}

func (s *UpdateLiveMPUTaskRequestTranscodeParams) GetBackground() *UpdateLiveMPUTaskRequestTranscodeParamsBackground {
	return s.Background
}

func (s *UpdateLiveMPUTaskRequestTranscodeParams) GetEncodeParams() *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams {
	return s.EncodeParams
}

func (s *UpdateLiveMPUTaskRequestTranscodeParams) GetLayout() *UpdateLiveMPUTaskRequestTranscodeParamsLayout {
	return s.Layout
}

func (s *UpdateLiveMPUTaskRequestTranscodeParams) GetUserInfos() []*UpdateLiveMPUTaskRequestTranscodeParamsUserInfos {
	return s.UserInfos
}

func (s *UpdateLiveMPUTaskRequestTranscodeParams) SetBackground(v *UpdateLiveMPUTaskRequestTranscodeParamsBackground) *UpdateLiveMPUTaskRequestTranscodeParams {
	s.Background = v
	return s
}

func (s *UpdateLiveMPUTaskRequestTranscodeParams) SetEncodeParams(v *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams) *UpdateLiveMPUTaskRequestTranscodeParams {
	s.EncodeParams = v
	return s
}

func (s *UpdateLiveMPUTaskRequestTranscodeParams) SetLayout(v *UpdateLiveMPUTaskRequestTranscodeParamsLayout) *UpdateLiveMPUTaskRequestTranscodeParams {
	s.Layout = v
	return s
}

func (s *UpdateLiveMPUTaskRequestTranscodeParams) SetUserInfos(v []*UpdateLiveMPUTaskRequestTranscodeParamsUserInfos) *UpdateLiveMPUTaskRequestTranscodeParams {
	s.UserInfos = v
	return s
}

func (s *UpdateLiveMPUTaskRequestTranscodeParams) Validate() error {
	return dara.Validate(s)
}

type UpdateLiveMPUTaskRequestTranscodeParamsBackground struct {
	// The display mode of the global background image.
	//
	// 	- **0**: scales the background image proportionally to fit the view, with black bars displayed.
	//
	// 	- **1*	- (default): crops the background image to fit the view.
	//
	// example:
	//
	// 1
	RenderMode *string `json:"RenderMode,omitempty" xml:"RenderMode,omitempty"`
	// The URL of the global background image.
	//
	// example:
	//
	// yourImageUrl
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s UpdateLiveMPUTaskRequestTranscodeParamsBackground) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveMPUTaskRequestTranscodeParamsBackground) GoString() string {
	return s.String()
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsBackground) GetRenderMode() *string {
	return s.RenderMode
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsBackground) GetURL() *string {
	return s.URL
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsBackground) SetRenderMode(v string) *UpdateLiveMPUTaskRequestTranscodeParamsBackground {
	s.RenderMode = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsBackground) SetURL(v string) *UpdateLiveMPUTaskRequestTranscodeParamsBackground {
	s.URL = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsBackground) Validate() error {
	return dara.Validate(s)
}

type UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams struct {
	// The bitrate of the audio. Valid values: [8,500]. Unit: Kbit/s.
	//
	// example:
	//
	// 128
	AudioBitrate *string `json:"AudioBitrate,omitempty" xml:"AudioBitrate,omitempty"`
	// The number of sound channels. Valid values: 1 and 2.
	//
	// example:
	//
	// 2
	AudioChannels *string `json:"AudioChannels,omitempty" xml:"AudioChannels,omitempty"`
	// Specifies whether the output stream is an audio-only stream. Valid values:
	//
	// 	- **true**: The output stream is an audio-only stream. If you set this parameter to true, you need to configure only audio-related parameters under EncodeParams.
	//
	// 	- **false*	- (default): The output stream is not an audio-only stream. If you set this parameter to false, you need to configure all parameters under EncodeParams, except the VideoCodec and EnhancedParam parameters.
	//
	// example:
	//
	// false
	AudioOnly *string `json:"AudioOnly,omitempty" xml:"AudioOnly,omitempty"`
	// The audio sampling rate. Valid values: 8000, 16000, 32000, 44100, and 48000. Unit: Hz.
	//
	// example:
	//
	// 44100
	AudioSampleRate *string `json:"AudioSampleRate,omitempty" xml:"AudioSampleRate,omitempty"`
	// The parameter used for encoding enhancement, which is a JSON string. The parameter includes the optional profile and preset fields.
	//
	// 	- profile: the encoding level. If the video codec is H.264, the valid values of this field are baseline, main, and high. If the video codec is H.265, the valid value of this field is main.
	//
	// 	- preset: adjusts the trade-off between encoding speed and video quality. The valid values of this field are ultrafast, superfast, veryfast, faster, fast, medium, slow, slower, veryslow, and placebo. Each value specifies a level of trade-off between encoding speed and video quality. For example, the ultrafast preset has the fastest encoding speed but the lowest video quality, while the placebo preset sacrifices the encoding speed for the best video quality.
	//
	// >  A value of superfast for the preset field is suitable for real-time communication scenarios. We recommend that you not set the field if you are not a professional encoding engineer.
	//
	// example:
	//
	// {"profile": "high", "preset": "veryfast"}
	EnhancedParam *string `json:"EnhancedParam,omitempty" xml:"EnhancedParam,omitempty"`
	// The bitrate of the video. Valid values: [1,10000]. Unit: Kbit/s.
	//
	// example:
	//
	// 3500
	VideoBitrate *string `json:"VideoBitrate,omitempty" xml:"VideoBitrate,omitempty"`
	// The video codec. Valid values:
	//
	// 	- H.264 (default)
	//
	// 	- H.265
	//
	// example:
	//
	// H.264
	VideoCodec *string `json:"VideoCodec,omitempty" xml:"VideoCodec,omitempty"`
	// The frame rate of the video. Valid values: [1,60]. Unit: frames per second (FPS).
	//
	// example:
	//
	// 25
	VideoFramerate *string `json:"VideoFramerate,omitempty" xml:"VideoFramerate,omitempty"`
	// The group of pictures (GOP) size of the video. Valid values: [1,60].
	//
	// example:
	//
	// 20
	VideoGop *string `json:"VideoGop,omitempty" xml:"VideoGop,omitempty"`
	// The height of the video. Valid values: [0,1920]. Unit: pixels.
	//
	// example:
	//
	// 1000
	VideoHeight *string `json:"VideoHeight,omitempty" xml:"VideoHeight,omitempty"`
	// The width of the video. Valid values: [0,1920]. Unit: pixels.
	//
	// example:
	//
	// 1920
	VideoWidth *string `json:"VideoWidth,omitempty" xml:"VideoWidth,omitempty"`
}

func (s UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams) GoString() string {
	return s.String()
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams) GetAudioBitrate() *string {
	return s.AudioBitrate
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams) GetAudioChannels() *string {
	return s.AudioChannels
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams) GetAudioOnly() *string {
	return s.AudioOnly
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams) GetAudioSampleRate() *string {
	return s.AudioSampleRate
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams) GetEnhancedParam() *string {
	return s.EnhancedParam
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams) GetVideoBitrate() *string {
	return s.VideoBitrate
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams) GetVideoCodec() *string {
	return s.VideoCodec
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams) GetVideoFramerate() *string {
	return s.VideoFramerate
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams) GetVideoGop() *string {
	return s.VideoGop
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams) GetVideoHeight() *string {
	return s.VideoHeight
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams) GetVideoWidth() *string {
	return s.VideoWidth
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams) SetAudioBitrate(v string) *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams {
	s.AudioBitrate = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams) SetAudioChannels(v string) *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams {
	s.AudioChannels = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams) SetAudioOnly(v string) *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams {
	s.AudioOnly = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams) SetAudioSampleRate(v string) *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams {
	s.AudioSampleRate = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams) SetEnhancedParam(v string) *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams {
	s.EnhancedParam = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams) SetVideoBitrate(v string) *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams {
	s.VideoBitrate = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams) SetVideoCodec(v string) *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams {
	s.VideoCodec = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams) SetVideoFramerate(v string) *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams {
	s.VideoFramerate = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams) SetVideoGop(v string) *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams {
	s.VideoGop = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams) SetVideoHeight(v string) *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams {
	s.VideoHeight = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams) SetVideoWidth(v string) *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams {
	s.VideoWidth = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsEncodeParams) Validate() error {
	return dara.Validate(s)
}

type UpdateLiveMPUTaskRequestTranscodeParamsLayout struct {
	// The information about the panes.
	UserPanes []*UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanes `json:"UserPanes,omitempty" xml:"UserPanes,omitempty" type:"Repeated"`
}

func (s UpdateLiveMPUTaskRequestTranscodeParamsLayout) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveMPUTaskRequestTranscodeParamsLayout) GoString() string {
	return s.String()
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsLayout) GetUserPanes() []*UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanes {
	return s.UserPanes
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsLayout) SetUserPanes(v []*UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) *UpdateLiveMPUTaskRequestTranscodeParamsLayout {
	s.UserPanes = v
	return s
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsLayout) Validate() error {
	return dara.Validate(s)
}

type UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanes struct {
	// The URL of the background image of the pane. This image is displayed if the user turns off the camera or is not present in the channel.
	//
	// example:
	//
	// yourImageUrl
	BackgroundImageUrl *string `json:"BackgroundImageUrl,omitempty" xml:"BackgroundImageUrl,omitempty"`
	// The height of the pane. The value is normalized.
	//
	// example:
	//
	// 0.2632
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The display mode of the pane. Valid values:
	//
	// 	- **0**: scales the video proportionally to fit the view, with black bars displayed.
	//
	// 	- **1 (default)**: crops the video to fit the view.
	//
	// example:
	//
	// 1
	RenderMode *string `json:"RenderMode,omitempty" xml:"RenderMode,omitempty"`
	// The information about the user whose stream is played in the pane. If you leave this parameter empty, the system automatically sets this parameter based on the order in which streamers join the channel.
	//
	// >
	//
	// 	- If you specify the information about a user by using this parameter, the information about the user must also be specified by using the TranscodeParams.UserInfos parameter.
	//
	// 	- This parameter is valid only when you set StreamType to 0 or 2.
	UserInfo *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanesUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
	// The width of the pane. The value is normalized.
	//
	// example:
	//
	// 0.3564
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
	// The x-coordinate of the pane. The value is normalized.
	//
	// example:
	//
	// 0.2456
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	// The y-coordinate of the pane. The value is normalized.
	//
	// example:
	//
	// 0.3789
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
	// The layer in which the pane resides. A value of 0 indicates the bottom layer. Each increment of the value by 1 indicates the next upper layer.
	//
	// example:
	//
	// 0
	ZOrder *string `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) GoString() string {
	return s.String()
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) GetBackgroundImageUrl() *string {
	return s.BackgroundImageUrl
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) GetHeight() *string {
	return s.Height
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) GetRenderMode() *string {
	return s.RenderMode
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) GetUserInfo() *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanesUserInfo {
	return s.UserInfo
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) GetWidth() *string {
	return s.Width
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) GetX() *string {
	return s.X
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) GetY() *string {
	return s.Y
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) GetZOrder() *string {
	return s.ZOrder
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) SetBackgroundImageUrl(v string) *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanes {
	s.BackgroundImageUrl = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) SetHeight(v string) *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanes {
	s.Height = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) SetRenderMode(v string) *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanes {
	s.RenderMode = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) SetUserInfo(v *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanesUserInfo) *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanes {
	s.UserInfo = v
	return s
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) SetWidth(v string) *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanes {
	s.Width = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) SetX(v string) *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanes {
	s.X = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) SetY(v string) *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanes {
	s.Y = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) SetZOrder(v string) *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanes {
	s.ZOrder = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) Validate() error {
	return dara.Validate(s)
}

type UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanesUserInfo struct {
	// The ID of the channel where the user is. If the user is in the same channel, you can leave this parameter empty. We recommend that you specify this parameter when you perform stream mixing across channels.
	//
	// example:
	//
	// yourChannelId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The type of the video source. This parameter is valid only when you set StreamType to 2. Valid values:
	//
	// 	- **camera*	- (default)
	//
	// 	- **shareScreen**
	//
	// example:
	//
	// camera
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The user ID.
	//
	// example:
	//
	// yourSubUserId
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanesUserInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanesUserInfo) GoString() string {
	return s.String()
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanesUserInfo) GetChannelId() *string {
	return s.ChannelId
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanesUserInfo) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanesUserInfo) GetUserId() *string {
	return s.UserId
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanesUserInfo) SetChannelId(v string) *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanesUserInfo {
	s.ChannelId = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanesUserInfo) SetSourceType(v string) *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanesUserInfo {
	s.SourceType = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanesUserInfo) SetUserId(v string) *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanesUserInfo {
	s.UserId = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsLayoutUserPanesUserInfo) Validate() error {
	return dara.Validate(s)
}

type UpdateLiveMPUTaskRequestTranscodeParamsUserInfos struct {
	// The ID of the channel where the subscribed user is. If the user is in the same channel, you can leave this parameter empty. We recommend that you specify this parameter when you perform stream mixing across channels.
	//
	// example:
	//
	// yourChannelId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The type of the video source that is subscribed to. This parameter is valid only when you set StreamType to 2. Valid values:
	//
	// 	- **camera*	- (default)
	//
	// 	- **shareScreen**
	//
	// example:
	//
	// camera
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The type of the relayed stream that is subscribed to. Valid values:
	//
	// 	- **0*	- (default): original stream
	//
	// 	- **1**: only the audio track
	//
	// 	- **2**: only the video track
	//
	// example:
	//
	// 0
	StreamType *string `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	// The ID of the subscribed user.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourSubUserId
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateLiveMPUTaskRequestTranscodeParamsUserInfos) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveMPUTaskRequestTranscodeParamsUserInfos) GoString() string {
	return s.String()
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsUserInfos) GetChannelId() *string {
	return s.ChannelId
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsUserInfos) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsUserInfos) GetStreamType() *string {
	return s.StreamType
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsUserInfos) GetUserId() *string {
	return s.UserId
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsUserInfos) SetChannelId(v string) *UpdateLiveMPUTaskRequestTranscodeParamsUserInfos {
	s.ChannelId = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsUserInfos) SetSourceType(v string) *UpdateLiveMPUTaskRequestTranscodeParamsUserInfos {
	s.SourceType = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsUserInfos) SetStreamType(v string) *UpdateLiveMPUTaskRequestTranscodeParamsUserInfos {
	s.StreamType = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsUserInfos) SetUserId(v string) *UpdateLiveMPUTaskRequestTranscodeParamsUserInfos {
	s.UserId = &v
	return s
}

func (s *UpdateLiveMPUTaskRequestTranscodeParamsUserInfos) Validate() error {
	return dara.Validate(s)
}
