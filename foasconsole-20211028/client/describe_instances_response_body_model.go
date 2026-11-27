// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody
	GetInstances() []*DescribeInstancesResponseBodyInstances
	SetPageIndex(v int32) *DescribeInstancesResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *DescribeInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeInstancesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *DescribeInstancesResponseBody
	GetTotalCount() *int64
	SetTotalPage(v int32) *DescribeInstancesResponseBody
	GetTotalPage() *int32
}

type DescribeInstancesResponseBody struct {
	// The instance group details.
	Instances []*DescribeInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// The number of instances per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C8DF2A5B-6FBA-5651-A3D4-960F3664****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of instances.
	//
	// example:
	//
	// 15
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 2
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBody) GetInstances() []*DescribeInstancesResponseBodyInstances {
	return s.Instances
}

func (s *DescribeInstancesResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *DescribeInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeInstancesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeInstancesResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeInstancesResponseBody) SetInstances(v []*DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageIndex(v int32) *DescribeInstancesResponseBody {
	s.PageIndex = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageSize(v int32) *DescribeInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetRequestId(v string) *DescribeInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetSuccess(v bool) *DescribeInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetTotalCount(v int64) *DescribeInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetTotalPage(v int32) *DescribeInstancesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeInstancesResponseBody) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstances struct {
	// Indicates whether the network type is ANSM.
	//
	// example:
	//
	// true
	Ansm *bool `json:"Ansm,omitempty" xml:"Ansm,omitempty"`
	// The processor architecture.
	//
	// example:
	//
	// X86
	ArchitectureType *string `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// c20c******404
	AskClusterId *string `json:"AskClusterId,omitempty" xml:"AskClusterId,omitempty"`
	// The billing method. Valid values:
	//
	// - POST: pay-as-you-go
	//
	// - PRE: subscription
	//
	// example:
	//
	// PRE
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The cluster state.
	ClusterState *DescribeInstancesResponseBodyInstancesClusterState `json:"ClusterState,omitempty" xml:"ClusterState,omitempty" type:"Struct"`
	// The cluster status. Valid values:
	//
	// - CREATING: The cluster is being created.
	//
	// - RUNNING: The cluster is running.
	//
	// - DISABLE: The cluster is invalid.
	//
	// - DELETING: The cluster is being deleted.
	//
	// - DELETED: The cluster is deleted.
	//
	// example:
	//
	// RUNNING
	ClusterStatus *string `json:"ClusterStatus,omitempty" xml:"ClusterStatus,omitempty"`
	// The overall resource usage of the Flink service.
	ClusterUsedResources *DescribeInstancesResponseBodyInstancesClusterUsedResources `json:"ClusterUsedResources,omitempty" xml:"ClusterUsedResources,omitempty" type:"Struct"`
	// The information about the storage resources used by the cluster.
	ClusterUsedStorage *DescribeInstancesResponseBodyInstancesClusterUsedStorage `json:"ClusterUsedStorage,omitempty" xml:"ClusterUsedStorage,omitempty" type:"Struct"`
	// Indicates whether deletion protection is enabled.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// Indicates whether elastic resources are enabled.
	//
	// example:
	//
	// false
	Elastic *bool `json:"Elastic,omitempty" xml:"Elastic,omitempty"`
	// The elastic order ID.
	//
	// example:
	//
	// f-cn-e3afbd321
	ElasticInstanceId *string `json:"ElasticInstanceId,omitempty" xml:"ElasticInstanceId,omitempty"`
	// The status of the elastic resource order.
	//
	// example:
	//
	// NORMAL
	ElasticOrderState *string `json:"ElasticOrderState,omitempty" xml:"ElasticOrderState,omitempty"`
	// The specifications of pay-as-you-go elastic resources.
	ElasticResourceSpec *DescribeInstancesResponseBodyInstancesElasticResourceSpec `json:"ElasticResourceSpec,omitempty" xml:"ElasticResourceSpec,omitempty" type:"Struct"`
	// Indicates whether zone-disaster recovery resources are selected.
	//
	// example:
	//
	// true
	Ha *bool `json:"Ha,omitempty" xml:"Ha,omitempty"`
	// The description of zone-disaster recovery resources.
	HaResourceSpec *DescribeInstancesResponseBodyInstancesHaResourceSpec `json:"HaResourceSpec,omitempty" xml:"HaResourceSpec,omitempty" type:"Struct"`
	// The vSwitch group in the secondary zone for zone-disaster recovery.
	HaVSwitchIds []*string `json:"HaVSwitchIds,omitempty" xml:"HaVSwitchIds,omitempty" type:"Repeated"`
	// The information about the vSwitch group in the secondary zone for zone-disaster recovery.
	HaVSwitchInfo []*DescribeInstancesResponseBodyInstancesHaVSwitchInfo `json:"HaVSwitchInfo,omitempty" xml:"HaVSwitchInfo,omitempty" type:"Repeated"`
	// The secondary zone ID for zone-disaster recovery.
	//
	// example:
	//
	// cn-hongkong-b
	HaZoneId *string `json:"HaZoneId,omitempty" xml:"HaZoneId,omitempty"`
	// The domain name information added by the user.
	//
	// This parameter is required.
	HostAliases []*DescribeInstancesResponseBodyInstancesHostAliases `json:"HostAliases,omitempty" xml:"HostAliases,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// example:
	//
	// f-cn-zvp2q0z****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The workspace name.
	//
	// example:
	//
	// vvp1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The type of the monitoring and alerting service. Valid values: ARMS and CloudMonitor.
	//
	// example:
	//
	// TAIHAO
	MonitorType *string `json:"MonitorType,omitempty" xml:"MonitorType,omitempty"`
	// The order status of the Realtime Compute for Apache Flink service. Valid values:
	//
	// - NOT_INIT: The order is placed but components are not deployed.
	//
	// - NORMAL: Normal.
	//
	// - CEASE: Expired.
	//
	// - RELEASE: Overdue.
	//
	// example:
	//
	// NORMAL
	OrderState *string `json:"OrderState,omitempty" xml:"OrderState,omitempty"`
	// The access configuration for OSS.
	OssInfo *DescribeInstancesResponseBodyInstancesOssInfo `json:"OssInfo,omitempty" xml:"OssInfo,omitempty" type:"Struct"`
	// The region to which the instance belongs.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The time when the instance was created. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1629879567394
	ResourceCreateTime *int64 `json:"ResourceCreateTime,omitempty" xml:"ResourceCreateTime,omitempty"`
	// The time when the instance expires. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1637337600000
	ResourceExpiredTime *int64 `json:"ResourceExpiredTime,omitempty" xml:"ResourceExpiredTime,omitempty"`
	// The resource group.
	//
	// example:
	//
	// rg-***
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// b3690a1655****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource specification information.
	ResourceSpec *DescribeInstancesResponseBodyInstancesResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
	// The storage information.
	Storage *DescribeInstancesResponseBodyInstancesStorage `json:"Storage,omitempty" xml:"Storage,omitempty" type:"Struct"`
	// Specifies whether disaster recovery drills are supported.
	SupportDisasterRecoveryDrill *bool `json:"SupportDisasterRecoveryDrill,omitempty" xml:"SupportDisasterRecoveryDrill,omitempty"`
	// The list of tags.
	Tags []*DescribeInstancesResponseBodyInstancesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the user to which the instance belongs.
	//
	// example:
	//
	// 183899668736****
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// The list of vSwitch IDs.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The primary vSwitch group information.
	VSwitchInfo []*DescribeInstancesResponseBodyInstancesVSwitchInfo `json:"VSwitchInfo,omitempty" xml:"VSwitchInfo,omitempty" type:"Repeated"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-2ze9*******nxfmfcdi
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The VPC information.
	VpcInfo *DescribeInstancesResponseBodyInstancesVpcInfo `json:"VpcInfo,omitempty" xml:"VpcInfo,omitempty" type:"Struct"`
	// The zone ID of the instance.
	//
	// example:
	//
	// cn-beijing-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstances) GetAnsm() *bool {
	return s.Ansm
}

func (s *DescribeInstancesResponseBodyInstances) GetArchitectureType() *string {
	return s.ArchitectureType
}

func (s *DescribeInstancesResponseBodyInstances) GetAskClusterId() *string {
	return s.AskClusterId
}

func (s *DescribeInstancesResponseBodyInstances) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeInstancesResponseBodyInstances) GetClusterState() *DescribeInstancesResponseBodyInstancesClusterState {
	return s.ClusterState
}

func (s *DescribeInstancesResponseBodyInstances) GetClusterStatus() *string {
	return s.ClusterStatus
}

func (s *DescribeInstancesResponseBodyInstances) GetClusterUsedResources() *DescribeInstancesResponseBodyInstancesClusterUsedResources {
	return s.ClusterUsedResources
}

func (s *DescribeInstancesResponseBodyInstances) GetClusterUsedStorage() *DescribeInstancesResponseBodyInstancesClusterUsedStorage {
	return s.ClusterUsedStorage
}

func (s *DescribeInstancesResponseBodyInstances) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *DescribeInstancesResponseBodyInstances) GetElastic() *bool {
	return s.Elastic
}

func (s *DescribeInstancesResponseBodyInstances) GetElasticInstanceId() *string {
	return s.ElasticInstanceId
}

func (s *DescribeInstancesResponseBodyInstances) GetElasticOrderState() *string {
	return s.ElasticOrderState
}

func (s *DescribeInstancesResponseBodyInstances) GetElasticResourceSpec() *DescribeInstancesResponseBodyInstancesElasticResourceSpec {
	return s.ElasticResourceSpec
}

func (s *DescribeInstancesResponseBodyInstances) GetHa() *bool {
	return s.Ha
}

func (s *DescribeInstancesResponseBodyInstances) GetHaResourceSpec() *DescribeInstancesResponseBodyInstancesHaResourceSpec {
	return s.HaResourceSpec
}

func (s *DescribeInstancesResponseBodyInstances) GetHaVSwitchIds() []*string {
	return s.HaVSwitchIds
}

func (s *DescribeInstancesResponseBodyInstances) GetHaVSwitchInfo() []*DescribeInstancesResponseBodyInstancesHaVSwitchInfo {
	return s.HaVSwitchInfo
}

func (s *DescribeInstancesResponseBodyInstances) GetHaZoneId() *string {
	return s.HaZoneId
}

func (s *DescribeInstancesResponseBodyInstances) GetHostAliases() []*DescribeInstancesResponseBodyInstancesHostAliases {
	return s.HostAliases
}

func (s *DescribeInstancesResponseBodyInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstancesResponseBodyInstances) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstancesResponseBodyInstances) GetMonitorType() *string {
	return s.MonitorType
}

func (s *DescribeInstancesResponseBodyInstances) GetOrderState() *string {
	return s.OrderState
}

func (s *DescribeInstancesResponseBodyInstances) GetOssInfo() *DescribeInstancesResponseBodyInstancesOssInfo {
	return s.OssInfo
}

func (s *DescribeInstancesResponseBodyInstances) GetRegion() *string {
	return s.Region
}

func (s *DescribeInstancesResponseBodyInstances) GetResourceCreateTime() *int64 {
	return s.ResourceCreateTime
}

func (s *DescribeInstancesResponseBodyInstances) GetResourceExpiredTime() *int64 {
	return s.ResourceExpiredTime
}

func (s *DescribeInstancesResponseBodyInstances) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInstancesResponseBodyInstances) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeInstancesResponseBodyInstances) GetResourceSpec() *DescribeInstancesResponseBodyInstancesResourceSpec {
	return s.ResourceSpec
}

func (s *DescribeInstancesResponseBodyInstances) GetStorage() *DescribeInstancesResponseBodyInstancesStorage {
	return s.Storage
}

func (s *DescribeInstancesResponseBodyInstances) GetSupportDisasterRecoveryDrill() *bool {
	return s.SupportDisasterRecoveryDrill
}

func (s *DescribeInstancesResponseBodyInstances) GetTags() []*DescribeInstancesResponseBodyInstancesTags {
	return s.Tags
}

func (s *DescribeInstancesResponseBodyInstances) GetUid() *string {
	return s.Uid
}

func (s *DescribeInstancesResponseBodyInstances) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *DescribeInstancesResponseBodyInstances) GetVSwitchInfo() []*DescribeInstancesResponseBodyInstancesVSwitchInfo {
	return s.VSwitchInfo
}

func (s *DescribeInstancesResponseBodyInstances) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeInstancesResponseBodyInstances) GetVpcInfo() *DescribeInstancesResponseBodyInstancesVpcInfo {
	return s.VpcInfo
}

func (s *DescribeInstancesResponseBodyInstances) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeInstancesResponseBodyInstances) SetAnsm(v bool) *DescribeInstancesResponseBodyInstances {
	s.Ansm = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetArchitectureType(v string) *DescribeInstancesResponseBodyInstances {
	s.ArchitectureType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetAskClusterId(v string) *DescribeInstancesResponseBodyInstances {
	s.AskClusterId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetChargeType(v string) *DescribeInstancesResponseBodyInstances {
	s.ChargeType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetClusterState(v *DescribeInstancesResponseBodyInstancesClusterState) *DescribeInstancesResponseBodyInstances {
	s.ClusterState = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetClusterStatus(v string) *DescribeInstancesResponseBodyInstances {
	s.ClusterStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetClusterUsedResources(v *DescribeInstancesResponseBodyInstancesClusterUsedResources) *DescribeInstancesResponseBodyInstances {
	s.ClusterUsedResources = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetClusterUsedStorage(v *DescribeInstancesResponseBodyInstancesClusterUsedStorage) *DescribeInstancesResponseBodyInstances {
	s.ClusterUsedStorage = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetDeletionProtection(v bool) *DescribeInstancesResponseBodyInstances {
	s.DeletionProtection = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetElastic(v bool) *DescribeInstancesResponseBodyInstances {
	s.Elastic = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetElasticInstanceId(v string) *DescribeInstancesResponseBodyInstances {
	s.ElasticInstanceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetElasticOrderState(v string) *DescribeInstancesResponseBodyInstances {
	s.ElasticOrderState = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetElasticResourceSpec(v *DescribeInstancesResponseBodyInstancesElasticResourceSpec) *DescribeInstancesResponseBodyInstances {
	s.ElasticResourceSpec = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetHa(v bool) *DescribeInstancesResponseBodyInstances {
	s.Ha = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetHaResourceSpec(v *DescribeInstancesResponseBodyInstancesHaResourceSpec) *DescribeInstancesResponseBodyInstances {
	s.HaResourceSpec = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetHaVSwitchIds(v []*string) *DescribeInstancesResponseBodyInstances {
	s.HaVSwitchIds = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetHaVSwitchInfo(v []*DescribeInstancesResponseBodyInstancesHaVSwitchInfo) *DescribeInstancesResponseBodyInstances {
	s.HaVSwitchInfo = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetHaZoneId(v string) *DescribeInstancesResponseBodyInstances {
	s.HaZoneId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetHostAliases(v []*DescribeInstancesResponseBodyInstancesHostAliases) *DescribeInstancesResponseBodyInstances {
	s.HostAliases = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceId(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceName(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetMonitorType(v string) *DescribeInstancesResponseBodyInstances {
	s.MonitorType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetOrderState(v string) *DescribeInstancesResponseBodyInstances {
	s.OrderState = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetOssInfo(v *DescribeInstancesResponseBodyInstancesOssInfo) *DescribeInstancesResponseBodyInstances {
	s.OssInfo = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetRegion(v string) *DescribeInstancesResponseBodyInstances {
	s.Region = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetResourceCreateTime(v int64) *DescribeInstancesResponseBodyInstances {
	s.ResourceCreateTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetResourceExpiredTime(v int64) *DescribeInstancesResponseBodyInstances {
	s.ResourceExpiredTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetResourceGroupId(v string) *DescribeInstancesResponseBodyInstances {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetResourceId(v string) *DescribeInstancesResponseBodyInstances {
	s.ResourceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetResourceSpec(v *DescribeInstancesResponseBodyInstancesResourceSpec) *DescribeInstancesResponseBodyInstances {
	s.ResourceSpec = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetStorage(v *DescribeInstancesResponseBodyInstancesStorage) *DescribeInstancesResponseBodyInstances {
	s.Storage = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetSupportDisasterRecoveryDrill(v bool) *DescribeInstancesResponseBodyInstances {
	s.SupportDisasterRecoveryDrill = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetTags(v []*DescribeInstancesResponseBodyInstancesTags) *DescribeInstancesResponseBodyInstances {
	s.Tags = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetUid(v string) *DescribeInstancesResponseBodyInstances {
	s.Uid = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetVSwitchIds(v []*string) *DescribeInstancesResponseBodyInstances {
	s.VSwitchIds = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetVSwitchInfo(v []*DescribeInstancesResponseBodyInstancesVSwitchInfo) *DescribeInstancesResponseBodyInstances {
	s.VSwitchInfo = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetVpcId(v string) *DescribeInstancesResponseBodyInstances {
	s.VpcId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetVpcInfo(v *DescribeInstancesResponseBodyInstancesVpcInfo) *DescribeInstancesResponseBodyInstances {
	s.VpcInfo = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetZoneId(v string) *DescribeInstancesResponseBodyInstances {
	s.ZoneId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) Validate() error {
	if s.ClusterState != nil {
		if err := s.ClusterState.Validate(); err != nil {
			return err
		}
	}
	if s.ClusterUsedResources != nil {
		if err := s.ClusterUsedResources.Validate(); err != nil {
			return err
		}
	}
	if s.ClusterUsedStorage != nil {
		if err := s.ClusterUsedStorage.Validate(); err != nil {
			return err
		}
	}
	if s.ElasticResourceSpec != nil {
		if err := s.ElasticResourceSpec.Validate(); err != nil {
			return err
		}
	}
	if s.HaResourceSpec != nil {
		if err := s.HaResourceSpec.Validate(); err != nil {
			return err
		}
	}
	if s.HaVSwitchInfo != nil {
		for _, item := range s.HaVSwitchInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.HostAliases != nil {
		for _, item := range s.HostAliases {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OssInfo != nil {
		if err := s.OssInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ResourceSpec != nil {
		if err := s.ResourceSpec.Validate(); err != nil {
			return err
		}
	}
	if s.Storage != nil {
		if err := s.Storage.Validate(); err != nil {
			return err
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
	if s.VSwitchInfo != nil {
		for _, item := range s.VSwitchInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VpcInfo != nil {
		if err := s.VpcInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstancesClusterState struct {
	// The cluster ID.
	//
	// example:
	//
	// a6f35c0ff73448
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster stage information.
	ClusterStage *DescribeInstancesResponseBodyInstancesClusterStateClusterStage `json:"ClusterStage,omitempty" xml:"ClusterStage,omitempty" type:"Struct"`
	// Indicates whether the cluster creation timed out.
	//
	// example:
	//
	// true
	CreateTimeout *bool `json:"CreateTimeout,omitempty" xml:"CreateTimeout,omitempty"`
	// The cluster status. Valid values:
	//
	// - CREATING: The cluster is being created.
	//
	// - RUNNING: The cluster is running.
	//
	// - DISABLE: The cluster is invalid.
	//
	// - DELETING: The cluster is being deleted.
	//
	// - DELETED: The cluster is deleted.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The cluster status. Valid values:
	//
	// - CREATING: The cluster is being created.
	//
	// - RUNNING: The cluster is running.
	//
	// - DISABLE: The cluster is invalid.
	//
	// - DELETING: The cluster is being deleted.
	//
	// - DELETED: The cluster is deleted.
	//
	// example:
	//
	// RUNNING
	SubStatus *string `json:"SubStatus,omitempty" xml:"SubStatus,omitempty"`
	// The URLs related to the cluster.
	//
	// example:
	//
	// https://flink-cn-hongkong.data.aliyun.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The user-side load balancing information.
	UserSlbDto *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto `json:"UserSlbDto,omitempty" xml:"UserSlbDto,omitempty" type:"Struct"`
	// The CIDR block of the VPC where the cluster is located.
	//
	// example:
	//
	// 10.0.0.0/16
	VpcCidr *string `json:"VpcCidr,omitempty" xml:"VpcCidr,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesClusterState) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesClusterState) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesClusterState) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeInstancesResponseBodyInstancesClusterState) GetClusterStage() *DescribeInstancesResponseBodyInstancesClusterStateClusterStage {
	return s.ClusterStage
}

func (s *DescribeInstancesResponseBodyInstancesClusterState) GetCreateTimeout() *bool {
	return s.CreateTimeout
}

func (s *DescribeInstancesResponseBodyInstancesClusterState) GetStatus() *string {
	return s.Status
}

func (s *DescribeInstancesResponseBodyInstancesClusterState) GetSubStatus() *string {
	return s.SubStatus
}

func (s *DescribeInstancesResponseBodyInstancesClusterState) GetUrl() *string {
	return s.Url
}

func (s *DescribeInstancesResponseBodyInstancesClusterState) GetUserSlbDto() *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto {
	return s.UserSlbDto
}

func (s *DescribeInstancesResponseBodyInstancesClusterState) GetVpcCidr() *string {
	return s.VpcCidr
}

func (s *DescribeInstancesResponseBodyInstancesClusterState) SetClusterId(v string) *DescribeInstancesResponseBodyInstancesClusterState {
	s.ClusterId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterState) SetClusterStage(v *DescribeInstancesResponseBodyInstancesClusterStateClusterStage) *DescribeInstancesResponseBodyInstancesClusterState {
	s.ClusterStage = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterState) SetCreateTimeout(v bool) *DescribeInstancesResponseBodyInstancesClusterState {
	s.CreateTimeout = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterState) SetStatus(v string) *DescribeInstancesResponseBodyInstancesClusterState {
	s.Status = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterState) SetSubStatus(v string) *DescribeInstancesResponseBodyInstancesClusterState {
	s.SubStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterState) SetUrl(v string) *DescribeInstancesResponseBodyInstancesClusterState {
	s.Url = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterState) SetUserSlbDto(v *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto) *DescribeInstancesResponseBodyInstancesClusterState {
	s.UserSlbDto = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterState) SetVpcCidr(v string) *DescribeInstancesResponseBodyInstancesClusterState {
	s.VpcCidr = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterState) Validate() error {
	if s.ClusterStage != nil {
		if err := s.ClusterStage.Validate(); err != nil {
			return err
		}
	}
	if s.UserSlbDto != nil {
		if err := s.UserSlbDto.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstancesClusterStateClusterStage struct {
	// The cluster ID.
	//
	// example:
	//
	// f0957e82de6e47
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The sequence number of the current stage.
	//
	// example:
	//
	// 0
	CurrentStage *int32 `json:"CurrentStage,omitempty" xml:"CurrentStage,omitempty"`
	// The error message.
	//
	// example:
	//
	// Resource creation failed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The status of the workspace. Valid values:
	//
	// - CREATING: The workspace is being created.
	//
	// - DELETING: The workspace is being deleted.
	//
	// - MODIFYING: The resource specifications are being modified.
	//
	// - SUCCESS: The previous operation was successful.
	//
	// - FAILED: The previous operation failed.
	//
	// example:
	//
	// CREATING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of stages and weights.
	TotalStageWithWeight []*DescribeInstancesResponseBodyInstancesClusterStateClusterStageTotalStageWithWeight `json:"TotalStageWithWeight,omitempty" xml:"TotalStageWithWeight,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesClusterStateClusterStage) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesClusterStateClusterStage) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateClusterStage) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateClusterStage) GetCurrentStage() *int32 {
	return s.CurrentStage
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateClusterStage) GetMessage() *string {
	return s.Message
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateClusterStage) GetStatus() *string {
	return s.Status
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateClusterStage) GetTotalStageWithWeight() []*DescribeInstancesResponseBodyInstancesClusterStateClusterStageTotalStageWithWeight {
	return s.TotalStageWithWeight
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateClusterStage) SetClusterId(v string) *DescribeInstancesResponseBodyInstancesClusterStateClusterStage {
	s.ClusterId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateClusterStage) SetCurrentStage(v int32) *DescribeInstancesResponseBodyInstancesClusterStateClusterStage {
	s.CurrentStage = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateClusterStage) SetMessage(v string) *DescribeInstancesResponseBodyInstancesClusterStateClusterStage {
	s.Message = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateClusterStage) SetStatus(v string) *DescribeInstancesResponseBodyInstancesClusterStateClusterStage {
	s.Status = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateClusterStage) SetTotalStageWithWeight(v []*DescribeInstancesResponseBodyInstancesClusterStateClusterStageTotalStageWithWeight) *DescribeInstancesResponseBodyInstancesClusterStateClusterStage {
	s.TotalStageWithWeight = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateClusterStage) Validate() error {
	if s.TotalStageWithWeight != nil {
		for _, item := range s.TotalStageWithWeight {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstancesClusterStateClusterStageTotalStageWithWeight struct {
	// The stage sequence number.
	//
	// example:
	//
	// 0
	StepIndex *int32 `json:"StepIndex,omitempty" xml:"StepIndex,omitempty"`
	// The stage name.
	//
	// example:
	//
	// Start development console
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
	// The stage weight.
	//
	// example:
	//
	// 0
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesClusterStateClusterStageTotalStageWithWeight) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesClusterStateClusterStageTotalStageWithWeight) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateClusterStageTotalStageWithWeight) GetStepIndex() *int32 {
	return s.StepIndex
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateClusterStageTotalStageWithWeight) GetStepName() *string {
	return s.StepName
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateClusterStageTotalStageWithWeight) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateClusterStageTotalStageWithWeight) SetStepIndex(v int32) *DescribeInstancesResponseBodyInstancesClusterStateClusterStageTotalStageWithWeight {
	s.StepIndex = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateClusterStageTotalStageWithWeight) SetStepName(v string) *DescribeInstancesResponseBodyInstancesClusterStateClusterStageTotalStageWithWeight {
	s.StepName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateClusterStageTotalStageWithWeight) SetWeight(v int32) *DescribeInstancesResponseBodyInstancesClusterStateClusterStageTotalStageWithWeight {
	s.Weight = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateClusterStageTotalStageWithWeight) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto struct {
	// Indicates whether a user-side load balancing instance exists.
	ExistSlb *bool `json:"ExistSlb,omitempty" xml:"ExistSlb,omitempty"`
	// The load balancing instance ID.
	//
	// example:
	//
	// lb-3nss24wchlpaa75xxx
	SlbId *string `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	// The IP address of the load balancing instance.
	//
	// example:
	//
	// 192.168.1.10
	SlbIp *string `json:"SlbIp,omitempty" xml:"SlbIp,omitempty"`
	// The status of the load balancing instance.
	//
	// example:
	//
	// active
	SlbStatus *string `json:"SlbStatus,omitempty" xml:"SlbStatus,omitempty"`
	// The list of user-side load balancing listeners.
	UserSlbListeners []*DescribeInstancesResponseBodyInstancesClusterStateUserSlbDtoUserSlbListeners `json:"UserSlbListeners,omitempty" xml:"UserSlbListeners,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto) GetExistSlb() *bool {
	return s.ExistSlb
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto) GetSlbId() *string {
	return s.SlbId
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto) GetSlbIp() *string {
	return s.SlbIp
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto) GetSlbStatus() *string {
	return s.SlbStatus
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto) GetUserSlbListeners() []*DescribeInstancesResponseBodyInstancesClusterStateUserSlbDtoUserSlbListeners {
	return s.UserSlbListeners
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto) SetExistSlb(v bool) *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto {
	s.ExistSlb = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto) SetSlbId(v string) *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto {
	s.SlbId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto) SetSlbIp(v string) *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto {
	s.SlbIp = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto) SetSlbStatus(v string) *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto {
	s.SlbStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto) SetUserSlbListeners(v []*DescribeInstancesResponseBodyInstancesClusterStateUserSlbDtoUserSlbListeners) *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto {
	s.UserSlbListeners = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto) Validate() error {
	if s.UserSlbListeners != nil {
		for _, item := range s.UserSlbListeners {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstancesClusterStateUserSlbDtoUserSlbListeners struct {
	// The listener status.
	//
	// example:
	//
	// running
	ListenersStatus *string `json:"ListenersStatus,omitempty" xml:"ListenersStatus,omitempty"`
	// The listener port.
	//
	// example:
	//
	// 8080
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesClusterStateUserSlbDtoUserSlbListeners) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesClusterStateUserSlbDtoUserSlbListeners) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDtoUserSlbListeners) GetListenersStatus() *string {
	return s.ListenersStatus
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDtoUserSlbListeners) GetPort() *string {
	return s.Port
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDtoUserSlbListeners) SetListenersStatus(v string) *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDtoUserSlbListeners {
	s.ListenersStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDtoUserSlbListeners) SetPort(v string) *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDtoUserSlbListeners {
	s.Port = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDtoUserSlbListeners) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesClusterUsedResources struct {
	// The cluster ID.
	//
	// example:
	//
	// ea2e9545fdc942
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The number of elastic CPUs used.
	//
	// example:
	//
	// 0
	ElasticUsedCpu *float32 `json:"ElasticUsedCpu,omitempty" xml:"ElasticUsedCpu,omitempty"`
	// The amount of elastic memory used.
	//
	// example:
	//
	// 0
	ElasticUsedMemory *float32 `json:"ElasticUsedMemory,omitempty" xml:"ElasticUsedMemory,omitempty"`
	// The number of elastic resources used.
	//
	// example:
	//
	// 0
	ElasticUsedResource *float32 `json:"ElasticUsedResource,omitempty" xml:"ElasticUsedResource,omitempty"`
	// The number of guaranteed CPUs used.
	//
	// example:
	//
	// 0
	GuaranteedUsedCpu *float32 `json:"GuaranteedUsedCpu,omitempty" xml:"GuaranteedUsedCpu,omitempty"`
	// The amount of guaranteed memory used.
	//
	// example:
	//
	// 0
	GuaranteedUsedMemory *float32 `json:"GuaranteedUsedMemory,omitempty" xml:"GuaranteedUsedMemory,omitempty"`
	// The number of guaranteed resources used.
	//
	// example:
	//
	// 0
	GuaranteedUsedResource *float32 `json:"GuaranteedUsedResource,omitempty" xml:"GuaranteedUsedResource,omitempty"`
	// Indicates whether zone-disaster recovery resources are selected.
	//
	// example:
	//
	// true
	Ha *bool `json:"Ha,omitempty" xml:"Ha,omitempty"`
	// The number of zone-disaster recovery CPUs used.
	//
	// example:
	//
	// 0
	HaUsedCpu *float32 `json:"HaUsedCpu,omitempty" xml:"HaUsedCpu,omitempty"`
	// The amount of zone-disaster recovery memory used.
	//
	// example:
	//
	// 0
	HaUsedMemory *float32 `json:"HaUsedMemory,omitempty" xml:"HaUsedMemory,omitempty"`
	// The number of zone-disaster recovery compute units used.
	//
	// example:
	//
	// 0
	HaUsedResource *float32 `json:"HaUsedResource,omitempty" xml:"HaUsedResource,omitempty"`
	// The total number of CPUs used.
	//
	// example:
	//
	// 0
	UsedCpu *float32 `json:"UsedCpu,omitempty" xml:"UsedCpu,omitempty"`
	// The total amount of memory used.
	//
	// example:
	//
	// 0
	UsedMemory *float32 `json:"UsedMemory,omitempty" xml:"UsedMemory,omitempty"`
	// The total number of compute units used.
	//
	// example:
	//
	// 0
	UsedResource *float32 `json:"UsedResource,omitempty" xml:"UsedResource,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesClusterUsedResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesClusterUsedResources) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) GetElasticUsedCpu() *float32 {
	return s.ElasticUsedCpu
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) GetElasticUsedMemory() *float32 {
	return s.ElasticUsedMemory
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) GetElasticUsedResource() *float32 {
	return s.ElasticUsedResource
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) GetGuaranteedUsedCpu() *float32 {
	return s.GuaranteedUsedCpu
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) GetGuaranteedUsedMemory() *float32 {
	return s.GuaranteedUsedMemory
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) GetGuaranteedUsedResource() *float32 {
	return s.GuaranteedUsedResource
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) GetHa() *bool {
	return s.Ha
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) GetHaUsedCpu() *float32 {
	return s.HaUsedCpu
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) GetHaUsedMemory() *float32 {
	return s.HaUsedMemory
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) GetHaUsedResource() *float32 {
	return s.HaUsedResource
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) GetUsedCpu() *float32 {
	return s.UsedCpu
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) GetUsedMemory() *float32 {
	return s.UsedMemory
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) GetUsedResource() *float32 {
	return s.UsedResource
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) SetClusterId(v string) *DescribeInstancesResponseBodyInstancesClusterUsedResources {
	s.ClusterId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) SetElasticUsedCpu(v float32) *DescribeInstancesResponseBodyInstancesClusterUsedResources {
	s.ElasticUsedCpu = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) SetElasticUsedMemory(v float32) *DescribeInstancesResponseBodyInstancesClusterUsedResources {
	s.ElasticUsedMemory = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) SetElasticUsedResource(v float32) *DescribeInstancesResponseBodyInstancesClusterUsedResources {
	s.ElasticUsedResource = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) SetGuaranteedUsedCpu(v float32) *DescribeInstancesResponseBodyInstancesClusterUsedResources {
	s.GuaranteedUsedCpu = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) SetGuaranteedUsedMemory(v float32) *DescribeInstancesResponseBodyInstancesClusterUsedResources {
	s.GuaranteedUsedMemory = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) SetGuaranteedUsedResource(v float32) *DescribeInstancesResponseBodyInstancesClusterUsedResources {
	s.GuaranteedUsedResource = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) SetHa(v bool) *DescribeInstancesResponseBodyInstancesClusterUsedResources {
	s.Ha = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) SetHaUsedCpu(v float32) *DescribeInstancesResponseBodyInstancesClusterUsedResources {
	s.HaUsedCpu = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) SetHaUsedMemory(v float32) *DescribeInstancesResponseBodyInstancesClusterUsedResources {
	s.HaUsedMemory = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) SetHaUsedResource(v float32) *DescribeInstancesResponseBodyInstancesClusterUsedResources {
	s.HaUsedResource = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) SetUsedCpu(v float32) *DescribeInstancesResponseBodyInstancesClusterUsedResources {
	s.UsedCpu = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) SetUsedMemory(v float32) *DescribeInstancesResponseBodyInstancesClusterUsedResources {
	s.UsedMemory = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) SetUsedResource(v float32) *DescribeInstancesResponseBodyInstancesClusterUsedResources {
	s.UsedResource = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedResources) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesClusterUsedStorage struct {
	// The cluster ID.
	//
	// example:
	//
	// f0957e82de6e47
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The amount of storage used by the cluster.
	//
	// example:
	//
	// 0
	UsedStorage *float32 `json:"UsedStorage,omitempty" xml:"UsedStorage,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesClusterUsedStorage) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesClusterUsedStorage) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedStorage) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedStorage) GetUsedStorage() *float32 {
	return s.UsedStorage
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedStorage) SetClusterId(v string) *DescribeInstancesResponseBodyInstancesClusterUsedStorage {
	s.ClusterId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedStorage) SetUsedStorage(v float32) *DescribeInstancesResponseBodyInstancesClusterUsedStorage {
	s.UsedStorage = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesClusterUsedStorage) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesElasticResourceSpec struct {
	// The number of CPUs.
	//
	// example:
	//
	// 5
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The amount of memory used.
	//
	// example:
	//
	// 280
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesElasticResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesElasticResourceSpec) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesElasticResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *DescribeInstancesResponseBodyInstancesElasticResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *DescribeInstancesResponseBodyInstancesElasticResourceSpec) SetCpu(v int32) *DescribeInstancesResponseBodyInstancesElasticResourceSpec {
	s.Cpu = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesElasticResourceSpec) SetMemoryGB(v int32) *DescribeInstancesResponseBodyInstancesElasticResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesElasticResourceSpec) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesHaResourceSpec struct {
	// The number of zone-disaster recovery CPUs.
	//
	// example:
	//
	// 12
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The memory size of zone-disaster recovery resources.
	//
	// example:
	//
	// 48
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesHaResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesHaResourceSpec) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesHaResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *DescribeInstancesResponseBodyInstancesHaResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *DescribeInstancesResponseBodyInstancesHaResourceSpec) SetCpu(v int32) *DescribeInstancesResponseBodyInstancesHaResourceSpec {
	s.Cpu = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesHaResourceSpec) SetMemoryGB(v int32) *DescribeInstancesResponseBodyInstancesHaResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesHaResourceSpec) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesHaVSwitchInfo struct {
	// The number of available IP addresses in the vSwitch.
	//
	// example:
	//
	// 64
	AvailableIpAddressCount *int64 `json:"AvailableIpAddressCount,omitempty" xml:"AvailableIpAddressCount,omitempty"`
	// The description of the VPC.
	//
	// example:
	//
	// “”
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The vSwitch CIDR block.
	//
	// example:
	//
	// 192.168.16.0/24
	VSwitchCidr *string `json:"VSwitchCidr,omitempty" xml:"VSwitchCidr,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-wz9e9zt38b9f3l6*****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The vSwitch name.
	//
	// example:
	//
	// Test vSwitch
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-uf6d0akr6lzd75v****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID of the workspace.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesHaVSwitchInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesHaVSwitchInfo) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesHaVSwitchInfo) GetAvailableIpAddressCount() *int64 {
	return s.AvailableIpAddressCount
}

func (s *DescribeInstancesResponseBodyInstancesHaVSwitchInfo) GetDescription() *string {
	return s.Description
}

func (s *DescribeInstancesResponseBodyInstancesHaVSwitchInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstancesResponseBodyInstancesHaVSwitchInfo) GetVSwitchCidr() *string {
	return s.VSwitchCidr
}

func (s *DescribeInstancesResponseBodyInstancesHaVSwitchInfo) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeInstancesResponseBodyInstancesHaVSwitchInfo) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *DescribeInstancesResponseBodyInstancesHaVSwitchInfo) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeInstancesResponseBodyInstancesHaVSwitchInfo) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeInstancesResponseBodyInstancesHaVSwitchInfo) SetAvailableIpAddressCount(v int64) *DescribeInstancesResponseBodyInstancesHaVSwitchInfo {
	s.AvailableIpAddressCount = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesHaVSwitchInfo) SetDescription(v string) *DescribeInstancesResponseBodyInstancesHaVSwitchInfo {
	s.Description = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesHaVSwitchInfo) SetRegionId(v string) *DescribeInstancesResponseBodyInstancesHaVSwitchInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesHaVSwitchInfo) SetVSwitchCidr(v string) *DescribeInstancesResponseBodyInstancesHaVSwitchInfo {
	s.VSwitchCidr = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesHaVSwitchInfo) SetVSwitchId(v string) *DescribeInstancesResponseBodyInstancesHaVSwitchInfo {
	s.VSwitchId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesHaVSwitchInfo) SetVSwitchName(v string) *DescribeInstancesResponseBodyInstancesHaVSwitchInfo {
	s.VSwitchName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesHaVSwitchInfo) SetVpcId(v string) *DescribeInstancesResponseBodyInstancesHaVSwitchInfo {
	s.VpcId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesHaVSwitchInfo) SetZoneId(v string) *DescribeInstancesResponseBodyInstancesHaVSwitchInfo {
	s.ZoneId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesHaVSwitchInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesHostAliases struct {
	// The list of hostnames.
	//
	// This parameter is required.
	HostNames []*string `json:"HostNames,omitempty" xml:"HostNames,omitempty" type:"Repeated"`
	// The IP address that corresponds to the hostname.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.1.20
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesHostAliases) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesHostAliases) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesHostAliases) GetHostNames() []*string {
	return s.HostNames
}

func (s *DescribeInstancesResponseBodyInstancesHostAliases) GetIp() *string {
	return s.Ip
}

func (s *DescribeInstancesResponseBodyInstancesHostAliases) SetHostNames(v []*string) *DescribeInstancesResponseBodyInstancesHostAliases {
	s.HostNames = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesHostAliases) SetIp(v string) *DescribeInstancesResponseBodyInstancesHostAliases {
	s.Ip = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesHostAliases) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesOssInfo struct {
	// The AccessKey ID for OSS.
	//
	// example:
	//
	// XXXXXX
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// The AccessKey secret for OSS.
	//
	// example:
	//
	// XXXXXXXX
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The name of the bound OSS bucket.
	//
	// example:
	//
	// flink-log-john
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The versioning status of the OSS bucket.
	//
	// example:
	//
	// OFF
	BucketVersioningStatus *string `json:"BucketVersioningStatus,omitempty" xml:"BucketVersioningStatus,omitempty"`
	// The OSS endpoint.
	//
	// example:
	//
	// oss-cn-hangzhou-internal.aliyuncs.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesOssInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesOssInfo) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesOssInfo) GetAccessId() *string {
	return s.AccessId
}

func (s *DescribeInstancesResponseBodyInstancesOssInfo) GetAccessKey() *string {
	return s.AccessKey
}

func (s *DescribeInstancesResponseBodyInstancesOssInfo) GetBucket() *string {
	return s.Bucket
}

func (s *DescribeInstancesResponseBodyInstancesOssInfo) GetBucketVersioningStatus() *string {
	return s.BucketVersioningStatus
}

func (s *DescribeInstancesResponseBodyInstancesOssInfo) GetEndpoint() *string {
	return s.Endpoint
}

func (s *DescribeInstancesResponseBodyInstancesOssInfo) SetAccessId(v string) *DescribeInstancesResponseBodyInstancesOssInfo {
	s.AccessId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesOssInfo) SetAccessKey(v string) *DescribeInstancesResponseBodyInstancesOssInfo {
	s.AccessKey = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesOssInfo) SetBucket(v string) *DescribeInstancesResponseBodyInstancesOssInfo {
	s.Bucket = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesOssInfo) SetBucketVersioningStatus(v string) *DescribeInstancesResponseBodyInstancesOssInfo {
	s.BucketVersioningStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesOssInfo) SetEndpoint(v string) *DescribeInstancesResponseBodyInstancesOssInfo {
	s.Endpoint = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesOssInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesResourceSpec struct {
	// The number of CPUs.
	//
	// example:
	//
	// 10
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The configured memory size.
	//
	// example:
	//
	// 40
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesResourceSpec) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *DescribeInstancesResponseBodyInstancesResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *DescribeInstancesResponseBodyInstancesResourceSpec) SetCpu(v int32) *DescribeInstancesResponseBodyInstancesResourceSpec {
	s.Cpu = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesResourceSpec) SetMemoryGB(v int32) *DescribeInstancesResponseBodyInstancesResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesResourceSpec) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesStorage struct {
	// Specifies whether fully managed storage is selected. Valid values:
	//
	// - true: Fully managed storage is used.
	//
	// - false: Fully managed storage is not used.
	//
	// example:
	//
	// true
	FullyManaged *bool `json:"FullyManaged,omitempty" xml:"FullyManaged,omitempty"`
	// The order status of fully managed storage. Valid values:
	//
	// - NOT_INIT: The order is placed but components are not deployed.
	//
	// - NORMAL: Normal.
	//
	// - CEASE: Expired.
	//
	// - RELEASE: Overdue.
	//
	// example:
	//
	// NORMAL
	OrderState *string `json:"OrderState,omitempty" xml:"OrderState,omitempty"`
	// The OSS storage information.
	Oss *DescribeInstancesResponseBodyInstancesStorageOss `json:"Oss,omitempty" xml:"Oss,omitempty" type:"Struct"`
	// Specifies whether fully managed storage can be created.
	//
	// example:
	//
	// true
	SupportCreateFullyManagedStorage *bool `json:"SupportCreateFullyManagedStorage,omitempty" xml:"SupportCreateFullyManagedStorage,omitempty"`
	// Specifies whether migration progress detection is supported.
	//
	// example:
	//
	// true
	SupportMigrationProgressDetection *bool `json:"SupportMigrationProgressDetection,omitempty" xml:"SupportMigrationProgressDetection,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesStorage) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesStorage) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesStorage) GetFullyManaged() *bool {
	return s.FullyManaged
}

func (s *DescribeInstancesResponseBodyInstancesStorage) GetOrderState() *string {
	return s.OrderState
}

func (s *DescribeInstancesResponseBodyInstancesStorage) GetOss() *DescribeInstancesResponseBodyInstancesStorageOss {
	return s.Oss
}

func (s *DescribeInstancesResponseBodyInstancesStorage) GetSupportCreateFullyManagedStorage() *bool {
	return s.SupportCreateFullyManagedStorage
}

func (s *DescribeInstancesResponseBodyInstancesStorage) GetSupportMigrationProgressDetection() *bool {
	return s.SupportMigrationProgressDetection
}

func (s *DescribeInstancesResponseBodyInstancesStorage) SetFullyManaged(v bool) *DescribeInstancesResponseBodyInstancesStorage {
	s.FullyManaged = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesStorage) SetOrderState(v string) *DescribeInstancesResponseBodyInstancesStorage {
	s.OrderState = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesStorage) SetOss(v *DescribeInstancesResponseBodyInstancesStorageOss) *DescribeInstancesResponseBodyInstancesStorage {
	s.Oss = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesStorage) SetSupportCreateFullyManagedStorage(v bool) *DescribeInstancesResponseBodyInstancesStorage {
	s.SupportCreateFullyManagedStorage = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesStorage) SetSupportMigrationProgressDetection(v bool) *DescribeInstancesResponseBodyInstancesStorage {
	s.SupportMigrationProgressDetection = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesStorage) Validate() error {
	if s.Oss != nil {
		if err := s.Oss.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstancesStorageOss struct {
	// The name of the bound OSS bucket.
	//
	// example:
	//
	// oss_flink
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesStorageOss) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesStorageOss) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesStorageOss) GetBucket() *string {
	return s.Bucket
}

func (s *DescribeInstancesResponseBodyInstancesStorageOss) SetBucket(v string) *DescribeInstancesResponseBodyInstancesStorageOss {
	s.Bucket = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesStorageOss) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesTags struct {
	// The tag key.
	//
	// example:
	//
	// flink
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesTags) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesTags) GetKey() *string {
	return s.Key
}

func (s *DescribeInstancesResponseBodyInstancesTags) GetValue() *string {
	return s.Value
}

func (s *DescribeInstancesResponseBodyInstancesTags) SetKey(v string) *DescribeInstancesResponseBodyInstancesTags {
	s.Key = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesTags) SetValue(v string) *DescribeInstancesResponseBodyInstancesTags {
	s.Value = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesTags) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesVSwitchInfo struct {
	// The number of available IP addresses in the vSwitch.
	//
	// example:
	//
	// 256
	AvailableIpAddressCount *string `json:"AvailableIpAddressCount,omitempty" xml:"AvailableIpAddressCount,omitempty"`
	// The description of the vSwitch group.
	//
	// example:
	//
	// “”
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The vSwitch CIDR block.
	//
	// example:
	//
	// 192.168.16.0/24
	VSwitchCidr *string `json:"VSwitchCidr,omitempty" xml:"VSwitchCidr,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-8vb6jk75wfcwnuq****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The name of the vSwitch.
	//
	// example:
	//
	// “”
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp16h9ru9gvz2tygo*****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID of the workspace.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesVSwitchInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesVSwitchInfo) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesVSwitchInfo) GetAvailableIpAddressCount() *string {
	return s.AvailableIpAddressCount
}

func (s *DescribeInstancesResponseBodyInstancesVSwitchInfo) GetDescription() *string {
	return s.Description
}

func (s *DescribeInstancesResponseBodyInstancesVSwitchInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstancesResponseBodyInstancesVSwitchInfo) GetVSwitchCidr() *string {
	return s.VSwitchCidr
}

func (s *DescribeInstancesResponseBodyInstancesVSwitchInfo) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeInstancesResponseBodyInstancesVSwitchInfo) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *DescribeInstancesResponseBodyInstancesVSwitchInfo) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeInstancesResponseBodyInstancesVSwitchInfo) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeInstancesResponseBodyInstancesVSwitchInfo) SetAvailableIpAddressCount(v string) *DescribeInstancesResponseBodyInstancesVSwitchInfo {
	s.AvailableIpAddressCount = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesVSwitchInfo) SetDescription(v string) *DescribeInstancesResponseBodyInstancesVSwitchInfo {
	s.Description = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesVSwitchInfo) SetRegionId(v string) *DescribeInstancesResponseBodyInstancesVSwitchInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesVSwitchInfo) SetVSwitchCidr(v string) *DescribeInstancesResponseBodyInstancesVSwitchInfo {
	s.VSwitchCidr = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesVSwitchInfo) SetVSwitchId(v string) *DescribeInstancesResponseBodyInstancesVSwitchInfo {
	s.VSwitchId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesVSwitchInfo) SetVSwitchName(v string) *DescribeInstancesResponseBodyInstancesVSwitchInfo {
	s.VSwitchName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesVSwitchInfo) SetVpcId(v string) *DescribeInstancesResponseBodyInstancesVSwitchInfo {
	s.VpcId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesVSwitchInfo) SetZoneId(v string) *DescribeInstancesResponseBodyInstancesVSwitchInfo {
	s.ZoneId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesVSwitchInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstancesVpcInfo struct {
	// The VPC CIDR block.
	//
	// example:
	//
	// 172.16.5.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The VPC description.
	//
	// example:
	//
	// “”
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The VPC status. Valid values:
	//
	// - Pending: Configuring.
	//
	// - Available: Available.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-gw8c6nhjta3eq12y****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The VPC name.
	//
	// example:
	//
	// Flink cluster VPC
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesVpcInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesVpcInfo) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesVpcInfo) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeInstancesResponseBodyInstancesVpcInfo) GetDescription() *string {
	return s.Description
}

func (s *DescribeInstancesResponseBodyInstancesVpcInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstancesResponseBodyInstancesVpcInfo) GetStatus() *string {
	return s.Status
}

func (s *DescribeInstancesResponseBodyInstancesVpcInfo) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeInstancesResponseBodyInstancesVpcInfo) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeInstancesResponseBodyInstancesVpcInfo) SetCidrBlock(v string) *DescribeInstancesResponseBodyInstancesVpcInfo {
	s.CidrBlock = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesVpcInfo) SetDescription(v string) *DescribeInstancesResponseBodyInstancesVpcInfo {
	s.Description = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesVpcInfo) SetRegionId(v string) *DescribeInstancesResponseBodyInstancesVpcInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesVpcInfo) SetStatus(v string) *DescribeInstancesResponseBodyInstancesVpcInfo {
	s.Status = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesVpcInfo) SetVpcId(v string) *DescribeInstancesResponseBodyInstancesVpcInfo {
	s.VpcId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesVpcInfo) SetVpcName(v string) *DescribeInstancesResponseBodyInstancesVpcInfo {
	s.VpcName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesVpcInfo) Validate() error {
	return dara.Validate(s)
}
