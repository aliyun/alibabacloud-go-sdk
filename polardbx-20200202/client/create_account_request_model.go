// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountDescription(v string) *CreateAccountRequest
	GetAccountDescription() *string
	SetAccountName(v string) *CreateAccountRequest
	GetAccountName() *string
	SetAccountPassword(v string) *CreateAccountRequest
	GetAccountPassword() *string
	SetAccountPrivilege(v string) *CreateAccountRequest
	GetAccountPrivilege() *string
	SetDBInstanceName(v string) *CreateAccountRequest
	GetDBInstanceName() *string
	SetDBName(v string) *CreateAccountRequest
	GetDBName() *string
	SetRegionId(v string) *CreateAccountRequest
	GetRegionId() *string
	SetSecurityAccountName(v string) *CreateAccountRequest
	GetSecurityAccountName() *string
	SetSecurityAccountPassword(v string) *CreateAccountRequest
	GetSecurityAccountPassword() *string
}

type CreateAccountRequest struct {
	// example:
	//
	// test
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testAccount
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Test@1111
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// example:
	//
	// ReadWrite
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-**************
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// testdb
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// securityAccount
	SecurityAccountName *string `json:"SecurityAccountName,omitempty" xml:"SecurityAccountName,omitempty"`
	// example:
	//
	// securityPassword
	SecurityAccountPassword *string `json:"SecurityAccountPassword,omitempty" xml:"SecurityAccountPassword,omitempty"`
}

func (s CreateAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateAccountRequest) GetAccountDescription() *string {
	return s.AccountDescription
}

func (s *CreateAccountRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateAccountRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *CreateAccountRequest) GetAccountPrivilege() *string {
	return s.AccountPrivilege
}

func (s *CreateAccountRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateAccountRequest) GetDBName() *string {
	return s.DBName
}

func (s *CreateAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAccountRequest) GetSecurityAccountName() *string {
	return s.SecurityAccountName
}

func (s *CreateAccountRequest) GetSecurityAccountPassword() *string {
	return s.SecurityAccountPassword
}

func (s *CreateAccountRequest) SetAccountDescription(v string) *CreateAccountRequest {
	s.AccountDescription = &v
	return s
}

func (s *CreateAccountRequest) SetAccountName(v string) *CreateAccountRequest {
	s.AccountName = &v
	return s
}

func (s *CreateAccountRequest) SetAccountPassword(v string) *CreateAccountRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateAccountRequest) SetAccountPrivilege(v string) *CreateAccountRequest {
	s.AccountPrivilege = &v
	return s
}

func (s *CreateAccountRequest) SetDBInstanceName(v string) *CreateAccountRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateAccountRequest) SetDBName(v string) *CreateAccountRequest {
	s.DBName = &v
	return s
}

func (s *CreateAccountRequest) SetRegionId(v string) *CreateAccountRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAccountRequest) SetSecurityAccountName(v string) *CreateAccountRequest {
	s.SecurityAccountName = &v
	return s
}

func (s *CreateAccountRequest) SetSecurityAccountPassword(v string) *CreateAccountRequest {
	s.SecurityAccountPassword = &v
	return s
}

func (s *CreateAccountRequest) Validate() error {
	return dara.Validate(s)
}
