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
	// This parameter is required.
	//
	// example:
	//
	// ********-7074-****-9ef5-85c19a4*****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// room1024
	ChannelId          *string                                          `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	MaxIdleTime        *int64                                           `json:"MaxIdleTime,omitempty" xml:"MaxIdleTime,omitempty"`
	MixLayoutParams    *StartRtcCloudRecordingRequestMixLayoutParams    `json:"MixLayoutParams,omitempty" xml:"MixLayoutParams,omitempty" type:"Struct"`
	MixTranscodeParams *StartRtcCloudRecordingRequestMixTranscodeParams `json:"MixTranscodeParams,omitempty" xml:"MixTranscodeParams,omitempty" type:"Struct"`
	NotifyAuthKey      *string                                          `json:"NotifyAuthKey,omitempty" xml:"NotifyAuthKey,omitempty"`
	// example:
	//
	// http://xxxx/test/mycallback
	NotifyUrl *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	// This parameter is required.
	RecordParams *StartRtcCloudRecordingRequestRecordParams `json:"RecordParams,omitempty" xml:"RecordParams,omitempty" type:"Struct"`
	// This parameter is required.
	StorageParams *StartRtcCloudRecordingRequestStorageParams `json:"StorageParams,omitempty" xml:"StorageParams,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type StartRtcCloudRecordingRequestMixLayoutParams struct {
	MixBackground *StartRtcCloudRecordingRequestMixLayoutParamsMixBackground `json:"MixBackground,omitempty" xml:"MixBackground,omitempty" type:"Struct"`
	UserPanes     []*StartRtcCloudRecordingRequestMixLayoutParamsUserPanes   `json:"UserPanes,omitempty" xml:"UserPanes,omitempty" type:"Repeated"`
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
	return dara.Validate(s)
}

type StartRtcCloudRecordingRequestMixLayoutParamsMixBackground struct {
	// example:
	//
	// 0
	RenderMode *int32 `json:"RenderMode,omitempty" xml:"RenderMode,omitempty"`
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
	// example:
	//
	// 0.5
	Height *string `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 0
	SourceType    *int32                                                              `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	SubBackground *StartRtcCloudRecordingRequestMixLayoutParamsUserPanesSubBackground `json:"SubBackground,omitempty" xml:"SubBackground,omitempty" type:"Struct"`
	// example:
	//
	// userA
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 0.5
	Width *string `json:"Width,omitempty" xml:"Width,omitempty"`
	// example:
	//
	// 0
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 0
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
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
	return dara.Validate(s)
}

type StartRtcCloudRecordingRequestMixLayoutParamsUserPanesSubBackground struct {
	// example:
	//
	// 0
	RenderMode *int32 `json:"RenderMode,omitempty" xml:"RenderMode,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 300
	AudioBitrate *int64 `json:"AudioBitrate,omitempty" xml:"AudioBitrate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	AudioChannels *int32 `json:"AudioChannels,omitempty" xml:"AudioChannels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 32000
	AudioSampleRate *int64 `json:"AudioSampleRate,omitempty" xml:"AudioSampleRate,omitempty"`
	// example:
	//
	// 0
	FrameFillType *int32 `json:"FrameFillType,omitempty" xml:"FrameFillType,omitempty"`
	// example:
	//
	// 5000
	VideoBitrate *int32 `json:"VideoBitrate,omitempty" xml:"VideoBitrate,omitempty"`
	// example:
	//
	// H.264
	VideoCodec *string `json:"VideoCodec,omitempty" xml:"VideoCodec,omitempty"`
	// example:
	//
	// 30
	VideoFramerate *int32 `json:"VideoFramerate,omitempty" xml:"VideoFramerate,omitempty"`
	// example:
	//
	// 30
	VideoGop *int32 `json:"VideoGop,omitempty" xml:"VideoGop,omitempty"`
	// example:
	//
	// 480
	VideoHeight *int32 `json:"VideoHeight,omitempty" xml:"VideoHeight,omitempty"`
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
	// example:
	//
	// 7200
	MaxFileDuration *int64 `json:"MaxFileDuration,omitempty" xml:"MaxFileDuration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	RecordMode *int32 `json:"RecordMode,omitempty" xml:"RecordMode,omitempty"`
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
	FileInfo  []*StartRtcCloudRecordingRequestStorageParamsFileInfo `json:"FileInfo,omitempty" xml:"FileInfo,omitempty" type:"Repeated"`
	OSSParams *StartRtcCloudRecordingRequestStorageParamsOSSParams  `json:"OSSParams,omitempty" xml:"OSSParams,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	StorageType *int32                                               `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	VodParams   *StartRtcCloudRecordingRequestStorageParamsVodParams `json:"VodParams,omitempty" xml:"VodParams,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type StartRtcCloudRecordingRequestStorageParamsFileInfo struct {
	// example:
	//
	// {AppId}_{ChannelId}_{StartTime}_{UserId}
	FileNamePattern *string   `json:"FileNamePattern,omitempty" xml:"FileNamePattern,omitempty"`
	FilePathPrefix  []*string `json:"FilePathPrefix,omitempty" xml:"FilePathPrefix,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// HLS
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
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

func (s *StartRtcCloudRecordingRequestStorageParamsFileInfo) SetSliceNamePattern(v string) *StartRtcCloudRecordingRequestStorageParamsFileInfo {
	s.SliceNamePattern = &v
	return s
}

func (s *StartRtcCloudRecordingRequestStorageParamsFileInfo) Validate() error {
	return dara.Validate(s)
}

type StartRtcCloudRecordingRequestStorageParamsOSSParams struct {
	// This parameter is required.
	//
	// example:
	//
	// mytest-bucket
	OSSBucket *string `json:"OSSBucket,omitempty" xml:"OSSBucket,omitempty"`
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
	AutoCompose                *int32  `json:"AutoCompose,omitempty" xml:"AutoCompose,omitempty"`
	ComposeVodTranscodeGroupId *string `json:"ComposeVodTranscodeGroupId,omitempty" xml:"ComposeVodTranscodeGroupId,omitempty"`
	StorageLocation            *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	VodTranscodeGroupId        *string `json:"VodTranscodeGroupId,omitempty" xml:"VodTranscodeGroupId,omitempty"`
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
	return dara.Validate(s)
}

type StartRtcCloudRecordingRequestSubscribeParamsSubscribeUserIdList struct {
	// example:
	//
	// 0
	SourceType *int32 `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// 0
	StreamType *int32 `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
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
