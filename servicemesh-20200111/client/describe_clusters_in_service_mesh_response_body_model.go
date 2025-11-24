// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClustersInServiceMeshResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusters(v []*DescribeClustersInServiceMeshResponseBodyClusters) *DescribeClustersInServiceMeshResponseBody
	GetClusters() []*DescribeClustersInServiceMeshResponseBodyClusters
	SetRequestId(v string) *DescribeClustersInServiceMeshResponseBody
	GetRequestId() *string
}

type DescribeClustersInServiceMeshResponseBody struct {
	// The list of the clusters in the ASM instance.
	Clusters []*DescribeClustersInServiceMeshResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 31d3a0f0-07ed-4f6e-9004-1804498c****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClustersInServiceMeshResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClustersInServiceMeshResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClustersInServiceMeshResponseBody) GetClusters() []*DescribeClustersInServiceMeshResponseBodyClusters {
	return s.Clusters
}

func (s *DescribeClustersInServiceMeshResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClustersInServiceMeshResponseBody) SetClusters(v []*DescribeClustersInServiceMeshResponseBodyClusters) *DescribeClustersInServiceMeshResponseBody {
	s.Clusters = v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBody) SetRequestId(v string) *DescribeClustersInServiceMeshResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBody) Validate() error {
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

type DescribeClustersInServiceMeshResponseBodyClusters struct {
	// The configurations of access log collection.
	AccessLogDashboards []*DescribeClustersInServiceMeshResponseBodyClustersAccessLogDashboards `json:"AccessLogDashboards,omitempty" xml:"AccessLogDashboards,omitempty" type:"Repeated"`
	// The domain name of the cluster.
	//
	// example:
	//
	// example.com
	ClusterDomain *string `json:"ClusterDomain,omitempty" xml:"ClusterDomain,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// c80f45444b3da447da60a911390c2****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The type of the cluster.
	//
	// example:
	//
	// Ask
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The time when the cluster was created.
	//
	// example:
	//
	// 2020-05-12T15:38:16+08:00
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The error message.
	//
	// example:
	//
	// ,
	ErrorMessage       *string                                                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	GuestClusterConfig *DescribeClustersInServiceMeshResponseBodyClustersGuestClusterConfig `json:"GuestClusterConfig,omitempty" xml:"GuestClusterConfig,omitempty" type:"Struct"`
	// Indicates whether the Logtail component is installed in the cluster. Valid values:
	//
	// 	- `logtail_installed`: The Logtail component is installed.
	//
	// \\-`logtail_uninstalled`: The Logtail component is not installed.
	//
	// 	- `logtail_state_get_error`: The Logtail component failed to be installed.
	//
	// example:
	//
	// logtail_installed
	LogtailInstalledState *string `json:"LogtailInstalledState,omitempty" xml:"LogtailInstalledState,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// ask1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region in which the cluster resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-bp197668l6iupljy****
	SgId *string `json:"SgId,omitempty" xml:"SgId,omitempty"`
	// The status of the cluster.
	//
	// example:
	//
	// running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The time when the cluster was last modified.
	//
	// example:
	//
	// 2020-05-12T15:38:16+08:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The version number of the cluster.
	//
	// example:
	//
	// v1.16.6-aliyun.1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-2zew0rajjkmxy2369****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeClustersInServiceMeshResponseBodyClusters) String() string {
	return dara.Prettify(s)
}

func (s DescribeClustersInServiceMeshResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) GetAccessLogDashboards() []*DescribeClustersInServiceMeshResponseBodyClustersAccessLogDashboards {
	return s.AccessLogDashboards
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) GetClusterDomain() *string {
	return s.ClusterDomain
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) GetClusterType() *string {
	return s.ClusterType
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) GetGuestClusterConfig() *DescribeClustersInServiceMeshResponseBodyClustersGuestClusterConfig {
	return s.GuestClusterConfig
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) GetLogtailInstalledState() *string {
	return s.LogtailInstalledState
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) GetName() *string {
	return s.Name
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) GetSgId() *string {
	return s.SgId
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) GetState() *string {
	return s.State
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) GetVersion() *string {
	return s.Version
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) SetAccessLogDashboards(v []*DescribeClustersInServiceMeshResponseBodyClustersAccessLogDashboards) *DescribeClustersInServiceMeshResponseBodyClusters {
	s.AccessLogDashboards = v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) SetClusterDomain(v string) *DescribeClustersInServiceMeshResponseBodyClusters {
	s.ClusterDomain = &v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) SetClusterId(v string) *DescribeClustersInServiceMeshResponseBodyClusters {
	s.ClusterId = &v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) SetClusterType(v string) *DescribeClustersInServiceMeshResponseBodyClusters {
	s.ClusterType = &v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) SetCreationTime(v string) *DescribeClustersInServiceMeshResponseBodyClusters {
	s.CreationTime = &v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) SetErrorMessage(v string) *DescribeClustersInServiceMeshResponseBodyClusters {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) SetGuestClusterConfig(v *DescribeClustersInServiceMeshResponseBodyClustersGuestClusterConfig) *DescribeClustersInServiceMeshResponseBodyClusters {
	s.GuestClusterConfig = v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) SetLogtailInstalledState(v string) *DescribeClustersInServiceMeshResponseBodyClusters {
	s.LogtailInstalledState = &v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) SetName(v string) *DescribeClustersInServiceMeshResponseBodyClusters {
	s.Name = &v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) SetRegionId(v string) *DescribeClustersInServiceMeshResponseBodyClusters {
	s.RegionId = &v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) SetSgId(v string) *DescribeClustersInServiceMeshResponseBodyClusters {
	s.SgId = &v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) SetState(v string) *DescribeClustersInServiceMeshResponseBodyClusters {
	s.State = &v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) SetUpdateTime(v string) *DescribeClustersInServiceMeshResponseBodyClusters {
	s.UpdateTime = &v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) SetVersion(v string) *DescribeClustersInServiceMeshResponseBodyClusters {
	s.Version = &v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) SetVpcId(v string) *DescribeClustersInServiceMeshResponseBodyClusters {
	s.VpcId = &v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBodyClusters) Validate() error {
	if s.AccessLogDashboards != nil {
		for _, item := range s.AccessLogDashboards {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.GuestClusterConfig != nil {
		if err := s.GuestClusterConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClustersInServiceMeshResponseBodyClustersAccessLogDashboards struct {
	// The name of the dashboard for access logs.
	//
	// example:
	//
	// mesh-access-log_details_cn
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The URL of the dashboard for access logs.
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeClustersInServiceMeshResponseBodyClustersAccessLogDashboards) String() string {
	return dara.Prettify(s)
}

func (s DescribeClustersInServiceMeshResponseBodyClustersAccessLogDashboards) GoString() string {
	return s.String()
}

func (s *DescribeClustersInServiceMeshResponseBodyClustersAccessLogDashboards) GetTitle() *string {
	return s.Title
}

func (s *DescribeClustersInServiceMeshResponseBodyClustersAccessLogDashboards) GetUrl() *string {
	return s.Url
}

func (s *DescribeClustersInServiceMeshResponseBodyClustersAccessLogDashboards) SetTitle(v string) *DescribeClustersInServiceMeshResponseBodyClustersAccessLogDashboards {
	s.Title = &v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBodyClustersAccessLogDashboards) SetUrl(v string) *DescribeClustersInServiceMeshResponseBodyClustersAccessLogDashboards {
	s.Url = &v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBodyClustersAccessLogDashboards) Validate() error {
	return dara.Validate(s)
}

type DescribeClustersInServiceMeshResponseBodyClustersGuestClusterConfig struct {
	SMC *DescribeClustersInServiceMeshResponseBodyClustersGuestClusterConfigSMC `json:"SMC,omitempty" xml:"SMC,omitempty" type:"Struct"`
}

func (s DescribeClustersInServiceMeshResponseBodyClustersGuestClusterConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeClustersInServiceMeshResponseBodyClustersGuestClusterConfig) GoString() string {
	return s.String()
}

func (s *DescribeClustersInServiceMeshResponseBodyClustersGuestClusterConfig) GetSMC() *DescribeClustersInServiceMeshResponseBodyClustersGuestClusterConfigSMC {
	return s.SMC
}

func (s *DescribeClustersInServiceMeshResponseBodyClustersGuestClusterConfig) SetSMC(v *DescribeClustersInServiceMeshResponseBodyClustersGuestClusterConfigSMC) *DescribeClustersInServiceMeshResponseBodyClustersGuestClusterConfig {
	s.SMC = v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBodyClustersGuestClusterConfig) Validate() error {
	if s.SMC != nil {
		if err := s.SMC.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClustersInServiceMeshResponseBodyClustersGuestClusterConfigSMC struct {
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s DescribeClustersInServiceMeshResponseBodyClustersGuestClusterConfigSMC) String() string {
	return dara.Prettify(s)
}

func (s DescribeClustersInServiceMeshResponseBodyClustersGuestClusterConfigSMC) GoString() string {
	return s.String()
}

func (s *DescribeClustersInServiceMeshResponseBodyClustersGuestClusterConfigSMC) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeClustersInServiceMeshResponseBodyClustersGuestClusterConfigSMC) SetEnabled(v bool) *DescribeClustersInServiceMeshResponseBodyClustersGuestClusterConfigSMC {
	s.Enabled = &v
	return s
}

func (s *DescribeClustersInServiceMeshResponseBodyClustersGuestClusterConfigSMC) Validate() error {
	return dara.Validate(s)
}
