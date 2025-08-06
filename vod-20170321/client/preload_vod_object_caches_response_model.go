// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreloadVodObjectCachesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PreloadVodObjectCachesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PreloadVodObjectCachesResponse
	GetStatusCode() *int32
	SetBody(v *PreloadVodObjectCachesResponseBody) *PreloadVodObjectCachesResponse
	GetBody() *PreloadVodObjectCachesResponseBody
}

type PreloadVodObjectCachesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreloadVodObjectCachesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreloadVodObjectCachesResponse) String() string {
	return dara.Prettify(s)
}

func (s PreloadVodObjectCachesResponse) GoString() string {
	return s.String()
}

func (s *PreloadVodObjectCachesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PreloadVodObjectCachesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PreloadVodObjectCachesResponse) GetBody() *PreloadVodObjectCachesResponseBody {
	return s.Body
}

func (s *PreloadVodObjectCachesResponse) SetHeaders(v map[string]*string) *PreloadVodObjectCachesResponse {
	s.Headers = v
	return s
}

func (s *PreloadVodObjectCachesResponse) SetStatusCode(v int32) *PreloadVodObjectCachesResponse {
	s.StatusCode = &v
	return s
}

func (s *PreloadVodObjectCachesResponse) SetBody(v *PreloadVodObjectCachesResponseBody) *PreloadVodObjectCachesResponse {
	s.Body = v
	return s
}

func (s *PreloadVodObjectCachesResponse) Validate() error {
	return dara.Validate(s)
}
