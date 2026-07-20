// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMonthBillSplitGetShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillBatch(v string) *MonthBillSplitGetShrinkRequest
	GetBillBatch() *string
	SetBillMonth(v string) *MonthBillSplitGetShrinkRequest
	GetBillMonth() *string
	SetBillSplitKeyListShrink(v string) *MonthBillSplitGetShrinkRequest
	GetBillSplitKeyListShrink() *string
	SetBillSplitMode(v string) *MonthBillSplitGetShrinkRequest
	GetBillSplitMode() *string
}

type MonthBillSplitGetShrinkRequest struct {
	// example:
	//
	// 20240101
	BillBatch *string `json:"bill_batch,omitempty" xml:"bill_batch,omitempty"`
	// example:
	//
	// 202401
	BillMonth              *string `json:"bill_month,omitempty" xml:"bill_month,omitempty"`
	BillSplitKeyListShrink *string `json:"bill_split_key_list,omitempty" xml:"bill_split_key_list,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// by_invoice_third_part_id
	BillSplitMode *string `json:"bill_split_mode,omitempty" xml:"bill_split_mode,omitempty"`
}

func (s MonthBillSplitGetShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s MonthBillSplitGetShrinkRequest) GoString() string {
	return s.String()
}

func (s *MonthBillSplitGetShrinkRequest) GetBillBatch() *string {
	return s.BillBatch
}

func (s *MonthBillSplitGetShrinkRequest) GetBillMonth() *string {
	return s.BillMonth
}

func (s *MonthBillSplitGetShrinkRequest) GetBillSplitKeyListShrink() *string {
	return s.BillSplitKeyListShrink
}

func (s *MonthBillSplitGetShrinkRequest) GetBillSplitMode() *string {
	return s.BillSplitMode
}

func (s *MonthBillSplitGetShrinkRequest) SetBillBatch(v string) *MonthBillSplitGetShrinkRequest {
	s.BillBatch = &v
	return s
}

func (s *MonthBillSplitGetShrinkRequest) SetBillMonth(v string) *MonthBillSplitGetShrinkRequest {
	s.BillMonth = &v
	return s
}

func (s *MonthBillSplitGetShrinkRequest) SetBillSplitKeyListShrink(v string) *MonthBillSplitGetShrinkRequest {
	s.BillSplitKeyListShrink = &v
	return s
}

func (s *MonthBillSplitGetShrinkRequest) SetBillSplitMode(v string) *MonthBillSplitGetShrinkRequest {
	s.BillSplitMode = &v
	return s
}

func (s *MonthBillSplitGetShrinkRequest) Validate() error {
	return dara.Validate(s)
}
