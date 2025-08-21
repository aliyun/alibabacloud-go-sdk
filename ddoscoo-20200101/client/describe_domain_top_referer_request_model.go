// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainTopRefererRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeDomainTopRefererRequest
	GetDomain() *string
	SetEndTime(v int64) *DescribeDomainTopRefererRequest
	GetEndTime() *int64
	SetLimit(v int64) *DescribeDomainTopRefererRequest
	GetLimit() *int64
	SetRegion(v string) *DescribeDomainTopRefererRequest
	GetRegion() *string
	SetStartTime(v int64) *DescribeDomainTopRefererRequest
	GetStartTime() *int64
}

type DescribeDomainTopRefererRequest struct {
	// The domain name of the website.
	//
	// >  A forwarding rule must be configured for the domain name. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all domain names.
	//
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// >  This UNIX timestamp must indicate a point in time that is accurate to the minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1721561100
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The maximum number of entries to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
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
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// >  This UNIX timestamp must indicate a point in time that is accurate to the minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1701991920
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainTopRefererRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopRefererRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopRefererRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainTopRefererRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDomainTopRefererRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *DescribeDomainTopRefererRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeDomainTopRefererRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDomainTopRefererRequest) SetDomain(v string) *DescribeDomainTopRefererRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainTopRefererRequest) SetEndTime(v int64) *DescribeDomainTopRefererRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainTopRefererRequest) SetLimit(v int64) *DescribeDomainTopRefererRequest {
	s.Limit = &v
	return s
}

func (s *DescribeDomainTopRefererRequest) SetRegion(v string) *DescribeDomainTopRefererRequest {
	s.Region = &v
	return s
}

func (s *DescribeDomainTopRefererRequest) SetStartTime(v int64) *DescribeDomainTopRefererRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainTopRefererRequest) Validate() error {
	return dara.Validate(s)
}
