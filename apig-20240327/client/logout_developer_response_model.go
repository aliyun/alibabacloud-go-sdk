// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLogoutDeveloperResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LogoutDeveloperResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LogoutDeveloperResponse
	GetStatusCode() *int32
	SetBody(v *LogoutDeveloperResponseBody) *LogoutDeveloperResponse
	GetBody() *LogoutDeveloperResponseBody
}

type LogoutDeveloperResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LogoutDeveloperResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LogoutDeveloperResponse) String() string {
	return dara.Prettify(s)
}

func (s LogoutDeveloperResponse) GoString() string {
	return s.String()
}

func (s *LogoutDeveloperResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LogoutDeveloperResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LogoutDeveloperResponse) GetBody() *LogoutDeveloperResponseBody {
	return s.Body
}

func (s *LogoutDeveloperResponse) SetHeaders(v map[string]*string) *LogoutDeveloperResponse {
	s.Headers = v
	return s
}

func (s *LogoutDeveloperResponse) SetStatusCode(v int32) *LogoutDeveloperResponse {
	s.StatusCode = &v
	return s
}

func (s *LogoutDeveloperResponse) SetBody(v *LogoutDeveloperResponseBody) *LogoutDeveloperResponse {
	s.Body = v
	return s
}

func (s *LogoutDeveloperResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
