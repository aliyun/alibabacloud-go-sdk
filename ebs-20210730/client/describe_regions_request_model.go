// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DescribeRegionsRequest
	GetAcceptLanguage() *string
	SetRegionId(v string) *DescribeRegionsRequest
	GetRegionId() *string
	SetResourceType(v string) *DescribeRegionsRequest
	GetResourceType() *string
}

type DescribeRegionsRequest struct {
	// The language in which the regions and zones are named. This parameter corresponds to the `LocalName` response parameter. Valid values:
	//
	// 	- zh-CN: Chinese
	//
	// 	- en-US: English
	//
	// 	- ja: Japanese
	//
	// Default value: zh-CN.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of resource. Valid values:
	//
	// 	- ear: async replication
	//
	// 	- lens: CloudLens for EBS
	//
	// 	- dbsc: Dedicated Block Storage Cluster
	//
	// Default value: ear.
	//
	// example:
	//
	// ear
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DescribeRegionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRegionsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceType(v string) *DescribeRegionsRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeRegionsRequest) Validate() error {
	return dara.Validate(s)
}
