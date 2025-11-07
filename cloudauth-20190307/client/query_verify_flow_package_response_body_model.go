// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVerifyFlowPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryVerifyFlowPackageResponseBody
	GetCode() *string
	SetHttpStatusCode(v int64) *QueryVerifyFlowPackageResponseBody
	GetHttpStatusCode() *int64
	SetItems(v []*QueryVerifyFlowPackageResponseBodyItems) *QueryVerifyFlowPackageResponseBody
	GetItems() []*QueryVerifyFlowPackageResponseBodyItems
	SetRequestId(v string) *QueryVerifyFlowPackageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryVerifyFlowPackageResponseBody
	GetSuccess() *bool
	SetTotalCount(v string) *QueryVerifyFlowPackageResponseBody
	GetTotalCount() *string
}

type QueryVerifyFlowPackageResponseBody struct {
	// Return code
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// List of returned data.
	Items []*QueryVerifyFlowPackageResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// ID of the request
	//
	// example:
	//
	// 969434DF-926B-4997-9881-4DE94E39F805
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the response was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Total count.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryVerifyFlowPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryVerifyFlowPackageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryVerifyFlowPackageResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryVerifyFlowPackageResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *QueryVerifyFlowPackageResponseBody) GetItems() []*QueryVerifyFlowPackageResponseBodyItems {
	return s.Items
}

func (s *QueryVerifyFlowPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryVerifyFlowPackageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryVerifyFlowPackageResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *QueryVerifyFlowPackageResponseBody) SetCode(v string) *QueryVerifyFlowPackageResponseBody {
	s.Code = &v
	return s
}

func (s *QueryVerifyFlowPackageResponseBody) SetHttpStatusCode(v int64) *QueryVerifyFlowPackageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryVerifyFlowPackageResponseBody) SetItems(v []*QueryVerifyFlowPackageResponseBodyItems) *QueryVerifyFlowPackageResponseBody {
	s.Items = v
	return s
}

func (s *QueryVerifyFlowPackageResponseBody) SetRequestId(v string) *QueryVerifyFlowPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryVerifyFlowPackageResponseBody) SetSuccess(v bool) *QueryVerifyFlowPackageResponseBody {
	s.Success = &v
	return s
}

func (s *QueryVerifyFlowPackageResponseBody) SetTotalCount(v string) *QueryVerifyFlowPackageResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryVerifyFlowPackageResponseBody) Validate() error {
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

type QueryVerifyFlowPackageResponseBodyItems struct {
	// Name of the resource package.
	//
	// example:
	//
	// 实人认证流量包
	CommodityName *string `json:"CommodityName,omitempty" xml:"CommodityName,omitempty"`
	// Current available capacity.
	//
	// example:
	//
	// 0.0
	CurrCapacity *float64 `json:"CurrCapacity,omitempty" xml:"CurrCapacity,omitempty"`
	// Proportion of current remaining capacity to total capacity.
	//
	// example:
	//
	// 0%
	CurrProportion *string `json:"CurrProportion,omitempty" xml:"CurrProportion,omitempty"`
	// Details of the flow package.
	FlowDetails []*QueryVerifyFlowPackageResponseBodyItemsFlowDetails `json:"FlowDetails,omitempty" xml:"FlowDetails,omitempty" type:"Repeated"`
	// Total quota.
	//
	// example:
	//
	// 0.0
	TotalCapacity *float64 `json:"TotalCapacity,omitempty" xml:"TotalCapacity,omitempty"`
	// Used capacity.
	//
	// example:
	//
	// 0.0
	UsedCapacity *float64 `json:"UsedCapacity,omitempty" xml:"UsedCapacity,omitempty"`
}

func (s QueryVerifyFlowPackageResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s QueryVerifyFlowPackageResponseBodyItems) GoString() string {
	return s.String()
}

func (s *QueryVerifyFlowPackageResponseBodyItems) GetCommodityName() *string {
	return s.CommodityName
}

func (s *QueryVerifyFlowPackageResponseBodyItems) GetCurrCapacity() *float64 {
	return s.CurrCapacity
}

func (s *QueryVerifyFlowPackageResponseBodyItems) GetCurrProportion() *string {
	return s.CurrProportion
}

func (s *QueryVerifyFlowPackageResponseBodyItems) GetFlowDetails() []*QueryVerifyFlowPackageResponseBodyItemsFlowDetails {
	return s.FlowDetails
}

func (s *QueryVerifyFlowPackageResponseBodyItems) GetTotalCapacity() *float64 {
	return s.TotalCapacity
}

func (s *QueryVerifyFlowPackageResponseBodyItems) GetUsedCapacity() *float64 {
	return s.UsedCapacity
}

func (s *QueryVerifyFlowPackageResponseBodyItems) SetCommodityName(v string) *QueryVerifyFlowPackageResponseBodyItems {
	s.CommodityName = &v
	return s
}

func (s *QueryVerifyFlowPackageResponseBodyItems) SetCurrCapacity(v float64) *QueryVerifyFlowPackageResponseBodyItems {
	s.CurrCapacity = &v
	return s
}

func (s *QueryVerifyFlowPackageResponseBodyItems) SetCurrProportion(v string) *QueryVerifyFlowPackageResponseBodyItems {
	s.CurrProportion = &v
	return s
}

func (s *QueryVerifyFlowPackageResponseBodyItems) SetFlowDetails(v []*QueryVerifyFlowPackageResponseBodyItemsFlowDetails) *QueryVerifyFlowPackageResponseBodyItems {
	s.FlowDetails = v
	return s
}

func (s *QueryVerifyFlowPackageResponseBodyItems) SetTotalCapacity(v float64) *QueryVerifyFlowPackageResponseBodyItems {
	s.TotalCapacity = &v
	return s
}

func (s *QueryVerifyFlowPackageResponseBodyItems) SetUsedCapacity(v float64) *QueryVerifyFlowPackageResponseBodyItems {
	s.UsedCapacity = &v
	return s
}

func (s *QueryVerifyFlowPackageResponseBodyItems) Validate() error {
	if s.FlowDetails != nil {
		for _, item := range s.FlowDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryVerifyFlowPackageResponseBodyItemsFlowDetails struct {
	// Total amount.
	//
	// example:
	//
	// 0
	Capacity *float64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// Name of the flow package.
	//
	// example:
	//
	// 实人认证流量包
	CommodityName *string `json:"CommodityName,omitempty" xml:"CommodityName,omitempty"`
	// Remaining amount.
	//
	// example:
	//
	// 0.0
	CurrCapacity *float64 `json:"CurrCapacity,omitempty" xml:"CurrCapacity,omitempty"`
	// Proportion of remaining amount.
	//
	// example:
	//
	// 100%
	CurrProportion *string `json:"CurrProportion,omitempty" xml:"CurrProportion,omitempty"`
	// Expiration date.
	//
	// example:
	//
	// -
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// Instance name
	//
	// example:
	//
	// tf-testacccn-hangzhouapigate84369
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Status.
	//
	// example:
	//
	// -
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Effective date.
	//
	// example:
	//
	// -
	TakeEffectDate *string `json:"TakeEffectDate,omitempty" xml:"TakeEffectDate,omitempty"`
}

func (s QueryVerifyFlowPackageResponseBodyItemsFlowDetails) String() string {
	return dara.Prettify(s)
}

func (s QueryVerifyFlowPackageResponseBodyItemsFlowDetails) GoString() string {
	return s.String()
}

func (s *QueryVerifyFlowPackageResponseBodyItemsFlowDetails) GetCapacity() *float64 {
	return s.Capacity
}

func (s *QueryVerifyFlowPackageResponseBodyItemsFlowDetails) GetCommodityName() *string {
	return s.CommodityName
}

func (s *QueryVerifyFlowPackageResponseBodyItemsFlowDetails) GetCurrCapacity() *float64 {
	return s.CurrCapacity
}

func (s *QueryVerifyFlowPackageResponseBodyItemsFlowDetails) GetCurrProportion() *string {
	return s.CurrProportion
}

func (s *QueryVerifyFlowPackageResponseBodyItemsFlowDetails) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *QueryVerifyFlowPackageResponseBodyItemsFlowDetails) GetInstanceName() *string {
	return s.InstanceName
}

func (s *QueryVerifyFlowPackageResponseBodyItemsFlowDetails) GetStatus() *string {
	return s.Status
}

func (s *QueryVerifyFlowPackageResponseBodyItemsFlowDetails) GetTakeEffectDate() *string {
	return s.TakeEffectDate
}

func (s *QueryVerifyFlowPackageResponseBodyItemsFlowDetails) SetCapacity(v float64) *QueryVerifyFlowPackageResponseBodyItemsFlowDetails {
	s.Capacity = &v
	return s
}

func (s *QueryVerifyFlowPackageResponseBodyItemsFlowDetails) SetCommodityName(v string) *QueryVerifyFlowPackageResponseBodyItemsFlowDetails {
	s.CommodityName = &v
	return s
}

func (s *QueryVerifyFlowPackageResponseBodyItemsFlowDetails) SetCurrCapacity(v float64) *QueryVerifyFlowPackageResponseBodyItemsFlowDetails {
	s.CurrCapacity = &v
	return s
}

func (s *QueryVerifyFlowPackageResponseBodyItemsFlowDetails) SetCurrProportion(v string) *QueryVerifyFlowPackageResponseBodyItemsFlowDetails {
	s.CurrProportion = &v
	return s
}

func (s *QueryVerifyFlowPackageResponseBodyItemsFlowDetails) SetExpireDate(v string) *QueryVerifyFlowPackageResponseBodyItemsFlowDetails {
	s.ExpireDate = &v
	return s
}

func (s *QueryVerifyFlowPackageResponseBodyItemsFlowDetails) SetInstanceName(v string) *QueryVerifyFlowPackageResponseBodyItemsFlowDetails {
	s.InstanceName = &v
	return s
}

func (s *QueryVerifyFlowPackageResponseBodyItemsFlowDetails) SetStatus(v string) *QueryVerifyFlowPackageResponseBodyItemsFlowDetails {
	s.Status = &v
	return s
}

func (s *QueryVerifyFlowPackageResponseBodyItemsFlowDetails) SetTakeEffectDate(v string) *QueryVerifyFlowPackageResponseBodyItemsFlowDetails {
	s.TakeEffectDate = &v
	return s
}

func (s *QueryVerifyFlowPackageResponseBodyItemsFlowDetails) Validate() error {
	return dara.Validate(s)
}
