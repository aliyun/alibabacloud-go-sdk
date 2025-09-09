// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *CreateInstanceAccountRequest
	GetAccountName() *string
	SetDbPrivilege(v []*CreateInstanceAccountRequestDbPrivilege) *CreateInstanceAccountRequest
	GetDbPrivilege() []*CreateInstanceAccountRequestDbPrivilege
	SetDrdsInstanceId(v string) *CreateInstanceAccountRequest
	GetDrdsInstanceId() *string
	SetPassword(v string) *CreateInstanceAccountRequest
	GetPassword() *string
}

type CreateInstanceAccountRequest struct {
	// The username of the account you want to create.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds_sample_account
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// This parameter is required.
	DbPrivilege []*CreateInstanceAccountRequestDbPrivilege `json:"DbPrivilege,omitempty" xml:"DbPrivilege,omitempty" type:"Repeated"`
	// The ID of the PolarDB-X 1.0 instance for which you want to create the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// drdsjiii1b49****
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The password of the account you want to create.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds_sample_password
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s CreateInstanceAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceAccountRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateInstanceAccountRequest) GetDbPrivilege() []*CreateInstanceAccountRequestDbPrivilege {
	return s.DbPrivilege
}

func (s *CreateInstanceAccountRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *CreateInstanceAccountRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateInstanceAccountRequest) SetAccountName(v string) *CreateInstanceAccountRequest {
	s.AccountName = &v
	return s
}

func (s *CreateInstanceAccountRequest) SetDbPrivilege(v []*CreateInstanceAccountRequestDbPrivilege) *CreateInstanceAccountRequest {
	s.DbPrivilege = v
	return s
}

func (s *CreateInstanceAccountRequest) SetDrdsInstanceId(v string) *CreateInstanceAccountRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *CreateInstanceAccountRequest) SetPassword(v string) *CreateInstanceAccountRequest {
	s.Password = &v
	return s
}

func (s *CreateInstanceAccountRequest) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceAccountRequestDbPrivilege struct {
	// The name of the database that you want to manage by using the account to create.
	//
	// example:
	//
	// test123
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The permissions that you want to grant to the account to manage the database.
	//
	// example:
	//
	// DDL
	Privilege *string `json:"Privilege,omitempty" xml:"Privilege,omitempty"`
}

func (s CreateInstanceAccountRequestDbPrivilege) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceAccountRequestDbPrivilege) GoString() string {
	return s.String()
}

func (s *CreateInstanceAccountRequestDbPrivilege) GetDbName() *string {
	return s.DbName
}

func (s *CreateInstanceAccountRequestDbPrivilege) GetPrivilege() *string {
	return s.Privilege
}

func (s *CreateInstanceAccountRequestDbPrivilege) SetDbName(v string) *CreateInstanceAccountRequestDbPrivilege {
	s.DbName = &v
	return s
}

func (s *CreateInstanceAccountRequestDbPrivilege) SetPrivilege(v string) *CreateInstanceAccountRequestDbPrivilege {
	s.Privilege = &v
	return s
}

func (s *CreateInstanceAccountRequestDbPrivilege) Validate() error {
	return dara.Validate(s)
}
