// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySslVpnServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySslVpnServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySslVpnServerResponse
	GetStatusCode() *int32
	SetBody(v *ModifySslVpnServerResponseBody) *ModifySslVpnServerResponse
	GetBody() *ModifySslVpnServerResponseBody
}

type ModifySslVpnServerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySslVpnServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySslVpnServerResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySslVpnServerResponse) GoString() string {
	return s.String()
}

func (s *ModifySslVpnServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySslVpnServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySslVpnServerResponse) GetBody() *ModifySslVpnServerResponseBody {
	return s.Body
}

func (s *ModifySslVpnServerResponse) SetHeaders(v map[string]*string) *ModifySslVpnServerResponse {
	s.Headers = v
	return s
}

func (s *ModifySslVpnServerResponse) SetStatusCode(v int32) *ModifySslVpnServerResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySslVpnServerResponse) SetBody(v *ModifySslVpnServerResponseBody) *ModifySslVpnServerResponse {
	s.Body = v
	return s
}

func (s *ModifySslVpnServerResponse) Validate() error {
	return dara.Validate(s)
}
