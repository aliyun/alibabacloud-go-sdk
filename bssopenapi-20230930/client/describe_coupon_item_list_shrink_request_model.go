// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCouponItemListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCouponId(v int64) *DescribeCouponItemListShrinkRequest
	GetCouponId() *int64
	SetCurrentPage(v int32) *DescribeCouponItemListShrinkRequest
	GetCurrentPage() *int32
	SetEcIdAccountIdsShrink(v string) *DescribeCouponItemListShrinkRequest
	GetEcIdAccountIdsShrink() *string
	SetName(v string) *DescribeCouponItemListShrinkRequest
	GetName() *string
	SetNbid(v string) *DescribeCouponItemListShrinkRequest
	GetNbid() *string
	SetPageSize(v int32) *DescribeCouponItemListShrinkRequest
	GetPageSize() *int32
}

type DescribeCouponItemListShrinkRequest struct {
	// example:
	//
	// 59104570
	CouponId *int64 `json:"CouponId,omitempty" xml:"CouponId,omitempty"`
	// example:
	//
	// 1
	CurrentPage          *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EcIdAccountIdsShrink *string `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeCouponItemListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCouponItemListShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeCouponItemListShrinkRequest) GetCouponId() *int64 {
	return s.CouponId
}

func (s *DescribeCouponItemListShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCouponItemListShrinkRequest) GetEcIdAccountIdsShrink() *string {
	return s.EcIdAccountIdsShrink
}

func (s *DescribeCouponItemListShrinkRequest) GetName() *string {
	return s.Name
}

func (s *DescribeCouponItemListShrinkRequest) GetNbid() *string {
	return s.Nbid
}

func (s *DescribeCouponItemListShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCouponItemListShrinkRequest) SetCouponId(v int64) *DescribeCouponItemListShrinkRequest {
	s.CouponId = &v
	return s
}

func (s *DescribeCouponItemListShrinkRequest) SetCurrentPage(v int32) *DescribeCouponItemListShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCouponItemListShrinkRequest) SetEcIdAccountIdsShrink(v string) *DescribeCouponItemListShrinkRequest {
	s.EcIdAccountIdsShrink = &v
	return s
}

func (s *DescribeCouponItemListShrinkRequest) SetName(v string) *DescribeCouponItemListShrinkRequest {
	s.Name = &v
	return s
}

func (s *DescribeCouponItemListShrinkRequest) SetNbid(v string) *DescribeCouponItemListShrinkRequest {
	s.Nbid = &v
	return s
}

func (s *DescribeCouponItemListShrinkRequest) SetPageSize(v int32) *DescribeCouponItemListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCouponItemListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
