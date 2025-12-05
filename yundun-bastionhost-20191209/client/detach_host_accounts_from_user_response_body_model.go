// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachHostAccountsFromUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachHostAccountsFromUserResponseBody
	GetRequestId() *string
	SetResults(v []*DetachHostAccountsFromUserResponseBodyResults) *DetachHostAccountsFromUserResponseBody
	GetResults() []*DetachHostAccountsFromUserResponseBodyResults
}

type DetachHostAccountsFromUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*DetachHostAccountsFromUserResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s DetachHostAccountsFromUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachHostAccountsFromUserResponseBody) GoString() string {
	return s.String()
}

func (s *DetachHostAccountsFromUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachHostAccountsFromUserResponseBody) GetResults() []*DetachHostAccountsFromUserResponseBodyResults {
	return s.Results
}

func (s *DetachHostAccountsFromUserResponseBody) SetRequestId(v string) *DetachHostAccountsFromUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachHostAccountsFromUserResponseBody) SetResults(v []*DetachHostAccountsFromUserResponseBodyResults) *DetachHostAccountsFromUserResponseBody {
	s.Results = v
	return s
}

func (s *DetachHostAccountsFromUserResponseBody) Validate() error {
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

type DetachHostAccountsFromUserResponseBodyResults struct {
	// The return code that indicates whether the call was successful. Valid values:
	//
	// 	- **OK**: The call was successful.
	//
	// 	- **UNEXPECTED**: An unknown error occurred.
	//
	// 	- **INVALID_ARGUMENT**: A request parameter is invalid.
	//
	// 	- **OBJECT_NOT_FOUND**: The specified object on which you want to perform the operation does not exist.
	//
	// 	- **OBJECT_AlREADY_EXISTS**: The specified object on which you want to perform the operation already exists.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of revoking permissions on the specified host accounts from the user.
	HostAccounts []*DetachHostAccountsFromUserResponseBodyResultsHostAccounts `json:"HostAccounts,omitempty" xml:"HostAccounts,omitempty" type:"Repeated"`
	// The host ID.
	//
	// example:
	//
	// 1
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// This parameter is deprecated.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 1
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DetachHostAccountsFromUserResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s DetachHostAccountsFromUserResponseBodyResults) GoString() string {
	return s.String()
}

func (s *DetachHostAccountsFromUserResponseBodyResults) GetCode() *string {
	return s.Code
}

func (s *DetachHostAccountsFromUserResponseBodyResults) GetHostAccounts() []*DetachHostAccountsFromUserResponseBodyResultsHostAccounts {
	return s.HostAccounts
}

func (s *DetachHostAccountsFromUserResponseBodyResults) GetHostId() *string {
	return s.HostId
}

func (s *DetachHostAccountsFromUserResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *DetachHostAccountsFromUserResponseBodyResults) GetUserId() *string {
	return s.UserId
}

func (s *DetachHostAccountsFromUserResponseBodyResults) SetCode(v string) *DetachHostAccountsFromUserResponseBodyResults {
	s.Code = &v
	return s
}

func (s *DetachHostAccountsFromUserResponseBodyResults) SetHostAccounts(v []*DetachHostAccountsFromUserResponseBodyResultsHostAccounts) *DetachHostAccountsFromUserResponseBodyResults {
	s.HostAccounts = v
	return s
}

func (s *DetachHostAccountsFromUserResponseBodyResults) SetHostId(v string) *DetachHostAccountsFromUserResponseBodyResults {
	s.HostId = &v
	return s
}

func (s *DetachHostAccountsFromUserResponseBodyResults) SetMessage(v string) *DetachHostAccountsFromUserResponseBodyResults {
	s.Message = &v
	return s
}

func (s *DetachHostAccountsFromUserResponseBodyResults) SetUserId(v string) *DetachHostAccountsFromUserResponseBodyResults {
	s.UserId = &v
	return s
}

func (s *DetachHostAccountsFromUserResponseBodyResults) Validate() error {
	if s.HostAccounts != nil {
		for _, item := range s.HostAccounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DetachHostAccountsFromUserResponseBodyResultsHostAccounts struct {
	// The return code that indicates whether permissions on the specified host accounts were revoked from the user. Valid values:
	//
	// 	- **OK**: The call was successful.
	//
	// 	- **UNEXPECTED**: An unknown error occurred.
	//
	// 	- **INVALID_ARGUMENT**: A request parameter is invalid.
	//
	// 	- **OBJECT_NOT_FOUND**: The specified object on which you want to perform the operation does not exist.
	//
	// 	- **OBJECT_AlREADY_EXISTS**: The specified object on which you want to perform the operation already exists.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The host account ID.
	//
	// example:
	//
	// 1
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	// This parameter is deprecated.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DetachHostAccountsFromUserResponseBodyResultsHostAccounts) String() string {
	return dara.Prettify(s)
}

func (s DetachHostAccountsFromUserResponseBodyResultsHostAccounts) GoString() string {
	return s.String()
}

func (s *DetachHostAccountsFromUserResponseBodyResultsHostAccounts) GetCode() *string {
	return s.Code
}

func (s *DetachHostAccountsFromUserResponseBodyResultsHostAccounts) GetHostAccountId() *string {
	return s.HostAccountId
}

func (s *DetachHostAccountsFromUserResponseBodyResultsHostAccounts) GetMessage() *string {
	return s.Message
}

func (s *DetachHostAccountsFromUserResponseBodyResultsHostAccounts) SetCode(v string) *DetachHostAccountsFromUserResponseBodyResultsHostAccounts {
	s.Code = &v
	return s
}

func (s *DetachHostAccountsFromUserResponseBodyResultsHostAccounts) SetHostAccountId(v string) *DetachHostAccountsFromUserResponseBodyResultsHostAccounts {
	s.HostAccountId = &v
	return s
}

func (s *DetachHostAccountsFromUserResponseBodyResultsHostAccounts) SetMessage(v string) *DetachHostAccountsFromUserResponseBodyResultsHostAccounts {
	s.Message = &v
	return s
}

func (s *DetachHostAccountsFromUserResponseBodyResultsHostAccounts) Validate() error {
	return dara.Validate(s)
}
