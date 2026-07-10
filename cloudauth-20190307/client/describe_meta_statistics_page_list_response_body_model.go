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
	// Current page number.
	//
	// example:
	//
	// 3
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Paginated list data.
	Items []*DescribeMetaStatisticsPageListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// Number of data entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of this request.
	//
	// example:
	//
	// C379C9E4-4DA0-5D0B-821B-25E2B8693D48
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 7
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Total number of pages.
	//
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
	// API.
	//
	// example:
	//
	// ID_PERIOD
	Api *string `json:"Api,omitempty" xml:"Api,omitempty"`
	// Name corresponding to the API:
	//
	// - **ID_CARD_2_META**: ID Card Two-Element Verification
	//
	// - **ID_PERIOD**: ID Card Validity Verification Period
	//
	// - **MOBILE_ONLINE_LENGTH**: Mobile Online Duration
	//
	// - **MOBILE_ONLINE_STATUS**: Mobile Online Status
	//
	// - **MOBILE_3_META_SIMPLE**: Mobile Number Three-Element Verification (Simple)
	//
	// - **MOBILE_3_META**: Mobile Number Three-Element Verification (Detailed)
	//
	// - **MOBILE_2_META**: Mobile Number Two-Element Verification
	//
	// - **BANK_CARD_N_META**: Bank Card Verification (Detailed)
	//
	// - **MOBILE_DETECT**: Number Detection
	//
	//  -**VEHICLE_N_META**: Vehicle Element Verification (Enhanced)
	//
	// - **VEHICLE_PENTA_INFO**: Vehicle Five-Element Information Recognition
	//
	// - **VEHICLE_LICENSE_INFO**: Vehicle Information Recognition
	//
	// - **VEHICLE_INSURE_DATE**: Vehicle Insurance Date Query
	//
	// - **VEHICLE_CHECK**: Vehicle Element Verification
	//
	// example:
	//
	// 车辆要素核验增强版
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// Number of hits (billed).
	//
	// example:
	//
	// 10
	BillCount *int64 `json:"BillCount,omitempty" xml:"BillCount,omitempty"`
	// Hit rate (%).
	//
	// example:
	//
	// 80
	BillRate *string `json:"BillRate,omitempty" xml:"BillRate,omitempty"`
	// Number of successful mobile number queries (exclusive to Number Detection).
	//
	// example:
	//
	// 1
	ChargeCount *int64 `json:"ChargeCount,omitempty" xml:"ChargeCount,omitempty"`
	// Date.
	//
	// example:
	//
	// 11/8
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// Carrier name:
	//
	// - **CMCC**: China Mobile
	//
	// - **CUCC**: China Unicom
	//
	// - **CTCC**: China Telecom
	//
	// example:
	//
	// CMCC
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	// Number of calls with no record information.
	//
	// example:
	//
	// 12
	NoRecordCount *int64 `json:"NoRecordCount,omitempty" xml:"NoRecordCount,omitempty"`
	// Number of passed authentications.
	//
	// example:
	//
	// 21
	PassedCount *int64 `json:"PassedCount,omitempty" xml:"PassedCount,omitempty"`
	// Authentication pass rate (%).
	//
	// example:
	//
	// 80
	PassedRate *string `json:"PassedRate,omitempty" xml:"PassedRate,omitempty"`
	// Number of successful requests.
	//
	// example:
	//
	// 1
	SuccessCount *int64 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// Call success rate (%).
	//
	// example:
	//
	// 100
	SuccessRate *string `json:"SuccessRate,omitempty" xml:"SuccessRate,omitempty"`
	// Total number of entries.
	//
	// example:
	//
	// 4
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Number of failed calls.
	//
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
