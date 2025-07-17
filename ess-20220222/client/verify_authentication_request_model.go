// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyAuthenticationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOnlyCheck(v bool) *VerifyAuthenticationRequest
	GetOnlyCheck() *bool
	SetOwnerId(v int64) *VerifyAuthenticationRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *VerifyAuthenticationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *VerifyAuthenticationRequest
	GetResourceOwnerId() *int64
	SetUid(v int64) *VerifyAuthenticationRequest
	GetUid() *int64
}

type VerifyAuthenticationRequest struct {
	// Specifies whether to check only the authorization status. Valid values:
	//
	// 	- true: checks only the authorization status. The service-linked role is not created.
	//
	// 	- false (default): checks the authorization status and resource usage.
	//
	// example:
	//
	// false
	OnlyCheck            *bool   `json:"OnlyCheck,omitempty" xml:"OnlyCheck,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of your Alibaba Cloud account.
	//
	// example:
	//
	// 12345678123*****
	Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s VerifyAuthenticationRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifyAuthenticationRequest) GoString() string {
	return s.String()
}

func (s *VerifyAuthenticationRequest) GetOnlyCheck() *bool {
	return s.OnlyCheck
}

func (s *VerifyAuthenticationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *VerifyAuthenticationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *VerifyAuthenticationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *VerifyAuthenticationRequest) GetUid() *int64 {
	return s.Uid
}

func (s *VerifyAuthenticationRequest) SetOnlyCheck(v bool) *VerifyAuthenticationRequest {
	s.OnlyCheck = &v
	return s
}

func (s *VerifyAuthenticationRequest) SetOwnerId(v int64) *VerifyAuthenticationRequest {
	s.OwnerId = &v
	return s
}

func (s *VerifyAuthenticationRequest) SetResourceOwnerAccount(v string) *VerifyAuthenticationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *VerifyAuthenticationRequest) SetResourceOwnerId(v int64) *VerifyAuthenticationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *VerifyAuthenticationRequest) SetUid(v int64) *VerifyAuthenticationRequest {
	s.Uid = &v
	return s
}

func (s *VerifyAuthenticationRequest) Validate() error {
	return dara.Validate(s)
}
