// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetAccountPasswordRestrictRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *ResetAccountPasswordRestrictRequest
	GetAccountName() *string
	SetAccountPassword(v string) *ResetAccountPasswordRestrictRequest
	GetAccountPassword() *string
	SetDBInstanceName(v string) *ResetAccountPasswordRestrictRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *ResetAccountPasswordRestrictRequest
	GetRegionId() *string
	SetSecurityAccountName(v string) *ResetAccountPasswordRestrictRequest
	GetSecurityAccountName() *string
	SetSecurityAccountPassword(v string) *ResetAccountPasswordRestrictRequest
	GetSecurityAccountPassword() *string
}

type ResetAccountPasswordRestrictRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// testaccount
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
	// pxc-********
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

func (s ResetAccountPasswordRestrictRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetAccountPasswordRestrictRequest) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordRestrictRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ResetAccountPasswordRestrictRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *ResetAccountPasswordRestrictRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ResetAccountPasswordRestrictRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResetAccountPasswordRestrictRequest) GetSecurityAccountName() *string {
	return s.SecurityAccountName
}

func (s *ResetAccountPasswordRestrictRequest) GetSecurityAccountPassword() *string {
	return s.SecurityAccountPassword
}

func (s *ResetAccountPasswordRestrictRequest) SetAccountName(v string) *ResetAccountPasswordRestrictRequest {
	s.AccountName = &v
	return s
}

func (s *ResetAccountPasswordRestrictRequest) SetAccountPassword(v string) *ResetAccountPasswordRestrictRequest {
	s.AccountPassword = &v
	return s
}

func (s *ResetAccountPasswordRestrictRequest) SetDBInstanceName(v string) *ResetAccountPasswordRestrictRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ResetAccountPasswordRestrictRequest) SetRegionId(v string) *ResetAccountPasswordRestrictRequest {
	s.RegionId = &v
	return s
}

func (s *ResetAccountPasswordRestrictRequest) SetSecurityAccountName(v string) *ResetAccountPasswordRestrictRequest {
	s.SecurityAccountName = &v
	return s
}

func (s *ResetAccountPasswordRestrictRequest) SetSecurityAccountPassword(v string) *ResetAccountPasswordRestrictRequest {
	s.SecurityAccountPassword = &v
	return s
}

func (s *ResetAccountPasswordRestrictRequest) Validate() error {
	return dara.Validate(s)
}
