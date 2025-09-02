// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLineageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *ListLineageRequest
	GetDirection() *string
	SetEntityQualifiedName(v string) *ListLineageRequest
	GetEntityQualifiedName() *string
	SetKeyword(v string) *ListLineageRequest
	GetKeyword() *string
	SetNextToken(v string) *ListLineageRequest
	GetNextToken() *string
	SetPageSize(v int32) *ListLineageRequest
	GetPageSize() *int32
}

type ListLineageRequest struct {
	// The lineage type. Valid values:
	//
	// 	- up: ancestor lineage
	//
	// 	- down: descendant lineage
	//
	// This parameter is required.
	//
	// example:
	//
	// up
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The unique identifier of the entity.
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute-table.project.table
	EntityQualifiedName *string `json:"EntityQualifiedName,omitempty" xml:"EntityQualifiedName,omitempty"`
	// The keyword of the entity name.
	//
	// example:
	//
	// name-keyword
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// next-token-from-previous-request
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListLineageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLineageRequest) GoString() string {
	return s.String()
}

func (s *ListLineageRequest) GetDirection() *string {
	return s.Direction
}

func (s *ListLineageRequest) GetEntityQualifiedName() *string {
	return s.EntityQualifiedName
}

func (s *ListLineageRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListLineageRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListLineageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLineageRequest) SetDirection(v string) *ListLineageRequest {
	s.Direction = &v
	return s
}

func (s *ListLineageRequest) SetEntityQualifiedName(v string) *ListLineageRequest {
	s.EntityQualifiedName = &v
	return s
}

func (s *ListLineageRequest) SetKeyword(v string) *ListLineageRequest {
	s.Keyword = &v
	return s
}

func (s *ListLineageRequest) SetNextToken(v string) *ListLineageRequest {
	s.NextToken = &v
	return s
}

func (s *ListLineageRequest) SetPageSize(v int32) *ListLineageRequest {
	s.PageSize = &v
	return s
}

func (s *ListLineageRequest) Validate() error {
	return dara.Validate(s)
}
