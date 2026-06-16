// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebFetchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *WebFetchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *WebFetchResponse
	GetStatusCode() *int32
	SetBody(v *WebFetchResponseBody) *WebFetchResponse
	GetBody() *WebFetchResponseBody
}

type WebFetchResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WebFetchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s WebFetchResponse) String() string {
	return dara.Prettify(s)
}

func (s WebFetchResponse) GoString() string {
	return s.String()
}

func (s *WebFetchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *WebFetchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *WebFetchResponse) GetBody() *WebFetchResponseBody {
	return s.Body
}

func (s *WebFetchResponse) SetHeaders(v map[string]*string) *WebFetchResponse {
	s.Headers = v
	return s
}

func (s *WebFetchResponse) SetStatusCode(v int32) *WebFetchResponse {
	s.StatusCode = &v
	return s
}

func (s *WebFetchResponse) SetBody(v *WebFetchResponseBody) *WebFetchResponse {
	s.Body = v
	return s
}

func (s *WebFetchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
