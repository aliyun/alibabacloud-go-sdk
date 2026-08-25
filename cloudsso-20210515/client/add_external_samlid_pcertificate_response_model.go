// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddExternalSAMLIdPCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddExternalSAMLIdPCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddExternalSAMLIdPCertificateResponse
	GetStatusCode() *int32
	SetBody(v *AddExternalSAMLIdPCertificateResponseBody) *AddExternalSAMLIdPCertificateResponse
	GetBody() *AddExternalSAMLIdPCertificateResponseBody
}

type AddExternalSAMLIdPCertificateResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddExternalSAMLIdPCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddExternalSAMLIdPCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s AddExternalSAMLIdPCertificateResponse) GoString() string {
	return s.String()
}

func (s *AddExternalSAMLIdPCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddExternalSAMLIdPCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddExternalSAMLIdPCertificateResponse) GetBody() *AddExternalSAMLIdPCertificateResponseBody {
	return s.Body
}

func (s *AddExternalSAMLIdPCertificateResponse) SetHeaders(v map[string]*string) *AddExternalSAMLIdPCertificateResponse {
	s.Headers = v
	return s
}

func (s *AddExternalSAMLIdPCertificateResponse) SetStatusCode(v int32) *AddExternalSAMLIdPCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *AddExternalSAMLIdPCertificateResponse) SetBody(v *AddExternalSAMLIdPCertificateResponseBody) *AddExternalSAMLIdPCertificateResponse {
	s.Body = v
	return s
}

func (s *AddExternalSAMLIdPCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
