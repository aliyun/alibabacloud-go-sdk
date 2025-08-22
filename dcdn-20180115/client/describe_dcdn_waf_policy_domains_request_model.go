// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafPolicyDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeDcdnWafPolicyDomainsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDcdnWafPolicyDomainsRequest
	GetPageSize() *int32
	SetPolicyId(v int64) *DescribeDcdnWafPolicyDomainsRequest
	GetPolicyId() *int64
}

type DescribeDcdnWafPolicyDomainsRequest struct {
	// The number of the page to return. Valid values: **1*	- to **100000**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of domain names to return on each page. Valid values: an integer from **1*	- to **500**. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the protection policy. You can specify only one ID in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100001
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s DescribeDcdnWafPolicyDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafPolicyDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafPolicyDomainsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDcdnWafPolicyDomainsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDcdnWafPolicyDomainsRequest) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *DescribeDcdnWafPolicyDomainsRequest) SetPageNumber(v int32) *DescribeDcdnWafPolicyDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnWafPolicyDomainsRequest) SetPageSize(v int32) *DescribeDcdnWafPolicyDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnWafPolicyDomainsRequest) SetPolicyId(v int64) *DescribeDcdnWafPolicyDomainsRequest {
	s.PolicyId = &v
	return s
}

func (s *DescribeDcdnWafPolicyDomainsRequest) Validate() error {
	return dara.Validate(s)
}
