// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBaaSAntChainBizChainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddBaaSAntChainBizChainResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *AddBaaSAntChainBizChainResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddBaaSAntChainBizChainResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddBaaSAntChainBizChainResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddBaaSAntChainBizChainResponseBody
	GetSuccess() *bool
}

type AddBaaSAntChainBizChainResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddBaaSAntChainBizChainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddBaaSAntChainBizChainResponseBody) GoString() string {
	return s.String()
}

func (s *AddBaaSAntChainBizChainResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddBaaSAntChainBizChainResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddBaaSAntChainBizChainResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddBaaSAntChainBizChainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddBaaSAntChainBizChainResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddBaaSAntChainBizChainResponseBody) SetCode(v string) *AddBaaSAntChainBizChainResponseBody {
	s.Code = &v
	return s
}

func (s *AddBaaSAntChainBizChainResponseBody) SetHttpStatusCode(v int32) *AddBaaSAntChainBizChainResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddBaaSAntChainBizChainResponseBody) SetMessage(v string) *AddBaaSAntChainBizChainResponseBody {
	s.Message = &v
	return s
}

func (s *AddBaaSAntChainBizChainResponseBody) SetRequestId(v string) *AddBaaSAntChainBizChainResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddBaaSAntChainBizChainResponseBody) SetSuccess(v bool) *AddBaaSAntChainBizChainResponseBody {
	s.Success = &v
	return s
}

func (s *AddBaaSAntChainBizChainResponseBody) Validate() error {
	return dara.Validate(s)
}
