// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteParameterGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DeleteParameterGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteParameterGroupRequest
	GetOwnerId() *int64
	SetParameterGroupId(v string) *DeleteParameterGroupRequest
	GetParameterGroupId() *string
	SetResourceOwnerAccount(v string) *DeleteParameterGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteParameterGroupRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DeleteParameterGroupRequest
	GetSecurityToken() *string
}

type DeleteParameterGroupRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the parameter template.
	//
	// This parameter is required.
	//
	// example:
	//
	// rpg-sys-00*****
	ParameterGroupId     *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteParameterGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteParameterGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteParameterGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteParameterGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteParameterGroupRequest) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *DeleteParameterGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteParameterGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteParameterGroupRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteParameterGroupRequest) SetOwnerAccount(v string) *DeleteParameterGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteParameterGroupRequest) SetOwnerId(v int64) *DeleteParameterGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteParameterGroupRequest) SetParameterGroupId(v string) *DeleteParameterGroupRequest {
	s.ParameterGroupId = &v
	return s
}

func (s *DeleteParameterGroupRequest) SetResourceOwnerAccount(v string) *DeleteParameterGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteParameterGroupRequest) SetResourceOwnerId(v int64) *DeleteParameterGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteParameterGroupRequest) SetSecurityToken(v string) *DeleteParameterGroupRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteParameterGroupRequest) Validate() error {
	return dara.Validate(s)
}
