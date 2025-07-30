// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceParameterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyInstanceParameterRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ModifyInstanceParameterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyInstanceParameterRequest
	GetOwnerId() *int64
	SetParameterGroupId(v string) *ModifyInstanceParameterRequest
	GetParameterGroupId() *string
	SetParameters(v string) *ModifyInstanceParameterRequest
	GetParameters() *string
	SetRegionId(v string) *ModifyInstanceParameterRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyInstanceParameterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceParameterRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyInstanceParameterRequest
	GetSecurityToken() *string
}

type ModifyInstanceParameterRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The parameter template ID.
	//
	// > You can view the list of parameter templates in the target region, including the parameter template ID, through the DescribeParameterGroups interface.
	//
	// example:
	//
	// g-idhwofwofewhf****
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	// The information about parameters.
	//
	// example:
	//
	// {"hz": "50"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyInstanceParameterRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceParameterRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceParameterRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceParameterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyInstanceParameterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceParameterRequest) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *ModifyInstanceParameterRequest) GetParameters() *string {
	return s.Parameters
}

func (s *ModifyInstanceParameterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceParameterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceParameterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceParameterRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyInstanceParameterRequest) SetInstanceId(v string) *ModifyInstanceParameterRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceParameterRequest) SetOwnerAccount(v string) *ModifyInstanceParameterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceParameterRequest) SetOwnerId(v int64) *ModifyInstanceParameterRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceParameterRequest) SetParameterGroupId(v string) *ModifyInstanceParameterRequest {
	s.ParameterGroupId = &v
	return s
}

func (s *ModifyInstanceParameterRequest) SetParameters(v string) *ModifyInstanceParameterRequest {
	s.Parameters = &v
	return s
}

func (s *ModifyInstanceParameterRequest) SetRegionId(v string) *ModifyInstanceParameterRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceParameterRequest) SetResourceOwnerAccount(v string) *ModifyInstanceParameterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceParameterRequest) SetResourceOwnerId(v int64) *ModifyInstanceParameterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceParameterRequest) SetSecurityToken(v string) *ModifyInstanceParameterRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyInstanceParameterRequest) Validate() error {
	return dara.Validate(s)
}
