// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMeteringDailyDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillPeriod(v string) *ListMeteringDailyDetailRequest
	GetBillPeriod() *string
	SetBillingItem(v string) *ListMeteringDailyDetailRequest
	GetBillingItem() *string
	SetEndTime(v string) *ListMeteringDailyDetailRequest
	GetEndTime() *string
	SetPageNumber(v int32) *ListMeteringDailyDetailRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMeteringDailyDetailRequest
	GetPageSize() *int32
	SetStartTime(v string) *ListMeteringDailyDetailRequest
	GetStartTime() *string
	SetSubAccountId(v string) *ListMeteringDailyDetailRequest
	GetSubAccountId() *string
}

type ListMeteringDailyDetailRequest struct {
	// example:
	//
	// 202506
	BillPeriod  *string `json:"billPeriod,omitempty" xml:"billPeriod,omitempty"`
	BillingItem *string `json:"billingItem,omitempty" xml:"billingItem,omitempty"`
	// example:
	//
	// 20240908
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 20240908
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 23432423423434
	SubAccountId *string `json:"subAccountId,omitempty" xml:"subAccountId,omitempty"`
}

func (s ListMeteringDailyDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMeteringDailyDetailRequest) GoString() string {
	return s.String()
}

func (s *ListMeteringDailyDetailRequest) GetBillPeriod() *string {
	return s.BillPeriod
}

func (s *ListMeteringDailyDetailRequest) GetBillingItem() *string {
	return s.BillingItem
}

func (s *ListMeteringDailyDetailRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListMeteringDailyDetailRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMeteringDailyDetailRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMeteringDailyDetailRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListMeteringDailyDetailRequest) GetSubAccountId() *string {
	return s.SubAccountId
}

func (s *ListMeteringDailyDetailRequest) SetBillPeriod(v string) *ListMeteringDailyDetailRequest {
	s.BillPeriod = &v
	return s
}

func (s *ListMeteringDailyDetailRequest) SetBillingItem(v string) *ListMeteringDailyDetailRequest {
	s.BillingItem = &v
	return s
}

func (s *ListMeteringDailyDetailRequest) SetEndTime(v string) *ListMeteringDailyDetailRequest {
	s.EndTime = &v
	return s
}

func (s *ListMeteringDailyDetailRequest) SetPageNumber(v int32) *ListMeteringDailyDetailRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMeteringDailyDetailRequest) SetPageSize(v int32) *ListMeteringDailyDetailRequest {
	s.PageSize = &v
	return s
}

func (s *ListMeteringDailyDetailRequest) SetStartTime(v string) *ListMeteringDailyDetailRequest {
	s.StartTime = &v
	return s
}

func (s *ListMeteringDailyDetailRequest) SetSubAccountId(v string) *ListMeteringDailyDetailRequest {
	s.SubAccountId = &v
	return s
}

func (s *ListMeteringDailyDetailRequest) Validate() error {
	return dara.Validate(s)
}
