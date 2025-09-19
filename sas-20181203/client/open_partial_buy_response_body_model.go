// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenPartialBuyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *OpenPartialBuyResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *OpenPartialBuyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *OpenPartialBuyResponseBody
	GetMessage() *string
	SetRequestId(v string) *OpenPartialBuyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OpenPartialBuyResponseBody
	GetSuccess() *bool
}

type OpenPartialBuyResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OpenPartialBuyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenPartialBuyResponseBody) GoString() string {
	return s.String()
}

func (s *OpenPartialBuyResponseBody) GetCode() *string {
	return s.Code
}

func (s *OpenPartialBuyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *OpenPartialBuyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OpenPartialBuyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenPartialBuyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OpenPartialBuyResponseBody) SetCode(v string) *OpenPartialBuyResponseBody {
	s.Code = &v
	return s
}

func (s *OpenPartialBuyResponseBody) SetHttpStatusCode(v int32) *OpenPartialBuyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *OpenPartialBuyResponseBody) SetMessage(v string) *OpenPartialBuyResponseBody {
	s.Message = &v
	return s
}

func (s *OpenPartialBuyResponseBody) SetRequestId(v string) *OpenPartialBuyResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenPartialBuyResponseBody) SetSuccess(v bool) *OpenPartialBuyResponseBody {
	s.Success = &v
	return s
}

func (s *OpenPartialBuyResponseBody) Validate() error {
	return dara.Validate(s)
}
