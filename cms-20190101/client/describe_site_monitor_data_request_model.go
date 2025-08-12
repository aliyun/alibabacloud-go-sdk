// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteMonitorDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeSiteMonitorDataRequest
	GetEndTime() *string
	SetLength(v int32) *DescribeSiteMonitorDataRequest
	GetLength() *int32
	SetMetricName(v string) *DescribeSiteMonitorDataRequest
	GetMetricName() *string
	SetNextToken(v string) *DescribeSiteMonitorDataRequest
	GetNextToken() *string
	SetPeriod(v string) *DescribeSiteMonitorDataRequest
	GetPeriod() *string
	SetRegionId(v string) *DescribeSiteMonitorDataRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeSiteMonitorDataRequest
	GetStartTime() *string
	SetTaskId(v string) *DescribeSiteMonitorDataRequest
	GetTaskId() *string
	SetType(v string) *DescribeSiteMonitorDataRequest
	GetType() *string
}

type DescribeSiteMonitorDataRequest struct {
	// The end of the time range to query. The following formats are supported:
	//
	// 	- UNIX timestamp: the number of milliseconds that have elapsed since 00:00:00 UTC on Thursday, January 1, 1970.
	//
	// 	- UTC time: the UTC time that follows the YYYY-MM-DDThh:mm:ssZ format.
	//
	// example:
	//
	// 1551581437000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of data points to return.
	//
	// example:
	//
	// 1000
	Length *int32 `json:"Length,omitempty" xml:"Length,omitempty"`
	// The metric name. Valid values:
	//
	// 	- Availability
	//
	// 	- ResponseTime
	//
	// This parameter is required.
	//
	// example:
	//
	// Availability
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// 49f7b317-7645-4cc9-94fd-ea42e5220930ea42e5220930ea42e522****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The statistical period. The value is an integral multiple of 60. Unit: seconds.
	//
	// >  The default value equals the minimum interval at which detection requests are sent to the monitored address.
	//
	// example:
	//
	// 60
	Period   *string `json:"Period,omitempty" xml:"Period,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start of the time range to query. The following formats are supported:
	//
	// 	- UNIX timestamp: the number of milliseconds that have elapsed since 00:00:00 UTC on Thursday, January 1, 1970.
	//
	// 	- UTC time: the UTC time that follows the YYYY-MM-DDThh:mm:ssZ format.
	//
	// example:
	//
	// 1551579637000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 49f7b317-7645-4cc9-94fd-ea42e522****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The type of the monitored object whose monitoring data is to be queried. Valid values:
	//
	// 	- metric
	//
	// 	- event
	//
	// example:
	//
	// metric
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSiteMonitorDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSiteMonitorDataRequest) GetLength() *int32 {
	return s.Length
}

func (s *DescribeSiteMonitorDataRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeSiteMonitorDataRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSiteMonitorDataRequest) GetPeriod() *string {
	return s.Period
}

func (s *DescribeSiteMonitorDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSiteMonitorDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSiteMonitorDataRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeSiteMonitorDataRequest) GetType() *string {
	return s.Type
}

func (s *DescribeSiteMonitorDataRequest) SetEndTime(v string) *DescribeSiteMonitorDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSiteMonitorDataRequest) SetLength(v int32) *DescribeSiteMonitorDataRequest {
	s.Length = &v
	return s
}

func (s *DescribeSiteMonitorDataRequest) SetMetricName(v string) *DescribeSiteMonitorDataRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeSiteMonitorDataRequest) SetNextToken(v string) *DescribeSiteMonitorDataRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeSiteMonitorDataRequest) SetPeriod(v string) *DescribeSiteMonitorDataRequest {
	s.Period = &v
	return s
}

func (s *DescribeSiteMonitorDataRequest) SetRegionId(v string) *DescribeSiteMonitorDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSiteMonitorDataRequest) SetStartTime(v string) *DescribeSiteMonitorDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSiteMonitorDataRequest) SetTaskId(v string) *DescribeSiteMonitorDataRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeSiteMonitorDataRequest) SetType(v string) *DescribeSiteMonitorDataRequest {
	s.Type = &v
	return s
}

func (s *DescribeSiteMonitorDataRequest) Validate() error {
	return dara.Validate(s)
}
