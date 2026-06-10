// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClustersForRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusters(v []*DescribeClustersForRegionResponseBodyClusters) *DescribeClustersForRegionResponseBody
	GetClusters() []*DescribeClustersForRegionResponseBodyClusters
	SetPageInfo(v *DescribeClustersForRegionResponseBodyPageInfo) *DescribeClustersForRegionResponseBody
	GetPageInfo() *DescribeClustersForRegionResponseBodyPageInfo
}

type DescribeClustersForRegionResponseBody struct {
	// A list of clusters.
	Clusters []*DescribeClustersForRegionResponseBodyClusters `json:"clusters,omitempty" xml:"clusters,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeClustersForRegionResponseBodyPageInfo `json:"page_info,omitempty" xml:"page_info,omitempty" type:"Struct"`
}

func (s DescribeClustersForRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClustersForRegionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClustersForRegionResponseBody) GetClusters() []*DescribeClustersForRegionResponseBodyClusters {
	return s.Clusters
}

func (s *DescribeClustersForRegionResponseBody) GetPageInfo() *DescribeClustersForRegionResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeClustersForRegionResponseBody) SetClusters(v []*DescribeClustersForRegionResponseBodyClusters) *DescribeClustersForRegionResponseBody {
	s.Clusters = v
	return s
}

func (s *DescribeClustersForRegionResponseBody) SetPageInfo(v *DescribeClustersForRegionResponseBodyPageInfo) *DescribeClustersForRegionResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeClustersForRegionResponseBody) Validate() error {
	if s.Clusters != nil {
		for _, item := range s.Clusters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClustersForRegionResponseBodyClusters struct {
	// The cluster domain.
	//
	// example:
	//
	// cluster.local
	ClusterDomain *string `json:"cluster_domain,omitempty" xml:"cluster_domain,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// c905d1364c2dd4b6284a3f41790c4****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The specification of the cluster. Valid values:
	//
	// - `ack.standard`: Basic Edition
	//
	// - `ack.pro.small`: Pro Edition
	//
	// - `ack.pro.xlarge`: Pro XL
	//
	// - `ack.pro.2xlarge`: Pro 2XL
	//
	// - `ack.pro.4xlarge`: Pro 4XL. This specification is available only to allowlisted users.
	//
	// Pro XL, Pro 2XL, and Pro 4XL are three specifications available for the <props="china">[ACK Pro provisioned control plane](https://help.aliyun.com/ack/ack-managed-and-ack-dedicated/user-guide/ack-pro-provisioned-control-plane)<props="intl">[ACK Pro provisioned control plane](https://www.alibabacloud.com/help/ack/ack-managed-and-ack-dedicated/user-guide/ack-pro-provisioned-control-plane). These specifications ensure a high and deterministic level of API concurrency and Pod scheduling capabilities by pre-allocating and dedicating control plane resources. They are suitable for AI training and inference, large-scale clusters, and mission-critical workloads.
	//
	// For information about the <props="china">[cluster management fee](https://help.aliyun.com/ack/ack-managed-and-ack-dedicated/product-overview/cluster-management-fee)<props="intl">[cluster management fee](https://www.alibabacloud.com/help/ack/ack-managed-and-ack-dedicated/product-overview/cluster-management-fee) for Pro Edition and ACK Pro provisioned control plane specifications, see the linked topic.
	//
	// example:
	//
	// ack.standard
	ClusterSpec *string `json:"cluster_spec,omitempty" xml:"cluster_spec,omitempty"`
	// The type of the cluster. Valid values:
	//
	// - `Kubernetes`: an ACK dedicated cluster.
	//
	// - `ManagedKubernetes`: an ACK managed cluster. This type includes ACK managed clusters (Pro and Basic editions), ACK Serverless clusters (Pro and Basic editions), ACK Edge clusters (Pro and Basic editions), and ACK Lingjun clusters (Pro edition).
	//
	// - `ExternalKubernetes`: a registered cluster.
	//
	// example:
	//
	// ManagedKubernetes
	ClusterType *string `json:"cluster_type,omitempty" xml:"cluster_type,omitempty"`
	// The CIDR block for Pods in the cluster.
	//
	// example:
	//
	// 172.20.0.0/16
	ContainerCidr *string `json:"container_cidr,omitempty" xml:"container_cidr,omitempty"`
	// The time the cluster was created.
	//
	// example:
	//
	// 2020-12-01T20:40:40+08:00
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// The current version of the cluster.
	//
	// example:
	//
	// 1.16.6-aliyun.1
	CurrentVersion *string `json:"current_version,omitempty" xml:"current_version,omitempty"`
	// Specifies whether deletion protection is enabled for the cluster. If enabled, you cannot delete the cluster from the console or by an API call. Valid values:
	//
	// - `true`: Deletion protection is enabled.
	//
	// - `false`: Deletion protection is disabled.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"deletion_protection,omitempty" xml:"deletion_protection,omitempty"`
	// The initial version of the cluster.
	//
	// example:
	//
	// 1.16.6-aliyun.1
	InitVersion *string `json:"init_version,omitempty" xml:"init_version,omitempty"`
	// The IP stack of the cluster.
	//
	// example:
	//
	// ipv4
	IpStack *string `json:"ip_stack,omitempty" xml:"ip_stack,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// test-cluster
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The available upgrade version.
	//
	// example:
	//
	// 1.18.8-aliyun.1
	NextVersion *string `json:"next_version,omitempty" xml:"next_version,omitempty"`
	// The subtype of the cluster. Valid values:
	//
	// - `Default`: An ACK managed cluster (Pro and Basic editions).
	//
	// - `Edge`: An ACK Edge cluster (Pro and Basic editions).
	//
	// - `Serverless`: An ACK Serverless cluster (Pro and Basic editions).
	//
	// - `LingJun`: An ACK Lingjun cluster (Pro edition).
	//
	// example:
	//
	// Default
	Profile *string `json:"profile,omitempty" xml:"profile,omitempty"`
	// The kube-proxy proxy mode of the cluster.
	//
	// example:
	//
	// ipvs
	ProxyMode *string `json:"proxy_mode,omitempty" xml:"proxy_mode,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing-a
	RegionId *string `json:"region_id,omitempty" xml:"region_id,omitempty"`
	// The ID of the resource group to which the cluster belongs.
	//
	// example:
	//
	// rg-acfmyvw3wjm****
	ResourceGroupId *string `json:"resource_group_id,omitempty" xml:"resource_group_id,omitempty"`
	// The security group ID of the cluster.
	//
	// example:
	//
	// sg-2zeihch86ooz9io4****
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id,omitempty"`
	// The CIDR block for the service network.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.21.0.0/20
	ServiceCidr *string `json:"service_cidr,omitempty" xml:"service_cidr,omitempty"`
	// The number of nodes in the cluster.
	//
	// example:
	//
	// 2
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// The state of the cluster. Valid values:
	//
	// - `initial`: The cluster is being created.
	//
	// - `failed`: Cluster creation failed.
	//
	// - `running`: The cluster is running.
	//
	// - `updating`: The cluster is being updated.
	//
	// - `upgrading`: The cluster is being upgraded.
	//
	// - `removing`: Nodes are being removed from the cluster.
	//
	// - `draining`: Node draining is in progress.
	//
	// - `scaling`: The cluster is being scaled.
	//
	// - `inactive`: The cluster is inactive.
	//
	// - `unavailable`: The cluster is unavailable.
	//
	// - `deleting`: The cluster is being deleted.
	//
	// - `deleted`: The cluster is deleted.
	//
	// - `delete_failed`: Cluster deletion failed.
	//
	// - `waiting`: The cluster is waiting for a connection.
	//
	// - `disconnected`: The cluster is disconnected.
	//
	// example:
	//
	// running
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The tags attached to the cluster.
	Tags []*Tag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The time zone of the cluster.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// The time the cluster was last updated.
	//
	// example:
	//
	// 2020-12-08T15:37:00+08:00
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
	// The VPC ID of the cluster.
	//
	// example:
	//
	// vpc-2zeg8nf1ukc0fcmvq****
	VpcId *string `json:"vpc_id,omitempty" xml:"vpc_id,omitempty"`
	// The IDs of the vSwitches for the control plane.
	VswitchIds []*string `json:"vswitch_ids,omitempty" xml:"vswitch_ids,omitempty" type:"Repeated"`
}

func (s DescribeClustersForRegionResponseBodyClusters) String() string {
	return dara.Prettify(s)
}

func (s DescribeClustersForRegionResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *DescribeClustersForRegionResponseBodyClusters) GetClusterDomain() *string {
	return s.ClusterDomain
}

func (s *DescribeClustersForRegionResponseBodyClusters) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClustersForRegionResponseBodyClusters) GetClusterSpec() *string {
	return s.ClusterSpec
}

func (s *DescribeClustersForRegionResponseBodyClusters) GetClusterType() *string {
	return s.ClusterType
}

func (s *DescribeClustersForRegionResponseBodyClusters) GetContainerCidr() *string {
	return s.ContainerCidr
}

func (s *DescribeClustersForRegionResponseBodyClusters) GetCreated() *string {
	return s.Created
}

func (s *DescribeClustersForRegionResponseBodyClusters) GetCurrentVersion() *string {
	return s.CurrentVersion
}

func (s *DescribeClustersForRegionResponseBodyClusters) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *DescribeClustersForRegionResponseBodyClusters) GetInitVersion() *string {
	return s.InitVersion
}

func (s *DescribeClustersForRegionResponseBodyClusters) GetIpStack() *string {
	return s.IpStack
}

func (s *DescribeClustersForRegionResponseBodyClusters) GetName() *string {
	return s.Name
}

func (s *DescribeClustersForRegionResponseBodyClusters) GetNextVersion() *string {
	return s.NextVersion
}

func (s *DescribeClustersForRegionResponseBodyClusters) GetProfile() *string {
	return s.Profile
}

func (s *DescribeClustersForRegionResponseBodyClusters) GetProxyMode() *string {
	return s.ProxyMode
}

func (s *DescribeClustersForRegionResponseBodyClusters) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeClustersForRegionResponseBodyClusters) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeClustersForRegionResponseBodyClusters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeClustersForRegionResponseBodyClusters) GetServiceCidr() *string {
	return s.ServiceCidr
}

func (s *DescribeClustersForRegionResponseBodyClusters) GetSize() *int64 {
	return s.Size
}

func (s *DescribeClustersForRegionResponseBodyClusters) GetState() *string {
	return s.State
}

func (s *DescribeClustersForRegionResponseBodyClusters) GetTags() []*Tag {
	return s.Tags
}

func (s *DescribeClustersForRegionResponseBodyClusters) GetTimezone() *string {
	return s.Timezone
}

func (s *DescribeClustersForRegionResponseBodyClusters) GetUpdated() *string {
	return s.Updated
}

func (s *DescribeClustersForRegionResponseBodyClusters) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeClustersForRegionResponseBodyClusters) GetVswitchIds() []*string {
	return s.VswitchIds
}

func (s *DescribeClustersForRegionResponseBodyClusters) SetClusterDomain(v string) *DescribeClustersForRegionResponseBodyClusters {
	s.ClusterDomain = &v
	return s
}

func (s *DescribeClustersForRegionResponseBodyClusters) SetClusterId(v string) *DescribeClustersForRegionResponseBodyClusters {
	s.ClusterId = &v
	return s
}

func (s *DescribeClustersForRegionResponseBodyClusters) SetClusterSpec(v string) *DescribeClustersForRegionResponseBodyClusters {
	s.ClusterSpec = &v
	return s
}

func (s *DescribeClustersForRegionResponseBodyClusters) SetClusterType(v string) *DescribeClustersForRegionResponseBodyClusters {
	s.ClusterType = &v
	return s
}

func (s *DescribeClustersForRegionResponseBodyClusters) SetContainerCidr(v string) *DescribeClustersForRegionResponseBodyClusters {
	s.ContainerCidr = &v
	return s
}

func (s *DescribeClustersForRegionResponseBodyClusters) SetCreated(v string) *DescribeClustersForRegionResponseBodyClusters {
	s.Created = &v
	return s
}

func (s *DescribeClustersForRegionResponseBodyClusters) SetCurrentVersion(v string) *DescribeClustersForRegionResponseBodyClusters {
	s.CurrentVersion = &v
	return s
}

func (s *DescribeClustersForRegionResponseBodyClusters) SetDeletionProtection(v bool) *DescribeClustersForRegionResponseBodyClusters {
	s.DeletionProtection = &v
	return s
}

func (s *DescribeClustersForRegionResponseBodyClusters) SetInitVersion(v string) *DescribeClustersForRegionResponseBodyClusters {
	s.InitVersion = &v
	return s
}

func (s *DescribeClustersForRegionResponseBodyClusters) SetIpStack(v string) *DescribeClustersForRegionResponseBodyClusters {
	s.IpStack = &v
	return s
}

func (s *DescribeClustersForRegionResponseBodyClusters) SetName(v string) *DescribeClustersForRegionResponseBodyClusters {
	s.Name = &v
	return s
}

func (s *DescribeClustersForRegionResponseBodyClusters) SetNextVersion(v string) *DescribeClustersForRegionResponseBodyClusters {
	s.NextVersion = &v
	return s
}

func (s *DescribeClustersForRegionResponseBodyClusters) SetProfile(v string) *DescribeClustersForRegionResponseBodyClusters {
	s.Profile = &v
	return s
}

func (s *DescribeClustersForRegionResponseBodyClusters) SetProxyMode(v string) *DescribeClustersForRegionResponseBodyClusters {
	s.ProxyMode = &v
	return s
}

func (s *DescribeClustersForRegionResponseBodyClusters) SetRegionId(v string) *DescribeClustersForRegionResponseBodyClusters {
	s.RegionId = &v
	return s
}

func (s *DescribeClustersForRegionResponseBodyClusters) SetResourceGroupId(v string) *DescribeClustersForRegionResponseBodyClusters {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeClustersForRegionResponseBodyClusters) SetSecurityGroupId(v string) *DescribeClustersForRegionResponseBodyClusters {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeClustersForRegionResponseBodyClusters) SetServiceCidr(v string) *DescribeClustersForRegionResponseBodyClusters {
	s.ServiceCidr = &v
	return s
}

func (s *DescribeClustersForRegionResponseBodyClusters) SetSize(v int64) *DescribeClustersForRegionResponseBodyClusters {
	s.Size = &v
	return s
}

func (s *DescribeClustersForRegionResponseBodyClusters) SetState(v string) *DescribeClustersForRegionResponseBodyClusters {
	s.State = &v
	return s
}

func (s *DescribeClustersForRegionResponseBodyClusters) SetTags(v []*Tag) *DescribeClustersForRegionResponseBodyClusters {
	s.Tags = v
	return s
}

func (s *DescribeClustersForRegionResponseBodyClusters) SetTimezone(v string) *DescribeClustersForRegionResponseBodyClusters {
	s.Timezone = &v
	return s
}

func (s *DescribeClustersForRegionResponseBodyClusters) SetUpdated(v string) *DescribeClustersForRegionResponseBodyClusters {
	s.Updated = &v
	return s
}

func (s *DescribeClustersForRegionResponseBodyClusters) SetVpcId(v string) *DescribeClustersForRegionResponseBodyClusters {
	s.VpcId = &v
	return s
}

func (s *DescribeClustersForRegionResponseBodyClusters) SetVswitchIds(v []*string) *DescribeClustersForRegionResponseBodyClusters {
	s.VswitchIds = v
	return s
}

func (s *DescribeClustersForRegionResponseBodyClusters) Validate() error {
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

type DescribeClustersForRegionResponseBodyPageInfo struct {
	// The returned page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// The total number of entries that match the query.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

func (s DescribeClustersForRegionResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeClustersForRegionResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeClustersForRegionResponseBodyPageInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeClustersForRegionResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeClustersForRegionResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeClustersForRegionResponseBodyPageInfo) SetPageNumber(v int32) *DescribeClustersForRegionResponseBodyPageInfo {
	s.PageNumber = &v
	return s
}

func (s *DescribeClustersForRegionResponseBodyPageInfo) SetPageSize(v int32) *DescribeClustersForRegionResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeClustersForRegionResponseBodyPageInfo) SetTotalCount(v int32) *DescribeClustersForRegionResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeClustersForRegionResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
