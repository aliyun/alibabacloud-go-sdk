// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTieredCacheRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCacheArchitectureMode(v string) *UpdateTieredCacheRequest
	GetCacheArchitectureMode() *string
	SetSiteId(v int64) *UpdateTieredCacheRequest
	GetSiteId() *int64
}

type UpdateTieredCacheRequest struct {
	// The tiered cache architecture mode. Valid values:
	//
	// 	- edge: edge tiered cache.
	//
	// 	- edge_smart: edge tiered cache + smart tiered cache.
	//
	// 	- edge_regional: edge tiered cache + regional tiered cache.
	//
	// 	- edge_regional_smart: edge tiered cache + regional tiered cache + smart tiered cache.
	//
	// This parameter is required.
	//
	// example:
	//
	// edge_smart
	CacheArchitectureMode *string `json:"CacheArchitectureMode,omitempty" xml:"CacheArchitectureMode,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateTieredCacheRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTieredCacheRequest) GoString() string {
	return s.String()
}

func (s *UpdateTieredCacheRequest) GetCacheArchitectureMode() *string {
	return s.CacheArchitectureMode
}

func (s *UpdateTieredCacheRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateTieredCacheRequest) SetCacheArchitectureMode(v string) *UpdateTieredCacheRequest {
	s.CacheArchitectureMode = &v
	return s
}

func (s *UpdateTieredCacheRequest) SetSiteId(v int64) *UpdateTieredCacheRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateTieredCacheRequest) Validate() error {
	return dara.Validate(s)
}
