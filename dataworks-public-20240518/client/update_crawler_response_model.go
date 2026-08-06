// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCrawlerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCrawlerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCrawlerResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCrawlerResponseBody) *UpdateCrawlerResponse
	GetBody() *UpdateCrawlerResponseBody
}

type UpdateCrawlerResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCrawlerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCrawlerResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCrawlerResponse) GoString() string {
	return s.String()
}

func (s *UpdateCrawlerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCrawlerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCrawlerResponse) GetBody() *UpdateCrawlerResponseBody {
	return s.Body
}

func (s *UpdateCrawlerResponse) SetHeaders(v map[string]*string) *UpdateCrawlerResponse {
	s.Headers = v
	return s
}

func (s *UpdateCrawlerResponse) SetStatusCode(v int32) *UpdateCrawlerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCrawlerResponse) SetBody(v *UpdateCrawlerResponseBody) *UpdateCrawlerResponse {
	s.Body = v
	return s
}

func (s *UpdateCrawlerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
