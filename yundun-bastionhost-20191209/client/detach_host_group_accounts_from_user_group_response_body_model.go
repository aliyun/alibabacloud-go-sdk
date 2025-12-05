// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachHostGroupAccountsFromUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachHostGroupAccountsFromUserGroupResponseBody
	GetRequestId() *string
	SetResults(v []*DetachHostGroupAccountsFromUserGroupResponseBodyResults) *DetachHostGroupAccountsFromUserGroupResponseBody
	GetResults() []*DetachHostGroupAccountsFromUserGroupResponseBodyResults
}

type DetachHostGroupAccountsFromUserGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*DetachHostGroupAccountsFromUserGroupResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s DetachHostGroupAccountsFromUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachHostGroupAccountsFromUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBody) GetResults() []*DetachHostGroupAccountsFromUserGroupResponseBodyResults {
	return s.Results
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBody) SetRequestId(v string) *DetachHostGroupAccountsFromUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBody) SetResults(v []*DetachHostGroupAccountsFromUserGroupResponseBodyResults) *DetachHostGroupAccountsFromUserGroupResponseBody {
	s.Results = v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBody) Validate() error {
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

type DetachHostGroupAccountsFromUserGroupResponseBodyResults struct {
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
	HostAccountNames []*DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames `json:"HostAccountNames,omitempty" xml:"HostAccountNames,omitempty" type:"Repeated"`
	// The ID of the host group.
	//
	// example:
	//
	// 1
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// This parameter is deprecated.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the group.
	//
	// example:
	//
	// 1
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s DetachHostGroupAccountsFromUserGroupResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s DetachHostGroupAccountsFromUserGroupResponseBodyResults) GoString() string {
	return s.String()
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBodyResults) GetCode() *string {
	return s.Code
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBodyResults) GetHostAccountNames() []*DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames {
	return s.HostAccountNames
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBodyResults) GetHostGroupId() *string {
	return s.HostGroupId
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBodyResults) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBodyResults) SetCode(v string) *DetachHostGroupAccountsFromUserGroupResponseBodyResults {
	s.Code = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBodyResults) SetHostAccountNames(v []*DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames) *DetachHostGroupAccountsFromUserGroupResponseBodyResults {
	s.HostAccountNames = v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBodyResults) SetHostGroupId(v string) *DetachHostGroupAccountsFromUserGroupResponseBodyResults {
	s.HostGroupId = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBodyResults) SetMessage(v string) *DetachHostGroupAccountsFromUserGroupResponseBodyResults {
	s.Message = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBodyResults) SetUserGroupId(v string) *DetachHostGroupAccountsFromUserGroupResponseBodyResults {
	s.UserGroupId = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBodyResults) Validate() error {
	if s.HostAccountNames != nil {
		for _, item := range s.HostAccountNames {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames struct {
	// The return code that indicates whether permissions on the specified host account were revoked from the specified user group. Valid values:
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
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames) String() string {
	return dara.Prettify(s)
}

func (s DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames) GoString() string {
	return s.String()
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames) GetCode() *string {
	return s.Code
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames) GetHostAccountName() *string {
	return s.HostAccountName
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames) GetMessage() *string {
	return s.Message
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames) SetCode(v string) *DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames {
	s.Code = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames) SetHostAccountName(v string) *DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames {
	s.HostAccountName = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames) SetMessage(v string) *DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames {
	s.Message = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames) Validate() error {
	return dara.Validate(s)
}
