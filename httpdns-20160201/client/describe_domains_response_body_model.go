// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomains(v *DescribeDomainsResponseBodyDomains) *DescribeDomainsResponseBody
	GetDomains() *DescribeDomainsResponseBodyDomains
	SetPageNumber(v int64) *DescribeDomainsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeDomainsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeDomainsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeDomainsResponseBody
	GetTotalCount() *int64
}

type DescribeDomainsResponseBody struct {
	Domains *DescribeDomainsResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// A6B3BB61-69CB-50E0-9DC0-0C1658D44A47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 12
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBody) GetDomains() *DescribeDomainsResponseBodyDomains {
	return s.Domains
}

func (s *DescribeDomainsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDomainsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDomainsResponseBody) SetDomains(v *DescribeDomainsResponseBodyDomains) *DescribeDomainsResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeDomainsResponseBody) SetPageNumber(v int64) *DescribeDomainsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainsResponseBody) SetPageSize(v int64) *DescribeDomainsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainsResponseBody) SetRequestId(v string) *DescribeDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainsResponseBody) SetTotalCount(v int64) *DescribeDomainsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDomainsResponseBody) Validate() error {
	if s.Domains != nil {
		if err := s.Domains.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDomainsResponseBodyDomains struct {
	Domain []*DescribeDomainsResponseBodyDomainsDomain `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Repeated"`
}

func (s DescribeDomainsResponseBodyDomains) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomains) GetDomain() []*DescribeDomainsResponseBodyDomainsDomain {
	return s.Domain
}

func (s *DescribeDomainsResponseBodyDomains) SetDomain(v []*DescribeDomainsResponseBodyDomainsDomain) *DescribeDomainsResponseBodyDomains {
	s.Domain = v
	return s
}

func (s *DescribeDomainsResponseBodyDomains) Validate() error {
	if s.Domain != nil {
		for _, item := range s.Domain {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainsResponseBodyDomainsDomain struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeDomainsResponseBodyDomainsDomain) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsResponseBodyDomainsDomain) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponseBodyDomainsDomain) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainsResponseBodyDomainsDomain) SetDomainName(v string) *DescribeDomainsResponseBodyDomainsDomain {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainsResponseBodyDomainsDomain) Validate() error {
	return dara.Validate(s)
}
