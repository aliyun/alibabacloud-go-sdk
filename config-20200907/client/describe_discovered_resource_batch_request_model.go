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
	// The regions where the resources reside. Separate multiple regions with commas (`,`).
	//
	// example:
	//
	// cn-shanghai,cn-hongkong,cn-zhangjiakou,cn-hangzhou
	Regions *string `json:"Regions,omitempty" xml:"Regions,omitempty"`
	// The resource IDs. Separate multiple resource IDs with commas (`,`).
	//
	// example:
	//
	// r-wz998f311e21****,r-wz97f4a73478****
	ResourceIds *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// The resource types. Separate multiple resource types with commas (`,`).
	//
	// example:
	//
	// ACS::ECS::Disk,ACS::ECS::Instance
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
