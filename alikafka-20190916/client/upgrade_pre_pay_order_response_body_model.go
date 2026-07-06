// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradePrePayOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpgradePrePayOrderResponseBody
	GetCode() *int32
	SetMessage(v string) *UpgradePrePayOrderResponseBody
	GetMessage() *string
	SetOrderId(v string) *UpgradePrePayOrderResponseBody
	GetOrderId() *string
	SetRequestId(v string) *UpgradePrePayOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpgradePrePayOrderResponseBody
	GetSuccess() *bool
}

type UpgradePrePayOrderResponseBody struct {
	// A return code of 200 indicates that the request succeeded.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 20497346575****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABA4A7FD-E10F-45C7-9774-A5236015***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpgradePrePayOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradePrePayOrderResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradePrePayOrderResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpgradePrePayOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpgradePrePayOrderResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *UpgradePrePayOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradePrePayOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpgradePrePayOrderResponseBody) SetCode(v int32) *UpgradePrePayOrderResponseBody {
	s.Code = &v
	return s
}

func (s *UpgradePrePayOrderResponseBody) SetMessage(v string) *UpgradePrePayOrderResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradePrePayOrderResponseBody) SetOrderId(v string) *UpgradePrePayOrderResponseBody {
	s.OrderId = &v
	return s
}

func (s *UpgradePrePayOrderResponseBody) SetRequestId(v string) *UpgradePrePayOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradePrePayOrderResponseBody) SetSuccess(v bool) *UpgradePrePayOrderResponseBody {
	s.Success = &v
	return s
}

func (s *UpgradePrePayOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
