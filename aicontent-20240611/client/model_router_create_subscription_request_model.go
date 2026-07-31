// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateSubscriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBalanceType(v string) *ModelRouterCreateSubscriptionRequest
	GetBalanceType() *string
	SetEffectiveTime(v int64) *ModelRouterCreateSubscriptionRequest
	GetEffectiveTime() *int64
	SetIdempotencyKey(v string) *ModelRouterCreateSubscriptionRequest
	GetIdempotencyKey() *string
	SetSubscriptionAmount(v float64) *ModelRouterCreateSubscriptionRequest
	GetSubscriptionAmount() *float64
}

type ModelRouterCreateSubscriptionRequest struct {
	// The balance pool to which the recharge is applied. Valid values:
	//
	// - permanent: the permanent balance pool.
	//
	// - monthly: the monthly balance pool.
	//
	// example:
	//
	// permanent
	BalanceType *string `json:"balanceType,omitempty" xml:"balanceType,omitempty"`
	// The effective period, in UNIX timestamp (seconds). Range: from 00:00 of today to 00:00 of the first day of the next month (Asia/Shanghai).
	//
	// example:
	//
	// 1719792000
	EffectiveTime *int64 `json:"effectiveTime,omitempty" xml:"effectiveTime,omitempty"`
	// The idempotency key. UUID v4 format without hyphens is recommended. This prevents duplicate subscription creation.
	//
	// example:
	//
	// 550e8400e29b41d4a716446655440000
	IdempotencyKey *string `json:"idempotencyKey,omitempty" xml:"idempotencyKey,omitempty"`
	// The subscription recharge amount.
	//
	// example:
	//
	// 100.00
	SubscriptionAmount *float64 `json:"subscriptionAmount,omitempty" xml:"subscriptionAmount,omitempty"`
}

func (s ModelRouterCreateSubscriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateSubscriptionRequest) GetBalanceType() *string {
	return s.BalanceType
}

func (s *ModelRouterCreateSubscriptionRequest) GetEffectiveTime() *int64 {
	return s.EffectiveTime
}

func (s *ModelRouterCreateSubscriptionRequest) GetIdempotencyKey() *string {
	return s.IdempotencyKey
}

func (s *ModelRouterCreateSubscriptionRequest) GetSubscriptionAmount() *float64 {
	return s.SubscriptionAmount
}

func (s *ModelRouterCreateSubscriptionRequest) SetBalanceType(v string) *ModelRouterCreateSubscriptionRequest {
	s.BalanceType = &v
	return s
}

func (s *ModelRouterCreateSubscriptionRequest) SetEffectiveTime(v int64) *ModelRouterCreateSubscriptionRequest {
	s.EffectiveTime = &v
	return s
}

func (s *ModelRouterCreateSubscriptionRequest) SetIdempotencyKey(v string) *ModelRouterCreateSubscriptionRequest {
	s.IdempotencyKey = &v
	return s
}

func (s *ModelRouterCreateSubscriptionRequest) SetSubscriptionAmount(v float64) *ModelRouterCreateSubscriptionRequest {
	s.SubscriptionAmount = &v
	return s
}

func (s *ModelRouterCreateSubscriptionRequest) Validate() error {
	return dara.Validate(s)
}
