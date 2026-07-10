// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetaStatisticsListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeMetaStatisticsListResponseBodyItems) *DescribeMetaStatisticsListResponseBody
	GetItems() []*DescribeMetaStatisticsListResponseBodyItems
	SetRequestId(v string) *DescribeMetaStatisticsListResponseBody
	GetRequestId() *string
}

type DescribeMetaStatisticsListResponseBody struct {
	// The list of statistics information.
	Items []*DescribeMetaStatisticsListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 3F860B3F-76B7-5555-A907-2F4433BF8868
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMetaStatisticsListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetaStatisticsListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetaStatisticsListResponseBody) GetItems() []*DescribeMetaStatisticsListResponseBodyItems {
	return s.Items
}

func (s *DescribeMetaStatisticsListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMetaStatisticsListResponseBody) SetItems(v []*DescribeMetaStatisticsListResponseBodyItems) *DescribeMetaStatisticsListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeMetaStatisticsListResponseBody) SetRequestId(v string) *DescribeMetaStatisticsListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetaStatisticsListResponseBody) Validate() error {
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

type DescribeMetaStatisticsListResponseBodyItems struct {
	// The commodity (product) code.
	//
	// example:
	//
	// MOBILE_ONLINE_LENGTH
	Api *string `json:"Api,omitempty" xml:"Api,omitempty"`
	// The name corresponding to the API. Valid values:
	//
	// - **ID_CARD_2_META**: ID card two-element verification
	//
	// - **ID_PERIOD**: ID card validity period verification
	//
	// - **MOBILE_ONLINE_LENGTH**: mobile number online duration
	//
	// - **MOBILE_ONLINE_STATUS**: mobile number online status
	//
	// - **MOBILE_3_META_SIMPLE**: mobile number three-element verification (simple edition)
	//
	// - **MOBILE_3_META**: mobile number three-element verification (detailed edition)
	//
	// - **MOBILE_2_META**: mobile number two-element verification
	//
	// - **BANK_CARD_N_META**: bank card verification (detailed edition)
	//
	// - **MOBILE_DETECT**: phone number detection
	//
	// - **VEHICLE_N_META**: vehicle element verification (enhanced edition)
	//
	// - **VEHICLE_PENTA_INFO**: vehicle five-element information recognition
	//
	// - **VEHICLE_LICENSE_INFO**: vehicle information recognition
	//
	// - **VEHICLE_INSURE_DATE**: vehicle insurance date query
	//
	// - **VEHICLE_CHECK**: vehicle element verification.
	//
	// example:
	//
	// 身份证二要素
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The number of successful queries (billable).
	//
	// example:
	//
	// 0
	BillCount *int64 `json:"BillCount,omitempty" xml:"BillCount,omitempty"`
	// The query hit rate (%).
	//
	// example:
	//
	// 0
	BillRate *string `json:"BillRate,omitempty" xml:"BillRate,omitempty"`
	// The number of successful phone number queries (exclusive to phone number detection).
	//
	// example:
	//
	// 0
	ChargeCount *int64 `json:"ChargeCount,omitempty" xml:"ChargeCount,omitempty"`
	// The date.
	//
	// example:
	//
	// 11/8
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// The name of the telecommunications service provider. Valid values:
	//
	// - **CMCC**: China Mobile
	//
	// - **CUCC**: China Unicom
	//
	// - **CTCC**: China Telecom.
	//
	// example:
	//
	// CUCC
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	// The number of calls with no education information found.
	//
	// example:
	//
	// 0
	NoRecordCount *int64 `json:"NoRecordCount,omitempty" xml:"NoRecordCount,omitempty"`
	// The number of authentication-passed transactions.
	//
	// example:
	//
	// 9
	PassedCount *int64 `json:"PassedCount,omitempty" xml:"PassedCount,omitempty"`
	// The authentication pass rate (%).
	//
	// example:
	//
	// 20
	PassedRate *string `json:"PassedRate,omitempty" xml:"PassedRate,omitempty"`
	// The number of successful calls.
	//
	// example:
	//
	// 1
	SuccessCount *int64 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// The call success rate (%).
	//
	// example:
	//
	// 100
	SuccessRate *string `json:"SuccessRate,omitempty" xml:"SuccessRate,omitempty"`
	// The total number of calls.
	//
	// example:
	//
	// 4
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The number of authentication-failed transactions.
	//
	// example:
	//
	// 2
	UnpassedCount *int64 `json:"UnpassedCount,omitempty" xml:"UnpassedCount,omitempty"`
}

func (s DescribeMetaStatisticsListResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetaStatisticsListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeMetaStatisticsListResponseBodyItems) GetApi() *string {
	return s.Api
}

func (s *DescribeMetaStatisticsListResponseBodyItems) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeMetaStatisticsListResponseBodyItems) GetBillCount() *int64 {
	return s.BillCount
}

func (s *DescribeMetaStatisticsListResponseBodyItems) GetBillRate() *string {
	return s.BillRate
}

func (s *DescribeMetaStatisticsListResponseBodyItems) GetChargeCount() *int64 {
	return s.ChargeCount
}

func (s *DescribeMetaStatisticsListResponseBodyItems) GetDate() *string {
	return s.Date
}

func (s *DescribeMetaStatisticsListResponseBodyItems) GetIspName() *string {
	return s.IspName
}

func (s *DescribeMetaStatisticsListResponseBodyItems) GetNoRecordCount() *int64 {
	return s.NoRecordCount
}

func (s *DescribeMetaStatisticsListResponseBodyItems) GetPassedCount() *int64 {
	return s.PassedCount
}

func (s *DescribeMetaStatisticsListResponseBodyItems) GetPassedRate() *string {
	return s.PassedRate
}

func (s *DescribeMetaStatisticsListResponseBodyItems) GetSuccessCount() *int64 {
	return s.SuccessCount
}

func (s *DescribeMetaStatisticsListResponseBodyItems) GetSuccessRate() *string {
	return s.SuccessRate
}

func (s *DescribeMetaStatisticsListResponseBodyItems) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeMetaStatisticsListResponseBodyItems) GetUnpassedCount() *int64 {
	return s.UnpassedCount
}

func (s *DescribeMetaStatisticsListResponseBodyItems) SetApi(v string) *DescribeMetaStatisticsListResponseBodyItems {
	s.Api = &v
	return s
}

func (s *DescribeMetaStatisticsListResponseBodyItems) SetApiName(v string) *DescribeMetaStatisticsListResponseBodyItems {
	s.ApiName = &v
	return s
}

func (s *DescribeMetaStatisticsListResponseBodyItems) SetBillCount(v int64) *DescribeMetaStatisticsListResponseBodyItems {
	s.BillCount = &v
	return s
}

func (s *DescribeMetaStatisticsListResponseBodyItems) SetBillRate(v string) *DescribeMetaStatisticsListResponseBodyItems {
	s.BillRate = &v
	return s
}

func (s *DescribeMetaStatisticsListResponseBodyItems) SetChargeCount(v int64) *DescribeMetaStatisticsListResponseBodyItems {
	s.ChargeCount = &v
	return s
}

func (s *DescribeMetaStatisticsListResponseBodyItems) SetDate(v string) *DescribeMetaStatisticsListResponseBodyItems {
	s.Date = &v
	return s
}

func (s *DescribeMetaStatisticsListResponseBodyItems) SetIspName(v string) *DescribeMetaStatisticsListResponseBodyItems {
	s.IspName = &v
	return s
}

func (s *DescribeMetaStatisticsListResponseBodyItems) SetNoRecordCount(v int64) *DescribeMetaStatisticsListResponseBodyItems {
	s.NoRecordCount = &v
	return s
}

func (s *DescribeMetaStatisticsListResponseBodyItems) SetPassedCount(v int64) *DescribeMetaStatisticsListResponseBodyItems {
	s.PassedCount = &v
	return s
}

func (s *DescribeMetaStatisticsListResponseBodyItems) SetPassedRate(v string) *DescribeMetaStatisticsListResponseBodyItems {
	s.PassedRate = &v
	return s
}

func (s *DescribeMetaStatisticsListResponseBodyItems) SetSuccessCount(v int64) *DescribeMetaStatisticsListResponseBodyItems {
	s.SuccessCount = &v
	return s
}

func (s *DescribeMetaStatisticsListResponseBodyItems) SetSuccessRate(v string) *DescribeMetaStatisticsListResponseBodyItems {
	s.SuccessRate = &v
	return s
}

func (s *DescribeMetaStatisticsListResponseBodyItems) SetTotalCount(v int64) *DescribeMetaStatisticsListResponseBodyItems {
	s.TotalCount = &v
	return s
}

func (s *DescribeMetaStatisticsListResponseBodyItems) SetUnpassedCount(v int64) *DescribeMetaStatisticsListResponseBodyItems {
	s.UnpassedCount = &v
	return s
}

func (s *DescribeMetaStatisticsListResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
