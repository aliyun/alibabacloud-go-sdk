// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExporterRuleListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeExporterRuleListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeExporterRuleListRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeExporterRuleListRequest
	GetRegionId() *string
}

type DescribeExporterRuleListRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 1000.
	//
	// example:
	//
	// 1000
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeExporterRuleListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExporterRuleListRequest) GoString() string {
	return s.String()
}

func (s *DescribeExporterRuleListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeExporterRuleListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeExporterRuleListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeExporterRuleListRequest) SetPageNumber(v int32) *DescribeExporterRuleListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeExporterRuleListRequest) SetPageSize(v int32) *DescribeExporterRuleListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeExporterRuleListRequest) SetRegionId(v string) *DescribeExporterRuleListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeExporterRuleListRequest) Validate() error {
	return dara.Validate(s)
}
