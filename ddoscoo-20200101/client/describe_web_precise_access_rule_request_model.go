// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebPreciseAccessRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomains(v []*string) *DescribeWebPreciseAccessRuleRequest
	GetDomains() []*string
	SetOwner(v string) *DescribeWebPreciseAccessRuleRequest
	GetOwner() *string
	SetResourceGroupId(v string) *DescribeWebPreciseAccessRuleRequest
	GetResourceGroupId() *string
}

type DescribeWebPreciseAccessRuleRequest struct {
	// The domain name of the website.
	//
	// > A forwarding rule must be configured for the domain name. You can call [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) to query all domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// The source of the rule. Valid values:
	//
	// - **manual*	- (default): The rule is manually created.
	//
	// - **clover**: The rule is automatically generated.
	//
	// example:
	//
	// manual
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Pro instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeWebPreciseAccessRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebPreciseAccessRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebPreciseAccessRuleRequest) GetDomains() []*string {
	return s.Domains
}

func (s *DescribeWebPreciseAccessRuleRequest) GetOwner() *string {
	return s.Owner
}

func (s *DescribeWebPreciseAccessRuleRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeWebPreciseAccessRuleRequest) SetDomains(v []*string) *DescribeWebPreciseAccessRuleRequest {
	s.Domains = v
	return s
}

func (s *DescribeWebPreciseAccessRuleRequest) SetOwner(v string) *DescribeWebPreciseAccessRuleRequest {
	s.Owner = &v
	return s
}

func (s *DescribeWebPreciseAccessRuleRequest) SetResourceGroupId(v string) *DescribeWebPreciseAccessRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeWebPreciseAccessRuleRequest) Validate() error {
	return dara.Validate(s)
}
