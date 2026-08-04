// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateMemberBalanceTransactionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v float64) *ModelRouterCreateMemberBalanceTransactionRequest
	GetAmount() *float64
	SetBalanceType(v string) *ModelRouterCreateMemberBalanceTransactionRequest
	GetBalanceType() *string
	SetIdempotencyKey(v string) *ModelRouterCreateMemberBalanceTransactionRequest
	GetIdempotencyKey() *string
	SetRemark(v string) *ModelRouterCreateMemberBalanceTransactionRequest
	GetRemark() *string
	SetType(v string) *ModelRouterCreateMemberBalanceTransactionRequest
	GetType() *string
}

type ModelRouterCreateMemberBalanceTransactionRequest struct {
	// The transaction amount.
	//
	// example:
	//
	// 100.00
	Amount *float64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// The balance type. Valid values:
	//
	// - permanent
	//
	// - monthly
	//
	// Default value: permanent.
	//
	// example:
	//
	// permanent
	BalanceType *string `json:"balanceType,omitempty" xml:"balanceType,omitempty"`
	// The idempotency key. UUID v4 format is recommended.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	IdempotencyKey *string `json:"idempotencyKey,omitempty" xml:"idempotencyKey,omitempty"`
	// The remark for the transaction.
	//
	// example:
	//
	// Recharge
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// The transaction type. Valid values: recharge, deduct, and transfer.
	//
	// example:
	//
	// recharge
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ModelRouterCreateMemberBalanceTransactionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateMemberBalanceTransactionRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateMemberBalanceTransactionRequest) GetAmount() *float64 {
	return s.Amount
}

func (s *ModelRouterCreateMemberBalanceTransactionRequest) GetBalanceType() *string {
	return s.BalanceType
}

func (s *ModelRouterCreateMemberBalanceTransactionRequest) GetIdempotencyKey() *string {
	return s.IdempotencyKey
}

func (s *ModelRouterCreateMemberBalanceTransactionRequest) GetRemark() *string {
	return s.Remark
}

func (s *ModelRouterCreateMemberBalanceTransactionRequest) GetType() *string {
	return s.Type
}

func (s *ModelRouterCreateMemberBalanceTransactionRequest) SetAmount(v float64) *ModelRouterCreateMemberBalanceTransactionRequest {
	s.Amount = &v
	return s
}

func (s *ModelRouterCreateMemberBalanceTransactionRequest) SetBalanceType(v string) *ModelRouterCreateMemberBalanceTransactionRequest {
	s.BalanceType = &v
	return s
}

func (s *ModelRouterCreateMemberBalanceTransactionRequest) SetIdempotencyKey(v string) *ModelRouterCreateMemberBalanceTransactionRequest {
	s.IdempotencyKey = &v
	return s
}

func (s *ModelRouterCreateMemberBalanceTransactionRequest) SetRemark(v string) *ModelRouterCreateMemberBalanceTransactionRequest {
	s.Remark = &v
	return s
}

func (s *ModelRouterCreateMemberBalanceTransactionRequest) SetType(v string) *ModelRouterCreateMemberBalanceTransactionRequest {
	s.Type = &v
	return s
}

func (s *ModelRouterCreateMemberBalanceTransactionRequest) Validate() error {
	return dara.Validate(s)
}
