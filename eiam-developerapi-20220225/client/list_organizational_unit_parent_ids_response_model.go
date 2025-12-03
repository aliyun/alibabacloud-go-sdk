// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrganizationalUnitParentIdsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOrganizationalUnitParentIdsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOrganizationalUnitParentIdsResponse
	GetStatusCode() *int32
	SetBody(v *ListOrganizationalUnitParentIdsResponseBody) *ListOrganizationalUnitParentIdsResponse
	GetBody() *ListOrganizationalUnitParentIdsResponseBody
}

type ListOrganizationalUnitParentIdsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOrganizationalUnitParentIdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOrganizationalUnitParentIdsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationalUnitParentIdsResponse) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitParentIdsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOrganizationalUnitParentIdsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOrganizationalUnitParentIdsResponse) GetBody() *ListOrganizationalUnitParentIdsResponseBody {
	return s.Body
}

func (s *ListOrganizationalUnitParentIdsResponse) SetHeaders(v map[string]*string) *ListOrganizationalUnitParentIdsResponse {
	s.Headers = v
	return s
}

func (s *ListOrganizationalUnitParentIdsResponse) SetStatusCode(v int32) *ListOrganizationalUnitParentIdsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOrganizationalUnitParentIdsResponse) SetBody(v *ListOrganizationalUnitParentIdsResponseBody) *ListOrganizationalUnitParentIdsResponse {
	s.Body = v
	return s
}

func (s *ListOrganizationalUnitParentIdsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
