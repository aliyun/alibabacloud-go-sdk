// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOrganizationalUnitParentIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateOrganizationalUnitParentIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateOrganizationalUnitParentIdResponse
	GetStatusCode() *int32
	SetBody(v *UpdateOrganizationalUnitParentIdResponseBody) *UpdateOrganizationalUnitParentIdResponse
	GetBody() *UpdateOrganizationalUnitParentIdResponseBody
}

type UpdateOrganizationalUnitParentIdResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOrganizationalUnitParentIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOrganizationalUnitParentIdResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateOrganizationalUnitParentIdResponse) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationalUnitParentIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateOrganizationalUnitParentIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateOrganizationalUnitParentIdResponse) GetBody() *UpdateOrganizationalUnitParentIdResponseBody {
	return s.Body
}

func (s *UpdateOrganizationalUnitParentIdResponse) SetHeaders(v map[string]*string) *UpdateOrganizationalUnitParentIdResponse {
	s.Headers = v
	return s
}

func (s *UpdateOrganizationalUnitParentIdResponse) SetStatusCode(v int32) *UpdateOrganizationalUnitParentIdResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOrganizationalUnitParentIdResponse) SetBody(v *UpdateOrganizationalUnitParentIdResponseBody) *UpdateOrganizationalUnitParentIdResponse {
	s.Body = v
	return s
}

func (s *UpdateOrganizationalUnitParentIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
