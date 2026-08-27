// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPasswordEncrypted(v string) *ResetPasswordRequest
	GetPasswordEncrypted() *string
	SetTenantId(v string) *ResetPasswordRequest
	GetTenantId() *string
	SetWnUserId(v string) *ResetPasswordRequest
	GetWnUserId() *string
}

type ResetPasswordRequest struct {
	// The base64-encoded password ciphertext encrypted with RSA-OAEP-SHA256. This parameter is required and cannot be empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	PasswordEncrypted *string `json:"passwordEncrypted,omitempty" xml:"passwordEncrypted,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 21577
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The ID of the target user (WINNEXO platform user ID).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	WnUserId *string `json:"wnUserId,omitempty" xml:"wnUserId,omitempty"`
}

func (s ResetPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetPasswordRequest) GetPasswordEncrypted() *string {
	return s.PasswordEncrypted
}

func (s *ResetPasswordRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ResetPasswordRequest) GetWnUserId() *string {
	return s.WnUserId
}

func (s *ResetPasswordRequest) SetPasswordEncrypted(v string) *ResetPasswordRequest {
	s.PasswordEncrypted = &v
	return s
}

func (s *ResetPasswordRequest) SetTenantId(v string) *ResetPasswordRequest {
	s.TenantId = &v
	return s
}

func (s *ResetPasswordRequest) SetWnUserId(v string) *ResetPasswordRequest {
	s.WnUserId = &v
	return s
}

func (s *ResetPasswordRequest) Validate() error {
	return dara.Validate(s)
}
