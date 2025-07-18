// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStreamingDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeStreamingDataSourceRequest
	GetDBInstanceId() *string
	SetDataSourceId(v int32) *DescribeStreamingDataSourceRequest
	GetDataSourceId() *int32
	SetRegionId(v string) *DescribeStreamingDataSourceRequest
	GetRegionId() *string
}

type DescribeStreamingDataSourceRequest struct {
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp10g78o9807yv9h3
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Data source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DataSourceId *int32 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// Region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) API to view available region IDs.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeStreamingDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeStreamingDataSourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeStreamingDataSourceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeStreamingDataSourceRequest) GetDataSourceId() *int32 {
	return s.DataSourceId
}

func (s *DescribeStreamingDataSourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeStreamingDataSourceRequest) SetDBInstanceId(v string) *DescribeStreamingDataSourceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeStreamingDataSourceRequest) SetDataSourceId(v int32) *DescribeStreamingDataSourceRequest {
	s.DataSourceId = &v
	return s
}

func (s *DescribeStreamingDataSourceRequest) SetRegionId(v string) *DescribeStreamingDataSourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeStreamingDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
