// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConnectionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListConnectionsRequest
	GetKeyword() *string
	SetLabelSelector(v []*string) *ListConnectionsRequest
	GetLabelSelector() []*string
	SetPageNumber(v int64) *ListConnectionsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListConnectionsRequest
	GetPageSize() *int64
}

type ListConnectionsRequest struct {
	// example:
	//
	// auto-
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

func (s ListConnectionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConnectionsRequest) GoString() string {
	return s.String()
}

func (s *ListConnectionsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListConnectionsRequest) GetLabelSelector() []*string {
	return s.LabelSelector
}

func (s *ListConnectionsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListConnectionsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListConnectionsRequest) SetKeyword(v string) *ListConnectionsRequest {
	s.Keyword = &v
	return s
}

func (s *ListConnectionsRequest) SetLabelSelector(v []*string) *ListConnectionsRequest {
	s.LabelSelector = v
	return s
}

func (s *ListConnectionsRequest) SetPageNumber(v int64) *ListConnectionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListConnectionsRequest) SetPageSize(v int64) *ListConnectionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListConnectionsRequest) Validate() error {
	return dara.Validate(s)
}
