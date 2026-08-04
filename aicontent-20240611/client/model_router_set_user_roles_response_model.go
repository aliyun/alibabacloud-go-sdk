// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterSetUserRolesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterSetUserRolesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterSetUserRolesResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterSetUserRolesResponseBody) *ModelRouterSetUserRolesResponse
	GetBody() *ModelRouterSetUserRolesResponseBody
}

type ModelRouterSetUserRolesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterSetUserRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterSetUserRolesResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterSetUserRolesResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterSetUserRolesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterSetUserRolesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterSetUserRolesResponse) GetBody() *ModelRouterSetUserRolesResponseBody {
	return s.Body
}

func (s *ModelRouterSetUserRolesResponse) SetHeaders(v map[string]*string) *ModelRouterSetUserRolesResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterSetUserRolesResponse) SetStatusCode(v int32) *ModelRouterSetUserRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterSetUserRolesResponse) SetBody(v *ModelRouterSetUserRolesResponseBody) *ModelRouterSetUserRolesResponse {
	s.Body = v
	return s
}

func (s *ModelRouterSetUserRolesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
