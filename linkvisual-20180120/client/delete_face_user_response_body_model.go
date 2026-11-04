// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteFaceUserResponseBody
	GetCode() *string
	SetErrorMessage(v string) *DeleteFaceUserResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteFaceUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteFaceUserResponseBody
	GetSuccess() *bool
}

type DeleteFaceUserResponseBody struct {
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
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFaceUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteFaceUserResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteFaceUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFaceUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteFaceUserResponseBody) SetCode(v string) *DeleteFaceUserResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFaceUserResponseBody) SetErrorMessage(v string) *DeleteFaceUserResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteFaceUserResponseBody) SetRequestId(v string) *DeleteFaceUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFaceUserResponseBody) SetSuccess(v bool) *DeleteFaceUserResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteFaceUserResponseBody) Validate() error {
	return dara.Validate(s)
}
