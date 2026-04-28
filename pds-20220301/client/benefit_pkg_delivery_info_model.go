// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBenefitPkgDeliveryInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int64) *BenefitPkgDeliveryInfo
	GetAmount() *int64
	SetCreatedAt(v string) *BenefitPkgDeliveryInfo
	GetCreatedAt() *string
	SetExpireTime(v string) *BenefitPkgDeliveryInfo
	GetExpireTime() *string
	SetIsPermanent(v bool) *BenefitPkgDeliveryInfo
	GetIsPermanent() *bool
}

type BenefitPkgDeliveryInfo struct {
	// Number of benefit packages delivered.
	//
	// example:
	//
	// 1
	Amount *int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// The creation time of the benefit package delivery.
	//
	// example:
	//
	// 2019-08-20T06:51:27.292Z
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// The expiration time of the benefit package delivery.
	//
	// If is_permit is set to false, a valid value is returned.
	//
	// example:
	//
	// 2019-08-20T06:51:27.292Z
	ExpireTime *string `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	// Whether it is permanently valid.
	//
	// example:
	//
	// false
	IsPermanent *bool `json:"is_permanent,omitempty" xml:"is_permanent,omitempty"`
}

func (s BenefitPkgDeliveryInfo) String() string {
	return dara.Prettify(s)
}

func (s BenefitPkgDeliveryInfo) GoString() string {
	return s.String()
}

func (s *BenefitPkgDeliveryInfo) GetAmount() *int64 {
	return s.Amount
}

func (s *BenefitPkgDeliveryInfo) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *BenefitPkgDeliveryInfo) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *BenefitPkgDeliveryInfo) GetIsPermanent() *bool {
	return s.IsPermanent
}

func (s *BenefitPkgDeliveryInfo) SetAmount(v int64) *BenefitPkgDeliveryInfo {
	s.Amount = &v
	return s
}

func (s *BenefitPkgDeliveryInfo) SetCreatedAt(v string) *BenefitPkgDeliveryInfo {
	s.CreatedAt = &v
	return s
}

func (s *BenefitPkgDeliveryInfo) SetExpireTime(v string) *BenefitPkgDeliveryInfo {
	s.ExpireTime = &v
	return s
}

func (s *BenefitPkgDeliveryInfo) SetIsPermanent(v bool) *BenefitPkgDeliveryInfo {
	s.IsPermanent = &v
	return s
}

func (s *BenefitPkgDeliveryInfo) Validate() error {
	return dara.Validate(s)
}
