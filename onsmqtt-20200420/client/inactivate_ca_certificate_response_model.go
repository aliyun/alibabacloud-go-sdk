// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInactivateCaCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InactivateCaCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InactivateCaCertificateResponse
	GetStatusCode() *int32
	SetBody(v *InactivateCaCertificateResponseBody) *InactivateCaCertificateResponse
	GetBody() *InactivateCaCertificateResponseBody
}

type InactivateCaCertificateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InactivateCaCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InactivateCaCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s InactivateCaCertificateResponse) GoString() string {
	return s.String()
}

func (s *InactivateCaCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InactivateCaCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InactivateCaCertificateResponse) GetBody() *InactivateCaCertificateResponseBody {
	return s.Body
}

func (s *InactivateCaCertificateResponse) SetHeaders(v map[string]*string) *InactivateCaCertificateResponse {
	s.Headers = v
	return s
}

func (s *InactivateCaCertificateResponse) SetStatusCode(v int32) *InactivateCaCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *InactivateCaCertificateResponse) SetBody(v *InactivateCaCertificateResponseBody) *InactivateCaCertificateResponse {
	s.Body = v
	return s
}

func (s *InactivateCaCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
