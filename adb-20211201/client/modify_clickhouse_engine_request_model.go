// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClickhouseEngineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCacheSize(v int32) *ModifyClickhouseEngineRequest
	GetCacheSize() *int32
	SetDBClusterId(v string) *ModifyClickhouseEngineRequest
	GetDBClusterId() *string
	SetEnabled(v bool) *ModifyClickhouseEngineRequest
	GetEnabled() *bool
	SetOwnerId(v string) *ModifyClickhouseEngineRequest
	GetOwnerId() *string
}

type ModifyClickhouseEngineRequest struct {
	// The disk cache size of the wide table engine. Unit: GB. Default value: 100. Valid values: 100 to 1000.
	//
	// example:
	//
	// 200
	CacheSize *int32 `json:"CacheSize,omitempty" xml:"CacheSize,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to enable the wide table engine feature. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// true
	Enabled *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s ModifyClickhouseEngineRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyClickhouseEngineRequest) GoString() string {
	return s.String()
}

func (s *ModifyClickhouseEngineRequest) GetCacheSize() *int32 {
	return s.CacheSize
}

func (s *ModifyClickhouseEngineRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyClickhouseEngineRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *ModifyClickhouseEngineRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ModifyClickhouseEngineRequest) SetCacheSize(v int32) *ModifyClickhouseEngineRequest {
	s.CacheSize = &v
	return s
}

func (s *ModifyClickhouseEngineRequest) SetDBClusterId(v string) *ModifyClickhouseEngineRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyClickhouseEngineRequest) SetEnabled(v bool) *ModifyClickhouseEngineRequest {
	s.Enabled = &v
	return s
}

func (s *ModifyClickhouseEngineRequest) SetOwnerId(v string) *ModifyClickhouseEngineRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyClickhouseEngineRequest) Validate() error {
	return dara.Validate(s)
}
