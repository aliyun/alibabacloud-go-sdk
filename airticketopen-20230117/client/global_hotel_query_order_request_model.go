// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalHotelQueryOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *GlobalHotelQueryOrderRequest
	GetAccountNo() *int64
	SetExternalOrderNo(v string) *GlobalHotelQueryOrderRequest
	GetExternalOrderNo() *string
	SetOrderNo(v string) *GlobalHotelQueryOrderRequest
	GetOrderNo() *string
	SetTracerId(v string) *GlobalHotelQueryOrderRequest
	GetTracerId() *string
}

type GlobalHotelQueryOrderRequest struct {
	// The distributor account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AccountNo *int64 `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
	// The external order number. You must specify at least one of ExternalOrderNo and OrderNo.
	//
	// example:
	//
	// EXT_ORDER_001
	ExternalOrderNo *string `json:"ExternalOrderNo,omitempty" xml:"ExternalOrderNo,omitempty"`
	// The sales order number. You must specify at least one of OrderNo and ExternalOrderNo.
	//
	// example:
	//
	// SO202606290001
	OrderNo *string `json:"OrderNo,omitempty" xml:"OrderNo,omitempty"`
	// TracerId
	//
	// example:
	//
	// TracerId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s GlobalHotelQueryOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s GlobalHotelQueryOrderRequest) GoString() string {
	return s.String()
}

func (s *GlobalHotelQueryOrderRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *GlobalHotelQueryOrderRequest) GetExternalOrderNo() *string {
	return s.ExternalOrderNo
}

func (s *GlobalHotelQueryOrderRequest) GetOrderNo() *string {
	return s.OrderNo
}

func (s *GlobalHotelQueryOrderRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *GlobalHotelQueryOrderRequest) SetAccountNo(v int64) *GlobalHotelQueryOrderRequest {
	s.AccountNo = &v
	return s
}

func (s *GlobalHotelQueryOrderRequest) SetExternalOrderNo(v string) *GlobalHotelQueryOrderRequest {
	s.ExternalOrderNo = &v
	return s
}

func (s *GlobalHotelQueryOrderRequest) SetOrderNo(v string) *GlobalHotelQueryOrderRequest {
	s.OrderNo = &v
	return s
}

func (s *GlobalHotelQueryOrderRequest) SetTracerId(v string) *GlobalHotelQueryOrderRequest {
	s.TracerId = &v
	return s
}

func (s *GlobalHotelQueryOrderRequest) Validate() error {
	return dara.Validate(s)
}
