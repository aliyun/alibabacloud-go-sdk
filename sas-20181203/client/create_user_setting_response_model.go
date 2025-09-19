// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUserSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUserSettingResponse
	GetStatusCode() *int32
	SetBody(v *CreateUserSettingResponseBody) *CreateUserSettingResponse
	GetBody() *CreateUserSettingResponseBody
}

type CreateUserSettingResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUserSettingResponse) GoString() string {
	return s.String()
}

func (s *CreateUserSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUserSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUserSettingResponse) GetBody() *CreateUserSettingResponseBody {
	return s.Body
}

func (s *CreateUserSettingResponse) SetHeaders(v map[string]*string) *CreateUserSettingResponse {
	s.Headers = v
	return s
}

func (s *CreateUserSettingResponse) SetStatusCode(v int32) *CreateUserSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserSettingResponse) SetBody(v *CreateUserSettingResponseBody) *CreateUserSettingResponse {
	s.Body = v
	return s
}

func (s *CreateUserSettingResponse) Validate() error {
	return dara.Validate(s)
}
