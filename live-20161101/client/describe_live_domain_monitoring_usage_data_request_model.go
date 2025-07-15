// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainMonitoringUsageDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveDomainMonitoringUsageDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveDomainMonitoringUsageDataRequest
	GetEndTime() *string
	SetInstanceId(v string) *DescribeLiveDomainMonitoringUsageDataRequest
	GetInstanceId() *string
	SetInterval(v string) *DescribeLiveDomainMonitoringUsageDataRequest
	GetInterval() *string
	SetOwnerId(v int64) *DescribeLiveDomainMonitoringUsageDataRequest
	GetOwnerId() *int64
	SetRegion(v string) *DescribeLiveDomainMonitoringUsageDataRequest
	GetRegion() *string
	SetRegionId(v string) *DescribeLiveDomainMonitoringUsageDataRequest
	GetRegionId() *string
	SetSplitBy(v string) *DescribeLiveDomainMonitoringUsageDataRequest
	GetSplitBy() *string
	SetStartTime(v string) *DescribeLiveDomainMonitoringUsageDataRequest
	GetStartTime() *string
}

type DescribeLiveDomainMonitoringUsageDataRequest struct {
	// The main streaming domain to query.
	//
	// 	- You can query one or more domain names. If you specify multiple domain names, separate them with commas (,).
	//
	// 	- If you leave this parameter empty, the data of all domain names within your Alibaba Cloud account is returned.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2022-12-10T22:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the monitoring session. If you leave this parameter empty, data of all monitoring sessions is queried by default. Separate multiple session IDs with commas (,).
	//
	// example:
	//
	// e62af24d-a354-3b0c-9f1f-da592c4b****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time granularity. Valid values: **3600*	- and **86400**. 3600 specifies that data is queried by hour and 86400 specifies that data is queried by day.
	//
	// example:
	//
	// 3600
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region of the live center. If you leave this parameter empty, data of all regions is queried by default. Separate multiple regions with commas (,).
	//
	// example:
	//
	// cn-shanghai
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The key that is used to group data. Valid values: **domain**, **region**, **instance**, and **resolution**. Default value: **resolution**. resolution specifies that data is grouped by resolution. Separate multiple values with commas (,).
	//
	// example:
	//
	// resolution
	SplitBy *string `json:"SplitBy,omitempty" xml:"SplitBy,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format.
	//
	// 	- The time must be in UTC.
	//
	// 	- The minimum data granularity is 1 hour.
	//
	// 	- If you leave this parameter empty, data in the previous 24 hours is queried.
	//
	// example:
	//
	// 2022-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveDomainMonitoringUsageDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainMonitoringUsageDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainMonitoringUsageDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainMonitoringUsageDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDomainMonitoringUsageDataRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeLiveDomainMonitoringUsageDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeLiveDomainMonitoringUsageDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveDomainMonitoringUsageDataRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeLiveDomainMonitoringUsageDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveDomainMonitoringUsageDataRequest) GetSplitBy() *string {
	return s.SplitBy
}

func (s *DescribeLiveDomainMonitoringUsageDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDomainMonitoringUsageDataRequest) SetDomainName(v string) *DescribeLiveDomainMonitoringUsageDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainMonitoringUsageDataRequest) SetEndTime(v string) *DescribeLiveDomainMonitoringUsageDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainMonitoringUsageDataRequest) SetInstanceId(v string) *DescribeLiveDomainMonitoringUsageDataRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeLiveDomainMonitoringUsageDataRequest) SetInterval(v string) *DescribeLiveDomainMonitoringUsageDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeLiveDomainMonitoringUsageDataRequest) SetOwnerId(v int64) *DescribeLiveDomainMonitoringUsageDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveDomainMonitoringUsageDataRequest) SetRegion(v string) *DescribeLiveDomainMonitoringUsageDataRequest {
	s.Region = &v
	return s
}

func (s *DescribeLiveDomainMonitoringUsageDataRequest) SetRegionId(v string) *DescribeLiveDomainMonitoringUsageDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveDomainMonitoringUsageDataRequest) SetSplitBy(v string) *DescribeLiveDomainMonitoringUsageDataRequest {
	s.SplitBy = &v
	return s
}

func (s *DescribeLiveDomainMonitoringUsageDataRequest) SetStartTime(v string) *DescribeLiveDomainMonitoringUsageDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainMonitoringUsageDataRequest) Validate() error {
	return dara.Validate(s)
}
