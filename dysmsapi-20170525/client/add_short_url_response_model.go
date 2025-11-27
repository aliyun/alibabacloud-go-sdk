// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddShortUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddShortUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddShortUrlResponse
	GetStatusCode() *int32
	SetBody(v *AddShortUrlResponseBody) *AddShortUrlResponse
	GetBody() *AddShortUrlResponseBody
}

type AddShortUrlResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddShortUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddShortUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s AddShortUrlResponse) GoString() string {
	return s.String()
}

func (s *AddShortUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddShortUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddShortUrlResponse) GetBody() *AddShortUrlResponseBody {
	return s.Body
}

func (s *AddShortUrlResponse) SetHeaders(v map[string]*string) *AddShortUrlResponse {
	s.Headers = v
	return s
}

func (s *AddShortUrlResponse) SetStatusCode(v int32) *AddShortUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *AddShortUrlResponse) SetBody(v *AddShortUrlResponseBody) *AddShortUrlResponse {
	s.Body = v
	return s
}

func (s *AddShortUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
