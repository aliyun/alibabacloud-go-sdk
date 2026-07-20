// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRbacOrgUnitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRbacOrgUnitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRbacOrgUnitResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRbacOrgUnitResponseBody) *UpdateRbacOrgUnitResponse
	GetBody() *UpdateRbacOrgUnitResponseBody
}

type UpdateRbacOrgUnitResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRbacOrgUnitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRbacOrgUnitResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRbacOrgUnitResponse) GoString() string {
	return s.String()
}

func (s *UpdateRbacOrgUnitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRbacOrgUnitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRbacOrgUnitResponse) GetBody() *UpdateRbacOrgUnitResponseBody {
	return s.Body
}

func (s *UpdateRbacOrgUnitResponse) SetHeaders(v map[string]*string) *UpdateRbacOrgUnitResponse {
	s.Headers = v
	return s
}

func (s *UpdateRbacOrgUnitResponse) SetStatusCode(v int32) *UpdateRbacOrgUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRbacOrgUnitResponse) SetBody(v *UpdateRbacOrgUnitResponseBody) *UpdateRbacOrgUnitResponse {
	s.Body = v
	return s
}

func (s *UpdateRbacOrgUnitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
