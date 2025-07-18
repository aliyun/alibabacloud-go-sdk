// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindDBResourceGroupWithRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindDBResourceGroupWithRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindDBResourceGroupWithRoleResponse
	GetStatusCode() *int32
	SetBody(v *BindDBResourceGroupWithRoleResponseBody) *BindDBResourceGroupWithRoleResponse
	GetBody() *BindDBResourceGroupWithRoleResponseBody
}

type BindDBResourceGroupWithRoleResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindDBResourceGroupWithRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindDBResourceGroupWithRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s BindDBResourceGroupWithRoleResponse) GoString() string {
	return s.String()
}

func (s *BindDBResourceGroupWithRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindDBResourceGroupWithRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindDBResourceGroupWithRoleResponse) GetBody() *BindDBResourceGroupWithRoleResponseBody {
	return s.Body
}

func (s *BindDBResourceGroupWithRoleResponse) SetHeaders(v map[string]*string) *BindDBResourceGroupWithRoleResponse {
	s.Headers = v
	return s
}

func (s *BindDBResourceGroupWithRoleResponse) SetStatusCode(v int32) *BindDBResourceGroupWithRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *BindDBResourceGroupWithRoleResponse) SetBody(v *BindDBResourceGroupWithRoleResponseBody) *BindDBResourceGroupWithRoleResponse {
	s.Body = v
	return s
}

func (s *BindDBResourceGroupWithRoleResponse) Validate() error {
	return dara.Validate(s)
}
