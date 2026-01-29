// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySavingsPlansDiscountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QuerySavingsPlansDiscountResponseBody
	GetCode() *string
	SetData(v *QuerySavingsPlansDiscountResponseBodyData) *QuerySavingsPlansDiscountResponseBody
	GetData() *QuerySavingsPlansDiscountResponseBodyData
	SetMessage(v string) *QuerySavingsPlansDiscountResponseBody
	GetMessage() *string
	SetRequestId(v string) *QuerySavingsPlansDiscountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QuerySavingsPlansDiscountResponseBody
	GetSuccess() *bool
}

type QuerySavingsPlansDiscountResponseBody struct {
	// The error code returned.
	//
	// example:
	//
	// PARAM_ERROR
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// data
	Data *QuerySavingsPlansDiscountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// This API is not applicable for caller.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6000EE23-274B-4E07-A697-FF2E999520A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySavingsPlansDiscountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySavingsPlansDiscountResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySavingsPlansDiscountResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySavingsPlansDiscountResponseBody) GetData() *QuerySavingsPlansDiscountResponseBodyData {
	return s.Data
}

func (s *QuerySavingsPlansDiscountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySavingsPlansDiscountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySavingsPlansDiscountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QuerySavingsPlansDiscountResponseBody) SetCode(v string) *QuerySavingsPlansDiscountResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySavingsPlansDiscountResponseBody) SetData(v *QuerySavingsPlansDiscountResponseBodyData) *QuerySavingsPlansDiscountResponseBody {
	s.Data = v
	return s
}

func (s *QuerySavingsPlansDiscountResponseBody) SetMessage(v string) *QuerySavingsPlansDiscountResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySavingsPlansDiscountResponseBody) SetRequestId(v string) *QuerySavingsPlansDiscountResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySavingsPlansDiscountResponseBody) SetSuccess(v bool) *QuerySavingsPlansDiscountResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySavingsPlansDiscountResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QuerySavingsPlansDiscountResponseBodyData struct {
	// The IP address of the request.
	//
	// example:
	//
	// 100.104.180.109
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The information about the discounts on saving plans.
	Items []*QuerySavingsPlansDiscountResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s QuerySavingsPlansDiscountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QuerySavingsPlansDiscountResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySavingsPlansDiscountResponseBodyData) GetHostId() *string {
	return s.HostId
}

func (s *QuerySavingsPlansDiscountResponseBodyData) GetItems() []*QuerySavingsPlansDiscountResponseBodyDataItems {
	return s.Items
}

func (s *QuerySavingsPlansDiscountResponseBodyData) SetHostId(v string) *QuerySavingsPlansDiscountResponseBodyData {
	s.HostId = &v
	return s
}

func (s *QuerySavingsPlansDiscountResponseBodyData) SetItems(v []*QuerySavingsPlansDiscountResponseBodyDataItems) *QuerySavingsPlansDiscountResponseBodyData {
	s.Items = v
	return s
}

func (s *QuerySavingsPlansDiscountResponseBodyData) Validate() error {
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

type QuerySavingsPlansDiscountResponseBodyDataItems struct {
	// The details of the service.
	//
	// example:
	//
	// Pay-as-you-go Elastic Compute Service (ECS) instance
	CommodityName *string `json:"CommodityName,omitempty" xml:"CommodityName,omitempty"`
	// The contracted discount.
	//
	// example:
	//
	// 0.85
	ContractDiscountRate *string `json:"ContractDiscountRate,omitempty" xml:"ContractDiscountRate,omitempty"`
	// The cycle based on which queries were performed.
	//
	// example:
	//
	// 1:Year
	Cycle *string `json:"Cycle,omitempty" xml:"Cycle,omitempty"`
	// The discount provided by the official website.
	//
	// example:
	//
	// 0.85
	DiscountRate *string `json:"DiscountRate,omitempty" xml:"DiscountRate,omitempty"`
	// The name of the pricing module.
	//
	// example:
	//
	// Instance
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// The payment mode. Valid values:
	//
	// 	- total: all upfront
	//
	// 	- half: half upfront
	//
	// 	- zero: no upfront
	//
	// example:
	//
	// total
	PayMode *string `json:"PayMode,omitempty" xml:"PayMode,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-zhangjiakou-na62-a01
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The region ID of the instance. You can call the [DescribeDBInstanceAttribute](https://help.aliyun.com/document_detail/26231.html) operation to query the region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// ecs.g6
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The type of the savings plan.
	//
	// example:
	//
	// universal
	SpnType *string `json:"SpnType,omitempty" xml:"SpnType,omitempty"`
}

func (s QuerySavingsPlansDiscountResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s QuerySavingsPlansDiscountResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *QuerySavingsPlansDiscountResponseBodyDataItems) GetCommodityName() *string {
	return s.CommodityName
}

func (s *QuerySavingsPlansDiscountResponseBodyDataItems) GetContractDiscountRate() *string {
	return s.ContractDiscountRate
}

func (s *QuerySavingsPlansDiscountResponseBodyDataItems) GetCycle() *string {
	return s.Cycle
}

func (s *QuerySavingsPlansDiscountResponseBodyDataItems) GetDiscountRate() *string {
	return s.DiscountRate
}

func (s *QuerySavingsPlansDiscountResponseBodyDataItems) GetModuleName() *string {
	return s.ModuleName
}

func (s *QuerySavingsPlansDiscountResponseBodyDataItems) GetPayMode() *string {
	return s.PayMode
}

func (s *QuerySavingsPlansDiscountResponseBodyDataItems) GetRegion() *string {
	return s.Region
}

func (s *QuerySavingsPlansDiscountResponseBodyDataItems) GetRegionCode() *string {
	return s.RegionCode
}

func (s *QuerySavingsPlansDiscountResponseBodyDataItems) GetSpec() *string {
	return s.Spec
}

func (s *QuerySavingsPlansDiscountResponseBodyDataItems) GetSpnType() *string {
	return s.SpnType
}

func (s *QuerySavingsPlansDiscountResponseBodyDataItems) SetCommodityName(v string) *QuerySavingsPlansDiscountResponseBodyDataItems {
	s.CommodityName = &v
	return s
}

func (s *QuerySavingsPlansDiscountResponseBodyDataItems) SetContractDiscountRate(v string) *QuerySavingsPlansDiscountResponseBodyDataItems {
	s.ContractDiscountRate = &v
	return s
}

func (s *QuerySavingsPlansDiscountResponseBodyDataItems) SetCycle(v string) *QuerySavingsPlansDiscountResponseBodyDataItems {
	s.Cycle = &v
	return s
}

func (s *QuerySavingsPlansDiscountResponseBodyDataItems) SetDiscountRate(v string) *QuerySavingsPlansDiscountResponseBodyDataItems {
	s.DiscountRate = &v
	return s
}

func (s *QuerySavingsPlansDiscountResponseBodyDataItems) SetModuleName(v string) *QuerySavingsPlansDiscountResponseBodyDataItems {
	s.ModuleName = &v
	return s
}

func (s *QuerySavingsPlansDiscountResponseBodyDataItems) SetPayMode(v string) *QuerySavingsPlansDiscountResponseBodyDataItems {
	s.PayMode = &v
	return s
}

func (s *QuerySavingsPlansDiscountResponseBodyDataItems) SetRegion(v string) *QuerySavingsPlansDiscountResponseBodyDataItems {
	s.Region = &v
	return s
}

func (s *QuerySavingsPlansDiscountResponseBodyDataItems) SetRegionCode(v string) *QuerySavingsPlansDiscountResponseBodyDataItems {
	s.RegionCode = &v
	return s
}

func (s *QuerySavingsPlansDiscountResponseBodyDataItems) SetSpec(v string) *QuerySavingsPlansDiscountResponseBodyDataItems {
	s.Spec = &v
	return s
}

func (s *QuerySavingsPlansDiscountResponseBodyDataItems) SetSpnType(v string) *QuerySavingsPlansDiscountResponseBodyDataItems {
	s.SpnType = &v
	return s
}

func (s *QuerySavingsPlansDiscountResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
