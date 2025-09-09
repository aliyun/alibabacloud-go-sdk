// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachDatabaseAccountsFromUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachDatabaseAccountsFromUserResponseBody
	GetRequestId() *string
	SetResults(v []*DetachDatabaseAccountsFromUserResponseBodyResults) *DetachDatabaseAccountsFromUserResponseBody
	GetResults() []*DetachDatabaseAccountsFromUserResponseBodyResults
}

type DetachDatabaseAccountsFromUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 79D7E114-CB52-5695-AB15-12776C308387
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*DetachDatabaseAccountsFromUserResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s DetachDatabaseAccountsFromUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachDatabaseAccountsFromUserResponseBody) GoString() string {
	return s.String()
}

func (s *DetachDatabaseAccountsFromUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachDatabaseAccountsFromUserResponseBody) GetResults() []*DetachDatabaseAccountsFromUserResponseBodyResults {
	return s.Results
}

func (s *DetachDatabaseAccountsFromUserResponseBody) SetRequestId(v string) *DetachDatabaseAccountsFromUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachDatabaseAccountsFromUserResponseBody) SetResults(v []*DetachDatabaseAccountsFromUserResponseBodyResults) *DetachDatabaseAccountsFromUserResponseBody {
	s.Results = v
	return s
}

func (s *DetachDatabaseAccountsFromUserResponseBody) Validate() error {
	return dara.Validate(s)
}

type DetachDatabaseAccountsFromUserResponseBodyResults struct {
	// The error code that is returned. If **OK*	- is returned, the operation was successful. If another error code is returned, the operation failed.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// A list that shows the operation results of the database accounts.
	DatabaseAccounts []*DetachDatabaseAccountsFromUserResponseBodyResultsDatabaseAccounts `json:"DatabaseAccounts,omitempty" xml:"DatabaseAccounts,omitempty" type:"Repeated"`
	// The ID of the database on which the permissions are revoked.
	//
	// example:
	//
	// 4
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// The error message that is returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 1
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DetachDatabaseAccountsFromUserResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s DetachDatabaseAccountsFromUserResponseBodyResults) GoString() string {
	return s.String()
}

func (s *DetachDatabaseAccountsFromUserResponseBodyResults) GetCode() *string {
	return s.Code
}

func (s *DetachDatabaseAccountsFromUserResponseBodyResults) GetDatabaseAccounts() []*DetachDatabaseAccountsFromUserResponseBodyResultsDatabaseAccounts {
	return s.DatabaseAccounts
}

func (s *DetachDatabaseAccountsFromUserResponseBodyResults) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *DetachDatabaseAccountsFromUserResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *DetachDatabaseAccountsFromUserResponseBodyResults) GetUserId() *string {
	return s.UserId
}

func (s *DetachDatabaseAccountsFromUserResponseBodyResults) SetCode(v string) *DetachDatabaseAccountsFromUserResponseBodyResults {
	s.Code = &v
	return s
}

func (s *DetachDatabaseAccountsFromUserResponseBodyResults) SetDatabaseAccounts(v []*DetachDatabaseAccountsFromUserResponseBodyResultsDatabaseAccounts) *DetachDatabaseAccountsFromUserResponseBodyResults {
	s.DatabaseAccounts = v
	return s
}

func (s *DetachDatabaseAccountsFromUserResponseBodyResults) SetDatabaseId(v string) *DetachDatabaseAccountsFromUserResponseBodyResults {
	s.DatabaseId = &v
	return s
}

func (s *DetachDatabaseAccountsFromUserResponseBodyResults) SetMessage(v string) *DetachDatabaseAccountsFromUserResponseBodyResults {
	s.Message = &v
	return s
}

func (s *DetachDatabaseAccountsFromUserResponseBodyResults) SetUserId(v string) *DetachDatabaseAccountsFromUserResponseBodyResults {
	s.UserId = &v
	return s
}

func (s *DetachDatabaseAccountsFromUserResponseBodyResults) Validate() error {
	return dara.Validate(s)
}

type DetachDatabaseAccountsFromUserResponseBodyResultsDatabaseAccounts struct {
	// The error code that is returned. If OK is returned, the operation was successful. If another error code is returned, the operation failed.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the database account on which the permissions are revoked.
	//
	// example:
	//
	// 9
	DatabaseAccountId *string `json:"DatabaseAccountId,omitempty" xml:"DatabaseAccountId,omitempty"`
	// The error message that is returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DetachDatabaseAccountsFromUserResponseBodyResultsDatabaseAccounts) String() string {
	return dara.Prettify(s)
}

func (s DetachDatabaseAccountsFromUserResponseBodyResultsDatabaseAccounts) GoString() string {
	return s.String()
}

func (s *DetachDatabaseAccountsFromUserResponseBodyResultsDatabaseAccounts) GetCode() *string {
	return s.Code
}

func (s *DetachDatabaseAccountsFromUserResponseBodyResultsDatabaseAccounts) GetDatabaseAccountId() *string {
	return s.DatabaseAccountId
}

func (s *DetachDatabaseAccountsFromUserResponseBodyResultsDatabaseAccounts) GetMessage() *string {
	return s.Message
}

func (s *DetachDatabaseAccountsFromUserResponseBodyResultsDatabaseAccounts) SetCode(v string) *DetachDatabaseAccountsFromUserResponseBodyResultsDatabaseAccounts {
	s.Code = &v
	return s
}

func (s *DetachDatabaseAccountsFromUserResponseBodyResultsDatabaseAccounts) SetDatabaseAccountId(v string) *DetachDatabaseAccountsFromUserResponseBodyResultsDatabaseAccounts {
	s.DatabaseAccountId = &v
	return s
}

func (s *DetachDatabaseAccountsFromUserResponseBodyResultsDatabaseAccounts) SetMessage(v string) *DetachDatabaseAccountsFromUserResponseBodyResultsDatabaseAccounts {
	s.Message = &v
	return s
}

func (s *DetachDatabaseAccountsFromUserResponseBodyResultsDatabaseAccounts) Validate() error {
	return dara.Validate(s)
}
