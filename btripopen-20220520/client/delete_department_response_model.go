// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDepartmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDepartmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDepartmentResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDepartmentResponseBody) *DeleteDepartmentResponse
	GetBody() *DeleteDepartmentResponseBody
}

type DeleteDepartmentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDepartmentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDepartmentResponse) GoString() string {
	return s.String()
}

func (s *DeleteDepartmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDepartmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDepartmentResponse) GetBody() *DeleteDepartmentResponseBody {
	return s.Body
}

func (s *DeleteDepartmentResponse) SetHeaders(v map[string]*string) *DeleteDepartmentResponse {
	s.Headers = v
	return s
}

func (s *DeleteDepartmentResponse) SetStatusCode(v int32) *DeleteDepartmentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDepartmentResponse) SetBody(v *DeleteDepartmentResponseBody) *DeleteDepartmentResponse {
	s.Body = v
	return s
}

func (s *DeleteDepartmentResponse) Validate() error {
	return dara.Validate(s)
}
