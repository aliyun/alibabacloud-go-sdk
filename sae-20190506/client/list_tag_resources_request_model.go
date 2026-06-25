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
	// A query can return a maximum of 50 results. If the number of results exceeds this limit, the response includes a NextToken. To retrieve the next page of results, pass this token in your next request.
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
	// The resource IDs, specified as a JSON array of strings. This parameter is required if the **Tags*	- parameter is not specified.
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
	// The tags used to filter resources. This parameter is required if the **ResourceIds*	- parameter is not specified. A tag is a key-value pair.
	//
	// - **key**: The tag key. The key can be 1 to 128 characters in length.
	//
	// - **value**: The tag value. The value can be 1 to 128 characters in length.
	//
	// Tag keys and tag values are case-sensitive. If you specify multiple tags, the operation returns only resources that have all the specified tags.
	//
	// A tag key cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
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
