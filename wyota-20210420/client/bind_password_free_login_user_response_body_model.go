// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindPasswordFreeLoginUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BindPasswordFreeLoginUserResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *BindPasswordFreeLoginUserResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *BindPasswordFreeLoginUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *BindPasswordFreeLoginUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BindPasswordFreeLoginUserResponseBody
	GetSuccess() *bool
}

type BindPasswordFreeLoginUserResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindPasswordFreeLoginUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindPasswordFreeLoginUserResponseBody) GoString() string {
	return s.String()
}

func (s *BindPasswordFreeLoginUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *BindPasswordFreeLoginUserResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *BindPasswordFreeLoginUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BindPasswordFreeLoginUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindPasswordFreeLoginUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BindPasswordFreeLoginUserResponseBody) SetCode(v string) *BindPasswordFreeLoginUserResponseBody {
	s.Code = &v
	return s
}

func (s *BindPasswordFreeLoginUserResponseBody) SetHttpStatusCode(v int32) *BindPasswordFreeLoginUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BindPasswordFreeLoginUserResponseBody) SetMessage(v string) *BindPasswordFreeLoginUserResponseBody {
	s.Message = &v
	return s
}

func (s *BindPasswordFreeLoginUserResponseBody) SetRequestId(v string) *BindPasswordFreeLoginUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindPasswordFreeLoginUserResponseBody) SetSuccess(v bool) *BindPasswordFreeLoginUserResponseBody {
	s.Success = &v
	return s
}

func (s *BindPasswordFreeLoginUserResponseBody) Validate() error {
	return dara.Validate(s)
}
