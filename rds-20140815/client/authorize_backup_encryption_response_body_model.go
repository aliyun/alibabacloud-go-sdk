// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeBackupEncryptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationState(v int32) *AuthorizeBackupEncryptionResponseBody
	GetAuthorizationState() *int32
	SetMessage(v string) *AuthorizeBackupEncryptionResponseBody
	GetMessage() *string
	SetRequestId(v string) *AuthorizeBackupEncryptionResponseBody
	GetRequestId() *string
	SetRoleARN(v string) *AuthorizeBackupEncryptionResponseBody
	GetRoleARN() *string
}

type AuthorizeBackupEncryptionResponseBody struct {
	// example:
	//
	// 1
	AuthorizationState *int32 `json:"AuthorizationState,omitempty" xml:"AuthorizationState,omitempty"`
	// example:
	//
	// create backup encrypt service linked role error.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1A1DD2A4-69F7-5848-AD56-********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// acs:ram::113991************:role/AliyunServiceRoleForRdsBackupEncryption
	RoleARN *string `json:"RoleARN,omitempty" xml:"RoleARN,omitempty"`
}

func (s AuthorizeBackupEncryptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeBackupEncryptionResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeBackupEncryptionResponseBody) GetAuthorizationState() *int32 {
	return s.AuthorizationState
}

func (s *AuthorizeBackupEncryptionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AuthorizeBackupEncryptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthorizeBackupEncryptionResponseBody) GetRoleARN() *string {
	return s.RoleARN
}

func (s *AuthorizeBackupEncryptionResponseBody) SetAuthorizationState(v int32) *AuthorizeBackupEncryptionResponseBody {
	s.AuthorizationState = &v
	return s
}

func (s *AuthorizeBackupEncryptionResponseBody) SetMessage(v string) *AuthorizeBackupEncryptionResponseBody {
	s.Message = &v
	return s
}

func (s *AuthorizeBackupEncryptionResponseBody) SetRequestId(v string) *AuthorizeBackupEncryptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthorizeBackupEncryptionResponseBody) SetRoleARN(v string) *AuthorizeBackupEncryptionResponseBody {
	s.RoleARN = &v
	return s
}

func (s *AuthorizeBackupEncryptionResponseBody) Validate() error {
	return dara.Validate(s)
}
