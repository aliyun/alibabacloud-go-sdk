// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusters(v []*ListClustersResponseBodyClusters) *ListClustersResponseBody
	GetClusters() []*ListClustersResponseBodyClusters
	SetPageNumber(v string) *ListClustersResponseBody
	GetPageNumber() *string
	SetPageSize(v int32) *ListClustersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListClustersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListClustersResponseBody
	GetTotalCount() *int32
}

type ListClustersResponseBody struct {
	// The list of clusters.
	Clusters []*ListClustersResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBody) GetClusters() []*ListClustersResponseBodyClusters {
	return s.Clusters
}

func (s *ListClustersResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListClustersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClustersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListClustersResponseBody) SetClusters(v []*ListClustersResponseBodyClusters) *ListClustersResponseBody {
	s.Clusters = v
	return s
}

func (s *ListClustersResponseBody) SetPageNumber(v string) *ListClustersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListClustersResponseBody) SetPageSize(v int32) *ListClustersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListClustersResponseBody) SetRequestId(v string) *ListClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClustersResponseBody) SetTotalCount(v int32) *ListClustersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListClustersResponseBody) Validate() error {
	if s.Clusters != nil {
		for _, item := range s.Clusters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListClustersResponseBodyClusters struct {
	// The information about installed software in the cluster.
	AdditionalPackages []*ListClustersResponseBodyClustersAdditionalPackages `json:"AdditionalPackages,omitempty" xml:"AdditionalPackages,omitempty" type:"Repeated"`
	// The information about the addons in the cluster.
	Addons []*ListClustersResponseBodyClustersAddons `json:"Addons,omitempty" xml:"Addons,omitempty" type:"Repeated"`
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
	// The logon credential type of the cluster. Valid values:
	//
	// 	- password: requires passwords for logons.
	//
	// 	- keypair: requires key pairs for logons.
	ClusterCredentials []*string `json:"ClusterCredentials,omitempty" xml:"ClusterCredentials,omitempty" type:"Repeated"`
	// The post-processing script used by the cluster.
	ClusterCustomConfiguration *ListClustersResponseBodyClustersClusterCustomConfiguration `json:"ClusterCustomConfiguration,omitempty" xml:"ClusterCustomConfiguration,omitempty" type:"Struct"`
	// The cluster description.
	//
	// example:
	//
	// Demo
	ClusterDescription *string `json:"ClusterDescription,omitempty" xml:"ClusterDescription,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-VMKe******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The deployment type of the cluster. Valid values:
	//
	// 	- Integrated: public cloud
	//
	// 	- Hybrid: hybrid cloud
	//
	// 	- Custom: a custom cluster
	//
	// example:
	//
	// Integrated
	ClusterMode *string `json:"ClusterMode,omitempty" xml:"ClusterMode,omitempty"`
	// The time when the cluster was modified. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mmZ format. The time is displayed in UTC. For more information, see [ISO 8601](https://help.aliyun.com/document_detail/25696.html).
	//
	// example:
	//
	// 2024-08-06T12:43:01.000Z
	ClusterModifyTime *string `json:"ClusterModifyTime,omitempty" xml:"ClusterModifyTime,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// slurm22.05.8-cluster-20240227
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
	// 	- Releasing: The cluster is being released.
	//
	// 	- stopping: The cluster is being stopped.
	//
	// 	- stopped: The cluster is stopped.
	//
	// 	- exception: The cluster has run into an exception.
	//
	// 	- pending: The cluster is waiting to be configured.
	//
	// example:
	//
	// running
	ClusterStatus *string `json:"ClusterStatus,omitempty" xml:"ClusterStatus,omitempty"`
	// The vCPU-hour usage of the cluster.
	//
	// example:
	//
	// 1000
	ClusterUsedCoreTime *float32 `json:"ClusterUsedCoreTime,omitempty" xml:"ClusterUsedCoreTime,omitempty"`
	// The ID of the vSwitch used by the cluster.
	//
	// example:
	//
	// vsw-f8za5p0mwzgdu3wgx****
	ClusterVSwitchId *string `json:"ClusterVSwitchId,omitempty" xml:"ClusterVSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC) used by the cluster.
	//
	// example:
	//
	// vpc-m5efjevmclc0xdmys****
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
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The Elastic High Performance Computing (E-HPC) version.
	//
	// example:
	//
	// 2.0.0
	EhpcVersion *string `json:"EhpcVersion,omitempty" xml:"EhpcVersion,omitempty"`
	// The configurations of the cluster management node.
	Manager *ListClustersResponseBodyClustersManager `json:"Manager,omitempty" xml:"Manager,omitempty" type:"Struct"`
	// The maximum total number of vCPUs used by the compute nodes that can be managed by the cluster.
	//
	// example:
	//
	// 10000
	MaxCoreCount *int64 `json:"MaxCoreCount,omitempty" xml:"MaxCoreCount,omitempty"`
	// The maximum number of compute nodes that can be managed by the cluster.
	//
	// example:
	//
	// 500
	MaxCount *int64 `json:"MaxCount,omitempty" xml:"MaxCount,omitempty"`
	// The node statistics of the cluster.
	Nodes *ListClustersResponseBodyClustersNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Struct"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxazb4ph****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the security group used by the cluster.
	//
	// example:
	//
	// sg-bp13n61xsydodfyg****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The user attribute information of the cluster.
	Users *ListClustersResponseBodyClustersUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
}

func (s ListClustersResponseBodyClusters) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClusters) GetAdditionalPackages() []*ListClustersResponseBodyClustersAdditionalPackages {
	return s.AdditionalPackages
}

func (s *ListClustersResponseBodyClusters) GetAddons() []*ListClustersResponseBodyClustersAddons {
	return s.Addons
}

func (s *ListClustersResponseBodyClusters) GetClusterCategory() *string {
	return s.ClusterCategory
}

func (s *ListClustersResponseBodyClusters) GetClusterCreateTime() *string {
	return s.ClusterCreateTime
}

func (s *ListClustersResponseBodyClusters) GetClusterCredentials() []*string {
	return s.ClusterCredentials
}

func (s *ListClustersResponseBodyClusters) GetClusterCustomConfiguration() *ListClustersResponseBodyClustersClusterCustomConfiguration {
	return s.ClusterCustomConfiguration
}

func (s *ListClustersResponseBodyClusters) GetClusterDescription() *string {
	return s.ClusterDescription
}

func (s *ListClustersResponseBodyClusters) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListClustersResponseBodyClusters) GetClusterMode() *string {
	return s.ClusterMode
}

func (s *ListClustersResponseBodyClusters) GetClusterModifyTime() *string {
	return s.ClusterModifyTime
}

func (s *ListClustersResponseBodyClusters) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListClustersResponseBodyClusters) GetClusterStatus() *string {
	return s.ClusterStatus
}

func (s *ListClustersResponseBodyClusters) GetClusterUsedCoreTime() *float32 {
	return s.ClusterUsedCoreTime
}

func (s *ListClustersResponseBodyClusters) GetClusterVSwitchId() *string {
	return s.ClusterVSwitchId
}

func (s *ListClustersResponseBodyClusters) GetClusterVpcId() *string {
	return s.ClusterVpcId
}

func (s *ListClustersResponseBodyClusters) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *ListClustersResponseBodyClusters) GetEhpcVersion() *string {
	return s.EhpcVersion
}

func (s *ListClustersResponseBodyClusters) GetManager() *ListClustersResponseBodyClustersManager {
	return s.Manager
}

func (s *ListClustersResponseBodyClusters) GetMaxCoreCount() *int64 {
	return s.MaxCoreCount
}

func (s *ListClustersResponseBodyClusters) GetMaxCount() *int64 {
	return s.MaxCount
}

func (s *ListClustersResponseBodyClusters) GetNodes() *ListClustersResponseBodyClustersNodes {
	return s.Nodes
}

func (s *ListClustersResponseBodyClusters) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListClustersResponseBodyClusters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListClustersResponseBodyClusters) GetUsers() *ListClustersResponseBodyClustersUsers {
	return s.Users
}

func (s *ListClustersResponseBodyClusters) SetAdditionalPackages(v []*ListClustersResponseBodyClustersAdditionalPackages) *ListClustersResponseBodyClusters {
	s.AdditionalPackages = v
	return s
}

func (s *ListClustersResponseBodyClusters) SetAddons(v []*ListClustersResponseBodyClustersAddons) *ListClustersResponseBodyClusters {
	s.Addons = v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterCategory(v string) *ListClustersResponseBodyClusters {
	s.ClusterCategory = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterCreateTime(v string) *ListClustersResponseBodyClusters {
	s.ClusterCreateTime = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterCredentials(v []*string) *ListClustersResponseBodyClusters {
	s.ClusterCredentials = v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterCustomConfiguration(v *ListClustersResponseBodyClustersClusterCustomConfiguration) *ListClustersResponseBodyClusters {
	s.ClusterCustomConfiguration = v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterDescription(v string) *ListClustersResponseBodyClusters {
	s.ClusterDescription = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterId(v string) *ListClustersResponseBodyClusters {
	s.ClusterId = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterMode(v string) *ListClustersResponseBodyClusters {
	s.ClusterMode = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterModifyTime(v string) *ListClustersResponseBodyClusters {
	s.ClusterModifyTime = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterName(v string) *ListClustersResponseBodyClusters {
	s.ClusterName = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterStatus(v string) *ListClustersResponseBodyClusters {
	s.ClusterStatus = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterUsedCoreTime(v float32) *ListClustersResponseBodyClusters {
	s.ClusterUsedCoreTime = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterVSwitchId(v string) *ListClustersResponseBodyClusters {
	s.ClusterVSwitchId = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterVpcId(v string) *ListClustersResponseBodyClusters {
	s.ClusterVpcId = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetDeletionProtection(v bool) *ListClustersResponseBodyClusters {
	s.DeletionProtection = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetEhpcVersion(v string) *ListClustersResponseBodyClusters {
	s.EhpcVersion = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetManager(v *ListClustersResponseBodyClustersManager) *ListClustersResponseBodyClusters {
	s.Manager = v
	return s
}

func (s *ListClustersResponseBodyClusters) SetMaxCoreCount(v int64) *ListClustersResponseBodyClusters {
	s.MaxCoreCount = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetMaxCount(v int64) *ListClustersResponseBodyClusters {
	s.MaxCount = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetNodes(v *ListClustersResponseBodyClustersNodes) *ListClustersResponseBodyClusters {
	s.Nodes = v
	return s
}

func (s *ListClustersResponseBodyClusters) SetResourceGroupId(v string) *ListClustersResponseBodyClusters {
	s.ResourceGroupId = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetSecurityGroupId(v string) *ListClustersResponseBodyClusters {
	s.SecurityGroupId = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetUsers(v *ListClustersResponseBodyClustersUsers) *ListClustersResponseBodyClusters {
	s.Users = v
	return s
}

func (s *ListClustersResponseBodyClusters) Validate() error {
	if s.AdditionalPackages != nil {
		for _, item := range s.AdditionalPackages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Addons != nil {
		for _, item := range s.Addons {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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
	if s.Nodes != nil {
		if err := s.Nodes.Validate(); err != nil {
			return err
		}
	}
	if s.Users != nil {
		if err := s.Users.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListClustersResponseBodyClustersAdditionalPackages struct {
	// The software name.
	//
	// example:
	//
	// gromacs
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The software version.
	//
	// example:
	//
	// 2024.1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListClustersResponseBodyClustersAdditionalPackages) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBodyClustersAdditionalPackages) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersAdditionalPackages) GetName() *string {
	return s.Name
}

func (s *ListClustersResponseBodyClustersAdditionalPackages) GetVersion() *string {
	return s.Version
}

func (s *ListClustersResponseBodyClustersAdditionalPackages) SetName(v string) *ListClustersResponseBodyClustersAdditionalPackages {
	s.Name = &v
	return s
}

func (s *ListClustersResponseBodyClustersAdditionalPackages) SetVersion(v string) *ListClustersResponseBodyClustersAdditionalPackages {
	s.Version = &v
	return s
}

func (s *ListClustersResponseBodyClustersAdditionalPackages) Validate() error {
	return dara.Validate(s)
}

type ListClustersResponseBodyClustersAddons struct {
	// The addon ID.
	//
	// example:
	//
	// Login-1.0-W2g****
	AddonId *string `json:"AddonId,omitempty" xml:"AddonId,omitempty"`
	// The addon description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The addon label.
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The addon name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Login
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The resource configurations of the addon.
	ResourcesSpec *ListClustersResponseBodyClustersAddonsResourcesSpec `json:"ResourcesSpec,omitempty" xml:"ResourcesSpec,omitempty" type:"Struct"`
	// The information about the addon services.
	ServicesSpec []*ListClustersResponseBodyClustersAddonsServicesSpec `json:"ServicesSpec,omitempty" xml:"ServicesSpec,omitempty" type:"Repeated"`
	// The addon state.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The addon version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListClustersResponseBodyClustersAddons) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBodyClustersAddons) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersAddons) GetAddonId() *string {
	return s.AddonId
}

func (s *ListClustersResponseBodyClustersAddons) GetDescription() *string {
	return s.Description
}

func (s *ListClustersResponseBodyClustersAddons) GetLabel() *string {
	return s.Label
}

func (s *ListClustersResponseBodyClustersAddons) GetName() *string {
	return s.Name
}

func (s *ListClustersResponseBodyClustersAddons) GetResourcesSpec() *ListClustersResponseBodyClustersAddonsResourcesSpec {
	return s.ResourcesSpec
}

func (s *ListClustersResponseBodyClustersAddons) GetServicesSpec() []*ListClustersResponseBodyClustersAddonsServicesSpec {
	return s.ServicesSpec
}

func (s *ListClustersResponseBodyClustersAddons) GetStatus() *string {
	return s.Status
}

func (s *ListClustersResponseBodyClustersAddons) GetVersion() *string {
	return s.Version
}

func (s *ListClustersResponseBodyClustersAddons) SetAddonId(v string) *ListClustersResponseBodyClustersAddons {
	s.AddonId = &v
	return s
}

func (s *ListClustersResponseBodyClustersAddons) SetDescription(v string) *ListClustersResponseBodyClustersAddons {
	s.Description = &v
	return s
}

func (s *ListClustersResponseBodyClustersAddons) SetLabel(v string) *ListClustersResponseBodyClustersAddons {
	s.Label = &v
	return s
}

func (s *ListClustersResponseBodyClustersAddons) SetName(v string) *ListClustersResponseBodyClustersAddons {
	s.Name = &v
	return s
}

func (s *ListClustersResponseBodyClustersAddons) SetResourcesSpec(v *ListClustersResponseBodyClustersAddonsResourcesSpec) *ListClustersResponseBodyClustersAddons {
	s.ResourcesSpec = v
	return s
}

func (s *ListClustersResponseBodyClustersAddons) SetServicesSpec(v []*ListClustersResponseBodyClustersAddonsServicesSpec) *ListClustersResponseBodyClustersAddons {
	s.ServicesSpec = v
	return s
}

func (s *ListClustersResponseBodyClustersAddons) SetStatus(v string) *ListClustersResponseBodyClustersAddons {
	s.Status = &v
	return s
}

func (s *ListClustersResponseBodyClustersAddons) SetVersion(v string) *ListClustersResponseBodyClustersAddons {
	s.Version = &v
	return s
}

func (s *ListClustersResponseBodyClustersAddons) Validate() error {
	if s.ResourcesSpec != nil {
		if err := s.ResourcesSpec.Validate(); err != nil {
			return err
		}
	}
	if s.ServicesSpec != nil {
		for _, item := range s.ServicesSpec {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListClustersResponseBodyClustersAddonsResourcesSpec struct {
	// The instance ID.
	//
	// example:
	//
	// i-bp1bg85d2q6laic8****
	EcsInstanceId *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	// The Elastic IP Address (EIP) ID.
	//
	// example:
	//
	// eip-bp1vi9124tbx1g3kr****
	EipInstanceId *string `json:"EipInstanceId,omitempty" xml:"EipInstanceId,omitempty"`
}

func (s ListClustersResponseBodyClustersAddonsResourcesSpec) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBodyClustersAddonsResourcesSpec) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersAddonsResourcesSpec) GetEcsInstanceId() *string {
	return s.EcsInstanceId
}

func (s *ListClustersResponseBodyClustersAddonsResourcesSpec) GetEipInstanceId() *string {
	return s.EipInstanceId
}

func (s *ListClustersResponseBodyClustersAddonsResourcesSpec) SetEcsInstanceId(v string) *ListClustersResponseBodyClustersAddonsResourcesSpec {
	s.EcsInstanceId = &v
	return s
}

func (s *ListClustersResponseBodyClustersAddonsResourcesSpec) SetEipInstanceId(v string) *ListClustersResponseBodyClustersAddonsResourcesSpec {
	s.EipInstanceId = &v
	return s
}

func (s *ListClustersResponseBodyClustersAddonsResourcesSpec) Validate() error {
	return dara.Validate(s)
}

type ListClustersResponseBodyClustersAddonsServicesSpec struct {
	// The service access type.
	//
	// example:
	//
	// URL
	ServiceAccessType *string `json:"ServiceAccessType,omitempty" xml:"ServiceAccessType,omitempty"`
	// The service access URL.
	//
	// example:
	//
	// https://47.96.xx.xx:12008
	ServiceAccessUrl *string `json:"ServiceAccessUrl,omitempty" xml:"ServiceAccessUrl,omitempty"`
	// The service name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Web Portal
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s ListClustersResponseBodyClustersAddonsServicesSpec) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBodyClustersAddonsServicesSpec) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersAddonsServicesSpec) GetServiceAccessType() *string {
	return s.ServiceAccessType
}

func (s *ListClustersResponseBodyClustersAddonsServicesSpec) GetServiceAccessUrl() *string {
	return s.ServiceAccessUrl
}

func (s *ListClustersResponseBodyClustersAddonsServicesSpec) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListClustersResponseBodyClustersAddonsServicesSpec) SetServiceAccessType(v string) *ListClustersResponseBodyClustersAddonsServicesSpec {
	s.ServiceAccessType = &v
	return s
}

func (s *ListClustersResponseBodyClustersAddonsServicesSpec) SetServiceAccessUrl(v string) *ListClustersResponseBodyClustersAddonsServicesSpec {
	s.ServiceAccessUrl = &v
	return s
}

func (s *ListClustersResponseBodyClustersAddonsServicesSpec) SetServiceName(v string) *ListClustersResponseBodyClustersAddonsServicesSpec {
	s.ServiceName = &v
	return s
}

func (s *ListClustersResponseBodyClustersAddonsServicesSpec) Validate() error {
	return dara.Validate(s)
}

type ListClustersResponseBodyClustersClusterCustomConfiguration struct {
	// The parameters of the post-processing script.
	//
	// example:
	//
	// demo
	Args *string `json:"Args,omitempty" xml:"Args,omitempty"`
	// The link to the post-processing script.
	//
	// example:
	//
	// https://xxxxx
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
}

func (s ListClustersResponseBodyClustersClusterCustomConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBodyClustersClusterCustomConfiguration) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersClusterCustomConfiguration) GetArgs() *string {
	return s.Args
}

func (s *ListClustersResponseBodyClustersClusterCustomConfiguration) GetScript() *string {
	return s.Script
}

func (s *ListClustersResponseBodyClustersClusterCustomConfiguration) SetArgs(v string) *ListClustersResponseBodyClustersClusterCustomConfiguration {
	s.Args = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterCustomConfiguration) SetScript(v string) *ListClustersResponseBodyClustersClusterCustomConfiguration {
	s.Script = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterCustomConfiguration) Validate() error {
	return dara.Validate(s)
}

type ListClustersResponseBodyClustersManager struct {
	// The configurations of the domain name resolution service.
	DNS *ListClustersResponseBodyClustersManagerDNS `json:"DNS,omitempty" xml:"DNS,omitempty" type:"Struct"`
	// The configurations of the directory service.
	DirectoryService *ListClustersResponseBodyClustersManagerDirectoryService `json:"DirectoryService,omitempty" xml:"DirectoryService,omitempty" type:"Struct"`
	// The configurations of the scheduler service.
	Scheduler *ListClustersResponseBodyClustersManagerScheduler `json:"Scheduler,omitempty" xml:"Scheduler,omitempty" type:"Struct"`
}

func (s ListClustersResponseBodyClustersManager) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBodyClustersManager) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersManager) GetDNS() *ListClustersResponseBodyClustersManagerDNS {
	return s.DNS
}

func (s *ListClustersResponseBodyClustersManager) GetDirectoryService() *ListClustersResponseBodyClustersManagerDirectoryService {
	return s.DirectoryService
}

func (s *ListClustersResponseBodyClustersManager) GetScheduler() *ListClustersResponseBodyClustersManagerScheduler {
	return s.Scheduler
}

func (s *ListClustersResponseBodyClustersManager) SetDNS(v *ListClustersResponseBodyClustersManagerDNS) *ListClustersResponseBodyClustersManager {
	s.DNS = v
	return s
}

func (s *ListClustersResponseBodyClustersManager) SetDirectoryService(v *ListClustersResponseBodyClustersManagerDirectoryService) *ListClustersResponseBodyClustersManager {
	s.DirectoryService = v
	return s
}

func (s *ListClustersResponseBodyClustersManager) SetScheduler(v *ListClustersResponseBodyClustersManagerScheduler) *ListClustersResponseBodyClustersManager {
	s.Scheduler = v
	return s
}

func (s *ListClustersResponseBodyClustersManager) Validate() error {
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
	if s.Scheduler != nil {
		if err := s.Scheduler.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListClustersResponseBodyClustersManagerDNS struct {
	// The resolution type.
	//
	// example:
	//
	// NIS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version of the domain name resolution service.
	//
	// example:
	//
	// 2.31
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListClustersResponseBodyClustersManagerDNS) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBodyClustersManagerDNS) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersManagerDNS) GetType() *string {
	return s.Type
}

func (s *ListClustersResponseBodyClustersManagerDNS) GetVersion() *string {
	return s.Version
}

func (s *ListClustersResponseBodyClustersManagerDNS) SetType(v string) *ListClustersResponseBodyClustersManagerDNS {
	s.Type = &v
	return s
}

func (s *ListClustersResponseBodyClustersManagerDNS) SetVersion(v string) *ListClustersResponseBodyClustersManagerDNS {
	s.Version = &v
	return s
}

func (s *ListClustersResponseBodyClustersManagerDNS) Validate() error {
	return dara.Validate(s)
}

type ListClustersResponseBodyClustersManagerDirectoryService struct {
	// The type of the domain account.
	//
	// example:
	//
	// NIS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version of the domain account service.
	//
	// example:
	//
	// 2.31
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListClustersResponseBodyClustersManagerDirectoryService) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBodyClustersManagerDirectoryService) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersManagerDirectoryService) GetType() *string {
	return s.Type
}

func (s *ListClustersResponseBodyClustersManagerDirectoryService) GetVersion() *string {
	return s.Version
}

func (s *ListClustersResponseBodyClustersManagerDirectoryService) SetType(v string) *ListClustersResponseBodyClustersManagerDirectoryService {
	s.Type = &v
	return s
}

func (s *ListClustersResponseBodyClustersManagerDirectoryService) SetVersion(v string) *ListClustersResponseBodyClustersManagerDirectoryService {
	s.Version = &v
	return s
}

func (s *ListClustersResponseBodyClustersManagerDirectoryService) Validate() error {
	return dara.Validate(s)
}

type ListClustersResponseBodyClustersManagerScheduler struct {
	// The scheduler type.
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

func (s ListClustersResponseBodyClustersManagerScheduler) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBodyClustersManagerScheduler) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersManagerScheduler) GetType() *string {
	return s.Type
}

func (s *ListClustersResponseBodyClustersManagerScheduler) GetVersion() *string {
	return s.Version
}

func (s *ListClustersResponseBodyClustersManagerScheduler) SetType(v string) *ListClustersResponseBodyClustersManagerScheduler {
	s.Type = &v
	return s
}

func (s *ListClustersResponseBodyClustersManagerScheduler) SetVersion(v string) *ListClustersResponseBodyClustersManagerScheduler {
	s.Version = &v
	return s
}

func (s *ListClustersResponseBodyClustersManagerScheduler) Validate() error {
	return dara.Validate(s)
}

type ListClustersResponseBodyClustersNodes struct {
	// The number of malfunctioning compute nodes.
	//
	// example:
	//
	// 0
	AbnormalCounts *int32 `json:"AbnormalCounts,omitempty" xml:"AbnormalCounts,omitempty"`
	// The number of compute nodes that are being created.
	//
	// example:
	//
	// 0
	CreatingCounts *int32 `json:"CreatingCounts,omitempty" xml:"CreatingCounts,omitempty"`
	// The number of running compute nodes.
	//
	// example:
	//
	// 1
	RunningCounts *int32 `json:"RunningCounts,omitempty" xml:"RunningCounts,omitempty"`
}

func (s ListClustersResponseBodyClustersNodes) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBodyClustersNodes) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersNodes) GetAbnormalCounts() *int32 {
	return s.AbnormalCounts
}

func (s *ListClustersResponseBodyClustersNodes) GetCreatingCounts() *int32 {
	return s.CreatingCounts
}

func (s *ListClustersResponseBodyClustersNodes) GetRunningCounts() *int32 {
	return s.RunningCounts
}

func (s *ListClustersResponseBodyClustersNodes) SetAbnormalCounts(v int32) *ListClustersResponseBodyClustersNodes {
	s.AbnormalCounts = &v
	return s
}

func (s *ListClustersResponseBodyClustersNodes) SetCreatingCounts(v int32) *ListClustersResponseBodyClustersNodes {
	s.CreatingCounts = &v
	return s
}

func (s *ListClustersResponseBodyClustersNodes) SetRunningCounts(v int32) *ListClustersResponseBodyClustersNodes {
	s.RunningCounts = &v
	return s
}

func (s *ListClustersResponseBodyClustersNodes) Validate() error {
	return dara.Validate(s)
}

type ListClustersResponseBodyClustersUsers struct {
	// The number of ordinary users.
	//
	// example:
	//
	// 2
	NormalCounts *int32 `json:"NormalCounts,omitempty" xml:"NormalCounts,omitempty"`
	// The number of administrators.
	//
	// example:
	//
	// 2
	SudoCounts *int32 `json:"SudoCounts,omitempty" xml:"SudoCounts,omitempty"`
}

func (s ListClustersResponseBodyClustersUsers) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBodyClustersUsers) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersUsers) GetNormalCounts() *int32 {
	return s.NormalCounts
}

func (s *ListClustersResponseBodyClustersUsers) GetSudoCounts() *int32 {
	return s.SudoCounts
}

func (s *ListClustersResponseBodyClustersUsers) SetNormalCounts(v int32) *ListClustersResponseBodyClustersUsers {
	s.NormalCounts = &v
	return s
}

func (s *ListClustersResponseBodyClustersUsers) SetSudoCounts(v int32) *ListClustersResponseBodyClustersUsers {
	s.SudoCounts = &v
	return s
}

func (s *ListClustersResponseBodyClustersUsers) Validate() error {
	return dara.Validate(s)
}
