// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeBackupSummaryRequest
	GetClusterId() *string
	SetPageNumber(v int32) *DescribeBackupSummaryRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBackupSummaryRequest
	GetPageSize() *int32
}

type DescribeBackupSummaryRequest struct {
	// The ID of the instance. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/144595.html) operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld-bp169l540vc6c****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries per page. Maximum value: 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeBackupSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupSummaryRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeBackupSummaryRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBackupSummaryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBackupSummaryRequest) SetClusterId(v string) *DescribeBackupSummaryRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeBackupSummaryRequest) SetPageNumber(v int32) *DescribeBackupSummaryRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupSummaryRequest) SetPageSize(v int32) *DescribeBackupSummaryRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupSummaryRequest) Validate() error {
	return dara.Validate(s)
}
