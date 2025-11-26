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
	EnableSSL        *bool   `json:"EnableSSL,omitempty" xml:"EnableSSL,omitempty"`
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
