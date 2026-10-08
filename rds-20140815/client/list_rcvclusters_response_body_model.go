// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRCVClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListRCVClustersResponseBody
	GetRequestId() *string
	SetVClusters(v []*ListRCVClustersResponseBodyVClusters) *ListRCVClustersResponseBody
	GetVClusters() []*ListRCVClustersResponseBodyVClusters
}

type ListRCVClustersResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VClusters []*ListRCVClustersResponseBodyVClusters `json:"VClusters,omitempty" xml:"VClusters,omitempty" type:"Repeated"`
}

func (s ListRCVClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRCVClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListRCVClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRCVClustersResponseBody) GetVClusters() []*ListRCVClustersResponseBodyVClusters {
	return s.VClusters
}

func (s *ListRCVClustersResponseBody) SetRequestId(v string) *ListRCVClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRCVClustersResponseBody) SetVClusters(v []*ListRCVClustersResponseBodyVClusters) *ListRCVClustersResponseBody {
	s.VClusters = v
	return s
}

func (s *ListRCVClustersResponseBody) Validate() error {
	if s.VClusters != nil {
		for _, item := range s.VClusters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRCVClustersResponseBodyVClusters struct {
	ClusterId                   *string                                            `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName                 *string                                            `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	InstanceCount               *int64                                             `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	MysqlOperator               *ListRCVClustersResponseBodyVClustersMysqlOperator `json:"MysqlOperator,omitempty" xml:"MysqlOperator,omitempty" type:"Struct"`
	RegionId                    *string                                            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status                      *string                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	SupportDiskPerformanceLevel []*string                                          `json:"SupportDiskPerformanceLevel,omitempty" xml:"SupportDiskPerformanceLevel,omitempty" type:"Repeated"`
	VpcId                       *string                                            `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListRCVClustersResponseBodyVClusters) String() string {
	return dara.Prettify(s)
}

func (s ListRCVClustersResponseBodyVClusters) GoString() string {
	return s.String()
}

func (s *ListRCVClustersResponseBodyVClusters) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListRCVClustersResponseBodyVClusters) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListRCVClustersResponseBodyVClusters) GetInstanceCount() *int64 {
	return s.InstanceCount
}

func (s *ListRCVClustersResponseBodyVClusters) GetMysqlOperator() *ListRCVClustersResponseBodyVClustersMysqlOperator {
	return s.MysqlOperator
}

func (s *ListRCVClustersResponseBodyVClusters) GetRegionId() *string {
	return s.RegionId
}

func (s *ListRCVClustersResponseBodyVClusters) GetStatus() *string {
	return s.Status
}

func (s *ListRCVClustersResponseBodyVClusters) GetSupportDiskPerformanceLevel() []*string {
	return s.SupportDiskPerformanceLevel
}

func (s *ListRCVClustersResponseBodyVClusters) GetVpcId() *string {
	return s.VpcId
}

func (s *ListRCVClustersResponseBodyVClusters) SetClusterId(v string) *ListRCVClustersResponseBodyVClusters {
	s.ClusterId = &v
	return s
}

func (s *ListRCVClustersResponseBodyVClusters) SetClusterName(v string) *ListRCVClustersResponseBodyVClusters {
	s.ClusterName = &v
	return s
}

func (s *ListRCVClustersResponseBodyVClusters) SetInstanceCount(v int64) *ListRCVClustersResponseBodyVClusters {
	s.InstanceCount = &v
	return s
}

func (s *ListRCVClustersResponseBodyVClusters) SetMysqlOperator(v *ListRCVClustersResponseBodyVClustersMysqlOperator) *ListRCVClustersResponseBodyVClusters {
	s.MysqlOperator = v
	return s
}

func (s *ListRCVClustersResponseBodyVClusters) SetRegionId(v string) *ListRCVClustersResponseBodyVClusters {
	s.RegionId = &v
	return s
}

func (s *ListRCVClustersResponseBodyVClusters) SetStatus(v string) *ListRCVClustersResponseBodyVClusters {
	s.Status = &v
	return s
}

func (s *ListRCVClustersResponseBodyVClusters) SetSupportDiskPerformanceLevel(v []*string) *ListRCVClustersResponseBodyVClusters {
	s.SupportDiskPerformanceLevel = v
	return s
}

func (s *ListRCVClustersResponseBodyVClusters) SetVpcId(v string) *ListRCVClustersResponseBodyVClusters {
	s.VpcId = &v
	return s
}

func (s *ListRCVClustersResponseBodyVClusters) Validate() error {
	if s.MysqlOperator != nil {
		if err := s.MysqlOperator.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRCVClustersResponseBodyVClustersMysqlOperator struct {
	DashboardPublicEndpoint *string `json:"DashboardPublicEndpoint,omitempty" xml:"DashboardPublicEndpoint,omitempty"`
	DashboardUsername       *string `json:"DashboardUsername,omitempty" xml:"DashboardUsername,omitempty"`
	DashboardVpcEndpoint    *string `json:"DashboardVpcEndpoint,omitempty" xml:"DashboardVpcEndpoint,omitempty"`
	DeployTime              *string `json:"DeployTime,omitempty" xml:"DeployTime,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListRCVClustersResponseBodyVClustersMysqlOperator) String() string {
	return dara.Prettify(s)
}

func (s ListRCVClustersResponseBodyVClustersMysqlOperator) GoString() string {
	return s.String()
}

func (s *ListRCVClustersResponseBodyVClustersMysqlOperator) GetDashboardPublicEndpoint() *string {
	return s.DashboardPublicEndpoint
}

func (s *ListRCVClustersResponseBodyVClustersMysqlOperator) GetDashboardUsername() *string {
	return s.DashboardUsername
}

func (s *ListRCVClustersResponseBodyVClustersMysqlOperator) GetDashboardVpcEndpoint() *string {
	return s.DashboardVpcEndpoint
}

func (s *ListRCVClustersResponseBodyVClustersMysqlOperator) GetDeployTime() *string {
	return s.DeployTime
}

func (s *ListRCVClustersResponseBodyVClustersMysqlOperator) GetStatus() *string {
	return s.Status
}

func (s *ListRCVClustersResponseBodyVClustersMysqlOperator) SetDashboardPublicEndpoint(v string) *ListRCVClustersResponseBodyVClustersMysqlOperator {
	s.DashboardPublicEndpoint = &v
	return s
}

func (s *ListRCVClustersResponseBodyVClustersMysqlOperator) SetDashboardUsername(v string) *ListRCVClustersResponseBodyVClustersMysqlOperator {
	s.DashboardUsername = &v
	return s
}

func (s *ListRCVClustersResponseBodyVClustersMysqlOperator) SetDashboardVpcEndpoint(v string) *ListRCVClustersResponseBodyVClustersMysqlOperator {
	s.DashboardVpcEndpoint = &v
	return s
}

func (s *ListRCVClustersResponseBodyVClustersMysqlOperator) SetDeployTime(v string) *ListRCVClustersResponseBodyVClustersMysqlOperator {
	s.DeployTime = &v
	return s
}

func (s *ListRCVClustersResponseBodyVClustersMysqlOperator) SetStatus(v string) *ListRCVClustersResponseBodyVClustersMysqlOperator {
	s.Status = &v
	return s
}

func (s *ListRCVClustersResponseBodyVClustersMysqlOperator) Validate() error {
	return dara.Validate(s)
}
