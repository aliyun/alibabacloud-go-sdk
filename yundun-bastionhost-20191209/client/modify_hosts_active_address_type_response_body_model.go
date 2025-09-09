// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHostsActiveAddressTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyHostsActiveAddressTypeResponseBody
	GetRequestId() *string
	SetResults(v []*ModifyHostsActiveAddressTypeResponseBodyResults) *ModifyHostsActiveAddressTypeResponseBody
	GetResults() []*ModifyHostsActiveAddressTypeResponseBodyResults
}

type ModifyHostsActiveAddressTypeResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*ModifyHostsActiveAddressTypeResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s ModifyHostsActiveAddressTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostsActiveAddressTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHostsActiveAddressTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHostsActiveAddressTypeResponseBody) GetResults() []*ModifyHostsActiveAddressTypeResponseBodyResults {
	return s.Results
}

func (s *ModifyHostsActiveAddressTypeResponseBody) SetRequestId(v string) *ModifyHostsActiveAddressTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHostsActiveAddressTypeResponseBody) SetResults(v []*ModifyHostsActiveAddressTypeResponseBodyResults) *ModifyHostsActiveAddressTypeResponseBody {
	s.Results = v
	return s
}

func (s *ModifyHostsActiveAddressTypeResponseBody) Validate() error {
	return dara.Validate(s)
}

type ModifyHostsActiveAddressTypeResponseBodyResults struct {
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
	// The ID of the host.
	//
	// example:
	//
	// 1
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// This parameter is deprecated.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ModifyHostsActiveAddressTypeResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostsActiveAddressTypeResponseBodyResults) GoString() string {
	return s.String()
}

func (s *ModifyHostsActiveAddressTypeResponseBodyResults) GetCode() *string {
	return s.Code
}

func (s *ModifyHostsActiveAddressTypeResponseBodyResults) GetHostId() *string {
	return s.HostId
}

func (s *ModifyHostsActiveAddressTypeResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *ModifyHostsActiveAddressTypeResponseBodyResults) SetCode(v string) *ModifyHostsActiveAddressTypeResponseBodyResults {
	s.Code = &v
	return s
}

func (s *ModifyHostsActiveAddressTypeResponseBodyResults) SetHostId(v string) *ModifyHostsActiveAddressTypeResponseBodyResults {
	s.HostId = &v
	return s
}

func (s *ModifyHostsActiveAddressTypeResponseBodyResults) SetMessage(v string) *ModifyHostsActiveAddressTypeResponseBodyResults {
	s.Message = &v
	return s
}

func (s *ModifyHostsActiveAddressTypeResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
