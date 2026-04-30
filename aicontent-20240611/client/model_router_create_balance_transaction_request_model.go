// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateBalanceTransactionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v float64) *ModelRouterCreateBalanceTransactionRequest
	GetAmount() *float64
	SetRemark(v string) *ModelRouterCreateBalanceTransactionRequest
	GetRemark() *string
	SetType(v string) *ModelRouterCreateBalanceTransactionRequest
	GetType() *string
}

type ModelRouterCreateBalanceTransactionRequest struct {
	// example:
	//
	// 100.00
	Amount *float64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// example:
	//
	// 充值
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// example:
	//
	// recharge
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ModelRouterCreateBalanceTransactionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateBalanceTransactionRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateBalanceTransactionRequest) GetAmount() *float64 {
	return s.Amount
}

func (s *ModelRouterCreateBalanceTransactionRequest) GetRemark() *string {
	return s.Remark
}

func (s *ModelRouterCreateBalanceTransactionRequest) GetType() *string {
	return s.Type
}

func (s *ModelRouterCreateBalanceTransactionRequest) SetAmount(v float64) *ModelRouterCreateBalanceTransactionRequest {
	s.Amount = &v
	return s
}

func (s *ModelRouterCreateBalanceTransactionRequest) SetRemark(v string) *ModelRouterCreateBalanceTransactionRequest {
	s.Remark = &v
	return s
}

func (s *ModelRouterCreateBalanceTransactionRequest) SetType(v string) *ModelRouterCreateBalanceTransactionRequest {
	s.Type = &v
	return s
}

func (s *ModelRouterCreateBalanceTransactionRequest) Validate() error {
	return dara.Validate(s)
}
