// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachDatabaseAccountsFromUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachDatabaseAccountsFromUserGroupResponseBody
	GetRequestId() *string
	SetResults(v []*DetachDatabaseAccountsFromUserGroupResponseBodyResults) *DetachDatabaseAccountsFromUserGroupResponseBody
	GetResults() []*DetachDatabaseAccountsFromUserGroupResponseBodyResults
}

type DetachDatabaseAccountsFromUserGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// AC528ED1-C302-56E5-9CB5-ADA625D64FF9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*DetachDatabaseAccountsFromUserGroupResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s DetachDatabaseAccountsFromUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachDatabaseAccountsFromUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DetachDatabaseAccountsFromUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachDatabaseAccountsFromUserGroupResponseBody) GetResults() []*DetachDatabaseAccountsFromUserGroupResponseBodyResults {
	return s.Results
}

func (s *DetachDatabaseAccountsFromUserGroupResponseBody) SetRequestId(v string) *DetachDatabaseAccountsFromUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachDatabaseAccountsFromUserGroupResponseBody) SetResults(v []*DetachDatabaseAccountsFromUserGroupResponseBodyResults) *DetachDatabaseAccountsFromUserGroupResponseBody {
	s.Results = v
	return s
}

func (s *DetachDatabaseAccountsFromUserGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type DetachDatabaseAccountsFromUserGroupResponseBodyResults struct {
	// The error code that is returned. If OK is returned, the operation was successful. If other error codes are returned, the operation failed.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// A list that shows the authorization results of the database accounts.
	DatabaseAccounts []*DetachDatabaseAccountsFromUserGroupResponseBodyResultsDatabaseAccounts `json:"DatabaseAccounts,omitempty" xml:"DatabaseAccounts,omitempty" type:"Repeated"`
	// The ID of the database on which the permissions are revoked.
	//
	// example:
	//
	// 27
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// The error message that is returned.
	//
	// example:
	//
	// N/A
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The user group ID.
	//
	// example:
	//
	// 3
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s DetachDatabaseAccountsFromUserGroupResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s DetachDatabaseAccountsFromUserGroupResponseBodyResults) GoString() string {
	return s.String()
}

func (s *DetachDatabaseAccountsFromUserGroupResponseBodyResults) GetCode() *string {
	return s.Code
}

func (s *DetachDatabaseAccountsFromUserGroupResponseBodyResults) GetDatabaseAccounts() []*DetachDatabaseAccountsFromUserGroupResponseBodyResultsDatabaseAccounts {
	return s.DatabaseAccounts
}

func (s *DetachDatabaseAccountsFromUserGroupResponseBodyResults) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *DetachDatabaseAccountsFromUserGroupResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *DetachDatabaseAccountsFromUserGroupResponseBodyResults) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *DetachDatabaseAccountsFromUserGroupResponseBodyResults) SetCode(v string) *DetachDatabaseAccountsFromUserGroupResponseBodyResults {
	s.Code = &v
	return s
}

func (s *DetachDatabaseAccountsFromUserGroupResponseBodyResults) SetDatabaseAccounts(v []*DetachDatabaseAccountsFromUserGroupResponseBodyResultsDatabaseAccounts) *DetachDatabaseAccountsFromUserGroupResponseBodyResults {
	s.DatabaseAccounts = v
	return s
}

func (s *DetachDatabaseAccountsFromUserGroupResponseBodyResults) SetDatabaseId(v string) *DetachDatabaseAccountsFromUserGroupResponseBodyResults {
	s.DatabaseId = &v
	return s
}

func (s *DetachDatabaseAccountsFromUserGroupResponseBodyResults) SetMessage(v string) *DetachDatabaseAccountsFromUserGroupResponseBodyResults {
	s.Message = &v
	return s
}

func (s *DetachDatabaseAccountsFromUserGroupResponseBodyResults) SetUserGroupId(v string) *DetachDatabaseAccountsFromUserGroupResponseBodyResults {
	s.UserGroupId = &v
	return s
}

func (s *DetachDatabaseAccountsFromUserGroupResponseBodyResults) Validate() error {
	return dara.Validate(s)
}

type DetachDatabaseAccountsFromUserGroupResponseBodyResultsDatabaseAccounts struct {
	// The error code that is returned. If OK is returned, the operation was successful. If other error codes are returned, the operation failed.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the database account on which the permissions are revoked.
	//
	// example:
	//
	// 5
	DatabaseAccountId *string `json:"DatabaseAccountId,omitempty" xml:"DatabaseAccountId,omitempty"`
	// The error message that is returned.
	//
	// example:
	//
	// N/A
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DetachDatabaseAccountsFromUserGroupResponseBodyResultsDatabaseAccounts) String() string {
	return dara.Prettify(s)
}

func (s DetachDatabaseAccountsFromUserGroupResponseBodyResultsDatabaseAccounts) GoString() string {
	return s.String()
}

func (s *DetachDatabaseAccountsFromUserGroupResponseBodyResultsDatabaseAccounts) GetCode() *string {
	return s.Code
}

func (s *DetachDatabaseAccountsFromUserGroupResponseBodyResultsDatabaseAccounts) GetDatabaseAccountId() *string {
	return s.DatabaseAccountId
}

func (s *DetachDatabaseAccountsFromUserGroupResponseBodyResultsDatabaseAccounts) GetMessage() *string {
	return s.Message
}

func (s *DetachDatabaseAccountsFromUserGroupResponseBodyResultsDatabaseAccounts) SetCode(v string) *DetachDatabaseAccountsFromUserGroupResponseBodyResultsDatabaseAccounts {
	s.Code = &v
	return s
}

func (s *DetachDatabaseAccountsFromUserGroupResponseBodyResultsDatabaseAccounts) SetDatabaseAccountId(v string) *DetachDatabaseAccountsFromUserGroupResponseBodyResultsDatabaseAccounts {
	s.DatabaseAccountId = &v
	return s
}

func (s *DetachDatabaseAccountsFromUserGroupResponseBodyResultsDatabaseAccounts) SetMessage(v string) *DetachDatabaseAccountsFromUserGroupResponseBodyResultsDatabaseAccounts {
	s.Message = &v
	return s
}

func (s *DetachDatabaseAccountsFromUserGroupResponseBodyResultsDatabaseAccounts) Validate() error {
	return dara.Validate(s)
}
