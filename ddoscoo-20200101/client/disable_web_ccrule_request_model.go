// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableWebCCRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DisableWebCCRuleRequest
	GetDomain() *string
	SetResourceGroupId(v string) *DisableWebCCRuleRequest
	GetResourceGroupId() *string
}

type DisableWebCCRuleRequest struct {
	// The domain name of the website.
	//
	// > A forwarding rule must be configured for a domain name. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all domain names.
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
}

func (s DisableWebCCRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableWebCCRuleRequest) GoString() string {
	return s.String()
}

func (s *DisableWebCCRuleRequest) GetDomain() *string {
	return s.Domain
}

func (s *DisableWebCCRuleRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DisableWebCCRuleRequest) SetDomain(v string) *DisableWebCCRuleRequest {
	s.Domain = &v
	return s
}

func (s *DisableWebCCRuleRequest) SetResourceGroupId(v string) *DisableWebCCRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DisableWebCCRuleRequest) Validate() error {
	return dara.Validate(s)
}
