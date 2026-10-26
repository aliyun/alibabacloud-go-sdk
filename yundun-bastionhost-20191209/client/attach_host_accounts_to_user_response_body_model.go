// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachHostAccountsToUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachHostAccountsToUserResponseBody
	GetRequestId() *string
	SetResults(v []*AttachHostAccountsToUserResponseBodyResults) *AttachHostAccountsToUserResponseBody
	GetResults() []*AttachHostAccountsToUserResponseBodyResults
}

type AttachHostAccountsToUserResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The results of the API call.
	Results []*AttachHostAccountsToUserResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s AttachHostAccountsToUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachHostAccountsToUserResponseBody) GoString() string {
	return s.String()
}

func (s *AttachHostAccountsToUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachHostAccountsToUserResponseBody) GetResults() []*AttachHostAccountsToUserResponseBodyResults {
	return s.Results
}

func (s *AttachHostAccountsToUserResponseBody) SetRequestId(v string) *AttachHostAccountsToUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachHostAccountsToUserResponseBody) SetResults(v []*AttachHostAccountsToUserResponseBodyResults) *AttachHostAccountsToUserResponseBody {
	s.Results = v
	return s
}

func (s *AttachHostAccountsToUserResponseBody) Validate() error {
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

type AttachHostAccountsToUserResponseBodyResults struct {
	// The result of the API call. Valid values:
	//
	// - **OK**: The operation is successful.
	//
	// - **UNEXPECTED**: An unknown error occurred.
	//
	// - **INVALID_ARGUMENT**: A request parameter is invalid.
	//
	//     > Check the request parameters, ensure that they are correct, and then call the operation again.
	//
	// - **OBJECT_NOT_FOUND**: The object on which the operation is performed does not exist.
	//
	//     > Check whether the ID of the Bastionhost instance exists, whether the host exists, and whether the host ID is correct, and then call the operation again.
	//
	// - **OBJECT_AlREADY_EXISTS**: The object on which the operation is performed already exists.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The results returned for the operation of authorizing host accounts for the user.
	HostAccounts []*AttachHostAccountsToUserResponseBodyResultsHostAccounts `json:"HostAccounts,omitempty" xml:"HostAccounts,omitempty" type:"Repeated"`
	// The ID of the host.
	//
	// example:
	//
	// 1
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// **[Deprecated]*	- This parameter is deprecated. You can ignore it.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// 1
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AttachHostAccountsToUserResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s AttachHostAccountsToUserResponseBodyResults) GoString() string {
	return s.String()
}

func (s *AttachHostAccountsToUserResponseBodyResults) GetCode() *string {
	return s.Code
}

func (s *AttachHostAccountsToUserResponseBodyResults) GetHostAccounts() []*AttachHostAccountsToUserResponseBodyResultsHostAccounts {
	return s.HostAccounts
}

func (s *AttachHostAccountsToUserResponseBodyResults) GetHostId() *string {
	return s.HostId
}

func (s *AttachHostAccountsToUserResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *AttachHostAccountsToUserResponseBodyResults) GetUserId() *string {
	return s.UserId
}

func (s *AttachHostAccountsToUserResponseBodyResults) SetCode(v string) *AttachHostAccountsToUserResponseBodyResults {
	s.Code = &v
	return s
}

func (s *AttachHostAccountsToUserResponseBodyResults) SetHostAccounts(v []*AttachHostAccountsToUserResponseBodyResultsHostAccounts) *AttachHostAccountsToUserResponseBodyResults {
	s.HostAccounts = v
	return s
}

func (s *AttachHostAccountsToUserResponseBodyResults) SetHostId(v string) *AttachHostAccountsToUserResponseBodyResults {
	s.HostId = &v
	return s
}

func (s *AttachHostAccountsToUserResponseBodyResults) SetMessage(v string) *AttachHostAccountsToUserResponseBodyResults {
	s.Message = &v
	return s
}

func (s *AttachHostAccountsToUserResponseBodyResults) SetUserId(v string) *AttachHostAccountsToUserResponseBodyResults {
	s.UserId = &v
	return s
}

func (s *AttachHostAccountsToUserResponseBodyResults) Validate() error {
	if s.HostAccounts != nil {
		for _, item := range s.HostAccounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AttachHostAccountsToUserResponseBodyResultsHostAccounts struct {
	// The result returned for the operation of authorizing host accounts for the user. Valid values:
	//
	// - **OK**: The operation is successful.
	//
	// - **UNEXPECTED**: An unknown error occurred.
	//
	// - **INVALID_ARGUMENT**: A request parameter is invalid.
	//
	// - **OBJECT_NOT_FOUND**: The object on which the operation is performed does not exist.
	//
	// - **OBJECT_AlREADY_EXISTS**: The object on which the operation is performed already exists.
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
	// **[Deprecated]*	- This parameter is deprecated. You can ignore it.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s AttachHostAccountsToUserResponseBodyResultsHostAccounts) String() string {
	return dara.Prettify(s)
}

func (s AttachHostAccountsToUserResponseBodyResultsHostAccounts) GoString() string {
	return s.String()
}

func (s *AttachHostAccountsToUserResponseBodyResultsHostAccounts) GetCode() *string {
	return s.Code
}

func (s *AttachHostAccountsToUserResponseBodyResultsHostAccounts) GetHostAccountId() *string {
	return s.HostAccountId
}

func (s *AttachHostAccountsToUserResponseBodyResultsHostAccounts) GetMessage() *string {
	return s.Message
}

func (s *AttachHostAccountsToUserResponseBodyResultsHostAccounts) SetCode(v string) *AttachHostAccountsToUserResponseBodyResultsHostAccounts {
	s.Code = &v
	return s
}

func (s *AttachHostAccountsToUserResponseBodyResultsHostAccounts) SetHostAccountId(v string) *AttachHostAccountsToUserResponseBodyResultsHostAccounts {
	s.HostAccountId = &v
	return s
}

func (s *AttachHostAccountsToUserResponseBodyResultsHostAccounts) SetMessage(v string) *AttachHostAccountsToUserResponseBodyResultsHostAccounts {
	s.Message = &v
	return s
}

func (s *AttachHostAccountsToUserResponseBodyResultsHostAccounts) Validate() error {
	return dara.Validate(s)
}
