// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIdentityProviderStatusCheckJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIdentityProviderStatusCheckJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIdentityProviderStatusCheckJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateIdentityProviderStatusCheckJobResponseBody) *CreateIdentityProviderStatusCheckJobResponse
	GetBody() *CreateIdentityProviderStatusCheckJobResponseBody
}

type CreateIdentityProviderStatusCheckJobResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIdentityProviderStatusCheckJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIdentityProviderStatusCheckJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderStatusCheckJobResponse) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderStatusCheckJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIdentityProviderStatusCheckJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIdentityProviderStatusCheckJobResponse) GetBody() *CreateIdentityProviderStatusCheckJobResponseBody {
	return s.Body
}

func (s *CreateIdentityProviderStatusCheckJobResponse) SetHeaders(v map[string]*string) *CreateIdentityProviderStatusCheckJobResponse {
	s.Headers = v
	return s
}

func (s *CreateIdentityProviderStatusCheckJobResponse) SetStatusCode(v int32) *CreateIdentityProviderStatusCheckJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIdentityProviderStatusCheckJobResponse) SetBody(v *CreateIdentityProviderStatusCheckJobResponseBody) *CreateIdentityProviderStatusCheckJobResponse {
	s.Body = v
	return s
}

func (s *CreateIdentityProviderStatusCheckJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
