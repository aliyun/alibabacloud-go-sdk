// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOrderRefundDetailQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OrderRefundDetailQueryResponseBody
	GetRequestId() *string
	SetCode(v string) *OrderRefundDetailQueryResponseBody
	GetCode() *string
	SetMessage(v string) *OrderRefundDetailQueryResponseBody
	GetMessage() *string
	SetModule(v *OrderRefundDetailQueryResponseBodyModule) *OrderRefundDetailQueryResponseBody
	GetModule() *OrderRefundDetailQueryResponseBodyModule
	SetSuccess(v bool) *OrderRefundDetailQueryResponseBody
	GetSuccess() *bool
}

type OrderRefundDetailQueryResponseBody struct {
	// example:
	//
	// 8CA36096-1FEE-5756-86DD-D195FEDE080E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// true
	Message *string                                   `json:"message,omitempty" xml:"message,omitempty"`
	Module  *OrderRefundDetailQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s OrderRefundDetailQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OrderRefundDetailQueryResponseBody) GoString() string {
	return s.String()
}

func (s *OrderRefundDetailQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OrderRefundDetailQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *OrderRefundDetailQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OrderRefundDetailQueryResponseBody) GetModule() *OrderRefundDetailQueryResponseBodyModule {
	return s.Module
}

func (s *OrderRefundDetailQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OrderRefundDetailQueryResponseBody) SetRequestId(v string) *OrderRefundDetailQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *OrderRefundDetailQueryResponseBody) SetCode(v string) *OrderRefundDetailQueryResponseBody {
	s.Code = &v
	return s
}

func (s *OrderRefundDetailQueryResponseBody) SetMessage(v string) *OrderRefundDetailQueryResponseBody {
	s.Message = &v
	return s
}

func (s *OrderRefundDetailQueryResponseBody) SetModule(v *OrderRefundDetailQueryResponseBodyModule) *OrderRefundDetailQueryResponseBody {
	s.Module = v
	return s
}

func (s *OrderRefundDetailQueryResponseBody) SetSuccess(v bool) *OrderRefundDetailQueryResponseBody {
	s.Success = &v
	return s
}

func (s *OrderRefundDetailQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type OrderRefundDetailQueryResponseBodyModule struct {
	// example:
	//
	// 1017002195370467138
	OrderId       *string                                                  `json:"order_id,omitempty" xml:"order_id,omitempty"`
	RefundDetails []*OrderRefundDetailQueryResponseBodyModuleRefundDetails `json:"refund_details,omitempty" xml:"refund_details,omitempty" type:"Repeated"`
	// example:
	//
	// 10000
	TotalAmount *int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
}

func (s OrderRefundDetailQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s OrderRefundDetailQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *OrderRefundDetailQueryResponseBodyModule) GetOrderId() *string {
	return s.OrderId
}

func (s *OrderRefundDetailQueryResponseBodyModule) GetRefundDetails() []*OrderRefundDetailQueryResponseBodyModuleRefundDetails {
	return s.RefundDetails
}

func (s *OrderRefundDetailQueryResponseBodyModule) GetTotalAmount() *int64 {
	return s.TotalAmount
}

func (s *OrderRefundDetailQueryResponseBodyModule) SetOrderId(v string) *OrderRefundDetailQueryResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *OrderRefundDetailQueryResponseBodyModule) SetRefundDetails(v []*OrderRefundDetailQueryResponseBodyModuleRefundDetails) *OrderRefundDetailQueryResponseBodyModule {
	s.RefundDetails = v
	return s
}

func (s *OrderRefundDetailQueryResponseBodyModule) SetTotalAmount(v int64) *OrderRefundDetailQueryResponseBodyModule {
	s.TotalAmount = &v
	return s
}

func (s *OrderRefundDetailQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type OrderRefundDetailQueryResponseBodyModuleRefundDetails struct {
	// example:
	//
	// ALIPAY
	PersonPayChannel *string `json:"person_pay_channel,omitempty" xml:"person_pay_channel,omitempty"`
	// example:
	//
	// 2025010223001423691442474885
	PersonRefundId *string `json:"person_refund_id,omitempty" xml:"person_refund_id,omitempty"`
	// example:
	//
	// 5000
	RefundAmount *int64 `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
	// example:
	//
	// 4800
	RefundAmountCorp *int64 `json:"refund_amount_corp,omitempty" xml:"refund_amount_corp,omitempty"`
	// example:
	//
	// 200
	RefundAmountPerson *int64 `json:"refund_amount_person,omitempty" xml:"refund_amount_person,omitempty"`
	// example:
	//
	// ZLJD12241231000002
	SupplierRefundId *string `json:"supplier_refund_id,omitempty" xml:"supplier_refund_id,omitempty"`
}

func (s OrderRefundDetailQueryResponseBodyModuleRefundDetails) String() string {
	return dara.Prettify(s)
}

func (s OrderRefundDetailQueryResponseBodyModuleRefundDetails) GoString() string {
	return s.String()
}

func (s *OrderRefundDetailQueryResponseBodyModuleRefundDetails) GetPersonPayChannel() *string {
	return s.PersonPayChannel
}

func (s *OrderRefundDetailQueryResponseBodyModuleRefundDetails) GetPersonRefundId() *string {
	return s.PersonRefundId
}

func (s *OrderRefundDetailQueryResponseBodyModuleRefundDetails) GetRefundAmount() *int64 {
	return s.RefundAmount
}

func (s *OrderRefundDetailQueryResponseBodyModuleRefundDetails) GetRefundAmountCorp() *int64 {
	return s.RefundAmountCorp
}

func (s *OrderRefundDetailQueryResponseBodyModuleRefundDetails) GetRefundAmountPerson() *int64 {
	return s.RefundAmountPerson
}

func (s *OrderRefundDetailQueryResponseBodyModuleRefundDetails) GetSupplierRefundId() *string {
	return s.SupplierRefundId
}

func (s *OrderRefundDetailQueryResponseBodyModuleRefundDetails) SetPersonPayChannel(v string) *OrderRefundDetailQueryResponseBodyModuleRefundDetails {
	s.PersonPayChannel = &v
	return s
}

func (s *OrderRefundDetailQueryResponseBodyModuleRefundDetails) SetPersonRefundId(v string) *OrderRefundDetailQueryResponseBodyModuleRefundDetails {
	s.PersonRefundId = &v
	return s
}

func (s *OrderRefundDetailQueryResponseBodyModuleRefundDetails) SetRefundAmount(v int64) *OrderRefundDetailQueryResponseBodyModuleRefundDetails {
	s.RefundAmount = &v
	return s
}

func (s *OrderRefundDetailQueryResponseBodyModuleRefundDetails) SetRefundAmountCorp(v int64) *OrderRefundDetailQueryResponseBodyModuleRefundDetails {
	s.RefundAmountCorp = &v
	return s
}

func (s *OrderRefundDetailQueryResponseBodyModuleRefundDetails) SetRefundAmountPerson(v int64) *OrderRefundDetailQueryResponseBodyModuleRefundDetails {
	s.RefundAmountPerson = &v
	return s
}

func (s *OrderRefundDetailQueryResponseBodyModuleRefundDetails) SetSupplierRefundId(v string) *OrderRefundDetailQueryResponseBodyModuleRefundDetails {
	s.SupplierRefundId = &v
	return s
}

func (s *OrderRefundDetailQueryResponseBodyModuleRefundDetails) Validate() error {
	return dara.Validate(s)
}
