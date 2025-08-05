// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHanaInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateHanaInstanceResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateHanaInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateHanaInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateHanaInstanceResponseBody
	GetSuccess() *bool
}

type UpdateHanaInstanceResponseBody struct {
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
	// A6AB6D5A-9D21-5529-9335-A894FB045ED6
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

func (s UpdateHanaInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateHanaInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHanaInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateHanaInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateHanaInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateHanaInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateHanaInstanceResponseBody) SetCode(v string) *UpdateHanaInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateHanaInstanceResponseBody) SetMessage(v string) *UpdateHanaInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateHanaInstanceResponseBody) SetRequestId(v string) *UpdateHanaInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHanaInstanceResponseBody) SetSuccess(v bool) *UpdateHanaInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateHanaInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
