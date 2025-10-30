// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindDBResourceGroupWithRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindDBResourceGroupWithRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindDBResourceGroupWithRoleResponse
	GetStatusCode() *int32
	SetBody(v *UnbindDBResourceGroupWithRoleResponseBody) *UnbindDBResourceGroupWithRoleResponse
	GetBody() *UnbindDBResourceGroupWithRoleResponseBody
}

type UnbindDBResourceGroupWithRoleResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindDBResourceGroupWithRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindDBResourceGroupWithRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindDBResourceGroupWithRoleResponse) GoString() string {
	return s.String()
}

func (s *UnbindDBResourceGroupWithRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindDBResourceGroupWithRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindDBResourceGroupWithRoleResponse) GetBody() *UnbindDBResourceGroupWithRoleResponseBody {
	return s.Body
}

func (s *UnbindDBResourceGroupWithRoleResponse) SetHeaders(v map[string]*string) *UnbindDBResourceGroupWithRoleResponse {
	s.Headers = v
	return s
}

func (s *UnbindDBResourceGroupWithRoleResponse) SetStatusCode(v int32) *UnbindDBResourceGroupWithRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindDBResourceGroupWithRoleResponse) SetBody(v *UnbindDBResourceGroupWithRoleResponseBody) *UnbindDBResourceGroupWithRoleResponse {
	s.Body = v
	return s
}

func (s *UnbindDBResourceGroupWithRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
