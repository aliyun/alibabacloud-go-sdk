// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUserStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUserStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUserStatusResponseBody) *UpdateUserStatusResponse
	GetBody() *UpdateUserStatusResponseBody
}

type UpdateUserStatusResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUserStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUserStatusResponse) GetBody() *UpdateUserStatusResponseBody {
	return s.Body
}

func (s *UpdateUserStatusResponse) SetHeaders(v map[string]*string) *UpdateUserStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserStatusResponse) SetStatusCode(v int32) *UpdateUserStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserStatusResponse) SetBody(v *UpdateUserStatusResponseBody) *UpdateUserStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateUserStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
