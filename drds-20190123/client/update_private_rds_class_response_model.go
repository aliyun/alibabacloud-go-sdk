// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrivateRdsClassResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePrivateRdsClassResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePrivateRdsClassResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePrivateRdsClassResponseBody) *UpdatePrivateRdsClassResponse
	GetBody() *UpdatePrivateRdsClassResponseBody
}

type UpdatePrivateRdsClassResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePrivateRdsClassResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePrivateRdsClassResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateRdsClassResponse) GoString() string {
	return s.String()
}

func (s *UpdatePrivateRdsClassResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePrivateRdsClassResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePrivateRdsClassResponse) GetBody() *UpdatePrivateRdsClassResponseBody {
	return s.Body
}

func (s *UpdatePrivateRdsClassResponse) SetHeaders(v map[string]*string) *UpdatePrivateRdsClassResponse {
	s.Headers = v
	return s
}

func (s *UpdatePrivateRdsClassResponse) SetStatusCode(v int32) *UpdatePrivateRdsClassResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePrivateRdsClassResponse) SetBody(v *UpdatePrivateRdsClassResponseBody) *UpdatePrivateRdsClassResponse {
	s.Body = v
	return s
}

func (s *UpdatePrivateRdsClassResponse) Validate() error {
	return dara.Validate(s)
}
