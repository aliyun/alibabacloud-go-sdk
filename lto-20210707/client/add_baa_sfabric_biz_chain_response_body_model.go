// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBaaSFabricBizChainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddBaaSFabricBizChainResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *AddBaaSFabricBizChainResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddBaaSFabricBizChainResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddBaaSFabricBizChainResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddBaaSFabricBizChainResponseBody
	GetSuccess() *bool
}

type AddBaaSFabricBizChainResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddBaaSFabricBizChainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddBaaSFabricBizChainResponseBody) GoString() string {
	return s.String()
}

func (s *AddBaaSFabricBizChainResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddBaaSFabricBizChainResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddBaaSFabricBizChainResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddBaaSFabricBizChainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddBaaSFabricBizChainResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddBaaSFabricBizChainResponseBody) SetCode(v string) *AddBaaSFabricBizChainResponseBody {
	s.Code = &v
	return s
}

func (s *AddBaaSFabricBizChainResponseBody) SetHttpStatusCode(v int32) *AddBaaSFabricBizChainResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddBaaSFabricBizChainResponseBody) SetMessage(v string) *AddBaaSFabricBizChainResponseBody {
	s.Message = &v
	return s
}

func (s *AddBaaSFabricBizChainResponseBody) SetRequestId(v string) *AddBaaSFabricBizChainResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddBaaSFabricBizChainResponseBody) SetSuccess(v bool) *AddBaaSFabricBizChainResponseBody {
	s.Success = &v
	return s
}

func (s *AddBaaSFabricBizChainResponseBody) Validate() error {
	return dara.Validate(s)
}
