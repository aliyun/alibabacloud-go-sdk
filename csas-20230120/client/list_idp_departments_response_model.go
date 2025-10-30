// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIdpDepartmentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIdpDepartmentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIdpDepartmentsResponse
	GetStatusCode() *int32
	SetBody(v *ListIdpDepartmentsResponseBody) *ListIdpDepartmentsResponse
	GetBody() *ListIdpDepartmentsResponseBody
}

type ListIdpDepartmentsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIdpDepartmentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIdpDepartmentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIdpDepartmentsResponse) GoString() string {
	return s.String()
}

func (s *ListIdpDepartmentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIdpDepartmentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIdpDepartmentsResponse) GetBody() *ListIdpDepartmentsResponseBody {
	return s.Body
}

func (s *ListIdpDepartmentsResponse) SetHeaders(v map[string]*string) *ListIdpDepartmentsResponse {
	s.Headers = v
	return s
}

func (s *ListIdpDepartmentsResponse) SetStatusCode(v int32) *ListIdpDepartmentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIdpDepartmentsResponse) SetBody(v *ListIdpDepartmentsResponseBody) *ListIdpDepartmentsResponse {
	s.Body = v
	return s
}

func (s *ListIdpDepartmentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
