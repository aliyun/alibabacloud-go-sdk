// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableApplicationFederatedCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableApplicationFederatedCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableApplicationFederatedCredentialResponse
	GetStatusCode() *int32
	SetBody(v *DisableApplicationFederatedCredentialResponseBody) *DisableApplicationFederatedCredentialResponse
	GetBody() *DisableApplicationFederatedCredentialResponseBody
}

type DisableApplicationFederatedCredentialResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableApplicationFederatedCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableApplicationFederatedCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableApplicationFederatedCredentialResponse) GoString() string {
	return s.String()
}

func (s *DisableApplicationFederatedCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableApplicationFederatedCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableApplicationFederatedCredentialResponse) GetBody() *DisableApplicationFederatedCredentialResponseBody {
	return s.Body
}

func (s *DisableApplicationFederatedCredentialResponse) SetHeaders(v map[string]*string) *DisableApplicationFederatedCredentialResponse {
	s.Headers = v
	return s
}

func (s *DisableApplicationFederatedCredentialResponse) SetStatusCode(v int32) *DisableApplicationFederatedCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableApplicationFederatedCredentialResponse) SetBody(v *DisableApplicationFederatedCredentialResponseBody) *DisableApplicationFederatedCredentialResponse {
	s.Body = v
	return s
}

func (s *DisableApplicationFederatedCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
