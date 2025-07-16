// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAccountPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *ResetAccountPasswordRequest
	GetAccountName() *string
	SetAccountPassword(v string) *ResetAccountPasswordRequest
	GetAccountPassword() *string
	SetDBInstanceName(v string) *ResetAccountPasswordRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *ResetAccountPasswordRequest
	GetRegionId() *string
	SetSecurityAccountName(v string) *ResetAccountPasswordRequest
	GetSecurityAccountName() *string
	SetSecurityAccountPassword(v string) *ResetAccountPasswordRequest
	GetSecurityAccountPassword() *string
}

type ResetAccountPasswordRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// account
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// *****
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-htri0ori2r4k9p
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// account_sec
	SecurityAccountName *string `json:"SecurityAccountName,omitempty" xml:"SecurityAccountName,omitempty"`
	// example:
	//
	// *****
	SecurityAccountPassword *string `json:"SecurityAccountPassword,omitempty" xml:"SecurityAccountPassword,omitempty"`
}

func (s ResetAccountPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ResetAccountPasswordRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *ResetAccountPasswordRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ResetAccountPasswordRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResetAccountPasswordRequest) GetSecurityAccountName() *string {
	return s.SecurityAccountName
}

func (s *ResetAccountPasswordRequest) GetSecurityAccountPassword() *string {
	return s.SecurityAccountPassword
}

func (s *ResetAccountPasswordRequest) SetAccountName(v string) *ResetAccountPasswordRequest {
	s.AccountName = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetAccountPassword(v string) *ResetAccountPasswordRequest {
	s.AccountPassword = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetDBInstanceName(v string) *ResetAccountPasswordRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetRegionId(v string) *ResetAccountPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetSecurityAccountName(v string) *ResetAccountPasswordRequest {
	s.SecurityAccountName = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetSecurityAccountPassword(v string) *ResetAccountPasswordRequest {
	s.SecurityAccountPassword = &v
	return s
}

func (s *ResetAccountPasswordRequest) Validate() error {
	return dara.Validate(s)
}
