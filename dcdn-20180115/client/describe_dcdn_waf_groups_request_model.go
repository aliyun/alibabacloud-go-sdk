// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *DescribeDcdnWafGroupsRequest
	GetLanguage() *string
	SetPageNumber(v int32) *DescribeDcdnWafGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDcdnWafGroupsRequest
	GetPageSize() *int32
	SetQueryArgs(v string) *DescribeDcdnWafGroupsRequest
	GetQueryArgs() *string
}

type DescribeDcdnWafGroupsRequest struct {
	// The language of the response. Valid values:
	//
	// 	- **en*	- (default): English.
	//
	// 	- **zh**: Chinese.
	//
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The query conditions. The value is a string in the JSON format. Format: `QueryArgs={"PolicyIds":"IDs of protection policies","RuleIds":"IDs of the protection rules","RuleNameLike":"Names of the protection rule","DomainNames":"Protected domain names","DefenseScenes":"waf_group","RuleStatus":"on","OrderBy":"GmtModified","Desc":"false"}`
	//
	// > If you do not specify this parameter, all protection rules are queried.
	//
	// example:
	//
	// {"RuleIds":"100001,200002"}
	QueryArgs *string `json:"QueryArgs,omitempty" xml:"QueryArgs,omitempty"`
}

func (s DescribeDcdnWafGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafGroupsRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeDcdnWafGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDcdnWafGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDcdnWafGroupsRequest) GetQueryArgs() *string {
	return s.QueryArgs
}

func (s *DescribeDcdnWafGroupsRequest) SetLanguage(v string) *DescribeDcdnWafGroupsRequest {
	s.Language = &v
	return s
}

func (s *DescribeDcdnWafGroupsRequest) SetPageNumber(v int32) *DescribeDcdnWafGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnWafGroupsRequest) SetPageSize(v int32) *DescribeDcdnWafGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnWafGroupsRequest) SetQueryArgs(v string) *DescribeDcdnWafGroupsRequest {
	s.QueryArgs = &v
	return s
}

func (s *DescribeDcdnWafGroupsRequest) Validate() error {
	return dara.Validate(s)
}
