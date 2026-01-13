// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMeteringSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillPeriod(v string) *GetMeteringSummaryRequest
	GetBillPeriod() *string
	SetBillingItem(v string) *GetMeteringSummaryRequest
	GetBillingItem() *string
	SetEndTime(v string) *GetMeteringSummaryRequest
	GetEndTime() *string
	SetStartTime(v string) *GetMeteringSummaryRequest
	GetStartTime() *string
	SetSubAccountId(v string) *GetMeteringSummaryRequest
	GetSubAccountId() *string
}

type GetMeteringSummaryRequest struct {
	// example:
	//
	// 202505
	BillPeriod  *string `json:"billPeriod,omitempty" xml:"billPeriod,omitempty"`
	BillingItem *string `json:"billingItem,omitempty" xml:"billingItem,omitempty"`
	// example:
	//
	// 20250522
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 20240920
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 2312312312312
	SubAccountId *string `json:"subAccountId,omitempty" xml:"subAccountId,omitempty"`
}

func (s GetMeteringSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMeteringSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetMeteringSummaryRequest) GetBillPeriod() *string {
	return s.BillPeriod
}

func (s *GetMeteringSummaryRequest) GetBillingItem() *string {
	return s.BillingItem
}

func (s *GetMeteringSummaryRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetMeteringSummaryRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetMeteringSummaryRequest) GetSubAccountId() *string {
	return s.SubAccountId
}

func (s *GetMeteringSummaryRequest) SetBillPeriod(v string) *GetMeteringSummaryRequest {
	s.BillPeriod = &v
	return s
}

func (s *GetMeteringSummaryRequest) SetBillingItem(v string) *GetMeteringSummaryRequest {
	s.BillingItem = &v
	return s
}

func (s *GetMeteringSummaryRequest) SetEndTime(v string) *GetMeteringSummaryRequest {
	s.EndTime = &v
	return s
}

func (s *GetMeteringSummaryRequest) SetStartTime(v string) *GetMeteringSummaryRequest {
	s.StartTime = &v
	return s
}

func (s *GetMeteringSummaryRequest) SetSubAccountId(v string) *GetMeteringSummaryRequest {
	s.SubAccountId = &v
	return s
}

func (s *GetMeteringSummaryRequest) Validate() error {
	return dara.Validate(s)
}
