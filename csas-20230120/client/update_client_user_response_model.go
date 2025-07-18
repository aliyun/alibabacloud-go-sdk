// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClientUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateClientUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateClientUserResponse
	GetStatusCode() *int32
	SetBody(v *UpdateClientUserResponseBody) *UpdateClientUserResponse
	GetBody() *UpdateClientUserResponseBody
}

type UpdateClientUserResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateClientUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateClientUserResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateClientUserResponse) GoString() string {
	return s.String()
}

func (s *UpdateClientUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateClientUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateClientUserResponse) GetBody() *UpdateClientUserResponseBody {
	return s.Body
}

func (s *UpdateClientUserResponse) SetHeaders(v map[string]*string) *UpdateClientUserResponse {
	s.Headers = v
	return s
}

func (s *UpdateClientUserResponse) SetStatusCode(v int32) *UpdateClientUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateClientUserResponse) SetBody(v *UpdateClientUserResponseBody) *UpdateClientUserResponse {
	s.Body = v
	return s
}

func (s *UpdateClientUserResponse) Validate() error {
	return dara.Validate(s)
}
