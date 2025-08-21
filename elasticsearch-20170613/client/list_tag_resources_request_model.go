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
	// The number of the returned page.
	//
	// example:
	//
	// 1d2db86sca4384811e0b5e8707e******
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Deprecated
	//
	// 1d2db86sca4384811e0b5e8707e\\*\\*\\*\\*\\*\\*
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ["es-cn-aaa","es-cn-bbb"]
	ResourceIds *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// [{"key":"env","value","dev"},{"key":"dev", "value":"IT"}]
	//
	// This parameter is required.
	//
	// example:
	//
	// INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Deprecated
	//
	// ["es-cn-aaa","es-cn-bbb"]
	//
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The header of the response. This parameter is empty and is for reference only. You cannot force this parameter to be relied on in the program.
	//
	// >  The return examples does not contain this parameter.
	//
	// example:
	//
	// [{"key":"env","value","dev"},{"key":"dev",  "value":"IT"}]
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
