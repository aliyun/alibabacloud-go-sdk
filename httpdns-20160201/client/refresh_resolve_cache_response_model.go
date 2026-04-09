// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshResolveCacheResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RefreshResolveCacheResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RefreshResolveCacheResponse
	GetStatusCode() *int32
	SetBody(v *RefreshResolveCacheResponseBody) *RefreshResolveCacheResponse
	GetBody() *RefreshResolveCacheResponseBody
}

type RefreshResolveCacheResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshResolveCacheResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshResolveCacheResponse) String() string {
	return dara.Prettify(s)
}

func (s RefreshResolveCacheResponse) GoString() string {
	return s.String()
}

func (s *RefreshResolveCacheResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RefreshResolveCacheResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RefreshResolveCacheResponse) GetBody() *RefreshResolveCacheResponseBody {
	return s.Body
}

func (s *RefreshResolveCacheResponse) SetHeaders(v map[string]*string) *RefreshResolveCacheResponse {
	s.Headers = v
	return s
}

func (s *RefreshResolveCacheResponse) SetStatusCode(v int32) *RefreshResolveCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshResolveCacheResponse) SetBody(v *RefreshResolveCacheResponseBody) *RefreshResolveCacheResponse {
	s.Body = v
	return s
}

func (s *RefreshResolveCacheResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
