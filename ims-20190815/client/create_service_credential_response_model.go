// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateServiceCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateServiceCredentialResponse
	GetStatusCode() *int32
	SetBody(v *CreateServiceCredentialResponseBody) *CreateServiceCredentialResponse
	GetBody() *CreateServiceCredentialResponseBody
}

type CreateServiceCredentialResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceCredentialResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateServiceCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateServiceCredentialResponse) GetBody() *CreateServiceCredentialResponseBody {
	return s.Body
}

func (s *CreateServiceCredentialResponse) SetHeaders(v map[string]*string) *CreateServiceCredentialResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceCredentialResponse) SetStatusCode(v int32) *CreateServiceCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceCredentialResponse) SetBody(v *CreateServiceCredentialResponseBody) *CreateServiceCredentialResponse {
	s.Body = v
	return s
}

func (s *CreateServiceCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
