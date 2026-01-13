// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMeteringDailyDetailResult interface {
	dara.Model
	String() string
	GoString() string
	SetMeteringDailyDetailList(v []*ListMeteringDailyDetailResultMeteringDailyDetailList) *ListMeteringDailyDetailResult
	GetMeteringDailyDetailList() []*ListMeteringDailyDetailResultMeteringDailyDetailList
}

type ListMeteringDailyDetailResult struct {
	MeteringDailyDetailList []*ListMeteringDailyDetailResultMeteringDailyDetailList `json:"meteringDailyDetailList,omitempty" xml:"meteringDailyDetailList,omitempty" type:"Repeated"`
}

func (s ListMeteringDailyDetailResult) String() string {
	return dara.Prettify(s)
}

func (s ListMeteringDailyDetailResult) GoString() string {
	return s.String()
}

func (s *ListMeteringDailyDetailResult) GetMeteringDailyDetailList() []*ListMeteringDailyDetailResultMeteringDailyDetailList {
	return s.MeteringDailyDetailList
}

func (s *ListMeteringDailyDetailResult) SetMeteringDailyDetailList(v []*ListMeteringDailyDetailResultMeteringDailyDetailList) *ListMeteringDailyDetailResult {
	s.MeteringDailyDetailList = v
	return s
}

func (s *ListMeteringDailyDetailResult) Validate() error {
	if s.MeteringDailyDetailList != nil {
		for _, item := range s.MeteringDailyDetailList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMeteringDailyDetailResultMeteringDailyDetailList struct {
	BillDate            *string `json:"billDate,omitempty" xml:"billDate,omitempty"`
	BillPeriod          *string `json:"billPeriod,omitempty" xml:"billPeriod,omitempty"`
	BillingFunctionItem *string `json:"billingFunctionItem,omitempty" xml:"billingFunctionItem,omitempty"`
	BillingItem         *string `json:"billingItem,omitempty" xml:"billingItem,omitempty"`
	BillingQps          *int32  `json:"billingQps,omitempty" xml:"billingQps,omitempty"`
	FailedCount         *int32  `json:"failedCount,omitempty" xml:"failedCount,omitempty"`
	LadderTypeCode      *int32  `json:"ladderTypeCode,omitempty" xml:"ladderTypeCode,omitempty"`
	SubAccountId        *string `json:"subAccountId,omitempty" xml:"subAccountId,omitempty"`
	SuccessCount        *int32  `json:"successCount,omitempty" xml:"successCount,omitempty"`
}

func (s ListMeteringDailyDetailResultMeteringDailyDetailList) String() string {
	return dara.Prettify(s)
}

func (s ListMeteringDailyDetailResultMeteringDailyDetailList) GoString() string {
	return s.String()
}

func (s *ListMeteringDailyDetailResultMeteringDailyDetailList) GetBillDate() *string {
	return s.BillDate
}

func (s *ListMeteringDailyDetailResultMeteringDailyDetailList) GetBillPeriod() *string {
	return s.BillPeriod
}

func (s *ListMeteringDailyDetailResultMeteringDailyDetailList) GetBillingFunctionItem() *string {
	return s.BillingFunctionItem
}

func (s *ListMeteringDailyDetailResultMeteringDailyDetailList) GetBillingItem() *string {
	return s.BillingItem
}

func (s *ListMeteringDailyDetailResultMeteringDailyDetailList) GetBillingQps() *int32 {
	return s.BillingQps
}

func (s *ListMeteringDailyDetailResultMeteringDailyDetailList) GetFailedCount() *int32 {
	return s.FailedCount
}

func (s *ListMeteringDailyDetailResultMeteringDailyDetailList) GetLadderTypeCode() *int32 {
	return s.LadderTypeCode
}

func (s *ListMeteringDailyDetailResultMeteringDailyDetailList) GetSubAccountId() *string {
	return s.SubAccountId
}

func (s *ListMeteringDailyDetailResultMeteringDailyDetailList) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *ListMeteringDailyDetailResultMeteringDailyDetailList) SetBillDate(v string) *ListMeteringDailyDetailResultMeteringDailyDetailList {
	s.BillDate = &v
	return s
}

func (s *ListMeteringDailyDetailResultMeteringDailyDetailList) SetBillPeriod(v string) *ListMeteringDailyDetailResultMeteringDailyDetailList {
	s.BillPeriod = &v
	return s
}

func (s *ListMeteringDailyDetailResultMeteringDailyDetailList) SetBillingFunctionItem(v string) *ListMeteringDailyDetailResultMeteringDailyDetailList {
	s.BillingFunctionItem = &v
	return s
}

func (s *ListMeteringDailyDetailResultMeteringDailyDetailList) SetBillingItem(v string) *ListMeteringDailyDetailResultMeteringDailyDetailList {
	s.BillingItem = &v
	return s
}

func (s *ListMeteringDailyDetailResultMeteringDailyDetailList) SetBillingQps(v int32) *ListMeteringDailyDetailResultMeteringDailyDetailList {
	s.BillingQps = &v
	return s
}

func (s *ListMeteringDailyDetailResultMeteringDailyDetailList) SetFailedCount(v int32) *ListMeteringDailyDetailResultMeteringDailyDetailList {
	s.FailedCount = &v
	return s
}

func (s *ListMeteringDailyDetailResultMeteringDailyDetailList) SetLadderTypeCode(v int32) *ListMeteringDailyDetailResultMeteringDailyDetailList {
	s.LadderTypeCode = &v
	return s
}

func (s *ListMeteringDailyDetailResultMeteringDailyDetailList) SetSubAccountId(v string) *ListMeteringDailyDetailResultMeteringDailyDetailList {
	s.SubAccountId = &v
	return s
}

func (s *ListMeteringDailyDetailResultMeteringDailyDetailList) SetSuccessCount(v int32) *ListMeteringDailyDetailResultMeteringDailyDetailList {
	s.SuccessCount = &v
	return s
}

func (s *ListMeteringDailyDetailResultMeteringDailyDetailList) Validate() error {
	return dara.Validate(s)
}
