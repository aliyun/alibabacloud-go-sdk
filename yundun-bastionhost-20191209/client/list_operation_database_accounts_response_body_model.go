// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationDatabaseAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseAccounts(v []*ListOperationDatabaseAccountsResponseBodyDatabaseAccounts) *ListOperationDatabaseAccountsResponseBody
	GetDatabaseAccounts() []*ListOperationDatabaseAccountsResponseBodyDatabaseAccounts
	SetRequestId(v string) *ListOperationDatabaseAccountsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListOperationDatabaseAccountsResponseBody
	GetTotalCount() *int64
}

type ListOperationDatabaseAccountsResponseBody struct {
	// The database accounts returned.
	DatabaseAccounts []*ListOperationDatabaseAccountsResponseBodyDatabaseAccounts `json:"DatabaseAccounts,omitempty" xml:"DatabaseAccounts,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 5
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOperationDatabaseAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOperationDatabaseAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOperationDatabaseAccountsResponseBody) GetDatabaseAccounts() []*ListOperationDatabaseAccountsResponseBodyDatabaseAccounts {
	return s.DatabaseAccounts
}

func (s *ListOperationDatabaseAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOperationDatabaseAccountsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListOperationDatabaseAccountsResponseBody) SetDatabaseAccounts(v []*ListOperationDatabaseAccountsResponseBodyDatabaseAccounts) *ListOperationDatabaseAccountsResponseBody {
	s.DatabaseAccounts = v
	return s
}

func (s *ListOperationDatabaseAccountsResponseBody) SetRequestId(v string) *ListOperationDatabaseAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOperationDatabaseAccountsResponseBody) SetTotalCount(v int64) *ListOperationDatabaseAccountsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOperationDatabaseAccountsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListOperationDatabaseAccountsResponseBodyDatabaseAccounts struct {
	// The name of the PostgreSQL or Oracle database.
	//
	// example:
	//
	// xe
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The database account ID.
	//
	// example:
	//
	// 3
	DatabaseAccountId *string `json:"DatabaseAccountId,omitempty" xml:"DatabaseAccountId,omitempty"`
	// The name of the database account.
	//
	// example:
	//
	// system
	DatabaseAccountName *string `json:"DatabaseAccountName,omitempty" xml:"DatabaseAccountName,omitempty"`
	// The database ID.
	//
	// example:
	//
	// 2
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// Indicates whether a password is configured for the database host account.
	//
	// example:
	//
	// true
	HasPassword *string `json:"HasPassword,omitempty" xml:"HasPassword,omitempty"`
	// The logon attribute. One of the following values is returned if the database engine is Oracle:
	//
	// 	- **SERVICENAME**
	//
	// 	- **SID**
	//
	// example:
	//
	// SID
	LoginAttribute *string `json:"LoginAttribute,omitempty" xml:"LoginAttribute,omitempty"`
	// The protocol that is used by the database account.
	//
	// example:
	//
	// MySQL
	ProtocolName *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
}

func (s ListOperationDatabaseAccountsResponseBodyDatabaseAccounts) String() string {
	return dara.Prettify(s)
}

func (s ListOperationDatabaseAccountsResponseBodyDatabaseAccounts) GoString() string {
	return s.String()
}

func (s *ListOperationDatabaseAccountsResponseBodyDatabaseAccounts) GetDBName() *string {
	return s.DBName
}

func (s *ListOperationDatabaseAccountsResponseBodyDatabaseAccounts) GetDatabaseAccountId() *string {
	return s.DatabaseAccountId
}

func (s *ListOperationDatabaseAccountsResponseBodyDatabaseAccounts) GetDatabaseAccountName() *string {
	return s.DatabaseAccountName
}

func (s *ListOperationDatabaseAccountsResponseBodyDatabaseAccounts) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *ListOperationDatabaseAccountsResponseBodyDatabaseAccounts) GetHasPassword() *string {
	return s.HasPassword
}

func (s *ListOperationDatabaseAccountsResponseBodyDatabaseAccounts) GetLoginAttribute() *string {
	return s.LoginAttribute
}

func (s *ListOperationDatabaseAccountsResponseBodyDatabaseAccounts) GetProtocolName() *string {
	return s.ProtocolName
}

func (s *ListOperationDatabaseAccountsResponseBodyDatabaseAccounts) SetDBName(v string) *ListOperationDatabaseAccountsResponseBodyDatabaseAccounts {
	s.DBName = &v
	return s
}

func (s *ListOperationDatabaseAccountsResponseBodyDatabaseAccounts) SetDatabaseAccountId(v string) *ListOperationDatabaseAccountsResponseBodyDatabaseAccounts {
	s.DatabaseAccountId = &v
	return s
}

func (s *ListOperationDatabaseAccountsResponseBodyDatabaseAccounts) SetDatabaseAccountName(v string) *ListOperationDatabaseAccountsResponseBodyDatabaseAccounts {
	s.DatabaseAccountName = &v
	return s
}

func (s *ListOperationDatabaseAccountsResponseBodyDatabaseAccounts) SetDatabaseId(v string) *ListOperationDatabaseAccountsResponseBodyDatabaseAccounts {
	s.DatabaseId = &v
	return s
}

func (s *ListOperationDatabaseAccountsResponseBodyDatabaseAccounts) SetHasPassword(v string) *ListOperationDatabaseAccountsResponseBodyDatabaseAccounts {
	s.HasPassword = &v
	return s
}

func (s *ListOperationDatabaseAccountsResponseBodyDatabaseAccounts) SetLoginAttribute(v string) *ListOperationDatabaseAccountsResponseBodyDatabaseAccounts {
	s.LoginAttribute = &v
	return s
}

func (s *ListOperationDatabaseAccountsResponseBodyDatabaseAccounts) SetProtocolName(v string) *ListOperationDatabaseAccountsResponseBodyDatabaseAccounts {
	s.ProtocolName = &v
	return s
}

func (s *ListOperationDatabaseAccountsResponseBodyDatabaseAccounts) Validate() error {
	return dara.Validate(s)
}
