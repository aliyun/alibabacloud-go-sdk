// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachHostGroupAccountsToUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachHostGroupAccountsToUserResponseBody
	GetRequestId() *string
	SetResults(v []*AttachHostGroupAccountsToUserResponseBodyResults) *AttachHostGroupAccountsToUserResponseBody
	GetResults() []*AttachHostGroupAccountsToUserResponseBodyResults
}

type AttachHostGroupAccountsToUserResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*AttachHostGroupAccountsToUserResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s AttachHostGroupAccountsToUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachHostGroupAccountsToUserResponseBody) GoString() string {
	return s.String()
}

func (s *AttachHostGroupAccountsToUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachHostGroupAccountsToUserResponseBody) GetResults() []*AttachHostGroupAccountsToUserResponseBodyResults {
	return s.Results
}

func (s *AttachHostGroupAccountsToUserResponseBody) SetRequestId(v string) *AttachHostGroupAccountsToUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachHostGroupAccountsToUserResponseBody) SetResults(v []*AttachHostGroupAccountsToUserResponseBodyResults) *AttachHostGroupAccountsToUserResponseBody {
	s.Results = v
	return s
}

func (s *AttachHostGroupAccountsToUserResponseBody) Validate() error {
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

type AttachHostGroupAccountsToUserResponseBodyResults struct {
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
	// The result of authorizing the user to manage the host accounts.
	HostAccountNames []*AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames `json:"HostAccountNames,omitempty" xml:"HostAccountNames,omitempty" type:"Repeated"`
	// The ID of the host group.
	//
	// example:
	//
	// 1
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// This parameter is deprecated.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// 1
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AttachHostGroupAccountsToUserResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s AttachHostGroupAccountsToUserResponseBodyResults) GoString() string {
	return s.String()
}

func (s *AttachHostGroupAccountsToUserResponseBodyResults) GetCode() *string {
	return s.Code
}

func (s *AttachHostGroupAccountsToUserResponseBodyResults) GetHostAccountNames() []*AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames {
	return s.HostAccountNames
}

func (s *AttachHostGroupAccountsToUserResponseBodyResults) GetHostGroupId() *string {
	return s.HostGroupId
}

func (s *AttachHostGroupAccountsToUserResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *AttachHostGroupAccountsToUserResponseBodyResults) GetUserId() *string {
	return s.UserId
}

func (s *AttachHostGroupAccountsToUserResponseBodyResults) SetCode(v string) *AttachHostGroupAccountsToUserResponseBodyResults {
	s.Code = &v
	return s
}

func (s *AttachHostGroupAccountsToUserResponseBodyResults) SetHostAccountNames(v []*AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames) *AttachHostGroupAccountsToUserResponseBodyResults {
	s.HostAccountNames = v
	return s
}

func (s *AttachHostGroupAccountsToUserResponseBodyResults) SetHostGroupId(v string) *AttachHostGroupAccountsToUserResponseBodyResults {
	s.HostGroupId = &v
	return s
}

func (s *AttachHostGroupAccountsToUserResponseBodyResults) SetMessage(v string) *AttachHostGroupAccountsToUserResponseBodyResults {
	s.Message = &v
	return s
}

func (s *AttachHostGroupAccountsToUserResponseBodyResults) SetUserId(v string) *AttachHostGroupAccountsToUserResponseBodyResults {
	s.UserId = &v
	return s
}

func (s *AttachHostGroupAccountsToUserResponseBodyResults) Validate() error {
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

type AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames struct {
	// The return code that indicates whether the user was authorized to manage the host account. Valid values:
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

func (s AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames) String() string {
	return dara.Prettify(s)
}

func (s AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames) GoString() string {
	return s.String()
}

func (s *AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames) GetCode() *string {
	return s.Code
}

func (s *AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames) GetHostAccountName() *string {
	return s.HostAccountName
}

func (s *AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames) GetMessage() *string {
	return s.Message
}

func (s *AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames) SetCode(v string) *AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames {
	s.Code = &v
	return s
}

func (s *AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames) SetHostAccountName(v string) *AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames {
	s.HostAccountName = &v
	return s
}

func (s *AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames) SetMessage(v string) *AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames {
	s.Message = &v
	return s
}

func (s *AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames) Validate() error {
	return dara.Validate(s)
}
