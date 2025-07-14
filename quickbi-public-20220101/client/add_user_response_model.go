// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddUserResponse
	GetStatusCode() *int32
	SetBody(v *AddUserResponseBody) *AddUserResponse
	GetBody() *AddUserResponseBody
}

type AddUserResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddUserResponse) String() string {
	return dara.Prettify(s)
}

func (s AddUserResponse) GoString() string {
	return s.String()
}

func (s *AddUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddUserResponse) GetBody() *AddUserResponseBody {
	return s.Body
}

func (s *AddUserResponse) SetHeaders(v map[string]*string) *AddUserResponse {
	s.Headers = v
	return s
}

func (s *AddUserResponse) SetStatusCode(v int32) *AddUserResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUserResponse) SetBody(v *AddUserResponseBody) *AddUserResponse {
	s.Body = v
	return s
}

func (s *AddUserResponse) Validate() error {
	return dara.Validate(s)
}
