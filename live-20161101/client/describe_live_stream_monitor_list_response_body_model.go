// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamMonitorListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLiveStreamMonitorList(v []*DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList) *DescribeLiveStreamMonitorListResponseBody
	GetLiveStreamMonitorList() []*DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList
	SetRequestId(v string) *DescribeLiveStreamMonitorListResponseBody
	GetRequestId() *string
	SetTotal(v int32) *DescribeLiveStreamMonitorListResponseBody
	GetTotal() *int32
}

type DescribeLiveStreamMonitorListResponseBody struct {
	// The list of monitoring sessions.
	LiveStreamMonitorList []*DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList `json:"LiveStreamMonitorList,omitempty" xml:"LiveStreamMonitorList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 2234baba-a586-46ea-8bd4-c8f7891abcdef
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of monitoring sessions.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeLiveStreamMonitorListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamMonitorListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamMonitorListResponseBody) GetLiveStreamMonitorList() []*DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList {
	return s.LiveStreamMonitorList
}

func (s *DescribeLiveStreamMonitorListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveStreamMonitorListResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeLiveStreamMonitorListResponseBody) SetLiveStreamMonitorList(v []*DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList) *DescribeLiveStreamMonitorListResponseBody {
	s.LiveStreamMonitorList = v
	return s
}

func (s *DescribeLiveStreamMonitorListResponseBody) SetRequestId(v string) *DescribeLiveStreamMonitorListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamMonitorListResponseBody) SetTotal(v int32) *DescribeLiveStreamMonitorListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeLiveStreamMonitorListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList struct {
	// The audio source in the layout.
	//
	// example:
	//
	// 1
	AudioFrom *int32 `json:"AudioFrom,omitempty" xml:"AudioFrom,omitempty"`
	// The callback URL that sends monitoring alerts.
	//
	// example:
	//
	// http://guide.aliyundoc.com/notify
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// The URL of the DingTalk chatbot.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=7a7d404056eee1f2fd944ace9bcfc361dc6448583e1d3d3baa****
	DingTalkWebHookUrl *string `json:"DingTalkWebHookUrl,omitempty" xml:"DingTalkWebHookUrl,omitempty"`
	// The domain name.
	//
	// example:
	//
	// demo.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The list of monitored input streams.
	InputList []*DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputList `json:"InputList,omitempty" xml:"InputList,omitempty" type:"Repeated"`
	// The monitoring alert thresholds. The following fields are included:
	//
	// 	- fpsLowThres: the video frame rate alert threshold. The value is a floating-point number.
	//
	// 	- brHighThres: the audio/video bitrate alert threshold. The value is a floating-point number.
	//
	// 	- eofDurationThresSec: the interruption duration alert threshold. The value is a floating-point number.
	//
	// example:
	//
	// "{\\"fpsLowThres\\": 0.6,\\"brLowThres\\": 1.1,\\"eofDurationThresSec\\": 10}"
	MonitorConfig *string `json:"MonitorConfig,omitempty" xml:"MonitorConfig,omitempty"`
	// The ID of the monitoring session.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	MonitorId *string `json:"MonitorId,omitempty" xml:"MonitorId,omitempty"`
	// The name of the monitoring session.
	//
	// example:
	//
	// liveMonito****
	MonitorName *string `json:"MonitorName,omitempty" xml:"MonitorName,omitempty"`
	// The output resolution template. Valid values:
	//
	// 	- **lp_ld**: low definition
	//
	// 	- **lp_sd**: standard definition
	//
	// 	- **lp_hd**: high definition
	//
	// 	- **lp_ud**: ultra-high definition
	//
	// example:
	//
	// lp_ud
	OutputTemplate *string `json:"OutputTemplate,omitempty" xml:"OutputTemplate,omitempty"`
	// The output URLs.
	OutputUrls *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListOutputUrls `json:"OutputUrls,omitempty" xml:"OutputUrls,omitempty" type:"Struct"`
	// The ID of the region. Valid values:
	//
	// 	- cn-shanghai: China (Shanghai)
	//
	// 	- cn-beijing: China (Beijing)
	//
	// 	- ap-southeast-1: Singapore
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The start time of live monitoring. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-01-11T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the monitoring session. Valid values:
	//
	// 	- 1: Monitoring
	//
	// 	- 0: Unmonitored
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The end time of live monitoring. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-01-11T12:00:00Z
	StopTime *string `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
}

func (s DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList) GetAudioFrom() *int32 {
	return s.AudioFrom
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList) GetDingTalkWebHookUrl() *string {
	return s.DingTalkWebHookUrl
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList) GetDomain() *string {
	return s.Domain
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList) GetInputList() []*DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputList {
	return s.InputList
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList) GetMonitorConfig() *string {
	return s.MonitorConfig
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList) GetMonitorId() *string {
	return s.MonitorId
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList) GetMonitorName() *string {
	return s.MonitorName
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList) GetOutputTemplate() *string {
	return s.OutputTemplate
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList) GetOutputUrls() *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListOutputUrls {
	return s.OutputUrls
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList) GetRegion() *string {
	return s.Region
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList) GetStopTime() *string {
	return s.StopTime
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList) SetAudioFrom(v int32) *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList {
	s.AudioFrom = &v
	return s
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList) SetCallbackUrl(v string) *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList {
	s.CallbackUrl = &v
	return s
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList) SetDingTalkWebHookUrl(v string) *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList {
	s.DingTalkWebHookUrl = &v
	return s
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList) SetDomain(v string) *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList {
	s.Domain = &v
	return s
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList) SetInputList(v []*DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputList) *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList {
	s.InputList = v
	return s
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList) SetMonitorConfig(v string) *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList {
	s.MonitorConfig = &v
	return s
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList) SetMonitorId(v string) *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList {
	s.MonitorId = &v
	return s
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList) SetMonitorName(v string) *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList {
	s.MonitorName = &v
	return s
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList) SetOutputTemplate(v string) *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList {
	s.OutputTemplate = &v
	return s
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList) SetOutputUrls(v *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListOutputUrls) *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList {
	s.OutputUrls = v
	return s
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList) SetRegion(v string) *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList {
	s.Region = &v
	return s
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList) SetStartTime(v string) *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList) SetStatus(v int32) *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList {
	s.Status = &v
	return s
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList) SetStopTime(v string) *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList {
	s.StopTime = &v
	return s
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorList) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputList struct {
	// The index.
	//
	// example:
	//
	// 1
	Index *int32 `json:"Index,omitempty" xml:"Index,omitempty"`
	// The URL of the input stream.
	//
	// example:
	//
	// demo.aliyundoc.com
	InputUrl *string `json:"InputUrl,omitempty" xml:"InputUrl,omitempty"`
	// The layout information.
	LayoutConfig *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputListLayoutConfig `json:"LayoutConfig,omitempty" xml:"LayoutConfig,omitempty" type:"Struct"`
	// The layout ID, which must start from 1.
	//
	// example:
	//
	// 1
	LayoutId *int32 `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	// The playback configurations.
	PlayConfig *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputListPlayConfig `json:"PlayConfig,omitempty" xml:"PlayConfig,omitempty" type:"Struct"`
	// The display name of the monitored stream.
	//
	// example:
	//
	// monitorStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
}

func (s DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputList) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputList) GetIndex() *int32 {
	return s.Index
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputList) GetInputUrl() *string {
	return s.InputUrl
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputList) GetLayoutConfig() *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputListLayoutConfig {
	return s.LayoutConfig
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputList) GetLayoutId() *int32 {
	return s.LayoutId
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputList) GetPlayConfig() *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputListPlayConfig {
	return s.PlayConfig
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputList) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputList) SetIndex(v int32) *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputList {
	s.Index = &v
	return s
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputList) SetInputUrl(v string) *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputList {
	s.InputUrl = &v
	return s
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputList) SetLayoutConfig(v *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputListLayoutConfig) *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputList {
	s.LayoutConfig = v
	return s
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputList) SetLayoutId(v int32) *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputList {
	s.LayoutId = &v
	return s
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputList) SetPlayConfig(v *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputListPlayConfig) *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputList {
	s.PlayConfig = v
	return s
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputList) SetStreamName(v string) *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputList {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputList) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputListLayoutConfig struct {
	// The fill type. Set this value to none.
	//
	// example:
	//
	// none
	FillMode *string `json:"FillMode,omitempty" xml:"FillMode,omitempty"`
	// The position of the layer, in the format of [unk][x,y][unk]. The values of x and y need to be normalized.
	PositionNormalized []*float32 `json:"PositionNormalized,omitempty" xml:"PositionNormalized,omitempty" type:"Repeated"`
	// The reference position of the element. Valid values:
	//
	// 	- topLeft
	//
	// 	- topRight
	//
	// 	- bottomLeft
	//
	// 	- bottomRight
	//
	// example:
	//
	// topLeft
	PositionRefer *string `json:"PositionRefer,omitempty" xml:"PositionRefer,omitempty"`
	// The size of the layer. Unit: bytes.
	SizeNormalized []*float32 `json:"SizeNormalized,omitempty" xml:"SizeNormalized,omitempty" type:"Repeated"`
}

func (s DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputListLayoutConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputListLayoutConfig) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputListLayoutConfig) GetFillMode() *string {
	return s.FillMode
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputListLayoutConfig) GetPositionNormalized() []*float32 {
	return s.PositionNormalized
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputListLayoutConfig) GetPositionRefer() *string {
	return s.PositionRefer
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputListLayoutConfig) GetSizeNormalized() []*float32 {
	return s.SizeNormalized
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputListLayoutConfig) SetFillMode(v string) *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputListLayoutConfig {
	s.FillMode = &v
	return s
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputListLayoutConfig) SetPositionNormalized(v []*float32) *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputListLayoutConfig {
	s.PositionNormalized = v
	return s
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputListLayoutConfig) SetPositionRefer(v string) *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputListLayoutConfig {
	s.PositionRefer = &v
	return s
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputListLayoutConfig) SetSizeNormalized(v []*float32) *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputListLayoutConfig {
	s.SizeNormalized = v
	return s
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputListLayoutConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputListPlayConfig struct {
	// The volume. Valid values: 0 to 1. The value is rounded to two decimal places.
	//
	// example:
	//
	// 0.50
	VolumeRate *float32 `json:"VolumeRate,omitempty" xml:"VolumeRate,omitempty"`
}

func (s DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputListPlayConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputListPlayConfig) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputListPlayConfig) GetVolumeRate() *float32 {
	return s.VolumeRate
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputListPlayConfig) SetVolumeRate(v float32) *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputListPlayConfig {
	s.VolumeRate = &v
	return s
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListInputListPlayConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListOutputUrls struct {
	// The output URL in the Flash Video (FLV) format.
	//
	// example:
	//
	// http://demo.aliyundoc.com/monitor/445409ec-7eaa-461d-8f29-4bec2eb9****.flv
	FlvUrl *string `json:"FlvUrl,omitempty" xml:"FlvUrl,omitempty"`
	// The output URL in the Real-Time Messaging Protocol (RTMP) format.
	//
	// example:
	//
	// rtmp://demo.aliyundoc.com/monitor/445409ec-7eaa-461d-8f29-4bec2eb9****
	RtmpUrl *string `json:"RtmpUrl,omitempty" xml:"RtmpUrl,omitempty"`
}

func (s DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListOutputUrls) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListOutputUrls) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListOutputUrls) GetFlvUrl() *string {
	return s.FlvUrl
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListOutputUrls) GetRtmpUrl() *string {
	return s.RtmpUrl
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListOutputUrls) SetFlvUrl(v string) *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListOutputUrls {
	s.FlvUrl = &v
	return s
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListOutputUrls) SetRtmpUrl(v string) *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListOutputUrls {
	s.RtmpUrl = &v
	return s
}

func (s *DescribeLiveStreamMonitorListResponseBodyLiveStreamMonitorListOutputUrls) Validate() error {
	return dara.Validate(s)
}
