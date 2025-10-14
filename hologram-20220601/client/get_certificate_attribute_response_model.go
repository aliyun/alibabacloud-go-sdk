// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCertificateAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCertificateAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCertificateAttributeResponse
	GetStatusCode() *int32
	SetBody(v *GetCertificateAttributeResponseBody) *GetCertificateAttributeResponse
	GetBody() *GetCertificateAttributeResponseBody
}

type GetCertificateAttributeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCertificateAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCertificateAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCertificateAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetCertificateAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCertificateAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCertificateAttributeResponse) GetBody() *GetCertificateAttributeResponseBody {
	return s.Body
}

func (s *GetCertificateAttributeResponse) SetHeaders(v map[string]*string) *GetCertificateAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetCertificateAttributeResponse) SetStatusCode(v int32) *GetCertificateAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCertificateAttributeResponse) SetBody(v *GetCertificateAttributeResponseBody) *GetCertificateAttributeResponse {
	s.Body = v
	return s
}

func (s *GetCertificateAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
