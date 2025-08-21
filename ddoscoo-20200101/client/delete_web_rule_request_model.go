// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DeleteWebRuleRequest
	GetDomain() *string
	SetResourceGroupId(v string) *DeleteWebRuleRequest
	GetResourceGroupId() *string
}

type DeleteWebRuleRequest struct {
	// The domain name of the website from which you want to delete the forwarding rule.
	//
	// > A forwarding rule must be configured for the domain name. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query the domain names for which forwarding rules are configured.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// For more information about resource groups, see [Create a resource group](https://help.aliyun.com/document_detail/94485.html).
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DeleteWebRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteWebRuleRequest) GetDomain() *string {
	return s.Domain
}

func (s *DeleteWebRuleRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteWebRuleRequest) SetDomain(v string) *DeleteWebRuleRequest {
	s.Domain = &v
	return s
}

func (s *DeleteWebRuleRequest) SetResourceGroupId(v string) *DeleteWebRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteWebRuleRequest) Validate() error {
	return dara.Validate(s)
}
