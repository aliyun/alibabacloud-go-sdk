// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterTransferToMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v float64) *ModelRouterTransferToMemberRequest
	GetAmount() *float64
	SetBalanceType(v string) *ModelRouterTransferToMemberRequest
	GetBalanceType() *string
	SetIdempotencyKey(v string) *ModelRouterTransferToMemberRequest
	GetIdempotencyKey() *string
	SetMonthlyQuota(v float64) *ModelRouterTransferToMemberRequest
	GetMonthlyQuota() *float64
	SetRemark(v string) *ModelRouterTransferToMemberRequest
	GetRemark() *string
}

type ModelRouterTransferToMemberRequest struct {
	// example:
	//
	// 100.00
	Amount *float64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// example:
	//
	// permanent
	BalanceType *string `json:"balanceType,omitempty" xml:"balanceType,omitempty"`
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	IdempotencyKey *string `json:"idempotencyKey,omitempty" xml:"idempotencyKey,omitempty"`
	// example:
	//
	// 0
	MonthlyQuota *float64 `json:"monthlyQuota,omitempty" xml:"monthlyQuota,omitempty"`
	// example:
	//
	// 转账
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s ModelRouterTransferToMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterTransferToMemberRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterTransferToMemberRequest) GetAmount() *float64 {
	return s.Amount
}

func (s *ModelRouterTransferToMemberRequest) GetBalanceType() *string {
	return s.BalanceType
}

func (s *ModelRouterTransferToMemberRequest) GetIdempotencyKey() *string {
	return s.IdempotencyKey
}

func (s *ModelRouterTransferToMemberRequest) GetMonthlyQuota() *float64 {
	return s.MonthlyQuota
}

func (s *ModelRouterTransferToMemberRequest) GetRemark() *string {
	return s.Remark
}

func (s *ModelRouterTransferToMemberRequest) SetAmount(v float64) *ModelRouterTransferToMemberRequest {
	s.Amount = &v
	return s
}

func (s *ModelRouterTransferToMemberRequest) SetBalanceType(v string) *ModelRouterTransferToMemberRequest {
	s.BalanceType = &v
	return s
}

func (s *ModelRouterTransferToMemberRequest) SetIdempotencyKey(v string) *ModelRouterTransferToMemberRequest {
	s.IdempotencyKey = &v
	return s
}

func (s *ModelRouterTransferToMemberRequest) SetMonthlyQuota(v float64) *ModelRouterTransferToMemberRequest {
	s.MonthlyQuota = &v
	return s
}

func (s *ModelRouterTransferToMemberRequest) SetRemark(v string) *ModelRouterTransferToMemberRequest {
	s.Remark = &v
	return s
}

func (s *ModelRouterTransferToMemberRequest) Validate() error {
	return dara.Validate(s)
}
