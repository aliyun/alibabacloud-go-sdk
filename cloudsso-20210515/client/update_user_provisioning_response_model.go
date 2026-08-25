// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserProvisioningResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUserProvisioningResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUserProvisioningResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUserProvisioningResponseBody) *UpdateUserProvisioningResponse
	GetBody() *UpdateUserProvisioningResponseBody
}

type UpdateUserProvisioningResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserProvisioningResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserProvisioningResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserProvisioningResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserProvisioningResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUserProvisioningResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUserProvisioningResponse) GetBody() *UpdateUserProvisioningResponseBody {
	return s.Body
}

func (s *UpdateUserProvisioningResponse) SetHeaders(v map[string]*string) *UpdateUserProvisioningResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserProvisioningResponse) SetStatusCode(v int32) *UpdateUserProvisioningResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserProvisioningResponse) SetBody(v *UpdateUserProvisioningResponseBody) *UpdateUserProvisioningResponse {
	s.Body = v
	return s
}

func (s *UpdateUserProvisioningResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
