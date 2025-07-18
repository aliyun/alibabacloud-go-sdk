// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIdpDepartmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIdpDepartmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIdpDepartmentResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIdpDepartmentResponseBody) *DeleteIdpDepartmentResponse
	GetBody() *DeleteIdpDepartmentResponseBody
}

type DeleteIdpDepartmentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIdpDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIdpDepartmentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIdpDepartmentResponse) GoString() string {
	return s.String()
}

func (s *DeleteIdpDepartmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIdpDepartmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIdpDepartmentResponse) GetBody() *DeleteIdpDepartmentResponseBody {
	return s.Body
}

func (s *DeleteIdpDepartmentResponse) SetHeaders(v map[string]*string) *DeleteIdpDepartmentResponse {
	s.Headers = v
	return s
}

func (s *DeleteIdpDepartmentResponse) SetStatusCode(v int32) *DeleteIdpDepartmentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIdpDepartmentResponse) SetBody(v *DeleteIdpDepartmentResponseBody) *DeleteIdpDepartmentResponse {
	s.Body = v
	return s
}

func (s *DeleteIdpDepartmentResponse) Validate() error {
	return dara.Validate(s)
}
