// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelOrRefundRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountNo(v int64) *CancelOrRefundRequest
	GetAccountNo() *int64
	SetOrderNo(v string) *CancelOrRefundRequest
	GetOrderNo() *string
	SetTracerId(v string) *CancelOrRefundRequest
	GetTracerId() *string
}

type CancelOrRefundRequest struct {
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

func (s CancelOrRefundRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelOrRefundRequest) GoString() string {
	return s.String()
}

func (s *CancelOrRefundRequest) GetAccountNo() *int64 {
	return s.AccountNo
}

func (s *CancelOrRefundRequest) GetOrderNo() *string {
	return s.OrderNo
}

func (s *CancelOrRefundRequest) GetTracerId() *string {
	return s.TracerId
}

func (s *CancelOrRefundRequest) SetAccountNo(v int64) *CancelOrRefundRequest {
	s.AccountNo = &v
	return s
}

func (s *CancelOrRefundRequest) SetOrderNo(v string) *CancelOrRefundRequest {
	s.OrderNo = &v
	return s
}

func (s *CancelOrRefundRequest) SetTracerId(v string) *CancelOrRefundRequest {
	s.TracerId = &v
	return s
}

func (s *CancelOrRefundRequest) Validate() error {
	return dara.Validate(s)
}
