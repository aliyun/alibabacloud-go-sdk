// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceLinkedRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteServiceLinkedRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteServiceLinkedRoleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteServiceLinkedRoleResponseBody) *DeleteServiceLinkedRoleResponse
	GetBody() *DeleteServiceLinkedRoleResponseBody
}

type DeleteServiceLinkedRoleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceLinkedRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceLinkedRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceLinkedRoleResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceLinkedRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteServiceLinkedRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteServiceLinkedRoleResponse) GetBody() *DeleteServiceLinkedRoleResponseBody {
	return s.Body
}

func (s *DeleteServiceLinkedRoleResponse) SetHeaders(v map[string]*string) *DeleteServiceLinkedRoleResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceLinkedRoleResponse) SetStatusCode(v int32) *DeleteServiceLinkedRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceLinkedRoleResponse) SetBody(v *DeleteServiceLinkedRoleResponseBody) *DeleteServiceLinkedRoleResponse {
	s.Body = v
	return s
}

func (s *DeleteServiceLinkedRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
