// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBsnFabricBizChainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddBsnFabricBizChainResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *AddBsnFabricBizChainResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddBsnFabricBizChainResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddBsnFabricBizChainResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddBsnFabricBizChainResponseBody
	GetSuccess() *bool
}

type AddBsnFabricBizChainResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddBsnFabricBizChainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddBsnFabricBizChainResponseBody) GoString() string {
	return s.String()
}

func (s *AddBsnFabricBizChainResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddBsnFabricBizChainResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddBsnFabricBizChainResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddBsnFabricBizChainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddBsnFabricBizChainResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddBsnFabricBizChainResponseBody) SetCode(v string) *AddBsnFabricBizChainResponseBody {
	s.Code = &v
	return s
}

func (s *AddBsnFabricBizChainResponseBody) SetHttpStatusCode(v int32) *AddBsnFabricBizChainResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddBsnFabricBizChainResponseBody) SetMessage(v string) *AddBsnFabricBizChainResponseBody {
	s.Message = &v
	return s
}

func (s *AddBsnFabricBizChainResponseBody) SetRequestId(v string) *AddBsnFabricBizChainResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddBsnFabricBizChainResponseBody) SetSuccess(v bool) *AddBsnFabricBizChainResponseBody {
	s.Success = &v
	return s
}

func (s *AddBsnFabricBizChainResponseBody) Validate() error {
	return dara.Validate(s)
}
