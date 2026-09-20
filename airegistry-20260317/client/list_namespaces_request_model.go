// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNamespacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListNamespacesRequest
	GetName() *string
	SetPageNo(v int32) *ListNamespacesRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListNamespacesRequest
	GetPageSize() *int32
	SetTags(v string) *ListNamespacesRequest
	GetTags() *string
}

type ListNamespacesRequest struct {
	// Performs a fuzzy search by name.
	//
	// example:
	//
	// test-namespace
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Filters by tags. Separate multiple tags with commas. Results are matched by intersection.
	//
	// example:
	//
	// production,customer-service
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListNamespacesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNamespacesRequest) GoString() string {
	return s.String()
}

func (s *ListNamespacesRequest) GetName() *string {
	return s.Name
}

func (s *ListNamespacesRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListNamespacesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNamespacesRequest) GetTags() *string {
	return s.Tags
}

func (s *ListNamespacesRequest) SetName(v string) *ListNamespacesRequest {
	s.Name = &v
	return s
}

func (s *ListNamespacesRequest) SetPageNo(v int32) *ListNamespacesRequest {
	s.PageNo = &v
	return s
}

func (s *ListNamespacesRequest) SetPageSize(v int32) *ListNamespacesRequest {
	s.PageSize = &v
	return s
}

func (s *ListNamespacesRequest) SetTags(v string) *ListNamespacesRequest {
	s.Tags = &v
	return s
}

func (s *ListNamespacesRequest) Validate() error {
	return dara.Validate(s)
}
