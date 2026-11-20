// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetInstanceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetInstanceResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *GetInstanceResponseBody
	GetHttpStatusCode() *string
	SetInstance(v *GetInstanceResponseBodyInstance) *GetInstanceResponseBody
	GetInstance() *GetInstanceResponseBodyInstance
	SetRequestId(v string) *GetInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetInstanceResponseBody
	GetSuccess() *bool
}

type GetInstanceResponseBody struct {
	// The error code.
	//
	// example:
	//
	// null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
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
	// The instance information.
	Instance *GetInstanceResponseBodyInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 865A02C2-B374-5DD4-9B34-0CA15DA1AEBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request result. This parameter indicates whether an exception occurred during the request and is independent of the business logic.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetInstanceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetInstanceResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *GetInstanceResponseBody) GetInstance() *GetInstanceResponseBodyInstance {
	return s.Instance
}

func (s *GetInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceResponseBody) GetSuccess() *bool {
	return s.Success
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

func (s *GetInstanceResponseBody) Validate() error {
	if s.Instance != nil {
		if err := s.Instance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceResponseBodyInstance struct {
	// Indicates whether auto-renewal is enabled.
	//
	// example:
	//
	// true
	AutoRenewal *string `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// The cold storage capacity.
	//
	// Unit: GB.
	//
	// SSD is equivalent to hot storage, which is also standard storage. HDD is equivalent to cold storage.
	//
	// example:
	//
	// 800
	ColdStorage *int64 `json:"ColdStorage,omitempty" xml:"ColdStorage,omitempty"`
	// The commodity code.
	//
	// example:
	//
	// hologram_combo_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The number of compute nodes.
	//
	// A typical configuration for a node is 16 vCPUs and 32 GB of memory.
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
	// The creation time. The time is displayed in UTC and formatted in accordance with the ISO 8601 standard.
	//
	// example:
	//
	// 2021-02-03T13:06:06Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The standard storage size.
	//
	// Unit: GB.
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
	// Indicates whether SSL is enabled.
	//
	// example:
	//
	// true
	EnableSSL *bool `json:"EnableSSL,omitempty" xml:"EnableSSL,omitempty"`
	// Indicates whether the instance sandbox feature is enabled.
	EnableSandbox *bool `json:"EnableSandbox,omitempty" xml:"EnableSandbox,omitempty"`
	// Indicates whether Serverless Computing is enabled.
	//
	// example:
	//
	// true
	EnableServerless *bool `json:"EnableServerless,omitempty" xml:"EnableServerless,omitempty"`
	// The list of endpoints.
	Endpoints []*GetInstanceResponseBodyInstanceEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// The expiration time. The time is displayed in UTC and formatted in accordance with the ISO 8601 standard.
	//
	// Pay-as-you-go instances do not have an expiration time.
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
	// The CPU resources of the gateway.
	//
	// Unit: Cores.
	//
	// example:
	//
	// 4
	GatewayCpu *int64 `json:"GatewayCpu,omitempty" xml:"GatewayCpu,omitempty"`
	// The memory resources of the gateway.
	//
	// Unit: GB.
	//
	// example:
	//
	// 16
	GatewayMemory *int64 `json:"GatewayMemory,omitempty" xml:"GatewayMemory,omitempty"`
	// The billing method of the instance.
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
	// The instance name.
	//
	// The name must be 2 to 64 characters in length.
	//
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance owner.
	//
	// example:
	//
	// 12345678900000
	InstanceOwner *string `json:"InstanceOwner,omitempty" xml:"InstanceOwner,omitempty"`
	// The instance status.
	//
	// example:
	//
	// Running
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The instance type.
	//
	// example:
	//
	// Standard
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The primary instance ID.
	//
	// example:
	//
	// hgpostcn-cn-i7m2ncd6w002
	LeaderInstanceId *string `json:"LeaderInstanceId,omitempty" xml:"LeaderInstanceId,omitempty"`
	// The memory size.
	//
	// Unit: GB.
	//
	// example:
	//
	// 128
	Memory *int64 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The role of the disaster recovery instance.
	//
	// Active: The primary disaster recovery instance.
	//
	// Passive: The secondary disaster recovery instance.
	//
	// PreActive: The primary disaster recovery instance that has not reached the desired state.
	//
	// example:
	//
	// Active
	ReplicaRole *string `json:"ReplicaRole,omitempty" xml:"ReplicaRole,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aekzuq7hpybze2i
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The storage type. Valid values:
	//
	// 	- redundant: 3AZ.
	//
	// 	- local: 1AZ.
	//
	// example:
	//
	// redundant
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The reason for suspension.
	//
	// example:
	//
	// Manual
	SuspendReason *string `json:"SuspendReason,omitempty" xml:"SuspendReason,omitempty"`
	// The instance tags.
	Tags []*GetInstanceResponseBodyInstanceTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The instance version.
	//
	// example:
	//
	// r1.3.37
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetInstanceResponseBodyInstance) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyInstance) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstance) GetAutoRenewal() *string {
	return s.AutoRenewal
}

func (s *GetInstanceResponseBodyInstance) GetColdStorage() *int64 {
	return s.ColdStorage
}

func (s *GetInstanceResponseBodyInstance) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *GetInstanceResponseBodyInstance) GetComputeNodeCount() *int64 {
	return s.ComputeNodeCount
}

func (s *GetInstanceResponseBodyInstance) GetCpu() *int64 {
	return s.Cpu
}

func (s *GetInstanceResponseBodyInstance) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetInstanceResponseBodyInstance) GetDisk() *string {
	return s.Disk
}

func (s *GetInstanceResponseBodyInstance) GetEnableHiveAccess() *string {
	return s.EnableHiveAccess
}

func (s *GetInstanceResponseBodyInstance) GetEnableSSL() *bool {
	return s.EnableSSL
}

func (s *GetInstanceResponseBodyInstance) GetEnableSandbox() *bool {
	return s.EnableSandbox
}

func (s *GetInstanceResponseBodyInstance) GetEnableServerless() *bool {
	return s.EnableServerless
}

func (s *GetInstanceResponseBodyInstance) GetEndpoints() []*GetInstanceResponseBodyInstanceEndpoints {
	return s.Endpoints
}

func (s *GetInstanceResponseBodyInstance) GetExpirationTime() *string {
	return s.ExpirationTime
}

func (s *GetInstanceResponseBodyInstance) GetGatewayCount() *int64 {
	return s.GatewayCount
}

func (s *GetInstanceResponseBodyInstance) GetGatewayCpu() *int64 {
	return s.GatewayCpu
}

func (s *GetInstanceResponseBodyInstance) GetGatewayMemory() *int64 {
	return s.GatewayMemory
}

func (s *GetInstanceResponseBodyInstance) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *GetInstanceResponseBodyInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceResponseBodyInstance) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetInstanceResponseBodyInstance) GetInstanceOwner() *string {
	return s.InstanceOwner
}

func (s *GetInstanceResponseBodyInstance) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *GetInstanceResponseBodyInstance) GetInstanceType() *string {
	return s.InstanceType
}

func (s *GetInstanceResponseBodyInstance) GetLeaderInstanceId() *string {
	return s.LeaderInstanceId
}

func (s *GetInstanceResponseBodyInstance) GetMemory() *int64 {
	return s.Memory
}

func (s *GetInstanceResponseBodyInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInstanceResponseBodyInstance) GetReplicaRole() *string {
	return s.ReplicaRole
}

func (s *GetInstanceResponseBodyInstance) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetInstanceResponseBodyInstance) GetStorageType() *string {
	return s.StorageType
}

func (s *GetInstanceResponseBodyInstance) GetSuspendReason() *string {
	return s.SuspendReason
}

func (s *GetInstanceResponseBodyInstance) GetTags() []*GetInstanceResponseBodyInstanceTags {
	return s.Tags
}

func (s *GetInstanceResponseBodyInstance) GetVersion() *string {
	return s.Version
}

func (s *GetInstanceResponseBodyInstance) GetZoneId() *string {
	return s.ZoneId
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

func (s *GetInstanceResponseBodyInstance) SetEnableSSL(v bool) *GetInstanceResponseBodyInstance {
	s.EnableSSL = &v
	return s
}

func (s *GetInstanceResponseBodyInstance) SetEnableSandbox(v bool) *GetInstanceResponseBodyInstance {
	s.EnableSandbox = &v
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

func (s *GetInstanceResponseBodyInstance) Validate() error {
	if s.Endpoints != nil {
		for _, item := range s.Endpoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetInstanceResponseBodyInstanceEndpoints struct {
	// Some legacy instances have both AnyTunnel and SingleTunnel enabled. When switching from AnyTunnel to SingleTunnel, the endpoints for both are retained. Therefore, an additional field is required to store the endpoints.
	//
	// example:
	//
	// hgprecn-cn-uqm362o1b001-cn-hangzhou-internal.hologres.aliyuncs.com:80
	AlternativeEndpoints *string `json:"AlternativeEndpoints,omitempty" xml:"AlternativeEndpoints,omitempty"`
	// Indicates whether the network is enabled.
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
	// example:
	//
	// Internet
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp1jqwp2ys6kp7tc9t983
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-uf66jjber3hgvwhki3wna
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The VPC instance ID.
	//
	// example:
	//
	// hgprecn-cn-uqm362o1b001-frontend-st
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
}

func (s GetInstanceResponseBodyInstanceEndpoints) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyInstanceEndpoints) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstanceEndpoints) GetAlternativeEndpoints() *string {
	return s.AlternativeEndpoints
}

func (s *GetInstanceResponseBodyInstanceEndpoints) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetInstanceResponseBodyInstanceEndpoints) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetInstanceResponseBodyInstanceEndpoints) GetType() *string {
	return s.Type
}

func (s *GetInstanceResponseBodyInstanceEndpoints) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetInstanceResponseBodyInstanceEndpoints) GetVpcId() *string {
	return s.VpcId
}

func (s *GetInstanceResponseBodyInstanceEndpoints) GetVpcInstanceId() *string {
	return s.VpcInstanceId
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

func (s *GetInstanceResponseBodyInstanceEndpoints) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyInstanceTags struct {
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

func (s GetInstanceResponseBodyInstanceTags) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyInstanceTags) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyInstanceTags) GetKey() *string {
	return s.Key
}

func (s *GetInstanceResponseBodyInstanceTags) GetValue() *string {
	return s.Value
}

func (s *GetInstanceResponseBodyInstanceTags) SetKey(v string) *GetInstanceResponseBodyInstanceTags {
	s.Key = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceTags) SetValue(v string) *GetInstanceResponseBodyInstanceTags {
	s.Value = &v
	return s
}

func (s *GetInstanceResponseBodyInstanceTags) Validate() error {
	return dara.Validate(s)
}
