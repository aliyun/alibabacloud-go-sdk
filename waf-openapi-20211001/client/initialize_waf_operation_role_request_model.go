// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeWafOperationRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *InitializeWafOperationRoleRequest
	GetInstanceId() *string
	SetRegionId(v string) *InitializeWafOperationRoleRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *InitializeWafOperationRoleRequest
	GetResourceManagerResourceGroupId() *string
}

type InitializeWafOperationRoleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s InitializeWafOperationRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s InitializeWafOperationRoleRequest) GoString() string {
	return s.String()
}

func (s *InitializeWafOperationRoleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *InitializeWafOperationRoleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *InitializeWafOperationRoleRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *InitializeWafOperationRoleRequest) SetInstanceId(v string) *InitializeWafOperationRoleRequest {
	s.InstanceId = &v
	return s
}

func (s *InitializeWafOperationRoleRequest) SetRegionId(v string) *InitializeWafOperationRoleRequest {
	s.RegionId = &v
	return s
}

func (s *InitializeWafOperationRoleRequest) SetResourceManagerResourceGroupId(v string) *InitializeWafOperationRoleRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *InitializeWafOperationRoleRequest) Validate() error {
	return dara.Validate(s)
}
