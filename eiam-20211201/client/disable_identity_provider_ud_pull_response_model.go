// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableIdentityProviderUdPullResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableIdentityProviderUdPullResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableIdentityProviderUdPullResponse
	GetStatusCode() *int32
	SetBody(v *DisableIdentityProviderUdPullResponseBody) *DisableIdentityProviderUdPullResponse
	GetBody() *DisableIdentityProviderUdPullResponseBody
}

type DisableIdentityProviderUdPullResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableIdentityProviderUdPullResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableIdentityProviderUdPullResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableIdentityProviderUdPullResponse) GoString() string {
	return s.String()
}

func (s *DisableIdentityProviderUdPullResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableIdentityProviderUdPullResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableIdentityProviderUdPullResponse) GetBody() *DisableIdentityProviderUdPullResponseBody {
	return s.Body
}

func (s *DisableIdentityProviderUdPullResponse) SetHeaders(v map[string]*string) *DisableIdentityProviderUdPullResponse {
	s.Headers = v
	return s
}

func (s *DisableIdentityProviderUdPullResponse) SetStatusCode(v int32) *DisableIdentityProviderUdPullResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableIdentityProviderUdPullResponse) SetBody(v *DisableIdentityProviderUdPullResponseBody) *DisableIdentityProviderUdPullResponse {
	s.Body = v
	return s
}

func (s *DisableIdentityProviderUdPullResponse) Validate() error {
	return dara.Validate(s)
}
