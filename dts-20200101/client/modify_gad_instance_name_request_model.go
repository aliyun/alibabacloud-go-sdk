// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGadInstanceNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyGadInstanceNameRequest
	GetInstanceId() *string
	SetInstanceName(v string) *ModifyGadInstanceNameRequest
	GetInstanceName() *string
	SetOwnerId(v string) *ModifyGadInstanceNameRequest
	GetOwnerId() *string
	SetRegionId(v string) *ModifyGadInstanceNameRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ModifyGadInstanceNameRequest
	GetResourceGroupId() *string
}

type ModifyGadInstanceNameRequest struct {
	// example:
	//
	// rm-bp1i99e8l7913****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// rm-uf6b0m001ir8mr9i9
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	OwnerId      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-aek26mat2ldb4oy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ModifyGadInstanceNameRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyGadInstanceNameRequest) GoString() string {
	return s.String()
}

func (s *ModifyGadInstanceNameRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyGadInstanceNameRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyGadInstanceNameRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ModifyGadInstanceNameRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyGadInstanceNameRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyGadInstanceNameRequest) SetInstanceId(v string) *ModifyGadInstanceNameRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyGadInstanceNameRequest) SetInstanceName(v string) *ModifyGadInstanceNameRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyGadInstanceNameRequest) SetOwnerId(v string) *ModifyGadInstanceNameRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyGadInstanceNameRequest) SetRegionId(v string) *ModifyGadInstanceNameRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyGadInstanceNameRequest) SetResourceGroupId(v string) *ModifyGadInstanceNameRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyGadInstanceNameRequest) Validate() error {
	return dara.Validate(s)
}
