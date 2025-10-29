// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadPageScrapeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReadPageScrapeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReadPageScrapeResponse
	GetStatusCode() *int32
	SetBody(v *ReadPageScrapeResponseBody) *ReadPageScrapeResponse
	GetBody() *ReadPageScrapeResponseBody
}

type ReadPageScrapeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadPageScrapeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadPageScrapeResponse) String() string {
	return dara.Prettify(s)
}

func (s ReadPageScrapeResponse) GoString() string {
	return s.String()
}

func (s *ReadPageScrapeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReadPageScrapeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReadPageScrapeResponse) GetBody() *ReadPageScrapeResponseBody {
	return s.Body
}

func (s *ReadPageScrapeResponse) SetHeaders(v map[string]*string) *ReadPageScrapeResponse {
	s.Headers = v
	return s
}

func (s *ReadPageScrapeResponse) SetStatusCode(v int32) *ReadPageScrapeResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadPageScrapeResponse) SetBody(v *ReadPageScrapeResponseBody) *ReadPageScrapeResponse {
	s.Body = v
	return s
}

func (s *ReadPageScrapeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
