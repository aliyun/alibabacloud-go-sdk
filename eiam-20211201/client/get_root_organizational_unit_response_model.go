// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRootOrganizationalUnitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRootOrganizationalUnitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRootOrganizationalUnitResponse
	GetStatusCode() *int32
	SetBody(v *GetRootOrganizationalUnitResponseBody) *GetRootOrganizationalUnitResponse
	GetBody() *GetRootOrganizationalUnitResponseBody
}

type GetRootOrganizationalUnitResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRootOrganizationalUnitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRootOrganizationalUnitResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRootOrganizationalUnitResponse) GoString() string {
	return s.String()
}

func (s *GetRootOrganizationalUnitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRootOrganizationalUnitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRootOrganizationalUnitResponse) GetBody() *GetRootOrganizationalUnitResponseBody {
	return s.Body
}

func (s *GetRootOrganizationalUnitResponse) SetHeaders(v map[string]*string) *GetRootOrganizationalUnitResponse {
	s.Headers = v
	return s
}

func (s *GetRootOrganizationalUnitResponse) SetStatusCode(v int32) *GetRootOrganizationalUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRootOrganizationalUnitResponse) SetBody(v *GetRootOrganizationalUnitResponseBody) *GetRootOrganizationalUnitResponse {
	s.Body = v
	return s
}

func (s *GetRootOrganizationalUnitResponse) Validate() error {
	return dara.Validate(s)
}
