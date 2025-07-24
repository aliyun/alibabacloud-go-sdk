// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoctorHiveTablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDoctorHiveTablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDoctorHiveTablesResponse
	GetStatusCode() *int32
	SetBody(v *ListDoctorHiveTablesResponseBody) *ListDoctorHiveTablesResponse
	GetBody() *ListDoctorHiveTablesResponseBody
}

type ListDoctorHiveTablesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDoctorHiveTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDoctorHiveTablesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveTablesResponse) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveTablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDoctorHiveTablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDoctorHiveTablesResponse) GetBody() *ListDoctorHiveTablesResponseBody {
	return s.Body
}

func (s *ListDoctorHiveTablesResponse) SetHeaders(v map[string]*string) *ListDoctorHiveTablesResponse {
	s.Headers = v
	return s
}

func (s *ListDoctorHiveTablesResponse) SetStatusCode(v int32) *ListDoctorHiveTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDoctorHiveTablesResponse) SetBody(v *ListDoctorHiveTablesResponseBody) *ListDoctorHiveTablesResponse {
	s.Body = v
	return s
}

func (s *ListDoctorHiveTablesResponse) Validate() error {
	return dara.Validate(s)
}
