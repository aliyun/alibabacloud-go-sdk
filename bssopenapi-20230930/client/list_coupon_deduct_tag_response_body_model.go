// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCouponDeductTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListCouponDeductTagResponseBodyData) *ListCouponDeductTagResponseBody
	GetData() []*ListCouponDeductTagResponseBodyData
	SetRequestId(v string) *ListCouponDeductTagResponseBody
	GetRequestId() *string
}

type ListCouponDeductTagResponseBody struct {
	Data      []*ListCouponDeductTagResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCouponDeductTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCouponDeductTagResponseBody) GoString() string {
	return s.String()
}

func (s *ListCouponDeductTagResponseBody) GetData() []*ListCouponDeductTagResponseBodyData {
	return s.Data
}

func (s *ListCouponDeductTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCouponDeductTagResponseBody) SetData(v []*ListCouponDeductTagResponseBodyData) *ListCouponDeductTagResponseBody {
	s.Data = v
	return s
}

func (s *ListCouponDeductTagResponseBody) SetRequestId(v string) *ListCouponDeductTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCouponDeductTagResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCouponDeductTagResponseBodyData struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListCouponDeductTagResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCouponDeductTagResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCouponDeductTagResponseBodyData) GetKey() *string {
	return s.Key
}

func (s *ListCouponDeductTagResponseBodyData) GetValue() *string {
	return s.Value
}

func (s *ListCouponDeductTagResponseBodyData) SetKey(v string) *ListCouponDeductTagResponseBodyData {
	s.Key = &v
	return s
}

func (s *ListCouponDeductTagResponseBodyData) SetValue(v string) *ListCouponDeductTagResponseBodyData {
	s.Value = &v
	return s
}

func (s *ListCouponDeductTagResponseBodyData) Validate() error {
	return dara.Validate(s)
}
