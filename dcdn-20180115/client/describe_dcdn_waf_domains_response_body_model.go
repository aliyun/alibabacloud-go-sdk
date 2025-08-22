// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomains(v []*DescribeDcdnWafDomainsResponseBodyDomains) *DescribeDcdnWafDomainsResponseBody
	GetDomains() []*DescribeDcdnWafDomainsResponseBodyDomains
	SetPageNumber(v int32) *DescribeDcdnWafDomainsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDcdnWafDomainsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDcdnWafDomainsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDcdnWafDomainsResponseBody
	GetTotalCount() *int32
}

type DescribeDcdnWafDomainsResponseBody struct {
	// The protected domain name.
	Domains []*DescribeDcdnWafDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
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
	// 153ca2cd-3c01-44be-b408-64dbc6c88630
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of protected domain names.
	//
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDcdnWafDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafDomainsResponseBody) GetDomains() []*DescribeDcdnWafDomainsResponseBodyDomains {
	return s.Domains
}

func (s *DescribeDcdnWafDomainsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDcdnWafDomainsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDcdnWafDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnWafDomainsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDcdnWafDomainsResponseBody) SetDomains(v []*DescribeDcdnWafDomainsResponseBodyDomains) *DescribeDcdnWafDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeDcdnWafDomainsResponseBody) SetPageNumber(v int32) *DescribeDcdnWafDomainsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnWafDomainsResponseBody) SetPageSize(v int32) *DescribeDcdnWafDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnWafDomainsResponseBody) SetRequestId(v string) *DescribeDcdnWafDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnWafDomainsResponseBody) SetTotalCount(v int32) *DescribeDcdnWafDomainsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDcdnWafDomainsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnWafDomainsResponseBodyDomains struct {
	// The header of IP address of the client that is connected to the point of presence (POP).
	//
	// example:
	//
	// X-Forwarded-For
	ClientIpTag *string `json:"ClientIpTag,omitempty" xml:"ClientIpTag,omitempty"`
	// The protected domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The number of protection policies that were configured for the protected domain name.
	//
	// example:
	//
	// 3
	PolicyCount *int32 `json:"PolicyCount,omitempty" xml:"PolicyCount,omitempty"`
}

func (s DescribeDcdnWafDomainsResponseBodyDomains) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafDomainsResponseBodyDomains) GetClientIpTag() *string {
	return s.ClientIpTag
}

func (s *DescribeDcdnWafDomainsResponseBodyDomains) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnWafDomainsResponseBodyDomains) GetPolicyCount() *int32 {
	return s.PolicyCount
}

func (s *DescribeDcdnWafDomainsResponseBodyDomains) SetClientIpTag(v string) *DescribeDcdnWafDomainsResponseBodyDomains {
	s.ClientIpTag = &v
	return s
}

func (s *DescribeDcdnWafDomainsResponseBodyDomains) SetDomainName(v string) *DescribeDcdnWafDomainsResponseBodyDomains {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnWafDomainsResponseBodyDomains) SetPolicyCount(v int32) *DescribeDcdnWafDomainsResponseBodyDomains {
	s.PolicyCount = &v
	return s
}

func (s *DescribeDcdnWafDomainsResponseBodyDomains) Validate() error {
	return dara.Validate(s)
}
