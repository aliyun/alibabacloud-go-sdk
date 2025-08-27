// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonthBillGetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillBatch(v string) *MonthBillGetRequest
	GetBillBatch() *string
	SetBillMonth(v string) *MonthBillGetRequest
	GetBillMonth() *string
}

type MonthBillGetRequest struct {
	BillBatch *string `json:"bill_batch,omitempty" xml:"bill_batch,omitempty"`
	// example:
	//
	// 202207
	BillMonth *string `json:"bill_month,omitempty" xml:"bill_month,omitempty"`
}

func (s MonthBillGetRequest) String() string {
	return dara.Prettify(s)
}

func (s MonthBillGetRequest) GoString() string {
	return s.String()
}

func (s *MonthBillGetRequest) GetBillBatch() *string {
	return s.BillBatch
}

func (s *MonthBillGetRequest) GetBillMonth() *string {
	return s.BillMonth
}

func (s *MonthBillGetRequest) SetBillBatch(v string) *MonthBillGetRequest {
	s.BillBatch = &v
	return s
}

func (s *MonthBillGetRequest) SetBillMonth(v string) *MonthBillGetRequest {
	s.BillMonth = &v
	return s
}

func (s *MonthBillGetRequest) Validate() error {
	return dara.Validate(s)
}
