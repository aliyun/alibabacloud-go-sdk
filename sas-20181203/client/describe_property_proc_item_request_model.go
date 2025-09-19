// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyProcItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribePropertyProcItemRequest
	GetCurrentPage() *int32
	SetForceFlush(v bool) *DescribePropertyProcItemRequest
	GetForceFlush() *bool
	SetName(v string) *DescribePropertyProcItemRequest
	GetName() *string
	SetPageSize(v int32) *DescribePropertyProcItemRequest
	GetPageSize() *int32
}

type DescribePropertyProcItemRequest struct {
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
	// The name of the process.
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

func (s DescribePropertyProcItemRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyProcItemRequest) GoString() string {
	return s.String()
}

func (s *DescribePropertyProcItemRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePropertyProcItemRequest) GetForceFlush() *bool {
	return s.ForceFlush
}

func (s *DescribePropertyProcItemRequest) GetName() *string {
	return s.Name
}

func (s *DescribePropertyProcItemRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePropertyProcItemRequest) SetCurrentPage(v int32) *DescribePropertyProcItemRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyProcItemRequest) SetForceFlush(v bool) *DescribePropertyProcItemRequest {
	s.ForceFlush = &v
	return s
}

func (s *DescribePropertyProcItemRequest) SetName(v string) *DescribePropertyProcItemRequest {
	s.Name = &v
	return s
}

func (s *DescribePropertyProcItemRequest) SetPageSize(v int32) *DescribePropertyProcItemRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePropertyProcItemRequest) Validate() error {
	return dara.Validate(s)
}
