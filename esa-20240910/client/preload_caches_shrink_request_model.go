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
	// The prefetch objects.
	ContentShrink *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The default header carried in a prefetch request is Accept-Encoding:gzip. If you want the prefetch request to carry other headers or implement multi-copy prefetching, use this parameter to specify custom prefetch headers.
	HeadersShrink *string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	// The site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
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
