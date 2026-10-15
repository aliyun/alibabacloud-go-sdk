// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDPUtilizationDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryDPUtilizationDetailResponseBody
	GetCode() *string
	SetData(v *QueryDPUtilizationDetailResponseBodyData) *QueryDPUtilizationDetailResponseBody
	GetData() *QueryDPUtilizationDetailResponseBodyData
	SetMessage(v string) *QueryDPUtilizationDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryDPUtilizationDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryDPUtilizationDetailResponseBody
	GetSuccess() *bool
}

type QueryDPUtilizationDetailResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *QueryDPUtilizationDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s QueryDPUtilizationDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDPUtilizationDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDPUtilizationDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryDPUtilizationDetailResponseBody) GetData() *QueryDPUtilizationDetailResponseBodyData {
	return s.Data
}

func (s *QueryDPUtilizationDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryDPUtilizationDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDPUtilizationDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryDPUtilizationDetailResponseBody) SetCode(v string) *QueryDPUtilizationDetailResponseBody {
	s.Code = &v
	return s
}

func (s *QueryDPUtilizationDetailResponseBody) SetData(v *QueryDPUtilizationDetailResponseBodyData) *QueryDPUtilizationDetailResponseBody {
	s.Data = v
	return s
}

func (s *QueryDPUtilizationDetailResponseBody) SetMessage(v string) *QueryDPUtilizationDetailResponseBody {
	s.Message = &v
	return s
}

func (s *QueryDPUtilizationDetailResponseBody) SetRequestId(v string) *QueryDPUtilizationDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDPUtilizationDetailResponseBody) SetSuccess(v bool) *QueryDPUtilizationDetailResponseBody {
	s.Success = &v
	return s
}

func (s *QueryDPUtilizationDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDPUtilizationDetailResponseBodyData struct {
	DetailList *QueryDPUtilizationDetailResponseBodyDataDetailList `json:"DetailList,omitempty" xml:"DetailList,omitempty" type:"Struct"`
	// The token that specifies the start position of the next query. Pass this value to the LastToken parameter in the next request. If this parameter is returned as null, the query is complete.
	//
	// example:
	//
	// CAESF***zNTAw
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s QueryDPUtilizationDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryDPUtilizationDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryDPUtilizationDetailResponseBodyData) GetDetailList() *QueryDPUtilizationDetailResponseBodyDataDetailList {
	return s.DetailList
}

func (s *QueryDPUtilizationDetailResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryDPUtilizationDetailResponseBodyData) SetDetailList(v *QueryDPUtilizationDetailResponseBodyDataDetailList) *QueryDPUtilizationDetailResponseBodyData {
	s.DetailList = v
	return s
}

func (s *QueryDPUtilizationDetailResponseBodyData) SetNextToken(v string) *QueryDPUtilizationDetailResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *QueryDPUtilizationDetailResponseBodyData) Validate() error {
	if s.DetailList != nil {
		if err := s.DetailList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDPUtilizationDetailResponseBodyDataDetailList struct {
	DetailList []*QueryDPUtilizationDetailResponseBodyDataDetailListDetailList `json:"DetailList,omitempty" xml:"DetailList,omitempty" type:"Repeated"`
}

func (s QueryDPUtilizationDetailResponseBodyDataDetailList) String() string {
	return dara.Prettify(s)
}

func (s QueryDPUtilizationDetailResponseBodyDataDetailList) GoString() string {
	return s.String()
}

func (s *QueryDPUtilizationDetailResponseBodyDataDetailList) GetDetailList() []*QueryDPUtilizationDetailResponseBodyDataDetailListDetailList {
	return s.DetailList
}

func (s *QueryDPUtilizationDetailResponseBodyDataDetailList) SetDetailList(v []*QueryDPUtilizationDetailResponseBodyDataDetailListDetailList) *QueryDPUtilizationDetailResponseBodyDataDetailList {
	s.DetailList = v
	return s
}

func (s *QueryDPUtilizationDetailResponseBodyDataDetailList) Validate() error {
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

type QueryDPUtilizationDetailResponseBodyDataDetailListDetailList struct {
	DeductDate            *string  `json:"DeductDate,omitempty" xml:"DeductDate,omitempty"`
	DeductFactorTotal     *float32 `json:"DeductFactorTotal,omitempty" xml:"DeductFactorTotal,omitempty"`
	DeductHours           *float32 `json:"DeductHours,omitempty" xml:"DeductHours,omitempty"`
	DeductMeasure         *float32 `json:"DeductMeasure,omitempty" xml:"DeductMeasure,omitempty"`
	DeductQuantity        *float32 `json:"DeductQuantity,omitempty" xml:"DeductQuantity,omitempty"`
	DeductedCommodityCode *string  `json:"DeductedCommodityCode,omitempty" xml:"DeductedCommodityCode,omitempty"`
	DeductedInstanceId    *string  `json:"DeductedInstanceId,omitempty" xml:"DeductedInstanceId,omitempty"`
	DeductedProductDetail *string  `json:"DeductedProductDetail,omitempty" xml:"DeductedProductDetail,omitempty"`
	InstanceId            *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceSpec          *string  `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	Region                *string  `json:"Region,omitempty" xml:"Region,omitempty"`
	ResCode               *string  `json:"ResCode,omitempty" xml:"ResCode,omitempty"`
	ShareUid              *int64   `json:"ShareUid,omitempty" xml:"ShareUid,omitempty"`
	Uid                   *int64   `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s QueryDPUtilizationDetailResponseBodyDataDetailListDetailList) String() string {
	return dara.Prettify(s)
}

func (s QueryDPUtilizationDetailResponseBodyDataDetailListDetailList) GoString() string {
	return s.String()
}

func (s *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList) GetDeductDate() *string {
	return s.DeductDate
}

func (s *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList) GetDeductFactorTotal() *float32 {
	return s.DeductFactorTotal
}

func (s *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList) GetDeductHours() *float32 {
	return s.DeductHours
}

func (s *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList) GetDeductMeasure() *float32 {
	return s.DeductMeasure
}

func (s *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList) GetDeductQuantity() *float32 {
	return s.DeductQuantity
}

func (s *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList) GetDeductedCommodityCode() *string {
	return s.DeductedCommodityCode
}

func (s *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList) GetDeductedInstanceId() *string {
	return s.DeductedInstanceId
}

func (s *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList) GetDeductedProductDetail() *string {
	return s.DeductedProductDetail
}

func (s *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList) GetRegion() *string {
	return s.Region
}

func (s *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList) GetResCode() *string {
	return s.ResCode
}

func (s *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList) GetShareUid() *int64 {
	return s.ShareUid
}

func (s *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList) GetUid() *int64 {
	return s.Uid
}

func (s *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList) SetDeductDate(v string) *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList {
	s.DeductDate = &v
	return s
}

func (s *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList) SetDeductFactorTotal(v float32) *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList {
	s.DeductFactorTotal = &v
	return s
}

func (s *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList) SetDeductHours(v float32) *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList {
	s.DeductHours = &v
	return s
}

func (s *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList) SetDeductMeasure(v float32) *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList {
	s.DeductMeasure = &v
	return s
}

func (s *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList) SetDeductQuantity(v float32) *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList {
	s.DeductQuantity = &v
	return s
}

func (s *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList) SetDeductedCommodityCode(v string) *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList {
	s.DeductedCommodityCode = &v
	return s
}

func (s *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList) SetDeductedInstanceId(v string) *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList {
	s.DeductedInstanceId = &v
	return s
}

func (s *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList) SetDeductedProductDetail(v string) *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList {
	s.DeductedProductDetail = &v
	return s
}

func (s *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList) SetInstanceId(v string) *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList {
	s.InstanceId = &v
	return s
}

func (s *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList) SetInstanceSpec(v string) *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList {
	s.InstanceSpec = &v
	return s
}

func (s *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList) SetRegion(v string) *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList {
	s.Region = &v
	return s
}

func (s *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList) SetResCode(v string) *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList {
	s.ResCode = &v
	return s
}

func (s *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList) SetShareUid(v int64) *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList {
	s.ShareUid = &v
	return s
}

func (s *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList) SetUid(v int64) *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList {
	s.Uid = &v
	return s
}

func (s *QueryDPUtilizationDetailResponseBodyDataDetailListDetailList) Validate() error {
	return dara.Validate(s)
}
