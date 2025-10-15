// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsForOrganizationalUnitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApplicationsForOrganizationalUnitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApplicationsForOrganizationalUnitResponse
	GetStatusCode() *int32
	SetBody(v *ListApplicationsForOrganizationalUnitResponseBody) *ListApplicationsForOrganizationalUnitResponse
	GetBody() *ListApplicationsForOrganizationalUnitResponseBody
}

type ListApplicationsForOrganizationalUnitResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationsForOrganizationalUnitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationsForOrganizationalUnitResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForOrganizationalUnitResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationsForOrganizationalUnitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApplicationsForOrganizationalUnitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApplicationsForOrganizationalUnitResponse) GetBody() *ListApplicationsForOrganizationalUnitResponseBody {
	return s.Body
}

func (s *ListApplicationsForOrganizationalUnitResponse) SetHeaders(v map[string]*string) *ListApplicationsForOrganizationalUnitResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationsForOrganizationalUnitResponse) SetStatusCode(v int32) *ListApplicationsForOrganizationalUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationsForOrganizationalUnitResponse) SetBody(v *ListApplicationsForOrganizationalUnitResponseBody) *ListApplicationsForOrganizationalUnitResponse {
	s.Body = v
	return s
}

func (s *ListApplicationsForOrganizationalUnitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
