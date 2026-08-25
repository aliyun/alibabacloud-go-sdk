// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveExternalSAMLIdPCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveExternalSAMLIdPCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveExternalSAMLIdPCertificateResponse
	GetStatusCode() *int32
	SetBody(v *RemoveExternalSAMLIdPCertificateResponseBody) *RemoveExternalSAMLIdPCertificateResponse
	GetBody() *RemoveExternalSAMLIdPCertificateResponseBody
}

type RemoveExternalSAMLIdPCertificateResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveExternalSAMLIdPCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveExternalSAMLIdPCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveExternalSAMLIdPCertificateResponse) GoString() string {
	return s.String()
}

func (s *RemoveExternalSAMLIdPCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveExternalSAMLIdPCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveExternalSAMLIdPCertificateResponse) GetBody() *RemoveExternalSAMLIdPCertificateResponseBody {
	return s.Body
}

func (s *RemoveExternalSAMLIdPCertificateResponse) SetHeaders(v map[string]*string) *RemoveExternalSAMLIdPCertificateResponse {
	s.Headers = v
	return s
}

func (s *RemoveExternalSAMLIdPCertificateResponse) SetStatusCode(v int32) *RemoveExternalSAMLIdPCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveExternalSAMLIdPCertificateResponse) SetBody(v *RemoveExternalSAMLIdPCertificateResponseBody) *RemoveExternalSAMLIdPCertificateResponse {
	s.Body = v
	return s
}

func (s *RemoveExternalSAMLIdPCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
