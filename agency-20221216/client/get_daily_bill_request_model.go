// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDailyBillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillOwner(v string) *GetDailyBillRequest
	GetBillOwner() *string
	SetBillType(v string) *GetDailyBillRequest
	GetBillType() *string
	SetDate(v string) *GetDailyBillRequest
	GetDate() *string
}

type GetDailyBillRequest struct {
	// Bill Owner type. Value Range:</br>
	//
	// 1: Master account</br>
	//
	// 2: Sub account</br>
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	BillOwner *string `json:"BillOwner,omitempty" xml:"BillOwner,omitempty"`
	// BillType. Value Range:</br>
	//
	// - DailyOrder(Deprecated)
	//
	// - DailyBill (Deprecated)
	//
	// - DailyInstanceBill (Deprecated)
	//
	// - DailyInstanceBillV2
	//
	// This parameter is required.
	//
	// example:
	//
	// DailyInstanceBillV2
	BillType *string `json:"BillType,omitempty" xml:"BillType,omitempty"`
	// Billing date. Format YYYY-MM-DD
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-11-24
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
}

func (s GetDailyBillRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDailyBillRequest) GoString() string {
	return s.String()
}

func (s *GetDailyBillRequest) GetBillOwner() *string {
	return s.BillOwner
}

func (s *GetDailyBillRequest) GetBillType() *string {
	return s.BillType
}

func (s *GetDailyBillRequest) GetDate() *string {
	return s.Date
}

func (s *GetDailyBillRequest) SetBillOwner(v string) *GetDailyBillRequest {
	s.BillOwner = &v
	return s
}

func (s *GetDailyBillRequest) SetBillType(v string) *GetDailyBillRequest {
	s.BillType = &v
	return s
}

func (s *GetDailyBillRequest) SetDate(v string) *GetDailyBillRequest {
	s.Date = &v
	return s
}

func (s *GetDailyBillRequest) Validate() error {
	return dara.Validate(s)
}
