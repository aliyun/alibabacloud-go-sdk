// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListParameterVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *ListParameterVersionsRequest
	GetId() *int64
	SetPageNumber(v int32) *ListParameterVersionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListParameterVersionsRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListParameterVersionsRequest
	GetSortBy() *string
}

type ListParameterVersionsRequest struct {
	// The parameter ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Specifies the page number of the results to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Specifies the number of entries to return per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies the sort criterion in the format \\<sort field> \\<sort order>. You can omit Asc for ascending order. The only valid value for the sort field is:
	//
	// - Version (Desc/Asc)
	//
	// example:
	//
	// Version Desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListParameterVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListParameterVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListParameterVersionsRequest) GetId() *int64 {
	return s.Id
}

func (s *ListParameterVersionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListParameterVersionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListParameterVersionsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListParameterVersionsRequest) SetId(v int64) *ListParameterVersionsRequest {
	s.Id = &v
	return s
}

func (s *ListParameterVersionsRequest) SetPageNumber(v int32) *ListParameterVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListParameterVersionsRequest) SetPageSize(v int32) *ListParameterVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListParameterVersionsRequest) SetSortBy(v string) *ListParameterVersionsRequest {
	s.SortBy = &v
	return s
}

func (s *ListParameterVersionsRequest) Validate() error {
	return dara.Validate(s)
}
