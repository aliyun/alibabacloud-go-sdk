// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCaCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCaCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCaCertificateResponse
	GetStatusCode() *int32
	SetBody(v *GetCaCertificateResponseBody) *GetCaCertificateResponse
	GetBody() *GetCaCertificateResponseBody
}

type GetCaCertificateResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCaCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCaCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCaCertificateResponse) GoString() string {
	return s.String()
}

func (s *GetCaCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCaCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCaCertificateResponse) GetBody() *GetCaCertificateResponseBody {
	return s.Body
}

func (s *GetCaCertificateResponse) SetHeaders(v map[string]*string) *GetCaCertificateResponse {
	s.Headers = v
	return s
}

func (s *GetCaCertificateResponse) SetStatusCode(v int32) *GetCaCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCaCertificateResponse) SetBody(v *GetCaCertificateResponseBody) *GetCaCertificateResponse {
	s.Body = v
	return s
}

func (s *GetCaCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
