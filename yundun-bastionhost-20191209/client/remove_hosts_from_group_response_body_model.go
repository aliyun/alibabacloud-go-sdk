// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveHostsFromGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveHostsFromGroupResponseBody
	GetRequestId() *string
	SetResults(v []*RemoveHostsFromGroupResponseBodyResults) *RemoveHostsFromGroupResponseBody
	GetResults() []*RemoveHostsFromGroupResponseBodyResults
}

type RemoveHostsFromGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*RemoveHostsFromGroupResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s RemoveHostsFromGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveHostsFromGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveHostsFromGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveHostsFromGroupResponseBody) GetResults() []*RemoveHostsFromGroupResponseBodyResults {
	return s.Results
}

func (s *RemoveHostsFromGroupResponseBody) SetRequestId(v string) *RemoveHostsFromGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveHostsFromGroupResponseBody) SetResults(v []*RemoveHostsFromGroupResponseBodyResults) *RemoveHostsFromGroupResponseBody {
	s.Results = v
	return s
}

func (s *RemoveHostsFromGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type RemoveHostsFromGroupResponseBodyResults struct {
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
	//     > Make sure that the specified bastion host ID and host IDs are valid and call the operation again.
	//
	// 	- **OBJECT_AlREADY_EXISTS**: The specified object on which you want to perform the operation already exists.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the asset group.
	//
	// example:
	//
	// 1
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
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
	// 无
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RemoveHostsFromGroupResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s RemoveHostsFromGroupResponseBodyResults) GoString() string {
	return s.String()
}

func (s *RemoveHostsFromGroupResponseBodyResults) GetCode() *string {
	return s.Code
}

func (s *RemoveHostsFromGroupResponseBodyResults) GetHostGroupId() *string {
	return s.HostGroupId
}

func (s *RemoveHostsFromGroupResponseBodyResults) GetHostId() *string {
	return s.HostId
}

func (s *RemoveHostsFromGroupResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *RemoveHostsFromGroupResponseBodyResults) SetCode(v string) *RemoveHostsFromGroupResponseBodyResults {
	s.Code = &v
	return s
}

func (s *RemoveHostsFromGroupResponseBodyResults) SetHostGroupId(v string) *RemoveHostsFromGroupResponseBodyResults {
	s.HostGroupId = &v
	return s
}

func (s *RemoveHostsFromGroupResponseBodyResults) SetHostId(v string) *RemoveHostsFromGroupResponseBodyResults {
	s.HostId = &v
	return s
}

func (s *RemoveHostsFromGroupResponseBodyResults) SetMessage(v string) *RemoveHostsFromGroupResponseBodyResults {
	s.Message = &v
	return s
}

func (s *RemoveHostsFromGroupResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
