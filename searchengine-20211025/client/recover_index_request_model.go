// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoverIndexRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBuildDeployId(v int32) *RecoverIndexRequest
	GetBuildDeployId() *int32
	SetDataSourceName(v string) *RecoverIndexRequest
	GetDataSourceName() *string
	SetGeneration(v string) *RecoverIndexRequest
	GetGeneration() *string
	SetIndexName(v string) *RecoverIndexRequest
	GetIndexName() *string
}

type RecoverIndexRequest struct {
	// The deployment ID of the data source.
	//
	// example:
	//
	// 277
	BuildDeployId *int32 `json:"buildDeployId,omitempty" xml:"buildDeployId,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// ha-cn-pl32rf0js04_odps_first
	DataSourceName *string `json:"dataSourceName,omitempty" xml:"dataSourceName,omitempty"`
	// The ID of the full index version.
	//
	// example:
	//
	// 1653018575
	Generation *string `json:"generation,omitempty" xml:"generation,omitempty"`
	// The index name.
	//
	// example:
	//
	// main_index
	IndexName *string `json:"indexName,omitempty" xml:"indexName,omitempty"`
}

func (s RecoverIndexRequest) String() string {
	return dara.Prettify(s)
}

func (s RecoverIndexRequest) GoString() string {
	return s.String()
}

func (s *RecoverIndexRequest) GetBuildDeployId() *int32 {
	return s.BuildDeployId
}

func (s *RecoverIndexRequest) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *RecoverIndexRequest) GetGeneration() *string {
	return s.Generation
}

func (s *RecoverIndexRequest) GetIndexName() *string {
	return s.IndexName
}

func (s *RecoverIndexRequest) SetBuildDeployId(v int32) *RecoverIndexRequest {
	s.BuildDeployId = &v
	return s
}

func (s *RecoverIndexRequest) SetDataSourceName(v string) *RecoverIndexRequest {
	s.DataSourceName = &v
	return s
}

func (s *RecoverIndexRequest) SetGeneration(v string) *RecoverIndexRequest {
	s.Generation = &v
	return s
}

func (s *RecoverIndexRequest) SetIndexName(v string) *RecoverIndexRequest {
	s.IndexName = &v
	return s
}

func (s *RecoverIndexRequest) Validate() error {
	return dara.Validate(s)
}
