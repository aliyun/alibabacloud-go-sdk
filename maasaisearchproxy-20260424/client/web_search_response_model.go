// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebSearchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *WebSearchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *WebSearchResponse
	GetStatusCode() *int32
	SetBody(v *WebSearchResponseBody) *WebSearchResponse
	GetBody() *WebSearchResponseBody
}

type WebSearchResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WebSearchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s WebSearchResponse) String() string {
	return dara.Prettify(s)
}

func (s WebSearchResponse) GoString() string {
	return s.String()
}

func (s *WebSearchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *WebSearchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *WebSearchResponse) GetBody() *WebSearchResponseBody {
	return s.Body
}

func (s *WebSearchResponse) SetHeaders(v map[string]*string) *WebSearchResponse {
	s.Headers = v
	return s
}

func (s *WebSearchResponse) SetStatusCode(v int32) *WebSearchResponse {
	s.StatusCode = &v
	return s
}

func (s *WebSearchResponse) SetBody(v *WebSearchResponseBody) *WebSearchResponse {
	s.Body = v
	return s
}

func (s *WebSearchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
