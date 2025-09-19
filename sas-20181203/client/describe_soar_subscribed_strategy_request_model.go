// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSoarSubscribedStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeSoarSubscribedStrategyRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSoarSubscribedStrategyRequest
	GetPageSize() *int32
}

type DescribeSoarSubscribedStrategyRequest struct {
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

func (s DescribeSoarSubscribedStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarSubscribedStrategyRequest) GoString() string {
	return s.String()
}

func (s *DescribeSoarSubscribedStrategyRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSoarSubscribedStrategyRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSoarSubscribedStrategyRequest) SetPageNumber(v int32) *DescribeSoarSubscribedStrategyRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSoarSubscribedStrategyRequest) SetPageSize(v int32) *DescribeSoarSubscribedStrategyRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSoarSubscribedStrategyRequest) Validate() error {
	return dara.Validate(s)
}
