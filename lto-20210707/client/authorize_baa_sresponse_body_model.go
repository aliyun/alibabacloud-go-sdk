// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeBaaSResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AuthorizeBaaSResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *AuthorizeBaaSResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AuthorizeBaaSResponseBody
	GetMessage() *string
	SetRequestId(v string) *AuthorizeBaaSResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AuthorizeBaaSResponseBody
	GetSuccess() *bool
}

type AuthorizeBaaSResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AuthorizeBaaSResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeBaaSResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeBaaSResponseBody) GetCode() *string {
	return s.Code
}

func (s *AuthorizeBaaSResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AuthorizeBaaSResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AuthorizeBaaSResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthorizeBaaSResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AuthorizeBaaSResponseBody) SetCode(v string) *AuthorizeBaaSResponseBody {
	s.Code = &v
	return s
}

func (s *AuthorizeBaaSResponseBody) SetHttpStatusCode(v int32) *AuthorizeBaaSResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AuthorizeBaaSResponseBody) SetMessage(v string) *AuthorizeBaaSResponseBody {
	s.Message = &v
	return s
}

func (s *AuthorizeBaaSResponseBody) SetRequestId(v string) *AuthorizeBaaSResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthorizeBaaSResponseBody) SetSuccess(v bool) *AuthorizeBaaSResponseBody {
	s.Success = &v
	return s
}

func (s *AuthorizeBaaSResponseBody) Validate() error {
	return dara.Validate(s)
}
