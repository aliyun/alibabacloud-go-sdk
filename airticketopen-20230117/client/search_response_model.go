// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchResponse
	GetStatusCode() *int32
	SetBody(v *SearchResponseBody) *SearchResponse
	GetBody() *SearchResponseBody
}

type SearchResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchResponse) GoString() string {
	return s.String()
}

func (s *SearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchResponse) GetBody() *SearchResponseBody {
	return s.Body
}

func (s *SearchResponse) SetHeaders(v map[string]*string) *SearchResponse {
	s.Headers = v
	return s
}

func (s *SearchResponse) SetStatusCode(v int32) *SearchResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchResponse) SetBody(v *SearchResponseBody) *SearchResponse {
	s.Body = v
	return s
}

func (s *SearchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
