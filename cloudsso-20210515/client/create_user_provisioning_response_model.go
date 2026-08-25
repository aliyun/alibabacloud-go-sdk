// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserProvisioningResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUserProvisioningResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUserProvisioningResponse
	GetStatusCode() *int32
	SetBody(v *CreateUserProvisioningResponseBody) *CreateUserProvisioningResponse
	GetBody() *CreateUserProvisioningResponseBody
}

type CreateUserProvisioningResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserProvisioningResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserProvisioningResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUserProvisioningResponse) GoString() string {
	return s.String()
}

func (s *CreateUserProvisioningResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUserProvisioningResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUserProvisioningResponse) GetBody() *CreateUserProvisioningResponseBody {
	return s.Body
}

func (s *CreateUserProvisioningResponse) SetHeaders(v map[string]*string) *CreateUserProvisioningResponse {
	s.Headers = v
	return s
}

func (s *CreateUserProvisioningResponse) SetStatusCode(v int32) *CreateUserProvisioningResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserProvisioningResponse) SetBody(v *CreateUserProvisioningResponseBody) *CreateUserProvisioningResponse {
	s.Body = v
	return s
}

func (s *CreateUserProvisioningResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
