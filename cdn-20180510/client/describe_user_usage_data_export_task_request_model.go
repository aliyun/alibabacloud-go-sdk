// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserUsageDataExportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v string) *DescribeUserUsageDataExportTaskRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeUserUsageDataExportTaskRequest
	GetPageSize() *string
}

type DescribeUserUsageDataExportTaskRequest struct {
	// The number of the page to return. Valid values: **1*	- to **100000**.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **20**. Maximum value: **50**.
	//
	// Valid values: an integer from **1*	- to **50**.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeUserUsageDataExportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserUsageDataExportTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserUsageDataExportTaskRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeUserUsageDataExportTaskRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeUserUsageDataExportTaskRequest) SetPageNumber(v string) *DescribeUserUsageDataExportTaskRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserUsageDataExportTaskRequest) SetPageSize(v string) *DescribeUserUsageDataExportTaskRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeUserUsageDataExportTaskRequest) Validate() error {
	return dara.Validate(s)
}
