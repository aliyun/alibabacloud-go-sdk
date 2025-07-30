// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrganizationalUnitsForApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOrganizationalUnitsForApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOrganizationalUnitsForApplicationResponse
	GetStatusCode() *int32
	SetBody(v *ListOrganizationalUnitsForApplicationResponseBody) *ListOrganizationalUnitsForApplicationResponse
	GetBody() *ListOrganizationalUnitsForApplicationResponseBody
}

type ListOrganizationalUnitsForApplicationResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOrganizationalUnitsForApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOrganizationalUnitsForApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationalUnitsForApplicationResponse) GoString() string {
	return s.String()
}

func (s *ListOrganizationalUnitsForApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOrganizationalUnitsForApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOrganizationalUnitsForApplicationResponse) GetBody() *ListOrganizationalUnitsForApplicationResponseBody {
	return s.Body
}

func (s *ListOrganizationalUnitsForApplicationResponse) SetHeaders(v map[string]*string) *ListOrganizationalUnitsForApplicationResponse {
	s.Headers = v
	return s
}

func (s *ListOrganizationalUnitsForApplicationResponse) SetStatusCode(v int32) *ListOrganizationalUnitsForApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOrganizationalUnitsForApplicationResponse) SetBody(v *ListOrganizationalUnitsForApplicationResponseBody) *ListOrganizationalUnitsForApplicationResponse {
	s.Body = v
	return s
}

func (s *ListOrganizationalUnitsForApplicationResponse) Validate() error {
	return dara.Validate(s)
}
