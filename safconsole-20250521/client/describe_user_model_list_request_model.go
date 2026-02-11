// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserModelListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthType(v string) *DescribeUserModelListRequest
	GetAuthType() *string
	SetCurrentPage(v string) *DescribeUserModelListRequest
	GetCurrentPage() *string
	SetName(v string) *DescribeUserModelListRequest
	GetName() *string
	SetPageSize(v string) *DescribeUserModelListRequest
	GetPageSize() *string
}

type DescribeUserModelListRequest struct {
	// Authorization type.
	//
	// example:
	//
	// READ
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Model name.
	//
	// example:
	//
	// StudyX
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Pagination parameter: number of items per page, default value 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeUserModelListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserModelListRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserModelListRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *DescribeUserModelListRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeUserModelListRequest) GetName() *string {
	return s.Name
}

func (s *DescribeUserModelListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeUserModelListRequest) SetAuthType(v string) *DescribeUserModelListRequest {
	s.AuthType = &v
	return s
}

func (s *DescribeUserModelListRequest) SetCurrentPage(v string) *DescribeUserModelListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeUserModelListRequest) SetName(v string) *DescribeUserModelListRequest {
	s.Name = &v
	return s
}

func (s *DescribeUserModelListRequest) SetPageSize(v string) *DescribeUserModelListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeUserModelListRequest) Validate() error {
	return dara.Validate(s)
}
