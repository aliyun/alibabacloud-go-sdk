// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCasterResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *UpdateCasterResourceGroupRequest
	GetCasterId() *string
	SetNewResourceGroupId(v string) *UpdateCasterResourceGroupRequest
	GetNewResourceGroupId() *string
	SetOwnerId(v int64) *UpdateCasterResourceGroupRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateCasterResourceGroupRequest
	GetRegionId() *string
}

type UpdateCasterResourceGroupRequest struct {
	// The ID of the production studio.
	//
	// This parameter is required.
	//
	// example:
	//
	// a41300c9-23d3-470e-b9bd-****663e0700
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The ID of the destination resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-aekz7***34cn5ty
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateCasterResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCasterResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateCasterResourceGroupRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *UpdateCasterResourceGroupRequest) GetNewResourceGroupId() *string {
	return s.NewResourceGroupId
}

func (s *UpdateCasterResourceGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateCasterResourceGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateCasterResourceGroupRequest) SetCasterId(v string) *UpdateCasterResourceGroupRequest {
	s.CasterId = &v
	return s
}

func (s *UpdateCasterResourceGroupRequest) SetNewResourceGroupId(v string) *UpdateCasterResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *UpdateCasterResourceGroupRequest) SetOwnerId(v int64) *UpdateCasterResourceGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateCasterResourceGroupRequest) SetRegionId(v string) *UpdateCasterResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateCasterResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
