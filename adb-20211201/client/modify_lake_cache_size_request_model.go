// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLakeCacheSizeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCapacity(v int64) *ModifyLakeCacheSizeRequest
	GetCapacity() *int64
	SetDBClusterId(v string) *ModifyLakeCacheSizeRequest
	GetDBClusterId() *string
	SetEnableLakeCache(v bool) *ModifyLakeCacheSizeRequest
	GetEnableLakeCache() *bool
}

type ModifyLakeCacheSizeRequest struct {
	// The lake cache size that you want to set. Unit: GB.
	//
	// example:
	//
	// 100
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
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
	// Specifies whether to enable the lake cache feature.
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
	EnableLakeCache *bool `json:"EnableLakeCache,omitempty" xml:"EnableLakeCache,omitempty"`
}

func (s ModifyLakeCacheSizeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLakeCacheSizeRequest) GoString() string {
	return s.String()
}

func (s *ModifyLakeCacheSizeRequest) GetCapacity() *int64 {
	return s.Capacity
}

func (s *ModifyLakeCacheSizeRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyLakeCacheSizeRequest) GetEnableLakeCache() *bool {
	return s.EnableLakeCache
}

func (s *ModifyLakeCacheSizeRequest) SetCapacity(v int64) *ModifyLakeCacheSizeRequest {
	s.Capacity = &v
	return s
}

func (s *ModifyLakeCacheSizeRequest) SetDBClusterId(v string) *ModifyLakeCacheSizeRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyLakeCacheSizeRequest) SetEnableLakeCache(v bool) *ModifyLakeCacheSizeRequest {
	s.EnableLakeCache = &v
	return s
}

func (s *ModifyLakeCacheSizeRequest) Validate() error {
	return dara.Validate(s)
}
