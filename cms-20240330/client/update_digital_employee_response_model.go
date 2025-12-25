// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDigitalEmployeeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDigitalEmployeeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDigitalEmployeeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDigitalEmployeeResponseBody) *UpdateDigitalEmployeeResponse
	GetBody() *UpdateDigitalEmployeeResponseBody
}

type UpdateDigitalEmployeeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDigitalEmployeeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDigitalEmployeeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDigitalEmployeeResponse) GoString() string {
	return s.String()
}

func (s *UpdateDigitalEmployeeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDigitalEmployeeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDigitalEmployeeResponse) GetBody() *UpdateDigitalEmployeeResponseBody {
	return s.Body
}

func (s *UpdateDigitalEmployeeResponse) SetHeaders(v map[string]*string) *UpdateDigitalEmployeeResponse {
	s.Headers = v
	return s
}

func (s *UpdateDigitalEmployeeResponse) SetStatusCode(v int32) *UpdateDigitalEmployeeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDigitalEmployeeResponse) SetBody(v *UpdateDigitalEmployeeResponseBody) *UpdateDigitalEmployeeResponse {
	s.Body = v
	return s
}

func (s *UpdateDigitalEmployeeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
