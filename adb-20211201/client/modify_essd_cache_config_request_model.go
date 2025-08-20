// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEssdCacheConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyEssdCacheConfigRequest
	GetDBClusterId() *string
	SetEnableEssdCache(v bool) *ModifyEssdCacheConfigRequest
	GetEnableEssdCache() *bool
	SetEssdCacheSize(v int32) *ModifyEssdCacheConfigRequest
	GetEssdCacheSize() *int32
}

type ModifyEssdCacheConfigRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp10yt0gva71ei7d
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to enable the disk cache feature.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableEssdCache *bool `json:"EnableEssdCache,omitempty" xml:"EnableEssdCache,omitempty"`
	// The disk cache size. Unit: GB.
	//
	// example:
	//
	// 500
	EssdCacheSize *int32 `json:"EssdCacheSize,omitempty" xml:"EssdCacheSize,omitempty"`
}

func (s ModifyEssdCacheConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyEssdCacheConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyEssdCacheConfigRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyEssdCacheConfigRequest) GetEnableEssdCache() *bool {
	return s.EnableEssdCache
}

func (s *ModifyEssdCacheConfigRequest) GetEssdCacheSize() *int32 {
	return s.EssdCacheSize
}

func (s *ModifyEssdCacheConfigRequest) SetDBClusterId(v string) *ModifyEssdCacheConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyEssdCacheConfigRequest) SetEnableEssdCache(v bool) *ModifyEssdCacheConfigRequest {
	s.EnableEssdCache = &v
	return s
}

func (s *ModifyEssdCacheConfigRequest) SetEssdCacheSize(v int32) *ModifyEssdCacheConfigRequest {
	s.EssdCacheSize = &v
	return s
}

func (s *ModifyEssdCacheConfigRequest) Validate() error {
	return dara.Validate(s)
}
