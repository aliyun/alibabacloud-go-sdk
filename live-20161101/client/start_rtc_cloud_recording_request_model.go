// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRtcCloudRecordingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StartRtcCloudRecordingRequest
	GetAppId() *string
	SetChannelId(v string) *StartRtcCloudRecordingRequest
	GetChannelId() *string
	SetMaxIdleTime(v int64) *StartRtcCloudRecordingRequest
	GetMaxIdleTime() *int64
	SetMixLayoutParams(v *StartRtcCloudRecordingRequestMixLayoutParams) *StartRtcCloudRecordingRequest
	GetMixLayoutParams() *StartRtcCloudRecordingRequestMixLayoutParams
	SetMixTranscodeParams(v *StartRtcCloudRecordingRequestMixTranscodeParams) *StartRtcCloudRecordingRequest
	GetMixTranscodeParams() *StartRtcCloudRecordingRequestMixTranscodeParams
	SetNotifyAuthKey(v string) *StartRtcCloudRecordingRequest
	GetNotifyAuthKey() *string
	SetNotifyFileUploadedFormat(v []*string) *StartRtcCloudRecordingRequest
	GetNotifyFileUploadedFormat() []*string
	SetNotifyUrl(v string) *StartRtcCloudRecordingRequest
	GetNotifyUrl() *string
	SetRecordParams(v *StartRtcCloudRecordingRequestRecordParams) *StartRtcCloudRecordingRequest
	GetRecordParams() *StartRtcCloudRecordingRequestRecordParams
	SetStorageParams(v *StartRtcCloudRecordingRequestStorageParams) *StartRtcCloudRecordingRequest
	GetStorageParams() *StartRtcCloudRecordingRequestStorageParams
	SetSubscribeParams(v *StartRtcCloudRecordingRequestSubscribeParams) *StartRtcCloudRecordingRequest
	GetSubscribeParams() *StartRtcCloudRecordingRequestSubscribeParams
}

type StartRtcCloudRecordingRequest struct {
	// The ID of the app to which the channel to be recorded belongs. The app must belong to the primary account associated with the current API caller\\"s account.
	//
	// This parameter is required.
	//
	// example:
	//
	// ********-7074-****-9ef5-85c19a4*****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the channel to be recorded. Make sure that the channel has active users when you call this operation. Otherwise, the recording task fails to be created.
	//
	// This parameter is required.
	//
	// example:
	//
	// room1024
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The idle timeout period. When the task remains idle for longer than MaxIdleTime, the task is automatically stopped. Unit: seconds. The value must be within [10,14400], which is a maximum of 4 hours. Default value: 300.
	//
	// example:
	//
	// 600
	MaxIdleTime *int64 `json:"MaxIdleTime,omitempty" xml:"MaxIdleTime,omitempty"`
	// The layout parameters. This parameter is not required in single-stream recording mode and is required in stream mixing recording mode when the output is not audio-only.
	MixLayoutParams *StartRtcCloudRecordingRequestMixLayoutParams `json:"MixLayoutParams,omitempty" xml:"MixLayoutParams,omitempty" type:"Struct"`
	// The transcoding parameters. This parameter is not required in single-stream recording mode and is required in stream mixing recording mode.
	MixTranscodeParams *StartRtcCloudRecordingRequestMixTranscodeParams `json:"MixTranscodeParams,omitempty" xml:"MixTranscodeParams,omitempty" type:"Struct"`
	// The authentication key for callback messages. Leave this parameter empty to skip authentication. If specified, the key must be 16 to 64 characters in length and consist of only uppercase and lowercase letters and digits.
	//
	// example:
	//
	// mytestkeymytestkey
	NotifyAuthKey *string `json:"NotifyAuthKey,omitempty" xml:"NotifyAuthKey,omitempty"`
	// The specified formats for which a callback message is sent when the recording file upload event (RecordFileUploaded) is triggered.
	NotifyFileUploadedFormat []*string `json:"NotifyFileUploadedFormat,omitempty" xml:"NotifyFileUploadedFormat,omitempty" type:"Repeated"`
	// The URL for receiving callback messages. Task status messages are pushed to this URL in JSON format by using the POST method. The maximum length is 2048 characters.
	//
	// example:
	//
	// http://xxxx/test/mycallback
	NotifyUrl *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	// The recording parameters.
	//
	// This parameter is required.
	RecordParams *StartRtcCloudRecordingRequestRecordParams `json:"RecordParams,omitempty" xml:"RecordParams,omitempty" type:"Struct"`
	// The storage parameters.
	//
	// This parameter is required.
	StorageParams *StartRtcCloudRecordingRequestStorageParams `json:"StorageParams,omitempty" xml:"StorageParams,omitempty" type:"Struct"`
	// The subscription parameters.
	//
	// This parameter is required.
	SubscribeParams *StartRtcCloudRecordingRequestSubscribeParams `json:"SubscribeParams,omitempty" xml:"SubscribeParams,omitempty" type:"Struct"`
}

func (s StartRtcCloudRecordingRequest) String() string {
	return dara.Prettify(s)
}

func (s StartRtcCloudRecordingRequest) GoString() string {
	return s.String()
}

func (s *StartRtcCloudRecordingRequest) GetAppId() *string {
	return s.AppId
}

func (s *StartRtcCloudRecordingRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *StartRtcCloudRecordingRequest) GetMaxIdleTime() *int64 {
	return s.MaxIdleTime
}

func (s *StartRtcCloudRecordingRequest) GetMixLayoutParams() *StartRtcCloudRecordingRequestMixLayoutParams {
	return s.MixLayoutParams
}

func (s *StartRtcCloudRecordingRequest) GetMixTranscodeParams() *StartRtcCloudRecordingRequestMixTranscodeParams {
	return s.MixTranscodeParams
}

func (s *StartRtcCloudRecordingRequest) GetNotifyAuthKey() *string {
	return s.NotifyAuthKey
}

func (s *StartRtcCloudRecordingRequest) GetNotifyFileUploadedFormat() []*string {
	return s.NotifyFileUploadedFormat
}

func (s *StartRtcCloudRecordingRequest) GetNotifyUrl() *string {
	return s.NotifyUrl
}

func (s *StartRtcCloudRecordingRequest) GetRecordParams() *StartRtcCloudRecordingRequestRecordParams {
	return s.RecordParams
}

func (s *StartRtcCloudRecordingRequest) GetStorageParams() *StartRtcCloudRecordingRequestStorageParams {
	return s.StorageParams
}

func (s *StartRtcCloudRecordingRequest) GetSubscribeParams() *StartRtcCloudRecordingRequestSubscribeParams {
	return s.SubscribeParams
}

func (s *StartRtcCloudRecordingRequest) SetAppId(v string) *StartRtcCloudRecordingRequest {
	s.AppId = &v
	return s
}

func (s *StartRtcCloudRecordingRequest) SetChannelId(v string) *StartRtcCloudRecordingRequest {
	s.ChannelId = &v
	return s
}

func (s *StartRtcCloudRecordingRequest) SetMaxIdleTime(v int64) *StartRtcCloudRecordingRequest {
	s.MaxIdleTime = &v
	return s
}

func (s *StartRtcCloudRecordingRequest) SetMixLayoutParams(v *StartRtcCloudRecordingRequestMixLayoutParams) *StartRtcCloudRecordingRequest {
	s.MixLayoutParams = v
	return s
}

func (s *StartRtcCloudRecordingRequest) SetMixTranscodeParams(v *StartRtcCloudRecordingRequestMixTranscodeParams) *StartRtcCloudRecordingRequest {
	s.MixTranscodeParams = v
	return s
}

func (s *StartRtcCloudRecordingRequest) SetNotifyAuthKey(v string) *StartRtcCloudRecordingRequest {
	s.NotifyAuthKey = &v
	return s
}

func (s *StartRtcCloudRecordingRequest) SetNotifyFileUploadedFormat(v []*string) *StartRtcCloudRecordingRequest {
	s.NotifyFileUploadedFormat = v
	return s
}

func (s *StartRtcCloudRecordingRequest) SetNotifyUrl(v string) *StartRtcCloudRecordingRequest {
	s.NotifyUrl = &v
	return s
}

func (s *StartRtcCloudRecordingRequest) SetRecordParams(v *StartRtcCloudRecordingRequestRecordParams) *StartRtcCloudRecordingRequest {
	s.RecordParams = v
	return s
}

func (s *StartRtcCloudRecordingRequest) SetStorageParams(v *StartRtcCloudRecordingRequestStorageParams) *StartRtcCloudRecordingRequest {
	s.StorageParams = v
	return s
}

func (s *StartRtcCloudRecordingRequest) SetSubscribeParams(v *StartRtcCloudRecordingRequestSubscribeParams) *StartRtcCloudRecordingRequest {
	s.SubscribeParams = v
	return s
}

func (s *StartRtcCloudRecordingRequest) Validate() error {
	if s.MixLayoutParams != nil {
		if err := s.MixLayoutParams.Validate(); err != nil {
			return err
		}
	}
	if s.MixTranscodeParams != nil {
		if err := s.MixTranscodeParams.Validate(); err != nil {
			return err
		}
	}
	if s.RecordParams != nil {
		if err := s.RecordParams.Validate(); err != nil {
			return err
		}
	}
	if s.StorageParams != nil {
		if err := s.StorageParams.Validate(); err != nil {
			return err
		}
	}
	if s.SubscribeParams != nil {
		if err := s.SubscribeParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartRtcCloudRecordingRequestMixLayoutParams struct {
	// The global background image for stream mixing.
	MixBackground *StartRtcCloudRecordingRequestMixLayoutParamsMixBackground `json:"MixBackground,omitempty" xml:"MixBackground,omitempty" type:"Struct"`
	// Specifies the window layout information for subscribed users. Only users whose UserId has layout information configured are included in the video. This parameter is required in stream mixing mode when recording non-audio-only files.
	UserPanes []*StartRtcCloudRecordingRequestMixLayoutParamsUserPanes `json:"UserPanes,omitempty" xml:"UserPanes,omitempty" type:"Repeated"`
}

func (s StartRtcCloudRecordingRequestMixLayoutParams) String() string {
	return dara.Prettify(s)
}

func (s StartRtcCloudRecordingRequestMixLayoutParams) GoString() string {
	return s.String()
}

func (s *StartRtcCloudRecordingRequestMixLayoutParams) GetMixBackground() *StartRtcCloudRecordingRequestMixLayoutParamsMixBackground {
	return s.MixBackground
}

func (s *StartRtcCloudRecordingRequestMixLayoutParams) GetUserPanes() []*StartRtcCloudRecordingRequestMixLayoutParamsUserPanes {
	return s.UserPanes
}

func (s *StartRtcCloudRecordingRequestMixLayoutParams) SetMixBackground(v *StartRtcCloudRecordingRequestMixLayoutParamsMixBackground) *StartRtcCloudRecordingRequestMixLayoutParams {
	s.MixBackground = v
	return s
}

func (s *StartRtcCloudRecordingRequestMixLayoutParams) SetUserPanes(v []*StartRtcCloudRecordingRequestMixLayoutParamsUserPanes) *StartRtcCloudRecordingRequestMixLayoutParams {
	s.UserPanes = v
	return s
}

func (s *StartRtcCloudRecordingRequestMixLayoutParams) Validate() error {
	if s.MixBackground != nil {
		if err := s.MixBackground.Validate(); err != nil {
			return err
		}
	}
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

type StartRtcCloudRecordingRequestMixLayoutParamsMixBackground struct {
	// The display mode for the output. Valid values:
	//
	// example:
	//
	// 0
	RenderMode *int32 `json:"RenderMode,omitempty" xml:"RenderMode,omitempty"`
	// The URL of the background image. The maximum length is 2048 characters.
	//
	// example:
	//
	// https://xxxx.com/photos/my-test-picture.png
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s StartRtcCloudRecordingRequestMixLayoutParamsMixBackground) String() string {
	return dara.Prettify(s)
}

func (s StartRtcCloudRecordingRequestMixLayoutParamsMixBackground) GoString() string {
	return s.String()
}

func (s *StartRtcCloudRecordingRequestMixLayoutParamsMixBackground) GetRenderMode() *int32 {
	return s.RenderMode
}

func (s *StartRtcCloudRecordingRequestMixLayoutParamsMixBackground) GetUrl() *string {
	return s.Url
}

func (s *StartRtcCloudRecordingRequestMixLayoutParamsMixBackground) SetRenderMode(v int32) *StartRtcCloudRecordingRequestMixLayoutParamsMixBackground {
	s.RenderMode = &v
	return s
}

func (s *StartRtcCloudRecordingRequestMixLayoutParamsMixBackground) SetUrl(v string) *StartRtcCloudRecordingRequestMixLayoutParamsMixBackground {
	s.Url = &v
	return s
}

func (s *StartRtcCloudRecordingRequestMixLayoutParamsMixBackground) Validate() error {
	return dara.Validate(s)
}

type StartRtcCloudRecordingRequestMixLayoutParamsUserPanes struct {
	// The pane height as a normalized percentage. The value must be within [0,1]. Default value: 0.
	//
	// example:
	//
	// 0.5
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// The video input stream type for this UserId. If UserId is not specified, this SourceType setting has no effect. Valid values:
	//
	// example:
	//
	// 0
	SourceType *int32 `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The background image for the sub-pane. When a user turns off the camera, has not published a stream after joining, or leaves the channel midway, the corresponding image fills the layout position.
	SubBackground *StartRtcCloudRecordingRequestMixLayoutParamsUserPanesSubBackground `json:"SubBackground,omitempty" xml:"SubBackground,omitempty" type:"Struct"`
	// The UserId corresponding to this window.
	//
	// example:
	//
	// userA
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The pane width as a normalized percentage. The value must be within [0,1]. Default value: 0.
	//
	// example:
	//
	// 0.5
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
	// The X coordinate as a normalized percentage. The value must be within [0,1]. Default value: 0.
	//
	// example:
	//
	// 0
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	// The Y coordinate as a normalized percentage. The value must be within [0,1]. Default value: 0.
	//
	// example:
	//
	// 0
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
	// The stacking order. 0 is the bottom layer, layer 1 is above layer 0, and so on. Default value: 0.
	//
	// example:
	//
	// 0
	ZOrder *int32 `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s StartRtcCloudRecordingRequestMixLayoutParamsUserPanes) String() string {
	return dara.Prettify(s)
}

func (s StartRtcCloudRecordingRequestMixLayoutParamsUserPanes) GoString() string {
	return s.String()
}

func (s *StartRtcCloudRecordingRequestMixLayoutParamsUserPanes) GetHeight() *string {
	return s.Height
}

func (s *StartRtcCloudRecordingRequestMixLayoutParamsUserPanes) GetSourceType() *int32 {
	return s.SourceType
}

func (s *StartRtcCloudRecordingRequestMixLayoutParamsUserPanes) GetSubBackground() *StartRtcCloudRecordingRequestMixLayoutParamsUserPanesSubBackground {
	return s.SubBackground
}

func (s *StartRtcCloudRecordingRequestMixLayoutParamsUserPanes) GetUserId() *string {
	return s.UserId
}

func (s *StartRtcCloudRecordingRequestMixLayoutParamsUserPanes) GetWidth() *string {
	return s.Width
}

func (s *StartRtcCloudRecordingRequestMixLayoutParamsUserPanes) GetX() *string {
	return s.X
}

func (s *StartRtcCloudRecordingRequestMixLayoutParamsUserPanes) GetY() *string {
	return s.Y
}

func (s *StartRtcCloudRecordingRequestMixLayoutParamsUserPanes) GetZOrder() *int32 {
	return s.ZOrder
}

func (s *StartRtcCloudRecordingRequestMixLayoutParamsUserPanes) SetHeight(v string) *StartRtcCloudRecordingRequestMixLayoutParamsUserPanes {
	s.Height = &v
	return s
}

func (s *StartRtcCloudRecordingRequestMixLayoutParamsUserPanes) SetSourceType(v int32) *StartRtcCloudRecordingRequestMixLayoutParamsUserPanes {
	s.SourceType = &v
	return s
}

func (s *StartRtcCloudRecordingRequestMixLayoutParamsUserPanes) SetSubBackground(v *StartRtcCloudRecordingRequestMixLayoutParamsUserPanesSubBackground) *StartRtcCloudRecordingRequestMixLayoutParamsUserPanes {
	s.SubBackground = v
	return s
}

func (s *StartRtcCloudRecordingRequestMixLayoutParamsUserPanes) SetUserId(v string) *StartRtcCloudRecordingRequestMixLayoutParamsUserPanes {
	s.UserId = &v
	return s
}

func (s *StartRtcCloudRecordingRequestMixLayoutParamsUserPanes) SetWidth(v string) *StartRtcCloudRecordingRequestMixLayoutParamsUserPanes {
	s.Width = &v
	return s
}

func (s *StartRtcCloudRecordingRequestMixLayoutParamsUserPanes) SetX(v string) *StartRtcCloudRecordingRequestMixLayoutParamsUserPanes {
	s.X = &v
	return s
}

func (s *StartRtcCloudRecordingRequestMixLayoutParamsUserPanes) SetY(v string) *StartRtcCloudRecordingRequestMixLayoutParamsUserPanes {
	s.Y = &v
	return s
}

func (s *StartRtcCloudRecordingRequestMixLayoutParamsUserPanes) SetZOrder(v int32) *StartRtcCloudRecordingRequestMixLayoutParamsUserPanes {
	s.ZOrder = &v
	return s
}

func (s *StartRtcCloudRecordingRequestMixLayoutParamsUserPanes) Validate() error {
	if s.SubBackground != nil {
		if err := s.SubBackground.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartRtcCloudRecordingRequestMixLayoutParamsUserPanesSubBackground struct {
	// The display mode for the sub-pane output. Valid values:
	//
	// example:
	//
	// 0
	RenderMode *int32 `json:"RenderMode,omitempty" xml:"RenderMode,omitempty"`
	// The URL of the background image. The maximum length is 2048 characters.
	//
	// example:
	//
	// https://xxxx.com/photos/my-test-pane-picture.png
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s StartRtcCloudRecordingRequestMixLayoutParamsUserPanesSubBackground) String() string {
	return dara.Prettify(s)
}

func (s StartRtcCloudRecordingRequestMixLayoutParamsUserPanesSubBackground) GoString() string {
	return s.String()
}

func (s *StartRtcCloudRecordingRequestMixLayoutParamsUserPanesSubBackground) GetRenderMode() *int32 {
	return s.RenderMode
}

func (s *StartRtcCloudRecordingRequestMixLayoutParamsUserPanesSubBackground) GetUrl() *string {
	return s.Url
}

func (s *StartRtcCloudRecordingRequestMixLayoutParamsUserPanesSubBackground) SetRenderMode(v int32) *StartRtcCloudRecordingRequestMixLayoutParamsUserPanesSubBackground {
	s.RenderMode = &v
	return s
}

func (s *StartRtcCloudRecordingRequestMixLayoutParamsUserPanesSubBackground) SetUrl(v string) *StartRtcCloudRecordingRequestMixLayoutParamsUserPanesSubBackground {
	s.Url = &v
	return s
}

func (s *StartRtcCloudRecordingRequestMixLayoutParamsUserPanesSubBackground) Validate() error {
	return dara.Validate(s)
}

type StartRtcCloudRecordingRequestMixTranscodeParams struct {
	// The audio bitrate in kbps. The value must be in the range of [8, 500]. This parameter is required in stream mixing mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// 300
	AudioBitrate *int64 `json:"AudioBitrate,omitempty" xml:"AudioBitrate,omitempty"`
	// The number of audio channels. Valid values:
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	AudioChannels *int32 `json:"AudioChannels,omitempty" xml:"AudioChannels,omitempty"`
	// The audio sample rate in Hz. Valid values:
	//
	// This parameter is required.
	//
	// example:
	//
	// 32000
	AudioSampleRate *int64 `json:"AudioSampleRate,omitempty" xml:"AudioSampleRate,omitempty"`
	// The frame fill type when a stream is interrupted. Valid values:
	//
	// example:
	//
	// 0
	FrameFillType *int32 `json:"FrameFillType,omitempty" xml:"FrameFillType,omitempty"`
	// The video bitrate in kbps. The value must be in the range of [1, 10000].
	//
	// example:
	//
	// 5000
	VideoBitrate *int32 `json:"VideoBitrate,omitempty" xml:"VideoBitrate,omitempty"`
	// The video encoding format. Valid values:
	//
	// example:
	//
	// H.264
	VideoCodec *string `json:"VideoCodec,omitempty" xml:"VideoCodec,omitempty"`
	// The video frame rate in fps. The value must be in the range of [1, 60].
	//
	// example:
	//
	// 30
	VideoFramerate *int32 `json:"VideoFramerate,omitempty" xml:"VideoFramerate,omitempty"`
	// The video GOP. An I-frame is inserted every VideoGop frames. The value must be in the range of [1, 60].
	//
	// example:
	//
	// 30
	VideoGop *int32 `json:"VideoGop,omitempty" xml:"VideoGop,omitempty"`
	// The video height in pixels. The value must be in the range of [0, 1920]. Default value: 0.
	//
	// example:
	//
	// 480
	VideoHeight *int32 `json:"VideoHeight,omitempty" xml:"VideoHeight,omitempty"`
	// The video width in pixels. The value must be in the range of [0, 1920]. Default value: 0.
	//
	// example:
	//
	// 640
	VideoWidth *int32 `json:"VideoWidth,omitempty" xml:"VideoWidth,omitempty"`
}

func (s StartRtcCloudRecordingRequestMixTranscodeParams) String() string {
	return dara.Prettify(s)
}

func (s StartRtcCloudRecordingRequestMixTranscodeParams) GoString() string {
	return s.String()
}

func (s *StartRtcCloudRecordingRequestMixTranscodeParams) GetAudioBitrate() *int64 {
	return s.AudioBitrate
}

func (s *StartRtcCloudRecordingRequestMixTranscodeParams) GetAudioChannels() *int32 {
	return s.AudioChannels
}

func (s *StartRtcCloudRecordingRequestMixTranscodeParams) GetAudioSampleRate() *int64 {
	return s.AudioSampleRate
}

func (s *StartRtcCloudRecordingRequestMixTranscodeParams) GetFrameFillType() *int32 {
	return s.FrameFillType
}

func (s *StartRtcCloudRecordingRequestMixTranscodeParams) GetVideoBitrate() *int32 {
	return s.VideoBitrate
}

func (s *StartRtcCloudRecordingRequestMixTranscodeParams) GetVideoCodec() *string {
	return s.VideoCodec
}

func (s *StartRtcCloudRecordingRequestMixTranscodeParams) GetVideoFramerate() *int32 {
	return s.VideoFramerate
}

func (s *StartRtcCloudRecordingRequestMixTranscodeParams) GetVideoGop() *int32 {
	return s.VideoGop
}

func (s *StartRtcCloudRecordingRequestMixTranscodeParams) GetVideoHeight() *int32 {
	return s.VideoHeight
}

func (s *StartRtcCloudRecordingRequestMixTranscodeParams) GetVideoWidth() *int32 {
	return s.VideoWidth
}

func (s *StartRtcCloudRecordingRequestMixTranscodeParams) SetAudioBitrate(v int64) *StartRtcCloudRecordingRequestMixTranscodeParams {
	s.AudioBitrate = &v
	return s
}

func (s *StartRtcCloudRecordingRequestMixTranscodeParams) SetAudioChannels(v int32) *StartRtcCloudRecordingRequestMixTranscodeParams {
	s.AudioChannels = &v
	return s
}

func (s *StartRtcCloudRecordingRequestMixTranscodeParams) SetAudioSampleRate(v int64) *StartRtcCloudRecordingRequestMixTranscodeParams {
	s.AudioSampleRate = &v
	return s
}

func (s *StartRtcCloudRecordingRequestMixTranscodeParams) SetFrameFillType(v int32) *StartRtcCloudRecordingRequestMixTranscodeParams {
	s.FrameFillType = &v
	return s
}

func (s *StartRtcCloudRecordingRequestMixTranscodeParams) SetVideoBitrate(v int32) *StartRtcCloudRecordingRequestMixTranscodeParams {
	s.VideoBitrate = &v
	return s
}

func (s *StartRtcCloudRecordingRequestMixTranscodeParams) SetVideoCodec(v string) *StartRtcCloudRecordingRequestMixTranscodeParams {
	s.VideoCodec = &v
	return s
}

func (s *StartRtcCloudRecordingRequestMixTranscodeParams) SetVideoFramerate(v int32) *StartRtcCloudRecordingRequestMixTranscodeParams {
	s.VideoFramerate = &v
	return s
}

func (s *StartRtcCloudRecordingRequestMixTranscodeParams) SetVideoGop(v int32) *StartRtcCloudRecordingRequestMixTranscodeParams {
	s.VideoGop = &v
	return s
}

func (s *StartRtcCloudRecordingRequestMixTranscodeParams) SetVideoHeight(v int32) *StartRtcCloudRecordingRequestMixTranscodeParams {
	s.VideoHeight = &v
	return s
}

func (s *StartRtcCloudRecordingRequestMixTranscodeParams) SetVideoWidth(v int32) *StartRtcCloudRecordingRequestMixTranscodeParams {
	s.VideoWidth = &v
	return s
}

func (s *StartRtcCloudRecordingRequestMixTranscodeParams) Validate() error {
	return dara.Validate(s)
}

type StartRtcCloudRecordingRequestRecordParams struct {
	// The maximum duration of a recording file, in seconds. A recording file that exceeds this duration is split. The value must be in the range of [180, 7200], which means a maximum of 2 hours. If this parameter is not specified, the default value is 7200 (2 hours).
	//
	// example:
	//
	// 7200
	MaxFileDuration *int64 `json:"MaxFileDuration,omitempty" xml:"MaxFileDuration,omitempty"`
	// The recording mode. Valid values:
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	RecordMode *int32 `json:"RecordMode,omitempty" xml:"RecordMode,omitempty"`
	// The media type of the output recording stream. Valid values:
	//
	// example:
	//
	// 0
	StreamType *int32 `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s StartRtcCloudRecordingRequestRecordParams) String() string {
	return dara.Prettify(s)
}

func (s StartRtcCloudRecordingRequestRecordParams) GoString() string {
	return s.String()
}

func (s *StartRtcCloudRecordingRequestRecordParams) GetMaxFileDuration() *int64 {
	return s.MaxFileDuration
}

func (s *StartRtcCloudRecordingRequestRecordParams) GetRecordMode() *int32 {
	return s.RecordMode
}

func (s *StartRtcCloudRecordingRequestRecordParams) GetStreamType() *int32 {
	return s.StreamType
}

func (s *StartRtcCloudRecordingRequestRecordParams) SetMaxFileDuration(v int64) *StartRtcCloudRecordingRequestRecordParams {
	s.MaxFileDuration = &v
	return s
}

func (s *StartRtcCloudRecordingRequestRecordParams) SetRecordMode(v int32) *StartRtcCloudRecordingRequestRecordParams {
	s.RecordMode = &v
	return s
}

func (s *StartRtcCloudRecordingRequestRecordParams) SetStreamType(v int32) *StartRtcCloudRecordingRequestRecordParams {
	s.StreamType = &v
	return s
}

func (s *StartRtcCloudRecordingRequestRecordParams) Validate() error {
	return dara.Validate(s)
}

type StartRtcCloudRecordingRequestStorageParams struct {
	// The file storage information, which specifies the format, storage location, and naming of recording files. This parameter takes effect only when StorageType is set to OSS.
	FileInfo []*StartRtcCloudRecordingRequestStorageParamsFileInfo `json:"FileInfo,omitempty" xml:"FileInfo,omitempty" type:"Repeated"`
	// The OSS storage configuration. This parameter is required when the storage method is OSS and is invalid when the storage method is VOD.
	OSSParams *StartRtcCloudRecordingRequestStorageParamsOSSParams `json:"OSSParams,omitempty" xml:"OSSParams,omitempty" type:"Struct"`
	// The storage method. Valid values:
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	StorageType *int32 `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The VOD storage configuration. This parameter is required when the storage method is VOD and is invalid when the storage method is OSS.
	VodParams *StartRtcCloudRecordingRequestStorageParamsVodParams `json:"VodParams,omitempty" xml:"VodParams,omitempty" type:"Struct"`
}

func (s StartRtcCloudRecordingRequestStorageParams) String() string {
	return dara.Prettify(s)
}

func (s StartRtcCloudRecordingRequestStorageParams) GoString() string {
	return s.String()
}

func (s *StartRtcCloudRecordingRequestStorageParams) GetFileInfo() []*StartRtcCloudRecordingRequestStorageParamsFileInfo {
	return s.FileInfo
}

func (s *StartRtcCloudRecordingRequestStorageParams) GetOSSParams() *StartRtcCloudRecordingRequestStorageParamsOSSParams {
	return s.OSSParams
}

func (s *StartRtcCloudRecordingRequestStorageParams) GetStorageType() *int32 {
	return s.StorageType
}

func (s *StartRtcCloudRecordingRequestStorageParams) GetVodParams() *StartRtcCloudRecordingRequestStorageParamsVodParams {
	return s.VodParams
}

func (s *StartRtcCloudRecordingRequestStorageParams) SetFileInfo(v []*StartRtcCloudRecordingRequestStorageParamsFileInfo) *StartRtcCloudRecordingRequestStorageParams {
	s.FileInfo = v
	return s
}

func (s *StartRtcCloudRecordingRequestStorageParams) SetOSSParams(v *StartRtcCloudRecordingRequestStorageParamsOSSParams) *StartRtcCloudRecordingRequestStorageParams {
	s.OSSParams = v
	return s
}

func (s *StartRtcCloudRecordingRequestStorageParams) SetStorageType(v int32) *StartRtcCloudRecordingRequestStorageParams {
	s.StorageType = &v
	return s
}

func (s *StartRtcCloudRecordingRequestStorageParams) SetVodParams(v *StartRtcCloudRecordingRequestStorageParamsVodParams) *StartRtcCloudRecordingRequestStorageParams {
	s.VodParams = v
	return s
}

func (s *StartRtcCloudRecordingRequestStorageParams) Validate() error {
	if s.FileInfo != nil {
		for _, item := range s.FileInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OSSParams != nil {
		if err := s.OSSParams.Validate(); err != nil {
			return err
		}
	}
	if s.VodParams != nil {
		if err := s.VodParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartRtcCloudRecordingRequestStorageParamsFileInfo struct {
	// The file naming format. You can select and combine the following variables in any order:
	//
	// example:
	//
	// {AppId}_{ChannelId}_{StartTime}_{UserId}
	FileNamePattern *string `json:"FileNamePattern,omitempty" xml:"FileNamePattern,omitempty"`
	// The file storage path. Each element in the array corresponds to a directory level. For example, if the value is ["dir1","dir2"], the xxx.m3u8 file is saved as dir1/dir2/TaskId/xxx.m3u8. If this parameter is empty, the file is saved as TaskId/xxx.m3u8.
	FilePathPrefix []*string `json:"FilePathPrefix,omitempty" xml:"FilePathPrefix,omitempty" type:"Repeated"`
	// The file storage format. Valid values:
	//
	// This parameter is required.
	//
	// example:
	//
	// HLS
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The segment length in seconds. This parameter takes effect only in HLS format. The value must be in the range of [10, 30]. Default value: 30.
	//
	// example:
	//
	// 30
	SliceDuration *int64 `json:"SliceDuration,omitempty" xml:"SliceDuration,omitempty"`
	// The segment naming format. This parameter takes effect only in HLS format. Similar to FileNamePattern, but with an additional variable Sequence:
	//
	// example:
	//
	// {AppId}_{ChannelId}_{StartTime}_{Sequence}
	SliceNamePattern *string `json:"SliceNamePattern,omitempty" xml:"SliceNamePattern,omitempty"`
}

func (s StartRtcCloudRecordingRequestStorageParamsFileInfo) String() string {
	return dara.Prettify(s)
}

func (s StartRtcCloudRecordingRequestStorageParamsFileInfo) GoString() string {
	return s.String()
}

func (s *StartRtcCloudRecordingRequestStorageParamsFileInfo) GetFileNamePattern() *string {
	return s.FileNamePattern
}

func (s *StartRtcCloudRecordingRequestStorageParamsFileInfo) GetFilePathPrefix() []*string {
	return s.FilePathPrefix
}

func (s *StartRtcCloudRecordingRequestStorageParamsFileInfo) GetFormat() *string {
	return s.Format
}

func (s *StartRtcCloudRecordingRequestStorageParamsFileInfo) GetSliceDuration() *int64 {
	return s.SliceDuration
}

func (s *StartRtcCloudRecordingRequestStorageParamsFileInfo) GetSliceNamePattern() *string {
	return s.SliceNamePattern
}

func (s *StartRtcCloudRecordingRequestStorageParamsFileInfo) SetFileNamePattern(v string) *StartRtcCloudRecordingRequestStorageParamsFileInfo {
	s.FileNamePattern = &v
	return s
}

func (s *StartRtcCloudRecordingRequestStorageParamsFileInfo) SetFilePathPrefix(v []*string) *StartRtcCloudRecordingRequestStorageParamsFileInfo {
	s.FilePathPrefix = v
	return s
}

func (s *StartRtcCloudRecordingRequestStorageParamsFileInfo) SetFormat(v string) *StartRtcCloudRecordingRequestStorageParamsFileInfo {
	s.Format = &v
	return s
}

func (s *StartRtcCloudRecordingRequestStorageParamsFileInfo) SetSliceDuration(v int64) *StartRtcCloudRecordingRequestStorageParamsFileInfo {
	s.SliceDuration = &v
	return s
}

func (s *StartRtcCloudRecordingRequestStorageParamsFileInfo) SetSliceNamePattern(v string) *StartRtcCloudRecordingRequestStorageParamsFileInfo {
	s.SliceNamePattern = &v
	return s
}

func (s *StartRtcCloudRecordingRequestStorageParamsFileInfo) Validate() error {
	return dara.Validate(s)
}

type StartRtcCloudRecordingRequestStorageParamsOSSParams struct {
	// The name of the OSS bucket. The bucket must belong to the primary account associated with the current API caller\\"s account.
	//
	// This parameter is required.
	//
	// example:
	//
	// mytest-bucket
	OSSBucket *string `json:"OSSBucket,omitempty" xml:"OSSBucket,omitempty"`
	// The endpoint of the OSS storage. The corresponding region ID must be consistent with the selected service registration endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss-cn-shanghai.aliyuncs.com
	OSSEndpoint *string `json:"OSSEndpoint,omitempty" xml:"OSSEndpoint,omitempty"`
}

func (s StartRtcCloudRecordingRequestStorageParamsOSSParams) String() string {
	return dara.Prettify(s)
}

func (s StartRtcCloudRecordingRequestStorageParamsOSSParams) GoString() string {
	return s.String()
}

func (s *StartRtcCloudRecordingRequestStorageParamsOSSParams) GetOSSBucket() *string {
	return s.OSSBucket
}

func (s *StartRtcCloudRecordingRequestStorageParamsOSSParams) GetOSSEndpoint() *string {
	return s.OSSEndpoint
}

func (s *StartRtcCloudRecordingRequestStorageParamsOSSParams) SetOSSBucket(v string) *StartRtcCloudRecordingRequestStorageParamsOSSParams {
	s.OSSBucket = &v
	return s
}

func (s *StartRtcCloudRecordingRequestStorageParamsOSSParams) SetOSSEndpoint(v string) *StartRtcCloudRecordingRequestStorageParamsOSSParams {
	s.OSSEndpoint = &v
	return s
}

func (s *StartRtcCloudRecordingRequestStorageParamsOSSParams) Validate() error {
	return dara.Validate(s)
}

type StartRtcCloudRecordingRequestStorageParamsVodParams struct {
	// Specifies whether to enable automatic composition. Valid values:
	//
	// example:
	//
	// 0
	AutoCompose *int32 `json:"AutoCompose,omitempty" xml:"AutoCompose,omitempty"`
	// The ID of the VOD transcoding template group used to transcode the automatically composed video in the VOD service.
	//
	// example:
	//
	// ****4c34112cfe68248f2f77759c****
	ComposeVodTranscodeGroupId *string `json:"ComposeVodTranscodeGroupId,omitempty" xml:"ComposeVodTranscodeGroupId,omitempty"`
	// The storage address configured in the ApsaraVideo VOD console under Media Asset Management > Storage Management. Recording files are first saved to this location and then uploaded to VOD.
	//
	// example:
	//
	// mytest.oss-cn-shenzhen.aliyuncs.com
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	// The ID of the VOD transcoding template group.
	//
	// example:
	//
	// ****8a914d3989e9825eb90530b2****
	VodTranscodeGroupId *string `json:"VodTranscodeGroupId,omitempty" xml:"VodTranscodeGroupId,omitempty"`
}

func (s StartRtcCloudRecordingRequestStorageParamsVodParams) String() string {
	return dara.Prettify(s)
}

func (s StartRtcCloudRecordingRequestStorageParamsVodParams) GoString() string {
	return s.String()
}

func (s *StartRtcCloudRecordingRequestStorageParamsVodParams) GetAutoCompose() *int32 {
	return s.AutoCompose
}

func (s *StartRtcCloudRecordingRequestStorageParamsVodParams) GetComposeVodTranscodeGroupId() *string {
	return s.ComposeVodTranscodeGroupId
}

func (s *StartRtcCloudRecordingRequestStorageParamsVodParams) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *StartRtcCloudRecordingRequestStorageParamsVodParams) GetVodTranscodeGroupId() *string {
	return s.VodTranscodeGroupId
}

func (s *StartRtcCloudRecordingRequestStorageParamsVodParams) SetAutoCompose(v int32) *StartRtcCloudRecordingRequestStorageParamsVodParams {
	s.AutoCompose = &v
	return s
}

func (s *StartRtcCloudRecordingRequestStorageParamsVodParams) SetComposeVodTranscodeGroupId(v string) *StartRtcCloudRecordingRequestStorageParamsVodParams {
	s.ComposeVodTranscodeGroupId = &v
	return s
}

func (s *StartRtcCloudRecordingRequestStorageParamsVodParams) SetStorageLocation(v string) *StartRtcCloudRecordingRequestStorageParamsVodParams {
	s.StorageLocation = &v
	return s
}

func (s *StartRtcCloudRecordingRequestStorageParamsVodParams) SetVodTranscodeGroupId(v string) *StartRtcCloudRecordingRequestStorageParamsVodParams {
	s.VodTranscodeGroupId = &v
	return s
}

func (s *StartRtcCloudRecordingRequestStorageParamsVodParams) Validate() error {
	return dara.Validate(s)
}

type StartRtcCloudRecordingRequestSubscribeParams struct {
	// The list of subscribed UserId entries. In single-stream recording mode, each UserId is recorded separately. In stream mixing recording mode, the audio and video of all UserIds are mixed into a single set of audio and video.
	//
	// This parameter is required.
	SubscribeUserIdList []*StartRtcCloudRecordingRequestSubscribeParamsSubscribeUserIdList `json:"SubscribeUserIdList,omitempty" xml:"SubscribeUserIdList,omitempty" type:"Repeated"`
}

func (s StartRtcCloudRecordingRequestSubscribeParams) String() string {
	return dara.Prettify(s)
}

func (s StartRtcCloudRecordingRequestSubscribeParams) GoString() string {
	return s.String()
}

func (s *StartRtcCloudRecordingRequestSubscribeParams) GetSubscribeUserIdList() []*StartRtcCloudRecordingRequestSubscribeParamsSubscribeUserIdList {
	return s.SubscribeUserIdList
}

func (s *StartRtcCloudRecordingRequestSubscribeParams) SetSubscribeUserIdList(v []*StartRtcCloudRecordingRequestSubscribeParamsSubscribeUserIdList) *StartRtcCloudRecordingRequestSubscribeParams {
	s.SubscribeUserIdList = v
	return s
}

func (s *StartRtcCloudRecordingRequestSubscribeParams) Validate() error {
	if s.SubscribeUserIdList != nil {
		for _, item := range s.SubscribeUserIdList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type StartRtcCloudRecordingRequestSubscribeParamsSubscribeUserIdList struct {
	// The video input stream type of the UserId. This parameter takes effect only when the subscription is not audio-only (StreamType != 1). Valid values:
	//
	// example:
	//
	// 0
	SourceType *int32 `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The media type of the subscribed UserId. Valid values:
	//
	// example:
	//
	// 0
	StreamType *int32 `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	// The subscribed UserId.
	//
	// This parameter is required.
	//
	// example:
	//
	// userA
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s StartRtcCloudRecordingRequestSubscribeParamsSubscribeUserIdList) String() string {
	return dara.Prettify(s)
}

func (s StartRtcCloudRecordingRequestSubscribeParamsSubscribeUserIdList) GoString() string {
	return s.String()
}

func (s *StartRtcCloudRecordingRequestSubscribeParamsSubscribeUserIdList) GetSourceType() *int32 {
	return s.SourceType
}

func (s *StartRtcCloudRecordingRequestSubscribeParamsSubscribeUserIdList) GetStreamType() *int32 {
	return s.StreamType
}

func (s *StartRtcCloudRecordingRequestSubscribeParamsSubscribeUserIdList) GetUserId() *string {
	return s.UserId
}

func (s *StartRtcCloudRecordingRequestSubscribeParamsSubscribeUserIdList) SetSourceType(v int32) *StartRtcCloudRecordingRequestSubscribeParamsSubscribeUserIdList {
	s.SourceType = &v
	return s
}

func (s *StartRtcCloudRecordingRequestSubscribeParamsSubscribeUserIdList) SetStreamType(v int32) *StartRtcCloudRecordingRequestSubscribeParamsSubscribeUserIdList {
	s.StreamType = &v
	return s
}

func (s *StartRtcCloudRecordingRequestSubscribeParamsSubscribeUserIdList) SetUserId(v string) *StartRtcCloudRecordingRequestSubscribeParamsSubscribeUserIdList {
	s.UserId = &v
	return s
}

func (s *StartRtcCloudRecordingRequestSubscribeParamsSubscribeUserIdList) Validate() error {
	return dara.Validate(s)
}
