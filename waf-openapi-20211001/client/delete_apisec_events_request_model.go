// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApisecEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteApisecEventsRequest
	GetClusterId() *string
	SetEventIds(v []*string) *DeleteApisecEventsRequest
	GetEventIds() []*string
	SetInstanceId(v string) *DeleteApisecEventsRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteApisecEventsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DeleteApisecEventsRequest
	GetResourceManagerResourceGroupId() *string
}

type DeleteApisecEventsRequest struct {
	// The ID of the hybrid cloud cluster.
	//
	// >For hybrid cloud scenarios only, you can call the [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) operation to query the hybrid cloud clusters.
	//
	// example:
	//
	// 428
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The IDs of the security events.
	//
	// This parameter is required.
	EventIds []*string `json:"EventIds,omitempty" xml:"EventIds,omitempty" type:"Repeated"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-g4t*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: the Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 阿里云资源组ID。
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DeleteApisecEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApisecEventsRequest) GoString() string {
	return s.String()
}

func (s *DeleteApisecEventsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteApisecEventsRequest) GetEventIds() []*string {
	return s.EventIds
}

func (s *DeleteApisecEventsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteApisecEventsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteApisecEventsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DeleteApisecEventsRequest) SetClusterId(v string) *DeleteApisecEventsRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteApisecEventsRequest) SetEventIds(v []*string) *DeleteApisecEventsRequest {
	s.EventIds = v
	return s
}

func (s *DeleteApisecEventsRequest) SetInstanceId(v string) *DeleteApisecEventsRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteApisecEventsRequest) SetRegionId(v string) *DeleteApisecEventsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteApisecEventsRequest) SetResourceManagerResourceGroupId(v string) *DeleteApisecEventsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DeleteApisecEventsRequest) Validate() error {
	return dara.Validate(s)
}
