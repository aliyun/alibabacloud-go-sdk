// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachDatabaseAccountsToUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachDatabaseAccountsToUserGroupResponseBody
	GetRequestId() *string
	SetResults(v []*AttachDatabaseAccountsToUserGroupResponseBodyResults) *AttachDatabaseAccountsToUserGroupResponseBody
	GetResults() []*AttachDatabaseAccountsToUserGroupResponseBodyResults
}

type AttachDatabaseAccountsToUserGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5D0EB759-CB0A-537D-A2CC-13A9854FA08D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*AttachDatabaseAccountsToUserGroupResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s AttachDatabaseAccountsToUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachDatabaseAccountsToUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AttachDatabaseAccountsToUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachDatabaseAccountsToUserGroupResponseBody) GetResults() []*AttachDatabaseAccountsToUserGroupResponseBodyResults {
	return s.Results
}

func (s *AttachDatabaseAccountsToUserGroupResponseBody) SetRequestId(v string) *AttachDatabaseAccountsToUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachDatabaseAccountsToUserGroupResponseBody) SetResults(v []*AttachDatabaseAccountsToUserGroupResponseBodyResults) *AttachDatabaseAccountsToUserGroupResponseBody {
	s.Results = v
	return s
}

func (s *AttachDatabaseAccountsToUserGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type AttachDatabaseAccountsToUserGroupResponseBodyResults struct {
	// The error code returned. If OK is returned, the authorization was successful. If another error code is returned, the authorization failed.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// A list that shows the authorization results of the database accounts.
	DatabaseAccounts []*AttachDatabaseAccountsToUserGroupResponseBodyResultsDatabaseAccounts `json:"DatabaseAccounts,omitempty" xml:"DatabaseAccounts,omitempty" type:"Repeated"`
	// The database ID.
	//
	// example:
	//
	// 2
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// The error message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The user group ID.
	//
	// example:
	//
	// 1
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s AttachDatabaseAccountsToUserGroupResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s AttachDatabaseAccountsToUserGroupResponseBodyResults) GoString() string {
	return s.String()
}

func (s *AttachDatabaseAccountsToUserGroupResponseBodyResults) GetCode() *string {
	return s.Code
}

func (s *AttachDatabaseAccountsToUserGroupResponseBodyResults) GetDatabaseAccounts() []*AttachDatabaseAccountsToUserGroupResponseBodyResultsDatabaseAccounts {
	return s.DatabaseAccounts
}

func (s *AttachDatabaseAccountsToUserGroupResponseBodyResults) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *AttachDatabaseAccountsToUserGroupResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *AttachDatabaseAccountsToUserGroupResponseBodyResults) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *AttachDatabaseAccountsToUserGroupResponseBodyResults) SetCode(v string) *AttachDatabaseAccountsToUserGroupResponseBodyResults {
	s.Code = &v
	return s
}

func (s *AttachDatabaseAccountsToUserGroupResponseBodyResults) SetDatabaseAccounts(v []*AttachDatabaseAccountsToUserGroupResponseBodyResultsDatabaseAccounts) *AttachDatabaseAccountsToUserGroupResponseBodyResults {
	s.DatabaseAccounts = v
	return s
}

func (s *AttachDatabaseAccountsToUserGroupResponseBodyResults) SetDatabaseId(v string) *AttachDatabaseAccountsToUserGroupResponseBodyResults {
	s.DatabaseId = &v
	return s
}

func (s *AttachDatabaseAccountsToUserGroupResponseBodyResults) SetMessage(v string) *AttachDatabaseAccountsToUserGroupResponseBodyResults {
	s.Message = &v
	return s
}

func (s *AttachDatabaseAccountsToUserGroupResponseBodyResults) SetUserGroupId(v string) *AttachDatabaseAccountsToUserGroupResponseBodyResults {
	s.UserGroupId = &v
	return s
}

func (s *AttachDatabaseAccountsToUserGroupResponseBodyResults) Validate() error {
	return dara.Validate(s)
}

type AttachDatabaseAccountsToUserGroupResponseBodyResultsDatabaseAccounts struct {
	// The error code returned. If OK is returned, the authorization was successful. If another error code is returned, the authorization failed.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The database account ID.
	//
	// example:
	//
	// 8
	DatabaseAccountId *string `json:"DatabaseAccountId,omitempty" xml:"DatabaseAccountId,omitempty"`
	// The error message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s AttachDatabaseAccountsToUserGroupResponseBodyResultsDatabaseAccounts) String() string {
	return dara.Prettify(s)
}

func (s AttachDatabaseAccountsToUserGroupResponseBodyResultsDatabaseAccounts) GoString() string {
	return s.String()
}

func (s *AttachDatabaseAccountsToUserGroupResponseBodyResultsDatabaseAccounts) GetCode() *string {
	return s.Code
}

func (s *AttachDatabaseAccountsToUserGroupResponseBodyResultsDatabaseAccounts) GetDatabaseAccountId() *string {
	return s.DatabaseAccountId
}

func (s *AttachDatabaseAccountsToUserGroupResponseBodyResultsDatabaseAccounts) GetMessage() *string {
	return s.Message
}

func (s *AttachDatabaseAccountsToUserGroupResponseBodyResultsDatabaseAccounts) SetCode(v string) *AttachDatabaseAccountsToUserGroupResponseBodyResultsDatabaseAccounts {
	s.Code = &v
	return s
}

func (s *AttachDatabaseAccountsToUserGroupResponseBodyResultsDatabaseAccounts) SetDatabaseAccountId(v string) *AttachDatabaseAccountsToUserGroupResponseBodyResultsDatabaseAccounts {
	s.DatabaseAccountId = &v
	return s
}

func (s *AttachDatabaseAccountsToUserGroupResponseBodyResultsDatabaseAccounts) SetMessage(v string) *AttachDatabaseAccountsToUserGroupResponseBodyResultsDatabaseAccounts {
	s.Message = &v
	return s
}

func (s *AttachDatabaseAccountsToUserGroupResponseBodyResultsDatabaseAccounts) Validate() error {
	return dara.Validate(s)
}
