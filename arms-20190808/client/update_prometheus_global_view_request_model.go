// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrometheusGlobalViewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllSubClustersSuccess(v bool) *UpdatePrometheusGlobalViewRequest
	GetAllSubClustersSuccess() *bool
	SetClusterId(v string) *UpdatePrometheusGlobalViewRequest
	GetClusterId() *string
	SetGroupName(v string) *UpdatePrometheusGlobalViewRequest
	GetGroupName() *string
	SetMostRegionId(v string) *UpdatePrometheusGlobalViewRequest
	GetMostRegionId() *string
	SetRegionId(v string) *UpdatePrometheusGlobalViewRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *UpdatePrometheusGlobalViewRequest
	GetResourceGroupId() *string
	SetSubClustersJson(v string) *UpdatePrometheusGlobalViewRequest
	GetSubClustersJson() *string
}

type UpdatePrometheusGlobalViewRequest struct {
	// To edit a GlobalView aggregated instance, do you require all passed child instances to be verified successfully before creating a GlobalView instance (optional, default to false):
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// true
	AllSubClustersSuccess *bool `json:"AllSubClustersSuccess,omitempty" xml:"AllSubClustersSuccess,omitempty"`
	// The ID of the Prometheus instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// global****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the global aggregation instance.
	//
	// example:
	//
	// zyGlobalView
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The region ID of the global aggregation instance.
	//
	// example:
	//
	// cn-hangzhou
	MostRegionId *string `json:"MostRegionId,omitempty" xml:"MostRegionId,omitempty"`
	// The ID of the region in which the Prometheus instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Prometheus instance belongs.
	//
	// example:
	//
	// rg-acfmxyexli2****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The data sources of the Prometheus instance for GlobalView.
	//
	// This parameter is required.
	//
	// example:
	//
	// [ { "headers":{ }, "regionId":"cn-hangzhou", "sourceType":"AlibabaPrometheus", "extras":{ }, "clusterId":"c39a1048921e04f***********", "sourceName":"arms-luyao-test", "dataSource":"", "userId":"1672753***********" }, { "headers":{ }, "regionId":"cn-beijing", "sourceType":"AlibabaPrometheus", "extras":{ }, "clusterId":"c6b6485496d5b40***********", "sourceName":"agent-321-test", "dataSource":"", "userId":"1672753***********" }, { "headers":{ }, "regionId":"cn-zhangjiakou", "sourceType":"AlibabaPrometheus", "extras":{ }, "clusterId":"c261a4f3200c446***********", "sourceName":"zaifeng-cardinality-01", "dataSource":"", "userId":"1672753***********" } ]
	SubClustersJson *string `json:"SubClustersJson,omitempty" xml:"SubClustersJson,omitempty"`
}

func (s UpdatePrometheusGlobalViewRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusGlobalViewRequest) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusGlobalViewRequest) GetAllSubClustersSuccess() *bool {
	return s.AllSubClustersSuccess
}

func (s *UpdatePrometheusGlobalViewRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdatePrometheusGlobalViewRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdatePrometheusGlobalViewRequest) GetMostRegionId() *string {
	return s.MostRegionId
}

func (s *UpdatePrometheusGlobalViewRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdatePrometheusGlobalViewRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdatePrometheusGlobalViewRequest) GetSubClustersJson() *string {
	return s.SubClustersJson
}

func (s *UpdatePrometheusGlobalViewRequest) SetAllSubClustersSuccess(v bool) *UpdatePrometheusGlobalViewRequest {
	s.AllSubClustersSuccess = &v
	return s
}

func (s *UpdatePrometheusGlobalViewRequest) SetClusterId(v string) *UpdatePrometheusGlobalViewRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdatePrometheusGlobalViewRequest) SetGroupName(v string) *UpdatePrometheusGlobalViewRequest {
	s.GroupName = &v
	return s
}

func (s *UpdatePrometheusGlobalViewRequest) SetMostRegionId(v string) *UpdatePrometheusGlobalViewRequest {
	s.MostRegionId = &v
	return s
}

func (s *UpdatePrometheusGlobalViewRequest) SetRegionId(v string) *UpdatePrometheusGlobalViewRequest {
	s.RegionId = &v
	return s
}

func (s *UpdatePrometheusGlobalViewRequest) SetResourceGroupId(v string) *UpdatePrometheusGlobalViewRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdatePrometheusGlobalViewRequest) SetSubClustersJson(v string) *UpdatePrometheusGlobalViewRequest {
	s.SubClustersJson = &v
	return s
}

func (s *UpdatePrometheusGlobalViewRequest) Validate() error {
	return dara.Validate(s)
}
