// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteUserGroupResponseBody
	GetCode() *string
	SetData(v bool) *DeleteUserGroupResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *DeleteUserGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteUserGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteUserGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteUserGroupResponseBody
	GetSuccess() *bool
}

type DeleteUserGroupResponseBody struct {
	// The error code. A value of OK indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The deletion result.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code returned by the backend.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteUserGroupResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteUserGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteUserGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUserGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteUserGroupResponseBody) SetCode(v string) *DeleteUserGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteUserGroupResponseBody) SetData(v bool) *DeleteUserGroupResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteUserGroupResponseBody) SetHttpStatusCode(v int32) *DeleteUserGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteUserGroupResponseBody) SetMessage(v string) *DeleteUserGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteUserGroupResponseBody) SetRequestId(v string) *DeleteUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserGroupResponseBody) SetSuccess(v bool) *DeleteUserGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteUserGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
