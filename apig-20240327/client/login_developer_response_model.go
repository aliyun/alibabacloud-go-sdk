// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoginDeveloperResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LoginDeveloperResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LoginDeveloperResponse
	GetStatusCode() *int32
	SetBody(v *LoginDeveloperResponseBody) *LoginDeveloperResponse
	GetBody() *LoginDeveloperResponseBody
}

type LoginDeveloperResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LoginDeveloperResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LoginDeveloperResponse) String() string {
	return dara.Prettify(s)
}

func (s LoginDeveloperResponse) GoString() string {
	return s.String()
}

func (s *LoginDeveloperResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LoginDeveloperResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LoginDeveloperResponse) GetBody() *LoginDeveloperResponseBody {
	return s.Body
}

func (s *LoginDeveloperResponse) SetHeaders(v map[string]*string) *LoginDeveloperResponse {
	s.Headers = v
	return s
}

func (s *LoginDeveloperResponse) SetStatusCode(v int32) *LoginDeveloperResponse {
	s.StatusCode = &v
	return s
}

func (s *LoginDeveloperResponse) SetBody(v *LoginDeveloperResponseBody) *LoginDeveloperResponse {
	s.Body = v
	return s
}

func (s *LoginDeveloperResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
