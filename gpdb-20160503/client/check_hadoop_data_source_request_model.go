// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckHadoopDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckDir(v string) *CheckHadoopDataSourceRequest
	GetCheckDir() *string
	SetDBInstanceId(v string) *CheckHadoopDataSourceRequest
	GetDBInstanceId() *string
	SetDataSourceId(v string) *CheckHadoopDataSourceRequest
	GetDataSourceId() *string
	SetRegionId(v string) *CheckHadoopDataSourceRequest
	GetRegionId() *string
}

type CheckHadoopDataSourceRequest struct {
	// The Hadoop path that you want to check.
	//
	// This parameter is required.
	//
	// example:
	//
	// tmp
	CheckDir *string `json:"CheckDir,omitempty" xml:"CheckDir,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The data source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CheckHadoopDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckHadoopDataSourceRequest) GoString() string {
	return s.String()
}

func (s *CheckHadoopDataSourceRequest) GetCheckDir() *string {
	return s.CheckDir
}

func (s *CheckHadoopDataSourceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CheckHadoopDataSourceRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *CheckHadoopDataSourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckHadoopDataSourceRequest) SetCheckDir(v string) *CheckHadoopDataSourceRequest {
	s.CheckDir = &v
	return s
}

func (s *CheckHadoopDataSourceRequest) SetDBInstanceId(v string) *CheckHadoopDataSourceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CheckHadoopDataSourceRequest) SetDataSourceId(v string) *CheckHadoopDataSourceRequest {
	s.DataSourceId = &v
	return s
}

func (s *CheckHadoopDataSourceRequest) SetRegionId(v string) *CheckHadoopDataSourceRequest {
	s.RegionId = &v
	return s
}

func (s *CheckHadoopDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
