// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataSourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeDataSourcesRequest
	GetClusterId() *string
	SetDataSourceId(v string) *DescribeDataSourcesRequest
	GetDataSourceId() *string
	SetDataSourceName(v string) *DescribeDataSourcesRequest
	GetDataSourceName() *string
	SetDataSourceType(v string) *DescribeDataSourcesRequest
	GetDataSourceType() *string
	SetPageNumber(v int32) *DescribeDataSourcesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDataSourcesRequest
	GetPageSize() *int32
}

type DescribeDataSourcesRequest struct {
	// example:
	//
	// cl-0003jyv******fsku5m
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// ds-000******2nqeo
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// example:
	//
	// MyLocalNas
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// example:
	//
	// COMMON_NAS
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDataSourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataSourcesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeDataSourcesRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *DescribeDataSourcesRequest) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *DescribeDataSourcesRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *DescribeDataSourcesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDataSourcesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDataSourcesRequest) SetClusterId(v string) *DescribeDataSourcesRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeDataSourcesRequest) SetDataSourceId(v string) *DescribeDataSourcesRequest {
	s.DataSourceId = &v
	return s
}

func (s *DescribeDataSourcesRequest) SetDataSourceName(v string) *DescribeDataSourcesRequest {
	s.DataSourceName = &v
	return s
}

func (s *DescribeDataSourcesRequest) SetDataSourceType(v string) *DescribeDataSourcesRequest {
	s.DataSourceType = &v
	return s
}

func (s *DescribeDataSourcesRequest) SetPageNumber(v int32) *DescribeDataSourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDataSourcesRequest) SetPageSize(v int32) *DescribeDataSourcesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDataSourcesRequest) Validate() error {
	return dara.Validate(s)
}
