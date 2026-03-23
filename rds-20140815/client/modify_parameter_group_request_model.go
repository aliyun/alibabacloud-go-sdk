// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyParameterGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModifyMode(v string) *ModifyParameterGroupRequest
	GetModifyMode() *string
	SetOwnerId(v int64) *ModifyParameterGroupRequest
	GetOwnerId() *int64
	SetParameterGroupDesc(v string) *ModifyParameterGroupRequest
	GetParameterGroupDesc() *string
	SetParameterGroupId(v string) *ModifyParameterGroupRequest
	GetParameterGroupId() *string
	SetParameterGroupName(v string) *ModifyParameterGroupRequest
	GetParameterGroupName() *string
	SetParameters(v string) *ModifyParameterGroupRequest
	GetParameters() *string
	SetRegionId(v string) *ModifyParameterGroupRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ModifyParameterGroupRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyParameterGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyParameterGroupRequest
	GetResourceOwnerId() *int64
}

type ModifyParameterGroupRequest struct {
	ModifyMode         *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ParameterGroupDesc *string `json:"ParameterGroupDesc,omitempty" xml:"ParameterGroupDesc,omitempty"`
	// This parameter is required.
	ParameterGroupId   *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	ParameterGroupName *string `json:"ParameterGroupName,omitempty" xml:"ParameterGroupName,omitempty"`
	Parameters         *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyParameterGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyParameterGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyParameterGroupRequest) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *ModifyParameterGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyParameterGroupRequest) GetParameterGroupDesc() *string {
	return s.ParameterGroupDesc
}

func (s *ModifyParameterGroupRequest) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *ModifyParameterGroupRequest) GetParameterGroupName() *string {
	return s.ParameterGroupName
}

func (s *ModifyParameterGroupRequest) GetParameters() *string {
	return s.Parameters
}

func (s *ModifyParameterGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyParameterGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyParameterGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyParameterGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyParameterGroupRequest) SetModifyMode(v string) *ModifyParameterGroupRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifyParameterGroupRequest) SetOwnerId(v int64) *ModifyParameterGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyParameterGroupRequest) SetParameterGroupDesc(v string) *ModifyParameterGroupRequest {
	s.ParameterGroupDesc = &v
	return s
}

func (s *ModifyParameterGroupRequest) SetParameterGroupId(v string) *ModifyParameterGroupRequest {
	s.ParameterGroupId = &v
	return s
}

func (s *ModifyParameterGroupRequest) SetParameterGroupName(v string) *ModifyParameterGroupRequest {
	s.ParameterGroupName = &v
	return s
}

func (s *ModifyParameterGroupRequest) SetParameters(v string) *ModifyParameterGroupRequest {
	s.Parameters = &v
	return s
}

func (s *ModifyParameterGroupRequest) SetRegionId(v string) *ModifyParameterGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyParameterGroupRequest) SetResourceGroupId(v string) *ModifyParameterGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyParameterGroupRequest) SetResourceOwnerAccount(v string) *ModifyParameterGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyParameterGroupRequest) SetResourceOwnerId(v int64) *ModifyParameterGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyParameterGroupRequest) Validate() error {
	return dara.Validate(s)
}
