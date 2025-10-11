// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationIds(v string) *DescribeApplicationsRequest
	GetApplicationIds() *string
	SetDBClusterId(v string) *DescribeApplicationsRequest
	GetDBClusterId() *string
	SetPageNumber(v int32) *DescribeApplicationsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeApplicationsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeApplicationsRequest
	GetRegionId() *string
}

type DescribeApplicationsRequest struct {
	// example:
	//
	// pa-**************
	ApplicationIds *string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty"`
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeApplicationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationsRequest) GetApplicationIds() *string {
	return s.ApplicationIds
}

func (s *DescribeApplicationsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeApplicationsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeApplicationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApplicationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApplicationsRequest) SetApplicationIds(v string) *DescribeApplicationsRequest {
	s.ApplicationIds = &v
	return s
}

func (s *DescribeApplicationsRequest) SetDBClusterId(v string) *DescribeApplicationsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeApplicationsRequest) SetPageNumber(v int32) *DescribeApplicationsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApplicationsRequest) SetPageSize(v int32) *DescribeApplicationsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApplicationsRequest) SetRegionId(v string) *DescribeApplicationsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApplicationsRequest) Validate() error {
	return dara.Validate(s)
}
