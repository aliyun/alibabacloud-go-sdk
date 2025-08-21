// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableWebCCRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DisableWebCCRequest
	GetDomain() *string
	SetResourceGroupId(v string) *DisableWebCCRequest
	GetResourceGroupId() *string
}

type DisableWebCCRequest struct {
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

func (s DisableWebCCRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableWebCCRequest) GoString() string {
	return s.String()
}

func (s *DisableWebCCRequest) GetDomain() *string {
	return s.Domain
}

func (s *DisableWebCCRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DisableWebCCRequest) SetDomain(v string) *DisableWebCCRequest {
	s.Domain = &v
	return s
}

func (s *DisableWebCCRequest) SetResourceGroupId(v string) *DisableWebCCRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DisableWebCCRequest) Validate() error {
	return dara.Validate(s)
}
