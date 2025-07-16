// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserUsageDetailDataExportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v string) *DescribeUserUsageDetailDataExportTaskRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeUserUsageDetailDataExportTaskRequest
	GetPageSize() *string
}

type DescribeUserUsageDetailDataExportTaskRequest struct {
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
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeUserUsageDetailDataExportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserUsageDetailDataExportTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserUsageDetailDataExportTaskRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeUserUsageDetailDataExportTaskRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeUserUsageDetailDataExportTaskRequest) SetPageNumber(v string) *DescribeUserUsageDetailDataExportTaskRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskRequest) SetPageSize(v string) *DescribeUserUsageDetailDataExportTaskRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeUserUsageDetailDataExportTaskRequest) Validate() error {
	return dara.Validate(s)
}
