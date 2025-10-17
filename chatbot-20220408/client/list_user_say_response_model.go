// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserSayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserSayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserSayResponse
	GetStatusCode() *int32
	SetBody(v *ListUserSayResponseBody) *ListUserSayResponse
	GetBody() *ListUserSayResponseBody
}

type ListUserSayResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserSayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserSayResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserSayResponse) GoString() string {
	return s.String()
}

func (s *ListUserSayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserSayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserSayResponse) GetBody() *ListUserSayResponseBody {
	return s.Body
}

func (s *ListUserSayResponse) SetHeaders(v map[string]*string) *ListUserSayResponse {
	s.Headers = v
	return s
}

func (s *ListUserSayResponse) SetStatusCode(v int32) *ListUserSayResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserSayResponse) SetBody(v *ListUserSayResponseBody) *ListUserSayResponse {
	s.Body = v
	return s
}

func (s *ListUserSayResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
