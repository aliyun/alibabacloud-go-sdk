// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckBackupEncryptionAuthorizedResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationState(v string) *CheckBackupEncryptionAuthorizedResponseBody
	GetAuthorizationState() *string
	SetRequestId(v string) *CheckBackupEncryptionAuthorizedResponseBody
	GetRequestId() *string
	SetRoleARN(v string) *CheckBackupEncryptionAuthorizedResponseBody
	GetRoleARN() *string
}

type CheckBackupEncryptionAuthorizedResponseBody struct {
	// example:
	//
	// 1
	AuthorizationState *string `json:"AuthorizationState,omitempty" xml:"AuthorizationState,omitempty"`
	// example:
	//
	// CB07C463-7428-50AA-9E39-********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// acs:ram::1139916************:role/AliyunServiceRoleForRdsBackupEncryption
	RoleARN *string `json:"RoleARN,omitempty" xml:"RoleARN,omitempty"`
}

func (s CheckBackupEncryptionAuthorizedResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckBackupEncryptionAuthorizedResponseBody) GoString() string {
	return s.String()
}

func (s *CheckBackupEncryptionAuthorizedResponseBody) GetAuthorizationState() *string {
	return s.AuthorizationState
}

func (s *CheckBackupEncryptionAuthorizedResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckBackupEncryptionAuthorizedResponseBody) GetRoleARN() *string {
	return s.RoleARN
}

func (s *CheckBackupEncryptionAuthorizedResponseBody) SetAuthorizationState(v string) *CheckBackupEncryptionAuthorizedResponseBody {
	s.AuthorizationState = &v
	return s
}

func (s *CheckBackupEncryptionAuthorizedResponseBody) SetRequestId(v string) *CheckBackupEncryptionAuthorizedResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckBackupEncryptionAuthorizedResponseBody) SetRoleARN(v string) *CheckBackupEncryptionAuthorizedResponseBody {
	s.RoleARN = &v
	return s
}

func (s *CheckBackupEncryptionAuthorizedResponseBody) Validate() error {
	return dara.Validate(s)
}
