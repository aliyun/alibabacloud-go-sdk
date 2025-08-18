// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyCertificateResponse
	GetStatusCode() *int32
	SetBody(v *ApplyCertificateResponseBody) *ApplyCertificateResponse
	GetBody() *ApplyCertificateResponseBody
}

type ApplyCertificateResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyCertificateResponse) GoString() string {
	return s.String()
}

func (s *ApplyCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyCertificateResponse) GetBody() *ApplyCertificateResponseBody {
	return s.Body
}

func (s *ApplyCertificateResponse) SetHeaders(v map[string]*string) *ApplyCertificateResponse {
	s.Headers = v
	return s
}

func (s *ApplyCertificateResponse) SetStatusCode(v int32) *ApplyCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyCertificateResponse) SetBody(v *ApplyCertificateResponseBody) *ApplyCertificateResponse {
	s.Body = v
	return s
}

func (s *ApplyCertificateResponse) Validate() error {
	return dara.Validate(s)
}
