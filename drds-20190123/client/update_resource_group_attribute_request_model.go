// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceGroupAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDrdsInstanceId(v string) *UpdateResourceGroupAttributeRequest
	GetDrdsInstanceId() *string
	SetNewResourceGroupId(v string) *UpdateResourceGroupAttributeRequest
	GetNewResourceGroupId() *string
	SetRegionId(v string) *UpdateResourceGroupAttributeRequest
	GetRegionId() *string
}

type UpdateResourceGroupAttributeRequest struct {
	// The ID of the instance that you want to transfer.
	//
	// >  You can call the [DescribeDrdsInstances](https://help.aliyun.com/document_detail/139284.html) operation to view the details of the instances under the account, including the instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds***********
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the resource group that you want to specify.
	//
	// >  You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation to view the details of the resource groups, including the resource group IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-***************
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	// The ID of the region where the instance you want to transfer is located.
	//
	// >  You can call the [DescribeDrdsInstances](https://help.aliyun.com/document_detail/139284.html) operation to view the details of the instances under the account, including the region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateResourceGroupAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupAttributeRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *UpdateResourceGroupAttributeRequest) GetNewResourceGroupId() *string {
	return s.NewResourceGroupId
}

func (s *UpdateResourceGroupAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateResourceGroupAttributeRequest) SetDrdsInstanceId(v string) *UpdateResourceGroupAttributeRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *UpdateResourceGroupAttributeRequest) SetNewResourceGroupId(v string) *UpdateResourceGroupAttributeRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *UpdateResourceGroupAttributeRequest) SetRegionId(v string) *UpdateResourceGroupAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateResourceGroupAttributeRequest) Validate() error {
	return dara.Validate(s)
}
