// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddHostsToGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddHostsToGroupResponseBody
	GetRequestId() *string
	SetResults(v []*AddHostsToGroupResponseBodyResults) *AddHostsToGroupResponseBody
	GetResults() []*AddHostsToGroupResponseBodyResults
}

type AddHostsToGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 81500666-d7f5-4143-8329-0223cc738105
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*AddHostsToGroupResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s AddHostsToGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddHostsToGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddHostsToGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddHostsToGroupResponseBody) GetResults() []*AddHostsToGroupResponseBodyResults {
	return s.Results
}

func (s *AddHostsToGroupResponseBody) SetRequestId(v string) *AddHostsToGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddHostsToGroupResponseBody) SetResults(v []*AddHostsToGroupResponseBodyResults) *AddHostsToGroupResponseBody {
	s.Results = v
	return s
}

func (s *AddHostsToGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddHostsToGroupResponseBodyResults struct {
	// The return code that indicates whether the call was successful. Valid values:
	//
	// 	- **OK**: The call was successful.
	//
	// 	- **UNEXPECTED**: An unknown error occurred.
	//
	// 	- **INVALID_ARGUMENT**: A request parameter is invalid.
	//
	//     >Make sure that the request parameters are valid and call the operation again.
	//
	// 	- **OBJECT_NOT_FOUND**: The specified object on which you want to perform the operation does not exist.
	//
	//     > Make sure that the specified bastion host ID and host IDs are valid. Then, call the operation again.
	//
	// 	- **OBJECT_AlREADY_EXISTS**: The specified object on which you want to perform the operation already exists.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The asset group ID.
	//
	// example:
	//
	// 1
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// The host ID.
	//
	// example:
	//
	// 1
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// N/A
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s AddHostsToGroupResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s AddHostsToGroupResponseBodyResults) GoString() string {
	return s.String()
}

func (s *AddHostsToGroupResponseBodyResults) GetCode() *string {
	return s.Code
}

func (s *AddHostsToGroupResponseBodyResults) GetHostGroupId() *string {
	return s.HostGroupId
}

func (s *AddHostsToGroupResponseBodyResults) GetHostId() *string {
	return s.HostId
}

func (s *AddHostsToGroupResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *AddHostsToGroupResponseBodyResults) SetCode(v string) *AddHostsToGroupResponseBodyResults {
	s.Code = &v
	return s
}

func (s *AddHostsToGroupResponseBodyResults) SetHostGroupId(v string) *AddHostsToGroupResponseBodyResults {
	s.HostGroupId = &v
	return s
}

func (s *AddHostsToGroupResponseBodyResults) SetHostId(v string) *AddHostsToGroupResponseBodyResults {
	s.HostId = &v
	return s
}

func (s *AddHostsToGroupResponseBodyResults) SetMessage(v string) *AddHostsToGroupResponseBodyResults {
	s.Message = &v
	return s
}

func (s *AddHostsToGroupResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
