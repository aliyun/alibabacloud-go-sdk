// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAdminPasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAdminPasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAdminPasswordResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAdminPasswordResponseBody) *UpdateAdminPasswordResponse
	GetBody() *UpdateAdminPasswordResponseBody
}

type UpdateAdminPasswordResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAdminPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAdminPasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAdminPasswordResponse) GoString() string {
	return s.String()
}

func (s *UpdateAdminPasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAdminPasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAdminPasswordResponse) GetBody() *UpdateAdminPasswordResponseBody {
	return s.Body
}

func (s *UpdateAdminPasswordResponse) SetHeaders(v map[string]*string) *UpdateAdminPasswordResponse {
	s.Headers = v
	return s
}

func (s *UpdateAdminPasswordResponse) SetStatusCode(v int32) *UpdateAdminPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAdminPasswordResponse) SetBody(v *UpdateAdminPasswordResponseBody) *UpdateAdminPasswordResponse {
	s.Body = v
	return s
}

func (s *UpdateAdminPasswordResponse) Validate() error {
	return dara.Validate(s)
}
