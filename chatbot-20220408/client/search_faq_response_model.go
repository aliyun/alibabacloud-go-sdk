// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchFaqResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchFaqResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchFaqResponse
	GetStatusCode() *int32
	SetBody(v *SearchFaqResponseBody) *SearchFaqResponse
	GetBody() *SearchFaqResponseBody
}

type SearchFaqResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchFaqResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchFaqResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchFaqResponse) GoString() string {
	return s.String()
}

func (s *SearchFaqResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchFaqResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchFaqResponse) GetBody() *SearchFaqResponseBody {
	return s.Body
}

func (s *SearchFaqResponse) SetHeaders(v map[string]*string) *SearchFaqResponse {
	s.Headers = v
	return s
}

func (s *SearchFaqResponse) SetStatusCode(v int32) *SearchFaqResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchFaqResponse) SetBody(v *SearchFaqResponseBody) *SearchFaqResponse {
	s.Body = v
	return s
}

func (s *SearchFaqResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
