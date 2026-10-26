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
	// The request ID. Alibaba Cloud generates a unique ID for each request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result information of the operation.
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
	// The result of the operation. Valid values:
	//
	// - **OK**: The operation was successful.
	//
	// - **UNEXPECTED**: An unknown error occurred.
	//
	// - **INVALID_ARGUMENT**: The request parameters are invalid.
	//
	//     > Check the request parameters and call the operation again.
	//
	// - **OBJECT_NOT_FOUND**: The specified object does not exist.
	//
	//     > Check whether the bastion host instance ID exists, whether the host exists, and whether the host ID is correct, and then call the operation again.
	//
	// - **OBJECT_AlREADY_EXISTS**: The specified object already exists.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// This parameter is deprecated and can be ignored.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The user group ID.
	//
	// example:
	//
	// 1
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// The user ID.
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
