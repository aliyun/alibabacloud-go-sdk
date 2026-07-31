// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartFormationCrawlerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartFormationCrawlerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartFormationCrawlerResponse
	GetStatusCode() *int32
	SetBody(v *StartFormationCrawlerResponseBody) *StartFormationCrawlerResponse
	GetBody() *StartFormationCrawlerResponseBody
}

type StartFormationCrawlerResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartFormationCrawlerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartFormationCrawlerResponse) String() string {
	return dara.Prettify(s)
}

func (s StartFormationCrawlerResponse) GoString() string {
	return s.String()
}

func (s *StartFormationCrawlerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartFormationCrawlerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartFormationCrawlerResponse) GetBody() *StartFormationCrawlerResponseBody {
	return s.Body
}

func (s *StartFormationCrawlerResponse) SetHeaders(v map[string]*string) *StartFormationCrawlerResponse {
	s.Headers = v
	return s
}

func (s *StartFormationCrawlerResponse) SetStatusCode(v int32) *StartFormationCrawlerResponse {
	s.StatusCode = &v
	return s
}

func (s *StartFormationCrawlerResponse) SetBody(v *StartFormationCrawlerResponseBody) *StartFormationCrawlerResponse {
	s.Body = v
	return s
}

func (s *StartFormationCrawlerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
