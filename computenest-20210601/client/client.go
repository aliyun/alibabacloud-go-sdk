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
	Order       *CommodityValueResultOrder     `json:"Order,omitempty" xml:"Order,omitempty" type:"Struct"`
	InquiryType *string                        `json:"InquiryType,omitempty" xml:"InquiryType,omitempty"`
	SubOrders   *CommodityValueResultSubOrders `json:"SubOrders,omitempty" xml:"SubOrders,omitempty" type:"Struct"`
	Coupons     []*CommodityValueResultCoupons `json:"Coupons,omitempty" xml:"Coupons,omitempty" type:"Repeated"`
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
	Currency       *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	TradeAmount    *string `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
	DiscountAmount *string `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
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
	ModuleId        *int64                                                            `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ModuleName      *string                                                           `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	ModuleCode      *string                                                           `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	TotalProductFee *float64                                                          `json:"TotalProductFee,omitempty" xml:"TotalProductFee,omitempty"`
	DiscountFee     *float64                                                          `json:"DiscountFee,omitempty" xml:"DiscountFee,omitempty"`
	PayFee          *float64                                                          `json:"PayFee,omitempty" xml:"PayFee,omitempty"`
	PriceUnit       *string                                                           `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
	IsPricingModule *bool                                                             `json:"IsPricingModule,omitempty" xml:"IsPricingModule,omitempty"`
	NeedOrderPay    *bool                                                             `json:"NeedOrderPay,omitempty" xml:"NeedOrderPay,omitempty"`
	PriceType       *string                                                           `json:"PriceType,omitempty" xml:"PriceType,omitempty"`
	ModuleAttrs     []*CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs `json:"ModuleAttrs,omitempty" xml:"ModuleAttrs,omitempty" type:"Repeated"`
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
	Type  *int64  `json:"Type,omitempty" xml:"Type,omitempty"`
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Code  *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Unit  *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
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
	CanPromFee     *float64 `json:"CanPromFee,omitempty" xml:"CanPromFee,omitempty"`
	CouponDesc     *string  `json:"CouponDesc,omitempty" xml:"CouponDesc,omitempty"`
	CouponName     *string  `json:"CouponName,omitempty" xml:"CouponName,omitempty"`
	CouponOptionNo *string  `json:"CouponOptionNo,omitempty" xml:"CouponOptionNo,omitempty"`
	Selected       *bool    `json:"Selected,omitempty" xml:"Selected,omitempty"`
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
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// true
	NeedDelete *bool `json:"NeedDelete,omitempty" xml:"NeedDelete,omitempty"`
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
	// 云账号报警联系人
	ContactGroup *string `json:"ContactGroup,omitempty" xml:"ContactGroup,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks the required parameters, request syntax, and limits.
	//
	// 	- **false*	- (default): sends the request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Indicates whether the service instance supports the operation feature.
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
	// Whether the resource pays automatically.Valid values:
	//
	// - true
	//
	// - false
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
	// 套餐一
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// The tags.
	Tag []*CreateServiceInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The name of the template.
	//
	// example:
	//
	// 模板1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The trial type of serviceInstance.
	//
	// Valid values:
	//
	// - Created:
	//
	// - Deploying
	//
	// - DeployedFailed
	//
	// - Deployed
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
	// Specifies whether to enable auto-renewal for the instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	AutoRenew *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	CouponId  *string `json:"CouponId,omitempty" xml:"CouponId,omitempty"`
	// The subscription duration.
	//
	// example:
	//
	// 1
	PayPeriod *int64 `json:"PayPeriod,omitempty" xml:"PayPeriod,omitempty"`
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
	// 云账号报警联系人
	ContactGroup *string `json:"ContactGroup,omitempty" xml:"ContactGroup,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks the required parameters, request syntax, and limits.
	//
	// 	- **false*	- (default): sends the request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Indicates whether the service instance supports the operation feature.
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
	// Whether the resource pays automatically.Valid values:
	//
	// - true
	//
	// - false
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
	// 套餐一
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// The tags.
	Tag []*CreateServiceInstanceShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The name of the template.
	//
	// example:
	//
	// 模板1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The trial type of serviceInstance.
	//
	// Valid values:
	//
	// - Created:
	//
	// - Deploying
	//
	// - DeployedFailed
	//
	// - Deployed
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
	// Specifies whether to enable auto-renewal for the instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	AutoRenew *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	CouponId  *string `json:"CouponId,omitempty" xml:"CouponId,omitempty"`
	// The subscription duration.
	//
	// example:
	//
	// 1
	PayPeriod *int64 `json:"PayPeriod,omitempty" xml:"PayPeriod,omitempty"`
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
	// Indicates whether the synchronization task was created. Valid values:
	//
	// 	- **1**: Created.
	//
	// 	- **0**: Creation failed. The tables in the synchronization task are duplicate. The duplicate tables are returned for the **RepeatedDbs*	- parameter.
	//
	// 	- **-1**: Creation failed. The cause why the creation failed is returned for the **ErrorMsg*	- parameter.
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
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// service-d6fc5f949a9246xxxxxx
	ServiceId       *string            `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
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
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// service-d6fc5f949a9246xxxxxx
	ServiceId             *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
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

type GetServiceEstimateCostRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// qwertyuiop
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The commodity details.
	Commodity *GetServiceEstimateCostRequestCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// The name of the configuration update operation.
	//
	// example:
	//
	// 修改游戏参数
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	// The parameters that are specified for service instance deployment.
	//
	// >  If you want to specify the region in which the service instance is deployed, you must specify the information in Parameters.
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
	// The package name.
	//
	// example:
	//
	// 自定义套餐
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// The name of the template. This parameter is returned only if you specify TemplateId.
	//
	// > -   If you specify TemplateVersion, the name of the template whose version is specified by TemplateVersion is returned.
	//
	// > -  If you not specify TemplateVersion, the name of the template whose version is the default version is returned.
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
	// The commodity details.
	CommodityShrink *string `json:"Commodity,omitempty" xml:"Commodity,omitempty"`
	// The name of the configuration update operation.
	//
	// example:
	//
	// 修改游戏参数
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	// The parameters that are specified for service instance deployment.
	//
	// >  If you want to specify the region in which the service instance is deployed, you must specify the information in Parameters.
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
	// The package name.
	//
	// example:
	//
	// 自定义套餐
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// The name of the template. This parameter is returned only if you specify TemplateId.
	//
	// > -   If you specify TemplateVersion, the name of the template whose version is specified by TemplateVersion is returned.
	//
	// > -  If you not specify TemplateVersion, the name of the template whose version is the default version is returned.
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
	// Estimated commodity cost.
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
	// Estimated resource cost.
	//
	// example:
	//
	// {
	//
	//     "EcsInstance" : {
	//
	//       "Type" : "ALIYUN::ECS::Instance",
	//
	//       "Success" : true,
	//
	//       "Result" : {
	//
	//         "Order" : {
	//
	//           "Currency" : "CNY",
	//
	//           "RuleIds" : [ "102111101338****" ],
	//
	//           "DetailInfos" : {
	//
	//             "ResourcePriceModel" : [ {
	//
	//               "OriginalPrice" : 0,
	//
	//               "DiscountPrice" : 0,
	//
	//               "SubRules" : {
	//
	//                 "Rule" : [ ]
	//
	//               },
	//
	//               "Resource" : "bandwidth",
	//
	//               "TradePrice" : 0
	//
	//             }, {
	//
	//               "OriginalPrice" : 0,
	//
	//               "DiscountPrice" : 0,
	//
	//               "SubRules" : {
	//
	//                 "Rule" : [ ]
	//
	//               },
	//
	//               "Resource" : "image",
	//
	//               "TradePrice" : 0
	//
	//             }, {
	//
	//               "OriginalPrice" : 0.366666,
	//
	//               "DiscountPrice" : 0.249012,
	//
	//               "SubRules" : {
	//
	//                 "Rule" : [ ]
	//
	//               },
	//
	//               "Resource" : "instanceType",
	//
	//               "TradePrice" : 0.117654
	//
	//             }, {
	//
	//               "OriginalPrice" : 0.055555,
	//
	//               "DiscountPrice" : 0.037729,
	//
	//               "SubRules" : {
	//
	//                 "Rule" : [ ]
	//
	//               },
	//
	//               "Resource" : "systemDisk",
	//
	//               "TradePrice" : 0.017826
	//
	//             } ]
	//
	//           },
	//
	//           "TradeAmount" : 0.135,
	//
	//           "OriginalAmount" : 0.422,
	//
	//           "Coupons" : {
	//
	//             "Coupon" : [ ]
	//
	//           },
	//
	//           "DiscountAmount" : 0.287
	//
	//         },
	//
	//         "OrderSupplement" : {
	//
	//           "PriceUnit" : "/Hour",
	//
	//           "ChargeType" : "PostPaid",
	//
	//           "Quantity" : 1,
	//
	//           "PriceType" : "Total"
	//
	//         },
	//
	//         "Rules" : {
	//
	//           "Rule" : [ {
	//
	//             "RuleDescId" : "102111101338****",
	//
	//             "Name" : "合同优惠_多计费项优惠_3.208750折"
	//
	//           } ]
	//
	//         }
	//
	//       }
	//
	//     }
	//
	//   }
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
	// 套餐一
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
	// 模板1
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
	// A公司
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// The URL of the service provider.
	//
	// example:
	//
	// http://example.com
	SupplierUrl *string `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	// The service version that can be updated.
	UpgradableServiceVersions []*string `json:"UpgradableServiceVersions,omitempty" xml:"UpgradableServiceVersions,omitempty" type:"Repeated"`
	// The metadata about the upgrade.
	//
	// example:
	//
	// {\\"Description\\":\\"开启服务升级\\",\\"SupportRollback\\":true,\\"SupportUpgradeFromVersions\\":[],\\"UpgradeComponents\\":[\\"Configuration\\",\\"Resource\\"]}
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
	// B数据库
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The description of the service.
	//
	// example:
	//
	// B是A公司自主设计并研发的开源分布式的关系型数据库
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

type ListServiceInstanceLogsRequest struct {
	Filter []*ListServiceInstanceLogsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The log source. Valid values:
	//
	// 	- computeNest (default): logs of the deployment and upgrade of the service instance.
	//
	// 	- application: logs generated by the application.
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

type ListServiceInstanceLogsRequestFilter struct {
	Name  *string   `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// The source of the service instance log. Valid values:
	//
	// 	- ros: The log is generated by Resource Orchestration Service (ROS).
	//
	// 	- computeNest: The log is generated by Compute Nest.
	//
	// example:
	//
	// ros
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
	// End time of resource usage.
	//
	// <notice	Note: Only supports querying service instances on private deployments.
	//
	// example:
	//
	// 2022-03-01T12:00:00
	ExpireTimeEnd *string `json:"ExpireTimeEnd,omitempty" xml:"ExpireTimeEnd,omitempty"`
	// Start time of resource usage.
	//
	// <notice	Note: Only supports querying service instances on private deployments.
	//
	// example:
	//
	// 2022-01-01T12:00:00
	ExpireTimeStart *string `json:"ExpireTimeStart,omitempty" xml:"ExpireTimeStart,omitempty"`
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
	// The billing method of the read-only instance. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go
	//
	// 	- **Prepaid**: subscription
	//
	// example:
	//
	// Subscription
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
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
	// The Alibaba Cloud Resource Name (ARN) of a resource.
	ResourceARN []*string `json:"ResourceARN,omitempty" xml:"ResourceARN,omitempty" type:"Repeated"`
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

func (s *ListServiceInstanceResourcesRequest) SetExpireTimeEnd(v string) *ListServiceInstanceResourcesRequest {
	s.ExpireTimeEnd = &v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetExpireTimeStart(v string) *ListServiceInstanceResourcesRequest {
	s.ExpireTimeStart = &v
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

func (s *ListServiceInstanceResourcesRequest) SetPayType(v string) *ListServiceInstanceResourcesRequest {
	s.PayType = &v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetRegionId(v string) *ListServiceInstanceResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetResourceARN(v []*string) *ListServiceInstanceResourcesRequest {
	s.ResourceARN = v
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
	Filter []*ListServiceUsagesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
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
	// example:
	//
	// ServiceId
	Name  *string   `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAYChudnQUoBH+mGWFpb6oP0=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 18AD0960-A9FE-1AC8-ADF8-22131Fxxxxxx
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceUsages []*ListServiceUsagesResponseBodyServiceUsages `json:"ServiceUsages,omitempty" xml:"ServiceUsages,omitempty" type:"Repeated"`
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

type ListServiceUsagesResponseBodyServiceUsages struct {
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// example:
	//
	// 2022-05-25T02:02:02Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// service-c9f36ec6d19b4exxxxxx
	ServiceId   *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// example:
	//
	// Submitted
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// example:
	//
	// 2022-05-25T02:02:02Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 127383705958xxxx
	UserAliUid      *int64             `json:"UserAliUid,omitempty" xml:"UserAliUid,omitempty"`
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

type UpdateServiceInstanceSpecRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The information about the order placed in Alibaba Cloud Marketplace. You do not need to specify this parameter if the service is not published in Alibaba Cloud Marketplace or uses the pay-as-you-go billing method.
	Commodity *UpdateServiceInstanceSpecRequestCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// Specifies whether to perform only a dry run for the request to check information such as the permissions and instance status. Valid values:
	//
	// 	- **true: performs a dry run for the request, but does not create a service instance.**
	//
	// 	- **false: performs a dry run for the request, and creates a service instance if the request passes the dry run.**
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to enable Prometheus on the customer side. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnableUserPrometheus *bool `json:"EnableUserPrometheus,omitempty" xml:"EnableUserPrometheus,omitempty"`
	// The name of the configuration update operation.
	//
	// example:
	//
	// package modify
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	// The configuration parameters of the service instance.
	//
	// example:
	//
	// {
	//
	//   "InstanceType": "ecs.g8ise.2xlarge"
	//
	// }
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The name of the specification package.
	//
	// example:
	//
	// package One
	PredefinedParametersName *string `json:"PredefinedParametersName,omitempty" xml:"PredefinedParametersName,omitempty"`
	// The service instance ID.
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
	// Specifies whether to automatically complete the payment. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
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
	// Specifies whether to perform only a dry run for the request to check information such as the permissions and instance status. Valid values:
	//
	// 	- **true: performs a dry run for the request, but does not create a service instance.**
	//
	// 	- **false: performs a dry run for the request, and creates a service instance if the request passes the dry run.**
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to enable Prometheus on the customer side. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnableUserPrometheus *bool `json:"EnableUserPrometheus,omitempty" xml:"EnableUserPrometheus,omitempty"`
	// The name of the configuration update operation.
	//
	// example:
	//
	// package modify
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	// The configuration parameters of the service instance.
	//
	// example:
	//
	// {
	//
	//   "InstanceType": "ecs.g8ise.2xlarge"
	//
	// }
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The name of the specification package.
	//
	// example:
	//
	// package One
	PredefinedParametersName *string `json:"PredefinedParametersName,omitempty" xml:"PredefinedParametersName,omitempty"`
	// The service instance ID.
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
	// Specifies whether to automatically complete the payment. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
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
	// example:
	//
	// AAAAAYChudnQUoBH+mGWFpb6oP0=
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// service-39f4f251e94843xxxxxx
	ServiceId       *string            `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
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
	// example:
	//
	// AAAAAYChudnQUoBH+mGWFpb6oP0=
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// service-39f4f251e94843xxxxxx
	ServiceId             *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
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
// 用户取消服务使用请求
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
// 用户取消服务使用请求
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
// 用户创建服务使用请求
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
// 用户创建服务使用请求
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
// 计算巢服务部署询价
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
// 计算巢服务部署询价
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
// Queries the information about a service instance.
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
// Queries the information about a service instance.
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
	if !tea.BoolValue(util.IsUnset(request.ExpireTimeEnd)) {
		query["ExpireTimeEnd"] = request.ExpireTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.ExpireTimeStart)) {
		query["ExpireTimeStart"] = request.ExpireTimeStart
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceARN)) {
		query["ResourceARN"] = request.ResourceARN
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
// 用户查询服务使用申请接口
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
// 用户查询服务使用申请接口
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
// 用户更新服务使用请求
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
// 用户更新服务使用请求
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
