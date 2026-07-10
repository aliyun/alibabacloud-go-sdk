// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSubmitPreBillShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppIp(v string) *BatchSubmitPreBillShrinkRequest
	GetAppIp() *string
	SetBillBatch(v string) *BatchSubmitPreBillShrinkRequest
	GetBillBatch() *string
	SetCustomerDecision(v int32) *BatchSubmitPreBillShrinkRequest
	GetCustomerDecision() *int32
	SetDimension(v int32) *BatchSubmitPreBillShrinkRequest
	GetDimension() *int32
	SetValuesShrink(v string) *BatchSubmitPreBillShrinkRequest
	GetValuesShrink() *string
}

type BatchSubmitPreBillShrinkRequest struct {
	// A system parameter. You do not need to manually specify this parameter.
	//
	// example:
	//
	// 100.66.54.114
	AppIp *string `json:"app_ip,omitempty" xml:"app_ip,omitempty"`
	// The bill batch date in the format of yyyy-MM-dd, such as 2026-06-21.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-06-21
	BillBatch *string `json:"bill_batch,omitempty" xml:"bill_batch,omitempty"`
	// The customer decision. Valid values:
	//
	// - 1: bill in the current period.
	//
	// - 2: deferred billing.
	//
	// - null: bill based on the current billing decision of the record.
	//
	// example:
	//
	// 1
	CustomerDecision *int32 `json:"customer_decision,omitempty" xml:"customer_decision,omitempty"`
	// The dimension type. Valid values:
	//
	// - 1: bill ID.
	//
	// - 2: order number.
	//
	// - 3: approval form.
	//
	// - 4: invoice title.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Dimension *int32 `json:"dimension,omitempty" xml:"dimension,omitempty"`
	// The values determined by the dimension parameter. For example, if dimension is set to 1, the values should be bill IDs.
	//
	// This parameter is required.
	ValuesShrink *string `json:"values,omitempty" xml:"values,omitempty"`
}

func (s BatchSubmitPreBillShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchSubmitPreBillShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchSubmitPreBillShrinkRequest) GetAppIp() *string {
	return s.AppIp
}

func (s *BatchSubmitPreBillShrinkRequest) GetBillBatch() *string {
	return s.BillBatch
}

func (s *BatchSubmitPreBillShrinkRequest) GetCustomerDecision() *int32 {
	return s.CustomerDecision
}

func (s *BatchSubmitPreBillShrinkRequest) GetDimension() *int32 {
	return s.Dimension
}

func (s *BatchSubmitPreBillShrinkRequest) GetValuesShrink() *string {
	return s.ValuesShrink
}

func (s *BatchSubmitPreBillShrinkRequest) SetAppIp(v string) *BatchSubmitPreBillShrinkRequest {
	s.AppIp = &v
	return s
}

func (s *BatchSubmitPreBillShrinkRequest) SetBillBatch(v string) *BatchSubmitPreBillShrinkRequest {
	s.BillBatch = &v
	return s
}

func (s *BatchSubmitPreBillShrinkRequest) SetCustomerDecision(v int32) *BatchSubmitPreBillShrinkRequest {
	s.CustomerDecision = &v
	return s
}

func (s *BatchSubmitPreBillShrinkRequest) SetDimension(v int32) *BatchSubmitPreBillShrinkRequest {
	s.Dimension = &v
	return s
}

func (s *BatchSubmitPreBillShrinkRequest) SetValuesShrink(v string) *BatchSubmitPreBillShrinkRequest {
	s.ValuesShrink = &v
	return s
}

func (s *BatchSubmitPreBillShrinkRequest) Validate() error {
	return dara.Validate(s)
}
