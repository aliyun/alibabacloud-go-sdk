// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDomainProxyTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDomainProxyTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDomainProxyTokenResponse
	GetStatusCode() *int32
	SetBody(v *CreateDomainProxyTokenResponseBody) *CreateDomainProxyTokenResponse
	GetBody() *CreateDomainProxyTokenResponseBody
}

type CreateDomainProxyTokenResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDomainProxyTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDomainProxyTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDomainProxyTokenResponse) GoString() string {
	return s.String()
}

func (s *CreateDomainProxyTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDomainProxyTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDomainProxyTokenResponse) GetBody() *CreateDomainProxyTokenResponseBody {
	return s.Body
}

func (s *CreateDomainProxyTokenResponse) SetHeaders(v map[string]*string) *CreateDomainProxyTokenResponse {
	s.Headers = v
	return s
}

func (s *CreateDomainProxyTokenResponse) SetStatusCode(v int32) *CreateDomainProxyTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDomainProxyTokenResponse) SetBody(v *CreateDomainProxyTokenResponseBody) *CreateDomainProxyTokenResponse {
	s.Body = v
	return s
}

func (s *CreateDomainProxyTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
