// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMemoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchMemoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchMemoriesResponse
	GetStatusCode() *int32
	SetBody(v *SearchMemoriesResponseBody) *SearchMemoriesResponse
	GetBody() *SearchMemoriesResponseBody
}

type SearchMemoriesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchMemoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchMemoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchMemoriesResponse) GoString() string {
	return s.String()
}

func (s *SearchMemoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchMemoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchMemoriesResponse) GetBody() *SearchMemoriesResponseBody {
	return s.Body
}

func (s *SearchMemoriesResponse) SetHeaders(v map[string]*string) *SearchMemoriesResponse {
	s.Headers = v
	return s
}

func (s *SearchMemoriesResponse) SetStatusCode(v int32) *SearchMemoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchMemoriesResponse) SetBody(v *SearchMemoriesResponseBody) *SearchMemoriesResponse {
	s.Body = v
	return s
}

func (s *SearchMemoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
