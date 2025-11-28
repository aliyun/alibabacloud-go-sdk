// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeMemberBizChainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AuthorizeMemberBizChainResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *AuthorizeMemberBizChainResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AuthorizeMemberBizChainResponseBody
	GetMessage() *string
	SetRequestId(v string) *AuthorizeMemberBizChainResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AuthorizeMemberBizChainResponseBody
	GetSuccess() *bool
}

type AuthorizeMemberBizChainResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AuthorizeMemberBizChainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeMemberBizChainResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeMemberBizChainResponseBody) GetCode() *string {
	return s.Code
}

func (s *AuthorizeMemberBizChainResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AuthorizeMemberBizChainResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AuthorizeMemberBizChainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthorizeMemberBizChainResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AuthorizeMemberBizChainResponseBody) SetCode(v string) *AuthorizeMemberBizChainResponseBody {
	s.Code = &v
	return s
}

func (s *AuthorizeMemberBizChainResponseBody) SetHttpStatusCode(v int32) *AuthorizeMemberBizChainResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AuthorizeMemberBizChainResponseBody) SetMessage(v string) *AuthorizeMemberBizChainResponseBody {
	s.Message = &v
	return s
}

func (s *AuthorizeMemberBizChainResponseBody) SetRequestId(v string) *AuthorizeMemberBizChainResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthorizeMemberBizChainResponseBody) SetSuccess(v bool) *AuthorizeMemberBizChainResponseBody {
	s.Success = &v
	return s
}

func (s *AuthorizeMemberBizChainResponseBody) Validate() error {
	return dara.Validate(s)
}
