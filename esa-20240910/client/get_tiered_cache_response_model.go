// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTieredCacheResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTieredCacheResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTieredCacheResponse
	GetStatusCode() *int32
	SetBody(v *GetTieredCacheResponseBody) *GetTieredCacheResponse
	GetBody() *GetTieredCacheResponseBody
}

type GetTieredCacheResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTieredCacheResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTieredCacheResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTieredCacheResponse) GoString() string {
	return s.String()
}

func (s *GetTieredCacheResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTieredCacheResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTieredCacheResponse) GetBody() *GetTieredCacheResponseBody {
	return s.Body
}

func (s *GetTieredCacheResponse) SetHeaders(v map[string]*string) *GetTieredCacheResponse {
	s.Headers = v
	return s
}

func (s *GetTieredCacheResponse) SetStatusCode(v int32) *GetTieredCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTieredCacheResponse) SetBody(v *GetTieredCacheResponseBody) *GetTieredCacheResponse {
	s.Body = v
	return s
}

func (s *GetTieredCacheResponse) Validate() error {
	return dara.Validate(s)
}
