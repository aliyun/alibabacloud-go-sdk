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
	SetLabelSelector(v []*string) *ListProjectsRequest
	GetLabelSelector() []*string
	SetPageNumber(v int64) *ListProjectsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListProjectsRequest
	GetPageSize() *int64
}

type ListProjectsRequest struct {
	// example:
	//
	// spring-boot
	Keyword       *string   `json:"keyword,omitempty" xml:"keyword,omitempty"`
	LabelSelector []*string `json:"labelSelector,omitempty" xml:"labelSelector,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
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

func (s *ListProjectsRequest) GetLabelSelector() []*string {
	return s.LabelSelector
}

func (s *ListProjectsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListProjectsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListProjectsRequest) SetKeyword(v string) *ListProjectsRequest {
	s.Keyword = &v
	return s
}

func (s *ListProjectsRequest) SetLabelSelector(v []*string) *ListProjectsRequest {
	s.LabelSelector = v
	return s
}

func (s *ListProjectsRequest) SetPageNumber(v int64) *ListProjectsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProjectsRequest) SetPageSize(v int64) *ListProjectsRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectsRequest) Validate() error {
	return dara.Validate(s)
}
