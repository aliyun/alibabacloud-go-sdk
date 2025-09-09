// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachHostGroupAccountsFromUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachHostGroupAccountsFromUserResponseBody
	GetRequestId() *string
	SetResults(v []*DetachHostGroupAccountsFromUserResponseBodyResults) *DetachHostGroupAccountsFromUserResponseBody
	GetResults() []*DetachHostGroupAccountsFromUserResponseBodyResults
}

type DetachHostGroupAccountsFromUserResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*DetachHostGroupAccountsFromUserResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s DetachHostGroupAccountsFromUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachHostGroupAccountsFromUserResponseBody) GoString() string {
	return s.String()
}

func (s *DetachHostGroupAccountsFromUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachHostGroupAccountsFromUserResponseBody) GetResults() []*DetachHostGroupAccountsFromUserResponseBodyResults {
	return s.Results
}

func (s *DetachHostGroupAccountsFromUserResponseBody) SetRequestId(v string) *DetachHostGroupAccountsFromUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserResponseBody) SetResults(v []*DetachHostGroupAccountsFromUserResponseBodyResults) *DetachHostGroupAccountsFromUserResponseBody {
	s.Results = v
	return s
}

func (s *DetachHostGroupAccountsFromUserResponseBody) Validate() error {
	return dara.Validate(s)
}

type DetachHostGroupAccountsFromUserResponseBodyResults struct {
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
	HostAccountNames []*DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames `json:"HostAccountNames,omitempty" xml:"HostAccountNames,omitempty" type:"Repeated"`
	// The ID of the host group.
	//
	// example:
	//
	// 1
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
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

func (s DetachHostGroupAccountsFromUserResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s DetachHostGroupAccountsFromUserResponseBodyResults) GoString() string {
	return s.String()
}

func (s *DetachHostGroupAccountsFromUserResponseBodyResults) GetCode() *string {
	return s.Code
}

func (s *DetachHostGroupAccountsFromUserResponseBodyResults) GetHostAccountNames() []*DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames {
	return s.HostAccountNames
}

func (s *DetachHostGroupAccountsFromUserResponseBodyResults) GetHostGroupId() *string {
	return s.HostGroupId
}

func (s *DetachHostGroupAccountsFromUserResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *DetachHostGroupAccountsFromUserResponseBodyResults) GetUserId() *string {
	return s.UserId
}

func (s *DetachHostGroupAccountsFromUserResponseBodyResults) SetCode(v string) *DetachHostGroupAccountsFromUserResponseBodyResults {
	s.Code = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserResponseBodyResults) SetHostAccountNames(v []*DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames) *DetachHostGroupAccountsFromUserResponseBodyResults {
	s.HostAccountNames = v
	return s
}

func (s *DetachHostGroupAccountsFromUserResponseBodyResults) SetHostGroupId(v string) *DetachHostGroupAccountsFromUserResponseBodyResults {
	s.HostGroupId = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserResponseBodyResults) SetMessage(v string) *DetachHostGroupAccountsFromUserResponseBodyResults {
	s.Message = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserResponseBodyResults) SetUserId(v string) *DetachHostGroupAccountsFromUserResponseBodyResults {
	s.UserId = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserResponseBodyResults) Validate() error {
	return dara.Validate(s)
}

type DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames struct {
	// The return code that indicates whether permissions on the specified host account were revoked from the user. Valid values:
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
	// root
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// N/A
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames) String() string {
	return dara.Prettify(s)
}

func (s DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames) GoString() string {
	return s.String()
}

func (s *DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames) GetCode() *string {
	return s.Code
}

func (s *DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames) GetHostAccountName() *string {
	return s.HostAccountName
}

func (s *DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames) GetMessage() *string {
	return s.Message
}

func (s *DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames) SetCode(v string) *DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames {
	s.Code = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames) SetHostAccountName(v string) *DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames {
	s.HostAccountName = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames) SetMessage(v string) *DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames {
	s.Message = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames) Validate() error {
	return dara.Validate(s)
}
