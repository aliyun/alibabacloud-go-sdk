// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAdditionalCertificateWithListenerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAdditionalCertificateWithListenerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAdditionalCertificateWithListenerResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAdditionalCertificateWithListenerResponseBody) *UpdateAdditionalCertificateWithListenerResponse
	GetBody() *UpdateAdditionalCertificateWithListenerResponseBody
}

type UpdateAdditionalCertificateWithListenerResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAdditionalCertificateWithListenerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAdditionalCertificateWithListenerResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAdditionalCertificateWithListenerResponse) GoString() string {
	return s.String()
}

func (s *UpdateAdditionalCertificateWithListenerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAdditionalCertificateWithListenerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAdditionalCertificateWithListenerResponse) GetBody() *UpdateAdditionalCertificateWithListenerResponseBody {
	return s.Body
}

func (s *UpdateAdditionalCertificateWithListenerResponse) SetHeaders(v map[string]*string) *UpdateAdditionalCertificateWithListenerResponse {
	s.Headers = v
	return s
}

func (s *UpdateAdditionalCertificateWithListenerResponse) SetStatusCode(v int32) *UpdateAdditionalCertificateWithListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAdditionalCertificateWithListenerResponse) SetBody(v *UpdateAdditionalCertificateWithListenerResponseBody) *UpdateAdditionalCertificateWithListenerResponse {
	s.Body = v
	return s
}

func (s *UpdateAdditionalCertificateWithListenerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
