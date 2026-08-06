// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCrawlerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCrawlerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCrawlerResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCrawlerResponseBody) *DeleteCrawlerResponse
	GetBody() *DeleteCrawlerResponseBody
}

type DeleteCrawlerResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCrawlerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCrawlerResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCrawlerResponse) GoString() string {
	return s.String()
}

func (s *DeleteCrawlerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCrawlerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCrawlerResponse) GetBody() *DeleteCrawlerResponseBody {
	return s.Body
}

func (s *DeleteCrawlerResponse) SetHeaders(v map[string]*string) *DeleteCrawlerResponse {
	s.Headers = v
	return s
}

func (s *DeleteCrawlerResponse) SetStatusCode(v int32) *DeleteCrawlerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCrawlerResponse) SetBody(v *DeleteCrawlerResponseBody) *DeleteCrawlerResponse {
	s.Body = v
	return s
}

func (s *DeleteCrawlerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
