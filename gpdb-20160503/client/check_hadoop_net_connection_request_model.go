// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckHadoopNetConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CheckHadoopNetConnectionRequest
	GetDBInstanceId() *string
	SetDataSourceId(v string) *CheckHadoopNetConnectionRequest
	GetDataSourceId() *string
	SetEmrInstanceId(v string) *CheckHadoopNetConnectionRequest
	GetEmrInstanceId() *string
	SetRegionId(v string) *CheckHadoopNetConnectionRequest
	GetRegionId() *string
}

type CheckHadoopNetConnectionRequest struct {
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// 1. Either DataSourceId or EmrInstanceId must be specified as input, otherwise an error will occur.
	//
	// 2. If both of the above parameters are specified, EmrInstanceId will be used preferentially.
	//
	// 3. If the data source specified by DataSourceId is a self-built Hadoop cluster, an error will occur directly.
	//
	// example:
	//
	// 126
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// EMR instance ID.
	//
	// example:
	//
	// c-xxx
	EmrInstanceId *string `json:"EmrInstanceId,omitempty" xml:"EmrInstanceId,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CheckHadoopNetConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckHadoopNetConnectionRequest) GoString() string {
	return s.String()
}

func (s *CheckHadoopNetConnectionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CheckHadoopNetConnectionRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *CheckHadoopNetConnectionRequest) GetEmrInstanceId() *string {
	return s.EmrInstanceId
}

func (s *CheckHadoopNetConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckHadoopNetConnectionRequest) SetDBInstanceId(v string) *CheckHadoopNetConnectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CheckHadoopNetConnectionRequest) SetDataSourceId(v string) *CheckHadoopNetConnectionRequest {
	s.DataSourceId = &v
	return s
}

func (s *CheckHadoopNetConnectionRequest) SetEmrInstanceId(v string) *CheckHadoopNetConnectionRequest {
	s.EmrInstanceId = &v
	return s
}

func (s *CheckHadoopNetConnectionRequest) SetRegionId(v string) *CheckHadoopNetConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *CheckHadoopNetConnectionRequest) Validate() error {
	return dara.Validate(s)
}
