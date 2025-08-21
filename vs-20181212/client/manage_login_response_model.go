// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManageLoginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ManageLoginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ManageLoginResponse
	GetStatusCode() *int32
	SetBody(v *ManageLoginResponseBody) *ManageLoginResponse
	GetBody() *ManageLoginResponseBody
}

type ManageLoginResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ManageLoginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ManageLoginResponse) String() string {
	return dara.Prettify(s)
}

func (s ManageLoginResponse) GoString() string {
	return s.String()
}

func (s *ManageLoginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ManageLoginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ManageLoginResponse) GetBody() *ManageLoginResponseBody {
	return s.Body
}

func (s *ManageLoginResponse) SetHeaders(v map[string]*string) *ManageLoginResponse {
	s.Headers = v
	return s
}

func (s *ManageLoginResponse) SetStatusCode(v int32) *ManageLoginResponse {
	s.StatusCode = &v
	return s
}

func (s *ManageLoginResponse) SetBody(v *ManageLoginResponseBody) *ManageLoginResponse {
	s.Body = v
	return s
}

func (s *ManageLoginResponse) Validate() error {
	return dara.Validate(s)
}
