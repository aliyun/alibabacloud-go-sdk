// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLindormV2InstanceParameterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateLindormV2InstanceParameterRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *UpdateLindormV2InstanceParameterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateLindormV2InstanceParameterRequest
	GetOwnerId() *int64
	SetParameterKey(v string) *UpdateLindormV2InstanceParameterRequest
	GetParameterKey() *string
	SetParameterValue(v string) *UpdateLindormV2InstanceParameterRequest
	GetParameterValue() *string
	SetResourceOwnerAccount(v string) *UpdateLindormV2InstanceParameterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateLindormV2InstanceParameterRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *UpdateLindormV2InstanceParameterRequest
	GetSecurityToken() *string
	SetUpdateType(v string) *UpdateLindormV2InstanceParameterRequest
	GetUpdateType() *string
}

type UpdateLindormV2InstanceParameterRequest struct {
	// This parameter is required.
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// This parameter is required.
	ParameterValue       *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	UpdateType           *string `json:"UpdateType,omitempty" xml:"UpdateType,omitempty"`
}

func (s UpdateLindormV2InstanceParameterRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLindormV2InstanceParameterRequest) GoString() string {
	return s.String()
}

func (s *UpdateLindormV2InstanceParameterRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateLindormV2InstanceParameterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateLindormV2InstanceParameterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateLindormV2InstanceParameterRequest) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *UpdateLindormV2InstanceParameterRequest) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *UpdateLindormV2InstanceParameterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateLindormV2InstanceParameterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateLindormV2InstanceParameterRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *UpdateLindormV2InstanceParameterRequest) GetUpdateType() *string {
	return s.UpdateType
}

func (s *UpdateLindormV2InstanceParameterRequest) SetInstanceId(v string) *UpdateLindormV2InstanceParameterRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateLindormV2InstanceParameterRequest) SetOwnerAccount(v string) *UpdateLindormV2InstanceParameterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateLindormV2InstanceParameterRequest) SetOwnerId(v int64) *UpdateLindormV2InstanceParameterRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateLindormV2InstanceParameterRequest) SetParameterKey(v string) *UpdateLindormV2InstanceParameterRequest {
	s.ParameterKey = &v
	return s
}

func (s *UpdateLindormV2InstanceParameterRequest) SetParameterValue(v string) *UpdateLindormV2InstanceParameterRequest {
	s.ParameterValue = &v
	return s
}

func (s *UpdateLindormV2InstanceParameterRequest) SetResourceOwnerAccount(v string) *UpdateLindormV2InstanceParameterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateLindormV2InstanceParameterRequest) SetResourceOwnerId(v int64) *UpdateLindormV2InstanceParameterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateLindormV2InstanceParameterRequest) SetSecurityToken(v string) *UpdateLindormV2InstanceParameterRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpdateLindormV2InstanceParameterRequest) SetUpdateType(v string) *UpdateLindormV2InstanceParameterRequest {
	s.UpdateType = &v
	return s
}

func (s *UpdateLindormV2InstanceParameterRequest) Validate() error {
	return dara.Validate(s)
}
