// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchStoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchStoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchStoriesResponse
	GetStatusCode() *int32
	SetBody(v *SearchStoriesResponseBody) *SearchStoriesResponse
	GetBody() *SearchStoriesResponseBody
}

type SearchStoriesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchStoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchStoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchStoriesResponse) GoString() string {
	return s.String()
}

func (s *SearchStoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchStoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchStoriesResponse) GetBody() *SearchStoriesResponseBody {
	return s.Body
}

func (s *SearchStoriesResponse) SetHeaders(v map[string]*string) *SearchStoriesResponse {
	s.Headers = v
	return s
}

func (s *SearchStoriesResponse) SetStatusCode(v int32) *SearchStoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchStoriesResponse) SetBody(v *SearchStoriesResponseBody) *SearchStoriesResponse {
	s.Body = v
	return s
}

func (s *SearchStoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
