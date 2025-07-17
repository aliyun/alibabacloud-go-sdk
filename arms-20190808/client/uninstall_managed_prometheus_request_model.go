// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallManagedPrometheusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UninstallManagedPrometheusRequest
	GetClusterId() *string
	SetClusterType(v string) *UninstallManagedPrometheusRequest
	GetClusterType() *string
	SetRegionId(v string) *UninstallManagedPrometheusRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *UninstallManagedPrometheusRequest
	GetResourceGroupId() *string
	SetVpcId(v string) *UninstallManagedPrometheusRequest
	GetVpcId() *string
}

type UninstallManagedPrometheusRequest struct {
	// The ID of the Container Service for Kubernetes (ACK) cluster. This parameter is required when the ClusterType parameter is set to ask or one.
	//
	// example:
	//
	// cc7a37ee31aea4ed1a059eff8034b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster type. Valid values: ask, ecs, and one.
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
	// The ID of the resource group to which the Prometheus instance belongs.
	//
	// example:
	//
	// rg-acfmxyexli2****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the virtual private cloud (VPC) where the cluster resides.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// vpc-rpn**********
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s UninstallManagedPrometheusRequest) String() string {
	return dara.Prettify(s)
}

func (s UninstallManagedPrometheusRequest) GoString() string {
	return s.String()
}

func (s *UninstallManagedPrometheusRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UninstallManagedPrometheusRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *UninstallManagedPrometheusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UninstallManagedPrometheusRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UninstallManagedPrometheusRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *UninstallManagedPrometheusRequest) SetClusterId(v string) *UninstallManagedPrometheusRequest {
	s.ClusterId = &v
	return s
}

func (s *UninstallManagedPrometheusRequest) SetClusterType(v string) *UninstallManagedPrometheusRequest {
	s.ClusterType = &v
	return s
}

func (s *UninstallManagedPrometheusRequest) SetRegionId(v string) *UninstallManagedPrometheusRequest {
	s.RegionId = &v
	return s
}

func (s *UninstallManagedPrometheusRequest) SetResourceGroupId(v string) *UninstallManagedPrometheusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UninstallManagedPrometheusRequest) SetVpcId(v string) *UninstallManagedPrometheusRequest {
	s.VpcId = &v
	return s
}

func (s *UninstallManagedPrometheusRequest) Validate() error {
	return dara.Validate(s)
}
