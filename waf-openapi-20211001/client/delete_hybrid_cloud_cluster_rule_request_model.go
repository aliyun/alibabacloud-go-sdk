// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHybridCloudClusterRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterRuleResourceId(v string) *DeleteHybridCloudClusterRuleRequest
	GetClusterRuleResourceId() *string
	SetInstanceId(v string) *DeleteHybridCloudClusterRuleRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteHybridCloudClusterRuleRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DeleteHybridCloudClusterRuleRequest
	GetResourceManagerResourceGroupId() *string
}

type DeleteHybridCloudClusterRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hdbc-clusterrule-*******m0w
	ClusterRuleResourceId *string `json:"ClusterRuleResourceId,omitempty" xml:"ClusterRuleResourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0xldbqt****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DeleteHybridCloudClusterRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHybridCloudClusterRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteHybridCloudClusterRuleRequest) GetClusterRuleResourceId() *string {
	return s.ClusterRuleResourceId
}

func (s *DeleteHybridCloudClusterRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteHybridCloudClusterRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteHybridCloudClusterRuleRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DeleteHybridCloudClusterRuleRequest) SetClusterRuleResourceId(v string) *DeleteHybridCloudClusterRuleRequest {
	s.ClusterRuleResourceId = &v
	return s
}

func (s *DeleteHybridCloudClusterRuleRequest) SetInstanceId(v string) *DeleteHybridCloudClusterRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteHybridCloudClusterRuleRequest) SetRegionId(v string) *DeleteHybridCloudClusterRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteHybridCloudClusterRuleRequest) SetResourceManagerResourceGroupId(v string) *DeleteHybridCloudClusterRuleRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DeleteHybridCloudClusterRuleRequest) Validate() error {
	return dara.Validate(s)
}
