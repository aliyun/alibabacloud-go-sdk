// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableCloudAccountRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableCloudAccountRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableCloudAccountRoleResponse
	GetStatusCode() *int32
	SetBody(v *DisableCloudAccountRoleResponseBody) *DisableCloudAccountRoleResponse
	GetBody() *DisableCloudAccountRoleResponseBody
}

type DisableCloudAccountRoleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableCloudAccountRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableCloudAccountRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableCloudAccountRoleResponse) GoString() string {
	return s.String()
}

func (s *DisableCloudAccountRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableCloudAccountRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableCloudAccountRoleResponse) GetBody() *DisableCloudAccountRoleResponseBody {
	return s.Body
}

func (s *DisableCloudAccountRoleResponse) SetHeaders(v map[string]*string) *DisableCloudAccountRoleResponse {
	s.Headers = v
	return s
}

func (s *DisableCloudAccountRoleResponse) SetStatusCode(v int32) *DisableCloudAccountRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableCloudAccountRoleResponse) SetBody(v *DisableCloudAccountRoleResponseBody) *DisableCloudAccountRoleResponse {
	s.Body = v
	return s
}

func (s *DisableCloudAccountRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
