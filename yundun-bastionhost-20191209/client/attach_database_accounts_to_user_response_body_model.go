// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachDatabaseAccountsToUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachDatabaseAccountsToUserResponseBody
	GetRequestId() *string
	SetResults(v []*AttachDatabaseAccountsToUserResponseBodyResults) *AttachDatabaseAccountsToUserResponseBody
	GetResults() []*AttachDatabaseAccountsToUserResponseBodyResults
}

type AttachDatabaseAccountsToUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 23120B8E-8737-50BD-A3A3-902A7821F04D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*AttachDatabaseAccountsToUserResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s AttachDatabaseAccountsToUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachDatabaseAccountsToUserResponseBody) GoString() string {
	return s.String()
}

func (s *AttachDatabaseAccountsToUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachDatabaseAccountsToUserResponseBody) GetResults() []*AttachDatabaseAccountsToUserResponseBodyResults {
	return s.Results
}

func (s *AttachDatabaseAccountsToUserResponseBody) SetRequestId(v string) *AttachDatabaseAccountsToUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachDatabaseAccountsToUserResponseBody) SetResults(v []*AttachDatabaseAccountsToUserResponseBodyResults) *AttachDatabaseAccountsToUserResponseBody {
	s.Results = v
	return s
}

func (s *AttachDatabaseAccountsToUserResponseBody) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AttachDatabaseAccountsToUserResponseBodyResults struct {
	// The error code that is returned. If **OK*	- is returned, the authorization was successful. If another error code is returned, the authorization failed.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// A list that shows the authorization results of the database accounts.
	DatabaseAccounts []*AttachDatabaseAccountsToUserResponseBodyResultsDatabaseAccounts `json:"DatabaseAccounts,omitempty" xml:"DatabaseAccounts,omitempty" type:"Repeated"`
	// The database ID.
	//
	// example:
	//
	// 22
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

func (s AttachDatabaseAccountsToUserResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s AttachDatabaseAccountsToUserResponseBodyResults) GoString() string {
	return s.String()
}

func (s *AttachDatabaseAccountsToUserResponseBodyResults) GetCode() *string {
	return s.Code
}

func (s *AttachDatabaseAccountsToUserResponseBodyResults) GetDatabaseAccounts() []*AttachDatabaseAccountsToUserResponseBodyResultsDatabaseAccounts {
	return s.DatabaseAccounts
}

func (s *AttachDatabaseAccountsToUserResponseBodyResults) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *AttachDatabaseAccountsToUserResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *AttachDatabaseAccountsToUserResponseBodyResults) GetUserId() *string {
	return s.UserId
}

func (s *AttachDatabaseAccountsToUserResponseBodyResults) SetCode(v string) *AttachDatabaseAccountsToUserResponseBodyResults {
	s.Code = &v
	return s
}

func (s *AttachDatabaseAccountsToUserResponseBodyResults) SetDatabaseAccounts(v []*AttachDatabaseAccountsToUserResponseBodyResultsDatabaseAccounts) *AttachDatabaseAccountsToUserResponseBodyResults {
	s.DatabaseAccounts = v
	return s
}

func (s *AttachDatabaseAccountsToUserResponseBodyResults) SetDatabaseId(v string) *AttachDatabaseAccountsToUserResponseBodyResults {
	s.DatabaseId = &v
	return s
}

func (s *AttachDatabaseAccountsToUserResponseBodyResults) SetMessage(v string) *AttachDatabaseAccountsToUserResponseBodyResults {
	s.Message = &v
	return s
}

func (s *AttachDatabaseAccountsToUserResponseBodyResults) SetUserId(v string) *AttachDatabaseAccountsToUserResponseBodyResults {
	s.UserId = &v
	return s
}

func (s *AttachDatabaseAccountsToUserResponseBodyResults) Validate() error {
	if s.DatabaseAccounts != nil {
		for _, item := range s.DatabaseAccounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AttachDatabaseAccountsToUserResponseBodyResultsDatabaseAccounts struct {
	// The error code that is returned. If OK is returned, the authorization was successful. If another error code is returned, the authorization failed.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The database account ID.
	//
	// example:
	//
	// 6
	DatabaseAccountId *string `json:"DatabaseAccountId,omitempty" xml:"DatabaseAccountId,omitempty"`
	// The error message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s AttachDatabaseAccountsToUserResponseBodyResultsDatabaseAccounts) String() string {
	return dara.Prettify(s)
}

func (s AttachDatabaseAccountsToUserResponseBodyResultsDatabaseAccounts) GoString() string {
	return s.String()
}

func (s *AttachDatabaseAccountsToUserResponseBodyResultsDatabaseAccounts) GetCode() *string {
	return s.Code
}

func (s *AttachDatabaseAccountsToUserResponseBodyResultsDatabaseAccounts) GetDatabaseAccountId() *string {
	return s.DatabaseAccountId
}

func (s *AttachDatabaseAccountsToUserResponseBodyResultsDatabaseAccounts) GetMessage() *string {
	return s.Message
}

func (s *AttachDatabaseAccountsToUserResponseBodyResultsDatabaseAccounts) SetCode(v string) *AttachDatabaseAccountsToUserResponseBodyResultsDatabaseAccounts {
	s.Code = &v
	return s
}

func (s *AttachDatabaseAccountsToUserResponseBodyResultsDatabaseAccounts) SetDatabaseAccountId(v string) *AttachDatabaseAccountsToUserResponseBodyResultsDatabaseAccounts {
	s.DatabaseAccountId = &v
	return s
}

func (s *AttachDatabaseAccountsToUserResponseBodyResultsDatabaseAccounts) SetMessage(v string) *AttachDatabaseAccountsToUserResponseBodyResultsDatabaseAccounts {
	s.Message = &v
	return s
}

func (s *AttachDatabaseAccountsToUserResponseBodyResultsDatabaseAccounts) Validate() error {
	return dara.Validate(s)
}
