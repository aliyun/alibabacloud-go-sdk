// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsSaleControlStockResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeEnsSaleControlStockResponseBody
	GetRequestId() *string
	SetSaleControl(v []*DescribeEnsSaleControlStockResponseBodySaleControl) *DescribeEnsSaleControlStockResponseBody
	GetSaleControl() []*DescribeEnsSaleControlStockResponseBodySaleControl
}

type DescribeEnsSaleControlStockResponseBody struct {
	RequestId   *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SaleControl []*DescribeEnsSaleControlStockResponseBodySaleControl `json:"SaleControl,omitempty" xml:"SaleControl,omitempty" type:"Repeated"`
}

func (s DescribeEnsSaleControlStockResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsSaleControlStockResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnsSaleControlStockResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEnsSaleControlStockResponseBody) GetSaleControl() []*DescribeEnsSaleControlStockResponseBodySaleControl {
	return s.SaleControl
}

func (s *DescribeEnsSaleControlStockResponseBody) SetRequestId(v string) *DescribeEnsSaleControlStockResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEnsSaleControlStockResponseBody) SetSaleControl(v []*DescribeEnsSaleControlStockResponseBodySaleControl) *DescribeEnsSaleControlStockResponseBody {
	s.SaleControl = v
	return s
}

func (s *DescribeEnsSaleControlStockResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEnsSaleControlStockResponseBodySaleControl struct {
	CommodityCode    *string                                                               `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	OrderType        *string                                                               `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	SaleControlItems []*DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItems `json:"SaleControlItems,omitempty" xml:"SaleControlItems,omitempty" type:"Repeated"`
}

func (s DescribeEnsSaleControlStockResponseBodySaleControl) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsSaleControlStockResponseBodySaleControl) GoString() string {
	return s.String()
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControl) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControl) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControl) GetSaleControlItems() []*DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItems {
	return s.SaleControlItems
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControl) SetCommodityCode(v string) *DescribeEnsSaleControlStockResponseBodySaleControl {
	s.CommodityCode = &v
	return s
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControl) SetOrderType(v string) *DescribeEnsSaleControlStockResponseBodySaleControl {
	s.OrderType = &v
	return s
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControl) SetSaleControlItems(v []*DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItems) *DescribeEnsSaleControlStockResponseBodySaleControl {
	s.SaleControlItems = v
	return s
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControl) Validate() error {
	return dara.Validate(s)
}

type DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItems struct {
	ModuleCode      *string                                                                            `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	SaleControlItem *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItem `json:"SaleControlItem,omitempty" xml:"SaleControlItem,omitempty" type:"Struct"`
}

func (s DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItems) GoString() string {
	return s.String()
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItems) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItems) GetSaleControlItem() *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItem {
	return s.SaleControlItem
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItems) SetModuleCode(v string) *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItems {
	s.ModuleCode = &v
	return s
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItems) SetSaleControlItem(v *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItem) *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItems {
	s.SaleControlItem = v
	return s
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItems) Validate() error {
	return dara.Validate(s)
}

type DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItem struct {
	BasicSaleControl     *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControl       `json:"BasicSaleControl,omitempty" xml:"BasicSaleControl,omitempty" type:"Struct"`
	ConditionSaleControl []*DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl `json:"ConditionSaleControl,omitempty" xml:"ConditionSaleControl,omitempty" type:"Repeated"`
}

func (s DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItem) GoString() string {
	return s.String()
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItem) GetBasicSaleControl() *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControl {
	return s.BasicSaleControl
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItem) GetConditionSaleControl() []*DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl {
	return s.ConditionSaleControl
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItem) SetBasicSaleControl(v *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControl) *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItem {
	s.BasicSaleControl = v
	return s
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItem) SetConditionSaleControl(v []*DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl) *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItem {
	s.ConditionSaleControl = v
	return s
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItem) Validate() error {
	return dara.Validate(s)
}

type DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControl struct {
	ModuleValue *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControlModuleValue `json:"ModuleValue,omitempty" xml:"ModuleValue,omitempty" type:"Struct"`
	StockValue  *string                                                                                                       `json:"StockValue,omitempty" xml:"StockValue,omitempty"`
}

func (s DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControl) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControl) GoString() string {
	return s.String()
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControl) GetModuleValue() *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControlModuleValue {
	return s.ModuleValue
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControl) GetStockValue() *string {
	return s.StockValue
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControl) SetModuleValue(v *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControlModuleValue) *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControl {
	s.ModuleValue = v
	return s
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControl) SetStockValue(v string) *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControl {
	s.StockValue = &v
	return s
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControl) Validate() error {
	return dara.Validate(s)
}

type DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControlModuleValue struct {
	ModuleMaxValue *string `json:"ModuleMaxValue,omitempty" xml:"ModuleMaxValue,omitempty"`
	ModuleMinValue *string `json:"ModuleMinValue,omitempty" xml:"ModuleMinValue,omitempty"`
}

func (s DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControlModuleValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControlModuleValue) GoString() string {
	return s.String()
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControlModuleValue) GetModuleMaxValue() *string {
	return s.ModuleMaxValue
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControlModuleValue) GetModuleMinValue() *string {
	return s.ModuleMinValue
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControlModuleValue) SetModuleMaxValue(v string) *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControlModuleValue {
	s.ModuleMaxValue = &v
	return s
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControlModuleValue) SetModuleMinValue(v string) *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControlModuleValue {
	s.ModuleMinValue = &v
	return s
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemBasicSaleControlModuleValue) Validate() error {
	return dara.Validate(s)
}

type DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl struct {
	ConditionControl *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlConditionControl `json:"ConditionControl,omitempty" xml:"ConditionControl,omitempty" type:"Struct"`
	ModuleValue      *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlModuleValue      `json:"ModuleValue,omitempty" xml:"ModuleValue,omitempty" type:"Struct"`
	StockValue       *string                                                                                                                `json:"StockValue,omitempty" xml:"StockValue,omitempty"`
}

func (s DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl) GoString() string {
	return s.String()
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl) GetConditionControl() *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlConditionControl {
	return s.ConditionControl
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl) GetModuleValue() *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlModuleValue {
	return s.ModuleValue
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl) GetStockValue() *string {
	return s.StockValue
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl) SetConditionControl(v *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlConditionControl) *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl {
	s.ConditionControl = v
	return s
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl) SetModuleValue(v *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlModuleValue) *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl {
	s.ModuleValue = v
	return s
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl) SetStockValue(v string) *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl {
	s.StockValue = &v
	return s
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControl) Validate() error {
	return dara.Validate(s)
}

type DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlConditionControl struct {
	ConditionControlModuleCode  *string `json:"ConditionControlModuleCode,omitempty" xml:"ConditionControlModuleCode,omitempty"`
	ConditionControlModuleValue *string `json:"ConditionControlModuleValue,omitempty" xml:"ConditionControlModuleValue,omitempty"`
}

func (s DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlConditionControl) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlConditionControl) GoString() string {
	return s.String()
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlConditionControl) GetConditionControlModuleCode() *string {
	return s.ConditionControlModuleCode
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlConditionControl) GetConditionControlModuleValue() *string {
	return s.ConditionControlModuleValue
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlConditionControl) SetConditionControlModuleCode(v string) *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlConditionControl {
	s.ConditionControlModuleCode = &v
	return s
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlConditionControl) SetConditionControlModuleValue(v string) *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlConditionControl {
	s.ConditionControlModuleValue = &v
	return s
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlConditionControl) Validate() error {
	return dara.Validate(s)
}

type DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlModuleValue struct {
	ModuleMaxValue *string `json:"ModuleMaxValue,omitempty" xml:"ModuleMaxValue,omitempty"`
	ModuleMinValue *string `json:"ModuleMinValue,omitempty" xml:"ModuleMinValue,omitempty"`
}

func (s DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlModuleValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlModuleValue) GoString() string {
	return s.String()
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlModuleValue) GetModuleMaxValue() *string {
	return s.ModuleMaxValue
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlModuleValue) GetModuleMinValue() *string {
	return s.ModuleMinValue
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlModuleValue) SetModuleMaxValue(v string) *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlModuleValue {
	s.ModuleMaxValue = &v
	return s
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlModuleValue) SetModuleMinValue(v string) *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlModuleValue {
	s.ModuleMinValue = &v
	return s
}

func (s *DescribeEnsSaleControlStockResponseBodySaleControlSaleControlItemsSaleControlItemConditionSaleControlModuleValue) Validate() error {
	return dara.Validate(s)
}
