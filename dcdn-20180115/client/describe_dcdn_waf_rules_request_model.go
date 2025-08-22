// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeDcdnWafRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDcdnWafRulesRequest
	GetPageSize() *int32
	SetQueryArgs(v string) *DescribeDcdnWafRulesRequest
	GetQueryArgs() *string
}

type DescribeDcdnWafRulesRequest struct {
	// The number of the page to return. Valid values: **1*	- to **100000**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of protection rules to return per page. Valid values: integers from **1*	- to **500**. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The query conditions. The value needs to be a JSON string in the following format: `QueryArgs={"PolicyIds":"The range of protection policy IDs","RuleIds":"The range of protection rule IDs","RuleNameLike":"The name of the protection rule","DomainNames":"The protected domain names","DefenseScenes":"waf_group","RuleStatus":"on","OrderBy":"GmtModified","Desc":"false"}`.
	//
	// > If you do not specify this parameter, all protection rules are queried.
	//
	// example:
	//
	// {"RuleIds":"100001,200002"}
	QueryArgs *string `json:"QueryArgs,omitempty" xml:"QueryArgs,omitempty"`
}

func (s DescribeDcdnWafRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDcdnWafRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDcdnWafRulesRequest) GetQueryArgs() *string {
	return s.QueryArgs
}

func (s *DescribeDcdnWafRulesRequest) SetPageNumber(v int32) *DescribeDcdnWafRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnWafRulesRequest) SetPageSize(v int32) *DescribeDcdnWafRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnWafRulesRequest) SetQueryArgs(v string) *DescribeDcdnWafRulesRequest {
	s.QueryArgs = &v
	return s
}

func (s *DescribeDcdnWafRulesRequest) Validate() error {
	return dara.Validate(s)
}
