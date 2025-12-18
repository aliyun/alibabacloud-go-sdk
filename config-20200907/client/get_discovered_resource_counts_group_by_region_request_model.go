// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDiscoveredResourceCountsGroupByRegionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceType(v string) *GetDiscoveredResourceCountsGroupByRegionRequest
	GetResourceType() *string
}

type GetDiscoveredResourceCountsGroupByRegionRequest struct {
	// The resource type.
	//
	// For more information about how to obtain the type of a resource, see [ListDiscoveredResources](https://help.aliyun.com/document_detail/169620.html).
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetDiscoveredResourceCountsGroupByRegionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDiscoveredResourceCountsGroupByRegionRequest) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceCountsGroupByRegionRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetDiscoveredResourceCountsGroupByRegionRequest) SetResourceType(v string) *GetDiscoveredResourceCountsGroupByRegionRequest {
	s.ResourceType = &v
	return s
}

func (s *GetDiscoveredResourceCountsGroupByRegionRequest) Validate() error {
	return dara.Validate(s)
}
