// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserAvatarResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUserAvatarResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUserAvatarResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUserAvatarResponseBody) *UpdateUserAvatarResponse
	GetBody() *UpdateUserAvatarResponseBody
}

type UpdateUserAvatarResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserAvatarResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserAvatarResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserAvatarResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserAvatarResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUserAvatarResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUserAvatarResponse) GetBody() *UpdateUserAvatarResponseBody {
	return s.Body
}

func (s *UpdateUserAvatarResponse) SetHeaders(v map[string]*string) *UpdateUserAvatarResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserAvatarResponse) SetStatusCode(v int32) *UpdateUserAvatarResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserAvatarResponse) SetBody(v *UpdateUserAvatarResponseBody) *UpdateUserAvatarResponse {
	s.Body = v
	return s
}

func (s *UpdateUserAvatarResponse) Validate() error {
	return dara.Validate(s)
}
