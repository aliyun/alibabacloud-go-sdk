// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIdpDepartmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateIdpDepartmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateIdpDepartmentResponse
	GetStatusCode() *int32
	SetBody(v *UpdateIdpDepartmentResponseBody) *UpdateIdpDepartmentResponse
	GetBody() *UpdateIdpDepartmentResponseBody
}

type UpdateIdpDepartmentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIdpDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIdpDepartmentResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateIdpDepartmentResponse) GoString() string {
	return s.String()
}

func (s *UpdateIdpDepartmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateIdpDepartmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateIdpDepartmentResponse) GetBody() *UpdateIdpDepartmentResponseBody {
	return s.Body
}

func (s *UpdateIdpDepartmentResponse) SetHeaders(v map[string]*string) *UpdateIdpDepartmentResponse {
	s.Headers = v
	return s
}

func (s *UpdateIdpDepartmentResponse) SetStatusCode(v int32) *UpdateIdpDepartmentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIdpDepartmentResponse) SetBody(v *UpdateIdpDepartmentResponseBody) *UpdateIdpDepartmentResponse {
	s.Body = v
	return s
}

func (s *UpdateIdpDepartmentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
