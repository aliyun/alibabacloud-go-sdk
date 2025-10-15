// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationFederatedCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateApplicationFederatedCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateApplicationFederatedCredentialResponse
	GetStatusCode() *int32
	SetBody(v *UpdateApplicationFederatedCredentialResponseBody) *UpdateApplicationFederatedCredentialResponse
	GetBody() *UpdateApplicationFederatedCredentialResponseBody
}

type UpdateApplicationFederatedCredentialResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApplicationFederatedCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApplicationFederatedCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationFederatedCredentialResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationFederatedCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateApplicationFederatedCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateApplicationFederatedCredentialResponse) GetBody() *UpdateApplicationFederatedCredentialResponseBody {
	return s.Body
}

func (s *UpdateApplicationFederatedCredentialResponse) SetHeaders(v map[string]*string) *UpdateApplicationFederatedCredentialResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationFederatedCredentialResponse) SetStatusCode(v int32) *UpdateApplicationFederatedCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApplicationFederatedCredentialResponse) SetBody(v *UpdateApplicationFederatedCredentialResponseBody) *UpdateApplicationFederatedCredentialResponse {
	s.Body = v
	return s
}

func (s *UpdateApplicationFederatedCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
