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
	// The IDs of resources. Separate multiple resource IDs with comma (,). This parameter is required if you do not specify the **Tags*	- parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["d42921c4-5433-4abd-8075-0e536f8b****"]
	ResourceIds *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// The type of the resource. Set the value to `application`.
	//
	// This parameter is required.
	//
	// example:
	//
	// application
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag in the format of a key-value pair. This parameter is required if you do not specify the **ResourceIds*	- parameter. The following parameters are involved:
	//
	// 	- **key**: the tag key. It cannot exceed 128 characters in length.
	//
	// 	- **value**: the tag value. It cannot exceed 128 characters in length.
	//
	// Tag keys and tag values are case-sensitive. If you specify multiple tags, the system adds all the tags to the specified resources. Each tag key on a resource can have only one tag value. If you create a tag that has the same key as an existing tag, the value of the existing tag is overwritten.
	//
	// Tag keys and tag values cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
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
