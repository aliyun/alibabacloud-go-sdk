// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityGroupPermissionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSecurityGroupPermissionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSecurityGroupPermissionsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSecurityGroupPermissionsResponseBody) *DeleteSecurityGroupPermissionsResponse
	GetBody() *DeleteSecurityGroupPermissionsResponseBody
}

type DeleteSecurityGroupPermissionsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSecurityGroupPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSecurityGroupPermissionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityGroupPermissionsResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecurityGroupPermissionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSecurityGroupPermissionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSecurityGroupPermissionsResponse) GetBody() *DeleteSecurityGroupPermissionsResponseBody {
	return s.Body
}

func (s *DeleteSecurityGroupPermissionsResponse) SetHeaders(v map[string]*string) *DeleteSecurityGroupPermissionsResponse {
	s.Headers = v
	return s
}

func (s *DeleteSecurityGroupPermissionsResponse) SetStatusCode(v int32) *DeleteSecurityGroupPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSecurityGroupPermissionsResponse) SetBody(v *DeleteSecurityGroupPermissionsResponseBody) *DeleteSecurityGroupPermissionsResponse {
	s.Body = v
	return s
}

func (s *DeleteSecurityGroupPermissionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
