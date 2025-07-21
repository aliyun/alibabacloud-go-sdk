// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindPasswordFreeLoginUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UnbindPasswordFreeLoginUserResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UnbindPasswordFreeLoginUserResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UnbindPasswordFreeLoginUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *UnbindPasswordFreeLoginUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UnbindPasswordFreeLoginUserResponseBody
	GetSuccess() *bool
}

type UnbindPasswordFreeLoginUserResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnbindPasswordFreeLoginUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindPasswordFreeLoginUserResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindPasswordFreeLoginUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *UnbindPasswordFreeLoginUserResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UnbindPasswordFreeLoginUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UnbindPasswordFreeLoginUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindPasswordFreeLoginUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UnbindPasswordFreeLoginUserResponseBody) SetCode(v string) *UnbindPasswordFreeLoginUserResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindPasswordFreeLoginUserResponseBody) SetHttpStatusCode(v int32) *UnbindPasswordFreeLoginUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UnbindPasswordFreeLoginUserResponseBody) SetMessage(v string) *UnbindPasswordFreeLoginUserResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindPasswordFreeLoginUserResponseBody) SetRequestId(v string) *UnbindPasswordFreeLoginUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindPasswordFreeLoginUserResponseBody) SetSuccess(v bool) *UnbindPasswordFreeLoginUserResponseBody {
	s.Success = &v
	return s
}

func (s *UnbindPasswordFreeLoginUserResponseBody) Validate() error {
	return dara.Validate(s)
}
