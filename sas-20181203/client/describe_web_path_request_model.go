// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebPathRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeWebPathRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeWebPathRequest
	GetPageSize() *int32
	SetType(v string) *DescribeWebPathRequest
	GetType() *string
}

type DescribeWebPathRequest struct {
	// The number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the alert configuration. Valid values:
	//
	// 	- **web_path**
	//
	// example:
	//
	// web_path
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeWebPathRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebPathRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebPathRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeWebPathRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeWebPathRequest) GetType() *string {
	return s.Type
}

func (s *DescribeWebPathRequest) SetCurrentPage(v int32) *DescribeWebPathRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWebPathRequest) SetPageSize(v int32) *DescribeWebPathRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeWebPathRequest) SetType(v string) *DescribeWebPathRequest {
	s.Type = &v
	return s
}

func (s *DescribeWebPathRequest) Validate() error {
	return dara.Validate(s)
}
