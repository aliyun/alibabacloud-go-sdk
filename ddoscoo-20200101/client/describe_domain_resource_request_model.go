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
	// The domain name of the website that you want to query.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// An array that consists of the IDs of instances to query.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The match mode. Valid values:
	//
	// 	- **fuzzy**: fuzzy match. This is the default value.
	//
	// 	- **exact**: exact match.
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
