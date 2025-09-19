// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSoarStrategyTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeSoarStrategyTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSoarStrategyTasksRequest
	GetPageSize() *int32
}

type DescribeSoarStrategyTasksRequest struct {
	// The page number. Pages start from page 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeSoarStrategyTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarStrategyTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeSoarStrategyTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSoarStrategyTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSoarStrategyTasksRequest) SetPageNumber(v int32) *DescribeSoarStrategyTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSoarStrategyTasksRequest) SetPageSize(v int32) *DescribeSoarStrategyTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSoarStrategyTasksRequest) Validate() error {
	return dara.Validate(s)
}
