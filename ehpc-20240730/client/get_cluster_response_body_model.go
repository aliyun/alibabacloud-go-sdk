// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClientVersion(v string) *GetClusterResponseBody
	GetClientVersion() *string
	SetClusterCategory(v string) *GetClusterResponseBody
	GetClusterCategory() *string
	SetClusterCreateTime(v string) *GetClusterResponseBody
	GetClusterCreateTime() *string
	SetClusterCustomConfiguration(v *GetClusterResponseBodyClusterCustomConfiguration) *GetClusterResponseBody
	GetClusterCustomConfiguration() *GetClusterResponseBodyClusterCustomConfiguration
	SetClusterId(v string) *GetClusterResponseBody
	GetClusterId() *string
	SetClusterMode(v string) *GetClusterResponseBody
	GetClusterMode() *string
	SetClusterModifyTime(v string) *GetClusterResponseBody
	GetClusterModifyTime() *string
	SetClusterName(v string) *GetClusterResponseBody
	GetClusterName() *string
	SetClusterStatus(v string) *GetClusterResponseBody
	GetClusterStatus() *string
	SetClusterVSwitchId(v string) *GetClusterResponseBody
	GetClusterVSwitchId() *string
	SetClusterVpcId(v string) *GetClusterResponseBody
	GetClusterVpcId() *string
	SetDeleteProtection(v string) *GetClusterResponseBody
	GetDeleteProtection() *string
	SetEhpcVersion(v string) *GetClusterResponseBody
	GetEhpcVersion() *string
	SetEnableScaleIn(v bool) *GetClusterResponseBody
	GetEnableScaleIn() *bool
	SetEnableScaleOut(v bool) *GetClusterResponseBody
	GetEnableScaleOut() *bool
	SetGrowInterval(v int32) *GetClusterResponseBody
	GetGrowInterval() *int32
	SetIdleInterval(v int32) *GetClusterResponseBody
	GetIdleInterval() *int32
	SetManager(v *GetClusterResponseBodyManager) *GetClusterResponseBody
	GetManager() *GetClusterResponseBodyManager
	SetMaxCoreCount(v string) *GetClusterResponseBody
	GetMaxCoreCount() *string
	SetMaxCount(v string) *GetClusterResponseBody
	GetMaxCount() *string
	SetMonitorSpec(v *GetClusterResponseBodyMonitorSpec) *GetClusterResponseBody
	GetMonitorSpec() *GetClusterResponseBodyMonitorSpec
	SetRequestId(v string) *GetClusterResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetClusterResponseBody
	GetResourceGroupId() *string
	SetSchedulerSpec(v *GetClusterResponseBodySchedulerSpec) *GetClusterResponseBody
	GetSchedulerSpec() *GetClusterResponseBodySchedulerSpec
	SetSecurityGroupId(v string) *GetClusterResponseBody
	GetSecurityGroupId() *string
}

type GetClusterResponseBody struct {
	// The E-HPC Util version.
	//
	// example:
	//
	// 2.0.31
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The cluster type. Valid values:
	//
	// 	- Standard
	//
	// 	- Serverless
	//
	// example:
	//
	// Standard
	ClusterCategory *string `json:"ClusterCategory,omitempty" xml:"ClusterCategory,omitempty"`
	// The time when the cluster was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time is displayed in UTC. For more information, see [ISO 8601](https://help.aliyun.com/document_detail/25696.html).
	//
	// example:
	//
	// 2024-08-06T12:43:01.000Z
	ClusterCreateTime *string `json:"ClusterCreateTime,omitempty" xml:"ClusterCreateTime,omitempty"`
	// The post-processing script of the cluster.
	ClusterCustomConfiguration *GetClusterResponseBodyClusterCustomConfiguration `json:"ClusterCustomConfiguration,omitempty" xml:"ClusterCustomConfiguration,omitempty" type:"Struct"`
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The deployment type of the cluster. Valid values:
	//
	// 	- Integrated: The cluster is deployed on a public cloud.
	//
	// 	- Hybrid: The cluster is deployed on a hybrid cloud.
	//
	// 	- Custom: The cluster is a custom cluster.
	//
	// example:
	//
	// Integrated
	ClusterMode *string `json:"ClusterMode,omitempty" xml:"ClusterMode,omitempty"`
	// The time when the cluster was last modified. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time is displayed in UTC. For more information, see [ISO 8601](https://help.aliyun.com/document_detail/25696.html).
	//
	// example:
	//
	// 2024-08-06T12:43:01.000Z
	ClusterModifyTime *string `json:"ClusterModifyTime,omitempty" xml:"ClusterModifyTime,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// slurm22.05.8-cluster-20240614
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The cluster state. Valid values:
	//
	// 	- uninit: The cluster is being installed.
	//
	// 	- creating: The cluster is being created.
	//
	// 	- initing: The cluster is being initialized.
	//
	// 	- running: The cluster is running.
	//
	// 	- exception: The cluster has run into an exception.
	//
	// 	- raleasing: The cluster is being released.
	//
	// 	- stopping: The cluster is being stopped.
	//
	// 	- stopped: The cluster is stopped.
	//
	// 	- pending: The cluster is waiting to be configured.
	//
	// example:
	//
	// running
	ClusterStatus *string `json:"ClusterStatus,omitempty" xml:"ClusterStatus,omitempty"`
	// The ID of the vSwitch used by the cluster.
	//
	// example:
	//
	// vsw-bp1p2uugqsjppno******
	ClusterVSwitchId *string `json:"ClusterVSwitchId,omitempty" xml:"ClusterVSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC) used by the cluster.
	//
	// example:
	//
	// vpc-uf6u3lk1pjy28eg*****
	ClusterVpcId *string `json:"ClusterVpcId,omitempty" xml:"ClusterVpcId,omitempty"`
	// Indicates whether deletion protection is enabled for the cluster. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	DeleteProtection *string `json:"DeleteProtection,omitempty" xml:"DeleteProtection,omitempty"`
	// The E-HPC version.
	//
	// example:
	//
	// 2.0.0
	EhpcVersion *string `json:"EhpcVersion,omitempty" xml:"EhpcVersion,omitempty"`
	// Indicates whether automatic scale-in is enabled for the cluster. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableScaleIn *bool `json:"EnableScaleIn,omitempty" xml:"EnableScaleIn,omitempty"`
	// Indicates whether automatic scale-out is enabled for the cluster. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableScaleOut *bool `json:"EnableScaleOut,omitempty" xml:"EnableScaleOut,omitempty"`
	// The interval at which the cluster is automatically scaled out.
	//
	// example:
	//
	// 2
	GrowInterval *int32 `json:"GrowInterval,omitempty" xml:"GrowInterval,omitempty"`
	// The idle duration of the compute nodes allowed by the cluster.
	//
	// example:
	//
	// 4
	IdleInterval *int32 `json:"IdleInterval,omitempty" xml:"IdleInterval,omitempty"`
	// The management node configurations.
	Manager *GetClusterResponseBodyManager `json:"Manager,omitempty" xml:"Manager,omitempty" type:"Struct"`
	// The maximum total number of vCPUs that can be used by all compute nodes managed by the cluster.
	//
	// example:
	//
	// 10000
	MaxCoreCount *string `json:"MaxCoreCount,omitempty" xml:"MaxCoreCount,omitempty"`
	// The maximum number of compute nodes that the cluster can manage.
	//
	// example:
	//
	// 100
	MaxCount *string `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	// The monitoring details of the cluster.
	MonitorSpec *GetClusterResponseBodyMonitorSpec `json:"MonitorSpec,omitempty" xml:"MonitorSpec,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxazb4ph****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The scheduler specifications of the cluster.
	SchedulerSpec *GetClusterResponseBodySchedulerSpec `json:"SchedulerSpec,omitempty" xml:"SchedulerSpec,omitempty" type:"Struct"`
	// The security group ID.
	//
	// example:
	//
	// sg-f8z9vb2zaezpkr69a21k
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s GetClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClusterResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBody) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *GetClusterResponseBody) GetClusterCategory() *string {
	return s.ClusterCategory
}

func (s *GetClusterResponseBody) GetClusterCreateTime() *string {
	return s.ClusterCreateTime
}

func (s *GetClusterResponseBody) GetClusterCustomConfiguration() *GetClusterResponseBodyClusterCustomConfiguration {
	return s.ClusterCustomConfiguration
}

func (s *GetClusterResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetClusterResponseBody) GetClusterMode() *string {
	return s.ClusterMode
}

func (s *GetClusterResponseBody) GetClusterModifyTime() *string {
	return s.ClusterModifyTime
}

func (s *GetClusterResponseBody) GetClusterName() *string {
	return s.ClusterName
}

func (s *GetClusterResponseBody) GetClusterStatus() *string {
	return s.ClusterStatus
}

func (s *GetClusterResponseBody) GetClusterVSwitchId() *string {
	return s.ClusterVSwitchId
}

func (s *GetClusterResponseBody) GetClusterVpcId() *string {
	return s.ClusterVpcId
}

func (s *GetClusterResponseBody) GetDeleteProtection() *string {
	return s.DeleteProtection
}

func (s *GetClusterResponseBody) GetEhpcVersion() *string {
	return s.EhpcVersion
}

func (s *GetClusterResponseBody) GetEnableScaleIn() *bool {
	return s.EnableScaleIn
}

func (s *GetClusterResponseBody) GetEnableScaleOut() *bool {
	return s.EnableScaleOut
}

func (s *GetClusterResponseBody) GetGrowInterval() *int32 {
	return s.GrowInterval
}

func (s *GetClusterResponseBody) GetIdleInterval() *int32 {
	return s.IdleInterval
}

func (s *GetClusterResponseBody) GetManager() *GetClusterResponseBodyManager {
	return s.Manager
}

func (s *GetClusterResponseBody) GetMaxCoreCount() *string {
	return s.MaxCoreCount
}

func (s *GetClusterResponseBody) GetMaxCount() *string {
	return s.MaxCount
}

func (s *GetClusterResponseBody) GetMonitorSpec() *GetClusterResponseBodyMonitorSpec {
	return s.MonitorSpec
}

func (s *GetClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClusterResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetClusterResponseBody) GetSchedulerSpec() *GetClusterResponseBodySchedulerSpec {
	return s.SchedulerSpec
}

func (s *GetClusterResponseBody) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetClusterResponseBody) SetClientVersion(v string) *GetClusterResponseBody {
	s.ClientVersion = &v
	return s
}

func (s *GetClusterResponseBody) SetClusterCategory(v string) *GetClusterResponseBody {
	s.ClusterCategory = &v
	return s
}

func (s *GetClusterResponseBody) SetClusterCreateTime(v string) *GetClusterResponseBody {
	s.ClusterCreateTime = &v
	return s
}

func (s *GetClusterResponseBody) SetClusterCustomConfiguration(v *GetClusterResponseBodyClusterCustomConfiguration) *GetClusterResponseBody {
	s.ClusterCustomConfiguration = v
	return s
}

func (s *GetClusterResponseBody) SetClusterId(v string) *GetClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *GetClusterResponseBody) SetClusterMode(v string) *GetClusterResponseBody {
	s.ClusterMode = &v
	return s
}

func (s *GetClusterResponseBody) SetClusterModifyTime(v string) *GetClusterResponseBody {
	s.ClusterModifyTime = &v
	return s
}

func (s *GetClusterResponseBody) SetClusterName(v string) *GetClusterResponseBody {
	s.ClusterName = &v
	return s
}

func (s *GetClusterResponseBody) SetClusterStatus(v string) *GetClusterResponseBody {
	s.ClusterStatus = &v
	return s
}

func (s *GetClusterResponseBody) SetClusterVSwitchId(v string) *GetClusterResponseBody {
	s.ClusterVSwitchId = &v
	return s
}

func (s *GetClusterResponseBody) SetClusterVpcId(v string) *GetClusterResponseBody {
	s.ClusterVpcId = &v
	return s
}

func (s *GetClusterResponseBody) SetDeleteProtection(v string) *GetClusterResponseBody {
	s.DeleteProtection = &v
	return s
}

func (s *GetClusterResponseBody) SetEhpcVersion(v string) *GetClusterResponseBody {
	s.EhpcVersion = &v
	return s
}

func (s *GetClusterResponseBody) SetEnableScaleIn(v bool) *GetClusterResponseBody {
	s.EnableScaleIn = &v
	return s
}

func (s *GetClusterResponseBody) SetEnableScaleOut(v bool) *GetClusterResponseBody {
	s.EnableScaleOut = &v
	return s
}

func (s *GetClusterResponseBody) SetGrowInterval(v int32) *GetClusterResponseBody {
	s.GrowInterval = &v
	return s
}

func (s *GetClusterResponseBody) SetIdleInterval(v int32) *GetClusterResponseBody {
	s.IdleInterval = &v
	return s
}

func (s *GetClusterResponseBody) SetManager(v *GetClusterResponseBodyManager) *GetClusterResponseBody {
	s.Manager = v
	return s
}

func (s *GetClusterResponseBody) SetMaxCoreCount(v string) *GetClusterResponseBody {
	s.MaxCoreCount = &v
	return s
}

func (s *GetClusterResponseBody) SetMaxCount(v string) *GetClusterResponseBody {
	s.MaxCount = &v
	return s
}

func (s *GetClusterResponseBody) SetMonitorSpec(v *GetClusterResponseBodyMonitorSpec) *GetClusterResponseBody {
	s.MonitorSpec = v
	return s
}

func (s *GetClusterResponseBody) SetRequestId(v string) *GetClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterResponseBody) SetResourceGroupId(v string) *GetClusterResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetClusterResponseBody) SetSchedulerSpec(v *GetClusterResponseBodySchedulerSpec) *GetClusterResponseBody {
	s.SchedulerSpec = v
	return s
}

func (s *GetClusterResponseBody) SetSecurityGroupId(v string) *GetClusterResponseBody {
	s.SecurityGroupId = &v
	return s
}

func (s *GetClusterResponseBody) Validate() error {
	if s.ClusterCustomConfiguration != nil {
		if err := s.ClusterCustomConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.Manager != nil {
		if err := s.Manager.Validate(); err != nil {
			return err
		}
	}
	if s.MonitorSpec != nil {
		if err := s.MonitorSpec.Validate(); err != nil {
			return err
		}
	}
	if s.SchedulerSpec != nil {
		if err := s.SchedulerSpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetClusterResponseBodyClusterCustomConfiguration struct {
	// The arguments that are used to run the script after the scrip is installed.
	//
	// example:
	//
	// E-HPC cn-hangzhou
	Args *string `json:"Args,omitempty" xml:"Args,omitempty"`
	// The URL that is used to download the post-processing script.
	//
	// example:
	//
	// http://*****
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
}

func (s GetClusterResponseBodyClusterCustomConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetClusterResponseBodyClusterCustomConfiguration) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyClusterCustomConfiguration) GetArgs() *string {
	return s.Args
}

func (s *GetClusterResponseBodyClusterCustomConfiguration) GetScript() *string {
	return s.Script
}

func (s *GetClusterResponseBodyClusterCustomConfiguration) SetArgs(v string) *GetClusterResponseBodyClusterCustomConfiguration {
	s.Args = &v
	return s
}

func (s *GetClusterResponseBodyClusterCustomConfiguration) SetScript(v string) *GetClusterResponseBodyClusterCustomConfiguration {
	s.Script = &v
	return s
}

func (s *GetClusterResponseBodyClusterCustomConfiguration) Validate() error {
	return dara.Validate(s)
}

type GetClusterResponseBodyManager struct {
	// The configurations of the domain name resolution service.
	DNS *GetClusterResponseBodyManagerDNS `json:"DNS,omitempty" xml:"DNS,omitempty" type:"Struct"`
	// The information about the domain account service.
	DirectoryService *GetClusterResponseBodyManagerDirectoryService `json:"DirectoryService,omitempty" xml:"DirectoryService,omitempty" type:"Struct"`
	// The configurations of the management node.
	ManagerNode *GetClusterResponseBodyManagerManagerNode `json:"ManagerNode,omitempty" xml:"ManagerNode,omitempty" type:"Struct"`
	// The information about the scheduler.
	Scheduler *GetClusterResponseBodyManagerScheduler `json:"Scheduler,omitempty" xml:"Scheduler,omitempty" type:"Struct"`
}

func (s GetClusterResponseBodyManager) String() string {
	return dara.Prettify(s)
}

func (s GetClusterResponseBodyManager) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyManager) GetDNS() *GetClusterResponseBodyManagerDNS {
	return s.DNS
}

func (s *GetClusterResponseBodyManager) GetDirectoryService() *GetClusterResponseBodyManagerDirectoryService {
	return s.DirectoryService
}

func (s *GetClusterResponseBodyManager) GetManagerNode() *GetClusterResponseBodyManagerManagerNode {
	return s.ManagerNode
}

func (s *GetClusterResponseBodyManager) GetScheduler() *GetClusterResponseBodyManagerScheduler {
	return s.Scheduler
}

func (s *GetClusterResponseBodyManager) SetDNS(v *GetClusterResponseBodyManagerDNS) *GetClusterResponseBodyManager {
	s.DNS = v
	return s
}

func (s *GetClusterResponseBodyManager) SetDirectoryService(v *GetClusterResponseBodyManagerDirectoryService) *GetClusterResponseBodyManager {
	s.DirectoryService = v
	return s
}

func (s *GetClusterResponseBodyManager) SetManagerNode(v *GetClusterResponseBodyManagerManagerNode) *GetClusterResponseBodyManager {
	s.ManagerNode = v
	return s
}

func (s *GetClusterResponseBodyManager) SetScheduler(v *GetClusterResponseBodyManagerScheduler) *GetClusterResponseBodyManager {
	s.Scheduler = v
	return s
}

func (s *GetClusterResponseBodyManager) Validate() error {
	if s.DNS != nil {
		if err := s.DNS.Validate(); err != nil {
			return err
		}
	}
	if s.DirectoryService != nil {
		if err := s.DirectoryService.Validate(); err != nil {
			return err
		}
	}
	if s.ManagerNode != nil {
		if err := s.ManagerNode.Validate(); err != nil {
			return err
		}
	}
	if s.Scheduler != nil {
		if err := s.Scheduler.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetClusterResponseBodyManagerDNS struct {
	// The state of the domain name resolution service. Valid values:
	//
	// 	- uninit: The service is being installed.
	//
	// 	- initing: The service is being initialized.
	//
	// 	- running: The service is running.
	//
	// 	- exception: The service has run into an exception.
	//
	// 	- releasing: The service is being released.
	//
	// 	- stopped: The service is stopped.
	//
	// 	- pending: The service is waiting to be configured.
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The resolution type.
	//
	// example:
	//
	// nis
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version of the resolution service.
	//
	// example:
	//
	// 2.31
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetClusterResponseBodyManagerDNS) String() string {
	return dara.Prettify(s)
}

func (s GetClusterResponseBodyManagerDNS) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyManagerDNS) GetStatus() *string {
	return s.Status
}

func (s *GetClusterResponseBodyManagerDNS) GetType() *string {
	return s.Type
}

func (s *GetClusterResponseBodyManagerDNS) GetVersion() *string {
	return s.Version
}

func (s *GetClusterResponseBodyManagerDNS) SetStatus(v string) *GetClusterResponseBodyManagerDNS {
	s.Status = &v
	return s
}

func (s *GetClusterResponseBodyManagerDNS) SetType(v string) *GetClusterResponseBodyManagerDNS {
	s.Type = &v
	return s
}

func (s *GetClusterResponseBodyManagerDNS) SetVersion(v string) *GetClusterResponseBodyManagerDNS {
	s.Version = &v
	return s
}

func (s *GetClusterResponseBodyManagerDNS) Validate() error {
	return dara.Validate(s)
}

type GetClusterResponseBodyManagerDirectoryService struct {
	// The state of the domain account service. Valid values:
	//
	// 	- uninit: The service is being installed.
	//
	// 	- initing: The service is being initialized.
	//
	// 	- running: The service is running.
	//
	// 	- exception: The service has run into an exception.
	//
	// 	- releasing: The service is being released.
	//
	// 	- stopped: The service is stopped.
	//
	// 	- pending: The service is waiting to be configured.
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the domain account.
	//
	// example:
	//
	// nis
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version of the domain account service.
	//
	// example:
	//
	// 2.31
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetClusterResponseBodyManagerDirectoryService) String() string {
	return dara.Prettify(s)
}

func (s GetClusterResponseBodyManagerDirectoryService) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyManagerDirectoryService) GetStatus() *string {
	return s.Status
}

func (s *GetClusterResponseBodyManagerDirectoryService) GetType() *string {
	return s.Type
}

func (s *GetClusterResponseBodyManagerDirectoryService) GetVersion() *string {
	return s.Version
}

func (s *GetClusterResponseBodyManagerDirectoryService) SetStatus(v string) *GetClusterResponseBodyManagerDirectoryService {
	s.Status = &v
	return s
}

func (s *GetClusterResponseBodyManagerDirectoryService) SetType(v string) *GetClusterResponseBodyManagerDirectoryService {
	s.Type = &v
	return s
}

func (s *GetClusterResponseBodyManagerDirectoryService) SetVersion(v string) *GetClusterResponseBodyManagerDirectoryService {
	s.Version = &v
	return s
}

func (s *GetClusterResponseBodyManagerDirectoryService) Validate() error {
	return dara.Validate(s)
}

type GetClusterResponseBodyManagerManagerNode struct {
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// example:
	//
	// Month
	AutoRenewPeriod *int64 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// example:
	//
	// 1
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	EnableHt *bool  `json:"EnableHt,omitempty" xml:"EnableHt,omitempty"`
	// The expiration time of the management node.
	//
	// example:
	//
	// 2099-12-31T15:59Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// example:
	//
	// aliyun_2_1903_x64_20G_alibase_20200324.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The instance billing method of the management node. Valid values:
	//
	// 	- PostPaid: pay-as-you-go
	//
	// 	- PrePaid: subscription
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance ID of the management node.
	//
	// example:
	//
	// i-bp1a170jgea1vl******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance type of the management node.
	//
	// example:
	//
	// ecs.g6.4xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// 1
	Period *int64 `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// example:
	//
	// 1
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// example:
	//
	// NoSpot
	SpotStrategy *string                                             `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	SystemDisk   *GetClusterResponseBodyManagerManagerNodeSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
}

func (s GetClusterResponseBodyManagerManagerNode) String() string {
	return dara.Prettify(s)
}

func (s GetClusterResponseBodyManagerManagerNode) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyManagerManagerNode) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *GetClusterResponseBodyManagerManagerNode) GetAutoRenewPeriod() *int64 {
	return s.AutoRenewPeriod
}

func (s *GetClusterResponseBodyManagerManagerNode) GetDuration() *int64 {
	return s.Duration
}

func (s *GetClusterResponseBodyManagerManagerNode) GetEnableHt() *bool {
	return s.EnableHt
}

func (s *GetClusterResponseBodyManagerManagerNode) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *GetClusterResponseBodyManagerManagerNode) GetImageId() *string {
	return s.ImageId
}

func (s *GetClusterResponseBodyManagerManagerNode) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *GetClusterResponseBodyManagerManagerNode) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetClusterResponseBodyManagerManagerNode) GetInstanceType() *string {
	return s.InstanceType
}

func (s *GetClusterResponseBodyManagerManagerNode) GetPeriod() *int64 {
	return s.Period
}

func (s *GetClusterResponseBodyManagerManagerNode) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *GetClusterResponseBodyManagerManagerNode) GetSpotPriceLimit() *float32 {
	return s.SpotPriceLimit
}

func (s *GetClusterResponseBodyManagerManagerNode) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *GetClusterResponseBodyManagerManagerNode) GetSystemDisk() *GetClusterResponseBodyManagerManagerNodeSystemDisk {
	return s.SystemDisk
}

func (s *GetClusterResponseBodyManagerManagerNode) SetAutoRenew(v bool) *GetClusterResponseBodyManagerManagerNode {
	s.AutoRenew = &v
	return s
}

func (s *GetClusterResponseBodyManagerManagerNode) SetAutoRenewPeriod(v int64) *GetClusterResponseBodyManagerManagerNode {
	s.AutoRenewPeriod = &v
	return s
}

func (s *GetClusterResponseBodyManagerManagerNode) SetDuration(v int64) *GetClusterResponseBodyManagerManagerNode {
	s.Duration = &v
	return s
}

func (s *GetClusterResponseBodyManagerManagerNode) SetEnableHt(v bool) *GetClusterResponseBodyManagerManagerNode {
	s.EnableHt = &v
	return s
}

func (s *GetClusterResponseBodyManagerManagerNode) SetExpiredTime(v string) *GetClusterResponseBodyManagerManagerNode {
	s.ExpiredTime = &v
	return s
}

func (s *GetClusterResponseBodyManagerManagerNode) SetImageId(v string) *GetClusterResponseBodyManagerManagerNode {
	s.ImageId = &v
	return s
}

func (s *GetClusterResponseBodyManagerManagerNode) SetInstanceChargeType(v string) *GetClusterResponseBodyManagerManagerNode {
	s.InstanceChargeType = &v
	return s
}

func (s *GetClusterResponseBodyManagerManagerNode) SetInstanceId(v string) *GetClusterResponseBodyManagerManagerNode {
	s.InstanceId = &v
	return s
}

func (s *GetClusterResponseBodyManagerManagerNode) SetInstanceType(v string) *GetClusterResponseBodyManagerManagerNode {
	s.InstanceType = &v
	return s
}

func (s *GetClusterResponseBodyManagerManagerNode) SetPeriod(v int64) *GetClusterResponseBodyManagerManagerNode {
	s.Period = &v
	return s
}

func (s *GetClusterResponseBodyManagerManagerNode) SetPeriodUnit(v string) *GetClusterResponseBodyManagerManagerNode {
	s.PeriodUnit = &v
	return s
}

func (s *GetClusterResponseBodyManagerManagerNode) SetSpotPriceLimit(v float32) *GetClusterResponseBodyManagerManagerNode {
	s.SpotPriceLimit = &v
	return s
}

func (s *GetClusterResponseBodyManagerManagerNode) SetSpotStrategy(v string) *GetClusterResponseBodyManagerManagerNode {
	s.SpotStrategy = &v
	return s
}

func (s *GetClusterResponseBodyManagerManagerNode) SetSystemDisk(v *GetClusterResponseBodyManagerManagerNodeSystemDisk) *GetClusterResponseBodyManagerManagerNode {
	s.SystemDisk = v
	return s
}

func (s *GetClusterResponseBodyManagerManagerNode) Validate() error {
	if s.SystemDisk != nil {
		if err := s.SystemDisk.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetClusterResponseBodyManagerManagerNodeSystemDisk struct {
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// PL0
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// 40
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s GetClusterResponseBodyManagerManagerNodeSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s GetClusterResponseBodyManagerManagerNodeSystemDisk) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyManagerManagerNodeSystemDisk) GetCategory() *string {
	return s.Category
}

func (s *GetClusterResponseBodyManagerManagerNodeSystemDisk) GetLevel() *string {
	return s.Level
}

func (s *GetClusterResponseBodyManagerManagerNodeSystemDisk) GetSize() *int64 {
	return s.Size
}

func (s *GetClusterResponseBodyManagerManagerNodeSystemDisk) SetCategory(v string) *GetClusterResponseBodyManagerManagerNodeSystemDisk {
	s.Category = &v
	return s
}

func (s *GetClusterResponseBodyManagerManagerNodeSystemDisk) SetLevel(v string) *GetClusterResponseBodyManagerManagerNodeSystemDisk {
	s.Level = &v
	return s
}

func (s *GetClusterResponseBodyManagerManagerNodeSystemDisk) SetSize(v int64) *GetClusterResponseBodyManagerManagerNodeSystemDisk {
	s.Size = &v
	return s
}

func (s *GetClusterResponseBodyManagerManagerNodeSystemDisk) Validate() error {
	return dara.Validate(s)
}

type GetClusterResponseBodyManagerScheduler struct {
	// The scheduler state. Valid values:
	//
	// 	- uninit: The scheduler is being installed.
	//
	// 	- initing: The scheduler is being initialized.
	//
	// 	- running: The scheduler is running.
	//
	// 	- exception: The scheduler has run into an exception.
	//
	// 	- releasing: The scheduler is being released.
	//
	// 	- stopped: The scheduler is stopped.
	//
	// 	- pending: The scheduler is waiting to be configured.
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The scheduler type. Valid values:
	//
	// 	- SLURM
	//
	// 	- PBS
	//
	// 	- OPENGRIDSCHEDULER
	//
	// 	- LSF_PLUGIN
	//
	// 	- PBS_PLUGIN
	//
	// example:
	//
	// SLURM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The scheduler version.
	//
	// example:
	//
	// 22.05.8
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetClusterResponseBodyManagerScheduler) String() string {
	return dara.Prettify(s)
}

func (s GetClusterResponseBodyManagerScheduler) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyManagerScheduler) GetStatus() *string {
	return s.Status
}

func (s *GetClusterResponseBodyManagerScheduler) GetType() *string {
	return s.Type
}

func (s *GetClusterResponseBodyManagerScheduler) GetVersion() *string {
	return s.Version
}

func (s *GetClusterResponseBodyManagerScheduler) SetStatus(v string) *GetClusterResponseBodyManagerScheduler {
	s.Status = &v
	return s
}

func (s *GetClusterResponseBodyManagerScheduler) SetType(v string) *GetClusterResponseBodyManagerScheduler {
	s.Type = &v
	return s
}

func (s *GetClusterResponseBodyManagerScheduler) SetVersion(v string) *GetClusterResponseBodyManagerScheduler {
	s.Version = &v
	return s
}

func (s *GetClusterResponseBodyManagerScheduler) Validate() error {
	return dara.Validate(s)
}

type GetClusterResponseBodyMonitorSpec struct {
	// Indicates whether the monitoring component of compute nodes is enabled for the cluster. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableComputeLoadMonitor *bool `json:"EnableComputeLoadMonitor,omitempty" xml:"EnableComputeLoadMonitor,omitempty"`
}

func (s GetClusterResponseBodyMonitorSpec) String() string {
	return dara.Prettify(s)
}

func (s GetClusterResponseBodyMonitorSpec) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyMonitorSpec) GetEnableComputeLoadMonitor() *bool {
	return s.EnableComputeLoadMonitor
}

func (s *GetClusterResponseBodyMonitorSpec) SetEnableComputeLoadMonitor(v bool) *GetClusterResponseBodyMonitorSpec {
	s.EnableComputeLoadMonitor = &v
	return s
}

func (s *GetClusterResponseBodyMonitorSpec) Validate() error {
	return dara.Validate(s)
}

type GetClusterResponseBodySchedulerSpec struct {
	// Indicates whether the topology awareness feature is enabled for the cluster. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableTopologyAwareness *bool `json:"EnableTopologyAwareness,omitempty" xml:"EnableTopologyAwareness,omitempty"`
}

func (s GetClusterResponseBodySchedulerSpec) String() string {
	return dara.Prettify(s)
}

func (s GetClusterResponseBodySchedulerSpec) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodySchedulerSpec) GetEnableTopologyAwareness() *bool {
	return s.EnableTopologyAwareness
}

func (s *GetClusterResponseBodySchedulerSpec) SetEnableTopologyAwareness(v bool) *GetClusterResponseBodySchedulerSpec {
	s.EnableTopologyAwareness = &v
	return s
}

func (s *GetClusterResponseBodySchedulerSpec) Validate() error {
	return dara.Validate(s)
}
