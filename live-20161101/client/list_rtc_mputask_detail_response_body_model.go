// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRtcMPUTaskDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMPUTasks(v []*ListRtcMPUTaskDetailResponseBodyMPUTasks) *ListRtcMPUTaskDetailResponseBody
	GetMPUTasks() []*ListRtcMPUTaskDetailResponseBodyMPUTasks
	SetRequestId(v string) *ListRtcMPUTaskDetailResponseBody
	GetRequestId() *string
}

type ListRtcMPUTaskDetailResponseBody struct {
	// The parameters that you configured when you called the StartLiveMPUTask operation to create the tasks.
	MPUTasks []*ListRtcMPUTaskDetailResponseBodyMPUTasks `json:"MPUTasks,omitempty" xml:"MPUTasks,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRtcMPUTaskDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRtcMPUTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s *ListRtcMPUTaskDetailResponseBody) GetMPUTasks() []*ListRtcMPUTaskDetailResponseBodyMPUTasks {
	return s.MPUTasks
}

func (s *ListRtcMPUTaskDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRtcMPUTaskDetailResponseBody) SetMPUTasks(v []*ListRtcMPUTaskDetailResponseBodyMPUTasks) *ListRtcMPUTaskDetailResponseBody {
	s.MPUTasks = v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBody) SetRequestId(v string) *ListRtcMPUTaskDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRtcMPUTaskDetailResponseBodyMPUTasks struct {
	// The ID of the application.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the channel.
	//
	// example:
	//
	// yourChannelId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The timeout period of an idle connection. Unit: seconds.
	//
	// >  If the task is idle for a period of time longer than the duration specified by the MaxIdleTime parameter, the task is automatically stopped. If the parameter is not specified, the task is stopped after the channel is closed.
	//
	// example:
	//
	// 10
	MaxIdleTime *string `json:"MaxIdleTime,omitempty" xml:"MaxIdleTime,omitempty"`
	// The stream mixing mode. Valid values:
	//
	// 	- 0: relays the original single stream without mixing streams. If the value of this parameter is 0, the TranscodeParams parameter is empty.
	//
	// 	- 1 (default): mixes multiple streams into a single stream and relays the mixed stream.
	//
	// example:
	//
	// 0
	MixMode *string `json:"MixMode,omitempty" xml:"MixMode,omitempty"`
	// The multiple ingest URLs relayed.
	MultiStreamURL []*ListRtcMPUTaskDetailResponseBodyMPUTasksMultiStreamURL `json:"MultiStreamURL,omitempty" xml:"MultiStreamURL,omitempty" type:"Repeated"`
	// The region in which the streams are mixed. Valid values:
	//
	// 	- **CN-shanghai**
	//
	// 	- **AP-Singapore (default)**
	//
	// 	- **EMAA-Saudi**
	//
	// example:
	//
	// CN-Shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The supplemental enhancement information (SEI) parameters.
	SeiParams *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParams `json:"SeiParams,omitempty" xml:"SeiParams,omitempty" type:"Struct"`
	// The parameters of the single-stream relay task.
	SingleSubParams *ListRtcMPUTaskDetailResponseBodyMPUTasksSingleSubParams `json:"SingleSubParams,omitempty" xml:"SingleSubParams,omitempty" type:"Struct"`
	// The ingest URL.
	//
	// example:
	//
	// rtmp://example.com/live/stream****
	StreamURL *string `json:"StreamURL,omitempty" xml:"StreamURL,omitempty"`
	// The ID of the stream relay task.
	//
	// example:
	//
	// yourTaskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The mixed-stream relay parameters.
	TranscodeParams *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParams `json:"TranscodeParams,omitempty" xml:"TranscodeParams,omitempty" type:"Struct"`
}

func (s ListRtcMPUTaskDetailResponseBodyMPUTasks) String() string {
	return dara.Prettify(s)
}

func (s ListRtcMPUTaskDetailResponseBodyMPUTasks) GoString() string {
	return s.String()
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasks) GetAppId() *string {
	return s.AppId
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasks) GetChannelId() *string {
	return s.ChannelId
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasks) GetMaxIdleTime() *string {
	return s.MaxIdleTime
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasks) GetMixMode() *string {
	return s.MixMode
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasks) GetMultiStreamURL() []*ListRtcMPUTaskDetailResponseBodyMPUTasksMultiStreamURL {
	return s.MultiStreamURL
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasks) GetRegion() *string {
	return s.Region
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasks) GetSeiParams() *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParams {
	return s.SeiParams
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasks) GetSingleSubParams() *ListRtcMPUTaskDetailResponseBodyMPUTasksSingleSubParams {
	return s.SingleSubParams
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasks) GetStreamURL() *string {
	return s.StreamURL
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasks) GetTranscodeParams() *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParams {
	return s.TranscodeParams
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasks) SetAppId(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasks {
	s.AppId = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasks) SetChannelId(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasks {
	s.ChannelId = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasks) SetMaxIdleTime(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasks {
	s.MaxIdleTime = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasks) SetMixMode(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasks {
	s.MixMode = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasks) SetMultiStreamURL(v []*ListRtcMPUTaskDetailResponseBodyMPUTasksMultiStreamURL) *ListRtcMPUTaskDetailResponseBodyMPUTasks {
	s.MultiStreamURL = v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasks) SetRegion(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasks {
	s.Region = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasks) SetSeiParams(v *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParams) *ListRtcMPUTaskDetailResponseBodyMPUTasks {
	s.SeiParams = v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasks) SetSingleSubParams(v *ListRtcMPUTaskDetailResponseBodyMPUTasksSingleSubParams) *ListRtcMPUTaskDetailResponseBodyMPUTasks {
	s.SingleSubParams = v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasks) SetStreamURL(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasks {
	s.StreamURL = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasks) SetTaskId(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasks {
	s.TaskId = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasks) SetTranscodeParams(v *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParams) *ListRtcMPUTaskDetailResponseBodyMPUTasks {
	s.TranscodeParams = v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasks) Validate() error {
	return dara.Validate(s)
}

type ListRtcMPUTaskDetailResponseBodyMPUTasksMultiStreamURL struct {
	// Indicates whether stream relay is performed by using Alibaba Cloud CDN. Valid values:
	//
	// 	- false: Stream relay is performed by using a CDN service that is not Alibaba Cloud CDN.
	//
	// 	- true: Stream relay is performed by using Alibaba Cloud CDN.
	//
	// example:
	//
	// false
	IsAliCdn *bool `json:"IsAliCdn,omitempty" xml:"IsAliCdn,omitempty"`
	// The ingest URL.
	//
	// example:
	//
	// rtmp://example.com/live/stream****
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s ListRtcMPUTaskDetailResponseBodyMPUTasksMultiStreamURL) String() string {
	return dara.Prettify(s)
}

func (s ListRtcMPUTaskDetailResponseBodyMPUTasksMultiStreamURL) GoString() string {
	return s.String()
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksMultiStreamURL) GetIsAliCdn() *bool {
	return s.IsAliCdn
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksMultiStreamURL) GetURL() *string {
	return s.URL
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksMultiStreamURL) SetIsAliCdn(v bool) *ListRtcMPUTaskDetailResponseBodyMPUTasksMultiStreamURL {
	s.IsAliCdn = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksMultiStreamURL) SetURL(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksMultiStreamURL {
	s.URL = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksMultiStreamURL) Validate() error {
	return dara.Validate(s)
}

type ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParams struct {
	// The layout and volume SEI. If the return value is an empty string, the default layout and volume SEI is used.
	LayoutVolume *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParamsLayoutVolume `json:"LayoutVolume,omitempty" xml:"LayoutVolume,omitempty" type:"Struct"`
	// The custom SEI.
	PassThrough *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParamsPassThrough `json:"PassThrough,omitempty" xml:"PassThrough,omitempty" type:"Struct"`
	// The custom payload type. Valid values: 100 to 254. Default value: 5.
	//
	// example:
	//
	// 100
	PayloadType *string `json:"PayloadType,omitempty" xml:"PayloadType,omitempty"`
}

func (s ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParams) String() string {
	return dara.Prettify(s)
}

func (s ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParams) GoString() string {
	return s.String()
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParams) GetLayoutVolume() *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParamsLayoutVolume {
	return s.LayoutVolume
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParams) GetPassThrough() *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParamsPassThrough {
	return s.PassThrough
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParams) GetPayloadType() *string {
	return s.PayloadType
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParams) SetLayoutVolume(v *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParamsLayoutVolume) *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParams {
	s.LayoutVolume = v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParams) SetPassThrough(v *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParamsPassThrough) *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParams {
	s.PassThrough = v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParams) SetPayloadType(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParams {
	s.PayloadType = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParams) Validate() error {
	return dara.Validate(s)
}

type ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParamsLayoutVolume struct {
	// Indicates whether to add SEI messages to Instantaneous Decoder Refresh (IDR) frames. Valid values:
	//
	// 	- 0: does not add SEI messages.
	//
	// 	- 1: adds SEI messages.
	//
	// example:
	//
	// 0
	FollowIdr *string `json:"FollowIdr,omitempty" xml:"FollowIdr,omitempty"`
	// The interval at which the SEI messages are added. Unit: milliseconds.
	//
	// example:
	//
	// 1000
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParamsLayoutVolume) String() string {
	return dara.Prettify(s)
}

func (s ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParamsLayoutVolume) GoString() string {
	return s.String()
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParamsLayoutVolume) GetFollowIdr() *string {
	return s.FollowIdr
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParamsLayoutVolume) GetInterval() *string {
	return s.Interval
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParamsLayoutVolume) SetFollowIdr(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParamsLayoutVolume {
	s.FollowIdr = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParamsLayoutVolume) SetInterval(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParamsLayoutVolume {
	s.Interval = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParamsLayoutVolume) Validate() error {
	return dara.Validate(s)
}

type ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParamsPassThrough struct {
	// Indicates whether to add SEI messages to Instantaneous Decoder Refresh (IDR) frames. Valid values:
	//
	// 	- 0: does not add SEI messages.
	//
	// 	- 1: adds SEI messages.
	//
	// example:
	//
	// 0
	FollowIdr *string `json:"FollowIdr,omitempty" xml:"FollowIdr,omitempty"`
	// The interval at which the SEI messages are added. Unit: milliseconds.
	//
	// example:
	//
	// 1000
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The payload content of the custom SEI.
	//
	// example:
	//
	// yourPayloadContent
	PayloadContent *string `json:"PayloadContent,omitempty" xml:"PayloadContent,omitempty"`
	// The key of the payload content. Default value: udd.
	//
	// example:
	//
	// yourPayloadContentKey
	PayloadContentKey *string `json:"PayloadContentKey,omitempty" xml:"PayloadContentKey,omitempty"`
}

func (s ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParamsPassThrough) String() string {
	return dara.Prettify(s)
}

func (s ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParamsPassThrough) GoString() string {
	return s.String()
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParamsPassThrough) GetFollowIdr() *string {
	return s.FollowIdr
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParamsPassThrough) GetInterval() *string {
	return s.Interval
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParamsPassThrough) GetPayloadContent() *string {
	return s.PayloadContent
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParamsPassThrough) GetPayloadContentKey() *string {
	return s.PayloadContentKey
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParamsPassThrough) SetFollowIdr(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParamsPassThrough {
	s.FollowIdr = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParamsPassThrough) SetInterval(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParamsPassThrough {
	s.Interval = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParamsPassThrough) SetPayloadContent(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParamsPassThrough {
	s.PayloadContent = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParamsPassThrough) SetPayloadContentKey(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParamsPassThrough {
	s.PayloadContentKey = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksSeiParamsPassThrough) Validate() error {
	return dara.Validate(s)
}

type ListRtcMPUTaskDetailResponseBodyMPUTasksSingleSubParams struct {
	// The source of the video. This parameter is valid only if you set StreamType to 2. Valid values:
	//
	// 	- camera (default): captures the video by using a camera.
	//
	// 	- shareScreen: captures the content displayed on a screen.
	//
	// example:
	//
	// camera
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The type of the stream that is relayed. Valid values:
	//
	// 	- 0 (default): the original stream.
	//
	// 	- 1: the audio-only stream.
	//
	// 	- 2: the video-only stream.
	//
	// example:
	//
	// 0
	StreamType *string `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	// The ID of the user whose stream is relayed. In single-stream relay mode, you can relay only one stream in a request.
	//
	// example:
	//
	// yourSubUserId
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListRtcMPUTaskDetailResponseBodyMPUTasksSingleSubParams) String() string {
	return dara.Prettify(s)
}

func (s ListRtcMPUTaskDetailResponseBodyMPUTasksSingleSubParams) GoString() string {
	return s.String()
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksSingleSubParams) GetSourceType() *string {
	return s.SourceType
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksSingleSubParams) GetStreamType() *string {
	return s.StreamType
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksSingleSubParams) GetUserId() *string {
	return s.UserId
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksSingleSubParams) SetSourceType(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksSingleSubParams {
	s.SourceType = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksSingleSubParams) SetStreamType(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksSingleSubParams {
	s.StreamType = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksSingleSubParams) SetUserId(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksSingleSubParams {
	s.UserId = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksSingleSubParams) Validate() error {
	return dara.Validate(s)
}

type ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParams struct {
	// The global background image.
	Background *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsBackground `json:"Background,omitempty" xml:"Background,omitempty" type:"Struct"`
	// The encoding parameters of the output stream.
	EncodeParams *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams `json:"EncodeParams,omitempty" xml:"EncodeParams,omitempty" type:"Struct"`
	// The video layout information.
	//
	// >  The video layout information includes the x-coordinate, y-coordinate, width, height, and layer of the pane. For audio-only transcoding, no video layout information is returned.
	Layout *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayout `json:"Layout,omitempty" xml:"Layout,omitempty" type:"Struct"`
	// The information about the user whose stream is mixed. If an empty value is returned, streams from all users are mixed.
	UserInfos []*ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsUserInfos `json:"UserInfos,omitempty" xml:"UserInfos,omitempty" type:"Repeated"`
}

func (s ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParams) String() string {
	return dara.Prettify(s)
}

func (s ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParams) GoString() string {
	return s.String()
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParams) GetBackground() *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsBackground {
	return s.Background
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParams) GetEncodeParams() *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams {
	return s.EncodeParams
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParams) GetLayout() *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayout {
	return s.Layout
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParams) GetUserInfos() []*ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsUserInfos {
	return s.UserInfos
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParams) SetBackground(v *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsBackground) *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParams {
	s.Background = v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParams) SetEncodeParams(v *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams) *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParams {
	s.EncodeParams = v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParams) SetLayout(v *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayout) *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParams {
	s.Layout = v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParams) SetUserInfos(v []*ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsUserInfos) *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParams {
	s.UserInfos = v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParams) Validate() error {
	return dara.Validate(s)
}

type ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsBackground struct {
	// The display mode. Valid values:
	//
	// 	- 0: proportionally scales the video or background image to fit the pane. Black bars are added to fill the extra space.
	//
	// 	- 1 (default): crops the video or background image to fit the pane.
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

func (s ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsBackground) String() string {
	return dara.Prettify(s)
}

func (s ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsBackground) GoString() string {
	return s.String()
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsBackground) GetRenderMode() *string {
	return s.RenderMode
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsBackground) GetURL() *string {
	return s.URL
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsBackground) SetRenderMode(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsBackground {
	s.RenderMode = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsBackground) SetURL(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsBackground {
	s.URL = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsBackground) Validate() error {
	return dara.Validate(s)
}

type ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams struct {
	// The bitrate of the audio. Unit: Kbit/s.
	//
	// example:
	//
	// 128
	AudioBitrate *string `json:"AudioBitrate,omitempty" xml:"AudioBitrate,omitempty"`
	// The number of audio channels. Valid values: 1 and 2.
	//
	// example:
	//
	// 2
	AudioChannels *string `json:"AudioChannels,omitempty" xml:"AudioChannels,omitempty"`
	// Indicates whether the output stream is an audio-only stream. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	AudioOnly *string `json:"AudioOnly,omitempty" xml:"AudioOnly,omitempty"`
	// The audio sampling rate. Unit: Hz.
	//
	// example:
	//
	// 44100
	AudioSampleRate *string `json:"AudioSampleRate,omitempty" xml:"AudioSampleRate,omitempty"`
	// The parameter for advanced video encoding. The value is a JSON string. Optional fields:
	//
	// 	- profile: the encoding level. If the video encoding format is set to H.264, the valid values of this field are baseline, main, and high.
	//
	// 	- preset: adjusts the trade-off between encoding speed and video quality. Valid values: ultrafast, superfast, veryfast, faster, fast, medium, slow, slower, veryslow, and placebo. Each value specifies a level of trade-off between encoding speed and video quality. For example, the ultrafast preset has the fastest encoding speed but the lowest video quality, while the placebo preset sacrifices the encoding speed for the best video quality.
	//
	// example:
	//
	// {"profile": "high", "preset": "veryfast"}
	EnhancedParam *string `json:"EnhancedParam,omitempty" xml:"EnhancedParam,omitempty"`
	// The bitrate of the video. Unit: Kbit/s.
	//
	// example:
	//
	// 3500
	VideoBitrate *string `json:"VideoBitrate,omitempty" xml:"VideoBitrate,omitempty"`
	// The video encoding format. Default value: H.264.
	//
	// example:
	//
	// H.264
	VideoCodec *string `json:"VideoCodec,omitempty" xml:"VideoCodec,omitempty"`
	// The frame rate of the video. Unit: frames per second (FPS).
	//
	// example:
	//
	// 25
	VideoFramerate *string `json:"VideoFramerate,omitempty" xml:"VideoFramerate,omitempty"`
	// The group of pictures (GOP) size of the video.
	//
	// example:
	//
	// 20
	VideoGop *string `json:"VideoGop,omitempty" xml:"VideoGop,omitempty"`
	// The height of the video. Unit: pixels.
	//
	// example:
	//
	// 1000
	VideoHeight *string `json:"VideoHeight,omitempty" xml:"VideoHeight,omitempty"`
	// The width of the video. Unit: pixels.
	//
	// example:
	//
	// 1920
	VideoWidth *string `json:"VideoWidth,omitempty" xml:"VideoWidth,omitempty"`
}

func (s ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams) String() string {
	return dara.Prettify(s)
}

func (s ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams) GoString() string {
	return s.String()
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams) GetAudioBitrate() *string {
	return s.AudioBitrate
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams) GetAudioChannels() *string {
	return s.AudioChannels
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams) GetAudioOnly() *string {
	return s.AudioOnly
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams) GetAudioSampleRate() *string {
	return s.AudioSampleRate
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams) GetEnhancedParam() *string {
	return s.EnhancedParam
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams) GetVideoBitrate() *string {
	return s.VideoBitrate
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams) GetVideoCodec() *string {
	return s.VideoCodec
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams) GetVideoFramerate() *string {
	return s.VideoFramerate
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams) GetVideoGop() *string {
	return s.VideoGop
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams) GetVideoHeight() *string {
	return s.VideoHeight
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams) GetVideoWidth() *string {
	return s.VideoWidth
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams) SetAudioBitrate(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams {
	s.AudioBitrate = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams) SetAudioChannels(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams {
	s.AudioChannels = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams) SetAudioOnly(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams {
	s.AudioOnly = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams) SetAudioSampleRate(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams {
	s.AudioSampleRate = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams) SetEnhancedParam(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams {
	s.EnhancedParam = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams) SetVideoBitrate(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams {
	s.VideoBitrate = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams) SetVideoCodec(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams {
	s.VideoCodec = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams) SetVideoFramerate(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams {
	s.VideoFramerate = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams) SetVideoGop(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams {
	s.VideoGop = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams) SetVideoHeight(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams {
	s.VideoHeight = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams) SetVideoWidth(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams {
	s.VideoWidth = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsEncodeParams) Validate() error {
	return dara.Validate(s)
}

type ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayout struct {
	// The information about the panes.
	UserPanes []*ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanes `json:"UserPanes,omitempty" xml:"UserPanes,omitempty" type:"Repeated"`
}

func (s ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayout) String() string {
	return dara.Prettify(s)
}

func (s ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayout) GoString() string {
	return s.String()
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayout) GetUserPanes() []*ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanes {
	return s.UserPanes
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayout) SetUserPanes(v []*ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanes) *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayout {
	s.UserPanes = v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayout) Validate() error {
	return dara.Validate(s)
}

type ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanes struct {
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
	// The display mode. Valid values:
	//
	// 	- 0: proportionally scales the video or background image to fit the pane. Black bars are added to fill the extra space.
	//
	// 	- 1 (default): crops the video or background image to fit the pane.
	//
	// example:
	//
	// 1
	RenderMode *string `json:"RenderMode,omitempty" xml:"RenderMode,omitempty"`
	// The information about the user whose stream is played in the pane.
	UserInfo *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanesUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
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
	// The layer of the pane. A value of 0 indicates that the pane is placed at the bottom layer. A larger value indicates a higher layer.
	//
	// example:
	//
	// 0
	ZOrder *string `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanes) String() string {
	return dara.Prettify(s)
}

func (s ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanes) GoString() string {
	return s.String()
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanes) GetBackgroundImageUrl() *string {
	return s.BackgroundImageUrl
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanes) GetHeight() *string {
	return s.Height
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanes) GetRenderMode() *string {
	return s.RenderMode
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanes) GetUserInfo() *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanesUserInfo {
	return s.UserInfo
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanes) GetWidth() *string {
	return s.Width
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanes) GetX() *string {
	return s.X
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanes) GetY() *string {
	return s.Y
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanes) GetZOrder() *string {
	return s.ZOrder
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanes) SetBackgroundImageUrl(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanes {
	s.BackgroundImageUrl = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanes) SetHeight(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanes {
	s.Height = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanes) SetRenderMode(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanes {
	s.RenderMode = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanes) SetUserInfo(v *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanesUserInfo) *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanes {
	s.UserInfo = v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanes) SetWidth(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanes {
	s.Width = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanes) SetX(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanes {
	s.X = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanes) SetY(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanes {
	s.Y = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanes) SetZOrder(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanes {
	s.ZOrder = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanes) Validate() error {
	return dara.Validate(s)
}

type ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanesUserInfo struct {
	// The ID of the channel where the user is.
	//
	// example:
	//
	// yourChannelId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The source of the video. This parameter is valid only if you set StreamType to 2. Valid values:
	//
	// 	- camera (default): captures the video by using a camera.
	//
	// 	- shareScreen: captures the content displayed on a screen.
	//
	// example:
	//
	// camera
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// yourSubUserId
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanesUserInfo) String() string {
	return dara.Prettify(s)
}

func (s ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanesUserInfo) GoString() string {
	return s.String()
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanesUserInfo) GetChannelId() *string {
	return s.ChannelId
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanesUserInfo) GetSourceType() *string {
	return s.SourceType
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanesUserInfo) GetUserId() *string {
	return s.UserId
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanesUserInfo) SetChannelId(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanesUserInfo {
	s.ChannelId = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanesUserInfo) SetSourceType(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanesUserInfo {
	s.SourceType = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanesUserInfo) SetUserId(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanesUserInfo {
	s.UserId = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsLayoutUserPanesUserInfo) Validate() error {
	return dara.Validate(s)
}

type ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsUserInfos struct {
	// The ID of the channel where the user is.
	//
	// example:
	//
	// yourChannelId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The source of the video. This parameter is valid only if you set StreamType to 2. Valid values:
	//
	// 	- camera (default): captures the video by using a camera.
	//
	// 	- shareScreen: captures the content displayed on a screen.
	//
	// example:
	//
	// camera
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The type of the stream that is relayed. Valid values:
	//
	// 	- 0 (default): the original stream.
	//
	// 	- 1: the audio-only stream.
	//
	// 	- 2: the video-only stream.
	//
	// example:
	//
	// 0
	StreamType *string `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// yourSubUserId
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsUserInfos) String() string {
	return dara.Prettify(s)
}

func (s ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsUserInfos) GoString() string {
	return s.String()
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsUserInfos) GetChannelId() *string {
	return s.ChannelId
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsUserInfos) GetSourceType() *string {
	return s.SourceType
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsUserInfos) GetStreamType() *string {
	return s.StreamType
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsUserInfos) GetUserId() *string {
	return s.UserId
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsUserInfos) SetChannelId(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsUserInfos {
	s.ChannelId = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsUserInfos) SetSourceType(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsUserInfos {
	s.SourceType = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsUserInfos) SetStreamType(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsUserInfos {
	s.StreamType = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsUserInfos) SetUserId(v string) *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsUserInfos {
	s.UserId = &v
	return s
}

func (s *ListRtcMPUTaskDetailResponseBodyMPUTasksTranscodeParamsUserInfos) Validate() error {
	return dara.Validate(s)
}
