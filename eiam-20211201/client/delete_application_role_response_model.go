// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteApplicationRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteApplicationRoleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteApplicationRoleResponseBody) *DeleteApplicationRoleResponse
	GetBody() *DeleteApplicationRoleResponseBody
}

type DeleteApplicationRoleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApplicationRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApplicationRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationRoleResponse) GoString() string {
	return s.String()
}

func (s *DeleteApplicationRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteApplicationRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteApplicationRoleResponse) GetBody() *DeleteApplicationRoleResponseBody {
	return s.Body
}

func (s *DeleteApplicationRoleResponse) SetHeaders(v map[string]*string) *DeleteApplicationRoleResponse {
	s.Headers = v
	return s
}

func (s *DeleteApplicationRoleResponse) SetStatusCode(v int32) *DeleteApplicationRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApplicationRoleResponse) SetBody(v *DeleteApplicationRoleResponseBody) *DeleteApplicationRoleResponse {
	s.Body = v
	return s
}

func (s *DeleteApplicationRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
