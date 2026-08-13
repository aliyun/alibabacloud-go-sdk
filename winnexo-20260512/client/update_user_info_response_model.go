// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUserInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUserInfoResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUserInfoResponseBody) *UpdateUserInfoResponse
	GetBody() *UpdateUserInfoResponseBody
}

type UpdateUserInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUserInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUserInfoResponse) GetBody() *UpdateUserInfoResponseBody {
	return s.Body
}

func (s *UpdateUserInfoResponse) SetHeaders(v map[string]*string) *UpdateUserInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserInfoResponse) SetStatusCode(v int32) *UpdateUserInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserInfoResponse) SetBody(v *UpdateUserInfoResponseBody) *UpdateUserInfoResponse {
	s.Body = v
	return s
}

func (s *UpdateUserInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
