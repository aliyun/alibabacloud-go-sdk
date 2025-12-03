// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestoreSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeRestoreSummaryRequest
	GetClusterId() *string
	SetPageNumber(v int32) *DescribeRestoreSummaryRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRestoreSummaryRequest
	GetPageSize() *int32
}

type DescribeRestoreSummaryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-m5eznlga4k5bcxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeRestoreSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeRestoreSummaryRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeRestoreSummaryRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRestoreSummaryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRestoreSummaryRequest) SetClusterId(v string) *DescribeRestoreSummaryRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeRestoreSummaryRequest) SetPageNumber(v int32) *DescribeRestoreSummaryRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRestoreSummaryRequest) SetPageSize(v int32) *DescribeRestoreSummaryRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRestoreSummaryRequest) Validate() error {
	return dara.Validate(s)
}
