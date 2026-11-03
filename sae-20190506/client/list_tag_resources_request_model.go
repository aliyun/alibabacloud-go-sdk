// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListTagResourcesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListTagResourcesRequest
	GetRegionId() *string
	SetResourceIds(v string) *ListTagResourcesRequest
	GetResourceIds() *string
	SetResourceType(v string) *ListTagResourcesRequest
	GetResourceType() *string
	SetTags(v string) *ListTagResourcesRequest
	GetTags() *string
}

type ListTagResourcesRequest struct {
	// The pagination token. A single query returns up to 50 results. If the total number of results exceeds 50, a NextToken is returned. You can use this token in the next query to retrieve the remaining results.
	//
	// example:
	//
	// A2RN
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource IDs. You can specify multiple resource IDs separated by commas (,). This parameter is required if you do not specify the **Tags*	- parameter.
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
	// The tag key-value pairs. This parameter is required if you do not specify the **ResourceIds*	- parameter.
	//
	// - **key**: the tag key. The value must be 1 to 128 characters in length.
	//
	// - **value**: the tag value. The value must be 1 to 128 characters in length.
	//
	// Tag keys and values are case-sensitive. If you specify multiple tags, multiple tags are created and bound to the specified resource at the same time. A tag key on a resource can map to only one tag value. If you add a tag with an existing key, the tag value is updated to the new value.
	//
	// The key and value cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// [{"key":"k1","value":"v1"}]
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListTagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTagResourcesRequest) GetResourceIds() *string {
	return s.ResourceIds
}

func (s *ListTagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagResourcesRequest) GetTags() *string {
	return s.Tags
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceIds(v string) *ListTagResourcesRequest {
	s.ResourceIds = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTags(v string) *ListTagResourcesRequest {
	s.Tags = &v
	return s
}

func (s *ListTagResourcesRequest) Validate() error {
	return dara.Validate(s)
}
