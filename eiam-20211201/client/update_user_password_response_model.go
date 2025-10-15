// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserPasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUserPasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUserPasswordResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUserPasswordResponseBody) *UpdateUserPasswordResponse
	GetBody() *UpdateUserPasswordResponseBody
}

type UpdateUserPasswordResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserPasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserPasswordResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserPasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUserPasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUserPasswordResponse) GetBody() *UpdateUserPasswordResponseBody {
	return s.Body
}

func (s *UpdateUserPasswordResponse) SetHeaders(v map[string]*string) *UpdateUserPasswordResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserPasswordResponse) SetStatusCode(v int32) *UpdateUserPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserPasswordResponse) SetBody(v *UpdateUserPasswordResponseBody) *UpdateUserPasswordResponse {
	s.Body = v
	return s
}

func (s *UpdateUserPasswordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
