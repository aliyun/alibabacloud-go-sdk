// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyPortItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribePropertyPortItemRequest
	GetCurrentPage() *int32
	SetForceFlush(v bool) *DescribePropertyPortItemRequest
	GetForceFlush() *bool
	SetPageSize(v int32) *DescribePropertyPortItemRequest
	GetPageSize() *int32
	SetPort(v string) *DescribePropertyPortItemRequest
	GetPort() *string
}

type DescribePropertyPortItemRequest struct {
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
	// The number of entries to return on each page.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The port number.
	//
	// example:
	//
	// 22
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribePropertyPortItemRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyPortItemRequest) GoString() string {
	return s.String()
}

func (s *DescribePropertyPortItemRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePropertyPortItemRequest) GetForceFlush() *bool {
	return s.ForceFlush
}

func (s *DescribePropertyPortItemRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePropertyPortItemRequest) GetPort() *string {
	return s.Port
}

func (s *DescribePropertyPortItemRequest) SetCurrentPage(v int32) *DescribePropertyPortItemRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyPortItemRequest) SetForceFlush(v bool) *DescribePropertyPortItemRequest {
	s.ForceFlush = &v
	return s
}

func (s *DescribePropertyPortItemRequest) SetPageSize(v int32) *DescribePropertyPortItemRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePropertyPortItemRequest) SetPort(v string) *DescribePropertyPortItemRequest {
	s.Port = &v
	return s
}

func (s *DescribePropertyPortItemRequest) Validate() error {
	return dara.Validate(s)
}
