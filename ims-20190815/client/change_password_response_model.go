// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangePasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangePasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangePasswordResponse
	GetStatusCode() *int32
	SetBody(v *ChangePasswordResponseBody) *ChangePasswordResponse
	GetBody() *ChangePasswordResponseBody
}

type ChangePasswordResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangePasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangePasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangePasswordResponse) GoString() string {
	return s.String()
}

func (s *ChangePasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangePasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangePasswordResponse) GetBody() *ChangePasswordResponseBody {
	return s.Body
}

func (s *ChangePasswordResponse) SetHeaders(v map[string]*string) *ChangePasswordResponse {
	s.Headers = v
	return s
}

func (s *ChangePasswordResponse) SetStatusCode(v int32) *ChangePasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangePasswordResponse) SetBody(v *ChangePasswordResponseBody) *ChangePasswordResponse {
	s.Body = v
	return s
}

func (s *ChangePasswordResponse) Validate() error {
	return dara.Validate(s)
}
