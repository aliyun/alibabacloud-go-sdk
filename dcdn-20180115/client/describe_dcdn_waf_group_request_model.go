// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DescribeDcdnWafGroupRequest
	GetId() *int64
	SetLanguage(v string) *DescribeDcdnWafGroupRequest
	GetLanguage() *string
	SetPageNumber(v int32) *DescribeDcdnWafGroupRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDcdnWafGroupRequest
	GetPageSize() *int32
	SetQueryArgs(v string) *DescribeDcdnWafGroupRequest
	GetQueryArgs() *string
	SetScope(v string) *DescribeDcdnWafGroupRequest
	GetScope() *string
}

type DescribeDcdnWafGroupRequest struct {
	// The ID of the WAF rule group. You can query the ID by calling the [DescribeDcdnWafGroups](~~DescribeDcdnWafGroups~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1012
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the response. Valid values:
	//
	// 	- **en**: English.
	//
	// 	- **zh**: Chinese.
	//
	// This parameter is required.
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
	// The query conditions. The value is a JSON string in the following format:
	//
	// `QueryArgs={"PolicyIds":"The range of protection policy IDs","RuleIds":"The range of protection rule IDs","RuleNameLike":"The name of the protection rule","DomainNames":"The protected domain names","DefenseScenes":"waf_group","RuleStatus":"on","OrderBy":"GmtModified","Desc":"false"}`
	//
	// >  If you do not specify this parameter, all protection rules are queried.
	//
	// example:
	//
	// {\\"RiskLevel\\":\\"\\",\\"ProtectionType\\":\\"\\",\\"ApplicationType\\":\\"\\",\\"RuleIdLike\\":\\"\\"}
	QueryArgs *string `json:"QueryArgs,omitempty" xml:"QueryArgs,omitempty"`
	// The range of the rule group to be queried.
	//
	// 	- **in**: Rules in the rule group are returned.
	//
	// 	- **out**: Rules that are in the full rule set but are not in the rule group are returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// in
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s DescribeDcdnWafGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafGroupRequest) GetId() *int64 {
	return s.Id
}

func (s *DescribeDcdnWafGroupRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeDcdnWafGroupRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDcdnWafGroupRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDcdnWafGroupRequest) GetQueryArgs() *string {
	return s.QueryArgs
}

func (s *DescribeDcdnWafGroupRequest) GetScope() *string {
	return s.Scope
}

func (s *DescribeDcdnWafGroupRequest) SetId(v int64) *DescribeDcdnWafGroupRequest {
	s.Id = &v
	return s
}

func (s *DescribeDcdnWafGroupRequest) SetLanguage(v string) *DescribeDcdnWafGroupRequest {
	s.Language = &v
	return s
}

func (s *DescribeDcdnWafGroupRequest) SetPageNumber(v int32) *DescribeDcdnWafGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnWafGroupRequest) SetPageSize(v int32) *DescribeDcdnWafGroupRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnWafGroupRequest) SetQueryArgs(v string) *DescribeDcdnWafGroupRequest {
	s.QueryArgs = &v
	return s
}

func (s *DescribeDcdnWafGroupRequest) SetScope(v string) *DescribeDcdnWafGroupRequest {
	s.Scope = &v
	return s
}

func (s *DescribeDcdnWafGroupRequest) Validate() error {
	return dara.Validate(s)
}
