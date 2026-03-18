// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckGrantRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsSlr(v bool) *CheckGrantRequest
	GetIsSlr() *bool
	SetRegionId(v string) *CheckGrantRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CheckGrantRequest
	GetResourceGroupId() *string
}

type CheckGrantRequest struct {
	// Specifies whether to allow Anti-DDoS Origin to check the service-linked role. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsSlr *bool `json:"IsSlr,omitempty" xml:"IsSlr,omitempty"`
	// The ID of the region where the Anti-DDoS Origin instance resides.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/118703.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Origin instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CheckGrantRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckGrantRequest) GoString() string {
	return s.String()
}

func (s *CheckGrantRequest) GetIsSlr() *bool {
	return s.IsSlr
}

func (s *CheckGrantRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckGrantRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CheckGrantRequest) SetIsSlr(v bool) *CheckGrantRequest {
	s.IsSlr = &v
	return s
}

func (s *CheckGrantRequest) SetRegionId(v string) *CheckGrantRequest {
	s.RegionId = &v
	return s
}

func (s *CheckGrantRequest) SetResourceGroupId(v string) *CheckGrantRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CheckGrantRequest) Validate() error {
	return dara.Validate(s)
}
