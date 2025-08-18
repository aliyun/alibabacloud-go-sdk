// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreloadCachesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PreloadCachesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PreloadCachesResponse
	GetStatusCode() *int32
	SetBody(v *PreloadCachesResponseBody) *PreloadCachesResponse
	GetBody() *PreloadCachesResponseBody
}

type PreloadCachesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreloadCachesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreloadCachesResponse) String() string {
	return dara.Prettify(s)
}

func (s PreloadCachesResponse) GoString() string {
	return s.String()
}

func (s *PreloadCachesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PreloadCachesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PreloadCachesResponse) GetBody() *PreloadCachesResponseBody {
	return s.Body
}

func (s *PreloadCachesResponse) SetHeaders(v map[string]*string) *PreloadCachesResponse {
	s.Headers = v
	return s
}

func (s *PreloadCachesResponse) SetStatusCode(v int32) *PreloadCachesResponse {
	s.StatusCode = &v
	return s
}

func (s *PreloadCachesResponse) SetBody(v *PreloadCachesResponseBody) *PreloadCachesResponse {
	s.Body = v
	return s
}

func (s *PreloadCachesResponse) Validate() error {
	return dara.Validate(s)
}
