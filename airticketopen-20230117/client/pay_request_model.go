// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *PayRequest
	GetAccountNo() *int64
	SetOrderNo(v string) *PayRequest
	GetOrderNo() *string
	SetTracerId(v string) *PayRequest
	GetTracerId() *string
}

type PayRequest struct {
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

func (s PayRequest) String() string {
	return dara.Prettify(s)
}

func (s PayRequest) GoString() string {
	return s.String()
}

func (s *PayRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *PayRequest) GetOrderNo() *string {
	return s.OrderNo
}

func (s *PayRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *PayRequest) SetAccountNo(v int64) *PayRequest {
	s.AccountNo = &v
	return s
}

func (s *PayRequest) SetOrderNo(v string) *PayRequest {
	s.OrderNo = &v
	return s
}

func (s *PayRequest) SetTracerId(v string) *PayRequest {
	s.TracerId = &v
	return s
}

func (s *PayRequest) Validate() error {
	return dara.Validate(s)
}
