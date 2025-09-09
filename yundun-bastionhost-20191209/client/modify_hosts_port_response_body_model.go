// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHostsPortResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyHostsPortResponseBody
	GetRequestId() *string
	SetResults(v []*ModifyHostsPortResponseBodyResults) *ModifyHostsPortResponseBody
	GetResults() []*ModifyHostsPortResponseBodyResults
}

type ModifyHostsPortResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*ModifyHostsPortResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s ModifyHostsPortResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostsPortResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHostsPortResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHostsPortResponseBody) GetResults() []*ModifyHostsPortResponseBodyResults {
	return s.Results
}

func (s *ModifyHostsPortResponseBody) SetRequestId(v string) *ModifyHostsPortResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHostsPortResponseBody) SetResults(v []*ModifyHostsPortResponseBodyResults) *ModifyHostsPortResponseBody {
	s.Results = v
	return s
}

func (s *ModifyHostsPortResponseBody) Validate() error {
	return dara.Validate(s)
}

type ModifyHostsPortResponseBodyResults struct {
	// The return code that indicates whether the call was successful. Valid values:
	//
	// 	- **OK**: The call was successful.
	//
	// 	- **UNEXPECTED**: An unknown error occurred.
	//
	// 	- **INVALID_ARGUMENT**: A request parameter is invalid.
	//
	//     > Make sure that the request parameters are valid and call the operation again.
	//
	// 	- **OBJECT_NOT_FOUND**: The specified object on which you want to perform the operation does not exist.
	//
	//     > Check whether the specified ID of the bastion host exists, whether the specified hosts exist, and whether the specified host IDs are valid. Then, call the operation again.
	//
	// 	- **OBJECT_AlREADY_EXISTS**: The specified object on which you want to perform the operation already exists.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the host.
	//
	// example:
	//
	// 1
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// -
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ModifyHostsPortResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostsPortResponseBodyResults) GoString() string {
	return s.String()
}

func (s *ModifyHostsPortResponseBodyResults) GetCode() *string {
	return s.Code
}

func (s *ModifyHostsPortResponseBodyResults) GetHostId() *string {
	return s.HostId
}

func (s *ModifyHostsPortResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *ModifyHostsPortResponseBodyResults) SetCode(v string) *ModifyHostsPortResponseBodyResults {
	s.Code = &v
	return s
}

func (s *ModifyHostsPortResponseBodyResults) SetHostId(v string) *ModifyHostsPortResponseBodyResults {
	s.HostId = &v
	return s
}

func (s *ModifyHostsPortResponseBodyResults) SetMessage(v string) *ModifyHostsPortResponseBodyResults {
	s.Message = &v
	return s
}

func (s *ModifyHostsPortResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
