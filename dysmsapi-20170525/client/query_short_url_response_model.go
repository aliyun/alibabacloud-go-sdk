// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryShortUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryShortUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryShortUrlResponse
	GetStatusCode() *int32
	SetBody(v *QueryShortUrlResponseBody) *QueryShortUrlResponse
	GetBody() *QueryShortUrlResponseBody
}

type QueryShortUrlResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryShortUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryShortUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryShortUrlResponse) GoString() string {
	return s.String()
}

func (s *QueryShortUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryShortUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryShortUrlResponse) GetBody() *QueryShortUrlResponseBody {
	return s.Body
}

func (s *QueryShortUrlResponse) SetHeaders(v map[string]*string) *QueryShortUrlResponse {
	s.Headers = v
	return s
}

func (s *QueryShortUrlResponse) SetStatusCode(v int32) *QueryShortUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryShortUrlResponse) SetBody(v *QueryShortUrlResponseBody) *QueryShortUrlResponse {
	s.Body = v
	return s
}

func (s *QueryShortUrlResponse) Validate() error {
	return dara.Validate(s)
}
