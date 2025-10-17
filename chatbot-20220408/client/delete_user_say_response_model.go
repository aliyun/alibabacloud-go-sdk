// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserSayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUserSayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUserSayResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUserSayResponseBody) *DeleteUserSayResponse
	GetBody() *DeleteUserSayResponseBody
}

type DeleteUserSayResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserSayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserSayResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserSayResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserSayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUserSayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUserSayResponse) GetBody() *DeleteUserSayResponseBody {
	return s.Body
}

func (s *DeleteUserSayResponse) SetHeaders(v map[string]*string) *DeleteUserSayResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserSayResponse) SetStatusCode(v int32) *DeleteUserSayResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserSayResponse) SetBody(v *DeleteUserSayResponseBody) *DeleteUserSayResponse {
	s.Body = v
	return s
}

func (s *DeleteUserSayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
