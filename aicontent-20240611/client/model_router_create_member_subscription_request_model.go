// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateMemberSubscriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v float64) *ModelRouterCreateMemberSubscriptionRequest
	GetAmount() *float64
	SetBalanceType(v string) *ModelRouterCreateMemberSubscriptionRequest
	GetBalanceType() *string
	SetEffectiveTime(v int64) *ModelRouterCreateMemberSubscriptionRequest
	GetEffectiveTime() *int64
	SetIdempotencyKey(v string) *ModelRouterCreateMemberSubscriptionRequest
	GetIdempotencyKey() *string
}

type ModelRouterCreateMemberSubscriptionRequest struct {
	// The subscription amount.
	//
	// example:
	//
	// 100.00
	Amount *float64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// The balance type. Valid values:
	//
	// - permanent: permanent balance.
	//
	// - monthly: monthly balance.
	//
	// example:
	//
	// monthly
	BalanceType *string `json:"balanceType,omitempty" xml:"balanceType,omitempty"`
	// The effective period in UNIX timestamp (seconds).
	//
	// example:
	//
	// 1753858800
	EffectiveTime *int64 `json:"effectiveTime,omitempty" xml:"effectiveTime,omitempty"`
	// The idempotency key. UUID v4 format is recommended.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	IdempotencyKey *string `json:"idempotencyKey,omitempty" xml:"idempotencyKey,omitempty"`
}

func (s ModelRouterCreateMemberSubscriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateMemberSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateMemberSubscriptionRequest) GetAmount() *float64 {
	return s.Amount
}

func (s *ModelRouterCreateMemberSubscriptionRequest) GetBalanceType() *string {
	return s.BalanceType
}

func (s *ModelRouterCreateMemberSubscriptionRequest) GetEffectiveTime() *int64 {
	return s.EffectiveTime
}

func (s *ModelRouterCreateMemberSubscriptionRequest) GetIdempotencyKey() *string {
	return s.IdempotencyKey
}

func (s *ModelRouterCreateMemberSubscriptionRequest) SetAmount(v float64) *ModelRouterCreateMemberSubscriptionRequest {
	s.Amount = &v
	return s
}

func (s *ModelRouterCreateMemberSubscriptionRequest) SetBalanceType(v string) *ModelRouterCreateMemberSubscriptionRequest {
	s.BalanceType = &v
	return s
}

func (s *ModelRouterCreateMemberSubscriptionRequest) SetEffectiveTime(v int64) *ModelRouterCreateMemberSubscriptionRequest {
	s.EffectiveTime = &v
	return s
}

func (s *ModelRouterCreateMemberSubscriptionRequest) SetIdempotencyKey(v string) *ModelRouterCreateMemberSubscriptionRequest {
	s.IdempotencyKey = &v
	return s
}

func (s *ModelRouterCreateMemberSubscriptionRequest) Validate() error {
	return dara.Validate(s)
}
