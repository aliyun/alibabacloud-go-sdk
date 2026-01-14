// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommodityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommodityCode(v string) *DescribeCommodityResponseBody
	GetCommodityCode() *string
	SetCommodityName(v string) *DescribeCommodityResponseBody
	GetCommodityName() *string
	SetComponents(v []*DescribeCommodityResponseBodyComponents) *DescribeCommodityResponseBody
	GetComponents() []*DescribeCommodityResponseBodyComponents
	SetRequestId(v string) *DescribeCommodityResponseBody
	GetRequestId() *string
}

type DescribeCommodityResponseBody struct {
	// The commodity code.
	//
	// Examples for the China site (aliyun.com):
	//
	// 	- **ga_gapluspre_public_cn**: GA instance.
	//
	// 	- **ga_plusbwppre_public_cn**: basic bandwidth plan.
	//
	// Examples for the international site (alibabacloud.com):
	//
	// 	- **ga_pluspre_public_intl**: GA instance.
	//
	// 	- **ga_bwppreintl_public_intl**: basic bandwidth plan.
	//
	// example:
	//
	// ga_gapluspre_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The name of the commodity.
	//
	// example:
	//
	// Global Accelerator_Instance Type (Subscription)
	CommodityName *string `json:"CommodityName,omitempty" xml:"CommodityName,omitempty"`
	// The information about the commodity modules.
	//
	// The returned information varies based on the commodity.
	Components []*DescribeCommodityResponseBodyComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCommodityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommodityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCommodityResponseBody) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeCommodityResponseBody) GetCommodityName() *string {
	return s.CommodityName
}

func (s *DescribeCommodityResponseBody) GetComponents() []*DescribeCommodityResponseBodyComponents {
	return s.Components
}

func (s *DescribeCommodityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCommodityResponseBody) SetCommodityCode(v string) *DescribeCommodityResponseBody {
	s.CommodityCode = &v
	return s
}

func (s *DescribeCommodityResponseBody) SetCommodityName(v string) *DescribeCommodityResponseBody {
	s.CommodityName = &v
	return s
}

func (s *DescribeCommodityResponseBody) SetComponents(v []*DescribeCommodityResponseBodyComponents) *DescribeCommodityResponseBody {
	s.Components = v
	return s
}

func (s *DescribeCommodityResponseBody) SetRequestId(v string) *DescribeCommodityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCommodityResponseBody) Validate() error {
	if s.Components != nil {
		for _, item := range s.Components {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCommodityResponseBodyComponents struct {
	// The code of the commodity module.
	//
	// The returned information varies based on the commodity module.
	//
	// example:
	//
	// Duration
	ComponentCode *string `json:"ComponentCode,omitempty" xml:"ComponentCode,omitempty"`
	// The name of the commodity module.
	//
	// The returned information varies based on the commodity module.
	//
	// example:
	//
	// Duration
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// The attributes of the commodity module.
	//
	// The returned information varies based on the commodity module.
	Properties []*DescribeCommodityResponseBodyComponentsProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Repeated"`
}

func (s DescribeCommodityResponseBodyComponents) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommodityResponseBodyComponents) GoString() string {
	return s.String()
}

func (s *DescribeCommodityResponseBodyComponents) GetComponentCode() *string {
	return s.ComponentCode
}

func (s *DescribeCommodityResponseBodyComponents) GetComponentName() *string {
	return s.ComponentName
}

func (s *DescribeCommodityResponseBodyComponents) GetProperties() []*DescribeCommodityResponseBodyComponentsProperties {
	return s.Properties
}

func (s *DescribeCommodityResponseBodyComponents) SetComponentCode(v string) *DescribeCommodityResponseBodyComponents {
	s.ComponentCode = &v
	return s
}

func (s *DescribeCommodityResponseBodyComponents) SetComponentName(v string) *DescribeCommodityResponseBodyComponents {
	s.ComponentName = &v
	return s
}

func (s *DescribeCommodityResponseBodyComponents) SetProperties(v []*DescribeCommodityResponseBodyComponentsProperties) *DescribeCommodityResponseBodyComponents {
	s.Properties = v
	return s
}

func (s *DescribeCommodityResponseBodyComponents) Validate() error {
	if s.Properties != nil {
		for _, item := range s.Properties {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCommodityResponseBodyComponentsProperties struct {
	// The code of the attribute.
	//
	// The returned information varies based on the commodity module.
	//
	// example:
	//
	// ord_time
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The name of the attribute.
	//
	// The returned information varies based on the commodity module.
	//
	// example:
	//
	// Duration
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of attribute values of the commodity module.
	//
	// The returned information varies based on the commodity module.
	PropertyValueList []*DescribeCommodityResponseBodyComponentsPropertiesPropertyValueList `json:"PropertyValueList,omitempty" xml:"PropertyValueList,omitempty" type:"Repeated"`
}

func (s DescribeCommodityResponseBodyComponentsProperties) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommodityResponseBodyComponentsProperties) GoString() string {
	return s.String()
}

func (s *DescribeCommodityResponseBodyComponentsProperties) GetCode() *string {
	return s.Code
}

func (s *DescribeCommodityResponseBodyComponentsProperties) GetName() *string {
	return s.Name
}

func (s *DescribeCommodityResponseBodyComponentsProperties) GetPropertyValueList() []*DescribeCommodityResponseBodyComponentsPropertiesPropertyValueList {
	return s.PropertyValueList
}

func (s *DescribeCommodityResponseBodyComponentsProperties) SetCode(v string) *DescribeCommodityResponseBodyComponentsProperties {
	s.Code = &v
	return s
}

func (s *DescribeCommodityResponseBodyComponentsProperties) SetName(v string) *DescribeCommodityResponseBodyComponentsProperties {
	s.Name = &v
	return s
}

func (s *DescribeCommodityResponseBodyComponentsProperties) SetPropertyValueList(v []*DescribeCommodityResponseBodyComponentsPropertiesPropertyValueList) *DescribeCommodityResponseBodyComponentsProperties {
	s.PropertyValueList = v
	return s
}

func (s *DescribeCommodityResponseBodyComponentsProperties) Validate() error {
	if s.PropertyValueList != nil {
		for _, item := range s.PropertyValueList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCommodityResponseBodyComponentsPropertiesPropertyValueList struct {
	// The sequence number of the attribute.
	//
	// The returned information varies based on the commodity module.
	//
	// example:
	//
	// 1
	OrderIndex *int64 `json:"OrderIndex,omitempty" xml:"OrderIndex,omitempty"`
	// The content of the attribute.
	//
	// The returned information varies based on the commodity module.
	//
	// example:
	//
	// 1 Month
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The message of the attribute.
	//
	// The returned information varies based on the commodity module.
	//
	// example:
	//
	// 1 Month
	Tips *string `json:"Tips,omitempty" xml:"Tips,omitempty"`
	// The value of the attribute.
	//
	// The returned information varies based on the commodity module.
	//
	// example:
	//
	// 1:Month
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCommodityResponseBodyComponentsPropertiesPropertyValueList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommodityResponseBodyComponentsPropertiesPropertyValueList) GoString() string {
	return s.String()
}

func (s *DescribeCommodityResponseBodyComponentsPropertiesPropertyValueList) GetOrderIndex() *int64 {
	return s.OrderIndex
}

func (s *DescribeCommodityResponseBodyComponentsPropertiesPropertyValueList) GetText() *string {
	return s.Text
}

func (s *DescribeCommodityResponseBodyComponentsPropertiesPropertyValueList) GetTips() *string {
	return s.Tips
}

func (s *DescribeCommodityResponseBodyComponentsPropertiesPropertyValueList) GetValue() *string {
	return s.Value
}

func (s *DescribeCommodityResponseBodyComponentsPropertiesPropertyValueList) SetOrderIndex(v int64) *DescribeCommodityResponseBodyComponentsPropertiesPropertyValueList {
	s.OrderIndex = &v
	return s
}

func (s *DescribeCommodityResponseBodyComponentsPropertiesPropertyValueList) SetText(v string) *DescribeCommodityResponseBodyComponentsPropertiesPropertyValueList {
	s.Text = &v
	return s
}

func (s *DescribeCommodityResponseBodyComponentsPropertiesPropertyValueList) SetTips(v string) *DescribeCommodityResponseBodyComponentsPropertiesPropertyValueList {
	s.Tips = &v
	return s
}

func (s *DescribeCommodityResponseBodyComponentsPropertiesPropertyValueList) SetValue(v string) *DescribeCommodityResponseBodyComponentsPropertiesPropertyValueList {
	s.Value = &v
	return s
}

func (s *DescribeCommodityResponseBodyComponentsPropertiesPropertyValueList) Validate() error {
	return dara.Validate(s)
}
