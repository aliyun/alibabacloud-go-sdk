// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDowngradePrePayOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DowngradePrePayOrderResponseBody
	GetCode() *int32
	SetMessage(v string) *DowngradePrePayOrderResponseBody
	GetMessage() *string
	SetOrderId(v string) *DowngradePrePayOrderResponseBody
	GetOrderId() *string
	SetRequestId(v string) *DowngradePrePayOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DowngradePrePayOrderResponseBody
	GetSuccess() *bool
}

type DowngradePrePayOrderResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DowngradePrePayOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DowngradePrePayOrderResponseBody) GoString() string {
	return s.String()
}

func (s *DowngradePrePayOrderResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DowngradePrePayOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DowngradePrePayOrderResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *DowngradePrePayOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DowngradePrePayOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DowngradePrePayOrderResponseBody) SetCode(v int32) *DowngradePrePayOrderResponseBody {
	s.Code = &v
	return s
}

func (s *DowngradePrePayOrderResponseBody) SetMessage(v string) *DowngradePrePayOrderResponseBody {
	s.Message = &v
	return s
}

func (s *DowngradePrePayOrderResponseBody) SetOrderId(v string) *DowngradePrePayOrderResponseBody {
	s.OrderId = &v
	return s
}

func (s *DowngradePrePayOrderResponseBody) SetRequestId(v string) *DowngradePrePayOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DowngradePrePayOrderResponseBody) SetSuccess(v bool) *DowngradePrePayOrderResponseBody {
	s.Success = &v
	return s
}

func (s *DowngradePrePayOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
