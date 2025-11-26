// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoctorHiveDatabasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDoctorHiveDatabasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDoctorHiveDatabasesResponse
	GetStatusCode() *int32
	SetBody(v *ListDoctorHiveDatabasesResponseBody) *ListDoctorHiveDatabasesResponse
	GetBody() *ListDoctorHiveDatabasesResponseBody
}

type ListDoctorHiveDatabasesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDoctorHiveDatabasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDoctorHiveDatabasesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesResponse) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDoctorHiveDatabasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDoctorHiveDatabasesResponse) GetBody() *ListDoctorHiveDatabasesResponseBody {
	return s.Body
}

func (s *ListDoctorHiveDatabasesResponse) SetHeaders(v map[string]*string) *ListDoctorHiveDatabasesResponse {
	s.Headers = v
	return s
}

func (s *ListDoctorHiveDatabasesResponse) SetStatusCode(v int32) *ListDoctorHiveDatabasesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDoctorHiveDatabasesResponse) SetBody(v *ListDoctorHiveDatabasesResponseBody) *ListDoctorHiveDatabasesResponse {
	s.Body = v
	return s
}

func (s *ListDoctorHiveDatabasesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
