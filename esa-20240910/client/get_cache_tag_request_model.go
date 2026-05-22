// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCacheTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *GetCacheTagRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *GetCacheTagRequest
	GetSiteVersion() *int32
}

type GetCacheTagRequest struct {
	// This parameter is required.
	SiteId      *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s GetCacheTagRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCacheTagRequest) GoString() string {
	return s.String()
}

func (s *GetCacheTagRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetCacheTagRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *GetCacheTagRequest) SetSiteId(v int64) *GetCacheTagRequest {
	s.SiteId = &v
	return s
}

func (s *GetCacheTagRequest) SetSiteVersion(v int32) *GetCacheTagRequest {
	s.SiteVersion = &v
	return s
}

func (s *GetCacheTagRequest) Validate() error {
	return dara.Validate(s)
}
