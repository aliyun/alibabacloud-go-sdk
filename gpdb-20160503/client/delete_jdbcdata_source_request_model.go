// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJDBCDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DeleteJDBCDataSourceRequest
	GetDBInstanceId() *string
	SetDataSourceId(v string) *DeleteJDBCDataSourceRequest
	GetDataSourceId() *string
	SetRegionId(v string) *DeleteJDBCDataSourceRequest
	GetRegionId() *string
}

type DeleteJDBCDataSourceRequest struct {
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Data source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// Region ID.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) interface to view available region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteJDBCDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteJDBCDataSourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteJDBCDataSourceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteJDBCDataSourceRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *DeleteJDBCDataSourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteJDBCDataSourceRequest) SetDBInstanceId(v string) *DeleteJDBCDataSourceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteJDBCDataSourceRequest) SetDataSourceId(v string) *DeleteJDBCDataSourceRequest {
	s.DataSourceId = &v
	return s
}

func (s *DeleteJDBCDataSourceRequest) SetRegionId(v string) *DeleteJDBCDataSourceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteJDBCDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
