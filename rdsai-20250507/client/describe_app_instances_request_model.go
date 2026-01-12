// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *DescribeAppInstancesRequest
	GetAppType() *string
	SetDBInstanceName(v string) *DescribeAppInstancesRequest
	GetDBInstanceName() *string
	SetPageNumber(v int64) *DescribeAppInstancesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeAppInstancesRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeAppInstancesRequest
	GetRegionId() *string
}

type DescribeAppInstancesRequest struct {
	// The ID of the RDS for PostgreSQL instance with which the RDS Supabase instances are associated. If you specify this parameter, the RDS Supabase instances associated with the specified RDS for PostgreSQL instance are queried.
	//
	// example:
	//
	// supabase
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// pgm-2ze49qv594vi****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The number of records per page. Valid values: **1 to 50**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The application type. Only **supabase*	- is supported. For more information, see [RDS Supabase](https://help.aliyun.com/document_detail/2938735.html).
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The operation that you want to perform. Set the value to **DescribeAppInstances**.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAppInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppInstancesRequest) GetAppType() *string {
	return s.AppType
}

func (s *DescribeAppInstancesRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeAppInstancesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeAppInstancesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeAppInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAppInstancesRequest) SetAppType(v string) *DescribeAppInstancesRequest {
	s.AppType = &v
	return s
}

func (s *DescribeAppInstancesRequest) SetDBInstanceName(v string) *DescribeAppInstancesRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeAppInstancesRequest) SetPageNumber(v int64) *DescribeAppInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAppInstancesRequest) SetPageSize(v int64) *DescribeAppInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAppInstancesRequest) SetRegionId(v string) *DescribeAppInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAppInstancesRequest) Validate() error {
	return dara.Validate(s)
}
