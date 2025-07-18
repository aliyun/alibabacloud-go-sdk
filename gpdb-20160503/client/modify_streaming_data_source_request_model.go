// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStreamingDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyStreamingDataSourceRequest
	GetDBInstanceId() *string
	SetDataSourceConfig(v string) *ModifyStreamingDataSourceRequest
	GetDataSourceConfig() *string
	SetDataSourceDescription(v string) *ModifyStreamingDataSourceRequest
	GetDataSourceDescription() *string
	SetDataSourceId(v string) *ModifyStreamingDataSourceRequest
	GetDataSourceId() *string
	SetRegionId(v string) *ModifyStreamingDataSourceRequest
	GetRegionId() *string
}

type ModifyStreamingDataSourceRequest struct {
	// The instance ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/196830.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-k2j36a3172b102593
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The configurations of the data source.
	//
	// example:
	//
	// {"brokers":"broker0:9091,broker1:9091","topic":"topic"}
	DataSourceConfig *string `json:"DataSourceConfig,omitempty" xml:"DataSourceConfig,omitempty"`
	// The description of the data source.
	//
	// example:
	//
	// test-kafka
	DataSourceDescription *string `json:"DataSourceDescription,omitempty" xml:"DataSourceDescription,omitempty"`
	// The data source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 57
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyStreamingDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyStreamingDataSourceRequest) GoString() string {
	return s.String()
}

func (s *ModifyStreamingDataSourceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyStreamingDataSourceRequest) GetDataSourceConfig() *string {
	return s.DataSourceConfig
}

func (s *ModifyStreamingDataSourceRequest) GetDataSourceDescription() *string {
	return s.DataSourceDescription
}

func (s *ModifyStreamingDataSourceRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *ModifyStreamingDataSourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyStreamingDataSourceRequest) SetDBInstanceId(v string) *ModifyStreamingDataSourceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyStreamingDataSourceRequest) SetDataSourceConfig(v string) *ModifyStreamingDataSourceRequest {
	s.DataSourceConfig = &v
	return s
}

func (s *ModifyStreamingDataSourceRequest) SetDataSourceDescription(v string) *ModifyStreamingDataSourceRequest {
	s.DataSourceDescription = &v
	return s
}

func (s *ModifyStreamingDataSourceRequest) SetDataSourceId(v string) *ModifyStreamingDataSourceRequest {
	s.DataSourceId = &v
	return s
}

func (s *ModifyStreamingDataSourceRequest) SetRegionId(v string) *ModifyStreamingDataSourceRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyStreamingDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
