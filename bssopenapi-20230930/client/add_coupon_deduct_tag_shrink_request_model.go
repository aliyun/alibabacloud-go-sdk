// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCouponDeductTagShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCouponId(v string) *AddCouponDeductTagShrinkRequest
	GetCouponId() *string
	SetEcIdAccountIdsShrink(v string) *AddCouponDeductTagShrinkRequest
	GetEcIdAccountIdsShrink() *string
	SetNbid(v string) *AddCouponDeductTagShrinkRequest
	GetNbid() *string
	SetTagsShrink(v string) *AddCouponDeductTagShrinkRequest
	GetTagsShrink() *string
}

type AddCouponDeductTagShrinkRequest struct {
	CouponId             *string `json:"CouponId,omitempty" xml:"CouponId,omitempty"`
	EcIdAccountIdsShrink *string `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty"`
	Nbid                 *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	TagsShrink           *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s AddCouponDeductTagShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCouponDeductTagShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddCouponDeductTagShrinkRequest) GetCouponId() *string {
	return s.CouponId
}

func (s *AddCouponDeductTagShrinkRequest) GetEcIdAccountIdsShrink() *string {
	return s.EcIdAccountIdsShrink
}

func (s *AddCouponDeductTagShrinkRequest) GetNbid() *string {
	return s.Nbid
}

func (s *AddCouponDeductTagShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *AddCouponDeductTagShrinkRequest) SetCouponId(v string) *AddCouponDeductTagShrinkRequest {
	s.CouponId = &v
	return s
}

func (s *AddCouponDeductTagShrinkRequest) SetEcIdAccountIdsShrink(v string) *AddCouponDeductTagShrinkRequest {
	s.EcIdAccountIdsShrink = &v
	return s
}

func (s *AddCouponDeductTagShrinkRequest) SetNbid(v string) *AddCouponDeductTagShrinkRequest {
	s.Nbid = &v
	return s
}

func (s *AddCouponDeductTagShrinkRequest) SetTagsShrink(v string) *AddCouponDeductTagShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *AddCouponDeductTagShrinkRequest) Validate() error {
	return dara.Validate(s)
}
