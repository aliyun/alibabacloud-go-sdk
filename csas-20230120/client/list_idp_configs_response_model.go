// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIdpConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIdpConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIdpConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListIdpConfigsResponseBody) *ListIdpConfigsResponse
	GetBody() *ListIdpConfigsResponseBody
}

type ListIdpConfigsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIdpConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIdpConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIdpConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListIdpConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIdpConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIdpConfigsResponse) GetBody() *ListIdpConfigsResponseBody {
	return s.Body
}

func (s *ListIdpConfigsResponse) SetHeaders(v map[string]*string) *ListIdpConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListIdpConfigsResponse) SetStatusCode(v int32) *ListIdpConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIdpConfigsResponse) SetBody(v *ListIdpConfigsResponseBody) *ListIdpConfigsResponse {
	s.Body = v
	return s
}

func (s *ListIdpConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
