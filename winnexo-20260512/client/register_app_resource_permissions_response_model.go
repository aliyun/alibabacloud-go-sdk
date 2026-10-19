// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterAppResourcePermissionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RegisterAppResourcePermissionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RegisterAppResourcePermissionsResponse
	GetStatusCode() *int32
	SetBody(v *RegisterAppResourcePermissionsResponseBody) *RegisterAppResourcePermissionsResponse
	GetBody() *RegisterAppResourcePermissionsResponseBody
}

type RegisterAppResourcePermissionsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterAppResourcePermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterAppResourcePermissionsResponse) String() string {
	return dara.Prettify(s)
}

func (s RegisterAppResourcePermissionsResponse) GoString() string {
	return s.String()
}

func (s *RegisterAppResourcePermissionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RegisterAppResourcePermissionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RegisterAppResourcePermissionsResponse) GetBody() *RegisterAppResourcePermissionsResponseBody {
	return s.Body
}

func (s *RegisterAppResourcePermissionsResponse) SetHeaders(v map[string]*string) *RegisterAppResourcePermissionsResponse {
	s.Headers = v
	return s
}

func (s *RegisterAppResourcePermissionsResponse) SetStatusCode(v int32) *RegisterAppResourcePermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterAppResourcePermissionsResponse) SetBody(v *RegisterAppResourcePermissionsResponseBody) *RegisterAppResourcePermissionsResponse {
	s.Body = v
	return s
}

func (s *RegisterAppResourcePermissionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
