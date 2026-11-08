// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExporterOutputListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeExporterOutputListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeExporterOutputListRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeExporterOutputListRequest
	GetRegionId() *string
}

type DescribeExporterOutputListRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeExporterOutputListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExporterOutputListRequest) GoString() string {
	return s.String()
}

func (s *DescribeExporterOutputListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeExporterOutputListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeExporterOutputListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeExporterOutputListRequest) SetPageNumber(v int32) *DescribeExporterOutputListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeExporterOutputListRequest) SetPageSize(v int32) *DescribeExporterOutputListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeExporterOutputListRequest) SetRegionId(v string) *DescribeExporterOutputListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeExporterOutputListRequest) Validate() error {
	return dara.Validate(s)
}
