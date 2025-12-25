// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDigitalEmployeesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDigitalEmployeesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDigitalEmployeesResponse
	GetStatusCode() *int32
	SetBody(v *ListDigitalEmployeesResponseBody) *ListDigitalEmployeesResponse
	GetBody() *ListDigitalEmployeesResponseBody
}

type ListDigitalEmployeesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDigitalEmployeesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDigitalEmployeesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDigitalEmployeesResponse) GoString() string {
	return s.String()
}

func (s *ListDigitalEmployeesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDigitalEmployeesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDigitalEmployeesResponse) GetBody() *ListDigitalEmployeesResponseBody {
	return s.Body
}

func (s *ListDigitalEmployeesResponse) SetHeaders(v map[string]*string) *ListDigitalEmployeesResponse {
	s.Headers = v
	return s
}

func (s *ListDigitalEmployeesResponse) SetStatusCode(v int32) *ListDigitalEmployeesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDigitalEmployeesResponse) SetBody(v *ListDigitalEmployeesResponseBody) *ListDigitalEmployeesResponse {
	s.Body = v
	return s
}

func (s *ListDigitalEmployeesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
