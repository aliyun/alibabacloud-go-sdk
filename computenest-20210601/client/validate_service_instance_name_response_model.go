// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateServiceInstanceNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValidateServiceInstanceNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValidateServiceInstanceNameResponse
	GetStatusCode() *int32
	SetBody(v *ValidateServiceInstanceNameResponseBody) *ValidateServiceInstanceNameResponse
	GetBody() *ValidateServiceInstanceNameResponseBody
}

type ValidateServiceInstanceNameResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateServiceInstanceNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateServiceInstanceNameResponse) String() string {
	return dara.Prettify(s)
}

func (s ValidateServiceInstanceNameResponse) GoString() string {
	return s.String()
}

func (s *ValidateServiceInstanceNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValidateServiceInstanceNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValidateServiceInstanceNameResponse) GetBody() *ValidateServiceInstanceNameResponseBody {
	return s.Body
}

func (s *ValidateServiceInstanceNameResponse) SetHeaders(v map[string]*string) *ValidateServiceInstanceNameResponse {
	s.Headers = v
	return s
}

func (s *ValidateServiceInstanceNameResponse) SetStatusCode(v int32) *ValidateServiceInstanceNameResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateServiceInstanceNameResponse) SetBody(v *ValidateServiceInstanceNameResponseBody) *ValidateServiceInstanceNameResponse {
	s.Body = v
	return s
}

func (s *ValidateServiceInstanceNameResponse) Validate() error {
	return dara.Validate(s)
}
