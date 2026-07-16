// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCacheReserveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCacheReserveInstanceId(v string) *UpdateCacheReserveRequest
	GetCacheReserveInstanceId() *string
	SetEnable(v string) *UpdateCacheReserveRequest
	GetEnable() *string
	SetSiteId(v int64) *UpdateCacheReserveRequest
	GetSiteId() *int64
}

type UpdateCacheReserveRequest struct {
	// The cache reserve instance ID.
	//
	// example:
	//
	// cr_hk_123456789
	CacheReserveInstanceId *string `json:"CacheReserveInstanceId,omitempty" xml:"CacheReserveInstanceId,omitempty"`
	// The switch. Valid values:
	//
	// - **on**: enabled.
	//
	// - **off**: disabled.
	//
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The site ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateCacheReserveRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCacheReserveRequest) GoString() string {
	return s.String()
}

func (s *UpdateCacheReserveRequest) GetCacheReserveInstanceId() *string {
	return s.CacheReserveInstanceId
}

func (s *UpdateCacheReserveRequest) GetEnable() *string {
	return s.Enable
}

func (s *UpdateCacheReserveRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateCacheReserveRequest) SetCacheReserveInstanceId(v string) *UpdateCacheReserveRequest {
	s.CacheReserveInstanceId = &v
	return s
}

func (s *UpdateCacheReserveRequest) SetEnable(v string) *UpdateCacheReserveRequest {
	s.Enable = &v
	return s
}

func (s *UpdateCacheReserveRequest) SetSiteId(v int64) *UpdateCacheReserveRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateCacheReserveRequest) Validate() error {
	return dara.Validate(s)
}
