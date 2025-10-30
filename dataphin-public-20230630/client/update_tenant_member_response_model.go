// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTenantMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTenantMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTenantMemberResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTenantMemberResponseBody) *UpdateTenantMemberResponse
	GetBody() *UpdateTenantMemberResponseBody
}

type UpdateTenantMemberResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTenantMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTenantMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTenantMemberResponse) GoString() string {
	return s.String()
}

func (s *UpdateTenantMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTenantMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTenantMemberResponse) GetBody() *UpdateTenantMemberResponseBody {
	return s.Body
}

func (s *UpdateTenantMemberResponse) SetHeaders(v map[string]*string) *UpdateTenantMemberResponse {
	s.Headers = v
	return s
}

func (s *UpdateTenantMemberResponse) SetStatusCode(v int32) *UpdateTenantMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTenantMemberResponse) SetBody(v *UpdateTenantMemberResponseBody) *UpdateTenantMemberResponse {
	s.Body = v
	return s
}

func (s *UpdateTenantMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
