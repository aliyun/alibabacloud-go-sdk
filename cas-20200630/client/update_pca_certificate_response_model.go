// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePcaCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePcaCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePcaCertificateResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePcaCertificateResponseBody) *UpdatePcaCertificateResponse
	GetBody() *UpdatePcaCertificateResponseBody
}

type UpdatePcaCertificateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePcaCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePcaCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePcaCertificateResponse) GoString() string {
	return s.String()
}

func (s *UpdatePcaCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePcaCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePcaCertificateResponse) GetBody() *UpdatePcaCertificateResponseBody {
	return s.Body
}

func (s *UpdatePcaCertificateResponse) SetHeaders(v map[string]*string) *UpdatePcaCertificateResponse {
	s.Headers = v
	return s
}

func (s *UpdatePcaCertificateResponse) SetStatusCode(v int32) *UpdatePcaCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePcaCertificateResponse) SetBody(v *UpdatePcaCertificateResponseBody) *UpdatePcaCertificateResponse {
	s.Body = v
	return s
}

func (s *UpdatePcaCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
