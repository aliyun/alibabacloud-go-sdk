// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetOpenSearchPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *ResetOpenSearchPasswordRequest
	GetAccountName() *string
	SetAccountPassword(v string) *ResetOpenSearchPasswordRequest
	GetAccountPassword() *string
	SetDBInstanceName(v string) *ResetOpenSearchPasswordRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *ResetOpenSearchPasswordRequest
	GetRegionId() *string
}

type ResetOpenSearchPasswordRequest struct {
	// The account name.
	//
	// This parameter is required.
	//
	// example:
	//
	// testAccount
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The new password. The password must be 6 to 32 characters in length and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters (!@#$&%^*()_+-=).
	//
	// This parameter is required.
	//
	// example:
	//
	// *****
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ResetOpenSearchPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetOpenSearchPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetOpenSearchPasswordRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ResetOpenSearchPasswordRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *ResetOpenSearchPasswordRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ResetOpenSearchPasswordRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResetOpenSearchPasswordRequest) SetAccountName(v string) *ResetOpenSearchPasswordRequest {
	s.AccountName = &v
	return s
}

func (s *ResetOpenSearchPasswordRequest) SetAccountPassword(v string) *ResetOpenSearchPasswordRequest {
	s.AccountPassword = &v
	return s
}

func (s *ResetOpenSearchPasswordRequest) SetDBInstanceName(v string) *ResetOpenSearchPasswordRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ResetOpenSearchPasswordRequest) SetRegionId(v string) *ResetOpenSearchPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *ResetOpenSearchPasswordRequest) Validate() error {
	return dara.Validate(s)
}
