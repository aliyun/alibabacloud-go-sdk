// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v string) *ListConfigsRequest
	GetFilter() *string
	SetPage(v int32) *ListConfigsRequest
	GetPage() *int32
	SetPageSize(v int32) *ListConfigsRequest
	GetPageSize() *int32
}

type ListConfigsRequest struct {
	// The field-level equality filter condition. The value is a URL-encoded JSON string (which decodes to a {"fieldName": value} object). Multiple fields have an AND relationship, meaning all conditions must be met for a result to be returned.
	//
	// example:
	//
	// %7B%22enabled%22%3Atrue%7D
	Filter *string `json:"filter,omitempty" xml:"filter,omitempty"`
	// page
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// pageSize
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListConfigsRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListConfigsRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListConfigsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListConfigsRequest) SetFilter(v string) *ListConfigsRequest {
	s.Filter = &v
	return s
}

func (s *ListConfigsRequest) SetPage(v int32) *ListConfigsRequest {
	s.Page = &v
	return s
}

func (s *ListConfigsRequest) SetPageSize(v int32) *ListConfigsRequest {
	s.PageSize = &v
	return s
}

func (s *ListConfigsRequest) Validate() error {
	return dara.Validate(s)
}
