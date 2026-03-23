// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateParameterGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEngine(v string) *CreateParameterGroupRequest
	GetEngine() *string
	SetEngineVersion(v string) *CreateParameterGroupRequest
	GetEngineVersion() *string
	SetOwnerId(v int64) *CreateParameterGroupRequest
	GetOwnerId() *int64
	SetParameterGroupDesc(v string) *CreateParameterGroupRequest
	GetParameterGroupDesc() *string
	SetParameterGroupName(v string) *CreateParameterGroupRequest
	GetParameterGroupName() *string
	SetParameters(v string) *CreateParameterGroupRequest
	GetParameters() *string
	SetRegionId(v string) *CreateParameterGroupRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateParameterGroupRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateParameterGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateParameterGroupRequest
	GetResourceOwnerId() *int64
}

type CreateParameterGroupRequest struct {
	// This parameter is required.
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// This parameter is required.
	EngineVersion      *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ParameterGroupDesc *string `json:"ParameterGroupDesc,omitempty" xml:"ParameterGroupDesc,omitempty"`
	// This parameter is required.
	ParameterGroupName *string `json:"ParameterGroupName,omitempty" xml:"ParameterGroupName,omitempty"`
	// This parameter is required.
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateParameterGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateParameterGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateParameterGroupRequest) GetEngine() *string {
	return s.Engine
}

func (s *CreateParameterGroupRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CreateParameterGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateParameterGroupRequest) GetParameterGroupDesc() *string {
	return s.ParameterGroupDesc
}

func (s *CreateParameterGroupRequest) GetParameterGroupName() *string {
	return s.ParameterGroupName
}

func (s *CreateParameterGroupRequest) GetParameters() *string {
	return s.Parameters
}

func (s *CreateParameterGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateParameterGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateParameterGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateParameterGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateParameterGroupRequest) SetEngine(v string) *CreateParameterGroupRequest {
	s.Engine = &v
	return s
}

func (s *CreateParameterGroupRequest) SetEngineVersion(v string) *CreateParameterGroupRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateParameterGroupRequest) SetOwnerId(v int64) *CreateParameterGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateParameterGroupRequest) SetParameterGroupDesc(v string) *CreateParameterGroupRequest {
	s.ParameterGroupDesc = &v
	return s
}

func (s *CreateParameterGroupRequest) SetParameterGroupName(v string) *CreateParameterGroupRequest {
	s.ParameterGroupName = &v
	return s
}

func (s *CreateParameterGroupRequest) SetParameters(v string) *CreateParameterGroupRequest {
	s.Parameters = &v
	return s
}

func (s *CreateParameterGroupRequest) SetRegionId(v string) *CreateParameterGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateParameterGroupRequest) SetResourceGroupId(v string) *CreateParameterGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateParameterGroupRequest) SetResourceOwnerAccount(v string) *CreateParameterGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateParameterGroupRequest) SetResourceOwnerId(v int64) *CreateParameterGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateParameterGroupRequest) Validate() error {
	return dara.Validate(s)
}
