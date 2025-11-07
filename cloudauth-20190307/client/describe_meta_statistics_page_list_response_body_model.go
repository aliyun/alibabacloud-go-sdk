// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetaStatisticsPageListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeMetaStatisticsPageListResponseBody
	GetCurrentPage() *int32
	SetItems(v []*DescribeMetaStatisticsPageListResponseBodyItems) *DescribeMetaStatisticsPageListResponseBody
	GetItems() []*DescribeMetaStatisticsPageListResponseBodyItems
	SetPageSize(v int32) *DescribeMetaStatisticsPageListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeMetaStatisticsPageListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeMetaStatisticsPageListResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *DescribeMetaStatisticsPageListResponseBody
	GetTotalPage() *int32
}

type DescribeMetaStatisticsPageListResponseBody struct {
	// example:
	//
	// 3
	CurrentPage *int32                                             `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Items       []*DescribeMetaStatisticsPageListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// C379C9E4-4DA0-5D0B-821B-25E2B8693D48
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 7
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 1
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeMetaStatisticsPageListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetaStatisticsPageListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetaStatisticsPageListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeMetaStatisticsPageListResponseBody) GetItems() []*DescribeMetaStatisticsPageListResponseBodyItems {
	return s.Items
}

func (s *DescribeMetaStatisticsPageListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMetaStatisticsPageListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMetaStatisticsPageListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeMetaStatisticsPageListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeMetaStatisticsPageListResponseBody) SetCurrentPage(v int32) *DescribeMetaStatisticsPageListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeMetaStatisticsPageListResponseBody) SetItems(v []*DescribeMetaStatisticsPageListResponseBodyItems) *DescribeMetaStatisticsPageListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeMetaStatisticsPageListResponseBody) SetPageSize(v int32) *DescribeMetaStatisticsPageListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeMetaStatisticsPageListResponseBody) SetRequestId(v string) *DescribeMetaStatisticsPageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetaStatisticsPageListResponseBody) SetTotalCount(v int32) *DescribeMetaStatisticsPageListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeMetaStatisticsPageListResponseBody) SetTotalPage(v int32) *DescribeMetaStatisticsPageListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeMetaStatisticsPageListResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMetaStatisticsPageListResponseBodyItems struct {
	// Api。
	//
	// example:
	//
	// ID_PERIOD
	Api     *string `json:"Api,omitempty" xml:"Api,omitempty"`
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// example:
	//
	// 10
	BillCount *int64 `json:"BillCount,omitempty" xml:"BillCount,omitempty"`
	// example:
	//
	// 80
	BillRate *string `json:"BillRate,omitempty" xml:"BillRate,omitempty"`
	// example:
	//
	// 1
	ChargeCount *int64 `json:"ChargeCount,omitempty" xml:"ChargeCount,omitempty"`
	// example:
	//
	// 11/8
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// example:
	//
	// CMCC
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	// example:
	//
	// 12
	NoRecordCount *int64 `json:"NoRecordCount,omitempty" xml:"NoRecordCount,omitempty"`
	// example:
	//
	// 21
	PassedCount *int64 `json:"PassedCount,omitempty" xml:"PassedCount,omitempty"`
	// example:
	//
	// 80
	PassedRate *string `json:"PassedRate,omitempty" xml:"PassedRate,omitempty"`
	// example:
	//
	// 1
	SuccessCount *int64 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// example:
	//
	// 100
	SuccessRate *string `json:"SuccessRate,omitempty" xml:"SuccessRate,omitempty"`
	// example:
	//
	// 4
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 2
	UnpassedCount *int64 `json:"UnpassedCount,omitempty" xml:"UnpassedCount,omitempty"`
}

func (s DescribeMetaStatisticsPageListResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetaStatisticsPageListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeMetaStatisticsPageListResponseBodyItems) GetApi() *string {
	return s.Api
}

func (s *DescribeMetaStatisticsPageListResponseBodyItems) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeMetaStatisticsPageListResponseBodyItems) GetBillCount() *int64 {
	return s.BillCount
}

func (s *DescribeMetaStatisticsPageListResponseBodyItems) GetBillRate() *string {
	return s.BillRate
}

func (s *DescribeMetaStatisticsPageListResponseBodyItems) GetChargeCount() *int64 {
	return s.ChargeCount
}

func (s *DescribeMetaStatisticsPageListResponseBodyItems) GetDate() *string {
	return s.Date
}

func (s *DescribeMetaStatisticsPageListResponseBodyItems) GetIspName() *string {
	return s.IspName
}

func (s *DescribeMetaStatisticsPageListResponseBodyItems) GetNoRecordCount() *int64 {
	return s.NoRecordCount
}

func (s *DescribeMetaStatisticsPageListResponseBodyItems) GetPassedCount() *int64 {
	return s.PassedCount
}

func (s *DescribeMetaStatisticsPageListResponseBodyItems) GetPassedRate() *string {
	return s.PassedRate
}

func (s *DescribeMetaStatisticsPageListResponseBodyItems) GetSuccessCount() *int64 {
	return s.SuccessCount
}

func (s *DescribeMetaStatisticsPageListResponseBodyItems) GetSuccessRate() *string {
	return s.SuccessRate
}

func (s *DescribeMetaStatisticsPageListResponseBodyItems) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeMetaStatisticsPageListResponseBodyItems) GetUnpassedCount() *int64 {
	return s.UnpassedCount
}

func (s *DescribeMetaStatisticsPageListResponseBodyItems) SetApi(v string) *DescribeMetaStatisticsPageListResponseBodyItems {
	s.Api = &v
	return s
}

func (s *DescribeMetaStatisticsPageListResponseBodyItems) SetApiName(v string) *DescribeMetaStatisticsPageListResponseBodyItems {
	s.ApiName = &v
	return s
}

func (s *DescribeMetaStatisticsPageListResponseBodyItems) SetBillCount(v int64) *DescribeMetaStatisticsPageListResponseBodyItems {
	s.BillCount = &v
	return s
}

func (s *DescribeMetaStatisticsPageListResponseBodyItems) SetBillRate(v string) *DescribeMetaStatisticsPageListResponseBodyItems {
	s.BillRate = &v
	return s
}

func (s *DescribeMetaStatisticsPageListResponseBodyItems) SetChargeCount(v int64) *DescribeMetaStatisticsPageListResponseBodyItems {
	s.ChargeCount = &v
	return s
}

func (s *DescribeMetaStatisticsPageListResponseBodyItems) SetDate(v string) *DescribeMetaStatisticsPageListResponseBodyItems {
	s.Date = &v
	return s
}

func (s *DescribeMetaStatisticsPageListResponseBodyItems) SetIspName(v string) *DescribeMetaStatisticsPageListResponseBodyItems {
	s.IspName = &v
	return s
}

func (s *DescribeMetaStatisticsPageListResponseBodyItems) SetNoRecordCount(v int64) *DescribeMetaStatisticsPageListResponseBodyItems {
	s.NoRecordCount = &v
	return s
}

func (s *DescribeMetaStatisticsPageListResponseBodyItems) SetPassedCount(v int64) *DescribeMetaStatisticsPageListResponseBodyItems {
	s.PassedCount = &v
	return s
}

func (s *DescribeMetaStatisticsPageListResponseBodyItems) SetPassedRate(v string) *DescribeMetaStatisticsPageListResponseBodyItems {
	s.PassedRate = &v
	return s
}

func (s *DescribeMetaStatisticsPageListResponseBodyItems) SetSuccessCount(v int64) *DescribeMetaStatisticsPageListResponseBodyItems {
	s.SuccessCount = &v
	return s
}

func (s *DescribeMetaStatisticsPageListResponseBodyItems) SetSuccessRate(v string) *DescribeMetaStatisticsPageListResponseBodyItems {
	s.SuccessRate = &v
	return s
}

func (s *DescribeMetaStatisticsPageListResponseBodyItems) SetTotalCount(v int64) *DescribeMetaStatisticsPageListResponseBodyItems {
	s.TotalCount = &v
	return s
}

func (s *DescribeMetaStatisticsPageListResponseBodyItems) SetUnpassedCount(v int64) *DescribeMetaStatisticsPageListResponseBodyItems {
	s.UnpassedCount = &v
	return s
}

func (s *DescribeMetaStatisticsPageListResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
