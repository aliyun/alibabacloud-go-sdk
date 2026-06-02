// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserAuthConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUserAuthConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUserAuthConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUserAuthConfigResponseBody) *UpdateUserAuthConfigResponse
	GetBody() *UpdateUserAuthConfigResponseBody
}

type UpdateUserAuthConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserAuthConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserAuthConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserAuthConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserAuthConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUserAuthConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUserAuthConfigResponse) GetBody() *UpdateUserAuthConfigResponseBody {
	return s.Body
}

func (s *UpdateUserAuthConfigResponse) SetHeaders(v map[string]*string) *UpdateUserAuthConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserAuthConfigResponse) SetStatusCode(v int32) *UpdateUserAuthConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserAuthConfigResponse) SetBody(v *UpdateUserAuthConfigResponseBody) *UpdateUserAuthConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateUserAuthConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
