// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBComputeClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComputeClusterId(v string) *DescribeAgenticDBComputeClustersRequest
	GetComputeClusterId() *string
	SetDBClusterId(v string) *DescribeAgenticDBComputeClustersRequest
	GetDBClusterId() *string
	SetPageNumber(v int32) *DescribeAgenticDBComputeClustersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAgenticDBComputeClustersRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeAgenticDBComputeClustersRequest
	GetRegionId() *string
	SetStatus(v string) *DescribeAgenticDBComputeClustersRequest
	GetStatus() *string
}

type DescribeAgenticDBComputeClustersRequest struct {
	// example:
	//
	// pc-2ze8k
	ComputeClusterId *string `json:"ComputeClusterId,omitempty" xml:"ComputeClusterId,omitempty"`
	// The AgenticDB cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pagc-bp1abcdef1234567
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 30. Maximum value: 100.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status filter. Valid values: Running, Stopped, and Waiting.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAgenticDBComputeClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBComputeClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBComputeClustersRequest) GetComputeClusterId() *string {
	return s.ComputeClusterId
}

func (s *DescribeAgenticDBComputeClustersRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAgenticDBComputeClustersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAgenticDBComputeClustersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAgenticDBComputeClustersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAgenticDBComputeClustersRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeAgenticDBComputeClustersRequest) SetComputeClusterId(v string) *DescribeAgenticDBComputeClustersRequest {
	s.ComputeClusterId = &v
	return s
}

func (s *DescribeAgenticDBComputeClustersRequest) SetDBClusterId(v string) *DescribeAgenticDBComputeClustersRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAgenticDBComputeClustersRequest) SetPageNumber(v int32) *DescribeAgenticDBComputeClustersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAgenticDBComputeClustersRequest) SetPageSize(v int32) *DescribeAgenticDBComputeClustersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAgenticDBComputeClustersRequest) SetRegionId(v string) *DescribeAgenticDBComputeClustersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAgenticDBComputeClustersRequest) SetStatus(v string) *DescribeAgenticDBComputeClustersRequest {
	s.Status = &v
	return s
}

func (s *DescribeAgenticDBComputeClustersRequest) Validate() error {
	return dara.Validate(s)
}
