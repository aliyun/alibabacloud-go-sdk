// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCredentialProviderDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCredentialProviderDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCredentialProviderDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCredentialProviderDescriptionResponseBody) *UpdateCredentialProviderDescriptionResponse
	GetBody() *UpdateCredentialProviderDescriptionResponseBody
}

type UpdateCredentialProviderDescriptionResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCredentialProviderDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCredentialProviderDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialProviderDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateCredentialProviderDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCredentialProviderDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCredentialProviderDescriptionResponse) GetBody() *UpdateCredentialProviderDescriptionResponseBody {
	return s.Body
}

func (s *UpdateCredentialProviderDescriptionResponse) SetHeaders(v map[string]*string) *UpdateCredentialProviderDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateCredentialProviderDescriptionResponse) SetStatusCode(v int32) *UpdateCredentialProviderDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCredentialProviderDescriptionResponse) SetBody(v *UpdateCredentialProviderDescriptionResponseBody) *UpdateCredentialProviderDescriptionResponse {
	s.Body = v
	return s
}

func (s *UpdateCredentialProviderDescriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
