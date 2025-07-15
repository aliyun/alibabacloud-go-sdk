// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartLiveMPUTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StartLiveMPUTaskRequest
	GetAppId() *string
	SetChannelId(v string) *StartLiveMPUTaskRequest
	GetChannelId() *string
	SetMaxIdleTime(v string) *StartLiveMPUTaskRequest
	GetMaxIdleTime() *string
	SetMixMode(v string) *StartLiveMPUTaskRequest
	GetMixMode() *string
	SetMultiStreamURL(v []*StartLiveMPUTaskRequestMultiStreamURL) *StartLiveMPUTaskRequest
	GetMultiStreamURL() []*StartLiveMPUTaskRequestMultiStreamURL
	SetRegion(v string) *StartLiveMPUTaskRequest
	GetRegion() *string
	SetSeiParams(v *StartLiveMPUTaskRequestSeiParams) *StartLiveMPUTaskRequest
	GetSeiParams() *StartLiveMPUTaskRequestSeiParams
	SetSingleSubParams(v *StartLiveMPUTaskRequestSingleSubParams) *StartLiveMPUTaskRequest
	GetSingleSubParams() *StartLiveMPUTaskRequestSingleSubParams
	SetStreamURL(v string) *StartLiveMPUTaskRequest
	GetStreamURL() *string
	SetTaskId(v string) *StartLiveMPUTaskRequest
	GetTaskId() *string
	SetTranscodeParams(v *StartLiveMPUTaskRequestTranscodeParams) *StartLiveMPUTaskRequest
	GetTranscodeParams() *StartLiveMPUTaskRequestTranscodeParams
}

type StartLiveMPUTaskRequest struct {
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
	// The timeout period of an idle connection. Unit: seconds. Valid values: [10,86400].
	//
	// >  If the task is idle for a period of time longer than the duration specified by the MaxIdleTime parameter, the task is automatically stopped. If the parameter is not specified, the task is stopped after the channel is closed.
	//
	// example:
	//
	// 10
	MaxIdleTime *string `json:"MaxIdleTime,omitempty" xml:"MaxIdleTime,omitempty"`
	// The stream mixing mode. Valid values:
	//
	// 	- **0**: the single-stream relay mode. In this mode, the service only relays the original single stream, but does not transcode mixed streams. You do not need to set parameters for mixed-stream transcoding.
	//
	// 	- **1*	- (default): the mixed-stream relay mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	MixMode *string `json:"MixMode,omitempty" xml:"MixMode,omitempty"`
	// The multiple ingest URLs to relay. This parameter allows you to specify multiple ingest URLs.
	//
	// >  The StreamURL and MultiStreamURL parameters are mutually exclusive. You must specify one of the two parameters.
	MultiStreamURL []*StartLiveMPUTaskRequestMultiStreamURL `json:"MultiStreamURL,omitempty" xml:"MultiStreamURL,omitempty" type:"Repeated"`
	// The region in which the streams are mixed. Valid values:
	//
	// 	- **CN-Shanghai**
	//
	// 	- **AP-Singapore*	- (default)
	//
	// 	- **EMAA-Saudi**
	//
	// example:
	//
	// CN-Shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The supplemental enhancement information (SEI) parameters.
	SeiParams *StartLiveMPUTaskRequestSeiParams `json:"SeiParams,omitempty" xml:"SeiParams,omitempty" type:"Struct"`
	// The single-stream relay parameters. These parameters are required if you set MixMode to 0. Leave these parameters empty in the mixed-stream relay mode.
	SingleSubParams *StartLiveMPUTaskRequestSingleSubParams `json:"SingleSubParams,omitempty" xml:"SingleSubParams,omitempty" type:"Struct"`
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
	// The mixed-stream relay parameters. These parameters are required if you set MixMode to 1. Leave these parameters empty if you use the single-stream relay mode.
	TranscodeParams *StartLiveMPUTaskRequestTranscodeParams `json:"TranscodeParams,omitempty" xml:"TranscodeParams,omitempty" type:"Struct"`
}

func (s StartLiveMPUTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StartLiveMPUTaskRequest) GoString() string {
	return s.String()
}

func (s *StartLiveMPUTaskRequest) GetAppId() *string {
	return s.AppId
}

func (s *StartLiveMPUTaskRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *StartLiveMPUTaskRequest) GetMaxIdleTime() *string {
	return s.MaxIdleTime
}

func (s *StartLiveMPUTaskRequest) GetMixMode() *string {
	return s.MixMode
}

func (s *StartLiveMPUTaskRequest) GetMultiStreamURL() []*StartLiveMPUTaskRequestMultiStreamURL {
	return s.MultiStreamURL
}

func (s *StartLiveMPUTaskRequest) GetRegion() *string {
	return s.Region
}

func (s *StartLiveMPUTaskRequest) GetSeiParams() *StartLiveMPUTaskRequestSeiParams {
	return s.SeiParams
}

func (s *StartLiveMPUTaskRequest) GetSingleSubParams() *StartLiveMPUTaskRequestSingleSubParams {
	return s.SingleSubParams
}

func (s *StartLiveMPUTaskRequest) GetStreamURL() *string {
	return s.StreamURL
}

func (s *StartLiveMPUTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *StartLiveMPUTaskRequest) GetTranscodeParams() *StartLiveMPUTaskRequestTranscodeParams {
	return s.TranscodeParams
}

func (s *StartLiveMPUTaskRequest) SetAppId(v string) *StartLiveMPUTaskRequest {
	s.AppId = &v
	return s
}

func (s *StartLiveMPUTaskRequest) SetChannelId(v string) *StartLiveMPUTaskRequest {
	s.ChannelId = &v
	return s
}

func (s *StartLiveMPUTaskRequest) SetMaxIdleTime(v string) *StartLiveMPUTaskRequest {
	s.MaxIdleTime = &v
	return s
}

func (s *StartLiveMPUTaskRequest) SetMixMode(v string) *StartLiveMPUTaskRequest {
	s.MixMode = &v
	return s
}

func (s *StartLiveMPUTaskRequest) SetMultiStreamURL(v []*StartLiveMPUTaskRequestMultiStreamURL) *StartLiveMPUTaskRequest {
	s.MultiStreamURL = v
	return s
}

func (s *StartLiveMPUTaskRequest) SetRegion(v string) *StartLiveMPUTaskRequest {
	s.Region = &v
	return s
}

func (s *StartLiveMPUTaskRequest) SetSeiParams(v *StartLiveMPUTaskRequestSeiParams) *StartLiveMPUTaskRequest {
	s.SeiParams = v
	return s
}

func (s *StartLiveMPUTaskRequest) SetSingleSubParams(v *StartLiveMPUTaskRequestSingleSubParams) *StartLiveMPUTaskRequest {
	s.SingleSubParams = v
	return s
}

func (s *StartLiveMPUTaskRequest) SetStreamURL(v string) *StartLiveMPUTaskRequest {
	s.StreamURL = &v
	return s
}

func (s *StartLiveMPUTaskRequest) SetTaskId(v string) *StartLiveMPUTaskRequest {
	s.TaskId = &v
	return s
}

func (s *StartLiveMPUTaskRequest) SetTranscodeParams(v *StartLiveMPUTaskRequestTranscodeParams) *StartLiveMPUTaskRequest {
	s.TranscodeParams = v
	return s
}

func (s *StartLiveMPUTaskRequest) Validate() error {
	return dara.Validate(s)
}

type StartLiveMPUTaskRequestMultiStreamURL struct {
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

func (s StartLiveMPUTaskRequestMultiStreamURL) String() string {
	return dara.Prettify(s)
}

func (s StartLiveMPUTaskRequestMultiStreamURL) GoString() string {
	return s.String()
}

func (s *StartLiveMPUTaskRequestMultiStreamURL) GetIsAliCdn() *bool {
	return s.IsAliCdn
}

func (s *StartLiveMPUTaskRequestMultiStreamURL) GetURL() *string {
	return s.URL
}

func (s *StartLiveMPUTaskRequestMultiStreamURL) SetIsAliCdn(v bool) *StartLiveMPUTaskRequestMultiStreamURL {
	s.IsAliCdn = &v
	return s
}

func (s *StartLiveMPUTaskRequestMultiStreamURL) SetURL(v string) *StartLiveMPUTaskRequestMultiStreamURL {
	s.URL = &v
	return s
}

func (s *StartLiveMPUTaskRequestMultiStreamURL) Validate() error {
	return dara.Validate(s)
}

type StartLiveMPUTaskRequestSeiParams struct {
	// The layout and volume SEI. If you leave this parameter empty, the default layout and volume SEI is used.
	LayoutVolume *StartLiveMPUTaskRequestSeiParamsLayoutVolume `json:"LayoutVolume,omitempty" xml:"LayoutVolume,omitempty" type:"Struct"`
	// Specifies whether to pass through the SEI.
	PassThrough *StartLiveMPUTaskRequestSeiParamsPassThrough `json:"PassThrough,omitempty" xml:"PassThrough,omitempty" type:"Struct"`
	// The custom payload_type of the SEI. Valid values: 100 to 254. If you do not specify this parameter, the default value 5 is used.
	//
	// example:
	//
	// 100
	PayloadType *string `json:"PayloadType,omitempty" xml:"PayloadType,omitempty"`
}

func (s StartLiveMPUTaskRequestSeiParams) String() string {
	return dara.Prettify(s)
}

func (s StartLiveMPUTaskRequestSeiParams) GoString() string {
	return s.String()
}

func (s *StartLiveMPUTaskRequestSeiParams) GetLayoutVolume() *StartLiveMPUTaskRequestSeiParamsLayoutVolume {
	return s.LayoutVolume
}

func (s *StartLiveMPUTaskRequestSeiParams) GetPassThrough() *StartLiveMPUTaskRequestSeiParamsPassThrough {
	return s.PassThrough
}

func (s *StartLiveMPUTaskRequestSeiParams) GetPayloadType() *string {
	return s.PayloadType
}

func (s *StartLiveMPUTaskRequestSeiParams) SetLayoutVolume(v *StartLiveMPUTaskRequestSeiParamsLayoutVolume) *StartLiveMPUTaskRequestSeiParams {
	s.LayoutVolume = v
	return s
}

func (s *StartLiveMPUTaskRequestSeiParams) SetPassThrough(v *StartLiveMPUTaskRequestSeiParamsPassThrough) *StartLiveMPUTaskRequestSeiParams {
	s.PassThrough = v
	return s
}

func (s *StartLiveMPUTaskRequestSeiParams) SetPayloadType(v string) *StartLiveMPUTaskRequestSeiParams {
	s.PayloadType = &v
	return s
}

func (s *StartLiveMPUTaskRequestSeiParams) Validate() error {
	return dara.Validate(s)
}

type StartLiveMPUTaskRequestSeiParamsLayoutVolume struct {
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

func (s StartLiveMPUTaskRequestSeiParamsLayoutVolume) String() string {
	return dara.Prettify(s)
}

func (s StartLiveMPUTaskRequestSeiParamsLayoutVolume) GoString() string {
	return s.String()
}

func (s *StartLiveMPUTaskRequestSeiParamsLayoutVolume) GetFollowIdr() *string {
	return s.FollowIdr
}

func (s *StartLiveMPUTaskRequestSeiParamsLayoutVolume) GetInterval() *string {
	return s.Interval
}

func (s *StartLiveMPUTaskRequestSeiParamsLayoutVolume) SetFollowIdr(v string) *StartLiveMPUTaskRequestSeiParamsLayoutVolume {
	s.FollowIdr = &v
	return s
}

func (s *StartLiveMPUTaskRequestSeiParamsLayoutVolume) SetInterval(v string) *StartLiveMPUTaskRequestSeiParamsLayoutVolume {
	s.Interval = &v
	return s
}

func (s *StartLiveMPUTaskRequestSeiParamsLayoutVolume) Validate() error {
	return dara.Validate(s)
}

type StartLiveMPUTaskRequestSeiParamsPassThrough struct {
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

func (s StartLiveMPUTaskRequestSeiParamsPassThrough) String() string {
	return dara.Prettify(s)
}

func (s StartLiveMPUTaskRequestSeiParamsPassThrough) GoString() string {
	return s.String()
}

func (s *StartLiveMPUTaskRequestSeiParamsPassThrough) GetFollowIdr() *string {
	return s.FollowIdr
}

func (s *StartLiveMPUTaskRequestSeiParamsPassThrough) GetInterval() *string {
	return s.Interval
}

func (s *StartLiveMPUTaskRequestSeiParamsPassThrough) GetPayloadContent() *string {
	return s.PayloadContent
}

func (s *StartLiveMPUTaskRequestSeiParamsPassThrough) GetPayloadContentKey() *string {
	return s.PayloadContentKey
}

func (s *StartLiveMPUTaskRequestSeiParamsPassThrough) SetFollowIdr(v string) *StartLiveMPUTaskRequestSeiParamsPassThrough {
	s.FollowIdr = &v
	return s
}

func (s *StartLiveMPUTaskRequestSeiParamsPassThrough) SetInterval(v string) *StartLiveMPUTaskRequestSeiParamsPassThrough {
	s.Interval = &v
	return s
}

func (s *StartLiveMPUTaskRequestSeiParamsPassThrough) SetPayloadContent(v string) *StartLiveMPUTaskRequestSeiParamsPassThrough {
	s.PayloadContent = &v
	return s
}

func (s *StartLiveMPUTaskRequestSeiParamsPassThrough) SetPayloadContentKey(v string) *StartLiveMPUTaskRequestSeiParamsPassThrough {
	s.PayloadContentKey = &v
	return s
}

func (s *StartLiveMPUTaskRequestSeiParamsPassThrough) Validate() error {
	return dara.Validate(s)
}

type StartLiveMPUTaskRequestSingleSubParams struct {
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

func (s StartLiveMPUTaskRequestSingleSubParams) String() string {
	return dara.Prettify(s)
}

func (s StartLiveMPUTaskRequestSingleSubParams) GoString() string {
	return s.String()
}

func (s *StartLiveMPUTaskRequestSingleSubParams) GetSourceType() *string {
	return s.SourceType
}

func (s *StartLiveMPUTaskRequestSingleSubParams) GetStreamType() *string {
	return s.StreamType
}

func (s *StartLiveMPUTaskRequestSingleSubParams) GetUserId() *string {
	return s.UserId
}

func (s *StartLiveMPUTaskRequestSingleSubParams) SetSourceType(v string) *StartLiveMPUTaskRequestSingleSubParams {
	s.SourceType = &v
	return s
}

func (s *StartLiveMPUTaskRequestSingleSubParams) SetStreamType(v string) *StartLiveMPUTaskRequestSingleSubParams {
	s.StreamType = &v
	return s
}

func (s *StartLiveMPUTaskRequestSingleSubParams) SetUserId(v string) *StartLiveMPUTaskRequestSingleSubParams {
	s.UserId = &v
	return s
}

func (s *StartLiveMPUTaskRequestSingleSubParams) Validate() error {
	return dara.Validate(s)
}

type StartLiveMPUTaskRequestTranscodeParams struct {
	// The global background image.
	Background *StartLiveMPUTaskRequestTranscodeParamsBackground `json:"Background,omitempty" xml:"Background,omitempty" type:"Struct"`
	// The encoding parameters for the output stream.
	EncodeParams *StartLiveMPUTaskRequestTranscodeParamsEncodeParams `json:"EncodeParams,omitempty" xml:"EncodeParams,omitempty" type:"Struct"`
	// The video layout information.
	//
	// >  If video transcoding is required, you must specify the video layout information, including the x-coordinate and y-coordinate, the width and height, and the layer. For audio-only transcoding, leave the video layout information empty.
	Layout *StartLiveMPUTaskRequestTranscodeParamsLayout `json:"Layout,omitempty" xml:"Layout,omitempty" type:"Struct"`
	// The information about the users whose streams are subscribed to. If you leave this parameter empty, streams from all users are mixed.
	UserInfos []*StartLiveMPUTaskRequestTranscodeParamsUserInfos `json:"UserInfos,omitempty" xml:"UserInfos,omitempty" type:"Repeated"`
}

func (s StartLiveMPUTaskRequestTranscodeParams) String() string {
	return dara.Prettify(s)
}

func (s StartLiveMPUTaskRequestTranscodeParams) GoString() string {
	return s.String()
}

func (s *StartLiveMPUTaskRequestTranscodeParams) GetBackground() *StartLiveMPUTaskRequestTranscodeParamsBackground {
	return s.Background
}

func (s *StartLiveMPUTaskRequestTranscodeParams) GetEncodeParams() *StartLiveMPUTaskRequestTranscodeParamsEncodeParams {
	return s.EncodeParams
}

func (s *StartLiveMPUTaskRequestTranscodeParams) GetLayout() *StartLiveMPUTaskRequestTranscodeParamsLayout {
	return s.Layout
}

func (s *StartLiveMPUTaskRequestTranscodeParams) GetUserInfos() []*StartLiveMPUTaskRequestTranscodeParamsUserInfos {
	return s.UserInfos
}

func (s *StartLiveMPUTaskRequestTranscodeParams) SetBackground(v *StartLiveMPUTaskRequestTranscodeParamsBackground) *StartLiveMPUTaskRequestTranscodeParams {
	s.Background = v
	return s
}

func (s *StartLiveMPUTaskRequestTranscodeParams) SetEncodeParams(v *StartLiveMPUTaskRequestTranscodeParamsEncodeParams) *StartLiveMPUTaskRequestTranscodeParams {
	s.EncodeParams = v
	return s
}

func (s *StartLiveMPUTaskRequestTranscodeParams) SetLayout(v *StartLiveMPUTaskRequestTranscodeParamsLayout) *StartLiveMPUTaskRequestTranscodeParams {
	s.Layout = v
	return s
}

func (s *StartLiveMPUTaskRequestTranscodeParams) SetUserInfos(v []*StartLiveMPUTaskRequestTranscodeParamsUserInfos) *StartLiveMPUTaskRequestTranscodeParams {
	s.UserInfos = v
	return s
}

func (s *StartLiveMPUTaskRequestTranscodeParams) Validate() error {
	return dara.Validate(s)
}

type StartLiveMPUTaskRequestTranscodeParamsBackground struct {
	// The display mode of the global background image. Valid values:
	//
	// 	- **0**: scales the background image proportionally to fit the view, with black bars displayed.
	//
	// 	- **1*	- (default): crops the background image to fit the view.
	//
	// example:
	//
	// 1
	RenderMode *string `json:"RenderMode,omitempty" xml:"RenderMode,omitempty"`
	// The URL of the global background image. The URL can be up to 2,048 characters in length.
	//
	// example:
	//
	// yourImageUrl
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s StartLiveMPUTaskRequestTranscodeParamsBackground) String() string {
	return dara.Prettify(s)
}

func (s StartLiveMPUTaskRequestTranscodeParamsBackground) GoString() string {
	return s.String()
}

func (s *StartLiveMPUTaskRequestTranscodeParamsBackground) GetRenderMode() *string {
	return s.RenderMode
}

func (s *StartLiveMPUTaskRequestTranscodeParamsBackground) GetURL() *string {
	return s.URL
}

func (s *StartLiveMPUTaskRequestTranscodeParamsBackground) SetRenderMode(v string) *StartLiveMPUTaskRequestTranscodeParamsBackground {
	s.RenderMode = &v
	return s
}

func (s *StartLiveMPUTaskRequestTranscodeParamsBackground) SetURL(v string) *StartLiveMPUTaskRequestTranscodeParamsBackground {
	s.URL = &v
	return s
}

func (s *StartLiveMPUTaskRequestTranscodeParamsBackground) Validate() error {
	return dara.Validate(s)
}

type StartLiveMPUTaskRequestTranscodeParamsEncodeParams struct {
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

func (s StartLiveMPUTaskRequestTranscodeParamsEncodeParams) String() string {
	return dara.Prettify(s)
}

func (s StartLiveMPUTaskRequestTranscodeParamsEncodeParams) GoString() string {
	return s.String()
}

func (s *StartLiveMPUTaskRequestTranscodeParamsEncodeParams) GetAudioBitrate() *string {
	return s.AudioBitrate
}

func (s *StartLiveMPUTaskRequestTranscodeParamsEncodeParams) GetAudioChannels() *string {
	return s.AudioChannels
}

func (s *StartLiveMPUTaskRequestTranscodeParamsEncodeParams) GetAudioOnly() *string {
	return s.AudioOnly
}

func (s *StartLiveMPUTaskRequestTranscodeParamsEncodeParams) GetAudioSampleRate() *string {
	return s.AudioSampleRate
}

func (s *StartLiveMPUTaskRequestTranscodeParamsEncodeParams) GetEnhancedParam() *string {
	return s.EnhancedParam
}

func (s *StartLiveMPUTaskRequestTranscodeParamsEncodeParams) GetVideoBitrate() *string {
	return s.VideoBitrate
}

func (s *StartLiveMPUTaskRequestTranscodeParamsEncodeParams) GetVideoCodec() *string {
	return s.VideoCodec
}

func (s *StartLiveMPUTaskRequestTranscodeParamsEncodeParams) GetVideoFramerate() *string {
	return s.VideoFramerate
}

func (s *StartLiveMPUTaskRequestTranscodeParamsEncodeParams) GetVideoGop() *string {
	return s.VideoGop
}

func (s *StartLiveMPUTaskRequestTranscodeParamsEncodeParams) GetVideoHeight() *string {
	return s.VideoHeight
}

func (s *StartLiveMPUTaskRequestTranscodeParamsEncodeParams) GetVideoWidth() *string {
	return s.VideoWidth
}

func (s *StartLiveMPUTaskRequestTranscodeParamsEncodeParams) SetAudioBitrate(v string) *StartLiveMPUTaskRequestTranscodeParamsEncodeParams {
	s.AudioBitrate = &v
	return s
}

func (s *StartLiveMPUTaskRequestTranscodeParamsEncodeParams) SetAudioChannels(v string) *StartLiveMPUTaskRequestTranscodeParamsEncodeParams {
	s.AudioChannels = &v
	return s
}

func (s *StartLiveMPUTaskRequestTranscodeParamsEncodeParams) SetAudioOnly(v string) *StartLiveMPUTaskRequestTranscodeParamsEncodeParams {
	s.AudioOnly = &v
	return s
}

func (s *StartLiveMPUTaskRequestTranscodeParamsEncodeParams) SetAudioSampleRate(v string) *StartLiveMPUTaskRequestTranscodeParamsEncodeParams {
	s.AudioSampleRate = &v
	return s
}

func (s *StartLiveMPUTaskRequestTranscodeParamsEncodeParams) SetEnhancedParam(v string) *StartLiveMPUTaskRequestTranscodeParamsEncodeParams {
	s.EnhancedParam = &v
	return s
}

func (s *StartLiveMPUTaskRequestTranscodeParamsEncodeParams) SetVideoBitrate(v string) *StartLiveMPUTaskRequestTranscodeParamsEncodeParams {
	s.VideoBitrate = &v
	return s
}

func (s *StartLiveMPUTaskRequestTranscodeParamsEncodeParams) SetVideoCodec(v string) *StartLiveMPUTaskRequestTranscodeParamsEncodeParams {
	s.VideoCodec = &v
	return s
}

func (s *StartLiveMPUTaskRequestTranscodeParamsEncodeParams) SetVideoFramerate(v string) *StartLiveMPUTaskRequestTranscodeParamsEncodeParams {
	s.VideoFramerate = &v
	return s
}

func (s *StartLiveMPUTaskRequestTranscodeParamsEncodeParams) SetVideoGop(v string) *StartLiveMPUTaskRequestTranscodeParamsEncodeParams {
	s.VideoGop = &v
	return s
}

func (s *StartLiveMPUTaskRequestTranscodeParamsEncodeParams) SetVideoHeight(v string) *StartLiveMPUTaskRequestTranscodeParamsEncodeParams {
	s.VideoHeight = &v
	return s
}

func (s *StartLiveMPUTaskRequestTranscodeParamsEncodeParams) SetVideoWidth(v string) *StartLiveMPUTaskRequestTranscodeParamsEncodeParams {
	s.VideoWidth = &v
	return s
}

func (s *StartLiveMPUTaskRequestTranscodeParamsEncodeParams) Validate() error {
	return dara.Validate(s)
}

type StartLiveMPUTaskRequestTranscodeParamsLayout struct {
	// The information about the panes.
	UserPanes []*StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanes `json:"UserPanes,omitempty" xml:"UserPanes,omitempty" type:"Repeated"`
}

func (s StartLiveMPUTaskRequestTranscodeParamsLayout) String() string {
	return dara.Prettify(s)
}

func (s StartLiveMPUTaskRequestTranscodeParamsLayout) GoString() string {
	return s.String()
}

func (s *StartLiveMPUTaskRequestTranscodeParamsLayout) GetUserPanes() []*StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanes {
	return s.UserPanes
}

func (s *StartLiveMPUTaskRequestTranscodeParamsLayout) SetUserPanes(v []*StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) *StartLiveMPUTaskRequestTranscodeParamsLayout {
	s.UserPanes = v
	return s
}

func (s *StartLiveMPUTaskRequestTranscodeParamsLayout) Validate() error {
	return dara.Validate(s)
}

type StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanes struct {
	// The URL of the background image of the pane. The URL can be up to 2,048 characters in length. This image is displayed if the user turns off the camera or is not present in the channel.
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
	UserInfo *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanesUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
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

func (s StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) String() string {
	return dara.Prettify(s)
}

func (s StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) GoString() string {
	return s.String()
}

func (s *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) GetBackgroundImageUrl() *string {
	return s.BackgroundImageUrl
}

func (s *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) GetHeight() *string {
	return s.Height
}

func (s *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) GetRenderMode() *string {
	return s.RenderMode
}

func (s *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) GetUserInfo() *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanesUserInfo {
	return s.UserInfo
}

func (s *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) GetWidth() *string {
	return s.Width
}

func (s *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) GetX() *string {
	return s.X
}

func (s *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) GetY() *string {
	return s.Y
}

func (s *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) GetZOrder() *string {
	return s.ZOrder
}

func (s *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) SetBackgroundImageUrl(v string) *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanes {
	s.BackgroundImageUrl = &v
	return s
}

func (s *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) SetHeight(v string) *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanes {
	s.Height = &v
	return s
}

func (s *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) SetRenderMode(v string) *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanes {
	s.RenderMode = &v
	return s
}

func (s *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) SetUserInfo(v *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanesUserInfo) *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanes {
	s.UserInfo = v
	return s
}

func (s *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) SetWidth(v string) *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanes {
	s.Width = &v
	return s
}

func (s *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) SetX(v string) *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanes {
	s.X = &v
	return s
}

func (s *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) SetY(v string) *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanes {
	s.Y = &v
	return s
}

func (s *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) SetZOrder(v string) *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanes {
	s.ZOrder = &v
	return s
}

func (s *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanes) Validate() error {
	return dara.Validate(s)
}

type StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanesUserInfo struct {
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

func (s StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanesUserInfo) String() string {
	return dara.Prettify(s)
}

func (s StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanesUserInfo) GoString() string {
	return s.String()
}

func (s *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanesUserInfo) GetChannelId() *string {
	return s.ChannelId
}

func (s *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanesUserInfo) GetSourceType() *string {
	return s.SourceType
}

func (s *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanesUserInfo) GetUserId() *string {
	return s.UserId
}

func (s *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanesUserInfo) SetChannelId(v string) *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanesUserInfo {
	s.ChannelId = &v
	return s
}

func (s *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanesUserInfo) SetSourceType(v string) *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanesUserInfo {
	s.SourceType = &v
	return s
}

func (s *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanesUserInfo) SetUserId(v string) *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanesUserInfo {
	s.UserId = &v
	return s
}

func (s *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanesUserInfo) Validate() error {
	return dara.Validate(s)
}

type StartLiveMPUTaskRequestTranscodeParamsUserInfos struct {
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

func (s StartLiveMPUTaskRequestTranscodeParamsUserInfos) String() string {
	return dara.Prettify(s)
}

func (s StartLiveMPUTaskRequestTranscodeParamsUserInfos) GoString() string {
	return s.String()
}

func (s *StartLiveMPUTaskRequestTranscodeParamsUserInfos) GetChannelId() *string {
	return s.ChannelId
}

func (s *StartLiveMPUTaskRequestTranscodeParamsUserInfos) GetSourceType() *string {
	return s.SourceType
}

func (s *StartLiveMPUTaskRequestTranscodeParamsUserInfos) GetStreamType() *string {
	return s.StreamType
}

func (s *StartLiveMPUTaskRequestTranscodeParamsUserInfos) GetUserId() *string {
	return s.UserId
}

func (s *StartLiveMPUTaskRequestTranscodeParamsUserInfos) SetChannelId(v string) *StartLiveMPUTaskRequestTranscodeParamsUserInfos {
	s.ChannelId = &v
	return s
}

func (s *StartLiveMPUTaskRequestTranscodeParamsUserInfos) SetSourceType(v string) *StartLiveMPUTaskRequestTranscodeParamsUserInfos {
	s.SourceType = &v
	return s
}

func (s *StartLiveMPUTaskRequestTranscodeParamsUserInfos) SetStreamType(v string) *StartLiveMPUTaskRequestTranscodeParamsUserInfos {
	s.StreamType = &v
	return s
}

func (s *StartLiveMPUTaskRequestTranscodeParamsUserInfos) SetUserId(v string) *StartLiveMPUTaskRequestTranscodeParamsUserInfos {
	s.UserId = &v
	return s
}

func (s *StartLiveMPUTaskRequestTranscodeParamsUserInfos) Validate() error {
	return dara.Validate(s)
}
