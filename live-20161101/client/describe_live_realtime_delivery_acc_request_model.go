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
	// 	- You can query one or more domain names. If you specify multiple domain names, separate them with commas (,).
	//
	// 	- If you leave this parameter empty, the data of all domain names within your Alibaba Cloud account is returned.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// The end time must be later than the start time. The maximum time range that can be specified is one year.
	//
	// example:
	//
	// 2015-12-10T21:05:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the query. Unit: seconds. Valid values:
	//
	// 	- **300**
	//
	// 	- **3600**
	//
	// 	- **86400**
	//
	// If you specify an invalid value or do not specify this parameter, the default value is used. If the specified time range is no more than three days, the default value is 300. If the specified time range is more than three days and no more than 30 days, the default value is 3600. If the specified time range is more than 30 days, the default value is 86400.
	//
	// example:
	//
	// 3600
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The name of the Logstore to which log entries are delivered. If you leave this parameter empty, the data of all Logstores is returned.
	//
	// example:
	//
	// logstore_example
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the Log Service project that is used for real-time log delivery. If you leave this parameter empty, the data of all Log Service projects is returned.
	//
	// example:
	//
	// project_example
	Project  *string `json:"Project,omitempty" xml:"Project,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
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
