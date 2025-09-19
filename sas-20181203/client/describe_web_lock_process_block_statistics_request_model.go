// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebLockProcessBlockStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeWebLockProcessBlockStatisticsRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeWebLockProcessBlockStatisticsRequest
	GetPageSize() *int32
}

type DescribeWebLockProcessBlockStatisticsRequest struct {
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries to return on each page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeWebLockProcessBlockStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockProcessBlockStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebLockProcessBlockStatisticsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeWebLockProcessBlockStatisticsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeWebLockProcessBlockStatisticsRequest) SetCurrentPage(v int32) *DescribeWebLockProcessBlockStatisticsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWebLockProcessBlockStatisticsRequest) SetPageSize(v int32) *DescribeWebLockProcessBlockStatisticsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeWebLockProcessBlockStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
