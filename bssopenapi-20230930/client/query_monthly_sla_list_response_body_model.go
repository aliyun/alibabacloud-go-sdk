// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMonthlySlaListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *QueryMonthlySlaListResponseBody
	GetCurrentPage() *int32
	SetData(v []*QueryMonthlySlaListResponseBodyData) *QueryMonthlySlaListResponseBody
	GetData() []*QueryMonthlySlaListResponseBodyData
	SetMetadata(v interface{}) *QueryMonthlySlaListResponseBody
	GetMetadata() interface{}
	SetPageSize(v int32) *QueryMonthlySlaListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *QueryMonthlySlaListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *QueryMonthlySlaListResponseBody
	GetTotalCount() *int32
}

type QueryMonthlySlaListResponseBody struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The data list.
	Data []*QueryMonthlySlaListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The response struct metadata.
	//
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6000EE23-274B-4E07-A697-FF2E999520A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryMonthlySlaListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMonthlySlaListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMonthlySlaListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryMonthlySlaListResponseBody) GetData() []*QueryMonthlySlaListResponseBodyData {
	return s.Data
}

func (s *QueryMonthlySlaListResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *QueryMonthlySlaListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryMonthlySlaListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMonthlySlaListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryMonthlySlaListResponseBody) SetCurrentPage(v int32) *QueryMonthlySlaListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryMonthlySlaListResponseBody) SetData(v []*QueryMonthlySlaListResponseBodyData) *QueryMonthlySlaListResponseBody {
	s.Data = v
	return s
}

func (s *QueryMonthlySlaListResponseBody) SetMetadata(v interface{}) *QueryMonthlySlaListResponseBody {
	s.Metadata = v
	return s
}

func (s *QueryMonthlySlaListResponseBody) SetPageSize(v int32) *QueryMonthlySlaListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryMonthlySlaListResponseBody) SetRequestId(v string) *QueryMonthlySlaListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMonthlySlaListResponseBody) SetTotalCount(v int32) *QueryMonthlySlaListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryMonthlySlaListResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMonthlySlaListResponseBodyData struct {
	// The service availability.
	//
	// example:
	//
	// 99.9155
	AvailableRate *float64 `json:"AvailableRate,omitempty" xml:"AvailableRate,omitempty"`
	// The unique ID of the damage record, used for targeted claims.
	//
	// example:
	//
	// 9b7***9413
	DamagedId *string `json:"DamagedId,omitempty" xml:"DamagedId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// OSSBAG-cn-0xl0n****003
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The month in yyyyMM format.
	//
	// example:
	//
	// 202603
	Month *int32 `json:"Month,omitempty" xml:"Month,omitempty"`
	// The monthly service fee.
	//
	// example:
	//
	// 365.00
	MonthlyServiceCharge *float64 `json:"MonthlyServiceCharge,omitempty" xml:"MonthlyServiceCharge,omitempty"`
	// The compensation description.
	//
	// example:
	//
	// SLA compensation
	PayDescription *string `json:"PayDescription,omitempty" xml:"PayDescription,omitempty"`
	// The compensation ratio, in percentage (%).
	//
	// example:
	//
	// 10
	PayRate *float64 `json:"PayRate,omitempty" xml:"PayRate,omitempty"`
	// The compensation status. Valid values:
	//
	// - 0: not compensated.
	//
	// - 1: compensated.
	//
	// - 2: no compensation required.
	//
	// example:
	//
	// 1
	PayStatus *int32 `json:"PayStatus,omitempty" xml:"PayStatus,omitempty"`
	// The product code.
	//
	// example:
	//
	// oss
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The compensation amount that should be paid.
	//
	// example:
	//
	// 36.5
	ShouldPaySum *float64 `json:"ShouldPaySum,omitempty" xml:"ShouldPaySum,omitempty"`
}

func (s QueryMonthlySlaListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryMonthlySlaListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryMonthlySlaListResponseBodyData) GetAvailableRate() *float64 {
	return s.AvailableRate
}

func (s *QueryMonthlySlaListResponseBodyData) GetDamagedId() *string {
	return s.DamagedId
}

func (s *QueryMonthlySlaListResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryMonthlySlaListResponseBodyData) GetMonth() *int32 {
	return s.Month
}

func (s *QueryMonthlySlaListResponseBodyData) GetMonthlyServiceCharge() *float64 {
	return s.MonthlyServiceCharge
}

func (s *QueryMonthlySlaListResponseBodyData) GetPayDescription() *string {
	return s.PayDescription
}

func (s *QueryMonthlySlaListResponseBodyData) GetPayRate() *float64 {
	return s.PayRate
}

func (s *QueryMonthlySlaListResponseBodyData) GetPayStatus() *int32 {
	return s.PayStatus
}

func (s *QueryMonthlySlaListResponseBodyData) GetProductCode() *string {
	return s.ProductCode
}

func (s *QueryMonthlySlaListResponseBodyData) GetShouldPaySum() *float64 {
	return s.ShouldPaySum
}

func (s *QueryMonthlySlaListResponseBodyData) SetAvailableRate(v float64) *QueryMonthlySlaListResponseBodyData {
	s.AvailableRate = &v
	return s
}

func (s *QueryMonthlySlaListResponseBodyData) SetDamagedId(v string) *QueryMonthlySlaListResponseBodyData {
	s.DamagedId = &v
	return s
}

func (s *QueryMonthlySlaListResponseBodyData) SetInstanceId(v string) *QueryMonthlySlaListResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *QueryMonthlySlaListResponseBodyData) SetMonth(v int32) *QueryMonthlySlaListResponseBodyData {
	s.Month = &v
	return s
}

func (s *QueryMonthlySlaListResponseBodyData) SetMonthlyServiceCharge(v float64) *QueryMonthlySlaListResponseBodyData {
	s.MonthlyServiceCharge = &v
	return s
}

func (s *QueryMonthlySlaListResponseBodyData) SetPayDescription(v string) *QueryMonthlySlaListResponseBodyData {
	s.PayDescription = &v
	return s
}

func (s *QueryMonthlySlaListResponseBodyData) SetPayRate(v float64) *QueryMonthlySlaListResponseBodyData {
	s.PayRate = &v
	return s
}

func (s *QueryMonthlySlaListResponseBodyData) SetPayStatus(v int32) *QueryMonthlySlaListResponseBodyData {
	s.PayStatus = &v
	return s
}

func (s *QueryMonthlySlaListResponseBodyData) SetProductCode(v string) *QueryMonthlySlaListResponseBodyData {
	s.ProductCode = &v
	return s
}

func (s *QueryMonthlySlaListResponseBodyData) SetShouldPaySum(v float64) *QueryMonthlySlaListResponseBodyData {
	s.ShouldPaySum = &v
	return s
}

func (s *QueryMonthlySlaListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
