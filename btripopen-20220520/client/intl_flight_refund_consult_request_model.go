// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightRefundConsultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *IntlFlightRefundConsultRequest
	GetOrderId() *string
	SetOutOrderId(v string) *IntlFlightRefundConsultRequest
	GetOutOrderId() *string
}

type IntlFlightRefundConsultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1002145190081005400
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 3753197470069161984
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
}

func (s IntlFlightRefundConsultRequest) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightRefundConsultRequest) GoString() string {
	return s.String()
}

func (s *IntlFlightRefundConsultRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *IntlFlightRefundConsultRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *IntlFlightRefundConsultRequest) SetOrderId(v string) *IntlFlightRefundConsultRequest {
	s.OrderId = &v
	return s
}

func (s *IntlFlightRefundConsultRequest) SetOutOrderId(v string) *IntlFlightRefundConsultRequest {
	s.OutOrderId = &v
	return s
}

func (s *IntlFlightRefundConsultRequest) Validate() error {
	return dara.Validate(s)
}
