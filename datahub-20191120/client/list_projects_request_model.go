// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListProjectsRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListProjectsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListProjectsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListProjectsRequest
	GetRegionId() *string
}

type ListProjectsRequest struct {
	// Search keyword
	//
	// example:
	//
	// search_key
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// Current page number
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Region ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListProjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListProjectsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProjectsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProjectsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListProjectsRequest) SetKeyword(v string) *ListProjectsRequest {
	s.Keyword = &v
	return s
}

func (s *ListProjectsRequest) SetPageNumber(v int32) *ListProjectsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProjectsRequest) SetPageSize(v int32) *ListProjectsRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectsRequest) SetRegionId(v string) *ListProjectsRequest {
	s.RegionId = &v
	return s
}

func (s *ListProjectsRequest) Validate() error {
	return dara.Validate(s)
}
