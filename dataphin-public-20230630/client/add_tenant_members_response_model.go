// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTenantMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddTenantMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddTenantMembersResponse
	GetStatusCode() *int32
	SetBody(v *AddTenantMembersResponseBody) *AddTenantMembersResponse
	GetBody() *AddTenantMembersResponseBody
}

type AddTenantMembersResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddTenantMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddTenantMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s AddTenantMembersResponse) GoString() string {
	return s.String()
}

func (s *AddTenantMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddTenantMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddTenantMembersResponse) GetBody() *AddTenantMembersResponseBody {
	return s.Body
}

func (s *AddTenantMembersResponse) SetHeaders(v map[string]*string) *AddTenantMembersResponse {
	s.Headers = v
	return s
}

func (s *AddTenantMembersResponse) SetStatusCode(v int32) *AddTenantMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *AddTenantMembersResponse) SetBody(v *AddTenantMembersResponseBody) *AddTenantMembersResponse {
	s.Body = v
	return s
}

func (s *AddTenantMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
