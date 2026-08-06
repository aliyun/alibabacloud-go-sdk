// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCrawlerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCrawlerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCrawlerResponse
	GetStatusCode() *int32
	SetBody(v *GetCrawlerResponseBody) *GetCrawlerResponse
	GetBody() *GetCrawlerResponseBody
}

type GetCrawlerResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCrawlerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCrawlerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCrawlerResponse) GoString() string {
	return s.String()
}

func (s *GetCrawlerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCrawlerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCrawlerResponse) GetBody() *GetCrawlerResponseBody {
	return s.Body
}

func (s *GetCrawlerResponse) SetHeaders(v map[string]*string) *GetCrawlerResponse {
	s.Headers = v
	return s
}

func (s *GetCrawlerResponse) SetStatusCode(v int32) *GetCrawlerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCrawlerResponse) SetBody(v *GetCrawlerResponseBody) *GetCrawlerResponse {
	s.Body = v
	return s
}

func (s *GetCrawlerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
