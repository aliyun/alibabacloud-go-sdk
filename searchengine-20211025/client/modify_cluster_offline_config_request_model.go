// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterOfflineConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBuildMode(v string) *ModifyClusterOfflineConfigRequest
	GetBuildMode() *string
	SetConfig(v map[string]*int32) *ModifyClusterOfflineConfigRequest
	GetConfig() map[string]*int32
	SetDataSourceName(v string) *ModifyClusterOfflineConfigRequest
	GetDataSourceName() *string
	SetDataSourceType(v string) *ModifyClusterOfflineConfigRequest
	GetDataSourceType() *string
	SetDataTimeSec(v int32) *ModifyClusterOfflineConfigRequest
	GetDataTimeSec() *int32
	SetDomain(v string) *ModifyClusterOfflineConfigRequest
	GetDomain() *string
	SetGeneration(v int64) *ModifyClusterOfflineConfigRequest
	GetGeneration() *int64
	SetPartition(v string) *ModifyClusterOfflineConfigRequest
	GetPartition() *string
	SetPushMode(v string) *ModifyClusterOfflineConfigRequest
	GetPushMode() *string
}

type ModifyClusterOfflineConfigRequest struct {
	// The reindexing method. Valid values: api: API data source. indexRecover: data recovery by using indexing.
	//
	// example:
	//
	// indexRecover
	BuildMode *string `json:"buildMode,omitempty" xml:"buildMode,omitempty"`
	// The configuration name, which is stored as a key.
	Config map[string]*int32 `json:"config,omitempty" xml:"config,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// test1
	DataSourceName *string `json:"dataSourceName,omitempty" xml:"dataSourceName,omitempty"`
	// The type of the data source. Valid values: odps: MaxCompute. swift: Swift. unKnow: unknown type.
	//
	// example:
	//
	// maxComputer
	DataSourceType *string `json:"dataSourceType,omitempty" xml:"dataSourceType,omitempty"`
	// This parameter is required when index building by using API data sources is triggered.
	//
	// example:
	//
	// 1640867288
	DataTimeSec *int32 `json:"dataTimeSec,omitempty" xml:"dataTimeSec,omitempty"`
	// The data center in which the data source is deployed.
	//
	// example:
	//
	// vpc_hz_domain_1
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The ID of the full index version.
	//
	// example:
	//
	// 160142641
	Generation *int64 `json:"generation,omitempty" xml:"generation,omitempty"`
	// This parameter is required when index building for full data in a MaxCompute data source is triggered.
	//
	// example:
	//
	// 20211202
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
	// The push mode of the configuration. By default, only the configuration is pushed.
	//
	// example:
	//
	// PUSH_ONLY
	PushMode *string `json:"pushMode,omitempty" xml:"pushMode,omitempty"`
}

func (s ModifyClusterOfflineConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterOfflineConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterOfflineConfigRequest) GetBuildMode() *string {
	return s.BuildMode
}

func (s *ModifyClusterOfflineConfigRequest) GetConfig() map[string]*int32 {
	return s.Config
}

func (s *ModifyClusterOfflineConfigRequest) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *ModifyClusterOfflineConfigRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *ModifyClusterOfflineConfigRequest) GetDataTimeSec() *int32 {
	return s.DataTimeSec
}

func (s *ModifyClusterOfflineConfigRequest) GetDomain() *string {
	return s.Domain
}

func (s *ModifyClusterOfflineConfigRequest) GetGeneration() *int64 {
	return s.Generation
}

func (s *ModifyClusterOfflineConfigRequest) GetPartition() *string {
	return s.Partition
}

func (s *ModifyClusterOfflineConfigRequest) GetPushMode() *string {
	return s.PushMode
}

func (s *ModifyClusterOfflineConfigRequest) SetBuildMode(v string) *ModifyClusterOfflineConfigRequest {
	s.BuildMode = &v
	return s
}

func (s *ModifyClusterOfflineConfigRequest) SetConfig(v map[string]*int32) *ModifyClusterOfflineConfigRequest {
	s.Config = v
	return s
}

func (s *ModifyClusterOfflineConfigRequest) SetDataSourceName(v string) *ModifyClusterOfflineConfigRequest {
	s.DataSourceName = &v
	return s
}

func (s *ModifyClusterOfflineConfigRequest) SetDataSourceType(v string) *ModifyClusterOfflineConfigRequest {
	s.DataSourceType = &v
	return s
}

func (s *ModifyClusterOfflineConfigRequest) SetDataTimeSec(v int32) *ModifyClusterOfflineConfigRequest {
	s.DataTimeSec = &v
	return s
}

func (s *ModifyClusterOfflineConfigRequest) SetDomain(v string) *ModifyClusterOfflineConfigRequest {
	s.Domain = &v
	return s
}

func (s *ModifyClusterOfflineConfigRequest) SetGeneration(v int64) *ModifyClusterOfflineConfigRequest {
	s.Generation = &v
	return s
}

func (s *ModifyClusterOfflineConfigRequest) SetPartition(v string) *ModifyClusterOfflineConfigRequest {
	s.Partition = &v
	return s
}

func (s *ModifyClusterOfflineConfigRequest) SetPushMode(v string) *ModifyClusterOfflineConfigRequest {
	s.PushMode = &v
	return s
}

func (s *ModifyClusterOfflineConfigRequest) Validate() error {
	return dara.Validate(s)
}
