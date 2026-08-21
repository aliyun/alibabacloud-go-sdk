// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodPlayerCollectDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeVodPlayerCollectDataRequest
	GetAppId() *string
	SetEndTime(v string) *DescribeVodPlayerCollectDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeVodPlayerCollectDataRequest
	GetInterval() *string
	SetMetrics(v string) *DescribeVodPlayerCollectDataRequest
	GetMetrics() *string
	SetOs(v string) *DescribeVodPlayerCollectDataRequest
	GetOs() *string
	SetPeriod(v string) *DescribeVodPlayerCollectDataRequest
	GetPeriod() *string
	SetStartTime(v string) *DescribeVodPlayerCollectDataRequest
	GetStartTime() *string
	SetTerminalType(v string) *DescribeVodPlayerCollectDataRequest
	GetTerminalType() *string
}

type DescribeVodPlayerCollectDataRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// app-1000000
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The end time of the query. Format: yyyy-mm-ddthh:mm:ssz (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025-06-05T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity for the query data. Valid values: **5m**, **1h**, and **1d**. The supported time granularity varies based on the time span specified by `StartTime` and `EndTime`:
	//
	// - Within 3 days: **5m**, **1h**, and **1d**.
	//
	// - 4 to 7 days: **1h*	- and **1d**.
	//
	// - More than 7 days: **1d**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1d
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The metric type. You can specify up to 3 metrics.
	//
	// >
	//
	// > - Percentage data is returned in decimal format.
	//
	// Playback quality (QoS) metrics:
	//
	// - Vv: play count.
	//
	// - RealVv: actual play count.
	//
	// - FirstFrame: first frame time.
	//
	// - SecondPlayRate: instant play rate.
	//
	// - SlowPlayRate: slow play rate.
	//
	// - StuckCountRate: stuttering rate by count.
	//
	// - SeekDuration: seek duration.
	//
	// - StuckDuration100s: stuttering duration per 100 seconds.
	//
	// - StuckCount100s: stuttering count per 100 seconds.
	//
	// - PlayFailRate: playback failure rate.
	//
	// - SeedFailRate: non-play rate.
	//
	// - AvgPlayBitrate: average playback bitrate.
	//
	// - AvgStartBitrate: average initial bitrate.
	//
	// - ErrorCount100s: error count per 100 seconds.
	//
	// Playback experience (QoE) metrics:
	//
	// - Uv: unique viewers.
	//
	// - AvgPerVv: average plays per user.
	//
	// - AvgVideoDuration: average video duration.
	//
	// - AvgPerPlayDuration: average playback duration per user.
	//
	// - AvgPerCompletionVv: average completion count per user.
	//
	// - CompletionVv: completion count.
	//
	// - CompletionRate: completion rate.
	//
	// - AvgPlayDuration: average playback duration.
	//
	// - JumpRate5s: 5-second bounce rate.
	//
	// This parameter is required.
	//
	// example:
	//
	// Vv,Uv,AvgPerVv
	Metrics *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// The operating system of the playback device. Specify this parameter to perform a filtered query for playback data of a specific operating system. Valid values: **Android**, **iOS**, **Harmony**, **Windows**, **MacOS**, and **Linux**.
	//
	// The available values vary by terminal type:
	//
	// - **native**: Android, iOS, Harmony.
	//
	// - **web**: Android, iOS, Harmony, Windows, MacOs, Linux.
	//
	// Separate multiple values with #_#.
	//
	// example:
	//
	// Android、iOS、Windows
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// The time range for period-over-period analysis, in days (d).
	//
	// For example, if you set this parameter to 1d (1 day), the period-over-period data is retrieved from the time range of StartTime-1d to EndTime-1d.
	//
	// example:
	//
	// 1d
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The start time of the query. Format: <i>yyyy-mm-dd</i>t<i>hh:mm:ss</i>z (UTC).
	//
	// >
	//
	// > - Playback data from the last year can be queried.
	//
	// > - The time range for a single query cannot exceed 31 days.
	//
	// > - The time interval is left-closed and right-open [StartTime, EndTime).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025-06-24T00:55:06Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The terminal type. Valid values:
	//
	// - **web**: web.
	//
	// - **mobile**: native.
	//
	// This parameter is required.
	//
	// example:
	//
	// web
	TerminalType *string `json:"TerminalType,omitempty" xml:"TerminalType,omitempty"`
}

func (s DescribeVodPlayerCollectDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodPlayerCollectDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodPlayerCollectDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeVodPlayerCollectDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodPlayerCollectDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeVodPlayerCollectDataRequest) GetMetrics() *string {
	return s.Metrics
}

func (s *DescribeVodPlayerCollectDataRequest) GetOs() *string {
	return s.Os
}

func (s *DescribeVodPlayerCollectDataRequest) GetPeriod() *string {
	return s.Period
}

func (s *DescribeVodPlayerCollectDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodPlayerCollectDataRequest) GetTerminalType() *string {
	return s.TerminalType
}

func (s *DescribeVodPlayerCollectDataRequest) SetAppId(v string) *DescribeVodPlayerCollectDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeVodPlayerCollectDataRequest) SetEndTime(v string) *DescribeVodPlayerCollectDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodPlayerCollectDataRequest) SetInterval(v string) *DescribeVodPlayerCollectDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVodPlayerCollectDataRequest) SetMetrics(v string) *DescribeVodPlayerCollectDataRequest {
	s.Metrics = &v
	return s
}

func (s *DescribeVodPlayerCollectDataRequest) SetOs(v string) *DescribeVodPlayerCollectDataRequest {
	s.Os = &v
	return s
}

func (s *DescribeVodPlayerCollectDataRequest) SetPeriod(v string) *DescribeVodPlayerCollectDataRequest {
	s.Period = &v
	return s
}

func (s *DescribeVodPlayerCollectDataRequest) SetStartTime(v string) *DescribeVodPlayerCollectDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodPlayerCollectDataRequest) SetTerminalType(v string) *DescribeVodPlayerCollectDataRequest {
	s.TerminalType = &v
	return s
}

func (s *DescribeVodPlayerCollectDataRequest) Validate() error {
	return dara.Validate(s)
}
