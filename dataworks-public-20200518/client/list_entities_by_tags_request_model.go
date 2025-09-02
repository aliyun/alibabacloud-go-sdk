// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEntitiesByTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntityType(v string) *ListEntitiesByTagsRequest
	GetEntityType() *string
	SetNextToken(v string) *ListEntitiesByTagsRequest
	GetNextToken() *string
	SetPageSize(v int32) *ListEntitiesByTagsRequest
	GetPageSize() *int32
	SetTags(v []*UserEntityTag) *ListEntitiesByTagsRequest
	GetTags() []*UserEntityTag
}

type ListEntitiesByTagsRequest struct {
	// The type of the entity.
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute-table
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 12345
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries per page. Default value: 10. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The tags.
	//
	// This parameter is required.
	Tags []*UserEntityTag `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListEntitiesByTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEntitiesByTagsRequest) GoString() string {
	return s.String()
}

func (s *ListEntitiesByTagsRequest) GetEntityType() *string {
	return s.EntityType
}

func (s *ListEntitiesByTagsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEntitiesByTagsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEntitiesByTagsRequest) GetTags() []*UserEntityTag {
	return s.Tags
}

func (s *ListEntitiesByTagsRequest) SetEntityType(v string) *ListEntitiesByTagsRequest {
	s.EntityType = &v
	return s
}

func (s *ListEntitiesByTagsRequest) SetNextToken(v string) *ListEntitiesByTagsRequest {
	s.NextToken = &v
	return s
}

func (s *ListEntitiesByTagsRequest) SetPageSize(v int32) *ListEntitiesByTagsRequest {
	s.PageSize = &v
	return s
}

func (s *ListEntitiesByTagsRequest) SetTags(v []*UserEntityTag) *ListEntitiesByTagsRequest {
	s.Tags = v
	return s
}

func (s *ListEntitiesByTagsRequest) Validate() error {
	return dara.Validate(s)
}
