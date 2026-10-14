// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainTopHttpMethodRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeDomainTopHttpMethodRequest
	GetDomain() *string
	SetEndTime(v int64) *DescribeDomainTopHttpMethodRequest
	GetEndTime() *int64
	SetLimit(v int64) *DescribeDomainTopHttpMethodRequest
	GetLimit() *int64
	SetQueryType(v string) *DescribeDomainTopHttpMethodRequest
	GetQueryType() *string
	SetRegion(v string) *DescribeDomainTopHttpMethodRequest
	GetRegion() *string
	SetStartTime(v int64) *DescribeDomainTopHttpMethodRequest
	GetStartTime() *int64
}

type DescribeDomainTopHttpMethodRequest struct {
	// The domain name of the website.
	//
	// > The domain name must have a website forwarding rule configured. You can call [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) to query all domain names.
	//
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The end time of the query. Specify a UNIX timestamp. Unit: seconds.
	//
	// > The value must be on the hour or on the minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1722339300
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of entries to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	Limit     *int64  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// The region of the service. Valid values:
	//
	// - **cn**: the Chinese mainland.
	//
	// - **cn-hongkong**: outside the Chinese mainland.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The start time of the query. Specify a UNIX timestamp. Unit: seconds.
	//
	// > The value must be on the hour or on the minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1712449710
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainTopHttpMethodRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopHttpMethodRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopHttpMethodRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainTopHttpMethodRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDomainTopHttpMethodRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *DescribeDomainTopHttpMethodRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *DescribeDomainTopHttpMethodRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeDomainTopHttpMethodRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDomainTopHttpMethodRequest) SetDomain(v string) *DescribeDomainTopHttpMethodRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainTopHttpMethodRequest) SetEndTime(v int64) *DescribeDomainTopHttpMethodRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainTopHttpMethodRequest) SetLimit(v int64) *DescribeDomainTopHttpMethodRequest {
	s.Limit = &v
	return s
}

func (s *DescribeDomainTopHttpMethodRequest) SetQueryType(v string) *DescribeDomainTopHttpMethodRequest {
	s.QueryType = &v
	return s
}

func (s *DescribeDomainTopHttpMethodRequest) SetRegion(v string) *DescribeDomainTopHttpMethodRequest {
	s.Region = &v
	return s
}

func (s *DescribeDomainTopHttpMethodRequest) SetStartTime(v int64) *DescribeDomainTopHttpMethodRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainTopHttpMethodRequest) Validate() error {
	return dara.Validate(s)
}
