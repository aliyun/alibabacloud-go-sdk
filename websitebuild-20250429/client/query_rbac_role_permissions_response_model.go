// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRbacRolePermissionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRbacRolePermissionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRbacRolePermissionsResponse
	GetStatusCode() *int32
	SetBody(v *QueryRbacRolePermissionsResponseBody) *QueryRbacRolePermissionsResponse
	GetBody() *QueryRbacRolePermissionsResponseBody
}

type QueryRbacRolePermissionsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRbacRolePermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRbacRolePermissionsResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRbacRolePermissionsResponse) GoString() string {
	return s.String()
}

func (s *QueryRbacRolePermissionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRbacRolePermissionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRbacRolePermissionsResponse) GetBody() *QueryRbacRolePermissionsResponseBody {
	return s.Body
}

func (s *QueryRbacRolePermissionsResponse) SetHeaders(v map[string]*string) *QueryRbacRolePermissionsResponse {
	s.Headers = v
	return s
}

func (s *QueryRbacRolePermissionsResponse) SetStatusCode(v int32) *QueryRbacRolePermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRbacRolePermissionsResponse) SetBody(v *QueryRbacRolePermissionsResponseBody) *QueryRbacRolePermissionsResponse {
	s.Body = v
	return s
}

func (s *QueryRbacRolePermissionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
