// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteShortUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteShortUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteShortUrlResponse
	GetStatusCode() *int32
	SetBody(v *DeleteShortUrlResponseBody) *DeleteShortUrlResponse
	GetBody() *DeleteShortUrlResponseBody
}

type DeleteShortUrlResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteShortUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteShortUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteShortUrlResponse) GoString() string {
	return s.String()
}

func (s *DeleteShortUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteShortUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteShortUrlResponse) GetBody() *DeleteShortUrlResponseBody {
	return s.Body
}

func (s *DeleteShortUrlResponse) SetHeaders(v map[string]*string) *DeleteShortUrlResponse {
	s.Headers = v
	return s
}

func (s *DeleteShortUrlResponse) SetStatusCode(v int32) *DeleteShortUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteShortUrlResponse) SetBody(v *DeleteShortUrlResponseBody) *DeleteShortUrlResponse {
	s.Body = v
	return s
}

func (s *DeleteShortUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
