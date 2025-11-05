// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCouponRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCouponId(v int64) *CancelCouponRequest
	GetCouponId() *int64
}

type CancelCouponRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 501001340370350
	CouponId *int64 `json:"CouponId,omitempty" xml:"CouponId,omitempty"`
}

func (s CancelCouponRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelCouponRequest) GoString() string {
	return s.String()
}

func (s *CancelCouponRequest) GetCouponId() *int64 {
	return s.CouponId
}

func (s *CancelCouponRequest) SetCouponId(v int64) *CancelCouponRequest {
	s.CouponId = &v
	return s
}

func (s *CancelCouponRequest) Validate() error {
	return dara.Validate(s)
}
