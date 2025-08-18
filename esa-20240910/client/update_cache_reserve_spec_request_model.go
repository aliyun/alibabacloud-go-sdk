// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCacheReserveSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *UpdateCacheReserveSpecRequest
	GetAutoPay() *bool
	SetChargeType(v string) *UpdateCacheReserveSpecRequest
	GetChargeType() *string
	SetInstanceId(v string) *UpdateCacheReserveSpecRequest
	GetInstanceId() *string
	SetTargetQuotaGb(v int64) *UpdateCacheReserveSpecRequest
	GetTargetQuotaGb() *int64
}

type UpdateCacheReserveSpecRequest struct {
	// Specifies whether to enable auto payment.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// example:
	//
	// PREPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// esa-cr-9tuv*********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1000
	TargetQuotaGb *int64 `json:"TargetQuotaGb,omitempty" xml:"TargetQuotaGb,omitempty"`
}

func (s UpdateCacheReserveSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCacheReserveSpecRequest) GoString() string {
	return s.String()
}

func (s *UpdateCacheReserveSpecRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *UpdateCacheReserveSpecRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *UpdateCacheReserveSpecRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateCacheReserveSpecRequest) GetTargetQuotaGb() *int64 {
	return s.TargetQuotaGb
}

func (s *UpdateCacheReserveSpecRequest) SetAutoPay(v bool) *UpdateCacheReserveSpecRequest {
	s.AutoPay = &v
	return s
}

func (s *UpdateCacheReserveSpecRequest) SetChargeType(v string) *UpdateCacheReserveSpecRequest {
	s.ChargeType = &v
	return s
}

func (s *UpdateCacheReserveSpecRequest) SetInstanceId(v string) *UpdateCacheReserveSpecRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateCacheReserveSpecRequest) SetTargetQuotaGb(v int64) *UpdateCacheReserveSpecRequest {
	s.TargetQuotaGb = &v
	return s
}

func (s *UpdateCacheReserveSpecRequest) Validate() error {
	return dara.Validate(s)
}
