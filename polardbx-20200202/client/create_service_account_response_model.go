// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateServiceAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateServiceAccountResponse
	GetStatusCode() *int32
	SetBody(v *CreateServiceAccountResponseBody) *CreateServiceAccountResponse
	GetBody() *CreateServiceAccountResponseBody
}

type CreateServiceAccountResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateServiceAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateServiceAccountResponse) GetBody() *CreateServiceAccountResponseBody {
	return s.Body
}

func (s *CreateServiceAccountResponse) SetHeaders(v map[string]*string) *CreateServiceAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceAccountResponse) SetStatusCode(v int32) *CreateServiceAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceAccountResponse) SetBody(v *CreateServiceAccountResponseBody) *CreateServiceAccountResponse {
	s.Body = v
	return s
}

func (s *CreateServiceAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
