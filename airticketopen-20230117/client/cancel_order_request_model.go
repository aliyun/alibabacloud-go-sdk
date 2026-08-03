// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *CancelOrderRequest
	GetAccountNo() *int64
	SetOrderNo(v string) *CancelOrderRequest
	GetOrderNo() *string
	SetTracerId(v string) *CancelOrderRequest
	GetTracerId() *string
}

type CancelOrderRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AccountNo *int64 `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SO202606290001
	OrderNo *string `json:"OrderNo,omitempty" xml:"OrderNo,omitempty"`
	// example:
	//
	// TraceId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s CancelOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelOrderRequest) GoString() string {
	return s.String()
}

func (s *CancelOrderRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *CancelOrderRequest) GetOrderNo() *string {
	return s.OrderNo
}

func (s *CancelOrderRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *CancelOrderRequest) SetAccountNo(v int64) *CancelOrderRequest {
	s.AccountNo = &v
	return s
}

func (s *CancelOrderRequest) SetOrderNo(v string) *CancelOrderRequest {
	s.OrderNo = &v
	return s
}

func (s *CancelOrderRequest) SetTracerId(v string) *CancelOrderRequest {
	s.TracerId = &v
	return s
}

func (s *CancelOrderRequest) Validate() error {
	return dara.Validate(s)
}
