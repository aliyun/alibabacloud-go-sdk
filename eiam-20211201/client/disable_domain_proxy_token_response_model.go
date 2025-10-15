// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDomainProxyTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableDomainProxyTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableDomainProxyTokenResponse
	GetStatusCode() *int32
	SetBody(v *DisableDomainProxyTokenResponseBody) *DisableDomainProxyTokenResponse
	GetBody() *DisableDomainProxyTokenResponseBody
}

type DisableDomainProxyTokenResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableDomainProxyTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableDomainProxyTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableDomainProxyTokenResponse) GoString() string {
	return s.String()
}

func (s *DisableDomainProxyTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableDomainProxyTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableDomainProxyTokenResponse) GetBody() *DisableDomainProxyTokenResponseBody {
	return s.Body
}

func (s *DisableDomainProxyTokenResponse) SetHeaders(v map[string]*string) *DisableDomainProxyTokenResponse {
	s.Headers = v
	return s
}

func (s *DisableDomainProxyTokenResponse) SetStatusCode(v int32) *DisableDomainProxyTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableDomainProxyTokenResponse) SetBody(v *DisableDomainProxyTokenResponseBody) *DisableDomainProxyTokenResponse {
	s.Body = v
	return s
}

func (s *DisableDomainProxyTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
