// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySplitItemBillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QuerySplitItemBillResponseBody
	GetCode() *string
	SetData(v *QuerySplitItemBillResponseBodyData) *QuerySplitItemBillResponseBody
	GetData() *QuerySplitItemBillResponseBodyData
	SetMessage(v string) *QuerySplitItemBillResponseBody
	GetMessage() *string
	SetRequestId(v string) *QuerySplitItemBillResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QuerySplitItemBillResponseBody
	GetSuccess() *bool
}

type QuerySplitItemBillResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *QuerySplitItemBillResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 79EE7556-0CFD-44EB-9CD6-B3B526E3A85F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySplitItemBillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySplitItemBillResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySplitItemBillResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySplitItemBillResponseBody) GetData() *QuerySplitItemBillResponseBodyData {
	return s.Data
}

func (s *QuerySplitItemBillResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySplitItemBillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySplitItemBillResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QuerySplitItemBillResponseBody) SetCode(v string) *QuerySplitItemBillResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySplitItemBillResponseBody) SetData(v *QuerySplitItemBillResponseBodyData) *QuerySplitItemBillResponseBody {
	s.Data = v
	return s
}

func (s *QuerySplitItemBillResponseBody) SetMessage(v string) *QuerySplitItemBillResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySplitItemBillResponseBody) SetRequestId(v string) *QuerySplitItemBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySplitItemBillResponseBody) SetSuccess(v bool) *QuerySplitItemBillResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySplitItemBillResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QuerySplitItemBillResponseBodyData struct {
	// The account ID.
	//
	// example:
	//
	// 185xxxx3489
	AccountID *string `json:"AccountID,omitempty" xml:"AccountID,omitempty"`
	// The user account.
	//
	// example:
	//
	// test@test.aliyunid.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The billing cycle. Format: YYYY-MM.
	//
	// example:
	//
	// 2020-03
	BillingCycle *string                                  `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	Items        *QuerySplitItemBillResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: 20. Maximum value: 300.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QuerySplitItemBillResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QuerySplitItemBillResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySplitItemBillResponseBodyData) GetAccountID() *string {
	return s.AccountID
}

func (s *QuerySplitItemBillResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *QuerySplitItemBillResponseBodyData) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *QuerySplitItemBillResponseBodyData) GetItems() *QuerySplitItemBillResponseBodyDataItems {
	return s.Items
}

func (s *QuerySplitItemBillResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QuerySplitItemBillResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QuerySplitItemBillResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QuerySplitItemBillResponseBodyData) SetAccountID(v string) *QuerySplitItemBillResponseBodyData {
	s.AccountID = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyData) SetAccountName(v string) *QuerySplitItemBillResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyData) SetBillingCycle(v string) *QuerySplitItemBillResponseBodyData {
	s.BillingCycle = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyData) SetItems(v *QuerySplitItemBillResponseBodyDataItems) *QuerySplitItemBillResponseBodyData {
	s.Items = v
	return s
}

func (s *QuerySplitItemBillResponseBodyData) SetPageNum(v int32) *QuerySplitItemBillResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyData) SetPageSize(v int32) *QuerySplitItemBillResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyData) SetTotalCount(v int32) *QuerySplitItemBillResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyData) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QuerySplitItemBillResponseBodyDataItems struct {
	Item []*QuerySplitItemBillResponseBodyDataItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s QuerySplitItemBillResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s QuerySplitItemBillResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *QuerySplitItemBillResponseBodyDataItems) GetItem() []*QuerySplitItemBillResponseBodyDataItemsItem {
	return s.Item
}

func (s *QuerySplitItemBillResponseBodyDataItems) SetItem(v []*QuerySplitItemBillResponseBodyDataItemsItem) *QuerySplitItemBillResponseBodyDataItems {
	s.Item = v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItems) Validate() error {
	if s.Item != nil {
		for _, item := range s.Item {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QuerySplitItemBillResponseBodyDataItemsItem struct {
	AdjustAmount              *float32 `json:"AdjustAmount,omitempty" xml:"AdjustAmount,omitempty"`
	BillingDate               *string  `json:"BillingDate,omitempty" xml:"BillingDate,omitempty"`
	BillingItem               *string  `json:"BillingItem,omitempty" xml:"BillingItem,omitempty"`
	BillingType               *string  `json:"BillingType,omitempty" xml:"BillingType,omitempty"`
	CashAmount                *float32 `json:"CashAmount,omitempty" xml:"CashAmount,omitempty"`
	CommodityCode             *string  `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CostUnit                  *string  `json:"CostUnit,omitempty" xml:"CostUnit,omitempty"`
	Currency                  *string  `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DeductedByCashCoupons     *float32 `json:"DeductedByCashCoupons,omitempty" xml:"DeductedByCashCoupons,omitempty"`
	DeductedByCoupons         *float32 `json:"DeductedByCoupons,omitempty" xml:"DeductedByCoupons,omitempty"`
	DeductedByPrepaidCard     *float32 `json:"DeductedByPrepaidCard,omitempty" xml:"DeductedByPrepaidCard,omitempty"`
	DeductedByResourcePackage *string  `json:"DeductedByResourcePackage,omitempty" xml:"DeductedByResourcePackage,omitempty"`
	InstanceConfig            *string  `json:"InstanceConfig,omitempty" xml:"InstanceConfig,omitempty"`
	InstanceID                *string  `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	InstanceSpec              *string  `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	InternetIP                *string  `json:"InternetIP,omitempty" xml:"InternetIP,omitempty"`
	IntranetIP                *string  `json:"IntranetIP,omitempty" xml:"IntranetIP,omitempty"`
	InvoiceDiscount           *float32 `json:"InvoiceDiscount,omitempty" xml:"InvoiceDiscount,omitempty"`
	Item                      *string  `json:"Item,omitempty" xml:"Item,omitempty"`
	ListPrice                 *string  `json:"ListPrice,omitempty" xml:"ListPrice,omitempty"`
	ListPriceUnit             *string  `json:"ListPriceUnit,omitempty" xml:"ListPriceUnit,omitempty"`
	NickName                  *string  `json:"NickName,omitempty" xml:"NickName,omitempty"`
	OutstandingAmount         *float32 `json:"OutstandingAmount,omitempty" xml:"OutstandingAmount,omitempty"`
	OwnerID                   *string  `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	PaymentAmount             *float32 `json:"PaymentAmount,omitempty" xml:"PaymentAmount,omitempty"`
	PipCode                   *string  `json:"PipCode,omitempty" xml:"PipCode,omitempty"`
	PretaxAmount              *float32 `json:"PretaxAmount,omitempty" xml:"PretaxAmount,omitempty"`
	PretaxGrossAmount         *float32 `json:"PretaxGrossAmount,omitempty" xml:"PretaxGrossAmount,omitempty"`
	ProductCode               *string  `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductDetail             *string  `json:"ProductDetail,omitempty" xml:"ProductDetail,omitempty"`
	ProductName               *string  `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	ProductType               *string  `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	Region                    *string  `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceGroup             *string  `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	ServicePeriod             *string  `json:"ServicePeriod,omitempty" xml:"ServicePeriod,omitempty"`
	ServicePeriodUnit         *string  `json:"ServicePeriodUnit,omitempty" xml:"ServicePeriodUnit,omitempty"`
	SplitAccountID            *string  `json:"SplitAccountID,omitempty" xml:"SplitAccountID,omitempty"`
	SplitAccountName          *string  `json:"SplitAccountName,omitempty" xml:"SplitAccountName,omitempty"`
	SplitBillingCycle         *string  `json:"SplitBillingCycle,omitempty" xml:"SplitBillingCycle,omitempty"`
	SplitCommodityCode        *string  `json:"SplitCommodityCode,omitempty" xml:"SplitCommodityCode,omitempty"`
	SplitItemID               *string  `json:"SplitItemID,omitempty" xml:"SplitItemID,omitempty"`
	SplitItemName             *string  `json:"SplitItemName,omitempty" xml:"SplitItemName,omitempty"`
	SplitProductDetail        *string  `json:"SplitProductDetail,omitempty" xml:"SplitProductDetail,omitempty"`
	SubscriptionType          *string  `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	Tag                       *string  `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Usage                     *string  `json:"Usage,omitempty" xml:"Usage,omitempty"`
	UsageUnit                 *string  `json:"UsageUnit,omitempty" xml:"UsageUnit,omitempty"`
	Zone                      *string  `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s QuerySplitItemBillResponseBodyDataItemsItem) String() string {
	return dara.Prettify(s)
}

func (s QuerySplitItemBillResponseBodyDataItemsItem) GoString() string {
	return s.String()
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetAdjustAmount() *float32 {
	return s.AdjustAmount
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetBillingDate() *string {
	return s.BillingDate
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetBillingItem() *string {
	return s.BillingItem
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetBillingType() *string {
	return s.BillingType
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetCashAmount() *float32 {
	return s.CashAmount
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetCostUnit() *string {
	return s.CostUnit
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetCurrency() *string {
	return s.Currency
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetDeductedByCashCoupons() *float32 {
	return s.DeductedByCashCoupons
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetDeductedByCoupons() *float32 {
	return s.DeductedByCoupons
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetDeductedByPrepaidCard() *float32 {
	return s.DeductedByPrepaidCard
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetDeductedByResourcePackage() *string {
	return s.DeductedByResourcePackage
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetInstanceConfig() *string {
	return s.InstanceConfig
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetInstanceID() *string {
	return s.InstanceID
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetInternetIP() *string {
	return s.InternetIP
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetIntranetIP() *string {
	return s.IntranetIP
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetInvoiceDiscount() *float32 {
	return s.InvoiceDiscount
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetItem() *string {
	return s.Item
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetListPrice() *string {
	return s.ListPrice
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetListPriceUnit() *string {
	return s.ListPriceUnit
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetNickName() *string {
	return s.NickName
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetOutstandingAmount() *float32 {
	return s.OutstandingAmount
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetOwnerID() *string {
	return s.OwnerID
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetPaymentAmount() *float32 {
	return s.PaymentAmount
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetPipCode() *string {
	return s.PipCode
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetPretaxAmount() *float32 {
	return s.PretaxAmount
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetPretaxGrossAmount() *float32 {
	return s.PretaxGrossAmount
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetProductCode() *string {
	return s.ProductCode
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetProductDetail() *string {
	return s.ProductDetail
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetProductName() *string {
	return s.ProductName
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetProductType() *string {
	return s.ProductType
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetRegion() *string {
	return s.Region
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetServicePeriod() *string {
	return s.ServicePeriod
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetServicePeriodUnit() *string {
	return s.ServicePeriodUnit
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetSplitAccountID() *string {
	return s.SplitAccountID
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetSplitAccountName() *string {
	return s.SplitAccountName
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetSplitBillingCycle() *string {
	return s.SplitBillingCycle
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetSplitCommodityCode() *string {
	return s.SplitCommodityCode
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetSplitItemID() *string {
	return s.SplitItemID
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetSplitItemName() *string {
	return s.SplitItemName
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetSplitProductDetail() *string {
	return s.SplitProductDetail
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetTag() *string {
	return s.Tag
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetUsage() *string {
	return s.Usage
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetUsageUnit() *string {
	return s.UsageUnit
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) GetZone() *string {
	return s.Zone
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetAdjustAmount(v float32) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.AdjustAmount = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetBillingDate(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.BillingDate = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetBillingItem(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.BillingItem = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetBillingType(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.BillingType = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetCashAmount(v float32) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.CashAmount = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetCommodityCode(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.CommodityCode = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetCostUnit(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.CostUnit = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetCurrency(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.Currency = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetDeductedByCashCoupons(v float32) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.DeductedByCashCoupons = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetDeductedByCoupons(v float32) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.DeductedByCoupons = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetDeductedByPrepaidCard(v float32) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.DeductedByPrepaidCard = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetDeductedByResourcePackage(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.DeductedByResourcePackage = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetInstanceConfig(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.InstanceConfig = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetInstanceID(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.InstanceID = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetInstanceSpec(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.InstanceSpec = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetInternetIP(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.InternetIP = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetIntranetIP(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.IntranetIP = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetInvoiceDiscount(v float32) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.InvoiceDiscount = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetItem(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.Item = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetListPrice(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.ListPrice = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetListPriceUnit(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.ListPriceUnit = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetNickName(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.NickName = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetOutstandingAmount(v float32) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.OutstandingAmount = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetOwnerID(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.OwnerID = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetPaymentAmount(v float32) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.PaymentAmount = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetPipCode(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.PipCode = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetPretaxAmount(v float32) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.PretaxAmount = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetPretaxGrossAmount(v float32) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.PretaxGrossAmount = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetProductCode(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.ProductCode = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetProductDetail(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.ProductDetail = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetProductName(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.ProductName = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetProductType(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.ProductType = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetRegion(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.Region = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetResourceGroup(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.ResourceGroup = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetServicePeriod(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.ServicePeriod = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetServicePeriodUnit(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.ServicePeriodUnit = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetSplitAccountID(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.SplitAccountID = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetSplitAccountName(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.SplitAccountName = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetSplitBillingCycle(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.SplitBillingCycle = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetSplitCommodityCode(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.SplitCommodityCode = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetSplitItemID(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.SplitItemID = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetSplitItemName(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.SplitItemName = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetSplitProductDetail(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.SplitProductDetail = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetSubscriptionType(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.SubscriptionType = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetTag(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.Tag = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetUsage(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.Usage = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetUsageUnit(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.UsageUnit = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) SetZone(v string) *QuerySplitItemBillResponseBodyDataItemsItem {
	s.Zone = &v
	return s
}

func (s *QuerySplitItemBillResponseBodyDataItemsItem) Validate() error {
	return dara.Validate(s)
}
