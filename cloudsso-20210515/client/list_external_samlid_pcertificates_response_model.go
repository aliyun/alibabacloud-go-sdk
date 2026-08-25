// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExternalSAMLIdPCertificatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListExternalSAMLIdPCertificatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListExternalSAMLIdPCertificatesResponse
	GetStatusCode() *int32
	SetBody(v *ListExternalSAMLIdPCertificatesResponseBody) *ListExternalSAMLIdPCertificatesResponse
	GetBody() *ListExternalSAMLIdPCertificatesResponseBody
}

type ListExternalSAMLIdPCertificatesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExternalSAMLIdPCertificatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExternalSAMLIdPCertificatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListExternalSAMLIdPCertificatesResponse) GoString() string {
	return s.String()
}

func (s *ListExternalSAMLIdPCertificatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListExternalSAMLIdPCertificatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListExternalSAMLIdPCertificatesResponse) GetBody() *ListExternalSAMLIdPCertificatesResponseBody {
	return s.Body
}

func (s *ListExternalSAMLIdPCertificatesResponse) SetHeaders(v map[string]*string) *ListExternalSAMLIdPCertificatesResponse {
	s.Headers = v
	return s
}

func (s *ListExternalSAMLIdPCertificatesResponse) SetStatusCode(v int32) *ListExternalSAMLIdPCertificatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExternalSAMLIdPCertificatesResponse) SetBody(v *ListExternalSAMLIdPCertificatesResponseBody) *ListExternalSAMLIdPCertificatesResponse {
	s.Body = v
	return s
}

func (s *ListExternalSAMLIdPCertificatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
