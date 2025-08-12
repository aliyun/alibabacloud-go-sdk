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
	// The tags. Tags are used to filter services.
	//
	// You can filter services only by the tag whose `name` is `product`. Example: {"name":"product","value":"ECS"}.
	//
	// > We recommend that you do not use the special tags in the CloudMonitor console.
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
	// > The value of this parameter is not limited. You can view a large number of entries per page.
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
