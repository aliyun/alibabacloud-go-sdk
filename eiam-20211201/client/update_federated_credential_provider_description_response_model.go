// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFederatedCredentialProviderDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFederatedCredentialProviderDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFederatedCredentialProviderDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFederatedCredentialProviderDescriptionResponseBody) *UpdateFederatedCredentialProviderDescriptionResponse
	GetBody() *UpdateFederatedCredentialProviderDescriptionResponseBody
}

type UpdateFederatedCredentialProviderDescriptionResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFederatedCredentialProviderDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFederatedCredentialProviderDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFederatedCredentialProviderDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateFederatedCredentialProviderDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFederatedCredentialProviderDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFederatedCredentialProviderDescriptionResponse) GetBody() *UpdateFederatedCredentialProviderDescriptionResponseBody {
	return s.Body
}

func (s *UpdateFederatedCredentialProviderDescriptionResponse) SetHeaders(v map[string]*string) *UpdateFederatedCredentialProviderDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateFederatedCredentialProviderDescriptionResponse) SetStatusCode(v int32) *UpdateFederatedCredentialProviderDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFederatedCredentialProviderDescriptionResponse) SetBody(v *UpdateFederatedCredentialProviderDescriptionResponseBody) *UpdateFederatedCredentialProviderDescriptionResponse {
	s.Body = v
	return s
}

func (s *UpdateFederatedCredentialProviderDescriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
