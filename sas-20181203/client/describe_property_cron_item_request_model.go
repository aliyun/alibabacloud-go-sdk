// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyCronItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribePropertyCronItemRequest
	GetCurrentPage() *int32
	SetForceFlush(v bool) *DescribePropertyCronItemRequest
	GetForceFlush() *bool
	SetPageSize(v int32) *DescribePropertyCronItemRequest
	GetPageSize() *int32
	SetSource(v string) *DescribePropertyCronItemRequest
	GetSource() *string
}

type DescribePropertyCronItemRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Specifies whether to forcefully refresh the data that you want to query. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// false
	ForceFlush *bool `json:"ForceFlush,omitempty" xml:"ForceFlush,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The path to the scheduled task.
	//
	// example:
	//
	// /data
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DescribePropertyCronItemRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyCronItemRequest) GoString() string {
	return s.String()
}

func (s *DescribePropertyCronItemRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePropertyCronItemRequest) GetForceFlush() *bool {
	return s.ForceFlush
}

func (s *DescribePropertyCronItemRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePropertyCronItemRequest) GetSource() *string {
	return s.Source
}

func (s *DescribePropertyCronItemRequest) SetCurrentPage(v int32) *DescribePropertyCronItemRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyCronItemRequest) SetForceFlush(v bool) *DescribePropertyCronItemRequest {
	s.ForceFlush = &v
	return s
}

func (s *DescribePropertyCronItemRequest) SetPageSize(v int32) *DescribePropertyCronItemRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePropertyCronItemRequest) SetSource(v string) *DescribePropertyCronItemRequest {
	s.Source = &v
	return s
}

func (s *DescribePropertyCronItemRequest) Validate() error {
	return dara.Validate(s)
}
