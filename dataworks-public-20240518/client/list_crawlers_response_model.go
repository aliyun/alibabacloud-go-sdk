// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCrawlersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCrawlersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCrawlersResponse
	GetStatusCode() *int32
	SetBody(v *ListCrawlersResponseBody) *ListCrawlersResponse
	GetBody() *ListCrawlersResponseBody
}

type ListCrawlersResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCrawlersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCrawlersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCrawlersResponse) GoString() string {
	return s.String()
}

func (s *ListCrawlersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCrawlersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCrawlersResponse) GetBody() *ListCrawlersResponseBody {
	return s.Body
}

func (s *ListCrawlersResponse) SetHeaders(v map[string]*string) *ListCrawlersResponse {
	s.Headers = v
	return s
}

func (s *ListCrawlersResponse) SetStatusCode(v int32) *ListCrawlersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCrawlersResponse) SetBody(v *ListCrawlersResponseBody) *ListCrawlersResponse {
	s.Body = v
	return s
}

func (s *ListCrawlersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
