// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDepartmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDepartmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDepartmentResponse
	GetStatusCode() *int32
	SetBody(v *AddDepartmentResponseBody) *AddDepartmentResponse
	GetBody() *AddDepartmentResponseBody
}

type AddDepartmentResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDepartmentResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDepartmentResponse) GoString() string {
	return s.String()
}

func (s *AddDepartmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDepartmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDepartmentResponse) GetBody() *AddDepartmentResponseBody {
	return s.Body
}

func (s *AddDepartmentResponse) SetHeaders(v map[string]*string) *AddDepartmentResponse {
	s.Headers = v
	return s
}

func (s *AddDepartmentResponse) SetStatusCode(v int32) *AddDepartmentResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDepartmentResponse) SetBody(v *AddDepartmentResponseBody) *AddDepartmentResponse {
	s.Body = v
	return s
}

func (s *AddDepartmentResponse) Validate() error {
	return dara.Validate(s)
}
