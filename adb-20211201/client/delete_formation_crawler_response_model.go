// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFormationCrawlerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFormationCrawlerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFormationCrawlerResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFormationCrawlerResponseBody) *DeleteFormationCrawlerResponse
	GetBody() *DeleteFormationCrawlerResponseBody
}

type DeleteFormationCrawlerResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFormationCrawlerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFormationCrawlerResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFormationCrawlerResponse) GoString() string {
	return s.String()
}

func (s *DeleteFormationCrawlerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFormationCrawlerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFormationCrawlerResponse) GetBody() *DeleteFormationCrawlerResponseBody {
	return s.Body
}

func (s *DeleteFormationCrawlerResponse) SetHeaders(v map[string]*string) *DeleteFormationCrawlerResponse {
	s.Headers = v
	return s
}

func (s *DeleteFormationCrawlerResponse) SetStatusCode(v int32) *DeleteFormationCrawlerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFormationCrawlerResponse) SetBody(v *DeleteFormationCrawlerResponseBody) *DeleteFormationCrawlerResponse {
	s.Body = v
	return s
}

func (s *DeleteFormationCrawlerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
