// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLockUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *LockUsersResponseBody
	GetRequestId() *string
	SetResults(v []*LockUsersResponseBodyResults) *LockUsersResponseBody
	GetResults() []*LockUsersResponseBodyResults
}

type LockUsersResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*LockUsersResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s LockUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LockUsersResponseBody) GoString() string {
	return s.String()
}

func (s *LockUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LockUsersResponseBody) GetResults() []*LockUsersResponseBodyResults {
	return s.Results
}

func (s *LockUsersResponseBody) SetRequestId(v string) *LockUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *LockUsersResponseBody) SetResults(v []*LockUsersResponseBodyResults) *LockUsersResponseBody {
	s.Results = v
	return s
}

func (s *LockUsersResponseBody) Validate() error {
	return dara.Validate(s)
}

type LockUsersResponseBodyResults struct {
	// The return code that indicates whether the call was successful. Valid values:
	//
	// 	- **OK**: The call was successful.
	//
	// 	- **UNEXPECTED**: An unknown error occurred.
	//
	// 	- **INVALID_ARGUMENT**: A request parameter is invalid.
	//
	// >Make sure that the request parameters are valid and call the operation again.
	//
	// 	- **OBJECT_NOT_FOUND**: The specified object on which you want to perform the operation does not exist.
	//
	// >Check whether the specified ID of the bastion host exists, whether the specified hosts exist, and whether the specified host IDs are valid. Then, call the operation again.
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
	// N/A
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// 1
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s LockUsersResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s LockUsersResponseBodyResults) GoString() string {
	return s.String()
}

func (s *LockUsersResponseBodyResults) GetCode() *string {
	return s.Code
}

func (s *LockUsersResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *LockUsersResponseBodyResults) GetUserId() *string {
	return s.UserId
}

func (s *LockUsersResponseBodyResults) SetCode(v string) *LockUsersResponseBodyResults {
	s.Code = &v
	return s
}

func (s *LockUsersResponseBodyResults) SetMessage(v string) *LockUsersResponseBodyResults {
	s.Message = &v
	return s
}

func (s *LockUsersResponseBodyResults) SetUserId(v string) *LockUsersResponseBodyResults {
	s.UserId = &v
	return s
}

func (s *LockUsersResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
