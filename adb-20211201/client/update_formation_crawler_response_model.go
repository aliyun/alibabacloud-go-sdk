// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFormationCrawlerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFormationCrawlerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFormationCrawlerResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFormationCrawlerResponseBody) *UpdateFormationCrawlerResponse
	GetBody() *UpdateFormationCrawlerResponseBody
}

type UpdateFormationCrawlerResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFormationCrawlerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFormationCrawlerResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFormationCrawlerResponse) GoString() string {
	return s.String()
}

func (s *UpdateFormationCrawlerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFormationCrawlerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFormationCrawlerResponse) GetBody() *UpdateFormationCrawlerResponseBody {
	return s.Body
}

func (s *UpdateFormationCrawlerResponse) SetHeaders(v map[string]*string) *UpdateFormationCrawlerResponse {
	s.Headers = v
	return s
}

func (s *UpdateFormationCrawlerResponse) SetStatusCode(v int32) *UpdateFormationCrawlerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFormationCrawlerResponse) SetBody(v *UpdateFormationCrawlerResponseBody) *UpdateFormationCrawlerResponse {
	s.Body = v
	return s
}

func (s *UpdateFormationCrawlerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
