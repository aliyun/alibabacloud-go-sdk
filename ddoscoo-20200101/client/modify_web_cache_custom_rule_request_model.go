// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebCacheCustomRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *ModifyWebCacheCustomRuleRequest
	GetDomain() *string
	SetResourceGroupId(v string) *ModifyWebCacheCustomRuleRequest
	GetResourceGroupId() *string
	SetRules(v string) *ModifyWebCacheCustomRuleRequest
	GetRules() *string
}

type ModifyWebCacheCustomRuleRequest struct {
	// The domain name of the website.
	//
	// > A forwarding rule must be configured for the domain name, and the domain name must be associated with an instance that uses the Enhanced function plan. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The details of the custom rule. This parameter is a JSON string. The string contains the following fields:
	//
	// 	- **Name**: the name of the rule. This field is required and must be of the string type.
	//
	// 	- **Uri**: the path to the cached page. This field is required and must be of the STRING type.
	//
	// 	- **Mode**: the cache mode. This field is required and must be of the STRING type. Valid values:
	//
	//     	- **standard**: uses the standard mode.
	//
	//     	- **aggressive**: uses the enhanced mode.
	//
	//     	- **bypass**: No data is cached.
	//
	// 	- **CacheTtl**: the expiration time of the page cache. This field is required and must be of the INTEGER type. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"Name": "test","Uri": "/a","Mode": "standard","CacheTtl": 3600}]
	Rules *string `json:"Rules,omitempty" xml:"Rules,omitempty"`
}

func (s ModifyWebCacheCustomRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebCacheCustomRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebCacheCustomRuleRequest) GetDomain() *string {
	return s.Domain
}

func (s *ModifyWebCacheCustomRuleRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyWebCacheCustomRuleRequest) GetRules() *string {
	return s.Rules
}

func (s *ModifyWebCacheCustomRuleRequest) SetDomain(v string) *ModifyWebCacheCustomRuleRequest {
	s.Domain = &v
	return s
}

func (s *ModifyWebCacheCustomRuleRequest) SetResourceGroupId(v string) *ModifyWebCacheCustomRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyWebCacheCustomRuleRequest) SetRules(v string) *ModifyWebCacheCustomRuleRequest {
	s.Rules = &v
	return s
}

func (s *ModifyWebCacheCustomRuleRequest) Validate() error {
	return dara.Validate(s)
}
