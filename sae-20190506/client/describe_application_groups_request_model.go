// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeApplicationGroupsRequest
	GetAppId() *string
	SetCurrentPage(v int32) *DescribeApplicationGroupsRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeApplicationGroupsRequest
	GetPageSize() *int32
}

type DescribeApplicationGroupsRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// d700e680-aa4d-4ec1-afc2-6566b5ff****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeApplicationGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationGroupsRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeApplicationGroupsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeApplicationGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApplicationGroupsRequest) SetAppId(v string) *DescribeApplicationGroupsRequest {
	s.AppId = &v
	return s
}

func (s *DescribeApplicationGroupsRequest) SetCurrentPage(v int32) *DescribeApplicationGroupsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeApplicationGroupsRequest) SetPageSize(v int32) *DescribeApplicationGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApplicationGroupsRequest) Validate() error {
	return dara.Validate(s)
}
