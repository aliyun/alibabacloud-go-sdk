// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserAdOrganizationUnitsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserAdOrganizationUnitsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserAdOrganizationUnitsResponse
	GetStatusCode() *int32
	SetBody(v *ListUserAdOrganizationUnitsResponseBody) *ListUserAdOrganizationUnitsResponse
	GetBody() *ListUserAdOrganizationUnitsResponseBody
}

type ListUserAdOrganizationUnitsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserAdOrganizationUnitsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserAdOrganizationUnitsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserAdOrganizationUnitsResponse) GoString() string {
	return s.String()
}

func (s *ListUserAdOrganizationUnitsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserAdOrganizationUnitsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserAdOrganizationUnitsResponse) GetBody() *ListUserAdOrganizationUnitsResponseBody {
	return s.Body
}

func (s *ListUserAdOrganizationUnitsResponse) SetHeaders(v map[string]*string) *ListUserAdOrganizationUnitsResponse {
	s.Headers = v
	return s
}

func (s *ListUserAdOrganizationUnitsResponse) SetStatusCode(v int32) *ListUserAdOrganizationUnitsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserAdOrganizationUnitsResponse) SetBody(v *ListUserAdOrganizationUnitsResponseBody) *ListUserAdOrganizationUnitsResponse {
	s.Body = v
	return s
}

func (s *ListUserAdOrganizationUnitsResponse) Validate() error {
	return dara.Validate(s)
}
