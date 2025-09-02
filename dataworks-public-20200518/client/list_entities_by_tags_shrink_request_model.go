// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEntitiesByTagsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntityType(v string) *ListEntitiesByTagsShrinkRequest
	GetEntityType() *string
	SetNextToken(v string) *ListEntitiesByTagsShrinkRequest
	GetNextToken() *string
	SetPageSize(v int32) *ListEntitiesByTagsShrinkRequest
	GetPageSize() *int32
	SetTagsShrink(v string) *ListEntitiesByTagsShrinkRequest
	GetTagsShrink() *string
}

type ListEntitiesByTagsShrinkRequest struct {
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
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListEntitiesByTagsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEntitiesByTagsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListEntitiesByTagsShrinkRequest) GetEntityType() *string {
	return s.EntityType
}

func (s *ListEntitiesByTagsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEntitiesByTagsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEntitiesByTagsShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *ListEntitiesByTagsShrinkRequest) SetEntityType(v string) *ListEntitiesByTagsShrinkRequest {
	s.EntityType = &v
	return s
}

func (s *ListEntitiesByTagsShrinkRequest) SetNextToken(v string) *ListEntitiesByTagsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListEntitiesByTagsShrinkRequest) SetPageSize(v int32) *ListEntitiesByTagsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListEntitiesByTagsShrinkRequest) SetTagsShrink(v string) *ListEntitiesByTagsShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListEntitiesByTagsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
