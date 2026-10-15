// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRIUtilizationDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryRIUtilizationDetailResponseBody
	GetCode() *string
	SetData(v *QueryRIUtilizationDetailResponseBodyData) *QueryRIUtilizationDetailResponseBody
	GetData() *QueryRIUtilizationDetailResponseBodyData
	SetMessage(v string) *QueryRIUtilizationDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryRIUtilizationDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryRIUtilizationDetailResponseBody
	GetSuccess() *bool
}

type QueryRIUtilizationDetailResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *QueryRIUtilizationDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// Successful！
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DF58589C-A06C-4224-8615-7797E6474FA3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRIUtilizationDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRIUtilizationDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRIUtilizationDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryRIUtilizationDetailResponseBody) GetData() *QueryRIUtilizationDetailResponseBodyData {
	return s.Data
}

func (s *QueryRIUtilizationDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryRIUtilizationDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRIUtilizationDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryRIUtilizationDetailResponseBody) SetCode(v string) *QueryRIUtilizationDetailResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRIUtilizationDetailResponseBody) SetData(v *QueryRIUtilizationDetailResponseBodyData) *QueryRIUtilizationDetailResponseBody {
	s.Data = v
	return s
}

func (s *QueryRIUtilizationDetailResponseBody) SetMessage(v string) *QueryRIUtilizationDetailResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRIUtilizationDetailResponseBody) SetRequestId(v string) *QueryRIUtilizationDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRIUtilizationDetailResponseBody) SetSuccess(v bool) *QueryRIUtilizationDetailResponseBody {
	s.Success = &v
	return s
}

func (s *QueryRIUtilizationDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryRIUtilizationDetailResponseBodyData struct {
	DetailList *QueryRIUtilizationDetailResponseBodyDataDetailList `json:"DetailList,omitempty" xml:"DetailList,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 0
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryRIUtilizationDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryRIUtilizationDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryRIUtilizationDetailResponseBodyData) GetDetailList() *QueryRIUtilizationDetailResponseBodyDataDetailList {
	return s.DetailList
}

func (s *QueryRIUtilizationDetailResponseBodyData) GetPageNum() *int64 {
	return s.PageNum
}

func (s *QueryRIUtilizationDetailResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryRIUtilizationDetailResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *QueryRIUtilizationDetailResponseBodyData) SetDetailList(v *QueryRIUtilizationDetailResponseBodyDataDetailList) *QueryRIUtilizationDetailResponseBodyData {
	s.DetailList = v
	return s
}

func (s *QueryRIUtilizationDetailResponseBodyData) SetPageNum(v int64) *QueryRIUtilizationDetailResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryRIUtilizationDetailResponseBodyData) SetPageSize(v int64) *QueryRIUtilizationDetailResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryRIUtilizationDetailResponseBodyData) SetTotalCount(v int64) *QueryRIUtilizationDetailResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryRIUtilizationDetailResponseBodyData) Validate() error {
	if s.DetailList != nil {
		if err := s.DetailList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryRIUtilizationDetailResponseBodyDataDetailList struct {
	DetailList []*QueryRIUtilizationDetailResponseBodyDataDetailListDetailList `json:"DetailList,omitempty" xml:"DetailList,omitempty" type:"Repeated"`
}

func (s QueryRIUtilizationDetailResponseBodyDataDetailList) String() string {
	return dara.Prettify(s)
}

func (s QueryRIUtilizationDetailResponseBodyDataDetailList) GoString() string {
	return s.String()
}

func (s *QueryRIUtilizationDetailResponseBodyDataDetailList) GetDetailList() []*QueryRIUtilizationDetailResponseBodyDataDetailListDetailList {
	return s.DetailList
}

func (s *QueryRIUtilizationDetailResponseBodyDataDetailList) SetDetailList(v []*QueryRIUtilizationDetailResponseBodyDataDetailListDetailList) *QueryRIUtilizationDetailResponseBodyDataDetailList {
	s.DetailList = v
	return s
}

func (s *QueryRIUtilizationDetailResponseBodyDataDetailList) Validate() error {
	if s.DetailList != nil {
		for _, item := range s.DetailList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryRIUtilizationDetailResponseBodyDataDetailListDetailList struct {
	DeductDate            *string  `json:"DeductDate,omitempty" xml:"DeductDate,omitempty"`
	DeductFactorTotal     *float32 `json:"DeductFactorTotal,omitempty" xml:"DeductFactorTotal,omitempty"`
	DeductHours           *string  `json:"DeductHours,omitempty" xml:"DeductHours,omitempty"`
	DeductQuantity        *float32 `json:"DeductQuantity,omitempty" xml:"DeductQuantity,omitempty"`
	DeductedCommodityCode *string  `json:"DeductedCommodityCode,omitempty" xml:"DeductedCommodityCode,omitempty"`
	DeductedInstanceId    *string  `json:"DeductedInstanceId,omitempty" xml:"DeductedInstanceId,omitempty"`
	DeductedProductDetail *string  `json:"DeductedProductDetail,omitempty" xml:"DeductedProductDetail,omitempty"`
	InstanceSpec          *string  `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	RIInstanceId          *string  `json:"RIInstanceId,omitempty" xml:"RIInstanceId,omitempty"`
}

func (s QueryRIUtilizationDetailResponseBodyDataDetailListDetailList) String() string {
	return dara.Prettify(s)
}

func (s QueryRIUtilizationDetailResponseBodyDataDetailListDetailList) GoString() string {
	return s.String()
}

func (s *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList) GetDeductDate() *string {
	return s.DeductDate
}

func (s *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList) GetDeductFactorTotal() *float32 {
	return s.DeductFactorTotal
}

func (s *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList) GetDeductHours() *string {
	return s.DeductHours
}

func (s *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList) GetDeductQuantity() *float32 {
	return s.DeductQuantity
}

func (s *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList) GetDeductedCommodityCode() *string {
	return s.DeductedCommodityCode
}

func (s *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList) GetDeductedInstanceId() *string {
	return s.DeductedInstanceId
}

func (s *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList) GetDeductedProductDetail() *string {
	return s.DeductedProductDetail
}

func (s *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList) GetRIInstanceId() *string {
	return s.RIInstanceId
}

func (s *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList) SetDeductDate(v string) *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList {
	s.DeductDate = &v
	return s
}

func (s *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList) SetDeductFactorTotal(v float32) *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList {
	s.DeductFactorTotal = &v
	return s
}

func (s *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList) SetDeductHours(v string) *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList {
	s.DeductHours = &v
	return s
}

func (s *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList) SetDeductQuantity(v float32) *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList {
	s.DeductQuantity = &v
	return s
}

func (s *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList) SetDeductedCommodityCode(v string) *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList {
	s.DeductedCommodityCode = &v
	return s
}

func (s *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList) SetDeductedInstanceId(v string) *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList {
	s.DeductedInstanceId = &v
	return s
}

func (s *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList) SetDeductedProductDetail(v string) *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList {
	s.DeductedProductDetail = &v
	return s
}

func (s *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList) SetInstanceSpec(v string) *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList {
	s.InstanceSpec = &v
	return s
}

func (s *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList) SetRIInstanceId(v string) *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList {
	s.RIInstanceId = &v
	return s
}

func (s *QueryRIUtilizationDetailResponseBodyDataDetailListDetailList) Validate() error {
	return dara.Validate(s)
}
