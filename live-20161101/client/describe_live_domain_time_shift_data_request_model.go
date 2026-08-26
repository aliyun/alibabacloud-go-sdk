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
	// The streaming domain to query.
	//
	// >
	//
	// > - When you specify DomainName, make sure that the specified domain is a live streaming domain and that the user calling this operation has permissions on the specified domain.
	//
	// > - You can specify a single domain or multiple domains. Separate multiple domains with commas (,).
	//
	// > - If this parameter is left empty, the merged data of all live streaming domains is returned by default.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end time. The end time must be later than the start time. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format in UTC.
	//
	// example:
	//
	// 2021-03-03T02:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time granularity of the queried data. Unit: seconds. Valid values:
	//
	// - 300.
	//
	// - 3600.
	//
	// - 86400.
	//
	// If you do not specify this parameter or specify an unsupported value, the default value 300 is used.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format in UTC.
	//
	// > - The minimum data granularity is 5 minutes.
	//
	// > - If you do not specify this parameter, data of the last 24 hours is returned by default.
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
