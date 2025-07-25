// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDNSSLBStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDNSSLBStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDNSSLBStatusResponse
	GetStatusCode() *int32
	SetBody(v *SetDNSSLBStatusResponseBody) *SetDNSSLBStatusResponse
	GetBody() *SetDNSSLBStatusResponseBody
}

type SetDNSSLBStatusResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDNSSLBStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDNSSLBStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDNSSLBStatusResponse) GoString() string {
	return s.String()
}

func (s *SetDNSSLBStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDNSSLBStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDNSSLBStatusResponse) GetBody() *SetDNSSLBStatusResponseBody {
	return s.Body
}

func (s *SetDNSSLBStatusResponse) SetHeaders(v map[string]*string) *SetDNSSLBStatusResponse {
	s.Headers = v
	return s
}

func (s *SetDNSSLBStatusResponse) SetStatusCode(v int32) *SetDNSSLBStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDNSSLBStatusResponse) SetBody(v *SetDNSSLBStatusResponseBody) *SetDNSSLBStatusResponse {
	s.Body = v
	return s
}

func (s *SetDNSSLBStatusResponse) Validate() error {
	return dara.Validate(s)
}
