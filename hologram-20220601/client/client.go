// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ChangeResourceGroupRequest struct {
	// The instance ID.
	//
	// example:
	//
	// hgprecn-cn-zvp25ysv3006
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// new resource group id
	//
	// example:
	//
	// rg-acfmxwerqwerasfd
	NewResourceGroupId *string `json:"newResourceGroupId,omitempty" xml:"newResourceGroupId,omitempty"`
}

func (s ChangeResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupRequest) SetInstanceId(v string) *ChangeResourceGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetNewResourceGroupId(v string) *ChangeResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

type ChangeResourceGroupResponseBody struct {
	// The returned data.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// AB71198A-2DB1-511B-AE4D-690BAA97F076
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponseBody) SetData(v bool) *ChangeResourceGroupResponseBody {
	s.Data = &v
	return s
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

type CreateHoloWarehouseRequest struct {
	// The specifications of the virtual warehouse. The number of vCPUs must be an integer multiple of 16 CPUs. Minimum value: 16.
	//
	// This parameter is required.
	//
	// example:
	//
	// 32
	Cpu *string `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// The name of the virtual warehouse.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_warehouse
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateHoloWarehouseRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHoloWarehouseRequest) GoString() string {
	return s.String()
}

func (s *CreateHoloWarehouseRequest) SetCpu(v string) *CreateHoloWarehouseRequest {
	s.Cpu = &v
	return s
}

func (s *CreateHoloWarehouseRequest) SetName(v string) *CreateHoloWarehouseRequest {
	s.Name = &v
	return s
}

type CreateHoloWarehouseResponseBody struct {
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2A8DEF6E-067E-5DB0-BAE1-2894266E6C6A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHoloWarehouseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateHoloWarehouseResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHoloWarehouseResponseBody) SetData(v bool) *CreateHoloWarehouseResponseBody {
	s.Data = &v
	return s
}

func (s *CreateHoloWarehouseResponseBody) SetRequestId(v string) *CreateHoloWarehouseResponseBody {
	s.RequestId = &v
	return s
}

type CreateHoloWarehouseResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHoloWarehouseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHoloWarehouseResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateHoloWarehouseResponse) GoString() string {
	return s.String()
}

func (s *CreateHoloWarehouseResponse) SetHeaders(v map[string]*string) *CreateHoloWarehouseResponse {
	s.Headers = v
	return s
}

func (s *CreateHoloWarehouseResponse) SetStatusCode(v int32) *CreateHoloWarehouseResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHoloWarehouseResponse) SetBody(v *CreateHoloWarehouseResponseBody) *CreateHoloWarehouseResponse {
	s.Body = v
	return s
}

type CreateInstanceRequest struct {
	// Specifies whether to enable auto-payment. Default value: true. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// >  The default value is true. If the balance of your account is insufficient, you can set this parameter to false. In this case, an unpaid order is generated. You can log on to the Expenses and Costs console to pay for the order.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"autoPay,omitempty" xml:"autoPay,omitempty"`
	// Specifies whether to enable monthly auto-renewal. The default value is false. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- PrePaid: subscription
	//
	// 	- PostPaid: pay-as-you-go
	//
	// >  This parameter is invalid for Hologres Shared Cluster instances. Hologres Shared Cluster instances have fixed specifications and are pay-as-you-go instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The infrequent access (IA) storage space of the instance. Unit: GB.
	//
	// >  This parameter is invalid for pay-as-you-go instances.
	//
	// example:
	//
	// 500
	ColdStorageSize *int64 `json:"coldStorageSize,omitempty" xml:"coldStorageSize,omitempty"`
	// The instance specifications. Valid values:
	//
	// 	- 8-core 32GB (number of compute nodes: 1)
	//
	// 	- 32-core 128GB (number of compute nodes: 2)
	//
	// 	- 64-core 256GB (number of compute nodes: 4)
	//
	// 	- 96-core 384GB (number of compute nodes: 6)
	//
	// 	- 128-core 512GB (number of compute nodes: 8)
	//
	// 	- Others
	//
	// >
	//
	// 	- Set this parameter to the number of cores.
	//
	// 	- If you want to set this parameter to specifications with more than 1,024 GB, you must submit a ticket.
	//
	// 	- This parameter is invalid for Hologres Shared Cluster instances.
	//
	// 	- The specifications of 8-core 32GB (number of compute nodes: 1) are for trial use only and cannot be used for production.
	//
	// example:
	//
	// 64
	Cpu *int64 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// The validity period of the instance that you want to purchase. For example, you can specify a validity period of two months.
	//
	// >  You do not need to configure this parameter for pay-as-you-go instances.
	//
	// example:
	//
	// 2
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// Specifies whether to enable the Serverless Computing feature.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableServerlessComputing *bool `json:"enableServerlessComputing,omitempty" xml:"enableServerlessComputing,omitempty"`
	// The number of gateways. Valid values: 2 to 50.
	//
	// >  This parameter is required only for virtual warehouse instances.
	//
	// example:
	//
	// 4
	GatewayCount *int64 `json:"gatewayCount,omitempty" xml:"gatewayCount,omitempty"`
	// The initial database.
	//
	// example:
	//
	// chatbot
	InitialDatabases *string `json:"initialDatabases,omitempty" xml:"initialDatabases,omitempty"`
	// The name of the instance. The name must be 2 to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_holo
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// The category of the instance. Valid values:
	//
	// 	- Standard: general-purpose instance
	//
	// 	- Follower: read-only secondary instance
	//
	// 	- Warehouse: virtual warehouse instance
	//
	// 	- Shared: Hologres Shared Cluster instance
	//
	// This parameter is required.
	//
	// example:
	//
	// Standard
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// The ID of the primary instance. This parameter is required for read-only secondary instances.
	//
	// >  The primary and secondary instances must meet the following requirements:
	//
	// 	- The primary instance is in the Running state.
	//
	// 	- The primary instance and secondary instances are deployed in the same region.
	//
	// 	- The primary instance and secondary instances are deployed in the same zone.
	//
	// 	- Less than 10 secondary instances are associated with the primary instance.
	//
	// 	- The primary instance and secondary instances belong to the same Alibaba Cloud account.
	//
	// example:
	//
	// hgpostcn-cn-lbj3aworq112
	LeaderInstanceId *string `json:"leaderInstanceId,omitempty" xml:"leaderInstanceId,omitempty"`
	// The billing cycle. Valid values:
	//
	// 	- Month
	//
	// 	- Hour
	//
	// >
	//
	// 	- This parameter can only be set to Month for subscription instances.
	//
	// 	- This parameter can only be set to Hour for pay-as-you-go instances.
	//
	// 	- By default, this parameter is set to Hour for Hologres Shared Cluster instances.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"pricingCycle,omitempty" xml:"pricingCycle,omitempty"`
	// The ID of the region. You can obtain region IDs in [Endpoints](https://www.alibabacloud.com/help/en/maxcompute/user-guide/endpoints).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The ID of the resource group. If you do not specify this parameter, the default resource group of the account is used.
	//
	// example:
	//
	// ""
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The standard storage space of the instance. Unit: GB.
	//
	// >  This parameter is invalid for pay-as-you-go instances.
	//
	// example:
	//
	// 500
	StorageSize *int64  `json:"storageSize,omitempty" xml:"storageSize,omitempty"`
	StorageType *string `json:"storageType,omitempty" xml:"storageType,omitempty"`
	// The ID of the vSwitch. The zone in which the vSwitch resides must be the same as the zone in which the Hologres instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-2vccsiymtxxxxxx
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC). The region in which the VPC resides must be the same as the region in which the Hologres instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-t4netc3y5xxxx
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// The ID of the zone. For more information, see the "Operation description" section in this topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetAutoPay(v bool) *CreateInstanceRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateInstanceRequest) SetAutoRenew(v bool) *CreateInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateInstanceRequest) SetChargeType(v string) *CreateInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateInstanceRequest) SetColdStorageSize(v int64) *CreateInstanceRequest {
	s.ColdStorageSize = &v
	return s
}

func (s *CreateInstanceRequest) SetCpu(v int64) *CreateInstanceRequest {
	s.Cpu = &v
	return s
}

func (s *CreateInstanceRequest) SetDuration(v int64) *CreateInstanceRequest {
	s.Duration = &v
	return s
}

func (s *CreateInstanceRequest) SetEnableServerlessComputing(v bool) *CreateInstanceRequest {
	s.EnableServerlessComputing = &v
	return s
}

func (s *CreateInstanceRequest) SetGatewayCount(v int64) *CreateInstanceRequest {
	s.GatewayCount = &v
	return s
}

func (s *CreateInstanceRequest) SetInitialDatabases(v string) *CreateInstanceRequest {
	s.InitialDatabases = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceName(v string) *CreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceType(v string) *CreateInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateInstanceRequest) SetLeaderInstanceId(v string) *CreateInstanceRequest {
	s.LeaderInstanceId = &v
	return s
}

func (s *CreateInstanceRequest) SetPricingCycle(v string) *CreateInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateInstanceRequest) SetRegionId(v string) *CreateInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateInstanceRequest) SetResourceGroupId(v string) *CreateInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateInstanceRequest) SetStorageSize(v int64) *CreateInstanceRequest {
	s.StorageSize = &v
	return s
}

func (s *CreateInstanceRequest) SetStorageType(v string) *CreateInstanceRequest {
	s.StorageType = &v
	return s
}

func (s *CreateInstanceRequest) SetVSwitchId(v string) *CreateInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateInstanceRequest) SetVpcId(v string) *CreateInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateInstanceRequest) SetZoneId(v string) *CreateInstanceRequest {
	s.ZoneId = &v
	return s
}

type CreateInstanceResponseBody struct {
	// The returned data.
	Data *CreateInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned.
	//
	// example:
	//
	// null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// null
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9CC37B9F-F4B4-5FF1-939B-AEE78DC70130
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) SetData(v *CreateInstanceResponseBodyData) *CreateInstanceResponseBody {
	s.Data = v
	return s
}

func (s *CreateInstanceResponseBody) SetErrorCode(v string) *CreateInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateInstanceResponseBody) SetErrorMessage(v string) *CreateInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateInstanceResponseBody) SetHttpStatusCode(v string) *CreateInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

type CreateInstanceResponseBodyData struct {
	// The error code returned.
	//
	// example:
	//
	// InvalidVpcOrVSwitch.NotAvailable
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// hgpostcn-cn-xxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The error details.
	//
	// example:
	//
	// Vpc is not available
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 217523224780172
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Indicates whether the instance was created.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBodyData) SetCode(v string) *CreateInstanceResponseBodyData {
	s.Code = &v
	return s
}

func (s *CreateInstanceResponseBodyData) SetInstanceId(v string) *CreateInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceResponseBodyData) SetMessage(v string) *CreateInstanceResponseBodyData {
	s.Message = &v
	return s
}

func (s *CreateInstanceResponseBodyData) SetOrderId(v string) *CreateInstanceResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *CreateInstanceResponseBodyData) SetSuccess(v string) *CreateInstanceResponseBodyData {
	s.Success = &v
	return s
}

type CreateInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponse) SetHeaders(v map[string]*string) *CreateInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceResponse) SetStatusCode(v int32) *CreateInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceResponse) SetBody(v *CreateInstanceResponseBody) *CreateInstanceResponse {
	s.Body = v
	return s
}

type DeleteHoloWarehouseRequest struct {
	// The name of the virtual warehouse.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_warehouse
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DeleteHoloWarehouseRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteHoloWarehouseRequest) GoString() string {
	return s.String()
}

func (s *DeleteHoloWarehouseRequest) SetName(v string) *DeleteHoloWarehouseRequest {
	s.Name = &v
	return s
}

type DeleteHoloWarehouseResponseBody struct {
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2A8DEF6E-067E-5DB0-BAE1-2894266E6C6A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHoloWarehouseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteHoloWarehouseResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHoloWarehouseResponseBody) SetData(v bool) *DeleteHoloWarehouseResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteHoloWarehouseResponseBody) SetRequestId(v string) *DeleteHoloWarehouseResponseBody {
	s.RequestId = &v
	return s
}

type DeleteHoloWarehouseResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHoloWarehouseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHoloWarehouseResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteHoloWarehouseResponse) GoString() string {
	return s.String()
}

func (s *DeleteHoloWarehouseResponse) SetHeaders(v map[string]*string) *DeleteHoloWarehouseResponse {
	s.Headers = v
	return s
}

func (s *DeleteHoloWarehouseResponse) SetStatusCode(v int32) *DeleteHoloWarehouseResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHoloWarehouseResponse) SetBody(v *DeleteHoloWarehouseResponseBody) *DeleteHoloWarehouseResponse {
	s.Body = v
	return s
}

type DeleteInstanceRequest struct {
	// The ID of the region in which the Hologres instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequest) SetRegionId(v string) *DeleteInstanceRequest {
	s.RegionId = &v
	return s
}

type DeleteInstanceResponseBody struct {
	// The returned result, which indicates whether the request was successful.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned.
	//
	// example:
	//
	// null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// null
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status Code
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CB13FFDD-2DF8-5396-A848-2D6A31245B6D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponseBody) SetData(v bool) *DeleteInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetErrorCode(v string) *DeleteInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetErrorMessage(v string) *DeleteInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetHttpStatusCode(v string) *DeleteInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetRequestId(v string) *DeleteInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetSuccess(v bool) *DeleteInstanceResponseBody {
	s.Success = &v
	return s
}

type DeleteInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponse) SetHeaders(v map[string]*string) *DeleteInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceResponse) SetStatusCode(v int32) *DeleteInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceResponse) SetBody(v *DeleteInstanceResponseBody) *DeleteInstanceResponse {
	s.Body = v
	return s
}

type DisableHiveAccessRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DisableHiveAccessRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableHiveAccessRequest) GoString() string {
	return s.String()
}

func (s *DisableHiveAccessRequest) SetRegionId(v string) *DisableHiveAccessRequest {
	s.RegionId = &v
	return s
}

type DisableHiveAccessResponseBody struct {
	// The returned result.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// 404
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// Internal server error.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 82B7A554-4D00-50DF-95D9-B59E7B4D5489
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableHiveAccessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableHiveAccessResponseBody) GoString() string {
	return s.String()
}

func (s *DisableHiveAccessResponseBody) SetData(v bool) *DisableHiveAccessResponseBody {
	s.Data = &v
	return s
}

func (s *DisableHiveAccessResponseBody) SetErrorCode(v string) *DisableHiveAccessResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DisableHiveAccessResponseBody) SetErrorMessage(v string) *DisableHiveAccessResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DisableHiveAccessResponseBody) SetHttpStatusCode(v string) *DisableHiveAccessResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DisableHiveAccessResponseBody) SetRequestId(v string) *DisableHiveAccessResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableHiveAccessResponseBody) SetSuccess(v bool) *DisableHiveAccessResponseBody {
	s.Success = &v
	return s
}

type DisableHiveAccessResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableHiveAccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableHiveAccessResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableHiveAccessResponse) GoString() string {
	return s.String()
}

func (s *DisableHiveAccessResponse) SetHeaders(v map[string]*string) *DisableHiveAccessResponse {
	s.Headers = v
	return s
}

func (s *DisableHiveAccessResponse) SetStatusCode(v int32) *DisableHiveAccessResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableHiveAccessResponse) SetBody(v *DisableHiveAccessResponseBody) *DisableHiveAccessResponse {
	s.Body = v
	return s
}

type EnableHiveAccessRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s EnableHiveAccessRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableHiveAccessRequest) GoString() string {
	return s.String()
}

func (s *EnableHiveAccessRequest) SetRegionId(v string) *EnableHiveAccessRequest {
	s.RegionId = &v
	return s
}

type EnableHiveAccessResponseBody struct {
	// The returned data.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// 404
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// Internal server error.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EA8F0084-5831-5907-BB31-BD05D2617844
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableHiveAccessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableHiveAccessResponseBody) GoString() string {
	return s.String()
}

func (s *EnableHiveAccessResponseBody) SetData(v bool) *EnableHiveAccessResponseBody {
	s.Data = &v
	return s
}

func (s *EnableHiveAccessResponseBody) SetErrorCode(v string) *EnableHiveAccessResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *EnableHiveAccessResponseBody) SetErrorMessage(v string) *EnableHiveAccessResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *EnableHiveAccessResponseBody) SetHttpStatusCode(v string) *EnableHiveAccessResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *EnableHiveAccessResponseBody) SetRequestId(v string) *EnableHiveAccessResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableHiveAccessResponseBody) SetSuccess(v bool) *EnableHiveAccessResponseBody {
	s.Success = &v
	return s
}

type EnableHiveAccessResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableHiveAccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableHiveAccessResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableHiveAccessResponse) GoString() string {
	return s.String()
}

func (s *EnableHiveAccessResponse) SetHeaders(v map[string]*string) *EnableHiveAccessResponse {
	s.Headers = v
	return s
}

func (s *EnableHiveAccessResponse) SetStatusCode(v int32) *EnableHiveAccessResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableHiveAccessResponse) SetBody(v *EnableHiveAccessResponseBody) *EnableHiveAccessResponse {
	s.Body = v
	return s
}

type GetInstanceResponseBody struct {
	// The error code that is returned if the request failed.
	//
	// example:
	//
	// 404
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Internal server error.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The information about the instance.
	Instance *GetInstanceResponseBodyInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 865A02C2-B374-5DD4-9B34-0CA15DA1AEBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request result, which indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) SetErrorCode(v string) *GetInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetInstanceResponseBody) SetErrorMessage(v string) *GetInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetInstanceResponseBody) SetHttpStatusCode(v string) *GetInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstance(v *GetInstanceResponseBodyInstance) *GetInstanceResponseBody {
	s.Instance = v
	return s
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponseBody) SetSuccess(v bool) *GetInstanceResponseBody {
	s.Success = &v
	return s
}

type GetInstanceResponseBodyInstance struct {
	// Indicates whether auto-renewal is enabled.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	AutoRenewal *string `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// The cold storage capacity of the instance. Unit: GB. Standard SSD is used for hot storage, and HDD is used for cold storage.
	//
	// example:
	//
	// 800
	ColdStorage *int64 `json:"ColdStorage,omitempty" xml:"ColdStorage,omitempty"`
	// The commodity code.
	//
	// Valid values:
	//
	// 	- hologram_maxcomputeAccelerate_public_cn
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     China site/Lakehouse Acceleration Edition
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- hologram_combo_public_cn
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     China site/Subscription
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- hologram_prepay_public_intl
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     International site/Subscription
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- hologram_storage_dp_cn
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     China site/Storage plan
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- hologram_postpay_public_cn
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     China site/Pay-as-you-go
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- hologram_postpay_public_intl
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     International site/Pay-as-you-go
	//
	//     <!-- -->
	//
	// 	- hologram_maxcomputeAccelerate_public_intl
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     International site/Lakehouse Acceleration Edition
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- hologram_cu_dp_cn
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     China site/Compute plan
	//
	//     <!-- -->
	//
	// example:
	//
	// hologram_combo_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The number of compute nodes. In a typical configuration, a node has 16 CPU cores and 32 GB of memory.
	//
	// example:
	//
	// 2
	ComputeNodeCount *int64 `json:"ComputeNodeCount,omitempty" xml:"ComputeNodeCount,omitempty"`
	// The number of CPU cores.
	//
	// example:
	//
	// 32
	Cpu *int64 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2021-02-03T13:06:06Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The amount of data that can be stored in the disk of the Standard storage class. Unit: GB.
	//
	// example:
	//
	// 500
	Disk *string `json:"Disk,omitempty" xml:"Disk,omitempty"`
	// Indicates whether data lake acceleration is enabled.
	//
	// example:
	//
	// true
	EnableHiveAccess *string `json:"EnableHiveAccess,omitempty" xml:"EnableHiveAccess,omitempty"`
	// EnableServerless
	//
	// example:
	//
	// true
	EnableServerless *bool `json:"EnableServerless,omitempty" xml:"EnableServerless,omitempty"`
	// The list of endpoints.
	Endpoints []*GetInstanceResponseBodyInstanceEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// The expiration time. This parameter is invalid for pay-as-you-go instances.
	//
	// example:
	//
	// 2021-02-03T13:06:06Z
	ExpirationTime *string `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// The number of gateway nodes.
	//
	// example:
	//
	// 2
	GatewayCount *int64 `json:"GatewayCount,omitempty" xml:"GatewayCount,omitempty"`
	// The number of CPU cores of the gateway. Unit: core.
	//
	// example:
	//
	// 4
	GatewayCpu *int64 `json:"GatewayCpu,omitempty" xml:"GatewayCpu,omitempty"`
	// The size of memory resources of the gateway. Unit: GB.
	//
	// example:
	//
	// 16
	GatewayMemory *int64 `json:"GatewayMemory,omitempty" xml:"GatewayMemory,omitempty"`
	// The billing method of the instance.
	//
	// Valid values:
	//
	// 	- PostPaid
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     pay-as-you-go
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- PrePaid
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     subscription
	//
	//     <!-- -->
	//
	//     .
	//
	// example:
	//
	// PrePaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// hgpostcn-cn-tl32s6cgw00b
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name. The instance name must be 2 to 64 characters in length.
	//
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The owner of the instance.
	//
	// example:
	//
	// 12345678900000
	InstanceOwner *string `json:"InstanceOwner,omitempty" xml:"InstanceOwner,omitempty"`
	// The status of the instance.
	//
	// Valid values:
	//
	// 	- Creating
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Running
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Suspended
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Allocating
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Running
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The type of the instance.
	//
	// Valid values:
	//
	// 	- Follower
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     read-only secondary instance
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- Standard
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     normal instance
	//
	//     <!-- -->
	//
	//     .
	//
	// example:
	//
	// Standard
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The ID of the primary instance.
	//
	// example:
	//
	// hgpostcn-cn-i7m2ncd6w002
	LeaderInstanceId *string `json:"LeaderInstanceId,omitempty" xml:"LeaderInstanceId,omitempty"`
	// The memory size. Unit: GB.
	//
	// example:
	//
	// 128
	Memory *int64 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Disaster recovery instance role.
	//
	// 	- Active: Primary disaster recovery instance.
	//
	// 	- Passive: Disaster tolerance instance.
	//
	// 	- PreActive: Primary disaster recovery instance not yet in final state.
	//
	// example:
	//
	// Active
	ReplicaRole *string `json:"ReplicaRole,omitempty" xml:"ReplicaRole,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzuq7hpybze2i
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The storage type.
	//
	// 	- redundant: 3 copies
	//
	// 	- local: single copy
	//
	// example:
	//
	// redundant
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The reason for the suspension.
	//
	// Valid values:
	//
	// 	- Indebet
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     The instance has an overdue payment
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- Manual
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     The instance is manually suspended
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- Overdue
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     The instance has expired
	//
	//     <!-- -->
	//
	//     .
	//
	// example:
	//
	// Manual
	SuspendReason *string `json:"SuspendReason,omitempty" xml:"SuspendReason,omitempty"`
	// The instance tag.
	Tags []*GetInstanceResponseBodyInstanceTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The instance version.
	//
	// example:
	//
	// r1.3.37
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The ID of the zone where the instance resides.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetInstanceResponseBodyInstance) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyInstance) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstance) SetAutoRenewal(v string) *GetInstanceResponseBodyInstance {
	s.AutoRenewal = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetColdStorage(v int64) *GetInstanceResponseBodyInstance {
	s.ColdStorage = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetCommodityCode(v string) *GetInstanceResponseBodyInstance {
	s.CommodityCode = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetComputeNodeCount(v int64) *GetInstanceResponseBodyInstance {
	s.ComputeNodeCount = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetCpu(v int64) *GetInstanceResponseBodyInstance {
	s.Cpu = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetCreationTime(v string) *GetInstanceResponseBodyInstance {
	s.CreationTime = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetDisk(v string) *GetInstanceResponseBodyInstance {
	s.Disk = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetEnableHiveAccess(v string) *GetInstanceResponseBodyInstance {
	s.EnableHiveAccess = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetEnableServerless(v bool) *GetInstanceResponseBodyInstance {
	s.EnableServerless = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetEndpoints(v []*GetInstanceResponseBodyInstanceEndpoints) *GetInstanceResponseBodyInstance {
	s.Endpoints = v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetExpirationTime(v string) *GetInstanceResponseBodyInstance {
	s.ExpirationTime = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetGatewayCount(v int64) *GetInstanceResponseBodyInstance {
	s.GatewayCount = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetGatewayCpu(v int64) *GetInstanceResponseBodyInstance {
	s.GatewayCpu = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetGatewayMemory(v int64) *GetInstanceResponseBodyInstance {
	s.GatewayMemory = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetInstanceChargeType(v string) *GetInstanceResponseBodyInstance {
	s.InstanceChargeType = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetInstanceId(v string) *GetInstanceResponseBodyInstance {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetInstanceName(v string) *GetInstanceResponseBodyInstance {
	s.InstanceName = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetInstanceOwner(v string) *GetInstanceResponseBodyInstance {
	s.InstanceOwner = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetInstanceStatus(v string) *GetInstanceResponseBodyInstance {
	s.InstanceStatus = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetInstanceType(v string) *GetInstanceResponseBodyInstance {
	s.InstanceType = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetLeaderInstanceId(v string) *GetInstanceResponseBodyInstance {
	s.LeaderInstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetMemory(v int64) *GetInstanceResponseBodyInstance {
	s.Memory = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetRegionId(v string) *GetInstanceResponseBodyInstance {
	s.RegionId = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetReplicaRole(v string) *GetInstanceResponseBodyInstance {
	s.ReplicaRole = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetResourceGroupId(v string) *GetInstanceResponseBodyInstance {
	s.ResourceGroupId = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetStorageType(v string) *GetInstanceResponseBodyInstance {
	s.StorageType = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetSuspendReason(v string) *GetInstanceResponseBodyInstance {
	s.SuspendReason = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetTags(v []*GetInstanceResponseBodyInstanceTags) *GetInstanceResponseBodyInstance {
	s.Tags = v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetVersion(v string) *GetInstanceResponseBodyInstance {
	s.Version = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetZoneId(v string) *GetInstanceResponseBodyInstance {
	s.ZoneId = &v
	return s
}

type GetInstanceResponseBodyInstanceEndpoints struct {
	// The endpoint. This parameter is returned if both the AnyTunnel and SingleTunnel modes are enabled for an instance, and the instance is switched from the AnyTunnel mode to the SingleTunnel mode. In this case, two endpoints are returned.
	//
	// example:
	//
	// hgprecn-cn-uqm362o1b001-cn-hangzhou-internal.hologres.aliyuncs.com:80
	AlternativeEndpoints *string `json:"AlternativeEndpoints,omitempty" xml:"AlternativeEndpoints,omitempty"`
	// Indicates whether the network is enabled.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The endpoint.
	//
	// example:
	//
	// hgprecn-cn-uqm362o1b001-cn-hangzhou-internal.hologres.aliyuncs.com:80
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The network type.
	//
	// Valid values:
	//
	// 	- VPCSingleTunnel
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     virtual private cloud (VPC)
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- Intranet
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     internal network
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- VPCAnyTunnel
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     not supported by new instances
	//
	//     <!-- -->
	//
	// 	- Internet
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     Internet
	//
	//     <!-- -->
	//
	//     .
	//
	// example:
	//
	// Internet
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-bp1jqwp2ys6kp7tc9t983
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC to which the instance belongs.
	//
	// example:
	//
	// vpc-uf66jjber3hgvwhki3wna
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the instance that is deployed in the VPC.
	//
	// example:
	//
	// hgprecn-cn-uqm362o1b001-frontend-st
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
}

func (s GetInstanceResponseBodyInstanceEndpoints) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyInstanceEndpoints) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstanceEndpoints) SetAlternativeEndpoints(v string) *GetInstanceResponseBodyInstanceEndpoints {
	s.AlternativeEndpoints = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceEndpoints) SetEnabled(v bool) *GetInstanceResponseBodyInstanceEndpoints {
	s.Enabled = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceEndpoints) SetEndpoint(v string) *GetInstanceResponseBodyInstanceEndpoints {
	s.Endpoint = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceEndpoints) SetType(v string) *GetInstanceResponseBodyInstanceEndpoints {
	s.Type = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceEndpoints) SetVSwitchId(v string) *GetInstanceResponseBodyInstanceEndpoints {
	s.VSwitchId = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceEndpoints) SetVpcId(v string) *GetInstanceResponseBodyInstanceEndpoints {
	s.VpcId = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceEndpoints) SetVpcInstanceId(v string) *GetInstanceResponseBodyInstanceEndpoints {
	s.VpcInstanceId = &v
	return s
}

type GetInstanceResponseBodyInstanceTags struct {
	// The key of tag N.
	//
	// example:
	//
	// tag
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetInstanceResponseBodyInstanceTags) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyInstanceTags) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstanceTags) SetKey(v string) *GetInstanceResponseBodyInstanceTags {
	s.Key = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceTags) SetValue(v string) *GetInstanceResponseBodyInstanceTags {
	s.Value = &v
	return s
}

type GetInstanceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceResponse) SetHeaders(v map[string]*string) *GetInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceResponse) SetStatusCode(v int32) *GetInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceResponse) SetBody(v *GetInstanceResponseBody) *GetInstanceResponse {
	s.Body = v
	return s
}

type GetWarehouseDetailResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D3AE84AB-0873-5FC7-A4C4-8CF869D2FA70
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned values.
	WarehouseDetail *GetWarehouseDetailResponseBodyWarehouseDetail `json:"WarehouseDetail,omitempty" xml:"WarehouseDetail,omitempty" type:"Struct"`
}

func (s GetWarehouseDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWarehouseDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetWarehouseDetailResponseBody) SetRequestId(v string) *GetWarehouseDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWarehouseDetailResponseBody) SetWarehouseDetail(v *GetWarehouseDetailResponseBodyWarehouseDetail) *GetWarehouseDetailResponseBody {
	s.WarehouseDetail = v
	return s
}

type GetWarehouseDetailResponseBodyWarehouseDetail struct {
	// The remaining unallocated computing resources of the virtual warehouse instance.
	//
	// example:
	//
	// 32
	RemainingCpu *string `json:"RemainingCpu,omitempty" xml:"RemainingCpu,omitempty"`
	// The reserved computing resources. The amount of computing resources in all running virtual warehouses in an instance cannot exceed the amount of reserved computing resources in the virtual warehouses.
	//
	// example:
	//
	// 64
	ReservedCpu     *string `json:"ReservedCpu,omitempty" xml:"ReservedCpu,omitempty"`
	TimedElasticCpu *string `json:"TimedElasticCpu,omitempty" xml:"TimedElasticCpu,omitempty"`
	// The list of virtual warehouses.
	WarehouseList []*GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList `json:"WarehouseList,omitempty" xml:"WarehouseList,omitempty" type:"Repeated"`
}

func (s GetWarehouseDetailResponseBodyWarehouseDetail) String() string {
	return tea.Prettify(s)
}

func (s GetWarehouseDetailResponseBodyWarehouseDetail) GoString() string {
	return s.String()
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetail) SetRemainingCpu(v string) *GetWarehouseDetailResponseBodyWarehouseDetail {
	s.RemainingCpu = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetail) SetReservedCpu(v string) *GetWarehouseDetailResponseBodyWarehouseDetail {
	s.ReservedCpu = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetail) SetTimedElasticCpu(v string) *GetWarehouseDetailResponseBodyWarehouseDetail {
	s.TimedElasticCpu = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetail) SetWarehouseList(v []*GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) *GetWarehouseDetailResponseBodyWarehouseDetail {
	s.WarehouseList = v
	return s
}

type GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList struct {
	// The number of CPU cores.
	//
	// example:
	//
	// 32
	Cpu              *int64 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	DefaultWarehouse *bool  `json:"DefaultWarehouse,omitempty" xml:"DefaultWarehouse,omitempty"`
	ElasticCpu       *int64 `json:"ElasticCpu,omitempty" xml:"ElasticCpu,omitempty"`
	// The ID.
	//
	// example:
	//
	// 2
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The memory capacity.
	//
	// example:
	//
	// 128
	Mem *int64 `json:"Mem,omitempty" xml:"Mem,omitempty"`
	// The name of the virtual warehouse instance.
	//
	// example:
	//
	// MyWarehouse
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of compute nodes.
	//
	// example:
	//
	// 2
	NodeCount       *int64  `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	RebalanceStatus *string `json:"RebalanceStatus,omitempty" xml:"RebalanceStatus,omitempty"`
	// The status.
	//
	// Valid values:
	//
	// 	- kRunning
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- kSuspended
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- kInit
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- kFailed
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- kAllocating
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// kRunning
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) String() string {
	return tea.Prettify(s)
}

func (s GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) GoString() string {
	return s.String()
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) SetCpu(v int64) *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	s.Cpu = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) SetDefaultWarehouse(v bool) *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	s.DefaultWarehouse = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) SetElasticCpu(v int64) *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	s.ElasticCpu = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) SetId(v int64) *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	s.Id = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) SetMem(v int64) *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	s.Mem = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) SetName(v string) *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	s.Name = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) SetNodeCount(v int64) *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	s.NodeCount = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) SetRebalanceStatus(v string) *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	s.RebalanceStatus = &v
	return s
}

func (s *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList) SetStatus(v string) *GetWarehouseDetailResponseBodyWarehouseDetailWarehouseList {
	s.Status = &v
	return s
}

type GetWarehouseDetailResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWarehouseDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWarehouseDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWarehouseDetailResponse) GoString() string {
	return s.String()
}

func (s *GetWarehouseDetailResponse) SetHeaders(v map[string]*string) *GetWarehouseDetailResponse {
	s.Headers = v
	return s
}

func (s *GetWarehouseDetailResponse) SetStatusCode(v int32) *GetWarehouseDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWarehouseDetailResponse) SetBody(v *GetWarehouseDetailResponseBody) *GetWarehouseDetailResponse {
	s.Body = v
	return s
}

type ListBackupDataRequest struct {
	// The backup type. Specific backup data is filtered based on the type. If you leave this parameter empty, all backup data is returned.
	//
	// Valid values:
	//
	// 	- redundant_remote
	//
	// 	- remote
	//
	// 	- redundant
	//
	// 	- full_remote
	//
	// 	- local
	//
	// 	- full
	//
	// example:
	//
	// redundant
	BackupType *string `json:"backupType,omitempty" xml:"backupType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// hgprecn-cn-wwoxxxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s ListBackupDataRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBackupDataRequest) GoString() string {
	return s.String()
}

func (s *ListBackupDataRequest) SetBackupType(v string) *ListBackupDataRequest {
	s.BackupType = &v
	return s
}

func (s *ListBackupDataRequest) SetInstanceId(v string) *ListBackupDataRequest {
	s.InstanceId = &v
	return s
}

type ListBackupDataResponseBody struct {
	// The backups.
	BackupDataList []*ListBackupDataResponseBodyBackupDataList `json:"BackupDataList,omitempty" xml:"BackupDataList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 4AA0C48F-B5BB-5FF9-A43B-6B91E0715D46
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListBackupDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBackupDataResponseBody) GoString() string {
	return s.String()
}

func (s *ListBackupDataResponseBody) SetBackupDataList(v []*ListBackupDataResponseBodyBackupDataList) *ListBackupDataResponseBody {
	s.BackupDataList = v
	return s
}

func (s *ListBackupDataResponseBody) SetRequestId(v string) *ListBackupDataResponseBody {
	s.RequestId = &v
	return s
}

type ListBackupDataResponseBodyBackupDataList struct {
	// The backup type. In general, the following two types are supported: local backup and remote backup. In the local backup type, snapshots reside in the same region as your instance. The following two sub-types are available: full (single backup, single replica) and redundant (zone-redundant storage, multiple replicas). In the remote backup type, snapshots and your instance reside in different regions. Remote backups are the replicas of the backups of the full or redundant type in another region. The values local and remote do not represent specific types, but are used only for data filtering. The value local indicates all local backups, and the value remote indicates all remote backups.
	//
	// example:
	//
	// redundant
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// The size of cold data. Unit: bytes.
	//
	// example:
	//
	// 32413521
	ColdDataSize *int64 `json:"ColdDataSize,omitempty" xml:"ColdDataSize,omitempty"`
	// The description of the backup data.
	//
	// example:
	//
	// demo
	DataDesc *string `json:"DataDesc,omitempty" xml:"DataDesc,omitempty"`
	// The backup granularity.
	//
	// Valid values:
	//
	// 	- instance
	//
	// example:
	//
	// instance
	DataGran *string `json:"DataGran,omitempty" xml:"DataGran,omitempty"`
	// The size of the backup data. Unit: bytes.
	//
	// example:
	//
	// 76085723136
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// The snapshot time. The value format of this parameter follows the same standard as that of the StartTime parameter.
	//
	// example:
	//
	// 2024-10-28T12:23:37.000+00:00
	DataTime *string `json:"DataTime,omitempty" xml:"DataTime,omitempty"`
	// The end time of the backup task. The value format of this parameter follows the same standard as that of the StartTime parameter.
	//
	// example:
	//
	// 2024-10-28T12:27:34.000+00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The unique ID of the backup.
	//
	// example:
	//
	// 1780805690994479105
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// hgpostcn-cn-pe33jdxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// my-hologres-dw
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	InstanceRegion *string `json:"InstanceRegion,omitempty" xml:"InstanceRegion,omitempty"`
	// The type of the instance.
	//
	// Valid values:
	//
	// 	- Warehouse: virtual warehouse instance
	//
	// 	- Standard: general-purpose instance
	//
	// example:
	//
	// Warehouse
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The zone in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou-j
	InstanceZoneId *string `json:"InstanceZoneId,omitempty" xml:"InstanceZoneId,omitempty"`
	// The region in which the backup data resides.
	//
	// example:
	//
	// cn-hangzhou
	SnapshotRegion *string `json:"SnapshotRegion,omitempty" xml:"SnapshotRegion,omitempty"`
	// The zone in which the backup data resides. In zone-redundant storage mode, backup data is stored in different zones, including the current zone.
	//
	// example:
	//
	// cn-hangzhou-j
	SnapshotZoneId *string `json:"SnapshotZoneId,omitempty" xml:"SnapshotZoneId,omitempty"`
	// The start time of the backup task. The time follows the ISO 8601 standard in the YYYY-MM-DDTHH:mm:ss.SSSTZ format. The time is displayed in UTC (the same below).
	//
	// example:
	//
	// 2024-10-28T11:19:56.000+00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the backup task.
	//
	// Valid values:
	//
	// 	- processing
	//
	// 	- completed
	//
	// 	- failed
	//
	// example:
	//
	// completed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The mode in which the backup task is triggered.
	//
	// Valid values:
	//
	// 	- scheduled: periodic backup
	//
	// 	- manual: manual backup
	//
	// example:
	//
	// scheduled
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s ListBackupDataResponseBodyBackupDataList) String() string {
	return tea.Prettify(s)
}

func (s ListBackupDataResponseBodyBackupDataList) GoString() string {
	return s.String()
}

func (s *ListBackupDataResponseBodyBackupDataList) SetBackupType(v string) *ListBackupDataResponseBodyBackupDataList {
	s.BackupType = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) SetColdDataSize(v int64) *ListBackupDataResponseBodyBackupDataList {
	s.ColdDataSize = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) SetDataDesc(v string) *ListBackupDataResponseBodyBackupDataList {
	s.DataDesc = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) SetDataGran(v string) *ListBackupDataResponseBodyBackupDataList {
	s.DataGran = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) SetDataSize(v int64) *ListBackupDataResponseBodyBackupDataList {
	s.DataSize = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) SetDataTime(v string) *ListBackupDataResponseBodyBackupDataList {
	s.DataTime = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) SetEndTime(v string) *ListBackupDataResponseBodyBackupDataList {
	s.EndTime = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) SetId(v int64) *ListBackupDataResponseBodyBackupDataList {
	s.Id = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) SetInstanceId(v string) *ListBackupDataResponseBodyBackupDataList {
	s.InstanceId = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) SetInstanceName(v string) *ListBackupDataResponseBodyBackupDataList {
	s.InstanceName = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) SetInstanceRegion(v string) *ListBackupDataResponseBodyBackupDataList {
	s.InstanceRegion = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) SetInstanceType(v string) *ListBackupDataResponseBodyBackupDataList {
	s.InstanceType = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) SetInstanceZoneId(v string) *ListBackupDataResponseBodyBackupDataList {
	s.InstanceZoneId = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) SetSnapshotRegion(v string) *ListBackupDataResponseBodyBackupDataList {
	s.SnapshotRegion = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) SetSnapshotZoneId(v string) *ListBackupDataResponseBodyBackupDataList {
	s.SnapshotZoneId = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) SetStartTime(v string) *ListBackupDataResponseBodyBackupDataList {
	s.StartTime = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) SetStatus(v string) *ListBackupDataResponseBodyBackupDataList {
	s.Status = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) SetTriggerType(v string) *ListBackupDataResponseBodyBackupDataList {
	s.TriggerType = &v
	return s
}

type ListBackupDataResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBackupDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBackupDataResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBackupDataResponse) GoString() string {
	return s.String()
}

func (s *ListBackupDataResponse) SetHeaders(v map[string]*string) *ListBackupDataResponse {
	s.Headers = v
	return s
}

func (s *ListBackupDataResponse) SetStatusCode(v int32) *ListBackupDataResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBackupDataResponse) SetBody(v *ListBackupDataResponseBody) *ListBackupDataResponse {
	s.Body = v
	return s
}

type ListInstancesRequest struct {
	// example:
	//
	// standard
	CmsInstanceType *string `json:"cmsInstanceType,omitempty" xml:"cmsInstanceType,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmvscak73zmby
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The tags to add to the resource.
	Tag []*ListInstancesRequestTag `json:"tag,omitempty" xml:"tag,omitempty" type:"Repeated"`
}

func (s ListInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) SetCmsInstanceType(v string) *ListInstancesRequest {
	s.CmsInstanceType = &v
	return s
}

func (s *ListInstancesRequest) SetResourceGroupId(v string) *ListInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesRequest) SetTag(v []*ListInstancesRequestTag) *ListInstancesRequest {
	s.Tag = v
	return s
}

type ListInstancesRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// mytag
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListInstancesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *ListInstancesRequestTag) SetKey(v string) *ListInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *ListInstancesRequestTag) SetValue(v string) *ListInstancesRequestTag {
	s.Value = &v
	return s
}

type ListInstancesResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// 404
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// Internal server error.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The instances.
	InstanceList []*ListInstancesResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D1303CD4-AA70-5998-8025-F55B22C50840
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) SetErrorCode(v string) *ListInstancesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListInstancesResponseBody) SetErrorMessage(v string) *ListInstancesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListInstancesResponseBody) SetHttpStatusCode(v string) *ListInstancesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInstancesResponseBody) SetInstanceList(v []*ListInstancesResponseBodyInstanceList) *ListInstancesResponseBody {
	s.InstanceList = v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetSuccess(v string) *ListInstancesResponseBody {
	s.Success = &v
	return s
}

type ListInstancesResponseBodyInstanceList struct {
	// The commodity code, which is the same as that on the Billing Management page.
	//
	// example:
	//
	// hologram_postpay_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The time when the cluster was created.
	//
	// example:
	//
	// 2022-12-16T02:24:05Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Indicates whether lakehouse acceleration is enabled.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	EnableHiveAccess *string `json:"EnableHiveAccess,omitempty" xml:"EnableHiveAccess,omitempty"`
	// The list of endpoints.
	Endpoints []*ListInstancesResponseBodyInstanceListEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// The time when the cluster expires.
	//
	// example:
	//
	// 2023-05-04T16:00:00.000Z
	ExpirationTime *string `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// Valid values:
	//
	// 	- PostPaid
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     pay-as-you-go
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- PrePaid
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     subscription
	//
	//     <!-- -->
	//
	//     .
	//
	// example:
	//
	// PrePaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// hgpostcn-cn-aaab9ad2d8fb
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// test_instance
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The status of the instance.
	//
	// Valid values:
	//
	// 	- Creating
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Running
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Suspended
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Allocating
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Running
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The type of the instance.
	//
	// Valid values:
	//
	// 	- Follower
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     read-only secondary instance
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- Standard
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     normal instance
	//
	//     <!-- -->
	//
	//     .
	//
	// example:
	//
	// Standard
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The ID of the primary instance.
	//
	// example:
	//
	// hgprecn-cn-2r42sqvxm006
	LeaderInstanceId *string `json:"LeaderInstanceId,omitempty" xml:"LeaderInstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmvscak73zmby
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StorageType     *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The reason for the suspension.
	//
	// example:
	//
	// Manual
	SuspendReason *string `json:"SuspendReason,omitempty" xml:"SuspendReason,omitempty"`
	// The tags that are added to the resource.
	Tags []*ListInstancesResponseBodyInstanceListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The version of the cluster.
	//
	// example:
	//
	// 1.3.37
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListInstancesResponseBodyInstanceList) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstanceList) SetCommodityCode(v string) *ListInstancesResponseBodyInstanceList {
	s.CommodityCode = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetCreationTime(v string) *ListInstancesResponseBodyInstanceList {
	s.CreationTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetEnableHiveAccess(v string) *ListInstancesResponseBodyInstanceList {
	s.EnableHiveAccess = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetEndpoints(v []*ListInstancesResponseBodyInstanceListEndpoints) *ListInstancesResponseBodyInstanceList {
	s.Endpoints = v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetExpirationTime(v string) *ListInstancesResponseBodyInstanceList {
	s.ExpirationTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetInstanceChargeType(v string) *ListInstancesResponseBodyInstanceList {
	s.InstanceChargeType = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetInstanceId(v string) *ListInstancesResponseBodyInstanceList {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetInstanceName(v string) *ListInstancesResponseBodyInstanceList {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetInstanceStatus(v string) *ListInstancesResponseBodyInstanceList {
	s.InstanceStatus = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetInstanceType(v string) *ListInstancesResponseBodyInstanceList {
	s.InstanceType = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetLeaderInstanceId(v string) *ListInstancesResponseBodyInstanceList {
	s.LeaderInstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetRegionId(v string) *ListInstancesResponseBodyInstanceList {
	s.RegionId = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetResourceGroupId(v string) *ListInstancesResponseBodyInstanceList {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetStorageType(v string) *ListInstancesResponseBodyInstanceList {
	s.StorageType = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetSuspendReason(v string) *ListInstancesResponseBodyInstanceList {
	s.SuspendReason = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetTags(v []*ListInstancesResponseBodyInstanceListTags) *ListInstancesResponseBodyInstanceList {
	s.Tags = v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetVersion(v string) *ListInstancesResponseBodyInstanceList {
	s.Version = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetZoneId(v string) *ListInstancesResponseBodyInstanceList {
	s.ZoneId = &v
	return s
}

type ListInstancesResponseBodyInstanceListEndpoints struct {
	// Indicates whether the endpoint is enabled.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The endpoint.
	//
	// example:
	//
	// hgpostcn-cn-aaab9ad2d8fb-cn-hangzhou-internal.hologres.aliyuncs.com:80
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The network type.
	//
	// Valid values:
	//
	// 	- VPCSingleTunnel
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     virtual private cloud (VPC)
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- Intranet
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     internal network
	//
	//     <!-- -->
	//
	// 	- VPCAnyTunnel
	//
	//     <!-- -->
	//
	//     : This value is not supported by new instances
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- Internet
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     Internet
	//
	//     <!-- -->
	//
	//     .
	//
	// example:
	//
	// Internet
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-wz9oap28raidjevhuszg4
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-uf6mrahzyu7uorlqqpz5f
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the VPC to which the instance belongs.
	//
	// example:
	//
	// hgpostcn-cn-wwo3665tx004-frontend-st
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
}

func (s ListInstancesResponseBodyInstanceListEndpoints) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstanceListEndpoints) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstanceListEndpoints) SetEnabled(v bool) *ListInstancesResponseBodyInstanceListEndpoints {
	s.Enabled = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListEndpoints) SetEndpoint(v string) *ListInstancesResponseBodyInstanceListEndpoints {
	s.Endpoint = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListEndpoints) SetType(v string) *ListInstancesResponseBodyInstanceListEndpoints {
	s.Type = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListEndpoints) SetVSwitchId(v string) *ListInstancesResponseBodyInstanceListEndpoints {
	s.VSwitchId = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListEndpoints) SetVpcId(v string) *ListInstancesResponseBodyInstanceListEndpoints {
	s.VpcId = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListEndpoints) SetVpcInstanceId(v string) *ListInstancesResponseBodyInstanceListEndpoints {
	s.VpcInstanceId = &v
	return s
}

type ListInstancesResponseBodyInstanceListTags struct {
	// The tag key.
	//
	// example:
	//
	// tag
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListInstancesResponseBodyInstanceListTags) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstanceListTags) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstanceListTags) SetKey(v string) *ListInstancesResponseBodyInstanceListTags {
	s.Key = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceListTags) SetValue(v string) *ListInstancesResponseBodyInstanceListTags {
	s.Value = &v
	return s
}

type ListInstancesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesResponse) SetHeaders(v map[string]*string) *ListInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListInstancesResponse) SetStatusCode(v int32) *ListInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancesResponse) SetBody(v *ListInstancesResponseBody) *ListInstancesResponse {
	s.Body = v
	return s
}

type ListWarehousesResponseBody struct {
	// The list of virtual warehouse instances.
	WarehouseList []*ListWarehousesResponseBodyWarehouseList `json:"WarehouseList,omitempty" xml:"WarehouseList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 819A7F0F-2951-540F-BD94-6A41ECF0281F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListWarehousesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWarehousesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWarehousesResponseBody) SetWarehouseList(v []*ListWarehousesResponseBodyWarehouseList) *ListWarehousesResponseBody {
	s.WarehouseList = v
	return s
}

func (s *ListWarehousesResponseBody) SetRequestId(v string) *ListWarehousesResponseBody {
	s.RequestId = &v
	return s
}

type ListWarehousesResponseBodyWarehouseList struct {
	// The number of CPU cores.
	//
	// example:
	//
	// 32
	Cpu *int64 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The ID.
	//
	// example:
	//
	// 3
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The memory capacity.
	//
	// example:
	//
	// 128
	Mem *int64 `json:"Mem,omitempty" xml:"Mem,omitempty"`
	// The name of the virtual warehouse instance.
	//
	// example:
	//
	// MyWarehouse
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of compute nodes.
	//
	// example:
	//
	// 2
	NodeCount *int64 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// The status.
	//
	// Valid values:
	//
	// 	- kRunning
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- kSuspended
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- kInit
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- kFailed
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- kAllocating
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// kRunning
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListWarehousesResponseBodyWarehouseList) String() string {
	return tea.Prettify(s)
}

func (s ListWarehousesResponseBodyWarehouseList) GoString() string {
	return s.String()
}

func (s *ListWarehousesResponseBodyWarehouseList) SetCpu(v int64) *ListWarehousesResponseBodyWarehouseList {
	s.Cpu = &v
	return s
}

func (s *ListWarehousesResponseBodyWarehouseList) SetId(v int64) *ListWarehousesResponseBodyWarehouseList {
	s.Id = &v
	return s
}

func (s *ListWarehousesResponseBodyWarehouseList) SetMem(v int64) *ListWarehousesResponseBodyWarehouseList {
	s.Mem = &v
	return s
}

func (s *ListWarehousesResponseBodyWarehouseList) SetName(v string) *ListWarehousesResponseBodyWarehouseList {
	s.Name = &v
	return s
}

func (s *ListWarehousesResponseBodyWarehouseList) SetNodeCount(v int64) *ListWarehousesResponseBodyWarehouseList {
	s.NodeCount = &v
	return s
}

func (s *ListWarehousesResponseBodyWarehouseList) SetStatus(v string) *ListWarehousesResponseBodyWarehouseList {
	s.Status = &v
	return s
}

type ListWarehousesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWarehousesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWarehousesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWarehousesResponse) GoString() string {
	return s.String()
}

func (s *ListWarehousesResponse) SetHeaders(v map[string]*string) *ListWarehousesResponse {
	s.Headers = v
	return s
}

func (s *ListWarehousesResponse) SetStatusCode(v int32) *ListWarehousesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWarehousesResponse) SetBody(v *ListWarehousesResponseBody) *ListWarehousesResponse {
	s.Body = v
	return s
}

type RebalanceHoloWarehouseRequest struct {
	// The name of the virtual warehouse.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_oss
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s RebalanceHoloWarehouseRequest) String() string {
	return tea.Prettify(s)
}

func (s RebalanceHoloWarehouseRequest) GoString() string {
	return s.String()
}

func (s *RebalanceHoloWarehouseRequest) SetName(v string) *RebalanceHoloWarehouseRequest {
	s.Name = &v
	return s
}

type RebalanceHoloWarehouseResponseBody struct {
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// C0EA5844-AB00-5653-8711-CD9FD1798412
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebalanceHoloWarehouseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RebalanceHoloWarehouseResponseBody) GoString() string {
	return s.String()
}

func (s *RebalanceHoloWarehouseResponseBody) SetData(v string) *RebalanceHoloWarehouseResponseBody {
	s.Data = &v
	return s
}

func (s *RebalanceHoloWarehouseResponseBody) SetRequestId(v string) *RebalanceHoloWarehouseResponseBody {
	s.RequestId = &v
	return s
}

type RebalanceHoloWarehouseResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebalanceHoloWarehouseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebalanceHoloWarehouseResponse) String() string {
	return tea.Prettify(s)
}

func (s RebalanceHoloWarehouseResponse) GoString() string {
	return s.String()
}

func (s *RebalanceHoloWarehouseResponse) SetHeaders(v map[string]*string) *RebalanceHoloWarehouseResponse {
	s.Headers = v
	return s
}

func (s *RebalanceHoloWarehouseResponse) SetStatusCode(v int32) *RebalanceHoloWarehouseResponse {
	s.StatusCode = &v
	return s
}

func (s *RebalanceHoloWarehouseResponse) SetBody(v *RebalanceHoloWarehouseResponseBody) *RebalanceHoloWarehouseResponse {
	s.Body = v
	return s
}

type RenameHoloWarehouseRequest struct {
	// The original name of the virtual warehouse.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_warehouse
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The new name of the virtual warehouse.
	//
	// This parameter is required.
	//
	// example:
	//
	// new_name
	NewWarehouseName *string `json:"newWarehouseName,omitempty" xml:"newWarehouseName,omitempty"`
}

func (s RenameHoloWarehouseRequest) String() string {
	return tea.Prettify(s)
}

func (s RenameHoloWarehouseRequest) GoString() string {
	return s.String()
}

func (s *RenameHoloWarehouseRequest) SetName(v string) *RenameHoloWarehouseRequest {
	s.Name = &v
	return s
}

func (s *RenameHoloWarehouseRequest) SetNewWarehouseName(v string) *RenameHoloWarehouseRequest {
	s.NewWarehouseName = &v
	return s
}

type RenameHoloWarehouseResponseBody struct {
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2A8DEF6E-067E-5DB0-BAE1-2894266E6C6A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenameHoloWarehouseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenameHoloWarehouseResponseBody) GoString() string {
	return s.String()
}

func (s *RenameHoloWarehouseResponseBody) SetData(v bool) *RenameHoloWarehouseResponseBody {
	s.Data = &v
	return s
}

func (s *RenameHoloWarehouseResponseBody) SetRequestId(v string) *RenameHoloWarehouseResponseBody {
	s.RequestId = &v
	return s
}

type RenameHoloWarehouseResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenameHoloWarehouseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenameHoloWarehouseResponse) String() string {
	return tea.Prettify(s)
}

func (s RenameHoloWarehouseResponse) GoString() string {
	return s.String()
}

func (s *RenameHoloWarehouseResponse) SetHeaders(v map[string]*string) *RenameHoloWarehouseResponse {
	s.Headers = v
	return s
}

func (s *RenameHoloWarehouseResponse) SetStatusCode(v int32) *RenameHoloWarehouseResponse {
	s.StatusCode = &v
	return s
}

func (s *RenameHoloWarehouseResponse) SetBody(v *RenameHoloWarehouseResponseBody) *RenameHoloWarehouseResponse {
	s.Body = v
	return s
}

type RenewInstanceRequest struct {
	// Specifies whether to enable monthly auto-renewal. The default value is false. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// >  If you enable auto-renewal for an instance for which auto-renewal is enabled, an error is reported.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
	// The renewal duration. Unit: month.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	Duration *int32 `json:"duration,omitempty" xml:"duration,omitempty"`
}

func (s RenewInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewInstanceRequest) SetAutoRenew(v bool) *RenewInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *RenewInstanceRequest) SetDuration(v int32) *RenewInstanceRequest {
	s.Duration = &v
	return s
}

type RenewInstanceResponseBody struct {
	// The returned data.
	Data *RenewInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned.
	//
	// example:
	//
	// null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// null
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D3AE84AB-0873-5FC7-A4C4-8CF869D2FA70
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request result, which indicates whether the request was successful.
	//
	// example:
	//
	// false
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RenewInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponseBody) SetData(v *RenewInstanceResponseBodyData) *RenewInstanceResponseBody {
	s.Data = v
	return s
}

func (s *RenewInstanceResponseBody) SetErrorCode(v string) *RenewInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RenewInstanceResponseBody) SetErrorMessage(v string) *RenewInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RenewInstanceResponseBody) SetHttpStatusCode(v string) *RenewInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RenewInstanceResponseBody) SetRequestId(v string) *RenewInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewInstanceResponseBody) SetSuccess(v string) *RenewInstanceResponseBody {
	s.Success = &v
	return s
}

type RenewInstanceResponseBodyData struct {
	// The error code returned.
	//
	// example:
	//
	// InvalidChargeType.UnRenewable
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error details.
	//
	// example:
	//
	// InvalidChargeType.UnRenewable
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 221625608580893
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Indicates whether the renewal was successful.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RenewInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponseBodyData) SetCode(v string) *RenewInstanceResponseBodyData {
	s.Code = &v
	return s
}

func (s *RenewInstanceResponseBodyData) SetMessage(v string) *RenewInstanceResponseBodyData {
	s.Message = &v
	return s
}

func (s *RenewInstanceResponseBodyData) SetOrderId(v string) *RenewInstanceResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *RenewInstanceResponseBodyData) SetSuccess(v bool) *RenewInstanceResponseBodyData {
	s.Success = &v
	return s
}

type RenewInstanceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceResponse) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponse) SetHeaders(v map[string]*string) *RenewInstanceResponse {
	s.Headers = v
	return s
}

func (s *RenewInstanceResponse) SetStatusCode(v int32) *RenewInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewInstanceResponse) SetBody(v *RenewInstanceResponseBody) *RenewInstanceResponse {
	s.Body = v
	return s
}

type RestartHoloWarehouseRequest struct {
	// The name of the virtual warehouse.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_warehouse
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s RestartHoloWarehouseRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartHoloWarehouseRequest) GoString() string {
	return s.String()
}

func (s *RestartHoloWarehouseRequest) SetName(v string) *RestartHoloWarehouseRequest {
	s.Name = &v
	return s
}

type RestartHoloWarehouseResponseBody struct {
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2A8DEF6E-067E-5DB0-BAE1-2894266E6C6A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartHoloWarehouseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartHoloWarehouseResponseBody) GoString() string {
	return s.String()
}

func (s *RestartHoloWarehouseResponseBody) SetData(v bool) *RestartHoloWarehouseResponseBody {
	s.Data = &v
	return s
}

func (s *RestartHoloWarehouseResponseBody) SetRequestId(v string) *RestartHoloWarehouseResponseBody {
	s.RequestId = &v
	return s
}

type RestartHoloWarehouseResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartHoloWarehouseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartHoloWarehouseResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartHoloWarehouseResponse) GoString() string {
	return s.String()
}

func (s *RestartHoloWarehouseResponse) SetHeaders(v map[string]*string) *RestartHoloWarehouseResponse {
	s.Headers = v
	return s
}

func (s *RestartHoloWarehouseResponse) SetStatusCode(v int32) *RestartHoloWarehouseResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartHoloWarehouseResponse) SetBody(v *RestartHoloWarehouseResponseBody) *RestartHoloWarehouseResponse {
	s.Body = v
	return s
}

type RestartInstanceResponseBody struct {
	// Indicates whether the operation was successful.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// 404
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Internal server error.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 36291497-CDB0-53DC-8CD7-762E054F57A6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request result, which indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RestartInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponseBody) SetData(v bool) *RestartInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *RestartInstanceResponseBody) SetErrorCode(v string) *RestartInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RestartInstanceResponseBody) SetErrorMessage(v string) *RestartInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RestartInstanceResponseBody) SetHttpStatusCode(v string) *RestartInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RestartInstanceResponseBody) SetRequestId(v string) *RestartInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartInstanceResponseBody) SetSuccess(v bool) *RestartInstanceResponseBody {
	s.Success = &v
	return s
}

type RestartInstanceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestartInstanceResponse) SetHeaders(v map[string]*string) *RestartInstanceResponse {
	s.Headers = v
	return s
}

func (s *RestartInstanceResponse) SetStatusCode(v int32) *RestartInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartInstanceResponse) SetBody(v *RestartInstanceResponseBody) *RestartInstanceResponse {
	s.Body = v
	return s
}

type ResumeHoloWarehouseRequest struct {
	// The name of the virtual warehouse.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_warehouse
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ResumeHoloWarehouseRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumeHoloWarehouseRequest) GoString() string {
	return s.String()
}

func (s *ResumeHoloWarehouseRequest) SetName(v string) *ResumeHoloWarehouseRequest {
	s.Name = &v
	return s
}

type ResumeHoloWarehouseResponseBody struct {
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2A8DEF6E-067E-5DB0-BAE1-2894266E6C6A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResumeHoloWarehouseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResumeHoloWarehouseResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeHoloWarehouseResponseBody) SetData(v bool) *ResumeHoloWarehouseResponseBody {
	s.Data = &v
	return s
}

func (s *ResumeHoloWarehouseResponseBody) SetRequestId(v string) *ResumeHoloWarehouseResponseBody {
	s.RequestId = &v
	return s
}

type ResumeHoloWarehouseResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeHoloWarehouseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeHoloWarehouseResponse) String() string {
	return tea.Prettify(s)
}

func (s ResumeHoloWarehouseResponse) GoString() string {
	return s.String()
}

func (s *ResumeHoloWarehouseResponse) SetHeaders(v map[string]*string) *ResumeHoloWarehouseResponse {
	s.Headers = v
	return s
}

func (s *ResumeHoloWarehouseResponse) SetStatusCode(v int32) *ResumeHoloWarehouseResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeHoloWarehouseResponse) SetBody(v *ResumeHoloWarehouseResponseBody) *ResumeHoloWarehouseResponse {
	s.Body = v
	return s
}

type ResumeInstanceResponseBody struct {
	// The returned result, which indicates whether the operation was successful.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// 404
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Internal server error.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D3AE84AB-0873-5FC7-A4C4-8CF869D2FA70
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request result, which indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ResumeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResumeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeInstanceResponseBody) SetData(v bool) *ResumeInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *ResumeInstanceResponseBody) SetErrorCode(v string) *ResumeInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ResumeInstanceResponseBody) SetErrorMessage(v string) *ResumeInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ResumeInstanceResponseBody) SetHttpStatusCode(v string) *ResumeInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ResumeInstanceResponseBody) SetRequestId(v string) *ResumeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeInstanceResponseBody) SetSuccess(v bool) *ResumeInstanceResponseBody {
	s.Success = &v
	return s
}

type ResumeInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ResumeInstanceResponse) GoString() string {
	return s.String()
}

func (s *ResumeInstanceResponse) SetHeaders(v map[string]*string) *ResumeInstanceResponse {
	s.Headers = v
	return s
}

func (s *ResumeInstanceResponse) SetStatusCode(v int32) *ResumeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeInstanceResponse) SetBody(v *ResumeInstanceResponseBody) *ResumeInstanceResponse {
	s.Body = v
	return s
}

type ScaleHoloWarehouseRequest struct {
	// The specifications of the virtual warehouse. The number of vCPUs must be an integer multiple of 16.
	//
	// This parameter is required.
	//
	// example:
	//
	// 64
	Cpu *int64 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// The name of the virtual warehouse.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_warehouse
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ScaleHoloWarehouseRequest) String() string {
	return tea.Prettify(s)
}

func (s ScaleHoloWarehouseRequest) GoString() string {
	return s.String()
}

func (s *ScaleHoloWarehouseRequest) SetCpu(v int64) *ScaleHoloWarehouseRequest {
	s.Cpu = &v
	return s
}

func (s *ScaleHoloWarehouseRequest) SetName(v string) *ScaleHoloWarehouseRequest {
	s.Name = &v
	return s
}

type ScaleHoloWarehouseResponseBody struct {
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2A8DEF6E-067E-5DB0-BAE1-2894266E6C6A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ScaleHoloWarehouseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScaleHoloWarehouseResponseBody) GoString() string {
	return s.String()
}

func (s *ScaleHoloWarehouseResponseBody) SetData(v bool) *ScaleHoloWarehouseResponseBody {
	s.Data = &v
	return s
}

func (s *ScaleHoloWarehouseResponseBody) SetRequestId(v string) *ScaleHoloWarehouseResponseBody {
	s.RequestId = &v
	return s
}

type ScaleHoloWarehouseResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScaleHoloWarehouseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScaleHoloWarehouseResponse) String() string {
	return tea.Prettify(s)
}

func (s ScaleHoloWarehouseResponse) GoString() string {
	return s.String()
}

func (s *ScaleHoloWarehouseResponse) SetHeaders(v map[string]*string) *ScaleHoloWarehouseResponse {
	s.Headers = v
	return s
}

func (s *ScaleHoloWarehouseResponse) SetStatusCode(v int32) *ScaleHoloWarehouseResponse {
	s.StatusCode = &v
	return s
}

func (s *ScaleHoloWarehouseResponse) SetBody(v *ScaleHoloWarehouseResponseBody) *ScaleHoloWarehouseResponse {
	s.Body = v
	return s
}

type ScaleInstanceRequest struct {
	// The infrequent access (IA) storage space of the instance. Unit: GB.
	//
	// > Ignore this parameter for pay-as-you-go instances.
	//
	// example:
	//
	// 1000G
	ColdStorageSize *int64 `json:"coldStorageSize,omitempty" xml:"coldStorageSize,omitempty"`
	// The specifications of the instance. Valid values:
	//
	// 	- 8-core 32GB (number of compute nodes: 1)
	//
	// 	- 16-core 64GB (number of compute nodes: 1)
	//
	// 	- 32-core 128GB (number of compute nodes: 2)
	//
	// 	- 64-core 256GB (number of compute nodes: 4)
	//
	// 	- 96-core 384GB (number of compute nodes: 6)
	//
	// 	- 128-core 512GB (number of compute nodes: 8)
	//
	// 	- Others
	//
	// >
	//
	// 	- Set this parameter to the number of cores.
	//
	// 	- If you want to set this parameter to specifications with more than 1,024 compute units (CUs), you must submit a ticket.
	//
	// 	- This parameter is invalid for Hologres Shared Cluster instances.
	//
	// 	- The specifications of 8-core 32GB (number of compute nodes: 1) are for trial use only and cannot be used for production.
	//
	// example:
	//
	// 128
	Cpu *int64 `json:"cpu,omitempty" xml:"cpu,omitempty"`
	// 是否开启ServerlessComputing
	//
	// example:
	//
	// true
	EnableServerlessComputing *bool `json:"enableServerlessComputing,omitempty" xml:"enableServerlessComputing,omitempty"`
	// The number of gateways. Valid values: 2 to 50.
	//
	// > This parameter is required only for virtual warehouse instances.
	//
	// example:
	//
	// 4
	GatewayCount *int64 `json:"gatewayCount,omitempty" xml:"gatewayCount,omitempty"`
	// The specification change type. Valid values:
	//
	// 	- UPGRADE
	//
	// 	- DOWNGRADE
	//
	// >
	//
	// 	- If you set this parameter to UPGRADE, the new specifications must be higher than the original specifications. You must configure at least one of the cpu, storageSize, and coldStorageSize parameters. If you leave a parameter empty, the related configuration remains unchanged.
	//
	// 	- If you set this parameter to DOWNGRADE, the new specifications must be lower than the original specifications. You must configure at least one of the cpu, storageSize, and coldStorageSize parameters. If you leave a parameter empty, the related configuration remains unchanged.
	//
	// This parameter is required.
	//
	// example:
	//
	// UPGRADE
	ScaleType *string `json:"scaleType,omitempty" xml:"scaleType,omitempty"`
	// The standard storage space of the instance. Unit: GB.
	//
	// > Ignore this parameter for pay-as-you-go instances.
	//
	// example:
	//
	// 1000G
	StorageSize *int64 `json:"storageSize,omitempty" xml:"storageSize,omitempty"`
}

func (s ScaleInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ScaleInstanceRequest) GoString() string {
	return s.String()
}

func (s *ScaleInstanceRequest) SetColdStorageSize(v int64) *ScaleInstanceRequest {
	s.ColdStorageSize = &v
	return s
}

func (s *ScaleInstanceRequest) SetCpu(v int64) *ScaleInstanceRequest {
	s.Cpu = &v
	return s
}

func (s *ScaleInstanceRequest) SetEnableServerlessComputing(v bool) *ScaleInstanceRequest {
	s.EnableServerlessComputing = &v
	return s
}

func (s *ScaleInstanceRequest) SetGatewayCount(v int64) *ScaleInstanceRequest {
	s.GatewayCount = &v
	return s
}

func (s *ScaleInstanceRequest) SetScaleType(v string) *ScaleInstanceRequest {
	s.ScaleType = &v
	return s
}

func (s *ScaleInstanceRequest) SetStorageSize(v int64) *ScaleInstanceRequest {
	s.StorageSize = &v
	return s
}

type ScaleInstanceResponseBody struct {
	// The returned data.
	Data *ScaleInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned.
	//
	// example:
	//
	// null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// null
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D3AE84AB-0873-5FC7-A4C4-8CF869D2FA70
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ScaleInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScaleInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ScaleInstanceResponseBody) SetData(v *ScaleInstanceResponseBodyData) *ScaleInstanceResponseBody {
	s.Data = v
	return s
}

func (s *ScaleInstanceResponseBody) SetErrorCode(v string) *ScaleInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ScaleInstanceResponseBody) SetErrorMessage(v string) *ScaleInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ScaleInstanceResponseBody) SetHttpStatusCode(v string) *ScaleInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ScaleInstanceResponseBody) SetRequestId(v string) *ScaleInstanceResponseBody {
	s.RequestId = &v
	return s
}

type ScaleInstanceResponseBodyData struct {
	// The error code returned.
	//
	// example:
	//
	// InvalidScaleType.Unsupported
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error details.
	//
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 219183853450000
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Indicates whether the change to specifications was successful.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ScaleInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ScaleInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ScaleInstanceResponseBodyData) SetCode(v string) *ScaleInstanceResponseBodyData {
	s.Code = &v
	return s
}

func (s *ScaleInstanceResponseBodyData) SetMessage(v string) *ScaleInstanceResponseBodyData {
	s.Message = &v
	return s
}

func (s *ScaleInstanceResponseBodyData) SetOrderId(v string) *ScaleInstanceResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *ScaleInstanceResponseBodyData) SetSuccess(v bool) *ScaleInstanceResponseBodyData {
	s.Success = &v
	return s
}

type ScaleInstanceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScaleInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScaleInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ScaleInstanceResponse) GoString() string {
	return s.String()
}

func (s *ScaleInstanceResponse) SetHeaders(v map[string]*string) *ScaleInstanceResponse {
	s.Headers = v
	return s
}

func (s *ScaleInstanceResponse) SetStatusCode(v int32) *ScaleInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ScaleInstanceResponse) SetBody(v *ScaleInstanceResponseBody) *ScaleInstanceResponse {
	s.Body = v
	return s
}

type StopInstanceResponseBody struct {
	// The returned result, which indicates whether the operation was successful.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// 404
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Internal server error.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2A8DEF6E-067E-5DB0-BAE1-2894266E6C6A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request result, which indicates whether the request was successful.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopInstanceResponseBody) SetData(v bool) *StopInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *StopInstanceResponseBody) SetErrorCode(v string) *StopInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StopInstanceResponseBody) SetErrorMessage(v string) *StopInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StopInstanceResponseBody) SetHttpStatusCode(v string) *StopInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StopInstanceResponseBody) SetRequestId(v string) *StopInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopInstanceResponseBody) SetSuccess(v bool) *StopInstanceResponseBody {
	s.Success = &v
	return s
}

type StopInstanceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopInstanceResponse) SetHeaders(v map[string]*string) *StopInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopInstanceResponse) SetStatusCode(v int32) *StopInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopInstanceResponse) SetBody(v *StopInstanceResponseBody) *StopInstanceResponse {
	s.Body = v
	return s
}

type SuspendHoloWarehouseRequest struct {
	// The name of the virtual warehouse.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_warehouse
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s SuspendHoloWarehouseRequest) String() string {
	return tea.Prettify(s)
}

func (s SuspendHoloWarehouseRequest) GoString() string {
	return s.String()
}

func (s *SuspendHoloWarehouseRequest) SetName(v string) *SuspendHoloWarehouseRequest {
	s.Name = &v
	return s
}

type SuspendHoloWarehouseResponseBody struct {
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2A8DEF6E-067E-5DB0-BAE1-2894266E6C6A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SuspendHoloWarehouseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SuspendHoloWarehouseResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendHoloWarehouseResponseBody) SetData(v string) *SuspendHoloWarehouseResponseBody {
	s.Data = &v
	return s
}

func (s *SuspendHoloWarehouseResponseBody) SetRequestId(v string) *SuspendHoloWarehouseResponseBody {
	s.RequestId = &v
	return s
}

type SuspendHoloWarehouseResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SuspendHoloWarehouseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SuspendHoloWarehouseResponse) String() string {
	return tea.Prettify(s)
}

func (s SuspendHoloWarehouseResponse) GoString() string {
	return s.String()
}

func (s *SuspendHoloWarehouseResponse) SetHeaders(v map[string]*string) *SuspendHoloWarehouseResponse {
	s.Headers = v
	return s
}

func (s *SuspendHoloWarehouseResponse) SetStatusCode(v int32) *SuspendHoloWarehouseResponse {
	s.StatusCode = &v
	return s
}

func (s *SuspendHoloWarehouseResponse) SetBody(v *SuspendHoloWarehouseResponseBody) *SuspendHoloWarehouseResponse {
	s.Body = v
	return s
}

type UpdateInstanceNameRequest struct {
	// The new name of the instance. The name must be 2 to 64 characters in length.
	//
	// example:
	//
	// new_name
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
}

func (s UpdateInstanceNameRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNameRequest) SetInstanceName(v string) *UpdateInstanceNameRequest {
	s.InstanceName = &v
	return s
}

type UpdateInstanceNameResponseBody struct {
	// The returned result, which indicates whether the operation was successful.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// 404
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Internal server error.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C6B55032-D41A-5FE0-9C07-8BD81C88422E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request result, which indicates whether the request was successful.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateInstanceNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNameResponseBody) SetData(v bool) *UpdateInstanceNameResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetErrorCode(v string) *UpdateInstanceNameResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetErrorMessage(v string) *UpdateInstanceNameResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetHttpStatusCode(v string) *UpdateInstanceNameResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetRequestId(v string) *UpdateInstanceNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceNameResponseBody) SetSuccess(v bool) *UpdateInstanceNameResponseBody {
	s.Success = &v
	return s
}

type UpdateInstanceNameResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceNameResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNameResponse) SetHeaders(v map[string]*string) *UpdateInstanceNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceNameResponse) SetStatusCode(v int32) *UpdateInstanceNameResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceNameResponse) SetBody(v *UpdateInstanceNameResponseBody) *UpdateInstanceNameResponse {
	s.Body = v
	return s
}

type UpdateInstanceNetworkTypeRequest struct {
	// Specifies whether to change the network type from AnyTunnel to SingleTunnel. This parameter is invalid for new instances. For new instances, this parameter is set to null by default.
	//
	// Valid values:
	//
	// 	- others/null: The network type is not changed from AnyTunnel to SingleTunnel.
	//
	// 	- true: The network type is changed from AnyTunnel to SingleTunnel.
	//
	// example:
	//
	// true
	AnyTunnelToSingleTunnel *string `json:"anyTunnelToSingleTunnel,omitempty" xml:"anyTunnelToSingleTunnel,omitempty"`
	// A list of network types that you want to enable. The network types are randomly ordered in the list. For example, the Internet, Intranet, and VPCSingleTunnel network types are enabled. If you want to disable the Internet type, set this parameter to Intranet,VPCSingleTunnel.
	//
	// Valid values:
	//
	// 	- VPCSingleTunnel: virtual private cloud (VPC).
	//
	// 	- Intranet: internal network.
	//
	// 	- VPCAnyTunnel: compatibility requirements. This value is not supported by new instances.
	//
	// 	- Internet: Internet.
	//
	// example:
	//
	// Internet,VPCSingleTunnel
	NetworkTypes *string `json:"networkTypes,omitempty" xml:"networkTypes,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-2vccsiymtqr9aavew0vo3
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-t4netc3y5etlondfb5ra7
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// The ID of the Alibaba Cloud account to which the VPC belongs.
	//
	// example:
	//
	// 1999365732646672
	VpcOwnerId *string `json:"vpcOwnerId,omitempty" xml:"vpcOwnerId,omitempty"`
	// The region in which the VPC resides.
	//
	// example:
	//
	// cn-hangzhou
	VpcRegionId *string `json:"vpcRegionId,omitempty" xml:"vpcRegionId,omitempty"`
}

func (s UpdateInstanceNetworkTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceNetworkTypeRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNetworkTypeRequest) SetAnyTunnelToSingleTunnel(v string) *UpdateInstanceNetworkTypeRequest {
	s.AnyTunnelToSingleTunnel = &v
	return s
}

func (s *UpdateInstanceNetworkTypeRequest) SetNetworkTypes(v string) *UpdateInstanceNetworkTypeRequest {
	s.NetworkTypes = &v
	return s
}

func (s *UpdateInstanceNetworkTypeRequest) SetVSwitchId(v string) *UpdateInstanceNetworkTypeRequest {
	s.VSwitchId = &v
	return s
}

func (s *UpdateInstanceNetworkTypeRequest) SetVpcId(v string) *UpdateInstanceNetworkTypeRequest {
	s.VpcId = &v
	return s
}

func (s *UpdateInstanceNetworkTypeRequest) SetVpcOwnerId(v string) *UpdateInstanceNetworkTypeRequest {
	s.VpcOwnerId = &v
	return s
}

func (s *UpdateInstanceNetworkTypeRequest) SetVpcRegionId(v string) *UpdateInstanceNetworkTypeRequest {
	s.VpcRegionId = &v
	return s
}

type UpdateInstanceNetworkTypeResponseBody struct {
	// The returned result, which indicates whether the operation was successful.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// 404
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Internal server error.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9CC37B9F-F4B4-5FF1-939B-AEE78DC70130
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateInstanceNetworkTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceNetworkTypeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNetworkTypeResponseBody) SetData(v string) *UpdateInstanceNetworkTypeResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateInstanceNetworkTypeResponseBody) SetErrorCode(v string) *UpdateInstanceNetworkTypeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateInstanceNetworkTypeResponseBody) SetErrorMessage(v string) *UpdateInstanceNetworkTypeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateInstanceNetworkTypeResponseBody) SetHttpStatusCode(v string) *UpdateInstanceNetworkTypeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateInstanceNetworkTypeResponseBody) SetRequestId(v string) *UpdateInstanceNetworkTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceNetworkTypeResponseBody) SetSuccess(v bool) *UpdateInstanceNetworkTypeResponseBody {
	s.Success = &v
	return s
}

type UpdateInstanceNetworkTypeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceNetworkTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceNetworkTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceNetworkTypeResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNetworkTypeResponse) SetHeaders(v map[string]*string) *UpdateInstanceNetworkTypeResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceNetworkTypeResponse) SetStatusCode(v int32) *UpdateInstanceNetworkTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceNetworkTypeResponse) SetBody(v *UpdateInstanceNetworkTypeResponseBody) *UpdateInstanceNetworkTypeResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("hologram"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Updates a resource group.
//
// @param request - ChangeResourceGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NewResourceGroupId)) {
		body["newResourceGroupId"] = request.NewResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeResourceGroup"),
		Version:     tea.String("2022-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/tag/changeResourceGroup"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ChangeResourceGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ChangeResourceGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Updates a resource group.
//
// @param request - ChangeResourceGroupRequest
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroup(request *ChangeResourceGroupRequest) (_result *ChangeResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.ChangeResourceGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a virtual warehouse.
//
// @param request - CreateHoloWarehouseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHoloWarehouseResponse
func (client *Client) CreateHoloWarehouseWithOptions(instanceId *string, request *CreateHoloWarehouseRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateHoloWarehouseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cpu)) {
		body["cpu"] = request.Cpu
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateHoloWarehouse"),
		Version:     tea.String("2022-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/createHoloWarehouse"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateHoloWarehouseResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateHoloWarehouseResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Creates a virtual warehouse.
//
// @param request - CreateHoloWarehouseRequest
//
// @return CreateHoloWarehouseResponse
func (client *Client) CreateHoloWarehouse(instanceId *string, request *CreateHoloWarehouseRequest) (_result *CreateHoloWarehouseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateHoloWarehouseResponse{}
	_body, _err := client.CreateHoloWarehouseWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Hologres instance.
//
// Description:
//
// > Before you call this operation, make sure that you understand the billing method and pricing of Hologres because this operation is charged.
//
// 	- For more information about the billing details of Hologres, see [Pricing](https://www.alibabacloud.com/help/en/hologres/developer-reference/api-hologram-2022-06-01-createinstance).
//
// 	- When you purchase a Hologres instance, you must specify the region and zone in which the Hologres instance resides. A region may correspond to multiple zones. Example:
//
// <!---->
//
//     cn-hangzhou: cn-hangzhou-h, cn-hangzhou-j
//
//        cn-shanghai: cn-shanghai-e, cn-shanghai-f
//
//        cn-beijing: cn-beijing-i, cn-beijing-g
//
//        cn-zhangjiakou: cn-zhangjiakou-b
//
//        cn-shenzhen: cn-shenzhen-e
//
//        cn-hongkong: cn-hongkong-b
//
//        cn-shanghai-finance-1: cn-shanghai-finance-1z
//
//        ap-northeast-1: ap-northeast-1a
//
//        ap-southeast-1: ap-southeast-1c
//
//        ap-southeast-3: ap-southeast-3b
//
//        ap-southeast-5: ap-southeast-5b
//
//        ap-south-1: ap-south-1b
//
//        eu-central-1: eu-central-1a
//
//        us-east-1: us-east-1a
//
//        us-west-1: us-west-1b
//
// @param request - CreateInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResponse
func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		body["autoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		body["autoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		body["chargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.ColdStorageSize)) {
		body["coldStorageSize"] = request.ColdStorageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Cpu)) {
		body["cpu"] = request.Cpu
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		body["duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.EnableServerlessComputing)) {
		body["enableServerlessComputing"] = request.EnableServerlessComputing
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayCount)) {
		body["gatewayCount"] = request.GatewayCount
	}

	if !tea.BoolValue(util.IsUnset(request.InitialDatabases)) {
		body["initialDatabases"] = request.InitialDatabases
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["instanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		body["instanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.LeaderInstanceId)) {
		body["leaderInstanceId"] = request.LeaderInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PricingCycle)) {
		body["pricingCycle"] = request.PricingCycle
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["regionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StorageSize)) {
		body["storageSize"] = request.StorageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StorageType)) {
		body["storageType"] = request.StorageType
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		body["vSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["vpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		body["zoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstance"),
		Version:     tea.String("2022-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Creates a Hologres instance.
//
// Description:
//
// > Before you call this operation, make sure that you understand the billing method and pricing of Hologres because this operation is charged.
//
// 	- For more information about the billing details of Hologres, see [Pricing](https://www.alibabacloud.com/help/en/hologres/developer-reference/api-hologram-2022-06-01-createinstance).
//
// 	- When you purchase a Hologres instance, you must specify the region and zone in which the Hologres instance resides. A region may correspond to multiple zones. Example:
//
// <!---->
//
//     cn-hangzhou: cn-hangzhou-h, cn-hangzhou-j
//
//        cn-shanghai: cn-shanghai-e, cn-shanghai-f
//
//        cn-beijing: cn-beijing-i, cn-beijing-g
//
//        cn-zhangjiakou: cn-zhangjiakou-b
//
//        cn-shenzhen: cn-shenzhen-e
//
//        cn-hongkong: cn-hongkong-b
//
//        cn-shanghai-finance-1: cn-shanghai-finance-1z
//
//        ap-northeast-1: ap-northeast-1a
//
//        ap-southeast-1: ap-southeast-1c
//
//        ap-southeast-3: ap-southeast-3b
//
//        ap-southeast-5: ap-southeast-5b
//
//        ap-south-1: ap-south-1b
//
//        eu-central-1: eu-central-1a
//
//        us-east-1: us-east-1a
//
//        us-west-1: us-west-1b
//
// @param request - CreateInstanceRequest
//
// @return CreateInstanceResponse
func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a virtual warehouse.
//
// @param request - DeleteHoloWarehouseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHoloWarehouseResponse
func (client *Client) DeleteHoloWarehouseWithOptions(instanceId *string, request *DeleteHoloWarehouseRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteHoloWarehouseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteHoloWarehouse"),
		Version:     tea.String("2022-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/deleteHoloWarehouse"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteHoloWarehouseResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteHoloWarehouseResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Deletes a virtual warehouse.
//
// @param request - DeleteHoloWarehouseRequest
//
// @return DeleteHoloWarehouseResponse
func (client *Client) DeleteHoloWarehouse(instanceId *string, request *DeleteHoloWarehouseRequest) (_result *DeleteHoloWarehouseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteHoloWarehouseResponse{}
	_body, _err := client.DeleteHoloWarehouseWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a Hologres instance.
//
// Description:
//
// > Before you call this operation, read the documentation and make sure that you understand the prerequisites and impacts of this operation.
//
// 	- After you delete a Hologres instance, data and objects in the instance cannot be restored. Proceed with caution. For more information, see [Billing overview](https://www.alibabacloud.com/help/zh/hologres/product-overview/billing-overview?spm=a2c63.p38356.0.0.efc33b87i5pDl7).
//
// 	- You can delete only pay-as-you-go instances.
//
// @param request - DeleteInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithOptions(instanceId *string, request *DeleteInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstance"),
		Version:     tea.String("2022-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Deletes a Hologres instance.
//
// Description:
//
// > Before you call this operation, read the documentation and make sure that you understand the prerequisites and impacts of this operation.
//
// 	- After you delete a Hologres instance, data and objects in the instance cannot be restored. Proceed with caution. For more information, see [Billing overview](https://www.alibabacloud.com/help/zh/hologres/product-overview/billing-overview?spm=a2c63.p38356.0.0.efc33b87i5pDl7).
//
// 	- You can delete only pay-as-you-go instances.
//
// @param request - DeleteInstanceRequest
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstance(instanceId *string, request *DeleteInstanceRequest) (_result *DeleteInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables data lake acceleration.
//
// @param request - DisableHiveAccessRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableHiveAccessResponse
func (client *Client) DisableHiveAccessWithOptions(instanceId *string, request *DisableHiveAccessRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DisableHiveAccessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableHiveAccess"),
		Version:     tea.String("2022-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/disableHiveAccess"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DisableHiveAccessResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DisableHiveAccessResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Disables data lake acceleration.
//
// @param request - DisableHiveAccessRequest
//
// @return DisableHiveAccessResponse
func (client *Client) DisableHiveAccess(instanceId *string, request *DisableHiveAccessRequest) (_result *DisableHiveAccessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DisableHiveAccessResponse{}
	_body, _err := client.DisableHiveAccessWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables data lake acceleration.
//
// @param request - EnableHiveAccessRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableHiveAccessResponse
func (client *Client) EnableHiveAccessWithOptions(instanceId *string, request *EnableHiveAccessRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *EnableHiveAccessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableHiveAccess"),
		Version:     tea.String("2022-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/enableHiveAccess"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &EnableHiveAccessResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &EnableHiveAccessResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Enables data lake acceleration.
//
// @param request - EnableHiveAccessRequest
//
// @return EnableHiveAccessResponse
func (client *Client) EnableHiveAccess(instanceId *string, request *EnableHiveAccessRequest) (_result *EnableHiveAccessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EnableHiveAccessResponse{}
	_body, _err := client.EnableHiveAccessWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the details of an instance.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceResponse
func (client *Client) GetInstanceWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstance"),
		Version:     tea.String("2022-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Obtains the details of an instance.
//
// @return GetInstanceResponse
func (client *Client) GetInstance(instanceId *string) (_result *GetInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceResponse{}
	_body, _err := client.GetInstanceWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries details of a virtual warehouse instance.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWarehouseDetailResponse
func (client *Client) GetWarehouseDetailWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetWarehouseDetailResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetWarehouseDetail"),
		Version:     tea.String("2022-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/getWarehouseDetail"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetWarehouseDetailResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetWarehouseDetailResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries details of a virtual warehouse instance.
//
// @return GetWarehouseDetailResponse
func (client *Client) GetWarehouseDetail(instanceId *string) (_result *GetWarehouseDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetWarehouseDetailResponse{}
	_body, _err := client.GetWarehouseDetailWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of backups. A backup is a full data snapshot of an instance at the end of the snapshot time. You can purchase another instance to completely restore the original data.
//
// @param request - ListBackupDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBackupDataResponse
func (client *Client) ListBackupDataWithOptions(request *ListBackupDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListBackupDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupType)) {
		query["backupType"] = request.BackupType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["instanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBackupData"),
		Version:     tea.String("2022-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/backups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListBackupDataResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListBackupDataResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries a list of backups. A backup is a full data snapshot of an instance at the end of the snapshot time. You can purchase another instance to completely restore the original data.
//
// @param request - ListBackupDataRequest
//
// @return ListBackupDataResponse
func (client *Client) ListBackupData(request *ListBackupDataRequest) (_result *ListBackupDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListBackupDataResponse{}
	_body, _err := client.ListBackupDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of instances.
//
// @param request - ListInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CmsInstanceType)) {
		body["cmsInstanceType"] = request.CmsInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		body["tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstances"),
		Version:     tea.String("2022-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListInstancesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListInstancesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries a list of instances.
//
// @param request - ListInstancesRequest
//
// @return ListInstancesResponse
func (client *Client) ListInstances(request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of virtual warehouse instances.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWarehousesResponse
func (client *Client) ListWarehousesWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListWarehousesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListWarehouses"),
		Version:     tea.String("2022-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/listWarehouses"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListWarehousesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListWarehousesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the list of virtual warehouse instances.
//
// @return ListWarehousesResponse
func (client *Client) ListWarehouses(instanceId *string) (_result *ListWarehousesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListWarehousesResponse{}
	_body, _err := client.ListWarehousesWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Triggers shard rebalancing for a virtual warehouse.
//
// @param request - RebalanceHoloWarehouseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebalanceHoloWarehouseResponse
func (client *Client) RebalanceHoloWarehouseWithOptions(instanceId *string, request *RebalanceHoloWarehouseRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RebalanceHoloWarehouseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RebalanceHoloWarehouse"),
		Version:     tea.String("2022-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/rebalanceHoloWarehouse"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RebalanceHoloWarehouseResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RebalanceHoloWarehouseResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Triggers shard rebalancing for a virtual warehouse.
//
// @param request - RebalanceHoloWarehouseRequest
//
// @return RebalanceHoloWarehouseResponse
func (client *Client) RebalanceHoloWarehouse(instanceId *string, request *RebalanceHoloWarehouseRequest) (_result *RebalanceHoloWarehouseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RebalanceHoloWarehouseResponse{}
	_body, _err := client.RebalanceHoloWarehouseWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Renames a virtual warehouse.
//
// @param request - RenameHoloWarehouseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenameHoloWarehouseResponse
func (client *Client) RenameHoloWarehouseWithOptions(instanceId *string, request *RenameHoloWarehouseRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RenameHoloWarehouseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NewWarehouseName)) {
		body["newWarehouseName"] = request.NewWarehouseName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RenameHoloWarehouse"),
		Version:     tea.String("2022-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/renameHoloWarehouse"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RenameHoloWarehouseResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RenameHoloWarehouseResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Renames a virtual warehouse.
//
// @param request - RenameHoloWarehouseRequest
//
// @return RenameHoloWarehouseResponse
func (client *Client) RenameHoloWarehouse(instanceId *string, request *RenameHoloWarehouseRequest) (_result *RenameHoloWarehouseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RenameHoloWarehouseResponse{}
	_body, _err := client.RenameHoloWarehouseWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Manually renews a Hologres instance. You can enable monthly auto-renewal when you renew a Hologres instance.
//
// Description:
//
// >  Before you call this operation, make sure that you understand the billing method and pricing of Hologres because this operation is charged.
//
// 	- For more information about the billing of Hologres, see [Billing overview](https://www.alibabacloud.com/help/zh/hologres/product-overview/billing-overview).
//
// 	- For more information about how to renew a Hologres instance, see [Manage renewals](https://www.alibabacloud.com/help/zh/hologres/product-overview/manage-renewals?spm=a2c63.p38356.0.0.38e731c9VAwtDP).
//
// 	- You can renew only subscription instances.
//
// @param request - RenewInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewInstanceResponse
func (client *Client) RenewInstanceWithOptions(instanceId *string, request *RenewInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RenewInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		body["autoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		body["duration"] = request.Duration
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RenewInstance"),
		Version:     tea.String("2022-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/renew"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RenewInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RenewInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Manually renews a Hologres instance. You can enable monthly auto-renewal when you renew a Hologres instance.
//
// Description:
//
// >  Before you call this operation, make sure that you understand the billing method and pricing of Hologres because this operation is charged.
//
// 	- For more information about the billing of Hologres, see [Billing overview](https://www.alibabacloud.com/help/zh/hologres/product-overview/billing-overview).
//
// 	- For more information about how to renew a Hologres instance, see [Manage renewals](https://www.alibabacloud.com/help/zh/hologres/product-overview/manage-renewals?spm=a2c63.p38356.0.0.38e731c9VAwtDP).
//
// 	- You can renew only subscription instances.
//
// @param request - RenewInstanceRequest
//
// @return RenewInstanceResponse
func (client *Client) RenewInstance(instanceId *string, request *RenewInstanceRequest) (_result *RenewInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RenewInstanceResponse{}
	_body, _err := client.RenewInstanceWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restarts a virtual warehouse.
//
// @param request - RestartHoloWarehouseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartHoloWarehouseResponse
func (client *Client) RestartHoloWarehouseWithOptions(instanceId *string, request *RestartHoloWarehouseRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RestartHoloWarehouseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RestartHoloWarehouse"),
		Version:     tea.String("2022-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/restartHoloWarehouse"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RestartHoloWarehouseResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RestartHoloWarehouseResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Restarts a virtual warehouse.
//
// @param request - RestartHoloWarehouseRequest
//
// @return RestartHoloWarehouseResponse
func (client *Client) RestartHoloWarehouse(instanceId *string, request *RestartHoloWarehouseRequest) (_result *RestartHoloWarehouseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RestartHoloWarehouseResponse{}
	_body, _err := client.RestartHoloWarehouseWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restarts an instance.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartInstanceResponse
func (client *Client) RestartInstanceWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RestartInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RestartInstance"),
		Version:     tea.String("2022-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/restart"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RestartInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RestartInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Restarts an instance.
//
// @return RestartInstanceResponse
func (client *Client) RestartInstance(instanceId *string) (_result *RestartInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RestartInstanceResponse{}
	_body, _err := client.RestartInstanceWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resumes a virtual warehouse.
//
// @param request - ResumeHoloWarehouseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeHoloWarehouseResponse
func (client *Client) ResumeHoloWarehouseWithOptions(instanceId *string, request *ResumeHoloWarehouseRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ResumeHoloWarehouseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ResumeHoloWarehouse"),
		Version:     tea.String("2022-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/resumeHoloWarehouse"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ResumeHoloWarehouseResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ResumeHoloWarehouseResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Resumes a virtual warehouse.
//
// @param request - ResumeHoloWarehouseRequest
//
// @return ResumeHoloWarehouseResponse
func (client *Client) ResumeHoloWarehouse(instanceId *string, request *ResumeHoloWarehouseRequest) (_result *ResumeHoloWarehouseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ResumeHoloWarehouseResponse{}
	_body, _err := client.ResumeHoloWarehouseWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resumes an instance.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeInstanceResponse
func (client *Client) ResumeInstanceWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ResumeInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ResumeInstance"),
		Version:     tea.String("2022-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/resume"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ResumeInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ResumeInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Resumes an instance.
//
// @return ResumeInstanceResponse
func (client *Client) ResumeInstance(instanceId *string) (_result *ResumeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ResumeInstanceResponse{}
	_body, _err := client.ResumeInstanceWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Scales in or out a virtual warehouse.
//
// @param request - ScaleHoloWarehouseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ScaleHoloWarehouseResponse
func (client *Client) ScaleHoloWarehouseWithOptions(instanceId *string, request *ScaleHoloWarehouseRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ScaleHoloWarehouseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cpu)) {
		body["cpu"] = request.Cpu
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ScaleHoloWarehouse"),
		Version:     tea.String("2022-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/scaleHoloWarehouse"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ScaleHoloWarehouseResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ScaleHoloWarehouseResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Scales in or out a virtual warehouse.
//
// @param request - ScaleHoloWarehouseRequest
//
// @return ScaleHoloWarehouseResponse
func (client *Client) ScaleHoloWarehouse(instanceId *string, request *ScaleHoloWarehouseRequest) (_result *ScaleHoloWarehouseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ScaleHoloWarehouseResponse{}
	_body, _err := client.ScaleHoloWarehouseWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the specifications and storage space of a Hologres instance.
//
// Description:
//
// > Before you call this operation, make sure that you understand the billing method and pricing of Hologres because this operation is charged.
//
// 	- For more information about the billing of Hologres, see [Billing overview](https://www.alibabacloud.com/help/zh/hologres/product-overview/billing-overview).
//
// 	- During the change of computing resource specifications of a Hologres instance, the instance is unavailable. During the change of storage resource specifications of a Hologres instance, the instance can work normally. Do not frequently change instance specifications. For more information, see [Upgrade or downgrade instance specifications](https://www.alibabacloud.com/help/en/hologres/product-overview/upgrade-or-downgrade-instance-specifications).
//
// @param request - ScaleInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ScaleInstanceResponse
func (client *Client) ScaleInstanceWithOptions(instanceId *string, request *ScaleInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ScaleInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ColdStorageSize)) {
		body["coldStorageSize"] = request.ColdStorageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Cpu)) {
		body["cpu"] = request.Cpu
	}

	if !tea.BoolValue(util.IsUnset(request.EnableServerlessComputing)) {
		body["enableServerlessComputing"] = request.EnableServerlessComputing
	}

	if !tea.BoolValue(util.IsUnset(request.GatewayCount)) {
		body["gatewayCount"] = request.GatewayCount
	}

	if !tea.BoolValue(util.IsUnset(request.ScaleType)) {
		body["scaleType"] = request.ScaleType
	}

	if !tea.BoolValue(util.IsUnset(request.StorageSize)) {
		body["storageSize"] = request.StorageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ScaleInstance"),
		Version:     tea.String("2022-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/scale"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ScaleInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ScaleInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Changes the specifications and storage space of a Hologres instance.
//
// Description:
//
// > Before you call this operation, make sure that you understand the billing method and pricing of Hologres because this operation is charged.
//
// 	- For more information about the billing of Hologres, see [Billing overview](https://www.alibabacloud.com/help/zh/hologres/product-overview/billing-overview).
//
// 	- During the change of computing resource specifications of a Hologres instance, the instance is unavailable. During the change of storage resource specifications of a Hologres instance, the instance can work normally. Do not frequently change instance specifications. For more information, see [Upgrade or downgrade instance specifications](https://www.alibabacloud.com/help/en/hologres/product-overview/upgrade-or-downgrade-instance-specifications).
//
// @param request - ScaleInstanceRequest
//
// @return ScaleInstanceResponse
func (client *Client) ScaleInstance(instanceId *string, request *ScaleInstanceRequest) (_result *ScaleInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ScaleInstanceResponse{}
	_body, _err := client.ScaleInstanceWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops an instance.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopInstanceResponse
func (client *Client) StopInstanceWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StopInstance"),
		Version:     tea.String("2022-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/stop"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &StopInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &StopInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Stops an instance.
//
// @return StopInstanceResponse
func (client *Client) StopInstance(instanceId *string) (_result *StopInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopInstanceResponse{}
	_body, _err := client.StopInstanceWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Suspends a virtual warehouse.
//
// @param request - SuspendHoloWarehouseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuspendHoloWarehouseResponse
func (client *Client) SuspendHoloWarehouseWithOptions(instanceId *string, request *SuspendHoloWarehouseRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SuspendHoloWarehouseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SuspendHoloWarehouse"),
		Version:     tea.String("2022-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/suspendHoloWarehouse"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SuspendHoloWarehouseResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SuspendHoloWarehouseResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Suspends a virtual warehouse.
//
// @param request - SuspendHoloWarehouseRequest
//
// @return SuspendHoloWarehouseResponse
func (client *Client) SuspendHoloWarehouse(instanceId *string, request *SuspendHoloWarehouseRequest) (_result *SuspendHoloWarehouseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SuspendHoloWarehouseResponse{}
	_body, _err := client.SuspendHoloWarehouseWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the name of an instance.
//
// @param request - UpdateInstanceNameRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceNameResponse
func (client *Client) UpdateInstanceNameWithOptions(instanceId *string, request *UpdateInstanceNameRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateInstanceNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["instanceName"] = request.InstanceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstanceName"),
		Version:     tea.String("2022-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/instanceName"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateInstanceNameResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateInstanceNameResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Changes the name of an instance.
//
// @param request - UpdateInstanceNameRequest
//
// @return UpdateInstanceNameResponse
func (client *Client) UpdateInstanceName(instanceId *string, request *UpdateInstanceNameRequest) (_result *UpdateInstanceNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstanceNameResponse{}
	_body, _err := client.UpdateInstanceNameWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the network configuration of an instance.
//
// @param request - UpdateInstanceNetworkTypeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceNetworkTypeResponse
func (client *Client) UpdateInstanceNetworkTypeWithOptions(instanceId *string, request *UpdateInstanceNetworkTypeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateInstanceNetworkTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnyTunnelToSingleTunnel)) {
		body["anyTunnelToSingleTunnel"] = request.AnyTunnelToSingleTunnel
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkTypes)) {
		body["networkTypes"] = request.NetworkTypes
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		body["vSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["vpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcOwnerId)) {
		body["vpcOwnerId"] = request.VpcOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcRegionId)) {
		body["vpcRegionId"] = request.VpcRegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstanceNetworkType"),
		Version:     tea.String("2022-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/network"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateInstanceNetworkTypeResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateInstanceNetworkTypeResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Modifies the network configuration of an instance.
//
// @param request - UpdateInstanceNetworkTypeRequest
//
// @return UpdateInstanceNetworkTypeResponse
func (client *Client) UpdateInstanceNetworkType(instanceId *string, request *UpdateInstanceNetworkTypeRequest) (_result *UpdateInstanceNetworkTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstanceNetworkTypeResponse{}
	_body, _err := client.UpdateInstanceNetworkTypeWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
