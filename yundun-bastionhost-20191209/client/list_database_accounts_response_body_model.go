// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabaseAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseAccounts(v []*ListDatabaseAccountsResponseBodyDatabaseAccounts) *ListDatabaseAccountsResponseBody
	GetDatabaseAccounts() []*ListDatabaseAccountsResponseBodyDatabaseAccounts
	SetRequestId(v string) *ListDatabaseAccountsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListDatabaseAccountsResponseBody
	GetTotalCount() *int64
}

type ListDatabaseAccountsResponseBody struct {
	// The returned database accounts.
	DatabaseAccounts []*ListDatabaseAccountsResponseBodyDatabaseAccounts `json:"DatabaseAccounts,omitempty" xml:"DatabaseAccounts,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 4F6C075F-FC86-476E-943B-097BD4E12948
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of database accounts returned.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDatabaseAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDatabaseAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatabaseAccountsResponseBody) GetDatabaseAccounts() []*ListDatabaseAccountsResponseBodyDatabaseAccounts {
	return s.DatabaseAccounts
}

func (s *ListDatabaseAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDatabaseAccountsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDatabaseAccountsResponseBody) SetDatabaseAccounts(v []*ListDatabaseAccountsResponseBodyDatabaseAccounts) *ListDatabaseAccountsResponseBody {
	s.DatabaseAccounts = v
	return s
}

func (s *ListDatabaseAccountsResponseBody) SetRequestId(v string) *ListDatabaseAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatabaseAccountsResponseBody) SetTotalCount(v int64) *ListDatabaseAccountsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDatabaseAccountsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDatabaseAccountsResponseBodyDatabaseAccounts struct {
	// The database account ID.
	//
	// example:
	//
	// 59
	DatabaseAccountId *string `json:"DatabaseAccountId,omitempty" xml:"DatabaseAccountId,omitempty"`
	// The name of the database account.
	//
	// example:
	//
	// test
	DatabaseAccountName *string `json:"DatabaseAccountName,omitempty" xml:"DatabaseAccountName,omitempty"`
	// The database ID.
	//
	// example:
	//
	// 4
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// The name of the database. A value is returned for this parameter if the engine of the database with the specified database ID is PostgreSQL or Oracle.
	//
	// example:
	//
	// orcl
	DatabaseSchema *string `json:"DatabaseSchema,omitempty" xml:"DatabaseSchema,omitempty"`
	// Indicates whether the database account has a password. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	HasPassword *string `json:"HasPassword,omitempty" xml:"HasPassword,omitempty"`
}

func (s ListDatabaseAccountsResponseBodyDatabaseAccounts) String() string {
	return dara.Prettify(s)
}

func (s ListDatabaseAccountsResponseBodyDatabaseAccounts) GoString() string {
	return s.String()
}

func (s *ListDatabaseAccountsResponseBodyDatabaseAccounts) GetDatabaseAccountId() *string {
	return s.DatabaseAccountId
}

func (s *ListDatabaseAccountsResponseBodyDatabaseAccounts) GetDatabaseAccountName() *string {
	return s.DatabaseAccountName
}

func (s *ListDatabaseAccountsResponseBodyDatabaseAccounts) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *ListDatabaseAccountsResponseBodyDatabaseAccounts) GetDatabaseSchema() *string {
	return s.DatabaseSchema
}

func (s *ListDatabaseAccountsResponseBodyDatabaseAccounts) GetHasPassword() *string {
	return s.HasPassword
}

func (s *ListDatabaseAccountsResponseBodyDatabaseAccounts) SetDatabaseAccountId(v string) *ListDatabaseAccountsResponseBodyDatabaseAccounts {
	s.DatabaseAccountId = &v
	return s
}

func (s *ListDatabaseAccountsResponseBodyDatabaseAccounts) SetDatabaseAccountName(v string) *ListDatabaseAccountsResponseBodyDatabaseAccounts {
	s.DatabaseAccountName = &v
	return s
}

func (s *ListDatabaseAccountsResponseBodyDatabaseAccounts) SetDatabaseId(v string) *ListDatabaseAccountsResponseBodyDatabaseAccounts {
	s.DatabaseId = &v
	return s
}

func (s *ListDatabaseAccountsResponseBodyDatabaseAccounts) SetDatabaseSchema(v string) *ListDatabaseAccountsResponseBodyDatabaseAccounts {
	s.DatabaseSchema = &v
	return s
}

func (s *ListDatabaseAccountsResponseBodyDatabaseAccounts) SetHasPassword(v string) *ListDatabaseAccountsResponseBodyDatabaseAccounts {
	s.HasPassword = &v
	return s
}

func (s *ListDatabaseAccountsResponseBodyDatabaseAccounts) Validate() error {
	return dara.Validate(s)
}
