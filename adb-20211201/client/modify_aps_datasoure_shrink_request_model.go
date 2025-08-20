// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApsDatasoureShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyApsDatasoureShrinkRequest
	GetDBClusterId() *string
	SetDatasourceDescription(v string) *ModifyApsDatasoureShrinkRequest
	GetDatasourceDescription() *string
	SetDatasourceId(v int64) *ModifyApsDatasoureShrinkRequest
	GetDatasourceId() *int64
	SetDatasourceName(v string) *ModifyApsDatasoureShrinkRequest
	GetDatasourceName() *string
	SetKafkaInfoShrink(v string) *ModifyApsDatasoureShrinkRequest
	GetKafkaInfoShrink() *string
	SetLakehouseIdShrink(v string) *ModifyApsDatasoureShrinkRequest
	GetLakehouseIdShrink() *string
	SetPolarDBMysqlInfoShrink(v string) *ModifyApsDatasoureShrinkRequest
	GetPolarDBMysqlInfoShrink() *string
	SetRdsMysqlInfoShrink(v string) *ModifyApsDatasoureShrinkRequest
	GetRdsMysqlInfoShrink() *string
	SetRegionId(v string) *ModifyApsDatasoureShrinkRequest
	GetRegionId() *string
	SetSlsInfoShrink(v string) *ModifyApsDatasoureShrinkRequest
	GetSlsInfoShrink() *string
}

type ModifyApsDatasoureShrinkRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-*******
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The description of the data source.
	//
	// example:
	//
	// test
	DatasourceDescription *string `json:"DatasourceDescription,omitempty" xml:"DatasourceDescription,omitempty"`
	// The data source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DatasourceId *int64 `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// sls-******
	DatasourceName *string `json:"DatasourceName,omitempty" xml:"DatasourceName,omitempty"`
	// The information about the Kafka instance.
	//
	// example:
	//
	// -
	KafkaInfoShrink *string `json:"KafkaInfo,omitempty" xml:"KafkaInfo,omitempty"`
	// The lakehouse ID.
	//
	// example:
	//
	// 123
	LakehouseIdShrink *string `json:"LakehouseId,omitempty" xml:"LakehouseId,omitempty"`
	// The parameter is no longer supported.
	//
	// example:
	//
	// -
	PolarDBMysqlInfoShrink *string `json:"PolarDBMysqlInfo,omitempty" xml:"PolarDBMysqlInfo,omitempty"`
	// The parameter is no longer supported.
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
	// The information about Simple Log Service (SLS).
	//
	// example:
	//
	// -
	SlsInfoShrink *string `json:"SlsInfo,omitempty" xml:"SlsInfo,omitempty"`
}

func (s ModifyApsDatasoureShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyApsDatasoureShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyApsDatasoureShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyApsDatasoureShrinkRequest) GetDatasourceDescription() *string {
	return s.DatasourceDescription
}

func (s *ModifyApsDatasoureShrinkRequest) GetDatasourceId() *int64 {
	return s.DatasourceId
}

func (s *ModifyApsDatasoureShrinkRequest) GetDatasourceName() *string {
	return s.DatasourceName
}

func (s *ModifyApsDatasoureShrinkRequest) GetKafkaInfoShrink() *string {
	return s.KafkaInfoShrink
}

func (s *ModifyApsDatasoureShrinkRequest) GetLakehouseIdShrink() *string {
	return s.LakehouseIdShrink
}

func (s *ModifyApsDatasoureShrinkRequest) GetPolarDBMysqlInfoShrink() *string {
	return s.PolarDBMysqlInfoShrink
}

func (s *ModifyApsDatasoureShrinkRequest) GetRdsMysqlInfoShrink() *string {
	return s.RdsMysqlInfoShrink
}

func (s *ModifyApsDatasoureShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyApsDatasoureShrinkRequest) GetSlsInfoShrink() *string {
	return s.SlsInfoShrink
}

func (s *ModifyApsDatasoureShrinkRequest) SetDBClusterId(v string) *ModifyApsDatasoureShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyApsDatasoureShrinkRequest) SetDatasourceDescription(v string) *ModifyApsDatasoureShrinkRequest {
	s.DatasourceDescription = &v
	return s
}

func (s *ModifyApsDatasoureShrinkRequest) SetDatasourceId(v int64) *ModifyApsDatasoureShrinkRequest {
	s.DatasourceId = &v
	return s
}

func (s *ModifyApsDatasoureShrinkRequest) SetDatasourceName(v string) *ModifyApsDatasoureShrinkRequest {
	s.DatasourceName = &v
	return s
}

func (s *ModifyApsDatasoureShrinkRequest) SetKafkaInfoShrink(v string) *ModifyApsDatasoureShrinkRequest {
	s.KafkaInfoShrink = &v
	return s
}

func (s *ModifyApsDatasoureShrinkRequest) SetLakehouseIdShrink(v string) *ModifyApsDatasoureShrinkRequest {
	s.LakehouseIdShrink = &v
	return s
}

func (s *ModifyApsDatasoureShrinkRequest) SetPolarDBMysqlInfoShrink(v string) *ModifyApsDatasoureShrinkRequest {
	s.PolarDBMysqlInfoShrink = &v
	return s
}

func (s *ModifyApsDatasoureShrinkRequest) SetRdsMysqlInfoShrink(v string) *ModifyApsDatasoureShrinkRequest {
	s.RdsMysqlInfoShrink = &v
	return s
}

func (s *ModifyApsDatasoureShrinkRequest) SetRegionId(v string) *ModifyApsDatasoureShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyApsDatasoureShrinkRequest) SetSlsInfoShrink(v string) *ModifyApsDatasoureShrinkRequest {
	s.SlsInfoShrink = &v
	return s
}

func (s *ModifyApsDatasoureShrinkRequest) Validate() error {
	return dara.Validate(s)
}
