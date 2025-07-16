// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSuperAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountDescription(v string) *CreateSuperAccountRequest
	GetAccountDescription() *string
	SetAccountName(v string) *CreateSuperAccountRequest
	GetAccountName() *string
	SetAccountPassword(v string) *CreateSuperAccountRequest
	GetAccountPassword() *string
	SetDBInstanceName(v string) *CreateSuperAccountRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *CreateSuperAccountRequest
	GetRegionId() *string
}

type CreateSuperAccountRequest struct {
	// example:
	//
	// testdbadescription
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dba
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testdbapassword
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-************
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateSuperAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSuperAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateSuperAccountRequest) GetAccountDescription() *string {
	return s.AccountDescription
}

func (s *CreateSuperAccountRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateSuperAccountRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *CreateSuperAccountRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateSuperAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSuperAccountRequest) SetAccountDescription(v string) *CreateSuperAccountRequest {
	s.AccountDescription = &v
	return s
}

func (s *CreateSuperAccountRequest) SetAccountName(v string) *CreateSuperAccountRequest {
	s.AccountName = &v
	return s
}

func (s *CreateSuperAccountRequest) SetAccountPassword(v string) *CreateSuperAccountRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateSuperAccountRequest) SetDBInstanceName(v string) *CreateSuperAccountRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateSuperAccountRequest) SetRegionId(v string) *CreateSuperAccountRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSuperAccountRequest) Validate() error {
	return dara.Validate(s)
}
