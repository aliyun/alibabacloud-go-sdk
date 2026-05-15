// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAllDepartmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAllDepartmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAllDepartmentResponse
	GetStatusCode() *int32
	SetBody(v *GetAllDepartmentResponseBody) *GetAllDepartmentResponse
	GetBody() *GetAllDepartmentResponseBody
}

type GetAllDepartmentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAllDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAllDepartmentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAllDepartmentResponse) GoString() string {
	return s.String()
}

func (s *GetAllDepartmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAllDepartmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAllDepartmentResponse) GetBody() *GetAllDepartmentResponseBody {
	return s.Body
}

func (s *GetAllDepartmentResponse) SetHeaders(v map[string]*string) *GetAllDepartmentResponse {
	s.Headers = v
	return s
}

func (s *GetAllDepartmentResponse) SetStatusCode(v int32) *GetAllDepartmentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAllDepartmentResponse) SetBody(v *GetAllDepartmentResponseBody) *GetAllDepartmentResponse {
	s.Body = v
	return s
}

func (s *GetAllDepartmentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
