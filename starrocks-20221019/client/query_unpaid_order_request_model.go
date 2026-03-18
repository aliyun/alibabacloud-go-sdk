// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUnpaidOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillingInstanceId(v string) *QueryUnpaidOrderRequest
	GetBillingInstanceId() *string
	SetInstanceId(v string) *QueryUnpaidOrderRequest
	GetInstanceId() *string
	SetOrderType(v string) *QueryUnpaidOrderRequest
	GetOrderType() *string
}

type QueryUnpaidOrderRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ng-12zxs328sha2
	BillingInstanceId *string `json:"BillingInstanceId,omitempty" xml:"BillingInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// BUY
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
}

func (s QueryUnpaidOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryUnpaidOrderRequest) GoString() string {
	return s.String()
}

func (s *QueryUnpaidOrderRequest) GetBillingInstanceId() *string {
	return s.BillingInstanceId
}

func (s *QueryUnpaidOrderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryUnpaidOrderRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *QueryUnpaidOrderRequest) SetBillingInstanceId(v string) *QueryUnpaidOrderRequest {
	s.BillingInstanceId = &v
	return s
}

func (s *QueryUnpaidOrderRequest) SetInstanceId(v string) *QueryUnpaidOrderRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryUnpaidOrderRequest) SetOrderType(v string) *QueryUnpaidOrderRequest {
	s.OrderType = &v
	return s
}

func (s *QueryUnpaidOrderRequest) Validate() error {
	return dara.Validate(s)
}
