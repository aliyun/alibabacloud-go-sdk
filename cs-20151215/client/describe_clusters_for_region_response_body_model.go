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
	// The information about the queried clusters.
	Clusters []*DescribeClustersForRegionResponseBodyClusters `json:"clusters,omitempty" xml:"clusters,omitempty" type:"Repeated"`
	// The pagination details.
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
	return dara.Validate(s)
}

type DescribeClustersForRegionResponseBodyClusters struct {
	// The domain name of the cluster.
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
	// The types of ACK managed clusters:
	//
	// 	- ack.pro.small: ACK Pro cluster
	//
	// 	- ack.standard: ACK Basic cluster
	//
	// example:
	//
	// ack.standard
	ClusterSpec *string `json:"cluster_spec,omitempty" xml:"cluster_spec,omitempty"`
	// The type of the cluster. Valid values:
	//
	// 	- Kubernetes: ACK dedicated cluster
	//
	// 	- ManagedKubernetes: ACK managed clusters. ACK managed clusters include ACK Basic clusters, ACK Pro clusters, ACK Serverless Basic clusters, ACK Serverless Pro clusters, ACK Edge Basic clusters, ACK Edge Pro clusters, and ACK Lingjun Pro clusters.
	//
	// 	- ExternalKubernetes: registered cluster
	//
	// example:
	//
	// ManagedKubernetes
	ClusterType *string `json:"cluster_type,omitempty" xml:"cluster_type,omitempty"`
	// The CIDR block of pods in the cluster.
	//
	// example:
	//
	// 172.20.0.0/16
	ContainerCidr *string `json:"container_cidr,omitempty" xml:"container_cidr,omitempty"`
	// The time at which the instance is created.
	//
	// example:
	//
	// 2020-12-01T20:40:40+08:00
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// The current Kubernetes version of the cluster.
	//
	// example:
	//
	// 1.16.6-aliyun.1
	CurrentVersion *string `json:"current_version,omitempty" xml:"current_version,omitempty"`
	// Specifies whether to enable cluster deletion protection. If you enable this option, the cluster cannot be deleted in the console or by calling API operations. You can obtain the terminal ID by calling one of the following operations:
	//
	// 	- true: enables deletion protection for the cluster. This way, the cluster cannot be deleted in the ACK console or by calling API operations.
	//
	// 	- false: disables deletion protection for the cluster. This way, the cluster can be deleted in the ACK console or by calling API operations.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"deletion_protection,omitempty" xml:"deletion_protection,omitempty"`
	// The initial Kubernetes version of the cluster.
	//
	// example:
	//
	// 1.16.6-aliyun.1
	InitVersion *string `json:"init_version,omitempty" xml:"init_version,omitempty"`
	// The IP protocol stack of the cluster.
	//
	// example:
	//
	// ipv4
	IpStack *string `json:"ip_stack,omitempty" xml:"ip_stack,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// test-cluster
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The Kubernetes version to which the cluster can be updated.
	//
	// example:
	//
	// 1.18.8-aliyun.1
	NextVersion *string `json:"next_version,omitempty" xml:"next_version,omitempty"`
	// The subtype of the clusters. Valid values:
	//
	// 	- Default: ACK managed clusters. ACK managed clusters include ACK Basic clusters and ACK Pro clusters.
	//
	// 	- Edge: ACK Edge clusters. ACK Edge clusters include ACK Edge Basic clusters and ACK Edge Pro clusters.
	//
	// 	- Serverless: ACK Serverless clusters. ACK Serverless clusters include ACK Serverless Basic clusters and ACK Serverless Pro clusters.
	//
	// 	- Lingjun: ACK Lingjun Pro clusters.
	//
	// example:
	//
	// Default
	Profile *string `json:"profile,omitempty" xml:"profile,omitempty"`
	// The kube-proxy mode of the cluster.
	//
	// Valid value:
	//
	// 	- iptables: iptables.
	//
	// 	- ipvs: ipvs.
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
	// The ID of the cluster resource group.
	//
	// example:
	//
	// rg-acfmyvw3wjm****
	ResourceGroupId *string `json:"resource_group_id,omitempty" xml:"resource_group_id,omitempty"`
	// The ID of the security group of the cluster.
	//
	// example:
	//
	// sg-2zeihch86ooz9io4****
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id,omitempty"`
	// The CIDR block of the service network.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.21.0.0/20
	ServiceCidr *string `json:"service_cidr,omitempty" xml:"service_cidr,omitempty"`
	// The number of nodes in the ACK cluster.
	//
	// example:
	//
	// 2
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// The status of the cluster. Valid values:
	//
	// 	- initial: The cluster is being created.
	//
	// 	- failed: The cluster failed to be created.
	//
	// 	- running: The cluster is running.
	//
	// 	- Upgrading: The cluster is being updated.
	//
	// 	- scaling: The cluster is being scaled.
	//
	// 	- waiting: The cluster is waiting for connection requests.
	//
	// 	- disconnected: The cluster is disconnected.
	//
	// 	- inactive: The cluster is inactive.
	//
	// 	- unavailable: The cluster is unavailable.
	//
	// 	- deleting: The cluster is being deleted.
	//
	// 	- deleted: The ACK cluster is deleted.
	//
	// 	- delete_failed: The cluster failed to be deleted.
	//
	// example:
	//
	// running
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The list of cluster tags.
	Tags []*Tag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The time zone.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// The time when the cluster was updated.
	//
	// example:
	//
	// 2020-12-08T15:37:00+08:00
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
	// The ID of the virtual private cloud (VPC) to which the cluster belongs.
	//
	// example:
	//
	// vpc-2zeg8nf1ukc0fcmvq****
	VpcId *string `json:"vpc_id,omitempty" xml:"vpc_id,omitempty"`
	// The list of vSwitches on the control plane of the cluster.
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
	return dara.Validate(s)
}

type DescribeClustersForRegionResponseBodyPageInfo struct {
	// The number of pages.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// The number of records on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// The total number of entries returned.
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
