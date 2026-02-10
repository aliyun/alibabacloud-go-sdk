// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrderDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetOrderDetailResponseBody
	GetCode() *string
	SetData(v *GetOrderDetailResponseBodyData) *GetOrderDetailResponseBody
	GetData() *GetOrderDetailResponseBodyData
	SetMessage(v string) *GetOrderDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetOrderDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetOrderDetailResponseBody
	GetSuccess() *bool
}

type GetOrderDetailResponseBody struct {
	// example:
	//
	// Success
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetOrderDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D6E068C3-25BC-455A-85FE-45F0B22ECB1F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOrderDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOrderDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrderDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetOrderDetailResponseBody) GetData() *GetOrderDetailResponseBodyData {
	return s.Data
}

func (s *GetOrderDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetOrderDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOrderDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetOrderDetailResponseBody) SetCode(v string) *GetOrderDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetOrderDetailResponseBody) SetData(v *GetOrderDetailResponseBodyData) *GetOrderDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetOrderDetailResponseBody) SetMessage(v string) *GetOrderDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetOrderDetailResponseBody) SetRequestId(v string) *GetOrderDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOrderDetailResponseBody) SetSuccess(v bool) *GetOrderDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetOrderDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOrderDetailResponseBodyData struct {
	// example:
	//
	// test
	HostName  *string                                  `json:"HostName,omitempty" xml:"HostName,omitempty"`
	OrderList *GetOrderDetailResponseBodyDataOrderList `json:"OrderList,omitempty" xml:"OrderList,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 400
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetOrderDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetOrderDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOrderDetailResponseBodyData) GetHostName() *string {
	return s.HostName
}

func (s *GetOrderDetailResponseBodyData) GetOrderList() *GetOrderDetailResponseBodyDataOrderList {
	return s.OrderList
}

func (s *GetOrderDetailResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetOrderDetailResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetOrderDetailResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetOrderDetailResponseBodyData) SetHostName(v string) *GetOrderDetailResponseBodyData {
	s.HostName = &v
	return s
}

func (s *GetOrderDetailResponseBodyData) SetOrderList(v *GetOrderDetailResponseBodyDataOrderList) *GetOrderDetailResponseBodyData {
	s.OrderList = v
	return s
}

func (s *GetOrderDetailResponseBodyData) SetPageNum(v int32) *GetOrderDetailResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetOrderDetailResponseBodyData) SetPageSize(v int32) *GetOrderDetailResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetOrderDetailResponseBodyData) SetTotalCount(v int32) *GetOrderDetailResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetOrderDetailResponseBodyData) Validate() error {
	if s.OrderList != nil {
		if err := s.OrderList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOrderDetailResponseBodyDataOrderList struct {
	Order []*GetOrderDetailResponseBodyDataOrderListOrder `json:"Order,omitempty" xml:"Order,omitempty" type:"Repeated"`
}

func (s GetOrderDetailResponseBodyDataOrderList) String() string {
	return dara.Prettify(s)
}

func (s GetOrderDetailResponseBodyDataOrderList) GoString() string {
	return s.String()
}

func (s *GetOrderDetailResponseBodyDataOrderList) GetOrder() []*GetOrderDetailResponseBodyDataOrderListOrder {
	return s.Order
}

func (s *GetOrderDetailResponseBodyDataOrderList) SetOrder(v []*GetOrderDetailResponseBodyDataOrderListOrder) *GetOrderDetailResponseBodyDataOrderList {
	s.Order = v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderList) Validate() error {
	if s.Order != nil {
		for _, item := range s.Order {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetOrderDetailResponseBodyDataOrderListOrder struct {
	AfterTaxAmount       *string                                                           `json:"AfterTaxAmount,omitempty" xml:"AfterTaxAmount,omitempty"`
	BillModuleConfig     *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfig     `json:"BillModuleConfig,omitempty" xml:"BillModuleConfig,omitempty" type:"Struct"`
	CommodityCode        *string                                                           `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	Config               *string                                                           `json:"Config,omitempty" xml:"Config,omitempty"`
	CreateTime           *string                                                           `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Currency             *string                                                           `json:"Currency,omitempty" xml:"Currency,omitempty"`
	ExtendInfos          map[string]*string                                                `json:"ExtendInfos,omitempty" xml:"ExtendInfos,omitempty"`
	InstanceIds          *string                                                           `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	Operator             *string                                                           `json:"Operator,omitempty" xml:"Operator,omitempty"`
	OrderId              *string                                                           `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrderSubType         *string                                                           `json:"OrderSubType,omitempty" xml:"OrderSubType,omitempty"`
	OrderType            *string                                                           `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OriginalConfig       *string                                                           `json:"OriginalConfig,omitempty" xml:"OriginalConfig,omitempty"`
	OriginalModuleConfig *GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfig `json:"OriginalModuleConfig,omitempty" xml:"OriginalModuleConfig,omitempty" type:"Struct"`
	PaymentCurrency      *string                                                           `json:"PaymentCurrency,omitempty" xml:"PaymentCurrency,omitempty"`
	PaymentStatus        *string                                                           `json:"PaymentStatus,omitempty" xml:"PaymentStatus,omitempty"`
	PaymentTime          *string                                                           `json:"PaymentTime,omitempty" xml:"PaymentTime,omitempty"`
	PretaxAmount         *string                                                           `json:"PretaxAmount,omitempty" xml:"PretaxAmount,omitempty"`
	PretaxAmountLocal    *string                                                           `json:"PretaxAmountLocal,omitempty" xml:"PretaxAmountLocal,omitempty"`
	PretaxGrossAmount    *string                                                           `json:"PretaxGrossAmount,omitempty" xml:"PretaxGrossAmount,omitempty"`
	ProductCode          *string                                                           `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductType          *string                                                           `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	Quantity             *string                                                           `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	Region               *string                                                           `json:"Region,omitempty" xml:"Region,omitempty"`
	RelatedOrderId       *string                                                           `json:"RelatedOrderId,omitempty" xml:"RelatedOrderId,omitempty"`
	SubOrderId           *string                                                           `json:"SubOrderId,omitempty" xml:"SubOrderId,omitempty"`
	SubscriptionType     *string                                                           `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	Tax                  *string                                                           `json:"Tax,omitempty" xml:"Tax,omitempty"`
	UsageEndTime         *string                                                           `json:"UsageEndTime,omitempty" xml:"UsageEndTime,omitempty"`
	UsageStartTime       *string                                                           `json:"UsageStartTime,omitempty" xml:"UsageStartTime,omitempty"`
}

func (s GetOrderDetailResponseBodyDataOrderListOrder) String() string {
	return dara.Prettify(s)
}

func (s GetOrderDetailResponseBodyDataOrderListOrder) GoString() string {
	return s.String()
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) GetAfterTaxAmount() *string {
	return s.AfterTaxAmount
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) GetBillModuleConfig() *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfig {
	return s.BillModuleConfig
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) GetConfig() *string {
	return s.Config
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) GetCurrency() *string {
	return s.Currency
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) GetExtendInfos() map[string]*string {
	return s.ExtendInfos
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) GetOperator() *string {
	return s.Operator
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) GetOrderId() *string {
	return s.OrderId
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) GetOrderSubType() *string {
	return s.OrderSubType
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) GetOrderType() *string {
	return s.OrderType
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) GetOriginalConfig() *string {
	return s.OriginalConfig
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) GetOriginalModuleConfig() *GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfig {
	return s.OriginalModuleConfig
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) GetPaymentCurrency() *string {
	return s.PaymentCurrency
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) GetPaymentStatus() *string {
	return s.PaymentStatus
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) GetPaymentTime() *string {
	return s.PaymentTime
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) GetPretaxAmount() *string {
	return s.PretaxAmount
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) GetPretaxAmountLocal() *string {
	return s.PretaxAmountLocal
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) GetPretaxGrossAmount() *string {
	return s.PretaxGrossAmount
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) GetProductCode() *string {
	return s.ProductCode
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) GetProductType() *string {
	return s.ProductType
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) GetQuantity() *string {
	return s.Quantity
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) GetRegion() *string {
	return s.Region
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) GetRelatedOrderId() *string {
	return s.RelatedOrderId
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) GetSubOrderId() *string {
	return s.SubOrderId
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) GetTax() *string {
	return s.Tax
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) GetUsageEndTime() *string {
	return s.UsageEndTime
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) GetUsageStartTime() *string {
	return s.UsageStartTime
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetAfterTaxAmount(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.AfterTaxAmount = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetBillModuleConfig(v *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfig) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.BillModuleConfig = v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetCommodityCode(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.CommodityCode = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetConfig(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.Config = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetCreateTime(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.CreateTime = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetCurrency(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.Currency = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetExtendInfos(v map[string]*string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.ExtendInfos = v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetInstanceIds(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.InstanceIds = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetOperator(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.Operator = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetOrderId(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.OrderId = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetOrderSubType(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.OrderSubType = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetOrderType(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.OrderType = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetOriginalConfig(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.OriginalConfig = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetOriginalModuleConfig(v *GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfig) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.OriginalModuleConfig = v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetPaymentCurrency(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.PaymentCurrency = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetPaymentStatus(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.PaymentStatus = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetPaymentTime(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.PaymentTime = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetPretaxAmount(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.PretaxAmount = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetPretaxAmountLocal(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.PretaxAmountLocal = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetPretaxGrossAmount(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.PretaxGrossAmount = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetProductCode(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.ProductCode = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetProductType(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.ProductType = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetQuantity(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.Quantity = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetRegion(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.Region = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetRelatedOrderId(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.RelatedOrderId = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetSubOrderId(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.SubOrderId = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetSubscriptionType(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.SubscriptionType = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetTax(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.Tax = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetUsageEndTime(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.UsageEndTime = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) SetUsageStartTime(v string) *GetOrderDetailResponseBodyDataOrderListOrder {
	s.UsageStartTime = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrder) Validate() error {
	if s.BillModuleConfig != nil {
		if err := s.BillModuleConfig.Validate(); err != nil {
			return err
		}
	}
	if s.OriginalModuleConfig != nil {
		if err := s.OriginalModuleConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfig struct {
	BillModuleConfig []*GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfig `json:"billModuleConfig,omitempty" xml:"billModuleConfig,omitempty" type:"Repeated"`
}

func (s GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfig) String() string {
	return dara.Prettify(s)
}

func (s GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfig) GoString() string {
	return s.String()
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfig) GetBillModuleConfig() []*GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfig {
	return s.BillModuleConfig
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfig) SetBillModuleConfig(v []*GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfig) *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfig {
	s.BillModuleConfig = v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfig) Validate() error {
	if s.BillModuleConfig != nil {
		for _, item := range s.BillModuleConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfig struct {
	ApiCode              *string                                                                                           `json:"ApiCode,omitempty" xml:"ApiCode,omitempty"`
	BillModuleProperties *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfigBillModuleProperties `json:"BillModuleProperties,omitempty" xml:"BillModuleProperties,omitempty" type:"Struct"`
	Code                 *string                                                                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Name                 *string                                                                                           `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfig) String() string {
	return dara.Prettify(s)
}

func (s GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfig) GoString() string {
	return s.String()
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfig) GetApiCode() *string {
	return s.ApiCode
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfig) GetBillModuleProperties() *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfigBillModuleProperties {
	return s.BillModuleProperties
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfig) GetCode() *string {
	return s.Code
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfig) GetName() *string {
	return s.Name
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfig) SetApiCode(v string) *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfig {
	s.ApiCode = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfig) SetBillModuleProperties(v *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfigBillModuleProperties) *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfig {
	s.BillModuleProperties = v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfig) SetCode(v string) *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfig {
	s.Code = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfig) SetName(v string) *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfig {
	s.Name = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfig) Validate() error {
	if s.BillModuleProperties != nil {
		if err := s.BillModuleProperties.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfigBillModuleProperties struct {
	BillModuleProperties []*GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfigBillModulePropertiesBillModuleProperties `json:"billModuleProperties,omitempty" xml:"billModuleProperties,omitempty" type:"Repeated"`
}

func (s GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfigBillModuleProperties) String() string {
	return dara.Prettify(s)
}

func (s GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfigBillModuleProperties) GoString() string {
	return s.String()
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfigBillModuleProperties) GetBillModuleProperties() []*GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfigBillModulePropertiesBillModuleProperties {
	return s.BillModuleProperties
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfigBillModuleProperties) SetBillModuleProperties(v []*GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfigBillModulePropertiesBillModuleProperties) *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfigBillModuleProperties {
	s.BillModuleProperties = v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfigBillModuleProperties) Validate() error {
	if s.BillModuleProperties != nil {
		for _, item := range s.BillModuleProperties {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfigBillModulePropertiesBillModuleProperties struct {
	AttrApiCode   *string `json:"AttrApiCode,omitempty" xml:"AttrApiCode,omitempty"`
	ModuleApiCode *string `json:"ModuleApiCode,omitempty" xml:"ModuleApiCode,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfigBillModulePropertiesBillModuleProperties) String() string {
	return dara.Prettify(s)
}

func (s GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfigBillModulePropertiesBillModuleProperties) GoString() string {
	return s.String()
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfigBillModulePropertiesBillModuleProperties) GetAttrApiCode() *string {
	return s.AttrApiCode
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfigBillModulePropertiesBillModuleProperties) GetModuleApiCode() *string {
	return s.ModuleApiCode
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfigBillModulePropertiesBillModuleProperties) GetValue() *string {
	return s.Value
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfigBillModulePropertiesBillModuleProperties) SetAttrApiCode(v string) *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfigBillModulePropertiesBillModuleProperties {
	s.AttrApiCode = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfigBillModulePropertiesBillModuleProperties) SetModuleApiCode(v string) *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfigBillModulePropertiesBillModuleProperties {
	s.ModuleApiCode = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfigBillModulePropertiesBillModuleProperties) SetValue(v string) *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfigBillModulePropertiesBillModuleProperties {
	s.Value = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderBillModuleConfigBillModuleConfigBillModulePropertiesBillModuleProperties) Validate() error {
	return dara.Validate(s)
}

type GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfig struct {
	OriginalModuleConfig []*GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfig `json:"originalModuleConfig,omitempty" xml:"originalModuleConfig,omitempty" type:"Repeated"`
}

func (s GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfig) String() string {
	return dara.Prettify(s)
}

func (s GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfig) GoString() string {
	return s.String()
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfig) GetOriginalModuleConfig() []*GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfig {
	return s.OriginalModuleConfig
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfig) SetOriginalModuleConfig(v []*GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfig) *GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfig {
	s.OriginalModuleConfig = v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfig) Validate() error {
	if s.OriginalModuleConfig != nil {
		for _, item := range s.OriginalModuleConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfig struct {
	Code             *string                                                                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	ModuleProperties *GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfigModuleProperties `json:"ModuleProperties,omitempty" xml:"ModuleProperties,omitempty" type:"Struct"`
	Name             *string                                                                                               `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfig) String() string {
	return dara.Prettify(s)
}

func (s GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfig) GoString() string {
	return s.String()
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfig) GetCode() *string {
	return s.Code
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfig) GetModuleProperties() *GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfigModuleProperties {
	return s.ModuleProperties
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfig) GetName() *string {
	return s.Name
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfig) SetCode(v string) *GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfig {
	s.Code = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfig) SetModuleProperties(v *GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfigModuleProperties) *GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfig {
	s.ModuleProperties = v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfig) SetName(v string) *GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfig {
	s.Name = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfig) Validate() error {
	if s.ModuleProperties != nil {
		if err := s.ModuleProperties.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfigModuleProperties struct {
	ModuleProperties []*GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfigModulePropertiesModuleProperties `json:"moduleProperties,omitempty" xml:"moduleProperties,omitempty" type:"Repeated"`
}

func (s GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfigModuleProperties) String() string {
	return dara.Prettify(s)
}

func (s GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfigModuleProperties) GoString() string {
	return s.String()
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfigModuleProperties) GetModuleProperties() []*GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfigModulePropertiesModuleProperties {
	return s.ModuleProperties
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfigModuleProperties) SetModuleProperties(v []*GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfigModulePropertiesModuleProperties) *GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfigModuleProperties {
	s.ModuleProperties = v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfigModuleProperties) Validate() error {
	if s.ModuleProperties != nil {
		for _, item := range s.ModuleProperties {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfigModulePropertiesModuleProperties struct {
	Code  *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfigModulePropertiesModuleProperties) String() string {
	return dara.Prettify(s)
}

func (s GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfigModulePropertiesModuleProperties) GoString() string {
	return s.String()
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfigModulePropertiesModuleProperties) GetCode() *string {
	return s.Code
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfigModulePropertiesModuleProperties) GetName() *string {
	return s.Name
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfigModulePropertiesModuleProperties) GetValue() *string {
	return s.Value
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfigModulePropertiesModuleProperties) SetCode(v string) *GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfigModulePropertiesModuleProperties {
	s.Code = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfigModulePropertiesModuleProperties) SetName(v string) *GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfigModulePropertiesModuleProperties {
	s.Name = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfigModulePropertiesModuleProperties) SetValue(v string) *GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfigModulePropertiesModuleProperties {
	s.Value = &v
	return s
}

func (s *GetOrderDetailResponseBodyDataOrderListOrderOriginalModuleConfigOriginalModuleConfigModulePropertiesModuleProperties) Validate() error {
	return dara.Validate(s)
}
