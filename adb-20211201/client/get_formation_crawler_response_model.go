// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFormationCrawlerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFormationCrawlerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFormationCrawlerResponse
	GetStatusCode() *int32
	SetBody(v *GetFormationCrawlerResponseBody) *GetFormationCrawlerResponse
	GetBody() *GetFormationCrawlerResponseBody
}

type GetFormationCrawlerResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFormationCrawlerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFormationCrawlerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFormationCrawlerResponse) GoString() string {
	return s.String()
}

func (s *GetFormationCrawlerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFormationCrawlerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFormationCrawlerResponse) GetBody() *GetFormationCrawlerResponseBody {
	return s.Body
}

func (s *GetFormationCrawlerResponse) SetHeaders(v map[string]*string) *GetFormationCrawlerResponse {
	s.Headers = v
	return s
}

func (s *GetFormationCrawlerResponse) SetStatusCode(v int32) *GetFormationCrawlerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFormationCrawlerResponse) SetBody(v *GetFormationCrawlerResponseBody) *GetFormationCrawlerResponse {
	s.Body = v
	return s
}

func (s *GetFormationCrawlerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
