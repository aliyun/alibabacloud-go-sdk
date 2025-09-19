// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebLockFileChangeStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeWebLockFileChangeStatisticsRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeWebLockFileChangeStatisticsRequest
	GetPageSize() *int32
}

type DescribeWebLockFileChangeStatisticsRequest struct {
	// The number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries to return on each page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeWebLockFileChangeStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockFileChangeStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebLockFileChangeStatisticsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeWebLockFileChangeStatisticsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeWebLockFileChangeStatisticsRequest) SetCurrentPage(v int32) *DescribeWebLockFileChangeStatisticsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWebLockFileChangeStatisticsRequest) SetPageSize(v int32) *DescribeWebLockFileChangeStatisticsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeWebLockFileChangeStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
