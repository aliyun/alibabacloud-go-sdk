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
	// The application ID. Only one ID is supported. It can contain uppercase letters, lowercase letters, digits, underscores (_), and hyphens (-). The maximum length is 64 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The channel ID. Only one ID is supported. It can contain uppercase letters, lowercase letters, digits, underscores (_), and hyphens (-). The maximum length is 64 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourChannelId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The idle timeout period. Unit: seconds. The value must be in the range of [10, 86400].
	//
	// > If you set this parameter, the task is automatically stopped when it has been idle for a period longer than MaxIdleTime. If you do not set this parameter, the task is stopped immediately after the channel is closed.
	//
	// example:
	//
	// 10
	MaxIdleTime *string `json:"MaxIdleTime,omitempty" xml:"MaxIdleTime,omitempty"`
	// The stream mixing mode. Valid values:
	//
	// - **0**: Single-stream ingest. The original single stream is ingested without stream mixing or transcoding. You do not need to configure stream mixing and transcoding parameters.
	//
	// - **1*	- (default): Stream mixing and transcoding.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	MixMode *string `json:"MixMode,omitempty" xml:"MixMode,omitempty"`
	// The parameters for ingesting to multiple URLs. You can specify multiple live ingest URLs.
	//
	// > When you set the ingest URL for a task, you must configure either the StreamURL parameter or the MultiStreamURL parameter, but not both.
	MultiStreamURL []*StartLiveMPUTaskRequestMultiStreamURL `json:"MultiStreamURL,omitempty" xml:"MultiStreamURL,omitempty" type:"Repeated"`
	// The region where the stream mixing service is located. Valid values:
	//
	// - **CN-Shanghai<props="china">(default)**: Shanghai.
	//
	// - **AP-Singapore<props="intl">(default)**: Singapore.
	//
	// - **EMAA-Saudi**: Saudi Arabia.
	//
	// example:
	//
	// CN-Shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The SEI configuration parameters.
	SeiParams *StartLiveMPUTaskRequestSeiParams `json:"SeiParams,omitempty" xml:"SeiParams,omitempty" type:"Struct"`
	// The parameters for single-stream ingest. This parameter is required when MixMode is set to 0. Do not set this parameter for stream mixing and transcoding.
	SingleSubParams *StartLiveMPUTaskRequestSingleSubParams `json:"SingleSubParams,omitempty" xml:"SingleSubParams,omitempty" type:"Struct"`
	// The live ingest URL. Only the RTMP protocol is supported. Only one URL is supported. The maximum length is 2048 characters. For information about how to generate the URL, see [Ingest URLs and playback URLs](https://help.aliyun.com/document_detail/199339.html).
	//
	// > - For domain names with hotlink protection enabled, the ingest URL must include an access token.
	//
	// - Do not use the same StreamURL in different tasks at the same time.
	//
	// - Do not use the same StreamURL within 10 seconds after a task stops.
	//
	// example:
	//
	// rtmp://example.com/live/stream
	StreamURL *string `json:"StreamURL,omitempty" xml:"StreamURL,omitempty"`
	// The task ID. Only one ID is supported. It can contain uppercase letters, lowercase letters, digits, underscores (_), and hyphens (-). The maximum length is 55 characters. This ID is the unique identifier for the bypass ingest task.
	//
	// If a task with the same ID still exists and has not been cleared when you start a new task, \\`InvalidParam\\` is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourTaskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The parameters for stream mixing and transcoding. This parameter is required when MixMode is set to 1. Do not set this parameter for single-stream ingest.
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
	if s.MultiStreamURL != nil {
		for _, item := range s.MultiStreamURL {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SeiParams != nil {
		if err := s.SeiParams.Validate(); err != nil {
			return err
		}
	}
	if s.SingleSubParams != nil {
		if err := s.SingleSubParams.Validate(); err != nil {
			return err
		}
	}
	if s.TranscodeParams != nil {
		if err := s.TranscodeParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartLiveMPUTaskRequestMultiStreamURL struct {
	// Specifies whether to ingest the stream to Alibaba Cloud CDN.
	//
	// - false: Ingest to a non-Alibaba Cloud CDN.
	//
	// - true: Ingest to Alibaba Cloud CDN.
	//
	// > The default value is false.
	//
	// example:
	//
	// false
	IsAliCdn *bool `json:"IsAliCdn,omitempty" xml:"IsAliCdn,omitempty"`
	// The live ingest URL. Only the RTMP protocol is supported. The maximum length is 2048 characters. For information about how to generate the URL, see [Ingest URLs and playback URLs](https://help.aliyun.com/document_detail/199339.html).
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
	// The layout and volume SEI. The content of this parameter can be empty, which means the default layout and volume SEI is carried.
	LayoutVolume *StartLiveMPUTaskRequestSeiParamsLayoutVolume `json:"LayoutVolume,omitempty" xml:"LayoutVolume,omitempty" type:"Struct"`
	// The pass-through SEI.
	PassThrough *StartLiveMPUTaskRequestSeiParamsPassThrough `json:"PassThrough,omitempty" xml:"PassThrough,omitempty" type:"Struct"`
	// The custom payload_type of the SEI message. The value must be in the range of 100-254. If not set, the default payload_type is 5.
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
	if s.LayoutVolume != nil {
		if err := s.LayoutVolume.Validate(); err != nil {
			return err
		}
	}
	if s.PassThrough != nil {
		if err := s.PassThrough.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartLiveMPUTaskRequestSeiParamsLayoutVolume struct {
	// Specifies whether to ensure that SEI is carried when sending an IDR keyframe. Valid values:
	//
	// - **0**: Does not ensure SEI is carried.
	//
	// - **1**: Ensures SEI is carried.
	//
	// example:
	//
	// 0
	FollowIdr *string `json:"FollowIdr,omitempty" xml:"FollowIdr,omitempty"`
	// The SEI sending interval. Unit: milliseconds. The value must be in the range of [1000, 5000].
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
	// Specifies whether to ensure that SEI is carried when sending an IDR keyframe. Valid values:
	//
	// - **0**: Does not ensure SEI is carried.
	//
	// - **1**: Ensures SEI is carried.
	//
	// example:
	//
	// 0
	FollowIdr *string `json:"FollowIdr,omitempty" xml:"FollowIdr,omitempty"`
	// The SEI sending interval. Unit: milliseconds. The value must be in the range of [1000, 5000].
	//
	// example:
	//
	// 1000
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The payload content of the pass-through SEI.
	//
	// example:
	//
	// yourPayloadContent
	PayloadContent *string `json:"PayloadContent,omitempty" xml:"PayloadContent,omitempty"`
	// The key corresponding to the payload content of the pass-through SEI. If not set, the default key is \\`udd\\`.
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
	// The type of video input stream in single-stream ingest mode. This parameter is valid only for video streams (StreamType=2). Valid values:
	//
	// - **camera*	- (default): Camera stream.
	//
	// - **shareScreen**: Screen sharing stream.
	//
	// example:
	//
	// camera
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The type of stream to ingest in single-stream ingest mode. Valid values:
	//
	// - **0*	- (default): Ingest the original stream.
	//
	// - **1**: Ingest only the audio stream.
	//
	// - **2**: Ingest only the video stream.
	//
	// example:
	//
	// 0
	StreamType *string `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	// The ID of the user whose stream is ingested. Only one stream can be ingested at a time.
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
	// The global background image for the mixed stream.
	Background *StartLiveMPUTaskRequestTranscodeParamsBackground `json:"Background,omitempty" xml:"Background,omitempty" type:"Struct"`
	// The encoding parameters for the output stream.
	EncodeParams *StartLiveMPUTaskRequestTranscodeParamsEncodeParams `json:"EncodeParams,omitempty" xml:"EncodeParams,omitempty" type:"Struct"`
	// The video layout information.
	//
	// > For video transcoding, you must specify the video layout information, including coordinates (X, Y), pane dimensions (Width, Height), and stacking order (ZOrder). For audio-only transcoding, do not specify video layout information.
	Layout *StartLiveMPUTaskRequestTranscodeParamsLayout `json:"Layout,omitempty" xml:"Layout,omitempty" type:"Struct"`
	// The information about the users to subscribe to for stream mixing. If you do not specify users, all users are included in the mixed stream.
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
	if s.Background != nil {
		if err := s.Background.Validate(); err != nil {
			return err
		}
	}
	if s.EncodeParams != nil {
		if err := s.EncodeParams.Validate(); err != nil {
			return err
		}
	}
	if s.Layout != nil {
		if err := s.Layout.Validate(); err != nil {
			return err
		}
	}
	if s.UserInfos != nil {
		for _, item := range s.UserInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type StartLiveMPUTaskRequestTranscodeParamsBackground struct {
	// The display mode of the output video. Valid values:
	//
	// - **0**: Scale and display a black background.
	//
	// - **1*	- (default): Clip.
	//
	// example:
	//
	// 1
	RenderMode *string `json:"RenderMode,omitempty" xml:"RenderMode,omitempty"`
	// The URL of the global background image. The maximum length is 2048 characters.
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
	// The audio bitrate. Unit: kbps. The value must be in the range of [8, 500].
	//
	// example:
	//
	// 128
	AudioBitrate *string `json:"AudioBitrate,omitempty" xml:"AudioBitrate,omitempty"`
	// The number of audio channels. Valid values: 1, 2.
	//
	// example:
	//
	// 2
	AudioChannels *string `json:"AudioChannels,omitempty" xml:"AudioChannels,omitempty"`
	// Specifies whether the stream is audio-only. Valid values:
	//
	// - **true**: Audio-only. You only need to set audio-related parameters.
	//
	// - **false*	- (default): Not audio-only. All parameters except VideoCodec and EnhancedParam must be specified.
	//
	// example:
	//
	// false
	AudioOnly *string `json:"AudioOnly,omitempty" xml:"AudioOnly,omitempty"`
	// The audio sampling rate. Unit: Hz. Valid values: 8000, 16000, 32000, 44100, 48000.
	//
	// example:
	//
	// 44100
	AudioSampleRate *string `json:"AudioSampleRate,omitempty" xml:"AudioSampleRate,omitempty"`
	// The enhanced encoding parameters. This is a JSON string. The supported optional configurations include \\`profile\\` and \\`preset\\`.
	//
	// - \\`profile\\`: The encoding profile. If the video encoding format is H.264, valid values for \\`profile\\` include "baseline", "main", and "high". If the video encoding format is H.265, the valid value for \\`profile\\` is "main".
	//
	// - \\`preset\\`: Balances encoding speed and quality. Valid values for \\`preset\\` include "ultrafast", "superfast", "veryfast", "faster", "fast", "medium", "slow", "slower", "veryslow", and "placebo". Each value represents a strategy for balancing encoding speed and output video quality, from "ultrafast" (fastest encoding speed) to "placebo" (highest quality, slowest encoding speed).
	//
	// > For example, "superfast" is mainly used for real-time communication. If you are not an expert in encoders, do not set this option.
	//
	// example:
	//
	// {"profile": "high", "preset": "veryfast"}
	EnhancedParam *string `json:"EnhancedParam,omitempty" xml:"EnhancedParam,omitempty"`
	// The video bitrate. Unit: kbps. The value must be in the range of [1, 10000].
	//
	// example:
	//
	// 3500
	VideoBitrate *string `json:"VideoBitrate,omitempty" xml:"VideoBitrate,omitempty"`
	// The video encoding format. Valid values:
	//
	// - H.264 (default).
	//
	// - H.265.
	//
	// example:
	//
	// H.264
	VideoCodec *string `json:"VideoCodec,omitempty" xml:"VideoCodec,omitempty"`
	// The video frame rate. Unit: fps. The value must be in the range of [1, 60].
	//
	// example:
	//
	// 25
	VideoFramerate *string `json:"VideoFramerate,omitempty" xml:"VideoFramerate,omitempty"`
	// The video GOP size. The value must be in the range of [1, 60].
	//
	// example:
	//
	// 20
	VideoGop *string `json:"VideoGop,omitempty" xml:"VideoGop,omitempty"`
	// The video height. Unit: pixels. The value must be in the range of [0, 1920].
	//
	// example:
	//
	// 1000
	VideoHeight *string `json:"VideoHeight,omitempty" xml:"VideoHeight,omitempty"`
	// The video width. Unit: pixels. The value must be in the range of [0, 1920].
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
	// The information about user panes in the mixed stream.
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
	if s.UserPanes != nil {
		for _, item := range s.UserPanes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanes struct {
	// The URL of the background image for the video pane. The maximum length is 2048 characters. When a user turns off their camera or has not joined the channel, this image is displayed in their layout position.
	//
	// example:
	//
	// yourImageUrl
	BackgroundImageUrl *string `json:"BackgroundImageUrl,omitempty" xml:"BackgroundImageUrl,omitempty"`
	// The height of the pane, as a normalized percentage.
	//
	// example:
	//
	// 0.2632
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The display mode of the output video pane. Valid values:
	//
	// - **0**: Scale and display a black background.
	//
	// - **1*	- (default): Clip.
	//
	// example:
	//
	// 1
	RenderMode *string `json:"RenderMode,omitempty" xml:"RenderMode,omitempty"`
	// The information about the user corresponding to this pane. If you do not set this parameter, the system automatically fills it based on the order in which streamers join the channel.
	//
	// > - If you specify user information, that user must already be configured in the \\`TranscodeParams.UserInfos\\` parameter.
	//
	// - This parameter is valid only for original streams and video streams.
	UserInfo *StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanesUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
	// The width of the pane, as a normalized percentage.
	//
	// example:
	//
	// 0.3564
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
	// The X-coordinate, as a normalized percentage.
	//
	// example:
	//
	// 0.2456
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	// The Y-coordinate, as a normalized percentage.
	//
	// example:
	//
	// 0.3789
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
	// The stacking order. 0 is the bottom layer. Layer 1 is on top of layer 0, and so on.
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
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartLiveMPUTaskRequestTranscodeParamsLayoutUserPanesUserInfo struct {
	// The ID of the channel where the user is located. You do not need to set this parameter for users in the same channel. For cross-channel stream mixing, set this parameter.
	//
	// example:
	//
	// yourChannelId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The type of video input stream in stream mixing and transcoding mode. This parameter is valid only for video streams (StreamType=2). Valid values:
	//
	// - **camera*	- (default): Camera stream.
	//
	// - **shareScreen**: Screen sharing stream.
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
	// The ID of the channel where the subscribed user is located. You do not need to set this parameter for users in the same channel. For cross-channel stream mixing, set this parameter.
	//
	// example:
	//
	// yourChannelId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The type of video input stream to subscribe to for stream mixing. This parameter is valid only for video streams (StreamType=2). Valid values:
	//
	// - **camera*	- (default): Camera stream.
	//
	// - **shareScreen**: Screen sharing stream.
	//
	// example:
	//
	// camera
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The type of stream to subscribe to for stream mixing. Valid values:
	//
	// - **0*	- (default): Ingest the original stream.
	//
	// - **1**: Ingest only the audio stream.
	//
	// - **2**: Ingest only the video stream.
	//
	// example:
	//
	// 0
	StreamType *string `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	// The ID of the user to subscribe to for stream mixing.
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
