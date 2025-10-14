// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClientCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteClientCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteClientCertificateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteClientCertificateResponseBody) *DeleteClientCertificateResponse
	GetBody() *DeleteClientCertificateResponseBody
}

type DeleteClientCertificateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteClientCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteClientCertificateResponse) GoString() string {
	return s.String()
}

func (s *DeleteClientCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteClientCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteClientCertificateResponse) GetBody() *DeleteClientCertificateResponseBody {
	return s.Body
}

func (s *DeleteClientCertificateResponse) SetHeaders(v map[string]*string) *DeleteClientCertificateResponse {
	s.Headers = v
	return s
}

func (s *DeleteClientCertificateResponse) SetStatusCode(v int32) *DeleteClientCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteClientCertificateResponse) SetBody(v *DeleteClientCertificateResponseBody) *DeleteClientCertificateResponse {
	s.Body = v
	return s
}

func (s *DeleteClientCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
