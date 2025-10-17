// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserSayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUserSayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUserSayResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUserSayResponseBody) *UpdateUserSayResponse
	GetBody() *UpdateUserSayResponseBody
}

type UpdateUserSayResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserSayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserSayResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserSayResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserSayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUserSayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUserSayResponse) GetBody() *UpdateUserSayResponseBody {
	return s.Body
}

func (s *UpdateUserSayResponse) SetHeaders(v map[string]*string) *UpdateUserSayResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserSayResponse) SetStatusCode(v int32) *UpdateUserSayResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserSayResponse) SetBody(v *UpdateUserSayResponseBody) *UpdateUserSayResponse {
	s.Body = v
	return s
}

func (s *UpdateUserSayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
