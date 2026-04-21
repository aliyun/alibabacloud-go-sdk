// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetYikeUserRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetYikeUserRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetYikeUserRoleResponse
	GetStatusCode() *int32
	SetBody(v *SetYikeUserRoleResponseBody) *SetYikeUserRoleResponse
	GetBody() *SetYikeUserRoleResponseBody
}

type SetYikeUserRoleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetYikeUserRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetYikeUserRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s SetYikeUserRoleResponse) GoString() string {
	return s.String()
}

func (s *SetYikeUserRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetYikeUserRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetYikeUserRoleResponse) GetBody() *SetYikeUserRoleResponseBody {
	return s.Body
}

func (s *SetYikeUserRoleResponse) SetHeaders(v map[string]*string) *SetYikeUserRoleResponse {
	s.Headers = v
	return s
}

func (s *SetYikeUserRoleResponse) SetStatusCode(v int32) *SetYikeUserRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *SetYikeUserRoleResponse) SetBody(v *SetYikeUserRoleResponseBody) *SetYikeUserRoleResponse {
	s.Body = v
	return s
}

func (s *SetYikeUserRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
