// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHadoopDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DeleteHadoopDataSourceRequest
	GetDBInstanceId() *string
	SetDataSourceId(v string) *DeleteHadoopDataSourceRequest
	GetDataSourceId() *string
	SetRegionId(v string) *DeleteHadoopDataSourceRequest
	GetRegionId() *string
}

type DeleteHadoopDataSourceRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteHadoopDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHadoopDataSourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteHadoopDataSourceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteHadoopDataSourceRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *DeleteHadoopDataSourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteHadoopDataSourceRequest) SetDBInstanceId(v string) *DeleteHadoopDataSourceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteHadoopDataSourceRequest) SetDataSourceId(v string) *DeleteHadoopDataSourceRequest {
	s.DataSourceId = &v
	return s
}

func (s *DeleteHadoopDataSourceRequest) SetRegionId(v string) *DeleteHadoopDataSourceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteHadoopDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
