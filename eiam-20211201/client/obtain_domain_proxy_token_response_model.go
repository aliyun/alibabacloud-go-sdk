// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObtainDomainProxyTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ObtainDomainProxyTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ObtainDomainProxyTokenResponse
	GetStatusCode() *int32
	SetBody(v *ObtainDomainProxyTokenResponseBody) *ObtainDomainProxyTokenResponse
	GetBody() *ObtainDomainProxyTokenResponseBody
}

type ObtainDomainProxyTokenResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ObtainDomainProxyTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ObtainDomainProxyTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s ObtainDomainProxyTokenResponse) GoString() string {
	return s.String()
}

func (s *ObtainDomainProxyTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ObtainDomainProxyTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ObtainDomainProxyTokenResponse) GetBody() *ObtainDomainProxyTokenResponseBody {
	return s.Body
}

func (s *ObtainDomainProxyTokenResponse) SetHeaders(v map[string]*string) *ObtainDomainProxyTokenResponse {
	s.Headers = v
	return s
}

func (s *ObtainDomainProxyTokenResponse) SetStatusCode(v int32) *ObtainDomainProxyTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *ObtainDomainProxyTokenResponse) SetBody(v *ObtainDomainProxyTokenResponseBody) *ObtainDomainProxyTokenResponse {
	s.Body = v
	return s
}

func (s *ObtainDomainProxyTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
