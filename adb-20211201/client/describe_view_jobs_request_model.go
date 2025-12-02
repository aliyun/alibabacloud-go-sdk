// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeViewJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeViewJobsRequest
	GetDBClusterId() *string
	SetFilterOwner(v string) *DescribeViewJobsRequest
	GetFilterOwner() *string
	SetFilterViewName(v string) *DescribeViewJobsRequest
	GetFilterViewName() *string
	SetFilterViewType(v string) *DescribeViewJobsRequest
	GetFilterViewType() *string
	SetOrderBy(v string) *DescribeViewJobsRequest
	GetOrderBy() *string
	SetPageNumber(v int32) *DescribeViewJobsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeViewJobsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeViewJobsRequest
	GetRegionId() *string
	SetSchemaName(v string) *DescribeViewJobsRequest
	GetSchemaName() *string
}

type DescribeViewJobsRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// am-bp1ub9grke1****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The owner of the view.
	//
	// example:
	//
	// admin
	FilterOwner *string `json:"FilterOwner,omitempty" xml:"FilterOwner,omitempty"`
	// The name of the view.
	//
	// example:
	//
	// test_mv
	FilterViewName *string `json:"FilterViewName,omitempty" xml:"FilterViewName,omitempty"`
	// The type of the view.
	//
	// Valid values:
	//
	// \\-VIRTUAL_VIEW
	//
	// \\-MATERIALIZED_VIEW
	//
	// This parameter is empty by default.
	//
	// example:
	//
	// MATERIALIZED_VIEW
	FilterViewType *string `json:"FilterViewType,omitempty" xml:"FilterViewType,omitempty"`
	// The field used for sorting. Valid values for Type:
	//
	// 	- Asc.
	//
	// 	- Desc.
	//
	// Valid values for Field:
	//
	// 	- StartTime.
	//
	// 	- EndTime;
	//
	// 	- ScheduledStartTime;
	//
	// example:
	//
	// {\\"Field\\":\\"StartTime\\",\\"Type\\":\\"Desc\\"}
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The database name.
	//
	// example:
	//
	// demo_db
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
}

func (s DescribeViewJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeViewJobsRequest) GoString() string {
	return s.String()
}

func (s *DescribeViewJobsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeViewJobsRequest) GetFilterOwner() *string {
	return s.FilterOwner
}

func (s *DescribeViewJobsRequest) GetFilterViewName() *string {
	return s.FilterViewName
}

func (s *DescribeViewJobsRequest) GetFilterViewType() *string {
	return s.FilterViewType
}

func (s *DescribeViewJobsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *DescribeViewJobsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeViewJobsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeViewJobsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeViewJobsRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeViewJobsRequest) SetDBClusterId(v string) *DescribeViewJobsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeViewJobsRequest) SetFilterOwner(v string) *DescribeViewJobsRequest {
	s.FilterOwner = &v
	return s
}

func (s *DescribeViewJobsRequest) SetFilterViewName(v string) *DescribeViewJobsRequest {
	s.FilterViewName = &v
	return s
}

func (s *DescribeViewJobsRequest) SetFilterViewType(v string) *DescribeViewJobsRequest {
	s.FilterViewType = &v
	return s
}

func (s *DescribeViewJobsRequest) SetOrderBy(v string) *DescribeViewJobsRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeViewJobsRequest) SetPageNumber(v int32) *DescribeViewJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeViewJobsRequest) SetPageSize(v int32) *DescribeViewJobsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeViewJobsRequest) SetRegionId(v string) *DescribeViewJobsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeViewJobsRequest) SetSchemaName(v string) *DescribeViewJobsRequest {
	s.SchemaName = &v
	return s
}

func (s *DescribeViewJobsRequest) Validate() error {
	return dara.Validate(s)
}
