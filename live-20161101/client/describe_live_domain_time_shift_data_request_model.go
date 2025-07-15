// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainTimeShiftDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveDomainTimeShiftDataRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveDomainTimeShiftDataRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeLiveDomainTimeShiftDataRequest
	GetInterval() *string
	SetOwnerId(v int64) *DescribeLiveDomainTimeShiftDataRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveDomainTimeShiftDataRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeLiveDomainTimeShiftDataRequest
	GetStartTime() *string
}

type DescribeLiveDomainTimeShiftDataRequest struct {
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
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2021-03-03T02:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the query. Unit: seconds. Valid values:
	//
	// 	- 300
	//
	// 	- 3600
	//
	// 	- 86400
	//
	// If you do not specify this parameter or specify an invalid value, the default value 300 is used.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// >
	//
	// 	- The minimum data granularity is 5 minutes.
	//
	// 	- If you leave this parameter empty, data in the last 24 hours is queried.
	//
	// example:
	//
	// 2021-03-03T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveDomainTimeShiftDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainTimeShiftDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainTimeShiftDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainTimeShiftDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDomainTimeShiftDataRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeLiveDomainTimeShiftDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveDomainTimeShiftDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveDomainTimeShiftDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDomainTimeShiftDataRequest) SetDomainName(v string) *DescribeLiveDomainTimeShiftDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainTimeShiftDataRequest) SetEndTime(v string) *DescribeLiveDomainTimeShiftDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainTimeShiftDataRequest) SetInterval(v string) *DescribeLiveDomainTimeShiftDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeLiveDomainTimeShiftDataRequest) SetOwnerId(v int64) *DescribeLiveDomainTimeShiftDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveDomainTimeShiftDataRequest) SetRegionId(v string) *DescribeLiveDomainTimeShiftDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveDomainTimeShiftDataRequest) SetStartTime(v string) *DescribeLiveDomainTimeShiftDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainTimeShiftDataRequest) Validate() error {
	return dara.Validate(s)
}
