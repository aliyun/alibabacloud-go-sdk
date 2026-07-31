// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFormationCrawlerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFormationCrawlerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFormationCrawlerResponse
	GetStatusCode() *int32
	SetBody(v *CreateFormationCrawlerResponseBody) *CreateFormationCrawlerResponse
	GetBody() *CreateFormationCrawlerResponseBody
}

type CreateFormationCrawlerResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFormationCrawlerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFormationCrawlerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFormationCrawlerResponse) GoString() string {
	return s.String()
}

func (s *CreateFormationCrawlerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFormationCrawlerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFormationCrawlerResponse) GetBody() *CreateFormationCrawlerResponseBody {
	return s.Body
}

func (s *CreateFormationCrawlerResponse) SetHeaders(v map[string]*string) *CreateFormationCrawlerResponse {
	s.Headers = v
	return s
}

func (s *CreateFormationCrawlerResponse) SetStatusCode(v int32) *CreateFormationCrawlerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFormationCrawlerResponse) SetBody(v *CreateFormationCrawlerResponseBody) *CreateFormationCrawlerResponse {
	s.Body = v
	return s
}

func (s *CreateFormationCrawlerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
