// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindIdentityProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UnbindIdentityProviderResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UnbindIdentityProviderResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UnbindIdentityProviderResponseBody
	GetMessage() *string
	SetRequestId(v string) *UnbindIdentityProviderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UnbindIdentityProviderResponseBody
	GetSuccess() *bool
}

type UnbindIdentityProviderResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnbindIdentityProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindIdentityProviderResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindIdentityProviderResponseBody) GetCode() *string {
	return s.Code
}

func (s *UnbindIdentityProviderResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UnbindIdentityProviderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UnbindIdentityProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindIdentityProviderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UnbindIdentityProviderResponseBody) SetCode(v string) *UnbindIdentityProviderResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindIdentityProviderResponseBody) SetHttpStatusCode(v int32) *UnbindIdentityProviderResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UnbindIdentityProviderResponseBody) SetMessage(v string) *UnbindIdentityProviderResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindIdentityProviderResponseBody) SetRequestId(v string) *UnbindIdentityProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindIdentityProviderResponseBody) SetSuccess(v bool) *UnbindIdentityProviderResponseBody {
	s.Success = &v
	return s
}

func (s *UnbindIdentityProviderResponseBody) Validate() error {
	return dara.Validate(s)
}
