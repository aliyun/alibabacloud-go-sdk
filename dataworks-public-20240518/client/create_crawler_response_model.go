// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCrawlerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCrawlerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCrawlerResponse
	GetStatusCode() *int32
	SetBody(v *CreateCrawlerResponseBody) *CreateCrawlerResponse
	GetBody() *CreateCrawlerResponseBody
}

type CreateCrawlerResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCrawlerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCrawlerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCrawlerResponse) GoString() string {
	return s.String()
}

func (s *CreateCrawlerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCrawlerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCrawlerResponse) GetBody() *CreateCrawlerResponseBody {
	return s.Body
}

func (s *CreateCrawlerResponse) SetHeaders(v map[string]*string) *CreateCrawlerResponse {
	s.Headers = v
	return s
}

func (s *CreateCrawlerResponse) SetStatusCode(v int32) *CreateCrawlerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCrawlerResponse) SetBody(v *CreateCrawlerResponseBody) *CreateCrawlerResponse {
	s.Body = v
	return s
}

func (s *CreateCrawlerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
