// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFaceUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateFaceUserResponseBody
	GetCode() *string
	SetErrorMessage(v string) *UpdateFaceUserResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateFaceUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateFaceUserResponseBody
	GetSuccess() *bool
}

type UpdateFaceUserResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// success
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0A13279A-5640-45E7-87AA-83750541AD0E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateFaceUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFaceUserResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFaceUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateFaceUserResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateFaceUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFaceUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateFaceUserResponseBody) SetCode(v string) *UpdateFaceUserResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateFaceUserResponseBody) SetErrorMessage(v string) *UpdateFaceUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateFaceUserResponseBody) SetRequestId(v string) *UpdateFaceUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFaceUserResponseBody) SetSuccess(v bool) *UpdateFaceUserResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateFaceUserResponseBody) Validate() error {
	return dara.Validate(s)
}
