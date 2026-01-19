// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImageCacheResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteImageCacheResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteImageCacheResponse
	GetStatusCode() *int32
	SetBody(v *DeleteImageCacheResponseBody) *DeleteImageCacheResponse
	GetBody() *DeleteImageCacheResponseBody
}

type DeleteImageCacheResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteImageCacheResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteImageCacheResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteImageCacheResponse) GoString() string {
	return s.String()
}

func (s *DeleteImageCacheResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteImageCacheResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteImageCacheResponse) GetBody() *DeleteImageCacheResponseBody {
	return s.Body
}

func (s *DeleteImageCacheResponse) SetHeaders(v map[string]*string) *DeleteImageCacheResponse {
	s.Headers = v
	return s
}

func (s *DeleteImageCacheResponse) SetStatusCode(v int32) *DeleteImageCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteImageCacheResponse) SetBody(v *DeleteImageCacheResponseBody) *DeleteImageCacheResponse {
	s.Body = v
	return s
}

func (s *DeleteImageCacheResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
