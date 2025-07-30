// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceVpcAuthModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyInstanceVpcAuthModeRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ModifyInstanceVpcAuthModeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyInstanceVpcAuthModeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyInstanceVpcAuthModeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceVpcAuthModeRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyInstanceVpcAuthModeRequest
	GetSecurityToken() *string
	SetVpcAuthMode(v string) *ModifyInstanceVpcAuthModeRequest
	GetVpcAuthMode() *string
}

type ModifyInstanceVpcAuthModeRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// Specifies whether to disable password-free access. Valid values:
	//
	// 	- **Open**: disables password-free access.
	//
	// 	- **Close**: enables password-free access.
	//
	// >  Default value: **Open**.
	//
	// This parameter is required.
	//
	// example:
	//
	// Close
	VpcAuthMode *string `json:"VpcAuthMode,omitempty" xml:"VpcAuthMode,omitempty"`
}

func (s ModifyInstanceVpcAuthModeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceVpcAuthModeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceVpcAuthModeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceVpcAuthModeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyInstanceVpcAuthModeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceVpcAuthModeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceVpcAuthModeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceVpcAuthModeRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyInstanceVpcAuthModeRequest) GetVpcAuthMode() *string {
	return s.VpcAuthMode
}

func (s *ModifyInstanceVpcAuthModeRequest) SetInstanceId(v string) *ModifyInstanceVpcAuthModeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceVpcAuthModeRequest) SetOwnerAccount(v string) *ModifyInstanceVpcAuthModeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceVpcAuthModeRequest) SetOwnerId(v int64) *ModifyInstanceVpcAuthModeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceVpcAuthModeRequest) SetResourceOwnerAccount(v string) *ModifyInstanceVpcAuthModeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceVpcAuthModeRequest) SetResourceOwnerId(v int64) *ModifyInstanceVpcAuthModeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceVpcAuthModeRequest) SetSecurityToken(v string) *ModifyInstanceVpcAuthModeRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyInstanceVpcAuthModeRequest) SetVpcAuthMode(v string) *ModifyInstanceVpcAuthModeRequest {
	s.VpcAuthMode = &v
	return s
}

func (s *ModifyInstanceVpcAuthModeRequest) Validate() error {
	return dara.Validate(s)
}
