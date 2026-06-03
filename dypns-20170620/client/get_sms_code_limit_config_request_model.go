// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmsCodeLimitConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *GetSmsCodeLimitConfigRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetSmsCodeLimitConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetSmsCodeLimitConfigRequest
	GetResourceOwnerId() *int64
}

type GetSmsCodeLimitConfigRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetSmsCodeLimitConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSmsCodeLimitConfigRequest) GoString() string {
	return s.String()
}

func (s *GetSmsCodeLimitConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetSmsCodeLimitConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetSmsCodeLimitConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetSmsCodeLimitConfigRequest) SetOwnerId(v int64) *GetSmsCodeLimitConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *GetSmsCodeLimitConfigRequest) SetResourceOwnerAccount(v string) *GetSmsCodeLimitConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetSmsCodeLimitConfigRequest) SetResourceOwnerId(v int64) *GetSmsCodeLimitConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetSmsCodeLimitConfigRequest) Validate() error {
	return dara.Validate(s)
}
