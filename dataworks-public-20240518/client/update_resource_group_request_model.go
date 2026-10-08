// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunResourceGroupId(v string) *UpdateResourceGroupRequest
	GetAliyunResourceGroupId() *string
	SetId(v string) *UpdateResourceGroupRequest
	GetId() *string
	SetName(v string) *UpdateResourceGroupRequest
	GetName() *string
	SetRemark(v string) *UpdateResourceGroupRequest
	GetRemark() *string
}

type UpdateResourceGroupRequest struct {
	// The ID of the new Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-aek2kqofrgXXXXX
	AliyunResourceGroupId *string `json:"AliyunResourceGroupId,omitempty" xml:"AliyunResourceGroupId,omitempty"`
	// The ID of the resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// Serverless_res_group_524257424564736_6831777003XXXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The new name that you want to change for the resource group.
	//
	// example:
	//
	// common_resource_group
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The new remarks that you want to modify for the resource group.
	//
	// example:
	//
	// Create a common resource group for common tasks
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s UpdateResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceGroupRequest) GetAliyunResourceGroupId() *string {
	return s.AliyunResourceGroupId
}

func (s *UpdateResourceGroupRequest) GetId() *string {
	return s.Id
}

func (s *UpdateResourceGroupRequest) GetName() *string {
	return s.Name
}

func (s *UpdateResourceGroupRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateResourceGroupRequest) SetAliyunResourceGroupId(v string) *UpdateResourceGroupRequest {
	s.AliyunResourceGroupId = &v
	return s
}

func (s *UpdateResourceGroupRequest) SetId(v string) *UpdateResourceGroupRequest {
	s.Id = &v
	return s
}

func (s *UpdateResourceGroupRequest) SetName(v string) *UpdateResourceGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateResourceGroupRequest) SetRemark(v string) *UpdateResourceGroupRequest {
	s.Remark = &v
	return s
}

func (s *UpdateResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
