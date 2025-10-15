// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationFederatedCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApplicationFederatedCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApplicationFederatedCredentialResponse
	GetStatusCode() *int32
	SetBody(v *GetApplicationFederatedCredentialResponseBody) *GetApplicationFederatedCredentialResponse
	GetBody() *GetApplicationFederatedCredentialResponseBody
}

type GetApplicationFederatedCredentialResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApplicationFederatedCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApplicationFederatedCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationFederatedCredentialResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationFederatedCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApplicationFederatedCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApplicationFederatedCredentialResponse) GetBody() *GetApplicationFederatedCredentialResponseBody {
	return s.Body
}

func (s *GetApplicationFederatedCredentialResponse) SetHeaders(v map[string]*string) *GetApplicationFederatedCredentialResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationFederatedCredentialResponse) SetStatusCode(v int32) *GetApplicationFederatedCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplicationFederatedCredentialResponse) SetBody(v *GetApplicationFederatedCredentialResponseBody) *GetApplicationFederatedCredentialResponse {
	s.Body = v
	return s
}

func (s *GetApplicationFederatedCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
