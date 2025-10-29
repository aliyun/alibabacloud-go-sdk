// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCertificateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCertificateResponseBody) *DeleteCertificateResponse
	GetBody() *DeleteCertificateResponseBody
}

type DeleteCertificateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCertificateResponse) GoString() string {
	return s.String()
}

func (s *DeleteCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCertificateResponse) GetBody() *DeleteCertificateResponseBody {
	return s.Body
}

func (s *DeleteCertificateResponse) SetHeaders(v map[string]*string) *DeleteCertificateResponse {
	s.Headers = v
	return s
}

func (s *DeleteCertificateResponse) SetStatusCode(v int32) *DeleteCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCertificateResponse) SetBody(v *DeleteCertificateResponseBody) *DeleteCertificateResponse {
	s.Body = v
	return s
}

func (s *DeleteCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
