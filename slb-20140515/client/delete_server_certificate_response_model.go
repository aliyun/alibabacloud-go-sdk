// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServerCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteServerCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteServerCertificateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteServerCertificateResponseBody) *DeleteServerCertificateResponse
	GetBody() *DeleteServerCertificateResponseBody
}

type DeleteServerCertificateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServerCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServerCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteServerCertificateResponse) GoString() string {
	return s.String()
}

func (s *DeleteServerCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteServerCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteServerCertificateResponse) GetBody() *DeleteServerCertificateResponseBody {
	return s.Body
}

func (s *DeleteServerCertificateResponse) SetHeaders(v map[string]*string) *DeleteServerCertificateResponse {
	s.Headers = v
	return s
}

func (s *DeleteServerCertificateResponse) SetStatusCode(v int32) *DeleteServerCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServerCertificateResponse) SetBody(v *DeleteServerCertificateResponseBody) *DeleteServerCertificateResponse {
	s.Body = v
	return s
}

func (s *DeleteServerCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
