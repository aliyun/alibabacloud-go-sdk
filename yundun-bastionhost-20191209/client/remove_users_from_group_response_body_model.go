// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUsersFromGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveUsersFromGroupResponseBody
	GetRequestId() *string
	SetResults(v []*RemoveUsersFromGroupResponseBodyResults) *RemoveUsersFromGroupResponseBody
	GetResults() []*RemoveUsersFromGroupResponseBodyResults
}

type RemoveUsersFromGroupResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*RemoveUsersFromGroupResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s RemoveUsersFromGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveUsersFromGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUsersFromGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveUsersFromGroupResponseBody) GetResults() []*RemoveUsersFromGroupResponseBodyResults {
	return s.Results
}

func (s *RemoveUsersFromGroupResponseBody) SetRequestId(v string) *RemoveUsersFromGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveUsersFromGroupResponseBody) SetResults(v []*RemoveUsersFromGroupResponseBodyResults) *RemoveUsersFromGroupResponseBody {
	s.Results = v
	return s
}

func (s *RemoveUsersFromGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type RemoveUsersFromGroupResponseBodyResults struct {
	// The return code that indicates whether the call was successful. Valid values:
	//
	// 	- **OK**: The call was successful.
	//
	// 	- **UNEXPECTED**: An unknown error occurred.
	//
	// 	- **INVALID_ARGUMENT**: A request parameter is invalid.
	//
	// > Make sure that the request parameters are valid and call the operation again.
	//
	// 	- **OBJECT_NOT_FOUND**: The specified object on which you want to perform the operation does not exist.
	//
	// > Check whether the specified ID of the bastion host exists, whether the specified hosts exist, and whether the specified host IDs are valid. Then, call the operation again.
	//
	// 	- **OBJECT_AlREADY_EXISTS**: The specified object on which you want to perform the operation already exists.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the group.
	//
	// example:
	//
	// 1
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// 1
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RemoveUsersFromGroupResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s RemoveUsersFromGroupResponseBodyResults) GoString() string {
	return s.String()
}

func (s *RemoveUsersFromGroupResponseBodyResults) GetCode() *string {
	return s.Code
}

func (s *RemoveUsersFromGroupResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *RemoveUsersFromGroupResponseBodyResults) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *RemoveUsersFromGroupResponseBodyResults) GetUserId() *string {
	return s.UserId
}

func (s *RemoveUsersFromGroupResponseBodyResults) SetCode(v string) *RemoveUsersFromGroupResponseBodyResults {
	s.Code = &v
	return s
}

func (s *RemoveUsersFromGroupResponseBodyResults) SetMessage(v string) *RemoveUsersFromGroupResponseBodyResults {
	s.Message = &v
	return s
}

func (s *RemoveUsersFromGroupResponseBodyResults) SetUserGroupId(v string) *RemoveUsersFromGroupResponseBodyResults {
	s.UserGroupId = &v
	return s
}

func (s *RemoveUsersFromGroupResponseBodyResults) SetUserId(v string) *RemoveUsersFromGroupResponseBodyResults {
	s.UserId = &v
	return s
}

func (s *RemoveUsersFromGroupResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
