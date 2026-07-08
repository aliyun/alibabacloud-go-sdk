// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeWHClientCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeWHClientCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeWHClientCertificateResponse
	GetStatusCode() *int32
	SetBody(v *RevokeWHClientCertificateResponseBody) *RevokeWHClientCertificateResponse
	GetBody() *RevokeWHClientCertificateResponseBody
}

type RevokeWHClientCertificateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeWHClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeWHClientCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeWHClientCertificateResponse) GoString() string {
	return s.String()
}

func (s *RevokeWHClientCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeWHClientCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeWHClientCertificateResponse) GetBody() *RevokeWHClientCertificateResponseBody {
	return s.Body
}

func (s *RevokeWHClientCertificateResponse) SetHeaders(v map[string]*string) *RevokeWHClientCertificateResponse {
	s.Headers = v
	return s
}

func (s *RevokeWHClientCertificateResponse) SetStatusCode(v int32) *RevokeWHClientCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeWHClientCertificateResponse) SetBody(v *RevokeWHClientCertificateResponseBody) *RevokeWHClientCertificateResponse {
	s.Body = v
	return s
}

func (s *RevokeWHClientCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
