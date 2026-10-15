// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppTemplateDictsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDictType(v string) *ListAppTemplateDictsRequest
	GetDictType() *string
	SetMaxResults(v int32) *ListAppTemplateDictsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAppTemplateDictsRequest
	GetNextToken() *string
}

type ListAppTemplateDictsRequest struct {
	// The dictionary type.
	//
	// example:
	//
	// product_version
	DictType *string `json:"DictType,omitempty" xml:"DictType,omitempty"`
	// The number of entries per query.
	//
	// Valid values: 10 to 100. Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. This parameter is empty if no more results exist.
	//
	// example:
	//
	// FFh3Xqm+JgZ/U9Jyb7wdVr9LWk80Tghn5UZjbcWEVEderBcbVF+Y6PS0i8PpCL4PQZ3e0C9oEH0Asd4tJEuGtkl2WuKdiWZpEwadNydQdJPFM=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListAppTemplateDictsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppTemplateDictsRequest) GoString() string {
	return s.String()
}

func (s *ListAppTemplateDictsRequest) GetDictType() *string {
	return s.DictType
}

func (s *ListAppTemplateDictsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAppTemplateDictsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAppTemplateDictsRequest) SetDictType(v string) *ListAppTemplateDictsRequest {
	s.DictType = &v
	return s
}

func (s *ListAppTemplateDictsRequest) SetMaxResults(v int32) *ListAppTemplateDictsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAppTemplateDictsRequest) SetNextToken(v string) *ListAppTemplateDictsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAppTemplateDictsRequest) Validate() error {
	return dara.Validate(s)
}
