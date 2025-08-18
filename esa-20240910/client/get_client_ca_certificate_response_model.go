// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientCaCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetClientCaCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetClientCaCertificateResponse
	GetStatusCode() *int32
	SetBody(v *GetClientCaCertificateResponseBody) *GetClientCaCertificateResponse
	GetBody() *GetClientCaCertificateResponseBody
}

type GetClientCaCertificateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClientCaCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClientCaCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetClientCaCertificateResponse) GoString() string {
	return s.String()
}

func (s *GetClientCaCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetClientCaCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetClientCaCertificateResponse) GetBody() *GetClientCaCertificateResponseBody {
	return s.Body
}

func (s *GetClientCaCertificateResponse) SetHeaders(v map[string]*string) *GetClientCaCertificateResponse {
	s.Headers = v
	return s
}

func (s *GetClientCaCertificateResponse) SetStatusCode(v int32) *GetClientCaCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClientCaCertificateResponse) SetBody(v *GetClientCaCertificateResponseBody) *GetClientCaCertificateResponse {
	s.Body = v
	return s
}

func (s *GetClientCaCertificateResponse) Validate() error {
	return dara.Validate(s)
}
