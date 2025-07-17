// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallManagedPrometheusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *InstallManagedPrometheusRequest
	GetClusterId() *string
	SetClusterName(v string) *InstallManagedPrometheusRequest
	GetClusterName() *string
	SetClusterType(v string) *InstallManagedPrometheusRequest
	GetClusterType() *string
	SetGrafanaInstanceId(v string) *InstallManagedPrometheusRequest
	GetGrafanaInstanceId() *string
	SetKubeConfig(v string) *InstallManagedPrometheusRequest
	GetKubeConfig() *string
	SetRegionId(v string) *InstallManagedPrometheusRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *InstallManagedPrometheusRequest
	GetResourceGroupId() *string
	SetSecurityGroupId(v string) *InstallManagedPrometheusRequest
	GetSecurityGroupId() *string
	SetVSwitchId(v string) *InstallManagedPrometheusRequest
	GetVSwitchId() *string
	SetVcExtraInfo(v string) *InstallManagedPrometheusRequest
	GetVcExtraInfo() *string
	SetVpcId(v string) *InstallManagedPrometheusRequest
	GetVpcId() *string
}

type InstallManagedPrometheusRequest struct {
	// The ID of the ASK cluster.
	//
	// example:
	//
	// cc7a37ee31aea4ed1a059eff8034b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the cluster. This parameter is required if the ClusterType parameter is set to ecs.
	//
	// example:
	//
	// prd-ecs
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The cluster type.
	//
	// Valid values:
	//
	// 	- ecs: ECS
	//
	// 	- one: ACK One
	//
	// 	- ask: ASK
	//
	// 	- pro: Container Monitoring Pro
	//
	// This parameter is required.
	//
	// example:
	//
	// ask
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The ID of the managed Grafana workspace that is associated with the cluster. If you set this parameter to free or leave this parameter empty, the cluster is associated with a shared Grafana workspace.
	//
	// example:
	//
	// grafana-bp1*****
	GrafanaInstanceId *string `json:"GrafanaInstanceId,omitempty" xml:"GrafanaInstanceId,omitempty"`
	// This parameter is not supported.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// -
	KubeConfig *string `json:"KubeConfig,omitempty" xml:"KubeConfig,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfmxyexli2****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the security group to which the cluster belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-bp1********
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The ID of the vSwitch that is used by the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1*********
	VSwitchId   *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VcExtraInfo *string `json:"VcExtraInfo,omitempty" xml:"VcExtraInfo,omitempty"`
	// The virtual private cloud (VPC) where the cluster resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-xxxxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s InstallManagedPrometheusRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallManagedPrometheusRequest) GoString() string {
	return s.String()
}

func (s *InstallManagedPrometheusRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *InstallManagedPrometheusRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *InstallManagedPrometheusRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *InstallManagedPrometheusRequest) GetGrafanaInstanceId() *string {
	return s.GrafanaInstanceId
}

func (s *InstallManagedPrometheusRequest) GetKubeConfig() *string {
	return s.KubeConfig
}

func (s *InstallManagedPrometheusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *InstallManagedPrometheusRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *InstallManagedPrometheusRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *InstallManagedPrometheusRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *InstallManagedPrometheusRequest) GetVcExtraInfo() *string {
	return s.VcExtraInfo
}

func (s *InstallManagedPrometheusRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *InstallManagedPrometheusRequest) SetClusterId(v string) *InstallManagedPrometheusRequest {
	s.ClusterId = &v
	return s
}

func (s *InstallManagedPrometheusRequest) SetClusterName(v string) *InstallManagedPrometheusRequest {
	s.ClusterName = &v
	return s
}

func (s *InstallManagedPrometheusRequest) SetClusterType(v string) *InstallManagedPrometheusRequest {
	s.ClusterType = &v
	return s
}

func (s *InstallManagedPrometheusRequest) SetGrafanaInstanceId(v string) *InstallManagedPrometheusRequest {
	s.GrafanaInstanceId = &v
	return s
}

func (s *InstallManagedPrometheusRequest) SetKubeConfig(v string) *InstallManagedPrometheusRequest {
	s.KubeConfig = &v
	return s
}

func (s *InstallManagedPrometheusRequest) SetRegionId(v string) *InstallManagedPrometheusRequest {
	s.RegionId = &v
	return s
}

func (s *InstallManagedPrometheusRequest) SetResourceGroupId(v string) *InstallManagedPrometheusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *InstallManagedPrometheusRequest) SetSecurityGroupId(v string) *InstallManagedPrometheusRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *InstallManagedPrometheusRequest) SetVSwitchId(v string) *InstallManagedPrometheusRequest {
	s.VSwitchId = &v
	return s
}

func (s *InstallManagedPrometheusRequest) SetVcExtraInfo(v string) *InstallManagedPrometheusRequest {
	s.VcExtraInfo = &v
	return s
}

func (s *InstallManagedPrometheusRequest) SetVpcId(v string) *InstallManagedPrometheusRequest {
	s.VpcId = &v
	return s
}

func (s *InstallManagedPrometheusRequest) Validate() error {
	return dara.Validate(s)
}
