// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTenantMembersBySourceUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddTenantMembersBySourceUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddTenantMembersBySourceUserResponse
	GetStatusCode() *int32
	SetBody(v *AddTenantMembersBySourceUserResponseBody) *AddTenantMembersBySourceUserResponse
	GetBody() *AddTenantMembersBySourceUserResponseBody
}

type AddTenantMembersBySourceUserResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddTenantMembersBySourceUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddTenantMembersBySourceUserResponse) String() string {
	return dara.Prettify(s)
}

func (s AddTenantMembersBySourceUserResponse) GoString() string {
	return s.String()
}

func (s *AddTenantMembersBySourceUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddTenantMembersBySourceUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddTenantMembersBySourceUserResponse) GetBody() *AddTenantMembersBySourceUserResponseBody {
	return s.Body
}

func (s *AddTenantMembersBySourceUserResponse) SetHeaders(v map[string]*string) *AddTenantMembersBySourceUserResponse {
	s.Headers = v
	return s
}

func (s *AddTenantMembersBySourceUserResponse) SetStatusCode(v int32) *AddTenantMembersBySourceUserResponse {
	s.StatusCode = &v
	return s
}

func (s *AddTenantMembersBySourceUserResponse) SetBody(v *AddTenantMembersBySourceUserResponseBody) *AddTenantMembersBySourceUserResponse {
	s.Body = v
	return s
}

func (s *AddTenantMembersBySourceUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
