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
	Items []*DescribeMetaStatisticsListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
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
	// example:
	//
	// MOBILE_ONLINE_LENGTH
	Api     *string `json:"Api,omitempty" xml:"Api,omitempty"`
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// example:
	//
	// 0
	BillCount *int64 `json:"BillCount,omitempty" xml:"BillCount,omitempty"`
	// example:
	//
	// 0
	BillRate *string `json:"BillRate,omitempty" xml:"BillRate,omitempty"`
	// example:
	//
	// 0
	ChargeCount *int64 `json:"ChargeCount,omitempty" xml:"ChargeCount,omitempty"`
	// example:
	//
	// 11/8
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// example:
	//
	// CUCC
	IspName *string `json:"IspName,omitempty" xml:"IspName,omitempty"`
	// example:
	//
	// 0
	NoRecordCount *int64 `json:"NoRecordCount,omitempty" xml:"NoRecordCount,omitempty"`
	// example:
	//
	// 9
	PassedCount *int64 `json:"PassedCount,omitempty" xml:"PassedCount,omitempty"`
	// example:
	//
	// 20
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
