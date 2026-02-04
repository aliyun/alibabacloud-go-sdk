// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderStatusCheckJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIdentityProviderStatusCheckJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIdentityProviderStatusCheckJobResponse
	GetStatusCode() *int32
	SetBody(v *GetIdentityProviderStatusCheckJobResponseBody) *GetIdentityProviderStatusCheckJobResponse
	GetBody() *GetIdentityProviderStatusCheckJobResponseBody
}

type GetIdentityProviderStatusCheckJobResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIdentityProviderStatusCheckJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIdentityProviderStatusCheckJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderStatusCheckJobResponse) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderStatusCheckJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIdentityProviderStatusCheckJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIdentityProviderStatusCheckJobResponse) GetBody() *GetIdentityProviderStatusCheckJobResponseBody {
	return s.Body
}

func (s *GetIdentityProviderStatusCheckJobResponse) SetHeaders(v map[string]*string) *GetIdentityProviderStatusCheckJobResponse {
	s.Headers = v
	return s
}

func (s *GetIdentityProviderStatusCheckJobResponse) SetStatusCode(v int32) *GetIdentityProviderStatusCheckJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIdentityProviderStatusCheckJobResponse) SetBody(v *GetIdentityProviderStatusCheckJobResponseBody) *GetIdentityProviderStatusCheckJobResponse {
	s.Body = v
	return s
}

func (s *GetIdentityProviderStatusCheckJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
