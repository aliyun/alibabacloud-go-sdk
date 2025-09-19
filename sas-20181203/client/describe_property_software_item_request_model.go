// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertySoftwareItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribePropertySoftwareItemRequest
	GetCurrentPage() *int32
	SetForceFlush(v bool) *DescribePropertySoftwareItemRequest
	GetForceFlush() *bool
	SetName(v string) *DescribePropertySoftwareItemRequest
	GetName() *string
	SetPageSize(v int32) *DescribePropertySoftwareItemRequest
	GetPageSize() *int32
}

type DescribePropertySoftwareItemRequest struct {
	// The number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Specifies whether to forcefully refresh the data that you want to query.
	//
	// example:
	//
	// true
	ForceFlush *bool `json:"ForceFlush,omitempty" xml:"ForceFlush,omitempty"`
	// The name of the software that you want to query.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribePropertySoftwareItemRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertySoftwareItemRequest) GoString() string {
	return s.String()
}

func (s *DescribePropertySoftwareItemRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePropertySoftwareItemRequest) GetForceFlush() *bool {
	return s.ForceFlush
}

func (s *DescribePropertySoftwareItemRequest) GetName() *string {
	return s.Name
}

func (s *DescribePropertySoftwareItemRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePropertySoftwareItemRequest) SetCurrentPage(v int32) *DescribePropertySoftwareItemRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertySoftwareItemRequest) SetForceFlush(v bool) *DescribePropertySoftwareItemRequest {
	s.ForceFlush = &v
	return s
}

func (s *DescribePropertySoftwareItemRequest) SetName(v string) *DescribePropertySoftwareItemRequest {
	s.Name = &v
	return s
}

func (s *DescribePropertySoftwareItemRequest) SetPageSize(v int32) *DescribePropertySoftwareItemRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePropertySoftwareItemRequest) Validate() error {
	return dara.Validate(s)
}
