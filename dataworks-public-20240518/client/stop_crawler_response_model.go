// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCrawlerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopCrawlerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopCrawlerResponse
	GetStatusCode() *int32
	SetBody(v *StopCrawlerResponseBody) *StopCrawlerResponse
	GetBody() *StopCrawlerResponseBody
}

type StopCrawlerResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopCrawlerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopCrawlerResponse) String() string {
	return dara.Prettify(s)
}

func (s StopCrawlerResponse) GoString() string {
	return s.String()
}

func (s *StopCrawlerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopCrawlerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopCrawlerResponse) GetBody() *StopCrawlerResponseBody {
	return s.Body
}

func (s *StopCrawlerResponse) SetHeaders(v map[string]*string) *StopCrawlerResponse {
	s.Headers = v
	return s
}

func (s *StopCrawlerResponse) SetStatusCode(v int32) *StopCrawlerResponse {
	s.StatusCode = &v
	return s
}

func (s *StopCrawlerResponse) SetBody(v *StopCrawlerResponseBody) *StopCrawlerResponse {
	s.Body = v
	return s
}

func (s *StopCrawlerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
