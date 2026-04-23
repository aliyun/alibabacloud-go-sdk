// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSSLDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSSLDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSSLDetailsResponse
	GetStatusCode() *int32
	SetBody(v *ListSSLDetailsResponseBody) *ListSSLDetailsResponse
	GetBody() *ListSSLDetailsResponseBody
}

type ListSSLDetailsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSSLDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSSLDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSSLDetailsResponse) GoString() string {
	return s.String()
}

func (s *ListSSLDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSSLDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSSLDetailsResponse) GetBody() *ListSSLDetailsResponseBody {
	return s.Body
}

func (s *ListSSLDetailsResponse) SetHeaders(v map[string]*string) *ListSSLDetailsResponse {
	s.Headers = v
	return s
}

func (s *ListSSLDetailsResponse) SetStatusCode(v int32) *ListSSLDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSSLDetailsResponse) SetBody(v *ListSSLDetailsResponseBody) *ListSSLDetailsResponse {
	s.Body = v
	return s
}

func (s *ListSSLDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
