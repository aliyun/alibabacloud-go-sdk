// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebPreciseAccessRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DeleteWebPreciseAccessRuleRequest
	GetDomain() *string
	SetResourceGroupId(v string) *DeleteWebPreciseAccessRuleRequest
	GetResourceGroupId() *string
	SetRuleNames(v []*string) *DeleteWebPreciseAccessRuleRequest
	GetRuleNames() []*string
}

type DeleteWebPreciseAccessRuleRequest struct {
	// The domain name of the website.
	//
	// > A forwarding rule must be configured for the domain name. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all domain names.
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
	// An array that consists of the names of rules to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// testrule
	RuleNames []*string `json:"RuleNames,omitempty" xml:"RuleNames,omitempty" type:"Repeated"`
}

func (s DeleteWebPreciseAccessRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebPreciseAccessRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteWebPreciseAccessRuleRequest) GetDomain() *string {
	return s.Domain
}

func (s *DeleteWebPreciseAccessRuleRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteWebPreciseAccessRuleRequest) GetRuleNames() []*string {
	return s.RuleNames
}

func (s *DeleteWebPreciseAccessRuleRequest) SetDomain(v string) *DeleteWebPreciseAccessRuleRequest {
	s.Domain = &v
	return s
}

func (s *DeleteWebPreciseAccessRuleRequest) SetResourceGroupId(v string) *DeleteWebPreciseAccessRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteWebPreciseAccessRuleRequest) SetRuleNames(v []*string) *DeleteWebPreciseAccessRuleRequest {
	s.RuleNames = v
	return s
}

func (s *DeleteWebPreciseAccessRuleRequest) Validate() error {
	return dara.Validate(s)
}
