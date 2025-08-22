// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafPolicyDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomains(v []*DescribeDcdnWafPolicyDomainsResponseBodyDomains) *DescribeDcdnWafPolicyDomainsResponseBody
	GetDomains() []*DescribeDcdnWafPolicyDomainsResponseBodyDomains
	SetPageNumber(v int32) *DescribeDcdnWafPolicyDomainsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDcdnWafPolicyDomainsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDcdnWafPolicyDomainsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDcdnWafPolicyDomainsResponseBody
	GetTotalCount() *int32
}

type DescribeDcdnWafPolicyDomainsResponseBody struct {
	// The accelerated domain names.
	Domains []*DescribeDcdnWafPolicyDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// The page number of the returned page, which is the same as the PageNumber parameter in request parameters.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of domain names returned per page, which is the same as the PageSize parameter in request parameters.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 153ca2cd-3c01-44be-b480-64dbc6c88630
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of domain names returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDcdnWafPolicyDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafPolicyDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafPolicyDomainsResponseBody) GetDomains() []*DescribeDcdnWafPolicyDomainsResponseBodyDomains {
	return s.Domains
}

func (s *DescribeDcdnWafPolicyDomainsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDcdnWafPolicyDomainsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDcdnWafPolicyDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnWafPolicyDomainsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDcdnWafPolicyDomainsResponseBody) SetDomains(v []*DescribeDcdnWafPolicyDomainsResponseBodyDomains) *DescribeDcdnWafPolicyDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeDcdnWafPolicyDomainsResponseBody) SetPageNumber(v int32) *DescribeDcdnWafPolicyDomainsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnWafPolicyDomainsResponseBody) SetPageSize(v int32) *DescribeDcdnWafPolicyDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnWafPolicyDomainsResponseBody) SetRequestId(v string) *DescribeDcdnWafPolicyDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnWafPolicyDomainsResponseBody) SetTotalCount(v int32) *DescribeDcdnWafPolicyDomainsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDcdnWafPolicyDomainsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnWafPolicyDomainsResponseBodyDomains struct {
	// The accelerated domain name that is protected by the specified protection policy.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeDcdnWafPolicyDomainsResponseBodyDomains) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafPolicyDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafPolicyDomainsResponseBodyDomains) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnWafPolicyDomainsResponseBodyDomains) SetDomainName(v string) *DescribeDcdnWafPolicyDomainsResponseBodyDomains {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnWafPolicyDomainsResponseBodyDomains) Validate() error {
	return dara.Validate(s)
}
