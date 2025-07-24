// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoctorHBaseTablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDoctorHBaseTablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDoctorHBaseTablesResponse
	GetStatusCode() *int32
	SetBody(v *ListDoctorHBaseTablesResponseBody) *ListDoctorHBaseTablesResponse
	GetBody() *ListDoctorHBaseTablesResponseBody
}

type ListDoctorHBaseTablesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDoctorHBaseTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDoctorHBaseTablesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseTablesResponse) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseTablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDoctorHBaseTablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDoctorHBaseTablesResponse) GetBody() *ListDoctorHBaseTablesResponseBody {
	return s.Body
}

func (s *ListDoctorHBaseTablesResponse) SetHeaders(v map[string]*string) *ListDoctorHBaseTablesResponse {
	s.Headers = v
	return s
}

func (s *ListDoctorHBaseTablesResponse) SetStatusCode(v int32) *ListDoctorHBaseTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDoctorHBaseTablesResponse) SetBody(v *ListDoctorHBaseTablesResponseBody) *ListDoctorHBaseTablesResponse {
	s.Body = v
	return s
}

func (s *ListDoctorHBaseTablesResponse) Validate() error {
	return dara.Validate(s)
}
