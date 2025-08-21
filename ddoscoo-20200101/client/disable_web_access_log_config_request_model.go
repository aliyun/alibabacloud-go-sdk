// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableWebAccessLogConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DisableWebAccessLogConfigRequest
	GetDomain() *string
	SetResourceGroupId(v string) *DisableWebAccessLogConfigRequest
	GetResourceGroupId() *string
}

type DisableWebAccessLogConfigRequest struct {
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
}

func (s DisableWebAccessLogConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableWebAccessLogConfigRequest) GoString() string {
	return s.String()
}

func (s *DisableWebAccessLogConfigRequest) GetDomain() *string {
	return s.Domain
}

func (s *DisableWebAccessLogConfigRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DisableWebAccessLogConfigRequest) SetDomain(v string) *DisableWebAccessLogConfigRequest {
	s.Domain = &v
	return s
}

func (s *DisableWebAccessLogConfigRequest) SetResourceGroupId(v string) *DisableWebAccessLogConfigRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DisableWebAccessLogConfigRequest) Validate() error {
	return dara.Validate(s)
}
