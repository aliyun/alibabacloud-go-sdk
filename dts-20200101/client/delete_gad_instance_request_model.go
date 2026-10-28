// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGadInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteGadInstanceRequest
	GetInstanceId() *string
	SetOwnerId(v string) *DeleteGadInstanceRequest
	GetOwnerId() *string
	SetRegionId(v string) *DeleteGadInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DeleteGadInstanceRequest
	GetResourceGroupId() *string
}

type DeleteGadInstanceRequest struct {
	// The instance ID.
	//
	// example:
	//
	// gad-pfx10jt8b8***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId    *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hongkong
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DeleteGadInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGadInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteGadInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteGadInstanceRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DeleteGadInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteGadInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteGadInstanceRequest) SetInstanceId(v string) *DeleteGadInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteGadInstanceRequest) SetOwnerId(v string) *DeleteGadInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteGadInstanceRequest) SetRegionId(v string) *DeleteGadInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteGadInstanceRequest) SetResourceGroupId(v string) *DeleteGadInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteGadInstanceRequest) Validate() error {
	return dara.Validate(s)
}
