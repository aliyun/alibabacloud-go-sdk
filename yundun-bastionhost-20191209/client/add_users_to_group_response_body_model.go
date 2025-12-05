// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUsersToGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddUsersToGroupResponseBody
	GetRequestId() *string
	SetResults(v []*AddUsersToGroupResponseBodyResults) *AddUsersToGroupResponseBody
	GetResults() []*AddUsersToGroupResponseBodyResults
}

type AddUsersToGroupResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*AddUsersToGroupResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s AddUsersToGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddUsersToGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddUsersToGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddUsersToGroupResponseBody) GetResults() []*AddUsersToGroupResponseBodyResults {
	return s.Results
}

func (s *AddUsersToGroupResponseBody) SetRequestId(v string) *AddUsersToGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddUsersToGroupResponseBody) SetResults(v []*AddUsersToGroupResponseBodyResults) *AddUsersToGroupResponseBody {
	s.Results = v
	return s
}

func (s *AddUsersToGroupResponseBody) Validate() error {
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

type AddUsersToGroupResponseBodyResults struct {
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

func (s AddUsersToGroupResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s AddUsersToGroupResponseBodyResults) GoString() string {
	return s.String()
}

func (s *AddUsersToGroupResponseBodyResults) GetCode() *string {
	return s.Code
}

func (s *AddUsersToGroupResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *AddUsersToGroupResponseBodyResults) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *AddUsersToGroupResponseBodyResults) GetUserId() *string {
	return s.UserId
}

func (s *AddUsersToGroupResponseBodyResults) SetCode(v string) *AddUsersToGroupResponseBodyResults {
	s.Code = &v
	return s
}

func (s *AddUsersToGroupResponseBodyResults) SetMessage(v string) *AddUsersToGroupResponseBodyResults {
	s.Message = &v
	return s
}

func (s *AddUsersToGroupResponseBodyResults) SetUserGroupId(v string) *AddUsersToGroupResponseBodyResults {
	s.UserGroupId = &v
	return s
}

func (s *AddUsersToGroupResponseBodyResults) SetUserId(v string) *AddUsersToGroupResponseBodyResults {
	s.UserId = &v
	return s
}

func (s *AddUsersToGroupResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
