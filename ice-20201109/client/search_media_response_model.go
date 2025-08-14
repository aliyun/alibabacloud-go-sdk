// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMediaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchMediaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchMediaResponse
	GetStatusCode() *int32
	SetBody(v *SearchMediaResponseBody) *SearchMediaResponse
	GetBody() *SearchMediaResponseBody
}

type SearchMediaResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchMediaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchMediaResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaResponse) GoString() string {
	return s.String()
}

func (s *SearchMediaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchMediaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchMediaResponse) GetBody() *SearchMediaResponseBody {
	return s.Body
}

func (s *SearchMediaResponse) SetHeaders(v map[string]*string) *SearchMediaResponse {
	s.Headers = v
	return s
}

func (s *SearchMediaResponse) SetStatusCode(v int32) *SearchMediaResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchMediaResponse) SetBody(v *SearchMediaResponseBody) *SearchMediaResponse {
	s.Body = v
	return s
}

func (s *SearchMediaResponse) Validate() error {
	return dara.Validate(s)
}
