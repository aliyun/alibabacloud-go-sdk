// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClustersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClustersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClustersResponse
	GetStatusCode() *int32
	SetBody(v []*DescribeClustersResponseBody) *DescribeClustersResponse
	GetBody() []*DescribeClustersResponseBody
}

type DescribeClustersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*DescribeClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s DescribeClustersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClustersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClustersResponse) GetBody() []*DescribeClustersResponseBody {
	return s.Body
}

func (s *DescribeClustersResponse) SetHeaders(v map[string]*string) *DescribeClustersResponse {
	s.Headers = v
	return s
}

func (s *DescribeClustersResponse) SetStatusCode(v int32) *DescribeClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClustersResponse) SetBody(v []*DescribeClustersResponseBody) *DescribeClustersResponse {
	s.Body = v
	return s
}

func (s *DescribeClustersResponse) Validate() error {
	if s.Body != nil {
		for _, item := range s.Body {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeClustersResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// c0123456789abcdef0123456789abcdef0
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The cluster type.
	//
	// example:
	//
	// ManagedKubernetes
	ClusterType *string `json:"cluster_type,omitempty" xml:"cluster_type,omitempty"`
	// The time when the cluster was created.
	//
	// example:
	//
	// 2026-10-01T08:00:00+08:00
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// The current version of the cluster.
	//
	// example:
	//
	// 1.32.1-aliyun.1
	CurrentVersion *string `json:"current_version,omitempty" xml:"current_version,omitempty"`
	// The data disk category of the node.
	DataDiskCategory *string `json:"data_disk_category,omitempty" xml:"data_disk_category,omitempty"`
	// The data disk size of the node.
	DataDiskSize *int64 `json:"data_disk_size,omitempty" xml:"data_disk_size,omitempty"`
	// Indicates whether deletion protection is enabled for the cluster.
	DeletionProtection *bool `json:"deletion_protection,omitempty" xml:"deletion_protection,omitempty"`
	// The Docker version of the cluster.
	//
	// example:
	//
	// 20.10.17
	DockerVersion *string `json:"docker_version,omitempty" xml:"docker_version,omitempty"`
	// The ID of the external load balancing instance used by the cluster.
	//
	// example:
	//
	// lb-a1b2c3d4e5f6g7h8i9j0k
	ExternalLoadbalancerId *string `json:"external_loadbalancer_id,omitempty" xml:"external_loadbalancer_id,omitempty"`
	// The initial version of the cluster.
	//
	// example:
	//
	// 1.32.1-aliyun.1
	InitVersion *string `json:"init_version,omitempty" xml:"init_version,omitempty"`
	// The access address of the cluster API server, including the internal network access address and the public network access address.
	//
	// example:
	//
	// {"api_server_endpoint":"","intranet_api_server_endpoint":"https://192.0.2.10:6443"}
	MasterUrl *string `json:"master_url,omitempty" xml:"master_url,omitempty"`
	// The metadata of the cluster.
	//
	// example:
	//
	// {"addons":[{"name":"kube-proxy","version":"v1.32.1"}]}
	MetaData *string `json:"meta_data,omitempty" xml:"meta_data,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// example-cluster
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The network mode of the cluster.
	//
	// example:
	//
	// vpc
	NetworkMode *string `json:"network_mode,omitempty" xml:"network_mode,omitempty"`
	// Indicates whether Private Zone is enabled for the cluster. Default value: false.
	PrivateZone *bool `json:"private_zone,omitempty" xml:"private_zone,omitempty"`
	// The subtype of the cluster.
	//
	// example:
	//
	// Default
	Profile *string `json:"profile,omitempty" xml:"profile,omitempty"`
	// The region ID of the cluster.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"region_id,omitempty" xml:"region_id,omitempty"`
	// The resource group ID of the cluster.
	//
	// example:
	//
	// rg-a1b2c3d4e5f6g7h
	ResourceGroupId *string `json:"resource_group_id,omitempty" xml:"resource_group_id,omitempty"`
	// The security group ID of the cluster.
	//
	// example:
	//
	// sg-a1b2c3d4e5f6g7h8i9j0
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id,omitempty"`
	// The current number of nodes in the cluster, including master and worker nodes.
	//
	// example:
	//
	// 3
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// The state of the cluster.
	//
	// example:
	//
	// running
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The subnet CIDR of the cluster.
	//
	// example:
	//
	// 172.16.0.0/16
	SubnetCidr *string `json:"subnet_cidr,omitempty" xml:"subnet_cidr,omitempty"`
	// The tags of the cluster.
	Tags []*DescribeClustersResponseBodyTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The time when the cluster was updated.
	//
	// example:
	//
	// 2026-10-02T08:00:00+08:00
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
	// The VPC ID of the cluster.
	//
	// example:
	//
	// vpc-a1b2c3d4e5f6g7h8i9j0k
	VpcId *string `json:"vpc_id,omitempty" xml:"vpc_id,omitempty"`
	// The vSwitch CIDR.
	VswitchCidr *string `json:"vswitch_cidr,omitempty" xml:"vswitch_cidr,omitempty"`
	// The ID of the vSwitch where the cluster is deployed.
	//
	// example:
	//
	// vsw-a1b2c3d4e5f6g7h8i9j0k
	VswitchId *string `json:"vswitch_id,omitempty" xml:"vswitch_id,omitempty"`
	// The name of the RAM role used by the worker nodes.
	//
	// example:
	//
	// example-worker-ram-role
	WorkerRamRoleName *string `json:"worker_ram_role_name,omitempty" xml:"worker_ram_role_name,omitempty"`
	// The zone ID of the cluster.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"zone_id,omitempty" xml:"zone_id,omitempty"`
}

func (s DescribeClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClustersResponseBody) GetClusterType() *string {
	return s.ClusterType
}

func (s *DescribeClustersResponseBody) GetCreated() *string {
	return s.Created
}

func (s *DescribeClustersResponseBody) GetCurrentVersion() *string {
	return s.CurrentVersion
}

func (s *DescribeClustersResponseBody) GetDataDiskCategory() *string {
	return s.DataDiskCategory
}

func (s *DescribeClustersResponseBody) GetDataDiskSize() *int64 {
	return s.DataDiskSize
}

func (s *DescribeClustersResponseBody) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *DescribeClustersResponseBody) GetDockerVersion() *string {
	return s.DockerVersion
}

func (s *DescribeClustersResponseBody) GetExternalLoadbalancerId() *string {
	return s.ExternalLoadbalancerId
}

func (s *DescribeClustersResponseBody) GetInitVersion() *string {
	return s.InitVersion
}

func (s *DescribeClustersResponseBody) GetMasterUrl() *string {
	return s.MasterUrl
}

func (s *DescribeClustersResponseBody) GetMetaData() *string {
	return s.MetaData
}

func (s *DescribeClustersResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeClustersResponseBody) GetNetworkMode() *string {
	return s.NetworkMode
}

func (s *DescribeClustersResponseBody) GetPrivateZone() *bool {
	return s.PrivateZone
}

func (s *DescribeClustersResponseBody) GetProfile() *string {
	return s.Profile
}

func (s *DescribeClustersResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeClustersResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeClustersResponseBody) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeClustersResponseBody) GetSize() *int64 {
	return s.Size
}

func (s *DescribeClustersResponseBody) GetState() *string {
	return s.State
}

func (s *DescribeClustersResponseBody) GetSubnetCidr() *string {
	return s.SubnetCidr
}

func (s *DescribeClustersResponseBody) GetTags() []*DescribeClustersResponseBodyTags {
	return s.Tags
}

func (s *DescribeClustersResponseBody) GetUpdated() *string {
	return s.Updated
}

func (s *DescribeClustersResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeClustersResponseBody) GetVswitchCidr() *string {
	return s.VswitchCidr
}

func (s *DescribeClustersResponseBody) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeClustersResponseBody) GetWorkerRamRoleName() *string {
	return s.WorkerRamRoleName
}

func (s *DescribeClustersResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeClustersResponseBody) SetClusterId(v string) *DescribeClustersResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeClustersResponseBody) SetClusterType(v string) *DescribeClustersResponseBody {
	s.ClusterType = &v
	return s
}

func (s *DescribeClustersResponseBody) SetCreated(v string) *DescribeClustersResponseBody {
	s.Created = &v
	return s
}

func (s *DescribeClustersResponseBody) SetCurrentVersion(v string) *DescribeClustersResponseBody {
	s.CurrentVersion = &v
	return s
}

func (s *DescribeClustersResponseBody) SetDataDiskCategory(v string) *DescribeClustersResponseBody {
	s.DataDiskCategory = &v
	return s
}

func (s *DescribeClustersResponseBody) SetDataDiskSize(v int64) *DescribeClustersResponseBody {
	s.DataDiskSize = &v
	return s
}

func (s *DescribeClustersResponseBody) SetDeletionProtection(v bool) *DescribeClustersResponseBody {
	s.DeletionProtection = &v
	return s
}

func (s *DescribeClustersResponseBody) SetDockerVersion(v string) *DescribeClustersResponseBody {
	s.DockerVersion = &v
	return s
}

func (s *DescribeClustersResponseBody) SetExternalLoadbalancerId(v string) *DescribeClustersResponseBody {
	s.ExternalLoadbalancerId = &v
	return s
}

func (s *DescribeClustersResponseBody) SetInitVersion(v string) *DescribeClustersResponseBody {
	s.InitVersion = &v
	return s
}

func (s *DescribeClustersResponseBody) SetMasterUrl(v string) *DescribeClustersResponseBody {
	s.MasterUrl = &v
	return s
}

func (s *DescribeClustersResponseBody) SetMetaData(v string) *DescribeClustersResponseBody {
	s.MetaData = &v
	return s
}

func (s *DescribeClustersResponseBody) SetName(v string) *DescribeClustersResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeClustersResponseBody) SetNetworkMode(v string) *DescribeClustersResponseBody {
	s.NetworkMode = &v
	return s
}

func (s *DescribeClustersResponseBody) SetPrivateZone(v bool) *DescribeClustersResponseBody {
	s.PrivateZone = &v
	return s
}

func (s *DescribeClustersResponseBody) SetProfile(v string) *DescribeClustersResponseBody {
	s.Profile = &v
	return s
}

func (s *DescribeClustersResponseBody) SetRegionId(v string) *DescribeClustersResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeClustersResponseBody) SetResourceGroupId(v string) *DescribeClustersResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeClustersResponseBody) SetSecurityGroupId(v string) *DescribeClustersResponseBody {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeClustersResponseBody) SetSize(v int64) *DescribeClustersResponseBody {
	s.Size = &v
	return s
}

func (s *DescribeClustersResponseBody) SetState(v string) *DescribeClustersResponseBody {
	s.State = &v
	return s
}

func (s *DescribeClustersResponseBody) SetSubnetCidr(v string) *DescribeClustersResponseBody {
	s.SubnetCidr = &v
	return s
}

func (s *DescribeClustersResponseBody) SetTags(v []*DescribeClustersResponseBodyTags) *DescribeClustersResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeClustersResponseBody) SetUpdated(v string) *DescribeClustersResponseBody {
	s.Updated = &v
	return s
}

func (s *DescribeClustersResponseBody) SetVpcId(v string) *DescribeClustersResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeClustersResponseBody) SetVswitchCidr(v string) *DescribeClustersResponseBody {
	s.VswitchCidr = &v
	return s
}

func (s *DescribeClustersResponseBody) SetVswitchId(v string) *DescribeClustersResponseBody {
	s.VswitchId = &v
	return s
}

func (s *DescribeClustersResponseBody) SetWorkerRamRoleName(v string) *DescribeClustersResponseBody {
	s.WorkerRamRoleName = &v
	return s
}

func (s *DescribeClustersResponseBody) SetZoneId(v string) *DescribeClustersResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DescribeClustersResponseBody) Validate() error {
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

type DescribeClustersResponseBodyTags struct {
	// The key of the cluster tag.
	//
	// example:
	//
	// example.tag.key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of the cluster tag.
	//
	// example:
	//
	// example-tag-value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeClustersResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeClustersResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *DescribeClustersResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *DescribeClustersResponseBodyTags) SetKey(v string) *DescribeClustersResponseBodyTags {
	s.Key = &v
	return s
}

func (s *DescribeClustersResponseBodyTags) SetValue(v string) *DescribeClustersResponseBodyTags {
	s.Value = &v
	return s
}

func (s *DescribeClustersResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
