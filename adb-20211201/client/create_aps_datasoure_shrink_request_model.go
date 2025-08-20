// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApsDatasoureShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateApsDatasoureShrinkRequest
	GetDBClusterId() *string
	SetDatabricksInfoShrink(v string) *CreateApsDatasoureShrinkRequest
	GetDatabricksInfoShrink() *string
	SetDatasourceDescription(v string) *CreateApsDatasoureShrinkRequest
	GetDatasourceDescription() *string
	SetDatasourceName(v string) *CreateApsDatasoureShrinkRequest
	GetDatasourceName() *string
	SetDatasourceType(v string) *CreateApsDatasoureShrinkRequest
	GetDatasourceType() *string
	SetHiveInfoShrink(v string) *CreateApsDatasoureShrinkRequest
	GetHiveInfoShrink() *string
	SetKafkaInfoShrink(v string) *CreateApsDatasoureShrinkRequest
	GetKafkaInfoShrink() *string
	SetMode(v string) *CreateApsDatasoureShrinkRequest
	GetMode() *string
	SetPolarDBMysqlInfoShrink(v string) *CreateApsDatasoureShrinkRequest
	GetPolarDBMysqlInfoShrink() *string
	SetPolarDBXInfoShrink(v string) *CreateApsDatasoureShrinkRequest
	GetPolarDBXInfoShrink() *string
	SetRdsMysqlInfoShrink(v string) *CreateApsDatasoureShrinkRequest
	GetRdsMysqlInfoShrink() *string
	SetRegionId(v string) *CreateApsDatasoureShrinkRequest
	GetRegionId() *string
	SetSlsInfoShrink(v string) *CreateApsDatasoureShrinkRequest
	GetSlsInfoShrink() *string
}

type CreateApsDatasoureShrinkRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-xxxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The information about the Databricks data source.
	//
	// example:
	//
	// -
	DatabricksInfoShrink *string `json:"DatabricksInfo,omitempty" xml:"DatabricksInfo,omitempty"`
	// The description of the data source.
	//
	// example:
	//
	// description
	DatasourceDescription *string `json:"DatasourceDescription,omitempty" xml:"DatasourceDescription,omitempty"`
	// The name of the data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// sls-******
	DatasourceName *string `json:"DatasourceName,omitempty" xml:"DatasourceName,omitempty"`
	// The type of the data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// KAFKA
	DatasourceType *string `json:"DatasourceType,omitempty" xml:"DatasourceType,omitempty"`
	// The information about the Hive data source.
	//
	// example:
	//
	// -
	HiveInfoShrink *string `json:"HiveInfo,omitempty" xml:"HiveInfo,omitempty"`
	// The information about the source Apache Kafka instance.
	//
	// example:
	//
	// -
	KafkaInfoShrink *string `json:"KafkaInfo,omitempty" xml:"KafkaInfo,omitempty"`
	// The mode.
	//
	// example:
	//
	// ALI_CLOUD_INSTANCE
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The information about the source PolarDB for MySQL cluster.
	//
	// example:
	//
	// -
	PolarDBMysqlInfoShrink *string `json:"PolarDBMysqlInfo,omitempty" xml:"PolarDBMysqlInfo,omitempty"`
	// The information about the source PolarDB-X instance.
	//
	// example:
	//
	// -
	PolarDBXInfoShrink *string `json:"PolarDBXInfo,omitempty" xml:"PolarDBXInfo,omitempty"`
	// The information about the source ApsaraDB RDS for MySQL instance.
	//
	// example:
	//
	// -
	RdsMysqlInfoShrink *string `json:"RdsMysqlInfo,omitempty" xml:"RdsMysqlInfo,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The information about the source Simple Log Service (SLS) instance or cluster.
	//
	// example:
	//
	// -
	SlsInfoShrink *string `json:"SlsInfo,omitempty" xml:"SlsInfo,omitempty"`
}

func (s CreateApsDatasoureShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApsDatasoureShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateApsDatasoureShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateApsDatasoureShrinkRequest) GetDatabricksInfoShrink() *string {
	return s.DatabricksInfoShrink
}

func (s *CreateApsDatasoureShrinkRequest) GetDatasourceDescription() *string {
	return s.DatasourceDescription
}

func (s *CreateApsDatasoureShrinkRequest) GetDatasourceName() *string {
	return s.DatasourceName
}

func (s *CreateApsDatasoureShrinkRequest) GetDatasourceType() *string {
	return s.DatasourceType
}

func (s *CreateApsDatasoureShrinkRequest) GetHiveInfoShrink() *string {
	return s.HiveInfoShrink
}

func (s *CreateApsDatasoureShrinkRequest) GetKafkaInfoShrink() *string {
	return s.KafkaInfoShrink
}

func (s *CreateApsDatasoureShrinkRequest) GetMode() *string {
	return s.Mode
}

func (s *CreateApsDatasoureShrinkRequest) GetPolarDBMysqlInfoShrink() *string {
	return s.PolarDBMysqlInfoShrink
}

func (s *CreateApsDatasoureShrinkRequest) GetPolarDBXInfoShrink() *string {
	return s.PolarDBXInfoShrink
}

func (s *CreateApsDatasoureShrinkRequest) GetRdsMysqlInfoShrink() *string {
	return s.RdsMysqlInfoShrink
}

func (s *CreateApsDatasoureShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateApsDatasoureShrinkRequest) GetSlsInfoShrink() *string {
	return s.SlsInfoShrink
}

func (s *CreateApsDatasoureShrinkRequest) SetDBClusterId(v string) *CreateApsDatasoureShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateApsDatasoureShrinkRequest) SetDatabricksInfoShrink(v string) *CreateApsDatasoureShrinkRequest {
	s.DatabricksInfoShrink = &v
	return s
}

func (s *CreateApsDatasoureShrinkRequest) SetDatasourceDescription(v string) *CreateApsDatasoureShrinkRequest {
	s.DatasourceDescription = &v
	return s
}

func (s *CreateApsDatasoureShrinkRequest) SetDatasourceName(v string) *CreateApsDatasoureShrinkRequest {
	s.DatasourceName = &v
	return s
}

func (s *CreateApsDatasoureShrinkRequest) SetDatasourceType(v string) *CreateApsDatasoureShrinkRequest {
	s.DatasourceType = &v
	return s
}

func (s *CreateApsDatasoureShrinkRequest) SetHiveInfoShrink(v string) *CreateApsDatasoureShrinkRequest {
	s.HiveInfoShrink = &v
	return s
}

func (s *CreateApsDatasoureShrinkRequest) SetKafkaInfoShrink(v string) *CreateApsDatasoureShrinkRequest {
	s.KafkaInfoShrink = &v
	return s
}

func (s *CreateApsDatasoureShrinkRequest) SetMode(v string) *CreateApsDatasoureShrinkRequest {
	s.Mode = &v
	return s
}

func (s *CreateApsDatasoureShrinkRequest) SetPolarDBMysqlInfoShrink(v string) *CreateApsDatasoureShrinkRequest {
	s.PolarDBMysqlInfoShrink = &v
	return s
}

func (s *CreateApsDatasoureShrinkRequest) SetPolarDBXInfoShrink(v string) *CreateApsDatasoureShrinkRequest {
	s.PolarDBXInfoShrink = &v
	return s
}

func (s *CreateApsDatasoureShrinkRequest) SetRdsMysqlInfoShrink(v string) *CreateApsDatasoureShrinkRequest {
	s.RdsMysqlInfoShrink = &v
	return s
}

func (s *CreateApsDatasoureShrinkRequest) SetRegionId(v string) *CreateApsDatasoureShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateApsDatasoureShrinkRequest) SetSlsInfoShrink(v string) *CreateApsDatasoureShrinkRequest {
	s.SlsInfoShrink = &v
	return s
}

func (s *CreateApsDatasoureShrinkRequest) Validate() error {
	return dara.Validate(s)
}
