// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInstanceGaapCostResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryInstanceGaapCostResponseBody
	GetCode() *string
	SetData(v *QueryInstanceGaapCostResponseBodyData) *QueryInstanceGaapCostResponseBody
	GetData() *QueryInstanceGaapCostResponseBodyData
	SetMessage(v string) *QueryInstanceGaapCostResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryInstanceGaapCostResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryInstanceGaapCostResponseBody
	GetSuccess() *bool
}

type QueryInstanceGaapCostResponseBody struct {
	// example:
	//
	// Success
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *QueryInstanceGaapCostResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// CCBB1BB9-22F1-4177-867B-7A75D665B488
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryInstanceGaapCostResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceGaapCostResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInstanceGaapCostResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryInstanceGaapCostResponseBody) GetData() *QueryInstanceGaapCostResponseBodyData {
	return s.Data
}

func (s *QueryInstanceGaapCostResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryInstanceGaapCostResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryInstanceGaapCostResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryInstanceGaapCostResponseBody) SetCode(v string) *QueryInstanceGaapCostResponseBody {
	s.Code = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBody) SetData(v *QueryInstanceGaapCostResponseBodyData) *QueryInstanceGaapCostResponseBody {
	s.Data = v
	return s
}

func (s *QueryInstanceGaapCostResponseBody) SetMessage(v string) *QueryInstanceGaapCostResponseBody {
	s.Message = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBody) SetRequestId(v string) *QueryInstanceGaapCostResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBody) SetSuccess(v bool) *QueryInstanceGaapCostResponseBody {
	s.Success = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryInstanceGaapCostResponseBodyData struct {
	// example:
	//
	// 1.1.1.1
	HostId  *string                                       `json:"HostId,omitempty" xml:"HostId,omitempty"`
	Modules *QueryInstanceGaapCostResponseBodyDataModules `json:"Modules,omitempty" xml:"Modules,omitempty" type:"Struct"`
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
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryInstanceGaapCostResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceGaapCostResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryInstanceGaapCostResponseBodyData) GetHostId() *string {
	return s.HostId
}

func (s *QueryInstanceGaapCostResponseBodyData) GetModules() *QueryInstanceGaapCostResponseBodyDataModules {
	return s.Modules
}

func (s *QueryInstanceGaapCostResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryInstanceGaapCostResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryInstanceGaapCostResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryInstanceGaapCostResponseBodyData) SetHostId(v string) *QueryInstanceGaapCostResponseBodyData {
	s.HostId = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyData) SetModules(v *QueryInstanceGaapCostResponseBodyDataModules) *QueryInstanceGaapCostResponseBodyData {
	s.Modules = v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyData) SetPageNum(v int32) *QueryInstanceGaapCostResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyData) SetPageSize(v int32) *QueryInstanceGaapCostResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyData) SetTotalCount(v int32) *QueryInstanceGaapCostResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type QueryInstanceGaapCostResponseBodyDataModules struct {
	Module []*QueryInstanceGaapCostResponseBodyDataModulesModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Repeated"`
}

func (s QueryInstanceGaapCostResponseBodyDataModules) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceGaapCostResponseBodyDataModules) GoString() string {
	return s.String()
}

func (s *QueryInstanceGaapCostResponseBodyDataModules) GetModule() []*QueryInstanceGaapCostResponseBodyDataModulesModule {
	return s.Module
}

func (s *QueryInstanceGaapCostResponseBodyDataModules) SetModule(v []*QueryInstanceGaapCostResponseBodyDataModulesModule) *QueryInstanceGaapCostResponseBodyDataModules {
	s.Module = v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModules) Validate() error {
	return dara.Validate(s)
}

type QueryInstanceGaapCostResponseBodyDataModulesModule struct {
	AccountingUnit *string `json:"AccountingUnit,omitempty" xml:"AccountingUnit,omitempty"`
	// example:
	//
	// SubscriptionOrder
	BillType *string `json:"BillType,omitempty" xml:"BillType,omitempty"`
	// example:
	//
	// 2020-03
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 0
	DeductedByCashCoupons *string `json:"DeductedByCashCoupons,omitempty" xml:"DeductedByCashCoupons,omitempty"`
	// example:
	//
	// 0
	DeductedByCoupons *string `json:"DeductedByCoupons,omitempty" xml:"DeductedByCoupons,omitempty"`
	// example:
	//
	// 0
	DeductedByPrepaidCard *string `json:"DeductedByPrepaidCard,omitempty" xml:"DeductedByPrepaidCard,omitempty"`
	// example:
	//
	// 0
	GaapDeductedByCashCoupons *string `json:"GaapDeductedByCashCoupons,omitempty" xml:"GaapDeductedByCashCoupons,omitempty"`
	// example:
	//
	// 0
	GaapDeductedByCoupons *string `json:"GaapDeductedByCoupons,omitempty" xml:"GaapDeductedByCoupons,omitempty"`
	// example:
	//
	// 0
	GaapDeductedByPrepaidCard *string `json:"GaapDeductedByPrepaidCard,omitempty" xml:"GaapDeductedByPrepaidCard,omitempty"`
	// example:
	//
	// 0
	GaapPaymentAmount *string `json:"GaapPaymentAmount,omitempty" xml:"GaapPaymentAmount,omitempty"`
	// example:
	//
	// 0
	GaapPretaxAmount *string `json:"GaapPretaxAmount,omitempty" xml:"GaapPretaxAmount,omitempty"`
	// example:
	//
	// 0
	GaapPretaxAmountLocal *string `json:"GaapPretaxAmountLocal,omitempty" xml:"GaapPretaxAmountLocal,omitempty"`
	// example:
	//
	// 0
	GaapPretaxGrossAmount *string `json:"GaapPretaxGrossAmount,omitempty" xml:"GaapPretaxGrossAmount,omitempty"`
	// example:
	//
	// 0
	GaapPricingDiscount *string `json:"GaapPricingDiscount,omitempty" xml:"GaapPricingDiscount,omitempty"`
	// example:
	//
	// OSSBAG-cn-0xl0xxxxxx
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// example:
	//
	// 0.75
	MonthGaapDeductedByCashCoupons *string `json:"MonthGaapDeductedByCashCoupons,omitempty" xml:"MonthGaapDeductedByCashCoupons,omitempty"`
	// example:
	//
	// 0
	MonthGaapDeductedByCoupons *string `json:"MonthGaapDeductedByCoupons,omitempty" xml:"MonthGaapDeductedByCoupons,omitempty"`
	// example:
	//
	// 0
	MonthGaapDeductedByPrepaidCard *string `json:"MonthGaapDeductedByPrepaidCard,omitempty" xml:"MonthGaapDeductedByPrepaidCard,omitempty"`
	// example:
	//
	// 0
	MonthGaapPaymentAmount *string `json:"MonthGaapPaymentAmount,omitempty" xml:"MonthGaapPaymentAmount,omitempty"`
	// example:
	//
	// 0
	MonthGaapPretaxAmount *string `json:"MonthGaapPretaxAmount,omitempty" xml:"MonthGaapPretaxAmount,omitempty"`
	// example:
	//
	// 0.99
	MonthGaapPretaxAmountLocal *string `json:"MonthGaapPretaxAmountLocal,omitempty" xml:"MonthGaapPretaxAmountLocal,omitempty"`
	// example:
	//
	// 0.99
	MonthGaapPretaxGrossAmount *string `json:"MonthGaapPretaxGrossAmount,omitempty" xml:"MonthGaapPretaxGrossAmount,omitempty"`
	// example:
	//
	// 0,.25
	MonthGaapPricingDiscount *string `json:"MonthGaapPricingDiscount,omitempty" xml:"MonthGaapPricingDiscount,omitempty"`
	// example:
	//
	// 213123213123
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// New
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// example:
	//
	// 123213123123
	OwnerID *string `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	// example:
	//
	// 2018-06-15 15:59:57
	PayTime *string `json:"PayTime,omitempty" xml:"PayTime,omitempty"`
	// example:
	//
	// 23534534
	PayerAccount *string `json:"PayerAccount,omitempty" xml:"PayerAccount,omitempty"`
	// example:
	//
	// 0
	PaymentAmount *string `json:"PaymentAmount,omitempty" xml:"PaymentAmount,omitempty"`
	// example:
	//
	// CNY
	PaymentCurrency *string `json:"PaymentCurrency,omitempty" xml:"PaymentCurrency,omitempty"`
	// example:
	//
	// 0
	PretaxAmount *string `json:"PretaxAmount,omitempty" xml:"PretaxAmount,omitempty"`
	// example:
	//
	// 0
	PretaxAmountLocal *string `json:"PretaxAmountLocal,omitempty" xml:"PretaxAmountLocal,omitempty"`
	// example:
	//
	// 123
	PretaxGrossAmount *string `json:"PretaxGrossAmount,omitempty" xml:"PretaxGrossAmount,omitempty"`
	// example:
	//
	// 1
	PricingDiscount *string `json:"PricingDiscount,omitempty" xml:"PricingDiscount,omitempty"`
	// example:
	//
	// rds
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// rds
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// example:
	//
	// 12434345
	SubOrderId *string `json:"SubOrderId,omitempty" xml:"SubOrderId,omitempty"`
	// example:
	//
	// PayAsYouGo
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	// example:
	//
	// tag
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// example:
	//
	// 0
	UnallocatedDeductedByCashCoupons *string `json:"UnallocatedDeductedByCashCoupons,omitempty" xml:"UnallocatedDeductedByCashCoupons,omitempty"`
	// example:
	//
	// 0
	UnallocatedDeductedByCoupons *string `json:"UnallocatedDeductedByCoupons,omitempty" xml:"UnallocatedDeductedByCoupons,omitempty"`
	// example:
	//
	// 0
	UnallocatedDeductedByPrepaidCard *string `json:"UnallocatedDeductedByPrepaidCard,omitempty" xml:"UnallocatedDeductedByPrepaidCard,omitempty"`
	// example:
	//
	// 0
	UnallocatedPaymentAmount *string `json:"UnallocatedPaymentAmount,omitempty" xml:"UnallocatedPaymentAmount,omitempty"`
	// example:
	//
	// 0
	UnallocatedPretaxAmount *string `json:"UnallocatedPretaxAmount,omitempty" xml:"UnallocatedPretaxAmount,omitempty"`
	// example:
	//
	// 0
	UnallocatedPretaxAmountLocal *string `json:"UnallocatedPretaxAmountLocal,omitempty" xml:"UnallocatedPretaxAmountLocal,omitempty"`
	// example:
	//
	// 0
	UnallocatedPretaxGrossAmount *string `json:"UnallocatedPretaxGrossAmount,omitempty" xml:"UnallocatedPretaxGrossAmount,omitempty"`
	// example:
	//
	// 0
	UnallocatedPricingDiscount *string `json:"UnallocatedPricingDiscount,omitempty" xml:"UnallocatedPricingDiscount,omitempty"`
	// example:
	//
	// 2019-05-01 00:00:00
	UsageEndDate *string `json:"UsageEndDate,omitempty" xml:"UsageEndDate,omitempty"`
	// example:
	//
	// 2019-04-01 00:00:00
	UsageStartDate *string `json:"UsageStartDate,omitempty" xml:"UsageStartDate,omitempty"`
}

func (s QueryInstanceGaapCostResponseBodyDataModulesModule) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceGaapCostResponseBodyDataModulesModule) GoString() string {
	return s.String()
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetAccountingUnit() *string {
	return s.AccountingUnit
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetBillType() *string {
	return s.BillType
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetCurrency() *string {
	return s.Currency
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetDeductedByCashCoupons() *string {
	return s.DeductedByCashCoupons
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetDeductedByCoupons() *string {
	return s.DeductedByCoupons
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetDeductedByPrepaidCard() *string {
	return s.DeductedByPrepaidCard
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetGaapDeductedByCashCoupons() *string {
	return s.GaapDeductedByCashCoupons
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetGaapDeductedByCoupons() *string {
	return s.GaapDeductedByCoupons
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetGaapDeductedByPrepaidCard() *string {
	return s.GaapDeductedByPrepaidCard
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetGaapPaymentAmount() *string {
	return s.GaapPaymentAmount
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetGaapPretaxAmount() *string {
	return s.GaapPretaxAmount
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetGaapPretaxAmountLocal() *string {
	return s.GaapPretaxAmountLocal
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetGaapPretaxGrossAmount() *string {
	return s.GaapPretaxGrossAmount
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetGaapPricingDiscount() *string {
	return s.GaapPricingDiscount
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetInstanceID() *string {
	return s.InstanceID
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetMonthGaapDeductedByCashCoupons() *string {
	return s.MonthGaapDeductedByCashCoupons
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetMonthGaapDeductedByCoupons() *string {
	return s.MonthGaapDeductedByCoupons
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetMonthGaapDeductedByPrepaidCard() *string {
	return s.MonthGaapDeductedByPrepaidCard
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetMonthGaapPaymentAmount() *string {
	return s.MonthGaapPaymentAmount
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetMonthGaapPretaxAmount() *string {
	return s.MonthGaapPretaxAmount
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetMonthGaapPretaxAmountLocal() *string {
	return s.MonthGaapPretaxAmountLocal
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetMonthGaapPretaxGrossAmount() *string {
	return s.MonthGaapPretaxGrossAmount
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetMonthGaapPricingDiscount() *string {
	return s.MonthGaapPricingDiscount
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetOrderId() *string {
	return s.OrderId
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetOrderType() *string {
	return s.OrderType
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetOwnerID() *string {
	return s.OwnerID
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetPayTime() *string {
	return s.PayTime
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetPayerAccount() *string {
	return s.PayerAccount
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetPaymentAmount() *string {
	return s.PaymentAmount
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetPaymentCurrency() *string {
	return s.PaymentCurrency
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetPretaxAmount() *string {
	return s.PretaxAmount
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetPretaxAmountLocal() *string {
	return s.PretaxAmountLocal
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetPretaxGrossAmount() *string {
	return s.PretaxGrossAmount
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetPricingDiscount() *string {
	return s.PricingDiscount
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetProductCode() *string {
	return s.ProductCode
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetProductType() *string {
	return s.ProductType
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetRegion() *string {
	return s.Region
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetSubOrderId() *string {
	return s.SubOrderId
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetTag() *string {
	return s.Tag
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetUnallocatedDeductedByCashCoupons() *string {
	return s.UnallocatedDeductedByCashCoupons
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetUnallocatedDeductedByCoupons() *string {
	return s.UnallocatedDeductedByCoupons
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetUnallocatedDeductedByPrepaidCard() *string {
	return s.UnallocatedDeductedByPrepaidCard
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetUnallocatedPaymentAmount() *string {
	return s.UnallocatedPaymentAmount
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetUnallocatedPretaxAmount() *string {
	return s.UnallocatedPretaxAmount
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetUnallocatedPretaxAmountLocal() *string {
	return s.UnallocatedPretaxAmountLocal
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetUnallocatedPretaxGrossAmount() *string {
	return s.UnallocatedPretaxGrossAmount
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetUnallocatedPricingDiscount() *string {
	return s.UnallocatedPricingDiscount
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetUsageEndDate() *string {
	return s.UsageEndDate
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) GetUsageStartDate() *string {
	return s.UsageStartDate
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetAccountingUnit(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.AccountingUnit = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetBillType(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.BillType = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetBillingCycle(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.BillingCycle = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetCurrency(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.Currency = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetDeductedByCashCoupons(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.DeductedByCashCoupons = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetDeductedByCoupons(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.DeductedByCoupons = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetDeductedByPrepaidCard(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.DeductedByPrepaidCard = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetGaapDeductedByCashCoupons(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.GaapDeductedByCashCoupons = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetGaapDeductedByCoupons(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.GaapDeductedByCoupons = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetGaapDeductedByPrepaidCard(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.GaapDeductedByPrepaidCard = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetGaapPaymentAmount(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.GaapPaymentAmount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetGaapPretaxAmount(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.GaapPretaxAmount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetGaapPretaxAmountLocal(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.GaapPretaxAmountLocal = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetGaapPretaxGrossAmount(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.GaapPretaxGrossAmount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetGaapPricingDiscount(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.GaapPricingDiscount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetInstanceID(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.InstanceID = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetMonthGaapDeductedByCashCoupons(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.MonthGaapDeductedByCashCoupons = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetMonthGaapDeductedByCoupons(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.MonthGaapDeductedByCoupons = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetMonthGaapDeductedByPrepaidCard(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.MonthGaapDeductedByPrepaidCard = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetMonthGaapPaymentAmount(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.MonthGaapPaymentAmount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetMonthGaapPretaxAmount(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.MonthGaapPretaxAmount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetMonthGaapPretaxAmountLocal(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.MonthGaapPretaxAmountLocal = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetMonthGaapPretaxGrossAmount(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.MonthGaapPretaxGrossAmount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetMonthGaapPricingDiscount(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.MonthGaapPricingDiscount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetOrderId(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.OrderId = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetOrderType(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.OrderType = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetOwnerID(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.OwnerID = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetPayTime(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.PayTime = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetPayerAccount(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.PayerAccount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetPaymentAmount(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.PaymentAmount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetPaymentCurrency(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.PaymentCurrency = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetPretaxAmount(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.PretaxAmount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetPretaxAmountLocal(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.PretaxAmountLocal = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetPretaxGrossAmount(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.PretaxGrossAmount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetPricingDiscount(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.PricingDiscount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetProductCode(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.ProductCode = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetProductType(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.ProductType = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetRegion(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.Region = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetResourceGroup(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.ResourceGroup = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetSubOrderId(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.SubOrderId = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetSubscriptionType(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.SubscriptionType = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetTag(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.Tag = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetUnallocatedDeductedByCashCoupons(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.UnallocatedDeductedByCashCoupons = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetUnallocatedDeductedByCoupons(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.UnallocatedDeductedByCoupons = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetUnallocatedDeductedByPrepaidCard(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.UnallocatedDeductedByPrepaidCard = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetUnallocatedPaymentAmount(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.UnallocatedPaymentAmount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetUnallocatedPretaxAmount(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.UnallocatedPretaxAmount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetUnallocatedPretaxAmountLocal(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.UnallocatedPretaxAmountLocal = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetUnallocatedPretaxGrossAmount(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.UnallocatedPretaxGrossAmount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetUnallocatedPricingDiscount(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.UnallocatedPricingDiscount = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetUsageEndDate(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.UsageEndDate = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) SetUsageStartDate(v string) *QueryInstanceGaapCostResponseBodyDataModulesModule {
	s.UsageStartDate = &v
	return s
}

func (s *QueryInstanceGaapCostResponseBodyDataModulesModule) Validate() error {
	return dara.Validate(s)
}
