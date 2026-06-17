// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProjectMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabels(v string) *DescribeProjectMetaRequest
	GetLabels() *string
	SetPageNumber(v int32) *DescribeProjectMetaRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeProjectMetaRequest
	GetPageSize() *int32
}

type DescribeProjectMetaRequest struct {
	// The tags. Tags are used to filter alerts, and each alert can be marked with special tags.
	//
	// Currently, only filtering by product is supported. That is, the `name` is `product`. For example: {"name":"product","value":"ECS"}.
	//
	// >We do not recommend that you use the special tags for the CloudMonitor console in Alibaba Cloud.
	//
	// example:
	//
	// [{"name":"product","value":"ECS"}]
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The page number.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 10000.
	//
	// Default value: 30.
	//
	// >Currently, Alibaba Cloud does not impose a limit on this parameter. If you need to obtain all results, set the page size to a large value.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeProjectMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectMetaRequest) GoString() string {
	return s.String()
}

func (s *DescribeProjectMetaRequest) GetLabels() *string {
	return s.Labels
}

func (s *DescribeProjectMetaRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeProjectMetaRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeProjectMetaRequest) SetLabels(v string) *DescribeProjectMetaRequest {
	s.Labels = &v
	return s
}

func (s *DescribeProjectMetaRequest) SetPageNumber(v int32) *DescribeProjectMetaRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeProjectMetaRequest) SetPageSize(v int32) *DescribeProjectMetaRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeProjectMetaRequest) Validate() error {
	return dara.Validate(s)
}
