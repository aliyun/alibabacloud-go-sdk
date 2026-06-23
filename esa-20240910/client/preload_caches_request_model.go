// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreloadCachesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v []*string) *PreloadCachesRequest
	GetContent() []*string
	SetHeaders(v map[string]*string) *PreloadCachesRequest
	GetHeaders() map[string]*string
	SetSiteId(v int64) *PreloadCachesRequest
	GetSiteId() *int64
}

type PreloadCachesRequest struct {
	// The prefetch objects.
	Content []*string `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// The default header carried in a prefetch request is Accept-Encoding:gzip. If you want the prefetch request to carry other headers or implement multi-copy prefetching, use this parameter to specify custom prefetch headers.
	Headers map[string]*string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	// The site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s PreloadCachesRequest) String() string {
	return dara.Prettify(s)
}

func (s PreloadCachesRequest) GoString() string {
	return s.String()
}

func (s *PreloadCachesRequest) GetContent() []*string {
	return s.Content
}

func (s *PreloadCachesRequest) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PreloadCachesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *PreloadCachesRequest) SetContent(v []*string) *PreloadCachesRequest {
	s.Content = v
	return s
}

func (s *PreloadCachesRequest) SetHeaders(v map[string]*string) *PreloadCachesRequest {
	s.Headers = v
	return s
}

func (s *PreloadCachesRequest) SetSiteId(v int64) *PreloadCachesRequest {
	s.SiteId = &v
	return s
}

func (s *PreloadCachesRequest) Validate() error {
	return dara.Validate(s)
}
