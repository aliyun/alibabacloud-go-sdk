// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRtcChannelMetricResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChannelMetricInfo(v *DescribeRtcChannelMetricResponseBodyChannelMetricInfo) *DescribeRtcChannelMetricResponseBody
	GetChannelMetricInfo() *DescribeRtcChannelMetricResponseBodyChannelMetricInfo
	SetRequestId(v string) *DescribeRtcChannelMetricResponseBody
	GetRequestId() *string
}

type DescribeRtcChannelMetricResponseBody struct {
	ChannelMetricInfo *DescribeRtcChannelMetricResponseBodyChannelMetricInfo `json:"ChannelMetricInfo,omitempty" xml:"ChannelMetricInfo,omitempty" type:"Struct"`
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRtcChannelMetricResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcChannelMetricResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelMetricResponseBody) GetChannelMetricInfo() *DescribeRtcChannelMetricResponseBodyChannelMetricInfo {
	return s.ChannelMetricInfo
}

func (s *DescribeRtcChannelMetricResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRtcChannelMetricResponseBody) SetChannelMetricInfo(v *DescribeRtcChannelMetricResponseBodyChannelMetricInfo) *DescribeRtcChannelMetricResponseBody {
	s.ChannelMetricInfo = v
	return s
}

func (s *DescribeRtcChannelMetricResponseBody) SetRequestId(v string) *DescribeRtcChannelMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRtcChannelMetricResponseBodyChannelMetricInfo struct {
	ChannelMetric *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric `json:"ChannelMetric,omitempty" xml:"ChannelMetric,omitempty" type:"Struct"`
	Duration      *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDuration      `json:"Duration,omitempty" xml:"Duration,omitempty" type:"Struct"`
}

func (s DescribeRtcChannelMetricResponseBodyChannelMetricInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcChannelMetricResponseBodyChannelMetricInfo) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfo) GetChannelMetric() *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric {
	return s.ChannelMetric
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfo) GetDuration() *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDuration {
	return s.Duration
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfo) SetChannelMetric(v *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric) *DescribeRtcChannelMetricResponseBodyChannelMetricInfo {
	s.ChannelMetric = v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfo) SetDuration(v *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDuration) *DescribeRtcChannelMetricResponseBodyChannelMetricInfo {
	s.Duration = v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric struct {
	// example:
	//
	// example_channel
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// 2019-06-06T18:57:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 10
	PubUserCount *int32 `json:"PubUserCount,omitempty" xml:"PubUserCount,omitempty"`
	// example:
	//
	// 2019-06-06T17:57:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 25
	SubUserCount *int32 `json:"SubUserCount,omitempty" xml:"SubUserCount,omitempty"`
	// example:
	//
	// 30
	UserCount *int32 `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
}

func (s DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric) GetPubUserCount() *int32 {
	return s.PubUserCount
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric) GetSubUserCount() *int32 {
	return s.SubUserCount
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric) GetUserCount() *int32 {
	return s.UserCount
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric) SetChannelId(v string) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric {
	s.ChannelId = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric) SetEndTime(v string) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric {
	s.EndTime = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric) SetPubUserCount(v int32) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric {
	s.PubUserCount = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric) SetStartTime(v string) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric {
	s.StartTime = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric) SetSubUserCount(v int32) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric {
	s.SubUserCount = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric) SetUserCount(v int32) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric {
	s.UserCount = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric) Validate() error {
	return dara.Validate(s)
}

type DescribeRtcChannelMetricResponseBodyChannelMetricInfoDuration struct {
	PubDuration *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration `json:"PubDuration,omitempty" xml:"PubDuration,omitempty" type:"Struct"`
	SubDuration *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration `json:"SubDuration,omitempty" xml:"SubDuration,omitempty" type:"Struct"`
}

func (s DescribeRtcChannelMetricResponseBodyChannelMetricInfoDuration) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcChannelMetricResponseBodyChannelMetricInfoDuration) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDuration) GetPubDuration() *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration {
	return s.PubDuration
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDuration) GetSubDuration() *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration {
	return s.SubDuration
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDuration) SetPubDuration(v *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDuration {
	s.PubDuration = v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDuration) SetSubDuration(v *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDuration {
	s.SubDuration = v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDuration) Validate() error {
	return dara.Validate(s)
}

type DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration struct {
	// example:
	//
	// 100
	Audio *int32 `json:"Audio,omitempty" xml:"Audio,omitempty"`
	// example:
	//
	// 100
	Content *int32 `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 100
	Video1080 *int32 `json:"Video1080,omitempty" xml:"Video1080,omitempty"`
	// example:
	//
	// 100
	Video360 *int32 `json:"Video360,omitempty" xml:"Video360,omitempty"`
	// example:
	//
	// 100
	Video720 *int32 `json:"Video720,omitempty" xml:"Video720,omitempty"`
}

func (s DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration) GetAudio() *int32 {
	return s.Audio
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration) GetContent() *int32 {
	return s.Content
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration) GetVideo1080() *int32 {
	return s.Video1080
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration) GetVideo360() *int32 {
	return s.Video360
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration) GetVideo720() *int32 {
	return s.Video720
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration) SetAudio(v int32) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration {
	s.Audio = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration) SetContent(v int32) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration {
	s.Content = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration) SetVideo1080(v int32) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration {
	s.Video1080 = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration) SetVideo360(v int32) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration {
	s.Video360 = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration) SetVideo720(v int32) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration {
	s.Video720 = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration) Validate() error {
	return dara.Validate(s)
}

type DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration struct {
	// example:
	//
	// 100
	Audio *int32 `json:"Audio,omitempty" xml:"Audio,omitempty"`
	// example:
	//
	// 100
	Content *int32 `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 100
	Video1080 *int32 `json:"Video1080,omitempty" xml:"Video1080,omitempty"`
	// example:
	//
	// 100
	Video360 *int32 `json:"Video360,omitempty" xml:"Video360,omitempty"`
	// example:
	//
	// 100
	Video720 *int32 `json:"Video720,omitempty" xml:"Video720,omitempty"`
}

func (s DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration) GetAudio() *int32 {
	return s.Audio
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration) GetContent() *int32 {
	return s.Content
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration) GetVideo1080() *int32 {
	return s.Video1080
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration) GetVideo360() *int32 {
	return s.Video360
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration) GetVideo720() *int32 {
	return s.Video720
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration) SetAudio(v int32) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration {
	s.Audio = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration) SetContent(v int32) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration {
	s.Content = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration) SetVideo1080(v int32) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration {
	s.Video1080 = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration) SetVideo360(v int32) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration {
	s.Video360 = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration) SetVideo720(v int32) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration {
	s.Video720 = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration) Validate() error {
	return dara.Validate(s)
}
