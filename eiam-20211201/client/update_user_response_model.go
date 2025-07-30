// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUserResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUserResponseBody) *UpdateUserResponse
	GetBody() *UpdateUserResponseBody
}

type UpdateUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUserResponse) GetBody() *UpdateUserResponseBody {
	return s.Body
}

func (s *UpdateUserResponse) SetHeaders(v map[string]*string) *UpdateUserResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserResponse) SetStatusCode(v int32) *UpdateUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserResponse) SetBody(v *UpdateUserResponseBody) *UpdateUserResponse {
	s.Body = v
	return s
}

func (s *UpdateUserResponse) Validate() error {
	return dara.Validate(s)
}
