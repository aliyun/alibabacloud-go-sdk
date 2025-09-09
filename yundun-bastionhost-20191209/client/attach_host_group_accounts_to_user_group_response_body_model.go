// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachHostGroupAccountsToUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachHostGroupAccountsToUserGroupResponseBody
	GetRequestId() *string
	SetResults(v []*AttachHostGroupAccountsToUserGroupResponseBodyResults) *AttachHostGroupAccountsToUserGroupResponseBody
	GetResults() []*AttachHostGroupAccountsToUserGroupResponseBodyResults
}

type AttachHostGroupAccountsToUserGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*AttachHostGroupAccountsToUserGroupResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s AttachHostGroupAccountsToUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachHostGroupAccountsToUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AttachHostGroupAccountsToUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachHostGroupAccountsToUserGroupResponseBody) GetResults() []*AttachHostGroupAccountsToUserGroupResponseBodyResults {
	return s.Results
}

func (s *AttachHostGroupAccountsToUserGroupResponseBody) SetRequestId(v string) *AttachHostGroupAccountsToUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupResponseBody) SetResults(v []*AttachHostGroupAccountsToUserGroupResponseBodyResults) *AttachHostGroupAccountsToUserGroupResponseBody {
	s.Results = v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type AttachHostGroupAccountsToUserGroupResponseBodyResults struct {
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
	// The result of authorizing the user group to manage the specified host accounts.
	HostAccountNames []*AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames `json:"HostAccountNames,omitempty" xml:"HostAccountNames,omitempty" type:"Repeated"`
	// The ID of the host group.
	//
	// example:
	//
	// 1
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// This parameter is deprecated.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the user group.
	//
	// example:
	//
	// 1
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s AttachHostGroupAccountsToUserGroupResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s AttachHostGroupAccountsToUserGroupResponseBodyResults) GoString() string {
	return s.String()
}

func (s *AttachHostGroupAccountsToUserGroupResponseBodyResults) GetCode() *string {
	return s.Code
}

func (s *AttachHostGroupAccountsToUserGroupResponseBodyResults) GetHostAccountNames() []*AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames {
	return s.HostAccountNames
}

func (s *AttachHostGroupAccountsToUserGroupResponseBodyResults) GetHostGroupId() *string {
	return s.HostGroupId
}

func (s *AttachHostGroupAccountsToUserGroupResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *AttachHostGroupAccountsToUserGroupResponseBodyResults) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *AttachHostGroupAccountsToUserGroupResponseBodyResults) SetCode(v string) *AttachHostGroupAccountsToUserGroupResponseBodyResults {
	s.Code = &v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupResponseBodyResults) SetHostAccountNames(v []*AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames) *AttachHostGroupAccountsToUserGroupResponseBodyResults {
	s.HostAccountNames = v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupResponseBodyResults) SetHostGroupId(v string) *AttachHostGroupAccountsToUserGroupResponseBodyResults {
	s.HostGroupId = &v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupResponseBodyResults) SetMessage(v string) *AttachHostGroupAccountsToUserGroupResponseBodyResults {
	s.Message = &v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupResponseBodyResults) SetUserGroupId(v string) *AttachHostGroupAccountsToUserGroupResponseBodyResults {
	s.UserGroupId = &v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupResponseBodyResults) Validate() error {
	return dara.Validate(s)
}

type AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames struct {
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
	// The name of the host account.
	//
	// example:
	//
	// abc
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	// This parameter is deprecated.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames) String() string {
	return dara.Prettify(s)
}

func (s AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames) GoString() string {
	return s.String()
}

func (s *AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames) GetCode() *string {
	return s.Code
}

func (s *AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames) GetHostAccountName() *string {
	return s.HostAccountName
}

func (s *AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames) GetMessage() *string {
	return s.Message
}

func (s *AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames) SetCode(v string) *AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames {
	s.Code = &v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames) SetHostAccountName(v string) *AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames {
	s.HostAccountName = &v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames) SetMessage(v string) *AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames {
	s.Message = &v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames) Validate() error {
	return dara.Validate(s)
}
