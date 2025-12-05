// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatabaseAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseAccount(v *GetDatabaseAccountResponseBodyDatabaseAccount) *GetDatabaseAccountResponseBody
	GetDatabaseAccount() *GetDatabaseAccountResponseBodyDatabaseAccount
	SetRequestId(v string) *GetDatabaseAccountResponseBody
	GetRequestId() *string
}

type GetDatabaseAccountResponseBody struct {
	// The returned information about the database account.
	DatabaseAccount *GetDatabaseAccountResponseBodyDatabaseAccount `json:"DatabaseAccount,omitempty" xml:"DatabaseAccount,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// FA06D274-8D0A-59FB-8B7E-584C0EEBBFFF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDatabaseAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDatabaseAccountResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatabaseAccountResponseBody) GetDatabaseAccount() *GetDatabaseAccountResponseBodyDatabaseAccount {
	return s.DatabaseAccount
}

func (s *GetDatabaseAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDatabaseAccountResponseBody) SetDatabaseAccount(v *GetDatabaseAccountResponseBodyDatabaseAccount) *GetDatabaseAccountResponseBody {
	s.DatabaseAccount = v
	return s
}

func (s *GetDatabaseAccountResponseBody) SetRequestId(v string) *GetDatabaseAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDatabaseAccountResponseBody) Validate() error {
	if s.DatabaseAccount != nil {
		if err := s.DatabaseAccount.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDatabaseAccountResponseBodyDatabaseAccount struct {
	// The database account ID.
	//
	// example:
	//
	// 9
	DatabaseAccountId *string `json:"DatabaseAccountId,omitempty" xml:"DatabaseAccountId,omitempty"`
	// The name of the database account.
	//
	// example:
	//
	// uac
	DatabaseAccountName *string `json:"DatabaseAccountName,omitempty" xml:"DatabaseAccountName,omitempty"`
	// The database name. A value is returned for this parameter if the database engine is PostgreSQL or Oracle.
	//
	// example:
	//
	// orcl
	DatabaseSchema *string `json:"DatabaseSchema,omitempty" xml:"DatabaseSchema,omitempty"`
	// Indicates whether the database account has a password.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	HasPassword *bool `json:"HasPassword,omitempty" xml:"HasPassword,omitempty"`
	// The logon attribute. A value is returned for this parameter if the database engine is Oracle. Valid values:
	//
	// 	- SERVICENAME
	//
	// 	- SID
	//
	// example:
	//
	// SID
	LoginAttribute *string `json:"LoginAttribute,omitempty" xml:"LoginAttribute,omitempty"`
}

func (s GetDatabaseAccountResponseBodyDatabaseAccount) String() string {
	return dara.Prettify(s)
}

func (s GetDatabaseAccountResponseBodyDatabaseAccount) GoString() string {
	return s.String()
}

func (s *GetDatabaseAccountResponseBodyDatabaseAccount) GetDatabaseAccountId() *string {
	return s.DatabaseAccountId
}

func (s *GetDatabaseAccountResponseBodyDatabaseAccount) GetDatabaseAccountName() *string {
	return s.DatabaseAccountName
}

func (s *GetDatabaseAccountResponseBodyDatabaseAccount) GetDatabaseSchema() *string {
	return s.DatabaseSchema
}

func (s *GetDatabaseAccountResponseBodyDatabaseAccount) GetHasPassword() *bool {
	return s.HasPassword
}

func (s *GetDatabaseAccountResponseBodyDatabaseAccount) GetLoginAttribute() *string {
	return s.LoginAttribute
}

func (s *GetDatabaseAccountResponseBodyDatabaseAccount) SetDatabaseAccountId(v string) *GetDatabaseAccountResponseBodyDatabaseAccount {
	s.DatabaseAccountId = &v
	return s
}

func (s *GetDatabaseAccountResponseBodyDatabaseAccount) SetDatabaseAccountName(v string) *GetDatabaseAccountResponseBodyDatabaseAccount {
	s.DatabaseAccountName = &v
	return s
}

func (s *GetDatabaseAccountResponseBodyDatabaseAccount) SetDatabaseSchema(v string) *GetDatabaseAccountResponseBodyDatabaseAccount {
	s.DatabaseSchema = &v
	return s
}

func (s *GetDatabaseAccountResponseBodyDatabaseAccount) SetHasPassword(v bool) *GetDatabaseAccountResponseBodyDatabaseAccount {
	s.HasPassword = &v
	return s
}

func (s *GetDatabaseAccountResponseBodyDatabaseAccount) SetLoginAttribute(v string) *GetDatabaseAccountResponseBodyDatabaseAccount {
	s.LoginAttribute = &v
	return s
}

func (s *GetDatabaseAccountResponseBodyDatabaseAccount) Validate() error {
	return dara.Validate(s)
}
