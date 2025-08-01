// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeServiceLinkRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitializeServiceLinkRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitializeServiceLinkRoleResponse
	GetStatusCode() *int32
	SetBody(v *InitializeServiceLinkRoleResponseBody) *InitializeServiceLinkRoleResponse
	GetBody() *InitializeServiceLinkRoleResponseBody
}

type InitializeServiceLinkRoleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitializeServiceLinkRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitializeServiceLinkRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s InitializeServiceLinkRoleResponse) GoString() string {
	return s.String()
}

func (s *InitializeServiceLinkRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitializeServiceLinkRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitializeServiceLinkRoleResponse) GetBody() *InitializeServiceLinkRoleResponseBody {
	return s.Body
}

func (s *InitializeServiceLinkRoleResponse) SetHeaders(v map[string]*string) *InitializeServiceLinkRoleResponse {
	s.Headers = v
	return s
}

func (s *InitializeServiceLinkRoleResponse) SetStatusCode(v int32) *InitializeServiceLinkRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *InitializeServiceLinkRoleResponse) SetBody(v *InitializeServiceLinkRoleResponseBody) *InitializeServiceLinkRoleResponse {
	s.Body = v
	return s
}

func (s *InitializeServiceLinkRoleResponse) Validate() error {
	return dara.Validate(s)
}
