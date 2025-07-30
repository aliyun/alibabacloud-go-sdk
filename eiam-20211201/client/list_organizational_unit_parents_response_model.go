// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrganizationalUnitParentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOrganizationalUnitParentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOrganizationalUnitParentsResponse
	GetStatusCode() *int32
	SetBody(v *ListOrganizationalUnitParentsResponseBody) *ListOrganizationalUnitParentsResponse
	GetBody() *ListOrganizationalUnitParentsResponseBody
}

type ListOrganizationalUnitParentsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOrganizationalUnitParentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOrganizationalUnitParentsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationalUnitParentsResponse) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitParentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOrganizationalUnitParentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOrganizationalUnitParentsResponse) GetBody() *ListOrganizationalUnitParentsResponseBody {
	return s.Body
}

func (s *ListOrganizationalUnitParentsResponse) SetHeaders(v map[string]*string) *ListOrganizationalUnitParentsResponse {
	s.Headers = v
	return s
}

func (s *ListOrganizationalUnitParentsResponse) SetStatusCode(v int32) *ListOrganizationalUnitParentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOrganizationalUnitParentsResponse) SetBody(v *ListOrganizationalUnitParentsResponseBody) *ListOrganizationalUnitParentsResponse {
	s.Body = v
	return s
}

func (s *ListOrganizationalUnitParentsResponse) Validate() error {
	return dara.Validate(s)
}
