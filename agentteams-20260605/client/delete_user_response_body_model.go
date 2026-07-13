// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteUserResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteUserResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteUserResponseBody
	GetSuccess() *bool
}

type DeleteUserResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteUserResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteUserResponseBody) SetCode(v string) *DeleteUserResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteUserResponseBody) SetHttpStatusCode(v int32) *DeleteUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteUserResponseBody) SetMessage(v string) *DeleteUserResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteUserResponseBody) SetRequestId(v string) *DeleteUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserResponseBody) SetSuccess(v bool) *DeleteUserResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteUserResponseBody) Validate() error {
	return dara.Validate(s)
}
