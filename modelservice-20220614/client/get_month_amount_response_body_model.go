// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMonthAmountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackgroundAmount(v int32) *GetMonthAmountResponseBody
	GetBackgroundAmount() *int32
	SetPostpayAmount(v int32) *GetMonthAmountResponseBody
	GetPostpayAmount() *int32
	SetPrepayAmount(v int32) *GetMonthAmountResponseBody
	GetPrepayAmount() *int32
	SetRequestId(v string) *GetMonthAmountResponseBody
	GetRequestId() *string
	SetTotalAmount(v int32) *GetMonthAmountResponseBody
	GetTotalAmount() *int32
}

type GetMonthAmountResponseBody struct {
	BackgroundAmount *int32  `json:"BackgroundAmount,omitempty" xml:"BackgroundAmount,omitempty"`
	PostpayAmount    *int32  `json:"PostpayAmount,omitempty" xml:"PostpayAmount,omitempty"`
	PrepayAmount     *int32  `json:"PrepayAmount,omitempty" xml:"PrepayAmount,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalAmount      *int32  `json:"TotalAmount,omitempty" xml:"TotalAmount,omitempty"`
}

func (s GetMonthAmountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMonthAmountResponseBody) GoString() string {
	return s.String()
}

func (s *GetMonthAmountResponseBody) GetBackgroundAmount() *int32 {
	return s.BackgroundAmount
}

func (s *GetMonthAmountResponseBody) GetPostpayAmount() *int32 {
	return s.PostpayAmount
}

func (s *GetMonthAmountResponseBody) GetPrepayAmount() *int32 {
	return s.PrepayAmount
}

func (s *GetMonthAmountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMonthAmountResponseBody) GetTotalAmount() *int32 {
	return s.TotalAmount
}

func (s *GetMonthAmountResponseBody) SetBackgroundAmount(v int32) *GetMonthAmountResponseBody {
	s.BackgroundAmount = &v
	return s
}

func (s *GetMonthAmountResponseBody) SetPostpayAmount(v int32) *GetMonthAmountResponseBody {
	s.PostpayAmount = &v
	return s
}

func (s *GetMonthAmountResponseBody) SetPrepayAmount(v int32) *GetMonthAmountResponseBody {
	s.PrepayAmount = &v
	return s
}

func (s *GetMonthAmountResponseBody) SetRequestId(v string) *GetMonthAmountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMonthAmountResponseBody) SetTotalAmount(v int32) *GetMonthAmountResponseBody {
	s.TotalAmount = &v
	return s
}

func (s *GetMonthAmountResponseBody) Validate() error {
	return dara.Validate(s)
}
