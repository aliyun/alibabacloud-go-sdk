// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CommodityValue struct {
	// Result模型。
	Result *CommodityValueResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CommodityValue) String() string {
	return tea.Prettify(s)
}

func (s CommodityValue) GoString() string {
	return s.String()
}

func (s *CommodityValue) SetResult(v *CommodityValueResult) *CommodityValue {
	s.Result = v
	return s
}

type CommodityValueResult struct {
	// 订单信息。
	Order *CommodityValueResultOrder `json:"Order,omitempty" xml:"Order,omitempty" type:"Struct"`
	// 询价类型，可选值：
	//
	// 1. Buy：新购询价。
	//
	// 2. ModificationBuy：变配询价。
	//
	// example:
	//
	// Buy
	InquiryType *string `json:"InquiryType,omitempty" xml:"InquiryType,omitempty"`
	// 订单子项。
	SubOrders *CommodityValueResultSubOrders `json:"SubOrders,omitempty" xml:"SubOrders,omitempty" type:"Struct"`
	// 优惠券。
	Coupons []*CommodityValueResultCoupons `json:"Coupons,omitempty" xml:"Coupons,omitempty" type:"Repeated"`
}

func (s CommodityValueResult) String() string {
	return tea.Prettify(s)
}

func (s CommodityValueResult) GoString() string {
	return s.String()
}

func (s *CommodityValueResult) SetOrder(v *CommodityValueResultOrder) *CommodityValueResult {
	s.Order = v
	return s
}

func (s *CommodityValueResult) SetInquiryType(v string) *CommodityValueResult {
	s.InquiryType = &v
	return s
}

func (s *CommodityValueResult) SetSubOrders(v *CommodityValueResultSubOrders) *CommodityValueResult {
	s.SubOrders = v
	return s
}

func (s *CommodityValueResult) SetCoupons(v []*CommodityValueResultCoupons) *CommodityValueResult {
	s.Coupons = v
	return s
}

type CommodityValueResultOrder struct {
	// 货币代码。
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// 优惠后。
	//
	// example:
	//
	// 9.99
	TradeAmount *string `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
	// 抵扣金额。
	//
	// example:
	//
	// 1.99
	DiscountAmount *string `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// 优惠前。
	//
	// example:
	//
	// 11.98
	OriginalAmount *string `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
}

func (s CommodityValueResultOrder) String() string {
	return tea.Prettify(s)
}

func (s CommodityValueResultOrder) GoString() string {
	return s.String()
}

func (s *CommodityValueResultOrder) SetCurrency(v string) *CommodityValueResultOrder {
	s.Currency = &v
	return s
}

func (s *CommodityValueResultOrder) SetTradeAmount(v string) *CommodityValueResultOrder {
	s.TradeAmount = &v
	return s
}

func (s *CommodityValueResultOrder) SetDiscountAmount(v string) *CommodityValueResultOrder {
	s.DiscountAmount = &v
	return s
}

func (s *CommodityValueResultOrder) SetOriginalAmount(v string) *CommodityValueResultOrder {
	s.OriginalAmount = &v
	return s
}

type CommodityValueResultSubOrders struct {
	// 订单子项。
	SubOrder []*CommodityValueResultSubOrdersSubOrder `json:"SubOrder,omitempty" xml:"SubOrder,omitempty" type:"Repeated"`
}

func (s CommodityValueResultSubOrders) String() string {
	return tea.Prettify(s)
}

func (s CommodityValueResultSubOrders) GoString() string {
	return s.String()
}

func (s *CommodityValueResultSubOrders) SetSubOrder(v []*CommodityValueResultSubOrdersSubOrder) *CommodityValueResultSubOrders {
	s.SubOrder = v
	return s
}

type CommodityValueResultSubOrdersSubOrder struct {
	// 模块（实例）信息。
	ModuleInstance []*CommodityValueResultSubOrdersSubOrderModuleInstance `json:"ModuleInstance,omitempty" xml:"ModuleInstance,omitempty" type:"Repeated"`
}

func (s CommodityValueResultSubOrdersSubOrder) String() string {
	return tea.Prettify(s)
}

func (s CommodityValueResultSubOrdersSubOrder) GoString() string {
	return s.String()
}

func (s *CommodityValueResultSubOrdersSubOrder) SetModuleInstance(v []*CommodityValueResultSubOrdersSubOrderModuleInstance) *CommodityValueResultSubOrdersSubOrder {
	s.ModuleInstance = v
	return s
}

type CommodityValueResultSubOrdersSubOrderModuleInstance struct {
	// 模块ID。
	//
	// example:
	//
	// 1234
	ModuleId *int64 `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	// 模块名称。
	//
	// example:
	//
	// Rds
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// 模块代码。
	//
	// example:
	//
	// rds_dbtype
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	// 产品原价（元）。
	//
	// example:
	//
	// 10.00
	TotalProductFee *float64 `json:"TotalProductFee,omitempty" xml:"TotalProductFee,omitempty"`
	// 折扣费用（元）。
	//
	// example:
	//
	// 1.99
	DiscountFee *float64 `json:"DiscountFee,omitempty" xml:"DiscountFee,omitempty"`
	// 实付金额（元）。
	//
	// example:
	//
	// 8.01
	PayFee *float64 `json:"PayFee,omitempty" xml:"PayFee,omitempty"`
	// 价格单位。
	//
	// example:
	//
	// 元/GB/小时
	PriceUnit *string `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
	// 是否计价项。
	//
	// example:
	//
	// true
	IsPricingModule *bool `json:"IsPricingModule,omitempty" xml:"IsPricingModule,omitempty"`
	// 在订单中是否需要支付。
	//
	// example:
	//
	// true
	NeedOrderPay *bool `json:"NeedOrderPay,omitempty" xml:"NeedOrderPay,omitempty"`
	// 定价类型。
	//
	// example:
	//
	// hourPrice
	PriceType *string `json:"PriceType,omitempty" xml:"PriceType,omitempty"`
	// 模块属性。
	ModuleAttrs []*CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs `json:"ModuleAttrs,omitempty" xml:"ModuleAttrs,omitempty" type:"Repeated"`
}

func (s CommodityValueResultSubOrdersSubOrderModuleInstance) String() string {
	return tea.Prettify(s)
}

func (s CommodityValueResultSubOrdersSubOrderModuleInstance) GoString() string {
	return s.String()
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) SetModuleId(v int64) *CommodityValueResultSubOrdersSubOrderModuleInstance {
	s.ModuleId = &v
	return s
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) SetModuleName(v string) *CommodityValueResultSubOrdersSubOrderModuleInstance {
	s.ModuleName = &v
	return s
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) SetModuleCode(v string) *CommodityValueResultSubOrdersSubOrderModuleInstance {
	s.ModuleCode = &v
	return s
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) SetTotalProductFee(v float64) *CommodityValueResultSubOrdersSubOrderModuleInstance {
	s.TotalProductFee = &v
	return s
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) SetDiscountFee(v float64) *CommodityValueResultSubOrdersSubOrderModuleInstance {
	s.DiscountFee = &v
	return s
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) SetPayFee(v float64) *CommodityValueResultSubOrdersSubOrderModuleInstance {
	s.PayFee = &v
	return s
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) SetPriceUnit(v string) *CommodityValueResultSubOrdersSubOrderModuleInstance {
	s.PriceUnit = &v
	return s
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) SetIsPricingModule(v bool) *CommodityValueResultSubOrdersSubOrderModuleInstance {
	s.IsPricingModule = &v
	return s
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) SetNeedOrderPay(v bool) *CommodityValueResultSubOrdersSubOrderModuleInstance {
	s.NeedOrderPay = &v
	return s
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) SetPriceType(v string) *CommodityValueResultSubOrdersSubOrderModuleInstance {
	s.PriceType = &v
	return s
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) SetModuleAttrs(v []*CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs) *CommodityValueResultSubOrdersSubOrderModuleInstance {
	s.ModuleAttrs = v
	return s
}

type CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs struct {
	// 属性类型，可选值：
	//
	// 1. 1：商品属性
	//
	// 2. 2：规格属性
	//
	// 3. 3：模块属性
	//
	// 4. 4：外部参数（备用）
	//
	// example:
	//
	// 3
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
	// Name
	//
	// example:
	//
	// 20GB
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Module attr code
	//
	// example:
	//
	// rds_storage
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Value
	//
	// example:
	//
	// 20
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// Unit
	//
	// example:
	//
	// GB
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs) String() string {
	return tea.Prettify(s)
}

func (s CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs) GoString() string {
	return s.String()
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs) SetType(v int64) *CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs {
	s.Type = &v
	return s
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs) SetName(v string) *CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs {
	s.Name = &v
	return s
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs) SetCode(v string) *CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs {
	s.Code = &v
	return s
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs) SetValue(v string) *CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs {
	s.Value = &v
	return s
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs) SetUnit(v string) *CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs {
	s.Unit = &v
	return s
}

type CommodityValueResultCoupons struct {
	// 可支付金额。
	//
	// example:
	//
	// 9.99
	CanPromFee *float64 `json:"CanPromFee,omitempty" xml:"CanPromFee,omitempty"`
	// Coupon Description
	//
	// example:
	//
	// 10元优惠券（有效期至2024年9月8日）
	CouponDesc *string `json:"CouponDesc,omitempty" xml:"CouponDesc,omitempty"`
	// Coupon Name
	//
	// example:
	//
	// 10元优惠券
	CouponName *string `json:"CouponName,omitempty" xml:"CouponName,omitempty"`
	// Coupon OptionNo
	//
	// example:
	//
	// 50008800000xxxx
	CouponOptionNo *string `json:"CouponOptionNo,omitempty" xml:"CouponOptionNo,omitempty"`
	// 是否选中。
	//
	// example:
	//
	// true
	Selected *bool `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s CommodityValueResultCoupons) String() string {
	return tea.Prettify(s)
}

func (s CommodityValueResultCoupons) GoString() string {
	return s.String()
}

func (s *CommodityValueResultCoupons) SetCanPromFee(v float64) *CommodityValueResultCoupons {
	s.CanPromFee = &v
	return s
}

func (s *CommodityValueResultCoupons) SetCouponDesc(v string) *CommodityValueResultCoupons {
	s.CouponDesc = &v
	return s
}

func (s *CommodityValueResultCoupons) SetCouponName(v string) *CommodityValueResultCoupons {
	s.CouponName = &v
	return s
}

func (s *CommodityValueResultCoupons) SetCouponOptionNo(v string) *CommodityValueResultCoupons {
	s.CouponOptionNo = &v
	return s
}

func (s *CommodityValueResultCoupons) SetSelected(v bool) *CommodityValueResultCoupons {
	s.Selected = &v
	return s
}

type CancelServiceUsageRequest struct {
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to delete the application.
	//
	// >  After you delete the application, you must re-enter the application information the next time you submit an application.
	//
	// example:
	//
	// true
	NeedDelete *bool `json:"NeedDelete,omitempty" xml:"NeedDelete,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-d6fc5f949a9246xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s CancelServiceUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelServiceUsageRequest) GoString() string {
	return s.String()
}

func (s *CancelServiceUsageRequest) SetClientToken(v string) *CancelServiceUsageRequest {
	s.ClientToken = &v
	return s
}

func (s *CancelServiceUsageRequest) SetNeedDelete(v bool) *CancelServiceUsageRequest {
	s.NeedDelete = &v
	return s
}

func (s *CancelServiceUsageRequest) SetServiceId(v string) *CancelServiceUsageRequest {
	s.ServiceId = &v
	return s
}

type CancelServiceUsageResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelServiceUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelServiceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *CancelServiceUsageResponseBody) SetRequestId(v string) *CancelServiceUsageResponseBody {
	s.RequestId = &v
	return s
}

type CancelServiceUsageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelServiceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelServiceUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelServiceUsageResponse) GoString() string {
	return s.String()
}

func (s *CancelServiceUsageResponse) SetHeaders(v map[string]*string) *CancelServiceUsageResponse {
	s.Headers = v
	return s
}

func (s *CancelServiceUsageResponse) SetStatusCode(v int32) *CancelServiceUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelServiceUsageResponse) SetBody(v *CancelServiceUsageResponseBody) *CancelServiceUsageResponse {
	s.Body = v
	return s
}

type ChangeResourceGroupRequest struct {
	// The ID of the new resource group.
	//
	// You can view resource group IDs in the [Resource Management console](https://resourcemanager.console.aliyun.com/resource-groups) .
	//
	// example:
	//
	// rg-acfmzmhzo******
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the cloud resource that you want to move to a new resource group.
	//
	// example:
	//
	// si-5dc794a7fd254e******
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type. Valid values:
	//
	// 	- service: service
	//
	// 	- serviceinstance: service instance
	//
	// example:
	//
	// service
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ChangeResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupRequest) SetNewResourceGroupId(v string) *ChangeResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetRegionId(v string) *ChangeResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceId(v string) *ChangeResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceType(v string) *ChangeResourceGroupRequest {
	s.ResourceType = &v
	return s
}

type ChangeResourceGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 464C8CB6-A548-5206-B83C-D32A8E43EC21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponseBody) SetRequestId(v string) *ChangeResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type ChangeResourceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponse) SetHeaders(v map[string]*string) *ChangeResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ChangeResourceGroupResponse) SetStatusCode(v int32) *ChangeResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeResourceGroupResponse) SetBody(v *ChangeResourceGroupResponseBody) *ChangeResourceGroupResponse {
	s.Body = v
	return s
}

type CheckServiceDeployableRequest struct {
	// Total amount of postpaid.
	//
	// example:
	//
	// 1.29
	PostPaidAmount *string `json:"PostPaidAmount,omitempty" xml:"PostPaidAmount,omitempty"`
	// Total amount of prepayment.
	//
	// example:
	//
	// 0.0
	PrePaidAmount *string `json:"PrePaidAmount,omitempty" xml:"PrePaidAmount,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-0e6fca6a51a544xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The trial type of the service instance. Valid values:
	//
	// 	- **Trial**: Trials are supported.
	//
	// 	- **NotTrial**: Trials are not supported.
	//
	// example:
	//
	// NotTrial
	TrialType *string `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
}

func (s CheckServiceDeployableRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceDeployableRequest) GoString() string {
	return s.String()
}

func (s *CheckServiceDeployableRequest) SetPostPaidAmount(v string) *CheckServiceDeployableRequest {
	s.PostPaidAmount = &v
	return s
}

func (s *CheckServiceDeployableRequest) SetPrePaidAmount(v string) *CheckServiceDeployableRequest {
	s.PrePaidAmount = &v
	return s
}

func (s *CheckServiceDeployableRequest) SetRegionId(v string) *CheckServiceDeployableRequest {
	s.RegionId = &v
	return s
}

func (s *CheckServiceDeployableRequest) SetServiceId(v string) *CheckServiceDeployableRequest {
	s.ServiceId = &v
	return s
}

func (s *CheckServiceDeployableRequest) SetServiceVersion(v string) *CheckServiceDeployableRequest {
	s.ServiceVersion = &v
	return s
}

func (s *CheckServiceDeployableRequest) SetTrialType(v string) *CheckServiceDeployableRequest {
	s.TrialType = &v
	return s
}

type CheckServiceDeployableResponseBody struct {
	// Inspection result.
	CheckResults []*CheckServiceDeployableResponseBodyCheckResults `json:"CheckResults,omitempty" xml:"CheckResults,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 06BF8F22-02DC-4750-83DF-3FFC11C065EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckServiceDeployableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceDeployableResponseBody) GoString() string {
	return s.String()
}

func (s *CheckServiceDeployableResponseBody) SetCheckResults(v []*CheckServiceDeployableResponseBodyCheckResults) *CheckServiceDeployableResponseBody {
	s.CheckResults = v
	return s
}

func (s *CheckServiceDeployableResponseBody) SetRequestId(v string) *CheckServiceDeployableResponseBody {
	s.RequestId = &v
	return s
}

type CheckServiceDeployableResponseBodyCheckResults struct {
	// Returns a hint message for the result.
	//
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Check type, invalid values:
	//
	// - Balance ：Account balance.
	//
	// - Quota:  Account quota.
	//
	// example:
	//
	// Balance
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Inspection result.
	//
	// example:
	//
	// true
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CheckServiceDeployableResponseBodyCheckResults) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceDeployableResponseBodyCheckResults) GoString() string {
	return s.String()
}

func (s *CheckServiceDeployableResponseBodyCheckResults) SetMessage(v string) *CheckServiceDeployableResponseBodyCheckResults {
	s.Message = &v
	return s
}

func (s *CheckServiceDeployableResponseBodyCheckResults) SetType(v string) *CheckServiceDeployableResponseBodyCheckResults {
	s.Type = &v
	return s
}

func (s *CheckServiceDeployableResponseBodyCheckResults) SetValue(v string) *CheckServiceDeployableResponseBodyCheckResults {
	s.Value = &v
	return s
}

type CheckServiceDeployableResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckServiceDeployableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckServiceDeployableResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceDeployableResponse) GoString() string {
	return s.String()
}

func (s *CheckServiceDeployableResponse) SetHeaders(v map[string]*string) *CheckServiceDeployableResponse {
	s.Headers = v
	return s
}

func (s *CheckServiceDeployableResponse) SetStatusCode(v int32) *CheckServiceDeployableResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckServiceDeployableResponse) SetBody(v *CheckServiceDeployableResponseBody) *CheckServiceDeployableResponse {
	s.Body = v
	return s
}

type ContinueDeployServiceInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run for the request to check information such as the permissions and instance status. Valid values:
	//
	// 	- true: performs a dry run for the request, but does not create a service instance.
	//
	// 	- false: performs a dry run for the request, and creates a service instance if the request passes the dry run.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The options that the system adopts when the system continues to create the service instance.
	Option []*string `json:"Option,omitempty" xml:"Option,omitempty" type:"Repeated"`
	// The parameters configured for the service instance.
	//
	// example:
	//
	// {"NodeCount": 3, "SystemDiskSize": 40, "InstancePassword": "******"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-0e6fca6a51a54420****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s ContinueDeployServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ContinueDeployServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *ContinueDeployServiceInstanceRequest) SetClientToken(v string) *ContinueDeployServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *ContinueDeployServiceInstanceRequest) SetDryRun(v bool) *ContinueDeployServiceInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *ContinueDeployServiceInstanceRequest) SetOption(v []*string) *ContinueDeployServiceInstanceRequest {
	s.Option = v
	return s
}

func (s *ContinueDeployServiceInstanceRequest) SetParameters(v string) *ContinueDeployServiceInstanceRequest {
	s.Parameters = &v
	return s
}

func (s *ContinueDeployServiceInstanceRequest) SetRegionId(v string) *ContinueDeployServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *ContinueDeployServiceInstanceRequest) SetServiceInstanceId(v string) *ContinueDeployServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

type ContinueDeployServiceInstanceResponseBody struct {
	// The dry run result.
	DryRunResult *ContinueDeployServiceInstanceResponseBodyDryRunResult `json:"DryRunResult,omitempty" xml:"DryRunResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the service instance.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s ContinueDeployServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ContinueDeployServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ContinueDeployServiceInstanceResponseBody) SetDryRunResult(v *ContinueDeployServiceInstanceResponseBodyDryRunResult) *ContinueDeployServiceInstanceResponseBody {
	s.DryRunResult = v
	return s
}

func (s *ContinueDeployServiceInstanceResponseBody) SetRequestId(v string) *ContinueDeployServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ContinueDeployServiceInstanceResponseBody) SetServiceInstanceId(v string) *ContinueDeployServiceInstanceResponseBody {
	s.ServiceInstanceId = &v
	return s
}

type ContinueDeployServiceInstanceResponseBodyDryRunResult struct {
	// The parameters that can be modified. The operation that is performed to modify the parameters does not cause a validation error.
	//
	// > This parameter is returned only if DryRun is set to true.
	ParametersAllowedToBeModified []*string `json:"ParametersAllowedToBeModified,omitempty" xml:"ParametersAllowedToBeModified,omitempty" type:"Repeated"`
	// The parameters that can be modified under specific conditions. The new values of the parameters determine whether the operation that is performed to modify the parameters causes a validation error.
	//
	// > This parameter is returned only if DryRun is set to true.
	ParametersConditionallyAllowedToBeModified []*string `json:"ParametersConditionallyAllowedToBeModified,omitempty" xml:"ParametersConditionallyAllowedToBeModified,omitempty" type:"Repeated"`
	// The parameters that cannot be modified. The operation that is performed to modify the parameters causes a validation error.
	//
	// > This parameter is returned only if DryRun is set to true.
	ParametersNotAllowedToBeModified []*string `json:"ParametersNotAllowedToBeModified,omitempty" xml:"ParametersNotAllowedToBeModified,omitempty" type:"Repeated"`
}

func (s ContinueDeployServiceInstanceResponseBodyDryRunResult) String() string {
	return tea.Prettify(s)
}

func (s ContinueDeployServiceInstanceResponseBodyDryRunResult) GoString() string {
	return s.String()
}

func (s *ContinueDeployServiceInstanceResponseBodyDryRunResult) SetParametersAllowedToBeModified(v []*string) *ContinueDeployServiceInstanceResponseBodyDryRunResult {
	s.ParametersAllowedToBeModified = v
	return s
}

func (s *ContinueDeployServiceInstanceResponseBodyDryRunResult) SetParametersConditionallyAllowedToBeModified(v []*string) *ContinueDeployServiceInstanceResponseBodyDryRunResult {
	s.ParametersConditionallyAllowedToBeModified = v
	return s
}

func (s *ContinueDeployServiceInstanceResponseBodyDryRunResult) SetParametersNotAllowedToBeModified(v []*string) *ContinueDeployServiceInstanceResponseBodyDryRunResult {
	s.ParametersNotAllowedToBeModified = v
	return s
}

type ContinueDeployServiceInstanceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ContinueDeployServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ContinueDeployServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ContinueDeployServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *ContinueDeployServiceInstanceResponse) SetHeaders(v map[string]*string) *ContinueDeployServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *ContinueDeployServiceInstanceResponse) SetStatusCode(v int32) *ContinueDeployServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ContinueDeployServiceInstanceResponse) SetBody(v *ContinueDeployServiceInstanceResponseBody) *ContinueDeployServiceInstanceResponse {
	s.Body = v
	return s
}

type CreateServiceInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The information about the order placed in Alibaba Cloud Marketplace. You do not need to specify this parameter if the service is not published in Alibaba Cloud Marketplace or uses the pay-as-you-go billing method.
	Commodity *CreateServiceInstanceRequestCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// The alert contact group.
	//
	// example:
	//
	// Default Group
	ContactGroup *string `json:"ContactGroup,omitempty" xml:"ContactGroup,omitempty"`
	// Specifies whether to perform only a dry run for the request to check information such as the permissions and instance status. Valid values:
	//
	// 	- **true**: performs a dry run for the request, but does not create a service instance.
	//
	// 	- **false**: performs a dry run for the request, and creates a service instance if the request passes the dry run.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether the service instance supports the hosted O\\&M feature. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnableInstanceOps *bool `json:"EnableInstanceOps,omitempty" xml:"EnableInstanceOps,omitempty"`
	// Specifies whether to enable the Prometheus monitoring feature. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnableUserPrometheus *bool `json:"EnableUserPrometheus,omitempty" xml:"EnableUserPrometheus,omitempty"`
	// The serviceInstance name.
	//
	// example:
	//
	// TestName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operation metadata.
	OperationMetadata *CreateServiceInstanceRequestOperationMetadata `json:"OperationMetadata,omitempty" xml:"OperationMetadata,omitempty" type:"Struct"`
	// The parameters that the customer specifies to deploy the service instance.
	//
	// >  If region information is required to create a service instance, you must specify the region ID in the value of Parameters.
	//
	// example:
	//
	// {"NodeCount": 3, "SystemDiskSize": 40, "InstancePassword": "******"}
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID. Valid values:
	//
	// 	- cn-hangzhou: China (Hangzhou).
	//
	// 	- ap-southeast-1: Singapore.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to automatically deduct the resource fees from the account balance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ResourceAutoPay *bool `json:"ResourceAutoPay,omitempty" xml:"ResourceAutoPay,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-0e6fca6a51a54420****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// Specification code.
	//
	// example:
	//
	// yuncode5425200001
	SpecificationCode *string `json:"SpecificationCode,omitempty" xml:"SpecificationCode,omitempty"`
	// The package name.
	//
	// example:
	//
	// Default Ppackage
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// The tags.
	Tag []*CreateServiceInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The name of the template.
	//
	// example:
	//
	// ECS Template
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The trial type of the service instance. Valid values:
	//
	// 	- **Trial**: Trials are supported.
	//
	// 	- **NotTrial**: Trials are not supported.
	//
	// example:
	//
	// Trial
	TrialType *string `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
}

func (s CreateServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceRequest) SetClientToken(v string) *CreateServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetCommodity(v *CreateServiceInstanceRequestCommodity) *CreateServiceInstanceRequest {
	s.Commodity = v
	return s
}

func (s *CreateServiceInstanceRequest) SetContactGroup(v string) *CreateServiceInstanceRequest {
	s.ContactGroup = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetDryRun(v bool) *CreateServiceInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetEnableInstanceOps(v bool) *CreateServiceInstanceRequest {
	s.EnableInstanceOps = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetEnableUserPrometheus(v bool) *CreateServiceInstanceRequest {
	s.EnableUserPrometheus = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetName(v string) *CreateServiceInstanceRequest {
	s.Name = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetOperationMetadata(v *CreateServiceInstanceRequestOperationMetadata) *CreateServiceInstanceRequest {
	s.OperationMetadata = v
	return s
}

func (s *CreateServiceInstanceRequest) SetParameters(v map[string]interface{}) *CreateServiceInstanceRequest {
	s.Parameters = v
	return s
}

func (s *CreateServiceInstanceRequest) SetRegionId(v string) *CreateServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetResourceAutoPay(v bool) *CreateServiceInstanceRequest {
	s.ResourceAutoPay = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetResourceGroupId(v string) *CreateServiceInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetServiceId(v string) *CreateServiceInstanceRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetServiceVersion(v string) *CreateServiceInstanceRequest {
	s.ServiceVersion = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetSpecificationCode(v string) *CreateServiceInstanceRequest {
	s.SpecificationCode = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetSpecificationName(v string) *CreateServiceInstanceRequest {
	s.SpecificationName = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetTag(v []*CreateServiceInstanceRequestTag) *CreateServiceInstanceRequest {
	s.Tag = v
	return s
}

func (s *CreateServiceInstanceRequest) SetTemplateName(v string) *CreateServiceInstanceRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetTrialType(v string) *CreateServiceInstanceRequest {
	s.TrialType = &v
	return s
}

type CreateServiceInstanceRequestCommodity struct {
	// Specifies whether to automatically complete the payment. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal for the service instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The coupon ID.
	//
	// example:
	//
	// 302070970220
	CouponId *string `json:"CouponId,omitempty" xml:"CouponId,omitempty"`
	// The subscription duration.
	//
	// example:
	//
	// 1
	PayPeriod *int64 `json:"PayPeriod,omitempty" xml:"PayPeriod,omitempty"`
	// The unit of the subscription duration. Valid values:
	//
	// 	- **Year**
	//
	// 	- **Month**
	//
	// 	- **Day**
	//
	// example:
	//
	// Year
	PayPeriodUnit *string `json:"PayPeriodUnit,omitempty" xml:"PayPeriodUnit,omitempty"`
}

func (s CreateServiceInstanceRequestCommodity) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceRequestCommodity) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceRequestCommodity) SetAutoPay(v bool) *CreateServiceInstanceRequestCommodity {
	s.AutoPay = &v
	return s
}

func (s *CreateServiceInstanceRequestCommodity) SetAutoRenew(v bool) *CreateServiceInstanceRequestCommodity {
	s.AutoRenew = &v
	return s
}

func (s *CreateServiceInstanceRequestCommodity) SetCouponId(v string) *CreateServiceInstanceRequestCommodity {
	s.CouponId = &v
	return s
}

func (s *CreateServiceInstanceRequestCommodity) SetPayPeriod(v int64) *CreateServiceInstanceRequestCommodity {
	s.PayPeriod = &v
	return s
}

func (s *CreateServiceInstanceRequestCommodity) SetPayPeriodUnit(v string) *CreateServiceInstanceRequestCommodity {
	s.PayPeriodUnit = &v
	return s
}

type CreateServiceInstanceRequestOperationMetadata struct {
	// The operation end time.
	//
	// example:
	//
	// 2022-01-28T06:48:56Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The additional information.
	//
	// example:
	//
	// ```json
	//
	//   {
	//
	//     "vncInfo": [
	//
	//       {
	//
	//         "instanceId": "i-001",
	//
	//         "username": "admin",
	//
	//         "password": "******",
	//
	//         "vncPassword": "******"
	//
	//       }
	//
	//     ]
	//
	//   }
	//
	//   ```
	ExtraInfo *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// Imported resource.
	//
	// example:
	//
	// {   "RegionId": "cn-hangzhou",   "Type": "ResourceIds",   "ResourceIds": {     "ALIYUN::ECS::INSTANCE": ["i-xxx", "i-yyy"],     "ALIYUN::RDS::INSTANCE": ["rm-xxx", "rm-yyy"],     "ALIYUN::VPC::VPC": ["vpc-xxx", "vpc-yyy"],     "ALIYUN::SLB::INSTANCE": ["lb-xxx", "lb-yyy"]   } }
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// The ID of the service instance.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The operation start time.
	//
	// example:
	//
	// 2021-12-29T06:48:56Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateServiceInstanceRequestOperationMetadata) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceRequestOperationMetadata) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceRequestOperationMetadata) SetEndTime(v string) *CreateServiceInstanceRequestOperationMetadata {
	s.EndTime = &v
	return s
}

func (s *CreateServiceInstanceRequestOperationMetadata) SetExtraInfo(v string) *CreateServiceInstanceRequestOperationMetadata {
	s.ExtraInfo = &v
	return s
}

func (s *CreateServiceInstanceRequestOperationMetadata) SetResources(v string) *CreateServiceInstanceRequestOperationMetadata {
	s.Resources = &v
	return s
}

func (s *CreateServiceInstanceRequestOperationMetadata) SetServiceInstanceId(v string) *CreateServiceInstanceRequestOperationMetadata {
	s.ServiceInstanceId = &v
	return s
}

func (s *CreateServiceInstanceRequestOperationMetadata) SetStartTime(v string) *CreateServiceInstanceRequestOperationMetadata {
	s.StartTime = &v
	return s
}

type CreateServiceInstanceRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateServiceInstanceRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceRequestTag) SetKey(v string) *CreateServiceInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateServiceInstanceRequestTag) SetValue(v string) *CreateServiceInstanceRequestTag {
	s.Value = &v
	return s
}

type CreateServiceInstanceShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The information about the order placed in Alibaba Cloud Marketplace. You do not need to specify this parameter if the service is not published in Alibaba Cloud Marketplace or uses the pay-as-you-go billing method.
	Commodity *CreateServiceInstanceShrinkRequestCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// The alert contact group.
	//
	// example:
	//
	// Default Group
	ContactGroup *string `json:"ContactGroup,omitempty" xml:"ContactGroup,omitempty"`
	// Specifies whether to perform only a dry run for the request to check information such as the permissions and instance status. Valid values:
	//
	// 	- **true**: performs a dry run for the request, but does not create a service instance.
	//
	// 	- **false**: performs a dry run for the request, and creates a service instance if the request passes the dry run.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether the service instance supports the hosted O\\&M feature. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnableInstanceOps *bool `json:"EnableInstanceOps,omitempty" xml:"EnableInstanceOps,omitempty"`
	// Specifies whether to enable the Prometheus monitoring feature. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnableUserPrometheus *bool `json:"EnableUserPrometheus,omitempty" xml:"EnableUserPrometheus,omitempty"`
	// The serviceInstance name.
	//
	// example:
	//
	// TestName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operation metadata.
	OperationMetadata *CreateServiceInstanceShrinkRequestOperationMetadata `json:"OperationMetadata,omitempty" xml:"OperationMetadata,omitempty" type:"Struct"`
	// The parameters that the customer specifies to deploy the service instance.
	//
	// >  If region information is required to create a service instance, you must specify the region ID in the value of Parameters.
	//
	// example:
	//
	// {"NodeCount": 3, "SystemDiskSize": 40, "InstancePassword": "******"}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID. Valid values:
	//
	// 	- cn-hangzhou: China (Hangzhou).
	//
	// 	- ap-southeast-1: Singapore.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to automatically deduct the resource fees from the account balance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ResourceAutoPay *bool `json:"ResourceAutoPay,omitempty" xml:"ResourceAutoPay,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-0e6fca6a51a54420****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// Specification code.
	//
	// example:
	//
	// yuncode5425200001
	SpecificationCode *string `json:"SpecificationCode,omitempty" xml:"SpecificationCode,omitempty"`
	// The package name.
	//
	// example:
	//
	// Default Ppackage
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// The tags.
	Tag []*CreateServiceInstanceShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The name of the template.
	//
	// example:
	//
	// ECS Template
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The trial type of the service instance. Valid values:
	//
	// 	- **Trial**: Trials are supported.
	//
	// 	- **NotTrial**: Trials are not supported.
	//
	// example:
	//
	// Trial
	TrialType *string `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
}

func (s CreateServiceInstanceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceShrinkRequest) SetClientToken(v string) *CreateServiceInstanceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetCommodity(v *CreateServiceInstanceShrinkRequestCommodity) *CreateServiceInstanceShrinkRequest {
	s.Commodity = v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetContactGroup(v string) *CreateServiceInstanceShrinkRequest {
	s.ContactGroup = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetDryRun(v bool) *CreateServiceInstanceShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetEnableInstanceOps(v bool) *CreateServiceInstanceShrinkRequest {
	s.EnableInstanceOps = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetEnableUserPrometheus(v bool) *CreateServiceInstanceShrinkRequest {
	s.EnableUserPrometheus = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetName(v string) *CreateServiceInstanceShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetOperationMetadata(v *CreateServiceInstanceShrinkRequestOperationMetadata) *CreateServiceInstanceShrinkRequest {
	s.OperationMetadata = v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetParametersShrink(v string) *CreateServiceInstanceShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetRegionId(v string) *CreateServiceInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetResourceAutoPay(v bool) *CreateServiceInstanceShrinkRequest {
	s.ResourceAutoPay = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetResourceGroupId(v string) *CreateServiceInstanceShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetServiceId(v string) *CreateServiceInstanceShrinkRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetServiceVersion(v string) *CreateServiceInstanceShrinkRequest {
	s.ServiceVersion = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetSpecificationCode(v string) *CreateServiceInstanceShrinkRequest {
	s.SpecificationCode = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetSpecificationName(v string) *CreateServiceInstanceShrinkRequest {
	s.SpecificationName = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetTag(v []*CreateServiceInstanceShrinkRequestTag) *CreateServiceInstanceShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetTemplateName(v string) *CreateServiceInstanceShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetTrialType(v string) *CreateServiceInstanceShrinkRequest {
	s.TrialType = &v
	return s
}

type CreateServiceInstanceShrinkRequestCommodity struct {
	// Specifies whether to automatically complete the payment. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal for the service instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The coupon ID.
	//
	// example:
	//
	// 302070970220
	CouponId *string `json:"CouponId,omitempty" xml:"CouponId,omitempty"`
	// The subscription duration.
	//
	// example:
	//
	// 1
	PayPeriod *int64 `json:"PayPeriod,omitempty" xml:"PayPeriod,omitempty"`
	// The unit of the subscription duration. Valid values:
	//
	// 	- **Year**
	//
	// 	- **Month**
	//
	// 	- **Day**
	//
	// example:
	//
	// Year
	PayPeriodUnit *string `json:"PayPeriodUnit,omitempty" xml:"PayPeriodUnit,omitempty"`
}

func (s CreateServiceInstanceShrinkRequestCommodity) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceShrinkRequestCommodity) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceShrinkRequestCommodity) SetAutoPay(v bool) *CreateServiceInstanceShrinkRequestCommodity {
	s.AutoPay = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequestCommodity) SetAutoRenew(v bool) *CreateServiceInstanceShrinkRequestCommodity {
	s.AutoRenew = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequestCommodity) SetCouponId(v string) *CreateServiceInstanceShrinkRequestCommodity {
	s.CouponId = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequestCommodity) SetPayPeriod(v int64) *CreateServiceInstanceShrinkRequestCommodity {
	s.PayPeriod = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequestCommodity) SetPayPeriodUnit(v string) *CreateServiceInstanceShrinkRequestCommodity {
	s.PayPeriodUnit = &v
	return s
}

type CreateServiceInstanceShrinkRequestOperationMetadata struct {
	// The operation end time.
	//
	// example:
	//
	// 2022-01-28T06:48:56Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The additional information.
	//
	// example:
	//
	// ```json
	//
	//   {
	//
	//     "vncInfo": [
	//
	//       {
	//
	//         "instanceId": "i-001",
	//
	//         "username": "admin",
	//
	//         "password": "******",
	//
	//         "vncPassword": "******"
	//
	//       }
	//
	//     ]
	//
	//   }
	//
	//   ```
	ExtraInfo *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// Imported resource.
	//
	// example:
	//
	// {   "RegionId": "cn-hangzhou",   "Type": "ResourceIds",   "ResourceIds": {     "ALIYUN::ECS::INSTANCE": ["i-xxx", "i-yyy"],     "ALIYUN::RDS::INSTANCE": ["rm-xxx", "rm-yyy"],     "ALIYUN::VPC::VPC": ["vpc-xxx", "vpc-yyy"],     "ALIYUN::SLB::INSTANCE": ["lb-xxx", "lb-yyy"]   } }
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// The ID of the service instance.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The operation start time.
	//
	// example:
	//
	// 2021-12-29T06:48:56Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateServiceInstanceShrinkRequestOperationMetadata) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceShrinkRequestOperationMetadata) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceShrinkRequestOperationMetadata) SetEndTime(v string) *CreateServiceInstanceShrinkRequestOperationMetadata {
	s.EndTime = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequestOperationMetadata) SetExtraInfo(v string) *CreateServiceInstanceShrinkRequestOperationMetadata {
	s.ExtraInfo = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequestOperationMetadata) SetResources(v string) *CreateServiceInstanceShrinkRequestOperationMetadata {
	s.Resources = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequestOperationMetadata) SetServiceInstanceId(v string) *CreateServiceInstanceShrinkRequestOperationMetadata {
	s.ServiceInstanceId = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequestOperationMetadata) SetStartTime(v string) *CreateServiceInstanceShrinkRequestOperationMetadata {
	s.StartTime = &v
	return s
}

type CreateServiceInstanceShrinkRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateServiceInstanceShrinkRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceShrinkRequestTag) SetKey(v string) *CreateServiceInstanceShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequestTag) SetValue(v string) *CreateServiceInstanceShrinkRequestTag {
	s.Value = &v
	return s
}

type CreateServiceInstanceResponseBody struct {
	// The MartketInstance ID.
	//
	// example:
	//
	// 786***45
	MarketInstanceId *string `json:"MarketInstanceId,omitempty" xml:"MarketInstanceId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 2306175xxxxxxxx
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service instance ID.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The status of the service instance. Valid values:
	//
	// 	- **Created**
	//
	// 	- **Deploying**
	//
	// 	- **DeployedFailed**
	//
	// 	- **Deployed**
	//
	// 	- **Upgrading**
	//
	// 	- **Deleting**
	//
	// 	- **Deleted**
	//
	// 	- **DeletedFailed**
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceResponseBody) SetMarketInstanceId(v string) *CreateServiceInstanceResponseBody {
	s.MarketInstanceId = &v
	return s
}

func (s *CreateServiceInstanceResponseBody) SetOrderId(v string) *CreateServiceInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateServiceInstanceResponseBody) SetRequestId(v string) *CreateServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceInstanceResponseBody) SetServiceInstanceId(v string) *CreateServiceInstanceResponseBody {
	s.ServiceInstanceId = &v
	return s
}

func (s *CreateServiceInstanceResponseBody) SetStatus(v string) *CreateServiceInstanceResponseBody {
	s.Status = &v
	return s
}

type CreateServiceInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceResponse) SetHeaders(v map[string]*string) *CreateServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceInstanceResponse) SetStatusCode(v int32) *CreateServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceInstanceResponse) SetBody(v *CreateServiceInstanceResponseBody) *CreateServiceInstanceResponse {
	s.Body = v
	return s
}

type CreateServiceUsageRequest struct {
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-d6fc5f949a9246xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The application information.
	UserInformation map[string]*string `json:"UserInformation,omitempty" xml:"UserInformation,omitempty"`
}

func (s CreateServiceUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceUsageRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceUsageRequest) SetClientToken(v string) *CreateServiceUsageRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceUsageRequest) SetServiceId(v string) *CreateServiceUsageRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateServiceUsageRequest) SetUserInformation(v map[string]*string) *CreateServiceUsageRequest {
	s.UserInformation = v
	return s
}

type CreateServiceUsageShrinkRequest struct {
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-d6fc5f949a9246xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The application information.
	UserInformationShrink *string `json:"UserInformation,omitempty" xml:"UserInformation,omitempty"`
}

func (s CreateServiceUsageShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceUsageShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceUsageShrinkRequest) SetClientToken(v string) *CreateServiceUsageShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceUsageShrinkRequest) SetServiceId(v string) *CreateServiceUsageShrinkRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateServiceUsageShrinkRequest) SetUserInformationShrink(v string) *CreateServiceUsageShrinkRequest {
	s.UserInformationShrink = &v
	return s
}

type CreateServiceUsageResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateServiceUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceUsageResponseBody) SetRequestId(v string) *CreateServiceUsageResponseBody {
	s.RequestId = &v
	return s
}

type CreateServiceUsageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceUsageResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceUsageResponse) SetHeaders(v map[string]*string) *CreateServiceUsageResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceUsageResponse) SetStatusCode(v int32) *CreateServiceUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceUsageResponse) SetBody(v *CreateServiceUsageResponseBody) *CreateServiceUsageResponse {
	s.Body = v
	return s
}

type DeleteServiceInstancesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the service instances.
	//
	// This parameter is required.
	ServiceInstanceId []*string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty" type:"Repeated"`
}

func (s DeleteServiceInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceInstancesRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceInstancesRequest) SetClientToken(v string) *DeleteServiceInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteServiceInstancesRequest) SetRegionId(v string) *DeleteServiceInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteServiceInstancesRequest) SetServiceInstanceId(v []*string) *DeleteServiceInstancesRequest {
	s.ServiceInstanceId = v
	return s
}

type DeleteServiceInstancesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceInstancesResponseBody) SetRequestId(v string) *DeleteServiceInstancesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServiceInstancesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceInstancesResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceInstancesResponse) SetHeaders(v map[string]*string) *DeleteServiceInstancesResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceInstancesResponse) SetStatusCode(v int32) *DeleteServiceInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceInstancesResponse) SetBody(v *DeleteServiceInstancesResponseBody) *DeleteServiceInstancesResponse {
	s.Body = v
	return s
}

type DeployServiceInstanceRequest struct {
	// Ensures idempotency of the request. Generate a unique value for this parameter from your client to ensure it is unique across different requests. ClientToken supports only ASCII characters and cannot exceed 64 characters.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Region ID. Allowed values:
	//
	// - cn-hangzhou: East China 1 (Hangzhou).
	//
	// - ap-southeast-1: Singapore.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Service instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-b58c874912fc4294****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s DeployServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeployServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeployServiceInstanceRequest) SetClientToken(v string) *DeployServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeployServiceInstanceRequest) SetRegionId(v string) *DeployServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DeployServiceInstanceRequest) SetServiceInstanceId(v string) *DeployServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

type DeployServiceInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeployServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeployServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeployServiceInstanceResponseBody) SetRequestId(v string) *DeployServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DeployServiceInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeployServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeployServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeployServiceInstanceResponse) SetHeaders(v map[string]*string) *DeployServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeployServiceInstanceResponse) SetStatusCode(v int32) *DeployServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployServiceInstanceResponse) SetBody(v *DeployServiceInstanceResponseBody) *DeployServiceInstanceResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponseBody struct {
	// The available regions.
	Regions []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 06BF8F22-02DC-4750-83DF-3FFC11C065EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	// The region endpoint.
	//
	// example:
	//
	// computenest.cn-hangzhou.aliyuncs.com
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type GenerateServicePolicyRequest struct {
	// The type of operation N for which you want to generate the policy information.
	//
	// Valid values:
	//
	// 	- CreateServiceInstance: creates a serviceInstance by calling the CreateServiceInstance operation.
	//
	// 	- UpdateServiceInstance: updates a serviceInstance by calling the UpdateServiceInstance operation.
	//
	// 	- DeleteServiceInstance: deletes a serviceInstance by calling the DeleteServiceInstance operation.
	//
	// >  The default value is the combination of all valid values.
	OperationTypes []*string `json:"OperationTypes,omitempty" xml:"OperationTypes,omitempty" type:"Repeated"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-b3e9ed878b0c4xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// draft
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// GPU-单机版
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The trial policy. Valid values:
	//
	// 	- Trial: Trials are supported.
	//
	// 	- NotTrial: Trials are not supported.
	//
	// example:
	//
	// NotTrial
	TrialType *string `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
}

func (s GenerateServicePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateServicePolicyRequest) GoString() string {
	return s.String()
}

func (s *GenerateServicePolicyRequest) SetOperationTypes(v []*string) *GenerateServicePolicyRequest {
	s.OperationTypes = v
	return s
}

func (s *GenerateServicePolicyRequest) SetRegionId(v string) *GenerateServicePolicyRequest {
	s.RegionId = &v
	return s
}

func (s *GenerateServicePolicyRequest) SetServiceId(v string) *GenerateServicePolicyRequest {
	s.ServiceId = &v
	return s
}

func (s *GenerateServicePolicyRequest) SetServiceVersion(v string) *GenerateServicePolicyRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GenerateServicePolicyRequest) SetTemplateName(v string) *GenerateServicePolicyRequest {
	s.TemplateName = &v
	return s
}

func (s *GenerateServicePolicyRequest) SetTrialType(v string) *GenerateServicePolicyRequest {
	s.TrialType = &v
	return s
}

type GenerateServicePolicyResponseBody struct {
	// The policies that are missing.
	MissingPolicy []*GenerateServicePolicyResponseBodyMissingPolicy `json:"MissingPolicy,omitempty" xml:"MissingPolicy,omitempty" type:"Repeated"`
	// The RAM policy.
	//
	// example:
	//
	// {Statement": [{ "Action": ["oos:*"], "Effect": "Allow", "Resource": "*"},{ "Action": ["ecs:DescribeInstances"], "Effect": "Allow", "Resource": "*"},{ "Action": ["ecs:RunInstance"], "Effect": "Allow", "Resource": "*"}], "Version": "1"}
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5040BE9E-8DA2-5C9D-9B70-0EE6027A14BC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateServicePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateServicePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateServicePolicyResponseBody) SetMissingPolicy(v []*GenerateServicePolicyResponseBodyMissingPolicy) *GenerateServicePolicyResponseBody {
	s.MissingPolicy = v
	return s
}

func (s *GenerateServicePolicyResponseBody) SetPolicy(v string) *GenerateServicePolicyResponseBody {
	s.Policy = &v
	return s
}

func (s *GenerateServicePolicyResponseBody) SetRequestId(v string) *GenerateServicePolicyResponseBody {
	s.RequestId = &v
	return s
}

type GenerateServicePolicyResponseBodyMissingPolicy struct {
	// Operations on specific resources.
	Action []*string `json:"Action,omitempty" xml:"Action,omitempty" type:"Repeated"`
	// The specific objects authorized. An asterisk (*) denotes all resources.
	//
	// example:
	//
	// *
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// ecs
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s GenerateServicePolicyResponseBodyMissingPolicy) String() string {
	return tea.Prettify(s)
}

func (s GenerateServicePolicyResponseBodyMissingPolicy) GoString() string {
	return s.String()
}

func (s *GenerateServicePolicyResponseBodyMissingPolicy) SetAction(v []*string) *GenerateServicePolicyResponseBodyMissingPolicy {
	s.Action = v
	return s
}

func (s *GenerateServicePolicyResponseBodyMissingPolicy) SetResource(v string) *GenerateServicePolicyResponseBodyMissingPolicy {
	s.Resource = &v
	return s
}

func (s *GenerateServicePolicyResponseBodyMissingPolicy) SetServiceName(v string) *GenerateServicePolicyResponseBodyMissingPolicy {
	s.ServiceName = &v
	return s
}

type GenerateServicePolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateServicePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateServicePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateServicePolicyResponse) GoString() string {
	return s.String()
}

func (s *GenerateServicePolicyResponse) SetHeaders(v map[string]*string) *GenerateServicePolicyResponse {
	s.Headers = v
	return s
}

func (s *GenerateServicePolicyResponse) SetStatusCode(v int32) *GenerateServicePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateServicePolicyResponse) SetBody(v *GenerateServicePolicyResponseBody) *GenerateServicePolicyResponse {
	s.Body = v
	return s
}

type GetServiceRequest struct {
	// Region Id.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-0e6fca6a51a544xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service instance id.
	//
	// example:
	//
	// si-b58c874912fc4294****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The service name.
	//
	// example:
	//
	// Wordpress
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1.0
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// Whether to disclose service details.
	ShowDetails []*string `json:"ShowDetails,omitempty" xml:"ShowDetails,omitempty" type:"Repeated"`
}

func (s GetServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceRequest) GoString() string {
	return s.String()
}

func (s *GetServiceRequest) SetRegionId(v string) *GetServiceRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceRequest) SetServiceId(v string) *GetServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *GetServiceRequest) SetServiceInstanceId(v string) *GetServiceRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *GetServiceRequest) SetServiceName(v string) *GetServiceRequest {
	s.ServiceName = &v
	return s
}

func (s *GetServiceRequest) SetServiceVersion(v string) *GetServiceRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GetServiceRequest) SetShowDetails(v []*string) *GetServiceRequest {
	s.ShowDetails = v
	return s
}

type GetServiceResponseBody struct {
	// The alert configurations of the service.
	//
	// >  This parameter takes effect only when you specify an alert policy for **PolicyNames**.
	//
	// example:
	//
	// { "TemplateUrl": "http://template.file.url", "ApplicationGroups": [ { "Name": "applicationGroup1", "TemplateUrl": "url1" }, { "Name": "applicationGroup2", "TemplateUrl": "url2" } ] }
	AlarmMetadata *string `json:"AlarmMetadata,omitempty" xml:"AlarmMetadata,omitempty"`
	// The categories of the Flow.
	//
	// example:
	//
	// AI
	Categories *string `json:"Categories,omitempty" xml:"Categories,omitempty"`
	// The information about the order placed in Alibaba Cloud Marketplace.
	Commodity *GetServiceResponseBodyCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// Service deployment approach, Valid values：
	//
	// - NoWhere
	//
	// - Marketplace
	//
	// - ComputeNest
	//
	// example:
	//
	// Marketplace
	DeployFrom *string `json:"DeployFrom,omitempty" xml:"DeployFrom,omitempty"`
	// The storage configurations of the service. The format in which the deployment information of a service is stored varies based on the deployment type of the service. In this case, the deployment information is stored in the JSON string format.
	//
	// example:
	//
	// {\\"TemplateUrl\\": \\"http://tidbRosFile\\"}
	DeployMetadata *string `json:"DeployMetadata,omitempty" xml:"DeployMetadata,omitempty"`
	// The deployment type of the service. Valid values:
	//
	// 	- ros: The service is deployed by using Resource Orchestration Service (ROS).
	//
	// 	- terraform: The service is deployed by using Terraform.
	//
	// example:
	//
	// ros
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// The duration for which hosted O\\&M is implemented. Unit: seconds.
	//
	// example:
	//
	// 259200
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Information about the ram role created in the service template.
	InstanceRoleInfos []*GetServiceResponseBodyInstanceRoleInfos `json:"InstanceRoleInfos,omitempty" xml:"InstanceRoleInfos,omitempty" type:"Repeated"`
	// Indicates whether the hosted O\\&M feature is enabled for the service. Default value: false. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// >  This parameter is returned if you set **ServiceType*	- to **private**.
	//
	// example:
	//
	// false
	IsSupportOperated *bool `json:"IsSupportOperated,omitempty" xml:"IsSupportOperated,omitempty"`
	// The license metadata.
	//
	// example:
	//
	// {\\"PayType\\":\\"CustomFixTime\\",\\"DefaultLicenseDays\\":7,\\"CustomMetadata\\":[{\\"TemplateName\\":\\"ECS\\",\\"SpecificationName\\":\\"bandwith-0\\",\\"CustomData\\":\\"1\\"}]}
	LicenseMetadata *string `json:"LicenseMetadata,omitempty" xml:"LicenseMetadata,omitempty"`
	// The logging configurations.
	//
	// example:
	//
	// {\\"Logstores\\":[]}
	LogMetadata *string `json:"LogMetadata,omitempty" xml:"LogMetadata,omitempty"`
	// The operation metadata.
	//
	// example:
	//
	// {\\"PrometheusConfigMap\\":{\\"New_Vpc_Ack_And_Jumpserver\\":{}}}
	OperationMetadata *string `json:"OperationMetadata,omitempty" xml:"OperationMetadata,omitempty"`
	// The permissions on the service. Valid values:
	//
	// 	- Deployable: Permissions to deploy the service.
	//
	// 	- Accessible: Permissions to access the service.
	//
	// example:
	//
	// Deployable
	Permission *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// The policy name. The name can be up to 128 characters in length. Separate multiple names with commas (,). Only hosted O\\&M policies are supported.
	//
	// example:
	//
	// policyName1, policyName2
	PolicyNames *string `json:"PolicyNames,omitempty" xml:"PolicyNames,omitempty"`
	// The time when the service was published.
	//
	// example:
	//
	// 2021-05-21T00:00:00Z
	PublishTime *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06BF8F22-02DC-4750-83DF-3FFC11C065EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Service document information.
	ServiceDocumentInfos []*GetServiceResponseBodyServiceDocumentInfos `json:"ServiceDocumentInfos,omitempty" xml:"ServiceDocumentInfos,omitempty" type:"Repeated"`
	// The service ID.
	//
	// example:
	//
	// service-0e6fca6a51a544xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The information about the service.
	ServiceInfos []*GetServiceResponseBodyServiceInfos `json:"ServiceInfos,omitempty" xml:"ServiceInfos,omitempty" type:"Repeated"`
	// The URL of the service page.
	//
	// example:
	//
	// http://example1.com
	ServiceProductUrl *string `json:"ServiceProductUrl,omitempty" xml:"ServiceProductUrl,omitempty"`
	// The type of the service. Valid values:
	//
	// - private: The service is a private service and is deployed within the account of a customer.
	//
	// - managed: The service is a fully managed service and is deployed within the account of a service provider.
	//
	// - operation: The service is a hosted O&M service.
	//
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The permission type of the deployment URL. Valid values:
	//
	// 	- Public: All users can go to the URL to create a service instance or a trial service instance.
	//
	// 	- Restricted: Only users in the whitelist can go to the URL to create a service instance or a trial service instance.
	//
	// 	- OnlyFormalRestricted: Only users in the whitelist can go to the URL to create a service instance.
	//
	// 	- OnlyTrailRestricted: Only users in the whitelist can go to the URL to create a trial service instance.
	//
	// 	- Hidden: Users not in the whitelist cannot see the service details page when they go to the URL and cannot request deployment permissions.
	//
	// example:
	//
	// Public
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The deploy status of the service. Valid values:
	//
	// - Draft
	//
	// - Beta
	//
	// - Submitted
	//
	// - Approved
	//
	// - Launching
	//
	// - Online
	//
	// - Offline
	//
	// - Creating
	//
	// example:
	//
	// Online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The description of service provider.
	//
	// example:
	//
	// Computing Nest Community service
	SupplierDesc *string `json:"SupplierDesc,omitempty" xml:"SupplierDesc,omitempty"`
	// The Logo of service provider.
	//
	// example:
	//
	// https://service-info-public.oss-cn-hangzhou.aliyuncs.com/xxx/service-image/xxx.png
	SupplierLogo *string `json:"SupplierLogo,omitempty" xml:"SupplierLogo,omitempty"`
	// The name of the service provider.
	//
	// example:
	//
	// Alibaba Cloud
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// The Alibaba Cloud account ID of the service provider.
	//
	// example:
	//
	// 158927391332*****
	SupplierUid *int64 `json:"SupplierUid,omitempty" xml:"SupplierUid,omitempty"`
	// The URL of the service provider.
	//
	// example:
	//
	// http://example.com
	SupplierUrl *string `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	// Contact information of the service provider
	SupportContacts []*GetServiceResponseBodySupportContacts `json:"SupportContacts,omitempty" xml:"SupportContacts,omitempty" type:"Repeated"`
	// The tags.
	Tags []*GetServiceResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The type of the tenant. Valid values:
	//
	// 	- SingleTenant
	//
	// 	- MultiTenant
	//
	// example:
	//
	// SingleTenant
	TenantType *string `json:"TenantType,omitempty" xml:"TenantType,omitempty"`
	// The trial duration. Unit: day. The maximum trial duration cannot exceed 30 days.
	//
	// example:
	//
	// 7
	TrialDuration *int64 `json:"TrialDuration,omitempty" xml:"TrialDuration,omitempty"`
	// The trial policy. Valid values:
	//
	// 	- Trial: Trials are supported.
	//
	// 	- NotTrial: Trials are not supported.
	//
	// example:
	//
	// Trial
	TrialType *string `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The version name.
	//
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBody) SetAlarmMetadata(v string) *GetServiceResponseBody {
	s.AlarmMetadata = &v
	return s
}

func (s *GetServiceResponseBody) SetCategories(v string) *GetServiceResponseBody {
	s.Categories = &v
	return s
}

func (s *GetServiceResponseBody) SetCommodity(v *GetServiceResponseBodyCommodity) *GetServiceResponseBody {
	s.Commodity = v
	return s
}

func (s *GetServiceResponseBody) SetDeployFrom(v string) *GetServiceResponseBody {
	s.DeployFrom = &v
	return s
}

func (s *GetServiceResponseBody) SetDeployMetadata(v string) *GetServiceResponseBody {
	s.DeployMetadata = &v
	return s
}

func (s *GetServiceResponseBody) SetDeployType(v string) *GetServiceResponseBody {
	s.DeployType = &v
	return s
}

func (s *GetServiceResponseBody) SetDuration(v int64) *GetServiceResponseBody {
	s.Duration = &v
	return s
}

func (s *GetServiceResponseBody) SetInstanceRoleInfos(v []*GetServiceResponseBodyInstanceRoleInfos) *GetServiceResponseBody {
	s.InstanceRoleInfos = v
	return s
}

func (s *GetServiceResponseBody) SetIsSupportOperated(v bool) *GetServiceResponseBody {
	s.IsSupportOperated = &v
	return s
}

func (s *GetServiceResponseBody) SetLicenseMetadata(v string) *GetServiceResponseBody {
	s.LicenseMetadata = &v
	return s
}

func (s *GetServiceResponseBody) SetLogMetadata(v string) *GetServiceResponseBody {
	s.LogMetadata = &v
	return s
}

func (s *GetServiceResponseBody) SetOperationMetadata(v string) *GetServiceResponseBody {
	s.OperationMetadata = &v
	return s
}

func (s *GetServiceResponseBody) SetPermission(v string) *GetServiceResponseBody {
	s.Permission = &v
	return s
}

func (s *GetServiceResponseBody) SetPolicyNames(v string) *GetServiceResponseBody {
	s.PolicyNames = &v
	return s
}

func (s *GetServiceResponseBody) SetPublishTime(v string) *GetServiceResponseBody {
	s.PublishTime = &v
	return s
}

func (s *GetServiceResponseBody) SetRequestId(v string) *GetServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceResponseBody) SetServiceDocumentInfos(v []*GetServiceResponseBodyServiceDocumentInfos) *GetServiceResponseBody {
	s.ServiceDocumentInfos = v
	return s
}

func (s *GetServiceResponseBody) SetServiceId(v string) *GetServiceResponseBody {
	s.ServiceId = &v
	return s
}

func (s *GetServiceResponseBody) SetServiceInfos(v []*GetServiceResponseBodyServiceInfos) *GetServiceResponseBody {
	s.ServiceInfos = v
	return s
}

func (s *GetServiceResponseBody) SetServiceProductUrl(v string) *GetServiceResponseBody {
	s.ServiceProductUrl = &v
	return s
}

func (s *GetServiceResponseBody) SetServiceType(v string) *GetServiceResponseBody {
	s.ServiceType = &v
	return s
}

func (s *GetServiceResponseBody) SetShareType(v string) *GetServiceResponseBody {
	s.ShareType = &v
	return s
}

func (s *GetServiceResponseBody) SetStatus(v string) *GetServiceResponseBody {
	s.Status = &v
	return s
}

func (s *GetServiceResponseBody) SetSupplierDesc(v string) *GetServiceResponseBody {
	s.SupplierDesc = &v
	return s
}

func (s *GetServiceResponseBody) SetSupplierLogo(v string) *GetServiceResponseBody {
	s.SupplierLogo = &v
	return s
}

func (s *GetServiceResponseBody) SetSupplierName(v string) *GetServiceResponseBody {
	s.SupplierName = &v
	return s
}

func (s *GetServiceResponseBody) SetSupplierUid(v int64) *GetServiceResponseBody {
	s.SupplierUid = &v
	return s
}

func (s *GetServiceResponseBody) SetSupplierUrl(v string) *GetServiceResponseBody {
	s.SupplierUrl = &v
	return s
}

func (s *GetServiceResponseBody) SetSupportContacts(v []*GetServiceResponseBodySupportContacts) *GetServiceResponseBody {
	s.SupportContacts = v
	return s
}

func (s *GetServiceResponseBody) SetTags(v []*GetServiceResponseBodyTags) *GetServiceResponseBody {
	s.Tags = v
	return s
}

func (s *GetServiceResponseBody) SetTenantType(v string) *GetServiceResponseBody {
	s.TenantType = &v
	return s
}

func (s *GetServiceResponseBody) SetTrialDuration(v int64) *GetServiceResponseBody {
	s.TrialDuration = &v
	return s
}

func (s *GetServiceResponseBody) SetTrialType(v string) *GetServiceResponseBody {
	s.TrialType = &v
	return s
}

func (s *GetServiceResponseBody) SetVersion(v string) *GetServiceResponseBody {
	s.Version = &v
	return s
}

func (s *GetServiceResponseBody) SetVersionName(v string) *GetServiceResponseBody {
	s.VersionName = &v
	return s
}

type GetServiceResponseBodyCommodity struct {
	// The billing method of the service. Valid values:
	//
	// 	- **PREPAY*	- (default): subscription.
	//
	// 	- **POSTPAY**: pay-as-you-go.
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The commodity code of the service in Alibaba Cloud Marketplace.
	//
	// example:
	//
	// cmjj00****
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The configuration metadata related to Lingxiao.
	CssMetadata *GetServiceResponseBodyCommodityCssMetadata `json:"CssMetadata,omitempty" xml:"CssMetadata,omitempty" type:"Struct"`
	// The deploy page.
	//
	// example:
	//
	// Order： Order page
	//
	// Detail： Detail page
	DeployPage *string `json:"DeployPage,omitempty" xml:"DeployPage,omitempty"`
	// The metadata of Alibaba Cloud Marketplace.
	MarketplaceMetadata *GetServiceResponseBodyCommodityMarketplaceMetadata `json:"MarketplaceMetadata,omitempty" xml:"MarketplaceMetadata,omitempty" type:"Struct"`
	// The order time.
	OrderTime map[string][]*string `json:"OrderTime,omitempty" xml:"OrderTime,omitempty"`
	// The configuration metadata related to Saas Boost.
	//
	// example:
	//
	// {
	//
	//     "Enabled":false    "PublicAccessUrl":"https://example.com"
	//
	// }
	SaasBoostMetadata *string `json:"SaasBoostMetadata,omitempty" xml:"SaasBoostMetadata,omitempty"`
	// The specification details of the service in Alibaba Cloud Marketplace.
	Specifications []*GetServiceResponseBodyCommoditySpecifications `json:"Specifications,omitempty" xml:"Specifications,omitempty" type:"Repeated"`
	// The service type. Valid values:
	//
	// 	- marketplace: Alibaba Cloud Marketplace.
	//
	// 	- Css: Lingxiao.
	//
	// example:
	//
	// Marketplace
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetServiceResponseBodyCommodity) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyCommodity) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodity) SetChargeType(v string) *GetServiceResponseBodyCommodity {
	s.ChargeType = &v
	return s
}

func (s *GetServiceResponseBodyCommodity) SetCommodityCode(v string) *GetServiceResponseBodyCommodity {
	s.CommodityCode = &v
	return s
}

func (s *GetServiceResponseBodyCommodity) SetCssMetadata(v *GetServiceResponseBodyCommodityCssMetadata) *GetServiceResponseBodyCommodity {
	s.CssMetadata = v
	return s
}

func (s *GetServiceResponseBodyCommodity) SetDeployPage(v string) *GetServiceResponseBodyCommodity {
	s.DeployPage = &v
	return s
}

func (s *GetServiceResponseBodyCommodity) SetMarketplaceMetadata(v *GetServiceResponseBodyCommodityMarketplaceMetadata) *GetServiceResponseBodyCommodity {
	s.MarketplaceMetadata = v
	return s
}

func (s *GetServiceResponseBodyCommodity) SetOrderTime(v map[string][]*string) *GetServiceResponseBodyCommodity {
	s.OrderTime = v
	return s
}

func (s *GetServiceResponseBodyCommodity) SetSaasBoostMetadata(v string) *GetServiceResponseBodyCommodity {
	s.SaasBoostMetadata = &v
	return s
}

func (s *GetServiceResponseBodyCommodity) SetSpecifications(v []*GetServiceResponseBodyCommoditySpecifications) *GetServiceResponseBodyCommodity {
	s.Specifications = v
	return s
}

func (s *GetServiceResponseBodyCommodity) SetType(v string) *GetServiceResponseBodyCommodity {
	s.Type = &v
	return s
}

type GetServiceResponseBodyCommodityCssMetadata struct {
	// The mapping information about the billing items.
	ComponentsMappings []*GetServiceResponseBodyCommodityCssMetadataComponentsMappings `json:"ComponentsMappings,omitempty" xml:"ComponentsMappings,omitempty" type:"Repeated"`
}

func (s GetServiceResponseBodyCommodityCssMetadata) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyCommodityCssMetadata) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodityCssMetadata) SetComponentsMappings(v []*GetServiceResponseBodyCommodityCssMetadataComponentsMappings) *GetServiceResponseBodyCommodityCssMetadata {
	s.ComponentsMappings = v
	return s
}

type GetServiceResponseBodyCommodityCssMetadataComponentsMappings struct {
	// The mappings.
	Mappings map[string]*string `json:"Mappings,omitempty" xml:"Mappings,omitempty"`
	// The template name.
	//
	// example:
	//
	// Template one.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GetServiceResponseBodyCommodityCssMetadataComponentsMappings) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyCommodityCssMetadataComponentsMappings) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodityCssMetadataComponentsMappings) SetMappings(v map[string]*string) *GetServiceResponseBodyCommodityCssMetadataComponentsMappings {
	s.Mappings = v
	return s
}

func (s *GetServiceResponseBodyCommodityCssMetadataComponentsMappings) SetTemplateName(v string) *GetServiceResponseBodyCommodityCssMetadataComponentsMappings {
	s.TemplateName = &v
	return s
}

type GetServiceResponseBodyCommodityMarketplaceMetadata struct {
	// The mappings between the service specifications and the template or package.
	SpecificationMappings []*GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings `json:"SpecificationMappings,omitempty" xml:"SpecificationMappings,omitempty" type:"Repeated"`
}

func (s GetServiceResponseBodyCommodityMarketplaceMetadata) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyCommodityMarketplaceMetadata) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadata) SetSpecificationMappings(v []*GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings) *GetServiceResponseBodyCommodityMarketplaceMetadata {
	s.SpecificationMappings = v
	return s
}

type GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings struct {
	// The specification code of the service in Alibaba Cloud Marketplace.
	//
	// example:
	//
	// cmjj00****
	SpecificationCode *string `json:"SpecificationCode,omitempty" xml:"SpecificationCode,omitempty"`
	// The package name.
	//
	// example:
	//
	// Package one.
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// The template name.
	//
	// example:
	//
	// Template one.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings) SetSpecificationCode(v string) *GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings {
	s.SpecificationCode = &v
	return s
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings) SetSpecificationName(v string) *GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings {
	s.SpecificationName = &v
	return s
}

func (s *GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings) SetTemplateName(v string) *GetServiceResponseBodyCommodityMarketplaceMetadataSpecificationMappings {
	s.TemplateName = &v
	return s
}

type GetServiceResponseBodyCommoditySpecifications struct {
	// The commodity code.
	//
	// example:
	//
	// cmjj00****
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The specification name.
	//
	// example:
	//
	// specifications1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The subscription duration. Unit: week or year.
	Times []*string `json:"Times,omitempty" xml:"Times,omitempty" type:"Repeated"`
}

func (s GetServiceResponseBodyCommoditySpecifications) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyCommoditySpecifications) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyCommoditySpecifications) SetCode(v string) *GetServiceResponseBodyCommoditySpecifications {
	s.Code = &v
	return s
}

func (s *GetServiceResponseBodyCommoditySpecifications) SetName(v string) *GetServiceResponseBodyCommoditySpecifications {
	s.Name = &v
	return s
}

func (s *GetServiceResponseBodyCommoditySpecifications) SetTimes(v []*string) *GetServiceResponseBodyCommoditySpecifications {
	s.Times = v
	return s
}

type GetServiceResponseBodyInstanceRoleInfos struct {
	// The content of the policy.
	//
	// example:
	//
	// {\\n  \\"Version\\": \\"1\\",\\n  \\"Statement\\": [\\n    {\\n      \\"Effect\\": \\"Allow\\",\\n      \\"Action\\": \\"*\\",\\n      \\"Principal\\": \\"*\\",\\n      \\"Resource\\": \\"*\\"\\n    }\\n  ]\\n}
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// The information of the RAM entity.
	Principals []*string `json:"Principals,omitempty" xml:"Principals,omitempty" type:"Repeated"`
	// The ram role name.
	//
	// example:
	//
	// ram-for-dts
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The template name.
	//
	// example:
	//
	// Template one.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GetServiceResponseBodyInstanceRoleInfos) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyInstanceRoleInfos) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyInstanceRoleInfos) SetPolicyDocument(v string) *GetServiceResponseBodyInstanceRoleInfos {
	s.PolicyDocument = &v
	return s
}

func (s *GetServiceResponseBodyInstanceRoleInfos) SetPrincipals(v []*string) *GetServiceResponseBodyInstanceRoleInfos {
	s.Principals = v
	return s
}

func (s *GetServiceResponseBodyInstanceRoleInfos) SetRoleName(v string) *GetServiceResponseBodyInstanceRoleInfos {
	s.RoleName = &v
	return s
}

func (s *GetServiceResponseBodyInstanceRoleInfos) SetTemplateName(v string) *GetServiceResponseBodyInstanceRoleInfos {
	s.TemplateName = &v
	return s
}

type GetServiceResponseBodyServiceDocumentInfos struct {
	// The URL that is used to access the document.
	//
	// example:
	//
	// https://help.aliyun.com/zh/compute-nest/use-cases/deploy-an-sd-painting-service-instance?spm=a2c4g.11186623.0.i2
	DocumentUrl *string `json:"DocumentUrl,omitempty" xml:"DocumentUrl,omitempty"`
	// The language that you use for the query. Valid values: zh-CN and en-US.
	//
	// example:
	//
	// zh-CN
	Locale *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	// The template name.
	//
	// example:
	//
	// Template one.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GetServiceResponseBodyServiceDocumentInfos) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyServiceDocumentInfos) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyServiceDocumentInfos) SetDocumentUrl(v string) *GetServiceResponseBodyServiceDocumentInfos {
	s.DocumentUrl = &v
	return s
}

func (s *GetServiceResponseBodyServiceDocumentInfos) SetLocale(v string) *GetServiceResponseBodyServiceDocumentInfos {
	s.Locale = &v
	return s
}

func (s *GetServiceResponseBodyServiceDocumentInfos) SetTemplateName(v string) *GetServiceResponseBodyServiceDocumentInfos {
	s.TemplateName = &v
	return s
}

type GetServiceResponseBodyServiceInfos struct {
	// The agreement information about the service.
	Agreements []*GetServiceResponseBodyServiceInfosAgreements `json:"Agreements,omitempty" xml:"Agreements,omitempty" type:"Repeated"`
	// The URL of the service icon.
	//
	// example:
	//
	// https://example.com/service-image/c1c4a559-cc60-4af1-b976-98f356602462.png
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The language of the service. Valid values:
	//
	// 	- zh-CN: Chinese
	//
	// 	- en-US: English
	//
	// example:
	//
	// zh-CN
	Locale *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	// The service name.
	//
	// example:
	//
	// Service document information.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The description of the service.
	//
	// example:
	//
	// Docker Community Edition (CE) is a free version of the Docker project, aimed at developers, enthusiasts, and individuals and organizations who want to use container technology.
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
	// The list of the software in the service.
	Softwares []*GetServiceResponseBodyServiceInfosSoftwares `json:"Softwares,omitempty" xml:"Softwares,omitempty" type:"Repeated"`
}

func (s GetServiceResponseBodyServiceInfos) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyServiceInfos) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyServiceInfos) SetAgreements(v []*GetServiceResponseBodyServiceInfosAgreements) *GetServiceResponseBodyServiceInfos {
	s.Agreements = v
	return s
}

func (s *GetServiceResponseBodyServiceInfos) SetImage(v string) *GetServiceResponseBodyServiceInfos {
	s.Image = &v
	return s
}

func (s *GetServiceResponseBodyServiceInfos) SetLocale(v string) *GetServiceResponseBodyServiceInfos {
	s.Locale = &v
	return s
}

func (s *GetServiceResponseBodyServiceInfos) SetName(v string) *GetServiceResponseBodyServiceInfos {
	s.Name = &v
	return s
}

func (s *GetServiceResponseBodyServiceInfos) SetShortDescription(v string) *GetServiceResponseBodyServiceInfos {
	s.ShortDescription = &v
	return s
}

func (s *GetServiceResponseBodyServiceInfos) SetSoftwares(v []*GetServiceResponseBodyServiceInfosSoftwares) *GetServiceResponseBodyServiceInfos {
	s.Softwares = v
	return s
}

type GetServiceResponseBodyServiceInfosAgreements struct {
	// The agreement name.
	//
	// example:
	//
	// User agreement
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The agreement URL.
	//
	// example:
	//
	// https://url
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetServiceResponseBodyServiceInfosAgreements) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyServiceInfosAgreements) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyServiceInfosAgreements) SetName(v string) *GetServiceResponseBodyServiceInfosAgreements {
	s.Name = &v
	return s
}

func (s *GetServiceResponseBodyServiceInfosAgreements) SetUrl(v string) *GetServiceResponseBodyServiceInfosAgreements {
	s.Url = &v
	return s
}

type GetServiceResponseBodyServiceInfosSoftwares struct {
	// The name of the Software.
	//
	// example:
	//
	// wordpress
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version of the software.
	//
	// example:
	//
	// 6.0.1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetServiceResponseBodyServiceInfosSoftwares) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyServiceInfosSoftwares) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyServiceInfosSoftwares) SetName(v string) *GetServiceResponseBodyServiceInfosSoftwares {
	s.Name = &v
	return s
}

func (s *GetServiceResponseBodyServiceInfosSoftwares) SetVersion(v string) *GetServiceResponseBodyServiceInfosSoftwares {
	s.Version = &v
	return s
}

type GetServiceResponseBodySupportContacts struct {
	// The type of contact information.
	//
	// example:
	//
	// Email
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of contact information.
	//
	// example:
	//
	// supplier@example.com
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetServiceResponseBodySupportContacts) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodySupportContacts) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodySupportContacts) SetType(v string) *GetServiceResponseBodySupportContacts {
	s.Type = &v
	return s
}

func (s *GetServiceResponseBodySupportContacts) SetValue(v string) *GetServiceResponseBodySupportContacts {
	s.Value = &v
	return s
}

type GetServiceResponseBodyTags struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetServiceResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetServiceResponseBodyTags) SetKey(v string) *GetServiceResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetServiceResponseBodyTags) SetValue(v string) *GetServiceResponseBodyTags {
	s.Value = &v
	return s
}

type GetServiceResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceResponse) GoString() string {
	return s.String()
}

func (s *GetServiceResponse) SetHeaders(v map[string]*string) *GetServiceResponse {
	s.Headers = v
	return s
}

func (s *GetServiceResponse) SetStatusCode(v int32) *GetServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceResponse) SetBody(v *GetServiceResponseBody) *GetServiceResponse {
	s.Body = v
	return s
}

type GetServiceEstimateCostRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// qwertyuiop
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The information about the subscription duration.
	Commodity *GetServiceEstimateCostRequestCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// The name of the configuration change operation.
	//
	// example:
	//
	// Parameter change
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	// The parameters that are specified to deploy the service instance.
	//
	// example:
	//
	// { \\"RegionId\\": \\"cn-hangzhou\\", \\"InstanceType\\": \\"ecs.g5.large\\"}
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-12xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service instance ID.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17xxxx
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The name of the package specification.
	//
	// example:
	//
	// Package 1
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// The template name.
	//
	// example:
	//
	// Template 1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The trial policy. Valid values:
	//
	// 	- Trial: Trials are supported.
	//
	// 	- NotTrial: Trials are not supported.
	//
	// example:
	//
	// NotTrial
	TrialType *string `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
}

func (s GetServiceEstimateCostRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceEstimateCostRequest) GoString() string {
	return s.String()
}

func (s *GetServiceEstimateCostRequest) SetClientToken(v string) *GetServiceEstimateCostRequest {
	s.ClientToken = &v
	return s
}

func (s *GetServiceEstimateCostRequest) SetCommodity(v *GetServiceEstimateCostRequestCommodity) *GetServiceEstimateCostRequest {
	s.Commodity = v
	return s
}

func (s *GetServiceEstimateCostRequest) SetOperationName(v string) *GetServiceEstimateCostRequest {
	s.OperationName = &v
	return s
}

func (s *GetServiceEstimateCostRequest) SetParameters(v map[string]interface{}) *GetServiceEstimateCostRequest {
	s.Parameters = v
	return s
}

func (s *GetServiceEstimateCostRequest) SetRegionId(v string) *GetServiceEstimateCostRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceEstimateCostRequest) SetServiceId(v string) *GetServiceEstimateCostRequest {
	s.ServiceId = &v
	return s
}

func (s *GetServiceEstimateCostRequest) SetServiceInstanceId(v string) *GetServiceEstimateCostRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *GetServiceEstimateCostRequest) SetServiceVersion(v string) *GetServiceEstimateCostRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GetServiceEstimateCostRequest) SetSpecificationName(v string) *GetServiceEstimateCostRequest {
	s.SpecificationName = &v
	return s
}

func (s *GetServiceEstimateCostRequest) SetTemplateName(v string) *GetServiceEstimateCostRequest {
	s.TemplateName = &v
	return s
}

func (s *GetServiceEstimateCostRequest) SetTrialType(v string) *GetServiceEstimateCostRequest {
	s.TrialType = &v
	return s
}

type GetServiceEstimateCostRequestCommodity struct {
	// 优惠券ID
	//
	// example:
	//
	// 302070970220
	CouponId *string `json:"CouponId,omitempty" xml:"CouponId,omitempty"`
	// The subscription duration.
	//
	// example:
	//
	// 1
	PayPeriod *int32 `json:"PayPeriod,omitempty" xml:"PayPeriod,omitempty"`
	// The unit of the subscription duration. Valid values:
	//
	// 	- Year
	//
	// 	- Month
	//
	// 	- Day
	//
	// example:
	//
	// Year
	PayPeriodUnit *string `json:"PayPeriodUnit,omitempty" xml:"PayPeriodUnit,omitempty"`
}

func (s GetServiceEstimateCostRequestCommodity) String() string {
	return tea.Prettify(s)
}

func (s GetServiceEstimateCostRequestCommodity) GoString() string {
	return s.String()
}

func (s *GetServiceEstimateCostRequestCommodity) SetCouponId(v string) *GetServiceEstimateCostRequestCommodity {
	s.CouponId = &v
	return s
}

func (s *GetServiceEstimateCostRequestCommodity) SetPayPeriod(v int32) *GetServiceEstimateCostRequestCommodity {
	s.PayPeriod = &v
	return s
}

func (s *GetServiceEstimateCostRequestCommodity) SetPayPeriodUnit(v string) *GetServiceEstimateCostRequestCommodity {
	s.PayPeriodUnit = &v
	return s
}

type GetServiceEstimateCostShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// qwertyuiop
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The information about the subscription duration.
	CommodityShrink *string `json:"Commodity,omitempty" xml:"Commodity,omitempty"`
	// The name of the configuration change operation.
	//
	// example:
	//
	// Parameter change
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	// The parameters that are specified to deploy the service instance.
	//
	// example:
	//
	// { \\"RegionId\\": \\"cn-hangzhou\\", \\"InstanceType\\": \\"ecs.g5.large\\"}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-12xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service instance ID.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17xxxx
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The name of the package specification.
	//
	// example:
	//
	// Package 1
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// The template name.
	//
	// example:
	//
	// Template 1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The trial policy. Valid values:
	//
	// 	- Trial: Trials are supported.
	//
	// 	- NotTrial: Trials are not supported.
	//
	// example:
	//
	// NotTrial
	TrialType *string `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
}

func (s GetServiceEstimateCostShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceEstimateCostShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetServiceEstimateCostShrinkRequest) SetClientToken(v string) *GetServiceEstimateCostShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetCommodityShrink(v string) *GetServiceEstimateCostShrinkRequest {
	s.CommodityShrink = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetOperationName(v string) *GetServiceEstimateCostShrinkRequest {
	s.OperationName = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetParametersShrink(v string) *GetServiceEstimateCostShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetRegionId(v string) *GetServiceEstimateCostShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetServiceId(v string) *GetServiceEstimateCostShrinkRequest {
	s.ServiceId = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetServiceInstanceId(v string) *GetServiceEstimateCostShrinkRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetServiceVersion(v string) *GetServiceEstimateCostShrinkRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetSpecificationName(v string) *GetServiceEstimateCostShrinkRequest {
	s.SpecificationName = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetTemplateName(v string) *GetServiceEstimateCostShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *GetServiceEstimateCostShrinkRequest) SetTrialType(v string) *GetServiceEstimateCostShrinkRequest {
	s.TrialType = &v
	return s
}

type GetServiceEstimateCostResponseBody struct {
	// The estimated price.
	//
	// example:
	//
	// {\\"cmgj00059839\\": {\\"Result\\": {\\"InquiryType\\": \\"Buy\\", \\"Order\\": {\\"Currency\\": \\"CNY\\", \\"DiscountAmount\\": \\"0.0\\", \\"TradeAmount\\": \\"0.01\\", \\"OriginalAmount\\": \\"0.01\\"}}}}
	Commodity map[string]*CommodityValue `json:"Commodity,omitempty" xml:"Commodity,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 08ABBB67-39C9-5EE7-98E5-80486F75CE8D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resources.
	//
	// example:
	//
	// { "EcsInstance" : { "Type" : "ALIYUN::ECS::Instance", "Success" : true, "Result" : { "Order" : { "Currency" : "CNY", "RuleIds" : [ "102111101338\\*\\*\\*\\*" ], "DetailInfos" : { "ResourcePriceModel" : [ { "OriginalPrice" : 0, "DiscountPrice" : 0, "SubRules" : { "Rule" : [ ] }, "Resource" : "bandwidth", "TradePrice" : 0 }, { "OriginalPrice" : 0, "DiscountPrice" : 0, "SubRules" : { "Rule" : [ ] }, "Resource" : "image", "TradePrice" : 0 }, { "OriginalPrice" : 0.366666, "DiscountPrice" : 0.249012, "SubRules" : { "Rule" : [ ] }, "Resource" : "instanceType", "TradePrice" : 0.117654 }, { "OriginalPrice" : 0.055555, "DiscountPrice" : 0.037729, "SubRules" : { "Rule" : [ ] }, "Resource" : "systemDisk", "TradePrice" : 0.017826 } ] }, "TradeAmount" : 0.135, "OriginalAmount" : 0.422, "Coupons" : { "Coupon" : [ ] }, "DiscountAmount" : 0.287 }, "OrderSupplement" : { "PriceUnit" : "/Hour", "ChargeType" : "PostPaid", "Quantity" : 1, "PriceType" : "Total" }, "Rules" : { "Rule" : [ { "RuleDescId" : "102111101338\\*\\*\\*\\*", "Name" : "contract discount_multi-billing item discount_3.208750 discount" } ] } } } }
	Resources map[string]interface{} `json:"Resources,omitempty" xml:"Resources,omitempty"`
}

func (s GetServiceEstimateCostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceEstimateCostResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceEstimateCostResponseBody) SetCommodity(v map[string]*CommodityValue) *GetServiceEstimateCostResponseBody {
	s.Commodity = v
	return s
}

func (s *GetServiceEstimateCostResponseBody) SetRequestId(v string) *GetServiceEstimateCostResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceEstimateCostResponseBody) SetResources(v map[string]interface{}) *GetServiceEstimateCostResponseBody {
	s.Resources = v
	return s
}

type GetServiceEstimateCostResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceEstimateCostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceEstimateCostResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceEstimateCostResponse) GoString() string {
	return s.String()
}

func (s *GetServiceEstimateCostResponse) SetHeaders(v map[string]*string) *GetServiceEstimateCostResponse {
	s.Headers = v
	return s
}

func (s *GetServiceEstimateCostResponse) SetStatusCode(v int32) *GetServiceEstimateCostResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceEstimateCostResponse) SetBody(v *GetServiceEstimateCostResponseBody) *GetServiceEstimateCostResponse {
	s.Body = v
	return s
}

type GetServiceInstanceRequest struct {
	// The MarketInstance ID.
	//
	// example:
	//
	// 704***59
	MarketInstanceId *string `json:"MarketInstanceId,omitempty" xml:"MarketInstanceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service instance ID.
	//
	// >  You must specify either `ServiceInstanceId` or `MarketInstanceId`. Otherwise, the operation fails.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s GetServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceRequest) SetMarketInstanceId(v string) *GetServiceInstanceRequest {
	s.MarketInstanceId = &v
	return s
}

func (s *GetServiceInstanceRequest) SetRegionId(v string) *GetServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceInstanceRequest) SetServiceInstanceId(v string) *GetServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

type GetServiceInstanceResponseBody struct {
	// The business state of the service instance. Valid values:
	//
	// 	- Normal
	//
	// 	- Renewing
	//
	// 	- RenewFailed
	//
	// 	- Expired
	//
	// example:
	//
	// Normal
	BizStatus *string `json:"BizStatus,omitempty" xml:"BizStatus,omitempty"`
	// Cloud Marketplace additional billing items.
	//
	// example:
	//
	// {"TiKVServerCount":"3","package_version":"yuncode5398300001","PDServerCount":"3","TiDBServerCount":"2"}
	Components *string `json:"Components,omitempty" xml:"Components,omitempty"`
	// The time when the serviceInstance was created.
	//
	// example:
	//
	// 2021-05-20T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the service instance supports the operation feature. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableInstanceOps *bool `json:"EnableInstanceOps,omitempty" xml:"EnableInstanceOps,omitempty"`
	// Whether to enable Prometheus monitoring.
	//
	// example:
	//
	// true
	EnableUserPrometheus *bool `json:"EnableUserPrometheus,omitempty" xml:"EnableUserPrometheus,omitempty"`
	// The expiration time of service instance.
	//
	// example:
	//
	// 2022-01-01T12:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The URL of the Grafana dashboard.
	//
	// example:
	//
	// https://g.console.aliyun.com/d/xxxxxxxx-cn-mariadb/mysql-xxxxxx-xxxxxxxx-and-dashboard?orgId=355401&refresh=10s
	GrafanaDashBoardUrl *string `json:"GrafanaDashBoardUrl,omitempty" xml:"GrafanaDashBoardUrl,omitempty"`
	// Indicates whether the hosted O\\&M feature is enabled for the service instance. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsOperated *bool `json:"IsOperated,omitempty" xml:"IsOperated,omitempty"`
	// The expiration time of licence.
	//
	// example:
	//
	// 2022-01-01T12:00:00
	LicenseEndTime *string `json:"LicenseEndTime,omitempty" xml:"LicenseEndTime,omitempty"`
	// The market Instance ID.
	//
	// example:
	//
	// 704***59
	MarketInstanceId *string `json:"MarketInstanceId,omitempty" xml:"MarketInstanceId,omitempty"`
	// The name of the service instance.
	//
	// example:
	//
	// TestName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The network configurations.
	//
	// >  This parameter is discontinued.
	NetworkConfig *GetServiceInstanceResponseBodyNetworkConfig `json:"NetworkConfig,omitempty" xml:"NetworkConfig,omitempty" type:"Struct"`
	// The serviceInstance  to be operated.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	OperatedServiceInstanceId *string `json:"OperatedServiceInstanceId,omitempty" xml:"OperatedServiceInstanceId,omitempty"`
	// The operation end time.
	//
	// example:
	//
	// 2022-01-28T06:48:56Z
	OperationEndTime *string `json:"OperationEndTime,omitempty" xml:"OperationEndTime,omitempty"`
	// The operation start time.
	//
	// example:
	//
	// 2021-12-29T06:48:56Z
	OperationStartTime *string `json:"OperationStartTime,omitempty" xml:"OperationStartTime,omitempty"`
	// The outputs returned from creating the service instance.
	//
	// 	- If the service is deployed by using a ROS template, all output fields of the template are returned.
	//
	// 	- If the service is deployed by calling an SPI operation, the output fields of the service provider and for the Compute Nest additional features are returned.
	//
	// example:
	//
	// {"InstanceIds":["i-hp38ofxl0dsyfi7z****"]}
	Outputs *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// The parameters configured for the service instance.
	//
	// example:
	//
	// {"param":"value"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The billing method of the instance for market. Valid values:
	//
	// 	- Permanent: Permanent purchase
	//
	// 	- Subscription: Subscription.
	//
	// 	- PayAsYouGo: Pay-as-you-go.
	//
	// 	- CustomFixTime: Merchant custom fixed duration.
	//
	// example:
	//
	// Subscription
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The package name.
	//
	// example:
	//
	// Package one
	PredefinedParameterName *string `json:"PredefinedParameterName,omitempty" xml:"PredefinedParameterName,omitempty"`
	// The deployment progress of the service instance. Unit: percentage.
	//
	// example:
	//
	// 90
	Progress *int64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resources.
	//
	// example:
	//
	// [{"StackId": "stack-xxx"}]
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// The information about the service to which the service instance belongs.
	Service *GetServiceInstanceResponseBodyService `json:"Service,omitempty" xml:"Service,omitempty" type:"Struct"`
	// The ID of the service instance.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The type of the service. Valid values:
	//
	// - private: The service is a private service and is deployed within the account of a customer.
	//
	// - managed: The service is a fully managed service and is deployed within the account of a service provider.
	//
	// - operation: The service is a hosted O&M service.
	//
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The source of the serviceInstance. Valid values:
	//
	// - User
	//
	// - Market
	//
	// - Supplier
	//
	// example:
	//
	// User
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The deploy status of the serviceInstance. Valid values:
	//
	// - Created
	//
	// - Deploying
	//
	// - DeployedFailed
	//
	// - Deployed
	//
	// - Upgrading
	//
	// - Deleting
	//
	// - Deleted
	//
	// - DeletedFailed
	//
	// example:
	//
	// Deployed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The description of the deployment state of the service instance.
	//
	// example:
	//
	// deploy successfully
	StatusDetail *string `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
	// The Alibaba Cloud account ID of the service provider.
	//
	// example:
	//
	// 158927391332*****
	SupplierUid *int64 `json:"SupplierUid,omitempty" xml:"SupplierUid,omitempty"`
	// The tags.
	Tags []*GetServiceInstanceResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The template name.
	//
	// example:
	//
	// Template one
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The time when the serviceInstance  was last updated.
	//
	// example:
	//
	// 2021-05-20T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The AliUid of the user.
	//
	// example:
	//
	// 130920852836****
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBody) SetBizStatus(v string) *GetServiceInstanceResponseBody {
	s.BizStatus = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetComponents(v string) *GetServiceInstanceResponseBody {
	s.Components = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetCreateTime(v string) *GetServiceInstanceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetEnableInstanceOps(v bool) *GetServiceInstanceResponseBody {
	s.EnableInstanceOps = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetEnableUserPrometheus(v bool) *GetServiceInstanceResponseBody {
	s.EnableUserPrometheus = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetEndTime(v string) *GetServiceInstanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetGrafanaDashBoardUrl(v string) *GetServiceInstanceResponseBody {
	s.GrafanaDashBoardUrl = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetIsOperated(v bool) *GetServiceInstanceResponseBody {
	s.IsOperated = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetLicenseEndTime(v string) *GetServiceInstanceResponseBody {
	s.LicenseEndTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetMarketInstanceId(v string) *GetServiceInstanceResponseBody {
	s.MarketInstanceId = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetName(v string) *GetServiceInstanceResponseBody {
	s.Name = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetNetworkConfig(v *GetServiceInstanceResponseBodyNetworkConfig) *GetServiceInstanceResponseBody {
	s.NetworkConfig = v
	return s
}

func (s *GetServiceInstanceResponseBody) SetOperatedServiceInstanceId(v string) *GetServiceInstanceResponseBody {
	s.OperatedServiceInstanceId = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetOperationEndTime(v string) *GetServiceInstanceResponseBody {
	s.OperationEndTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetOperationStartTime(v string) *GetServiceInstanceResponseBody {
	s.OperationStartTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetOutputs(v string) *GetServiceInstanceResponseBody {
	s.Outputs = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetParameters(v string) *GetServiceInstanceResponseBody {
	s.Parameters = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetPayType(v string) *GetServiceInstanceResponseBody {
	s.PayType = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetPredefinedParameterName(v string) *GetServiceInstanceResponseBody {
	s.PredefinedParameterName = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetProgress(v int64) *GetServiceInstanceResponseBody {
	s.Progress = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetRequestId(v string) *GetServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetResourceGroupId(v string) *GetServiceInstanceResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetResources(v string) *GetServiceInstanceResponseBody {
	s.Resources = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetService(v *GetServiceInstanceResponseBodyService) *GetServiceInstanceResponseBody {
	s.Service = v
	return s
}

func (s *GetServiceInstanceResponseBody) SetServiceInstanceId(v string) *GetServiceInstanceResponseBody {
	s.ServiceInstanceId = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetServiceType(v string) *GetServiceInstanceResponseBody {
	s.ServiceType = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetSource(v string) *GetServiceInstanceResponseBody {
	s.Source = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetStatus(v string) *GetServiceInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetStatusDetail(v string) *GetServiceInstanceResponseBody {
	s.StatusDetail = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetSupplierUid(v int64) *GetServiceInstanceResponseBody {
	s.SupplierUid = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetTags(v []*GetServiceInstanceResponseBodyTags) *GetServiceInstanceResponseBody {
	s.Tags = v
	return s
}

func (s *GetServiceInstanceResponseBody) SetTemplateName(v string) *GetServiceInstanceResponseBody {
	s.TemplateName = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetUpdateTime(v string) *GetServiceInstanceResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetUserId(v int64) *GetServiceInstanceResponseBody {
	s.UserId = &v
	return s
}

type GetServiceInstanceResponseBodyNetworkConfig struct {
	// The ID of the endpoint for the private connection.
	//
	// >  This parameter is discontinued.
	//
	// example:
	//
	// ep-m5ei37240541816b****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The information about private connections.
	PrivateVpcConnections []*GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections `json:"PrivateVpcConnections,omitempty" xml:"PrivateVpcConnections,omitempty" type:"Repeated"`
	// The PrivateZone ID.
	//
	// example:
	//
	// cb7f214f80ac348d87daaeac1f35****
	PrivateZoneId *string `json:"PrivateZoneId,omitempty" xml:"PrivateZoneId,omitempty"`
	// The information about the reverse private connection.
	ReversePrivateVpcConnections []*GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections `json:"ReversePrivateVpcConnections,omitempty" xml:"ReversePrivateVpcConnections,omitempty" type:"Repeated"`
}

func (s GetServiceInstanceResponseBodyNetworkConfig) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyNetworkConfig) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) SetEndpointId(v string) *GetServiceInstanceResponseBodyNetworkConfig {
	s.EndpointId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) SetPrivateVpcConnections(v []*GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) *GetServiceInstanceResponseBodyNetworkConfig {
	s.PrivateVpcConnections = v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) SetPrivateZoneId(v string) *GetServiceInstanceResponseBodyNetworkConfig {
	s.PrivateZoneId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) SetReversePrivateVpcConnections(v []*GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) *GetServiceInstanceResponseBodyNetworkConfig {
	s.ReversePrivateVpcConnections = v
	return s
}

type GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections struct {
	// The network configurations, which are mainly used for private connections.
	ConnectionConfigs []*GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs `json:"ConnectionConfigs,omitempty" xml:"ConnectionConfigs,omitempty" type:"Repeated"`
	// The endpoint ID of the private connection.
	//
	// example:
	//
	// ep-m5ei37240541816b****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The ID of the private zone of the custom private domain name.
	//
	// example:
	//
	// cb7f214f80ac348d87daaeac1f35****
	PrivateZoneId *string `json:"PrivateZoneId,omitempty" xml:"PrivateZoneId,omitempty"`
	// The custom domain name.
	//
	// example:
	//
	// test.computenest.aliyuncs.com
	PrivateZoneName *string `json:"PrivateZoneName,omitempty" xml:"PrivateZoneName,omitempty"`
	// The region ID of the endpoint of the PrivateLink connection.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) SetConnectionConfigs(v []*GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	s.ConnectionConfigs = v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) SetEndpointId(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	s.EndpointId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) SetPrivateZoneId(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	s.PrivateZoneId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) SetPrivateZoneName(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	s.PrivateZoneName = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) SetRegionId(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	s.RegionId = &v
	return s
}

type GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs struct {
	// The bandwidth limit for the private connection established based on the private network interconnection mode of Compute Nest.
	//
	// example:
	//
	// 1536Mbps
	ConnectBandwidth *int32 `json:"ConnectBandwidth,omitempty" xml:"ConnectBandwidth,omitempty"`
	// The domain name.
	//
	// example:
	//
	// ie-569a9be34f5534f6bc6559b5c1xxxxxx.service-51f80502802e48xxxxxx.cn-hangzhou.computenest.aliyuncs.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The IP addresses of the endpoints of the private connections.
	EndpointIps []*string `json:"EndpointIps,omitempty" xml:"EndpointIps,omitempty" type:"Repeated"`
	// The state of the ingress endpoint. Valid values:
	//
	// 	- Ready: The ingress endpoint is connected.
	//
	// 	- Pending: The ingress endpoint is being connected.
	//
	// 	- Failed: The ingress endpoint fails to be connected.
	//
	// 	- Deleted: The ingress endpoint is deleted.
	//
	// 	- Deleting: The ingress endpoint is being deleted.
	//
	// example:
	//
	// Ready
	IngressEndpointStatus *string `json:"IngressEndpointStatus,omitempty" xml:"IngressEndpointStatus,omitempty"`
	// The state of the network service. Valid values:
	//
	// 	- Ready: The network service is connected.
	//
	// 	- Pending: The network service is being connected.
	//
	// 	- Failed: The network service fails to be connected.
	//
	// 	- Deleted: The network service is deleted.
	//
	// 	- Deleting: The network service is being deleted.
	//
	// example:
	//
	// Ready
	NetworkServiceStatus *string `json:"NetworkServiceStatus,omitempty" xml:"NetworkServiceStatus,omitempty"`
	// The region ID of the VPC to which the endpoint of the private connection established based on the private network interconnection mode of Compute Nest belongs.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The names of the security groups.
	SecurityGroups []*string `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty" type:"Repeated"`
	// The names of the vSwitches.
	VSwitches []*string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-bp1vwnn14rqpyiczj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetConnectBandwidth(v int32) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.ConnectBandwidth = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetDomainName(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.DomainName = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetEndpointIps(v []*string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.EndpointIps = v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetIngressEndpointStatus(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.IngressEndpointStatus = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetNetworkServiceStatus(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.NetworkServiceStatus = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetRegionId(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.RegionId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetSecurityGroups(v []*string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.SecurityGroups = v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetVSwitches(v []*string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.VSwitches = v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetVpcId(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.VpcId = &v
	return s
}

type GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections struct {
	// The endpoint ID of the reverse private connection.
	//
	// example:
	//
	// ep-m5ei42370541816b****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
}

func (s GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) SetEndpointId(v string) *GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections {
	s.EndpointId = &v
	return s
}

type GetServiceInstanceResponseBodyService struct {
	// The storage configurations of the service. The format in which the deployment information of a service is stored varies based on the deployment type of the service. In this case, the deployment information is stored in the JSON string format.
	DeployMetadata *string `json:"DeployMetadata,omitempty" xml:"DeployMetadata,omitempty"`
	// The deployment type of the service. Valid values:
	//
	// 	- ros: The service is deployed by using Resource Orchestration Service (ROS).
	//
	// 	- terraform: The service is deployed by using Terraform.
	//
	// 	- ack: The service is deployed by using Container Service for Kubernetes (ACK).
	//
	// 	- spi: The service is deployed by calling a service provider interface (SPI).
	//
	// 	- operation: The service is deployed by using a hosted O\\&M service.
	//
	// example:
	//
	// ros
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// The time when the service version was published.
	//
	// example:
	//
	// 2021-05-21T00:00:00Z
	PublishTime *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	// The URL of the service documentation.
	//
	// example:
	//
	// http://example.com
	ServiceDocUrl *string `json:"ServiceDocUrl,omitempty" xml:"ServiceDocUrl,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-9c8a3522528b4fe8****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The information about the service.
	ServiceInfos []*GetServiceInstanceResponseBodyServiceServiceInfos `json:"ServiceInfos,omitempty" xml:"ServiceInfos,omitempty" type:"Repeated"`
	// The URL of the service page.
	//
	// example:
	//
	// https://service-info-private.oss-cn-hangzhou.aliyuncs.com
	ServiceProductUrl *string `json:"ServiceProductUrl,omitempty" xml:"ServiceProductUrl,omitempty"`
	// The type of the service. Valid values:
	//
	// 	- private: The service is a private service and is deployed within the account of a customer.
	//
	// 	- managed: The service is a fully managed service and is deployed within the account of a service provider.
	//
	// 	- operation: The service is a hosted O\\&M service.
	//
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The status of the service. Valid values:
	//
	// 	- Draft
	//
	// 	- Submited
	//
	// 	- Approved
	//
	// 	- Online
	//
	// 	- Offline
	//
	// 	- Deleted
	//
	// 	- Launching
	//
	// 	- Beta
	//
	// example:
	//
	// Online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the service provider.
	//
	// example:
	//
	// Alibaba Cloud
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// The URL of the service provider.
	//
	// example:
	//
	// http://example.com
	SupplierUrl *string `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	// The upgradable service version.
	UpgradableServiceInfos []*GetServiceInstanceResponseBodyServiceUpgradableServiceInfos `json:"UpgradableServiceInfos,omitempty" xml:"UpgradableServiceInfos,omitempty" type:"Repeated"`
	// The service version that can be updated.
	UpgradableServiceVersions []*string `json:"UpgradableServiceVersions,omitempty" xml:"UpgradableServiceVersions,omitempty" type:"Repeated"`
	// The metadata about the upgrade.
	//
	// example:
	//
	// {
	//
	//   "Type": "OOS",
	//
	//   "Description": "Changelog or something description",
	//
	//   "SupportUpgradeFromVersions": [1, 2],
	//
	//   "UpgradeSteps": {
	//
	//     "PreUpgradeStage": {
	//
	//       "Description": "初始化数据库",
	//
	//       "Type": "RunCommand",
	//
	//       "ResourceName": "EcsRole1",
	//
	//       "CommandType": "runShellScript",
	//
	//       "CommandContent": "echo hello"
	//
	//     },
	//
	//     "UpgradeStage": [{
	//
	//       "Description": "更新EcsRole1实例",
	//
	//       "Type": "RunCommand",
	//
	//       "ResourceName": "EcsRole1",
	//
	//       "ArtifactsDownloadDirectory": "/home/admin",
	//
	//       "CommandType": "runShellScript",
	//
	//       "CommandContent": "echo hello"
	//
	//     }],
	//
	//     "PostUpgradeStage": {
	//
	//       "Description": "部署后post check",
	//
	//       "Type": "None/RunCommand",
	//
	//       "ResourceName": "EcsRole1",
	//
	//       "CommandType": "runShellScript",
	//
	//       "CommandContent": "echo hello"
	//
	//     }
	//
	//   }
	//
	// }
	UpgradeMetadata *string `json:"UpgradeMetadata,omitempty" xml:"UpgradeMetadata,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The custom version name defined by the service provider.
	//
	// example:
	//
	// 1.1.0
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetServiceInstanceResponseBodyService) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyService) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyService) SetDeployMetadata(v string) *GetServiceInstanceResponseBodyService {
	s.DeployMetadata = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetDeployType(v string) *GetServiceInstanceResponseBodyService {
	s.DeployType = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetPublishTime(v string) *GetServiceInstanceResponseBodyService {
	s.PublishTime = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceDocUrl(v string) *GetServiceInstanceResponseBodyService {
	s.ServiceDocUrl = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceId(v string) *GetServiceInstanceResponseBodyService {
	s.ServiceId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceInfos(v []*GetServiceInstanceResponseBodyServiceServiceInfos) *GetServiceInstanceResponseBodyService {
	s.ServiceInfos = v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceProductUrl(v string) *GetServiceInstanceResponseBodyService {
	s.ServiceProductUrl = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceType(v string) *GetServiceInstanceResponseBodyService {
	s.ServiceType = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetStatus(v string) *GetServiceInstanceResponseBodyService {
	s.Status = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetSupplierName(v string) *GetServiceInstanceResponseBodyService {
	s.SupplierName = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetSupplierUrl(v string) *GetServiceInstanceResponseBodyService {
	s.SupplierUrl = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetUpgradableServiceInfos(v []*GetServiceInstanceResponseBodyServiceUpgradableServiceInfos) *GetServiceInstanceResponseBodyService {
	s.UpgradableServiceInfos = v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetUpgradableServiceVersions(v []*string) *GetServiceInstanceResponseBodyService {
	s.UpgradableServiceVersions = v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetUpgradeMetadata(v string) *GetServiceInstanceResponseBodyService {
	s.UpgradeMetadata = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetVersion(v string) *GetServiceInstanceResponseBodyService {
	s.Version = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetVersionName(v string) *GetServiceInstanceResponseBodyService {
	s.VersionName = &v
	return s
}

type GetServiceInstanceResponseBodyServiceServiceInfos struct {
	// The URL of the service icon.
	//
	// example:
	//
	// https://example.com/service-image/c1c4a559-cc60-4af1-b976-98f356602462.png
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The language of the service instance.
	//
	// example:
	//
	// zh-CN
	Locale *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// Docker Community Edition
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The description of the service.
	//
	// example:
	//
	// Docker Community Edition (CE) is a free version of the Docker project, aimed at developers, enthusiasts, and individuals and organizations who want to use container technology.
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
}

func (s GetServiceInstanceResponseBodyServiceServiceInfos) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyServiceServiceInfos) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) SetImage(v string) *GetServiceInstanceResponseBodyServiceServiceInfos {
	s.Image = &v
	return s
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) SetLocale(v string) *GetServiceInstanceResponseBodyServiceServiceInfos {
	s.Locale = &v
	return s
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) SetName(v string) *GetServiceInstanceResponseBodyServiceServiceInfos {
	s.Name = &v
	return s
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) SetShortDescription(v string) *GetServiceInstanceResponseBodyServiceServiceInfos {
	s.ShortDescription = &v
	return s
}

type GetServiceInstanceResponseBodyServiceUpgradableServiceInfos struct {
	// An upgradable service version.
	//
	// example:
	//
	// draft
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The version name of an upgradable service version.
	//
	// example:
	//
	// 0.1.0
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetServiceInstanceResponseBodyServiceUpgradableServiceInfos) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyServiceUpgradableServiceInfos) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyServiceUpgradableServiceInfos) SetVersion(v string) *GetServiceInstanceResponseBodyServiceUpgradableServiceInfos {
	s.Version = &v
	return s
}

func (s *GetServiceInstanceResponseBodyServiceUpgradableServiceInfos) SetVersionName(v string) *GetServiceInstanceResponseBodyServiceUpgradableServiceInfos {
	s.VersionName = &v
	return s
}

type GetServiceInstanceResponseBodyTags struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetServiceInstanceResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyTags) SetKey(v string) *GetServiceInstanceResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetServiceInstanceResponseBodyTags) SetValue(v string) *GetServiceInstanceResponseBodyTags {
	s.Value = &v
	return s
}

type GetServiceInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponse) SetHeaders(v map[string]*string) *GetServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetServiceInstanceResponse) SetStatusCode(v int32) *GetServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceInstanceResponse) SetBody(v *GetServiceInstanceResponseBody) *GetServiceInstanceResponse {
	s.Body = v
	return s
}

type GetServiceInstanceSubscriptionEstimateCostRequest struct {
	// Ensures idempotence of the request. Generate a parameter value from your client to ensure its uniqueness across different requests. **ClientToken*	- supports only ASCII characters and cannot exceed 64 characters.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Order type. Possible value: Renewal.
	//
	// This parameter is required.
	//
	// example:
	//
	// Renewal
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The renewal duration for all prepaid resources of the service instance. The unit is specified by PeriodUnit.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The time unit for the renewal duration of all prepaid resources of the service instance, which is the unit of the Period parameter. Valid values: Month, Year. Default value: Month.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource renewal configuration.
	ResourcePeriod []*GetServiceInstanceSubscriptionEstimateCostRequestResourcePeriod `json:"ResourcePeriod,omitempty" xml:"ResourcePeriod,omitempty" type:"Repeated"`
	// Service instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s GetServiceInstanceSubscriptionEstimateCostRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceSubscriptionEstimateCostRequest) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceSubscriptionEstimateCostRequest) SetClientToken(v string) *GetServiceInstanceSubscriptionEstimateCostRequest {
	s.ClientToken = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostRequest) SetOrderType(v string) *GetServiceInstanceSubscriptionEstimateCostRequest {
	s.OrderType = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostRequest) SetPeriod(v int32) *GetServiceInstanceSubscriptionEstimateCostRequest {
	s.Period = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostRequest) SetPeriodUnit(v string) *GetServiceInstanceSubscriptionEstimateCostRequest {
	s.PeriodUnit = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostRequest) SetRegionId(v string) *GetServiceInstanceSubscriptionEstimateCostRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostRequest) SetResourcePeriod(v []*GetServiceInstanceSubscriptionEstimateCostRequestResourcePeriod) *GetServiceInstanceSubscriptionEstimateCostRequest {
	s.ResourcePeriod = v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostRequest) SetServiceInstanceId(v string) *GetServiceInstanceSubscriptionEstimateCostRequest {
	s.ServiceInstanceId = &v
	return s
}

type GetServiceInstanceSubscriptionEstimateCostRequestResourcePeriod struct {
	// Renewal duration. The unit is specified by PeriodUnit.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The time unit for the resource renewal duration, which is the unit of the Period parameter. Valid values: Month, Year. Default value: Month.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// Resource ARN (Aliyun Resource Name).
	//
	// example:
	//
	// acs:ecs:cn-guangzhou:1361753504587228:instance/i-7xv9pgeqvhxg10jji3vd
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
}

func (s GetServiceInstanceSubscriptionEstimateCostRequestResourcePeriod) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceSubscriptionEstimateCostRequestResourcePeriod) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceSubscriptionEstimateCostRequestResourcePeriod) SetPeriod(v int32) *GetServiceInstanceSubscriptionEstimateCostRequestResourcePeriod {
	s.Period = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostRequestResourcePeriod) SetPeriodUnit(v string) *GetServiceInstanceSubscriptionEstimateCostRequestResourcePeriod {
	s.PeriodUnit = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostRequestResourcePeriod) SetResourceArn(v string) *GetServiceInstanceSubscriptionEstimateCostRequestResourcePeriod {
	s.ResourceArn = &v
	return s
}

type GetServiceInstanceSubscriptionEstimateCostResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 464C8CB6-A548-5206-B83C-D32A8E43EC21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// List of resource price information.
	ResourcePrices []*GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices `json:"ResourcePrices,omitempty" xml:"ResourcePrices,omitempty" type:"Repeated"`
}

func (s GetServiceInstanceSubscriptionEstimateCostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceSubscriptionEstimateCostResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBody) SetRequestId(v string) *GetServiceInstanceSubscriptionEstimateCostResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBody) SetResourcePrices(v []*GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) *GetServiceInstanceSubscriptionEstimateCostResponseBody {
	s.ResourcePrices = v
	return s
}

type GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices struct {
	// Currency. Valid values:
	//
	// - CNY: Chinese Yuan.
	//
	// - USD: US Dollar.
	//
	// - JPY: Japanese Yen.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The price details of the pricing module.
	DetailInfos []*GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesDetailInfos `json:"DetailInfos,omitempty" xml:"DetailInfos,omitempty" type:"Repeated"`
	// Discount.
	//
	// example:
	//
	// 100
	DiscountAmount *float32 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// Original price.
	//
	// example:
	//
	// 900
	OriginalAmount *float32 `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	// Renewal duration. The unit is specified by PeriodUnit.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The time unit for the renewal duration, which is the unit of the Period parameter. Valid values: Month, Year. Default value: Month.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// Resource ARN (Aliyun Resource Name).
	//
	// example:
	//
	// acs:ecs:cn-hongkong:1488317743351199:instance/i-j6c6f3lbky38o8rpeqw2
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
	// Promotion details.
	Rules []*GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// Discounted price.
	//
	// example:
	//
	// 500
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) SetCurrency(v string) *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices {
	s.Currency = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) SetDetailInfos(v []*GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesDetailInfos) *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices {
	s.DetailInfos = v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) SetDiscountAmount(v float32) *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices {
	s.DiscountAmount = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) SetOriginalAmount(v float32) *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices {
	s.OriginalAmount = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) SetPeriod(v int32) *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices {
	s.Period = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) SetPeriodUnit(v string) *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices {
	s.PeriodUnit = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) SetResourceArn(v string) *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices {
	s.ResourceArn = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) SetRules(v []*GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesRules) *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices {
	s.Rules = v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices) SetTradeAmount(v float32) *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePrices {
	s.TradeAmount = &v
	return s
}

type GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesDetailInfos struct {
	// Discount amount.
	//
	// example:
	//
	// 100
	DiscountAmount *float32 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// Original price.
	//
	// example:
	//
	// 900
	OriginalAmount *float32 `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	// Pricing module identifier.
	//
	// example:
	//
	// instance
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// Discounted price.
	//
	// example:
	//
	// 500
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesDetailInfos) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesDetailInfos) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesDetailInfos) SetDiscountAmount(v float32) *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesDetailInfos {
	s.DiscountAmount = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesDetailInfos) SetOriginalAmount(v float32) *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesDetailInfos {
	s.OriginalAmount = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesDetailInfos) SetResource(v string) *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesDetailInfos {
	s.Resource = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesDetailInfos) SetTradeAmount(v float32) *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesDetailInfos {
	s.TradeAmount = &v
	return s
}

type GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesRules struct {
	// Promotion description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Promotion name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Promotion ID.
	//
	// example:
	//
	// 1021199213
	RuleDescId *int64 `json:"RuleDescId,omitempty" xml:"RuleDescId,omitempty"`
}

func (s GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesRules) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesRules) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesRules) SetDescription(v string) *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesRules {
	s.Description = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesRules) SetName(v string) *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesRules {
	s.Name = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesRules) SetRuleDescId(v int64) *GetServiceInstanceSubscriptionEstimateCostResponseBodyResourcePricesRules {
	s.RuleDescId = &v
	return s
}

type GetServiceInstanceSubscriptionEstimateCostResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceInstanceSubscriptionEstimateCostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceInstanceSubscriptionEstimateCostResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceSubscriptionEstimateCostResponse) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponse) SetHeaders(v map[string]*string) *GetServiceInstanceSubscriptionEstimateCostResponse {
	s.Headers = v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponse) SetStatusCode(v int32) *GetServiceInstanceSubscriptionEstimateCostResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceInstanceSubscriptionEstimateCostResponse) SetBody(v *GetServiceInstanceSubscriptionEstimateCostResponseBody) *GetServiceInstanceSubscriptionEstimateCostResponse {
	s.Body = v
	return s
}

type GetServiceProvisionsRequest struct {
	// The parameters configured for the service instance.
	//
	// example:
	//
	// {\\"RegionId\\":\\"cn-hangzhou\\",\\"ZoneId\\":\\"cn-hangzhou-g\\",\\"EcsInstanceType\\":\\"ecs.g5.large\\",\\"InstancePassword\\":\\"xxxxxxxx\\",\\"PayType\\":\\"PostPaid\\",\\"PayPeriodUnit\\":\\"Month\\",\\"PayPeriod\\":1}
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-0efc0db451794bxxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The version of the service.
	//
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The template name.
	//
	// example:
	//
	// ECS
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The trial policy. Valid values:
	//
	// 	- Trial: Trials are supported.
	//
	// 	- NotTrial: Trials are not supported.
	//
	// example:
	//
	// Trial
	TrialType *string `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
}

func (s GetServiceProvisionsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceProvisionsRequest) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsRequest) SetParameters(v map[string]interface{}) *GetServiceProvisionsRequest {
	s.Parameters = v
	return s
}

func (s *GetServiceProvisionsRequest) SetRegionId(v string) *GetServiceProvisionsRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceProvisionsRequest) SetServiceId(v string) *GetServiceProvisionsRequest {
	s.ServiceId = &v
	return s
}

func (s *GetServiceProvisionsRequest) SetServiceVersion(v string) *GetServiceProvisionsRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GetServiceProvisionsRequest) SetTemplateName(v string) *GetServiceProvisionsRequest {
	s.TemplateName = &v
	return s
}

func (s *GetServiceProvisionsRequest) SetTrialType(v string) *GetServiceProvisionsRequest {
	s.TrialType = &v
	return s
}

type GetServiceProvisionsShrinkRequest struct {
	// The parameters configured for the service instance.
	//
	// example:
	//
	// {\\"RegionId\\":\\"cn-hangzhou\\",\\"ZoneId\\":\\"cn-hangzhou-g\\",\\"EcsInstanceType\\":\\"ecs.g5.large\\",\\"InstancePassword\\":\\"xxxxxxxx\\",\\"PayType\\":\\"PostPaid\\",\\"PayPeriodUnit\\":\\"Month\\",\\"PayPeriod\\":1}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-0efc0db451794bxxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The version of the service.
	//
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The template name.
	//
	// example:
	//
	// ECS
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The trial policy. Valid values:
	//
	// 	- Trial: Trials are supported.
	//
	// 	- NotTrial: Trials are not supported.
	//
	// example:
	//
	// Trial
	TrialType *string `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
}

func (s GetServiceProvisionsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceProvisionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsShrinkRequest) SetParametersShrink(v string) *GetServiceProvisionsShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *GetServiceProvisionsShrinkRequest) SetRegionId(v string) *GetServiceProvisionsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceProvisionsShrinkRequest) SetServiceId(v string) *GetServiceProvisionsShrinkRequest {
	s.ServiceId = &v
	return s
}

func (s *GetServiceProvisionsShrinkRequest) SetServiceVersion(v string) *GetServiceProvisionsShrinkRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GetServiceProvisionsShrinkRequest) SetTemplateName(v string) *GetServiceProvisionsShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *GetServiceProvisionsShrinkRequest) SetTrialType(v string) *GetServiceProvisionsShrinkRequest {
	s.TrialType = &v
	return s
}

type GetServiceProvisionsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8C27145F-C9F4-545D-A355-DCDDAD63D548
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the services.
	ServiceProvisions []*GetServiceProvisionsResponseBodyServiceProvisions `json:"ServiceProvisions,omitempty" xml:"ServiceProvisions,omitempty" type:"Repeated"`
}

func (s GetServiceProvisionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceProvisionsResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsResponseBody) SetRequestId(v string) *GetServiceProvisionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceProvisionsResponseBody) SetServiceProvisions(v []*GetServiceProvisionsResponseBodyServiceProvisions) *GetServiceProvisionsResponseBody {
	s.ServiceProvisions = v
	return s
}

type GetServiceProvisionsResponseBodyServiceProvisions struct {
	// Indicates whether automatic activation for the service is defined in the template. Valid values:
	//
	// 	- true: Automatic activation for the service is defined in the template.
	//
	// 	- false: Manual activation for the service is defined in the template.
	//
	// example:
	//
	// true
	AutoEnableService *bool `json:"AutoEnableService,omitempty" xml:"AutoEnableService,omitempty"`
	// The URL that points to the activation page of the service.
	//
	// > This parameter is returned if Status is set to Disabled.
	//
	// example:
	//
	// https://common-buy.aliyun.com/?commodityCode=sls
	EnableURL *string `json:"EnableURL,omitempty" xml:"EnableURL,omitempty"`
	// The information about the RAM roles of the service. If this parameter is empty, no RAM role is associated with the service.
	RoleProvision *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision `json:"RoleProvision,omitempty" xml:"RoleProvision,omitempty" type:"Struct"`
	// The service name.
	//
	// example:
	//
	// CS
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The activation status of the service. Valid values:
	//
	// 	- Enabled: The service is activated.
	//
	// 	- Disabled: The service is not activated.
	//
	// 	- Unknown: The activation status of the service is unknown.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The reason why the service is in the Disabled or Unknown state.
	//
	// > This parameter is returned if Status is set to Disabled or Unknown.
	//
	// example:
	//
	// No permission
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
}

func (s GetServiceProvisionsResponseBodyServiceProvisions) String() string {
	return tea.Prettify(s)
}

func (s GetServiceProvisionsResponseBodyServiceProvisions) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) SetAutoEnableService(v bool) *GetServiceProvisionsResponseBodyServiceProvisions {
	s.AutoEnableService = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) SetEnableURL(v string) *GetServiceProvisionsResponseBodyServiceProvisions {
	s.EnableURL = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) SetRoleProvision(v *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision) *GetServiceProvisionsResponseBodyServiceProvisions {
	s.RoleProvision = v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) SetServiceName(v string) *GetServiceProvisionsResponseBodyServiceProvisions {
	s.ServiceName = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) SetStatus(v string) *GetServiceProvisionsResponseBodyServiceProvisions {
	s.Status = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) SetStatusReason(v string) *GetServiceProvisionsResponseBodyServiceProvisions {
	s.StatusReason = &v
	return s
}

type GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision struct {
	// The authorization URL of the RAM role.
	//
	// > This parameter is returned if Created is set to false.
	//
	// example:
	//
	// https://ram.console.aliyun.com/role/authorization?request={"Services":[{"Service":"CS","Roles":[{"RoleName":"AliyunCSManagedVKRole","TemplateId":"AliyunCSManagedVKRole"},{"RoleName":"AliyunCSDefaultRole","TemplateId":"Default"}]}],"ReturnUrl":"https://cs.console.aliyun.com/"}
	AuthorizationURL *string `json:"AuthorizationURL,omitempty" xml:"AuthorizationURL,omitempty"`
	// The RAM roles of the service.
	Roles []*GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
}

func (s GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision) String() string {
	return tea.Prettify(s)
}

func (s GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision) SetAuthorizationURL(v string) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision {
	s.AuthorizationURL = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision) SetRoles(v []*GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision {
	s.Roles = v
	return s
}

type GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles struct {
	// The information about the API operation that is used to create the RAM role.
	ApiForCreation *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation `json:"ApiForCreation,omitempty" xml:"ApiForCreation,omitempty" type:"Struct"`
	// Indicates whether the RAM role is created. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Created *bool `json:"Created,omitempty" xml:"Created,omitempty"`
	// The purpose for which the RAM role is used. Default value: Default. A value of Default indicates that the RAM role is the default role of the service.
	//
	// example:
	//
	// Default
	Function *string `json:"Function,omitempty" xml:"Function,omitempty"`
	// The name of the role.
	//
	// example:
	//
	// AliyunCSManagedVKRole
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles) String() string {
	return tea.Prettify(s)
}

func (s GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles) SetApiForCreation(v *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles {
	s.ApiForCreation = v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles) SetCreated(v bool) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles {
	s.Created = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles) SetFunction(v string) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles {
	s.Function = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles) SetRoleName(v string) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles {
	s.RoleName = &v
	return s
}

type GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation struct {
	// The name of the API operation.
	//
	// example:
	//
	// CreateServiceLinkedRole
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The ID of the Alibaba Cloud service to which the API operation belongs.
	//
	// example:
	//
	// rds
	ApiProductId *string `json:"ApiProductId,omitempty" xml:"ApiProductId,omitempty"`
	// The type of the API operation. Valid values:
	//
	// 	- Open: public
	//
	// 	- Inner: private
	//
	// example:
	//
	// Open
	ApiType *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// The ROS parameters of the cluster.
	//
	// example:
	//
	// { "ServiceLinkedRole": "AliyunServiceRoleForRdsPgsqlOnEcs", "RegionId": "${RegionId}" }
	Parameters map[string]interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
}

func (s GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation) String() string {
	return tea.Prettify(s)
}

func (s GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation) SetApiName(v string) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation {
	s.ApiName = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation) SetApiProductId(v string) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation {
	s.ApiProductId = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation) SetApiType(v string) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation {
	s.ApiType = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation) SetParameters(v map[string]interface{}) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation {
	s.Parameters = v
	return s
}

type GetServiceProvisionsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceProvisionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceProvisionsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceProvisionsResponse) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsResponse) SetHeaders(v map[string]*string) *GetServiceProvisionsResponse {
	s.Headers = v
	return s
}

func (s *GetServiceProvisionsResponse) SetStatusCode(v int32) *GetServiceProvisionsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceProvisionsResponse) SetBody(v *GetServiceProvisionsResponseBody) *GetServiceProvisionsResponse {
	s.Body = v
	return s
}

type GetServiceTemplateParameterConstraintsRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	DeployRegionId *string `json:"DeployRegionId,omitempty" xml:"DeployRegionId,omitempty"`
	// Specifies whether to enable the private connection. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnablePrivateVpcConnection *bool `json:"EnablePrivateVpcConnection,omitempty" xml:"EnablePrivateVpcConnection,omitempty"`
	// The configuration parameters of the service instance.
	Parameters []*GetServiceTemplateParameterConstraintsRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-731f788406024axxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service instance ID.
	//
	// example:
	//
	// si-461ee95f46ca46xxxxxx
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The name of the specification package.
	//
	// example:
	//
	// 套餐一
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// The template name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 模板1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The trial policy. Valid values:
	//
	// 	- Trial: Trials are supported.
	//
	// 	- NotTrial: Trials are not supported.
	//
	// example:
	//
	// NotTrial
	TrialType *string `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
}

func (s GetServiceTemplateParameterConstraintsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceTemplateParameterConstraintsRequest) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetClientToken(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.ClientToken = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetDeployRegionId(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.DeployRegionId = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetEnablePrivateVpcConnection(v bool) *GetServiceTemplateParameterConstraintsRequest {
	s.EnablePrivateVpcConnection = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetParameters(v []*GetServiceTemplateParameterConstraintsRequestParameters) *GetServiceTemplateParameterConstraintsRequest {
	s.Parameters = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetRegionId(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetServiceId(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.ServiceId = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetServiceInstanceId(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetServiceVersion(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetSpecificationName(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.SpecificationName = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetTemplateName(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.TemplateName = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetTrialType(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.TrialType = &v
	return s
}

type GetServiceTemplateParameterConstraintsRequestParameters struct {
	// The name of the parameter. If you do not specify Parameters, the parameters and values in the template are used.
	//
	// >  Parameters is an optional parameter. ParameterKey is required if you specify Parameters.
	//
	// example:
	//
	// InstanceType
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The parameter value that is defined in the template.
	//
	// >  Parameters is an optional parameter. ParameterValue is required if you specify Parameters.
	//
	// example:
	//
	// cn-hangzhou-j
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetServiceTemplateParameterConstraintsRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s GetServiceTemplateParameterConstraintsRequestParameters) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateParameterConstraintsRequestParameters) SetParameterKey(v string) *GetServiceTemplateParameterConstraintsRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequestParameters) SetParameterValue(v string) *GetServiceTemplateParameterConstraintsRequestParameters {
	s.ParameterValue = &v
	return s
}

type GetServiceTemplateParameterConstraintsResponseBody struct {
	// The package family constraints.
	FamilyConstraints []*string `json:"FamilyConstraints,omitempty" xml:"FamilyConstraints,omitempty" type:"Repeated"`
	// The constraints on the parameters.
	ParameterConstraints []*GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints `json:"ParameterConstraints,omitempty" xml:"ParameterConstraints,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 449DC03D-A039-56A6-8D6F-6EBCCCE0EE2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetServiceTemplateParameterConstraintsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceTemplateParameterConstraintsResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateParameterConstraintsResponseBody) SetFamilyConstraints(v []*string) *GetServiceTemplateParameterConstraintsResponseBody {
	s.FamilyConstraints = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBody) SetParameterConstraints(v []*GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) *GetServiceTemplateParameterConstraintsResponseBody {
	s.ParameterConstraints = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBody) SetRequestId(v string) *GetServiceTemplateParameterConstraintsResponseBody {
	s.RequestId = &v
	return s
}

type GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints struct {
	// The valid values of the parameter.
	AllowedValues []*string `json:"AllowedValues,omitempty" xml:"AllowedValues,omitempty" type:"Repeated"`
	// The names of the associated parameters.
	AssociationParameterNames []*string `json:"AssociationParameterNames,omitempty" xml:"AssociationParameterNames,omitempty" type:"Repeated"`
	// The behavior of the parameter. Valid values:
	//
	// 	- NoLimit: No limit is imposed on the value of this parameter.
	//
	// 	- NotSupport: The value of this parameter cannot be queried.
	//
	// 	- QueryError: This parameter failed to be queried.
	//
	// >  If AllowedValues is not returned, Behavior and BehaviorReason are returned, which indicate the behavior of the parameter and the reason for the behavior.
	//
	// example:
	//
	// NoLimit
	Behavior *string `json:"Behavior,omitempty" xml:"Behavior,omitempty"`
	// The reason why the behavior of the parameter is returned.
	//
	// example:
	//
	// No resource property refer to the parameter
	BehaviorReason *string `json:"BehaviorReason,omitempty" xml:"BehaviorReason,omitempty"`
	// The original constraint information.
	OriginalConstraints []*GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints `json:"OriginalConstraints,omitempty" xml:"OriginalConstraints,omitempty" type:"Repeated"`
	// The name of the parameter.
	//
	// example:
	//
	// ZoneInfo
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The error details that are returned if the request fails.
	QueryErrors []*GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors `json:"QueryErrors,omitempty" xml:"QueryErrors,omitempty" type:"Repeated"`
	// The data type of the parameter.
	//
	// example:
	//
	// String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) String() string {
	return tea.Prettify(s)
}

func (s GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetAllowedValues(v []*string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.AllowedValues = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetAssociationParameterNames(v []*string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.AssociationParameterNames = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetBehavior(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.Behavior = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetBehaviorReason(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.BehaviorReason = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetOriginalConstraints(v []*GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.OriginalConstraints = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetParameterKey(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.ParameterKey = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetQueryErrors(v []*GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.QueryErrors = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetType(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.Type = &v
	return s
}

type GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints struct {
	// The valid values of the parameter.
	AllowedValues []*string `json:"AllowedValues,omitempty" xml:"AllowedValues,omitempty" type:"Repeated"`
	// The property name.
	//
	// example:
	//
	// ZoneId
	PropertyName *string `json:"PropertyName,omitempty" xml:"PropertyName,omitempty"`
	// The name of the resource that is defined in the template.
	//
	// example:
	//
	// MyECS
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ALIYUN::ECS::InstanceGroup
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) String() string {
	return tea.Prettify(s)
}

func (s GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) SetAllowedValues(v []*string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints {
	s.AllowedValues = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) SetPropertyName(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints {
	s.PropertyName = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) SetResourceName(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints {
	s.ResourceName = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) SetResourceType(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints {
	s.ResourceType = &v
	return s
}

type GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors struct {
	// The error message.
	//
	// example:
	//
	// record not exist
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The resource name.
	//
	// example:
	//
	// MyECS
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ALIYUN::ECS::InstanceGroup
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors) String() string {
	return tea.Prettify(s)
}

func (s GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors) SetErrorMessage(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors {
	s.ErrorMessage = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors) SetResourceName(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors {
	s.ResourceName = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors) SetResourceType(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors {
	s.ResourceType = &v
	return s
}

type GetServiceTemplateParameterConstraintsResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceTemplateParameterConstraintsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceTemplateParameterConstraintsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceTemplateParameterConstraintsResponse) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateParameterConstraintsResponse) SetHeaders(v map[string]*string) *GetServiceTemplateParameterConstraintsResponse {
	s.Headers = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponse) SetStatusCode(v int32) *GetServiceTemplateParameterConstraintsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponse) SetBody(v *GetServiceTemplateParameterConstraintsResponseBody) *GetServiceTemplateParameterConstraintsResponse {
	s.Body = v
	return s
}

type GetUserInformationRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetUserInformationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserInformationRequest) GoString() string {
	return s.String()
}

func (s *GetUserInformationRequest) SetRegionId(v string) *GetUserInformationRequest {
	s.RegionId = &v
	return s
}

type GetUserInformationResponseBody struct {
	// The delivery settings.
	DeliverySettings *GetUserInformationResponseBodyDeliverySettings `json:"DeliverySettings,omitempty" xml:"DeliverySettings,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 52EBAF16-22F6-53DB-AE1E-44764FC62AF0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUserInformationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserInformationResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserInformationResponseBody) SetDeliverySettings(v *GetUserInformationResponseBodyDeliverySettings) *GetUserInformationResponseBody {
	s.DeliverySettings = v
	return s
}

func (s *GetUserInformationResponseBody) SetRequestId(v string) *GetUserInformationResponseBody {
	s.RequestId = &v
	return s
}

type GetUserInformationResponseBodyDeliverySettings struct {
	// Indicates whether screencast delivery to OSS is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	ActiontrailDeliveryToOssEnabled *bool `json:"ActiontrailDeliveryToOssEnabled,omitempty" xml:"ActiontrailDeliveryToOssEnabled,omitempty"`
	// The name of the OSS bucket.
	//
	// example:
	//
	// 0101data
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// Indicates whether screencast delivery to Object Storage Service (OSS) is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	OssEnabled *bool `json:"OssEnabled,omitempty" xml:"OssEnabled,omitempty"`
	// The number of days for which the screencasts are saved.
	//
	// example:
	//
	// 7
	OssExpirationDays *int64 `json:"OssExpirationDays,omitempty" xml:"OssExpirationDays,omitempty"`
	// The OSS path.
	//
	// example:
	//
	// /test
	OssPath *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
}

func (s GetUserInformationResponseBodyDeliverySettings) String() string {
	return tea.Prettify(s)
}

func (s GetUserInformationResponseBodyDeliverySettings) GoString() string {
	return s.String()
}

func (s *GetUserInformationResponseBodyDeliverySettings) SetActiontrailDeliveryToOssEnabled(v bool) *GetUserInformationResponseBodyDeliverySettings {
	s.ActiontrailDeliveryToOssEnabled = &v
	return s
}

func (s *GetUserInformationResponseBodyDeliverySettings) SetOssBucketName(v string) *GetUserInformationResponseBodyDeliverySettings {
	s.OssBucketName = &v
	return s
}

func (s *GetUserInformationResponseBodyDeliverySettings) SetOssEnabled(v bool) *GetUserInformationResponseBodyDeliverySettings {
	s.OssEnabled = &v
	return s
}

func (s *GetUserInformationResponseBodyDeliverySettings) SetOssExpirationDays(v int64) *GetUserInformationResponseBodyDeliverySettings {
	s.OssExpirationDays = &v
	return s
}

func (s *GetUserInformationResponseBodyDeliverySettings) SetOssPath(v string) *GetUserInformationResponseBodyDeliverySettings {
	s.OssPath = &v
	return s
}

type GetUserInformationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserInformationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserInformationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserInformationResponse) GoString() string {
	return s.String()
}

func (s *GetUserInformationResponse) SetHeaders(v map[string]*string) *GetUserInformationResponse {
	s.Headers = v
	return s
}

func (s *GetUserInformationResponse) SetStatusCode(v int32) *GetUserInformationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserInformationResponse) SetBody(v *GetUserInformationResponseBody) *GetUserInformationResponse {
	s.Body = v
	return s
}

type ListPoliciesRequest struct {
	// Page size.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Token for the next query, an empty nextToken indicates there is no next page.
	//
	// example:
	//
	// BBBAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListPoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListPoliciesRequest) SetMaxResults(v int32) *ListPoliciesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPoliciesRequest) SetNextToken(v string) *ListPoliciesRequest {
	s.NextToken = &v
	return s
}

func (s *ListPoliciesRequest) SetRegionId(v string) *ListPoliciesRequest {
	s.RegionId = &v
	return s
}

type ListPoliciesResponseBody struct {
	// 分页大小。
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Next Token
	//
	// example:
	//
	// AAAAAZ9FmxgN6wKfeK/GOKRnnjU=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Permission policy list
	Policies []*ListPoliciesResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// Request ID
	//
	// example:
	//
	// 51945B04-6AA6-410D-93BA-236E0248B104
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponseBody) SetMaxResults(v int32) *ListPoliciesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPoliciesResponseBody) SetNextToken(v string) *ListPoliciesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPoliciesResponseBody) SetPolicies(v []*ListPoliciesResponseBodyPolicies) *ListPoliciesResponseBody {
	s.Policies = v
	return s
}

func (s *ListPoliciesResponseBody) SetRequestId(v string) *ListPoliciesResponseBody {
	s.RequestId = &v
	return s
}

type ListPoliciesResponseBodyPolicies struct {
	// Permission policy description.
	//
	// example:
	//
	// Only read permission policy
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Policy content.
	//
	// example:
	//
	// {"Action":["*:Describe*","*:List*","*:Get*","*:BatchGet*","*:Query*","*:BatchQuery*","actiontrail:LookupEvents"],"Resource":["*"],"Effect":"Allow"}
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// Permission policy name.
	//
	// example:
	//
	// AliyunComputeNestPolicyForReadOnly
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// Permission policy type.
	//
	// - Custom: Custom policy.
	//
	// - System: System policy.
	//
	// example:
	//
	// System
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ListPoliciesResponseBodyPolicies) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponseBodyPolicies) SetDescription(v string) *ListPoliciesResponseBodyPolicies {
	s.Description = &v
	return s
}

func (s *ListPoliciesResponseBodyPolicies) SetPolicyDocument(v string) *ListPoliciesResponseBodyPolicies {
	s.PolicyDocument = &v
	return s
}

func (s *ListPoliciesResponseBodyPolicies) SetPolicyName(v string) *ListPoliciesResponseBodyPolicies {
	s.PolicyName = &v
	return s
}

func (s *ListPoliciesResponseBodyPolicies) SetPolicyType(v string) *ListPoliciesResponseBodyPolicies {
	s.PolicyType = &v
	return s
}

type ListPoliciesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponse) SetHeaders(v map[string]*string) *ListPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListPoliciesResponse) SetStatusCode(v int32) *ListPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPoliciesResponse) SetBody(v *ListPoliciesResponseBody) *ListPoliciesResponse {
	s.Body = v
	return s
}

type ListServiceCategoriesResponseBody struct {
	// The category list of the service.
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 13FE89A5-C036-56BF-A0FF-A31C59819FD7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListServiceCategoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceCategoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceCategoriesResponseBody) SetCategories(v []*string) *ListServiceCategoriesResponseBody {
	s.Categories = v
	return s
}

func (s *ListServiceCategoriesResponseBody) SetRequestId(v string) *ListServiceCategoriesResponseBody {
	s.RequestId = &v
	return s
}

type ListServiceCategoriesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceCategoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceCategoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceCategoriesResponse) GoString() string {
	return s.String()
}

func (s *ListServiceCategoriesResponse) SetHeaders(v map[string]*string) *ListServiceCategoriesResponse {
	s.Headers = v
	return s
}

func (s *ListServiceCategoriesResponse) SetStatusCode(v int32) *ListServiceCategoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceCategoriesResponse) SetBody(v *ListServiceCategoriesResponseBody) *ListServiceCategoriesResponse {
	s.Body = v
	return s
}

type ListServiceInstanceBillRequest struct {
	// The billing cycle. Format: YYYY-MM.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-05
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The billing date. This parameter is required only if the **Granularity*	- parameter is set to DAILY. Format: YYYY-MM-DD.
	//
	// example:
	//
	// 2025-04-01
	BillingDate *string `json:"BillingDate,omitempty" xml:"BillingDate,omitempty"`
	// The granularity at which bills are queried. Valid values:
	//
	// 	- MONTHLY: queries bills by month. The data queried is consistent with the data that is displayed for the specified billing cycle on the Billing Details tab of the Bill Details page in User Center.
	//
	// 	- DAILY: queries bills by day. The data queried is consistent with the data that is displayed for the specified day on the Billing Details tab of the Bill Details page in User Center.
	//
	// You must set the **BillingDate*	- parameter before you can set the Granularity parameter to DAILY.
	//
	// example:
	//
	// DAILY
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// The number of entries page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token.
	//
	// example:
	//
	// AAAAAVz7BQqj2xtiNSC3d3RAD38=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The service instance ID.
	//
	// example:
	//
	// si-49793f3bfa034ec6a990
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s ListServiceInstanceBillRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceBillRequest) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceBillRequest) SetBillingCycle(v string) *ListServiceInstanceBillRequest {
	s.BillingCycle = &v
	return s
}

func (s *ListServiceInstanceBillRequest) SetBillingDate(v string) *ListServiceInstanceBillRequest {
	s.BillingDate = &v
	return s
}

func (s *ListServiceInstanceBillRequest) SetGranularity(v string) *ListServiceInstanceBillRequest {
	s.Granularity = &v
	return s
}

func (s *ListServiceInstanceBillRequest) SetMaxResults(v int32) *ListServiceInstanceBillRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceBillRequest) SetNextToken(v string) *ListServiceInstanceBillRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceBillRequest) SetServiceInstanceId(v string) *ListServiceInstanceBillRequest {
	s.ServiceInstanceId = &v
	return s
}

type ListServiceInstanceBillResponseBody struct {
	// The billing information of the backup schedule.
	Item []*ListServiceInstanceBillResponseBodyItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAc3HCuYhJi/wvpk4xOr0VLbAx7BkQzyYC+ONO+WudHGKEdB0uWSY7AGnM3qCgm/Ynge7zU6NWdbj0Tegyajyqyc=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2849EE73-AFFA-5AFD-9575-12FA886451DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 0
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServiceInstanceBillResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceBillResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceBillResponseBody) SetItem(v []*ListServiceInstanceBillResponseBodyItem) *ListServiceInstanceBillResponseBody {
	s.Item = v
	return s
}

func (s *ListServiceInstanceBillResponseBody) SetMaxResults(v int32) *ListServiceInstanceBillResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceBillResponseBody) SetNextToken(v string) *ListServiceInstanceBillResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceBillResponseBody) SetRequestId(v string) *ListServiceInstanceBillResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceInstanceBillResponseBody) SetTotalCount(v int64) *ListServiceInstanceBillResponseBody {
	s.TotalCount = &v
	return s
}

type ListServiceInstanceBillResponseBodyItem struct {
	// The billing cycle. Format: YYYY-MM.
	//
	// example:
	//
	// 2025-02
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The billing date. This parameter is required only if the **Granularity*	- parameter is set to DAILY. Format: YYYY-MM-DD.
	//
	// example:
	//
	// 2024-10-23
	BillingDate *string `json:"BillingDate,omitempty" xml:"BillingDate,omitempty"`
	// The billable item.
	//
	// example:
	//
	// Bandwidth
	BillingItem *string `json:"BillingItem,omitempty" xml:"BillingItem,omitempty"`
	// The code of the billable item.
	//
	// example:
	//
	// disk
	BillingItemCode *string `json:"BillingItemCode,omitempty" xml:"BillingItemCode,omitempty"`
	// The currency unit.
	//
	// 	- China site: **CNY**.
	//
	// 	- International site: **USD**.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The amount deducted with resource plans.
	//
	// example:
	//
	// 0
	DeductedByResourcePackage *string `json:"DeductedByResourcePackage,omitempty" xml:"DeductedByResourcePackage,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// rm-bp1z88pb48487907u
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// The discount amount.
	//
	// example:
	//
	// 0
	InvoiceDiscount *string `json:"InvoiceDiscount,omitempty" xml:"InvoiceDiscount,omitempty"`
	// The unit price.
	//
	// example:
	//
	// 0.12
	ListPrice *string `json:"ListPrice,omitempty" xml:"ListPrice,omitempty"`
	// The unit of the unit price.
	//
	// example:
	//
	// CNY/GB
	ListPriceUnit *string `json:"ListPriceUnit,omitempty" xml:"ListPriceUnit,omitempty"`
	// The pretax amount.
	//
	// example:
	//
	// 0
	PretaxAmount *string `json:"PretaxAmount,omitempty" xml:"PretaxAmount,omitempty"`
	// The pretax gross amount.
	//
	// example:
	//
	// 0
	PretaxGrossAmount *string `json:"PretaxGrossAmount,omitempty" xml:"PretaxGrossAmount,omitempty"`
	// The code of the service.
	//
	// example:
	//
	// sls
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The specific service resource.
	//
	// example:
	//
	// sls
	ProductDetail *string `json:"ProductDetail,omitempty" xml:"ProductDetail,omitempty"`
	// The name of the cloud service or the name of the service-linked role with which the cloud service is associated.
	//
	// example:
	//
	// NLB
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The billing cycle in which the bill is split.
	//
	// example:
	//
	// 2021-07
	SplitBillingCycle *string `json:"SplitBillingCycle,omitempty" xml:"SplitBillingCycle,omitempty"`
	// The billing method. Valid values:
	//
	// 	- Subscription: the subscription billing method.
	//
	// 	- PayAsYouGo: the pay-as-you-go billing method.
	//
	// example:
	//
	// Subscription
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	// The amount of resource usage.
	//
	// example:
	//
	// {\\"EmbeddingTokens\\": 314}
	Usage *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
	// The unit of usage.
	//
	// example:
	//
	// GB
	UsageUnit *string `json:"UsageUnit,omitempty" xml:"UsageUnit,omitempty"`
}

func (s ListServiceInstanceBillResponseBodyItem) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceBillResponseBodyItem) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceBillResponseBodyItem) SetBillingCycle(v string) *ListServiceInstanceBillResponseBodyItem {
	s.BillingCycle = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetBillingDate(v string) *ListServiceInstanceBillResponseBodyItem {
	s.BillingDate = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetBillingItem(v string) *ListServiceInstanceBillResponseBodyItem {
	s.BillingItem = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetBillingItemCode(v string) *ListServiceInstanceBillResponseBodyItem {
	s.BillingItemCode = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetCurrency(v string) *ListServiceInstanceBillResponseBodyItem {
	s.Currency = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetDeductedByResourcePackage(v string) *ListServiceInstanceBillResponseBodyItem {
	s.DeductedByResourcePackage = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetInstanceID(v string) *ListServiceInstanceBillResponseBodyItem {
	s.InstanceID = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetInvoiceDiscount(v string) *ListServiceInstanceBillResponseBodyItem {
	s.InvoiceDiscount = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetListPrice(v string) *ListServiceInstanceBillResponseBodyItem {
	s.ListPrice = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetListPriceUnit(v string) *ListServiceInstanceBillResponseBodyItem {
	s.ListPriceUnit = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetPretaxAmount(v string) *ListServiceInstanceBillResponseBodyItem {
	s.PretaxAmount = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetPretaxGrossAmount(v string) *ListServiceInstanceBillResponseBodyItem {
	s.PretaxGrossAmount = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetProductCode(v string) *ListServiceInstanceBillResponseBodyItem {
	s.ProductCode = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetProductDetail(v string) *ListServiceInstanceBillResponseBodyItem {
	s.ProductDetail = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetProductName(v string) *ListServiceInstanceBillResponseBodyItem {
	s.ProductName = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetSplitBillingCycle(v string) *ListServiceInstanceBillResponseBodyItem {
	s.SplitBillingCycle = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetSubscriptionType(v string) *ListServiceInstanceBillResponseBodyItem {
	s.SubscriptionType = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetUsage(v string) *ListServiceInstanceBillResponseBodyItem {
	s.Usage = &v
	return s
}

func (s *ListServiceInstanceBillResponseBodyItem) SetUsageUnit(v string) *ListServiceInstanceBillResponseBodyItem {
	s.UsageUnit = &v
	return s
}

type ListServiceInstanceBillResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceInstanceBillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceInstanceBillResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceBillResponse) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceBillResponse) SetHeaders(v map[string]*string) *ListServiceInstanceBillResponse {
	s.Headers = v
	return s
}

func (s *ListServiceInstanceBillResponse) SetStatusCode(v int32) *ListServiceInstanceBillResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceInstanceBillResponse) SetBody(v *ListServiceInstanceBillResponseBody) *ListServiceInstanceBillResponse {
	s.Body = v
	return s
}

type ListServiceInstanceLogsRequest struct {
	// The filters.
	Filter []*ListServiceInstanceLogsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The log source. When this field is empty, query logs with the source set to computeNest and ros. Valid values:
	//
	// computeNest : logs of the deployment and upgrade of the service instance.
	//
	// application: logs generated by the application.
	//
	// actionTrail: logs generated by ActionTrail.
	//
	// compliancePack: Logs originating from the compliance package.
	//
	// ros: Logs originating from ROS.
	//
	// meteringData：Logs originating from the pay-as-you-go model.
	//
	// example:
	//
	// computeNest
	LogSource *string `json:"LogSource,omitempty" xml:"LogSource,omitempty"`
	// The Logstore. You must specify this parameter if you set LogSource to application.
	//
	// example:
	//
	// logabc
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// BBBAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID. Valid values:
	//
	// 	- cn-hangzhou: China (Hangzhou).
	//
	// 	- ap-southeast-1: Singapore.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-70a3b15bb626435b****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The order in which you want to sort the results. Valid values:
	//
	// 	- Ascending
	//
	// 	- (Default) Descending
	//
	// example:
	//
	// Ascending: Ascending order
	//
	// Descending (default value): Descending order
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListServiceInstanceLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceLogsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceLogsRequest) SetFilter(v []*ListServiceInstanceLogsRequestFilter) *ListServiceInstanceLogsRequest {
	s.Filter = v
	return s
}

func (s *ListServiceInstanceLogsRequest) SetLogSource(v string) *ListServiceInstanceLogsRequest {
	s.LogSource = &v
	return s
}

func (s *ListServiceInstanceLogsRequest) SetLogstore(v string) *ListServiceInstanceLogsRequest {
	s.Logstore = &v
	return s
}

func (s *ListServiceInstanceLogsRequest) SetMaxResults(v int32) *ListServiceInstanceLogsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceLogsRequest) SetNextToken(v string) *ListServiceInstanceLogsRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceLogsRequest) SetRegionId(v string) *ListServiceInstanceLogsRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceInstanceLogsRequest) SetServiceInstanceId(v string) *ListServiceInstanceLogsRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *ListServiceInstanceLogsRequest) SetSortOrder(v string) *ListServiceInstanceLogsRequest {
	s.SortOrder = &v
	return s
}

type ListServiceInstanceLogsRequestFilter struct {
	// The parameter name of the filter. You can specify one or more filters. Valid values:
	//
	// 	- StartTime: the start time of the log event.
	//
	// 	- EndTime: the end time of the ActionTrail event.
	//
	// 	- EventName: the name of the ActionTrail event.
	//
	// 	- ResourceName: the name of the ActionTrail resource.
	//
	// 	- ApplicationGroupName: the name of the application group.
	//
	// example:
	//
	// - StartTime
	//
	// - EndTime
	//
	// - EventName
	//
	// - ResourceName
	//
	// - ApplicationGroupName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameter value N of the filter. Valid values of N: 1 to 10.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListServiceInstanceLogsRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceLogsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceLogsRequestFilter) SetName(v string) *ListServiceInstanceLogsRequestFilter {
	s.Name = &v
	return s
}

func (s *ListServiceInstanceLogsRequestFilter) SetValue(v []*string) *ListServiceInstanceLogsRequestFilter {
	s.Value = v
	return s
}

type ListServiceInstanceLogsResponseBody struct {
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 51945B04-6AA6-410D-93BA-236E0248B104
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The logs of the service instance.
	ServiceInstancesLogs []*ListServiceInstanceLogsResponseBodyServiceInstancesLogs `json:"ServiceInstancesLogs,omitempty" xml:"ServiceInstancesLogs,omitempty" type:"Repeated"`
}

func (s ListServiceInstanceLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceLogsResponseBody) SetMaxResults(v int32) *ListServiceInstanceLogsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBody) SetNextToken(v string) *ListServiceInstanceLogsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBody) SetRequestId(v string) *ListServiceInstanceLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBody) SetServiceInstancesLogs(v []*ListServiceInstanceLogsResponseBodyServiceInstancesLogs) *ListServiceInstanceLogsResponseBody {
	s.ServiceInstancesLogs = v
	return s
}

type ListServiceInstanceLogsResponseBodyServiceInstancesLogs struct {
	// Compliance package risk types. For example, data security checks within a VPC, such as VpcDataRisk
	//
	// example:
	//
	// VpcDataRisk
	CompliancePackType *string `json:"CompliancePackType,omitempty" xml:"CompliancePackType,omitempty"`
	// Specific risk rule names for the compliance package. For example, ECS instance migration out of VPC - ecs-move-out-vpc.
	//
	// example:
	//
	// ecs-move-out-vpc
	ComplianceRuleName *string `json:"ComplianceRuleName,omitempty" xml:"ComplianceRuleName,omitempty"`
	// The log content.
	//
	// example:
	//
	// Start creating service instance
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The log type. Valid values:
	//
	// 	- serviceInstance: log generated by the service instance.
	//
	// 	- resource: log generated by ROS resources.
	//
	// example:
	//
	// serviceInstance
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// si-5c6525c0589545c3****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ROS.Stack
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The log source. Valid values:
	//
	// computeNest : logs of the deployment and upgrade of the service instance.
	//
	// application: logs generated by the application.
	//
	// actionTrail: logs generated by ActionTrail.
	//
	// compliancePack: Logs originating from the compliance package.
	//
	// ros: Logs originating from ROS.
	//
	// meteringData：Logs originating from the pay-as-you-go model.
	//
	// example:
	//
	// computeNest
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The state of the service instance. Valid values:
	//
	// 	- Creating: The service instance is being created.
	//
	// 	- Created: The service instance is created.
	//
	// 	- Deploying: The service instance is being deployed.
	//
	// 	- Deployed: The service instance is deployed.
	//
	// 	- DeployedFailed: The service instance failed to be deployed.
	//
	// 	- Expired: The service instance expired.
	//
	// 	- ExtendSuccess: The service instance is renewed.
	//
	// 	- Upgrading: The service instance is being updated.
	//
	// 	- UpgradeSuccess: The service instance is updated.
	//
	// example:
	//
	// Creating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The timestamp of the service instance log.
	//
	// example:
	//
	// 2021-05-21T00:00:00Z
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s ListServiceInstanceLogsResponseBodyServiceInstancesLogs) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceLogsResponseBodyServiceInstancesLogs) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetCompliancePackType(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.CompliancePackType = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetComplianceRuleName(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.ComplianceRuleName = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetContent(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.Content = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetLogType(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.LogType = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetResourceId(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.ResourceId = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetResourceType(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.ResourceType = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetSource(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.Source = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetStatus(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.Status = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetTimestamp(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.Timestamp = &v
	return s
}

type ListServiceInstanceLogsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceInstanceLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceInstanceLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceLogsResponse) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceLogsResponse) SetHeaders(v map[string]*string) *ListServiceInstanceLogsResponse {
	s.Headers = v
	return s
}

func (s *ListServiceInstanceLogsResponse) SetStatusCode(v int32) *ListServiceInstanceLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceInstanceLogsResponse) SetBody(v *ListServiceInstanceLogsResponseBody) *ListServiceInstanceLogsResponse {
	s.Body = v
	return s
}

type ListServiceInstanceResourcesRequest struct {
	// The filter conditions. Vaild values:
	//
	// - ExpireTimeStart：
	//
	// Query start time for Subscription resource expiration.
	//
	// <notice>Notice Note: Only supports querying service instances on private deployments.	Notice:
	//
	// - ExpireTimeEnd：Query end time for Subscription resource expiration.
	//
	// <notice>Notice Note: Only supports querying service instances on private deployments.	Notice:
	//
	// - PayType：The billing method of the read-only instance.
	//
	// <notice>Notice Note: Only supports querying service instances on private deployments.<notice>
	//
	//    Valid values:
	//
	//    - PayAsYouGo
	//
	//    - Subscription
	//
	// - ResourceARN：The Alibaba Cloud Resource Name (ARN) of a resource.
	Filters []*ListServiceInstanceResourcesRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the next query. Valid values:
	//
	// 	- If **NextToken*	- is not returned, it indicates that no additional results exist.
	//
	// 	- If **NextToken*	- was returned in the previous query, specify the value to obtain the next set of results.
	//
	// example:
	//
	// AAAAAc3HCuYhJi/wvpk4xOr0VLbAx7BkQzyYC+ONO+WudHGKEdB0uWSY7AGnM3qCgm/Ynge7zU6NWdbj0Tegyajyqyc=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID. Valid values:
	//
	// 	- cn-hangzhou: China (Hangzhou).
	//
	// 	- ap-southeast-1: Singapore.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-d8a0cc2a1ee04dce****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// Service Instance resource type，include AliyunResource and ContainerResource.
	//
	// example:
	//
	// AliyunResource
	ServiceInstanceResourceType *string `json:"ServiceInstanceResourceType,omitempty" xml:"ServiceInstanceResourceType,omitempty"`
	// The tag key and value.
	Tag []*ListServiceInstanceResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListServiceInstanceResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceResourcesRequest) SetFilters(v []*ListServiceInstanceResourcesRequestFilters) *ListServiceInstanceResourcesRequest {
	s.Filters = v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetMaxResults(v int32) *ListServiceInstanceResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetNextToken(v string) *ListServiceInstanceResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetRegionId(v string) *ListServiceInstanceResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetServiceInstanceId(v string) *ListServiceInstanceResourcesRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetServiceInstanceResourceType(v string) *ListServiceInstanceResourcesRequest {
	s.ServiceInstanceResourceType = &v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetTag(v []*ListServiceInstanceResourcesRequestTag) *ListServiceInstanceResourcesRequest {
	s.Tag = v
	return s
}

type ListServiceInstanceResourcesRequestFilters struct {
	// Vaild values:
	//
	// - ExpireTimeStart
	//
	// - ExpireTimeEnd
	//
	// - PayType
	//
	// - ResourceARN
	//
	// example:
	//
	// ExpireTimeStart
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the filter condition.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListServiceInstanceResourcesRequestFilters) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceResourcesRequestFilters) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceResourcesRequestFilters) SetName(v string) *ListServiceInstanceResourcesRequestFilters {
	s.Name = &v
	return s
}

func (s *ListServiceInstanceResourcesRequestFilters) SetValues(v []*string) *ListServiceInstanceResourcesRequestFilters {
	s.Values = v
	return s
}

type ListServiceInstanceResourcesRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServiceInstanceResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceResourcesRequestTag) SetKey(v string) *ListServiceInstanceResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListServiceInstanceResourcesRequestTag) SetValue(v string) *ListServiceInstanceResourcesRequestTag {
	s.Value = &v
	return s
}

type ListServiceInstanceResourcesResponseBody struct {
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token.
	//
	// example:
	//
	// AAAAAc3HCuYhJi/wvpk4xOr0VLbAx7BkQzyYC+ONO+WudHGKEdB0uWSY7AGnM3qCgm/Ynge7zU6NWdbj0Tegyajyqyc=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resources.
	Resources []*ListServiceInstanceResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s ListServiceInstanceResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceResourcesResponseBody) SetMaxResults(v int32) *ListServiceInstanceResourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBody) SetNextToken(v string) *ListServiceInstanceResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBody) SetRequestId(v string) *ListServiceInstanceResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBody) SetResources(v []*ListServiceInstanceResourcesResponseBodyResources) *ListServiceInstanceResourcesResponseBody {
	s.Resources = v
	return s
}

type ListServiceInstanceResourcesResponseBodyResources struct {
	// The time when the resource was created.
	//
	// example:
	//
	// 2022-01-01T12:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the resource expires.
	//
	// example:
	//
	// 2022-03-01T12:00:00
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The billing method. Valid values:
	//
	// 	- Subscription
	//
	// 	- PayAsYouGo
	//
	// example:
	//
	// Subscription
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The code of the cloud service.
	//
	// example:
	//
	// rds
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The type of the cloud service.
	//
	// example:
	//
	// RDS
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The renewal state. Valid values:
	//
	// 	- AutoRenewal
	//
	// 	- ManualRenewal
	//
	// 	- NotRenewal
	//
	// example:
	//
	// AutoRenewal
	RenewStatus *string `json:"RenewStatus,omitempty" xml:"RenewStatus,omitempty"`
	// The renewal period.
	//
	// example:
	//
	// 1
	RenewalPeriod *int32 `json:"RenewalPeriod,omitempty" xml:"RenewalPeriod,omitempty"`
	// The unit of the renewal period. Valid values:
	//
	// 	- Month
	//
	// 	- Year
	//
	// example:
	//
	// Month
	RenewalPeriodUnit *string `json:"RenewalPeriodUnit,omitempty" xml:"RenewalPeriodUnit,omitempty"`
	// The ARN of the resource.
	//
	// example:
	//
	// arn:acs:sag:cn-hangzhou:130920852836****:ccn/ccn-b3qf0q439sq2de****
	ResourceARN *string `json:"ResourceARN,omitempty" xml:"ResourceARN,omitempty"`
	// The state of the resource. Valid values:
	//
	// 	- running
	//
	// 	- waiting
	//
	// 	- terminated
	//
	// >  This parameter is returned only for containers.
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListServiceInstanceResourcesResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetCreateTime(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.CreateTime = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetExpireTime(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.ExpireTime = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetPayType(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.PayType = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetProductCode(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.ProductCode = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetProductType(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.ProductType = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetRenewStatus(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.RenewStatus = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetRenewalPeriod(v int32) *ListServiceInstanceResourcesResponseBodyResources {
	s.RenewalPeriod = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetRenewalPeriodUnit(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.RenewalPeriodUnit = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetResourceARN(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.ResourceARN = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetStatus(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.Status = &v
	return s
}

type ListServiceInstanceResourcesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceInstanceResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceInstanceResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceResourcesResponse) SetHeaders(v map[string]*string) *ListServiceInstanceResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListServiceInstanceResourcesResponse) SetStatusCode(v int32) *ListServiceInstanceResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceInstanceResourcesResponse) SetBody(v *ListServiceInstanceResourcesResponseBody) *ListServiceInstanceResourcesResponse {
	s.Body = v
	return s
}

type ListServiceInstanceUpgradeHistoryRequest struct {
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// BBBAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-70a3b15bb62643xxxxxx
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s ListServiceInstanceUpgradeHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceUpgradeHistoryRequest) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceUpgradeHistoryRequest) SetMaxResults(v int32) *ListServiceInstanceUpgradeHistoryRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryRequest) SetNextToken(v string) *ListServiceInstanceUpgradeHistoryRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryRequest) SetRegionId(v string) *ListServiceInstanceUpgradeHistoryRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryRequest) SetServiceInstanceId(v string) *ListServiceInstanceUpgradeHistoryRequest {
	s.ServiceInstanceId = &v
	return s
}

type ListServiceInstanceUpgradeHistoryResponseBody struct {
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAfu+XtuBE55iRLHEYYuojI41
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EE3EDF4E-B3B1-19B6-BD01-30D4D00F6E5D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The upgrade history.
	UpgradeHistory []*ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory `json:"UpgradeHistory,omitempty" xml:"UpgradeHistory,omitempty" type:"Repeated"`
}

func (s ListServiceInstanceUpgradeHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceUpgradeHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceUpgradeHistoryResponseBody) SetMaxResults(v int32) *ListServiceInstanceUpgradeHistoryResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBody) SetNextToken(v string) *ListServiceInstanceUpgradeHistoryResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBody) SetRequestId(v string) *ListServiceInstanceUpgradeHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBody) SetTotalCount(v int64) *ListServiceInstanceUpgradeHistoryResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBody) SetUpgradeHistory(v []*ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) *ListServiceInstanceUpgradeHistoryResponseBody {
	s.UpgradeHistory = v
	return s
}

type ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory struct {
	// The time when the update ended.
	//
	// example:
	//
	// 2022-04-26T09:09:51Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The version before the upgrade.
	//
	// example:
	//
	// 1
	FromVersion *string `json:"FromVersion,omitempty" xml:"FromVersion,omitempty"`
	// The upgrade result.
	//
	// example:
	//
	// {\\"PreUpgradeExecutionId\\":\\"exec-123\\"}
	Results *string `json:"Results,omitempty" xml:"Results,omitempty"`
	// The time when the update started.
	//
	// example:
	//
	// 2022-04-26T08:09:51Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the update. Valid values:
	//
	// 	- upgrading: The service instance is being upgraded.
	//
	// 	- UpgradeSuccessful: The service instance is upgraded.
	//
	// 	- UpgradeFailed: The service instance failed to be upgraded.
	//
	// example:
	//
	// UpgradeFailed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The version after the upgrade.
	//
	// example:
	//
	// 3
	ToVersion *string `json:"ToVersion,omitempty" xml:"ToVersion,omitempty"`
	// The update type.
	//
	// example:
	//
	// Upgrade
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the upgrade record.
	//
	// example:
	//
	// exec-123
	UpgradeHistoryId *string `json:"UpgradeHistoryId,omitempty" xml:"UpgradeHistoryId,omitempty"`
}

func (s ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) SetEndTime(v string) *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory {
	s.EndTime = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) SetFromVersion(v string) *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory {
	s.FromVersion = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) SetResults(v string) *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory {
	s.Results = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) SetStartTime(v string) *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory {
	s.StartTime = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) SetStatus(v string) *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory {
	s.Status = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) SetToVersion(v string) *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory {
	s.ToVersion = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) SetType(v string) *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory {
	s.Type = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory) SetUpgradeHistoryId(v string) *ListServiceInstanceUpgradeHistoryResponseBodyUpgradeHistory {
	s.UpgradeHistoryId = &v
	return s
}

type ListServiceInstanceUpgradeHistoryResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceInstanceUpgradeHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceInstanceUpgradeHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceUpgradeHistoryResponse) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceUpgradeHistoryResponse) SetHeaders(v map[string]*string) *ListServiceInstanceUpgradeHistoryResponse {
	s.Headers = v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponse) SetStatusCode(v int32) *ListServiceInstanceUpgradeHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceInstanceUpgradeHistoryResponse) SetBody(v *ListServiceInstanceUpgradeHistoryResponseBody) *ListServiceInstanceUpgradeHistoryResponse {
	s.Body = v
	return s
}

type ListServiceInstancesRequest struct {
	// The filter.
	Filter []*ListServiceInstancesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// BBBAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tag key and value.
	Tag []*ListServiceInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListServiceInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesRequest) SetFilter(v []*ListServiceInstancesRequestFilter) *ListServiceInstancesRequest {
	s.Filter = v
	return s
}

func (s *ListServiceInstancesRequest) SetMaxResults(v int32) *ListServiceInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstancesRequest) SetNextToken(v string) *ListServiceInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstancesRequest) SetRegionId(v string) *ListServiceInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceInstancesRequest) SetResourceGroupId(v string) *ListServiceInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServiceInstancesRequest) SetTag(v []*ListServiceInstancesRequestTag) *ListServiceInstancesRequest {
	s.Tag = v
	return s
}

type ListServiceInstancesRequestFilter struct {
	// The parameter name of the filter. You can specify one or more parameter names to query services. Valid values:
	//
	// - Name：Query by service name.
	//
	// - ServiceInstanceName：Query by service  instance name.
	//
	// - ServiceInstanceId：Query by service  instance ID.
	//
	// - ServiceId：Query by service ID.
	//
	// - Version：Query by service version.
	//
	// - Status：Query by service status.
	//
	// - DeployType: Query by service deployType.
	//
	// - ServiceType：Query by service deployType.
	//
	// example:
	//
	// ServiceInstanceId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameter values of the filter.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListServiceInstancesRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesRequestFilter) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesRequestFilter) SetName(v string) *ListServiceInstancesRequestFilter {
	s.Name = &v
	return s
}

func (s *ListServiceInstancesRequestFilter) SetValue(v []*string) *ListServiceInstancesRequestFilter {
	s.Value = v
	return s
}

type ListServiceInstancesRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServiceInstancesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesRequestTag) SetKey(v string) *ListServiceInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *ListServiceInstancesRequestTag) SetValue(v string) *ListServiceInstancesRequestTag {
	s.Value = &v
	return s
}

type ListServiceInstancesResponseBody struct {
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token.
	//
	// example:
	//
	// AAAAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E50287CB-AABF-4877-92C0-289B339A1546
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the service instances.
	ServiceInstances []*ListServiceInstancesResponseBodyServiceInstances `json:"ServiceInstances,omitempty" xml:"ServiceInstances,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServiceInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBody) SetMaxResults(v int32) *ListServiceInstancesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstancesResponseBody) SetNextToken(v string) *ListServiceInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstancesResponseBody) SetRequestId(v string) *ListServiceInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceInstancesResponseBody) SetServiceInstances(v []*ListServiceInstancesResponseBodyServiceInstances) *ListServiceInstancesResponseBody {
	s.ServiceInstances = v
	return s
}

func (s *ListServiceInstancesResponseBody) SetTotalCount(v int64) *ListServiceInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListServiceInstancesResponseBodyServiceInstances struct {
	// The business state of the service instance. Valid values:
	//
	// 	- Normal
	//
	// 	- Renewing
	//
	// 	- RenewFailed
	//
	// 	- Expired
	//
	// example:
	//
	// Normal
	BizStatus *string `json:"BizStatus,omitempty" xml:"BizStatus,omitempty"`
	// The time when the service instance was created.
	//
	// example:
	//
	// 2021-05-20T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the service instance supports the hosted O\\&M feature. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableInstanceOps *bool `json:"EnableInstanceOps,omitempty" xml:"EnableInstanceOps,omitempty"`
	// The time when the service instance expires.
	//
	// example:
	//
	// 2022-01-01T12:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the Alibaba Cloud Marketplace instance.
	//
	// example:
	//
	// 5827****
	MarketInstanceId *string `json:"MarketInstanceId,omitempty" xml:"MarketInstanceId,omitempty"`
	// The name of the service instance.
	//
	// example:
	//
	// TestName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the managed service instance.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	OperatedServiceInstanceId *string `json:"OperatedServiceInstanceId,omitempty" xml:"OperatedServiceInstanceId,omitempty"`
	// The end of the time range during which hosted O\\&M is implemented.
	//
	// example:
	//
	// 2022-01-28T06:48:56Z
	OperationEndTime *string `json:"OperationEndTime,omitempty" xml:"OperationEndTime,omitempty"`
	// The beginning of the time range during which hosted O\\&M is implemented.
	//
	// example:
	//
	// 2021-12-29T06:48:56Z
	OperationStartTime *string `json:"OperationStartTime,omitempty" xml:"OperationStartTime,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 2306175xxxxxxxx
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The information returned after the service instance is created.
	//
	// example:
	//
	// {"managementUrl": "http://xx.xx"}
	Outputs *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// The parameters of the service instance.
	//
	// example:
	//
	// {"param":"value"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The billing method. Valid values:
	//
	// 	- Permanent: Once you purchase the service, you can use it permanently.
	//
	// 	- Subscription: You purchase the service from Alibaba Cloud Marketplace and are charged for the service on a subscription basis.
	//
	// 	- PayAsYouGo: You purchase the service from Alibaba Cloud Marketplace and are charged for the service on a pay-as-you-go basis.
	//
	// 	- CustomFixTime: You are charged for the service based on a custom duration fixed by the service provider.
	//
	// example:
	//
	// Subscription
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The deployment progress of the service instance, in percentage.
	//
	// example:
	//
	// 90
	Progress *int64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aekz6scpcxxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resources.
	//
	// example:
	//
	// [{"StackId": "stack-xxx"}]
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// The services.
	Service *ListServiceInstancesResponseBodyServiceInstancesService `json:"Service,omitempty" xml:"Service,omitempty" type:"Struct"`
	// The service instance ID.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The type of the service. Valid values:
	//
	// 	- private: The service is a private service and is deployed within the account of a customer.
	//
	// 	- managed: The service is a fully managed service and is deployed within the account of a service provider.
	//
	// 	- operation: The service is a hosted O\\&M service.
	//
	// 	- poc: The service is a trial service.
	//
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The source from which the service instance is created.
	//
	// example:
	//
	// Supplier
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The state of the service instance. Valid values:
	//
	// 	- Created
	//
	// 	- Deploying
	//
	// 	- DeployedFailed
	//
	// 	- Deployed
	//
	// 	- Upgrading
	//
	// 	- Deleting
	//
	// 	- Deleted
	//
	// 	- DeletedFailed
	//
	// example:
	//
	// Deployed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The description of the deployment of the service instance.
	//
	// example:
	//
	// deploy successfully
	StatusDetail *string `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
	// The custom tags.
	Tags []*ListServiceInstancesResponseBodyServiceInstancesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The template name.
	//
	// example:
	//
	// 模板1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The time when the service instance was updated.
	//
	// example:
	//
	// 2021-05-20T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListServiceInstancesResponseBodyServiceInstances) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstances) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetBizStatus(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.BizStatus = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetCreateTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.CreateTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetEnableInstanceOps(v bool) *ListServiceInstancesResponseBodyServiceInstances {
	s.EnableInstanceOps = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetEndTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.EndTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetMarketInstanceId(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.MarketInstanceId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetName(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Name = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetOperatedServiceInstanceId(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.OperatedServiceInstanceId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetOperationEndTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.OperationEndTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetOperationStartTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.OperationStartTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetOrderId(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.OrderId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetOutputs(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Outputs = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetParameters(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Parameters = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetPayType(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.PayType = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetProgress(v int64) *ListServiceInstancesResponseBodyServiceInstances {
	s.Progress = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetResourceGroupId(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetResources(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Resources = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetService(v *ListServiceInstancesResponseBodyServiceInstancesService) *ListServiceInstancesResponseBodyServiceInstances {
	s.Service = v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetServiceInstanceId(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.ServiceInstanceId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetServiceType(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.ServiceType = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetSource(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Source = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetStatus(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Status = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetStatusDetail(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.StatusDetail = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetTags(v []*ListServiceInstancesResponseBodyServiceInstancesTags) *ListServiceInstancesResponseBodyServiceInstances {
	s.Tags = v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetTemplateName(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.TemplateName = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetUpdateTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.UpdateTime = &v
	return s
}

type ListServiceInstancesResponseBodyServiceInstancesService struct {
	// The commodity details.
	Commodity *ListServiceInstancesResponseBodyServiceInstancesServiceCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// The deployment type of the service. Valid values:
	//
	// 	- ros: The service is deployed by using Resource Orchestration Service (ROS).
	//
	// 	- terraform: The service is deployed by using Terraform.
	//
	// 	- ack: The service is deployed by using Alibaba Cloud Container Service for Kubernetes (ACK).
	//
	// 	- spi: The service is deployed by calling the Service Provider Interface (SPI).
	//
	// 	- operation: The service is deployed by using a hosted O\\&M service.
	//
	// example:
	//
	// ros
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// The time when the service was published.
	//
	// example:
	//
	// 2021-05-21T00:00:00Z
	PublishTime *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-0e6fca6a51a54420****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The information about the service.
	ServiceInfos []*ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos `json:"ServiceInfos,omitempty" xml:"ServiceInfos,omitempty" type:"Repeated"`
	// The type of the service. Valid values:
	//
	// 	- private: The service is a private service and is deployed within the account of a customer.
	//
	// 	- managed: The service is a fully managed service and is deployed within the account of a service provider.
	//
	// 	- operation: The service is a hosted O\\&M service.
	//
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The service state.
	//
	// example:
	//
	// Online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the service provider.
	//
	// example:
	//
	// Alibaba Cloud
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// The URL of the service provider.
	//
	// example:
	//
	// http://example.com
	SupplierUrl *string `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The custom version name defined by the service provider.
	//
	// example:
	//
	// 1.1.0
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s ListServiceInstancesResponseBodyServiceInstancesService) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstancesService) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetCommodity(v *ListServiceInstancesResponseBodyServiceInstancesServiceCommodity) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.Commodity = v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetDeployType(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.DeployType = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetPublishTime(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.PublishTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetServiceId(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.ServiceId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetServiceInfos(v []*ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.ServiceInfos = v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetServiceType(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.ServiceType = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetStatus(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.Status = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetSupplierName(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.SupplierName = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetSupplierUrl(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.SupplierUrl = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetVersion(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.Version = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetVersionName(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.VersionName = &v
	return s
}

type ListServiceInstancesResponseBodyServiceInstancesServiceCommodity struct {
	// The configuration metadata related to SaaS Boost.
	//
	// example:
	//
	// { // Specifies whether to associate the service with the SaaS Boost commodity. Default value: false. "Enabled":true/false // The public endpoint of the SaaS Boost instance. "PublicAccessUrl":"https://example.com" }
	SaasBoostMetadata *string `json:"SaasBoostMetadata,omitempty" xml:"SaasBoostMetadata,omitempty"`
	// The platform type. Valid values:
	//
	// 	- marketplace: Alibaba Cloud Marketplace.
	//
	// 	- Css: Lingxiao.
	//
	// 	- SaasBoost: SaaS Boost.
	//
	// example:
	//
	// Marketplace
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListServiceInstancesResponseBodyServiceInstancesServiceCommodity) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstancesServiceCommodity) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceCommodity) SetSaasBoostMetadata(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceCommodity {
	s.SaasBoostMetadata = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceCommodity) SetType(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceCommodity {
	s.Type = &v
	return s
}

type ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos struct {
	// The URL of the service icon.
	//
	// example:
	//
	// https://example.com/service-image/c1c4a559-cc60-4af1-b976-98f356602462.png
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The language of the service instance.
	//
	// example:
	//
	// zh-CN
	Locale *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// wordpress
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The description of the service.
	//
	// example:
	//
	// B是A公司自主设计并研发的开源分布式的关系型数据库
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
}

func (s ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) SetImage(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos {
	s.Image = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) SetLocale(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos {
	s.Locale = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) SetName(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos {
	s.Name = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) SetShortDescription(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos {
	s.ShortDescription = &v
	return s
}

type ListServiceInstancesResponseBodyServiceInstancesTags struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServiceInstancesResponseBodyServiceInstancesTags) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstancesTags) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstancesTags) SetKey(v string) *ListServiceInstancesResponseBodyServiceInstancesTags {
	s.Key = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesTags) SetValue(v string) *ListServiceInstancesResponseBodyServiceInstancesTags {
	s.Value = &v
	return s
}

type ListServiceInstancesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponse) SetHeaders(v map[string]*string) *ListServiceInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListServiceInstancesResponse) SetStatusCode(v int32) *ListServiceInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceInstancesResponse) SetBody(v *ListServiceInstancesResponseBody) *ListServiceInstancesResponse {
	s.Body = v
	return s
}

type ListServiceUsagesRequest struct {
	// The filters.
	Filter []*ListServiceUsagesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAWns8w4MmhzeptXVRG0PUEU=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListServiceUsagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceUsagesRequest) GoString() string {
	return s.String()
}

func (s *ListServiceUsagesRequest) SetFilter(v []*ListServiceUsagesRequestFilter) *ListServiceUsagesRequest {
	s.Filter = v
	return s
}

func (s *ListServiceUsagesRequest) SetMaxResults(v int32) *ListServiceUsagesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceUsagesRequest) SetNextToken(v string) *ListServiceUsagesRequest {
	s.NextToken = &v
	return s
}

type ListServiceUsagesRequestFilter struct {
	// The parameter name of the filter. You can specify one or more filters. Valid values:
	//
	// 	- ServiceId: the ID of the service.
	//
	// 	- ServiceName: the service name.
	//
	// 	- Status: the state of the service.
	//
	// 	- SupplierName: the name of the service provider.
	//
	// example:
	//
	// ServiceId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameter values of the filter.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListServiceUsagesRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListServiceUsagesRequestFilter) GoString() string {
	return s.String()
}

func (s *ListServiceUsagesRequestFilter) SetName(v string) *ListServiceUsagesRequestFilter {
	s.Name = &v
	return s
}

func (s *ListServiceUsagesRequestFilter) SetValue(v []*string) *ListServiceUsagesRequestFilter {
	s.Value = v
	return s
}

type ListServiceUsagesResponseBody struct {
	// The maximum number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAYChudnQUoBH+mGWFpb6oP0=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 18AD0960-A9FE-1AC8-ADF8-22131Fxxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service applications.
	ServiceUsages []*ListServiceUsagesResponseBodyServiceUsages `json:"ServiceUsages,omitempty" xml:"ServiceUsages,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServiceUsagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceUsagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceUsagesResponseBody) SetMaxResults(v int32) *ListServiceUsagesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceUsagesResponseBody) SetNextToken(v string) *ListServiceUsagesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceUsagesResponseBody) SetRequestId(v string) *ListServiceUsagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceUsagesResponseBody) SetServiceUsages(v []*ListServiceUsagesResponseBodyServiceUsages) *ListServiceUsagesResponseBody {
	s.ServiceUsages = v
	return s
}

func (s *ListServiceUsagesResponseBody) SetTotalCount(v int32) *ListServiceUsagesResponseBody {
	s.TotalCount = &v
	return s
}

type ListServiceUsagesResponseBodyServiceUsages struct {
	// The review comment.
	//
	// example:
	//
	// Approved
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The time when the application was created.
	//
	// example:
	//
	// 2022-05-25T02:02:02Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-c9f36ec6d19b4exxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service name.
	//
	// example:
	//
	// LobelChat
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The state of the service application. Valid values:
	//
	// 	- Submitted: The application is submitted for review.
	//
	// 	- Approved: The application is approved.
	//
	// 	- Rejected: The application is rejected.
	//
	// 	- Canceled: The application is canceled.
	//
	// example:
	//
	// Submitted
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the service provider.
	//
	// example:
	//
	// TestSupplier
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// The time when the application was updated.
	//
	// example:
	//
	// 2022-05-25T02:02:02Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 127383705958xxxx
	UserAliUid *int64 `json:"UserAliUid,omitempty" xml:"UserAliUid,omitempty"`
	// The information about the applicants.
	UserInformation map[string]*string `json:"UserInformation,omitempty" xml:"UserInformation,omitempty"`
}

func (s ListServiceUsagesResponseBodyServiceUsages) String() string {
	return tea.Prettify(s)
}

func (s ListServiceUsagesResponseBodyServiceUsages) GoString() string {
	return s.String()
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetComments(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.Comments = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetCreateTime(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.CreateTime = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetServiceId(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.ServiceId = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetServiceName(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.ServiceName = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetStatus(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.Status = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetSupplierName(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.SupplierName = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetUpdateTime(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.UpdateTime = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetUserAliUid(v int64) *ListServiceUsagesResponseBodyServiceUsages {
	s.UserAliUid = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetUserInformation(v map[string]*string) *ListServiceUsagesResponseBodyServiceUsages {
	s.UserInformation = v
	return s
}

type ListServiceUsagesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceUsagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceUsagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceUsagesResponse) GoString() string {
	return s.String()
}

func (s *ListServiceUsagesResponse) SetHeaders(v map[string]*string) *ListServiceUsagesResponse {
	s.Headers = v
	return s
}

func (s *ListServiceUsagesResponse) SetStatusCode(v int32) *ListServiceUsagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceUsagesResponse) SetBody(v *ListServiceUsagesResponseBody) *ListServiceUsagesResponse {
	s.Body = v
	return s
}

type ListServicesRequest struct {
	// The filter.
	Filter []*ListServicesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// Keyword fuzzy query.
	//
	// example:
	//
	// name
	FuzzyKeyword *string `json:"FuzzyKeyword,omitempty" xml:"FuzzyKeyword,omitempty"`
	// Whether it is used. Optional values:
	//
	//
	//
	// - false: not being used.
	//
	//
	//
	// - true: already in use.
	//
	// example:
	//
	// false
	InUsed *bool `json:"InUsed,omitempty" xml:"InUsed,omitempty"`
	// The number of entries page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// BBBAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Service ordering type.
	//
	// example:
	//
	// UpdateTime
	OrderByType *string `json:"OrderByType,omitempty" xml:"OrderByType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Service access type.
	//
	// example:
	//
	// All
	ServiceAccessType *string `json:"ServiceAccessType,omitempty" xml:"ServiceAccessType,omitempty"`
	// The tags.
	Tag []*ListServicesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServicesRequest) GoString() string {
	return s.String()
}

func (s *ListServicesRequest) SetFilter(v []*ListServicesRequestFilter) *ListServicesRequest {
	s.Filter = v
	return s
}

func (s *ListServicesRequest) SetFuzzyKeyword(v string) *ListServicesRequest {
	s.FuzzyKeyword = &v
	return s
}

func (s *ListServicesRequest) SetInUsed(v bool) *ListServicesRequest {
	s.InUsed = &v
	return s
}

func (s *ListServicesRequest) SetMaxResults(v int32) *ListServicesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServicesRequest) SetNextToken(v string) *ListServicesRequest {
	s.NextToken = &v
	return s
}

func (s *ListServicesRequest) SetOrderByType(v string) *ListServicesRequest {
	s.OrderByType = &v
	return s
}

func (s *ListServicesRequest) SetRegionId(v string) *ListServicesRequest {
	s.RegionId = &v
	return s
}

func (s *ListServicesRequest) SetServiceAccessType(v string) *ListServicesRequest {
	s.ServiceAccessType = &v
	return s
}

func (s *ListServicesRequest) SetTag(v []*ListServicesRequestTag) *ListServicesRequest {
	s.Tag = v
	return s
}

type ListServicesRequestFilter struct {
	// The parameter name of the filter. You can specify one or more parameter names to query services. Valid values:
	//
	// 	- ServiceId: the ID of the service.
	//
	// 	- Name: the name of the service.
	//
	// 	- Status: the state of the service.
	//
	// 	- SupplierName: the name of the service provider.
	//
	// example:
	//
	// Status
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// A value of the filter condition.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListServicesRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListServicesRequestFilter) GoString() string {
	return s.String()
}

func (s *ListServicesRequestFilter) SetName(v string) *ListServicesRequestFilter {
	s.Name = &v
	return s
}

func (s *ListServicesRequestFilter) SetValue(v []*string) *ListServicesRequestFilter {
	s.Value = v
	return s
}

type ListServicesRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServicesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListServicesRequestTag) GoString() string {
	return s.String()
}

func (s *ListServicesRequestTag) SetKey(v string) *ListServicesRequestTag {
	s.Key = &v
	return s
}

func (s *ListServicesRequestTag) SetValue(v string) *ListServicesRequestTag {
	s.Value = &v
	return s
}

type ListServicesResponseBody struct {
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token.
	//
	// example:
	//
	// AAAAAfu+XtuBE55iRLHEYYuojI41
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3F976EF8-C10A-57DC-917C-BB7BEB508FFB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The services.
	Services []*ListServicesResponseBodyServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBody) SetMaxResults(v int32) *ListServicesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServicesResponseBody) SetNextToken(v string) *ListServicesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServicesResponseBody) SetRequestId(v string) *ListServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServicesResponseBody) SetServices(v []*ListServicesResponseBodyServices) *ListServicesResponseBody {
	s.Services = v
	return s
}

func (s *ListServicesResponseBody) SetTotalCount(v int32) *ListServicesResponseBody {
	s.TotalCount = &v
	return s
}

type ListServicesResponseBodyServices struct {
	// The category of the service.
	//
	// example:
	//
	// cloud_ssd
	Categories *string `json:"Categories,omitempty" xml:"Categories,omitempty"`
	// The commodity details.
	Commodity *ListServicesResponseBodyServicesCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// The commodity code of the service in Alibaba Cloud Marketplace.
	//
	// example:
	//
	// cmjj00xxxx
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// Service deployment approach. Valid values:
	//
	// - NoWhere
	//
	// - Marketplace
	//
	// - ComputeNest
	//
	// example:
	//
	// ComputeNest
	DeployFrom *string `json:"DeployFrom,omitempty" xml:"DeployFrom,omitempty"`
	// The deployment type of the service. Valid values:
	//
	// 	- ros: The service is deployed by using Resource Orchestration Service (ROS).
	//
	// 	- terraform: The service is deployed by using Terraform.
	//
	// example:
	//
	// ros
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// The time when the service was published.
	//
	// example:
	//
	// 2022-01-21T10:35:44Z
	PublishTime *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	// Service recommendation score.
	//
	// example:
	//
	// 5
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-6b5d632edd394dxxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The information about the service.
	ServiceInfos []*ListServicesResponseBodyServicesServiceInfos `json:"ServiceInfos,omitempty" xml:"ServiceInfos,omitempty" type:"Repeated"`
	// The URL of the service page.
	//
	// example:
	//
	// http://example1.com
	ServiceProductUrl *string `json:"ServiceProductUrl,omitempty" xml:"ServiceProductUrl,omitempty"`
	// The type of the service. Valid values:
	//
	// - private: The service is a private service and is deployed within the account of a customer.
	//
	// - managed: The service is a fully managed service and is deployed within the account of a service provider.
	//
	// - operation: The service is a hosted O&M service.
	//
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The state of the service. Valid values:
	//
	// 	- Draft: The service is a draft.
	//
	// 	- Submitted: The service is submitted for review. You cannot modify services in this state.
	//
	// 	- Approved: The service is approved. You cannot modify services in this state. You can publish services in this state.
	//
	// 	- Launching: The service is being published.
	//
	// 	- Online: The service is published.
	//
	// 	- Offline: The service is unpublished.
	//
	// example:
	//
	// Online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the service provider.
	//
	// example:
	//
	// Alibaba Cloud
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// The name of service provider.
	//
	// example:
	//
	// Alibaba Cloud
	SupplierNameEng *string `json:"SupplierNameEng,omitempty" xml:"SupplierNameEng,omitempty"`
	// The Alibaba Cloud account ID of the service provider.
	//
	// example:
	//
	// 1436322xxxxx
	SupplierUid *int64 `json:"SupplierUid,omitempty" xml:"SupplierUid,omitempty"`
	// The URL of the service provider.
	//
	// example:
	//
	// http://example.com
	SupplierUrl *string `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	// The tags.
	Tags []*ListServicesResponseBodyServicesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The tenant type of the managed service. Valid values:
	//
	// 	- SingleTenant
	//
	// 	- MultiTenant
	//
	// example:
	//
	// SingleTenant
	TenantType *string `json:"TenantType,omitempty" xml:"TenantType,omitempty"`
	// The trial duration. Unit: day. The maximum trial duration cannot exceed 30 days.
	//
	// example:
	//
	// 7
	TrialDuration *string `json:"TrialDuration,omitempty" xml:"TrialDuration,omitempty"`
	// The trial policy. Valid values:
	//
	// 	- Trial: Trials are supported.
	//
	// 	- NotTrial: Trials are not supported.
	//
	// example:
	//
	// Trial
	TrialType *string `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
	// The version of the service.
	//
	// example:
	//
	// 4
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The custom version name defined by the service provider.
	//
	// example:
	//
	// v2.0.0
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
	// Indicates whether the service is a virtual Internet service. Valid values:
	//
	// 	- false
	//
	// 	- true
	//
	// example:
	//
	// false
	VirtualInternetService *string `json:"VirtualInternetService,omitempty" xml:"VirtualInternetService,omitempty"`
}

func (s ListServicesResponseBodyServices) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBodyServices) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyServices) SetCategories(v string) *ListServicesResponseBodyServices {
	s.Categories = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetCommodity(v *ListServicesResponseBodyServicesCommodity) *ListServicesResponseBodyServices {
	s.Commodity = v
	return s
}

func (s *ListServicesResponseBodyServices) SetCommodityCode(v string) *ListServicesResponseBodyServices {
	s.CommodityCode = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetDeployFrom(v string) *ListServicesResponseBodyServices {
	s.DeployFrom = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetDeployType(v string) *ListServicesResponseBodyServices {
	s.DeployType = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetPublishTime(v string) *ListServicesResponseBodyServices {
	s.PublishTime = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetScore(v int32) *ListServicesResponseBodyServices {
	s.Score = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetServiceId(v string) *ListServicesResponseBodyServices {
	s.ServiceId = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetServiceInfos(v []*ListServicesResponseBodyServicesServiceInfos) *ListServicesResponseBodyServices {
	s.ServiceInfos = v
	return s
}

func (s *ListServicesResponseBodyServices) SetServiceProductUrl(v string) *ListServicesResponseBodyServices {
	s.ServiceProductUrl = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetServiceType(v string) *ListServicesResponseBodyServices {
	s.ServiceType = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetStatus(v string) *ListServicesResponseBodyServices {
	s.Status = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetSupplierName(v string) *ListServicesResponseBodyServices {
	s.SupplierName = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetSupplierNameEng(v string) *ListServicesResponseBodyServices {
	s.SupplierNameEng = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetSupplierUid(v int64) *ListServicesResponseBodyServices {
	s.SupplierUid = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetSupplierUrl(v string) *ListServicesResponseBodyServices {
	s.SupplierUrl = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetTags(v []*ListServicesResponseBodyServicesTags) *ListServicesResponseBodyServices {
	s.Tags = v
	return s
}

func (s *ListServicesResponseBodyServices) SetTenantType(v string) *ListServicesResponseBodyServices {
	s.TenantType = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetTrialDuration(v string) *ListServicesResponseBodyServices {
	s.TrialDuration = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetTrialType(v string) *ListServicesResponseBodyServices {
	s.TrialType = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetVersion(v string) *ListServicesResponseBodyServices {
	s.Version = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetVersionName(v string) *ListServicesResponseBodyServices {
	s.VersionName = &v
	return s
}

func (s *ListServicesResponseBodyServices) SetVirtualInternetService(v string) *ListServicesResponseBodyServices {
	s.VirtualInternetService = &v
	return s
}

type ListServicesResponseBodyServicesCommodity struct {
	// The commodity code.
	//
	// example:
	//
	// cmjj00****
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// Deploy Page.
	//
	// example:
	//
	// Order： Order Page
	//
	// Detail： Detail Page
	DeployPage *string `json:"DeployPage,omitempty" xml:"DeployPage,omitempty"`
}

func (s ListServicesResponseBodyServicesCommodity) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBodyServicesCommodity) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyServicesCommodity) SetCommodityCode(v string) *ListServicesResponseBodyServicesCommodity {
	s.CommodityCode = &v
	return s
}

func (s *ListServicesResponseBodyServicesCommodity) SetDeployPage(v string) *ListServicesResponseBodyServicesCommodity {
	s.DeployPage = &v
	return s
}

type ListServicesResponseBodyServicesServiceInfos struct {
	// The URL of the service icon.
	//
	// example:
	//
	// http://img.tidb.oss.url
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The language of the service. Valid values:
	//
	// 	- zh-CN: Chinese.
	//
	// 	- en-US: English.
	//
	// example:
	//
	// zh-CN
	Locale *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// Docker Community Edition
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The description of the service.
	//
	// example:
	//
	// Docker Community Edition (CE) is a free version of the Docker project, aimed at developers, enthusiasts, and individuals and organizations who want to use container technology.
	ShortDescription *string                                                  `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
	Softwares        []*ListServicesResponseBodyServicesServiceInfosSoftwares `json:"Softwares,omitempty" xml:"Softwares,omitempty" type:"Repeated"`
}

func (s ListServicesResponseBodyServicesServiceInfos) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBodyServicesServiceInfos) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyServicesServiceInfos) SetImage(v string) *ListServicesResponseBodyServicesServiceInfos {
	s.Image = &v
	return s
}

func (s *ListServicesResponseBodyServicesServiceInfos) SetLocale(v string) *ListServicesResponseBodyServicesServiceInfos {
	s.Locale = &v
	return s
}

func (s *ListServicesResponseBodyServicesServiceInfos) SetName(v string) *ListServicesResponseBodyServicesServiceInfos {
	s.Name = &v
	return s
}

func (s *ListServicesResponseBodyServicesServiceInfos) SetShortDescription(v string) *ListServicesResponseBodyServicesServiceInfos {
	s.ShortDescription = &v
	return s
}

func (s *ListServicesResponseBodyServicesServiceInfos) SetSoftwares(v []*ListServicesResponseBodyServicesServiceInfosSoftwares) *ListServicesResponseBodyServicesServiceInfos {
	s.Softwares = v
	return s
}

type ListServicesResponseBodyServicesServiceInfosSoftwares struct {
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListServicesResponseBodyServicesServiceInfosSoftwares) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBodyServicesServiceInfosSoftwares) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyServicesServiceInfosSoftwares) SetName(v string) *ListServicesResponseBodyServicesServiceInfosSoftwares {
	s.Name = &v
	return s
}

func (s *ListServicesResponseBodyServicesServiceInfosSoftwares) SetVersion(v string) *ListServicesResponseBodyServicesServiceInfosSoftwares {
	s.Version = &v
	return s
}

type ListServicesResponseBodyServicesTags struct {
	// The tag key.
	//
	// >  This parameter is required.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServicesResponseBodyServicesTags) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBodyServicesTags) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyServicesTags) SetKey(v string) *ListServicesResponseBodyServicesTags {
	s.Key = &v
	return s
}

func (s *ListServicesResponseBodyServicesTags) SetValue(v string) *ListServicesResponseBodyServicesTags {
	s.Value = &v
	return s
}

type ListServicesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponse) GoString() string {
	return s.String()
}

func (s *ListServicesResponse) SetHeaders(v map[string]*string) *ListServicesResponse {
	s.Headers = v
	return s
}

func (s *ListServicesResponse) SetStatusCode(v int32) *ListServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServicesResponse) SetBody(v *ListServicesResponseBody) *ListServicesResponse {
	s.Body = v
	return s
}

type ListTagKeysRequest struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource type. Valid values:
	//
	// 	- service: service
	//
	// 	- serviceinstance: service instance
	//
	// 	- artifact: artifact
	//
	// 	- dataset: dataset
	//
	// This parameter is required.
	//
	// example:
	//
	// service
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetNextToken(v string) *ListTagKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysRequest) SetRegionId(v string) *ListTagKeysRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceType(v string) *ListTagKeysRequest {
	s.ResourceType = &v
	return s
}

type ListTagKeysResponseBody struct {
	// Details of the tag keys.
	Keys []*string `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAZ9FmxgN6wKfeK/GOKRnnjU=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8872ACE6-0297-54A4-8AAD-3A8623EC6C5D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) SetKeys(v []*string) *ListTagKeysResponseBody {
	s.Keys = v
	return s
}

func (s *ListTagKeysResponseBody) SetNextToken(v string) *ListTagKeysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

type ListTagKeysResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponse) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponse) SetHeaders(v map[string]*string) *ListTagKeysResponse {
	s.Headers = v
	return s
}

func (s *ListTagKeysResponse) SetStatusCode(v int32) *ListTagKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagKeysResponse) SetBody(v *ListTagKeysResponseBody) *ListTagKeysResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource IDs. You can specify up to 50 resource IDs.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type. Valid values:
	//
	// 	- service: service
	//
	// 	- serviceinstance: service instance
	//
	// 	- artifact: artifact
	//
	// This parameter is required.
	//
	// example:
	//
	// service
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAVz7BQqj2xtiNSC3d3RAD38=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06BF8F22-02DC-4750-83DF-3FFC11C065EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of resources that have tags.
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	// The resource ID.
	//
	// example:
	//
	// si-44b9923be2d048eb8f5f
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type. Valid values:
	//
	// 	- service: service
	//
	// 	- serviceinstance: service instance
	//
	// 	- artifact: artifact
	//
	// example:
	//
	// serviceinstance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of the tag.
	//
	// example:
	//
	// test
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The values of the tags.
	//
	// example:
	//
	// major
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ListTagValuesRequest struct {
	// The tag key.
	//
	// >  This parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAfmTH5rcd4YFfob4P0uDAAc=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource type. Valid values:
	//
	// 	- service: service
	//
	// 	- serviceinstance: service instance
	//
	// 	- artifact: artifact
	//
	// 	- dataset: dataset
	//
	// This parameter is required.
	//
	// example:
	//
	// service
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListTagValuesRequest) SetKey(v string) *ListTagValuesRequest {
	s.Key = &v
	return s
}

func (s *ListTagValuesRequest) SetNextToken(v string) *ListTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesRequest) SetRegionId(v string) *ListTagValuesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceType(v string) *ListTagValuesRequest {
	s.ResourceType = &v
	return s
}

type ListTagValuesResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAVz7BQqj2xtiNSC3d3RAD38=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0631D623-D917-1C2D-ACD6-5B3B19XXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the tag values.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListTagValuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponseBody) SetNextToken(v string) *ListTagValuesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesResponseBody) SetRequestId(v string) *ListTagValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagValuesResponseBody) SetValues(v []*string) *ListTagValuesResponseBody {
	s.Values = v
	return s
}

type ListTagValuesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagValuesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponse) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponse) SetHeaders(v map[string]*string) *ListTagValuesResponse {
	s.Headers = v
	return s
}

func (s *ListTagValuesResponse) SetStatusCode(v int32) *ListTagValuesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagValuesResponse) SetBody(v *ListTagValuesResponseBody) *ListTagValuesResponse {
	s.Body = v
	return s
}

type RenewServiceInstanceResourcesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The renewal duration for all prepaid resources in the service instance. The unit is specified by PeriodUnit.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The time unit for the renewal duration of all prepaid resources in the service instance, which is the unit for the Period parameter. Valid values: Month, Year. Default value: Month.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// List of resource renewals.
	ResourcePeriod []*RenewServiceInstanceResourcesRequestResourcePeriod `json:"ResourcePeriod,omitempty" xml:"ResourcePeriod,omitempty" type:"Repeated"`
	// Service instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-b58c874912fc4294****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s RenewServiceInstanceResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewServiceInstanceResourcesRequest) GoString() string {
	return s.String()
}

func (s *RenewServiceInstanceResourcesRequest) SetClientToken(v string) *RenewServiceInstanceResourcesRequest {
	s.ClientToken = &v
	return s
}

func (s *RenewServiceInstanceResourcesRequest) SetPeriod(v int32) *RenewServiceInstanceResourcesRequest {
	s.Period = &v
	return s
}

func (s *RenewServiceInstanceResourcesRequest) SetPeriodUnit(v string) *RenewServiceInstanceResourcesRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RenewServiceInstanceResourcesRequest) SetRegionId(v string) *RenewServiceInstanceResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *RenewServiceInstanceResourcesRequest) SetResourcePeriod(v []*RenewServiceInstanceResourcesRequestResourcePeriod) *RenewServiceInstanceResourcesRequest {
	s.ResourcePeriod = v
	return s
}

func (s *RenewServiceInstanceResourcesRequest) SetServiceInstanceId(v string) *RenewServiceInstanceResourcesRequest {
	s.ServiceInstanceId = &v
	return s
}

type RenewServiceInstanceResourcesRequestResourcePeriod struct {
	// The renewal duration for the resource. The unit is specified by PeriodUnit.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The time unit for the renewal duration of the resource, which is the unit for the Period parameter. Valid values: Month, Year. Default value: Month.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// Resource ARN (Aliyun Resource Name).
	//
	// example:
	//
	// acs:ecs:cn-hongkong:1488317743351199:instance/i-j6c6f3lbky38o8rpeqw2
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
}

func (s RenewServiceInstanceResourcesRequestResourcePeriod) String() string {
	return tea.Prettify(s)
}

func (s RenewServiceInstanceResourcesRequestResourcePeriod) GoString() string {
	return s.String()
}

func (s *RenewServiceInstanceResourcesRequestResourcePeriod) SetPeriod(v int32) *RenewServiceInstanceResourcesRequestResourcePeriod {
	s.Period = &v
	return s
}

func (s *RenewServiceInstanceResourcesRequestResourcePeriod) SetPeriodUnit(v string) *RenewServiceInstanceResourcesRequestResourcePeriod {
	s.PeriodUnit = &v
	return s
}

func (s *RenewServiceInstanceResourcesRequestResourcePeriod) SetResourceArn(v string) *RenewServiceInstanceResourcesRequestResourcePeriod {
	s.ResourceArn = &v
	return s
}

type RenewServiceInstanceResourcesResponseBody struct {
	// Details of failed renewals.
	FailureDetails []*RenewServiceInstanceResourcesResponseBodyFailureDetails `json:"FailureDetails,omitempty" xml:"FailureDetails,omitempty" type:"Repeated"`
	// Renewal result.
	RenewalResult *RenewServiceInstanceResourcesResponseBodyRenewalResult `json:"RenewalResult,omitempty" xml:"RenewalResult,omitempty" type:"Struct"`
	// Request ID.
	//
	// example:
	//
	// 13FE89A5-C036-56BF-A0FF-A31C59819FD7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewServiceInstanceResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewServiceInstanceResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *RenewServiceInstanceResourcesResponseBody) SetFailureDetails(v []*RenewServiceInstanceResourcesResponseBodyFailureDetails) *RenewServiceInstanceResourcesResponseBody {
	s.FailureDetails = v
	return s
}

func (s *RenewServiceInstanceResourcesResponseBody) SetRenewalResult(v *RenewServiceInstanceResourcesResponseBodyRenewalResult) *RenewServiceInstanceResourcesResponseBody {
	s.RenewalResult = v
	return s
}

func (s *RenewServiceInstanceResourcesResponseBody) SetRequestId(v string) *RenewServiceInstanceResourcesResponseBody {
	s.RequestId = &v
	return s
}

type RenewServiceInstanceResourcesResponseBodyFailureDetails struct {
	// Error code.
	//
	// example:
	//
	// InvalidPeriod
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Error message.
	//
	// example:
	//
	// Error message
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Resource ARN (Aliyun Resource Name).
	//
	// example:
	//
	// acs:ecs:cn-hongkong:1488317743351199:instance/i-j6c6f3lbky38o8rpeqw2
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
}

func (s RenewServiceInstanceResourcesResponseBodyFailureDetails) String() string {
	return tea.Prettify(s)
}

func (s RenewServiceInstanceResourcesResponseBodyFailureDetails) GoString() string {
	return s.String()
}

func (s *RenewServiceInstanceResourcesResponseBodyFailureDetails) SetErrorCode(v string) *RenewServiceInstanceResourcesResponseBodyFailureDetails {
	s.ErrorCode = &v
	return s
}

func (s *RenewServiceInstanceResourcesResponseBodyFailureDetails) SetErrorMessage(v string) *RenewServiceInstanceResourcesResponseBodyFailureDetails {
	s.ErrorMessage = &v
	return s
}

func (s *RenewServiceInstanceResourcesResponseBodyFailureDetails) SetResourceArn(v string) *RenewServiceInstanceResourcesResponseBodyFailureDetails {
	s.ResourceArn = &v
	return s
}

type RenewServiceInstanceResourcesResponseBodyRenewalResult struct {
	// Number of failed renewals.
	//
	// example:
	//
	// 1
	Failed *int32 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// Number of successfully renewed resources.
	//
	// example:
	//
	// 9
	Succeeded *int32 `json:"Succeeded,omitempty" xml:"Succeeded,omitempty"`
	// Number of renewed resources.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s RenewServiceInstanceResourcesResponseBodyRenewalResult) String() string {
	return tea.Prettify(s)
}

func (s RenewServiceInstanceResourcesResponseBodyRenewalResult) GoString() string {
	return s.String()
}

func (s *RenewServiceInstanceResourcesResponseBodyRenewalResult) SetFailed(v int32) *RenewServiceInstanceResourcesResponseBodyRenewalResult {
	s.Failed = &v
	return s
}

func (s *RenewServiceInstanceResourcesResponseBodyRenewalResult) SetSucceeded(v int32) *RenewServiceInstanceResourcesResponseBodyRenewalResult {
	s.Succeeded = &v
	return s
}

func (s *RenewServiceInstanceResourcesResponseBodyRenewalResult) SetTotalCount(v int32) *RenewServiceInstanceResourcesResponseBodyRenewalResult {
	s.TotalCount = &v
	return s
}

type RenewServiceInstanceResourcesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewServiceInstanceResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewServiceInstanceResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewServiceInstanceResourcesResponse) GoString() string {
	return s.String()
}

func (s *RenewServiceInstanceResourcesResponse) SetHeaders(v map[string]*string) *RenewServiceInstanceResourcesResponse {
	s.Headers = v
	return s
}

func (s *RenewServiceInstanceResourcesResponse) SetStatusCode(v int32) *RenewServiceInstanceResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewServiceInstanceResourcesResponse) SetBody(v *RenewServiceInstanceResourcesResponseBody) *RenewServiceInstanceResourcesResponse {
	s.Body = v
	return s
}

type RestartServiceInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID where the service instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s RestartServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartServiceInstanceRequest) SetClientToken(v string) *RestartServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RestartServiceInstanceRequest) SetRegionId(v string) *RestartServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *RestartServiceInstanceRequest) SetServiceInstanceId(v string) *RestartServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

type RestartServiceInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 13FE89A5-C036-56BF-A0FF-A31C59819FD7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartServiceInstanceResponseBody) SetRequestId(v string) *RestartServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RestartServiceInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestartServiceInstanceResponse) SetHeaders(v map[string]*string) *RestartServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *RestartServiceInstanceResponse) SetStatusCode(v int32) *RestartServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartServiceInstanceResponse) SetBody(v *RestartServiceInstanceResponseBody) *RestartServiceInstanceResponse {
	s.Body = v
	return s
}

type RollbackServiceInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service instance ID.
	//
	// example:
	//
	// si-d6ab3a63ccbb4bxxxxxx
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s RollbackServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RollbackServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *RollbackServiceInstanceRequest) SetClientToken(v string) *RollbackServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RollbackServiceInstanceRequest) SetRegionId(v string) *RollbackServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *RollbackServiceInstanceRequest) SetServiceInstanceId(v string) *RollbackServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

type RollbackServiceInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RollbackServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RollbackServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackServiceInstanceResponseBody) SetRequestId(v string) *RollbackServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RollbackServiceInstanceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RollbackServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RollbackServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RollbackServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *RollbackServiceInstanceResponse) SetHeaders(v map[string]*string) *RollbackServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *RollbackServiceInstanceResponse) SetStatusCode(v int32) *RollbackServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackServiceInstanceResponse) SetBody(v *RollbackServiceInstanceResponseBody) *RollbackServiceInstanceResponse {
	s.Body = v
	return s
}

type StartServiceInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID where the service instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s StartServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartServiceInstanceRequest) SetClientToken(v string) *StartServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *StartServiceInstanceRequest) SetRegionId(v string) *StartServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *StartServiceInstanceRequest) SetServiceInstanceId(v string) *StartServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

type StartServiceInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 464C8CB6-A548-5206-B83C-D32A8E43EC21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartServiceInstanceResponseBody) SetRequestId(v string) *StartServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

type StartServiceInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartServiceInstanceResponse) SetHeaders(v map[string]*string) *StartServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartServiceInstanceResponse) SetStatusCode(v int32) *StartServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartServiceInstanceResponse) SetBody(v *StartServiceInstanceResponseBody) *StartServiceInstanceResponse {
	s.Body = v
	return s
}

type StopServiceInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region where the service instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-b58c874912fc4294****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s StopServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StopServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopServiceInstanceRequest) SetClientToken(v string) *StopServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *StopServiceInstanceRequest) SetRegionId(v string) *StopServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *StopServiceInstanceRequest) SetServiceInstanceId(v string) *StopServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

type StopServiceInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopServiceInstanceResponseBody) SetRequestId(v string) *StopServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

type StopServiceInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StopServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopServiceInstanceResponse) SetHeaders(v map[string]*string) *StopServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopServiceInstanceResponse) SetStatusCode(v int32) *StopServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopServiceInstanceResponse) SetBody(v *StopServiceInstanceResponseBody) *StopServiceInstanceResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource IDs. You can specify up to 50 resource IDs.
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type. Valid values:
	//
	// 	- service: service
	//
	// 	- serviceinstance: service instance
	//
	// 	- artifact: artifact
	//
	// 	- dataset: dataset
	//
	// This parameter is required.
	//
	// example:
	//
	// service
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag key and value.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// Key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// Value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 06BF8F22-02DC-4750-83DF-3FFC11C065EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UnTagResourcesRequest struct {
	// Specifies whether to remove all tags from the resource. Valid values:
	//
	// 	- true: All tags are removed from the resource.
	//
	// 	- false (default): The specified tags are removed from the resource.
	//
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource ID.
	//
	// You can remove tags from up to 50 resources at a time.
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type. Valid values:
	//
	// 	- service: service
	//
	// 	- serviceinstance: service instance
	//
	// 	- artifact: artifact
	//
	// 	- dataset: dataset
	//
	// This parameter is required.
	//
	// example:
	//
	// service
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag keys.
	//
	// You can specify a maximum of 20 tag keys.
	//
	// > If you set the `All` parameter to `true`, you do not need to specify tag keys.
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UnTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UnTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UnTagResourcesRequest) SetAll(v bool) *UnTagResourcesRequest {
	s.All = &v
	return s
}

func (s *UnTagResourcesRequest) SetRegionId(v string) *UnTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UnTagResourcesRequest) SetResourceId(v []*string) *UnTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UnTagResourcesRequest) SetResourceType(v string) *UnTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UnTagResourcesRequest) SetTagKey(v []*string) *UnTagResourcesRequest {
	s.TagKey = v
	return s
}

type UnTagResourcesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 06BF8F22-02DC-4750-83DF-3FFC11C065EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UnTagResourcesResponseBody) SetRequestId(v string) *UnTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UnTagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UnTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UnTagResourcesResponse) SetHeaders(v map[string]*string) *UnTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UnTagResourcesResponse) SetStatusCode(v int32) *UnTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UnTagResourcesResponse) SetBody(v *UnTagResourcesResponseBody) *UnTagResourcesResponse {
	s.Body = v
	return s
}

type UpdateServiceInstanceAttributesRequest struct {
	// Specifies whether to authorize the service provider to perform O\\&M operations on the service instance.
	//
	// example:
	//
	// true
	EnableOperation *bool `json:"EnableOperation,omitempty" xml:"EnableOperation,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the service instance.
	//
	// You can call the [ListServiceInstances](https://help.aliyun.com/document_detail/396200.html) operation to obtain the ID of the service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17xxxx
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s UpdateServiceInstanceAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceAttributesRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceAttributesRequest) SetEnableOperation(v bool) *UpdateServiceInstanceAttributesRequest {
	s.EnableOperation = &v
	return s
}

func (s *UpdateServiceInstanceAttributesRequest) SetRegionId(v string) *UpdateServiceInstanceAttributesRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateServiceInstanceAttributesRequest) SetServiceInstanceId(v string) *UpdateServiceInstanceAttributesRequest {
	s.ServiceInstanceId = &v
	return s
}

type UpdateServiceInstanceAttributesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 06BF8F22-02DC-4750-83DF-3FFC11C065EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceInstanceAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceAttributesResponseBody) SetRequestId(v string) *UpdateServiceInstanceAttributesResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServiceInstanceAttributesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceInstanceAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceInstanceAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceAttributesResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceAttributesResponse) SetHeaders(v map[string]*string) *UpdateServiceInstanceAttributesResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceInstanceAttributesResponse) SetStatusCode(v int32) *UpdateServiceInstanceAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceInstanceAttributesResponse) SetBody(v *UpdateServiceInstanceAttributesResponseBody) *UpdateServiceInstanceAttributesResponse {
	s.Body = v
	return s
}

type UpdateServiceInstanceSpecRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The information about the order placed in Alibaba Cloud Marketplace. You do not need to specify this parameter if the service is not published in Alibaba Cloud Marketplace or uses the pay-as-you-go billing method.
	Commodity *UpdateServiceInstanceSpecRequestCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// Specifies whether to perform only a dry run, without performing the actual request. A dry run includes checks on the permissions and instance state.
	//
	// Valid values:
	//
	// 	- true: performs a dry run but does not create a service instance.
	//
	// 	- false: performs a dry run and creates a service instance if the request passes the dry run.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to enable Prometheus monitoring on the user side.
	//
	// Valid values:
	//
	// true
	//
	// false
	//
	// example:
	//
	// true
	EnableUserPrometheus *bool `json:"EnableUserPrometheus,omitempty" xml:"EnableUserPrometheus,omitempty"`
	// The name of the configuration change operation.
	//
	// To obtain the names and content of the configuration change operations of the service, you can call the [GetService](https://help.aliyun.com/document_detail/2340828.html) operation. In the response, check the value of **ModifyParametersConfig*	- in the value of **OperationMetadata**.
	//
	// example:
	//
	// package modify
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	// The configuration parameter.
	//
	// This parameter is available if the service provider set **Method*	- to **Change Parameter*	- when configuring configuration change operations.
	//
	// >
	//
	// 	- To obtain the parameters of the service that support configuration change, you can call the [GetService](https://help.aliyun.com/document_detail/2340828.html) operation. In the response, check the value of **ModifyParametersConfig*	- in the value of **OperationMetadata**.
	//
	// 	- You can also view the parameters of the service that support configuration change in the **configuration change*	- dialog box in the [Compute Nest console](https://computenest.console.aliyun.com/service/instance/cn-hangzhou).
	//
	// For example, if the service supports Elastic Compute Service (ECS) instance type upgrade, you must specify an instance type that has higher specifications than the current one.
	//
	// example:
	//
	// {
	//
	//   "InstanceType": "ecs.g8ise.2xlarge"
	//
	// }
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The name of the configuration plan.
	//
	// This parameter is available if the service provider set **Method*	- to **Change Plan*	- when configuring configuration change operations.
	//
	// To obtain the configuration plan names of the service, you can call the [GetService](https://help.aliyun.com/document_detail/2340828.html) operation. In the response, check the value of **PredefinedParameters*	- in the value of **DeployMetadata**.
	//
	// example:
	//
	// package One
	PredefinedParametersName *string `json:"PredefinedParametersName,omitempty" xml:"PredefinedParametersName,omitempty"`
	// The ID of the service instance.
	//
	// You can call the [ListServiceInstances](https://help.aliyun.com/document_detail/396200.html) operation to obtain the ID of the service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s UpdateServiceInstanceSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceSpecRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceSpecRequest) SetClientToken(v string) *UpdateServiceInstanceSpecRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetCommodity(v *UpdateServiceInstanceSpecRequestCommodity) *UpdateServiceInstanceSpecRequest {
	s.Commodity = v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetDryRun(v bool) *UpdateServiceInstanceSpecRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetEnableUserPrometheus(v bool) *UpdateServiceInstanceSpecRequest {
	s.EnableUserPrometheus = &v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetOperationName(v string) *UpdateServiceInstanceSpecRequest {
	s.OperationName = &v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetParameters(v map[string]interface{}) *UpdateServiceInstanceSpecRequest {
	s.Parameters = v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetPredefinedParametersName(v string) *UpdateServiceInstanceSpecRequest {
	s.PredefinedParametersName = &v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetServiceInstanceId(v string) *UpdateServiceInstanceSpecRequest {
	s.ServiceInstanceId = &v
	return s
}

type UpdateServiceInstanceSpecRequestCommodity struct {
	// Specifies whether to enable automatic payment.
	//
	// Valid values:
	//
	// 	- **true (default)**: automatically completes the payment. You must make sure that your account balance is sufficient.
	//
	// 	- **false**: does not automatically complete the payment. An unpaid order is generated. If your account balance is insufficient, you can set AutoPay to false. In this case, an unpaid order is generated. You can complete the payment in the Expenses and Costs console.[](https://rdsnext.console.aliyun.com/dashboard/cn-beijing)
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
}

func (s UpdateServiceInstanceSpecRequestCommodity) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceSpecRequestCommodity) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceSpecRequestCommodity) SetAutoPay(v bool) *UpdateServiceInstanceSpecRequestCommodity {
	s.AutoPay = &v
	return s
}

type UpdateServiceInstanceSpecShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The information about the order placed in Alibaba Cloud Marketplace. You do not need to specify this parameter if the service is not published in Alibaba Cloud Marketplace or uses the pay-as-you-go billing method.
	Commodity *UpdateServiceInstanceSpecShrinkRequestCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// Specifies whether to perform only a dry run, without performing the actual request. A dry run includes checks on the permissions and instance state.
	//
	// Valid values:
	//
	// 	- true: performs a dry run but does not create a service instance.
	//
	// 	- false: performs a dry run and creates a service instance if the request passes the dry run.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to enable Prometheus monitoring on the user side.
	//
	// Valid values:
	//
	// true
	//
	// false
	//
	// example:
	//
	// true
	EnableUserPrometheus *bool `json:"EnableUserPrometheus,omitempty" xml:"EnableUserPrometheus,omitempty"`
	// The name of the configuration change operation.
	//
	// To obtain the names and content of the configuration change operations of the service, you can call the [GetService](https://help.aliyun.com/document_detail/2340828.html) operation. In the response, check the value of **ModifyParametersConfig*	- in the value of **OperationMetadata**.
	//
	// example:
	//
	// package modify
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	// The configuration parameter.
	//
	// This parameter is available if the service provider set **Method*	- to **Change Parameter*	- when configuring configuration change operations.
	//
	// >
	//
	// 	- To obtain the parameters of the service that support configuration change, you can call the [GetService](https://help.aliyun.com/document_detail/2340828.html) operation. In the response, check the value of **ModifyParametersConfig*	- in the value of **OperationMetadata**.
	//
	// 	- You can also view the parameters of the service that support configuration change in the **configuration change*	- dialog box in the [Compute Nest console](https://computenest.console.aliyun.com/service/instance/cn-hangzhou).
	//
	// For example, if the service supports Elastic Compute Service (ECS) instance type upgrade, you must specify an instance type that has higher specifications than the current one.
	//
	// example:
	//
	// {
	//
	//   "InstanceType": "ecs.g8ise.2xlarge"
	//
	// }
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The name of the configuration plan.
	//
	// This parameter is available if the service provider set **Method*	- to **Change Plan*	- when configuring configuration change operations.
	//
	// To obtain the configuration plan names of the service, you can call the [GetService](https://help.aliyun.com/document_detail/2340828.html) operation. In the response, check the value of **PredefinedParameters*	- in the value of **DeployMetadata**.
	//
	// example:
	//
	// package One
	PredefinedParametersName *string `json:"PredefinedParametersName,omitempty" xml:"PredefinedParametersName,omitempty"`
	// The ID of the service instance.
	//
	// You can call the [ListServiceInstances](https://help.aliyun.com/document_detail/396200.html) operation to obtain the ID of the service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s UpdateServiceInstanceSpecShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceSpecShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetClientToken(v string) *UpdateServiceInstanceSpecShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetCommodity(v *UpdateServiceInstanceSpecShrinkRequestCommodity) *UpdateServiceInstanceSpecShrinkRequest {
	s.Commodity = v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetDryRun(v bool) *UpdateServiceInstanceSpecShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetEnableUserPrometheus(v bool) *UpdateServiceInstanceSpecShrinkRequest {
	s.EnableUserPrometheus = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetOperationName(v string) *UpdateServiceInstanceSpecShrinkRequest {
	s.OperationName = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetParametersShrink(v string) *UpdateServiceInstanceSpecShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetPredefinedParametersName(v string) *UpdateServiceInstanceSpecShrinkRequest {
	s.PredefinedParametersName = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetServiceInstanceId(v string) *UpdateServiceInstanceSpecShrinkRequest {
	s.ServiceInstanceId = &v
	return s
}

type UpdateServiceInstanceSpecShrinkRequestCommodity struct {
	// Specifies whether to enable automatic payment.
	//
	// Valid values:
	//
	// 	- **true (default)**: automatically completes the payment. You must make sure that your account balance is sufficient.
	//
	// 	- **false**: does not automatically complete the payment. An unpaid order is generated. If your account balance is insufficient, you can set AutoPay to false. In this case, an unpaid order is generated. You can complete the payment in the Expenses and Costs console.[](https://rdsnext.console.aliyun.com/dashboard/cn-beijing)
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
}

func (s UpdateServiceInstanceSpecShrinkRequestCommodity) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceSpecShrinkRequestCommodity) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceSpecShrinkRequestCommodity) SetAutoPay(v bool) *UpdateServiceInstanceSpecShrinkRequestCommodity {
	s.AutoPay = &v
	return s
}

type UpdateServiceInstanceSpecResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// 23396265896****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06BF8F22-02DC-4750-83DF-3FFC11C065EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceInstanceSpecResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceSpecResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceSpecResponseBody) SetOrderId(v string) *UpdateServiceInstanceSpecResponseBody {
	s.OrderId = &v
	return s
}

func (s *UpdateServiceInstanceSpecResponseBody) SetRequestId(v string) *UpdateServiceInstanceSpecResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServiceInstanceSpecResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceInstanceSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceInstanceSpecResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceSpecResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceSpecResponse) SetHeaders(v map[string]*string) *UpdateServiceInstanceSpecResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceInstanceSpecResponse) SetStatusCode(v int32) *UpdateServiceInstanceSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceInstanceSpecResponse) SetBody(v *UpdateServiceInstanceSpecResponseBody) *UpdateServiceInstanceSpecResponse {
	s.Body = v
	return s
}

type UpdateServiceUsageRequest struct {
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// AAAAAYChudnQUoBH+mGWFpb6oP0=
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-39f4f251e94843xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The information about the applicant.
	UserInformation map[string]*string `json:"UserInformation,omitempty" xml:"UserInformation,omitempty"`
}

func (s UpdateServiceUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceUsageRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceUsageRequest) SetClientToken(v string) *UpdateServiceUsageRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServiceUsageRequest) SetServiceId(v string) *UpdateServiceUsageRequest {
	s.ServiceId = &v
	return s
}

func (s *UpdateServiceUsageRequest) SetUserInformation(v map[string]*string) *UpdateServiceUsageRequest {
	s.UserInformation = v
	return s
}

type UpdateServiceUsageShrinkRequest struct {
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// AAAAAYChudnQUoBH+mGWFpb6oP0=
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-39f4f251e94843xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The information about the applicant.
	UserInformationShrink *string `json:"UserInformation,omitempty" xml:"UserInformation,omitempty"`
}

func (s UpdateServiceUsageShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceUsageShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceUsageShrinkRequest) SetClientToken(v string) *UpdateServiceUsageShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServiceUsageShrinkRequest) SetServiceId(v string) *UpdateServiceUsageShrinkRequest {
	s.ServiceId = &v
	return s
}

func (s *UpdateServiceUsageShrinkRequest) SetUserInformationShrink(v string) *UpdateServiceUsageShrinkRequest {
	s.UserInformationShrink = &v
	return s
}

type UpdateServiceUsageResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceUsageResponseBody) SetRequestId(v string) *UpdateServiceUsageResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServiceUsageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceUsageResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceUsageResponse) SetHeaders(v map[string]*string) *UpdateServiceUsageResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceUsageResponse) SetStatusCode(v int32) *UpdateServiceUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceUsageResponse) SetBody(v *UpdateServiceUsageResponseBody) *UpdateServiceUsageResponse {
	s.Body = v
	return s
}

type UpdateUserInformationRequest struct {
	// The modified delivery settings.
	DeliverySettings *UpdateUserInformationRequestDeliverySettings `json:"DeliverySettings,omitempty" xml:"DeliverySettings,omitempty" type:"Struct"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateUserInformationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserInformationRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserInformationRequest) SetDeliverySettings(v *UpdateUserInformationRequestDeliverySettings) *UpdateUserInformationRequest {
	s.DeliverySettings = v
	return s
}

func (s *UpdateUserInformationRequest) SetRegionId(v string) *UpdateUserInformationRequest {
	s.RegionId = &v
	return s
}

type UpdateUserInformationRequestDeliverySettings struct {
	// Specifies whether to enable screencast delivery to OSS. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	ActiontrailDeliveryToOssEnabled *bool `json:"ActiontrailDeliveryToOssEnabled,omitempty" xml:"ActiontrailDeliveryToOssEnabled,omitempty"`
	// The name of the OSS bucket.
	//
	// example:
	//
	// "mybucket"
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// Specifies whether to enable screencast delivery to Object Storage Service (OSS). Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	OssEnabled *bool `json:"OssEnabled,omitempty" xml:"OssEnabled,omitempty"`
	// The number of days for which the screencasts are saved.
	//
	// example:
	//
	// 7
	OssExpirationDays *int64 `json:"OssExpirationDays,omitempty" xml:"OssExpirationDays,omitempty"`
	// The OSS path.
	//
	// example:
	//
	// "path1/path2/"
	OssPath *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
}

func (s UpdateUserInformationRequestDeliverySettings) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserInformationRequestDeliverySettings) GoString() string {
	return s.String()
}

func (s *UpdateUserInformationRequestDeliverySettings) SetActiontrailDeliveryToOssEnabled(v bool) *UpdateUserInformationRequestDeliverySettings {
	s.ActiontrailDeliveryToOssEnabled = &v
	return s
}

func (s *UpdateUserInformationRequestDeliverySettings) SetOssBucketName(v string) *UpdateUserInformationRequestDeliverySettings {
	s.OssBucketName = &v
	return s
}

func (s *UpdateUserInformationRequestDeliverySettings) SetOssEnabled(v bool) *UpdateUserInformationRequestDeliverySettings {
	s.OssEnabled = &v
	return s
}

func (s *UpdateUserInformationRequestDeliverySettings) SetOssExpirationDays(v int64) *UpdateUserInformationRequestDeliverySettings {
	s.OssExpirationDays = &v
	return s
}

func (s *UpdateUserInformationRequestDeliverySettings) SetOssPath(v string) *UpdateUserInformationRequestDeliverySettings {
	s.OssPath = &v
	return s
}

type UpdateUserInformationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 51945B04-6AA6-410D-93BA-236E0248B104
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUserInformationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserInformationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserInformationResponseBody) SetRequestId(v string) *UpdateUserInformationResponseBody {
	s.RequestId = &v
	return s
}

type UpdateUserInformationResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserInformationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserInformationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserInformationResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserInformationResponse) SetHeaders(v map[string]*string) *UpdateUserInformationResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserInformationResponse) SetStatusCode(v int32) *UpdateUserInformationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserInformationResponse) SetBody(v *UpdateUserInformationResponseBody) *UpdateUserInformationResponse {
	s.Body = v
	return s
}

type UpgradeServiceInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run for the request to check information such as the permissions and instance status. Valid values:
	//
	// 	- **true**: performs a dry run for the request, but does not upgrade service instance.
	//
	// 	- **false**: performs a dry run for the request, and upgrade service instance if the request passes the dry run.
	//
	// example:
	//
	// true
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The parameters required for the upgrade. This parameter is required if the destination version of the service has new parameters.
	//
	// example:
	//
	// { \\"RegionId\\": \\"cn-hangzhou\\", \\"InstanceType\\": \\"ecs.g5.large\\"}
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the service instance.
	//
	// example:
	//
	// si-d6ab3a63ccbb4bxxxxxx
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The destination version.
	//
	// example:
	//
	// 2
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s UpgradeServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpgradeServiceInstanceRequest) SetClientToken(v string) *UpgradeServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *UpgradeServiceInstanceRequest) SetDryRun(v string) *UpgradeServiceInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *UpgradeServiceInstanceRequest) SetParameters(v map[string]interface{}) *UpgradeServiceInstanceRequest {
	s.Parameters = v
	return s
}

func (s *UpgradeServiceInstanceRequest) SetRegionId(v string) *UpgradeServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeServiceInstanceRequest) SetServiceInstanceId(v string) *UpgradeServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *UpgradeServiceInstanceRequest) SetServiceVersion(v string) *UpgradeServiceInstanceRequest {
	s.ServiceVersion = &v
	return s
}

type UpgradeServiceInstanceShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run for the request to check information such as the permissions and instance status. Valid values:
	//
	// 	- **true**: performs a dry run for the request, but does not upgrade service instance.
	//
	// 	- **false**: performs a dry run for the request, and upgrade service instance if the request passes the dry run.
	//
	// example:
	//
	// true
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The parameters required for the upgrade. This parameter is required if the destination version of the service has new parameters.
	//
	// example:
	//
	// { \\"RegionId\\": \\"cn-hangzhou\\", \\"InstanceType\\": \\"ecs.g5.large\\"}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the service instance.
	//
	// example:
	//
	// si-d6ab3a63ccbb4bxxxxxx
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The destination version.
	//
	// example:
	//
	// 2
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s UpgradeServiceInstanceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeServiceInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpgradeServiceInstanceShrinkRequest) SetClientToken(v string) *UpgradeServiceInstanceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpgradeServiceInstanceShrinkRequest) SetDryRun(v string) *UpgradeServiceInstanceShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *UpgradeServiceInstanceShrinkRequest) SetParametersShrink(v string) *UpgradeServiceInstanceShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *UpgradeServiceInstanceShrinkRequest) SetRegionId(v string) *UpgradeServiceInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeServiceInstanceShrinkRequest) SetServiceInstanceId(v string) *UpgradeServiceInstanceShrinkRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *UpgradeServiceInstanceShrinkRequest) SetServiceVersion(v string) *UpgradeServiceInstanceShrinkRequest {
	s.ServiceVersion = &v
	return s
}

type UpgradeServiceInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The parameters required for the upgrade. This parameter is returned only if DryRun is set to true in the request. You can specify the required parameters based on the returned value when you perform an actual request.
	UpgradeRequiredParameters []*string `json:"UpgradeRequiredParameters,omitempty" xml:"UpgradeRequiredParameters,omitempty" type:"Repeated"`
}

func (s UpgradeServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeServiceInstanceResponseBody) SetRequestId(v string) *UpgradeServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeServiceInstanceResponseBody) SetUpgradeRequiredParameters(v []*string) *UpgradeServiceInstanceResponseBody {
	s.UpgradeRequiredParameters = v
	return s
}

type UpgradeServiceInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpgradeServiceInstanceResponse) SetHeaders(v map[string]*string) *UpgradeServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpgradeServiceInstanceResponse) SetStatusCode(v int32) *UpgradeServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeServiceInstanceResponse) SetBody(v *UpgradeServiceInstanceResponseBody) *UpgradeServiceInstanceResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("computenest"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels the application for using a service.
//
// @param request - CancelServiceUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelServiceUsageResponse
func (client *Client) CancelServiceUsageWithOptions(request *CancelServiceUsageRequest, runtime *util.RuntimeOptions) (_result *CancelServiceUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.NeedDelete)) {
		query["NeedDelete"] = request.NeedDelete
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelServiceUsage"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelServiceUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the application for using a service.
//
// @param request - CancelServiceUsageRequest
//
// @return CancelServiceUsageResponse
func (client *Client) CancelServiceUsage(request *CancelServiceUsageRequest) (_result *CancelServiceUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelServiceUsageResponse{}
	_body, _err := client.CancelServiceUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the resource group to which a cloud resource belongs.
//
// @param request - ChangeResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, runtime *util.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewResourceGroupId)) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeResourceGroup"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the resource group to which a cloud resource belongs.
//
// @param request - ChangeResourceGroupRequest
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroup(request *ChangeResourceGroupRequest) (_result *ChangeResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.ChangeResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 服务实例部署前的预检查
//
// @param request - CheckServiceDeployableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckServiceDeployableResponse
func (client *Client) CheckServiceDeployableWithOptions(request *CheckServiceDeployableRequest, runtime *util.RuntimeOptions) (_result *CheckServiceDeployableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PostPaidAmount)) {
		query["PostPaidAmount"] = request.PostPaidAmount
	}

	if !tea.BoolValue(util.IsUnset(request.PrePaidAmount)) {
		query["PrePaidAmount"] = request.PrePaidAmount
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !tea.BoolValue(util.IsUnset(request.TrialType)) {
		query["TrialType"] = request.TrialType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckServiceDeployable"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckServiceDeployableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务实例部署前的预检查
//
// @param request - CheckServiceDeployableRequest
//
// @return CheckServiceDeployableResponse
func (client *Client) CheckServiceDeployable(request *CheckServiceDeployableRequest) (_result *CheckServiceDeployableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckServiceDeployableResponse{}
	_body, _err := client.CheckServiceDeployableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Continues to deploy a service instance after the service instance failed to be deployed.
//
// Description:
//
// This operation is available only for service instances that belong to private services deployed by using Resource Orchestration Service (ROS).
//
// @param request - ContinueDeployServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ContinueDeployServiceInstanceResponse
func (client *Client) ContinueDeployServiceInstanceWithOptions(request *ContinueDeployServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *ContinueDeployServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Option)) {
		query["Option"] = request.Option
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ContinueDeployServiceInstance"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ContinueDeployServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Continues to deploy a service instance after the service instance failed to be deployed.
//
// Description:
//
// This operation is available only for service instances that belong to private services deployed by using Resource Orchestration Service (ROS).
//
// @param request - ContinueDeployServiceInstanceRequest
//
// @return ContinueDeployServiceInstanceResponse
func (client *Client) ContinueDeployServiceInstance(request *ContinueDeployServiceInstanceRequest) (_result *ContinueDeployServiceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ContinueDeployServiceInstanceResponse{}
	_body, _err := client.ContinueDeployServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates and deploys a service instance.
//
// @param tmpReq - CreateServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceInstanceResponse
func (client *Client) CreateServiceInstanceWithOptions(tmpReq *CreateServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateServiceInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Parameters)) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, tea.String("Parameters"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Commodity)) {
		query["Commodity"] = request.Commodity
	}

	if !tea.BoolValue(util.IsUnset(request.ContactGroup)) {
		query["ContactGroup"] = request.ContactGroup
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EnableInstanceOps)) {
		query["EnableInstanceOps"] = request.EnableInstanceOps
	}

	if !tea.BoolValue(util.IsUnset(request.EnableUserPrometheus)) {
		query["EnableUserPrometheus"] = request.EnableUserPrometheus
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OperationMetadata)) {
		query["OperationMetadata"] = request.OperationMetadata
	}

	if !tea.BoolValue(util.IsUnset(request.ParametersShrink)) {
		query["Parameters"] = request.ParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceAutoPay)) {
		query["ResourceAutoPay"] = request.ResourceAutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !tea.BoolValue(util.IsUnset(request.SpecificationCode)) {
		query["SpecificationCode"] = request.SpecificationCode
	}

	if !tea.BoolValue(util.IsUnset(request.SpecificationName)) {
		query["SpecificationName"] = request.SpecificationName
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TrialType)) {
		query["TrialType"] = request.TrialType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServiceInstance"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates and deploys a service instance.
//
// @param request - CreateServiceInstanceRequest
//
// @return CreateServiceInstanceResponse
func (client *Client) CreateServiceInstance(request *CreateServiceInstanceRequest) (_result *CreateServiceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServiceInstanceResponse{}
	_body, _err := client.CreateServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an application for using a service.
//
// @param tmpReq - CreateServiceUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceUsageResponse
func (client *Client) CreateServiceUsageWithOptions(tmpReq *CreateServiceUsageRequest, runtime *util.RuntimeOptions) (_result *CreateServiceUsageResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateServiceUsageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.UserInformation)) {
		request.UserInformationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInformation, tea.String("UserInformation"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserInformationShrink)) {
		query["UserInformation"] = request.UserInformationShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServiceUsage"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an application for using a service.
//
// @param request - CreateServiceUsageRequest
//
// @return CreateServiceUsageResponse
func (client *Client) CreateServiceUsage(request *CreateServiceUsageRequest) (_result *CreateServiceUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServiceUsageResponse{}
	_body, _err := client.CreateServiceUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete service instances.
//
// @param request - DeleteServiceInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServiceInstancesResponse
func (client *Client) DeleteServiceInstancesWithOptions(request *DeleteServiceInstancesRequest, runtime *util.RuntimeOptions) (_result *DeleteServiceInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteServiceInstances"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServiceInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete service instances.
//
// @param request - DeleteServiceInstancesRequest
//
// @return DeleteServiceInstancesResponse
func (client *Client) DeleteServiceInstances(request *DeleteServiceInstancesRequest) (_result *DeleteServiceInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteServiceInstancesResponse{}
	_body, _err := client.DeleteServiceInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deploy service instance in Created status.
//
// @param request - DeployServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployServiceInstanceResponse
func (client *Client) DeployServiceInstanceWithOptions(request *DeployServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *DeployServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeployServiceInstance"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeployServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deploy service instance in Created status.
//
// @param request - DeployServiceInstanceRequest
//
// @return DeployServiceInstanceResponse
func (client *Client) DeployServiceInstance(request *DeployServiceInstanceRequest) (_result *DeployServiceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeployServiceInstanceResponse{}
	_body, _err := client.DeployServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// List available regions.
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// List available regions.
//
// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 生成并校验服务创建stack所需要的权限
//
// @param request - GenerateServicePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateServicePolicyResponse
func (client *Client) GenerateServicePolicyWithOptions(request *GenerateServicePolicyRequest, runtime *util.RuntimeOptions) (_result *GenerateServicePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperationTypes)) {
		query["OperationTypes"] = request.OperationTypes
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TrialType)) {
		query["TrialType"] = request.TrialType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateServicePolicy"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateServicePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成并校验服务创建stack所需要的权限
//
// @param request - GenerateServicePolicyRequest
//
// @return GenerateServicePolicyResponse
func (client *Client) GenerateServicePolicy(request *GenerateServicePolicyRequest) (_result *GenerateServicePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateServicePolicyResponse{}
	_body, _err := client.GenerateServicePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a service.
//
// @param request - GetServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceResponse
func (client *Client) GetServiceWithOptions(request *GetServiceRequest, runtime *util.RuntimeOptions) (_result *GetServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ShowDetails)) {
		query["ShowDetails"] = request.ShowDetails
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetService"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a service.
//
// @param request - GetServiceRequest
//
// @return GetServiceResponse
func (client *Client) GetService(request *GetServiceRequest) (_result *GetServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetServiceResponse{}
	_body, _err := client.GetServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the estimated price for creating a service instance.
//
// @param tmpReq - GetServiceEstimateCostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceEstimateCostResponse
func (client *Client) GetServiceEstimateCostWithOptions(tmpReq *GetServiceEstimateCostRequest, runtime *util.RuntimeOptions) (_result *GetServiceEstimateCostResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetServiceEstimateCostShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Commodity)) {
		request.CommodityShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Commodity, tea.String("Commodity"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Parameters)) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, tea.String("Parameters"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CommodityShrink)) {
		query["Commodity"] = request.CommodityShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OperationName)) {
		query["OperationName"] = request.OperationName
	}

	if !tea.BoolValue(util.IsUnset(request.ParametersShrink)) {
		query["Parameters"] = request.ParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !tea.BoolValue(util.IsUnset(request.SpecificationName)) {
		query["SpecificationName"] = request.SpecificationName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TrialType)) {
		query["TrialType"] = request.TrialType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceEstimateCost"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceEstimateCostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the estimated price for creating a service instance.
//
// @param request - GetServiceEstimateCostRequest
//
// @return GetServiceEstimateCostResponse
func (client *Client) GetServiceEstimateCost(request *GetServiceEstimateCostRequest) (_result *GetServiceEstimateCostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetServiceEstimateCostResponse{}
	_body, _err := client.GetServiceEstimateCostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a service instance based on the region ID and the ID of the service instance or the Alibaba Cloud Marketplace instance. Information including the service status, template name, and involved resources are returned.
//
// @param request - GetServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceInstanceResponse
func (client *Client) GetServiceInstanceWithOptions(request *GetServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *GetServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MarketInstanceId)) {
		query["MarketInstanceId"] = request.MarketInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceInstance"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a service instance based on the region ID and the ID of the service instance or the Alibaba Cloud Marketplace instance. Information including the service status, template name, and involved resources are returned.
//
// @param request - GetServiceInstanceRequest
//
// @return GetServiceInstanceResponse
func (client *Client) GetServiceInstance(request *GetServiceInstanceRequest) (_result *GetServiceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetServiceInstanceResponse{}
	_body, _err := client.GetServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query renewal prices for prepaid resources of private deployment service instance. You can query renewal prices for all prepaid resources included in a service instance, or query renewal prices for specified resources. Only one of the two methods can be used.
//
// @param request - GetServiceInstanceSubscriptionEstimateCostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceInstanceSubscriptionEstimateCostResponse
func (client *Client) GetServiceInstanceSubscriptionEstimateCostWithOptions(request *GetServiceInstanceSubscriptionEstimateCostRequest, runtime *util.RuntimeOptions) (_result *GetServiceInstanceSubscriptionEstimateCostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrderType)) {
		query["OrderType"] = request.OrderType
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourcePeriod)) {
		query["ResourcePeriod"] = request.ResourcePeriod
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceInstanceSubscriptionEstimateCost"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceInstanceSubscriptionEstimateCostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query renewal prices for prepaid resources of private deployment service instance. You can query renewal prices for all prepaid resources included in a service instance, or query renewal prices for specified resources. Only one of the two methods can be used.
//
// @param request - GetServiceInstanceSubscriptionEstimateCostRequest
//
// @return GetServiceInstanceSubscriptionEstimateCostResponse
func (client *Client) GetServiceInstanceSubscriptionEstimateCost(request *GetServiceInstanceSubscriptionEstimateCostRequest) (_result *GetServiceInstanceSubscriptionEstimateCostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetServiceInstanceSubscriptionEstimateCostResponse{}
	_body, _err := client.GetServiceInstanceSubscriptionEstimateCostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 计算巢查询服务是否开通
//
// @param tmpReq - GetServiceProvisionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceProvisionsResponse
func (client *Client) GetServiceProvisionsWithOptions(tmpReq *GetServiceProvisionsRequest, runtime *util.RuntimeOptions) (_result *GetServiceProvisionsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetServiceProvisionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Parameters)) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, tea.String("Parameters"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ParametersShrink)) {
		query["Parameters"] = request.ParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TrialType)) {
		query["TrialType"] = request.TrialType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceProvisions"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceProvisionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 计算巢查询服务是否开通
//
// @param request - GetServiceProvisionsRequest
//
// @return GetServiceProvisionsResponse
func (client *Client) GetServiceProvisions(request *GetServiceProvisionsRequest) (_result *GetServiceProvisionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetServiceProvisionsResponse{}
	_body, _err := client.GetServiceProvisionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the constraints on the parameters in an Resource Orchestration Service (ROS) template.
//
// @param request - GetServiceTemplateParameterConstraintsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceTemplateParameterConstraintsResponse
func (client *Client) GetServiceTemplateParameterConstraintsWithOptions(request *GetServiceTemplateParameterConstraintsRequest, runtime *util.RuntimeOptions) (_result *GetServiceTemplateParameterConstraintsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DeployRegionId)) {
		query["DeployRegionId"] = request.DeployRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.EnablePrivateVpcConnection)) {
		query["EnablePrivateVpcConnection"] = request.EnablePrivateVpcConnection
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !tea.BoolValue(util.IsUnset(request.SpecificationName)) {
		query["SpecificationName"] = request.SpecificationName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TrialType)) {
		query["TrialType"] = request.TrialType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceTemplateParameterConstraints"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceTemplateParameterConstraintsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the constraints on the parameters in an Resource Orchestration Service (ROS) template.
//
// @param request - GetServiceTemplateParameterConstraintsRequest
//
// @return GetServiceTemplateParameterConstraintsResponse
func (client *Client) GetServiceTemplateParameterConstraints(request *GetServiceTemplateParameterConstraintsRequest) (_result *GetServiceTemplateParameterConstraintsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetServiceTemplateParameterConstraintsResponse{}
	_body, _err := client.GetServiceTemplateParameterConstraintsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a customer.
//
// @param request - GetUserInformationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserInformationResponse
func (client *Client) GetUserInformationWithOptions(request *GetUserInformationRequest, runtime *util.RuntimeOptions) (_result *GetUserInformationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserInformation"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserInformationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a customer.
//
// @param request - GetUserInformationRequest
//
// @return GetUserInformationResponse
func (client *Client) GetUserInformation(request *GetUserInformationRequest) (_result *GetUserInformationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserInformationResponse{}
	_body, _err := client.GetUserInformationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Permission Policy List
//
// @param request - ListPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPoliciesResponse
func (client *Client) ListPoliciesWithOptions(request *ListPoliciesRequest, runtime *util.RuntimeOptions) (_result *ListPoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPolicies"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Permission Policy List
//
// @param request - ListPoliciesRequest
//
// @return ListPoliciesResponse
func (client *Client) ListPolicies(request *ListPoliciesRequest) (_result *ListPoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPoliciesResponse{}
	_body, _err := client.ListPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询服务类别
//
// @param request - ListServiceCategoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceCategoriesResponse
func (client *Client) ListServiceCategoriesWithOptions(runtime *util.RuntimeOptions) (_result *ListServiceCategoriesResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListServiceCategories"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceCategoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询服务类别
//
// @return ListServiceCategoriesResponse
func (client *Client) ListServiceCategories() (_result *ListServiceCategoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServiceCategoriesResponse{}
	_body, _err := client.ListServiceCategoriesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 展示服务实例账单
//
// @param request - ListServiceInstanceBillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceInstanceBillResponse
func (client *Client) ListServiceInstanceBillWithOptions(request *ListServiceInstanceBillRequest, runtime *util.RuntimeOptions) (_result *ListServiceInstanceBillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillingCycle)) {
		query["BillingCycle"] = request.BillingCycle
	}

	if !tea.BoolValue(util.IsUnset(request.BillingDate)) {
		query["BillingDate"] = request.BillingDate
	}

	if !tea.BoolValue(util.IsUnset(request.Granularity)) {
		query["Granularity"] = request.Granularity
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceInstanceBill"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceInstanceBillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 展示服务实例账单
//
// @param request - ListServiceInstanceBillRequest
//
// @return ListServiceInstanceBillResponse
func (client *Client) ListServiceInstanceBill(request *ListServiceInstanceBillRequest) (_result *ListServiceInstanceBillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServiceInstanceBillResponse{}
	_body, _err := client.ListServiceInstanceBillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the deployment and upgrade logs of a service instance.
//
// @param request - ListServiceInstanceLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceInstanceLogsResponse
func (client *Client) ListServiceInstanceLogsWithOptions(request *ListServiceInstanceLogsRequest, runtime *util.RuntimeOptions) (_result *ListServiceInstanceLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.LogSource)) {
		query["LogSource"] = request.LogSource
	}

	if !tea.BoolValue(util.IsUnset(request.Logstore)) {
		query["Logstore"] = request.Logstore
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		query["SortOrder"] = request.SortOrder
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceInstanceLogs"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceInstanceLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the deployment and upgrade logs of a service instance.
//
// @param request - ListServiceInstanceLogsRequest
//
// @return ListServiceInstanceLogsResponse
func (client *Client) ListServiceInstanceLogs(request *ListServiceInstanceLogsRequest) (_result *ListServiceInstanceLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServiceInstanceLogsResponse{}
	_body, _err := client.ListServiceInstanceLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the resources contained in a service instance.
//
// @param request - ListServiceInstanceResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceInstanceResourcesResponse
func (client *Client) ListServiceInstanceResourcesWithOptions(request *ListServiceInstanceResourcesRequest, runtime *util.RuntimeOptions) (_result *ListServiceInstanceResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filters)) {
		query["Filters"] = request.Filters
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceResourceType)) {
		query["ServiceInstanceResourceType"] = request.ServiceInstanceResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceInstanceResources"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceInstanceResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resources contained in a service instance.
//
// @param request - ListServiceInstanceResourcesRequest
//
// @return ListServiceInstanceResourcesResponse
func (client *Client) ListServiceInstanceResources(request *ListServiceInstanceResourcesRequest) (_result *ListServiceInstanceResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServiceInstanceResourcesResponse{}
	_body, _err := client.ListServiceInstanceResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the upgrade history of a service instance.
//
// @param request - ListServiceInstanceUpgradeHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceInstanceUpgradeHistoryResponse
func (client *Client) ListServiceInstanceUpgradeHistoryWithOptions(request *ListServiceInstanceUpgradeHistoryRequest, runtime *util.RuntimeOptions) (_result *ListServiceInstanceUpgradeHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceInstanceUpgradeHistory"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceInstanceUpgradeHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the upgrade history of a service instance.
//
// @param request - ListServiceInstanceUpgradeHistoryRequest
//
// @return ListServiceInstanceUpgradeHistoryResponse
func (client *Client) ListServiceInstanceUpgradeHistory(request *ListServiceInstanceUpgradeHistoryRequest) (_result *ListServiceInstanceUpgradeHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServiceInstanceUpgradeHistoryResponse{}
	_body, _err := client.ListServiceInstanceUpgradeHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// {}
//
// @param request - ListServiceInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceInstancesResponse
func (client *Client) ListServiceInstancesWithOptions(request *ListServiceInstancesRequest, runtime *util.RuntimeOptions) (_result *ListServiceInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceInstances"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// {}
//
// @param request - ListServiceInstancesRequest
//
// @return ListServiceInstancesResponse
func (client *Client) ListServiceInstances(request *ListServiceInstancesRequest) (_result *ListServiceInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServiceInstancesResponse{}
	_body, _err := client.ListServiceInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the applications for using a service.
//
// @param request - ListServiceUsagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceUsagesResponse
func (client *Client) ListServiceUsagesWithOptions(request *ListServiceUsagesRequest, runtime *util.RuntimeOptions) (_result *ListServiceUsagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceUsages"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceUsagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the applications for using a service.
//
// @param request - ListServiceUsagesRequest
//
// @return ListServiceUsagesResponse
func (client *Client) ListServiceUsages(request *ListServiceUsagesRequest) (_result *ListServiceUsagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServiceUsagesResponse{}
	_body, _err := client.ListServiceUsagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of services.
//
// @param request - ListServicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServicesResponse
func (client *Client) ListServicesWithOptions(request *ListServicesRequest, runtime *util.RuntimeOptions) (_result *ListServicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.FuzzyKeyword)) {
		query["FuzzyKeyword"] = request.FuzzyKeyword
	}

	if !tea.BoolValue(util.IsUnset(request.InUsed)) {
		query["InUsed"] = request.InUsed
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrderByType)) {
		query["OrderByType"] = request.OrderByType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceAccessType)) {
		query["ServiceAccessType"] = request.ServiceAccessType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServices"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of services.
//
// @param request - ListServicesRequest
//
// @return ListServicesResponse
func (client *Client) ListServices(request *ListServicesRequest) (_result *ListServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServicesResponse{}
	_body, _err := client.ListServicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询标签键列表
//
// @param request - ListTagKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagKeysResponse
func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *util.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagKeys"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询标签键列表
//
// @param request - ListTagKeysRequest
//
// @return ListTagKeysResponse
func (client *Client) ListTagKeys(request *ListTagKeysRequest) (_result *ListTagKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagKeysResponse{}
	_body, _err := client.ListTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询标签资源列表
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询标签资源列表
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询标签值列表
//
// @param request - ListTagValuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagValuesResponse
func (client *Client) ListTagValuesWithOptions(request *ListTagValuesRequest, runtime *util.RuntimeOptions) (_result *ListTagValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagValues"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagValuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询标签值列表
//
// @param request - ListTagValuesRequest
//
// @return ListTagValuesResponse
func (client *Client) ListTagValues(request *ListTagValuesRequest) (_result *ListTagValuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagValuesResponse{}
	_body, _err := client.ListTagValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Renew the prepaid resources included in the private deployment service instance. You can renew all prepaid resources under the specified service instance ID, or you can renew the specified resources. Only one of the two renewal methods can be used.
//
// @param request - RenewServiceInstanceResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewServiceInstanceResourcesResponse
func (client *Client) RenewServiceInstanceResourcesWithOptions(request *RenewServiceInstanceResourcesRequest, runtime *util.RuntimeOptions) (_result *RenewServiceInstanceResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourcePeriod)) {
		query["ResourcePeriod"] = request.ResourcePeriod
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RenewServiceInstanceResources"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RenewServiceInstanceResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renew the prepaid resources included in the private deployment service instance. You can renew all prepaid resources under the specified service instance ID, or you can renew the specified resources. Only one of the two renewal methods can be used.
//
// @param request - RenewServiceInstanceResourcesRequest
//
// @return RenewServiceInstanceResourcesResponse
func (client *Client) RenewServiceInstanceResources(request *RenewServiceInstanceResourcesRequest) (_result *RenewServiceInstanceResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenewServiceInstanceResourcesResponse{}
	_body, _err := client.RenewServiceInstanceResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// When the service instance is Deployed, call the RestartServiceInstance interface to restart the service instance.
//
// @param request - RestartServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartServiceInstanceResponse
func (client *Client) RestartServiceInstanceWithOptions(request *RestartServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *RestartServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RestartServiceInstance"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RestartServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// When the service instance is Deployed, call the RestartServiceInstance interface to restart the service instance.
//
// @param request - RestartServiceInstanceRequest
//
// @return RestartServiceInstanceResponse
func (client *Client) RestartServiceInstance(request *RestartServiceInstanceRequest) (_result *RestartServiceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestartServiceInstanceResponse{}
	_body, _err := client.RestartServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Rolls back an upgraded service instance to the previous version.
//
// @param request - RollbackServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RollbackServiceInstanceResponse
func (client *Client) RollbackServiceInstanceWithOptions(request *RollbackServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *RollbackServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RollbackServiceInstance"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RollbackServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rolls back an upgraded service instance to the previous version.
//
// @param request - RollbackServiceInstanceRequest
//
// @return RollbackServiceInstanceResponse
func (client *Client) RollbackServiceInstance(request *RollbackServiceInstanceRequest) (_result *RollbackServiceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RollbackServiceInstanceResponse{}
	_body, _err := client.RollbackServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// When the service instance status is Stopped (Stopped) or StartFailed (Startup failed), the StartServiceInstance interface is invoked to start the service instance.
//
// @param request - StartServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartServiceInstanceResponse
func (client *Client) StartServiceInstanceWithOptions(request *StartServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *StartServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartServiceInstance"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// When the service instance status is Stopped (Stopped) or StartFailed (Startup failed), the StartServiceInstance interface is invoked to start the service instance.
//
// @param request - StartServiceInstanceRequest
//
// @return StartServiceInstanceResponse
func (client *Client) StartServiceInstance(request *StartServiceInstanceRequest) (_result *StartServiceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartServiceInstanceResponse{}
	_body, _err := client.StartServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// When the service instance is Deployed and StopFailed, call the StopServiceInstance interface to stop the service instance.
//
// @param request - StopServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopServiceInstanceResponse
func (client *Client) StopServiceInstanceWithOptions(request *StopServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *StopServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopServiceInstance"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// When the service instance is Deployed and StopFailed, call the StopServiceInstance interface to stop the service instance.
//
// @param request - StopServiceInstanceRequest
//
// @return StopServiceInstanceResponse
func (client *Client) StopServiceInstance(request *StopServiceInstanceRequest) (_result *StopServiceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopServiceInstanceResponse{}
	_body, _err := client.StopServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 给资源打标签
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 给资源打标签
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 给资源解除标签
//
// @param request - UnTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnTagResourcesResponse
func (client *Client) UnTagResourcesWithOptions(request *UnTagResourcesRequest, runtime *util.RuntimeOptions) (_result *UnTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnTagResources"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 给资源解除标签
//
// @param request - UnTagResourcesRequest
//
// @return UnTagResourcesResponse
func (client *Client) UnTagResources(request *UnTagResourcesRequest) (_result *UnTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnTagResourcesResponse{}
	_body, _err := client.UnTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the attributes of a service instance.
//
// @param request - UpdateServiceInstanceAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceInstanceAttributesResponse
func (client *Client) UpdateServiceInstanceAttributesWithOptions(request *UpdateServiceInstanceAttributesRequest, runtime *util.RuntimeOptions) (_result *UpdateServiceInstanceAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnableOperation)) {
		query["EnableOperation"] = request.EnableOperation
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServiceInstanceAttributes"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceInstanceAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the attributes of a service instance.
//
// @param request - UpdateServiceInstanceAttributesRequest
//
// @return UpdateServiceInstanceAttributesResponse
func (client *Client) UpdateServiceInstanceAttributes(request *UpdateServiceInstanceAttributesRequest) (_result *UpdateServiceInstanceAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateServiceInstanceAttributesResponse{}
	_body, _err := client.UpdateServiceInstanceAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the configurations of a service instance.
//
// Description:
//
// ### [](#)Prerequisites
//
// Configuration change is enabled and the related parameters are configured for the service by the service provider.
//
// @param tmpReq - UpdateServiceInstanceSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceInstanceSpecResponse
func (client *Client) UpdateServiceInstanceSpecWithOptions(tmpReq *UpdateServiceInstanceSpecRequest, runtime *util.RuntimeOptions) (_result *UpdateServiceInstanceSpecResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateServiceInstanceSpecShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Parameters)) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, tea.String("Parameters"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Commodity)) {
		query["Commodity"] = request.Commodity
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EnableUserPrometheus)) {
		query["EnableUserPrometheus"] = request.EnableUserPrometheus
	}

	if !tea.BoolValue(util.IsUnset(request.OperationName)) {
		query["OperationName"] = request.OperationName
	}

	if !tea.BoolValue(util.IsUnset(request.ParametersShrink)) {
		query["Parameters"] = request.ParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PredefinedParametersName)) {
		query["PredefinedParametersName"] = request.PredefinedParametersName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServiceInstanceSpec"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceInstanceSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the configurations of a service instance.
//
// Description:
//
// ### [](#)Prerequisites
//
// Configuration change is enabled and the related parameters are configured for the service by the service provider.
//
// @param request - UpdateServiceInstanceSpecRequest
//
// @return UpdateServiceInstanceSpecResponse
func (client *Client) UpdateServiceInstanceSpec(request *UpdateServiceInstanceSpecRequest) (_result *UpdateServiceInstanceSpecResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateServiceInstanceSpecResponse{}
	_body, _err := client.UpdateServiceInstanceSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the application for using a service.
//
// @param tmpReq - UpdateServiceUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceUsageResponse
func (client *Client) UpdateServiceUsageWithOptions(tmpReq *UpdateServiceUsageRequest, runtime *util.RuntimeOptions) (_result *UpdateServiceUsageResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateServiceUsageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.UserInformation)) {
		request.UserInformationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserInformation, tea.String("UserInformation"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserInformationShrink)) {
		query["UserInformation"] = request.UserInformationShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServiceUsage"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the application for using a service.
//
// @param request - UpdateServiceUsageRequest
//
// @return UpdateServiceUsageResponse
func (client *Client) UpdateServiceUsage(request *UpdateServiceUsageRequest) (_result *UpdateServiceUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateServiceUsageResponse{}
	_body, _err := client.UpdateServiceUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about a customer.
//
// @param request - UpdateUserInformationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserInformationResponse
func (client *Client) UpdateUserInformationWithOptions(request *UpdateUserInformationRequest, runtime *util.RuntimeOptions) (_result *UpdateUserInformationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeliverySettings)) {
		query["DeliverySettings"] = request.DeliverySettings
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUserInformation"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserInformationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a customer.
//
// @param request - UpdateUserInformationRequest
//
// @return UpdateUserInformationResponse
func (client *Client) UpdateUserInformation(request *UpdateUserInformationRequest) (_result *UpdateUserInformationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserInformationResponse{}
	_body, _err := client.UpdateUserInformationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Upgrades the version of a service instance.
//
// @param tmpReq - UpgradeServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeServiceInstanceResponse
func (client *Client) UpgradeServiceInstanceWithOptions(tmpReq *UpgradeServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *UpgradeServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpgradeServiceInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Parameters)) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, tea.String("Parameters"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.ParametersShrink)) {
		query["Parameters"] = request.ParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeServiceInstance"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades the version of a service instance.
//
// @param request - UpgradeServiceInstanceRequest
//
// @return UpgradeServiceInstanceResponse
func (client *Client) UpgradeServiceInstance(request *UpgradeServiceInstanceRequest) (_result *UpgradeServiceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeServiceInstanceResponse{}
	_body, _err := client.UpgradeServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
