// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopFormationCrawlerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopFormationCrawlerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopFormationCrawlerResponse
	GetStatusCode() *int32
	SetBody(v *StopFormationCrawlerResponseBody) *StopFormationCrawlerResponse
	GetBody() *StopFormationCrawlerResponseBody
}

type StopFormationCrawlerResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopFormationCrawlerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopFormationCrawlerResponse) String() string {
	return dara.Prettify(s)
}

func (s StopFormationCrawlerResponse) GoString() string {
	return s.String()
}

func (s *StopFormationCrawlerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopFormationCrawlerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopFormationCrawlerResponse) GetBody() *StopFormationCrawlerResponseBody {
	return s.Body
}

func (s *StopFormationCrawlerResponse) SetHeaders(v map[string]*string) *StopFormationCrawlerResponse {
	s.Headers = v
	return s
}

func (s *StopFormationCrawlerResponse) SetStatusCode(v int32) *StopFormationCrawlerResponse {
	s.StatusCode = &v
	return s
}

func (s *StopFormationCrawlerResponse) SetBody(v *StopFormationCrawlerResponseBody) *StopFormationCrawlerResponse {
	s.Body = v
	return s
}

func (s *StopFormationCrawlerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
