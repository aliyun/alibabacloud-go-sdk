// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOriginCaCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOriginCaCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOriginCaCertificateResponse
	GetStatusCode() *int32
	SetBody(v *GetOriginCaCertificateResponseBody) *GetOriginCaCertificateResponse
	GetBody() *GetOriginCaCertificateResponseBody
}

type GetOriginCaCertificateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOriginCaCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOriginCaCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOriginCaCertificateResponse) GoString() string {
	return s.String()
}

func (s *GetOriginCaCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOriginCaCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOriginCaCertificateResponse) GetBody() *GetOriginCaCertificateResponseBody {
	return s.Body
}

func (s *GetOriginCaCertificateResponse) SetHeaders(v map[string]*string) *GetOriginCaCertificateResponse {
	s.Headers = v
	return s
}

func (s *GetOriginCaCertificateResponse) SetStatusCode(v int32) *GetOriginCaCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOriginCaCertificateResponse) SetBody(v *GetOriginCaCertificateResponseBody) *GetOriginCaCertificateResponse {
	s.Body = v
	return s
}

func (s *GetOriginCaCertificateResponse) Validate() error {
	return dara.Validate(s)
}
