// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabaseAccountsForUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseAccounts(v []*ListDatabaseAccountsForUserResponseBodyDatabaseAccounts) *ListDatabaseAccountsForUserResponseBody
	GetDatabaseAccounts() []*ListDatabaseAccountsForUserResponseBodyDatabaseAccounts
	SetRequestId(v string) *ListDatabaseAccountsForUserResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListDatabaseAccountsForUserResponseBody
	GetTotalCount() *int64
}

type ListDatabaseAccountsForUserResponseBody struct {
	// The database accounts returned.
	DatabaseAccounts []*ListDatabaseAccountsForUserResponseBodyDatabaseAccounts `json:"DatabaseAccounts,omitempty" xml:"DatabaseAccounts,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 00E3701B-3616-55FE-93EC-E7CF5480B654
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of database accounts that are returned.
	//
	// example:
	//
	// 3
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDatabaseAccountsForUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDatabaseAccountsForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatabaseAccountsForUserResponseBody) GetDatabaseAccounts() []*ListDatabaseAccountsForUserResponseBodyDatabaseAccounts {
	return s.DatabaseAccounts
}

func (s *ListDatabaseAccountsForUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDatabaseAccountsForUserResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDatabaseAccountsForUserResponseBody) SetDatabaseAccounts(v []*ListDatabaseAccountsForUserResponseBodyDatabaseAccounts) *ListDatabaseAccountsForUserResponseBody {
	s.DatabaseAccounts = v
	return s
}

func (s *ListDatabaseAccountsForUserResponseBody) SetRequestId(v string) *ListDatabaseAccountsForUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatabaseAccountsForUserResponseBody) SetTotalCount(v int64) *ListDatabaseAccountsForUserResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDatabaseAccountsForUserResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDatabaseAccountsForUserResponseBodyDatabaseAccounts struct {
	// The database account ID.
	//
	// example:
	//
	// 6
	DatabaseAccountId *string `json:"DatabaseAccountId,omitempty" xml:"DatabaseAccountId,omitempty"`
	// The name of the database account.
	//
	// example:
	//
	// test
	DatabaseAccountName *string `json:"DatabaseAccountName,omitempty" xml:"DatabaseAccountName,omitempty"`
	// The ID of the database to which the database account belongs.
	//
	// example:
	//
	// 70
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// Indicates whether the user is authorized to manage the database account. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsAuthorized *bool `json:"IsAuthorized,omitempty" xml:"IsAuthorized,omitempty"`
	// The protocol used by the database account. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **Oracle**
	//
	// 	- **PostgreSQL**
	//
	// 	- **SQLServer**
	//
	// example:
	//
	// MySQL
	ProtocolName *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
}

func (s ListDatabaseAccountsForUserResponseBodyDatabaseAccounts) String() string {
	return dara.Prettify(s)
}

func (s ListDatabaseAccountsForUserResponseBodyDatabaseAccounts) GoString() string {
	return s.String()
}

func (s *ListDatabaseAccountsForUserResponseBodyDatabaseAccounts) GetDatabaseAccountId() *string {
	return s.DatabaseAccountId
}

func (s *ListDatabaseAccountsForUserResponseBodyDatabaseAccounts) GetDatabaseAccountName() *string {
	return s.DatabaseAccountName
}

func (s *ListDatabaseAccountsForUserResponseBodyDatabaseAccounts) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *ListDatabaseAccountsForUserResponseBodyDatabaseAccounts) GetIsAuthorized() *bool {
	return s.IsAuthorized
}

func (s *ListDatabaseAccountsForUserResponseBodyDatabaseAccounts) GetProtocolName() *string {
	return s.ProtocolName
}

func (s *ListDatabaseAccountsForUserResponseBodyDatabaseAccounts) SetDatabaseAccountId(v string) *ListDatabaseAccountsForUserResponseBodyDatabaseAccounts {
	s.DatabaseAccountId = &v
	return s
}

func (s *ListDatabaseAccountsForUserResponseBodyDatabaseAccounts) SetDatabaseAccountName(v string) *ListDatabaseAccountsForUserResponseBodyDatabaseAccounts {
	s.DatabaseAccountName = &v
	return s
}

func (s *ListDatabaseAccountsForUserResponseBodyDatabaseAccounts) SetDatabaseId(v string) *ListDatabaseAccountsForUserResponseBodyDatabaseAccounts {
	s.DatabaseId = &v
	return s
}

func (s *ListDatabaseAccountsForUserResponseBodyDatabaseAccounts) SetIsAuthorized(v bool) *ListDatabaseAccountsForUserResponseBodyDatabaseAccounts {
	s.IsAuthorized = &v
	return s
}

func (s *ListDatabaseAccountsForUserResponseBodyDatabaseAccounts) SetProtocolName(v string) *ListDatabaseAccountsForUserResponseBodyDatabaseAccounts {
	s.ProtocolName = &v
	return s
}

func (s *ListDatabaseAccountsForUserResponseBodyDatabaseAccounts) Validate() error {
	return dara.Validate(s)
}
