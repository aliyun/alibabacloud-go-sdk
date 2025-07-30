// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOrganizationalUnitDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateOrganizationalUnitDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateOrganizationalUnitDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateOrganizationalUnitDescriptionResponseBody) *UpdateOrganizationalUnitDescriptionResponse
	GetBody() *UpdateOrganizationalUnitDescriptionResponseBody
}

type UpdateOrganizationalUnitDescriptionResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOrganizationalUnitDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOrganizationalUnitDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateOrganizationalUnitDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationalUnitDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateOrganizationalUnitDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateOrganizationalUnitDescriptionResponse) GetBody() *UpdateOrganizationalUnitDescriptionResponseBody {
	return s.Body
}

func (s *UpdateOrganizationalUnitDescriptionResponse) SetHeaders(v map[string]*string) *UpdateOrganizationalUnitDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateOrganizationalUnitDescriptionResponse) SetStatusCode(v int32) *UpdateOrganizationalUnitDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOrganizationalUnitDescriptionResponse) SetBody(v *UpdateOrganizationalUnitDescriptionResponseBody) *UpdateOrganizationalUnitDescriptionResponse {
	s.Body = v
	return s
}

func (s *UpdateOrganizationalUnitDescriptionResponse) Validate() error {
	return dara.Validate(s)
}
