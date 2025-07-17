// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableUserResponse
	GetStatusCode() *int32
	SetBody(v *DisableUserResponseBody) *DisableUserResponse
	GetBody() *DisableUserResponseBody
}

type DisableUserResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableUserResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableUserResponse) GoString() string {
	return s.String()
}

func (s *DisableUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableUserResponse) GetBody() *DisableUserResponseBody {
	return s.Body
}

func (s *DisableUserResponse) SetHeaders(v map[string]*string) *DisableUserResponse {
	s.Headers = v
	return s
}

func (s *DisableUserResponse) SetStatusCode(v int32) *DisableUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableUserResponse) SetBody(v *DisableUserResponseBody) *DisableUserResponse {
	s.Body = v
	return s
}

func (s *DisableUserResponse) Validate() error {
	return dara.Validate(s)
}
