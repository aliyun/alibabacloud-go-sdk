// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchImageResponse
	GetStatusCode() *int32
	SetBody(v *SearchImageResponseBody) *SearchImageResponse
	GetBody() *SearchImageResponseBody
}

type SearchImageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchImageResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchImageResponse) GoString() string {
	return s.String()
}

func (s *SearchImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchImageResponse) GetBody() *SearchImageResponseBody {
	return s.Body
}

func (s *SearchImageResponse) SetHeaders(v map[string]*string) *SearchImageResponse {
	s.Headers = v
	return s
}

func (s *SearchImageResponse) SetStatusCode(v int32) *SearchImageResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchImageResponse) SetBody(v *SearchImageResponseBody) *SearchImageResponse {
	s.Body = v
	return s
}

func (s *SearchImageResponse) Validate() error {
	return dara.Validate(s)
}
