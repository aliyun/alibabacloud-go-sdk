// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreloadCachesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentShrink(v string) *PreloadCachesShrinkRequest
	GetContentShrink() *string
	SetHeadersShrink(v string) *PreloadCachesShrinkRequest
	GetHeadersShrink() *string
	SetSiteId(v int64) *PreloadCachesShrinkRequest
	GetSiteId() *int64
}

type PreloadCachesShrinkRequest struct {
	// The files to be prefetched.
	ContentShrink *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// By default, prefetch requests include the Accept-Encoding:gzip header. If you want a prefetch request to include other headers or implement multi-replica prefetch, you can specify a custom prefetch header by configuring the Headers parameter.
	HeadersShrink *string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	// The website ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the ID.
	//
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s PreloadCachesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PreloadCachesShrinkRequest) GoString() string {
	return s.String()
}

func (s *PreloadCachesShrinkRequest) GetContentShrink() *string {
	return s.ContentShrink
}

func (s *PreloadCachesShrinkRequest) GetHeadersShrink() *string {
	return s.HeadersShrink
}

func (s *PreloadCachesShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *PreloadCachesShrinkRequest) SetContentShrink(v string) *PreloadCachesShrinkRequest {
	s.ContentShrink = &v
	return s
}

func (s *PreloadCachesShrinkRequest) SetHeadersShrink(v string) *PreloadCachesShrinkRequest {
	s.HeadersShrink = &v
	return s
}

func (s *PreloadCachesShrinkRequest) SetSiteId(v int64) *PreloadCachesShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *PreloadCachesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
