// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *DeleteAccountRequest
	GetAccountName() *string
	SetDBInstanceName(v string) *DeleteAccountRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DeleteAccountRequest
	GetRegionId() *string
	SetSecurityAccountName(v string) *DeleteAccountRequest
	GetSecurityAccountName() *string
	SetSecurityAccountPassword(v string) *DeleteAccountRequest
	GetSecurityAccountPassword() *string
}

type DeleteAccountRequest struct {
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
	// pxc-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
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

func (s DeleteAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccountRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *DeleteAccountRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DeleteAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAccountRequest) GetSecurityAccountName() *string {
	return s.SecurityAccountName
}

func (s *DeleteAccountRequest) GetSecurityAccountPassword() *string {
	return s.SecurityAccountPassword
}

func (s *DeleteAccountRequest) SetAccountName(v string) *DeleteAccountRequest {
	s.AccountName = &v
	return s
}

func (s *DeleteAccountRequest) SetDBInstanceName(v string) *DeleteAccountRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DeleteAccountRequest) SetRegionId(v string) *DeleteAccountRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAccountRequest) SetSecurityAccountName(v string) *DeleteAccountRequest {
	s.SecurityAccountName = &v
	return s
}

func (s *DeleteAccountRequest) SetSecurityAccountPassword(v string) *DeleteAccountRequest {
	s.SecurityAccountPassword = &v
	return s
}

func (s *DeleteAccountRequest) Validate() error {
	return dara.Validate(s)
}
