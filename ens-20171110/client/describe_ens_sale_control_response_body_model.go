// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsSaleControlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeEnsSaleControlResponseBody
	GetRequestId() *string
	SetSaleControl(v []*DescribeEnsSaleControlResponseBodySaleControl) *DescribeEnsSaleControlResponseBody
	GetSaleControl() []*DescribeEnsSaleControlResponseBodySaleControl
}

type DescribeEnsSaleControlResponseBody struct {
	RequestId   *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SaleControl []*DescribeEnsSaleControlResponseBodySaleControl `json:"SaleControl,omitempty" xml:"SaleControl,omitempty" type:"Repeated"`
}

func (s DescribeEnsSaleControlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsSaleControlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnsSaleControlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEnsSaleControlResponseBody) GetSaleControl() []*DescribeEnsSaleControlResponseBodySaleControl {
	return s.SaleControl
}

func (s *DescribeEnsSaleControlResponseBody) SetRequestId(v string) *DescribeEnsSaleControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEnsSaleControlResponseBody) SetSaleControl(v []*DescribeEnsSaleControlResponseBodySaleControl) *DescribeEnsSaleControlResponseBody {
	s.SaleControl = v
	return s
}

func (s *DescribeEnsSaleControlResponseBody) Validate() error {
	if s.SaleControl != nil {
		for _, item := range s.SaleControl {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEnsSaleControlResponseBodySaleControl struct {
	CommodityCode    *string                                                          `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	OrderType        *string                                                          `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	SaleControlItems []*DescribeEnsSaleControlResponseBodySaleControlSaleControlItems `json:"SaleControlItems,omitempty" xml:"SaleControlItems,omitempty" type:"Repeated"`
}

func (s DescribeEnsSaleControlResponseBodySaleControl) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsSaleControlResponseBodySaleControl) GoString() string {
	return s.String()
}

func (s *DescribeEnsSaleControlResponseBodySaleControl) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeEnsSaleControlResponseBodySaleControl) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribeEnsSaleControlResponseBodySaleControl) GetSaleControlItems() []*DescribeEnsSaleControlResponseBodySaleControlSaleControlItems {
	return s.SaleControlItems
}

func (s *DescribeEnsSaleControlResponseBodySaleControl) SetCommodityCode(v string) *DescribeEnsSaleControlResponseBodySaleControl {
	s.CommodityCode = &v
	return s
}

func (s *DescribeEnsSaleControlResponseBodySaleControl) SetOrderType(v string) *DescribeEnsSaleControlResponseBodySaleControl {
	s.OrderType = &v
	return s
}

func (s *DescribeEnsSaleControlResponseBodySaleControl) SetSaleControlItems(v []*DescribeEnsSaleControlResponseBodySaleControlSaleControlItems) *DescribeEnsSaleControlResponseBodySaleControl {
	s.SaleControlItems = v
	return s
}

func (s *DescribeEnsSaleControlResponseBodySaleControl) Validate() error {
	if s.SaleControlItems != nil {
		for _, item := range s.SaleControlItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEnsSaleControlResponseBodySaleControlSaleControlItems struct {
	ModuleCode      *string                                                                       `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	SaleControlItem *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItem `json:"SaleControlItem,omitempty" xml:"SaleControlItem,omitempty" type:"Struct"`
}

func (s DescribeEnsSaleControlResponseBodySaleControlSaleControlItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsSaleControlResponseBodySaleControlSaleControlItems) GoString() string {
	return s.String()
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItems) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItems) GetSaleControlItem() *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItem {
	return s.SaleControlItem
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItems) SetModuleCode(v string) *DescribeEnsSaleControlResponseBodySaleControlSaleControlItems {
	s.ModuleCode = &v
	return s
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItems) SetSaleControlItem(v *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItem) *DescribeEnsSaleControlResponseBodySaleControlSaleControlItems {
	s.SaleControlItem = v
	return s
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItems) Validate() error {
	if s.SaleControlItem != nil {
		if err := s.SaleControlItem.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItem struct {
	BasicSaleControl     *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControl       `json:"BasicSaleControl,omitempty" xml:"BasicSaleControl,omitempty" type:"Struct"`
	ConditionSaleControl []*DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl `json:"ConditionSaleControl,omitempty" xml:"ConditionSaleControl,omitempty" type:"Repeated"`
}

func (s DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItem) GoString() string {
	return s.String()
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItem) GetBasicSaleControl() *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControl {
	return s.BasicSaleControl
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItem) GetConditionSaleControl() []*DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl {
	return s.ConditionSaleControl
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItem) SetBasicSaleControl(v *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControl) *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItem {
	s.BasicSaleControl = v
	return s
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItem) SetConditionSaleControl(v []*DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl) *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItem {
	s.ConditionSaleControl = v
	return s
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItem) Validate() error {
	if s.BasicSaleControl != nil {
		if err := s.BasicSaleControl.Validate(); err != nil {
			return err
		}
	}
	if s.ConditionSaleControl != nil {
		for _, item := range s.ConditionSaleControl {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControl struct {
	Description *string                                                                                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	ModuleValue *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControlModuleValue `json:"ModuleValue,omitempty" xml:"ModuleValue,omitempty" type:"Struct"`
	Operator    *string                                                                                                  `json:"Operator,omitempty" xml:"Operator,omitempty"`
}

func (s DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControl) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControl) GoString() string {
	return s.String()
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControl) GetDescription() *string {
	return s.Description
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControl) GetModuleValue() *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControlModuleValue {
	return s.ModuleValue
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControl) GetOperator() *string {
	return s.Operator
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControl) SetDescription(v string) *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControl {
	s.Description = &v
	return s
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControl) SetModuleValue(v *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControlModuleValue) *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControl {
	s.ModuleValue = v
	return s
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControl) SetOperator(v string) *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControl {
	s.Operator = &v
	return s
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControl) Validate() error {
	if s.ModuleValue != nil {
		if err := s.ModuleValue.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControlModuleValue struct {
	ModuleMaxValue *string   `json:"ModuleMaxValue,omitempty" xml:"ModuleMaxValue,omitempty"`
	ModuleMinValue *string   `json:"ModuleMinValue,omitempty" xml:"ModuleMinValue,omitempty"`
	ModuleValue    []*string `json:"ModuleValue,omitempty" xml:"ModuleValue,omitempty" type:"Repeated"`
}

func (s DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControlModuleValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControlModuleValue) GoString() string {
	return s.String()
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControlModuleValue) GetModuleMaxValue() *string {
	return s.ModuleMaxValue
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControlModuleValue) GetModuleMinValue() *string {
	return s.ModuleMinValue
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControlModuleValue) GetModuleValue() []*string {
	return s.ModuleValue
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControlModuleValue) SetModuleMaxValue(v string) *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControlModuleValue {
	s.ModuleMaxValue = &v
	return s
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControlModuleValue) SetModuleMinValue(v string) *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControlModuleValue {
	s.ModuleMinValue = &v
	return s
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControlModuleValue) SetModuleValue(v []*string) *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControlModuleValue {
	s.ModuleValue = v
	return s
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControlModuleValue) Validate() error {
	return dara.Validate(s)
}

type DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl struct {
	ConditionControl *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlConditionControl `json:"ConditionControl,omitempty" xml:"ConditionControl,omitempty" type:"Struct"`
	Description      *string                                                                                                           `json:"Description,omitempty" xml:"Description,omitempty"`
	ModuleValue      *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlModuleValue      `json:"ModuleValue,omitempty" xml:"ModuleValue,omitempty" type:"Struct"`
	Operator         *string                                                                                                           `json:"Operator,omitempty" xml:"Operator,omitempty"`
}

func (s DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl) GoString() string {
	return s.String()
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl) GetConditionControl() *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlConditionControl {
	return s.ConditionControl
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl) GetDescription() *string {
	return s.Description
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl) GetModuleValue() *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlModuleValue {
	return s.ModuleValue
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl) GetOperator() *string {
	return s.Operator
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl) SetConditionControl(v *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlConditionControl) *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl {
	s.ConditionControl = v
	return s
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl) SetDescription(v string) *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl {
	s.Description = &v
	return s
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl) SetModuleValue(v *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlModuleValue) *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl {
	s.ModuleValue = v
	return s
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl) SetOperator(v string) *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl {
	s.Operator = &v
	return s
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl) Validate() error {
	if s.ConditionControl != nil {
		if err := s.ConditionControl.Validate(); err != nil {
			return err
		}
	}
	if s.ModuleValue != nil {
		if err := s.ModuleValue.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlConditionControl struct {
	ConditionControlModuleCode  *string `json:"ConditionControlModuleCode,omitempty" xml:"ConditionControlModuleCode,omitempty"`
	ConditionControlModuleValue *string `json:"ConditionControlModuleValue,omitempty" xml:"ConditionControlModuleValue,omitempty"`
}

func (s DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlConditionControl) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlConditionControl) GoString() string {
	return s.String()
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlConditionControl) GetConditionControlModuleCode() *string {
	return s.ConditionControlModuleCode
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlConditionControl) GetConditionControlModuleValue() *string {
	return s.ConditionControlModuleValue
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlConditionControl) SetConditionControlModuleCode(v string) *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlConditionControl {
	s.ConditionControlModuleCode = &v
	return s
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlConditionControl) SetConditionControlModuleValue(v string) *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlConditionControl {
	s.ConditionControlModuleValue = &v
	return s
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlConditionControl) Validate() error {
	return dara.Validate(s)
}

type DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlModuleValue struct {
	ModuleMaxValue *string   `json:"ModuleMaxValue,omitempty" xml:"ModuleMaxValue,omitempty"`
	ModuleMinValue *string   `json:"ModuleMinValue,omitempty" xml:"ModuleMinValue,omitempty"`
	ModuleValue    []*string `json:"ModuleValue,omitempty" xml:"ModuleValue,omitempty" type:"Repeated"`
}

func (s DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlModuleValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlModuleValue) GoString() string {
	return s.String()
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlModuleValue) GetModuleMaxValue() *string {
	return s.ModuleMaxValue
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlModuleValue) GetModuleMinValue() *string {
	return s.ModuleMinValue
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlModuleValue) GetModuleValue() []*string {
	return s.ModuleValue
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlModuleValue) SetModuleMaxValue(v string) *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlModuleValue {
	s.ModuleMaxValue = &v
	return s
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlModuleValue) SetModuleMinValue(v string) *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlModuleValue {
	s.ModuleMinValue = &v
	return s
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlModuleValue) SetModuleValue(v []*string) *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlModuleValue {
	s.ModuleValue = v
	return s
}

func (s *DescribeEnsSaleControlResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlModuleValue) Validate() error {
	return dara.Validate(s)
}
