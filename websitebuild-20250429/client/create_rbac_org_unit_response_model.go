// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRbacOrgUnitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRbacOrgUnitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRbacOrgUnitResponse
	GetStatusCode() *int32
	SetBody(v *CreateRbacOrgUnitResponseBody) *CreateRbacOrgUnitResponse
	GetBody() *CreateRbacOrgUnitResponseBody
}

type CreateRbacOrgUnitResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRbacOrgUnitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRbacOrgUnitResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRbacOrgUnitResponse) GoString() string {
	return s.String()
}

func (s *CreateRbacOrgUnitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRbacOrgUnitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRbacOrgUnitResponse) GetBody() *CreateRbacOrgUnitResponseBody {
	return s.Body
}

func (s *CreateRbacOrgUnitResponse) SetHeaders(v map[string]*string) *CreateRbacOrgUnitResponse {
	s.Headers = v
	return s
}

func (s *CreateRbacOrgUnitResponse) SetStatusCode(v int32) *CreateRbacOrgUnitResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRbacOrgUnitResponse) SetBody(v *CreateRbacOrgUnitResponseBody) *CreateRbacOrgUnitResponse {
	s.Body = v
	return s
}

func (s *CreateRbacOrgUnitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
