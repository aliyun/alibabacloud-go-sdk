// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCouponDeductTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *AddCouponDeductTagResponseBody
	GetData() *bool
	SetRequestId(v string) *AddCouponDeductTagResponseBody
	GetRequestId() *string
}

type AddCouponDeductTagResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddCouponDeductTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddCouponDeductTagResponseBody) GoString() string {
	return s.String()
}

func (s *AddCouponDeductTagResponseBody) GetData() *bool {
	return s.Data
}

func (s *AddCouponDeductTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddCouponDeductTagResponseBody) SetData(v bool) *AddCouponDeductTagResponseBody {
	s.Data = &v
	return s
}

func (s *AddCouponDeductTagResponseBody) SetRequestId(v string) *AddCouponDeductTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCouponDeductTagResponseBody) Validate() error {
	return dara.Validate(s)
}
