// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateOrderResponseBody
	GetCode() *string
	SetData(v string) *CreateOrderResponseBody
	GetData() *string
	SetMessage(v string) *CreateOrderResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateOrderResponseBody
	GetSuccess() *bool
}

type CreateOrderResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrderResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateOrderResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateOrderResponseBody) SetCode(v string) *CreateOrderResponseBody {
	s.Code = &v
	return s
}

func (s *CreateOrderResponseBody) SetData(v string) *CreateOrderResponseBody {
	s.Data = &v
	return s
}

func (s *CreateOrderResponseBody) SetMessage(v string) *CreateOrderResponseBody {
	s.Message = &v
	return s
}

func (s *CreateOrderResponseBody) SetRequestId(v string) *CreateOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrderResponseBody) SetSuccess(v bool) *CreateOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CreateOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
