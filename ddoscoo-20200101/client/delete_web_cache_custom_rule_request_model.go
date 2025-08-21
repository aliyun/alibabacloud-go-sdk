// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebCacheCustomRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DeleteWebCacheCustomRuleRequest
	GetDomain() *string
	SetResourceGroupId(v string) *DeleteWebCacheCustomRuleRequest
	GetResourceGroupId() *string
	SetRuleNames(v []*string) *DeleteWebCacheCustomRuleRequest
	GetRuleNames() []*string
}

type DeleteWebCacheCustomRuleRequest struct {
	// The domain name for which you want to delete the custom rules of the Static Page Caching policy.
	//
	// > You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all the domain names that are added to Anti-DDoS Pro or Anti-DDoS Premium.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the resource group to which the instance belongs in Resource Management.
	//
	// If you do not configure this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// An array consisting of the names of the rules that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	RuleNames []*string `json:"RuleNames,omitempty" xml:"RuleNames,omitempty" type:"Repeated"`
}

func (s DeleteWebCacheCustomRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebCacheCustomRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteWebCacheCustomRuleRequest) GetDomain() *string {
	return s.Domain
}

func (s *DeleteWebCacheCustomRuleRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteWebCacheCustomRuleRequest) GetRuleNames() []*string {
	return s.RuleNames
}

func (s *DeleteWebCacheCustomRuleRequest) SetDomain(v string) *DeleteWebCacheCustomRuleRequest {
	s.Domain = &v
	return s
}

func (s *DeleteWebCacheCustomRuleRequest) SetResourceGroupId(v string) *DeleteWebCacheCustomRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteWebCacheCustomRuleRequest) SetRuleNames(v []*string) *DeleteWebCacheCustomRuleRequest {
	s.RuleNames = v
	return s
}

func (s *DeleteWebCacheCustomRuleRequest) Validate() error {
	return dara.Validate(s)
}
