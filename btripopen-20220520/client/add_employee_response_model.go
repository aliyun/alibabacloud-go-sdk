// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddEmployeeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddEmployeeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddEmployeeResponse
	GetStatusCode() *int32
	SetBody(v *AddEmployeeResponseBody) *AddEmployeeResponse
	GetBody() *AddEmployeeResponseBody
}

type AddEmployeeResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddEmployeeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddEmployeeResponse) String() string {
	return dara.Prettify(s)
}

func (s AddEmployeeResponse) GoString() string {
	return s.String()
}

func (s *AddEmployeeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddEmployeeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddEmployeeResponse) GetBody() *AddEmployeeResponseBody {
	return s.Body
}

func (s *AddEmployeeResponse) SetHeaders(v map[string]*string) *AddEmployeeResponse {
	s.Headers = v
	return s
}

func (s *AddEmployeeResponse) SetStatusCode(v int32) *AddEmployeeResponse {
	s.StatusCode = &v
	return s
}

func (s *AddEmployeeResponse) SetBody(v *AddEmployeeResponseBody) *AddEmployeeResponse {
	s.Body = v
	return s
}

func (s *AddEmployeeResponse) Validate() error {
	return dara.Validate(s)
}
