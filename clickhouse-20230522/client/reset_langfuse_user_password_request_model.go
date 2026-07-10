// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetLangfuseUserPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ResetLangfuseUserPasswordRequest
	GetDBInstanceId() *string
	SetEmail(v string) *ResetLangfuseUserPasswordRequest
	GetEmail() *string
	SetNewPassword(v string) *ResetLangfuseUserPasswordRequest
	GetNewPassword() *string
	SetRegionId(v string) *ResetLangfuseUserPasswordRequest
	GetRegionId() *string
}

type ResetLangfuseUserPasswordRequest struct {
	// The Langfuse instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lfs-****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The email address of the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// john@company.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The new user password. The password must meet the following rules:
	//
	// The password must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// The supported special characters are !@#$%^&*()_+-=.
	//
	// The password must be 8 to 32 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2F9e9@******
	NewPassword *string `json:"NewPassword,omitempty" xml:"NewPassword,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ResetLangfuseUserPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetLangfuseUserPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetLangfuseUserPasswordRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ResetLangfuseUserPasswordRequest) GetEmail() *string {
	return s.Email
}

func (s *ResetLangfuseUserPasswordRequest) GetNewPassword() *string {
	return s.NewPassword
}

func (s *ResetLangfuseUserPasswordRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResetLangfuseUserPasswordRequest) SetDBInstanceId(v string) *ResetLangfuseUserPasswordRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ResetLangfuseUserPasswordRequest) SetEmail(v string) *ResetLangfuseUserPasswordRequest {
	s.Email = &v
	return s
}

func (s *ResetLangfuseUserPasswordRequest) SetNewPassword(v string) *ResetLangfuseUserPasswordRequest {
	s.NewPassword = &v
	return s
}

func (s *ResetLangfuseUserPasswordRequest) SetRegionId(v string) *ResetLangfuseUserPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *ResetLangfuseUserPasswordRequest) Validate() error {
	return dara.Validate(s)
}
