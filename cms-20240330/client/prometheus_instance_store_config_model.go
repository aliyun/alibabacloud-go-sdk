// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrometheusInstanceStoreConfig interface {
	dara.Model
	String() string
	GoString() string
	SetParallelCountPerHost(v int32) *PrometheusInstanceStoreConfig
	GetParallelCountPerHost() *int32
	SetParallelEnable(v bool) *PrometheusInstanceStoreConfig
	GetParallelEnable() *bool
	SetParallelMode(v string) *PrometheusInstanceStoreConfig
	GetParallelMode() *string
	SetQueryCacheEnable(v bool) *PrometheusInstanceStoreConfig
	GetQueryCacheEnable() *bool
	SetTotalParallelCount(v int32) *PrometheusInstanceStoreConfig
	GetTotalParallelCount() *int32
}

type PrometheusInstanceStoreConfig struct {
	// The concurrency per host. If this parameter is not specified, the default value is 2. Valid values: 1 to 8.
	//
	// example:
	//
	// 2
	ParallelCountPerHost *int32 `json:"parallelCountPerHost,omitempty" xml:"parallelCountPerHost,omitempty"`
	// Specifies whether to enable parallel query. If this parameter is not specified, the value is considered as false.
	ParallelEnable *bool `json:"parallelEnable,omitempty" xml:"parallelEnable,omitempty"`
	// The parallel query mode. Valid values:
	//
	// - auto
	//
	// - static
	//
	// If this parameter is not specified, the default value is auto.
	//
	// example:
	//
	// static
	ParallelMode *string `json:"parallelMode,omitempty" xml:"parallelMode,omitempty"`
	// Specifies whether to enable query cache. If this parameter is not specified, the value is considered as false.
	QueryCacheEnable *bool `json:"queryCacheEnable,omitempty" xml:"queryCacheEnable,omitempty"`
	// The global concurrency. If this parameter is not specified, the default value is 8. Valid values: 2 to 64.
	//
	// example:
	//
	// 8
	TotalParallelCount *int32 `json:"totalParallelCount,omitempty" xml:"totalParallelCount,omitempty"`
}

func (s PrometheusInstanceStoreConfig) String() string {
	return dara.Prettify(s)
}

func (s PrometheusInstanceStoreConfig) GoString() string {
	return s.String()
}

func (s *PrometheusInstanceStoreConfig) GetParallelCountPerHost() *int32 {
	return s.ParallelCountPerHost
}

func (s *PrometheusInstanceStoreConfig) GetParallelEnable() *bool {
	return s.ParallelEnable
}

func (s *PrometheusInstanceStoreConfig) GetParallelMode() *string {
	return s.ParallelMode
}

func (s *PrometheusInstanceStoreConfig) GetQueryCacheEnable() *bool {
	return s.QueryCacheEnable
}

func (s *PrometheusInstanceStoreConfig) GetTotalParallelCount() *int32 {
	return s.TotalParallelCount
}

func (s *PrometheusInstanceStoreConfig) SetParallelCountPerHost(v int32) *PrometheusInstanceStoreConfig {
	s.ParallelCountPerHost = &v
	return s
}

func (s *PrometheusInstanceStoreConfig) SetParallelEnable(v bool) *PrometheusInstanceStoreConfig {
	s.ParallelEnable = &v
	return s
}

func (s *PrometheusInstanceStoreConfig) SetParallelMode(v string) *PrometheusInstanceStoreConfig {
	s.ParallelMode = &v
	return s
}

func (s *PrometheusInstanceStoreConfig) SetQueryCacheEnable(v bool) *PrometheusInstanceStoreConfig {
	s.QueryCacheEnable = &v
	return s
}

func (s *PrometheusInstanceStoreConfig) SetTotalParallelCount(v int32) *PrometheusInstanceStoreConfig {
	s.TotalParallelCount = &v
	return s
}

func (s *PrometheusInstanceStoreConfig) Validate() error {
	return dara.Validate(s)
}
