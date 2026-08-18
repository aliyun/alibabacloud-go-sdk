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
	SetBalanceType(v string) *ModelRouterCreateBalanceTransactionRequest
	GetBalanceType() *string
	SetIdempotencyKey(v string) *ModelRouterCreateBalanceTransactionRequest
	GetIdempotencyKey() *string
	SetRemark(v string) *ModelRouterCreateBalanceTransactionRequest
	GetRemark() *string
	SetType(v string) *ModelRouterCreateBalanceTransactionRequest
	GetType() *string
}

type ModelRouterCreateBalanceTransactionRequest struct {
	// The transaction amount.
	//
	// example:
	//
	// 100.00
	Amount *float64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// The target balance pool type. If not specified, the default value is permanent. Valid values: permanent: permanent balance pool (the balance never expires). monthly: monthly balance pool (the balance is automatically cleared at the beginning of each month).
	//
	// example:
	//
	// amount
	BalanceType *string `json:"balanceType,omitempty" xml:"balanceType,omitempty"`
	// The idempotency key. UUID v4 format is recommended. The maximum length is 32 characters. Duplicate submissions with the same key are not executed more than once.
	//
	// example:
	//
	// 550e8400e29b41d4a716446655440000
	IdempotencyKey *string `json:"idempotencyKey,omitempty" xml:"idempotencyKey,omitempty"`
	// The remarks.
	//
	// example:
	//
	// Recharge
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// The transaction type.
	//
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

func (s *ModelRouterCreateBalanceTransactionRequest) GetBalanceType() *string {
	return s.BalanceType
}

func (s *ModelRouterCreateBalanceTransactionRequest) GetIdempotencyKey() *string {
	return s.IdempotencyKey
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

func (s *ModelRouterCreateBalanceTransactionRequest) SetBalanceType(v string) *ModelRouterCreateBalanceTransactionRequest {
	s.BalanceType = &v
	return s
}

func (s *ModelRouterCreateBalanceTransactionRequest) SetIdempotencyKey(v string) *ModelRouterCreateBalanceTransactionRequest {
	s.IdempotencyKey = &v
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
