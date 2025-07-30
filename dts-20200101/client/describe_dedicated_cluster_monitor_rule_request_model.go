// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDedicatedClusterMonitorRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDedicatedClusterId(v string) *DescribeDedicatedClusterMonitorRuleRequest
	GetDedicatedClusterId() *string
	SetOwnerId(v string) *DescribeDedicatedClusterMonitorRuleRequest
	GetOwnerId() *string
	SetRegionId(v string) *DescribeDedicatedClusterMonitorRuleRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDedicatedClusterMonitorRuleRequest
	GetResourceGroupId() *string
}

type DescribeDedicatedClusterMonitorRuleRequest struct {
	// The ID of the cluster.
	//
	// example:
	//
	// dtsClustervcwn1oeyu5fx4yf
	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" xml:"DedicatedClusterId,omitempty"`
	OwnerId            *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeDedicatedClusterMonitorRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedClusterMonitorRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedClusterMonitorRuleRequest) GetDedicatedClusterId() *string {
	return s.DedicatedClusterId
}

func (s *DescribeDedicatedClusterMonitorRuleRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeDedicatedClusterMonitorRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDedicatedClusterMonitorRuleRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDedicatedClusterMonitorRuleRequest) SetDedicatedClusterId(v string) *DescribeDedicatedClusterMonitorRuleRequest {
	s.DedicatedClusterId = &v
	return s
}

func (s *DescribeDedicatedClusterMonitorRuleRequest) SetOwnerId(v string) *DescribeDedicatedClusterMonitorRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDedicatedClusterMonitorRuleRequest) SetRegionId(v string) *DescribeDedicatedClusterMonitorRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedClusterMonitorRuleRequest) SetResourceGroupId(v string) *DescribeDedicatedClusterMonitorRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDedicatedClusterMonitorRuleRequest) Validate() error {
	return dara.Validate(s)
}
