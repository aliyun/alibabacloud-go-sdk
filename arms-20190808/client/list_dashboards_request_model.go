// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDashboardsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListDashboardsRequest
	GetClusterId() *string
	SetClusterType(v string) *ListDashboardsRequest
	GetClusterType() *string
	SetDashboardName(v string) *ListDashboardsRequest
	GetDashboardName() *string
	SetLanguage(v string) *ListDashboardsRequest
	GetLanguage() *string
	SetProduct(v string) *ListDashboardsRequest
	GetProduct() *string
	SetRecreateSwitch(v bool) *ListDashboardsRequest
	GetRecreateSwitch() *bool
	SetRegionId(v string) *ListDashboardsRequest
	GetRegionId() *string
	SetTitle(v string) *ListDashboardsRequest
	GetTitle() *string
}

type ListDashboardsRequest struct {
	// The ID of the ACK cluster.
	//
	// example:
	//
	// cc7a37ee31aea4ed1a059eff8034b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Valid values: ACK, ASK, cloud-product-prometheus, and Node. You can query the dashboards of a virtual cluster by specifying the cluster type. For InfluxDB, set this parameter to `cloud-product-prometheus`.
	//
	// example:
	//
	// Node
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The unique names of the dashboards. You can query dashboards by specifying their names. The **dashboard title*	- can be changed whereas the **dashboard name*	- cannot. You can specify multiple names and separate them with commas (,), for example, `k8s-event,k8s-overview`. A dashboard may have multiple versions. If you want to specify a version, you can add version information after the name, for example, `k8s-event:v1,k8s-overview:latest`.
	//
	// example:
	//
	// k8s-node-overview
	DashboardName *string `json:"DashboardName,omitempty" xml:"DashboardName,omitempty"`
	// The language of the returned Grafana dashboard. Valid values: en and zh. Default value: en.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The cloud service code. This parameter is required if you set the ClusterType parameter to `cloud-product-prometheus`. The following cloud services are available: Serverless App Engine, Microservices Engine, Message Queue for Apache RocketMQ, Lindorm, Message Queue for Apache Kafka, ApsaraDB for ClickHouse, Data Lake Analytics, Message Queue for RabbitMQ, ApsaraDB for MongoDB, Time Series Database (TSDB) for InfluxDB, MSE Cloud-native Gateway, Grafana Service, SchedulerX, Global Transaction Service, Enterprise Distributed Application Service, Machine Learning Platform for AI - Elastic Algorithm Service (EAS), Application High Availability Service, and Performance Testing.
	//
	// example:
	//
	// xxxx
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// Specifies whether to create or query a virtual cluster. This parameter provides backward compatibility.
	//
	// example:
	//
	// false
	RecreateSwitch *bool `json:"RecreateSwitch,omitempty" xml:"RecreateSwitch,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The dashboard title. The dashboard title can be changed. We recommend that you specify the **DashboardName*	- parameter.
	//
	// example:
	//
	// ApiServer
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListDashboardsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDashboardsRequest) GoString() string {
	return s.String()
}

func (s *ListDashboardsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListDashboardsRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *ListDashboardsRequest) GetDashboardName() *string {
	return s.DashboardName
}

func (s *ListDashboardsRequest) GetLanguage() *string {
	return s.Language
}

func (s *ListDashboardsRequest) GetProduct() *string {
	return s.Product
}

func (s *ListDashboardsRequest) GetRecreateSwitch() *bool {
	return s.RecreateSwitch
}

func (s *ListDashboardsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDashboardsRequest) GetTitle() *string {
	return s.Title
}

func (s *ListDashboardsRequest) SetClusterId(v string) *ListDashboardsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListDashboardsRequest) SetClusterType(v string) *ListDashboardsRequest {
	s.ClusterType = &v
	return s
}

func (s *ListDashboardsRequest) SetDashboardName(v string) *ListDashboardsRequest {
	s.DashboardName = &v
	return s
}

func (s *ListDashboardsRequest) SetLanguage(v string) *ListDashboardsRequest {
	s.Language = &v
	return s
}

func (s *ListDashboardsRequest) SetProduct(v string) *ListDashboardsRequest {
	s.Product = &v
	return s
}

func (s *ListDashboardsRequest) SetRecreateSwitch(v bool) *ListDashboardsRequest {
	s.RecreateSwitch = &v
	return s
}

func (s *ListDashboardsRequest) SetRegionId(v string) *ListDashboardsRequest {
	s.RegionId = &v
	return s
}

func (s *ListDashboardsRequest) SetTitle(v string) *ListDashboardsRequest {
	s.Title = &v
	return s
}

func (s *ListDashboardsRequest) Validate() error {
	return dara.Validate(s)
}
