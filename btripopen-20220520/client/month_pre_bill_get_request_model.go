// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonthPreBillGetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillBatch(v string) *MonthPreBillGetRequest
	GetBillBatch() *string
	SetBillMonth(v string) *MonthPreBillGetRequest
	GetBillMonth() *string
}

type MonthPreBillGetRequest struct {
	BillBatch *string `json:"bill_batch,omitempty" xml:"bill_batch,omitempty"`
	BillMonth *string `json:"bill_month,omitempty" xml:"bill_month,omitempty"`
}

func (s MonthPreBillGetRequest) String() string {
	return dara.Prettify(s)
}

func (s MonthPreBillGetRequest) GoString() string {
	return s.String()
}

func (s *MonthPreBillGetRequest) GetBillBatch() *string {
	return s.BillBatch
}

func (s *MonthPreBillGetRequest) GetBillMonth() *string {
	return s.BillMonth
}

func (s *MonthPreBillGetRequest) SetBillBatch(v string) *MonthPreBillGetRequest {
	s.BillBatch = &v
	return s
}

func (s *MonthPreBillGetRequest) SetBillMonth(v string) *MonthPreBillGetRequest {
	s.BillMonth = &v
	return s
}

func (s *MonthPreBillGetRequest) Validate() error {
	return dara.Validate(s)
}
