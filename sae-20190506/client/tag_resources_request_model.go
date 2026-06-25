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
	SetResourceIds(v string) *TagResourcesRequest
	GetResourceIds() *string
	SetResourceType(v string) *TagResourcesRequest
	GetResourceType() *string
	SetTags(v string) *TagResourcesRequest
	GetTags() *string
}

type TagResourcesRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource IDs. You can specify up to 50 resource IDs in a JSON array. This parameter is required unless you specify the **Tags*	- parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["d42921c4-5433-4abd-8075-0e536f8b****"]
	ResourceIds *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// The resource type. Only `application` is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// application
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key-value pairs of the tags. This parameter is required unless you specify the **ResourceIds*	- parameter. The following rules apply:
	//
	// - **key**: The tag key. The key must be 1 to 128 characters in length.
	//
	// - **value**: The tag value. The value must be 1 to 128 characters in length.
	//
	// Tags are case-sensitive. If you specify multiple tags, they are created and bound to the specified resources. For a single resource, each tag key must be unique. If you specify a tag key that already exists for a resource, the operation updates the existing tag value.
	//
	// A tag key cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// [{"key":"k1","value":"v1"}]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
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

func (s *TagResourcesRequest) GetResourceIds() *string {
	return s.ResourceIds
}

func (s *TagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *TagResourcesRequest) GetTags() *string {
	return s.Tags
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceIds(v string) *TagResourcesRequest {
	s.ResourceIds = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTags(v string) *TagResourcesRequest {
	s.Tags = &v
	return s
}

func (s *TagResourcesRequest) Validate() error {
	return dara.Validate(s)
}
