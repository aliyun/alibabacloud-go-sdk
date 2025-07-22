// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCouponDeductTagShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCouponId(v string) *DeleteCouponDeductTagShrinkRequest
	GetCouponId() *string
	SetEcIdAccountIdsShrink(v string) *DeleteCouponDeductTagShrinkRequest
	GetEcIdAccountIdsShrink() *string
	SetNbid(v string) *DeleteCouponDeductTagShrinkRequest
	GetNbid() *string
	SetTagKeysShrink(v string) *DeleteCouponDeductTagShrinkRequest
	GetTagKeysShrink() *string
}

type DeleteCouponDeductTagShrinkRequest struct {
	CouponId             *string `json:"CouponId,omitempty" xml:"CouponId,omitempty"`
	EcIdAccountIdsShrink *string `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty"`
	Nbid                 *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	TagKeysShrink        *string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty"`
}

func (s DeleteCouponDeductTagShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCouponDeductTagShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteCouponDeductTagShrinkRequest) GetCouponId() *string {
	return s.CouponId
}

func (s *DeleteCouponDeductTagShrinkRequest) GetEcIdAccountIdsShrink() *string {
	return s.EcIdAccountIdsShrink
}

func (s *DeleteCouponDeductTagShrinkRequest) GetNbid() *string {
	return s.Nbid
}

func (s *DeleteCouponDeductTagShrinkRequest) GetTagKeysShrink() *string {
	return s.TagKeysShrink
}

func (s *DeleteCouponDeductTagShrinkRequest) SetCouponId(v string) *DeleteCouponDeductTagShrinkRequest {
	s.CouponId = &v
	return s
}

func (s *DeleteCouponDeductTagShrinkRequest) SetEcIdAccountIdsShrink(v string) *DeleteCouponDeductTagShrinkRequest {
	s.EcIdAccountIdsShrink = &v
	return s
}

func (s *DeleteCouponDeductTagShrinkRequest) SetNbid(v string) *DeleteCouponDeductTagShrinkRequest {
	s.Nbid = &v
	return s
}

func (s *DeleteCouponDeductTagShrinkRequest) SetTagKeysShrink(v string) *DeleteCouponDeductTagShrinkRequest {
	s.TagKeysShrink = &v
	return s
}

func (s *DeleteCouponDeductTagShrinkRequest) Validate() error {
	return dara.Validate(s)
}
