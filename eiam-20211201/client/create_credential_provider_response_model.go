// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCredentialProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCredentialProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCredentialProviderResponse
	GetStatusCode() *int32
	SetBody(v *CreateCredentialProviderResponseBody) *CreateCredentialProviderResponse
	GetBody() *CreateCredentialProviderResponseBody
}

type CreateCredentialProviderResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCredentialProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCredentialProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialProviderResponse) GoString() string {
	return s.String()
}

func (s *CreateCredentialProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCredentialProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCredentialProviderResponse) GetBody() *CreateCredentialProviderResponseBody {
	return s.Body
}

func (s *CreateCredentialProviderResponse) SetHeaders(v map[string]*string) *CreateCredentialProviderResponse {
	s.Headers = v
	return s
}

func (s *CreateCredentialProviderResponse) SetStatusCode(v int32) *CreateCredentialProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCredentialProviderResponse) SetBody(v *CreateCredentialProviderResponseBody) *CreateCredentialProviderResponse {
	s.Body = v
	return s
}

func (s *CreateCredentialProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
