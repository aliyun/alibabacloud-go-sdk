// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodPlayerMetricDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeVodPlayerMetricDataRequest
	GetAppId() *string
	SetEndTime(v string) *DescribeVodPlayerMetricDataRequest
	GetEndTime() *string
	SetFilters(v string) *DescribeVodPlayerMetricDataRequest
	GetFilters() *string
	SetInterval(v string) *DescribeVodPlayerMetricDataRequest
	GetInterval() *string
	SetLanguage(v string) *DescribeVodPlayerMetricDataRequest
	GetLanguage() *string
	SetMetrics(v string) *DescribeVodPlayerMetricDataRequest
	GetMetrics() *string
	SetOs(v string) *DescribeVodPlayerMetricDataRequest
	GetOs() *string
	SetPageNumber(v int64) *DescribeVodPlayerMetricDataRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeVodPlayerMetricDataRequest
	GetPageSize() *int64
	SetStartTime(v string) *DescribeVodPlayerMetricDataRequest
	GetStartTime() *string
	SetTerminalType(v string) *DescribeVodPlayerMetricDataRequest
	GetTerminalType() *string
	SetTop(v int64) *DescribeVodPlayerMetricDataRequest
	GetTop() *int64
}

type DescribeVodPlayerMetricDataRequest struct {
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
	// The metric dimension filters. A dimension consists of a dimension type (Field), an operator (Op), and a dimension value.
	//
	// > - A maximum of three dimensions can be specified.
	//
	// > - When the Metrics parameter includes the following four metrics, Filters do not take effect: Uv (playback users), AvgPerVv (average plays per user), AvgPerPlayDuration (average play duration per user), and AvgPerCompletionVv (average completion plays per user).
	//
	// > - For provinces and countries, pass the regionCode.
	//
	// > - Separate multiple values with #_#.
	//
	// Valid values for dimension type (Field):
	//
	// - SdkVersion: SDK version.
	//
	// - AppVersion: app version.
	//
	// - Codec: codec.
	//
	// - VideoType: video format.
	//
	// - Network: network type.
	//
	// - Country: country.
	//
	// - Isp: ISP.
	//
	// - VideoDefinition: resolution.
	//
	// - Domain: domain name.
	//
	// - Province: province.
	//
	// - IsHw: whether hardware decoding is used.
	//
	// - ErrorCode: error code.
	//
	// Valid values for operator (Op): = (equal to), > (greater than), < (less than), and != (not equal to).
	//
	// >
	//
	// > - SdkVersion and VideoDefinition support all four operators. Other metrics support only = (equal to) and != (not equal to).
	//
	// Retrieve dimension values by calling DescribeVodPlayerDimensionData.
	//
	// example:
	//
	// [
	//
	//   {
	//
	//     "Field": "codec",
	//
	//     "Op": "=",
	//
	//     "Value": "h265#_#h264"
	//
	//   },
	//
	//   {
	//
	//     "Field": "os",
	//
	//     "Op": "=",
	//
	//     "Value": "Android#_#iOS"
	//
	//   }
	//
	// ]
	Filters *string `json:"Filters,omitempty" xml:"Filters,omitempty"`
	// The time granularity for querying data. Valid values: **5m**, **1h**, and **1d**. The supported time granularity depends on the time span between `StartTime` and `EndTime`:
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
	// The language of the response. Valid values:
	//
	// - **zh*	- (**default**): Simplified Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The metric types. You can select multiple metrics (up to 3).
	//
	// >
	//
	// > - Percentage data is returned in decimal form.
	//
	// Quality of Service (QoS) metrics:
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
	// - PlayFailRate: play failure rate.
	//
	// - SeedFailRate: non-play rate.
	//
	// - AvgPlayBitrate: average playback bitrate.
	//
	// - AvgStartBitrate: average start bitrate.
	//
	// - ErrorCount100s: error count per 100 seconds.
	//
	// Quality of Experience (QoE) metrics:
	//
	// - Uv: playback users.
	//
	// - AvgPerVv: average plays per user.
	//
	// - AvgVideoDuration: average video duration.
	//
	// - AvgPerPlayDuration: average play duration per user.
	//
	// - AvgPerCompletionVv: average completion plays per user.
	//
	// - CompletionVv: completion count.
	//
	// - CompletionRate: completion rate.
	//
	// - AvgPlayDuration: average play duration.
	//
	// - JumpRate5s: 5-second bounce rate.
	//
	// This parameter is required.
	//
	// example:
	//
	// Vv,Uv,AvgPerVv
	Metrics *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// The operating system of the player. Specify this parameter to perform a filtered query for playback data of a specific operating system. Valid values: **Android**, **iOS**, **Harmony**, **Windows**, **MacOS**, and **Linux**.
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
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **5000**. Maximum value: **5000**.
	//
	// example:
	//
	// 5000
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The start time of the query. Format: <i>yyyy-mm-dd</i>t<i>hh:mm:ss</i>z (UTC).
	//
	// >
	//
	// > - Supports querying playback data history for the past year.
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
	// Returns data for the top N items ranked by play count. If this parameter is not specified, data for all dimensions is returned.
	//
	// example:
	//
	// 5
	Top *int64 `json:"Top,omitempty" xml:"Top,omitempty"`
}

func (s DescribeVodPlayerMetricDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodPlayerMetricDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodPlayerMetricDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeVodPlayerMetricDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodPlayerMetricDataRequest) GetFilters() *string {
	return s.Filters
}

func (s *DescribeVodPlayerMetricDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeVodPlayerMetricDataRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeVodPlayerMetricDataRequest) GetMetrics() *string {
	return s.Metrics
}

func (s *DescribeVodPlayerMetricDataRequest) GetOs() *string {
	return s.Os
}

func (s *DescribeVodPlayerMetricDataRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeVodPlayerMetricDataRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeVodPlayerMetricDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodPlayerMetricDataRequest) GetTerminalType() *string {
	return s.TerminalType
}

func (s *DescribeVodPlayerMetricDataRequest) GetTop() *int64 {
	return s.Top
}

func (s *DescribeVodPlayerMetricDataRequest) SetAppId(v string) *DescribeVodPlayerMetricDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeVodPlayerMetricDataRequest) SetEndTime(v string) *DescribeVodPlayerMetricDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodPlayerMetricDataRequest) SetFilters(v string) *DescribeVodPlayerMetricDataRequest {
	s.Filters = &v
	return s
}

func (s *DescribeVodPlayerMetricDataRequest) SetInterval(v string) *DescribeVodPlayerMetricDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeVodPlayerMetricDataRequest) SetLanguage(v string) *DescribeVodPlayerMetricDataRequest {
	s.Language = &v
	return s
}

func (s *DescribeVodPlayerMetricDataRequest) SetMetrics(v string) *DescribeVodPlayerMetricDataRequest {
	s.Metrics = &v
	return s
}

func (s *DescribeVodPlayerMetricDataRequest) SetOs(v string) *DescribeVodPlayerMetricDataRequest {
	s.Os = &v
	return s
}

func (s *DescribeVodPlayerMetricDataRequest) SetPageNumber(v int64) *DescribeVodPlayerMetricDataRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVodPlayerMetricDataRequest) SetPageSize(v int64) *DescribeVodPlayerMetricDataRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVodPlayerMetricDataRequest) SetStartTime(v string) *DescribeVodPlayerMetricDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodPlayerMetricDataRequest) SetTerminalType(v string) *DescribeVodPlayerMetricDataRequest {
	s.TerminalType = &v
	return s
}

func (s *DescribeVodPlayerMetricDataRequest) SetTop(v int64) *DescribeVodPlayerMetricDataRequest {
	s.Top = &v
	return s
}

func (s *DescribeVodPlayerMetricDataRequest) Validate() error {
	return dara.Validate(s)
}
