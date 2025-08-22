// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeDcdnWafPoliciesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDcdnWafPoliciesRequest
	GetPageSize() *int32
	SetQueryArgs(v string) *DescribeDcdnWafPoliciesRequest
	GetQueryArgs() *string
}

type DescribeDcdnWafPoliciesRequest struct {
	// The number of the page to return. Valid values: **1*	- to **100000**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of protection policies to return on each page. Valid values: an integer from **1*	- to **500**. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The query conditions. The value is a JSON string. The format is `QueryArgs={"PolicyIds":"The IDs of protection policies","RuleIds":"The IDs of protection rules","PolicyNameLike":"The name of the protection policy","DomainNames":"The protected domain names","PolicyType":"default","DefenseScenes":"waf_group","PolicyStatus":"on","OrderBy":"GmtModified","Desc":"false"}`
	//
	// > If you do not set this parameter, all protection policies are queried.
	//
	// example:
	//
	// {"PolicyNameLIike":"test_policy"}
	QueryArgs *string `json:"QueryArgs,omitempty" xml:"QueryArgs,omitempty"`
}

func (s DescribeDcdnWafPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafPoliciesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafPoliciesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDcdnWafPoliciesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDcdnWafPoliciesRequest) GetQueryArgs() *string {
	return s.QueryArgs
}

func (s *DescribeDcdnWafPoliciesRequest) SetPageNumber(v int32) *DescribeDcdnWafPoliciesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnWafPoliciesRequest) SetPageSize(v int32) *DescribeDcdnWafPoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnWafPoliciesRequest) SetQueryArgs(v string) *DescribeDcdnWafPoliciesRequest {
	s.QueryArgs = &v
	return s
}

func (s *DescribeDcdnWafPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
