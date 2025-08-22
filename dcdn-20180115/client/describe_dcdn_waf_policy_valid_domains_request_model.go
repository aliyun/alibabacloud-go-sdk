// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafPolicyValidDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefenseScene(v string) *DescribeDcdnWafPolicyValidDomainsRequest
	GetDefenseScene() *string
	SetDomainNameLike(v string) *DescribeDcdnWafPolicyValidDomainsRequest
	GetDomainNameLike() *string
	SetPageNumber(v int32) *DescribeDcdnWafPolicyValidDomainsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDcdnWafPolicyValidDomainsRequest
	GetPageSize() *int32
}

type DescribeDcdnWafPolicyValidDomainsRequest struct {
	// The type of the Web Application Firewall (WAF) protection policy. Valid values:
	//
	// 	- waf_group: basic web protection
	//
	// 	- custom_acl: custom protection
	//
	// 	- whitelist: IP address whitelist
	//
	// 	- ip_blacklist: IP address blacklist
	//
	// 	- region_block: region blacklist
	//
	// 	- bot: bot management
	//
	// This parameter is required.
	//
	// example:
	//
	// custom_acl
	DefenseScene *string `json:"DefenseScene,omitempty" xml:"DefenseScene,omitempty"`
	// The protected domain name. Fuzzy search is supported.
	//
	// example:
	//
	// example.com
	DomainNameLike *string `json:"DomainNameLike,omitempty" xml:"DomainNameLike,omitempty"`
	// The page number of the returned page. Valid values: **1*	- to **100000**. Default value: **1**.
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
}

func (s DescribeDcdnWafPolicyValidDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafPolicyValidDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafPolicyValidDomainsRequest) GetDefenseScene() *string {
	return s.DefenseScene
}

func (s *DescribeDcdnWafPolicyValidDomainsRequest) GetDomainNameLike() *string {
	return s.DomainNameLike
}

func (s *DescribeDcdnWafPolicyValidDomainsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDcdnWafPolicyValidDomainsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDcdnWafPolicyValidDomainsRequest) SetDefenseScene(v string) *DescribeDcdnWafPolicyValidDomainsRequest {
	s.DefenseScene = &v
	return s
}

func (s *DescribeDcdnWafPolicyValidDomainsRequest) SetDomainNameLike(v string) *DescribeDcdnWafPolicyValidDomainsRequest {
	s.DomainNameLike = &v
	return s
}

func (s *DescribeDcdnWafPolicyValidDomainsRequest) SetPageNumber(v int32) *DescribeDcdnWafPolicyValidDomainsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnWafPolicyValidDomainsRequest) SetPageSize(v int32) *DescribeDcdnWafPolicyValidDomainsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnWafPolicyValidDomainsRequest) Validate() error {
	return dara.Validate(s)
}
