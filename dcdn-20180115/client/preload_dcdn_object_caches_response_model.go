// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreloadDcdnObjectCachesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PreloadDcdnObjectCachesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PreloadDcdnObjectCachesResponse
	GetStatusCode() *int32
	SetBody(v *PreloadDcdnObjectCachesResponseBody) *PreloadDcdnObjectCachesResponse
	GetBody() *PreloadDcdnObjectCachesResponseBody
}

type PreloadDcdnObjectCachesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreloadDcdnObjectCachesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreloadDcdnObjectCachesResponse) String() string {
	return dara.Prettify(s)
}

func (s PreloadDcdnObjectCachesResponse) GoString() string {
	return s.String()
}

func (s *PreloadDcdnObjectCachesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PreloadDcdnObjectCachesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PreloadDcdnObjectCachesResponse) GetBody() *PreloadDcdnObjectCachesResponseBody {
	return s.Body
}

func (s *PreloadDcdnObjectCachesResponse) SetHeaders(v map[string]*string) *PreloadDcdnObjectCachesResponse {
	s.Headers = v
	return s
}

func (s *PreloadDcdnObjectCachesResponse) SetStatusCode(v int32) *PreloadDcdnObjectCachesResponse {
	s.StatusCode = &v
	return s
}

func (s *PreloadDcdnObjectCachesResponse) SetBody(v *PreloadDcdnObjectCachesResponseBody) *PreloadDcdnObjectCachesResponse {
	s.Body = v
	return s
}

func (s *PreloadDcdnObjectCachesResponse) Validate() error {
	return dara.Validate(s)
}
