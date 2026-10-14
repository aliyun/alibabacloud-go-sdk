// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeDomainResourceRequest
	GetDomain() *string
	SetInstanceIds(v []*string) *DescribeDomainResourceRequest
	GetInstanceIds() []*string
	SetPageNumber(v int32) *DescribeDomainResourceRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDomainResourceRequest
	GetPageSize() *int32
	SetQueryDomainPattern(v string) *DescribeDomainResourceRequest
	GetQueryDomainPattern() *string
}

type DescribeDomainResourceRequest struct {
	// The domain name to query.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The IDs of the Anti-DDoS Pro instances to query.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The number of the page to return. The default value is **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of forwarding rules per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The matching pattern for the query. Valid values:
	//
	// - **fuzzy*	- (default): Specifies a fuzzy search.
	//
	// - **exact**: Specifies an exact match.
	//
	// example:
	//
	// fuzzy
	QueryDomainPattern *string `json:"QueryDomainPattern,omitempty" xml:"QueryDomainPattern,omitempty"`
}

func (s DescribeDomainResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainResourceRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainResourceRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeDomainResourceRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDomainResourceRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDomainResourceRequest) GetQueryDomainPattern() *string {
	return s.QueryDomainPattern
}

func (s *DescribeDomainResourceRequest) SetDomain(v string) *DescribeDomainResourceRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainResourceRequest) SetInstanceIds(v []*string) *DescribeDomainResourceRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeDomainResourceRequest) SetPageNumber(v int32) *DescribeDomainResourceRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainResourceRequest) SetPageSize(v int32) *DescribeDomainResourceRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainResourceRequest) SetQueryDomainPattern(v string) *DescribeDomainResourceRequest {
	s.QueryDomainPattern = &v
	return s
}

func (s *DescribeDomainResourceRequest) Validate() error {
	return dara.Validate(s)
}
