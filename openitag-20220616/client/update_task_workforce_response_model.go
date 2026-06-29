// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskWorkforceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTaskWorkforceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTaskWorkforceResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTaskWorkforceResponseBody) *UpdateTaskWorkforceResponse
	GetBody() *UpdateTaskWorkforceResponseBody
}

type UpdateTaskWorkforceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTaskWorkforceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTaskWorkforceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskWorkforceResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskWorkforceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTaskWorkforceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTaskWorkforceResponse) GetBody() *UpdateTaskWorkforceResponseBody {
	return s.Body
}

func (s *UpdateTaskWorkforceResponse) SetHeaders(v map[string]*string) *UpdateTaskWorkforceResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskWorkforceResponse) SetStatusCode(v int32) *UpdateTaskWorkforceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskWorkforceResponse) SetBody(v *UpdateTaskWorkforceResponseBody) *UpdateTaskWorkforceResponse {
	s.Body = v
	return s
}

func (s *UpdateTaskWorkforceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
