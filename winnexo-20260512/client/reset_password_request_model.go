// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ResetPasswordRequest
	GetAppToken() *string
	SetOpsUserCode(v string) *ResetPasswordRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ResetPasswordRequest
	GetOpsUserId() *string
	SetPasswordEncrypted(v string) *ResetPasswordRequest
	GetPasswordEncrypted() *string
	SetTenantId(v string) *ResetPasswordRequest
	GetTenantId() *string
	SetWnUserId(v string) *ResetPasswordRequest
	GetWnUserId() *string
}

type ResetPasswordRequest struct {
	// The platform application token issued by the console. If this parameter is passed, the operation is performed as the user specified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The proxy user code, which is unique within the tenant. This parameter is used only when the application credential is configured as source_uid.
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// The proxy user ID, which is the WINNEXO platform user ID. This parameter is used only when the application credential is configured as winnexo_uid.
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The Base64-encoded password ciphertext encrypted by using RSA-OAEP-SHA256. This parameter is required and cannot be empty.
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
	// The target user ID, which is the WINNEXO platform user ID.
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

func (s *ResetPasswordRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ResetPasswordRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ResetPasswordRequest) GetOpsUserId() *string {
	return s.OpsUserId
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

func (s *ResetPasswordRequest) SetAppToken(v string) *ResetPasswordRequest {
	s.AppToken = &v
	return s
}

func (s *ResetPasswordRequest) SetOpsUserCode(v string) *ResetPasswordRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ResetPasswordRequest) SetOpsUserId(v string) *ResetPasswordRequest {
	s.OpsUserId = &v
	return s
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
