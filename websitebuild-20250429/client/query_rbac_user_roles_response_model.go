// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRbacUserRolesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryRbacUserRolesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryRbacUserRolesResponse
	GetStatusCode() *int32
	SetBody(v *QueryRbacUserRolesResponseBody) *QueryRbacUserRolesResponse
	GetBody() *QueryRbacUserRolesResponseBody
}

type QueryRbacUserRolesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRbacUserRolesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryRbacUserRolesResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryRbacUserRolesResponse) GoString() string {
	return s.String()
}

func (s *QueryRbacUserRolesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryRbacUserRolesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryRbacUserRolesResponse) GetBody() *QueryRbacUserRolesResponseBody {
	return s.Body
}

func (s *QueryRbacUserRolesResponse) SetHeaders(v map[string]*string) *QueryRbacUserRolesResponse {
	s.Headers = v
	return s
}

func (s *QueryRbacUserRolesResponse) SetStatusCode(v int32) *QueryRbacUserRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRbacUserRolesResponse) SetBody(v *QueryRbacUserRolesResponseBody) *QueryRbacUserRolesResponse {
	s.Body = v
	return s
}

func (s *QueryRbacUserRolesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
