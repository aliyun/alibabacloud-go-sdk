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
	// Site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// Version number of the site.
	//
	// example:
	//
	// 1
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
