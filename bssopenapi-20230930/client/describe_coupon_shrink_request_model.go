// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCouponShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCouponId(v int64) *DescribeCouponShrinkRequest
	GetCouponId() *int64
	SetCouponNo(v string) *DescribeCouponShrinkRequest
	GetCouponNo() *string
	SetCouponType(v string) *DescribeCouponShrinkRequest
	GetCouponType() *string
	SetCurrentPage(v int32) *DescribeCouponShrinkRequest
	GetCurrentPage() *int32
	SetEcIdAccountIdsShrink(v string) *DescribeCouponShrinkRequest
	GetEcIdAccountIdsShrink() *string
	SetEffectiveEndTime(v int64) *DescribeCouponShrinkRequest
	GetEffectiveEndTime() *int64
	SetEffectiveStartTime(v int64) *DescribeCouponShrinkRequest
	GetEffectiveStartTime() *int64
	SetExpireEndDate(v int64) *DescribeCouponShrinkRequest
	GetExpireEndDate() *int64
	SetExpireStartDate(v int64) *DescribeCouponShrinkRequest
	GetExpireStartDate() *int64
	SetNbid(v string) *DescribeCouponShrinkRequest
	GetNbid() *string
	SetPageSize(v int32) *DescribeCouponShrinkRequest
	GetPageSize() *int32
	SetStatus(v string) *DescribeCouponShrinkRequest
	GetStatus() *string
}

type DescribeCouponShrinkRequest struct {
	// example:
	//
	// 351430260343
	CouponId *int64 `json:"CouponId,omitempty" xml:"CouponId,omitempty"`
	// example:
	//
	// 554863270150
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// example:
	//
	// CERTAIN
	CouponType *string `json:"CouponType,omitempty" xml:"CouponType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage          *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EcIdAccountIdsShrink *string `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty"`
	// example:
	//
	// 1708423156000
	EffectiveEndTime *int64 `json:"EffectiveEndTime,omitempty" xml:"EffectiveEndTime,omitempty"`
	// example:
	//
	// 1684750028000
	EffectiveStartTime *int64 `json:"EffectiveStartTime,omitempty" xml:"EffectiveStartTime,omitempty"`
	// example:
	//
	// 1708423156000
	ExpireEndDate *int64 `json:"ExpireEndDate,omitempty" xml:"ExpireEndDate,omitempty"`
	// example:
	//
	// 1684750028000
	ExpireStartDate *int64 `json:"ExpireStartDate,omitempty" xml:"ExpireStartDate,omitempty"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// AVAILABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCouponShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCouponShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeCouponShrinkRequest) GetCouponId() *int64 {
	return s.CouponId
}

func (s *DescribeCouponShrinkRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *DescribeCouponShrinkRequest) GetCouponType() *string {
	return s.CouponType
}

func (s *DescribeCouponShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCouponShrinkRequest) GetEcIdAccountIdsShrink() *string {
	return s.EcIdAccountIdsShrink
}

func (s *DescribeCouponShrinkRequest) GetEffectiveEndTime() *int64 {
	return s.EffectiveEndTime
}

func (s *DescribeCouponShrinkRequest) GetEffectiveStartTime() *int64 {
	return s.EffectiveStartTime
}

func (s *DescribeCouponShrinkRequest) GetExpireEndDate() *int64 {
	return s.ExpireEndDate
}

func (s *DescribeCouponShrinkRequest) GetExpireStartDate() *int64 {
	return s.ExpireStartDate
}

func (s *DescribeCouponShrinkRequest) GetNbid() *string {
	return s.Nbid
}

func (s *DescribeCouponShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCouponShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeCouponShrinkRequest) SetCouponId(v int64) *DescribeCouponShrinkRequest {
	s.CouponId = &v
	return s
}

func (s *DescribeCouponShrinkRequest) SetCouponNo(v string) *DescribeCouponShrinkRequest {
	s.CouponNo = &v
	return s
}

func (s *DescribeCouponShrinkRequest) SetCouponType(v string) *DescribeCouponShrinkRequest {
	s.CouponType = &v
	return s
}

func (s *DescribeCouponShrinkRequest) SetCurrentPage(v int32) *DescribeCouponShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCouponShrinkRequest) SetEcIdAccountIdsShrink(v string) *DescribeCouponShrinkRequest {
	s.EcIdAccountIdsShrink = &v
	return s
}

func (s *DescribeCouponShrinkRequest) SetEffectiveEndTime(v int64) *DescribeCouponShrinkRequest {
	s.EffectiveEndTime = &v
	return s
}

func (s *DescribeCouponShrinkRequest) SetEffectiveStartTime(v int64) *DescribeCouponShrinkRequest {
	s.EffectiveStartTime = &v
	return s
}

func (s *DescribeCouponShrinkRequest) SetExpireEndDate(v int64) *DescribeCouponShrinkRequest {
	s.ExpireEndDate = &v
	return s
}

func (s *DescribeCouponShrinkRequest) SetExpireStartDate(v int64) *DescribeCouponShrinkRequest {
	s.ExpireStartDate = &v
	return s
}

func (s *DescribeCouponShrinkRequest) SetNbid(v string) *DescribeCouponShrinkRequest {
	s.Nbid = &v
	return s
}

func (s *DescribeCouponShrinkRequest) SetPageSize(v int32) *DescribeCouponShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCouponShrinkRequest) SetStatus(v string) *DescribeCouponShrinkRequest {
	s.Status = &v
	return s
}

func (s *DescribeCouponShrinkRequest) Validate() error {
	return dara.Validate(s)
}
