// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainBpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeDomainBpsRequest
	GetDomain() *string
	SetEndTime(v int64) *DescribeDomainBpsRequest
	GetEndTime() *int64
	SetInterval(v int64) *DescribeDomainBpsRequest
	GetInterval() *int64
	SetRegion(v string) *DescribeDomainBpsRequest
	GetRegion() *string
	SetStartTime(v int64) *DescribeDomainBpsRequest
	GetStartTime() *int64
}

type DescribeDomainBpsRequest struct {
	// The domain name of the website.
	//
	// >  A forwarding rule must be configured for the domain name. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all domain names.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The end of the time range to query. This value is a UNIX timestamp. Unit: seconds.
	//
	// >  This UNIX timestamp must indicate a point in time that is accurate to the minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1722339300
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The interval for returning data. Unit: seconds. Valid values are 300, 3600, and 86400. If the time span between StartTime and EndTime is less than 3 days, valid values are 300, 3600, and 86400. If the time span between StartTime and EndTime is from 3 to 30 days, valid values are 3600 and 86400. If the time span between StartTime and EndTime is 31 days or longer, the valid value is 86400. If you leave this parameter empty or specify an invalid value, the default value is used.
	//
	// This parameter is required.
	//
	// example:
	//
	// 600
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The region in which your service is deployed. Valid values:
	//
	// 	- **cn**: a region in the Chinese mainland.
	//
	// 	- **cn-hongkong**: a region outside the Chinese mainland.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The beginning of the time range to query. This value is a UNIX timestamp. Unit: seconds.
	//
	// >  This UNIX timestamp must indicate a point in time that is accurate to the minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1719211800
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainBpsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainBpsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainBpsRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainBpsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDomainBpsRequest) GetInterval() *int64 {
	return s.Interval
}

func (s *DescribeDomainBpsRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeDomainBpsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDomainBpsRequest) SetDomain(v string) *DescribeDomainBpsRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainBpsRequest) SetEndTime(v int64) *DescribeDomainBpsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainBpsRequest) SetInterval(v int64) *DescribeDomainBpsRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainBpsRequest) SetRegion(v string) *DescribeDomainBpsRequest {
	s.Region = &v
	return s
}

func (s *DescribeDomainBpsRequest) SetStartTime(v int64) *DescribeDomainBpsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainBpsRequest) Validate() error {
	return dara.Validate(s)
}
