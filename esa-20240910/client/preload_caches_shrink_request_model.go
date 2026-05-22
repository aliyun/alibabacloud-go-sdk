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
	ContentShrink *string `json:"Content,omitempty" xml:"Content,omitempty"`
	HeadersShrink *string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	SiteId        *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
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
