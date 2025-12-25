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
	Instances []*DescribeInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// C8DF2A5B-6FBA-5651-A3D4-960F36640E6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 15
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	Ansm             *bool   `json:"Ansm,omitempty" xml:"Ansm,omitempty"`
	ArchitectureType *string `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
	AskClusterId     *string `json:"AskClusterId,omitempty" xml:"AskClusterId,omitempty"`
	// example:
	//
	// PRE
	ChargeType   *string                                             `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ClusterState *DescribeInstancesResponseBodyInstancesClusterState `json:"ClusterState,omitempty" xml:"ClusterState,omitempty" type:"Struct"`
	// example:
	//
	// RUNNING
	ClusterStatus        *string                                                     `json:"ClusterStatus,omitempty" xml:"ClusterStatus,omitempty"`
	ClusterUsedResources *DescribeInstancesResponseBodyInstancesClusterUsedResources `json:"ClusterUsedResources,omitempty" xml:"ClusterUsedResources,omitempty" type:"Struct"`
	ClusterUsedStorage   *DescribeInstancesResponseBodyInstancesClusterUsedStorage   `json:"ClusterUsedStorage,omitempty" xml:"ClusterUsedStorage,omitempty" type:"Struct"`
	Elastic              *bool                                                       `json:"Elastic,omitempty" xml:"Elastic,omitempty"`
	// example:
	//
	// f-cn-e3afbd321
	ElasticInstanceId   *string                                                    `json:"ElasticInstanceId,omitempty" xml:"ElasticInstanceId,omitempty"`
	ElasticOrderState   *string                                                    `json:"ElasticOrderState,omitempty" xml:"ElasticOrderState,omitempty"`
	ElasticResourceSpec *DescribeInstancesResponseBodyInstancesElasticResourceSpec `json:"ElasticResourceSpec,omitempty" xml:"ElasticResourceSpec,omitempty" type:"Struct"`
	Ha                  *bool                                                      `json:"Ha,omitempty" xml:"Ha,omitempty"`
	HaResourceSpec      *DescribeInstancesResponseBodyInstancesHaResourceSpec      `json:"HaResourceSpec,omitempty" xml:"HaResourceSpec,omitempty" type:"Struct"`
	HaVSwitchIds        []*string                                                  `json:"HaVSwitchIds,omitempty" xml:"HaVSwitchIds,omitempty" type:"Repeated"`
	HaVSwitchInfo       []*DescribeInstancesResponseBodyInstancesHaVSwitchInfo     `json:"HaVSwitchInfo,omitempty" xml:"HaVSwitchInfo,omitempty" type:"Repeated"`
	HaZoneId            *string                                                    `json:"HaZoneId,omitempty" xml:"HaZoneId,omitempty"`
	// This parameter is required.
	HostAliases []*DescribeInstancesResponseBodyInstancesHostAliases `json:"HostAliases,omitempty" xml:"HostAliases,omitempty" type:"Repeated"`
	// example:
	//
	// f-cn-zvp2q0zik06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// vvp1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	MonitorType  *string `json:"MonitorType,omitempty" xml:"MonitorType,omitempty"`
	// example:
	//
	// NORMAL
	OrderState *string                                        `json:"OrderState,omitempty" xml:"OrderState,omitempty"`
	OssInfo    *DescribeInstancesResponseBodyInstancesOssInfo `json:"OssInfo,omitempty" xml:"OssInfo,omitempty" type:"Struct"`
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 1629879567394
	ResourceCreateTime *int64 `json:"ResourceCreateTime,omitempty" xml:"ResourceCreateTime,omitempty"`
	// example:
	//
	// 1637337600000
	ResourceExpiredTime *int64  `json:"ResourceExpiredTime,omitempty" xml:"ResourceExpiredTime,omitempty"`
	ResourceGroupId     *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// b3690a1655da47
	ResourceId   *string                                             `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceSpec *DescribeInstancesResponseBodyInstancesResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
	Storage      *DescribeInstancesResponseBodyInstancesStorage      `json:"Storage,omitempty" xml:"Storage,omitempty" type:"Struct"`
	Tags         []*DescribeInstancesResponseBodyInstancesTags       `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// 1838996687368452
	Uid         *string                                              `json:"Uid,omitempty" xml:"Uid,omitempty"`
	VSwitchIds  []*string                                            `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	VSwitchInfo []*DescribeInstancesResponseBodyInstancesVSwitchInfo `json:"VSwitchInfo,omitempty" xml:"VSwitchInfo,omitempty" type:"Repeated"`
	// example:
	//
	// vpc-2ze9*******nxfmfcdi
	VpcId   *string                                        `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcInfo *DescribeInstancesResponseBodyInstancesVpcInfo `json:"VpcInfo,omitempty" xml:"VpcInfo,omitempty" type:"Struct"`
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
	ClusterId     *string                                                         `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterStage  *DescribeInstancesResponseBodyInstancesClusterStateClusterStage `json:"ClusterStage,omitempty" xml:"ClusterStage,omitempty" type:"Struct"`
	CreateTimeout *bool                                                           `json:"CreateTimeout,omitempty" xml:"CreateTimeout,omitempty"`
	Status        *string                                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	SubStatus     *string                                                         `json:"SubStatus,omitempty" xml:"SubStatus,omitempty"`
	Url           *string                                                         `json:"Url,omitempty" xml:"Url,omitempty"`
	UserSlbDto    *DescribeInstancesResponseBodyInstancesClusterStateUserSlbDto   `json:"UserSlbDto,omitempty" xml:"UserSlbDto,omitempty" type:"Struct"`
	VpcCidr       *string                                                         `json:"VpcCidr,omitempty" xml:"VpcCidr,omitempty"`
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
	ClusterId            *string                                                                               `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CurrentStage         *int32                                                                                `json:"CurrentStage,omitempty" xml:"CurrentStage,omitempty"`
	Message              *string                                                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	Status               *string                                                                               `json:"Status,omitempty" xml:"Status,omitempty"`
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
	StepIndex *int32  `json:"StepIndex,omitempty" xml:"StepIndex,omitempty"`
	StepName  *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
	Weight    *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
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
	ExistSlb         *bool                                                                           `json:"ExistSlb,omitempty" xml:"ExistSlb,omitempty"`
	SlbId            *string                                                                         `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	SlbIp            *string                                                                         `json:"SlbIp,omitempty" xml:"SlbIp,omitempty"`
	SlbStatus        *string                                                                         `json:"SlbStatus,omitempty" xml:"SlbStatus,omitempty"`
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
	ListenersStatus *string `json:"ListenersStatus,omitempty" xml:"ListenersStatus,omitempty"`
	Port            *string `json:"Port,omitempty" xml:"Port,omitempty"`
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
	ClusterId              *string  `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ElasticUsedCpu         *float32 `json:"ElasticUsedCpu,omitempty" xml:"ElasticUsedCpu,omitempty"`
	ElasticUsedMemory      *float32 `json:"ElasticUsedMemory,omitempty" xml:"ElasticUsedMemory,omitempty"`
	ElasticUsedResource    *float32 `json:"ElasticUsedResource,omitempty" xml:"ElasticUsedResource,omitempty"`
	GuaranteedUsedCpu      *float32 `json:"GuaranteedUsedCpu,omitempty" xml:"GuaranteedUsedCpu,omitempty"`
	GuaranteedUsedMemory   *float32 `json:"GuaranteedUsedMemory,omitempty" xml:"GuaranteedUsedMemory,omitempty"`
	GuaranteedUsedResource *float32 `json:"GuaranteedUsedResource,omitempty" xml:"GuaranteedUsedResource,omitempty"`
	Ha                     *bool    `json:"Ha,omitempty" xml:"Ha,omitempty"`
	HaUsedCpu              *float32 `json:"HaUsedCpu,omitempty" xml:"HaUsedCpu,omitempty"`
	HaUsedMemory           *float32 `json:"HaUsedMemory,omitempty" xml:"HaUsedMemory,omitempty"`
	HaUsedResource         *float32 `json:"HaUsedResource,omitempty" xml:"HaUsedResource,omitempty"`
	UsedCpu                *float32 `json:"UsedCpu,omitempty" xml:"UsedCpu,omitempty"`
	UsedMemory             *float32 `json:"UsedMemory,omitempty" xml:"UsedMemory,omitempty"`
	UsedResource           *float32 `json:"UsedResource,omitempty" xml:"UsedResource,omitempty"`
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
	ClusterId   *string  `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
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
	Cpu      *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
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
	Cpu      *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
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
	AvailableIpAddressCount *int64  `json:"AvailableIpAddressCount,omitempty" xml:"AvailableIpAddressCount,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VSwitchCidr             *string `json:"VSwitchCidr,omitempty" xml:"VSwitchCidr,omitempty"`
	VSwitchId               *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VSwitchName             *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	VpcId                   *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId                  *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	// This parameter is required.
	HostNames []*string `json:"HostNames,omitempty" xml:"HostNames,omitempty" type:"Repeated"`
	// This parameter is required.
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
	AccessId               *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	AccessKey              *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	Bucket                 *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	BucketVersioningStatus *string `json:"BucketVersioningStatus,omitempty" xml:"BucketVersioningStatus,omitempty"`
	Endpoint               *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
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
	// example:
	//
	// 10
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
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
	FullyManaged                      *bool                                             `json:"FullyManaged,omitempty" xml:"FullyManaged,omitempty"`
	OrderState                        *string                                           `json:"OrderState,omitempty" xml:"OrderState,omitempty"`
	Oss                               *DescribeInstancesResponseBodyInstancesStorageOss `json:"Oss,omitempty" xml:"Oss,omitempty" type:"Struct"`
	SupportCreateFullyManagedStorage  *bool                                             `json:"SupportCreateFullyManagedStorage,omitempty" xml:"SupportCreateFullyManagedStorage,omitempty"`
	SupportMigrationProgressDetection *bool                                             `json:"SupportMigrationProgressDetection,omitempty" xml:"SupportMigrationProgressDetection,omitempty"`
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
	// example:
	//
	// flink
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	AvailableIpAddressCount *string `json:"AvailableIpAddressCount,omitempty" xml:"AvailableIpAddressCount,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VSwitchCidr             *string `json:"VSwitchCidr,omitempty" xml:"VSwitchCidr,omitempty"`
	VSwitchId               *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VSwitchName             *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	VpcId                   *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId                  *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	CidrBlock   *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VpcId       *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcName     *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
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
