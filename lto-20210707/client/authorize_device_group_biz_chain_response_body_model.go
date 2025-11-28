// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeDeviceGroupBizChainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AuthorizeDeviceGroupBizChainResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *AuthorizeDeviceGroupBizChainResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AuthorizeDeviceGroupBizChainResponseBody
	GetMessage() *string
	SetRequestId(v string) *AuthorizeDeviceGroupBizChainResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AuthorizeDeviceGroupBizChainResponseBody
	GetSuccess() *bool
}

type AuthorizeDeviceGroupBizChainResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AuthorizeDeviceGroupBizChainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeDeviceGroupBizChainResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeDeviceGroupBizChainResponseBody) GetCode() *string {
	return s.Code
}

func (s *AuthorizeDeviceGroupBizChainResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AuthorizeDeviceGroupBizChainResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AuthorizeDeviceGroupBizChainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AuthorizeDeviceGroupBizChainResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AuthorizeDeviceGroupBizChainResponseBody) SetCode(v string) *AuthorizeDeviceGroupBizChainResponseBody {
	s.Code = &v
	return s
}

func (s *AuthorizeDeviceGroupBizChainResponseBody) SetHttpStatusCode(v int32) *AuthorizeDeviceGroupBizChainResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AuthorizeDeviceGroupBizChainResponseBody) SetMessage(v string) *AuthorizeDeviceGroupBizChainResponseBody {
	s.Message = &v
	return s
}

func (s *AuthorizeDeviceGroupBizChainResponseBody) SetRequestId(v string) *AuthorizeDeviceGroupBizChainResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthorizeDeviceGroupBizChainResponseBody) SetSuccess(v bool) *AuthorizeDeviceGroupBizChainResponseBody {
	s.Success = &v
	return s
}

func (s *AuthorizeDeviceGroupBizChainResponseBody) Validate() error {
	return dara.Validate(s)
}
