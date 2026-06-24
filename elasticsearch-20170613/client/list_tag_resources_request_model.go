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
	SetPage(v int32) *ListTagResourcesRequest
	GetPage() *int32
	SetResourceIds(v string) *ListTagResourcesRequest
	GetResourceIds() *string
	SetResourceType(v string) *ListTagResourcesRequest
	GetResourceType() *string
	SetSize(v int32) *ListTagResourcesRequest
	GetSize() *int32
	SetTags(v string) *ListTagResourcesRequest
	GetTags() *string
}

type ListTagResourcesRequest struct {
	// The token for the next query.
	//
	// example:
	//
	// 1d2db86sca4384811e0b5e8707e******
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Deprecated
	//
	// The page number of the resource relationship list. This parameter is deprecated.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The list of instance IDs to query. The value is in JSON array format and can contain up to 20 items.
	//
	// example:
	//
	// ["es-cn-aaa","es-cn-bbb"]
	ResourceIds *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// The resource type definition.
	//
	// This parameter is required.
	//
	// example:
	//
	// INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Deprecated
	//
	// The number of entries per page in Settings for paged query and paging. This field is deprecated.
	//
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The list of tags to query. The value is in JSON string format and can contain up to 20 items.
	//
	// example:
	//
	// [{"key":"env", "value":"dev"},{"key":"dev", "value":"IT"}]
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

func (s *ListTagResourcesRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListTagResourcesRequest) GetResourceIds() *string {
	return s.ResourceIds
}

func (s *ListTagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagResourcesRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListTagResourcesRequest) GetTags() *string {
	return s.Tags
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetPage(v int32) *ListTagResourcesRequest {
	s.Page = &v
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

func (s *ListTagResourcesRequest) SetSize(v int32) *ListTagResourcesRequest {
	s.Size = &v
	return s
}

func (s *ListTagResourcesRequest) SetTags(v string) *ListTagResourcesRequest {
	s.Tags = &v
	return s
}

func (s *ListTagResourcesRequest) Validate() error {
	return dara.Validate(s)
}
