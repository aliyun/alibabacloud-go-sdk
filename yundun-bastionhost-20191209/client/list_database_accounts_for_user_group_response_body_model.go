// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabaseAccountsForUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseAccounts(v []*ListDatabaseAccountsForUserGroupResponseBodyDatabaseAccounts) *ListDatabaseAccountsForUserGroupResponseBody
	GetDatabaseAccounts() []*ListDatabaseAccountsForUserGroupResponseBodyDatabaseAccounts
	SetRequestId(v string) *ListDatabaseAccountsForUserGroupResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListDatabaseAccountsForUserGroupResponseBody
	GetTotalCount() *int64
}

type ListDatabaseAccountsForUserGroupResponseBody struct {
	// The database accounts returned.
	DatabaseAccounts []*ListDatabaseAccountsForUserGroupResponseBodyDatabaseAccounts `json:"DatabaseAccounts,omitempty" xml:"DatabaseAccounts,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of database accounts returned.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDatabaseAccountsForUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDatabaseAccountsForUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatabaseAccountsForUserGroupResponseBody) GetDatabaseAccounts() []*ListDatabaseAccountsForUserGroupResponseBodyDatabaseAccounts {
	return s.DatabaseAccounts
}

func (s *ListDatabaseAccountsForUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDatabaseAccountsForUserGroupResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDatabaseAccountsForUserGroupResponseBody) SetDatabaseAccounts(v []*ListDatabaseAccountsForUserGroupResponseBodyDatabaseAccounts) *ListDatabaseAccountsForUserGroupResponseBody {
	s.DatabaseAccounts = v
	return s
}

func (s *ListDatabaseAccountsForUserGroupResponseBody) SetRequestId(v string) *ListDatabaseAccountsForUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatabaseAccountsForUserGroupResponseBody) SetTotalCount(v int64) *ListDatabaseAccountsForUserGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDatabaseAccountsForUserGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDatabaseAccountsForUserGroupResponseBodyDatabaseAccounts struct {
	// The ID of the database account.
	//
	// example:
	//
	// 4
	DatabaseAccountId *string `json:"DatabaseAccountId,omitempty" xml:"DatabaseAccountId,omitempty"`
	// The name of the database account.
	//
	// example:
	//
	// root
	DatabaseAccountName *string `json:"DatabaseAccountName,omitempty" xml:"DatabaseAccountName,omitempty"`
	// The ID of the database to which the database account belongs.
	//
	// example:
	//
	// 11
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// Indicates whether the user group is authorized to manage the database account. Valid values:
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

func (s ListDatabaseAccountsForUserGroupResponseBodyDatabaseAccounts) String() string {
	return dara.Prettify(s)
}

func (s ListDatabaseAccountsForUserGroupResponseBodyDatabaseAccounts) GoString() string {
	return s.String()
}

func (s *ListDatabaseAccountsForUserGroupResponseBodyDatabaseAccounts) GetDatabaseAccountId() *string {
	return s.DatabaseAccountId
}

func (s *ListDatabaseAccountsForUserGroupResponseBodyDatabaseAccounts) GetDatabaseAccountName() *string {
	return s.DatabaseAccountName
}

func (s *ListDatabaseAccountsForUserGroupResponseBodyDatabaseAccounts) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *ListDatabaseAccountsForUserGroupResponseBodyDatabaseAccounts) GetIsAuthorized() *bool {
	return s.IsAuthorized
}

func (s *ListDatabaseAccountsForUserGroupResponseBodyDatabaseAccounts) GetProtocolName() *string {
	return s.ProtocolName
}

func (s *ListDatabaseAccountsForUserGroupResponseBodyDatabaseAccounts) SetDatabaseAccountId(v string) *ListDatabaseAccountsForUserGroupResponseBodyDatabaseAccounts {
	s.DatabaseAccountId = &v
	return s
}

func (s *ListDatabaseAccountsForUserGroupResponseBodyDatabaseAccounts) SetDatabaseAccountName(v string) *ListDatabaseAccountsForUserGroupResponseBodyDatabaseAccounts {
	s.DatabaseAccountName = &v
	return s
}

func (s *ListDatabaseAccountsForUserGroupResponseBodyDatabaseAccounts) SetDatabaseId(v string) *ListDatabaseAccountsForUserGroupResponseBodyDatabaseAccounts {
	s.DatabaseId = &v
	return s
}

func (s *ListDatabaseAccountsForUserGroupResponseBodyDatabaseAccounts) SetIsAuthorized(v bool) *ListDatabaseAccountsForUserGroupResponseBodyDatabaseAccounts {
	s.IsAuthorized = &v
	return s
}

func (s *ListDatabaseAccountsForUserGroupResponseBodyDatabaseAccounts) SetProtocolName(v string) *ListDatabaseAccountsForUserGroupResponseBodyDatabaseAccounts {
	s.ProtocolName = &v
	return s
}

func (s *ListDatabaseAccountsForUserGroupResponseBodyDatabaseAccounts) Validate() error {
	return dara.Validate(s)
}
