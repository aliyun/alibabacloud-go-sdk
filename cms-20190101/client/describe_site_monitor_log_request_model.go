// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteMonitorLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBrowser(v string) *DescribeSiteMonitorLogRequest
	GetBrowser() *string
	SetBrowserInfo(v string) *DescribeSiteMonitorLogRequest
	GetBrowserInfo() *string
	SetCity(v string) *DescribeSiteMonitorLogRequest
	GetCity() *string
	SetDevice(v string) *DescribeSiteMonitorLogRequest
	GetDevice() *string
	SetEndTime(v string) *DescribeSiteMonitorLogRequest
	GetEndTime() *string
	SetFilter(v string) *DescribeSiteMonitorLogRequest
	GetFilter() *string
	SetIsp(v string) *DescribeSiteMonitorLogRequest
	GetIsp() *string
	SetLength(v int32) *DescribeSiteMonitorLogRequest
	GetLength() *int32
	SetMetricName(v string) *DescribeSiteMonitorLogRequest
	GetMetricName() *string
	SetNextToken(v string) *DescribeSiteMonitorLogRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeSiteMonitorLogRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeSiteMonitorLogRequest
	GetStartTime() *string
	SetTaskIds(v string) *DescribeSiteMonitorLogRequest
	GetTaskIds() *string
}

type DescribeSiteMonitorLogRequest struct {
	// example:
	//
	// Chrome
	Browser     *string `json:"Browser,omitempty" xml:"Browser,omitempty"`
	BrowserInfo *string `json:"BrowserInfo,omitempty" xml:"BrowserInfo,omitempty"`
	// The city identification code.
	//
	// example:
	//
	// 546
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// example:
	//
	// laptop
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The end of the time range to query. Valid values:
	//
	// 	- UNIX timestamp: the number of milliseconds that have elapsed since 00:00:00 Thursday, January 1, 1970
	//
	// 	- UTC time: the UTC time that follows the YYYY-MM-DDThh:mm:ssZ format
	//
	// >  We recommend that you use UNIX timestamps to prevent time zone-related issues.
	//
	// example:
	//
	// 1638422475687
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The filter condition.
	//
	// You can specify a simple expression, for example, `TotalTime>100`. In this case, the operation returns only the data for instant test tasks whose total response time exceeds 100 milliseconds.
	//
	// example:
	//
	// TotalTime>100
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The carrier identification code.
	//
	// example:
	//
	// 465
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 1440.
	//
	// example:
	//
	// 1000
	Length *int32 `json:"Length,omitempty" xml:"Length,omitempty"`
	// The name of the metric.
	//
	// Only the `ProbeLog` metric is supported.
	//
	// example:
	//
	// ProbeLog
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The token that is used to initiate the next request if the response of the current request is truncated. You can use the token to initiate another request and obtain the remaining records.``
	//
	// example:
	//
	// IWBjqMYSy0is7zSMGu16****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start of the time range to query. The following formats are supported:
	//
	// 	- UNIX timestamp: the number of milliseconds that have elapsed since 00:00:00 Thursday, January 1, 1970
	//
	// 	- UTC time: the UTC time that follows the YYYY-MM-DDThh:mm:ssZ format
	//
	// >
	//
	// 	- The specified time range includes the end time and excludes the start time. The start time must be earlier than the end time.\\
	//
	//     We recommend that you use UNIX timestamps to prevent time zone-related issues.
	//
	// example:
	//
	// 1638422474389
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The IDs of the instant test tasks. Separate multiple task IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// afa5c3ce-f944-4363-9edb-ce919a29****
	TaskIds *string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty"`
}

func (s DescribeSiteMonitorLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorLogRequest) GetBrowser() *string {
	return s.Browser
}

func (s *DescribeSiteMonitorLogRequest) GetBrowserInfo() *string {
	return s.BrowserInfo
}

func (s *DescribeSiteMonitorLogRequest) GetCity() *string {
	return s.City
}

func (s *DescribeSiteMonitorLogRequest) GetDevice() *string {
	return s.Device
}

func (s *DescribeSiteMonitorLogRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSiteMonitorLogRequest) GetFilter() *string {
	return s.Filter
}

func (s *DescribeSiteMonitorLogRequest) GetIsp() *string {
	return s.Isp
}

func (s *DescribeSiteMonitorLogRequest) GetLength() *int32 {
	return s.Length
}

func (s *DescribeSiteMonitorLogRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeSiteMonitorLogRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSiteMonitorLogRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSiteMonitorLogRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSiteMonitorLogRequest) GetTaskIds() *string {
	return s.TaskIds
}

func (s *DescribeSiteMonitorLogRequest) SetBrowser(v string) *DescribeSiteMonitorLogRequest {
	s.Browser = &v
	return s
}

func (s *DescribeSiteMonitorLogRequest) SetBrowserInfo(v string) *DescribeSiteMonitorLogRequest {
	s.BrowserInfo = &v
	return s
}

func (s *DescribeSiteMonitorLogRequest) SetCity(v string) *DescribeSiteMonitorLogRequest {
	s.City = &v
	return s
}

func (s *DescribeSiteMonitorLogRequest) SetDevice(v string) *DescribeSiteMonitorLogRequest {
	s.Device = &v
	return s
}

func (s *DescribeSiteMonitorLogRequest) SetEndTime(v string) *DescribeSiteMonitorLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSiteMonitorLogRequest) SetFilter(v string) *DescribeSiteMonitorLogRequest {
	s.Filter = &v
	return s
}

func (s *DescribeSiteMonitorLogRequest) SetIsp(v string) *DescribeSiteMonitorLogRequest {
	s.Isp = &v
	return s
}

func (s *DescribeSiteMonitorLogRequest) SetLength(v int32) *DescribeSiteMonitorLogRequest {
	s.Length = &v
	return s
}

func (s *DescribeSiteMonitorLogRequest) SetMetricName(v string) *DescribeSiteMonitorLogRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeSiteMonitorLogRequest) SetNextToken(v string) *DescribeSiteMonitorLogRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeSiteMonitorLogRequest) SetRegionId(v string) *DescribeSiteMonitorLogRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSiteMonitorLogRequest) SetStartTime(v string) *DescribeSiteMonitorLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSiteMonitorLogRequest) SetTaskIds(v string) *DescribeSiteMonitorLogRequest {
	s.TaskIds = &v
	return s
}

func (s *DescribeSiteMonitorLogRequest) Validate() error {
	return dara.Validate(s)
}
