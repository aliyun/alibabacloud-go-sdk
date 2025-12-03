// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrderSummaryForPartnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *GetOrderSummaryForPartnerRequest
	GetOrderId() *string
}

type GetOrderSummaryForPartnerRequest struct {
	// example:
	//
	// 1672369049358
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s GetOrderSummaryForPartnerRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOrderSummaryForPartnerRequest) GoString() string {
	return s.String()
}

func (s *GetOrderSummaryForPartnerRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *GetOrderSummaryForPartnerRequest) SetOrderId(v string) *GetOrderSummaryForPartnerRequest {
	s.OrderId = &v
	return s
}

func (s *GetOrderSummaryForPartnerRequest) Validate() error {
	return dara.Validate(s)
}
