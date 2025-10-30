// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIdpDepartmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIdpDepartmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIdpDepartmentResponse
	GetStatusCode() *int32
	SetBody(v *CreateIdpDepartmentResponseBody) *CreateIdpDepartmentResponse
	GetBody() *CreateIdpDepartmentResponseBody
}

type CreateIdpDepartmentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIdpDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIdpDepartmentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIdpDepartmentResponse) GoString() string {
	return s.String()
}

func (s *CreateIdpDepartmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIdpDepartmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIdpDepartmentResponse) GetBody() *CreateIdpDepartmentResponseBody {
	return s.Body
}

func (s *CreateIdpDepartmentResponse) SetHeaders(v map[string]*string) *CreateIdpDepartmentResponse {
	s.Headers = v
	return s
}

func (s *CreateIdpDepartmentResponse) SetStatusCode(v int32) *CreateIdpDepartmentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIdpDepartmentResponse) SetBody(v *CreateIdpDepartmentResponseBody) *CreateIdpDepartmentResponse {
	s.Body = v
	return s
}

func (s *CreateIdpDepartmentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
