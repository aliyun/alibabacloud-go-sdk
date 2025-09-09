// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnlockUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnlockUsersResponseBody
	GetRequestId() *string
	SetResults(v []*UnlockUsersResponseBodyResults) *UnlockUsersResponseBody
	GetResults() []*UnlockUsersResponseBodyResults
}

type UnlockUsersResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of information about the result of the call.
	Results []*UnlockUsersResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s UnlockUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnlockUsersResponseBody) GoString() string {
	return s.String()
}

func (s *UnlockUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnlockUsersResponseBody) GetResults() []*UnlockUsersResponseBodyResults {
	return s.Results
}

func (s *UnlockUsersResponseBody) SetRequestId(v string) *UnlockUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnlockUsersResponseBody) SetResults(v []*UnlockUsersResponseBodyResults) *UnlockUsersResponseBody {
	s.Results = v
	return s
}

func (s *UnlockUsersResponseBody) Validate() error {
	return dara.Validate(s)
}

type UnlockUsersResponseBodyResults struct {
	// The result of the call. Valid values:
	//
	// 	- **OK**: The call was successful.
	//
	// 	- **UNEXPECTED**: An unknown error occurred.
	//
	// 	- **INVALID_ARGUMENT**: A request parameter is invalid.
	//
	//     **
	//
	//     **Note**Make sure that the request parameters are valid and call the operation again.
	//
	// 	- **OBJECT_NOT_FOUND**: The specified object on which you want to perform the operation does not exist.
	//
	//     **
	//
	//     **Note**Check whether the specified ID of the bastion host exists, whether the specified hosts exist, and whether the specified host IDs are valid. Then, call the operation again.
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
	// １
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UnlockUsersResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s UnlockUsersResponseBodyResults) GoString() string {
	return s.String()
}

func (s *UnlockUsersResponseBodyResults) GetCode() *string {
	return s.Code
}

func (s *UnlockUsersResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *UnlockUsersResponseBodyResults) GetUserId() *string {
	return s.UserId
}

func (s *UnlockUsersResponseBodyResults) SetCode(v string) *UnlockUsersResponseBodyResults {
	s.Code = &v
	return s
}

func (s *UnlockUsersResponseBodyResults) SetMessage(v string) *UnlockUsersResponseBodyResults {
	s.Message = &v
	return s
}

func (s *UnlockUsersResponseBodyResults) SetUserId(v string) *UnlockUsersResponseBodyResults {
	s.UserId = &v
	return s
}

func (s *UnlockUsersResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
