// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSSLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableSSLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableSSLResponse
	GetStatusCode() *int32
	SetBody(v *DisableSSLResponseBody) *DisableSSLResponse
	GetBody() *DisableSSLResponseBody
}

type DisableSSLResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableSSLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableSSLResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableSSLResponse) GoString() string {
	return s.String()
}

func (s *DisableSSLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableSSLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableSSLResponse) GetBody() *DisableSSLResponseBody {
	return s.Body
}

func (s *DisableSSLResponse) SetHeaders(v map[string]*string) *DisableSSLResponse {
	s.Headers = v
	return s
}

func (s *DisableSSLResponse) SetStatusCode(v int32) *DisableSSLResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableSSLResponse) SetBody(v *DisableSSLResponseBody) *DisableSSLResponse {
	s.Body = v
	return s
}

func (s *DisableSSLResponse) Validate() error {
	return dara.Validate(s)
}
