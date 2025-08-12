// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorGroupByResourceGroupIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateMonitorGroupByResourceGroupIdResponseBody
	GetCode() *string
	SetId(v int64) *CreateMonitorGroupByResourceGroupIdResponseBody
	GetId() *int64
	SetMessage(v string) *CreateMonitorGroupByResourceGroupIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateMonitorGroupByResourceGroupIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateMonitorGroupByResourceGroupIdResponseBody
	GetSuccess() *bool
}

type CreateMonitorGroupByResourceGroupIdResponseBody struct {
	// The status code.
	//
	// > The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the application group.
	//
	// example:
	//
	// 3607****
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified resource is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 784CAB3C-F613-5BCE-8469-6DCB29B18A20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMonitorGroupByResourceGroupIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorGroupByResourceGroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMonitorGroupByResourceGroupIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateMonitorGroupByResourceGroupIdResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateMonitorGroupByResourceGroupIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateMonitorGroupByResourceGroupIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMonitorGroupByResourceGroupIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMonitorGroupByResourceGroupIdResponseBody) SetCode(v string) *CreateMonitorGroupByResourceGroupIdResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMonitorGroupByResourceGroupIdResponseBody) SetId(v int64) *CreateMonitorGroupByResourceGroupIdResponseBody {
	s.Id = &v
	return s
}

func (s *CreateMonitorGroupByResourceGroupIdResponseBody) SetMessage(v string) *CreateMonitorGroupByResourceGroupIdResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMonitorGroupByResourceGroupIdResponseBody) SetRequestId(v string) *CreateMonitorGroupByResourceGroupIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMonitorGroupByResourceGroupIdResponseBody) SetSuccess(v bool) *CreateMonitorGroupByResourceGroupIdResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMonitorGroupByResourceGroupIdResponseBody) Validate() error {
	return dara.Validate(s)
}
