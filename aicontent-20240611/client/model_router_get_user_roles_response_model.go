// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterGetUserRolesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelRouterGetUserRolesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelRouterGetUserRolesResponse
	GetStatusCode() *int32
	SetBody(v *ModelRouterGetUserRolesResponseBody) *ModelRouterGetUserRolesResponse
	GetBody() *ModelRouterGetUserRolesResponseBody
}

type ModelRouterGetUserRolesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelRouterGetUserRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterGetUserRolesResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterGetUserRolesResponse) GoString() string {
	return s.String()
}

func (s *ModelRouterGetUserRolesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelRouterGetUserRolesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelRouterGetUserRolesResponse) GetBody() *ModelRouterGetUserRolesResponseBody {
	return s.Body
}

func (s *ModelRouterGetUserRolesResponse) SetHeaders(v map[string]*string) *ModelRouterGetUserRolesResponse {
	s.Headers = v
	return s
}

func (s *ModelRouterGetUserRolesResponse) SetStatusCode(v int32) *ModelRouterGetUserRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelRouterGetUserRolesResponse) SetBody(v *ModelRouterGetUserRolesResponseBody) *ModelRouterGetUserRolesResponse {
	s.Body = v
	return s
}

func (s *ModelRouterGetUserRolesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
