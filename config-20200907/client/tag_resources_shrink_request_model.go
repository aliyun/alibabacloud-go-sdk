// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagResourcesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *TagResourcesShrinkRequest
	GetRegionId() *string
	SetResourceId(v []*string) *TagResourcesShrinkRequest
	GetResourceId() []*string
	SetResourceType(v string) *TagResourcesShrinkRequest
	GetResourceType() *string
	SetTagShrink(v string) *TagResourcesShrinkRequest
	GetTagShrink() *string
}

type TagResourcesShrinkRequest struct {
	// The region ID of the tags.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource IDs. You can add tags to up to 50 resources.
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type. Valid values:
	//
	// 	- `ACS::Config::Rule`
	//
	// 	- `ACS::Config::AggregateConfigRule`
	//
	// 	- `ACS::Config::Aggregator`
	//
	// 	- `ACS::Config::CompliancePack`
	//
	// 	- `ACS::Config::AggregateCompliancePack`
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS::Config::Rule
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag value of the resource.
	//
	// You can specify up to 20 tag values.
	//
	// This parameter is required.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s TagResourcesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *TagResourcesShrinkRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *TagResourcesShrinkRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *TagResourcesShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *TagResourcesShrinkRequest) SetRegionId(v string) *TagResourcesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesShrinkRequest) SetResourceId(v []*string) *TagResourcesShrinkRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesShrinkRequest) SetResourceType(v string) *TagResourcesShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesShrinkRequest) SetTagShrink(v string) *TagResourcesShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *TagResourcesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
