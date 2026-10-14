// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainTopFingerprintRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeDomainTopFingerprintRequest
	GetDomain() *string
	SetEndTime(v int64) *DescribeDomainTopFingerprintRequest
	GetEndTime() *int64
	SetInterval(v int64) *DescribeDomainTopFingerprintRequest
	GetInterval() *int64
	SetLimit(v int64) *DescribeDomainTopFingerprintRequest
	GetLimit() *int64
	SetQueryType(v string) *DescribeDomainTopFingerprintRequest
	GetQueryType() *string
	SetRegion(v string) *DescribeDomainTopFingerprintRequest
	GetRegion() *string
	SetStartTime(v int64) *DescribeDomainTopFingerprintRequest
	GetStartTime() *int64
}

type DescribeDomainTopFingerprintRequest struct {
	// The domain name of the website.
	//
	// > The domain name must have a website traffic forwarding rule configured. You can call [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) to query all domain names.
	//
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The end time of the query. Specify a UNIX timestamp. Unit: seconds.
	//
	// > The value must be on the exact minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1723552200
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The step size for returned data. Unit: seconds. This specifies the interval between consecutive results.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The number of data entries to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
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
	// > The value must be on the exact minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1719211800
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainTopFingerprintRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainTopFingerprintRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainTopFingerprintRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainTopFingerprintRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDomainTopFingerprintRequest) GetInterval() *int64 {
	return s.Interval
}

func (s *DescribeDomainTopFingerprintRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *DescribeDomainTopFingerprintRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *DescribeDomainTopFingerprintRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeDomainTopFingerprintRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDomainTopFingerprintRequest) SetDomain(v string) *DescribeDomainTopFingerprintRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainTopFingerprintRequest) SetEndTime(v int64) *DescribeDomainTopFingerprintRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainTopFingerprintRequest) SetInterval(v int64) *DescribeDomainTopFingerprintRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainTopFingerprintRequest) SetLimit(v int64) *DescribeDomainTopFingerprintRequest {
	s.Limit = &v
	return s
}

func (s *DescribeDomainTopFingerprintRequest) SetQueryType(v string) *DescribeDomainTopFingerprintRequest {
	s.QueryType = &v
	return s
}

func (s *DescribeDomainTopFingerprintRequest) SetRegion(v string) *DescribeDomainTopFingerprintRequest {
	s.Region = &v
	return s
}

func (s *DescribeDomainTopFingerprintRequest) SetStartTime(v int64) *DescribeDomainTopFingerprintRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainTopFingerprintRequest) Validate() error {
	return dara.Validate(s)
}
