// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvironmentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListEnvironmentsRequest
	GetKeyword() *string
	SetLabelSelector(v []*string) *ListEnvironmentsRequest
	GetLabelSelector() []*string
	SetPageNumber(v int64) *ListEnvironmentsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListEnvironmentsRequest
	GetPageSize() *int64
}

type ListEnvironmentsRequest struct {
	// example:
	//
	// dev
	Keyword       *string   `json:"keyword,omitempty" xml:"keyword,omitempty"`
	LabelSelector []*string `json:"labelSelector,omitempty" xml:"labelSelector,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListEnvironmentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentsRequest) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListEnvironmentsRequest) GetLabelSelector() []*string {
	return s.LabelSelector
}

func (s *ListEnvironmentsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListEnvironmentsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListEnvironmentsRequest) SetKeyword(v string) *ListEnvironmentsRequest {
	s.Keyword = &v
	return s
}

func (s *ListEnvironmentsRequest) SetLabelSelector(v []*string) *ListEnvironmentsRequest {
	s.LabelSelector = v
	return s
}

func (s *ListEnvironmentsRequest) SetPageNumber(v int64) *ListEnvironmentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEnvironmentsRequest) SetPageSize(v int64) *ListEnvironmentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListEnvironmentsRequest) Validate() error {
	return dara.Validate(s)
}
