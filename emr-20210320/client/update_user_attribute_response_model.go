// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUserAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUserAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUserAttributeResponseBody) *UpdateUserAttributeResponse
	GetBody() *UpdateUserAttributeResponseBody
}

type UpdateUserAttributeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUserAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUserAttributeResponse) GetBody() *UpdateUserAttributeResponseBody {
	return s.Body
}

func (s *UpdateUserAttributeResponse) SetHeaders(v map[string]*string) *UpdateUserAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserAttributeResponse) SetStatusCode(v int32) *UpdateUserAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserAttributeResponse) SetBody(v *UpdateUserAttributeResponseBody) *UpdateUserAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateUserAttributeResponse) Validate() error {
	return dara.Validate(s)
}
