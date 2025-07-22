// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCouponDeductTagShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCouponId(v string) *ListCouponDeductTagShrinkRequest
	GetCouponId() *string
	SetEcIdAccountIdsShrink(v string) *ListCouponDeductTagShrinkRequest
	GetEcIdAccountIdsShrink() *string
	SetNbid(v string) *ListCouponDeductTagShrinkRequest
	GetNbid() *string
}

type ListCouponDeductTagShrinkRequest struct {
	CouponId             *string `json:"CouponId,omitempty" xml:"CouponId,omitempty"`
	EcIdAccountIdsShrink *string `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty"`
	Nbid                 *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
}

func (s ListCouponDeductTagShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCouponDeductTagShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListCouponDeductTagShrinkRequest) GetCouponId() *string {
	return s.CouponId
}

func (s *ListCouponDeductTagShrinkRequest) GetEcIdAccountIdsShrink() *string {
	return s.EcIdAccountIdsShrink
}

func (s *ListCouponDeductTagShrinkRequest) GetNbid() *string {
	return s.Nbid
}

func (s *ListCouponDeductTagShrinkRequest) SetCouponId(v string) *ListCouponDeductTagShrinkRequest {
	s.CouponId = &v
	return s
}

func (s *ListCouponDeductTagShrinkRequest) SetEcIdAccountIdsShrink(v string) *ListCouponDeductTagShrinkRequest {
	s.EcIdAccountIdsShrink = &v
	return s
}

func (s *ListCouponDeductTagShrinkRequest) SetNbid(v string) *ListCouponDeductTagShrinkRequest {
	s.Nbid = &v
	return s
}

func (s *ListCouponDeductTagShrinkRequest) Validate() error {
	return dara.Validate(s)
}
