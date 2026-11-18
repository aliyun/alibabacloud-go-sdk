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
	// The end time of the query. Format: yyyy-mm-ddThh:mm:ssZ (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025-06-05T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the queried data. Valid values: **5m**, **1h**, and **1d**. The supported time granularity varies based on the time span specified by `StartTime` and `EndTime`:
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
	// The metric type. You can specify multiple metrics, up to three.
	//
	// >
	//
	// > - Percentage data is returned in decimal format.
	//
	// Quality of Service (QoS) metrics for playback quality:
	//
	// - Vv: Number of plays.
	//
	// - RealVv: Actual play volume.
	//
	// - FirstFrame: First frame time.
	//
	// - SecondPlayRate: 1-second startup rate.
	//
	// - SlowPlayRate: Slow startup rate.
	//
	// - StuckCountRate: Stuttering count rate.
	//
	// - SeekDuration: Seek duration.
	//
	// - StuckDuration100s: Stuttering duration per 100 seconds.
	//
	// - StuckCount100s: Stuttering count per 100 seconds.
	//
	// - PlayFailRate: Playback failure rate.
	//
	// - SeedFailRate: Unstarted playback rate.
	//
	// - AvgPlayBitrate: Average playback bitrate.
	//
	// - AvgStartBitrate: Average startup bitrate.
	//
	// - ErrorCount100s: Error count per 100 seconds.
	//
	// Quality of Experience (QoE) metrics for playback experience:
	//
	// - Uv: Number of playback users.
	//
	// - AvgPerVv: Average number of plays per user.
	//
	// - AvgVideoDuration: Average video duration.
	//
	// - AvgPerPlayDuration: Average playback duration per user.
	//
	// - AvgPerCompletionVv: Average number of completed plays per user.
	//
	// - CompletionVv: Completed play volume.
	//
	// - CompletionRate: Completion rate.
	//
	// - AvgPlayDuration: Average playback duration.
	//
	// - JumpRate5s: 5-second bounce rate.
	//
	// This parameter is required.
	//
	// example:
	//
	// Vv,Uv,AvgPerVv
	Metrics *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// The operating system of the playback terminal. You can specify this parameter to perform a filtered query for the playback data of a specific operating system. Valid values: **Android**, **iOS**, **Harmony**, **Windows**, **macOS**, and **Linux**.
	//
	// The valid values vary based on the terminal type:
	//
	// - **native**: Android, iOS, and Harmony.
	//
	// - **web**: Android, iOS, Harmony, Windows, macOS, and Linux.
	//
	// Separate multiple operating systems with #_#.
	//
	// example:
	//
	// Android、iOS、Windows
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// The time range used to confirm the period-over-period analysis, in days (d).
	//
	// For example, if you set the value to 1d (1 day), the period-over-period time range is from StartTime-1d to EndTime-1d.
	//
	// example:
	//
	// 1d
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The start time of the query. Format: <i>yyyy-mm-dd</i>T<i>hh:mm:ss</i>Z (UTC).
	//
	// >
	//
	// > - You can query historical playback data for the last year.
	//
	// > - The time range for a single query cannot exceed 31 days.
	//
	// > - The time range is a left-closed and right-open interval [StartTime, EndTime).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025-06-24T00:55:06Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The terminal type. Valid values:
	//
	// - **web**: Web client.
	//
	// - **mobile**: Native client.
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
