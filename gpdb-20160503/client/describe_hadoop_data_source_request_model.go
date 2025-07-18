// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHadoopDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeHadoopDataSourceRequest
	GetDBInstanceId() *string
	SetDataSourceId(v string) *DescribeHadoopDataSourceRequest
	GetDataSourceId() *string
	SetRegionId(v string) *DescribeHadoopDataSourceRequest
	GetRegionId() *string
}

type DescribeHadoopDataSourceRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The data source ID.
	//
	// example:
	//
	// 123
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The region ID of the instance.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/2361846.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeHadoopDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHadoopDataSourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeHadoopDataSourceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeHadoopDataSourceRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *DescribeHadoopDataSourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHadoopDataSourceRequest) SetDBInstanceId(v string) *DescribeHadoopDataSourceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeHadoopDataSourceRequest) SetDataSourceId(v string) *DescribeHadoopDataSourceRequest {
	s.DataSourceId = &v
	return s
}

func (s *DescribeHadoopDataSourceRequest) SetRegionId(v string) *DescribeHadoopDataSourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHadoopDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
