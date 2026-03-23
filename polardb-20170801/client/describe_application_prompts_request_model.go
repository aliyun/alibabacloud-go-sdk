// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationPromptsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribeApplicationPromptsRequest
	GetApplicationId() *string
	SetPageNumber(v int32) *DescribeApplicationPromptsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApplicationPromptsRequest
	GetPageSize() *int32
}

type DescribeApplicationPromptsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeApplicationPromptsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationPromptsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationPromptsRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribeApplicationPromptsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApplicationPromptsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApplicationPromptsRequest) SetApplicationId(v string) *DescribeApplicationPromptsRequest {
	s.ApplicationId = &v
	return s
}

func (s *DescribeApplicationPromptsRequest) SetPageNumber(v int32) *DescribeApplicationPromptsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApplicationPromptsRequest) SetPageSize(v int32) *DescribeApplicationPromptsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApplicationPromptsRequest) Validate() error {
	return dara.Validate(s)
}
