// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebCCRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DeleteWebCCRuleRequest
	GetDomain() *string
	SetName(v string) *DeleteWebCCRuleRequest
	GetName() *string
	SetResourceGroupId(v string) *DeleteWebCCRuleRequest
	GetResourceGroupId() *string
}

type DeleteWebCCRuleRequest struct {
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
	// The name of the custom frequency control rule that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// wq
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DeleteWebCCRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebCCRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteWebCCRuleRequest) GetDomain() *string {
	return s.Domain
}

func (s *DeleteWebCCRuleRequest) GetName() *string {
	return s.Name
}

func (s *DeleteWebCCRuleRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteWebCCRuleRequest) SetDomain(v string) *DeleteWebCCRuleRequest {
	s.Domain = &v
	return s
}

func (s *DeleteWebCCRuleRequest) SetName(v string) *DeleteWebCCRuleRequest {
	s.Name = &v
	return s
}

func (s *DeleteWebCCRuleRequest) SetResourceGroupId(v string) *DeleteWebCCRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteWebCCRuleRequest) Validate() error {
	return dara.Validate(s)
}
