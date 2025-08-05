// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CheckRoleResponseBody
	GetCode() *string
	SetMessage(v string) *CheckRoleResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckRoleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CheckRoleResponseBody
	GetSuccess() *bool
}

type CheckRoleResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message. If the request was successful, "successful" is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2F63CA9B-744E-51C0-A638-27882BB03078
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

func (s CheckRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CheckRoleResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckRoleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckRoleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckRoleResponseBody) SetCode(v string) *CheckRoleResponseBody {
	s.Code = &v
	return s
}

func (s *CheckRoleResponseBody) SetMessage(v string) *CheckRoleResponseBody {
	s.Message = &v
	return s
}

func (s *CheckRoleResponseBody) SetRequestId(v string) *CheckRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckRoleResponseBody) SetSuccess(v bool) *CheckRoleResponseBody {
	s.Success = &v
	return s
}

func (s *CheckRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
