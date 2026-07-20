// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRbacRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRbacRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRbacRoleResponse
	GetStatusCode() *int32
	SetBody(v *QueryRbacRoleResponseBody) *QueryRbacRoleResponse
	GetBody() *QueryRbacRoleResponseBody
}

type QueryRbacRoleResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRbacRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRbacRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRbacRoleResponse) GoString() string {
	return s.String()
}

func (s *QueryRbacRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRbacRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRbacRoleResponse) GetBody() *QueryRbacRoleResponseBody {
	return s.Body
}

func (s *QueryRbacRoleResponse) SetHeaders(v map[string]*string) *QueryRbacRoleResponse {
	s.Headers = v
	return s
}

func (s *QueryRbacRoleResponse) SetStatusCode(v int32) *QueryRbacRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRbacRoleResponse) SetBody(v *QueryRbacRoleResponseBody) *QueryRbacRoleResponse {
	s.Body = v
	return s
}

func (s *QueryRbacRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
