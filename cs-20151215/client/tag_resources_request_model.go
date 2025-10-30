// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *TagResourcesRequest
	GetRegionId() *string
	SetResourceIds(v []*string) *TagResourcesRequest
	GetResourceIds() []*string
	SetResourceType(v string) *TagResourcesRequest
	GetResourceType() *string
	SetTags(v []*Tag) *TagResourcesRequest
	GetTags() []*Tag
}

type TagResourcesRequest struct {
	// The ID of the region in which the resource resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"region_id,omitempty" xml:"region_id,omitempty"`
	// The list of resource IDs.
	//
	// This parameter is required.
	ResourceIds []*string `json:"resource_ids,omitempty" xml:"resource_ids,omitempty" type:"Repeated"`
	// The type of resources that you want to label. Set the value to `CLUSTER`.
	//
	// This parameter is required.
	//
	// example:
	//
	// CLUSTER
	ResourceType *string `json:"resource_type,omitempty" xml:"resource_type,omitempty"`
	// The tags that you want to add to the resources in key-value pairs. You can add up to 20 key-value pairs. Note:
	//
	// 	- The values cannot be empty strings. A value must be 1 to 128 characters in length.
	//
	// 	- A key or value cannot start with `aliyun` or `acs:`.
	//
	// 	- A key or value cannot contain `http://` or `https://`.
	//
	// This parameter is required.
	Tags []*Tag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *TagResourcesRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *TagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *TagResourcesRequest) GetTags() []*Tag {
	return s.Tags
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceIds(v []*string) *TagResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTags(v []*Tag) *TagResourcesRequest {
	s.Tags = v
	return s
}

func (s *TagResourcesRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
