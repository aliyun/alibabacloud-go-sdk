// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSlaCouponApplyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMetadata(v interface{}) *SubmitSlaCouponApplyResponseBody
	GetMetadata() interface{}
	SetRequestId(v string) *SubmitSlaCouponApplyResponseBody
	GetRequestId() *string
	SetSumCoupon(v float64) *SubmitSlaCouponApplyResponseBody
	GetSumCoupon() *float64
	SetValidEndTime(v string) *SubmitSlaCouponApplyResponseBody
	GetValidEndTime() *string
	SetValidStartTime(v string) *SubmitSlaCouponApplyResponseBody
	GetValidStartTime() *string
}

type SubmitSlaCouponApplyResponseBody struct {
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 79EE7556-0CFD-44EB-9CD6-B3B526E3A85F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 6.4
	SumCoupon *float64 `json:"SumCoupon,omitempty" xml:"SumCoupon,omitempty"`
	// example:
	//
	// Mon Apr 27 00:00:00 CST 2026
	ValidEndTime *string `json:"ValidEndTime,omitempty" xml:"ValidEndTime,omitempty"`
	// example:
	//
	// Tue Oct 27 13:15:58 CST 2026
	ValidStartTime *string `json:"ValidStartTime,omitempty" xml:"ValidStartTime,omitempty"`
}

func (s SubmitSlaCouponApplyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitSlaCouponApplyResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSlaCouponApplyResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *SubmitSlaCouponApplyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitSlaCouponApplyResponseBody) GetSumCoupon() *float64 {
	return s.SumCoupon
}

func (s *SubmitSlaCouponApplyResponseBody) GetValidEndTime() *string {
	return s.ValidEndTime
}

func (s *SubmitSlaCouponApplyResponseBody) GetValidStartTime() *string {
	return s.ValidStartTime
}

func (s *SubmitSlaCouponApplyResponseBody) SetMetadata(v interface{}) *SubmitSlaCouponApplyResponseBody {
	s.Metadata = v
	return s
}

func (s *SubmitSlaCouponApplyResponseBody) SetRequestId(v string) *SubmitSlaCouponApplyResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSlaCouponApplyResponseBody) SetSumCoupon(v float64) *SubmitSlaCouponApplyResponseBody {
	s.SumCoupon = &v
	return s
}

func (s *SubmitSlaCouponApplyResponseBody) SetValidEndTime(v string) *SubmitSlaCouponApplyResponseBody {
	s.ValidEndTime = &v
	return s
}

func (s *SubmitSlaCouponApplyResponseBody) SetValidStartTime(v string) *SubmitSlaCouponApplyResponseBody {
	s.ValidStartTime = &v
	return s
}

func (s *SubmitSlaCouponApplyResponseBody) Validate() error {
	return dara.Validate(s)
}
