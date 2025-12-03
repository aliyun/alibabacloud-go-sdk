// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetServerCertificateNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetServerCertificateNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetServerCertificateNameResponse
	GetStatusCode() *int32
	SetBody(v *SetServerCertificateNameResponseBody) *SetServerCertificateNameResponse
	GetBody() *SetServerCertificateNameResponseBody
}

type SetServerCertificateNameResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetServerCertificateNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetServerCertificateNameResponse) String() string {
	return dara.Prettify(s)
}

func (s SetServerCertificateNameResponse) GoString() string {
	return s.String()
}

func (s *SetServerCertificateNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetServerCertificateNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetServerCertificateNameResponse) GetBody() *SetServerCertificateNameResponseBody {
	return s.Body
}

func (s *SetServerCertificateNameResponse) SetHeaders(v map[string]*string) *SetServerCertificateNameResponse {
	s.Headers = v
	return s
}

func (s *SetServerCertificateNameResponse) SetStatusCode(v int32) *SetServerCertificateNameResponse {
	s.StatusCode = &v
	return s
}

func (s *SetServerCertificateNameResponse) SetBody(v *SetServerCertificateNameResponseBody) *SetServerCertificateNameResponse {
	s.Body = v
	return s
}

func (s *SetServerCertificateNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
