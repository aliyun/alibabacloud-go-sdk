// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListToolsetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListToolsetsRequest
	GetKeyword() *string
	SetLabelSelector(v []*string) *ListToolsetsRequest
	GetLabelSelector() []*string
	SetPageNumber(v int64) *ListToolsetsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListToolsetsRequest
	GetPageSize() *int64
}

type ListToolsetsRequest struct {
	// example:
	//
	// demo
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

func (s ListToolsetsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListToolsetsRequest) GoString() string {
	return s.String()
}

func (s *ListToolsetsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListToolsetsRequest) GetLabelSelector() []*string {
	return s.LabelSelector
}

func (s *ListToolsetsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListToolsetsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListToolsetsRequest) SetKeyword(v string) *ListToolsetsRequest {
	s.Keyword = &v
	return s
}

func (s *ListToolsetsRequest) SetLabelSelector(v []*string) *ListToolsetsRequest {
	s.LabelSelector = v
	return s
}

func (s *ListToolsetsRequest) SetPageNumber(v int64) *ListToolsetsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListToolsetsRequest) SetPageSize(v int64) *ListToolsetsRequest {
	s.PageSize = &v
	return s
}

func (s *ListToolsetsRequest) Validate() error {
	return dara.Validate(s)
}
