// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateInstanceRequest struct {
	AutoPay          *bool   `json:"autoPay,omitempty" xml:"autoPay,omitempty"`
	AutoRenew        *bool   `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
	ChargeType       *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	ColdStorageSize  *int64  `json:"coldStorageSize,omitempty" xml:"coldStorageSize,omitempty"`
	Cpu              *int64  `json:"cpu,omitempty" xml:"cpu,omitempty"`
	Duration         *int64  `json:"duration,omitempty" xml:"duration,omitempty"`
	GatewayCount     *int64  `json:"gatewayCount,omitempty" xml:"gatewayCount,omitempty"`
	InstanceName     *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	InstanceType     *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	LeaderInstanceId *string `json:"leaderInstanceId,omitempty" xml:"leaderInstanceId,omitempty"`
	PricingCycle     *string `json:"pricingCycle,omitempty" xml:"pricingCycle,omitempty"`
	RegionId         *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	ResourceGroupId  *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	StorageSize      *int64  `json:"storageSize,omitempty" xml:"storageSize,omitempty"`
	VSwitchId        *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	VpcId            *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	ZoneId           *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
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

func (s *CreateInstanceRequest) SetGatewayCount(v int64) *CreateInstanceRequest {
	s.GatewayCount = &v
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
	Data           *CreateInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode      *string                         `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                         `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HttpStatusCode *string                         `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Id of the request
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
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	OrderId    *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DeleteInstanceRequest struct {
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
	Data           *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorCode      *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type GetInstanceResponseBody struct {
	// The error code returned if the request failed.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The details of the instance.
	Instance *GetInstanceResponseBodyInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request result, which indicates whether the request was successful.
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
	// *   true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// *   false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	AutoRenewal *string `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// The cold storage capacity of the instance. Unit: GB. Standard SSD is used for hot storage, and HDD is used for cold storage.
	ColdStorage *int64 `json:"ColdStorage,omitempty" xml:"ColdStorage,omitempty"`
	// The commodity code.
	//
	// Valid values:
	//
	// *   hologram_maxcomputeAccelerate_public_cn
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
	// *   hologram_combo_public_cn
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
	// *   hologram_prepay_public_intl
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
	// *   hologram_storage_dp_cn
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
	// *   hologram_postpay_public_cn
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
	// *   hologram_postpay_public_intl
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
	// *   hologram_maxcomputeAccelerate_public_intl
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
	// *   hologram_cu_dp_cn
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
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The number of compute nodes. In a typical configuration, a node has 16 vCPUs and 32 GB of memory.
	ComputeNodeCount *int64 `json:"ComputeNodeCount,omitempty" xml:"ComputeNodeCount,omitempty"`
	// The number of vCPUs.
	Cpu *int64 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The creation time.
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The amount of data that can be stored in the disk of the Standard storage class. Unit: GB.
	Disk *string `json:"Disk,omitempty" xml:"Disk,omitempty"`
	// Indicates whether data lake acceleration is enabled.
	EnableHiveAccess *string `json:"EnableHiveAccess,omitempty" xml:"EnableHiveAccess,omitempty"`
	// The list of endpoints.
	Endpoints []*GetInstanceResponseBodyInstanceEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// The expiration time. This parameter is invalid for pay-as-you-go instances.
	ExpirationTime *string `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// 网关节点数量。
	GatewayCount *int64 `json:"GatewayCount,omitempty" xml:"GatewayCount,omitempty"`
	// 网关cpu资源。
	// 单位：core。
	GatewayCpu *int64 `json:"GatewayCpu,omitempty" xml:"GatewayCpu,omitempty"`
	// 网关内存资源。
	// 单位：GB。
	GatewayMemory *int64 `json:"GatewayMemory,omitempty" xml:"GatewayMemory,omitempty"`
	// The billing method of the instance.
	//
	// Valid values:
	//
	// *   PostPaid
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
	// *   PrePaid
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
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance ID.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name. The instance name must be 2 to 64 characters in length.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The owner of the instance.
	InstanceOwner *string `json:"InstanceOwner,omitempty" xml:"InstanceOwner,omitempty"`
	// The status of the instance.
	//
	// Valid values:
	//
	// *   Creating
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// *   Running
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// *   Suspended
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// *   Allocating
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The type of the instance.
	//
	// Valid values:
	//
	// *   Follower
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
	// *   Standard
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
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The ID of the primary instance.
	LeaderInstanceId *string `json:"LeaderInstanceId,omitempty" xml:"LeaderInstanceId,omitempty"`
	// The memory size. Unit: GB.
	Memory   *int64  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The reason for the suspension.
	SuspendReason *string `json:"SuspendReason,omitempty" xml:"SuspendReason,omitempty"`
	// The instance tag.
	Tags []*GetInstanceResponseBodyInstanceTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The instance version.
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The ID of the zone where the instance resides.
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

func (s *GetInstanceResponseBodyInstance) SetResourceGroupId(v string) *GetInstanceResponseBodyInstance {
	s.ResourceGroupId = &v
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
	AlternativeEndpoints *string `json:"AlternativeEndpoints,omitempty" xml:"AlternativeEndpoints,omitempty"`
	// Indicates whether the network is enabled.
	//
	// Valid values:
	//
	// *   true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// *   false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The endpoint.
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The network type.
	//
	// Valid values:
	//
	// *   VPCSingleTunnel
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
	// *   Intranet
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
	// *   VPCAnyTunnel
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
	// *   Internet
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
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the vSwitch.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of VPC to which the instance belongs.
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
	// The tag key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ListInstancesRequest struct {
	CmsInstanceType *string `json:"cmsInstanceType,omitempty" xml:"cmsInstanceType,omitempty"`
	// The ID of the resource group.
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
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value.
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
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The list of queried instances.
	InstanceList []*ListInstancesResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
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
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The time when the cluster was created.
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Indicates whether lakehouse acceleration is enabled.
	//
	// Valid values:
	//
	// *   true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// *   false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	EnableHiveAccess *string `json:"EnableHiveAccess,omitempty" xml:"EnableHiveAccess,omitempty"`
	// The list of endpoints.
	Endpoints []*ListInstancesResponseBodyInstanceListEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// The time when the cluster expires.
	ExpirationTime *string `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// Valid values:
	//
	// *   PostPaid
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
	// *   PrePaid
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
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance ID.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The status of the instance.
	//
	// Valid values:
	//
	// *   Creating
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// *   Running
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// *   Suspended
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// *   Allocating
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The type of the instance.
	//
	// Valid values:
	//
	// *   Follower
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
	// *   Standard
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
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The ID of the primary instance.
	LeaderInstanceId *string `json:"LeaderInstanceId,omitempty" xml:"LeaderInstanceId,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The reason for the suspension.
	SuspendReason *string `json:"SuspendReason,omitempty" xml:"SuspendReason,omitempty"`
	// The tags that are added to the resource.
	Tags []*ListInstancesResponseBodyInstanceListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The version of the cluster.
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	ZoneId  *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	// *   true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// *   false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The endpoint.
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The network type.
	//
	// Valid values:
	//
	// *   VPCSingleTunnel
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
	// *   Intranet
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
	// *   VPCAnyTunnel
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
	// *   Internet
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
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The vSwitch ID.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the VPC to which the instance belongs.
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
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type RenewInstanceRequest struct {
	AutoRenew *bool  `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
	Duration  *int32 `json:"duration,omitempty" xml:"duration,omitempty"`
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
	Data           *RenewInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode      *string                        `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                        `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HttpStatusCode *string                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RenewInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type RestartInstanceResponseBody struct {
	// Indicates whether the operation was successful.
	//
	// Valid values:
	//
	// *   true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// *   false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned if the request failed.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request result, which indicates whether the request was successful.
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RestartInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ResumeInstanceResponseBody struct {
	// The returned result, which indicates whether the operation was successful.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned if the request failed.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request result, which indicates whether the request was successful.
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResumeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ScaleInstanceRequest struct {
	ColdStorageSize *int64  `json:"coldStorageSize,omitempty" xml:"coldStorageSize,omitempty"`
	Cpu             *int64  `json:"cpu,omitempty" xml:"cpu,omitempty"`
	ScaleType       *string `json:"scaleType,omitempty" xml:"scaleType,omitempty"`
	StorageSize     *int64  `json:"storageSize,omitempty" xml:"storageSize,omitempty"`
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

func (s *ScaleInstanceRequest) SetScaleType(v string) *ScaleInstanceRequest {
	s.ScaleType = &v
	return s
}

func (s *ScaleInstanceRequest) SetStorageSize(v int64) *ScaleInstanceRequest {
	s.StorageSize = &v
	return s
}

type ScaleInstanceResponseBody struct {
	Data           *ScaleInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorCode      *string                        `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                        `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HttpStatusCode *string                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Id of the request
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
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ScaleInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// *   true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// *   false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned if the request failed.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request result, which indicates whether the request was successful.
	//
	// Valid values:
	//
	// *   true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// *   false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type UpdateInstanceNameRequest struct {
	// The new name of the instance.
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
	// *   true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// *   false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned if the request failed.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request result, which indicates whether the request was successful.
	//
	// Valid values:
	//
	// *   true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// *   false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateInstanceNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// *   others/null
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// *   true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	AnyTunnelToSingleTunnel *string `json:"anyTunnelToSingleTunnel,omitempty" xml:"anyTunnelToSingleTunnel,omitempty"`
	// A list of network types that you want to enable. The list of enabled network types is randomly ordered. For example, the Internet, internal network, and VPCSingleTunnel network types are enabled. If you want to disable the Internet type, set this parameter to Intranet,VPCSingleTunnel.
	NetworkTypes *string `json:"networkTypes,omitempty" xml:"networkTypes,omitempty"`
	// The vSwitch ID.
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// The ID of the VPC to which the instance belongs.
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// The owner ID of the VPC, which is the ID of the Alibaba Cloud account.
	VpcOwnerId *string `json:"vpcOwnerId,omitempty" xml:"vpcOwnerId,omitempty"`
	// The region ID of the VPC.
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
	// *   true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// *   false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned if the request failed.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// Valid values:
	//
	// *   true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// *   false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateInstanceNetworkTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

	if !tea.BoolValue(util.IsUnset(request.GatewayCount)) {
		body["gatewayCount"] = request.GatewayCount
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
	_result = &CreateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &DeleteInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &GetInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ListInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &RenewInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &RestartInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ResumeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &ScaleInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &StopInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &UpdateInstanceNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
	_result = &UpdateInstanceNetworkTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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
