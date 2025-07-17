// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetManagedPrometheusStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetManagedPrometheusStatusRequest
	GetClusterId() *string
	SetClusterType(v string) *GetManagedPrometheusStatusRequest
	GetClusterType() *string
	SetRegionId(v string) *GetManagedPrometheusStatusRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *GetManagedPrometheusStatusRequest
	GetResourceGroupId() *string
	SetVpcId(v string) *GetManagedPrometheusStatusRequest
	GetVpcId() *string
}

type GetManagedPrometheusStatusRequest struct {
	// The cluster ID. This parameter is required if the ClusterType parameter is set to ask.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc7a37ee31aea4ed1a059eff8034b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The type of the cluster. Valid values: ask and ecs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ask
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group id of the Prometheus instance.
	//
	// example:
	//
	// rg-acfmxyexli2****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the virtual private cloud (VPC). This parameter is required if the ClusterType parameter is set to ecs.
	//
	// example:
	//
	// vpc-***
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetManagedPrometheusStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetManagedPrometheusStatusRequest) GoString() string {
	return s.String()
}

func (s *GetManagedPrometheusStatusRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetManagedPrometheusStatusRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *GetManagedPrometheusStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetManagedPrometheusStatusRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetManagedPrometheusStatusRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *GetManagedPrometheusStatusRequest) SetClusterId(v string) *GetManagedPrometheusStatusRequest {
	s.ClusterId = &v
	return s
}

func (s *GetManagedPrometheusStatusRequest) SetClusterType(v string) *GetManagedPrometheusStatusRequest {
	s.ClusterType = &v
	return s
}

func (s *GetManagedPrometheusStatusRequest) SetRegionId(v string) *GetManagedPrometheusStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetManagedPrometheusStatusRequest) SetResourceGroupId(v string) *GetManagedPrometheusStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *GetManagedPrometheusStatusRequest) SetVpcId(v string) *GetManagedPrometheusStatusRequest {
	s.VpcId = &v
	return s
}

func (s *GetManagedPrometheusStatusRequest) Validate() error {
	return dara.Validate(s)
}
