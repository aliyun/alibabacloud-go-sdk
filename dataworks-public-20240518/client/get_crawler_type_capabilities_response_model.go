// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCrawlerTypeCapabilitiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCrawlerTypeCapabilitiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCrawlerTypeCapabilitiesResponse
	GetStatusCode() *int32
	SetBody(v *GetCrawlerTypeCapabilitiesResponseBody) *GetCrawlerTypeCapabilitiesResponse
	GetBody() *GetCrawlerTypeCapabilitiesResponseBody
}

type GetCrawlerTypeCapabilitiesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCrawlerTypeCapabilitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCrawlerTypeCapabilitiesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCrawlerTypeCapabilitiesResponse) GoString() string {
	return s.String()
}

func (s *GetCrawlerTypeCapabilitiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCrawlerTypeCapabilitiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCrawlerTypeCapabilitiesResponse) GetBody() *GetCrawlerTypeCapabilitiesResponseBody {
	return s.Body
}

func (s *GetCrawlerTypeCapabilitiesResponse) SetHeaders(v map[string]*string) *GetCrawlerTypeCapabilitiesResponse {
	s.Headers = v
	return s
}

func (s *GetCrawlerTypeCapabilitiesResponse) SetStatusCode(v int32) *GetCrawlerTypeCapabilitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCrawlerTypeCapabilitiesResponse) SetBody(v *GetCrawlerTypeCapabilitiesResponseBody) *GetCrawlerTypeCapabilitiesResponse {
	s.Body = v
	return s
}

func (s *GetCrawlerTypeCapabilitiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
