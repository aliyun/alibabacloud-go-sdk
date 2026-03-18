// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCredentialProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCredentialProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCredentialProviderResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCredentialProviderResponseBody) *UpdateCredentialProviderResponse
	GetBody() *UpdateCredentialProviderResponseBody
}

type UpdateCredentialProviderResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCredentialProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCredentialProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialProviderResponse) GoString() string {
	return s.String()
}

func (s *UpdateCredentialProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCredentialProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCredentialProviderResponse) GetBody() *UpdateCredentialProviderResponseBody {
	return s.Body
}

func (s *UpdateCredentialProviderResponse) SetHeaders(v map[string]*string) *UpdateCredentialProviderResponse {
	s.Headers = v
	return s
}

func (s *UpdateCredentialProviderResponse) SetStatusCode(v int32) *UpdateCredentialProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCredentialProviderResponse) SetBody(v *UpdateCredentialProviderResponseBody) *UpdateCredentialProviderResponse {
	s.Body = v
	return s
}

func (s *UpdateCredentialProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
