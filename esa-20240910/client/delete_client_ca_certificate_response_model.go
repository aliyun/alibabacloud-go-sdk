// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClientCaCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteClientCaCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteClientCaCertificateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteClientCaCertificateResponseBody) *DeleteClientCaCertificateResponse
	GetBody() *DeleteClientCaCertificateResponseBody
}

type DeleteClientCaCertificateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteClientCaCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteClientCaCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteClientCaCertificateResponse) GoString() string {
	return s.String()
}

func (s *DeleteClientCaCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteClientCaCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteClientCaCertificateResponse) GetBody() *DeleteClientCaCertificateResponseBody {
	return s.Body
}

func (s *DeleteClientCaCertificateResponse) SetHeaders(v map[string]*string) *DeleteClientCaCertificateResponse {
	s.Headers = v
	return s
}

func (s *DeleteClientCaCertificateResponse) SetStatusCode(v int32) *DeleteClientCaCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteClientCaCertificateResponse) SetBody(v *DeleteClientCaCertificateResponseBody) *DeleteClientCaCertificateResponse {
	s.Body = v
	return s
}

func (s *DeleteClientCaCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
