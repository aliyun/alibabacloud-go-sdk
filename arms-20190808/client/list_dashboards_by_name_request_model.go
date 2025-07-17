// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDashboardsByNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListDashboardsByNameRequest
	GetClusterId() *string
	SetClusterType(v string) *ListDashboardsByNameRequest
	GetClusterType() *string
	SetDashBoardName(v string) *ListDashboardsByNameRequest
	GetDashBoardName() *string
	SetDashBoardVersion(v string) *ListDashboardsByNameRequest
	GetDashBoardVersion() *string
	SetDataSourceType(v string) *ListDashboardsByNameRequest
	GetDataSourceType() *string
	SetGroupName(v string) *ListDashboardsByNameRequest
	GetGroupName() *string
	SetOnlyQuery(v bool) *ListDashboardsByNameRequest
	GetOnlyQuery() *bool
	SetProductCode(v string) *ListDashboardsByNameRequest
	GetProductCode() *string
	SetRegionId(v string) *ListDashboardsByNameRequest
	GetRegionId() *string
}

type ListDashboardsByNameRequest struct {
	// The ID of the cluster. If the ClusterType parameter is not set to `cloud-product-prometheus` or `cms-enterprise-prometheus`, you must specify the ClusterId parameter.
	//
	// example:
	//
	// cc7a37ee31aea4ed1a059eff8034b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster type. Valid values:
	//
	// 	- vpc-prometheus
	//
	// 	- cloud-product-prometheus
	//
	// 	- cms-enterprise-prometheus
	//
	// 	- ExternalKubernetes
	//
	// 	- Ask
	//
	// 	- Kubernetes
	//
	// 	- ManagedKubernetes
	//
	// 	- remote-write-prometheus
	//
	// 	- GlobalViewV2
	//
	// example:
	//
	// cloud-product-prometheus
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The name of the dashboard.
	//
	// example:
	//
	// edas-ingress-url
	DashBoardName *string `json:"DashBoardName,omitempty" xml:"DashBoardName,omitempty"`
	// The version of the dashboard.
	//
	// example:
	//
	// latest
	DashBoardVersion *string `json:"DashBoardVersion,omitempty" xml:"DashBoardVersion,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- loki
	//
	// 	- prometheus
	//
	// example:
	//
	// loki
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The name of the dashboard group.
	//
	// example:
	//
	// EDAS
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Specifies whether to display the Grafana dashboard only in the Application Real-Time Monitoring Service (ARMS) console.
	//
	// example:
	//
	// true
	OnlyQuery *bool `json:"OnlyQuery,omitempty" xml:"OnlyQuery,omitempty"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// example:
	//
	// edas
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListDashboardsByNameRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDashboardsByNameRequest) GoString() string {
	return s.String()
}

func (s *ListDashboardsByNameRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListDashboardsByNameRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *ListDashboardsByNameRequest) GetDashBoardName() *string {
	return s.DashBoardName
}

func (s *ListDashboardsByNameRequest) GetDashBoardVersion() *string {
	return s.DashBoardVersion
}

func (s *ListDashboardsByNameRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *ListDashboardsByNameRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ListDashboardsByNameRequest) GetOnlyQuery() *bool {
	return s.OnlyQuery
}

func (s *ListDashboardsByNameRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListDashboardsByNameRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDashboardsByNameRequest) SetClusterId(v string) *ListDashboardsByNameRequest {
	s.ClusterId = &v
	return s
}

func (s *ListDashboardsByNameRequest) SetClusterType(v string) *ListDashboardsByNameRequest {
	s.ClusterType = &v
	return s
}

func (s *ListDashboardsByNameRequest) SetDashBoardName(v string) *ListDashboardsByNameRequest {
	s.DashBoardName = &v
	return s
}

func (s *ListDashboardsByNameRequest) SetDashBoardVersion(v string) *ListDashboardsByNameRequest {
	s.DashBoardVersion = &v
	return s
}

func (s *ListDashboardsByNameRequest) SetDataSourceType(v string) *ListDashboardsByNameRequest {
	s.DataSourceType = &v
	return s
}

func (s *ListDashboardsByNameRequest) SetGroupName(v string) *ListDashboardsByNameRequest {
	s.GroupName = &v
	return s
}

func (s *ListDashboardsByNameRequest) SetOnlyQuery(v bool) *ListDashboardsByNameRequest {
	s.OnlyQuery = &v
	return s
}

func (s *ListDashboardsByNameRequest) SetProductCode(v string) *ListDashboardsByNameRequest {
	s.ProductCode = &v
	return s
}

func (s *ListDashboardsByNameRequest) SetRegionId(v string) *ListDashboardsByNameRequest {
	s.RegionId = &v
	return s
}

func (s *ListDashboardsByNameRequest) Validate() error {
	return dara.Validate(s)
}
