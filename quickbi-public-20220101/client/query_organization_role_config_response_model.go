// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOrganizationRoleConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryOrganizationRoleConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryOrganizationRoleConfigResponse
	GetStatusCode() *int32
	SetBody(v *QueryOrganizationRoleConfigResponseBody) *QueryOrganizationRoleConfigResponse
	GetBody() *QueryOrganizationRoleConfigResponseBody
}

type QueryOrganizationRoleConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOrganizationRoleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOrganizationRoleConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryOrganizationRoleConfigResponse) GoString() string {
	return s.String()
}

func (s *QueryOrganizationRoleConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryOrganizationRoleConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryOrganizationRoleConfigResponse) GetBody() *QueryOrganizationRoleConfigResponseBody {
	return s.Body
}

func (s *QueryOrganizationRoleConfigResponse) SetHeaders(v map[string]*string) *QueryOrganizationRoleConfigResponse {
	s.Headers = v
	return s
}

func (s *QueryOrganizationRoleConfigResponse) SetStatusCode(v int32) *QueryOrganizationRoleConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOrganizationRoleConfigResponse) SetBody(v *QueryOrganizationRoleConfigResponseBody) *QueryOrganizationRoleConfigResponse {
	s.Body = v
	return s
}

func (s *QueryOrganizationRoleConfigResponse) Validate() error {
	return dara.Validate(s)
}
