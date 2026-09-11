// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachGadInstanceDbMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DetachGadInstanceDbMemberRequest
	GetInstanceId() *string
	SetOwnerId(v string) *DetachGadInstanceDbMemberRequest
	GetOwnerId() *string
	SetRegionId(v string) *DetachGadInstanceDbMemberRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DetachGadInstanceDbMemberRequest
	GetResourceGroupId() *string
	SetSlaveDbInstanceId(v string) *DetachGadInstanceDbMemberRequest
	GetSlaveDbInstanceId() *string
}

type DetachGadInstanceDbMemberRequest struct {
	// The instance ID of the active geo-redundancy instance group.
	//
	// example:
	//
	// gad-bp162d4tp0500****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId    *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. Specify this parameter to indicate the region where the instance resides. For more information, see the list of supported regions.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The instance ID of the secondary role.
	//
	// example:
	//
	// rm-sdfghjk****
	SlaveDbInstanceId *string `json:"SlaveDbInstanceId,omitempty" xml:"SlaveDbInstanceId,omitempty"`
}

func (s DetachGadInstanceDbMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachGadInstanceDbMemberRequest) GoString() string {
	return s.String()
}

func (s *DetachGadInstanceDbMemberRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DetachGadInstanceDbMemberRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DetachGadInstanceDbMemberRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachGadInstanceDbMemberRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DetachGadInstanceDbMemberRequest) GetSlaveDbInstanceId() *string {
	return s.SlaveDbInstanceId
}

func (s *DetachGadInstanceDbMemberRequest) SetInstanceId(v string) *DetachGadInstanceDbMemberRequest {
	s.InstanceId = &v
	return s
}

func (s *DetachGadInstanceDbMemberRequest) SetOwnerId(v string) *DetachGadInstanceDbMemberRequest {
	s.OwnerId = &v
	return s
}

func (s *DetachGadInstanceDbMemberRequest) SetRegionId(v string) *DetachGadInstanceDbMemberRequest {
	s.RegionId = &v
	return s
}

func (s *DetachGadInstanceDbMemberRequest) SetResourceGroupId(v string) *DetachGadInstanceDbMemberRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DetachGadInstanceDbMemberRequest) SetSlaveDbInstanceId(v string) *DetachGadInstanceDbMemberRequest {
	s.SlaveDbInstanceId = &v
	return s
}

func (s *DetachGadInstanceDbMemberRequest) Validate() error {
	return dara.Validate(s)
}
