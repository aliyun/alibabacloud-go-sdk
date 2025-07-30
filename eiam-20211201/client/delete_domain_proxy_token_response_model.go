// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDomainProxyTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDomainProxyTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDomainProxyTokenResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDomainProxyTokenResponseBody) *DeleteDomainProxyTokenResponse
	GetBody() *DeleteDomainProxyTokenResponseBody
}

type DeleteDomainProxyTokenResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDomainProxyTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDomainProxyTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDomainProxyTokenResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainProxyTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDomainProxyTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDomainProxyTokenResponse) GetBody() *DeleteDomainProxyTokenResponseBody {
	return s.Body
}

func (s *DeleteDomainProxyTokenResponse) SetHeaders(v map[string]*string) *DeleteDomainProxyTokenResponse {
	s.Headers = v
	return s
}

func (s *DeleteDomainProxyTokenResponse) SetStatusCode(v int32) *DeleteDomainProxyTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDomainProxyTokenResponse) SetBody(v *DeleteDomainProxyTokenResponseBody) *DeleteDomainProxyTokenResponse {
	s.Body = v
	return s
}

func (s *DeleteDomainProxyTokenResponse) Validate() error {
	return dara.Validate(s)
}
