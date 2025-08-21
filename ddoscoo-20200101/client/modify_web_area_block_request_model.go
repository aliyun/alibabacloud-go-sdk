// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebAreaBlockRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *ModifyWebAreaBlockRequest
	GetDomain() *string
	SetRegions(v []*string) *ModifyWebAreaBlockRequest
	GetRegions() []*string
	SetResourceGroupId(v string) *ModifyWebAreaBlockRequest
	GetResourceGroupId() *string
}

type ModifyWebAreaBlockRequest struct {
	// The domain name whose configurations you want to modify.
	//
	// > A forwarding rule must be configured for the domain name. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The regions from which you block requests.
	//
	// > If you do not configure this parameter, the Blocked Regions (Domain Names) policy is disabled.
	//
	// example:
	//
	// CN-SHANGHAI
	Regions []*string `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// For more information about resource groups, see [Create a resource group](https://help.aliyun.com/document_detail/94485.html).
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ModifyWebAreaBlockRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebAreaBlockRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebAreaBlockRequest) GetDomain() *string {
	return s.Domain
}

func (s *ModifyWebAreaBlockRequest) GetRegions() []*string {
	return s.Regions
}

func (s *ModifyWebAreaBlockRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyWebAreaBlockRequest) SetDomain(v string) *ModifyWebAreaBlockRequest {
	s.Domain = &v
	return s
}

func (s *ModifyWebAreaBlockRequest) SetRegions(v []*string) *ModifyWebAreaBlockRequest {
	s.Regions = v
	return s
}

func (s *ModifyWebAreaBlockRequest) SetResourceGroupId(v string) *ModifyWebAreaBlockRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyWebAreaBlockRequest) Validate() error {
	return dara.Validate(s)
}
