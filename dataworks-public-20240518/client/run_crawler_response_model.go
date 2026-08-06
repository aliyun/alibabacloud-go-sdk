// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCrawlerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunCrawlerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunCrawlerResponse
	GetStatusCode() *int32
	SetBody(v *RunCrawlerResponseBody) *RunCrawlerResponse
	GetBody() *RunCrawlerResponseBody
}

type RunCrawlerResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunCrawlerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunCrawlerResponse) String() string {
	return dara.Prettify(s)
}

func (s RunCrawlerResponse) GoString() string {
	return s.String()
}

func (s *RunCrawlerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunCrawlerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunCrawlerResponse) GetBody() *RunCrawlerResponseBody {
	return s.Body
}

func (s *RunCrawlerResponse) SetHeaders(v map[string]*string) *RunCrawlerResponse {
	s.Headers = v
	return s
}

func (s *RunCrawlerResponse) SetStatusCode(v int32) *RunCrawlerResponse {
	s.StatusCode = &v
	return s
}

func (s *RunCrawlerResponse) SetBody(v *RunCrawlerResponseBody) *RunCrawlerResponse {
	s.Body = v
	return s
}

func (s *RunCrawlerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
