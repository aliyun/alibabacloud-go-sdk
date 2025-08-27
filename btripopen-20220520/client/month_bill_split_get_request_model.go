// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonthBillSplitGetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillBatch(v string) *MonthBillSplitGetRequest
	GetBillBatch() *string
	SetBillMonth(v string) *MonthBillSplitGetRequest
	GetBillMonth() *string
	SetBillSplitKeyList(v []*string) *MonthBillSplitGetRequest
	GetBillSplitKeyList() []*string
	SetBillSplitMode(v string) *MonthBillSplitGetRequest
	GetBillSplitMode() *string
}

type MonthBillSplitGetRequest struct {
	// example:
	//
	// 20240101
	BillBatch *string `json:"bill_batch,omitempty" xml:"bill_batch,omitempty"`
	// example:
	//
	// 202401
	BillMonth        *string   `json:"bill_month,omitempty" xml:"bill_month,omitempty"`
	BillSplitKeyList []*string `json:"bill_split_key_list,omitempty" xml:"bill_split_key_list,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// by_invoice_third_part_id
	BillSplitMode *string `json:"bill_split_mode,omitempty" xml:"bill_split_mode,omitempty"`
}

func (s MonthBillSplitGetRequest) String() string {
	return dara.Prettify(s)
}

func (s MonthBillSplitGetRequest) GoString() string {
	return s.String()
}

func (s *MonthBillSplitGetRequest) GetBillBatch() *string {
	return s.BillBatch
}

func (s *MonthBillSplitGetRequest) GetBillMonth() *string {
	return s.BillMonth
}

func (s *MonthBillSplitGetRequest) GetBillSplitKeyList() []*string {
	return s.BillSplitKeyList
}

func (s *MonthBillSplitGetRequest) GetBillSplitMode() *string {
	return s.BillSplitMode
}

func (s *MonthBillSplitGetRequest) SetBillBatch(v string) *MonthBillSplitGetRequest {
	s.BillBatch = &v
	return s
}

func (s *MonthBillSplitGetRequest) SetBillMonth(v string) *MonthBillSplitGetRequest {
	s.BillMonth = &v
	return s
}

func (s *MonthBillSplitGetRequest) SetBillSplitKeyList(v []*string) *MonthBillSplitGetRequest {
	s.BillSplitKeyList = v
	return s
}

func (s *MonthBillSplitGetRequest) SetBillSplitMode(v string) *MonthBillSplitGetRequest {
	s.BillSplitMode = &v
	return s
}

func (s *MonthBillSplitGetRequest) Validate() error {
	return dara.Validate(s)
}
