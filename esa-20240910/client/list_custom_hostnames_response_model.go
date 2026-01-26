// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomHostnamesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCustomHostnamesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCustomHostnamesResponse
	GetStatusCode() *int32
	SetBody(v *ListCustomHostnamesResponseBody) *ListCustomHostnamesResponse
	GetBody() *ListCustomHostnamesResponseBody
}

type ListCustomHostnamesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomHostnamesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomHostnamesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCustomHostnamesResponse) GoString() string {
	return s.String()
}

func (s *ListCustomHostnamesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCustomHostnamesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCustomHostnamesResponse) GetBody() *ListCustomHostnamesResponseBody {
	return s.Body
}

func (s *ListCustomHostnamesResponse) SetHeaders(v map[string]*string) *ListCustomHostnamesResponse {
	s.Headers = v
	return s
}

func (s *ListCustomHostnamesResponse) SetStatusCode(v int32) *ListCustomHostnamesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomHostnamesResponse) SetBody(v *ListCustomHostnamesResponseBody) *ListCustomHostnamesResponse {
	s.Body = v
	return s
}

func (s *ListCustomHostnamesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
