// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateUserResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateUserResponseBody
	GetRequestId() *string
	SetWnUserId(v string) *UpdateUserResponseBody
	GetWnUserId() *string
}

type UpdateUserResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The status code description.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The WINNEXO platform user ID.
	//
	// example:
	//
	// 1
	WnUserId *string `json:"wnUserId,omitempty" xml:"wnUserId,omitempty"`
}

func (s UpdateUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUserResponseBody) GetWnUserId() *string {
	return s.WnUserId
}

func (s *UpdateUserResponseBody) SetCode(v string) *UpdateUserResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateUserResponseBody) SetMessage(v string) *UpdateUserResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateUserResponseBody) SetRequestId(v string) *UpdateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserResponseBody) SetWnUserId(v string) *UpdateUserResponseBody {
	s.WnUserId = &v
	return s
}

func (s *UpdateUserResponseBody) Validate() error {
	return dara.Validate(s)
}
