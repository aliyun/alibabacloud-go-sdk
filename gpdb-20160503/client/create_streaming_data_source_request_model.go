// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStreamingDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CreateStreamingDataSourceRequest
	GetDBInstanceId() *string
	SetDataSourceConfig(v string) *CreateStreamingDataSourceRequest
	GetDataSourceConfig() *string
	SetDataSourceDescription(v string) *CreateStreamingDataSourceRequest
	GetDataSourceDescription() *string
	SetDataSourceName(v string) *CreateStreamingDataSourceRequest
	GetDataSourceName() *string
	SetDataSourceType(v string) *CreateStreamingDataSourceRequest
	GetDataSourceType() *string
	SetRegionId(v string) *CreateStreamingDataSourceRequest
	GetRegionId() *string
	SetServiceId(v int32) *CreateStreamingDataSourceRequest
	GetServiceId() *int32
}

type CreateStreamingDataSourceRequest struct {
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// CreateExternalDataSource
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Data source configuration information.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"broker_list":"broker0:9091,broker1:9091","topic":"topic"}
	DataSourceConfig *string `json:"DataSourceConfig,omitempty" xml:"DataSourceConfig,omitempty"`
	// Data source description.
	//
	// example:
	//
	// test-kafka
	DataSourceDescription *string `json:"DataSourceDescription,omitempty" xml:"DataSourceDescription,omitempty"`
	// Data source name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-kafka
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// Data source type. Values:
	//
	//  -  kafka
	//
	// This parameter is required.
	//
	// example:
	//
	// kafka
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// Region ID.
	//
	// > You can view available region IDs through the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) interface.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Real-time data service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ServiceId *int32 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s CreateStreamingDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStreamingDataSourceRequest) GoString() string {
	return s.String()
}

func (s *CreateStreamingDataSourceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateStreamingDataSourceRequest) GetDataSourceConfig() *string {
	return s.DataSourceConfig
}

func (s *CreateStreamingDataSourceRequest) GetDataSourceDescription() *string {
	return s.DataSourceDescription
}

func (s *CreateStreamingDataSourceRequest) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *CreateStreamingDataSourceRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *CreateStreamingDataSourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateStreamingDataSourceRequest) GetServiceId() *int32 {
	return s.ServiceId
}

func (s *CreateStreamingDataSourceRequest) SetDBInstanceId(v string) *CreateStreamingDataSourceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateStreamingDataSourceRequest) SetDataSourceConfig(v string) *CreateStreamingDataSourceRequest {
	s.DataSourceConfig = &v
	return s
}

func (s *CreateStreamingDataSourceRequest) SetDataSourceDescription(v string) *CreateStreamingDataSourceRequest {
	s.DataSourceDescription = &v
	return s
}

func (s *CreateStreamingDataSourceRequest) SetDataSourceName(v string) *CreateStreamingDataSourceRequest {
	s.DataSourceName = &v
	return s
}

func (s *CreateStreamingDataSourceRequest) SetDataSourceType(v string) *CreateStreamingDataSourceRequest {
	s.DataSourceType = &v
	return s
}

func (s *CreateStreamingDataSourceRequest) SetRegionId(v string) *CreateStreamingDataSourceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateStreamingDataSourceRequest) SetServiceId(v int32) *CreateStreamingDataSourceRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateStreamingDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
