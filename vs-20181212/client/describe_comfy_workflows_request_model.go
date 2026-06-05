// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComfyWorkflowsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DescribeComfyWorkflowsRequest
	GetName() *string
	SetPageNumber(v int32) *DescribeComfyWorkflowsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeComfyWorkflowsRequest
	GetPageSize() *int32
}

type DescribeComfyWorkflowsRequest struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeComfyWorkflowsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeComfyWorkflowsRequest) GoString() string {
	return s.String()
}

func (s *DescribeComfyWorkflowsRequest) GetName() *string {
	return s.Name
}

func (s *DescribeComfyWorkflowsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeComfyWorkflowsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeComfyWorkflowsRequest) SetName(v string) *DescribeComfyWorkflowsRequest {
	s.Name = &v
	return s
}

func (s *DescribeComfyWorkflowsRequest) SetPageNumber(v int32) *DescribeComfyWorkflowsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeComfyWorkflowsRequest) SetPageSize(v int32) *DescribeComfyWorkflowsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeComfyWorkflowsRequest) Validate() error {
	return dara.Validate(s)
}
