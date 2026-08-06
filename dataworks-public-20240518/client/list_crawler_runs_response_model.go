// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCrawlerRunsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCrawlerRunsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCrawlerRunsResponse
	GetStatusCode() *int32
	SetBody(v *ListCrawlerRunsResponseBody) *ListCrawlerRunsResponse
	GetBody() *ListCrawlerRunsResponseBody
}

type ListCrawlerRunsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCrawlerRunsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCrawlerRunsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCrawlerRunsResponse) GoString() string {
	return s.String()
}

func (s *ListCrawlerRunsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCrawlerRunsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCrawlerRunsResponse) GetBody() *ListCrawlerRunsResponseBody {
	return s.Body
}

func (s *ListCrawlerRunsResponse) SetHeaders(v map[string]*string) *ListCrawlerRunsResponse {
	s.Headers = v
	return s
}

func (s *ListCrawlerRunsResponse) SetStatusCode(v int32) *ListCrawlerRunsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCrawlerRunsResponse) SetBody(v *ListCrawlerRunsResponseBody) *ListCrawlerRunsResponse {
	s.Body = v
	return s
}

func (s *ListCrawlerRunsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
