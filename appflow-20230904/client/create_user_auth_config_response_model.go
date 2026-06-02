// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserAuthConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUserAuthConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUserAuthConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateUserAuthConfigResponseBody) *CreateUserAuthConfigResponse
	GetBody() *CreateUserAuthConfigResponseBody
}

type CreateUserAuthConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserAuthConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserAuthConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUserAuthConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateUserAuthConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUserAuthConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUserAuthConfigResponse) GetBody() *CreateUserAuthConfigResponseBody {
	return s.Body
}

func (s *CreateUserAuthConfigResponse) SetHeaders(v map[string]*string) *CreateUserAuthConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateUserAuthConfigResponse) SetStatusCode(v int32) *CreateUserAuthConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserAuthConfigResponse) SetBody(v *CreateUserAuthConfigResponseBody) *CreateUserAuthConfigResponse {
	s.Body = v
	return s
}

func (s *CreateUserAuthConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
