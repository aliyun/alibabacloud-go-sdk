// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachHostAccountsToUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachHostAccountsToUserGroupResponseBody
	GetRequestId() *string
	SetResults(v []*AttachHostAccountsToUserGroupResponseBodyResults) *AttachHostAccountsToUserGroupResponseBody
	GetResults() []*AttachHostAccountsToUserGroupResponseBodyResults
}

type AttachHostAccountsToUserGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*AttachHostAccountsToUserGroupResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s AttachHostAccountsToUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachHostAccountsToUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AttachHostAccountsToUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachHostAccountsToUserGroupResponseBody) GetResults() []*AttachHostAccountsToUserGroupResponseBodyResults {
	return s.Results
}

func (s *AttachHostAccountsToUserGroupResponseBody) SetRequestId(v string) *AttachHostAccountsToUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachHostAccountsToUserGroupResponseBody) SetResults(v []*AttachHostAccountsToUserGroupResponseBodyResults) *AttachHostAccountsToUserGroupResponseBody {
	s.Results = v
	return s
}

func (s *AttachHostAccountsToUserGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type AttachHostAccountsToUserGroupResponseBodyResults struct {
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
	// The result of authorizing the specified user group to manage the specified host accounts.
	HostAccounts []*AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts `json:"HostAccounts,omitempty" xml:"HostAccounts,omitempty" type:"Repeated"`
	// The ID of the host.
	//
	// example:
	//
	// 1
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// N/A
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the user group.
	//
	// example:
	//
	// 1
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s AttachHostAccountsToUserGroupResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s AttachHostAccountsToUserGroupResponseBodyResults) GoString() string {
	return s.String()
}

func (s *AttachHostAccountsToUserGroupResponseBodyResults) GetCode() *string {
	return s.Code
}

func (s *AttachHostAccountsToUserGroupResponseBodyResults) GetHostAccounts() []*AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts {
	return s.HostAccounts
}

func (s *AttachHostAccountsToUserGroupResponseBodyResults) GetHostId() *string {
	return s.HostId
}

func (s *AttachHostAccountsToUserGroupResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *AttachHostAccountsToUserGroupResponseBodyResults) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *AttachHostAccountsToUserGroupResponseBodyResults) SetCode(v string) *AttachHostAccountsToUserGroupResponseBodyResults {
	s.Code = &v
	return s
}

func (s *AttachHostAccountsToUserGroupResponseBodyResults) SetHostAccounts(v []*AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts) *AttachHostAccountsToUserGroupResponseBodyResults {
	s.HostAccounts = v
	return s
}

func (s *AttachHostAccountsToUserGroupResponseBodyResults) SetHostId(v string) *AttachHostAccountsToUserGroupResponseBodyResults {
	s.HostId = &v
	return s
}

func (s *AttachHostAccountsToUserGroupResponseBodyResults) SetMessage(v string) *AttachHostAccountsToUserGroupResponseBodyResults {
	s.Message = &v
	return s
}

func (s *AttachHostAccountsToUserGroupResponseBodyResults) SetUserGroupId(v string) *AttachHostAccountsToUserGroupResponseBodyResults {
	s.UserGroupId = &v
	return s
}

func (s *AttachHostAccountsToUserGroupResponseBodyResults) Validate() error {
	return dara.Validate(s)
}

type AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts struct {
	// The return code that indicates whether the user group was authorized to manage the specified host account. Valid values:
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
	// 1
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// N/A
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts) String() string {
	return dara.Prettify(s)
}

func (s AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts) GoString() string {
	return s.String()
}

func (s *AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts) GetCode() *string {
	return s.Code
}

func (s *AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts) GetHostAccountId() *string {
	return s.HostAccountId
}

func (s *AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts) GetMessage() *string {
	return s.Message
}

func (s *AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts) SetCode(v string) *AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts {
	s.Code = &v
	return s
}

func (s *AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts) SetHostAccountId(v string) *AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts {
	s.HostAccountId = &v
	return s
}

func (s *AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts) SetMessage(v string) *AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts {
	s.Message = &v
	return s
}

func (s *AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts) Validate() error {
	return dara.Validate(s)
}
