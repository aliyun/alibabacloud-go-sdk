// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveRealtimeDeliveryAccRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveRealtimeDeliveryAccRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveRealtimeDeliveryAccRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeLiveRealtimeDeliveryAccRequest
	GetInterval() *string
	SetLogStore(v string) *DescribeLiveRealtimeDeliveryAccRequest
	GetLogStore() *string
	SetOwnerId(v int64) *DescribeLiveRealtimeDeliveryAccRequest
	GetOwnerId() *int64
	SetProject(v string) *DescribeLiveRealtimeDeliveryAccRequest
	GetProject() *string
	SetRegionId(v string) *DescribeLiveRealtimeDeliveryAccRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeLiveRealtimeDeliveryAccRequest
	GetStartTime() *string
}

type DescribeLiveRealtimeDeliveryAccRequest struct {
	// The streaming domain.
	//
	// - You can specify a single domain name or multiple domain names. Separate multiple domain names with commas (,).
	//
	// - If this parameter is not specified, the merged data of all live streaming domain names is returned by default.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end time. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// The end time must be later than the start time. The interval between the start time and end time cannot exceed one year.
	//
	// example:
	//
	// 2015-12-10T21:05:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the queried data. Unit: seconds. Valid values:
	//
	// - **300**
	//
	// - **3600**
	//
	// - **86400**
	//
	// If this parameter is not specified or the specified value is not supported, the default value is 300 seconds when the time span does not exceed 3 days, 3600 seconds when the time span exceeds 3 days, and 86400 seconds when the time span exceeds 30 days.
	//
	// example:
	//
	// 3600
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The Logstore for real-time log delivery. If this parameter is not specified, the merged data of all Logstores is returned by default.
	//
	// example:
	//
	// logstore_example
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The Project for real-time log delivery. If this parameter is not specified, the merged data of all Projects is returned by default.
	//
	// example:
	//
	// project_example
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC).
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveRealtimeDeliveryAccRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRealtimeDeliveryAccRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveRealtimeDeliveryAccRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveRealtimeDeliveryAccRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveRealtimeDeliveryAccRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeLiveRealtimeDeliveryAccRequest) GetLogStore() *string {
	return s.LogStore
}

func (s *DescribeLiveRealtimeDeliveryAccRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveRealtimeDeliveryAccRequest) GetProject() *string {
	return s.Project
}

func (s *DescribeLiveRealtimeDeliveryAccRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveRealtimeDeliveryAccRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveRealtimeDeliveryAccRequest) SetDomainName(v string) *DescribeLiveRealtimeDeliveryAccRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveRealtimeDeliveryAccRequest) SetEndTime(v string) *DescribeLiveRealtimeDeliveryAccRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveRealtimeDeliveryAccRequest) SetInterval(v string) *DescribeLiveRealtimeDeliveryAccRequest {
	s.Interval = &v
	return s
}

func (s *DescribeLiveRealtimeDeliveryAccRequest) SetLogStore(v string) *DescribeLiveRealtimeDeliveryAccRequest {
	s.LogStore = &v
	return s
}

func (s *DescribeLiveRealtimeDeliveryAccRequest) SetOwnerId(v int64) *DescribeLiveRealtimeDeliveryAccRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveRealtimeDeliveryAccRequest) SetProject(v string) *DescribeLiveRealtimeDeliveryAccRequest {
	s.Project = &v
	return s
}

func (s *DescribeLiveRealtimeDeliveryAccRequest) SetRegionId(v string) *DescribeLiveRealtimeDeliveryAccRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveRealtimeDeliveryAccRequest) SetStartTime(v string) *DescribeLiveRealtimeDeliveryAccRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveRealtimeDeliveryAccRequest) Validate() error {
	return dara.Validate(s)
}
