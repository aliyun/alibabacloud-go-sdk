// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExternalCACertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateExternalCACertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateExternalCACertificateResponse
	GetStatusCode() *int32
	SetBody(v *CreateExternalCACertificateResponseBody) *CreateExternalCACertificateResponse
	GetBody() *CreateExternalCACertificateResponseBody
}

type CreateExternalCACertificateResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExternalCACertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExternalCACertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalCACertificateResponse) GoString() string {
	return s.String()
}

func (s *CreateExternalCACertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateExternalCACertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateExternalCACertificateResponse) GetBody() *CreateExternalCACertificateResponseBody {
	return s.Body
}

func (s *CreateExternalCACertificateResponse) SetHeaders(v map[string]*string) *CreateExternalCACertificateResponse {
	s.Headers = v
	return s
}

func (s *CreateExternalCACertificateResponse) SetStatusCode(v int32) *CreateExternalCACertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExternalCACertificateResponse) SetBody(v *CreateExternalCACertificateResponseBody) *CreateExternalCACertificateResponse {
	s.Body = v
	return s
}

func (s *CreateExternalCACertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
