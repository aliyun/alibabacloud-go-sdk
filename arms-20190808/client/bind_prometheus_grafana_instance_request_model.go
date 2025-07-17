// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindPrometheusGrafanaInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *BindPrometheusGrafanaInstanceRequest
	GetClusterId() *string
	SetGrafanaInstanceId(v string) *BindPrometheusGrafanaInstanceRequest
	GetGrafanaInstanceId() *string
	SetRegionId(v string) *BindPrometheusGrafanaInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *BindPrometheusGrafanaInstanceRequest
	GetResourceGroupId() *string
}

type BindPrometheusGrafanaInstanceRequest struct {
	// The ID of the Prometheus instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc7a37ee31aea4ed1a059eff8034b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the Grafana workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// grafana-bp1*****
	GrafanaInstanceId *string `json:"GrafanaInstanceId,omitempty" xml:"GrafanaInstanceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Prometheus instance belongs.
	//
	// example:
	//
	// rg-acfmxyexli2****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s BindPrometheusGrafanaInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s BindPrometheusGrafanaInstanceRequest) GoString() string {
	return s.String()
}

func (s *BindPrometheusGrafanaInstanceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *BindPrometheusGrafanaInstanceRequest) GetGrafanaInstanceId() *string {
	return s.GrafanaInstanceId
}

func (s *BindPrometheusGrafanaInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *BindPrometheusGrafanaInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *BindPrometheusGrafanaInstanceRequest) SetClusterId(v string) *BindPrometheusGrafanaInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *BindPrometheusGrafanaInstanceRequest) SetGrafanaInstanceId(v string) *BindPrometheusGrafanaInstanceRequest {
	s.GrafanaInstanceId = &v
	return s
}

func (s *BindPrometheusGrafanaInstanceRequest) SetRegionId(v string) *BindPrometheusGrafanaInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *BindPrometheusGrafanaInstanceRequest) SetResourceGroupId(v string) *BindPrometheusGrafanaInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *BindPrometheusGrafanaInstanceRequest) Validate() error {
	return dara.Validate(s)
}
