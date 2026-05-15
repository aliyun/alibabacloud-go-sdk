// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDepartmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDepartmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDepartmentResponse
	GetStatusCode() *int32
	SetBody(v *CreateDepartmentResponseBody) *CreateDepartmentResponse
	GetBody() *CreateDepartmentResponseBody
}

type CreateDepartmentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDepartmentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDepartmentResponse) GoString() string {
	return s.String()
}

func (s *CreateDepartmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDepartmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDepartmentResponse) GetBody() *CreateDepartmentResponseBody {
	return s.Body
}

func (s *CreateDepartmentResponse) SetHeaders(v map[string]*string) *CreateDepartmentResponse {
	s.Headers = v
	return s
}

func (s *CreateDepartmentResponse) SetStatusCode(v int32) *CreateDepartmentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDepartmentResponse) SetBody(v *CreateDepartmentResponseBody) *CreateDepartmentResponse {
	s.Body = v
	return s
}

func (s *CreateDepartmentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
