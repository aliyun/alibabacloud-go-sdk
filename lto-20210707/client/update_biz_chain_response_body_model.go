// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBizChainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateBizChainResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateBizChainResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateBizChainResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateBizChainResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateBizChainResponseBody
	GetSuccess() *bool
}

type UpdateBizChainResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateBizChainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateBizChainResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBizChainResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateBizChainResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateBizChainResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateBizChainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateBizChainResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateBizChainResponseBody) SetCode(v string) *UpdateBizChainResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateBizChainResponseBody) SetHttpStatusCode(v int32) *UpdateBizChainResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateBizChainResponseBody) SetMessage(v string) *UpdateBizChainResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateBizChainResponseBody) SetRequestId(v string) *UpdateBizChainResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBizChainResponseBody) SetSuccess(v bool) *UpdateBizChainResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateBizChainResponseBody) Validate() error {
	return dara.Validate(s)
}
