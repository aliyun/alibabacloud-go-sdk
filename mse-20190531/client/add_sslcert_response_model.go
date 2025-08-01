// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSSLCertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddSSLCertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddSSLCertResponse
	GetStatusCode() *int32
	SetBody(v *AddSSLCertResponseBody) *AddSSLCertResponse
	GetBody() *AddSSLCertResponseBody
}

type AddSSLCertResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddSSLCertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddSSLCertResponse) String() string {
	return dara.Prettify(s)
}

func (s AddSSLCertResponse) GoString() string {
	return s.String()
}

func (s *AddSSLCertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddSSLCertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddSSLCertResponse) GetBody() *AddSSLCertResponseBody {
	return s.Body
}

func (s *AddSSLCertResponse) SetHeaders(v map[string]*string) *AddSSLCertResponse {
	s.Headers = v
	return s
}

func (s *AddSSLCertResponse) SetStatusCode(v int32) *AddSSLCertResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSSLCertResponse) SetBody(v *AddSSLCertResponseBody) *AddSSLCertResponse {
	s.Body = v
	return s
}

func (s *AddSSLCertResponse) Validate() error {
	return dara.Validate(s)
}
