// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCouponDeductTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteCouponDeductTagResponseBody
	GetData() *bool
	SetRequestId(v string) *DeleteCouponDeductTagResponseBody
	GetRequestId() *string
}

type DeleteCouponDeductTagResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCouponDeductTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCouponDeductTagResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCouponDeductTagResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteCouponDeductTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCouponDeductTagResponseBody) SetData(v bool) *DeleteCouponDeductTagResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteCouponDeductTagResponseBody) SetRequestId(v string) *DeleteCouponDeductTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCouponDeductTagResponseBody) Validate() error {
	return dara.Validate(s)
}
