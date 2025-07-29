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
	return dara.Validate(s)
}

type DescribeClustersResponseBody struct {
	ClusterId              *string                             `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	ClusterType            *string                             `json:"cluster_type,omitempty" xml:"cluster_type,omitempty"`
	Created                *string                             `json:"created,omitempty" xml:"created,omitempty"`
	CurrentVersion         *string                             `json:"current_version,omitempty" xml:"current_version,omitempty"`
	DataDiskCategory       *string                             `json:"data_disk_category,omitempty" xml:"data_disk_category,omitempty"`
	DataDiskSize           *int64                              `json:"data_disk_size,omitempty" xml:"data_disk_size,omitempty"`
	DeletionProtection     *bool                               `json:"deletion_protection,omitempty" xml:"deletion_protection,omitempty"`
	DockerVersion          *string                             `json:"docker_version,omitempty" xml:"docker_version,omitempty"`
	ExternalLoadbalancerId *string                             `json:"external_loadbalancer_id,omitempty" xml:"external_loadbalancer_id,omitempty"`
	InitVersion            *string                             `json:"init_version,omitempty" xml:"init_version,omitempty"`
	MasterUrl              *string                             `json:"master_url,omitempty" xml:"master_url,omitempty"`
	MetaData               *string                             `json:"meta_data,omitempty" xml:"meta_data,omitempty"`
	Name                   *string                             `json:"name,omitempty" xml:"name,omitempty"`
	NetworkMode            *string                             `json:"network_mode,omitempty" xml:"network_mode,omitempty"`
	PrivateZone            *bool                               `json:"private_zone,omitempty" xml:"private_zone,omitempty"`
	Profile                *string                             `json:"profile,omitempty" xml:"profile,omitempty"`
	RegionId               *string                             `json:"region_id,omitempty" xml:"region_id,omitempty"`
	ResourceGroupId        *string                             `json:"resource_group_id,omitempty" xml:"resource_group_id,omitempty"`
	SecurityGroupId        *string                             `json:"security_group_id,omitempty" xml:"security_group_id,omitempty"`
	Size                   *int64                              `json:"size,omitempty" xml:"size,omitempty"`
	State                  *string                             `json:"state,omitempty" xml:"state,omitempty"`
	SubnetCidr             *string                             `json:"subnet_cidr,omitempty" xml:"subnet_cidr,omitempty"`
	Tags                   []*DescribeClustersResponseBodyTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	Updated                *string                             `json:"updated,omitempty" xml:"updated,omitempty"`
	VpcId                  *string                             `json:"vpc_id,omitempty" xml:"vpc_id,omitempty"`
	VswitchCidr            *string                             `json:"vswitch_cidr,omitempty" xml:"vswitch_cidr,omitempty"`
	VswitchId              *string                             `json:"vswitch_id,omitempty" xml:"vswitch_id,omitempty"`
	WorkerRamRoleName      *string                             `json:"worker_ram_role_name,omitempty" xml:"worker_ram_role_name,omitempty"`
	ZoneId                 *string                             `json:"zone_id,omitempty" xml:"zone_id,omitempty"`
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
	return dara.Validate(s)
}

type DescribeClustersResponseBodyTags struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
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
