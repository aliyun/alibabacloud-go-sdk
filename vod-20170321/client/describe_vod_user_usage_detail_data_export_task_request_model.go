// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodUserUsageDetailDataExportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeVodUserUsageDetailDataExportTaskRequest
	GetOwnerId() *int64
	SetPageNumber(v string) *DescribeVodUserUsageDetailDataExportTaskRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeVodUserUsageDetailDataExportTaskRequest
	GetPageSize() *string
}

type DescribeVodUserUsageDetailDataExportTaskRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeVodUserUsageDetailDataExportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserUsageDetailDataExportTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodUserUsageDetailDataExportTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodUserUsageDetailDataExportTaskRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeVodUserUsageDetailDataExportTaskRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeVodUserUsageDetailDataExportTaskRequest) SetOwnerId(v int64) *DescribeVodUserUsageDetailDataExportTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodUserUsageDetailDataExportTaskRequest) SetPageNumber(v string) *DescribeVodUserUsageDetailDataExportTaskRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVodUserUsageDetailDataExportTaskRequest) SetPageSize(v string) *DescribeVodUserUsageDetailDataExportTaskRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVodUserUsageDetailDataExportTaskRequest) Validate() error {
	return dara.Validate(s)
}
