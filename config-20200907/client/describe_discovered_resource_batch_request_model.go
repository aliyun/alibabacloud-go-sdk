// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiscoveredResourceBatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegions(v string) *DescribeDiscoveredResourceBatchRequest
	GetRegions() *string
	SetResourceIds(v string) *DescribeDiscoveredResourceBatchRequest
	GetResourceIds() *string
	SetResourceTypes(v string) *DescribeDiscoveredResourceBatchRequest
	GetResourceTypes() *string
}

type DescribeDiscoveredResourceBatchRequest struct {
	// example:
	//
	// cn-shanghai,cn-hongkong,cn-zhangjiakou,cn-hangzhou
	Regions *string `json:"Regions,omitempty" xml:"Regions,omitempty"`
	// example:
	//
	// r-wz998f311e21exxx,r-wz97f4a734789xxx
	ResourceIds *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// example:
	//
	// ACS::ECS::Disk
	ResourceTypes *string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty"`
}

func (s DescribeDiscoveredResourceBatchRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiscoveredResourceBatchRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiscoveredResourceBatchRequest) GetRegions() *string {
	return s.Regions
}

func (s *DescribeDiscoveredResourceBatchRequest) GetResourceIds() *string {
	return s.ResourceIds
}

func (s *DescribeDiscoveredResourceBatchRequest) GetResourceTypes() *string {
	return s.ResourceTypes
}

func (s *DescribeDiscoveredResourceBatchRequest) SetRegions(v string) *DescribeDiscoveredResourceBatchRequest {
	s.Regions = &v
	return s
}

func (s *DescribeDiscoveredResourceBatchRequest) SetResourceIds(v string) *DescribeDiscoveredResourceBatchRequest {
	s.ResourceIds = &v
	return s
}

func (s *DescribeDiscoveredResourceBatchRequest) SetResourceTypes(v string) *DescribeDiscoveredResourceBatchRequest {
	s.ResourceTypes = &v
	return s
}

func (s *DescribeDiscoveredResourceBatchRequest) Validate() error {
	return dara.Validate(s)
}
