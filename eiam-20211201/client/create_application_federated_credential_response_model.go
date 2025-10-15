// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationFederatedCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateApplicationFederatedCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateApplicationFederatedCredentialResponse
	GetStatusCode() *int32
	SetBody(v *CreateApplicationFederatedCredentialResponseBody) *CreateApplicationFederatedCredentialResponse
	GetBody() *CreateApplicationFederatedCredentialResponseBody
}

type CreateApplicationFederatedCredentialResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApplicationFederatedCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApplicationFederatedCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationFederatedCredentialResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationFederatedCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateApplicationFederatedCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateApplicationFederatedCredentialResponse) GetBody() *CreateApplicationFederatedCredentialResponseBody {
	return s.Body
}

func (s *CreateApplicationFederatedCredentialResponse) SetHeaders(v map[string]*string) *CreateApplicationFederatedCredentialResponse {
	s.Headers = v
	return s
}

func (s *CreateApplicationFederatedCredentialResponse) SetStatusCode(v int32) *CreateApplicationFederatedCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApplicationFederatedCredentialResponse) SetBody(v *CreateApplicationFederatedCredentialResponseBody) *CreateApplicationFederatedCredentialResponse {
	s.Body = v
	return s
}

func (s *CreateApplicationFederatedCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
