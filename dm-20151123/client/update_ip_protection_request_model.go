// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIpProtectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpProtection(v string) *UpdateIpProtectionRequest
	GetIpProtection() *string
	SetOwnerId(v int64) *UpdateIpProtectionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpdateIpProtectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateIpProtectionRequest
	GetResourceOwnerId() *int64
}

type UpdateIpProtectionRequest struct {
	// IP protection switch, On: 1 Off: 0
	//
	// example:
	//
	// 0
	IpProtection         *string `json:"IpProtection,omitempty" xml:"IpProtection,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateIpProtectionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIpProtectionRequest) GoString() string {
	return s.String()
}

func (s *UpdateIpProtectionRequest) GetIpProtection() *string {
	return s.IpProtection
}

func (s *UpdateIpProtectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateIpProtectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateIpProtectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateIpProtectionRequest) SetIpProtection(v string) *UpdateIpProtectionRequest {
	s.IpProtection = &v
	return s
}

func (s *UpdateIpProtectionRequest) SetOwnerId(v int64) *UpdateIpProtectionRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateIpProtectionRequest) SetResourceOwnerAccount(v string) *UpdateIpProtectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateIpProtectionRequest) SetResourceOwnerId(v int64) *UpdateIpProtectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateIpProtectionRequest) Validate() error {
	return dara.Validate(s)
}
