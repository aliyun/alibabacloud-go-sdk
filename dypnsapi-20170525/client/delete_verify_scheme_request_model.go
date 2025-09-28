// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVerifySchemeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerId(v int64) *DeleteVerifySchemeRequest
	GetCustomerId() *int64
	SetOwnerId(v int64) *DeleteVerifySchemeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteVerifySchemeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteVerifySchemeRequest
	GetResourceOwnerId() *int64
	SetSchemeCode(v string) *DeleteVerifySchemeRequest
	GetSchemeCode() *string
}

type DeleteVerifySchemeRequest struct {
	// The user ID.
	//
	// example:
	//
	// 12345678
	CustomerId           *int64  `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The service code.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC10000014164****
	SchemeCode *string `json:"SchemeCode,omitempty" xml:"SchemeCode,omitempty"`
}

func (s DeleteVerifySchemeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVerifySchemeRequest) GoString() string {
	return s.String()
}

func (s *DeleteVerifySchemeRequest) GetCustomerId() *int64 {
	return s.CustomerId
}

func (s *DeleteVerifySchemeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteVerifySchemeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteVerifySchemeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteVerifySchemeRequest) GetSchemeCode() *string {
	return s.SchemeCode
}

func (s *DeleteVerifySchemeRequest) SetCustomerId(v int64) *DeleteVerifySchemeRequest {
	s.CustomerId = &v
	return s
}

func (s *DeleteVerifySchemeRequest) SetOwnerId(v int64) *DeleteVerifySchemeRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteVerifySchemeRequest) SetResourceOwnerAccount(v string) *DeleteVerifySchemeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteVerifySchemeRequest) SetResourceOwnerId(v int64) *DeleteVerifySchemeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteVerifySchemeRequest) SetSchemeCode(v string) *DeleteVerifySchemeRequest {
	s.SchemeCode = &v
	return s
}

func (s *DeleteVerifySchemeRequest) Validate() error {
	return dara.Validate(s)
}
