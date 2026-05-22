// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurchaseCacheReserveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *PurchaseCacheReserveRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *PurchaseCacheReserveRequest
	GetAutoRenew() *bool
	SetChargeType(v string) *PurchaseCacheReserveRequest
	GetChargeType() *string
	SetCrRegion(v string) *PurchaseCacheReserveRequest
	GetCrRegion() *string
	SetPeriod(v int32) *PurchaseCacheReserveRequest
	GetPeriod() *int32
	SetQuotaGb(v int64) *PurchaseCacheReserveRequest
	GetQuotaGb() *int64
}

type PurchaseCacheReserveRequest struct {
	// Whether to automatically pay. The default value is false.
	//
	// - true: Automatically pay.
	//
	// - false: Do not automatically pay.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Whether to auto-renew:
	//
	// - true: Auto-renew.
	//
	// - false: Do not auto-renew.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// Billing type
	//
	// - PREPAY: Prepaid.
	//
	// - POSTPAY: Postpaid.
	//
	// example:
	//
	// PREPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// Cache retention region
	//
	// - HK: Hong Kong, China
	//
	// - CN-beijing: Mainland China - Beijing
	//
	// example:
	//
	// HK
	CrRegion *string `json:"CrRegion,omitempty" xml:"CrRegion,omitempty"`
	// Purchase period (unit: month).
	//
	// example:
	//
	// 3
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// Cache retention specification (unit: GB).
	//
	// example:
	//
	// 512000
	QuotaGb *int64 `json:"QuotaGb,omitempty" xml:"QuotaGb,omitempty"`
}

func (s PurchaseCacheReserveRequest) String() string {
	return dara.Prettify(s)
}

func (s PurchaseCacheReserveRequest) GoString() string {
	return s.String()
}

func (s *PurchaseCacheReserveRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *PurchaseCacheReserveRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *PurchaseCacheReserveRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *PurchaseCacheReserveRequest) GetCrRegion() *string {
	return s.CrRegion
}

func (s *PurchaseCacheReserveRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *PurchaseCacheReserveRequest) GetQuotaGb() *int64 {
	return s.QuotaGb
}

func (s *PurchaseCacheReserveRequest) SetAutoPay(v bool) *PurchaseCacheReserveRequest {
	s.AutoPay = &v
	return s
}

func (s *PurchaseCacheReserveRequest) SetAutoRenew(v bool) *PurchaseCacheReserveRequest {
	s.AutoRenew = &v
	return s
}

func (s *PurchaseCacheReserveRequest) SetChargeType(v string) *PurchaseCacheReserveRequest {
	s.ChargeType = &v
	return s
}

func (s *PurchaseCacheReserveRequest) SetCrRegion(v string) *PurchaseCacheReserveRequest {
	s.CrRegion = &v
	return s
}

func (s *PurchaseCacheReserveRequest) SetPeriod(v int32) *PurchaseCacheReserveRequest {
	s.Period = &v
	return s
}

func (s *PurchaseCacheReserveRequest) SetQuotaGb(v int64) *PurchaseCacheReserveRequest {
	s.QuotaGb = &v
	return s
}

func (s *PurchaseCacheReserveRequest) Validate() error {
	return dara.Validate(s)
}
