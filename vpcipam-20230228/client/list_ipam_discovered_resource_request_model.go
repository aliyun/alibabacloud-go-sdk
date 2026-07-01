// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpamDiscoveredResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpamResourceDiscoveryId(v string) *ListIpamDiscoveredResourceRequest
	GetIpamResourceDiscoveryId() *string
	SetMaxResults(v int32) *ListIpamDiscoveredResourceRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListIpamDiscoveredResourceRequest
	GetNextToken() *string
	SetRegionId(v string) *ListIpamDiscoveredResourceRequest
	GetRegionId() *string
	SetResourceRegionId(v string) *ListIpamDiscoveredResourceRequest
	GetResourceRegionId() *string
	SetResourceType(v string) *ListIpamDiscoveredResourceRequest
	GetResourceType() *string
}

type ListIpamDiscoveredResourceRequest struct {
	// The resource discovery instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-res-disco-jt5f2af2u6nk2z321****
	IpamResourceDiscoveryId *string `json:"IpamResourceDiscoveryId,omitempty" xml:"IpamResourceDiscoveryId,omitempty"`
	// The maximum number of entries to return per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Valid values:
	//
	// - If this is the first query or no subsequent query exists, leave this parameter empty.
	//
	// - If a subsequent query exists, set this parameter to the **NextToken*	- value returned in the previous API call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The hosted region ID of the resource discovery instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to obtain the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The operating region of the resource discovery.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// The resource type. Valid values:
	//
	// - **VPC**: VPC.
	//
	// - **VSwitch**: vSwitch.
	//
	// example:
	//
	// VPC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListIpamDiscoveredResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIpamDiscoveredResourceRequest) GoString() string {
	return s.String()
}

func (s *ListIpamDiscoveredResourceRequest) GetIpamResourceDiscoveryId() *string {
	return s.IpamResourceDiscoveryId
}

func (s *ListIpamDiscoveredResourceRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListIpamDiscoveredResourceRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIpamDiscoveredResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListIpamDiscoveredResourceRequest) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *ListIpamDiscoveredResourceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListIpamDiscoveredResourceRequest) SetIpamResourceDiscoveryId(v string) *ListIpamDiscoveredResourceRequest {
	s.IpamResourceDiscoveryId = &v
	return s
}

func (s *ListIpamDiscoveredResourceRequest) SetMaxResults(v int32) *ListIpamDiscoveredResourceRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIpamDiscoveredResourceRequest) SetNextToken(v string) *ListIpamDiscoveredResourceRequest {
	s.NextToken = &v
	return s
}

func (s *ListIpamDiscoveredResourceRequest) SetRegionId(v string) *ListIpamDiscoveredResourceRequest {
	s.RegionId = &v
	return s
}

func (s *ListIpamDiscoveredResourceRequest) SetResourceRegionId(v string) *ListIpamDiscoveredResourceRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *ListIpamDiscoveredResourceRequest) SetResourceType(v string) *ListIpamDiscoveredResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *ListIpamDiscoveredResourceRequest) Validate() error {
	return dara.Validate(s)
}
