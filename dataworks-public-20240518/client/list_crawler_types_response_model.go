// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCrawlerTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCrawlerTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCrawlerTypesResponse
	GetStatusCode() *int32
	SetBody(v *ListCrawlerTypesResponseBody) *ListCrawlerTypesResponse
	GetBody() *ListCrawlerTypesResponseBody
}

type ListCrawlerTypesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCrawlerTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCrawlerTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCrawlerTypesResponse) GoString() string {
	return s.String()
}

func (s *ListCrawlerTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCrawlerTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCrawlerTypesResponse) GetBody() *ListCrawlerTypesResponseBody {
	return s.Body
}

func (s *ListCrawlerTypesResponse) SetHeaders(v map[string]*string) *ListCrawlerTypesResponse {
	s.Headers = v
	return s
}

func (s *ListCrawlerTypesResponse) SetStatusCode(v int32) *ListCrawlerTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCrawlerTypesResponse) SetBody(v *ListCrawlerTypesResponseBody) *ListCrawlerTypesResponse {
	s.Body = v
	return s
}

func (s *ListCrawlerTypesResponse) Validate() error {
	return dara.Validate(s)
}
