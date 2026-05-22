// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTieredCacheRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *GetTieredCacheRequest
	GetSiteId() *int64
}

type GetTieredCacheRequest struct {
	// Site ID, which can be obtained by calling the [ListSites](~~ListSites~~) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetTieredCacheRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTieredCacheRequest) GoString() string {
	return s.String()
}

func (s *GetTieredCacheRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetTieredCacheRequest) SetSiteId(v int64) *GetTieredCacheRequest {
	s.SiteId = &v
	return s
}

func (s *GetTieredCacheRequest) Validate() error {
	return dara.Validate(s)
}
