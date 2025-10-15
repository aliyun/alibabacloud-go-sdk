// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationFederatedCredentialDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateApplicationFederatedCredentialDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateApplicationFederatedCredentialDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateApplicationFederatedCredentialDescriptionResponseBody) *UpdateApplicationFederatedCredentialDescriptionResponse
	GetBody() *UpdateApplicationFederatedCredentialDescriptionResponseBody
}

type UpdateApplicationFederatedCredentialDescriptionResponse struct {
	Headers    map[string]*string                                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApplicationFederatedCredentialDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApplicationFederatedCredentialDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationFederatedCredentialDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationFederatedCredentialDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateApplicationFederatedCredentialDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateApplicationFederatedCredentialDescriptionResponse) GetBody() *UpdateApplicationFederatedCredentialDescriptionResponseBody {
	return s.Body
}

func (s *UpdateApplicationFederatedCredentialDescriptionResponse) SetHeaders(v map[string]*string) *UpdateApplicationFederatedCredentialDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationFederatedCredentialDescriptionResponse) SetStatusCode(v int32) *UpdateApplicationFederatedCredentialDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApplicationFederatedCredentialDescriptionResponse) SetBody(v *UpdateApplicationFederatedCredentialDescriptionResponseBody) *UpdateApplicationFederatedCredentialDescriptionResponse {
	s.Body = v
	return s
}

func (s *UpdateApplicationFederatedCredentialDescriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
