// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainViewTopUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeDomainViewTopUrlRequest
	GetDomain() *string
	SetEndTime(v int64) *DescribeDomainViewTopUrlRequest
	GetEndTime() *int64
	SetInerval(v int64) *DescribeDomainViewTopUrlRequest
	GetInerval() *int64
	SetResourceGroupId(v string) *DescribeDomainViewTopUrlRequest
	GetResourceGroupId() *string
	SetStartTime(v int64) *DescribeDomainViewTopUrlRequest
	GetStartTime() *int64
	SetTop(v int32) *DescribeDomainViewTopUrlRequest
	GetTop() *int32
}

type DescribeDomainViewTopUrlRequest struct {
	// The domain name of the website service.
	//
	// > The domain name must have a website service forwarding rule configured. You can invoke [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) to query all domain names that are added to Anti-DDoS Pro and Anti-DDoS Premium.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The end of the time range to query. Specify a UNIX timestamp. Unit: seconds.
	//
	// > The value must be on the exact minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1583683200
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The step size for returning data, in seconds. This parameter specifies the interval at which a result is returned. The supported step size varies based on the query time range (determined by **StartTime*	- and **EndTime**). Recommended settings:
	//
	// - Query time range of up to 1 hour: minimum step size of 60 seconds, maximum not exceeding the query time range.
	//
	// - Query time range of more than 1 hour but up to 6 hours: minimum step size of 600 seconds, maximum not exceeding the query time range.
	//
	// - Query time range of more than 6 hours but up to 24 hours: minimum step size of 1800 seconds, maximum not exceeding the query time range.
	//
	// - Query time range of more than 24 hours but up to 7 days: minimum step size of 3600 seconds, maximum not exceeding the query time range.
	//
	// - Query time range of more than 7 days but up to 15 days: minimum step size of 14400 seconds, maximum not exceeding the query time range.
	//
	// - Query time range of more than 15 days: minimum step size of 43200 seconds, maximum not exceeding the query time range.
	//
	// example:
	//
	// 60
	Inerval *int64 `json:"Inerval,omitempty" xml:"Inerval,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Pro or Anti-DDoS Premium instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The beginning of the time range to query. Specify a UNIX timestamp. Unit: seconds.
	//
	// > The value must be on the exact minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1582992000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The number of URLs to return. Valid values: **1*	- to **100**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	Top *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
}

func (s DescribeDomainViewTopUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainViewTopUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainViewTopUrlRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainViewTopUrlRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDomainViewTopUrlRequest) GetInerval() *int64 {
	return s.Inerval
}

func (s *DescribeDomainViewTopUrlRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDomainViewTopUrlRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDomainViewTopUrlRequest) GetTop() *int32 {
	return s.Top
}

func (s *DescribeDomainViewTopUrlRequest) SetDomain(v string) *DescribeDomainViewTopUrlRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainViewTopUrlRequest) SetEndTime(v int64) *DescribeDomainViewTopUrlRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainViewTopUrlRequest) SetInerval(v int64) *DescribeDomainViewTopUrlRequest {
	s.Inerval = &v
	return s
}

func (s *DescribeDomainViewTopUrlRequest) SetResourceGroupId(v string) *DescribeDomainViewTopUrlRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainViewTopUrlRequest) SetStartTime(v int64) *DescribeDomainViewTopUrlRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainViewTopUrlRequest) SetTop(v int32) *DescribeDomainViewTopUrlRequest {
	s.Top = &v
	return s
}

func (s *DescribeDomainViewTopUrlRequest) Validate() error {
	return dara.Validate(s)
}
