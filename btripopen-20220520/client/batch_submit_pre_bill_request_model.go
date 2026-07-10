// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSubmitPreBillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppIp(v string) *BatchSubmitPreBillRequest
	GetAppIp() *string
	SetBillBatch(v string) *BatchSubmitPreBillRequest
	GetBillBatch() *string
	SetCustomerDecision(v int32) *BatchSubmitPreBillRequest
	GetCustomerDecision() *int32
	SetDimension(v int32) *BatchSubmitPreBillRequest
	GetDimension() *int32
	SetValues(v []*string) *BatchSubmitPreBillRequest
	GetValues() []*string
}

type BatchSubmitPreBillRequest struct {
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
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s BatchSubmitPreBillRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchSubmitPreBillRequest) GoString() string {
	return s.String()
}

func (s *BatchSubmitPreBillRequest) GetAppIp() *string {
	return s.AppIp
}

func (s *BatchSubmitPreBillRequest) GetBillBatch() *string {
	return s.BillBatch
}

func (s *BatchSubmitPreBillRequest) GetCustomerDecision() *int32 {
	return s.CustomerDecision
}

func (s *BatchSubmitPreBillRequest) GetDimension() *int32 {
	return s.Dimension
}

func (s *BatchSubmitPreBillRequest) GetValues() []*string {
	return s.Values
}

func (s *BatchSubmitPreBillRequest) SetAppIp(v string) *BatchSubmitPreBillRequest {
	s.AppIp = &v
	return s
}

func (s *BatchSubmitPreBillRequest) SetBillBatch(v string) *BatchSubmitPreBillRequest {
	s.BillBatch = &v
	return s
}

func (s *BatchSubmitPreBillRequest) SetCustomerDecision(v int32) *BatchSubmitPreBillRequest {
	s.CustomerDecision = &v
	return s
}

func (s *BatchSubmitPreBillRequest) SetDimension(v int32) *BatchSubmitPreBillRequest {
	s.Dimension = &v
	return s
}

func (s *BatchSubmitPreBillRequest) SetValues(v []*string) *BatchSubmitPreBillRequest {
	s.Values = v
	return s
}

func (s *BatchSubmitPreBillRequest) Validate() error {
	return dara.Validate(s)
}
