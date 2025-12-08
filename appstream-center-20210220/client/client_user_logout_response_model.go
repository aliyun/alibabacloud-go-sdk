// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClientUserLogoutResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClientUserLogoutResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClientUserLogoutResponse
	GetStatusCode() *int32
	SetBody(v *ClientUserLogoutResponseBody) *ClientUserLogoutResponse
	GetBody() *ClientUserLogoutResponseBody
}

type ClientUserLogoutResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClientUserLogoutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClientUserLogoutResponse) String() string {
	return dara.Prettify(s)
}

func (s ClientUserLogoutResponse) GoString() string {
	return s.String()
}

func (s *ClientUserLogoutResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClientUserLogoutResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClientUserLogoutResponse) GetBody() *ClientUserLogoutResponseBody {
	return s.Body
}

func (s *ClientUserLogoutResponse) SetHeaders(v map[string]*string) *ClientUserLogoutResponse {
	s.Headers = v
	return s
}

func (s *ClientUserLogoutResponse) SetStatusCode(v int32) *ClientUserLogoutResponse {
	s.StatusCode = &v
	return s
}

func (s *ClientUserLogoutResponse) SetBody(v *ClientUserLogoutResponseBody) *ClientUserLogoutResponse {
	s.Body = v
	return s
}

func (s *ClientUserLogoutResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
