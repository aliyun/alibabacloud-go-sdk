// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachHostAccountsFromUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachHostAccountsFromUserGroupResponseBody
	GetRequestId() *string
	SetResults(v []*DetachHostAccountsFromUserGroupResponseBodyResults) *DetachHostAccountsFromUserGroupResponseBody
	GetResults() []*DetachHostAccountsFromUserGroupResponseBodyResults
}

type DetachHostAccountsFromUserGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*DetachHostAccountsFromUserGroupResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s DetachHostAccountsFromUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachHostAccountsFromUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DetachHostAccountsFromUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachHostAccountsFromUserGroupResponseBody) GetResults() []*DetachHostAccountsFromUserGroupResponseBodyResults {
	return s.Results
}

func (s *DetachHostAccountsFromUserGroupResponseBody) SetRequestId(v string) *DetachHostAccountsFromUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachHostAccountsFromUserGroupResponseBody) SetResults(v []*DetachHostAccountsFromUserGroupResponseBodyResults) *DetachHostAccountsFromUserGroupResponseBody {
	s.Results = v
	return s
}

func (s *DetachHostAccountsFromUserGroupResponseBody) Validate() error {
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

type DetachHostAccountsFromUserGroupResponseBodyResults struct {
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
	// The result of revoking permissions on the specified host accounts from the user group.
	HostAccounts []*DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts `json:"HostAccounts,omitempty" xml:"HostAccounts,omitempty" type:"Repeated"`
	// The ID of the host.
	//
	// example:
	//
	// １
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// This parameter is deprecated.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the group.
	//
	// example:
	//
	// １
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s DetachHostAccountsFromUserGroupResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s DetachHostAccountsFromUserGroupResponseBodyResults) GoString() string {
	return s.String()
}

func (s *DetachHostAccountsFromUserGroupResponseBodyResults) GetCode() *string {
	return s.Code
}

func (s *DetachHostAccountsFromUserGroupResponseBodyResults) GetHostAccounts() []*DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts {
	return s.HostAccounts
}

func (s *DetachHostAccountsFromUserGroupResponseBodyResults) GetHostId() *string {
	return s.HostId
}

func (s *DetachHostAccountsFromUserGroupResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *DetachHostAccountsFromUserGroupResponseBodyResults) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *DetachHostAccountsFromUserGroupResponseBodyResults) SetCode(v string) *DetachHostAccountsFromUserGroupResponseBodyResults {
	s.Code = &v
	return s
}

func (s *DetachHostAccountsFromUserGroupResponseBodyResults) SetHostAccounts(v []*DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts) *DetachHostAccountsFromUserGroupResponseBodyResults {
	s.HostAccounts = v
	return s
}

func (s *DetachHostAccountsFromUserGroupResponseBodyResults) SetHostId(v string) *DetachHostAccountsFromUserGroupResponseBodyResults {
	s.HostId = &v
	return s
}

func (s *DetachHostAccountsFromUserGroupResponseBodyResults) SetMessage(v string) *DetachHostAccountsFromUserGroupResponseBodyResults {
	s.Message = &v
	return s
}

func (s *DetachHostAccountsFromUserGroupResponseBodyResults) SetUserGroupId(v string) *DetachHostAccountsFromUserGroupResponseBodyResults {
	s.UserGroupId = &v
	return s
}

func (s *DetachHostAccountsFromUserGroupResponseBodyResults) Validate() error {
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

type DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts struct {
	// The return code that indicates whether permissions on the specified host account were revoked from the user group. Valid values:
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
	// The ID of the host account.
	//
	// example:
	//
	// １
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	// This parameter is deprecated.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts) String() string {
	return dara.Prettify(s)
}

func (s DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts) GoString() string {
	return s.String()
}

func (s *DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts) GetCode() *string {
	return s.Code
}

func (s *DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts) GetHostAccountId() *string {
	return s.HostAccountId
}

func (s *DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts) GetMessage() *string {
	return s.Message
}

func (s *DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts) SetCode(v string) *DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts {
	s.Code = &v
	return s
}

func (s *DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts) SetHostAccountId(v string) *DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts {
	s.HostAccountId = &v
	return s
}

func (s *DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts) SetMessage(v string) *DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts {
	s.Message = &v
	return s
}

func (s *DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts) Validate() error {
	return dara.Validate(s)
}
