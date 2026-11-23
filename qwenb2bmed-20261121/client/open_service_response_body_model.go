// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *OpenServiceResponseBody
	GetOrderId() *string
	SetSpInstanceId(v string) *OpenServiceResponseBody
	GetSpInstanceId() *string
	SetStatus(v string) *OpenServiceResponseBody
	GetStatus() *string
}

type OpenServiceResponseBody struct {
	// 关联本次开通申请的订单 ID，存在时返回。
	//
	// example:
	//
	// 2000000000000001
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// 服务已开通时返回的实例标识。
	//
	// example:
	//
	// 1234567890123456
	SpInstanceId *string `json:"spInstanceId,omitempty" xml:"spInstanceId,omitempty"`
	// TEST_OK 表示仅验证调用成功。OPENING 表示开通处理中。ACTIVE 表示已开通。ORDER_UNKNOWN 表示暂时无法确认开通结果。REJECTED 表示开通请求未获通过。
	//
	// This parameter is required.
	//
	// example:
	//
	// TEST_OK
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s OpenServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenServiceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *OpenServiceResponseBody) GetSpInstanceId() *string {
	return s.SpInstanceId
}

func (s *OpenServiceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *OpenServiceResponseBody) SetOrderId(v string) *OpenServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenServiceResponseBody) SetSpInstanceId(v string) *OpenServiceResponseBody {
	s.SpInstanceId = &v
	return s
}

func (s *OpenServiceResponseBody) SetStatus(v string) *OpenServiceResponseBody {
	s.Status = &v
	return s
}

func (s *OpenServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
