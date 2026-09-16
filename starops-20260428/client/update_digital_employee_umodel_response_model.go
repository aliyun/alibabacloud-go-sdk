// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDigitalEmployeeUmodelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDigitalEmployeeUmodelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDigitalEmployeeUmodelResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDigitalEmployeeUmodelResponseBody) *UpdateDigitalEmployeeUmodelResponse
	GetBody() *UpdateDigitalEmployeeUmodelResponseBody
}

type UpdateDigitalEmployeeUmodelResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDigitalEmployeeUmodelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDigitalEmployeeUmodelResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDigitalEmployeeUmodelResponse) GoString() string {
	return s.String()
}

func (s *UpdateDigitalEmployeeUmodelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDigitalEmployeeUmodelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDigitalEmployeeUmodelResponse) GetBody() *UpdateDigitalEmployeeUmodelResponseBody {
	return s.Body
}

func (s *UpdateDigitalEmployeeUmodelResponse) SetHeaders(v map[string]*string) *UpdateDigitalEmployeeUmodelResponse {
	s.Headers = v
	return s
}

func (s *UpdateDigitalEmployeeUmodelResponse) SetStatusCode(v int32) *UpdateDigitalEmployeeUmodelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDigitalEmployeeUmodelResponse) SetBody(v *UpdateDigitalEmployeeUmodelResponseBody) *UpdateDigitalEmployeeUmodelResponse {
	s.Body = v
	return s
}

func (s *UpdateDigitalEmployeeUmodelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
