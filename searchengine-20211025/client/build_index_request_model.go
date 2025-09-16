// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBuildIndexRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBuildMode(v string) *BuildIndexRequest
	GetBuildMode() *string
	SetDataSourceName(v string) *BuildIndexRequest
	GetDataSourceName() *string
	SetDataSourceType(v string) *BuildIndexRequest
	GetDataSourceType() *string
	SetDataTimeSec(v int32) *BuildIndexRequest
	GetDataTimeSec() *int32
	SetDomain(v string) *BuildIndexRequest
	GetDomain() *string
	SetGeneration(v int64) *BuildIndexRequest
	GetGeneration() *int64
	SetPartition(v string) *BuildIndexRequest
	GetPartition() *string
	SetPath(v string) *BuildIndexRequest
	GetPath() *string
	SetTag(v string) *BuildIndexRequest
	GetTag() *string
}

type BuildIndexRequest struct {
	// The reindexing method. Valid values: api: API data source. indexRecover: data recovery by using indexing.
	//
	// example:
	//
	// indexRecover
	BuildMode *string `json:"buildMode,omitempty" xml:"buildMode,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// my_data_source
	DataSourceName *string `json:"dataSourceName,omitempty" xml:"dataSourceName,omitempty"`
	// The type of the data source.
	//
	// example:
	//
	// swift
	DataSourceType *string `json:"dataSourceType,omitempty" xml:"dataSourceType,omitempty"`
	// The timestamp in seconds. The value must be of the INTEGER type. This parameter is required if you specify an API data source.
	//
	// example:
	//
	// 1640867288
	DataTimeSec *int32 `json:"dataTimeSec,omitempty" xml:"dataTimeSec,omitempty"`
	// The data center in which the data source is deployed.
	//
	// example:
	//
	// test
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The data restoration version.
	//
	// example:
	//
	// 160131146
	Generation *int64 `json:"generation,omitempty" xml:"generation,omitempty"`
	// The partition in the MaxCompute table. This parameter is required if type is set to odps.
	//
	// example:
	//
	// 20201010
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
	Path      *string `json:"path,omitempty" xml:"path,omitempty"`
	Tag       *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s BuildIndexRequest) String() string {
	return dara.Prettify(s)
}

func (s BuildIndexRequest) GoString() string {
	return s.String()
}

func (s *BuildIndexRequest) GetBuildMode() *string {
	return s.BuildMode
}

func (s *BuildIndexRequest) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *BuildIndexRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *BuildIndexRequest) GetDataTimeSec() *int32 {
	return s.DataTimeSec
}

func (s *BuildIndexRequest) GetDomain() *string {
	return s.Domain
}

func (s *BuildIndexRequest) GetGeneration() *int64 {
	return s.Generation
}

func (s *BuildIndexRequest) GetPartition() *string {
	return s.Partition
}

func (s *BuildIndexRequest) GetPath() *string {
	return s.Path
}

func (s *BuildIndexRequest) GetTag() *string {
	return s.Tag
}

func (s *BuildIndexRequest) SetBuildMode(v string) *BuildIndexRequest {
	s.BuildMode = &v
	return s
}

func (s *BuildIndexRequest) SetDataSourceName(v string) *BuildIndexRequest {
	s.DataSourceName = &v
	return s
}

func (s *BuildIndexRequest) SetDataSourceType(v string) *BuildIndexRequest {
	s.DataSourceType = &v
	return s
}

func (s *BuildIndexRequest) SetDataTimeSec(v int32) *BuildIndexRequest {
	s.DataTimeSec = &v
	return s
}

func (s *BuildIndexRequest) SetDomain(v string) *BuildIndexRequest {
	s.Domain = &v
	return s
}

func (s *BuildIndexRequest) SetGeneration(v int64) *BuildIndexRequest {
	s.Generation = &v
	return s
}

func (s *BuildIndexRequest) SetPartition(v string) *BuildIndexRequest {
	s.Partition = &v
	return s
}

func (s *BuildIndexRequest) SetPath(v string) *BuildIndexRequest {
	s.Path = &v
	return s
}

func (s *BuildIndexRequest) SetTag(v string) *BuildIndexRequest {
	s.Tag = &v
	return s
}

func (s *BuildIndexRequest) Validate() error {
	return dara.Validate(s)
}
