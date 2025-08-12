// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridMonitorSLSGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *DescribeHybridMonitorSLSGroupRequest
	GetKeyword() *string
	SetPageNumber(v string) *DescribeHybridMonitorSLSGroupRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeHybridMonitorSLSGroupRequest
	GetPageSize() *string
	SetRegionId(v string) *DescribeHybridMonitorSLSGroupRequest
	GetRegionId() *string
	SetSLSGroupName(v string) *DescribeHybridMonitorSLSGroupRequest
	GetSLSGroupName() *string
}

type DescribeHybridMonitorSLSGroupRequest struct {
	// The keyword that is used to search for Logstore groups.
	//
	// example:
	//
	// Logstore
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Minimum value: 1. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the Logstore group.
	//
	// example:
	//
	// Logstore_test
	SLSGroupName *string `json:"SLSGroupName,omitempty" xml:"SLSGroupName,omitempty"`
}

func (s DescribeHybridMonitorSLSGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridMonitorSLSGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeHybridMonitorSLSGroupRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeHybridMonitorSLSGroupRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeHybridMonitorSLSGroupRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeHybridMonitorSLSGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHybridMonitorSLSGroupRequest) GetSLSGroupName() *string {
	return s.SLSGroupName
}

func (s *DescribeHybridMonitorSLSGroupRequest) SetKeyword(v string) *DescribeHybridMonitorSLSGroupRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeHybridMonitorSLSGroupRequest) SetPageNumber(v string) *DescribeHybridMonitorSLSGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHybridMonitorSLSGroupRequest) SetPageSize(v string) *DescribeHybridMonitorSLSGroupRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHybridMonitorSLSGroupRequest) SetRegionId(v string) *DescribeHybridMonitorSLSGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHybridMonitorSLSGroupRequest) SetSLSGroupName(v string) *DescribeHybridMonitorSLSGroupRequest {
	s.SLSGroupName = &v
	return s
}

func (s *DescribeHybridMonitorSLSGroupRequest) Validate() error {
	return dara.Validate(s)
}
