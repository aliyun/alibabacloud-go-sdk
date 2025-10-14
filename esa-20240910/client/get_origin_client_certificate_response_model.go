// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOriginClientCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOriginClientCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOriginClientCertificateResponse
	GetStatusCode() *int32
	SetBody(v *GetOriginClientCertificateResponseBody) *GetOriginClientCertificateResponse
	GetBody() *GetOriginClientCertificateResponseBody
}

type GetOriginClientCertificateResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOriginClientCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOriginClientCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOriginClientCertificateResponse) GoString() string {
	return s.String()
}

func (s *GetOriginClientCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOriginClientCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOriginClientCertificateResponse) GetBody() *GetOriginClientCertificateResponseBody {
	return s.Body
}

func (s *GetOriginClientCertificateResponse) SetHeaders(v map[string]*string) *GetOriginClientCertificateResponse {
	s.Headers = v
	return s
}

func (s *GetOriginClientCertificateResponse) SetStatusCode(v int32) *GetOriginClientCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOriginClientCertificateResponse) SetBody(v *GetOriginClientCertificateResponseBody) *GetOriginClientCertificateResponse {
	s.Body = v
	return s
}

func (s *GetOriginClientCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
