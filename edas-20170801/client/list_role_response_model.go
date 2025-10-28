// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRoleResponse
	GetStatusCode() *int32
	SetBody(v *ListRoleResponseBody) *ListRoleResponse
	GetBody() *ListRoleResponseBody
}

type ListRoleResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRoleResponse) GoString() string {
	return s.String()
}

func (s *ListRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRoleResponse) GetBody() *ListRoleResponseBody {
	return s.Body
}

func (s *ListRoleResponse) SetHeaders(v map[string]*string) *ListRoleResponse {
	s.Headers = v
	return s
}

func (s *ListRoleResponse) SetStatusCode(v int32) *ListRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRoleResponse) SetBody(v *ListRoleResponseBody) *ListRoleResponse {
	s.Body = v
	return s
}

func (s *ListRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
