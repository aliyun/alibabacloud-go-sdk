// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSoarStrategiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeSoarStrategiesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSoarStrategiesRequest
	GetPageSize() *int32
}

type DescribeSoarStrategiesRequest struct {
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

func (s DescribeSoarStrategiesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarStrategiesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSoarStrategiesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSoarStrategiesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSoarStrategiesRequest) SetPageNumber(v int32) *DescribeSoarStrategiesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSoarStrategiesRequest) SetPageSize(v int32) *DescribeSoarStrategiesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSoarStrategiesRequest) Validate() error {
	return dara.Validate(s)
}
