// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckLdpsColumnarIndexStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CheckLdpsColumnarIndexStatusRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *CheckLdpsColumnarIndexStatusRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CheckLdpsColumnarIndexStatusRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CheckLdpsColumnarIndexStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CheckLdpsColumnarIndexStatusRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *CheckLdpsColumnarIndexStatusRequest
	GetSecurityToken() *string
}

type CheckLdpsColumnarIndexStatusRequest struct {
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CheckLdpsColumnarIndexStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckLdpsColumnarIndexStatusRequest) GoString() string {
	return s.String()
}

func (s *CheckLdpsColumnarIndexStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CheckLdpsColumnarIndexStatusRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CheckLdpsColumnarIndexStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CheckLdpsColumnarIndexStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CheckLdpsColumnarIndexStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CheckLdpsColumnarIndexStatusRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CheckLdpsColumnarIndexStatusRequest) SetInstanceId(v string) *CheckLdpsColumnarIndexStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *CheckLdpsColumnarIndexStatusRequest) SetOwnerAccount(v string) *CheckLdpsColumnarIndexStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckLdpsColumnarIndexStatusRequest) SetOwnerId(v int64) *CheckLdpsColumnarIndexStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckLdpsColumnarIndexStatusRequest) SetResourceOwnerAccount(v string) *CheckLdpsColumnarIndexStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckLdpsColumnarIndexStatusRequest) SetResourceOwnerId(v int64) *CheckLdpsColumnarIndexStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckLdpsColumnarIndexStatusRequest) SetSecurityToken(v string) *CheckLdpsColumnarIndexStatusRequest {
	s.SecurityToken = &v
	return s
}

func (s *CheckLdpsColumnarIndexStatusRequest) Validate() error {
	return dara.Validate(s)
}
