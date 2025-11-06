// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSSLCertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSSLCertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSSLCertResponse
	GetStatusCode() *int32
	SetBody(v *ListSSLCertResponseBody) *ListSSLCertResponse
	GetBody() *ListSSLCertResponseBody
}

type ListSSLCertResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSSLCertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSSLCertResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSSLCertResponse) GoString() string {
	return s.String()
}

func (s *ListSSLCertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSSLCertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSSLCertResponse) GetBody() *ListSSLCertResponseBody {
	return s.Body
}

func (s *ListSSLCertResponse) SetHeaders(v map[string]*string) *ListSSLCertResponse {
	s.Headers = v
	return s
}

func (s *ListSSLCertResponse) SetStatusCode(v int32) *ListSSLCertResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSSLCertResponse) SetBody(v *ListSSLCertResponseBody) *ListSSLCertResponse {
	s.Body = v
	return s
}

func (s *ListSSLCertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
