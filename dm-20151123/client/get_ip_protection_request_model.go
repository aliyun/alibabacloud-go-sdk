// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIpProtectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *GetIpProtectionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetIpProtectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetIpProtectionRequest
	GetResourceOwnerId() *int64
}

type GetIpProtectionRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetIpProtectionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIpProtectionRequest) GoString() string {
	return s.String()
}

func (s *GetIpProtectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetIpProtectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetIpProtectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetIpProtectionRequest) SetOwnerId(v int64) *GetIpProtectionRequest {
	s.OwnerId = &v
	return s
}

func (s *GetIpProtectionRequest) SetResourceOwnerAccount(v string) *GetIpProtectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetIpProtectionRequest) SetResourceOwnerId(v int64) *GetIpProtectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetIpProtectionRequest) Validate() error {
	return dara.Validate(s)
}
