// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSSLConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableSSLConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableSSLConnectionResponse
	GetStatusCode() *int32
	SetBody(v *DisableSSLConnectionResponseBody) *DisableSSLConnectionResponse
	GetBody() *DisableSSLConnectionResponseBody
}

type DisableSSLConnectionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableSSLConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableSSLConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableSSLConnectionResponse) GoString() string {
	return s.String()
}

func (s *DisableSSLConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableSSLConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableSSLConnectionResponse) GetBody() *DisableSSLConnectionResponseBody {
	return s.Body
}

func (s *DisableSSLConnectionResponse) SetHeaders(v map[string]*string) *DisableSSLConnectionResponse {
	s.Headers = v
	return s
}

func (s *DisableSSLConnectionResponse) SetStatusCode(v int32) *DisableSSLConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableSSLConnectionResponse) SetBody(v *DisableSSLConnectionResponseBody) *DisableSSLConnectionResponse {
	s.Body = v
	return s
}

func (s *DisableSSLConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
