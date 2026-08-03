// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *QueryOrderRequest
	GetAccountNo() *int64
	SetExternalOrderNo(v string) *QueryOrderRequest
	GetExternalOrderNo() *string
	SetOrderNo(v string) *QueryOrderRequest
	GetOrderNo() *string
	SetTracerId(v string) *QueryOrderRequest
	GetTracerId() *string
}

type QueryOrderRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AccountNo *int64 `json:"AccountNo,omitempty" xml:"AccountNo,omitempty"`
	// example:
	//
	// EXT_ORDER_001
	ExternalOrderNo *string `json:"ExternalOrderNo,omitempty" xml:"ExternalOrderNo,omitempty"`
	// example:
	//
	// SO202606290001
	OrderNo *string `json:"OrderNo,omitempty" xml:"OrderNo,omitempty"`
	// example:
	//
	// TracerId
	TracerId *string `json:"TracerId,omitempty" xml:"TracerId,omitempty"`
}

func (s QueryOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryOrderRequest) GoString() string {
	return s.String()
}

func (s *QueryOrderRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *QueryOrderRequest) GetExternalOrderNo() *string {
	return s.ExternalOrderNo
}

func (s *QueryOrderRequest) GetOrderNo() *string {
	return s.OrderNo
}

func (s *QueryOrderRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *QueryOrderRequest) SetAccountNo(v int64) *QueryOrderRequest {
	s.AccountNo = &v
	return s
}

func (s *QueryOrderRequest) SetExternalOrderNo(v string) *QueryOrderRequest {
	s.ExternalOrderNo = &v
	return s
}

func (s *QueryOrderRequest) SetOrderNo(v string) *QueryOrderRequest {
	s.OrderNo = &v
	return s
}

func (s *QueryOrderRequest) SetTracerId(v string) *QueryOrderRequest {
	s.TracerId = &v
	return s
}

func (s *QueryOrderRequest) Validate() error {
	return dara.Validate(s)
}
