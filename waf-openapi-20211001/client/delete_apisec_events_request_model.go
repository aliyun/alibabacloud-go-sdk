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
	SetEventScope(v string) *DeleteApisecEventsRequest
	GetEventScope() *string
	SetInstanceId(v string) *DeleteApisecEventsRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteApisecEventsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DeleteApisecEventsRequest
	GetResourceManagerResourceGroupId() *string
}

type DeleteApisecEventsRequest struct {
	// The hybrid cloud cluster ID.
	//
	// > This parameter applies only to hybrid cloud scenarios. You can call [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) to obtain hybrid cloud cluster information.
	//
	// example:
	//
	// 428
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The list of API security event IDs.
	//
	// This parameter is required.
	EventIds []*string `json:"EventIds,omitempty" xml:"EventIds,omitempty" type:"Repeated"`
	// The dimension of the security event.
	//
	// example:
	//
	// ip
	EventScope *string `json:"EventScope,omitempty" xml:"EventScope,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to query the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-g4t*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance is deployed. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Alibaba Cloud resource group ID.
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

func (s *DeleteApisecEventsRequest) GetEventScope() *string {
	return s.EventScope
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

func (s *DeleteApisecEventsRequest) SetEventScope(v string) *DeleteApisecEventsRequest {
	s.EventScope = &v
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
