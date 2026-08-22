// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpenSearchAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *CreateOpenSearchAccountRequest
	GetAccountName() *string
	SetAccountPassword(v string) *CreateOpenSearchAccountRequest
	GetAccountPassword() *string
	SetDBInstanceName(v string) *CreateOpenSearchAccountRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *CreateOpenSearchAccountRequest
	GetRegionId() *string
}

type CreateOpenSearchAccountRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// account_sec
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
}

func (s CreateOpenSearchAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOpenSearchAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateOpenSearchAccountRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateOpenSearchAccountRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *CreateOpenSearchAccountRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateOpenSearchAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateOpenSearchAccountRequest) SetAccountName(v string) *CreateOpenSearchAccountRequest {
	s.AccountName = &v
	return s
}

func (s *CreateOpenSearchAccountRequest) SetAccountPassword(v string) *CreateOpenSearchAccountRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateOpenSearchAccountRequest) SetDBInstanceName(v string) *CreateOpenSearchAccountRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateOpenSearchAccountRequest) SetRegionId(v string) *CreateOpenSearchAccountRequest {
	s.RegionId = &v
	return s
}

func (s *CreateOpenSearchAccountRequest) Validate() error {
	return dara.Validate(s)
}
